package moysklad

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

// VariantService
// Сервис для работы с модификациями.
type VariantService struct {
	Endpoint
	endpointGetList[Variant]
	endpointCreate[Variant]
	endpointCreateUpdateDeleteMany[Variant]
	endpointDelete
	endpointGetById[Variant]
	endpointUpdate[Variant]
	endpointMetadata[MetadataVariant]
	endpointImages
	endpointNamedFilter
}

func NewVariantService(client *Client) *VariantService {
	e := NewEndpoint(client, "entity/variant")
	return &VariantService{
		Endpoint:                       e,
		endpointGetList:                endpointGetList[Variant]{e},
		endpointCreate:                 endpointCreate[Variant]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[Variant]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[Variant]{e},
		endpointUpdate:                 endpointUpdate[Variant]{e},
		endpointMetadata:               endpointMetadata[MetadataVariant]{e},
		endpointImages:                 endpointImages{e},
		endpointNamedFilter:            endpointNamedFilter{e},
	}
}

// CreateCharacteristic Создать характеристику.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-harakteristiki-modifikacij-sozdat-harakteristiku
func (s *VariantService) CreateCharacteristic(ctx context.Context, characteristic *Characteristic) (*Characteristic, *Response, error) {
	path := "metadata/characteristics"
	return NewRequestBuilder[Characteristic](s.Endpoint, ctx).WithPath(path).WithBody(characteristic).Post()
}

// CreateCharacteristics Массовое создание Характеристик.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-harakteristiki-modifikacij-massowoe-sozdanie-harakteristik
func (s *VariantService) CreateCharacteristics(ctx context.Context, characteristics []*Characteristic) (*Slice[Characteristic], *Response, error) {
	path := "metadata/characteristics"
	return NewRequestBuilder[Slice[Characteristic]](s.Endpoint, ctx).WithPath(path).WithBody(characteristics).Post()
}

// GetCharacteristicById Получить Характеристику.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-harakteristiki-modifikacij-poluchit-harakteristiku
func (s *VariantService) GetCharacteristicById(ctx context.Context, id *uuid.UUID) (*Characteristic, *Response, error) {
	path := fmt.Sprintf("metadata/characteristics/%s", id)
	return NewRequestBuilder[Characteristic](s.Endpoint, ctx).WithPath(path).Get()
}

// UpdateCharacteristic Изменить характеристику.
func (s *VariantService) UpdateCharacteristic(ctx context.Context, id *uuid.UUID, characteristic *Characteristic) (*Characteristic, *Response, error) {
	path := fmt.Sprintf("metadata/characteristics/%s", id)
	return NewRequestBuilder[Characteristic](s.Endpoint, ctx).WithPath(path).WithBody(characteristic).Put()
}

// DeleteCharacteristic Удалить характеристику.
func (s *VariantService) DeleteCharacteristic(ctx context.Context, id *uuid.UUID) (bool, *Response, error) {
	path := fmt.Sprintf("metadata/characteristics/%s", id)
	return NewRequestBuilder[any](s.Endpoint, ctx).WithPath(path).Delete()
}
