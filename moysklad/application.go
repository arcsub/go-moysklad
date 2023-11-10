package moysklad

import (
	"github.com/google/uuid"
)

// Application Серверное приложение.
// Ключевое слово: application
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-obschie-swedeniq-serwernye-prilozheniq
type Application struct {
	AccountID *uuid.UUID `json:"accountId,omitempty"` // ID учетной записи
	ID        *uuid.UUID `json:"id,omitempty"`        // ID сущности
	Name      *string    `json:"name,omitempty"`      // Наименование
	Meta      *Meta      `json:"meta,omitempty"`      // Метаданные
	AppUid    *uuid.UUID `json:"appUid,omitempty"`    // UID приложения
}

func (a Application) String() string {
	return Stringify(a)
}

func (a Application) MetaType() MetaType {
	return MetaTypeApplication
}
