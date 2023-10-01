package moysklad

import (
	"fmt"
)

// ApiError Структура ошибки API МойСклад.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-obschie-swedeniq-obrabotka-oshibok
type ApiError struct {
	Message      string `json:"error"`         // Заголовок ошибки (обязательное при ответе)
	Parameter    string `json:"parameter"`     // Параметр, на котором произошла ошибка
	Code         int    `json:"code"`          // Код ошибки (Если поле ничего не содержит, смотрите HTTP status code)
	ErrorMessage string `json:"error_message"` // Сообщение, прилагаемое к ошибке
	MoreInfo     string `json:"moreInfo"`      // Ссылка на документацию с описанием полученной ошибки
	Line         int    `json:"line"`          // Строка JSON, на которой произошла ошибка
	Column       int    `json:"column"`        // Координата элемента в строке line, на котором произошла ошибка
	Dependencies []Meta `json:"dependencies"`  // Список метаданных зависимых сущностей или документов. Выводится при невозможности удаления сущности, документа, если имеются зависимости от удаляемой сущности, документа
	Meta         Meta   `json:"meta"`          // Метаданные сущности, документа на котором произошла ошибка
}

func (e ApiError) Error() string {
	return fmt.Sprintf("%v %v %v %v", e.Code, e.Message, e.ErrorMessage, e.MoreInfo)
}

type ApiErrors struct {
	ApiErrors []ApiError `json:"errors"`
}

func (e ApiErrors) Error() string {
	var message string
	for i, er := range e.ApiErrors {
		message += fmt.Sprintf("%d) %v %v %v %v\n", i+1, er.Code, er.Message, er.ErrorMessage, er.MoreInfo)
	}
	return message
}
