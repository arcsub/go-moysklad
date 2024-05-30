package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// PaymentInService
// Сервис для работы с входящими платежами.
type PaymentInService interface {
	GetList(ctx context.Context, params *Params) (*List[PaymentIn], *resty.Response, error)
	Create(ctx context.Context, paymentIn *PaymentIn, params *Params) (*PaymentIn, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, paymentInList []*PaymentIn, params *Params) (*[]PaymentIn, *resty.Response, error)
	DeleteMany(ctx context.Context, paymentInList []*PaymentIn) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*PaymentIn, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, paymentIn *PaymentIn, params *Params) (*PaymentIn, *resty.Response, error)
	//endpointTemplate[PaymentIn]
	//endpointTemplateBasedOn[PaymentIn, PaymentInTemplateArg]
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
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
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*PaymentIn, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewPaymentInService(client *Client) PaymentInService {
	e := NewEndpoint(client, "entity/paymentin")
	return newMainService[PaymentIn, any, MetadataAttributeSharedStates, any](e)
}
