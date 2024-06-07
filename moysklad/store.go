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
	Attributes   Attributes       `json:"attributes,omitempty"`
}

func (s Store) String() string {
	return Stringify(s)
}

func (s Store) MetaType() MetaType {
	return MetaTypeStore
}

// StoreService
// Сервис для работы со складами.
type StoreService interface {
	GetList(ctx context.Context, params *Params) (*List[Store], *resty.Response, error)
	Create(ctx context.Context, store *Store, params *Params) (*Store, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, storeList []*Store, params *Params) (*[]Store, *resty.Response, error)
	DeleteMany(ctx context.Context, storeList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetadataAttributeShared, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id *uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error)
	UpdateAttribute(ctx context.Context, id *uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Store, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, store *Store, params *Params) (*Store, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
	GetSlots(ctx context.Context, storeID *uuid.UUID) (*List[Slot], *resty.Response, error)
	CreateSlot(ctx context.Context, storeID *uuid.UUID, slot *Slot) (*Slot, *resty.Response, error)
	CreateOrUpdateSlots(ctx context.Context, storeID *uuid.UUID, slots []*Slot) (*[]Slot, *resty.Response, error)
	DeleteSlots(ctx context.Context, storeID *uuid.UUID, slots []*Slot) (*DeleteManyResponse, *resty.Response, error)
	GetSlotByID(ctx context.Context, storeID, slotID *uuid.UUID) (*Slot, *resty.Response, error)
	DeleteSlot(ctx context.Context, storeID, slotID *uuid.UUID) (bool, *resty.Response, error)
	UpdateSlot(ctx context.Context, storeID, slotID *uuid.UUID, slot *Slot) (*Slot, *resty.Response, error)
	GetZones(ctx context.Context, storeID *uuid.UUID) (*List[Zone], *resty.Response, error)
	CreateZone(ctx context.Context, storeID *uuid.UUID, zone *Zone) (*Zone, *resty.Response, error)
	CreateOrUpdateZones(ctx context.Context, storeID *uuid.UUID, zones []*Zone) (*[]Zone, *resty.Response, error)
	DeleteZones(ctx context.Context, storeID *uuid.UUID, zones []*Zone) (*DeleteManyResponse, *resty.Response, error)
	DeleteZone(ctx context.Context, storeID, zoneID *uuid.UUID) (bool, *resty.Response, error)
	GetZoneByID(ctx context.Context, storeID, zoneID *uuid.UUID) (*Zone, *resty.Response, error)
	UpdateZone(ctx context.Context, storeID, zoneID *uuid.UUID, zone *Zone) (*Zone, *resty.Response, error)
}

type storeService struct {
	Endpoint
	endpointGetList[Store]
	endpointCreate[Store]
	endpointCreateUpdateMany[Store]
	endpointDeleteMany[Store]
	endpointDelete
	endpointMetadata[MetadataAttributeShared]
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
		endpointMetadata:         endpointMetadata[MetadataAttributeShared]{e},
		endpointAttributes:       endpointAttributes{e},
		endpointGetByID:          endpointGetByID[Store]{e},
		endpointUpdate:           endpointUpdate[Store]{e},
		endpointNamedFilter:      endpointNamedFilter{e},
	}
}

// GetSlots Получить список всех Ячеек Склада.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-poluchit-qchejki-sklada
func (s *storeService) GetSlots(ctx context.Context, storeID *uuid.UUID) (*List[Slot], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/slots", s.uri, storeID)
	return NewRequestBuilder[List[Slot]](s.client, path).Get(ctx)
}

// CreateSlot Запрос на создание Ячейки Склада.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-sozdat-qchejku-sklada
func (s *storeService) CreateSlot(ctx context.Context, storeID *uuid.UUID, slot *Slot) (*Slot, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/slots", s.uri, storeID)
	return NewRequestBuilder[Slot](s.client, path).Post(ctx, slot)
}

// CreateOrUpdateSlots Запрос создания и обновления нескольких Ячеек Склада.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-massowoe-sozdanie-i-obnowlenie-qcheek-sklada
func (s *storeService) CreateOrUpdateSlots(ctx context.Context, storeID *uuid.UUID, slots []*Slot) (*[]Slot, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/slots", s.uri, storeID)
	return NewRequestBuilder[[]Slot](s.client, path).Post(ctx, slots)
}

// DeleteSlots Запрос на массовое удаление Ячеек склада.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-massowoe-udalenie-qcheek-sklada
func (s *storeService) DeleteSlots(ctx context.Context, storeID *uuid.UUID, slots []*Slot) (*DeleteManyResponse, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/slots/delete", s.uri, storeID)
	return NewRequestBuilder[DeleteManyResponse](s.client, path).Post(ctx, slots)
}

// GetSlotByID Запрос на получение отдельной Ячейки Склада с указанным id.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-poluchit-qchejku-sklada
func (s *storeService) GetSlotByID(ctx context.Context, storeID, slotID *uuid.UUID) (*Slot, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/slots/%s", s.uri, storeID, slotID)
	return NewRequestBuilder[Slot](s.client, path).Get(ctx)
}

// DeleteSlot Запрос на удаление Ячейки склада с указанным id.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-udalit-qchejku-sklada
func (s *storeService) DeleteSlot(ctx context.Context, storeID, slotID *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/slots/%s", s.uri, storeID, slotID)
	return NewRequestBuilder[any](s.client, path).Delete(ctx)
}

// UpdateSlot Запрос на обновление Ячейки склада.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-izmenit-qchejku-sklada
func (s *storeService) UpdateSlot(ctx context.Context, storeID, slotID *uuid.UUID, slot *Slot) (*Slot, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/slots/%s", s.uri, storeID, slotID)
	return NewRequestBuilder[Slot](s.client, path).Put(ctx, slot)
}

// GetZones Получить список всех Зон.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-poluchit-zony-sklada
func (s *storeService) GetZones(ctx context.Context, storeID *uuid.UUID) (*List[Zone], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/zones", s.uri, storeID)
	return NewRequestBuilder[List[Zone]](s.client, path).Get(ctx)
}

// CreateZone Запрос на создание Зоны склада.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-sozdat-zonu-sklada
func (s *storeService) CreateZone(ctx context.Context, storeID *uuid.UUID, zone *Zone) (*Zone, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/zones", s.uri, storeID)
	return NewRequestBuilder[Zone](s.client, path).Post(ctx, zone)
}

// CreateOrUpdateZones Запрос на создание и обновление нескольких Зон склада.
func (s *storeService) CreateOrUpdateZones(ctx context.Context, storeID *uuid.UUID, zones []*Zone) (*[]Zone, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/zones", s.uri, storeID)
	return NewRequestBuilder[[]Zone](s.client, path).Post(ctx, zones)
}

// DeleteZones Запрос на массовое удаление Зон склада.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-massowoe-udalenie-zon-sklada
func (s *storeService) DeleteZones(ctx context.Context, storeID *uuid.UUID, zones []*Zone) (*DeleteManyResponse, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/zones/delete", s.uri, storeID)
	return NewRequestBuilder[DeleteManyResponse](s.client, path).Post(ctx, zones)
}

// DeleteZone Запрос на удаление Зоны склада с указанным id.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-udalit-zonu-sklada
func (s *storeService) DeleteZone(ctx context.Context, storeID, zoneID *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/zones/%s", s.uri, storeID, zoneID)
	return NewRequestBuilder[any](s.client, path).Delete(ctx)
}

// GetZoneByID Запрос на получение отдельной Зоны Склада с указанным id.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-poluchit-zonu-sklada
func (s *storeService) GetZoneByID(ctx context.Context, storeID, zoneID *uuid.UUID) (*Zone, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/zones/%s", s.uri, storeID, zoneID)
	return NewRequestBuilder[Zone](s.client, path).Get(ctx)
}

// UpdateZone Запрос на обновление Зоны склада.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-izmenit-zonu-sklada
func (s *storeService) UpdateZone(ctx context.Context, storeID, zoneID *uuid.UUID, zone *Zone) (*Zone, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/zones/%s", s.uri, storeID, zoneID)
	return NewRequestBuilder[Zone](s.client, path).Put(ctx, zone)
}
