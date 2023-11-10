package moysklad

import (
	"github.com/google/uuid"
)

// Country Страна.
// Ключевое слово: country
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-strana
type Country struct {
	AccountID    *uuid.UUID `json:"accountId,omitempty"`    // ID учетной записи
	Code         *string    `json:"code,omitempty"`         // Код Страны
	Description  *string    `json:"description,omitempty"`  // Описание Страны
	ExternalCode *string    `json:"externalCode,omitempty"` // Внешний код Страны
	Group        *Group     `json:"group,omitempty"`        // Отдел-владелец
	ID           *uuid.UUID `json:"id,omitempty"`           // ID сущности
	Meta         *Meta      `json:"meta,omitempty"`         // Метаданные
	Name         *string    `json:"name,omitempty"`         // Наименование
	Owner        *Employee  `json:"owner,omitempty"`        // Сотрудник-владелец
	Shared       *bool      `json:"shared,omitempty"`       // Флаг Общий доступ
	Updated      *Timestamp `json:"updated,omitempty"`      // Момент последнего обновления сущности
}

func (c Country) String() string {
	return Stringify(c)
}

func (c Country) MetaType() MetaType {
	return MetaTypeCountry
}
