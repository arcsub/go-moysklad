package moysklad

import (
	"github.com/google/uuid"
)

// Variant Модификация.
// Ключевое слово: variant
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-modifikaciq
type Variant struct {
	AccountId          *uuid.UUID       `json:"accountId,omitempty"`          // ID учетной записи
	Barcodes           *Barcodes        `json:"barcodes,omitempty"`           // Штрихкоды
	Code               *string          `json:"code,omitempty"`               // Код
	Description        *string          `json:"description,omitempty"`        // Описание
	ExternalCode       *string          `json:"externalCode,omitempty"`       // Внешний код
	Id                 *uuid.UUID       `json:"id,omitempty"`                 // ID сущности
	Meta               *Meta            `json:"meta,omitempty"`               // Метаданные
	Name               *string          `json:"name,omitempty"`               // Наименование
	Archived           *bool            `json:"archived,omitempty"`           // Добавлен ли товар в архив
	BuyPrice           *BuyPrice        `json:"buyPrice,omitempty"`           // Закупочная цена
	Characteristics    *Characteristics `json:"characteristics,omitempty"`    // Характеристики Модификации
	DiscountProhibited *bool            `json:"discountProhibited,omitempty"` // Признак запрета скидок
	Images             *Images          `json:"images,omitempty"`             // Массив метаданных Изображений (Максимальное количество изображений - 10)
	MinPrice           *MinPrice        `json:"minPrice,omitempty"`           // Минимальная цена
	Packs              []VariantPack    `json:"packs,omitempty"`              // Упаковки модификации
	Product            *Product         `json:"product,omitempty"`            // Метаданные товара, к которому привязана Модификация
	SalePrices         *SalePrices      `json:"salePrices,omitempty"`         // Цены продажи
	Things             *Things          `json:"things,omitempty"`             // Серийные номера
	Updated            *Timestamp       `json:"updated,omitempty"`            // Момент последнего обновления сущности
}

func (v Variant) String() string {
	return Stringify(v)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (v Variant) GetMeta() *Meta {
	return v.Meta
}

func (v Variant) MetaType() MetaType {
	return MetaTypeVariant
}

func (v Variant) ConvertToAssortmentPosition() (*AssortmentPosition, error) {
	return convertToAssortmentPosition(v)
}

// VariantPack Упаковка модификации.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-modifikaciq-modifikacii-atributy-wlozhennyh-suschnostej-upakowki-modifikacii
type VariantPack struct {
	Barcodes   *Barcodes  `json:"barcodes,omitempty"`   // Массив штрихкодов упаковки модификации. Данный массив может содержать только один штрихкод
	Id         *uuid.UUID `json:"id,omitempty"`         // ID упаковки модификации
	ParentPack *Pack      `json:"parentpack,omitempty"` // Метаданные родительской упаковки (упаковки товара), для которой переопределяется штрихкод
}

func (v VariantPack) String() string {
	return Stringify(v)
}
