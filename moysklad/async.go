package moysklad

import (
	"github.com/google/uuid"
)

// Async Асинхронная задача.
// Ключевое слово: async
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-asinhronnyj-obmen-asinhronnaq-zadacha
type Async struct {
	AccountID    uuid.UUID       `json:"accountId,omitempty"`    // ID учетной записи
	DeletionDate Timestamp       `json:"deletionDate,omitempty"` // Дата, после которой результат выполнения задачи станет недоступен. Содержится в ответе, если поле state имеет значение DONE
	ApiErrors    Slice[ApiError] `json:"errors,omitempty"`       // json ошибки апи, если поле state имеет значение API_ERROR
	ID           uuid.UUID       `json:"id,omitempty"`           // ID Асинхронной задачи
	Meta         Meta            `json:"meta,omitempty"`         // Метаданные Асинхронной задачи
	Owner        Meta            `json:"owner,omitempty"`        // Пользователь или приложение, которые создали Асинхронную задачу
	RequestURL   string          `json:"request,omitempty"`      // URL запроса, по которому создана Асинхронная задача
	ResultURL    string          `json:"resultUrl,omitempty"`    // Ссылка на результат выполнения задачи. Содержится в ответе, если поле state имеет значение DONE
	State        AsyncState      `json:"state,omitempty"`        // Статус выполнения Асинхронной задачи
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
