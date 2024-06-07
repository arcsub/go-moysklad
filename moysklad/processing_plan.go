package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// ProcessingPlan Техкарта.
// Ключевое слово: processingplan
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tehkarta-tehkarty
type ProcessingPlan struct {
	AccountID         *uuid.UUID                         `json:"accountId,omitempty"`         // ID учетной записи
	Archived          *bool                              `json:"archived,omitempty"`          // Добавлена ли Тех. карта в архив
	Code              *string                            `json:"code,omitempty"`              // Код Тех. карты
	Cost              *float64                           `json:"cost,omitempty"`              // Стоимость производства
	ExternalCode      *string                            `json:"externalCode,omitempty"`      // Внешний код
	Group             *Group                             `json:"group,omitempty"`             // Отдел сотрудника
	ID                *uuid.UUID                         `json:"id,omitempty"`                // ID сущности
	Stages            *MetaArray[ProcessingStage]        `json:"stages,omitempty"`            // Коллекция метаданных этапов Тех. карты
	Materials         *Positions[ProcessingPlanMaterial] `json:"materials,omitempty"`         // Список Метаданных материалов Тех. операции
	Meta              *Meta                              `json:"meta,omitempty"`              // Метаданные
	Name              *string                            `json:"name,omitempty"`              // Наименование
	Owner             *Employee                          `json:"owner,omitempty"`             // Владелец (Сотрудник)
	Parent            *Group                             `json:"parent,omitempty"`            // Метаданные группы Тех. карты
	PathName          *string                            `json:"pathName,omitempty"`          // Наименование группы, в которую входит Тех. карта
	ProcessingProcess *ProcessingProcess                 `json:"processingProcess,omitempty"` // Метаданные Тех. процесса
	Products          *Positions[ProcessingPlanProduct]  `json:"products,omitempty"`          // Коллекция метаданных готовых продуктов Тех. карты
	Shared            *bool                              `json:"shared,omitempty"`            // Общий доступ
	Updated           *Timestamp                         `json:"updated,omitempty"`           // Момент последнего обновления
}

func (p ProcessingPlan) String() string {
	return Stringify(p)
}

func (p ProcessingPlan) MetaType() MetaType {
	return MetaTypeProcessingPlan
}

// ProcessingPlanService
// Сервис для работы с тех картами.
type ProcessingPlanService interface {
	GetList(ctx context.Context, params *Params) (*List[ProcessingPlan], *resty.Response, error)
	Create(ctx context.Context, processingPlan *ProcessingPlan, params *Params) (*ProcessingPlan, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, processingPlanList []*ProcessingPlan, params *Params) (*[]ProcessingPlan, *resty.Response, error)
	DeleteMany(ctx context.Context, processingPlanList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*ProcessingPlan, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, processingPlan *ProcessingPlan, params *Params) (*ProcessingPlan, *resty.Response, error)
	GetPositions(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[ProcessingPlanProduct], *resty.Response, error)
	GetPositionByID(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, params *Params) (*ProcessingPlanProduct, *resty.Response, error)
	UpdatePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, position *ProcessingPlanProduct, params *Params) (*ProcessingPlanProduct, *resty.Response, error)
	CreatePosition(ctx context.Context, id *uuid.UUID, position *ProcessingPlanProduct) (*ProcessingPlanProduct, *resty.Response, error)
	CreatePositions(ctx context.Context, id *uuid.UUID, positions []*ProcessingPlanProduct) (*[]ProcessingPlanProduct, *resty.Response, error)
	DeletePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID) (bool, *resty.Response, error)
	GetPositionTrackingCodes(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID) (*MetaArray[TrackingCode], *resty.Response, error)
	CreateOrUpdatePositionTrackingCodes(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, trackingCodes TrackingCodes) (*[]TrackingCode, *resty.Response, error)
	DeletePositionTrackingCodes(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, trackingCodes TrackingCodes) (*DeleteManyResponse, *resty.Response, error)
	MoveToTrash(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetStages(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[ProcessingStage], *resty.Response, error)
	GetStageById(ctx context.Context, id, stageID *uuid.UUID) (*ProcessingStage, *resty.Response, error)
	UpdateStage(ctx context.Context, id, stageID *uuid.UUID, stage *ProcessingStage) (*ProcessingStage, *resty.Response, error)
	GetMaterials(ctx context.Context, id *uuid.UUID) (*List[ProcessingPlanMaterial], *resty.Response, error)
	CreateMaterial(ctx context.Context, id *uuid.UUID, material *ProcessingPlanMaterial) (*ProcessingPlanMaterial, *resty.Response, error)
	GetMaterialById(ctx context.Context, id, materialID *uuid.UUID) (*ProcessingPlanMaterial, *resty.Response, error)
	UpdateMaterial(ctx context.Context, id, materialID *uuid.UUID, material *ProcessingPlanMaterial) (*ProcessingPlanMaterial, *resty.Response, error)
	DeleteMaterial(ctx context.Context, id, materialID *uuid.UUID) (bool, *resty.Response, error)
	GetProducts(ctx context.Context, id *uuid.UUID) (*List[ProcessingPlanProduct], *resty.Response, error)
	CreateProduct(ctx context.Context, id *uuid.UUID, product *ProcessingPlanProduct) (*ProcessingPlanProduct, *resty.Response, error)
	GetProductById(ctx context.Context, id, productID *uuid.UUID) (*ProcessingPlanProduct, *resty.Response, error)
	UpdateProduct(ctx context.Context, id, productID *uuid.UUID, product *ProcessingPlanProduct) (*ProcessingPlanProduct, *resty.Response, error)
	DeleteProduct(ctx context.Context, id, productID *uuid.UUID) (bool, *resty.Response, error)
}

type processingPlanService struct {
	Endpoint
	endpointGetList[ProcessingPlan]
	endpointCreate[ProcessingPlan]
	endpointCreateUpdateMany[ProcessingPlan]
	endpointDeleteMany[ProcessingPlan]
	endpointDelete
	endpointGetById[ProcessingPlan]
	endpointUpdate[ProcessingPlan]
	endpointPositions[ProcessingPlanProduct]
	endpointRemove
}

func NewProcessingPlanService(client *Client) ProcessingPlanService {
	e := NewEndpoint(client, "entity/processingplan")
	return &processingPlanService{
		Endpoint:                 e,
		endpointGetList:          endpointGetList[ProcessingPlan]{e},
		endpointCreate:           endpointCreate[ProcessingPlan]{e},
		endpointCreateUpdateMany: endpointCreateUpdateMany[ProcessingPlan]{e},
		endpointDeleteMany:       endpointDeleteMany[ProcessingPlan]{e},
		endpointDelete:           endpointDelete{e},
		endpointGetById:          endpointGetById[ProcessingPlan]{e},
		endpointUpdate:           endpointUpdate[ProcessingPlan]{e},
		endpointPositions:        endpointPositions[ProcessingPlanProduct]{e},
		endpointRemove:           endpointRemove{e},
	}
}

func (s *processingPlanService) GetStages(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[ProcessingStage], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/stages", s.uri, id)
	return NewRequestBuilder[MetaArray[ProcessingStage]](s.client, path).SetParams(params).Get(ctx)
}

func (s *processingPlanService) GetStageById(ctx context.Context, id, stageID *uuid.UUID) (*ProcessingStage, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/stages/%s", s.uri, id, stageID)
	return NewRequestBuilder[ProcessingStage](s.client, path).Get(ctx)
}

func (s *processingPlanService) UpdateStage(ctx context.Context, id, stageID *uuid.UUID, stage *ProcessingStage) (*ProcessingStage, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/stages/%s", s.uri, id, stageID)
	return NewRequestBuilder[ProcessingStage](s.client, path).Put(ctx, stage)
}

func (s *processingPlanService) GetMaterials(ctx context.Context, id *uuid.UUID) (*List[ProcessingPlanMaterial], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/materials", s.uri, id)
	return NewRequestBuilder[List[ProcessingPlanMaterial]](s.client, path).Get(ctx)
}

func (s *processingPlanService) CreateMaterial(ctx context.Context, id *uuid.UUID, material *ProcessingPlanMaterial) (*ProcessingPlanMaterial, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/materials", s.uri, id)
	return NewRequestBuilder[ProcessingPlanMaterial](s.client, path).Post(ctx, material)
}

func (s *processingPlanService) GetMaterialById(ctx context.Context, id, materialID *uuid.UUID) (*ProcessingPlanMaterial, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/materials/%s", s.uri, id, materialID)
	return NewRequestBuilder[ProcessingPlanMaterial](s.client, path).Get(ctx)
}

func (s *processingPlanService) UpdateMaterial(ctx context.Context, id, materialID *uuid.UUID, material *ProcessingPlanMaterial) (*ProcessingPlanMaterial, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/materials/%s", s.uri, id, materialID)
	return NewRequestBuilder[ProcessingPlanMaterial](s.client, path).Put(ctx, material)
}

func (s *processingPlanService) DeleteMaterial(ctx context.Context, id, materialID *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/materials/%s", s.uri, id, materialID)
	return NewRequestBuilder[any](s.client, path).Delete(ctx)
}

func (s *processingPlanService) GetProducts(ctx context.Context, id *uuid.UUID) (*List[ProcessingPlanProduct], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/products", s.uri, id)
	return NewRequestBuilder[List[ProcessingPlanProduct]](s.client, path).Get(ctx)
}

func (s *processingPlanService) CreateProduct(ctx context.Context, id *uuid.UUID, product *ProcessingPlanProduct) (*ProcessingPlanProduct, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/products", s.uri, id)
	return NewRequestBuilder[ProcessingPlanProduct](s.client, path).Post(ctx, product)
}

func (s *processingPlanService) GetProductById(ctx context.Context, id, productID *uuid.UUID) (*ProcessingPlanProduct, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/products/%s", s.uri, id, productID)
	return NewRequestBuilder[ProcessingPlanProduct](s.client, path).Get(ctx)
}

func (s *processingPlanService) UpdateProduct(ctx context.Context, id, productID *uuid.UUID, product *ProcessingPlanProduct) (*ProcessingPlanProduct, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/products/%s", s.uri, id, productID)
	return NewRequestBuilder[ProcessingPlanProduct](s.client, path).Put(ctx, product)
}

func (s *processingPlanService) DeleteProduct(ctx context.Context, id, productID *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/products/%s", s.uri, id, productID)
	return NewRequestBuilder[any](s.client, path).Delete(ctx)
}
