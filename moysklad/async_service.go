package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
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
func (s *AsyncService) GetStatuses(ctx context.Context, params *Params) (*List[Async], *resty.Response, error) {
	return NewRequestBuilder[List[Async]](s.client, s.uri).SetParams(params).Get(ctx)
}

// GetStatusById Получение статуса Асинхронной задачи.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-asinhronnyj-obmen-poluchenie-statusa-asinhronnoj-zadachi
func (s *AsyncService) GetStatusById(ctx context.Context, id *uuid.UUID, params *Params) (*Async, *resty.Response, error) {
	path := fmt.Sprintf("async/%s", id)
	return NewRequestBuilder[Async](s.client, path).SetParams(params).Get(ctx)
}

type AsyncResultService[T any] struct {
	req       *resty.Request
	statusURL *url.URL // URL статуса Асинхронной задачи.
	resultURL *url.URL // URL результата выполнения Асинхронной задачи.
}

// NewAsyncResultService Сервис для работы с асинхронной задачей.
func NewAsyncResultService[T any](req *resty.Request, resp *resty.Response) *AsyncResultService[T] {
	statusUrlStr := resp.Header().Get("Location")
	resultUrlStr := resp.Header().Get("Content-Location")
	statusUrl, _ := url.Parse(statusUrlStr)
	resultUrl, _ := url.Parse(resultUrlStr)

	return &AsyncResultService[T]{
		req:       req,
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
func (s *AsyncResultService[T]) Check(ctx context.Context) (bool, *resty.Response, error) {
	async := &Async{}
	resp, err := s.req.SetContext(ctx).SetBody(async).Get(s.StatusURL().String())
	if err != nil {
		return false, resp, err
	}
	ok := async.State == AsyncStateDone
	return ok, resp, nil
}

// Result Запрос на получение результата.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-asinhronnyj-obmen-poluchenie-rezul-tata-wypolneniq-asinhronnoj-zadachi
func (s *AsyncResultService[T]) Result(ctx context.Context) (*T, *resty.Response, error) {
	data := new(T)
	resp, err := s.req.SetContext(ctx).SetBody(data).Get(s.ResultURL().String())
	return data, resp, err
}

// Cancel Отмена Асинхронной задачи.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-asinhronnyj-obmen-otmena-asinhronnoj-zadachi
func (s *AsyncResultService[T]) Cancel(ctx context.Context) (bool, *resty.Response, error) {
	u, _ := s.StatusURL().Parse("/cancel")
	resp, err := s.req.SetContext(ctx).Post(u.String())
	if err != nil {
		return false, resp, err
	}

	ok := resp.StatusCode() == http.StatusNoContent
	return ok, resp, nil
}
