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
	AccountID            *uuid.UUID                         `json:"accountId,omitempty"`            // ID учетной записи
	Archived             *bool                              `json:"archived,omitempty"`             // Добавлена ли Тех. карта в архив
	Code                 *string                            `json:"code,omitempty"`                 // Код Тех. карты
	Cost                 *float64                           `json:"cost,omitempty"`                 // Стоимость производства
	ExternalCode         *string                            `json:"externalCode,omitempty"`         // Внешний код
	Group                *Group                             `json:"group,omitempty"`                // Отдел сотрудника
	ID                   *uuid.UUID                         `json:"id,omitempty"`                   // ID сущности
	Stages               *MetaArray[ProcessingStage]        `json:"stages,omitempty"`               // Коллекция метаданных этапов Тех. карты
	Materials            *Positions[ProcessingPlanMaterial] `json:"materials,omitempty"`            // Список Метаданных материалов Тех. операции
	Meta                 *Meta                              `json:"meta,omitempty"`                 // Метаданные
	Name                 *string                            `json:"name,omitempty"`                 // Наименование
	Owner                *Employee                          `json:"owner,omitempty"`                // Владелец (Сотрудник)
	Parent               *Group                             `json:"parent,omitempty"`               // Метаданные группы Тех. карты
	PathName             *string                            `json:"pathName,omitempty"`             // Наименование группы, в которую входит Тех. карта
	ProcessingProcess    *ProcessingProcess                 `json:"processingProcess,omitempty"`    // Метаданные Тех. процесса
	Products             *Positions[ProcessingPlanProduct]  `json:"products,omitempty"`             // Коллекция метаданных готовых продуктов Тех. карты
	Shared               *bool                              `json:"shared,omitempty"`               // Общий доступ
	Updated              *Timestamp                         `json:"updated,omitempty"`              // Момент последнего обновления
	CostDistributionType CostDistributionType               `json:"costDistributionType,omitempty"` // Тип распределения себестоимости. Возможные значения: BY_PRICE, BY_PRODUCTION
}

func (processingPlan ProcessingPlan) Clean() *ProcessingPlan {
	return &ProcessingPlan{Meta: processingPlan.Meta}
}

func (processingPlan ProcessingPlan) GetAccountID() uuid.UUID {
	return Deref(processingPlan.AccountID)
}

func (processingPlan ProcessingPlan) GetArchived() bool {
	return Deref(processingPlan.Archived)
}

func (processingPlan ProcessingPlan) GetCode() string {
	return Deref(processingPlan.Code)
}

func (processingPlan ProcessingPlan) GetCost() float64 {
	return Deref(processingPlan.Cost)
}

func (processingPlan ProcessingPlan) GetCostDistributionType() CostDistributionType {
	return processingPlan.CostDistributionType
}

func (processingPlan ProcessingPlan) GetExternalCode() string {
	return Deref(processingPlan.ExternalCode)
}

func (processingPlan ProcessingPlan) GetGroup() Group {
	return Deref(processingPlan.Group)
}

func (processingPlan ProcessingPlan) GetID() uuid.UUID {
	return Deref(processingPlan.ID)
}

func (processingPlan ProcessingPlan) GetStages() MetaArray[ProcessingStage] {
	return Deref(processingPlan.Stages)
}

func (processingPlan ProcessingPlan) GetMaterials() Positions[ProcessingPlanMaterial] {
	return Deref(processingPlan.Materials)
}

func (processingPlan ProcessingPlan) GetMeta() Meta {
	return Deref(processingPlan.Meta)
}

func (processingPlan ProcessingPlan) GetName() string {
	return Deref(processingPlan.Name)
}

func (processingPlan ProcessingPlan) GetOwner() Employee {
	return Deref(processingPlan.Owner)
}

func (processingPlan ProcessingPlan) GetParent() Group {
	return Deref(processingPlan.Parent)
}

func (processingPlan ProcessingPlan) GetPathName() string {
	return Deref(processingPlan.PathName)
}

func (processingPlan ProcessingPlan) GetProcessingProcess() ProcessingProcess {
	return Deref(processingPlan.ProcessingProcess)
}

func (processingPlan ProcessingPlan) GetProducts() Positions[ProcessingPlanProduct] {
	return Deref(processingPlan.Products)
}

func (processingPlan ProcessingPlan) GetShared() bool {
	return Deref(processingPlan.Shared)
}

func (processingPlan ProcessingPlan) GetUpdated() Timestamp {
	return Deref(processingPlan.Updated)
}

func (processingPlan *ProcessingPlan) SetArchived(archived bool) *ProcessingPlan {
	processingPlan.Archived = &archived
	return processingPlan
}

func (processingPlan *ProcessingPlan) SetCode(code string) *ProcessingPlan {
	processingPlan.Code = &code
	return processingPlan
}

func (processingPlan *ProcessingPlan) SetCost(cost float64) *ProcessingPlan {
	processingPlan.Cost = &cost
	return processingPlan
}

func (processingPlan *ProcessingPlan) SetExternalCode(externalCode string) *ProcessingPlan {
	processingPlan.ExternalCode = &externalCode
	return processingPlan
}

func (processingPlan *ProcessingPlan) SetGroup(group *Group) *ProcessingPlan {
	processingPlan.Group = group.Clean()
	return processingPlan
}

func (processingPlan *ProcessingPlan) SetStages(stages Slice[ProcessingStage]) *ProcessingPlan {
	processingPlan.Stages = NewMetaArrayRows(stages)
	return processingPlan
}

func (processingPlan *ProcessingPlan) SetMaterials(materials *Positions[ProcessingPlanMaterial]) *ProcessingPlan {
	processingPlan.Materials = materials
	return processingPlan
}

func (processingPlan *ProcessingPlan) SetMeta(meta *Meta) *ProcessingPlan {
	processingPlan.Meta = meta
	return processingPlan
}

func (processingPlan *ProcessingPlan) SetName(name string) *ProcessingPlan {
	processingPlan.Name = &name
	return processingPlan
}

func (processingPlan *ProcessingPlan) SetOwner(owner *Employee) *ProcessingPlan {
	processingPlan.Owner = owner.Clean()
	return processingPlan
}

func (processingPlan *ProcessingPlan) SetParent(parent *Group) *ProcessingPlan {
	processingPlan.Parent = parent.Clean()
	return processingPlan
}

func (processingPlan *ProcessingPlan) SetProcessingProcess(processingProcess *ProcessingProcess) *ProcessingPlan {
	processingPlan.ProcessingProcess = processingProcess.Clean()
	return processingPlan
}

func (processingPlan *ProcessingPlan) SetProducts(products *Positions[ProcessingPlanProduct]) *ProcessingPlan {
	processingPlan.Products = products
	return processingPlan
}

func (processingPlan *ProcessingPlan) SetShared(shared bool) *ProcessingPlan {
	processingPlan.Shared = &shared
	return processingPlan
}

func (processingPlan ProcessingPlan) String() string {
	return Stringify(processingPlan)
}

func (processingPlan ProcessingPlan) MetaType() MetaType {
	return MetaTypeProcessingPlan
}

// ProcessingPlanProduct Продукт Тех. карты.
// Ключевое слово: processingplanresult
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tehkarta-tehkarty-produkty-tehkarty
type ProcessingPlanProduct struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара или модификации позиции
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Product    *Product            `json:"product,omitempty"`    // Метаданные товара позиции. В случае, если в поле assortment указана модификация, то это поле содержит товар, к которому относится эта модификация
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров данного вида в позиции
}

func (processingPlanProduct ProcessingPlanProduct) GetAccountID() uuid.UUID {
	return Deref(processingPlanProduct.AccountID)
}

func (processingPlanProduct ProcessingPlanProduct) GetAssortment() AssortmentPosition {
	return Deref(processingPlanProduct.Assortment)
}

func (processingPlanProduct ProcessingPlanProduct) GetID() uuid.UUID {
	return Deref(processingPlanProduct.ID)
}

func (processingPlanProduct ProcessingPlanProduct) GetProduct() Product {
	return Deref(processingPlanProduct.Product)
}

func (processingPlanProduct ProcessingPlanProduct) GetQuantity() float64 {
	return Deref(processingPlanProduct.Quantity)
}

func (processingPlanProduct *ProcessingPlanProduct) SetAssortment(assortment AsAssortment) *ProcessingPlanProduct {
	processingPlanProduct.Assortment = assortment.AsAssortment()
	return processingPlanProduct
}

func (processingPlanProduct *ProcessingPlanProduct) SetProduct(product *Product) *ProcessingPlanProduct {
	processingPlanProduct.Product = product.Clean()
	return processingPlanProduct
}

func (processingPlanProduct *ProcessingPlanProduct) SetQuantity(quantity float64) *ProcessingPlanProduct {
	processingPlanProduct.Quantity = &quantity
	return processingPlanProduct
}

func (processingPlanProduct ProcessingPlanProduct) String() string {
	return Stringify(processingPlanProduct)
}

func (processingPlanProduct ProcessingPlanProduct) MetaType() MetaType {
	return MetaTypeProcessingPlanProduct
}

// ProcessingPlanMaterial Материал Тех. карты.
// Ключевое слово: processingplanmaterial
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tehkarta-tehkarty-materialy-tehkarty
type ProcessingPlanMaterial struct {
	AccountID                 *uuid.UUID          `json:"accountId,omitempty"`                 // ID учетной записи
	Assortment                *AssortmentPosition `json:"assortment,omitempty"`                // Метаданные товара или модификации позиции
	ID                        *uuid.UUID          `json:"id,omitempty"`                        // ID позиции
	Product                   *Product            `json:"product,omitempty"`                   // Метаданные товара позиции. В случае, если в поле assortment указана модификация, то это поле содержит товар, к которому относится эта модификация
	Quantity                  *float64            `json:"quantity,omitempty"`                  // Количество товаров данного вида в позиции
	ProcessingProcessPosition *Meta               `json:"processingProcessPosition,omitempty"` // Метаданные позиции Тех. процесса
	MaterialProcessingPlan    *Meta               `json:"materialProcessingPlan"`              // Метаданные техкарты материала [11-01-2024]
}

func (processingPlanMaterial ProcessingPlanMaterial) GetAccountID() uuid.UUID {
	return Deref(processingPlanMaterial.AccountID)
}

func (processingPlanMaterial ProcessingPlanMaterial) GetAssortment() AssortmentPosition {
	return Deref(processingPlanMaterial.Assortment)
}

func (processingPlanMaterial ProcessingPlanMaterial) GetID() uuid.UUID {
	return Deref(processingPlanMaterial.ID)
}

func (processingPlanMaterial ProcessingPlanMaterial) GetProduct() Product {
	return Deref(processingPlanMaterial.Product)
}

func (processingPlanMaterial ProcessingPlanMaterial) GetQuantity() float64 {
	return Deref(processingPlanMaterial.Quantity)
}

func (processingPlanMaterial ProcessingPlanMaterial) GetProcessingProcessPosition() Meta {
	return Deref(processingPlanMaterial.ProcessingProcessPosition)
}

func (processingPlanMaterial ProcessingPlanMaterial) GetMaterialProcessingPlan() Meta {
	return Deref(processingPlanMaterial.MaterialProcessingPlan)
}

func (processingPlanMaterial *ProcessingPlanMaterial) SetAssortment(assortment AsAssortment) *ProcessingPlanMaterial {
	processingPlanMaterial.Assortment = assortment.AsAssortment()
	return processingPlanMaterial
}

func (processingPlanMaterial *ProcessingPlanMaterial) SetProduct(product *Product) *ProcessingPlanMaterial {
	processingPlanMaterial.Product = product.Clean()
	return processingPlanMaterial
}

func (processingPlanMaterial *ProcessingPlanMaterial) SetQuantity(quantity float64) *ProcessingPlanMaterial {
	processingPlanMaterial.Quantity = &quantity
	return processingPlanMaterial
}

func (processingPlanMaterial *ProcessingPlanMaterial) SetProcessingProcessPosition(processingProcessPosition *ProcessingProcessPosition) *ProcessingPlanMaterial {
	processingPlanMaterial.ProcessingProcessPosition = processingProcessPosition.Meta
	return processingPlanMaterial
}

func (processingPlanMaterial ProcessingPlanMaterial) String() string {
	return Stringify(processingPlanMaterial)
}

func (processingPlanMaterial ProcessingPlanMaterial) MetaType() MetaType {
	return MetaTypeProcessingPlanMaterial
}

type CostDistributionType string

const (
	CostDistributionByPrice      CostDistributionType = "BY_PRICE"
	CostDistributionByProduction CostDistributionType = "BY_PRODUCTION"
)

// ProcessingPlanService
// Сервис для работы с тех картами.
type ProcessingPlanService interface {
	GetList(ctx context.Context, params *Params) (*List[ProcessingPlan], *resty.Response, error)
	Create(ctx context.Context, processingPlan *ProcessingPlan, params *Params) (*ProcessingPlan, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, processingPlanList Slice[ProcessingPlan], params *Params) (*Slice[ProcessingPlan], *resty.Response, error)
	DeleteMany(ctx context.Context, processingPlanList []MetaWrapper) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params *Params) (*ProcessingPlan, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, processingPlan *ProcessingPlan, params *Params) (*ProcessingPlan, *resty.Response, error)
	GetPositions(ctx context.Context, id uuid.UUID, params *Params) (*MetaArray[ProcessingPlanProduct], *resty.Response, error)
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params *Params) (*ProcessingPlanProduct, *resty.Response, error)
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *ProcessingPlanProduct, params *Params) (*ProcessingPlanProduct, *resty.Response, error)
	CreatePosition(ctx context.Context, id uuid.UUID, position *ProcessingPlanProduct) (*ProcessingPlanProduct, *resty.Response, error)
	CreatePositions(ctx context.Context, id uuid.UUID, positions Slice[ProcessingPlanProduct]) (*Slice[ProcessingPlanProduct], *resty.Response, error)
	DeletePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (bool, *resty.Response, error)
	GetPositionTrackingCodes(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (*MetaArray[TrackingCode], *resty.Response, error)
	CreateOrUpdatePositionTrackingCodes(ctx context.Context, id uuid.UUID, positionID uuid.UUID, trackingCodes Slice[TrackingCode]) (*Slice[TrackingCode], *resty.Response, error)
	DeletePositionTrackingCodes(ctx context.Context, id uuid.UUID, positionID uuid.UUID, trackingCodes Slice[TrackingCode]) (*DeleteManyResponse, *resty.Response, error)
	MoveToTrash(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetStages(ctx context.Context, id uuid.UUID, params *Params) (*MetaArray[ProcessingStage], *resty.Response, error)
	GetStageByID(ctx context.Context, id, stageID uuid.UUID) (*ProcessingStage, *resty.Response, error)
	UpdateStage(ctx context.Context, id, stageID uuid.UUID, stage *ProcessingStage) (*ProcessingStage, *resty.Response, error)
	GetMaterials(ctx context.Context, id uuid.UUID) (*List[ProcessingPlanMaterial], *resty.Response, error)
	CreateMaterial(ctx context.Context, id uuid.UUID, material *ProcessingPlanMaterial) (*ProcessingPlanMaterial, *resty.Response, error)
	GetMaterialByID(ctx context.Context, id, materialID uuid.UUID) (*ProcessingPlanMaterial, *resty.Response, error)
	UpdateMaterial(ctx context.Context, id, materialID uuid.UUID, material *ProcessingPlanMaterial) (*ProcessingPlanMaterial, *resty.Response, error)
	DeleteMaterial(ctx context.Context, id, materialID uuid.UUID) (bool, *resty.Response, error)
	GetProducts(ctx context.Context, id uuid.UUID) (*List[ProcessingPlanProduct], *resty.Response, error)
	CreateProduct(ctx context.Context, id uuid.UUID, product *ProcessingPlanProduct) (*ProcessingPlanProduct, *resty.Response, error)
	GetProductByID(ctx context.Context, id, productID uuid.UUID) (*ProcessingPlanProduct, *resty.Response, error)
	UpdateProduct(ctx context.Context, id, productID uuid.UUID, product *ProcessingPlanProduct) (*ProcessingPlanProduct, *resty.Response, error)
	DeleteProduct(ctx context.Context, id, productID uuid.UUID) (bool, *resty.Response, error)
}

type processingPlanService struct {
	Endpoint
	endpointGetList[ProcessingPlan]
	endpointCreate[ProcessingPlan]
	endpointCreateUpdateMany[ProcessingPlan]
	endpointDeleteMany[ProcessingPlan]
	endpointDelete
	endpointGetByID[ProcessingPlan]
	endpointUpdate[ProcessingPlan]
	endpointPositions[ProcessingPlanProduct]
	endpointTrash
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
		endpointGetByID:          endpointGetByID[ProcessingPlan]{e},
		endpointUpdate:           endpointUpdate[ProcessingPlan]{e},
		endpointPositions:        endpointPositions[ProcessingPlanProduct]{e},
		endpointTrash:            endpointTrash{e},
	}
}

func (service *processingPlanService) GetStages(ctx context.Context, id uuid.UUID, params *Params) (*MetaArray[ProcessingStage], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/stages", service.uri, id)
	return NewRequestBuilder[MetaArray[ProcessingStage]](service.client, path).SetParams(params).Get(ctx)
}

func (service *processingPlanService) GetStageByID(ctx context.Context, id, stageID uuid.UUID) (*ProcessingStage, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/stages/%s", service.uri, id, stageID)
	return NewRequestBuilder[ProcessingStage](service.client, path).Get(ctx)
}

func (service *processingPlanService) UpdateStage(ctx context.Context, id, stageID uuid.UUID, stage *ProcessingStage) (*ProcessingStage, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/stages/%s", service.uri, id, stageID)
	return NewRequestBuilder[ProcessingStage](service.client, path).Put(ctx, stage)
}

func (service *processingPlanService) GetMaterials(ctx context.Context, id uuid.UUID) (*List[ProcessingPlanMaterial], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/materials", service.uri, id)
	return NewRequestBuilder[List[ProcessingPlanMaterial]](service.client, path).Get(ctx)
}

func (service *processingPlanService) CreateMaterial(ctx context.Context, id uuid.UUID, material *ProcessingPlanMaterial) (*ProcessingPlanMaterial, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/materials", service.uri, id)
	return NewRequestBuilder[ProcessingPlanMaterial](service.client, path).Post(ctx, material)
}

func (service *processingPlanService) GetMaterialByID(ctx context.Context, id, materialID uuid.UUID) (*ProcessingPlanMaterial, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/materials/%s", service.uri, id, materialID)
	return NewRequestBuilder[ProcessingPlanMaterial](service.client, path).Get(ctx)
}

func (service *processingPlanService) UpdateMaterial(ctx context.Context, id, materialID uuid.UUID, material *ProcessingPlanMaterial) (*ProcessingPlanMaterial, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/materials/%s", service.uri, id, materialID)
	return NewRequestBuilder[ProcessingPlanMaterial](service.client, path).Put(ctx, material)
}

func (service *processingPlanService) DeleteMaterial(ctx context.Context, id, materialID uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/materials/%s", service.uri, id, materialID)
	return NewRequestBuilder[any](service.client, path).Delete(ctx)
}

func (service *processingPlanService) GetProducts(ctx context.Context, id uuid.UUID) (*List[ProcessingPlanProduct], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/products", service.uri, id)
	return NewRequestBuilder[List[ProcessingPlanProduct]](service.client, path).Get(ctx)
}

func (service *processingPlanService) CreateProduct(ctx context.Context, id uuid.UUID, product *ProcessingPlanProduct) (*ProcessingPlanProduct, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/products", service.uri, id)
	return NewRequestBuilder[ProcessingPlanProduct](service.client, path).Post(ctx, product)
}

func (service *processingPlanService) GetProductByID(ctx context.Context, id, productID uuid.UUID) (*ProcessingPlanProduct, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/products/%s", service.uri, id, productID)
	return NewRequestBuilder[ProcessingPlanProduct](service.client, path).Get(ctx)
}

func (service *processingPlanService) UpdateProduct(ctx context.Context, id, productID uuid.UUID, product *ProcessingPlanProduct) (*ProcessingPlanProduct, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/products/%s", service.uri, id, productID)
	return NewRequestBuilder[ProcessingPlanProduct](service.client, path).Put(ctx, product)
}

func (service *processingPlanService) DeleteProduct(ctx context.Context, id, productID uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/products/%s", service.uri, id, productID)
	return NewRequestBuilder[any](service.client, path).Delete(ctx)
}
