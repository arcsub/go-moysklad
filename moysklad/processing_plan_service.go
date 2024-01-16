package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// ProcessingPlanService
// Сервис для работы с тех картами.
type ProcessingPlanService struct {
	Endpoint
	endpointGetList[ProcessingPlan]
	endpointCreate[ProcessingPlan]
	endpointCreateUpdateDeleteMany[ProcessingPlan]
	endpointDelete
	endpointGetById[ProcessingPlan]
	endpointUpdate[ProcessingPlan]
	endpointPositions[ProcessingPlanProduct]
	endpointRemove
}

func NewProcessingPlanService(client *Client) *ProcessingPlanService {
	e := NewEndpoint(client, "entity/processingplan")
	return &ProcessingPlanService{
		Endpoint:                       e,
		endpointGetList:                endpointGetList[ProcessingPlan]{e},
		endpointCreate:                 endpointCreate[ProcessingPlan]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[ProcessingPlan]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[ProcessingPlan]{e},
		endpointUpdate:                 endpointUpdate[ProcessingPlan]{e},
		endpointPositions:              endpointPositions[ProcessingPlanProduct]{e},
		endpointRemove:                 endpointRemove{e},
	}
}

func (s *ProcessingPlanService) GetStages(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[ProcessingStage], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/stages", s.uri, id)
	return NewRequestBuilder[MetaArray[ProcessingStage]](s.client, path).SetParams(params).Get(ctx)
}

func (s *ProcessingPlanService) GetStageById(ctx context.Context, id, stageId uuid.UUID) (*ProcessingStage, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/stages/%s", s.uri, id, stageId)
	return NewRequestBuilder[ProcessingStage](s.client, path).Get(ctx)
}

func (s *ProcessingPlanService) UpdateStage(ctx context.Context, id, stageId uuid.UUID, stage *ProcessingStage) (*ProcessingStage, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/stages/%s", s.uri, id, stageId)
	return NewRequestBuilder[ProcessingStage](s.client, path).Put(ctx, stage)
}

func (s *ProcessingPlanService) GetMaterials(ctx context.Context, id *uuid.UUID) (*List[ProcessingPlanMaterial], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/materials", s.uri, id)
	return NewRequestBuilder[List[ProcessingPlanMaterial]](s.client, path).Get(ctx)
}

func (s *ProcessingPlanService) CreateMaterial(ctx context.Context, id *uuid.UUID, material *ProcessingPlanMaterial) (*ProcessingPlanMaterial, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/materials", s.uri, id)
	return NewRequestBuilder[ProcessingPlanMaterial](s.client, path).Post(ctx, material)
}

func (s *ProcessingPlanService) GetMaterialById(ctx context.Context, id, materialId uuid.UUID) (*ProcessingPlanMaterial, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/materials/%s", s.uri, id, materialId)
	return NewRequestBuilder[ProcessingPlanMaterial](s.client, path).Get(ctx)
}

func (s *ProcessingPlanService) UpdateMaterial(ctx context.Context, id, materialId uuid.UUID, material *ProcessingPlanMaterial) (*ProcessingPlanMaterial, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/materials/%s", s.uri, id, materialId)
	return NewRequestBuilder[ProcessingPlanMaterial](s.client, path).Put(ctx, material)
}

func (s *ProcessingPlanService) DeleteMaterial(ctx context.Context, id, materialId uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/materials/%s", s.uri, id, materialId)
	return NewRequestBuilder[any](s.client, path).Delete(ctx)
}

func (s *ProcessingPlanService) GetProducts(ctx context.Context, id *uuid.UUID) (*List[ProcessingPlanProduct], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/products", s.uri, id)
	return NewRequestBuilder[List[ProcessingPlanProduct]](s.client, path).Get(ctx)
}

func (s *ProcessingPlanService) CreateProduct(ctx context.Context, id *uuid.UUID, product *ProcessingPlanProduct) (*ProcessingPlanProduct, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/products", s.uri, id)
	return NewRequestBuilder[ProcessingPlanProduct](s.client, path).Post(ctx, product)
}

func (s *ProcessingPlanService) GetProductById(ctx context.Context, id, productId uuid.UUID) (*ProcessingPlanProduct, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/products/%s", s.uri, id, productId)
	return NewRequestBuilder[ProcessingPlanProduct](s.client, path).Get(ctx)
}

func (s *ProcessingPlanService) UpdateProduct(ctx context.Context, id, productId uuid.UUID, product *ProcessingPlanProduct) (*ProcessingPlanProduct, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/products/%s", s.uri, id, productId)
	return NewRequestBuilder[ProcessingPlanProduct](s.client, path).Put(ctx, product)
}

func (s *ProcessingPlanService) DeleteProduct(ctx context.Context, id, productId uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/products/%s", s.uri, id, productId)
	return NewRequestBuilder[any](s.client, path).Delete(ctx)
}
