package moysklad

import (
	"github.com/google/uuid"
)

// Async Асинхронная задача.
// Ключевое слово: async
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-asinhronnyj-obmen-asinhronnaq-zadacha
type Async struct {
	Meta         Meta            `json:"meta,omitempty"`
	Owner        Meta            `json:"owner,omitempty"`
	DeletionDate Timestamp       `json:"deletionDate,omitempty"`
	RequestURL   string          `json:"request,omitempty"`
	ResultURL    string          `json:"resultUrl,omitempty"`
	State        AsyncState      `json:"state,omitempty"`
	ApiErrors    Slice[ApiError] `json:"errors,omitempty"`
	AccountID    uuid.UUID       `json:"accountId,omitempty"`
	ID           uuid.UUID       `json:"id,omitempty"`
}

func (a Async) String() string {
	return Stringify(a)
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
