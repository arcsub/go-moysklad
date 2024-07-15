package moysklad

import (
	"github.com/goccy/go-json"
)

// ApiError Структура ошибки API МойСклад.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-obschie-swedeniq-obrabotka-oshibok
type ApiError struct {
	Meta         *Meta       `json:"meta,omitempty"`          // Метаданные сущности, документа на котором произошла ошибка
	Header       string      `json:"error,omitempty"`         // Заголовок ошибки
	Parameter    string      `json:"parameter,omitempty"`     // Параметр, на котором произошла ошибка
	Message      string      `json:"error_message,omitempty"` // Сообщение, прилагаемое к ошибке
	MoreInfo     string      `json:"moreInfo,omitempty"`      // Ссылка на документацию с описанием полученной ошибки
	Dependencies Slice[Meta] `json:"dependencies,omitempty"`  // Список метаданных зависимых сущностей или документов. Выводится при невозможности удаления сущности, документа, если имеются зависимости от удаляемой сущности, документа
	Code         int         `json:"code,omitempty"`          // Код ошибки (Если поле ничего не содержит, смотрите HTTP status cod
	Line         int         `json:"line,omitempty"`          // Строка JSON, на которой произошла ошибка
	Column       int         `json:"column,omitempty"`        // Координата элемента в строке line, на котором произошла ошибка
}

// Error выводит ошибку в формате JSON
func (apiError ApiError) Error() string {
	b, _ := json.Marshal(apiError)
	return string(b)
}

// ApiErrors Структура ошибок API МойСклад.
type ApiErrors struct {
	ApiErrors Slice[ApiError] `json:"errors"` // Список ошибок
}

func (apiErrors ApiErrors) Error() string {
	b, _ := json.Marshal(apiErrors)
	return string(b)
}
