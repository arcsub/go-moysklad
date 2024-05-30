package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"net/http"
)

// InventoryService
// Сервис для работы с инвентаризациями.
type InventoryService interface {
	GetList(ctx context.Context, params *Params) (*List[Inventory], *resty.Response, error)
	Create(ctx context.Context, inventory *Inventory, params *Params) (*Inventory, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, inventoryList []*Inventory, params *Params) (*[]Inventory, *resty.Response, error)
	DeleteMany(ctx context.Context, inventoryList []*Inventory) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Inventory, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, inventory *Inventory, params *Params) (*Inventory, *resty.Response, error)
	//endpointTemplate[Inventory]
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	GetPositions(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[InventoryPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, params *Params) (*InventoryPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, position *InventoryPosition, params *Params) (*InventoryPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id *uuid.UUID, position *InventoryPosition) (*InventoryPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id *uuid.UUID, positions []*InventoryPosition) (*[]InventoryPosition, *resty.Response, error)
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
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*Inventory, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	Recalculate(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

type inventoryService struct {
	Endpoint
	endpointGetList[Inventory]
	endpointCreate[Inventory]
	endpointCreateUpdateMany[Inventory]
	endpointDeleteMany[Inventory]
	endpointDelete
	endpointGetById[Inventory]
	endpointUpdate[Inventory]
	endpointTemplate[Inventory]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointPositions[InventoryPosition]
	endpointAttributes
	endpointSyncID[Inventory]
	endpointRemove
}

func NewInventoryService(client *Client) InventoryService {
	e := NewEndpoint(client, "entity/inventory")
	return &inventoryService{
		Endpoint:                 e,
		endpointGetList:          endpointGetList[Inventory]{e},
		endpointCreate:           endpointCreate[Inventory]{e},
		endpointCreateUpdateMany: endpointCreateUpdateMany[Inventory]{e},
		endpointDeleteMany:       endpointDeleteMany[Inventory]{e},
		endpointDelete:           endpointDelete{e},
		endpointGetById:          endpointGetById[Inventory]{e},
		endpointUpdate:           endpointUpdate[Inventory]{e},
		endpointTemplate:         endpointTemplate[Inventory]{e},
		endpointMetadata:         endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointPositions:        endpointPositions[InventoryPosition]{e},
		endpointAttributes:       endpointAttributes{e},
		endpointSyncID:           endpointSyncID[Inventory]{e},
		endpointRemove:           endpointRemove{e},
	}
}

// Recalculate Запрос на пересчёт расчётных остатков у позиций инвентаризации.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-obschie-swedeniq-pereschet-raschetnogo-ostatka-w-inwentarizacii
func (s *inventoryService) Recalculate(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("rpc/inventory/%s/recalcCalculatedQuantity", id)
	_, resp, err := NewRequestBuilder[any](s.client, path).Put(ctx, nil)
	if err != nil {
		return false, resp, err
	}
	return resp.StatusCode() == http.StatusCreated, resp, nil
}
