package moysklad

import (
	"github.com/google/uuid"
)

// Service Услуга.
// Ключевое слово: service
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-usluga
type Service struct {
	Files               *Files         `json:"files,omitempty"`
	Code                *string        `json:"code,omitempty"`
	Group               *Group         `json:"group,omitempty"`
	Description         *string        `json:"description,omitempty"`
	ExternalCode        *string        `json:"externalCode,omitempty"`
	ID                  *uuid.UUID     `json:"id,omitempty"`
	Meta                *Meta          `json:"meta,omitempty"`
	Name                *string        `json:"name,omitempty"`
	Archived            *bool          `json:"archived,omitempty"`
	Attributes          *Attributes    `json:"attributes,omitempty"`
	BuyPrice            *BuyPrice      `json:"buyPrice,omitempty"`
	DiscountProhibited  *bool          `json:"discountProhibited,omitempty"`
	EffectiveVat        *int           `json:"effectiveVat,omitempty"`
	EffectiveVatEnabled *bool          `json:"effectiveVatEnabled,omitempty"`
	VatEnabled          *bool          `json:"vatEnabled,omitempty"`
	Barcodes            *Barcodes      `json:"barcodes,omitempty"`
	MinPrice            *MinPrice      `json:"minPrice,omitempty"`
	Owner               *Employee      `json:"owner,omitempty"`
	PathName            *string        `json:"pathName,omitempty"`
	AccountID           *uuid.UUID     `json:"accountId,omitempty"`
	ProductFolder       *ProductFolder `json:"productFolder,omitempty"`
	SalePrices          *SalePrices    `json:"salePrices,omitempty"`
	Shared              *bool          `json:"shared,omitempty"`
	SyncID              *uuid.UUID     `json:"syncId,omitempty"`
	Vat                 *int           `json:"vat,omitempty"`
	Uom                 *Uom           `json:"uom,omitempty"`
	Updated             *Timestamp     `json:"updated,omitempty"`
	UseParentVat        *bool          `json:"useParentVat,omitempty"`
	TaxSystem           TaxSystem      `json:"taxSystem,omitempty"`
	PaymentItemType     PaymentItem    `json:"paymentItemType,omitempty"`
}

func (s Service) String() string {
	return Stringify(s)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (s Service) GetMeta() *Meta {
	return s.Meta
}

func (s Service) MetaType() MetaType {
	return MetaTypeService
}

func (s Service) ConvertToAssortmentPosition() (*AssortmentPosition, error) {
	return convertToAssortmentPosition(s)
}
