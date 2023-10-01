package moysklad

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type RequestBuilder[T any] struct {
	client        *Client
	ctx           context.Context
	params        *Params
	url           string
	uri           string
	body          any
	contentHeader bool
}

func NewRequestBuilder[T any](e Endpoint, ctx context.Context) *RequestBuilder[T] {
	return &RequestBuilder[T]{client: e.client, ctx: ctx, uri: e.uri}
}

func (rb *RequestBuilder[T]) WithContext(ctx context.Context) *RequestBuilder[T] {
	rb.ctx = ctx
	return rb
}

func (rb *RequestBuilder[T]) WithPath(path string) *RequestBuilder[T] {
	if strings.HasPrefix(path, "/") {
		rb.uri += path
	} else {
		rb.uri += "/" + path
	}
	return rb
}

func (rb *RequestBuilder[T]) WithParams(params *Params) *RequestBuilder[T] {
	rb.params = params
	return rb
}

func (rb *RequestBuilder[T]) WithBody(body any) *RequestBuilder[T] {
	rb.body = body
	return rb
}

func (rb *RequestBuilder[T]) WithURL(url *url.URL) *RequestBuilder[T] {
	rb.url = url.String()
	return rb
}

func (rb *RequestBuilder[T]) setContentHeader() *RequestBuilder[T] {
	rb.contentHeader = true
	return rb
}

func (rb *RequestBuilder[T]) buildUrlString() error {
	u, err := url.Parse(baseApiURL + rb.uri)
	if err != nil {
		return err
	}
	u.RawQuery = rb.params.QueryString()
	rb.url = u.String()
	return nil
}

func (rb *RequestBuilder[T]) setHeaders(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", defaultUserAgent)
	req.Header.Set("Accept-Encoding", "gzip")

	if rb.client.disableWebhookContent {
		req.Header.Set(headerWebHookDisable, "true")
	}

	if rb.contentHeader {
		req.Header.Set(headerGetContent, "true")
	}
}

type Endpoint struct {
	client *Client
	uri    string
}

func NewEndpoint(client *Client, uri string) Endpoint {
	return Endpoint{client: client, uri: uri}
}

// Создаёт запрос к API МойСклад.
func (rb *RequestBuilder[T]) newRequest(method string) (*http.Request, error) {
	buf := new(bytes.Buffer)
	if rb.body != nil {
		encoder := json.NewEncoder(buf)
		encoder.SetEscapeHTML(false)

		if err := encoder.Encode(rb.body); err != nil {
			return nil, err
		}
	}

	if err := rb.buildUrlString(); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, rb.url, buf)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// Отправляет запрос с необходимыми данными
func (rb *RequestBuilder[T]) do(method string) (*Response, error) {
	req, err := rb.newRequest(method)
	if err != nil {
		return nil, err
	}
	rb.setHeaders(req)
	return rb.client.Do(rb.ctx, req)
}

// Отправляет запрос с необходимыми данными и парсит результат.
func requestWithUnmarshal[T any](rb *RequestBuilder[T], method string) (*T, *Response, error) {
	resp, err := rb.do(method)
	if err != nil {
		return nil, nil, err
	}

	o, err := unmarshalResponse[T](resp)
	if err != nil {
		return nil, resp, err
	}
	return o, resp, nil
}

// Get Отправляет GET запрос и парсит ответ
func (rb *RequestBuilder[T]) Get() (*T, *Response, error) {
	return requestWithUnmarshal[T](rb, http.MethodGet)
}

// Post Отправляет POST запрос и парсит ответ
func (rb *RequestBuilder[T]) Post() (*T, *Response, error) {
	return requestWithUnmarshal[T](rb, http.MethodPost)
}

// Put Отправляет PUT запрос и парсит ответ
func (rb *RequestBuilder[T]) Put() (*T, *Response, error) {
	return requestWithUnmarshal[T](rb, http.MethodPut)
}

// Delete Отправляет DELETE запрос
// Возвращает true в случае успешного удаления
func (rb *RequestBuilder[T]) Delete() (bool, *Response, error) {
	resp, err := rb.do(http.MethodDelete)
	if err != nil {
		return false, resp, err
	}
	ok := resp.StatusCode == http.StatusOK
	return ok, resp, nil
}

// Async Отправляет запрос на создание асинхронной операции.
// Возвращает сервис *AsyncResultService, который можно использовать
// для получения результата выполнения асинхронного запроса.
func (rb *RequestBuilder[T]) Async() (*AsyncResultService[T], *Response, error) {
	// устанавливаем флаг запроса на создание асинхронной операции
	// async=true
	rb.params.withAsync()

	resp, err := rb.do(http.MethodGet)
	if err != nil {
		return nil, resp, err
	}
	async := NewAsyncResultService[T](rb, resp)
	return async, resp, nil
}

func unmarshalResponse[T any](resp *Response) (*T, error) {
	v := new(T)
	decErr := json.NewDecoder(resp.Body).Decode(v)
	// ignore EOF errors caused by empty response body
	if decErr != nil && !errors.Is(decErr, io.EOF) {
		return nil, decErr
	}
	return v, nil
}
