package moysklad

// Credits: https://github.com/hashicorp/go-retryablehttp
import (
	"bytes"
	"context"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var (
	// Default retry configuration
	defaultRetryWaitMin = 1 * time.Second
	defaultRetryWaitMax = 10 * time.Second
	defaultRetryMax     = 5

	// defaultLogger is the logger provided with defaultClient
	defaultLogger = log.New(os.Stderr, "", log.LstdFlags)

	// We need to consume response bodies to maintain http connections, but
	// limit the size we consume to respReadLimit.
	respReadLimit = int64(4096)

	// A regular expression to match the error returned by net/http when the
	// configured number of redirects is exhausted. This error isn't typed
	// specifically so we resort to matching on the error string.
	redirectsErrorRe = regexp.MustCompile(`stopped after \d+ redirects\z`)

	// A regular expression to match the error returned by net/http when the
	// scheme specified in the URL is invalid. This error isn't typed
	// specifically so we resort to matching on the error string.
	schemeErrorRe = regexp.MustCompile(`unsupported protocol scheme`)

	// A regular expression to match the error returned by net/http when the
	// TLS certificate is not trusted. This error isn't typed
	// specifically so we resort to matching on the error string.
	notTrustedErrorRe = regexp.MustCompile(`certificate is not trusted`)
)

func DefaultPooledTransport() *http.Transport {
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		ForceAttemptHTTP2:     true,
		MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
	}
	return transport
}

func DefaultPooledClient() *http.Client {
	return &http.Client{
		Transport: DefaultPooledTransport(),
	}
}

// ReaderFunc is the type of function that can be given natively to NewRequest
type ReaderFunc func() (io.Reader, error)

// ResponseHandlerFunc is a type of function that takes in a Response, and does something with it.
// The ResponseHandlerFunc is called when the HTTP client successfully receives a response and the
// CheckRetry function indicates that a retry of the base request is not necessary.
// If an error is returned from this function, the CheckRetry policy will be used to determine
// whether to retry the whole request (including this handler).
//
// Make sure to check status codes! Even if the request was completed it may have a non-2xx status code.
//
// The response body is not automatically closed. It must be closed either by the ResponseHandlerFunc or
// by the caller out-of-band. Failure to do so will result in a memory leak.
type ResponseHandlerFunc func(*http.Response) error

// LenReader is an interface implemented by many in-memory io.Reader's. Used
// for automatically sending the right Content-Length header when possible.
type LenReader interface {
	Len() int
}

// Request wraps the metadata needed to create HTTP requests.
type Request struct {
	// body is a seekable reader over the request body payload. This is
	// used to rewind the request data in between retries.
	body ReaderFunc

	responseHandler ResponseHandlerFunc

	// Embed an HTTP request directly. This makes a *Request act exactly
	// like an *http.Request so that all meta methods are supported.
	*http.Request
}

func getBodyReaderAndContentLength(rawBody any) (ReaderFunc, int64, error) {
	var bodyReader ReaderFunc
	var contentLength int64

	switch body := rawBody.(type) {
	// If they gave us a function already, great! Use it.
	case ReaderFunc:
		bodyReader = body
		tmp, err := body()
		if err != nil {
			return nil, 0, err
		}
		if lr, ok := tmp.(LenReader); ok {
			contentLength = int64(lr.Len())
		}
		if c, ok := tmp.(io.Closer); ok {
			_ = c.Close()
		}

	case func() (io.Reader, error):
		bodyReader = body
		tmp, err := body()
		if err != nil {
			return nil, 0, err
		}
		if lr, ok := tmp.(LenReader); ok {
			contentLength = int64(lr.Len())
		}
		if c, ok := tmp.(io.Closer); ok {
			_ = c.Close()
		}

	// If a regular byte slice, we can read it over and over via new
	// readers
	case []byte:
		buf := body
		bodyReader = func() (io.Reader, error) {
			return bytes.NewReader(buf), nil
		}
		contentLength = int64(len(buf))

	// If a bytes.Buffer we can read the underlying byte slice over and
	// over
	case *bytes.Buffer:
		buf := body
		bodyReader = func() (io.Reader, error) {
			return bytes.NewReader(buf.Bytes()), nil
		}
		contentLength = int64(buf.Len())

	// We prioritize *bytes.Reader here because we don't really want to
	// deal with it seeking so want it to match here instead of the
	// io.ReadSeeker case.
	case *bytes.Reader:
		buf, err := io.ReadAll(body)
		if err != nil {
			return nil, 0, err
		}
		bodyReader = func() (io.Reader, error) {
			return bytes.NewReader(buf), nil
		}
		contentLength = int64(len(buf))

	// Compat case
	case io.ReadSeeker:
		raw := body
		bodyReader = func() (io.Reader, error) {
			_, err := raw.Seek(0, 0)
			return io.NopCloser(raw), err
		}
		if lr, ok := raw.(LenReader); ok {
			contentLength = int64(lr.Len())
		}

	// Read all in so we can reset
	case io.Reader:
		buf, err := io.ReadAll(body)
		if err != nil {
			return nil, 0, err
		}
		if len(buf) == 0 {
			bodyReader = func() (io.Reader, error) {
				return http.NoBody, nil
			}
			contentLength = 0
		} else {
			bodyReader = func() (io.Reader, error) {
				return bytes.NewReader(buf), nil
			}
			contentLength = int64(len(buf))
		}

	// No body provided, nothing to do
	case nil:

	// Unrecognized type
	default:
		return nil, 0, fmt.Errorf("cannot handle type %T", rawBody)
	}
	return bodyReader, contentLength, nil
}

// FromRequest wraps an http.Request in a retryablehttp.Request
func FromRequest(r *http.Request) (*Request, error) {
	bodyReader, _, err := getBodyReaderAndContentLength(r.Body)
	if err != nil {
		return nil, err
	}
	// Could assert contentLength == r.ContentLength
	return &Request{body: bodyReader, Request: r}, nil
}

// Logger interface allows to use other loggers than
// standard log.Logger.
type Logger interface {
	Printf(string, ...any)
}

// LeveledLogger is an interface that can be implemented by any logger or a
// logger wrapper to provide leveled logging. The methods accept a message
// string and a variadic number of key-value pairs. For log.Printf style
// formatting where message string contains a format specifier, use Logger
// interface.
type LeveledLogger interface {
	Error(msg string, keysAndValues ...any)
	Info(msg string, keysAndValues ...any)
	Debug(msg string, keysAndValues ...any)
	Warn(msg string, keysAndValues ...any)
}

// hookLogger adapts an LeveledLogger to Logger for use by the existing hook functions
// without changing the API.
type hookLogger struct {
	LeveledLogger
}

func (h hookLogger) Printf(s string, args ...any) {
	h.Info(fmt.Sprintf(s, args...))
}

// RequestLogHook allows a function to run before each retry. The HTTP
// request which will be made, and the retry number (0 for the initial
// request) are available to users. The internal logger is exposed to
// consumers.
type RequestLogHook func(Logger, *http.Request, int)

// ResponseLogHook is like RequestLogHook, but allows running a function
// on each HTTP response. This function will be invoked at the end of
// every HTTP request executed, regardless of whether a subsequent retry
// needs to be performed or not. If the response body is read or closed
// from this method, this will affect the response returned from Do().
type ResponseLogHook func(Logger, *http.Response)

// CheckRetry specifies a policy for handling retries. It is called
// following each request with the response and error values returned by
// the http.Client. If CheckRetry returns false, the RetryableClient stops retrying
// and returns the response to the caller. If CheckRetry returns an error,
// that error value is returned in lieu of the error from the request. The
// RetryableClient will close any response body when retrying, but if the retry is
// aborted it is up to the CheckRetry callback to properly close any
// response body before returning.
type CheckRetry func(ctx context.Context, resp *http.Response, err error) (bool, error)

// Backoff specifies a policy for how long to wait between retries.
// It is called after a failing request to determine the amount of time
// that should pass before trying again.
type Backoff func(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration

// ErrorHandler is called if retries are expired, containing the last status
// from the http library. If not specified, default behavior for the library is
// to close the body and return an error indicating how many tries were
// attempted. If overriding this, be sure to close the body if needed.
type ErrorHandler func(resp *http.Response, err error, numTries int) (*http.Response, error)

// RetryableClient is used to make HTTP requests. It adds additional functionality
// like automatic retries to tolerate minor outages.
type RetryableClient struct {
	Logger          any
	HTTPClient      *http.Client
	RequestLogHook  RequestLogHook
	ResponseLogHook ResponseLogHook
	CheckRetry      CheckRetry
	Backoff         Backoff
	ErrorHandler    ErrorHandler
	RetryWaitMin    time.Duration
	RetryWaitMax    time.Duration
	RetryMax        int
	loggerInit      sync.Once
	clientInit      sync.Once
}

// newRetryableClient creates a new RetryableClient with default settings.
func newRetryableClient() *RetryableClient {
	return &RetryableClient{
		HTTPClient:   DefaultPooledClient(),
		Logger:       defaultLogger,
		RetryWaitMin: defaultRetryWaitMin,
		RetryWaitMax: defaultRetryWaitMax,
		RetryMax:     defaultRetryMax,
		CheckRetry:   ErrorPropagatedRetryPolicy,
		Backoff:      DefaultBackoff,
	}
}

func (c *RetryableClient) logger() any {
	c.loggerInit.Do(func() {
		if c.Logger == nil {
			return
		}

		switch c.Logger.(type) {
		case Logger, LeveledLogger:
			// ok
		default:
			// This should happen in dev when they are setting Logger and work on code, not in prod.
			panic(fmt.Sprintf("invalid logger type passed, must be Logger or LeveledLogger, was %T", c.Logger))
		}
	})

	return c.Logger
}

// ErrorPropagatedRetryPolicy is the same as DefaultRetryPolicy, except it
// propagates errors back instead of returning nil. This allows you to inspect
// why it decided to retry or not.
func ErrorPropagatedRetryPolicy(ctx context.Context, resp *http.Response, err error) (bool, error) {
	// do not retry on context.Canceled or context.DeadlineExceeded
	if ctx.Err() != nil {
		return false, ctx.Err()
	}

	return baseRetryPolicy(resp, err)
}

func baseRetryPolicy(resp *http.Response, err error) (bool, error) {
	if err != nil {
		if v, ok := err.(*url.Error); ok {
			// Don't retry if the error was due to too many redirects.
			if redirectsErrorRe.MatchString(v.Error()) {
				return false, v
			}

			// Don't retry if the error was due to an invalid protocol scheme.
			if schemeErrorRe.MatchString(v.Error()) {
				return false, v
			}

			// Don't retry if the error was due to TLS cert verification failure.
			if notTrustedErrorRe.MatchString(v.Error()) {
				return false, v
			}
			if _, ok := v.Err.(x509.UnknownAuthorityError); ok {
				return false, v
			}
		}

		// The error is likely recoverable so retry.
		return true, nil
	}

	// 429 Too Many Requests is recoverable. Sometimes the server puts
	// a Retry-After response header to indicate when the server is
	// available to start processing request from client.
	if resp.StatusCode == http.StatusTooManyRequests {
		return true, nil
	}

	// Check the response code. We retry on 500-range responses to allow
	// the server time to recover, as 500's are typically not permanent
	// errors and may relate to outages on the server side. This will catch
	// invalid response codes as well, like 0 and 999.
	if resp.StatusCode == 0 || (resp.StatusCode >= 500 && resp.StatusCode != http.StatusNotImplemented) {
		return true, fmt.Errorf("unexpected HTTP status %s", resp.Status)
	}

	return false, nil
}

// DefaultBackoff provides a default callback for RetryableClient.Backoff which
// will perform exponential backoff based on the attempt number and limited
// by the provided minimum and maximum durations.
//
// It also tries to parse Retry-After response header when a http.StatusTooManyRequests
// (HTTP Code 429) is found in the resp parameter. Hence it will return the number of
// seconds the server states it may be ready to process more requests from this client.
func DefaultBackoff(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
	if resp != nil {
		if resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode == http.StatusServiceUnavailable {
			if s, ok := resp.Header[headerRetryAfter]; ok {
				if sleep, err := strconv.ParseInt(s[0], 10, 64); err == nil {
					return time.Second * time.Duration(sleep)
				}
			}
		}
	}

	mult := math.Pow(2, float64(attemptNum)) * float64(min)
	sleep := time.Duration(mult)
	if float64(sleep) != mult || sleep > max {
		sleep = max
	}
	return sleep
}

// Do wraps calling an HTTP method with retries.
func (c *RetryableClient) Do(req *Request) (*http.Response, error) {
	c.clientInit.Do(func() {
		if c.HTTPClient == nil {
			c.HTTPClient = DefaultPooledClient()
		}
	})

	logger := c.logger()

	if logger != nil {
		switch v := logger.(type) {
		case LeveledLogger:
			v.Debug("performing request", "method", req.Method, "url", req.URL)
		case Logger:
			v.Printf("[DEBUG] %s %s", req.Method, req.URL)
		}
	}

	var resp *http.Response
	var attempt int
	var shouldRetry bool
	var doErr, respErr, checkErr error

	for i := 0; ; i++ {
		doErr, respErr = nil, nil
		attempt++

		// Always rewind the request body when non-nil.
		if req.body != nil {
			body, err := req.body()
			if err != nil {
				c.HTTPClient.CloseIdleConnections()
				return resp, err
			}
			if c, ok := body.(io.ReadCloser); ok {
				req.Body = c
			} else {
				req.Body = io.NopCloser(body)
			}
		}

		if c.RequestLogHook != nil {
			switch v := logger.(type) {
			case LeveledLogger:
				c.RequestLogHook(hookLogger{v}, req.Request, i)
			case Logger:
				c.RequestLogHook(v, req.Request, i)
			default:
				c.RequestLogHook(nil, req.Request, i)
			}
		}

		// Attempt the request
		resp, doErr = c.HTTPClient.Do(req.Request)

		// Check if we should continue with retries.
		shouldRetry, checkErr = c.CheckRetry(req.Context(), resp, doErr)
		if !shouldRetry && doErr == nil && req.responseHandler != nil {
			respErr = req.responseHandler(resp)
			shouldRetry, checkErr = c.CheckRetry(req.Context(), resp, respErr)
		}

		err := doErr
		if respErr != nil {
			err = respErr
		}
		if err != nil {
			switch v := logger.(type) {
			case LeveledLogger:
				v.Error("request failed", "error", err, "method", req.Method, "url", req.URL)
			case Logger:
				v.Printf("[ERR] %s %s request failed: %v", req.Method, req.URL, err)
			}
		} else {
			// Call this here to maintain the behavior of logging all requests,
			// even if CheckRetry signals to stop.
			if c.ResponseLogHook != nil {
				// Call the response logger function if provided.
				switch v := logger.(type) {
				case LeveledLogger:
					c.ResponseLogHook(hookLogger{v}, resp)
				case Logger:
					c.ResponseLogHook(v, resp)
				default:
					c.ResponseLogHook(nil, resp)
				}
			}
		}

		if !shouldRetry {
			break
		}

		// We do this before drainBody because there's no need for the I/O if
		// we're breaking out
		remain := c.RetryMax - i
		if remain <= 0 {
			break
		}

		// We're going to retry, consume any response to reuse the connection.
		if doErr == nil {
			c.drainBody(resp.Body)
		}

		wait := c.Backoff(c.RetryWaitMin, c.RetryWaitMax, i, resp)
		if logger != nil {
			desc := fmt.Sprintf("%s %s", req.Method, req.URL)
			if resp != nil {
				desc = fmt.Sprintf("%s (status: %d)", desc, resp.StatusCode)
			}
			switch v := logger.(type) {
			case LeveledLogger:
				v.Debug("retrying request", "request", desc, "timeout", wait, "remaining", remain)
			case Logger:
				v.Printf("[DEBUG] %s: retrying in %s (%d left)", desc, wait, remain)
			}
		}
		timer := time.NewTimer(wait)
		select {
		case <-req.Context().Done():
			timer.Stop()
			c.HTTPClient.CloseIdleConnections()
			return nil, req.Context().Err()
		case <-timer.C:
		}

		// Make shallow copy of http Request so that we can modify its body
		// without racing against the closeBody call in persistConn.writeLoop.
		httpreq := *req.Request
		req.Request = &httpreq
	}

	// this is the closest we have to success criteria
	if doErr == nil && respErr == nil && checkErr == nil && !shouldRetry {
		return resp, nil
	}

	defer c.HTTPClient.CloseIdleConnections()

	var err error
	if checkErr != nil {
		err = checkErr
	} else if respErr != nil {
		err = respErr
	} else {
		err = doErr
	}

	if c.ErrorHandler != nil {
		return c.ErrorHandler(resp, err, attempt)
	}

	// By default, we close the response body and return an error without
	// returning the response
	if resp != nil {
		c.drainBody(resp.Body)
	}

	// this means CheckRetry thought the request was a failure, but didn't
	// communicate why
	if err == nil {
		return nil, fmt.Errorf("%s %s giving up after %d attempt(s)",
			req.Method, req.URL, attempt)
	}

	return nil, fmt.Errorf("%s %s giving up after %d attempt(s): %w",
		req.Method, req.URL, attempt, err)
}

// Try to read the response body so we can reuse this connection.
func (c *RetryableClient) drainBody(body io.ReadCloser) {
	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(body)
	_, err := io.Copy(io.Discard, io.LimitReader(body, respReadLimit))
	if err != nil {
		if c.logger() != nil {
			switch v := c.logger().(type) {
			case LeveledLogger:
				v.Error("error reading response body", "error", err)
			case Logger:
				v.Printf("[ERR] error reading response body: %v", err)
			}
		}
	}
}
