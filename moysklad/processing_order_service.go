package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// ProcessingOrderService
// Сервис для работы с заказами на производство.
type ProcessingOrderService interface {
	GetList(ctx context.Context, params *Params) (*List[ProcessingOrder], *resty.Response, error)
	Create(ctx context.Context, processingOrder *ProcessingOrder, params *Params) (*ProcessingOrder, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, processingOrderList []*ProcessingOrder, params *Params) (*[]ProcessingOrder, *resty.Response, error)
	DeleteMany(ctx context.Context, processingOrderList []*ProcessingOrder) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*ProcessingOrder, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, processingOrder *ProcessingOrder, params *Params) (*ProcessingOrder, *resty.Response, error)
	//endpointTemplate[ProcessingOrder]
	//endpointTemplateBasedOn[ProcessingOrder, ProcessingOrderTemplateArg]
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	GetPositions(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[ProcessingOrderPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, params *Params) (*ProcessingOrderPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, position *ProcessingOrderPosition, params *Params) (*ProcessingOrderPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id *uuid.UUID, position *ProcessingOrderPosition) (*ProcessingOrderPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id *uuid.UUID, positions []*ProcessingOrderPosition) (*[]ProcessingOrderPosition, *resty.Response, error)
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
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*ProcessingOrder, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewProcessingOrderService(client *Client) ProcessingOrderService {
	e := NewEndpoint(client, "entity/processingorder")
	return newMainService[ProcessingOrder, ProcessingOrderPosition, MetadataAttributeSharedStates, any](e)
}
