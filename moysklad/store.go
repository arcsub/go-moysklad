package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Store Склад.
// Ключевое слово: store
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad
type Store struct {
	Owner        *Employee        `json:"owner,omitempty"`
	Slots        *MetaArray[Slot] `json:"slots,omitempty"`
	Group        *Group           `json:"group,omitempty"`
	Archived     *bool            `json:"archived,omitempty"`
	ID           *uuid.UUID       `json:"id,omitempty"`
	Code         *string          `json:"code,omitempty"`
	Description  *string          `json:"description,omitempty"`
	ExternalCode *string          `json:"externalCode,omitempty"`
	AddressFull  *Address         `json:"addressFull,omitempty"`
	Address      *string          `json:"address,omitempty"`
	Name         *string          `json:"name,omitempty"`
	Meta         *Meta            `json:"meta,omitempty"`
	AccountID    *uuid.UUID       `json:"accountId,omitempty"`
	Parent       *Store           `json:"parent,omitempty"`
	PathName     *string          `json:"pathName,omitempty"`
	Shared       *bool            `json:"shared,omitempty"`
	Updated      *Timestamp       `json:"updated,omitempty"`
	Zones        *MetaArray[Zone] `json:"zones,omitempty"`
	Attributes   Slice[Attribute] `json:"attributes,omitempty"`
}

// Clean возвращает сущность с единственным заполненным полем Meta
func (store Store) Clean() *Store {
	return &Store{Meta: store.Meta}
}

func (store Store) GetOwner() Employee {
	return Deref(store.Owner)
}

func (store Store) GetSlots() MetaArray[Slot] {
	return Deref(store.Slots)
}

func (store Store) GetGroup() Group {
	return Deref(store.Group)
}

func (store Store) GetArchived() bool {
	return Deref(store.Archived)
}

func (store Store) GetID() uuid.UUID {
	return Deref(store.ID)
}

func (store Store) GetCode() string {
	return Deref(store.Code)
}

func (store Store) GetDescription() string {
	return Deref(store.Description)
}

func (store Store) GetExternalCode() string {
	return Deref(store.ExternalCode)
}

func (store Store) GetAddressFull() Address {
	return Deref(store.AddressFull)
}

func (store Store) GetAddress() string {
	return Deref(store.Address)
}

func (store Store) GetName() string {
	return Deref(store.Name)
}

func (store Store) GetMeta() Meta {
	return Deref(store.Meta)
}

func (store Store) GetAccountID() uuid.UUID {
	return Deref(store.AccountID)
}

func (store Store) GetParent() Store {
	return Deref(store.Parent)
}

func (store Store) GetPathName() string {
	return Deref(store.PathName)
}

func (store Store) GetShared() bool {
	return Deref(store.Shared)
}

func (store Store) GetUpdated() Timestamp {
	return Deref(store.Updated)
}

func (store Store) GetZones() MetaArray[Zone] {
	return Deref(store.Zones)
}

func (store Store) GetAttributes() Slice[Attribute] {
	return store.Attributes
}

func (store *Store) SetOwner(owner *Employee) *Store {
	store.Owner = owner
	return store
}

func (store *Store) SetGroup(group *Group) *Store {
	store.Group = group
	return store
}

func (store *Store) SetArchived(archived bool) *Store {
	store.Archived = &archived
	return store
}

func (store *Store) SetCode(code string) *Store {
	store.Code = &code
	return store
}

func (store *Store) SetDescription(description string) *Store {
	store.Description = &description
	return store
}

func (store *Store) SetExternalCode(externalCode string) *Store {
	store.ExternalCode = &externalCode
	return store
}

func (store *Store) SetAddressFull(addressFull *Address) *Store {
	store.AddressFull = addressFull
	return store
}

func (store *Store) SetAddress(address string) *Store {
	store.Address = &address
	return store
}

func (store *Store) SetName(name string) *Store {
	store.Name = &name
	return store
}

func (store *Store) SetMeta(meta *Meta) *Store {
	store.Meta = meta
	return store
}

func (store *Store) SetParent(parent *Store) *Store {
	store.Parent = parent
	return store
}

func (store *Store) SetPathName(pathName string) *Store {
	store.PathName = &pathName
	return store
}

func (store *Store) SetShared(shared bool) *Store {
	store.Shared = &shared
	return store
}

func (store *Store) SetAttributes(attributes ...*Attribute) *Store {
	store.Attributes = attributes
	return store
}

func (store Store) String() string {
	return Stringify(store)
}

func (store Store) MetaType() MetaType {
	return MetaTypeStore
}

// Update shortcut
func (store Store) Update(ctx context.Context, client *Client, params ...*Params) (*Store, *resty.Response, error) {
	return client.Entity().Store().Update(ctx, store.GetID(), &store, params...)
}

// Create shortcut
func (store Store) Create(ctx context.Context, client *Client, params ...*Params) (*Store, *resty.Response, error) {
	return client.Entity().Store().Create(ctx, &store, params...)
}

// Delete shortcut
func (store Store) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return client.Entity().Store().Delete(ctx, store.GetID())
}

// Slot Ячейка склада.
// Ключевое слово: slot
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-yachejki-sklada
type Slot struct {
	AccountID    *uuid.UUID `json:"accountId,omitempty"`    // ID учетной записи
	Barcode      *string    `json:"barcode,omitempty"`      // Штрихкод ячейки [16-10-2023]
	ExternalCode *string    `json:"externalCode,omitempty"` // Внешний код Ячейки
	ID           *uuid.UUID `json:"id,omitempty"`           // ID Ячейки
	Meta         *Meta      `json:"meta,omitempty"`         // Метаданные Ячейки
	Name         *string    `json:"name,omitempty"`         // Наименование Ячейки
	Updated      *Timestamp `json:"updated,omitempty"`      // Момент последнего обновления Ячейки
	Zone         *Zone      `json:"zone,omitempty"`         // Зона ячейки
}

// Clean возвращает сущность с единственным заполненным полем Meta
func (slot Slot) Clean() *Slot {
	return &Slot{Meta: slot.Meta}
}

func (slot Slot) GetAccountID() uuid.UUID {
	return Deref(slot.AccountID)
}

func (slot Slot) GetBarcode() string {
	return Deref(slot.Barcode)
}

func (slot Slot) GetExternalCode() string {
	return Deref(slot.ExternalCode)
}

func (slot Slot) GetID() uuid.UUID {
	return Deref(slot.ID)
}

func (slot Slot) GetMeta() Meta {
	return Deref(slot.Meta)
}

func (slot Slot) GetName() string {
	return Deref(slot.Name)
}

func (slot Slot) GetUpdated() Timestamp {
	return Deref(slot.Updated)
}

func (slot Slot) GetZone() Zone {
	return Deref(slot.Zone)
}

func (slot *Slot) SetBarcode(barcode string) *Slot {
	slot.Barcode = &barcode
	return slot
}

func (slot *Slot) SetExternalCode(externalCode string) *Slot {
	slot.ExternalCode = &externalCode
	return slot
}

func (slot *Slot) SetMeta(meta *Meta) *Slot {
	slot.Meta = meta
	return slot
}

func (slot *Slot) SetName(name string) *Slot {
	slot.Name = &name
	return slot
}

func (slot Slot) String() string {
	return Stringify(slot)
}

func (slot Slot) MetaType() MetaType {
	return MetaTypeSlot
}

// Zone Зона склада.
// Ключевое слово: storezone
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-zony-sklada
type Zone struct {
	AccountID    *uuid.UUID `json:"accountId,omitempty"`    // ID учетной записи
	ExternalCode *string    `json:"externalCode,omitempty"` // Внешний код Зоны
	ID           *uuid.UUID `json:"id,omitempty"`           // ID Зоны
	Meta         *Meta      `json:"meta,omitempty"`         // Метаданные Зоны
	Name         *string    `json:"name,omitempty"`         // Наименование Зоны
	Updated      *Timestamp `json:"updated,omitempty"`      // Момент последнего обновления Зоны
}

func (zone Zone) GetAccountID() uuid.UUID {
	return Deref(zone.AccountID)
}

func (zone Zone) GetExternalCode() string {
	return Deref(zone.ExternalCode)
}

func (zone Zone) GetID() uuid.UUID {
	return Deref(zone.ID)
}

func (zone Zone) GetMeta() Meta {
	return Deref(zone.Meta)
}

func (zone Zone) GetName() string {
	return Deref(zone.Name)
}

func (zone Zone) GetUpdated() Timestamp {
	return Deref(zone.Updated)
}

func (zone *Zone) SetExternalCode(externalCode string) *Zone {
	zone.ExternalCode = &externalCode
	return zone
}

func (zone *Zone) SetMeta(meta *Meta) *Zone {
	zone.Meta = meta
	return zone
}

func (zone *Zone) SetName(name string) *Zone {
	zone.Name = &name
	return zone
}

func (zone Zone) String() string {
	return Stringify(zone)
}

func (zone Zone) MetaType() MetaType {
	return MetaTypeStoreZone
}

// StoreService
// Сервис для работы со складами.
type StoreService interface {
	GetList(ctx context.Context, params ...*Params) (*List[Store], *resty.Response, error)
	Create(ctx context.Context, store *Store, params ...*Params) (*Store, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, storeList Slice[Store], params ...*Params) (*Slice[Store], *resty.Response, error)
	DeleteMany(ctx context.Context, entities ...*Store) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetaAttributesSharedWrapper, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributeMany(ctx context.Context, attributes ...*Attribute) (*Slice[Attribute], *resty.Response, error)
	UpdateAttribute(ctx context.Context, id uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributeMany(ctx context.Context, attributes ...*Attribute) (*DeleteManyResponse, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Store, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, store *Store, params ...*Params) (*Store, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params ...*Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)
	GetSlots(ctx context.Context, storeID uuid.UUID) (*List[Slot], *resty.Response, error)
	CreateSlot(ctx context.Context, storeID uuid.UUID, slot *Slot) (*Slot, *resty.Response, error)
	CreateUpdateSlotMany(ctx context.Context, storeID uuid.UUID, slots ...*Slot) (*Slice[Slot], *resty.Response, error)
	DeleteSlotMany(ctx context.Context, storeID uuid.UUID, slots ...*Slot) (*DeleteManyResponse, *resty.Response, error)
	GetSlotByID(ctx context.Context, storeID, slotID uuid.UUID) (*Slot, *resty.Response, error)
	DeleteSlot(ctx context.Context, storeID, slotID uuid.UUID) (bool, *resty.Response, error)
	UpdateSlot(ctx context.Context, storeID, slotID uuid.UUID, slot *Slot) (*Slot, *resty.Response, error)
	GetZones(ctx context.Context, storeID uuid.UUID) (*List[Zone], *resty.Response, error)
	CreateZone(ctx context.Context, storeID uuid.UUID, zone *Zone) (*Zone, *resty.Response, error)
	CreateUpdateZoneMany(ctx context.Context, storeID uuid.UUID, zones ...*Zone) (*Slice[Zone], *resty.Response, error)
	DeleteZoneMany(ctx context.Context, storeID uuid.UUID, zones ...*Zone) (*DeleteManyResponse, *resty.Response, error)
	DeleteZone(ctx context.Context, storeID, zoneID uuid.UUID) (bool, *resty.Response, error)
	GetZoneByID(ctx context.Context, storeID, zoneID uuid.UUID) (*Zone, *resty.Response, error)
	UpdateZone(ctx context.Context, storeID, zoneID uuid.UUID, zone *Zone) (*Zone, *resty.Response, error)
}

type storeService struct {
	Endpoint
	endpointGetList[Store]
	endpointCreate[Store]
	endpointCreateUpdateMany[Store]
	endpointDeleteMany[Store]
	endpointDelete
	endpointMetadata[MetaAttributesSharedWrapper]
	endpointAttributes
	endpointGetByID[Store]
	endpointUpdate[Store]
	endpointNamedFilter
}

func NewStoreService(client *Client) StoreService {
	e := NewEndpoint(client, "entity/store")
	return &storeService{
		Endpoint:                 e,
		endpointGetList:          endpointGetList[Store]{e},
		endpointCreate:           endpointCreate[Store]{e},
		endpointCreateUpdateMany: endpointCreateUpdateMany[Store]{e},
		endpointDeleteMany:       endpointDeleteMany[Store]{e},
		endpointDelete:           endpointDelete{e},
		endpointMetadata:         endpointMetadata[MetaAttributesSharedWrapper]{e},
		endpointAttributes:       endpointAttributes{e},
		endpointGetByID:          endpointGetByID[Store]{e},
		endpointUpdate:           endpointUpdate[Store]{e},
		endpointNamedFilter:      endpointNamedFilter{e},
	}
}

// GetSlots Получить список всех Ячеек Склада.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-poluchit-qchejki-sklada
func (service *storeService) GetSlots(ctx context.Context, storeID uuid.UUID) (*List[Slot], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/slots", service.uri, storeID)
	return NewRequestBuilder[List[Slot]](service.client, path).Get(ctx)
}

// CreateSlot Запрос на создание Ячейки Склада.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-sozdat-qchejku-sklada
func (service *storeService) CreateSlot(ctx context.Context, storeID uuid.UUID, slot *Slot) (*Slot, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/slots", service.uri, storeID)
	return NewRequestBuilder[Slot](service.client, path).Post(ctx, slot)
}

// CreateUpdateSlotMany Запрос создания и обновления нескольких Ячеек Склада.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-massowoe-sozdanie-i-obnowlenie-qcheek-sklada
func (service *storeService) CreateUpdateSlotMany(ctx context.Context, storeID uuid.UUID, slots ...*Slot) (*Slice[Slot], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/slots", service.uri, storeID)
	return NewRequestBuilder[Slice[Slot]](service.client, path).Post(ctx, slots)
}

// DeleteSlotMany Запрос на массовое удаление Ячеек склада.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-massowoe-udalenie-qcheek-sklada
func (service *storeService) DeleteSlotMany(ctx context.Context, storeID uuid.UUID, slots ...*Slot) (*DeleteManyResponse, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/slots/delete", service.uri, storeID)
	return NewRequestBuilder[DeleteManyResponse](service.client, path).Post(ctx, AsMetaWrapperSlice(slots))
}

// GetSlotByID Запрос на получение отдельной Ячейки Склада с указанным id.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-poluchit-qchejku-sklada
func (service *storeService) GetSlotByID(ctx context.Context, storeID, slotID uuid.UUID) (*Slot, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/slots/%s", service.uri, storeID, slotID)
	return NewRequestBuilder[Slot](service.client, path).Get(ctx)
}

// DeleteSlot Запрос на удаление Ячейки склада с указанным id.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-udalit-qchejku-sklada
func (service *storeService) DeleteSlot(ctx context.Context, storeID, slotID uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/slots/%s", service.uri, storeID, slotID)
	return NewRequestBuilder[any](service.client, path).Delete(ctx)
}

// UpdateSlot Запрос на обновление Ячейки склада.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-izmenit-qchejku-sklada
func (service *storeService) UpdateSlot(ctx context.Context, storeID, slotID uuid.UUID, slot *Slot) (*Slot, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/slots/%s", service.uri, storeID, slotID)
	return NewRequestBuilder[Slot](service.client, path).Put(ctx, slot)
}

// GetZones Получить список всех Зон.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-poluchit-zony-sklada
func (service *storeService) GetZones(ctx context.Context, storeID uuid.UUID) (*List[Zone], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/zones", service.uri, storeID)
	return NewRequestBuilder[List[Zone]](service.client, path).Get(ctx)
}

// CreateZone Запрос на создание Зоны склада.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-sozdat-zonu-sklada
func (service *storeService) CreateZone(ctx context.Context, storeID uuid.UUID, zone *Zone) (*Zone, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/zones", service.uri, storeID)
	return NewRequestBuilder[Zone](service.client, path).Post(ctx, zone)
}

// CreateUpdateZoneMany Запрос на создание и обновление нескольких Зон склада.
func (service *storeService) CreateUpdateZoneMany(ctx context.Context, storeID uuid.UUID, zones ...*Zone) (*Slice[Zone], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/zones", service.uri, storeID)
	return NewRequestBuilder[Slice[Zone]](service.client, path).Post(ctx, zones)
}

// DeleteZoneMany Запрос на массовое удаление Зон склада.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-massowoe-udalenie-zon-sklada
func (service *storeService) DeleteZoneMany(ctx context.Context, storeID uuid.UUID, zones ...*Zone) (*DeleteManyResponse, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/zones/delete", service.uri, storeID)
	return NewRequestBuilder[DeleteManyResponse](service.client, path).Post(ctx, AsMetaWrapperSlice(zones))
}

// DeleteZone Запрос на удаление Зоны склада с указанным id.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-udalit-zonu-sklada
func (service *storeService) DeleteZone(ctx context.Context, storeID, zoneID uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/zones/%s", service.uri, storeID, zoneID)
	return NewRequestBuilder[any](service.client, path).Delete(ctx)
}

// GetZoneByID Запрос на получение отдельной Зоны Склада с указанным id.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-poluchit-zonu-sklada
func (service *storeService) GetZoneByID(ctx context.Context, storeID, zoneID uuid.UUID) (*Zone, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/zones/%s", service.uri, storeID, zoneID)
	return NewRequestBuilder[Zone](service.client, path).Get(ctx)
}

// UpdateZone Запрос на обновление Зоны склада.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-izmenit-zonu-sklada
func (service *storeService) UpdateZone(ctx context.Context, storeID, zoneID uuid.UUID, zone *Zone) (*Zone, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/zones/%s", service.uri, storeID, zoneID)
	return NewRequestBuilder[Zone](service.client, path).Put(ctx, zone)
}
