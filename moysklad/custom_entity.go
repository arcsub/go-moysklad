package moysklad

import (
	"github.com/google/uuid"
)

// CustomEntity Пользовательский справочник.
// Ключевое слово: customentity
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-pol-zowatel-skij-sprawochnik
type CustomEntity struct {
	ID   *uuid.UUID `json:"id,omitempty"`   // ID Пользовательского справочника
	Meta *Meta      `json:"meta,omitempty"` // Метаданные Пользовательского справочника
	Name *string    `json:"name,omitempty"` // Наименование Пользовательского справочника
}

func (c CustomEntity) String() string {
	return Stringify(c)
}

func (c CustomEntity) MetaType() MetaType {
	return MetaTypeCustomEntity
}

// CustomEntityElement Элемент Пользовательского справочника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-pol-zowatel-skij-sprawochnik-jelementy-pol-zowatel-skogo-sprawochnika
type CustomEntityElement struct {
	AccountID    *uuid.UUID `json:"accountId,omitempty"`    // ID учетной записи
	Code         *string    `json:"code,omitempty"`         // Код элемента Пользовательского справочника
	Description  *string    `json:"description,omitempty"`  // Описание элемента Пользовательского справочника
	ExternalCode *string    `json:"externalCode,omitempty"` // Внешний код элемента Пользовательского справочника
	ID           *uuid.UUID `json:"id,omitempty"`           // ID сущности
	Meta         *Meta      `json:"meta,omitempty"`         // Метаданные
	Name         *string    `json:"name,omitempty"`         // Наименование
	Updated      *Timestamp `json:"updated,omitempty"`      // Момент последнего обновления элементе Пользовательского справочника
	Group        *Group     `json:"group,omitempty"`        // Отдел сотрудника
	Owner        *Employee  `json:"owner,omitempty"`        // Владелец (Сотрудник)
	Shared       *bool      `json:"shared,omitempty"`       // Общий доступ
}

func (c CustomEntityElement) String() string {
	return Stringify(c)
}
