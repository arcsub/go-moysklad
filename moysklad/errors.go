package moysklad

import (
	"fmt"
)

// ApiError Структура ошибки API МойСклад.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-obschie-swedeniq-obrabotka-oshibok
type ApiError struct {
	Meta         Meta   `json:"meta"`
	Message      string `json:"error"`
	Parameter    string `json:"parameter"`
	ErrorMessage string `json:"error_message"`
	MoreInfo     string `json:"moreInfo"`
	Dependencies []Meta `json:"dependencies"`
	Code         int    `json:"code"`
	Line         int    `json:"line"`
	Column       int    `json:"column"`
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

type ErrUnknownEntity struct {
	t any
}

func (e ErrUnknownEntity) Error() string {
	return fmt.Sprintf("unrecognized entity: %v", e.t)
}

type ErrWrongMetaType struct {
	have, need MetaType
}

func (e ErrWrongMetaType) Error() string {
	return fmt.Sprintf("meta type mismatch! have %s, need %s", e.have, e.need)
}
