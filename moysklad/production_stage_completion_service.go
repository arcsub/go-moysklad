package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// ProductionStageCompletionService
// Сервис для работы с выполнениями этапов производства
type ProductionStageCompletionService struct {
	Endpoint
	endpointGetList[ProductionStageCompletion]
	endpointCreate[ProductionStageCompletion]
	endpointCreateUpdateDeleteMany[ProductionStageCompletion]
	endpointDelete
	endpointGetById[ProductionStageCompletion]
	endpointUpdate[ProductionStageCompletion]
}

func NewProductionStageCompletionService(client *Client) *ProductionStageCompletionService {
	e := NewEndpoint(client, "entity/productionstagecompletion")
	return &ProductionStageCompletionService{
		Endpoint:                       e,
		endpointGetList:                endpointGetList[ProductionStageCompletion]{e},
		endpointCreate:                 endpointCreate[ProductionStageCompletion]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[ProductionStageCompletion]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[ProductionStageCompletion]{e},
		endpointUpdate:                 endpointUpdate[ProductionStageCompletion]{e},
	}
}

// GetMaterials Получить Материалы выполнения этапа производства.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vypolnenie-atapa-proizwodstwa-poluchit-materialy-wypolneniq-atapa-proizwodstwa
func (s *ProductionStageCompletionService) GetMaterials(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[ProductionStageCompletionMaterial], *resty.Response, error) {
	path := fmt.Sprintf("%s/materials", id)
	return NewRequestBuilder[MetaArray[ProductionStageCompletionMaterial]](s.client, path).SetParams(params).Get(ctx)
}

// CreateMaterial Добавить Материал выполнения этапа производства.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vypolnenie-atapa-proizwodstwa-dobawit-material-wypolneniq-atapa-proizwodstwa
func (s *ProductionStageCompletionService) CreateMaterial(ctx context.Context, id *uuid.UUID, productionStageCompletionMaterial *ProductionStageCompletionMaterial, params *Params) (*ProductionStageCompletionMaterial, *resty.Response, error) {
	path := fmt.Sprintf("%s/materials", id)
	return NewRequestBuilder[ProductionStageCompletionMaterial](s.client, path).SetParams(params).Post(ctx, productionStageCompletionMaterial)
}

// UpdateMaterial Изменить Материал выполнения этапа производства.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vypolnenie-atapa-proizwodstwa-izmenit-material-wypolneniq-atapa-proizwodstwa
func (s *ProductionStageCompletionService) UpdateMaterial(ctx context.Context, id *uuid.UUID, materialID *uuid.UUID, productionStageCompletionMaterial *ProductionStageCompletionMaterial, params *Params) (*ProductionStageCompletionMaterial, *resty.Response, error) {
	path := fmt.Sprintf("%s/materials/%s", id, materialID)
	return NewRequestBuilder[ProductionStageCompletionMaterial](s.client, path).SetParams(params).Put(ctx, productionStageCompletionMaterial)
}

// GetProducts Получить Продукты выполнения этапа производства.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vypolnenie-atapa-proizwodstwa-poluchit-produkty-wypolneniq-atapa-proizwodstwa
func (s *ProductionStageCompletionService) GetProducts(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[ProductionStageCompletionResult], *resty.Response, error) {
	path := fmt.Sprintf("%s/products", id)
	return NewRequestBuilder[MetaArray[ProductionStageCompletionResult]](s.client, path).SetParams(params).Get(ctx)
}

// UpdateProduct Изменить Продукт выполнения этапа производства.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vypolnenie-atapa-proizwodstwa-izmenit-produkt-wypolneniq-atapa-proizwodstwa
func (s *ProductionStageCompletionService) UpdateProduct(ctx context.Context, id *uuid.UUID, productlID *uuid.UUID, productionStageCompletionResult *ProductionStageCompletionResult, params *Params) (*ProductionStageCompletionResult, *resty.Response, error) {
	path := fmt.Sprintf("%s/products/%s", id, productlID)
	return NewRequestBuilder[ProductionStageCompletionResult](s.client, path).SetParams(params).Put(ctx, productionStageCompletionResult)
}
