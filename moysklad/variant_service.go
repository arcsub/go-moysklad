package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
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
func (s *VariantService) CreateCharacteristic(ctx context.Context, characteristic *Characteristic) (*Characteristic, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/characteristics", s.uri)
	return NewRequestBuilder[Characteristic](s.client, path).Post(ctx, characteristic)
}

// CreateCharacteristics Массовое создание Характеристик.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-harakteristiki-modifikacij-massowoe-sozdanie-harakteristik
func (s *VariantService) CreateCharacteristics(ctx context.Context, characteristics []*Characteristic) (*[]Characteristic, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/characteristics", s.uri)
	return NewRequestBuilder[[]Characteristic](s.client, path).Post(ctx, characteristics)
}

// GetCharacteristicById Получить Характеристику.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-harakteristiki-modifikacij-poluchit-harakteristiku
func (s *VariantService) GetCharacteristicById(ctx context.Context, id *uuid.UUID) (*Characteristic, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/characteristics/%s", s.uri, id)
	return NewRequestBuilder[Characteristic](s.client, path).Get(ctx)
}

// UpdateCharacteristic Изменить характеристику.
func (s *VariantService) UpdateCharacteristic(ctx context.Context, id *uuid.UUID, characteristic *Characteristic) (*Characteristic, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/characteristics/%s", s.uri, id)
	return NewRequestBuilder[Characteristic](s.client, path).Put(ctx, characteristic)
}

// DeleteCharacteristic Удалить характеристику.
func (s *VariantService) DeleteCharacteristic(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/characteristics/%s", s.uri, id)
	return NewRequestBuilder[any](s.client, path).Delete(ctx)
}
