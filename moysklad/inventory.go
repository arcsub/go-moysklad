package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"net/http"
)

// Inventory Инвентаризация.
// Ключевое слово: inventory
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-inwentarizaciq
type Inventory struct {
	AccountID    *uuid.UUID                    `json:"accountId,omitempty"`    // ID учетной записи
	Attributes   *Attributes                   `json:"attributes,omitempty"`   // Коллекция метаданных доп. полей. Поля объекта
	Code         *string                       `json:"code,omitempty"`         // Код выданного
	Created      *Timestamp                    `json:"created,omitempty"`      // Дата создания
	Deleted      *Timestamp                    `json:"deleted,omitempty"`      // Момент последнего удаления Счета-фактуры полученного
	Description  *string                       `json:"description,omitempty"`  // Комментарий выданного Счета-фактуры полученного
	ExternalCode *string                       `json:"externalCode,omitempty"` // Внешний код выданного Счета-фактуры полученного
	Files        *Files                        `json:"files,omitempty"`        // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group        *Group                        `json:"group,omitempty"`        // Отдел сотрудника
	ID           *uuid.UUID                    `json:"id,omitempty"`           // ID сущности
	Meta         *Meta                         `json:"meta,omitempty"`         // Метаданные
	Moment       *Timestamp                    `json:"moment,omitempty"`       // Дата документа
	Name         *string                       `json:"name,omitempty"`         // Наименование
	Organization *Organization                 `json:"organization,omitempty"` // Метаданные юрлица
	Owner        *Employee                     `json:"owner,omitempty"`        // Владелец (Сотрудник)
	Positions    *Positions[InventoryPosition] `json:"positions,omitempty"`    // Метаданные позиций Инвентаризации
	Printed      *bool                         `json:"printed,omitempty"`      // Напечатан ли документ
	Published    *bool                         `json:"published,omitempty"`    // Опубликован ли документ
	Shared       *bool                         `json:"shared,omitempty"`       // Общий доступ
	State        *State                        `json:"state,omitempty"`        // Метаданные статуса
	Store        *Store                        `json:"store,omitempty"`        // Метаданные склада
	Sum          *Decimal                      `json:"sum,omitempty"`          // Сумма
	SyncID       *uuid.UUID                    `json:"syncId,omitempty"`       // ID синхронизации. После заполнения недоступен для изменения
	Updated      *Timestamp                    `json:"updated,omitempty"`      // Момент последнего обновления
}

func (i Inventory) String() string {
	return Stringify(i)
}

func (i Inventory) MetaType() MetaType {
	return MetaTypeInventory
}

// InventoryPosition Позиция Инвентаризации.
// Ключевое слово: inventoryposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-inwentarizaciq-inwentarizaciq-pozicii-inwentarizacii
type InventoryPosition struct {
	AccountID          *uuid.UUID          `json:"accountId,omitempty"`          // ID учетной записи
	Assortment         *AssortmentPosition `json:"assortment,omitempty"`         // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	CalculatedQuantity *float64            `json:"calculatedQuantity,omitempty"` // расчетный остаток
	CorrectionAmount   *float64            `json:"correctionAmount,omitempty"`   // разница между расчетным остатком и фактическимх
	CorrectionSum      *Decimal            `json:"correctionSum,omitempty"`      // избыток/недостача
	ID                 *uuid.UUID          `json:"id,omitempty"`                 // ID сущности
	Pack               *Pack               `json:"pack,omitempty"`               // Упаковка Товара
	Price              *Decimal            `json:"price,omitempty"`              // Цена товара/услуги в копейках
	Quantity           *float64            `json:"quantity,omitempty"`           // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
}

func (i InventoryPosition) String() string {
	return Stringify(i)
}

func (i InventoryPosition) MetaType() MetaType {
	return MetaTypeInventoryPosition
}

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
