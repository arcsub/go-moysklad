package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// ProductionTaskService
// Сервис для работы с производственными заданиями
type ProductionTaskService interface {
	GetList(ctx context.Context, params *Params) (*List[ProductionTask], *resty.Response, error)
	Create(ctx context.Context, productionTask *ProductionTask, params *Params) (*ProductionTask, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, productionTaskList []*ProductionTask, params *Params) (*[]ProductionTask, *resty.Response, error)
	DeleteMany(ctx context.Context, productionTaskList []*ProductionTask) (*DeleteManyResponse, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id *uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error)
	UpdateAttribute(ctx context.Context, id *uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList []*Attribute) (*DeleteManyResponse, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*ProductionTask, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, productionTask *ProductionTask, params *Params) (*ProductionTask, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetPositions(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[ProductionRow], *resty.Response, error)
	GetPositionByID(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, params *Params) (*ProductionRow, *resty.Response, error)
	UpdatePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, position *ProductionRow, params *Params) (*ProductionRow, *resty.Response, error)
	CreatePosition(ctx context.Context, id *uuid.UUID, position *ProductionRow) (*ProductionRow, *resty.Response, error)
	CreatePositions(ctx context.Context, id *uuid.UUID, positions []*ProductionRow) (*[]ProductionRow, *resty.Response, error)
	DeletePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID) (bool, *resty.Response, error)
	GetPositionTrackingCodes(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID) (*MetaArray[TrackingCode], *resty.Response, error)
	CreateOrUpdatePositionTrackingCodes(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, trackingCodes TrackingCodes) (*[]TrackingCode, *resty.Response, error)
	DeletePositionTrackingCodes(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, trackingCodes TrackingCodes) (*DeleteManyResponse, *resty.Response, error)
	GetProducts(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[ProductionTaskResult], *resty.Response, error)
	GetProductByID(ctx context.Context, id *uuid.UUID, productID *uuid.UUID, params *Params) (*ProductionTaskResult, *resty.Response, error)
	CreateProduct(ctx context.Context, id *uuid.UUID, productionTaskResult *ProductionTaskResult, params *Params) (*ProductionTaskResult, *resty.Response, error)
	UpdateProduct(ctx context.Context, id *uuid.UUID, productID *uuid.UUID, productionTaskResult *ProductionTaskResult, params *Params) (*ProductionTaskResult, *resty.Response, error)
	DeleteProduct(ctx context.Context, id *uuid.UUID, productID *uuid.UUID) (bool, *resty.Response, error)
	DeleteProductMany(ctx context.Context, id *uuid.UUID) (*DeleteManyResponse, *resty.Response, error)
}

type productionTaskService struct {
	Endpoint
	endpointGetList[ProductionTask]
	endpointCreate[ProductionTask]
	endpointCreateUpdateMany[ProductionTask]
	endpointDeleteMany[ProductionTask]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointAttributes
	endpointGetById[ProductionTask]
	endpointUpdate[ProductionTask]
	endpointDelete
	endpointPositions[ProductionRow]
}

func NewProductionTaskService(client *Client) ProductionTaskService {
	e := NewEndpoint(client, "entity/productiontask")
	return &productionTaskService{
		Endpoint:                 e,
		endpointGetList:          endpointGetList[ProductionTask]{e},
		endpointCreate:           endpointCreate[ProductionTask]{e},
		endpointCreateUpdateMany: endpointCreateUpdateMany[ProductionTask]{e},
		endpointDeleteMany:       endpointDeleteMany[ProductionTask]{e},
		endpointMetadata:         endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointGetById:          endpointGetById[ProductionTask]{e},
		endpointUpdate:           endpointUpdate[ProductionTask]{e},
		endpointPositions:        endpointPositions[ProductionRow]{e},
		endpointDelete:           endpointDelete{e},
		endpointAttributes:       endpointAttributes{e},
	}
}

// GetProducts Получить Продукты производственного задания.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-poluchit-produkty-proizwodstwennogo-zadaniq
func (s *productionTaskService) GetProducts(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[ProductionTaskResult], *resty.Response, error) {
	path := fmt.Sprintf("%s/products", id)
	return NewRequestBuilder[MetaArray[ProductionTaskResult]](s.client, path).SetParams(params).Get(ctx)
}

// GetProductByID Получить продукт производственного задания.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-produkt-proizwodstwennogo-zadaniq
func (s *productionTaskService) GetProductByID(ctx context.Context, id *uuid.UUID, productID *uuid.UUID, params *Params) (*ProductionTaskResult, *resty.Response, error) {
	path := fmt.Sprintf("%s/products/%s", id, productID)
	return NewRequestBuilder[ProductionTaskResult](s.client, path).SetParams(params).Get(ctx)
}

// CreateProduct Создать продукт.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-sozdat-produkt
func (s *productionTaskService) CreateProduct(ctx context.Context, id *uuid.UUID, productionTaskResult *ProductionTaskResult, params *Params) (*ProductionTaskResult, *resty.Response, error) {
	path := fmt.Sprintf("%s/products", id) // fixme:в документации указан endpoint без 's' на конце, вероятно ошибка
	return NewRequestBuilder[ProductionTaskResult](s.client, path).SetParams(params).Post(ctx, productionTaskResult)
}

// UpdateProduct Изменить продукт.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-izmenit-produkt
func (s *productionTaskService) UpdateProduct(ctx context.Context, id *uuid.UUID, productID *uuid.UUID, productionTaskResult *ProductionTaskResult, params *Params) (*ProductionTaskResult, *resty.Response, error) {
	path := fmt.Sprintf("%s/products/%s", id, productID)
	return NewRequestBuilder[ProductionTaskResult](s.client, path).SetParams(params).Put(ctx, productionTaskResult)
}

// DeleteProduct Удалить продукт.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-udalit-produkt
func (s *productionTaskService) DeleteProduct(ctx context.Context, id *uuid.UUID, productID *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/products/%s", id, productID)
	return NewRequestBuilder[any](s.client, path).Delete(ctx)
}

// DeleteProductMany Массовое удаление продуктов Производственного задания.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-udalit-produkt
func (s *productionTaskService) DeleteProductMany(ctx context.Context, id *uuid.UUID) (*DeleteManyResponse, *resty.Response, error) {
	path := fmt.Sprintf("%s/products/delete", id)
	return NewRequestBuilder[DeleteManyResponse](s.client, path).Post(ctx, nil)
}
