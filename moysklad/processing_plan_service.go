package moysklad

import (
	"context"
	"fmt"
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

func (s *ProcessingPlanService) GetStages(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[ProcessingStage], *Response, error) {
	path := fmt.Sprintf("%s/stages", id)
	return NewRequestBuilder[MetaArray[ProcessingStage]](s.Endpoint, ctx).WithPath(path).WithParams(params).Get()
}

func (s *ProcessingPlanService) GetStageById(ctx context.Context, id, stageId uuid.UUID) (*ProcessingStage, *Response, error) {
	path := fmt.Sprintf("%s/stages/%s", id, stageId)
	return NewRequestBuilder[ProcessingStage](s.Endpoint, ctx).WithPath(path).Get()
}

func (s *ProcessingPlanService) UpdateStage(ctx context.Context, id, stageId uuid.UUID, stage *ProcessingStage) (*ProcessingStage, *Response, error) {
	path := fmt.Sprintf("%s/stages/%s", id, stageId)
	return NewRequestBuilder[ProcessingStage](s.Endpoint, ctx).WithPath(path).Put()
}

func (s *ProcessingPlanService) GetMaterials(ctx context.Context, id *uuid.UUID) (*List[ProcessingPlanMaterial], *Response, error) {
	path := fmt.Sprintf("%s/materials", id)
	return NewRequestBuilder[List[ProcessingPlanMaterial]](s.Endpoint, ctx).WithPath(path).Put()
}

func (s *ProcessingPlanService) CreateMaterial(ctx context.Context, id *uuid.UUID, material *ProcessingPlanMaterial) (*ProcessingPlanMaterial, *Response, error) {
	path := fmt.Sprintf("%s/materials", id)
	return NewRequestBuilder[ProcessingPlanMaterial](s.Endpoint, ctx).WithPath(path).WithBody(material).Post()
}

func (s *ProcessingPlanService) GetMaterialById(ctx context.Context, id, materialId uuid.UUID) (*ProcessingPlanMaterial, *Response, error) {
	path := fmt.Sprintf("%s/materials/%s", id, materialId)
	return NewRequestBuilder[ProcessingPlanMaterial](s.Endpoint, ctx).WithPath(path).Get()
}

func (s *ProcessingPlanService) UpdateMaterial(ctx context.Context, id, materialId uuid.UUID, material *ProcessingPlanMaterial) (*ProcessingPlanMaterial, *Response, error) {
	path := fmt.Sprintf("%s/materials/%s", id, materialId)
	return NewRequestBuilder[ProcessingPlanMaterial](s.Endpoint, ctx).WithPath(path).WithBody(material).Put()
}

func (s *ProcessingPlanService) DeleteMaterial(ctx context.Context, id, materialId uuid.UUID) (bool, *Response, error) {
	path := fmt.Sprintf("%s/materials/%s", id, materialId)
	return NewRequestBuilder[any](s.Endpoint, ctx).WithPath(path).Delete()
}

func (s *ProcessingPlanService) GetProducts(ctx context.Context, id *uuid.UUID) (*List[ProcessingPlanProduct], *Response, error) {
	path := fmt.Sprintf("%s/products", id)
	return NewRequestBuilder[List[ProcessingPlanProduct]](s.Endpoint, ctx).WithPath(path).Get()
}

func (s *ProcessingPlanService) CreateProduct(ctx context.Context, id *uuid.UUID, product *ProcessingPlanProduct) (*ProcessingPlanProduct, *Response, error) {
	path := fmt.Sprintf("%s/products", id)
	return NewRequestBuilder[ProcessingPlanProduct](s.Endpoint, ctx).WithPath(path).WithBody(product).Post()
}

func (s *ProcessingPlanService) GetProductById(ctx context.Context, id, productId uuid.UUID) (*ProcessingPlanProduct, *Response, error) {
	path := fmt.Sprintf("%s/products/%s", id, productId)
	return NewRequestBuilder[ProcessingPlanProduct](s.Endpoint, ctx).WithPath(path).Get()
}

func (s *ProcessingPlanService) UpdateProduct(ctx context.Context, id, productId uuid.UUID, product *ProcessingPlanProduct) (*ProcessingPlanProduct, *Response, error) {
	path := fmt.Sprintf("%s/products/%s", id, productId)
	return NewRequestBuilder[ProcessingPlanProduct](s.Endpoint, ctx).WithPath(path).WithBody(product).Put()
}

func (s *ProcessingPlanService) DeleteProduct(ctx context.Context, id, productId uuid.UUID) (bool, *Response, error) {
	path := fmt.Sprintf("%s/products/%s", id, productId)
	return NewRequestBuilder[any](s.Endpoint, ctx).WithPath(path).Delete()
}
