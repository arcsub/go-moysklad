package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
	"net/http"
)

type Endpoint struct {
	client *Client
	uri    string
}

func NewEndpoint(client *Client, uri string) Endpoint {
	return Endpoint{client, uri}
}

type RequestBuilder[T any] struct {
	req *resty.Request
	uri string
}

func NewRequestBuilder[T any](client *Client, uri string) *RequestBuilder[T] {
	return &RequestBuilder[T]{client.client.R(), uri}
}

func (rb *RequestBuilder[T]) SetParams(params *Params) *RequestBuilder[T] {
	v, _ := query.Values(params)
	rb.req.SetQueryParamsFromValues(v)
	return rb
}

func (rb *RequestBuilder[T]) Get(ctx context.Context) (*T, *resty.Response, error) {
	data := new(T)
	apiErrors := new(ApiErrors)
	resp, err := rb.req.SetContext(ctx).SetResult(data).SetError(apiErrors).Get(rb.uri)

	if resp.Error() != nil {
		return data, resp, apiErrors
	}

	return data, resp, err
}

func (rb *RequestBuilder[T]) Put(ctx context.Context, body any) (*T, *resty.Response, error) {
	data := new(T)
	apiErrors := new(ApiErrors)
	resp, err := rb.req.SetContext(ctx).SetBody(body).SetResult(data).SetError(apiErrors).Put(rb.uri)

	if resp.Error() != nil {
		return data, resp, apiErrors
	}

	return data, resp, err
}

func (rb *RequestBuilder[T]) Post(ctx context.Context, body any) (*T, *resty.Response, error) {
	data := new(T)
	apiErrors := new(ApiErrors)
	resp, err := rb.req.SetContext(ctx).SetBody(body).SetResult(data).SetError(apiErrors).Post(rb.uri)

	if resp.Error() != nil {
		return data, resp, apiErrors
	}

	return data, resp, err
}

func (rb *RequestBuilder[T]) Delete(ctx context.Context) (bool, *resty.Response, error) {
	apiErrors := new(ApiErrors)
	resp, err := rb.req.SetContext(ctx).SetError(apiErrors).Delete(rb.uri)

	ok := resp.StatusCode() == http.StatusOK

	if resp.Error() != nil {
		return ok, resp, apiErrors
	}

	return ok, resp, err
}

func (rb *RequestBuilder[T]) Async(ctx context.Context) (*AsyncResultService[T], *resty.Response, error) {
	// устанавливаем флаг async=true на создание асинхронной операции
	rb.req.SetContext(ctx).SetQueryParam("async", "true")

	resp, err := rb.req.Get(rb.uri)
	if err != nil {
		return nil, resp, err
	}
	async := NewAsyncResultService[T](rb.req, resp)
	return async, resp, nil
}
