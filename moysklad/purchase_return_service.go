package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// PurchaseReturnService
// Сервис для работы с возвратами поставщикам.
type PurchaseReturnService interface {
	GetList(ctx context.Context, params *Params) (*List[PurchaseReturn], *resty.Response, error)
	Create(ctx context.Context, purchaseReturn *PurchaseReturn, params *Params) (*PurchaseReturn, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, purchaseReturnList []*PurchaseReturn, params *Params) (*[]PurchaseReturn, *resty.Response, error)
	DeleteMany(ctx context.Context, purchaseReturnList []*PurchaseReturn) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*PurchaseReturn, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, purchaseReturn *PurchaseReturn, params *Params) (*PurchaseReturn, *resty.Response, error)
	//endpointTemplate[PurchaseReturn]
	//endpointTemplateBasedOn[PurchaseReturn, PurchaseReturnTemplateArg]
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	GetPositions(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[PurchaseReturnPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, params *Params) (*PurchaseReturnPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, position *PurchaseReturnPosition, params *Params) (*PurchaseReturnPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id *uuid.UUID, position *PurchaseReturnPosition) (*PurchaseReturnPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id *uuid.UUID, positions []*PurchaseReturnPosition) (*[]PurchaseReturnPosition, *resty.Response, error)
	DeletePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID) (bool, *resty.Response, error)
	GetPositionTrackingCodes(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID) (*MetaArray[TrackingCode], *resty.Response, error)
	CreateOrUpdatePositionTrackingCodes(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, trackingCodes TrackingCodes) (*[]TrackingCode, *resty.Response, error)
	DeletePositionTrackingCodes(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, trackingCodes TrackingCodes) (*DeleteManyResponse, *resty.Response, error)
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
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*PurchaseReturn, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
	Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewPurchaseReturnService(client *Client) PurchaseReturnService {
	e := NewEndpoint(client, "entity/purchasereturn")
	return newMainService[PurchaseReturn, PurchaseReturnPosition, MetadataAttributeSharedStates, any](e)
}
