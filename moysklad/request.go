package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/goccy/go-json"
	"github.com/google/go-querystring/query"
	"log"
	"net/http"
	"reflect"
	"strings"
)

type Endpoint struct {
	client *Client
	uri    string
}

func NewEndpoint(client *Client, uri string) Endpoint {
	return Endpoint{client, uri}
}

type RequestBuilder[T any] struct {
	client *Client
	req    *resty.Request
	uri    string
}

func NewRequestBuilder[T any](client *Client, uri string) *RequestBuilder[T] {
	return &RequestBuilder[T]{client, client.R(), uri}
}

// TODO: improve
func parseResponse[T any](r *resty.Response) (*T, *resty.Response, error) {
	// check empty response body
	if r.Body() == nil {
		return nil, r, nil
	}

	var (
		bodyBytes  = r.Body()
		bodyString = string(bodyBytes)
		result     = *new(T)
		apiErrors  ApiErrors
	)

	// response body is slice
	if strings.HasPrefix(bodyString, "[") && strings.HasSuffix(bodyString, "]") &&
		reflect.ValueOf(result).Kind() == reflect.Slice {

		switch statusCode := r.StatusCode(); {

		case statusCode >= http.StatusOK && statusCode < http.StatusMultipleChoices: // ok
			if err := json.Unmarshal(bodyBytes, &result); err != nil {
				return nil, r, err
			}

		case statusCode >= http.StatusBadRequest: // error
			var rawSlice []any
			if err := json.Unmarshal(bodyBytes, &rawSlice); err != nil {
				return nil, r, err
			}

			resultType := reflect.TypeOf(result)

			if resultType.Kind() == reflect.Ptr {
				resultType = resultType.Elem()
			}

			if resultType.Kind() == reflect.Slice {
				resultType = resultType.Elem()
			}

			if resultType.Kind() != reflect.Struct {
				log.Printf("[DEBUG] resultType is not a struct: %v", resultType.Kind())
				return nil, r, nil
			}

			data := reflect.New(reflect.TypeOf(result)).Interface()
			dataType := reflect.TypeOf(data)

			if dataType.Kind() == reflect.Ptr {
				dataType = dataType.Elem()
			}

			if dataType.Kind() == reflect.Slice {
				dataType = dataType.Elem()
			} else {
				log.Printf("[DEBUG] dataType is %v", dataType.Kind())
				return nil, r, nil
			}

			dataValue := reflect.ValueOf(data)
			if dataValue.Kind() == reflect.Ptr {
				dataValue = dataValue.Elem()
			}

			if dataType.Kind() == reflect.Ptr {
				dataValue = dataValue.Elem()
			}

			elem := dataValue.Type().Elem()

			if elem.Kind() == reflect.Ptr {
				elem = elem.Elem()
			}

			if elem.Kind() == reflect.Slice {
				elem = elem.Elem()
			}

			for _, object := range rawSlice {
				o, err := json.Marshal(object)
				if err != nil {
					return nil, r, err
				}
				newElem := reflect.New(elem).Elem()
				if err := json.Unmarshal(o, newElem.Addr().Interface()); err == nil && !newElem.IsZero() {
					if newElem.Kind() == reflect.Ptr {
						newElem = newElem.Elem()
					}

					dataValue.Set(reflect.Append(dataValue, newElem))
					continue
				}

				var errs ApiErrors
				if err := json.Unmarshal(o, &errs); err == nil {
					apiErrors.ApiErrors = append(apiErrors.ApiErrors, errs.ApiErrors...)
				}
			}

			result = dataValue.Interface().(T)
		}
	} else { // response body is object
		switch statusCode := r.StatusCode(); {

		case statusCode >= http.StatusOK && statusCode < http.StatusMultipleChoices: // ok
			if err := json.Unmarshal(bodyBytes, &result); err == nil &&
				!reflect.ValueOf(&result).Elem().IsZero() {
				return &result, r, nil
			}

		case statusCode >= http.StatusBadRequest: // error
			if err := json.Unmarshal(bodyBytes, &apiErrors); err != nil {
				return nil, r, err
			}
		}
	}

	if len(apiErrors.ApiErrors) > 0 {
		return &result, r, apiErrors
	}

	return &result, r, nil
}

func (requestBuilder *RequestBuilder[T]) SetHeader(header, value string) *RequestBuilder[T] {
	requestBuilder.req.Header.Set(header, value)
	return requestBuilder
}

func (requestBuilder *RequestBuilder[T]) SetParams(params ...*Params) *RequestBuilder[T] {
	if len(params) > 0 {
		v, _ := query.Values(params[0])
		requestBuilder.req.SetQueryParamsFromValues(v)
	}
	return requestBuilder
}

func (requestBuilder *RequestBuilder[T]) Send(ctx context.Context, method string, body any) (*T, *resty.Response, error) {
	// Ограничения на количество запросов
	requestBuilder.client.limits.Wait()
	defer requestBuilder.client.limits.Done()

	resp, err := requestBuilder.req.SetContext(ctx).SetBody(body).Execute(method, requestBuilder.uri)
	if err != nil {
		return nil, resp, err
	}

	return parseResponse[T](resp)
}

func (requestBuilder *RequestBuilder[T]) Get(ctx context.Context) (*T, *resty.Response, error) {
	return requestBuilder.Send(ctx, http.MethodGet, nil)
}

func (requestBuilder *RequestBuilder[T]) Put(ctx context.Context, body any) (*T, *resty.Response, error) {
	return requestBuilder.Send(ctx, http.MethodPut, body)
}

func (requestBuilder *RequestBuilder[T]) Post(ctx context.Context, body any) (*T, *resty.Response, error) {
	return requestBuilder.Send(ctx, http.MethodPost, body)
}

func (requestBuilder *RequestBuilder[T]) Delete(ctx context.Context) (bool, *resty.Response, error) {
	// Ограничения на количество запросов
	requestBuilder.client.limits.Wait()
	defer requestBuilder.client.limits.Done()

	_, resp, err := requestBuilder.Send(ctx, http.MethodDelete, nil)
	return resp.StatusCode() == http.StatusOK, resp, err
}

func (requestBuilder *RequestBuilder[T]) Async(ctx context.Context) (AsyncResultService[T], *resty.Response, error) {
	// Ограничения на количество запросов
	requestBuilder.client.limits.Wait()
	defer requestBuilder.client.limits.Done()

	// устанавливаем флаг async=true на создание асинхронной операции
	requestBuilder.req.SetContext(ctx).SetQueryParam("async", "true")

	resp, err := requestBuilder.req.Get(requestBuilder.uri)
	if err != nil {
		return nil, resp, err
	}
	async := NewAsyncResultService[T](requestBuilder.client, resp)
	return async, resp, nil
}

// FetchMeta позволяет выполнить точечный запрос по переданному объекту [Meta].
//
// Необходимо точно указать обобщённый тип T, который ожидаем получить в ответ, иначе есть риск получить ошибку.
func FetchMeta[T any](ctx context.Context, client *Client, meta Meta, params ...*Params) (*T, *resty.Response, error) {
	return NewRequestBuilder[T](client, strings.ReplaceAll(meta.GetHref(), baseApiURL, "")).SetParams(params...).Get(ctx)
}

// Context объект, содержащий метаданные о выполнившем запрос сотруднике.
type Context struct {
	Employee MetaWrapper `json:"employee,omitempty"`
}

// String реализует интерфейс [fmt.Stringer].
func (context Context) String() string {
	return Stringify(context)
}

// List объект ответа на запрос списка сущностей T.
type List[T any] struct {
	Context Context `json:"context"` // Метаданные о выполнившем запрос сотруднике
	MetaArray[T]
}

// String реализует интерфейс [fmt.Stringer].
func (list List[T]) String() string {
	return Stringify(list)
}
