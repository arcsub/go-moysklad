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
	Name         *string                       `json:"name,omitempty"`
	Sum          *float64                      `json:"sum,omitempty"`
	Code         *string                       `json:"code,omitempty"`
	Created      *Timestamp                    `json:"created,omitempty"`
	Deleted      *Timestamp                    `json:"deleted,omitempty"`
	Description  *string                       `json:"description,omitempty"`
	ExternalCode *string                       `json:"externalCode,omitempty"`
	Files        *MetaArray[File]              `json:"files,omitempty"`
	Group        *Group                        `json:"group,omitempty"`
	ID           *uuid.UUID                    `json:"id,omitempty"`
	Updated      *Timestamp                    `json:"updated,omitempty"`
	Meta         *Meta                         `json:"meta,omitempty"`
	Owner        *Employee                     `json:"owner,omitempty"`
	Organization *Organization                 `json:"organization,omitempty"`
	AccountID    *uuid.UUID                    `json:"accountId,omitempty"`
	Positions    *Positions[InventoryPosition] `json:"positions,omitempty"`
	Printed      *bool                         `json:"printed,omitempty"`
	Published    *bool                         `json:"published,omitempty"`
	Shared       *bool                         `json:"shared,omitempty"`
	State        *NullValue[State]             `json:"state,omitempty"`
	Store        *Store                        `json:"store,omitempty"`
	Moment       *Timestamp                    `json:"moment,omitempty"`
	SyncID       *uuid.UUID                    `json:"syncId,omitempty"`
	Attributes   Slice[AttributeValue]         `json:"attributes,omitempty"`
}

// Clean возвращает сущность с единственным заполненным полем Meta
func (inventory Inventory) Clean() *Inventory {
	return &Inventory{Meta: inventory.Meta}
}

// AsTaskOperation реализует интерфейс AsTaskOperationInterface
func (inventory Inventory) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: inventory.Meta}
}

func (inventory Inventory) GetName() string {
	return Deref(inventory.Name)
}

func (inventory Inventory) GetSum() float64 {
	return Deref(inventory.Sum)
}

func (inventory Inventory) GetCode() string {
	return Deref(inventory.Code)
}

func (inventory Inventory) GetCreated() Timestamp {
	return Deref(inventory.Created)
}

func (inventory Inventory) GetDeleted() Timestamp {
	return Deref(inventory.Deleted)
}

func (inventory Inventory) GetDescription() string {
	return Deref(inventory.Description)
}

func (inventory Inventory) GetExternalCode() string {
	return Deref(inventory.ExternalCode)
}

func (inventory Inventory) GetFiles() MetaArray[File] {
	return Deref(inventory.Files)
}

func (inventory Inventory) GetGroup() Group {
	return Deref(inventory.Group)
}

func (inventory Inventory) GetID() uuid.UUID {
	return Deref(inventory.ID)
}

func (inventory Inventory) GetUpdated() Timestamp {
	return Deref(inventory.Updated)
}

func (inventory Inventory) GetMeta() Meta {
	return Deref(inventory.Meta)
}

func (inventory Inventory) GetOwner() Employee {
	return Deref(inventory.Owner)
}

func (inventory Inventory) GetOrganization() Organization {
	return Deref(inventory.Organization)
}

func (inventory Inventory) GetAccountID() uuid.UUID {
	return Deref(inventory.AccountID)
}

func (inventory Inventory) GetPositions() Positions[InventoryPosition] {
	return Deref(inventory.Positions)
}

func (inventory Inventory) GetPrinted() bool {
	return Deref(inventory.Printed)
}

func (inventory Inventory) GetPublished() bool {
	return Deref(inventory.Published)
}

func (inventory Inventory) GetShared() bool {
	return Deref(inventory.Shared)
}

func (inventory Inventory) GetState() State {
	return inventory.State.Get()
}

func (inventory Inventory) GetStore() Store {
	return Deref(inventory.Store)
}

func (inventory Inventory) GetMoment() Timestamp {
	return Deref(inventory.Moment)
}

func (inventory Inventory) GetSyncID() uuid.UUID {
	return Deref(inventory.SyncID)
}

func (inventory Inventory) GetAttributes() Slice[AttributeValue] {
	return inventory.Attributes
}

func (inventory *Inventory) SetName(name string) *Inventory {
	inventory.Name = &name
	return inventory
}

func (inventory *Inventory) SetCode(code string) *Inventory {
	inventory.Code = &code
	return inventory
}

func (inventory *Inventory) SetDescription(description string) *Inventory {
	inventory.Description = &description
	return inventory
}

func (inventory *Inventory) SetExternalCode(externalCode string) *Inventory {
	inventory.ExternalCode = &externalCode
	return inventory
}

func (inventory *Inventory) SetFiles(files Slice[File]) *Inventory {
	inventory.Files = NewMetaArrayRows(files)
	return inventory
}

func (inventory *Inventory) SetGroup(group *Group) *Inventory {
	inventory.Group = group.Clean()
	return inventory
}

func (inventory *Inventory) SetMeta(meta *Meta) *Inventory {
	inventory.Meta = meta
	return inventory
}

func (inventory *Inventory) SetOwner(owner *Employee) *Inventory {
	inventory.Owner = owner.Clean()
	return inventory
}

func (inventory *Inventory) SetOrganization(organization *Organization) *Inventory {
	inventory.Organization = organization.Clean()
	return inventory
}

func (inventory *Inventory) SetPositions(positions *Positions[InventoryPosition]) *Inventory {
	inventory.Positions = positions
	return inventory
}

func (inventory *Inventory) SetShared(shared bool) *Inventory {
	inventory.Shared = &shared
	return inventory
}

func (inventory *Inventory) SetState(state *State) *Inventory {
	inventory.State = NewNullValueWith(state.Clean())
	return inventory
}

func (inventory *Inventory) SetNullState() *Inventory {
	inventory.State = NewNullValue[State]()
	return inventory
}

func (inventory *Inventory) SetStore(store *Store) *Inventory {
	inventory.Store = store.Clean()
	return inventory
}

func (inventory *Inventory) SetMoment(moment *Timestamp) *Inventory {
	inventory.Moment = moment
	return inventory
}

func (inventory *Inventory) SetSyncID(syncID uuid.UUID) *Inventory {
	inventory.SyncID = &syncID
	return inventory
}

func (inventory *Inventory) SetAttributes(attributes Slice[AttributeValue]) *Inventory {
	inventory.Attributes = attributes
	return inventory
}

func (inventory Inventory) String() string {
	return Stringify(inventory)
}

func (inventory Inventory) MetaType() MetaType {
	return MetaTypeInventory
}

// Update shortcut
func (inventory Inventory) Update(ctx context.Context, client *Client, params ...*Params) (*Inventory, *resty.Response, error) {
	return client.Entity().Inventory().Update(ctx, inventory.GetID(), &inventory, params...)
}

// Create shortcut
func (inventory Inventory) Create(ctx context.Context, client *Client, params ...*Params) (*Inventory, *resty.Response, error) {
	return client.Entity().Inventory().Create(ctx, &inventory, params...)
}

// Delete shortcut
func (inventory Inventory) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return client.Entity().Inventory().Delete(ctx, inventory.GetID())
}

// InventoryPosition Позиция Инвентаризации.
// Ключевое слово: inventoryposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-inwentarizaciq-inwentarizaciq-pozicii-inwentarizacii
type InventoryPosition struct {
	AccountID          *uuid.UUID          `json:"accountId,omitempty"`          // ID учетной записи
	Assortment         *AssortmentPosition `json:"assortment,omitempty"`         // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	CalculatedQuantity *float64            `json:"calculatedQuantity,omitempty"` // расчетный остаток
	CorrectionAmount   *float64            `json:"correctionAmount,omitempty"`   // разница между расчетным остатком и фактическимх
	CorrectionSum      *float64            `json:"correctionSum,omitempty"`      // избыток/недостача
	ID                 *uuid.UUID          `json:"id,omitempty"`                 // ID сущности
	Pack               *Pack               `json:"pack,omitempty"`               // Упаковка Товара
	Price              *float64            `json:"price,omitempty"`              // Цена товара/услуги в копейках
	Quantity           *float64            `json:"quantity,omitempty"`           // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
}

func (inventoryPosition InventoryPosition) GetAccountID() uuid.UUID {
	return Deref(inventoryPosition.AccountID)
}

func (inventoryPosition InventoryPosition) GetAssortment() AssortmentPosition {
	return Deref(inventoryPosition.Assortment)
}

func (inventoryPosition InventoryPosition) GetCalculatedQuantity() float64 {
	return Deref(inventoryPosition.CalculatedQuantity)
}

func (inventoryPosition InventoryPosition) GetCorrectionAmount() float64 {
	return Deref(inventoryPosition.CorrectionAmount)
}

func (inventoryPosition InventoryPosition) GetCorrectionSum() float64 {
	return Deref(inventoryPosition.CorrectionSum)
}

func (inventoryPosition InventoryPosition) GetID() uuid.UUID {
	return Deref(inventoryPosition.ID)
}

func (inventoryPosition InventoryPosition) GetPack() Pack {
	return Deref(inventoryPosition.Pack)
}

func (inventoryPosition InventoryPosition) GetPrice() float64 {
	return Deref(inventoryPosition.Price)
}

func (inventoryPosition InventoryPosition) GetQuantity() float64 {
	return Deref(inventoryPosition.Quantity)
}

func (inventoryPosition *InventoryPosition) SetAssortment(assortment AsAssortment) *InventoryPosition {
	inventoryPosition.Assortment = assortment.AsAssortment()
	return inventoryPosition
}

func (inventoryPosition *InventoryPosition) SetCalculatedQuantity(calculatedQuantity float64) *InventoryPosition {
	inventoryPosition.CalculatedQuantity = &calculatedQuantity
	return inventoryPosition
}

func (inventoryPosition *InventoryPosition) SetPack(pack *Pack) *InventoryPosition {
	inventoryPosition.Pack = pack
	return inventoryPosition
}

func (inventoryPosition *InventoryPosition) SetPrice(price float64) *InventoryPosition {
	inventoryPosition.Price = &price
	return inventoryPosition
}

func (inventoryPosition *InventoryPosition) SetQuantity(quantity float64) *InventoryPosition {
	inventoryPosition.Quantity = &quantity
	return inventoryPosition
}

func (inventoryPosition InventoryPosition) String() string {
	return Stringify(inventoryPosition)
}

func (inventoryPosition InventoryPosition) MetaType() MetaType {
	return MetaTypeInventoryPosition
}

// InventoryService
// Сервис для работы с инвентаризациями.
type InventoryService interface {
	GetList(ctx context.Context, params ...*Params) (*List[Inventory], *resty.Response, error)
	Create(ctx context.Context, inventory *Inventory, params ...*Params) (*Inventory, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, inventoryList Slice[Inventory], params ...*Params) (*Slice[Inventory], *resty.Response, error)
	DeleteMany(ctx context.Context, entities ...Inventory) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Inventory, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, inventory *Inventory, params ...*Params) (*Inventory, *resty.Response, error)
	Template(ctx context.Context) (*Inventory, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetaAttributesSharedStatesWrapper, *resty.Response, error)
	GetPositions(ctx context.Context, id uuid.UUID, params ...*Params) (*MetaArray[InventoryPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params ...*Params) (*InventoryPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *InventoryPosition, params ...*Params) (*InventoryPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id uuid.UUID, position *InventoryPosition) (*InventoryPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id uuid.UUID, positions Slice[InventoryPosition]) (*Slice[InventoryPosition], *resty.Response, error)
	DeletePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (bool, *resty.Response, error)
	GetPositionTrackingCodes(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (*MetaArray[TrackingCode], *resty.Response, error)
	CreateOrUpdatePositionTrackingCodes(ctx context.Context, id uuid.UUID, positionID uuid.UUID, trackingCodes Slice[TrackingCode]) (*Slice[TrackingCode], *resty.Response, error)
	DeletePositionTrackingCodes(ctx context.Context, id uuid.UUID, positionID uuid.UUID, trackingCodes Slice[TrackingCode]) (*DeleteManyResponse, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList Slice[Attribute]) (*Slice[Attribute], *resty.Response, error)
	UpdateAttribute(ctx context.Context, id uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList []MetaWrapper) (*DeleteManyResponse, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*Inventory, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error)
	MoveToTrash(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	Recalculate(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetFiles(ctx context.Context, id uuid.UUID) (*MetaArray[File], *resty.Response, error)
	CreateFile(ctx context.Context, id uuid.UUID, file *File) (*Slice[File], *resty.Response, error)
	UpdateFiles(ctx context.Context, id uuid.UUID, files Slice[File]) (*Slice[File], *resty.Response, error)
	DeleteFile(ctx context.Context, id uuid.UUID, fileID uuid.UUID) (bool, *resty.Response, error)
	DeleteFiles(ctx context.Context, id uuid.UUID, files []MetaWrapper) (*DeleteManyResponse, *resty.Response, error)
	Evaluate(ctx context.Context, entity *Inventory, evaluate ...Evaluate) (*Inventory, *resty.Response, error)
}

type inventoryService struct {
	Endpoint
	endpointGetList[Inventory]
	endpointCreate[Inventory]
	endpointCreateUpdateMany[Inventory]
	endpointDeleteMany[Inventory]
	endpointDelete
	endpointGetByID[Inventory]
	endpointUpdate[Inventory]
	endpointTemplate[Inventory]
	endpointMetadata[MetaAttributesSharedStatesWrapper]
	endpointPositions[InventoryPosition]
	endpointAttributes
	endpointSyncID[Inventory]
	endpointTrash
	endpointStates
	endpointFiles
	endpointEvaluate[Inventory]
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
		endpointGetByID:          endpointGetByID[Inventory]{e},
		endpointUpdate:           endpointUpdate[Inventory]{e},
		endpointTemplate:         endpointTemplate[Inventory]{e},
		endpointMetadata:         endpointMetadata[MetaAttributesSharedStatesWrapper]{e},
		endpointPositions:        endpointPositions[InventoryPosition]{e},
		endpointAttributes:       endpointAttributes{e},
		endpointSyncID:           endpointSyncID[Inventory]{e},
		endpointTrash:            endpointTrash{e},
		endpointStates:           endpointStates{e},
		endpointFiles:            endpointFiles{e},
		endpointEvaluate:         endpointEvaluate[Inventory]{e},
	}
}

// Recalculate Запрос на пересчёт расчётных остатков у позиций инвентаризации.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-obschie-swedeniq-pereschet-raschetnogo-ostatka-w-inwentarizacii
func (service *inventoryService) Recalculate(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("rpc/inventory/%s/recalcCalculatedQuantity", id)
	_, resp, err := NewRequestBuilder[any](service.client, path).Put(ctx, nil)
	if err != nil {
		return false, resp, err
	}
	return resp.StatusCode() == http.StatusCreated, resp, nil
}
