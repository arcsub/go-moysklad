package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// StoreService
// Сервис для работы со складами.
type StoreService struct {
	Endpoint
	endpointGetList[Store]
	endpointCreate[Store]
	endpointCreateUpdateDeleteMany[Store]
	endpointDelete
	endpointMetadata[MetadataAttributeShared]
	endpointAttributes
	endpointGetById[Store]
	endpointUpdate[Store]
	endpointNamedFilter
}

func NewStoreService(client *Client) *StoreService {
	e := NewEndpoint(client, "entity/store")
	return &StoreService{
		Endpoint:                       e,
		endpointGetList:                endpointGetList[Store]{e},
		endpointCreate:                 endpointCreate[Store]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[Store]{e},
		endpointDelete:                 endpointDelete{e},
		endpointMetadata:               endpointMetadata[MetadataAttributeShared]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointGetById:                endpointGetById[Store]{e},
		endpointUpdate:                 endpointUpdate[Store]{e},
		endpointNamedFilter:            endpointNamedFilter{e},
	}
}

// GetSlots Получить список всех Ячеек Склада.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-poluchit-qchejki-sklada
func (s *StoreService) GetSlots(ctx context.Context, storeId uuid.UUID) (*List[Slot], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/slots", s.uri, storeId)
	return NewRequestBuilder[List[Slot]](s.client, path).Get(ctx)
}

// CreateSlot Запрос на создание Ячейки Склада.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-sozdat-qchejku-sklada
func (s *StoreService) CreateSlot(ctx context.Context, storeId uuid.UUID, slot *Slot) (*Slot, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/slots", s.uri, storeId)
	return NewRequestBuilder[Slot](s.client, path).Post(ctx, slot)
}

// CreateOrUpdateSlots Запрос создания и обновления нескольких Ячеек Склада.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-massowoe-sozdanie-i-obnowlenie-qcheek-sklada
func (s *StoreService) CreateOrUpdateSlots(ctx context.Context, storeId uuid.UUID, slots []*Slot) (*[]Slot, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/slots", s.uri, storeId)
	return NewRequestBuilder[[]Slot](s.client, path).Post(ctx, slots)
}

// DeleteSlots Запрос на массовое удаление Ячеек склада.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-massowoe-udalenie-qcheek-sklada
func (s *StoreService) DeleteSlots(ctx context.Context, storeId uuid.UUID, slots []*Slot) (*DeleteManyResponse, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/slots/delete", s.uri, storeId)
	return NewRequestBuilder[DeleteManyResponse](s.client, path).Post(ctx, slots)
}

// GetSlotById Запрос на получение отдельной Ячейки Склада с указанным id.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-poluchit-qchejku-sklada
func (s *StoreService) GetSlotById(ctx context.Context, storeId, slotId uuid.UUID) (*Slot, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/slots/%s", s.uri, storeId, slotId)
	return NewRequestBuilder[Slot](s.client, path).Get(ctx)
}

// DeleteSlot Запрос на удаление Ячейки склада с указанным id.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-udalit-qchejku-sklada
func (s *StoreService) DeleteSlot(ctx context.Context, storeId, slotId uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/slots/%s", s.uri, storeId, slotId)
	return NewRequestBuilder[any](s.client, path).Delete(ctx)
}

// UpdateSlot Запрос на обновление Ячейки склада.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-izmenit-qchejku-sklada
func (s *StoreService) UpdateSlot(ctx context.Context, storeId, slotId uuid.UUID, slot *Slot) (*Slot, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/slots/%s", s.uri, storeId, slotId)
	return NewRequestBuilder[Slot](s.client, path).Put(ctx, slot)
}

// GetZones Получить список всех Зон.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-poluchit-zony-sklada
func (s *StoreService) GetZones(ctx context.Context, storeId uuid.UUID) (*List[Zone], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/zones", s.uri, storeId)
	return NewRequestBuilder[List[Zone]](s.client, path).Get(ctx)
}

// CreateZone Запрос на создание Зоны склада.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-sozdat-zonu-sklada
func (s *StoreService) CreateZone(ctx context.Context, storeId uuid.UUID, zone *Zone) (*Zone, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/zones", s.uri, storeId)
	return NewRequestBuilder[Zone](s.client, path).Post(ctx, zone)
}

// CreateOrUpdateZones Запрос на создание и обновление нескольких Зон склада.
func (s *StoreService) CreateOrUpdateZones(ctx context.Context, storeId uuid.UUID, zones []*Zone) (*[]Zone, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/zones", s.uri, storeId)
	return NewRequestBuilder[[]Zone](s.client, path).Post(ctx, zones)
}

// DeleteZones Запрос на массовое удаление Зон склада.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-massowoe-udalenie-zon-sklada
func (s *StoreService) DeleteZones(ctx context.Context, storeId uuid.UUID, zones []*Zone) (*DeleteManyResponse, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/zones/delete", s.uri, storeId)
	return NewRequestBuilder[DeleteManyResponse](s.client, path).Post(ctx, zones)
}

// DeleteZone Запрос на удаление Зоны склада с указанным id.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-udalit-zonu-sklada
func (s *StoreService) DeleteZone(ctx context.Context, storeId, zoneId uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/zones/%s", s.uri, storeId, zoneId)
	return NewRequestBuilder[any](s.client, path).Delete(ctx)
}

// GetZoneById Запрос на получение отдельной Зоны Склада с указанным id.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-poluchit-zonu-sklada
func (s *StoreService) GetZoneById(ctx context.Context, storeId, zoneId uuid.UUID) (*Zone, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/zones/%s", s.uri, storeId, zoneId)
	return NewRequestBuilder[Zone](s.client, path).Get(ctx)
}

// UpdateZone Запрос на обновление Зоны склада.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-izmenit-zonu-sklada
func (s *StoreService) UpdateZone(ctx context.Context, storeId, zoneId uuid.UUID, zone *Zone) (*Zone, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/zones/%s", s.uri, storeId, zoneId)
	return NewRequestBuilder[Zone](s.client, path).Put(ctx, zone)
}
