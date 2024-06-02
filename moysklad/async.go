package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"net/http"
	"net/url"
)

// Async Асинхронная задача.
// Ключевое слово: async
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-asinhronnyj-obmen-asinhronnaq-zadacha
type Async struct {
	DeletionDate Timestamp  `json:"deletionDate,omitempty"`
	Meta         Meta       `json:"meta,omitempty"`
	Owner        Meta       `json:"owner,omitempty"`
	RequestURL   string     `json:"request,omitempty"`
	ResultURL    string     `json:"resultUrl,omitempty"`
	State        AsyncState `json:"state,omitempty"`
	ApiErrors    ApiErrors  `json:"errors,omitempty"`
	AccountID    uuid.UUID  `json:"accountId,omitempty"`
	ID           uuid.UUID  `json:"id,omitempty"`
}

func (async Async) String() string {
	return Stringify(async)
}

// AsyncState Статус выполнения Асинхронной задачи.
type AsyncState string

const (
	AsyncStatePending    AsyncState = "PENDING"    // Задача находится в очереди
	AsyncStateProcessing AsyncState = "PROCESSING" // Задача находится в обработке, результат еще не готов
	AsyncStateDone       AsyncState = "DONE"       // Задача выполнена успешно
	AsyncStateError      AsyncState = "ERROR"      // Задача не была выполнена в результате внутренней ошибки. В этом случае нужно попробовать запустить задачу заново
	AsyncStateCancel     AsyncState = "CANCEL"     // Задача была отменена
	AsyncStateApiError   AsyncState = "API_ERROR"  // Задача была завершена с ошибкой апи
)

// AsyncService Сервис для работы с асинхронными задачами.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-asinhronnyj-obmen
type AsyncService interface {
	// GetStatuses выполняет запрос на получение Асинхронных задач.
	//
	// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-asinhronnyj-obmen-statusy-asinhronnyh-zadach
	GetStatuses(ctx context.Context, params *Params) (*List[Async], *resty.Response, error)

	// GetStatusByID выполняет запрос на получение статуса Асинхронной задачи.
	//
	// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-asinhronnyj-obmen-poluchenie-statusa-asinhronnoj-zadachi
	GetStatusByID(ctx context.Context, id *uuid.UUID, params *Params) (*Async, *resty.Response, error)
}

type asyncService struct {
	Endpoint
}

func NewAsyncService(client *Client) AsyncService {
	e := NewEndpoint(client, "async")
	return &asyncService{e}
}

func (s *asyncService) GetStatuses(ctx context.Context, params *Params) (*List[Async], *resty.Response, error) {
	return NewRequestBuilder[List[Async]](s.client, s.uri).SetParams(params).Get(ctx)
}

func (s *asyncService) GetStatusByID(ctx context.Context, id *uuid.UUID, params *Params) (*Async, *resty.Response, error) {
	path := fmt.Sprintf("async/%s", id)
	return NewRequestBuilder[Async](s.client, path).SetParams(params).Get(ctx)
}

type AsyncResultService[T any] interface {
	// StatusURL возвращает URL проверки статуса асинхронной задачи.
	StatusURL() *url.URL

	// ResultURL возвращает URL результата выполнения асинхронной задачи.
	ResultURL() *url.URL

	// Check выполняет запрос на проверку статус асинхронной задачи.
	//
	// Если статус задачи = AsyncStateDone (DONE), возвращает true, иначе false
	Check(ctx context.Context) (bool, *resty.Response, error)

	// Result выполняет запрос на получение результата.
	//
	// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-asinhronnyj-obmen-poluchenie-rezul-tata-wypolneniq-asinhronnoj-zadachi
	Result(ctx context.Context) (*T, *resty.Response, error)

	// Cancel выполняет запрос на отмену Асинхронной задачи.
	//
	// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-asinhronnyj-obmen-otmena-asinhronnoj-zadachi
	Cancel(ctx context.Context) (bool, *resty.Response, error)
}

type asyncResultService[T any] struct {
	client    *Client
	statusURL *url.URL // URL статуса Асинхронной задачи.
	resultURL *url.URL // URL результата выполнения Асинхронной задачи.
}

// NewAsyncResultService Сервис для работы с асинхронной задачей.
func NewAsyncResultService[T any](client *Client, resp *resty.Response) AsyncResultService[T] {
	statusUrlStr := resp.Header().Get("Location")
	resultUrlStr := resp.Header().Get("Content-Location")
	statusUrl, _ := url.Parse(statusUrlStr)
	resultUrl, _ := url.Parse(resultUrlStr)

	return &asyncResultService[T]{
		client:    client,
		statusURL: statusUrl,
		resultURL: resultUrl,
	}
}

func (s *asyncResultService[T]) StatusURL() *url.URL {
	return s.statusURL
}

func (s *asyncResultService[T]) ResultURL() *url.URL {
	return s.resultURL
}

func (s *asyncResultService[T]) Check(ctx context.Context) (bool, *resty.Response, error) {
	async, resp, err := NewRequestBuilder[Async](s.client, s.StatusURL().String()).Get(ctx)
	if err != nil {
		return false, resp, err
	}
	return async.State == AsyncStateDone, resp, nil
}

func (s *asyncResultService[T]) Result(ctx context.Context) (*T, *resty.Response, error) {
	data, resp, err := NewRequestBuilder[T](s.client, s.ResultURL().String()).Get(ctx)
	if err != nil {
		return nil, resp, err
	}
	return data, resp, nil
}

func (s *asyncResultService[T]) Cancel(ctx context.Context) (bool, *resty.Response, error) {
	u, _ := s.StatusURL().Parse("/cancel")
	_, resp, err := NewRequestBuilder[any](s.client, u.String()).Post(ctx, nil)
	if err != nil {
		return false, resp, err
	}
	return resp.StatusCode() == http.StatusNoContent, resp, nil
}
