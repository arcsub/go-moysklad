package moysklad

import (
	"github.com/google/uuid"
)

// Region Регион.
// Ключевое слово: region
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-region
type Region struct {
	AccountId    *uuid.UUID `json:"accountId,omitempty"`    // ID учетной записи
	Code         *string    `json:"code,omitempty"`         // Код Региона
	ExternalCode *string    `json:"externalCode,omitempty"` // Внешний код Региона
	Id           *uuid.UUID `json:"id,omitempty"`           // ID Региона
	Meta         *Meta      `json:"meta,omitempty"`         // Метаданные
	Name         *string    `json:"name,omitempty"`         // Наименование
	Updated      *Timestamp `json:"updated,omitempty"`      // Момент последнего обновления сущности
	Version      *int       `json:"version,omitempty"`      // Версия сущности
}

func (r Region) String() string {
	return Stringify(r)
}

func (r Region) MetaType() MetaType {
	return MetaTypeRegion
}
