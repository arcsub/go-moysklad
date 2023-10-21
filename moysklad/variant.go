package moysklad

import (
	"github.com/google/uuid"
)

// Variant Модификация.
// Ключевое слово: variant
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-modifikaciq
type Variant struct {
	Archived           *bool            `json:"archived,omitempty"`
	Updated            *Timestamp       `json:"updated,omitempty"`
	AccountID          *uuid.UUID       `json:"accountId,omitempty"`
	Description        *string          `json:"description,omitempty"`
	ExternalCode       *string          `json:"externalCode,omitempty"`
	ID                 *uuid.UUID       `json:"id,omitempty"`
	Meta               *Meta            `json:"meta,omitempty"`
	Name               *string          `json:"name,omitempty"`
	Code               *string          `json:"code,omitempty"`
	Barcodes           *Barcodes        `json:"barcodes,omitempty"`
	DiscountProhibited *bool            `json:"discountProhibited,omitempty"`
	Characteristics    *Characteristics `json:"characteristics,omitempty"`
	Images             *Images          `json:"images,omitempty"`
	MinPrice           *MinPrice        `json:"minPrice,omitempty"`
	BuyPrice           *BuyPrice        `json:"buyPrice,omitempty"`
	Product            *Product         `json:"product,omitempty"`
	SalePrices         *SalePrices      `json:"salePrices,omitempty"`
	Things             *Things          `json:"things,omitempty"`
	Packs              []VariantPack    `json:"packs,omitempty"`
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
	ID         *uuid.UUID `json:"id,omitempty"`         // ID упаковки модификации
	ParentPack *Pack      `json:"parentpack,omitempty"` // Метаданные родительской упаковки (упаковки товара), для которой переопределяется штрихкод
}

func (v VariantPack) String() string {
	return Stringify(v)
}
