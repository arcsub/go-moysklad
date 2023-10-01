package moysklad

import (
	"context"
	"github.com/google/uuid"
	"net/http"
	"net/url"
)

// AsyncService Сервис для работы с асинхронными задачами.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-asinhronnyj-obmen
type AsyncService struct {
	Endpoint
}

func NewAsyncService(client *Client) *AsyncService {
	e := NewEndpoint(client, "async")
	return &AsyncService{e}
}

// GetStatuses Статусы Асинхронных задач.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-asinhronnyj-obmen-statusy-asinhronnyh-zadach
func (s *AsyncService) GetStatuses(ctx context.Context, params *Params) (*List[Async], *Response, error) {
	return NewRequestBuilder[List[Async]](s.Endpoint, ctx).WithParams(params).Get()
}

// GetStatusById Получение статуса Асинхронной задачи.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-asinhronnyj-obmen-poluchenie-statusa-asinhronnoj-zadachi
func (s *AsyncService) GetStatusById(ctx context.Context, id *uuid.UUID, params *Params) (*Async, *Response, error) {
	path := id.String()
	return NewRequestBuilder[Async](s.Endpoint, ctx).WithPath(path).WithParams(params).Get()
}

type AsyncResultService[T any] struct {
	rb        *RequestBuilder[T]
	statusURL *url.URL // URL статуса Асинхронной задачи.
	resultURL *url.URL // URL результата выполнения Асинхронной задачи.
}

// NewAsyncResultService Сервис для работы с асинхронной задачей.
func NewAsyncResultService[T any](rb *RequestBuilder[T], resp *Response) *AsyncResultService[T] {
	statusUrlStr := resp.Header.Get("Location")
	resultUrlStr := resp.Header.Get("Content-Location")
	statusUrl, _ := url.Parse(statusUrlStr)
	resultUrl, _ := url.Parse(resultUrlStr)

	return &AsyncResultService[T]{
		rb:        rb,
		statusURL: statusUrl,
		resultURL: resultUrl,
	}
}

// StatusURL возвращает URL проверки статуса асинхронной задачи.
func (s *AsyncResultService[T]) StatusURL() *url.URL {
	return s.statusURL
}

// ResultURL возвращает URL результата выполнения асинхронной задачи.
func (s *AsyncResultService[T]) ResultURL() *url.URL {
	return s.resultURL
}

// Check Проверяет статус асинхронной задачи.
// Если статус задачи = DONE, возвращает true, иначе false
func (s *AsyncResultService[T]) Check(ctx context.Context) (bool, *Response, error) {
	rb := NewRequestBuilder[Async](Endpoint{client: s.rb.client}, ctx).WithURL(s.StatusURL())
	res, response, err := rb.Get()
	if err != nil {
		return false, response, err
	}
	ok := res.State == AsyncStateDone
	return ok, response, nil
}

// Result Запрос на получение результата.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-asinhronnyj-obmen-poluchenie-rezul-tata-wypolneniq-asinhronnoj-zadachi
func (s *AsyncResultService[T]) Result(ctx context.Context) (*T, *Response, error) {
	rb := NewRequestBuilder[T](Endpoint{client: s.rb.client}, ctx).WithURL(s.ResultURL())
	return rb.Get()
}

// Cancel Отмена Асинхронной задачи.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-asinhronnyj-obmen-otmena-asinhronnoj-zadachi
func (s *AsyncResultService[T]) Cancel(ctx context.Context) (bool, *Response, error) {
	u, _ := s.StatusURL().Parse("/cancel")
	rb := NewRequestBuilder[any](Endpoint{client: s.rb.client}, ctx).WithURL(u)
	resp, err := rb.do(http.MethodPut)
	if err != nil {
		return false, resp, err
	}
	ok := resp.StatusCode == http.StatusNoContent
	return ok, resp, nil
}
