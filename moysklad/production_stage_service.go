package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// ProductionStageService
// Сервис для работы с производственными этапами
type ProductionStageService interface {
	Update(ctx context.Context, id *uuid.UUID, productionStage *ProductionStage, params *Params) (*ProductionStage, *resty.Response, error)
	GetProductStages(ctx context.Context, productionTaskID *uuid.UUID, params *Params) (*MetaArray[ProductionStage], *resty.Response, error)
	GetMaterials(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[ProductionTaskMaterial], *resty.Response, error)
	CreateMaterial(ctx context.Context, id *uuid.UUID, productionTaskMaterial *ProductionTaskMaterial, params *Params) (*ProductionTaskMaterial, *resty.Response, error)
	UpdateMaterial(ctx context.Context, id *uuid.UUID, materialID *uuid.UUID, productionTaskMaterial *ProductionTaskMaterial, params *Params) (*ProductionTaskMaterial, *resty.Response, error)
	DeleteMaterial(ctx context.Context, id *uuid.UUID, materialID *uuid.UUID) (bool, *resty.Response, error)
}

type productionStageService struct {
	Endpoint
	endpointUpdate[ProductionStage]
}

func NewProductionStageService(client *Client) ProductionStageService {
	e := NewEndpoint(client, "entity/productionstage")
	return &productionStageService{
		Endpoint:       e,
		endpointUpdate: endpointUpdate[ProductionStage]{e},
	}
}

// GetProductStages Получить список Производственных этапов Производственного задания.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-poluchit-spisok-proizwodstwennyh-atapow-proizwodstwennogo-zadaniq
func (s *productionStageService) GetProductStages(ctx context.Context, productionTaskID *uuid.UUID, params *Params) (*MetaArray[ProductionStage], *resty.Response, error) {
	ptURL := fmt.Sprintf("https://api.moysklad.ru/api/remap/1.2/entity/productiontask/%s", productionTaskID)
	params.WithFilterEquals("productionTask", ptURL)
	return NewRequestBuilder[MetaArray[ProductionStage]](s.client, s.uri).SetParams(params).Get(ctx)
}

// GetMaterials Получить Материалы производственного этапа.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-poluchit-materialy-proizwodstwennogo-atapa
func (s *productionStageService) GetMaterials(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[ProductionTaskMaterial], *resty.Response, error) {
	path := fmt.Sprintf("%s/materials", id)
	return NewRequestBuilder[MetaArray[ProductionTaskMaterial]](s.client, path).SetParams(params).Get(ctx)
}

// CreateMaterial Добавить Материал к производственному этапу.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-dobawit-material-k-proizwodstwennomu-atapu
func (s *productionStageService) CreateMaterial(ctx context.Context, id *uuid.UUID, productionTaskMaterial *ProductionTaskMaterial, params *Params) (*ProductionTaskMaterial, *resty.Response, error) {
	path := fmt.Sprintf("%s/materials", id)
	return NewRequestBuilder[ProductionTaskMaterial](s.client, path).SetParams(params).Post(ctx, productionTaskMaterial)
}

// UpdateMaterial Изменить Материал производственного этапа.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-izmenit-material-proizwodstwennogo-atapa
func (s *productionStageService) UpdateMaterial(ctx context.Context, id *uuid.UUID, materialID *uuid.UUID, productionTaskMaterial *ProductionTaskMaterial, params *Params) (*ProductionTaskMaterial, *resty.Response, error) {
	path := fmt.Sprintf("%s/materials/%s", id, materialID)
	return NewRequestBuilder[ProductionTaskMaterial](s.client, path).SetParams(params).Put(ctx, productionTaskMaterial)
}

// DeleteMaterial Удалить Материал производственного этапа.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-udalit-material-proizwodstwennogo-atapa
func (s *productionStageService) DeleteMaterial(ctx context.Context, id *uuid.UUID, materialID *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/materials/%s", id, materialID)
	return NewRequestBuilder[any](s.client, path).Delete(ctx)
}
