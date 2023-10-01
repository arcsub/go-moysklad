package moysklad

import (
	"github.com/google/uuid"
)

// Uom Единица измерения.
// Ключевое слово: uom
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-edinica-izmereniq
type Uom struct {
	AccountId    *uuid.UUID `json:"accountId,omitempty"`    // ID учетной записи
	Code         *string    `json:"code,omitempty"`         // Код Единицы измерения
	Description  *string    `json:"description,omitempty"`  // Описание Единциы измерения
	ExternalCode *string    `json:"externalCode,omitempty"` // Внешний код Единицы измерения
	Group        *Group     `json:"group,omitempty"`        // Отдел сотрудника
	Id           *uuid.UUID `json:"id,omitempty"`           // ID сущности
	Meta         *Meta      `json:"meta,omitempty"`         // Метаданные
	Name         *string    `json:"name,omitempty"`         // Наименование Единицы измерения
	Owner        *Employee  `json:"owner,omitempty"`        // Владелец (Сотрудник)
	Shared       *bool      `json:"shared,omitempty"`       // Общий доступ
	Updated      *Timestamp `json:"updated,omitempty"`      // Момент последнего обновления Единицы измерения
}

func (u Uom) String() string {
	return Stringify(u)
}

func (u Uom) MetaType() MetaType {
	return MetaTypeUom
}
