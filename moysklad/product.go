package moysklad

import (
	"github.com/google/uuid"
)

// Product Товар.
// Ключевое слово: product
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-towar
type Product struct {
	MinimumBalance      *float64       `json:"minimumBalance,omitempty"`
	UseParentVat        *bool          `json:"useParentVat,omitempty"`
	Code                *string        `json:"code,omitempty"`
	Description         *string        `json:"description,omitempty"`
	ExternalCode        *string        `json:"externalCode,omitempty"`
	ID                  *uuid.UUID     `json:"id,omitempty"`
	Meta                *Meta          `json:"meta,omitempty"`
	Name                *string        `json:"name,omitempty"`
	Alcoholic           *Alcoholic     `json:"alcoholic,omitempty"`
	Archived            *bool          `json:"archived,omitempty"`
	Article             *string        `json:"article,omitempty"`
	Attributes          *Attributes    `json:"attributes,omitempty"`
	BuyPrice            *BuyPrice      `json:"buyPrice,omitempty"`
	Country             *Country       `json:"country,omitempty"`
	DiscountProhibited  *bool          `json:"discountProhibited,omitempty"`
	EffectiveVat        *int           `json:"effectiveVat,omitempty"`
	EffectiveVatEnabled *bool          `json:"effectiveVatEnabled,omitempty"`
	Files               *Files         `json:"files,omitempty"`
	Group               *Group         `json:"group,omitempty"`
	Images              *Images        `json:"images,omitempty"`
	IsSerialTrackable   *bool          `json:"isSerialTrackable,omitempty"`
	MinPrice            *MinPrice      `json:"minPrice,omitempty"`
	Volume              *float64       `json:"volume,omitempty"`
	Barcodes            *Barcodes      `json:"barcodes,omitempty"`
	PathName            *string        `json:"pathName,omitempty"`
	Packs               *Packs         `json:"packs,omitempty"`
	PartialDisposal     *bool          `json:"partialDisposal,omitempty"`
	Owner               *Employee      `json:"owner,omitempty"`
	Weight              *float64       `json:"weight,omitempty"`
	PpeType             *string        `json:"ppeType,omitempty"`
	ProductFolder       *ProductFolder `json:"productFolder,omitempty"`
	SalePrices          *SalePrices    `json:"salePrices,omitempty"`
	Shared              *bool          `json:"shared,omitempty"`
	Supplier            *Counterparty  `json:"supplier,omitempty"`
	SyncID              *uuid.UUID     `json:"syncId,omitempty"`
	AccountID           *uuid.UUID     `json:"accountId,omitempty"`
	Things              *Things        `json:"things,omitempty"`
	Tnved               *string        `json:"tnved,omitempty"`
	VatEnabled          *bool          `json:"vatEnabled,omitempty"`
	Uom                 *Uom           `json:"uom,omitempty"`
	Updated             *Timestamp     `json:"updated,omitempty"`
	OnTap               *bool          `json:"onTap,omitempty"`
	VariantsCount       *int           `json:"variantsCount,omitempty"`
	Vat                 *int           `json:"vat,omitempty"`
	TrackingType        TrackingType   `json:"trackingType,omitempty"`
	TaxSystem           GoodTaxSystem  `json:"taxSystem,omitempty"`
	PaymentItemType     PaymentItem    `json:"paymentItemType,omitempty"`
}

func (p Product) String() string {
	return Stringify(p)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (p Product) GetMeta() *Meta {
	return p.Meta
}

func (p Product) MetaType() MetaType {
	return MetaTypeProduct
}

func (p Product) ConvertToAssortmentPosition() (*AssortmentPosition, error) {
	return convertToAssortmentPosition(p)
}

// Alcoholic Объект, содержащий поля алкогольной продукции
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-towar-towary-atributy-wlozhennyh-suschnostej-ob-ekt-soderzhaschij-polq-alkogol-noj-produkcii
type Alcoholic struct {
	Excise   *bool    `json:"excise,omitempty"`   // Содержит акцизную марку
	Type     *int     `json:"type,omitempty"`     // Код вида продукции
	Strength *float64 `json:"strength,omitempty"` // Крепость
	Volume   *float64 `json:"volume,omitempty"`   // Объём тары
}

func (a Alcoholic) String() string {
	return Stringify(a)
}
