package moysklad

import (
	"context"
	"fmt"
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
func (s *ProcessingService) GetMaterials(ctx context.Context, id *uuid.UUID) (*List[ProcessingPlanMaterial], *Response, error) {
	path := fmt.Sprintf("%s/materials", id)
	return NewRequestBuilder[List[ProcessingPlanMaterial]](s.Endpoint, ctx).WithPath(path).Get()
}

// CreateMaterial Создать материал Тех. карты.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-sozdat-material-teh-karty
func (s *ProcessingService) CreateMaterial(ctx context.Context, id *uuid.UUID, material *ProcessingPlanMaterial) (*ProcessingPlanMaterial, *Response, error) {
	path := fmt.Sprintf("%s/materials", id)
	return NewRequestBuilder[ProcessingPlanMaterial](s.Endpoint, ctx).WithPath(path).WithBody(material).Post()
}

// CreateMaterials Создать несколько материалов Тех. карты.
func (s *ProcessingService) CreateMaterials(ctx context.Context, id *uuid.UUID, materials []*ProcessingPlanMaterial) (*Slice[ProcessingPlanMaterial], *Response, error) {
	path := fmt.Sprintf("%s/materials", id)
	return NewRequestBuilder[Slice[ProcessingPlanMaterial]](s.Endpoint, ctx).WithPath(path).WithBody(materials).Post()
}

// GetMaterialById Получить материал.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-poluchit-material
func (s *ProcessingService) GetMaterialById(ctx context.Context, id, materialId uuid.UUID) (*ProcessingPlanMaterial, *Response, error) {
	path := fmt.Sprintf("%s/materials/%s", id, materialId)
	return NewRequestBuilder[ProcessingPlanMaterial](s.Endpoint, ctx).WithPath(path).Get()
}

// UpdateMaterial Изменить материал.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-izmenit-material
func (s *ProcessingService) UpdateMaterial(ctx context.Context, id, materialId uuid.UUID, material *ProcessingPlanMaterial) (*ProcessingPlanMaterial, *Response, error) {
	path := fmt.Sprintf("%s/materials/%s", id, materialId)
	return NewRequestBuilder[ProcessingPlanMaterial](s.Endpoint, ctx).WithPath(path).WithBody(material).Put()
}

// DeleteMaterial Удалить материал.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-udalit-material
func (s *ProcessingService) DeleteMaterial(ctx context.Context, id, materialId uuid.UUID) (bool, *Response, error) {
	path := fmt.Sprintf("%s/materials/%s", id, materialId)
	return NewRequestBuilder[any](s.Endpoint, ctx).WithPath(path).Delete()
}

// GetProducts Получить продукты Тех. карты.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-poluchit-produkty-teh-karty
func (s *ProcessingService) GetProducts(ctx context.Context, id *uuid.UUID) (*List[ProcessingPlanProduct], *Response, error) {
	path := fmt.Sprintf("%s/products", id)
	return NewRequestBuilder[List[ProcessingPlanProduct]](s.Endpoint, ctx).WithPath(path).Get()
}

// CreateProduct Создать продукт Тех. карты.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-sozdat-produkt-teh-karty
func (s *ProcessingService) CreateProduct(ctx context.Context, id *uuid.UUID, product *ProcessingPlanProduct) (*ProcessingPlanProduct, *Response, error) {
	path := fmt.Sprintf("%s/products", id)
	return NewRequestBuilder[ProcessingPlanProduct](s.Endpoint, ctx).WithPath(path).WithBody(product).Post()
}

// CreateProducts Создать несколько продуктов Тех. карты.
func (s *ProcessingService) CreateProducts(ctx context.Context, id *uuid.UUID, products []*ProcessingPlanProduct) (*Slice[ProcessingPlanProduct], *Response, error) {
	path := fmt.Sprintf("%s/products", id)
	return NewRequestBuilder[Slice[ProcessingPlanProduct]](s.Endpoint, ctx).WithPath(path).WithBody(products).Post()
}

// GetProductById Получить продукт.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-poluchit-produkt
func (s *ProcessingService) GetProductById(ctx context.Context, id, productId uuid.UUID) (*ProcessingPlanProduct, *Response, error) {
	path := fmt.Sprintf("%s/products/%s", id, productId)
	return NewRequestBuilder[ProcessingPlanProduct](s.Endpoint, ctx).WithPath(path).Get()
}

// UpdateProduct Изменить продукт.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-izmenit-produkt
func (s *ProcessingService) UpdateProduct(ctx context.Context, id, productId uuid.UUID, product *ProcessingPlanProduct) (*ProcessingPlanProduct, *Response, error) {
	path := fmt.Sprintf("%s/products/%s", id, productId)
	return NewRequestBuilder[ProcessingPlanProduct](s.Endpoint, ctx).WithPath(path).WithBody(product).Put()
}

// DeleteProduct Удалить продукт.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-udalit-produkt
func (s *ProcessingService) DeleteProduct(ctx context.Context, id, productId uuid.UUID) (bool, *Response, error) {
	path := fmt.Sprintf("%s/products/%s", id, productId)
	return NewRequestBuilder[any](s.Endpoint, ctx).WithPath(path).Delete()
}
