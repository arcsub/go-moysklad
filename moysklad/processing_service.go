package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// ProcessingService
// Сервис для работы с Техоперациями.
type ProcessingService interface {
	GetList(ctx context.Context, params *Params) (*List[Processing], *resty.Response, error)
	Create(ctx context.Context, processing *Processing, params *Params) (*Processing, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, processingList []*Processing, params *Params) (*[]Processing, *resty.Response, error)
	DeleteMany(ctx context.Context, processingList []*Processing) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Processing, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, processing *Processing, params *Params) (*Processing, *resty.Response, error)
	//endpointTemplate[Processing]
	//endpointTemplateBasedOn[Processing, ProcessingTemplateArg]
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id *uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error)
	UpdateAttribute(ctx context.Context, id *uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList []*Attribute) (*DeleteManyResponse, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*Processing, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetMaterials(ctx context.Context, id *uuid.UUID) (*List[ProcessingPlanMaterial], *resty.Response, error)
	CreateMaterial(ctx context.Context, id *uuid.UUID, material *ProcessingPlanMaterial) (*ProcessingPlanMaterial, *resty.Response, error)
	CreateMaterials(ctx context.Context, id *uuid.UUID, materials []*ProcessingPlanMaterial) (*[]ProcessingPlanMaterial, *resty.Response, error)
	GetMaterialById(ctx context.Context, id, materialID *uuid.UUID) (*ProcessingPlanMaterial, *resty.Response, error)
	UpdateMaterial(ctx context.Context, id, materialID *uuid.UUID, material *ProcessingPlanMaterial) (*ProcessingPlanMaterial, *resty.Response, error)
	DeleteMaterial(ctx context.Context, id, materialID *uuid.UUID) (bool, *resty.Response, error)
	GetProducts(ctx context.Context, id *uuid.UUID) (*List[ProcessingPlanProduct], *resty.Response, error)
	CreateProduct(ctx context.Context, id *uuid.UUID, product *ProcessingPlanProduct) (*ProcessingPlanProduct, *resty.Response, error)
	CreateProducts(ctx context.Context, id *uuid.UUID, products []*ProcessingPlanProduct) (*[]ProcessingPlanProduct, *resty.Response, error)
	GetProductById(ctx context.Context, id, productID *uuid.UUID) (*ProcessingPlanProduct, *resty.Response, error)
	UpdateProduct(ctx context.Context, id, productID *uuid.UUID, product *ProcessingPlanProduct) (*ProcessingPlanProduct, *resty.Response, error)
	DeleteProduct(ctx context.Context, id, productID *uuid.UUID) (bool, *resty.Response, error)
}

type processingService struct {
	Endpoint
	endpointGetList[Processing]
	endpointCreate[Processing]
	endpointCreateUpdateMany[Processing]
	endpointDeleteMany[Processing]
	endpointDelete
	endpointGetById[Processing]
	endpointUpdate[Processing]
	//endpointTemplate[Processing]
	//endpointTemplateBasedOn[Processing, ProcessingTemplateArg]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointAttributes
	endpointSyncID[Processing]
	endpointRemove
}

func NewProcessingService(client *Client) ProcessingService {
	e := NewEndpoint(client, "entity/processing")
	return &processingService{
		Endpoint:                 e,
		endpointGetList:          endpointGetList[Processing]{e},
		endpointCreate:           endpointCreate[Processing]{e},
		endpointCreateUpdateMany: endpointCreateUpdateMany[Processing]{e},
		endpointDeleteMany:       endpointDeleteMany[Processing]{e},
		endpointDelete:           endpointDelete{e},
		endpointGetById:          endpointGetById[Processing]{e},
		endpointUpdate:           endpointUpdate[Processing]{e},
		endpointMetadata:         endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointAttributes:       endpointAttributes{e},
		endpointSyncID:           endpointSyncID[Processing]{e},
		endpointRemove:           endpointRemove{e},
	}
}

// GetMaterials Получить материалы Тех. карты.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-poluchit-materialy-teh-karty
func (s *processingService) GetMaterials(ctx context.Context, id *uuid.UUID) (*List[ProcessingPlanMaterial], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/materials", s.uri, id)
	return NewRequestBuilder[List[ProcessingPlanMaterial]](s.client, path).Get(ctx)
}

// CreateMaterial Создать материал Тех. карты.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-sozdat-material-teh-karty
func (s *processingService) CreateMaterial(ctx context.Context, id *uuid.UUID, material *ProcessingPlanMaterial) (*ProcessingPlanMaterial, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/materials", s.uri, id)
	return NewRequestBuilder[ProcessingPlanMaterial](s.client, path).Post(ctx, material)
}

// CreateMaterials Создать несколько материалов Тех. карты.
func (s *processingService) CreateMaterials(ctx context.Context, id *uuid.UUID, materials []*ProcessingPlanMaterial) (*[]ProcessingPlanMaterial, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/materials", s.uri, id)
	return NewRequestBuilder[[]ProcessingPlanMaterial](s.client, path).Post(ctx, materials)
}

// GetMaterialById Получить материал.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-poluchit-material
func (s *processingService) GetMaterialById(ctx context.Context, id, materialId *uuid.UUID) (*ProcessingPlanMaterial, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/materials/%s", s.uri, id, materialId)
	return NewRequestBuilder[ProcessingPlanMaterial](s.client, path).Get(ctx)
}

// UpdateMaterial Изменить материал.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-izmenit-material
func (s *processingService) UpdateMaterial(ctx context.Context, id, materialId *uuid.UUID, material *ProcessingPlanMaterial) (*ProcessingPlanMaterial, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/materials/%s", s.uri, id, materialId)
	return NewRequestBuilder[ProcessingPlanMaterial](s.client, path).Put(ctx, material)
}

// DeleteMaterial Удалить материал.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-udalit-material
func (s *processingService) DeleteMaterial(ctx context.Context, id, materialId *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/materials/%s", s.uri, id, materialId)
	return NewRequestBuilder[any](s.client, path).Delete(ctx)
}

// GetProducts Получить продукты Тех. карты.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-poluchit-produkty-teh-karty
func (s *processingService) GetProducts(ctx context.Context, id *uuid.UUID) (*List[ProcessingPlanProduct], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/products", s.uri, id)
	return NewRequestBuilder[List[ProcessingPlanProduct]](s.client, path).Get(ctx)
}

// CreateProduct Создать продукт Тех. карты.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-sozdat-produkt-teh-karty
func (s *processingService) CreateProduct(ctx context.Context, id *uuid.UUID, product *ProcessingPlanProduct) (*ProcessingPlanProduct, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/products", s.uri, id)
	return NewRequestBuilder[ProcessingPlanProduct](s.client, path).Post(ctx, product)
}

// CreateProducts Создать несколько продуктов Тех. карты.
func (s *processingService) CreateProducts(ctx context.Context, id *uuid.UUID, products []*ProcessingPlanProduct) (*[]ProcessingPlanProduct, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/products", s.uri, id)
	return NewRequestBuilder[[]ProcessingPlanProduct](s.client, path).Post(ctx, products)
}

// GetProductById Получить продукт.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-poluchit-produkt
func (s *processingService) GetProductById(ctx context.Context, id, productId *uuid.UUID) (*ProcessingPlanProduct, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/products/%s", s.uri, id, productId)
	return NewRequestBuilder[ProcessingPlanProduct](s.client, path).Get(ctx)
}

// UpdateProduct Изменить продукт.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-izmenit-produkt
func (s *processingService) UpdateProduct(ctx context.Context, id, productId *uuid.UUID, product *ProcessingPlanProduct) (*ProcessingPlanProduct, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/products/%s", s.uri, id, productId)
	return NewRequestBuilder[ProcessingPlanProduct](s.client, path).Put(ctx, product)
}

// DeleteProduct Удалить продукт.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-udalit-produkt
func (s *processingService) DeleteProduct(ctx context.Context, id, productId *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/products/%s", s.uri, id, productId)
	return NewRequestBuilder[any](s.client, path).Delete(ctx)
}
