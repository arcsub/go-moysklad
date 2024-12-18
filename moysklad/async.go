package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"

	"net/http"
)

// Async Асинхронная задача.
//
// Код сущности: async
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-asinhronnyj-obmen-asinhronnaq-zadacha
type Async struct {
	Meta         Meta        `json:"meta,omitempty"`         // Метаданные Асинхронной задачи
	Owner        MetaWrapper `json:"owner,omitempty"`        // Пользователь или приложение, которые создали Асинхронную задачу
	DeletionDate Timestamp   `json:"deletionDate,omitempty"` // Дата, после которой результат выполнения задачи станет недоступен. Содержится в ответе, если поле state имеет значение AsyncStateDone (DONE)
	RequestURL   string      `json:"request,omitempty"`      // URL запроса, по которому создана Асинхронная задача
	ResultURL    string      `json:"resultUrl,omitempty"`    // Ссылка на результат выполнения задачи. Содержится в ответе, если поле state имеет значение AsyncStateDone (DONE)
	State        AsyncState  `json:"state,omitempty"`        // Статус выполнения Асинхронной задачи.
	AccountID    string      `json:"accountId,omitempty"`    // ID учётной записи
	ID           string      `json:"id,omitempty"`           // ID Асинхронной задачи
	Errors       ApiErrors   `json:"errors,omitempty"`       // Ошибки апи, если поле state имеет значение AsyncStateApiError (API_ERROR)
}

// String реализует интерфейс [fmt.Stringer].
func (async Async) String() string {
	return Stringify(async)
}

// AsyncState Статус выполнения Асинхронной задачи.
//
// Возможные значения:
//   - AsyncStatePending     – Задача находится в очереди
//   - AsyncStateProcessing  – Задача находится в обработке, результат еще не готов
//   - AsyncStateDone        – Задача выполнена успешно
//   - AsyncStateError       – Задача не была выполнена в результате внутренней ошибки. В этом случае нужно попробовать запустить задачу заново
//   - AsyncStateCancel      – Задача была отменена
//   - AsyncStateApiError    – Задача была завершена с ошибкой апи
type AsyncState string

const (
	AsyncStatePending    AsyncState = "PENDING"    // Задача находится в очереди
	AsyncStateProcessing AsyncState = "PROCESSING" // Задача находится в обработке, результат еще не готов
	AsyncStateDone       AsyncState = "DONE"       // Задача выполнена успешно
	AsyncStateError      AsyncState = "ERROR"      // Задача не была выполнена в результате внутренней ошибки. В этом случае нужно попробовать запустить задачу заново
	AsyncStateCancel     AsyncState = "CANCEL"     // Задача была отменена
	AsyncStateApiError   AsyncState = "API_ERROR"  // Задача была завершена с ошибкой апи
)

// AsyncService методы сервиса для работы с асинхронными задачами.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-asinhronnyj-obmen
type AsyncService interface {
	// GetStatuses выполняет запрос на получение списка статусов выполнения Асинхронных задач.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Доступна фильтрация по полям state, request, deletionDate.
	// Результат содержит статусы Асинхронных задач за последнюю неделю.
	// Возвращает объект List.
	GetStatuses(ctx context.Context, params ...func(*Params)) (*List[Async], *resty.Response, error)

	// GetStatusByID выполняет запрос на получение статуса Асинхронной задачи.
	// Принимает контекст и ID асинхронной задачи.
	// Возвращает объект Async.
	GetStatusByID(ctx context.Context, id string) (*Async, *resty.Response, error)
}

type asyncService struct {
	Endpoint
}

func (service *asyncService) GetStatuses(ctx context.Context, params ...func(*Params)) (*List[Async], *resty.Response, error) {
	return NewRequestBuilder[List[Async]](service.client, service.uri).SetParams(params).Get(ctx)
}

func (service *asyncService) GetStatusByID(ctx context.Context, id string) (*Async, *resty.Response, error) {
	path := fmt.Sprintf(EndpointAsyncID, id)
	return NewRequestBuilder[Async](service.client, path).Get(ctx)
}

// AsyncResultService методы сервиса для обработки асинхронного запроса.
type AsyncResultService[T any] interface {
	// StatusURL возвращает URL проверки статуса асинхронной задачи.
	StatusURL() string

	// ResultURL возвращает URL результата выполнения асинхронной задачи.
	ResultURL() string

	// Check выполняет запрос на проверку статус асинхронной задачи.
	// Возвращает true, если статус задачи имеет значение AsyncStateDone (DONE).
	Check(ctx context.Context) (bool, *resty.Response, error)

	// Result выполняет запрос на получение результата.
	// Возвращает объект обобщённого типа, который был указан при создании сервиса для обработки асинхронного запроса.
	Result(ctx context.Context) (*T, *resty.Response, error)

	// Cancel выполняет запрос на отмену Асинхронной задачи.
	// Возвращает true, если задача успешно отменена.
	Cancel(ctx context.Context) (bool, *resty.Response, error)
}

type asyncResultService[T any] struct {
	client    *Client // Клиент
	statusURL string  // URL статуса Асинхронной задачи.
	resultURL string  // URL результата выполнения Асинхронной задачи.
}

const (
	EndpointAsync       = string(MetaTypeAsync)
	EndpointAsyncID     = EndpointAsync + "/%s"
	EndpointAsyncCancel = "%s/cancel"
)

// NewAsyncResultService принимает [Client] и возвращает сервис для работы с асинхронной задачей.
func NewAsyncResultService[T any](client *Client, resp *resty.Response) AsyncResultService[T] {
	return &asyncResultService[T]{
		client:    client,
		statusURL: resp.Header().Get("Location"),
		resultURL: resp.Header().Get("Content-Location"),
	}
}

func (service *asyncResultService[T]) StatusURL() string {
	return service.statusURL
}

func (service *asyncResultService[T]) ResultURL() string {
	return service.resultURL
}

func (service *asyncResultService[T]) Check(ctx context.Context) (bool, *resty.Response, error) {
	async, resp, err := NewRequestBuilder[Async](service.client, service.StatusURL()).Get(ctx)
	if err != nil {
		return false, resp, err
	}
	return async.State == AsyncStateDone, resp, nil
}

func (service *asyncResultService[T]) Result(ctx context.Context) (*T, *resty.Response, error) {
	data, resp, err := NewRequestBuilder[T](service.client, service.ResultURL()).Get(ctx)
	if err != nil {
		return nil, resp, err
	}
	return data, resp, nil
}

func (service *asyncResultService[T]) Cancel(ctx context.Context) (bool, *resty.Response, error) {
	path := fmt.Sprintf(EndpointAsyncCancel, service.StatusURL())
	_, resp, err := NewRequestBuilder[any](service.client, path).Post(ctx, nil)
	if err != nil {
		return false, resp, err
	}
	return resp.StatusCode() == http.StatusNoContent, resp, nil
}

// NewAsyncService принимает [Client] и возвращает сервис для работы с асинхронными задачами.
func NewAsyncService(client *Client) AsyncService {
	return &asyncService{NewEndpoint(client, EndpointAsync)}
}
