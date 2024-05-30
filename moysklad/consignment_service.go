package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// ConsignmentService
// Сервис для работы с сериями.
type ConsignmentService interface {
	GetList(ctx context.Context, params *Params) (*List[Consignment], *resty.Response, error)
	Create(ctx context.Context, consignment *Consignment, params *Params) (*Consignment, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, consignmentList []*Consignment, params *Params) (*[]Consignment, *resty.Response, error)
	DeleteMany(ctx context.Context, consignmentList []*Consignment) (*DeleteManyResponse, *resty.Response, error)
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
	DeleteAttributes(ctx context.Context, attributeList []*Attribute) (*DeleteManyResponse, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
}

func NewConsignmentService(client *Client) ConsignmentService {
	e := NewEndpoint(client, "entity/consignment")
	return newMainService[Consignment, any, MetadataAttribute, any](e)

}
