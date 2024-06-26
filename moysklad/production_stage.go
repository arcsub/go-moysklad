package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// ProductionStage Производственный этап
// Ключевое слово: productionstage
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-proizwodstwennye-atapy
type ProductionStage struct {
	AccountID          *uuid.UUID                         `json:"accountId,omitempty"`          // ID учетной записи
	ID                 *uuid.UUID                         `json:"id,omitempty"`                 // ID Производственного этапа
	Meta               *Meta                              `json:"meta,omitempty"`               // Метаданные Производственного этапа
	LabourUnitCost     *float64                           `json:"labourUnitCost,omitempty"`     // Затраты на оплату труда за единицу объема производства
	Materials          *Positions[ProductionTaskMaterial] `json:"materials,omitempty"`          // Метаданные Материалов производственного этапа
	OrderingPosition   *int                               `json:"orderingPosition,omitempty"`   // Индекс Производственного этапа в Позиции производственного задания
	Stage              *ProductionStage                   `json:"stage,omitempty"`              // Метаданные Этапа производства
	ProductionRow      *ProductionRow                     `json:"productionRow,omitempty"`      // Метаданные Позиции производственного задания
	TotalQuantity      *float64                           `json:"totalQuantity,omitempty"`      // Объем Производственного этапа. Соответствует объему Позиции производственного задания
	CompletedQuantity  *float64                           `json:"completedQuantity,omitempty"`  // Выполненное количество
	AvailableQuantity  *float64                           `json:"availableQuantity,omitempty"`  // Количество, доступное к выполнению
	BlockedQuantity    *float64                           `json:"blockedQuantity,omitempty"`    // Количество, которое на данный момент выполнять нельзя. Например, ещё не выполнен предыдущий этап
	SkippedQuantity    *float64                           `json:"skippedQuantity,omitempty"`    // Количество, которое не будет выполнено. Например, из-за остановки производства
	ProcessingUnitCost *float64                           `json:"processingUnitCost,omitempty"` // Затраты на единицу объема производства
	StandardHourUnit   *float64                           `json:"standardHourUnit,omitempty"`   // Нормо-часы единицы объема производства
}

// Clean возвращает сущность с единственным заполненным полем Meta
func (productionStage ProductionStage) Clean() *ProductionStage {
	return &ProductionStage{Meta: productionStage.Meta}
}

func (productionStage ProductionStage) GetAccountID() uuid.UUID {
	return Deref(productionStage.AccountID)
}

func (productionStage ProductionStage) GetID() uuid.UUID {
	return Deref(productionStage.ID)
}

func (productionStage ProductionStage) GetMeta() Meta {
	return Deref(productionStage.Meta)
}

func (productionStage ProductionStage) GetLabourUnitCost() float64 {
	return Deref(productionStage.LabourUnitCost)
}

func (productionStage ProductionStage) GetMaterials() Positions[ProductionTaskMaterial] {
	return Deref(productionStage.Materials)
}

func (productionStage ProductionStage) GetOrderingPosition() int {
	return Deref(productionStage.OrderingPosition)
}

func (productionStage ProductionStage) GetStage() ProductionStage {
	return Deref(productionStage.Stage)
}

func (productionStage ProductionStage) GetProductionRow() ProductionRow {
	return Deref(productionStage.ProductionRow)
}

func (productionStage ProductionStage) GetTotalQuantity() float64 {
	return Deref(productionStage.TotalQuantity)
}

func (productionStage ProductionStage) GetCompletedQuantity() float64 {
	return Deref(productionStage.CompletedQuantity)
}

func (productionStage ProductionStage) GetAvailableQuantity() float64 {
	return Deref(productionStage.AvailableQuantity)
}

func (productionStage ProductionStage) GetBlockedQuantity() float64 {
	return Deref(productionStage.BlockedQuantity)
}

func (productionStage ProductionStage) GetSkippedQuantity() float64 {
	return Deref(productionStage.SkippedQuantity)
}

func (productionStage ProductionStage) GetProcessingUnitCost() float64 {
	return Deref(productionStage.ProcessingUnitCost)
}

func (productionStage ProductionStage) GetStandardHourUnit() float64 {
	return Deref(productionStage.StandardHourUnit)
}

func (productionStage *ProductionStage) SetMeta(meta *Meta) *ProductionStage {
	productionStage.Meta = meta
	return productionStage
}

func (productionStage *ProductionStage) SetLabourUnitCost(labourUnitCost float64) *ProductionStage {
	productionStage.LabourUnitCost = &labourUnitCost
	return productionStage
}

func (productionStage *ProductionStage) SetMaterials(materials ...*ProductionTaskMaterial) *ProductionStage {
	productionStage.Materials = NewPositionsFrom(materials)
	return productionStage
}

func (productionStage *ProductionStage) SetProcessingUnitCost(processingUnitCost float64) *ProductionStage {
	productionStage.ProcessingUnitCost = &processingUnitCost
	return productionStage
}

func (productionStage *ProductionStage) SetStandardHourUnit(standardHourUnit float64) *ProductionStage {
	productionStage.StandardHourUnit = &standardHourUnit
	return productionStage
}

func (productionStage ProductionStage) String() string {
	return Stringify(productionStage)
}

// MetaType возвращает тип сущности.
func (ProductionStage) MetaType() MetaType {
	return MetaTypeProductionStage
}

// ProductionTaskMaterial Материал Производственного этапа
// Ключевое слово: productiontaskmaterial
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-proizwodstwennye-atapy-materialy-proizwodstwennogo-atapa
type ProductionTaskMaterial struct {
	AccountID    *uuid.UUID          `json:"accountId,omitempty"`    // ID учетной записи
	Assortment   *AssortmentPosition `json:"assortment,omitempty"`   // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	ID           *uuid.UUID          `json:"id,omitempty"`           // ID позиции
	PlanQuantity *float64            `json:"planQuantity,omitempty"` // Количество товаров/модификаций данного вида в позиции
}

func (productionTaskMaterial ProductionTaskMaterial) GetAccountID() uuid.UUID {
	return Deref(productionTaskMaterial.AccountID)
}

func (productionTaskMaterial ProductionTaskMaterial) GetAssortment() AssortmentPosition {
	return Deref(productionTaskMaterial.Assortment)
}

func (productionTaskMaterial ProductionTaskMaterial) GetID() uuid.UUID {
	return Deref(productionTaskMaterial.ID)
}

func (productionTaskMaterial ProductionTaskMaterial) GetPlanQuantity() float64 {
	return Deref(productionTaskMaterial.PlanQuantity)
}

func (productionTaskMaterial *ProductionTaskMaterial) SetAssortment(assortment AsAssortment) *ProductionTaskMaterial {
	productionTaskMaterial.Assortment = assortment.AsAssortment()
	return productionTaskMaterial
}

func (productionTaskMaterial *ProductionTaskMaterial) SetPlanQuantity(planQuantity float64) *ProductionTaskMaterial {
	productionTaskMaterial.PlanQuantity = &planQuantity
	return productionTaskMaterial
}

func (productionTaskMaterial ProductionTaskMaterial) String() string {
	return Stringify(productionTaskMaterial)
}

// MetaType возвращает тип сущности.
func (ProductionTaskMaterial) MetaType() MetaType {
	return MetaTypeProductionTaskMaterial
}

// ProductionStageService
// Сервис для работы с производственными этапами
type ProductionStageService interface {
	Update(ctx context.Context, id uuid.UUID, productionStage *ProductionStage, params ...*Params) (*ProductionStage, *resty.Response, error)
	GetProductStages(ctx context.Context, productionTaskID uuid.UUID, params ...*Params) (*MetaArray[ProductionStage], *resty.Response, error)
	GetMaterials(ctx context.Context, id uuid.UUID, params ...*Params) (*MetaArray[ProductionTaskMaterial], *resty.Response, error)
	CreateMaterial(ctx context.Context, id uuid.UUID, productionTaskMaterial *ProductionTaskMaterial, params ...*Params) (*ProductionTaskMaterial, *resty.Response, error)
	UpdateMaterial(ctx context.Context, id uuid.UUID, materialID uuid.UUID, productionTaskMaterial *ProductionTaskMaterial, params ...*Params) (*ProductionTaskMaterial, *resty.Response, error)
	DeleteMaterial(ctx context.Context, id uuid.UUID, materialID uuid.UUID) (bool, *resty.Response, error)
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
func (service *productionStageService) GetProductStages(ctx context.Context, productionTaskID uuid.UUID, params ...*Params) (*MetaArray[ProductionStage], *resty.Response, error) {
	ptURL := fmt.Sprintf("https://api.moysklad.ru/api/remap/1.2/entity/productiontask/%s", productionTaskID)
	var param *Params
	if len(params) > 0 {
		param = params[0]
	} else {
		param = new(Params)
	}
	param.WithFilterEquals("productionTask", ptURL)
	return NewRequestBuilder[MetaArray[ProductionStage]](service.client, service.uri).SetParams(param).Get(ctx)
}

// GetMaterials Получить Материалы производственного этапа.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-poluchit-materialy-proizwodstwennogo-atapa
func (service *productionStageService) GetMaterials(ctx context.Context, id uuid.UUID, params ...*Params) (*MetaArray[ProductionTaskMaterial], *resty.Response, error) {
	path := fmt.Sprintf("%s/materials", id)
	return NewRequestBuilder[MetaArray[ProductionTaskMaterial]](service.client, path).SetParams(params...).Get(ctx)
}

// CreateMaterial Добавить Материал к производственному этапу.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-dobawit-material-k-proizwodstwennomu-atapu
func (service *productionStageService) CreateMaterial(ctx context.Context, id uuid.UUID, productionTaskMaterial *ProductionTaskMaterial, params ...*Params) (*ProductionTaskMaterial, *resty.Response, error) {
	path := fmt.Sprintf("%s/materials", id)
	return NewRequestBuilder[ProductionTaskMaterial](service.client, path).SetParams(params...).Post(ctx, productionTaskMaterial)
}

// UpdateMaterial Изменить Материал производственного этапа.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-izmenit-material-proizwodstwennogo-atapa
func (service *productionStageService) UpdateMaterial(ctx context.Context, id uuid.UUID, materialID uuid.UUID, productionTaskMaterial *ProductionTaskMaterial, params ...*Params) (*ProductionTaskMaterial, *resty.Response, error) {
	path := fmt.Sprintf("%s/materials/%s", id, materialID)
	return NewRequestBuilder[ProductionTaskMaterial](service.client, path).SetParams(params...).Put(ctx, productionTaskMaterial)
}

// DeleteMaterial Удалить Материал производственного этапа.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-udalit-material-proizwodstwennogo-atapa
func (service *productionStageService) DeleteMaterial(ctx context.Context, id uuid.UUID, materialID uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/materials/%s", id, materialID)
	return NewRequestBuilder[any](service.client, path).Delete(ctx)
}
