package moysklad

import (
	"github.com/google/uuid"
)

// PriceType Тип цены.
// Ключевое слово: pricetype
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tipy-cen
type PriceType struct {
	ExternalCode *string    `json:"externalCode,omitempty"` // Внешний код Типа цены
	ID           *uuid.UUID `json:"id,omitempty"`           // ID типа цены
	Meta         *Meta      `json:"meta,omitempty"`         // Метаданные Типа цены
	Name         *string    `json:"name,omitempty"`         // Наименование Типа цены
}

func (p PriceType) String() string {
	return Stringify(p)
}

func (p PriceType) MetaType() MetaType {
	return MetaTypePriceType
}

type PriceTypes = Slice[PriceType]
