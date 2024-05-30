package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// RetailDrawerCashInService
// Сервис для работы с внесениями денег.
type RetailDrawerCashInService interface {
	GetList(ctx context.Context, params *Params) (*List[RetailDrawerCashIn], *resty.Response, error)
	Create(ctx context.Context, retailDrawerCashIn *RetailDrawerCashIn, params *Params) (*RetailDrawerCashIn, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, retailDrawerCashInList []*RetailDrawerCashIn, params *Params) (*[]RetailDrawerCashIn, *resty.Response, error)
	DeleteMany(ctx context.Context, retailDrawerCashInList []*RetailDrawerCashIn) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*RetailDrawerCashIn, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, retailDrawerCashIn *RetailDrawerCashIn, params *Params) (*RetailDrawerCashIn, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	//endpointTemplate[RetailDrawerCashIn]
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id *uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error)
	UpdateAttribute(ctx context.Context, id *uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList []*Attribute) (*DeleteManyResponse, *resty.Response, error)
	GetPublications(ctx context.Context, id *uuid.UUID) (*MetaArray[Publication], *resty.Response, error)
	GetPublicationByID(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (*Publication, *resty.Response, error)
	Publish(ctx context.Context, id *uuid.UUID, template *Templater) (*Publication, *resty.Response, error)
	DeletePublication(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (bool, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*RetailDrawerCashIn, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
	Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewRetailDrawerCashInService(client *Client) RetailDrawerCashInService {
	e := NewEndpoint(client, "entity/retaildrawercashin")
	return newMainService[RetailDrawerCashIn, any, MetadataAttributeSharedStates, any](e)
}
