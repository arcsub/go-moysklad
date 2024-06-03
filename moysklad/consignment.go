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
	Meta         *Meta               `json:"meta,omitempty"`
	Barcodes     Barcodes            `json:"barcodes,omitempty"`
	Code         *string             `json:"code,omitempty"`
	Description  *string             `json:"description,omitempty"`
	ExternalCode *string             `json:"externalCode,omitempty"`
	ID           *uuid.UUID          `json:"id,omitempty"`
	AccountID    *uuid.UUID          `json:"accountId,omitempty"`
	Name         *string             `json:"name,omitempty"`
	Assortment   *AssortmentPosition `json:"assortment,omitempty"`
	Image        *Image              `json:"image,omitempty"`
	Label        *string             `json:"label,omitempty"`
	Updated      *Timestamp          `json:"updated,omitempty"`
	Attributes   Attributes          `json:"attributes,omitempty"`
}

func (c Consignment) String() string {
	return Stringify(c)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (c Consignment) GetMeta() Meta {
	return Deref(c.Meta)
}

func (c Consignment) MetaType() MetaType {
	return MetaTypeConsignment
}

// ConsignmentService
// Сервис для работы с сериями.
type ConsignmentService interface {
	GetList(ctx context.Context, params *Params) (*List[Consignment], *resty.Response, error)
	Create(ctx context.Context, consignment *Consignment, params *Params) (*Consignment, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, consignmentList []*Consignment, params *Params) (*[]Consignment, *resty.Response, error)
	DeleteMany(ctx context.Context, consignmentList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Consignment, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, consignment *Consignment, params *Params) (*Consignment, *resty.Response, error)
	GetMetadata(context.Context) (*MetadataAttribute, *resty.Response, error)
	GetAttributes(context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id *uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error)
	UpdateAttribute(ctx context.Context, id *uuid.UUID, attr *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
}

func NewConsignmentService(client *Client) ConsignmentService {
	e := NewEndpoint(client, "entity/consignment")
	return newMainService[Consignment, any, MetadataAttribute, any](e)

}
