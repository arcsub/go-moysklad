package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Consignment Серия.
// Ключевое слово: consignment
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-seriq
type Consignment struct {
	Meta         *Meta                 `json:"meta,omitempty"`
	Barcodes     Slice[Barcode]        `json:"barcodes,omitempty"`
	Code         *string               `json:"code,omitempty"`
	Description  *string               `json:"description,omitempty"`
	ExternalCode *string               `json:"externalCode,omitempty"`
	ID           *uuid.UUID            `json:"id,omitempty"`
	AccountID    *uuid.UUID            `json:"accountId,omitempty"`
	Name         *string               `json:"name,omitempty"`
	Assortment   *AssortmentPosition   `json:"assortment,omitempty"`
	Image        *Image                `json:"image,omitempty"`
	Label        *string               `json:"label,omitempty"`
	Updated      *Timestamp            `json:"updated,omitempty"`
	Attributes   Slice[AttributeValue] `json:"attributes,omitempty"`
}

func NewConsignmentFromAssortment(assortmentPosition AssortmentPosition) *Consignment {
	return unmarshalAsType[Consignment](assortmentPosition)
}

func (consignment Consignment) FromAssortment(assortmentPosition AssortmentPosition) *Consignment {
	return unmarshalAsType[Consignment](assortmentPosition)
}

func (consignment Consignment) AsAssortment() *AssortmentPosition {
	return &AssortmentPosition{Meta: consignment.GetMeta()}
}

func (consignment Consignment) GetMeta() Meta {
	return Deref(consignment.Meta)
}

func (consignment Consignment) GetBarcodes() Slice[Barcode] {
	return consignment.Barcodes
}

func (consignment Consignment) GetCode() string {
	return Deref(consignment.Code)
}

func (consignment Consignment) GetDescription() string {
	return Deref(consignment.Description)
}

func (consignment Consignment) GetExternalCode() string {
	return Deref(consignment.ExternalCode)
}

func (consignment Consignment) GetID() uuid.UUID {
	return Deref(consignment.ID)
}

func (consignment Consignment) GetAccountID() uuid.UUID {
	return Deref(consignment.AccountID)
}

func (consignment Consignment) GetName() string {
	return Deref(consignment.Name)
}

func (consignment Consignment) GetAssortment() AssortmentPosition {
	return Deref(consignment.Assortment)
}

func (consignment Consignment) GetImage() Image {
	return Deref(consignment.Image)
}

func (consignment Consignment) GetLabel() string {
	return Deref(consignment.Label)
}

func (consignment Consignment) GetUpdated() Timestamp {
	return Deref(consignment.Updated)
}

func (consignment Consignment) GetAttributes() Slice[AttributeValue] {
	return consignment.Attributes
}

func (consignment *Consignment) SetMeta(meta *Meta) *Consignment {
	consignment.Meta = meta
	return consignment
}

func (consignment *Consignment) SetBarcodes(barcodes Slice[Barcode]) *Consignment {
	consignment.Barcodes = barcodes
	return consignment
}

func (consignment *Consignment) SetCode(code string) *Consignment {
	consignment.Code = &code
	return consignment
}

func (consignment *Consignment) SetDescription(description string) *Consignment {
	consignment.Description = &description
	return consignment
}

func (consignment *Consignment) SetExternalCode(externalCode string) *Consignment {
	consignment.ExternalCode = &externalCode
	return consignment
}

func (consignment *Consignment) SetAssortment(assortment AsAssortment) *Consignment {
	consignment.Assortment = assortment.AsAssortment()
	return consignment
}

func (consignment *Consignment) SetImage(image *Image) *Consignment {
	consignment.Image = image
	return consignment
}

func (consignment *Consignment) SetLabel(label string) *Consignment {
	consignment.Label = &label
	return consignment
}

func (consignment *Consignment) SetAttributes(attributes Slice[AttributeValue]) *Consignment {
	consignment.Attributes = attributes
	return consignment
}

func (consignment Consignment) String() string {
	return Stringify(consignment)
}

func (consignment Consignment) MetaType() MetaType {
	return MetaTypeConsignment
}

// ConsignmentService
// Сервис для работы с сериями.
type ConsignmentService interface {
	GetList(ctx context.Context, params *Params) (*List[Consignment], *resty.Response, error)
	Create(ctx context.Context, consignment *Consignment, params *Params) (*Consignment, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, consignmentList []*Consignment, params *Params) (*[]Consignment, *resty.Response, error)
	DeleteMany(ctx context.Context, consignmentList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params *Params) (*Consignment, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, consignment *Consignment, params *Params) (*Consignment, *resty.Response, error)
	GetMetadata(context.Context) (*MetaAttributesWrapper, *resty.Response, error)
	GetAttributes(context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error)
	UpdateAttribute(ctx context.Context, id uuid.UUID, attr *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)
}

func NewConsignmentService(client *Client) ConsignmentService {
	e := NewEndpoint(client, "entity/consignment")
	return newMainService[Consignment, any, MetaAttributesWrapper, any](e)

}
