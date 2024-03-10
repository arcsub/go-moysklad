package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// ProductionTaskService
// Сервис для работы с производственными заданиями
type ProductionTaskService struct {
	Endpoint
	endpointGetList[ProductionTask]
	endpointCreate[ProductionTask]
	endpointCreateUpdateDeleteMany[ProductionTask]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointAttributes
	endpointGetById[ProductionTask]
	endpointUpdate[ProductionTask]
	endpointDelete
	endpointPositions[ProductionRow]
}

func NewProductionTaskService(client *Client) *ProductionTaskService {
	e := NewEndpoint(client, "entity/productiontask")
	return &ProductionTaskService{
		Endpoint:                       e,
		endpointGetList:                endpointGetList[ProductionTask]{e},
		endpointCreate:                 endpointCreate[ProductionTask]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[ProductionTask]{e},
		endpointMetadata:               endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointGetById:                endpointGetById[ProductionTask]{e},
		endpointUpdate:                 endpointUpdate[ProductionTask]{e},
		endpointDelete:                 endpointDelete{e},
		endpointPositions:              endpointPositions[ProductionRow]{e},
	}
}

// GetProducts Получить Продукты производственного задания.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-poluchit-produkty-proizwodstwennogo-zadaniq
func (s *ProductionTaskService) GetProducts(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[ProductionTaskResult], *resty.Response, error) {
	path := fmt.Sprintf("%s/products", id)
	return NewRequestBuilder[MetaArray[ProductionTaskResult]](s.client, path).SetParams(params).Get(ctx)
}

// GetProductByID Получить продукт производственного задания.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-produkt-proizwodstwennogo-zadaniq
func (s *ProductionTaskService) GetProductByID(ctx context.Context, id *uuid.UUID, productID *uuid.UUID, params *Params) (*ProductionTaskResult, *resty.Response, error) {
	path := fmt.Sprintf("%s/products/%s", id, productID)
	return NewRequestBuilder[ProductionTaskResult](s.client, path).SetParams(params).Get(ctx)
}

// CreateProduct Создать продукт.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-sozdat-produkt
func (s *ProductionTaskService) CreateProduct(ctx context.Context, id *uuid.UUID, productionTaskResult *ProductionTaskResult, params *Params) (*ProductionTaskResult, *resty.Response, error) {
	path := fmt.Sprintf("%s/products", id) // fixme:в документации указан endpoint без 's' на конце, вероятно ошибка
	return NewRequestBuilder[ProductionTaskResult](s.client, path).SetParams(params).Post(ctx, productionTaskResult)
}

// UpdateProduct Изменить продукт.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-izmenit-produkt
func (s *ProductionTaskService) UpdateProduct(ctx context.Context, id *uuid.UUID, productID *uuid.UUID, productionTaskResult *ProductionTaskResult, params *Params) (*ProductionTaskResult, *resty.Response, error) {
	path := fmt.Sprintf("%s/products/%s", id, productID)
	return NewRequestBuilder[ProductionTaskResult](s.client, path).SetParams(params).Put(ctx, productionTaskResult)
}

// DeleteProduct Удалить продукт.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-udalit-produkt
func (s *ProductionTaskService) DeleteProduct(ctx context.Context, id *uuid.UUID, productID *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/products/%s", id, productID)
	return NewRequestBuilder[any](s.client, path).Delete(ctx)
}

// DeleteProductMany Массовое удаление продуктов Производственного задания.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-udalit-produkt
func (s *ProductionTaskService) DeleteProductMany(ctx context.Context, id *uuid.UUID) (*DeleteManyResponse, *resty.Response, error) {
	path := fmt.Sprintf("%s/products/delete", id)
	return NewRequestBuilder[DeleteManyResponse](s.client, path).Post(ctx, nil)
}
