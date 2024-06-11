package moysklad

import (
	"encoding/json"
	"fmt"
)

// ApiError Структура ошибки API МойСклад.
//
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-obschie-swedeniq-obrabotka-oshibok
type ApiError struct {
	Meta         *Meta       `json:"meta,omitempty"`
	Message      string      `json:"error,omitempty"`
	Parameter    string      `json:"parameter,omitempty"`
	ErrorMessage string      `json:"error_message,omitempty"`
	MoreInfo     string      `json:"moreInfo,omitempty"`
	Dependencies Slice[Meta] `json:"dependencies,omitempty"`
	Code         int         `json:"code,omitempty"`
	Line         int         `json:"line,omitempty"`
	Column       int         `json:"column,omitempty"`
}

// Error выводит ошибку в формате JSON
func (apiError ApiError) Error() string {
	b, _ := json.Marshal(apiError)
	return string(b)
}

type ApiErrors struct {
	ApiErrors Slice[ApiError] `json:"errors"`
}

func (apiErrors ApiErrors) Error() string {
	var message string
	for _, er := range apiErrors.ApiErrors {
		message += fmt.Sprintf("%v\n", er.Error())
	}
	return message
}
