package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"time"
)

// Store Склад.
//
// Код сущности: store
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad
type Store struct {
	Owner        *Employee        `json:"owner,omitempty"`        // Метаданные владельца (Сотрудника)
	Slots        *MetaArray[Slot] `json:"slots,omitempty"`        // Ячейки склада
	Group        *Group           `json:"group,omitempty"`        // Отдел сотрудника
	Archived     *bool            `json:"archived,omitempty"`     // Добавлен ли Склад в архив
	ID           *uuid.UUID       `json:"id,omitempty"`           // ID Склада
	Code         *string          `json:"code,omitempty"`         // Код Склада
	Description  *string          `json:"description,omitempty"`  // Комментарий к Складу
	ExternalCode *string          `json:"externalCode,omitempty"` // Внешний код Склада
	AddressFull  *Address         `json:"addressFull,omitempty"`  // Адрес с детализацией по отдельным полям
	Address      *string          `json:"address,omitempty"`      // Адрес склада
	Name         *string          `json:"name,omitempty"`         // Наименование Склада
	Meta         *Meta            `json:"meta,omitempty"`         // Метаданные Склада
	AccountID    *uuid.UUID       `json:"accountId,omitempty"`    // ID учётной записи
	Parent       *Store           `json:"parent,omitempty"`       // Метаданные родительского склада (Группы)
	PathName     *string          `json:"pathName,omitempty"`     // Группа Склада
	Shared       *bool            `json:"shared,omitempty"`       // Общий доступ
	Updated      *Timestamp       `json:"updated,omitempty"`      // Момент последнего обновления Склада
	Zones        *MetaArray[Zone] `json:"zones,omitempty"`        // Зоны склада
	Attributes   Slice[Attribute] `json:"attributes,omitempty"`   // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (store Store) Clean() *Store {
	if store.Meta == nil {
		return nil
	}
	return &Store{Meta: store.Meta}
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (store Store) GetOwner() Employee {
	return Deref(store.Owner)
}

// GetSlots возвращает Ячейки склада.
func (store Store) GetSlots() MetaArray[Slot] {
	return Deref(store.Slots)
}

// GetGroup возвращает Отдел сотрудника.
func (store Store) GetGroup() Group {
	return Deref(store.Group)
}

// GetArchived возвращает флаг нахождения в архиве.
func (store Store) GetArchived() bool {
	return Deref(store.Archived)
}

// GetID возвращает ID Склада.
func (store Store) GetID() uuid.UUID {
	return Deref(store.ID)
}

// GetCode возвращает Код Склада.
func (store Store) GetCode() string {
	return Deref(store.Code)
}

// GetDescription возвращает Комментарий к Складу.
func (store Store) GetDescription() string {
	return Deref(store.Description)
}

// GetExternalCode возвращает Внешний код Склада.
func (store Store) GetExternalCode() string {
	return Deref(store.ExternalCode)
}

// GetAddressFull возвращает Адрес склада с детализацией по отдельным полям.
func (store Store) GetAddressFull() Address {
	return Deref(store.AddressFull)
}

// GetAddress возвращает Адрес склада.
func (store Store) GetAddress() string {
	return Deref(store.Address)
}

// GetName возвращает Наименование склада.
func (store Store) GetName() string {
	return Deref(store.Name)
}

// GetMeta возвращает Метаданные склада.
func (store Store) GetMeta() Meta {
	return Deref(store.Meta)
}

// GetAccountID возвращает ID учётной записи.
func (store Store) GetAccountID() uuid.UUID {
	return Deref(store.AccountID)
}

// GetParent возвращает Метаданные родительского склада (Группы).
func (store Store) GetParent() Store {
	return Deref(store.Parent)
}

// GetPathName возвращает Группу Склада.
func (store Store) GetPathName() string {
	return Deref(store.PathName)
}

// GetShared возвращает флаг Общего доступа.
func (store Store) GetShared() bool {
	return Deref(store.Shared)
}

// GetUpdated возвращает Момент последнего обновления Склада.
func (store Store) GetUpdated() time.Time {
	return Deref(store.Updated).Time()
}

// GetZones возвращает Зоны склада.
func (store Store) GetZones() MetaArray[Zone] {
	return Deref(store.Zones)
}

// GetAttributes возвращает Список метаданных доп. полей.
func (store Store) GetAttributes() Slice[Attribute] {
	return store.Attributes
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (store *Store) SetOwner(owner *Employee) *Store {
	if owner != nil {
		store.Owner = owner.Clean()
	}
	return store
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (store *Store) SetGroup(group *Group) *Store {
	if group != nil {
		store.Group = group.Clean()
	}
	return store
}

// SetArchived устанавливает флаг нахождения в архиве.
func (store *Store) SetArchived(archived bool) *Store {
	store.Archived = &archived
	return store
}

// SetCode устанавливает Код Склада.
func (store *Store) SetCode(code string) *Store {
	store.Code = &code
	return store
}

// SetDescription устанавливает Комментарий к Складу.
func (store *Store) SetDescription(description string) *Store {
	store.Description = &description
	return store
}

// SetExternalCode устанавливает Внешний код Склада.
func (store *Store) SetExternalCode(externalCode string) *Store {
	store.ExternalCode = &externalCode
	return store
}

// SetAddressFull устанавливает Адрес склада с детализацией по отдельным полям.
//
// Передача nil передаёт сброс значения (null).
func (store *Store) SetAddressFull(addressFull *Address) *Store {
	if addressFull == nil {
		store.SetAddress("")
	} else {
		store.AddressFull = addressFull
	}
	return store
}

// SetAddress устанавливает Адрес склада.
func (store *Store) SetAddress(address string) *Store {
	store.Address = &address
	return store
}

// SetName устанавливает Наименование склада.
func (store *Store) SetName(name string) *Store {
	store.Name = &name
	return store
}

// SetMeta устанавливает Метаданные склада.
func (store *Store) SetMeta(meta *Meta) *Store {
	store.Meta = meta
	return store
}

// SetParent устанавливает Метаданные родительского склада (Группы).
func (store *Store) SetParent(parent *Store) *Store {
	if parent != nil {
		store.Parent = parent.Clean()
	}
	return store
}

// SetShared устанавливает флаг общего доступа.
func (store *Store) SetShared(shared bool) *Store {
	store.Shared = &shared
	return store
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (store *Store) SetAttributes(attributes ...*Attribute) *Store {
	store.Attributes.Push(attributes...)
	return store
}

// String реализует интерфейс [fmt.Stringer].
func (store Store) String() string {
	return Stringify(store)
}

// MetaType возвращает код сущности.
func (Store) MetaType() MetaType {
	return MetaTypeStore
}

// Update shortcut
func (store Store) Update(ctx context.Context, client *Client, params ...*Params) (*Store, *resty.Response, error) {
	return NewStoreService(client).Update(ctx, store.GetID(), &store, params...)
}

// Create shortcut
func (store Store) Create(ctx context.Context, client *Client, params ...*Params) (*Store, *resty.Response, error) {
	return NewStoreService(client).Create(ctx, &store, params...)
}

// Delete shortcut
func (store Store) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewStoreService(client).Delete(ctx, store.GetID())
}

// Slot Ячейка склада.
//
// Код сущности: slot
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-yachejki-sklada
type Slot struct {
	AccountID    *uuid.UUID `json:"accountId,omitempty"`    // ID учётной записи
	Barcode      *string    `json:"barcode,omitempty"`      // Штрихкод ячейки
	ExternalCode *string    `json:"externalCode,omitempty"` // Внешний код Ячейки
	ID           *uuid.UUID `json:"id,omitempty"`           // ID Ячейки
	Meta         *Meta      `json:"meta,omitempty"`         // Метаданные Ячейки
	Name         *string    `json:"name,omitempty"`         // Наименование Ячейки
	Updated      *Timestamp `json:"updated,omitempty"`      // Момент последнего обновления Ячейки
	Zone         *Zone      `json:"zone,omitempty"`         // Зона ячейки
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (slot Slot) Clean() *Slot {
	if slot.Meta == nil {
		return nil
	}
	return &Slot{Meta: slot.Meta}
}

// GetAccountID возвращает ID учётной записи.
func (slot Slot) GetAccountID() uuid.UUID {
	return Deref(slot.AccountID)
}

// GetBarcode возвращает Штрихкод ячейки.
func (slot Slot) GetBarcode() string {
	return Deref(slot.Barcode)
}

// GetExternalCode возвращает Внешний код Ячейки.
func (slot Slot) GetExternalCode() string {
	return Deref(slot.ExternalCode)
}

// GetID возвращает ID Ячейки.
func (slot Slot) GetID() uuid.UUID {
	return Deref(slot.ID)
}

// GetMeta возвращает Метаданные Ячейки.
func (slot Slot) GetMeta() Meta {
	return Deref(slot.Meta)
}

// GetName возвращает Наименование Ячейки.
func (slot Slot) GetName() string {
	return Deref(slot.Name)
}

// GetUpdated возвращает Момент последнего обновления Ячейки.
func (slot Slot) GetUpdated() time.Time {
	return Deref(slot.Updated).Time()
}

// GetZone возвращает Зону ячейки.
func (slot Slot) GetZone() Zone {
	return Deref(slot.Zone)
}

// SetBarcode устанавливает Штрихкод ячейки.
func (slot *Slot) SetBarcode(barcode string) *Slot {
	slot.Barcode = &barcode
	return slot
}

// SetExternalCode устанавливает Внешний код ячейки.
func (slot *Slot) SetExternalCode(externalCode string) *Slot {
	slot.ExternalCode = &externalCode
	return slot
}

// SetMeta устанавливает Метаданные ячейки.
func (slot *Slot) SetMeta(meta *Meta) *Slot {
	slot.Meta = meta
	return slot
}

// SetName устанавливает Наименование ячейки.
func (slot *Slot) SetName(name string) *Slot {
	slot.Name = &name
	return slot
}

// String реализует интерфейс [fmt.Stringer].
func (slot Slot) String() string {
	return Stringify(slot)
}

// MetaType возвращает код сущности.
func (Slot) MetaType() MetaType {
	return MetaTypeSlot
}

// Zone Зона склада.
//
// Код сущности: storezone
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-zony-sklada
type Zone struct {
	AccountID    *uuid.UUID `json:"accountId,omitempty"`    // ID учётной записи
	ExternalCode *string    `json:"externalCode,omitempty"` // Внешний код Зоны
	ID           *uuid.UUID `json:"id,omitempty"`           // ID Зоны
	Meta         *Meta      `json:"meta,omitempty"`         // Метаданные Зоны
	Name         *string    `json:"name,omitempty"`         // Наименование Зоны
	Updated      *Timestamp `json:"updated,omitempty"`      // Момент последнего обновления Зоны
}

// GetAccountID возвращает ID учётной записи.
func (zone Zone) GetAccountID() uuid.UUID {
	return Deref(zone.AccountID)
}

// GetExternalCode возвращает Внешний код Зоны.
func (zone Zone) GetExternalCode() string {
	return Deref(zone.ExternalCode)
}

// GetID возвращает ID Зоны.
func (zone Zone) GetID() uuid.UUID {
	return Deref(zone.ID)
}

// GetMeta возвращает Метаданные Зоны.
func (zone Zone) GetMeta() Meta {
	return Deref(zone.Meta)
}

// GetName возвращает Наименование Зоны.
func (zone Zone) GetName() string {
	return Deref(zone.Name)
}

// GetUpdated возвращает Момент последнего обновления Зоны.
func (zone Zone) GetUpdated() time.Time {
	return Deref(zone.Updated).Time()
}

// SetExternalCode устанавливает Внешний код Зоны.
func (zone *Zone) SetExternalCode(externalCode string) *Zone {
	zone.ExternalCode = &externalCode
	return zone
}

// SetMeta устанавливает Метаданные Зоны.
func (zone *Zone) SetMeta(meta *Meta) *Zone {
	zone.Meta = meta
	return zone
}

// SetName устанавливает Наименование Зоны.
func (zone *Zone) SetName(name string) *Zone {
	zone.Name = &name
	return zone
}

// String реализует интерфейс [fmt.Stringer].
func (zone Zone) String() string {
	return Stringify(zone)
}

// MetaType возвращает код сущности.
func (Zone) MetaType() MetaType {
	return MetaTypeStoreZone
}

// StoreService описывает методы сервиса для работы со складами.
type StoreService interface {
	// GetList выполняет запрос на получение списка складов.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[Store], *resty.Response, error)

	// Create выполняет запрос на создание склада.
	// Обязательные поля для заполнения:
	//	- name (Наименования склада)
	// Принимает контекст, склад и опционально объект параметров запроса Params.
	// Возвращает созданный склад.
	Create(ctx context.Context, store *Store, params ...*Params) (*Store, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение складов.
	// Изменяемые склады должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список складов и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых складов.
	CreateUpdateMany(ctx context.Context, storeList Slice[Store], params ...*Params) (*Slice[Store], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление складов.
	// Принимает контекст и множество складов.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*Store) (*DeleteManyResponse, *resty.Response, error)

	// Delete выполняет запрос на удаление склада.
	// Принимает контекст и ID склада.
	// Возвращает «true» в случае успешного удаления склада.
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных складов.
	// Принимает контекст.
	// Возвращает объект метаданных MetaAttributesSharedWrapper.
	GetMetadata(ctx context.Context) (*MetaAttributesSharedWrapper, *resty.Response, error)

	// GetAttributeList выполняет запрос на получение списка доп полей.
	// Принимает контекст.
	// Возвращает объект List.
	GetAttributeList(ctx context.Context) (*List[Attribute], *resty.Response, error)

	// GetAttributeByID выполняет запрос на получение отдельного доп поля по ID.
	// Принимает контекст и ID доп поля.
	// Возвращает найденное доп поле.
	GetAttributeByID(ctx context.Context, id uuid.UUID) (*Attribute, *resty.Response, error)

	// CreateAttribute выполняет запрос на создание доп поля.
	// Принимает контекст и доп поле.
	// Возвращает созданное доп поле.
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)

	// CreateUpdateAttributeMany выполняет запрос на массовое создание и/или изменение доп полей.
	// Изменяемые доп поля должны содержать идентификатор в виде метаданных.
	// Принимает контекст и множество доп полей.
	// Возвращает список созданных и/или изменённых доп полей.
	CreateUpdateAttributeMany(ctx context.Context, attributes ...*Attribute) (*Slice[Attribute], *resty.Response, error)

	// UpdateAttribute выполняет запрос на изменения доп поля.
	// Принимает контекст, ID доп поля и доп поле.
	// Возвращает изменённое доп поле.
	UpdateAttribute(ctx context.Context, id uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)

	// DeleteAttribute выполняет запрос на удаление доп поля.
	// Принимает контекст и ID доп поля.
	// Возвращает «true» в случае успешного удаления доп поля.
	DeleteAttribute(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// DeleteAttributeMany выполняет запрос на массовое удаление доп полей.
	// Принимает контекст и множество доп полей.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteAttributeMany(ctx context.Context, attributes ...*Attribute) (*DeleteManyResponse, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельного склада по ID.
	// Принимает контекст, ID склада и опционально объект параметров запроса Params.
	// Возвращает склад.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Store, *resty.Response, error)

	// Update выполняет запрос на изменение склада.
	// Принимает контекст, ID склада, склад и опционально объект параметров запроса Params.
	// Возвращает изменённый склад.
	Update(ctx context.Context, id uuid.UUID, store *Store, params ...*Params) (*Store, *resty.Response, error)

	// GetNamedFilterList выполняет запрос на получение списка фильтров.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetNamedFilterList(ctx context.Context, params ...*Params) (*List[NamedFilter], *resty.Response, error)

	// GetNamedFilterByID выполняет запрос на получение отдельного фильтра по ID.
	// Принимает контекст и ID фильтра.
	// Возвращает найденный фильтр.
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)

	// GetSlotList выполняет запрос на получение списка всех ячеек склада.
	// Принимает контекст и ID склада.
	// Возвращает объект List.
	GetSlotList(ctx context.Context, storeID uuid.UUID) (*List[Slot], *resty.Response, error)

	// CreateSlot выполняет запрос на создание ячейки склада.
	// Обязательные поля для заполнения:
	//	- name (Наименование ячейки склада)
	// Принимает контекст, ID склада и ячейку склада.
	// Возвращает созданную ячейку склада.
	CreateSlot(ctx context.Context, storeID uuid.UUID, slot *Slot) (*Slot, *resty.Response, error)

	// CreateUpdateSlotMany выполняет запрос на массовое создание и/или изменение ячеек склада.
	// Изменяемые ячейки склада должны содержать идентификатор в виде метаданных.
	// Принимает контекст, ID склада и множество ячеек склада.
	// Возвращает список созданных и/или изменённых ячеек склада.
	CreateUpdateSlotMany(ctx context.Context, storeID uuid.UUID, slots ...*Slot) (*Slice[Slot], *resty.Response, error)

	// DeleteSlotMany выполняет запрос на массовое удаление ячеек склада.
	// Принимает контекст, ID склада и множество ячеек склада.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteSlotMany(ctx context.Context, storeID uuid.UUID, slots ...*Slot) (*DeleteManyResponse, *resty.Response, error)

	// DeleteSlot выполняет запрос на удаление ячейки склада.
	// Принимает контекст, ID склада и ID ячейки склада.
	// Возвращает «true» в случае успешного удаления ячейки склада.
	DeleteSlot(ctx context.Context, storeID, slotID uuid.UUID) (bool, *resty.Response, error)

	// GetSlotByID выполняет запрос на получение отдельной ячейки склада по ID.
	// Принимает контекст, ID склада и ID ячейки склада.
	// Возвращает ячейку склада.
	GetSlotByID(ctx context.Context, storeID, slotID uuid.UUID) (*Slot, *resty.Response, error)

	// UpdateSlot выполняет запрос на изменение ячейки склада.
	// Принимает контекст, ID склада, ID ячейки склада и ячейку склада.
	// Возвращает изменённую ячейку склада.
	UpdateSlot(ctx context.Context, storeID, slotID uuid.UUID, slot *Slot) (*Slot, *resty.Response, error)

	// GetZoneList выполняет запрос на получение списка всех зон склада.
	// Принимает контекст и ID склада.
	// Возвращает объект List.
	GetZoneList(ctx context.Context, storeID uuid.UUID) (*List[Zone], *resty.Response, error)

	// CreateZone выполняет запрос на создание зоны склада.
	// Обязательные поля для заполнения:
	//	- name (Наименование зоны склада)
	// Принимает контекст, ID склада и зону склада.
	// Возвращает созданную зону склада.
	CreateZone(ctx context.Context, storeID uuid.UUID, zone *Zone) (*Zone, *resty.Response, error)

	// CreateUpdateZoneMany выполняет запрос на массовое создание и/или изменение зон склада.
	// Изменяемые зоны склада должны содержать идентификатор в виде метаданных.
	// Принимает контекст, ID склада и множество зон склада.
	// Возвращает список созданных и/или изменённых зон склада.
	CreateUpdateZoneMany(ctx context.Context, storeID uuid.UUID, zones ...*Zone) (*Slice[Zone], *resty.Response, error)

	// DeleteZoneMany выполняет запрос на массовое удаление зон склада.
	// Принимает контекст, ID склада и множество зон склада.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteZoneMany(ctx context.Context, storeID uuid.UUID, zones ...*Zone) (*DeleteManyResponse, *resty.Response, error)

	// DeleteZone выполняет запрос на удаление зоны склада.
	// Принимает контекст, ID склада и ID зоны склада.
	// Возвращает «true» в случае успешного удаления зоны склада.
	DeleteZone(ctx context.Context, storeID, zoneID uuid.UUID) (bool, *resty.Response, error)

	// GetZoneByID выполняет запрос на получение отдельной зоны склада по ID.
	// Принимает контекст, ID склада и ID зоны склада.
	// Возвращает зону склада.
	GetZoneByID(ctx context.Context, storeID, zoneID uuid.UUID) (*Zone, *resty.Response, error)

	// UpdateZone выполняет запрос на изменение зоны склада.
	// Принимает контекст, ID склада, ID зоны склада и зону склада.
	// Возвращает изменённую зону склада.
	UpdateZone(ctx context.Context, storeID, zoneID uuid.UUID, zone *Zone) (*Zone, *resty.Response, error)
}

const (
	EndpointStore            = EndpointEntity + string(MetaTypeStore)
	EndpointStoreSlots       = EndpointStore + "/%s/slots"
	EndpointStoreSlotsID     = EndpointStoreSlots + "/%s"
	EndpointStoreSlotsDelete = EndpointStoreSlots + EndpointDelete
	EndpointStoreZones       = EndpointStore + "/%s/zones"
	EndpointStoreZonesID     = EndpointStoreZones + "/%s"
	EndpointStoreZonesDelete = EndpointStoreZones + EndpointDelete
)

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

func (service *storeService) GetSlotList(ctx context.Context, storeID uuid.UUID) (*List[Slot], *resty.Response, error) {
	path := fmt.Sprintf(EndpointStoreSlots, storeID)
	return NewRequestBuilder[List[Slot]](service.client, path).Get(ctx)
}

func (service *storeService) CreateSlot(ctx context.Context, storeID uuid.UUID, slot *Slot) (*Slot, *resty.Response, error) {
	path := fmt.Sprintf(EndpointStoreSlots, storeID)
	return NewRequestBuilder[Slot](service.client, path).Post(ctx, slot)
}

func (service *storeService) CreateUpdateSlotMany(ctx context.Context, storeID uuid.UUID, slots ...*Slot) (*Slice[Slot], *resty.Response, error) {
	path := fmt.Sprintf(EndpointStoreSlots, storeID)
	return NewRequestBuilder[Slice[Slot]](service.client, path).Post(ctx, slots)
}

func (service *storeService) DeleteSlotMany(ctx context.Context, storeID uuid.UUID, slots ...*Slot) (*DeleteManyResponse, *resty.Response, error) {
	path := fmt.Sprintf(EndpointStoreSlotsDelete, storeID)
	return NewRequestBuilder[DeleteManyResponse](service.client, path).Post(ctx, AsMetaWrapperSlice(slots))
}

func (service *storeService) GetSlotByID(ctx context.Context, storeID, slotID uuid.UUID) (*Slot, *resty.Response, error) {
	path := fmt.Sprintf(EndpointStoreSlotsID, storeID, slotID)
	return NewRequestBuilder[Slot](service.client, path).Get(ctx)
}

func (service *storeService) DeleteSlot(ctx context.Context, storeID, slotID uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf(EndpointStoreSlotsID, storeID, slotID)
	return NewRequestBuilder[any](service.client, path).Delete(ctx)
}

func (service *storeService) UpdateSlot(ctx context.Context, storeID, slotID uuid.UUID, slot *Slot) (*Slot, *resty.Response, error) {
	path := fmt.Sprintf(EndpointStoreSlotsID, storeID, slotID)
	return NewRequestBuilder[Slot](service.client, path).Put(ctx, slot)
}

func (service *storeService) GetZoneList(ctx context.Context, storeID uuid.UUID) (*List[Zone], *resty.Response, error) {
	path := fmt.Sprintf(EndpointStoreZones, storeID)
	return NewRequestBuilder[List[Zone]](service.client, path).Get(ctx)
}

func (service *storeService) CreateZone(ctx context.Context, storeID uuid.UUID, zone *Zone) (*Zone, *resty.Response, error) {
	path := fmt.Sprintf(EndpointStoreZones, storeID)
	return NewRequestBuilder[Zone](service.client, path).Post(ctx, zone)
}

func (service *storeService) CreateUpdateZoneMany(ctx context.Context, storeID uuid.UUID, zones ...*Zone) (*Slice[Zone], *resty.Response, error) {
	path := fmt.Sprintf(EndpointStoreZones, storeID)
	return NewRequestBuilder[Slice[Zone]](service.client, path).Post(ctx, zones)
}

func (service *storeService) DeleteZoneMany(ctx context.Context, storeID uuid.UUID, zones ...*Zone) (*DeleteManyResponse, *resty.Response, error) {
	path := fmt.Sprintf(EndpointStoreZonesDelete, storeID)
	return NewRequestBuilder[DeleteManyResponse](service.client, path).Post(ctx, AsMetaWrapperSlice(zones))
}

func (service *storeService) DeleteZone(ctx context.Context, storeID, zoneID uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf(EndpointStoreZonesID, storeID, zoneID)
	return NewRequestBuilder[any](service.client, path).Delete(ctx)
}

func (service *storeService) GetZoneByID(ctx context.Context, storeID, zoneID uuid.UUID) (*Zone, *resty.Response, error) {
	path := fmt.Sprintf(EndpointStoreZonesID, storeID, zoneID)
	return NewRequestBuilder[Zone](service.client, path).Get(ctx)
}

func (service *storeService) UpdateZone(ctx context.Context, storeID, zoneID uuid.UUID, zone *Zone) (*Zone, *resty.Response, error) {
	path := fmt.Sprintf(EndpointStoreZonesID, storeID, zoneID)
	return NewRequestBuilder[Zone](service.client, path).Put(ctx, zone)
}

// NewStoreService принимает [Client] и возвращает сервис для работы со складами.
func NewStoreService(client *Client) StoreService {
	e := NewEndpoint(client, EndpointStore)
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
