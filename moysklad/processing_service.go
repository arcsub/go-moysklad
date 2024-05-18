package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// ProcessingService
// Сервис для работы с Техоперациями.
type ProcessingService struct {
	Endpoint
	endpointGetList[Processing]
	endpointCreate[Processing]
	endpointCreateUpdateDeleteMany[Processing]
	endpointDelete
	endpointGetById[Processing]
	endpointUpdate[Processing]
	endpointTemplate[Processing]
	endpointTemplateBasedOn[Processing, ProcessingTemplateArg]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointAttributes
	endpointSyncID[Processing]
	endpointRemove
}

func NewProcessingService(client *Client) *ProcessingService {
	e := NewEndpoint(client, "entity/processing")
	return &ProcessingService{
		Endpoint:                       e,
		endpointGetList:                endpointGetList[Processing]{e},
		endpointCreate:                 endpointCreate[Processing]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[Processing]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[Processing]{e},
		endpointUpdate:                 endpointUpdate[Processing]{e},
		endpointTemplate:               endpointTemplate[Processing]{e},
		endpointTemplateBasedOn:        endpointTemplateBasedOn[Processing, ProcessingTemplateArg]{e},
		endpointMetadata:               endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointSyncID:                 endpointSyncID[Processing]{e},
		endpointRemove:                 endpointRemove{e},
	}
}

// GetMaterials Получить материалы Тех. карты.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-poluchit-materialy-teh-karty
func (s *ProcessingService) GetMaterials(ctx context.Context, id *uuid.UUID) (*List[ProcessingPlanMaterial], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/materials", s.uri, id)
	return NewRequestBuilder[List[ProcessingPlanMaterial]](s.client, path).Get(ctx)
}

// CreateMaterial Создать материал Тех. карты.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-sozdat-material-teh-karty
func (s *ProcessingService) CreateMaterial(ctx context.Context, id *uuid.UUID, material *ProcessingPlanMaterial) (*ProcessingPlanMaterial, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/materials", s.uri, id)
	return NewRequestBuilder[ProcessingPlanMaterial](s.client, path).Post(ctx, material)
}

// CreateMaterials Создать несколько материалов Тех. карты.
func (s *ProcessingService) CreateMaterials(ctx context.Context, id *uuid.UUID, materials []*ProcessingPlanMaterial) (*[]ProcessingPlanMaterial, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/materials", s.uri, id)
	return NewRequestBuilder[[]ProcessingPlanMaterial](s.client, path).Post(ctx, materials)
}

// GetMaterialById Получить материал.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-poluchit-material
func (s *ProcessingService) GetMaterialById(ctx context.Context, id, materialId *uuid.UUID) (*ProcessingPlanMaterial, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/materials/%s", s.uri, id, materialId)
	return NewRequestBuilder[ProcessingPlanMaterial](s.client, path).Get(ctx)
}

// UpdateMaterial Изменить материал.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-izmenit-material
func (s *ProcessingService) UpdateMaterial(ctx context.Context, id, materialId *uuid.UUID, material *ProcessingPlanMaterial) (*ProcessingPlanMaterial, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/materials/%s", s.uri, id, materialId)
	return NewRequestBuilder[ProcessingPlanMaterial](s.client, path).Put(ctx, material)
}

// DeleteMaterial Удалить материал.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-udalit-material
func (s *ProcessingService) DeleteMaterial(ctx context.Context, id, materialId *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/materials/%s", s.uri, id, materialId)
	return NewRequestBuilder[any](s.client, path).Delete(ctx)
}

// GetProducts Получить продукты Тех. карты.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-poluchit-produkty-teh-karty
func (s *ProcessingService) GetProducts(ctx context.Context, id *uuid.UUID) (*List[ProcessingPlanProduct], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/products", s.uri, id)
	return NewRequestBuilder[List[ProcessingPlanProduct]](s.client, path).Get(ctx)
}

// CreateProduct Создать продукт Тех. карты.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-sozdat-produkt-teh-karty
func (s *ProcessingService) CreateProduct(ctx context.Context, id *uuid.UUID, product *ProcessingPlanProduct) (*ProcessingPlanProduct, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/products", s.uri, id)
	return NewRequestBuilder[ProcessingPlanProduct](s.client, path).Post(ctx, product)
}

// CreateProducts Создать несколько продуктов Тех. карты.
func (s *ProcessingService) CreateProducts(ctx context.Context, id *uuid.UUID, products []*ProcessingPlanProduct) (*[]ProcessingPlanProduct, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/products", s.uri, id)
	return NewRequestBuilder[[]ProcessingPlanProduct](s.client, path).Post(ctx, products)
}

// GetProductById Получить продукт.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-poluchit-produkt
func (s *ProcessingService) GetProductById(ctx context.Context, id, productId *uuid.UUID) (*ProcessingPlanProduct, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/products/%s", s.uri, id, productId)
	return NewRequestBuilder[ProcessingPlanProduct](s.client, path).Get(ctx)
}

// UpdateProduct Изменить продукт.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-izmenit-produkt
func (s *ProcessingService) UpdateProduct(ctx context.Context, id, productId *uuid.UUID, product *ProcessingPlanProduct) (*ProcessingPlanProduct, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/products/%s", s.uri, id, productId)
	return NewRequestBuilder[ProcessingPlanProduct](s.client, path).Put(ctx, product)
}

// DeleteProduct Удалить продукт.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-udalit-produkt
func (s *ProcessingService) DeleteProduct(ctx context.Context, id, productId *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/products/%s", s.uri, id, productId)
	return NewRequestBuilder[any](s.client, path).Delete(ctx)
}
