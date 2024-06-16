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

// Clean возвращает сущность с единственным заполненным полем Meta
func (priceType PriceType) Clean() *PriceType {
	return &PriceType{Meta: priceType.Meta}
}

func (priceType PriceType) GetExternalCode() string {
	return Deref(priceType.ExternalCode)
}

func (priceType PriceType) GetID() uuid.UUID {
	return Deref(priceType.ID)
}

func (priceType PriceType) GetMeta() Meta {
	return Deref(priceType.Meta)
}

func (priceType PriceType) GetName() string {
	return Deref(priceType.Name)
}

func (priceType *PriceType) SetExternalCode(externalCode string) *PriceType {
	priceType.ExternalCode = &externalCode
	return priceType
}

func (priceType *PriceType) SetMeta(meta *Meta) *PriceType {
	priceType.Meta = meta
	return priceType
}

func (priceType *PriceType) SetName(name string) *PriceType {
	priceType.Name = &name
	return priceType
}

func (priceType PriceType) String() string {
	return Stringify(priceType)
}

func (priceType PriceType) MetaType() MetaType {
	return MetaTypePriceType
}
