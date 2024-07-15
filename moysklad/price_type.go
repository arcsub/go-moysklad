package moysklad

import (
	"github.com/google/uuid"
)

// PriceType Тип цены.
//
// Код сущности: pricetype
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tipy-cen
type PriceType struct {
	ExternalCode *string    `json:"externalCode,omitempty"` // Внешний код Типа цены
	ID           *uuid.UUID `json:"id,omitempty"`           // ID типа цены
	Meta         *Meta      `json:"meta,omitempty"`         // Метаданные Типа цены
	Name         *string    `json:"name,omitempty"`         // Наименование Типа цены
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (priceType PriceType) Clean() *PriceType {
	if priceType.Meta == nil {
		return nil
	}
	return &PriceType{Meta: priceType.Meta}
}

// GetExternalCode возвращает Внешний код Типа цены.
func (priceType PriceType) GetExternalCode() string {
	return Deref(priceType.ExternalCode)
}

// GetID возвращает ID типа цены.
func (priceType PriceType) GetID() uuid.UUID {
	return Deref(priceType.ID)
}

// GetMeta возвращает Метаданные Типа цены.
func (priceType PriceType) GetMeta() Meta {
	return Deref(priceType.Meta)
}

// GetName возвращает Наименование Типа цены.
func (priceType PriceType) GetName() string {
	return Deref(priceType.Name)
}

// SetExternalCode устанавливает Внешний код Типа цены.
func (priceType *PriceType) SetExternalCode(externalCode string) *PriceType {
	priceType.ExternalCode = &externalCode
	return priceType
}

// SetMeta устанавливает Метаданные Типа цены.
func (priceType *PriceType) SetMeta(meta *Meta) *PriceType {
	priceType.Meta = meta
	return priceType
}

// SetName устанавливает Наименование Типа цены.
func (priceType *PriceType) SetName(name string) *PriceType {
	priceType.Name = &name
	return priceType
}

// String реализует интерфейс [fmt.Stringer].
func (priceType PriceType) String() string {
	return Stringify(priceType)
}

// MetaType возвращает код сущности.
func (PriceType) MetaType() MetaType {
	return MetaTypePriceType
}
