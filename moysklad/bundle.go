package moysklad

import (
	"github.com/google/uuid"
)

// Bundle Комплект.
// Ключевое слово: bundle
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-komplekt
type Bundle struct {
	Images              *Images                     `json:"images,omitempty"`
	Updated             *Timestamp                  `json:"updated,omitempty"`
	Code                *string                     `json:"code,omitempty"`
	Description         *string                     `json:"description,omitempty"`
	ExternalCode        *string                     `json:"externalCode,omitempty"`
	ID                  *uuid.UUID                  `json:"id,omitempty"`
	Meta                *Meta                       `json:"meta,omitempty"`
	Name                *string                     `json:"name,omitempty"`
	Archived            *bool                       `json:"archived,omitempty"`
	Article             *string                     `json:"article,omitempty"`
	Attributes          *Attributes                 `json:"attributes,omitempty"`
	Components          *Positions[BundleComponent] `json:"components,omitempty"`
	Country             *Country                    `json:"country,omitempty"`
	Overhead            *BundleOverhead             `json:"overhead,omitempty"`
	EffectiveVat        *int                        `json:"effectiveVat,omitempty"`
	EffectiveVatEnabled *bool                       `json:"effectiveVatEnabled,omitempty"`
	Files               *Files                      `json:"files,omitempty"`
	Group               *Group                      `json:"group,omitempty"`
	Volume              *float64                    `json:"volume,omitempty"`
	Barcodes            *Barcodes                   `json:"barcodes,omitempty"`
	DiscountProhibited  *bool                       `json:"discountProhibited"`
	Owner               *Employee                   `json:"owner,omitempty"`
	PartialDisposal     *bool                       `json:"partialDisposal,omitempty"`
	PathName            *string                     `json:"pathName,omitempty"`
	Weight              *float64                    `json:"weight,omitempty"`
	SalePrices          *SalePrices                 `json:"salePrices,omitempty"`
	ProductFolder       *ProductFolder              `json:"productFolder,omitempty"`
	Shared              *bool                       `json:"shared,omitempty"`
	SyncID              *uuid.UUID                  `json:"syncId,omitempty"`
	AccountID           *uuid.UUID                  `json:"accountId,omitempty"`
	Tnved               *string                     `json:"tnved,omitempty"`
	VatEnabled          *bool                       `json:"vatEnabled,omitempty"`
	Uom                 *Uom                        `json:"uom,omitempty"`
	MinPrice            *MinPrice                   `json:"minPrice,omitempty"`
	UseParentVat        *bool                       `json:"useParentVat,omitempty"`
	Vat                 *int                        `json:"vat,omitempty"`
	TrackingType        TrackingType                `json:"trackingType,omitempty"`
	TaxSystem           GoodTaxSystem               `json:"taxSystem,omitempty"`
	PaymentItemType     PaymentItem                 `json:"paymentItemType,omitempty"`
}

func (b Bundle) String() string {
	return Stringify(b)
}

func (b Bundle) MetaType() MetaType {
	return MetaTypeBundle
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (b Bundle) GetMeta() *Meta {
	return b.Meta
}

func (b Bundle) ConvertToAssortmentPosition() (*AssortmentPosition, error) {
	return convertToAssortmentPosition(b)
}

// BundleOverhead Дополнительные расходы
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-komplekt-komplekty-atributy-wlozhennyh-suschnostej-dopolnitel-nye-rashody
type BundleOverhead struct {
	Value    *Decimal  `json:"value,omitempty"`    // Значение цены
	Currency *Currency `json:"currency,omitempty"` // Ссылка на валюту в формате Метаданных
}

func (b BundleOverhead) String() string {
	return Stringify(b)
}

// BundleComponent Компонент комплекта.
// Ключевое слово: bundlecomponent
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-komplekt-komplekty-komponenty-komplekta
type BundleComponent struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги, которую представляет собой компонент
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID сущности
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в компоненте
}

func (b BundleComponent) String() string {
	return Stringify(b)
}

func (b BundleComponent) MetaType() MetaType {
	return MetaTypeBundleComponent
}
