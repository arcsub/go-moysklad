package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
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
	TrackingType        *TrackingType  `json:"trackingType,omitempty"`
	TaxSystem           *GoodTaxSystem `json:"taxSystem,omitempty"`
	PaymentItemType     *PaymentItem   `json:"paymentItemType,omitempty"`
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

// ProductService
// Сервис для работы с товарами.
type ProductService interface {
	GetList(ctx context.Context, params *Params) (*List[Product], *resty.Response, error)
	Create(ctx context.Context, product *Product, params *Params) (*Product, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, productList []*Product, params *Params) (*[]Product, *resty.Response, error)
	DeleteMany(ctx context.Context, productList []*Product) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetadataAttributeShared, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id *uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error)
	UpdateAttribute(ctx context.Context, id *uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList []*Attribute) (*DeleteManyResponse, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Product, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, product *Product, params *Params) (*Product, *resty.Response, error)
	GetImages(ctx context.Context, id *uuid.UUID) (*MetaArray[Image], *resty.Response, error)
	CreateImage(ctx context.Context, id *uuid.UUID, image *Image) (*[]*Image, *resty.Response, error)
	UpdateImages(ctx context.Context, id *uuid.UUID, images []*Image) (*[]Image, *resty.Response, error)
	DeleteImage(ctx context.Context, id *uuid.UUID, imageId *uuid.UUID) (bool, *resty.Response, error)
	DeleteImages(ctx context.Context, id *uuid.UUID, images []*Image) (*DeleteManyResponse, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*Product, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
	GetAudit(ctx context.Context, id *uuid.UUID, params *Params) (*List[AuditEvent], *resty.Response, error)
	PrintLabel(ctx context.Context, id *uuid.UUID, PrintLabelArg *PrintLabelArg) (*PrintFile, *resty.Response, error)
}

func NewProductService(client *Client) ProductService {
	e := NewEndpoint(client, "entity/product")
	return newMainService[Product, any, MetadataAttributeShared, any](e)
}
