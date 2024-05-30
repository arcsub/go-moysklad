package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// SupplyService
// Сервис для работы с приёмками.
type SupplyService interface {
	GetList(ctx context.Context, params *Params) (*List[Supply], *resty.Response, error)
	Create(ctx context.Context, supply *Supply, params *Params) (*Supply, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, supplyList []*Supply, params *Params) (*[]Supply, *resty.Response, error)
	DeleteMany(ctx context.Context, supplyList []*Supply) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Supply, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, supply *Supply, params *Params) (*Supply, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	//endpointTemplate[Supply]
	GetPositions(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[SupplyPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, params *Params) (*SupplyPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, position *SupplyPosition, params *Params) (*SupplyPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id *uuid.UUID, position *SupplyPosition) (*SupplyPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id *uuid.UUID, positions []*SupplyPosition) (*[]SupplyPosition, *resty.Response, error)
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
	PrintDocument(ctx context.Context, id *uuid.UUID, PrintDocumentArg *PrintDocumentArg) (*PrintFile, *resty.Response, error)
	GetFiles(ctx context.Context, id *uuid.UUID) (*MetaArray[File], *resty.Response, error)
	CreateFile(ctx context.Context, id *uuid.UUID, file *File) (*[]File, *resty.Response, error)
	UpdateFiles(ctx context.Context, id *uuid.UUID, files []*File) (*[]File, *resty.Response, error)
	DeleteFile(ctx context.Context, id *uuid.UUID, fileId *uuid.UUID) (bool, *resty.Response, error)
	DeleteFiles(ctx context.Context, id *uuid.UUID, files []*File) (*DeleteManyResponse, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
	GetStateByID(ctx context.Context, id *uuid.UUID) (*State, *resty.Response, error)
	CreateState(ctx context.Context, state *State) (*State, *resty.Response, error)
	UpdateState(ctx context.Context, id *uuid.UUID, state *State) (*State, *resty.Response, error)
	CreateOrUpdateStates(ctx context.Context, id *uuid.UUID, states []*State) (*[]State, *resty.Response, error)
	DeleteState(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*Supply, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewSupplyService(client *Client) SupplyService {
	e := NewEndpoint(client, "entity/supply")
	return newMainService[Supply, SupplyPosition, MetadataAttributeSharedStates, any](e)
}
