package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// ProductionStageCompletion Выполнение этапа производства
// Ключевое слово: productionstagecompletion
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vypolnenie-atapa-proizwodstwa
type ProductionStageCompletion struct {
	AccountID          *uuid.UUID                                    `json:"accountId,omitempty"`          // ID учетной записи
	Created            *Timestamp                                    `json:"created,omitempty"`            // Дата создания
	ExternalCode       *string                                       `json:"externalCode,omitempty"`       // Внешний код Выполнения этапа производства
	Group              *Group                                        `json:"group,omitempty"`              // Отдел сотрудника
	ID                 *uuid.UUID                                    `json:"id,omitempty"`                 // ID Выполнения этапа производства
	LabourUnitCost     *float64                                      `json:"labourUnitCost,omitempty"`     // Оплата труда за единицу объема производства
	StandardHourUnit   *float64                                      `json:"standardHourUnit,omitempty"`   // Нормо-часы единицы объема производства
	Materials          *Positions[ProductionStageCompletionMaterial] `json:"materials,omitempty"`          // Метаданные Материалов выполнения этапа производства
	Meta               *Meta                                         `json:"meta,omitempty"`               // Метаданные Выполнения этапа производства
	Moment             *Timestamp                                    `json:"moment,omitempty"`             // Дата документа
	Name               *string                                       `json:"name,omitempty"`               // Наименование Выполнения этапа производства
	Owner              *Employee                                     `json:"owner,omitempty"`              // Владелец (Сотрудник)
	Performer          *Employee                                     `json:"performer,omitempty"`          // Исполнитель (Сотрудник)
	ProcessingUnitCost *float64                                      `json:"processingUnitCost,omitempty"` // Затраты на единицу объема производства
	ProductionStage    *ProductionStage                              `json:"productionStage,omitempty"`    // Производственный этап
	ProductionVolume   *float64                                      `json:"productionVolume,omitempty"`   // Объем производства
	Products           *Positions[ProductionStageCompletionResult]   `json:"products,omitempty"`           // Метаданные Продуктов выполнения этапа производства. Есть только у последнего этапа
	Shared             *bool                                         `json:"shared,omitempty"`             // Общий доступ
	Updated            *Timestamp                                    `json:"updated,omitempty"`            // Момент последнего обновления Выполнения этапа производства
}

// Clean возвращает сущность с единственным заполненным полем Meta
func (productionStageCompletion ProductionStageCompletion) Clean() *ProductionStageCompletion {
	return &ProductionStageCompletion{Meta: productionStageCompletion.Meta}
}

func (productionStageCompletion ProductionStageCompletion) GetAccountID() uuid.UUID {
	return Deref(productionStageCompletion.AccountID)
}

func (productionStageCompletion ProductionStageCompletion) GetCreated() Timestamp {
	return Deref(productionStageCompletion.Created)
}

func (productionStageCompletion ProductionStageCompletion) GetExternalCode() string {
	return Deref(productionStageCompletion.ExternalCode)
}

func (productionStageCompletion ProductionStageCompletion) GetGroup() Group {
	return Deref(productionStageCompletion.Group)
}

func (productionStageCompletion ProductionStageCompletion) GetID() uuid.UUID {
	return Deref(productionStageCompletion.ID)
}

func (productionStageCompletion ProductionStageCompletion) GetLabourUnitCost() float64 {
	return Deref(productionStageCompletion.LabourUnitCost)
}

func (productionStageCompletion ProductionStageCompletion) GetStandardHourUnit() float64 {
	return Deref(productionStageCompletion.StandardHourUnit)
}

func (productionStageCompletion ProductionStageCompletion) GetMaterials() Positions[ProductionStageCompletionMaterial] {
	return Deref(productionStageCompletion.Materials)
}

func (productionStageCompletion ProductionStageCompletion) GetMeta() Meta {
	return Deref(productionStageCompletion.Meta)
}

func (productionStageCompletion ProductionStageCompletion) GetMoment() Timestamp {
	return Deref(productionStageCompletion.Moment)
}

func (productionStageCompletion ProductionStageCompletion) GetName() string {
	return Deref(productionStageCompletion.Name)
}

func (productionStageCompletion ProductionStageCompletion) GetOwner() Employee {
	return Deref(productionStageCompletion.Owner)
}

func (productionStageCompletion ProductionStageCompletion) GetPerformer() Employee {
	return Deref(productionStageCompletion.Performer)
}

func (productionStageCompletion ProductionStageCompletion) GetProcessingUnitCost() float64 {
	return Deref(productionStageCompletion.ProcessingUnitCost)
}

func (productionStageCompletion ProductionStageCompletion) GetProductionStage() ProductionStage {
	return Deref(productionStageCompletion.ProductionStage)
}

func (productionStageCompletion ProductionStageCompletion) GetProductionVolume() float64 {
	return Deref(productionStageCompletion.ProductionVolume)
}

func (productionStageCompletion ProductionStageCompletion) GetProducts() Positions[ProductionStageCompletionResult] {
	return Deref(productionStageCompletion.Products)
}

func (productionStageCompletion ProductionStageCompletion) GetShared() bool {
	return Deref(productionStageCompletion.Shared)
}

func (productionStageCompletion ProductionStageCompletion) GetUpdated() Timestamp {
	return Deref(productionStageCompletion.Updated)
}

func (productionStageCompletion *ProductionStageCompletion) SetExternalCode(externalCode string) *ProductionStageCompletion {
	productionStageCompletion.ExternalCode = &externalCode
	return productionStageCompletion
}

func (productionStageCompletion *ProductionStageCompletion) SetGroup(group *Group) *ProductionStageCompletion {
	productionStageCompletion.Group = group.Clean()
	return productionStageCompletion
}

func (productionStageCompletion *ProductionStageCompletion) SetLabourUnitCost(labourUnitCost float64) *ProductionStageCompletion {
	productionStageCompletion.LabourUnitCost = &labourUnitCost
	return productionStageCompletion
}

func (productionStageCompletion *ProductionStageCompletion) SetStandardHourUnit(standardHourUnit float64) *ProductionStageCompletion {
	productionStageCompletion.StandardHourUnit = &standardHourUnit
	return productionStageCompletion
}

func (productionStageCompletion *ProductionStageCompletion) SetMaterials(materials *Positions[ProductionStageCompletionMaterial]) *ProductionStageCompletion {
	productionStageCompletion.Materials = materials
	return productionStageCompletion
}

func (productionStageCompletion *ProductionStageCompletion) SetMeta(meta *Meta) *ProductionStageCompletion {
	productionStageCompletion.Meta = meta
	return productionStageCompletion
}

func (productionStageCompletion *ProductionStageCompletion) SetMoment(moment *Timestamp) *ProductionStageCompletion {
	productionStageCompletion.Moment = moment
	return productionStageCompletion
}

func (productionStageCompletion *ProductionStageCompletion) SetName(name string) *ProductionStageCompletion {
	productionStageCompletion.Name = &name
	return productionStageCompletion
}

func (productionStageCompletion *ProductionStageCompletion) SetOwner(owner *Employee) *ProductionStageCompletion {
	productionStageCompletion.Owner = owner.Clean()
	return productionStageCompletion
}

func (productionStageCompletion *ProductionStageCompletion) SetPerformer(performer *Employee) *ProductionStageCompletion {
	productionStageCompletion.Performer = performer.Clean()
	return productionStageCompletion
}

func (productionStageCompletion *ProductionStageCompletion) SetProcessingUnitCost(processingUnitCost float64) *ProductionStageCompletion {
	productionStageCompletion.ProcessingUnitCost = &processingUnitCost
	return productionStageCompletion
}

func (productionStageCompletion *ProductionStageCompletion) SetProductionStage(productionStage *ProductionStage) *ProductionStageCompletion {
	productionStageCompletion.ProductionStage = productionStage.Clean()
	return productionStageCompletion
}

func (productionStageCompletion *ProductionStageCompletion) SetProductionVolume(productionVolume float64) *ProductionStageCompletion {
	productionStageCompletion.ProductionVolume = &productionVolume
	return productionStageCompletion
}

func (productionStageCompletion *ProductionStageCompletion) SetProducts(products *Positions[ProductionStageCompletionResult]) *ProductionStageCompletion {
	productionStageCompletion.Products = products
	return productionStageCompletion
}

func (productionStageCompletion *ProductionStageCompletion) SetShared(shared bool) *ProductionStageCompletion {
	productionStageCompletion.Shared = &shared
	return productionStageCompletion
}

func (productionStageCompletion ProductionStageCompletion) String() string {
	return Stringify(productionStageCompletion)
}

func (productionStageCompletion ProductionStageCompletion) MetaType() MetaType {
	return MetaTypeProductionStageCompletion
}

// Update shortcut
func (productionStageCompletion ProductionStageCompletion) Update(ctx context.Context, client *Client, params ...*Params) (*ProductionStageCompletion, *resty.Response, error) {
	return client.Entity().ProductionStageCompletion().Update(ctx, productionStageCompletion.GetID(), &productionStageCompletion, params...)
}

// Create shortcut
func (productionStageCompletion ProductionStageCompletion) Create(ctx context.Context, client *Client, params ...*Params) (*ProductionStageCompletion, *resty.Response, error) {
	return client.Entity().ProductionStageCompletion().Create(ctx, &productionStageCompletion, params...)
}

// Delete shortcut
func (productionStageCompletion ProductionStageCompletion) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return client.Entity().ProductionStageCompletion().Delete(ctx, productionStageCompletion.GetID())
}

// ProductionStageCompletionMaterial Материалы Выполнения этапа производства
// Ключевое слово: productionstagecompletionmaterial
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vypolnenie-atapa-proizwodstwa-izmenit-vypolnenie-atapa-proizwodstwa-materialy-vypolneniq-atapa-proizwodstwa
type ProductionStageCompletionMaterial struct {
	AccountID        *uuid.UUID          `json:"accountId,omitempty"`        // ID учетной записи
	Assortment       *AssortmentPosition `json:"assortment,omitempty"`       // Метаданные товара/модификации/серии, которую представляет собой позиция
	ConsumedQuantity *float64            `json:"consumedQuantity,omitempty"` // Количество товаров/модификаций данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе
	ID               *uuid.UUID          `json:"id,omitempty"`               // ID позиции
	Things           Slice[string]       `json:"things,omitempty"`           // Серийные номера. Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете. В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута
}

func (productionStageCompletionMaterial ProductionStageCompletionMaterial) GetAccountID() uuid.UUID {
	return Deref(productionStageCompletionMaterial.AccountID)
}

func (productionStageCompletionMaterial ProductionStageCompletionMaterial) GetAssortment() AssortmentPosition {
	return Deref(productionStageCompletionMaterial.Assortment)
}

func (productionStageCompletionMaterial ProductionStageCompletionMaterial) GetConsumedQuantity() float64 {
	return Deref(productionStageCompletionMaterial.ConsumedQuantity)
}

func (productionStageCompletionMaterial ProductionStageCompletionMaterial) GetID() uuid.UUID {
	return Deref(productionStageCompletionMaterial.ID)
}

func (productionStageCompletionMaterial ProductionStageCompletionMaterial) GetThings() Slice[string] {
	return productionStageCompletionMaterial.Things
}

func (productionStageCompletionMaterial *ProductionStageCompletionMaterial) SetAssortment(assortment AsAssortment) *ProductionStageCompletionMaterial {
	productionStageCompletionMaterial.Assortment = assortment.AsAssortment()
	return productionStageCompletionMaterial
}

func (productionStageCompletionMaterial *ProductionStageCompletionMaterial) SetConsumedQuantity(consumedQuantity float64) *ProductionStageCompletionMaterial {
	productionStageCompletionMaterial.ConsumedQuantity = &consumedQuantity
	return productionStageCompletionMaterial
}

func (productionStageCompletionMaterial *ProductionStageCompletionMaterial) SetThings(things ...string) *ProductionStageCompletionMaterial {
	productionStageCompletionMaterial.Things = NewSliceFrom(things)
	return productionStageCompletionMaterial
}

func (productionStageCompletionMaterial ProductionStageCompletionMaterial) String() string {
	return Stringify(productionStageCompletionMaterial)
}

func (productionStageCompletionMaterial ProductionStageCompletionMaterial) MetaType() MetaType {
	return MetaTypeProductionStageCompletionMaterial
}

// ProductionStageCompletionResult Продукт Выполнения этапа производства
// Ключевое слово: productionstagecompletionresult
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vypolnenie-atapa-proizwodstwa-izmenit-vypolnenie-atapa-proizwodstwa-produkty-vypolneniq-atapa-proizwodstwa
type ProductionStageCompletionResult struct {
	AccountID        *uuid.UUID          `json:"accountId,omitempty"`        // ID учетной записи
	Assortment       *AssortmentPosition `json:"assortment,omitempty"`       // Метаданные товара/модификации/серии, которую представляет собой позиция
	ID               *uuid.UUID          `json:"id,omitempty"`               // ID позиции
	ProducedQuantity *float64            `json:"producedQuantity,omitempty"` // Количество товаров/модификаций данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе
	Things           Slice[string]       `json:"things,omitempty"`           // Серийные номера. Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете. В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута
}

func (productionStageCompletionResult ProductionStageCompletionResult) GetAccountID() uuid.UUID {
	return Deref(productionStageCompletionResult.AccountID)
}

func (productionStageCompletionResult ProductionStageCompletionResult) GetAssortment() AssortmentPosition {
	return Deref(productionStageCompletionResult.Assortment)
}

func (productionStageCompletionResult ProductionStageCompletionResult) GetID() uuid.UUID {
	return Deref(productionStageCompletionResult.ID)
}

func (productionStageCompletionResult ProductionStageCompletionResult) GetProducedQuantity() float64 {
	return Deref(productionStageCompletionResult.ProducedQuantity)
}

func (productionStageCompletionResult ProductionStageCompletionResult) GetThings() Slice[string] {
	return productionStageCompletionResult.Things
}

func (productionStageCompletionResult *ProductionStageCompletionResult) SetAssortment(assortment AsAssortment) *ProductionStageCompletionResult {
	productionStageCompletionResult.Assortment = assortment.AsAssortment()
	return productionStageCompletionResult
}

func (productionStageCompletionResult *ProductionStageCompletionResult) SetProducedQuantity(producedQuantity float64) *ProductionStageCompletionResult {
	productionStageCompletionResult.ProducedQuantity = &producedQuantity
	return productionStageCompletionResult
}

func (productionStageCompletionResult *ProductionStageCompletionResult) SetThings(things ...string) *ProductionStageCompletionResult {
	productionStageCompletionResult.Things = NewSliceFrom(things)
	return productionStageCompletionResult
}

func (productionStageCompletionResult ProductionStageCompletionResult) String() string {
	return Stringify(productionStageCompletionResult)
}

func (productionStageCompletionResult ProductionStageCompletionResult) MetaType() MetaType {
	return MetaTypeProductionStageCompletionResult
}

// ProductionStageCompletionService
// Сервис для работы с выполнениями этапов производства
type ProductionStageCompletionService interface {
	GetList(ctx context.Context, params ...*Params) (*List[ProductionStageCompletion], *resty.Response, error)
	Create(ctx context.Context, productionStageCompletion *ProductionStageCompletion, params ...*Params) (*ProductionStageCompletion, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, productionStageCompletionList Slice[ProductionStageCompletion], params ...*Params) (*Slice[ProductionStageCompletion], *resty.Response, error)
	DeleteMany(ctx context.Context, entities ...*ProductionStageCompletion) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*ProductionStageCompletion, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, productionStageCompletion *ProductionStageCompletion, params ...*Params) (*ProductionStageCompletion, *resty.Response, error)
	GetMaterials(ctx context.Context, id uuid.UUID, params ...*Params) (*MetaArray[ProductionStageCompletionMaterial], *resty.Response, error)
	CreateMaterial(ctx context.Context, id uuid.UUID, productionStageCompletionMaterial *ProductionStageCompletionMaterial, params ...*Params) (*ProductionStageCompletionMaterial, *resty.Response, error)
	UpdateMaterial(ctx context.Context, id uuid.UUID, materialID uuid.UUID, productionStageCompletionMaterial *ProductionStageCompletionMaterial, params ...*Params) (*ProductionStageCompletionMaterial, *resty.Response, error)
	GetProducts(ctx context.Context, id uuid.UUID, params ...*Params) (*MetaArray[ProductionStageCompletionResult], *resty.Response, error)
	UpdateProduct(ctx context.Context, id uuid.UUID, productID uuid.UUID, productionStageCompletionResult *ProductionStageCompletionResult, params ...*Params) (*ProductionStageCompletionResult, *resty.Response, error)
}

type productionStageCompletionService struct {
	Endpoint
	endpointGetList[ProductionStageCompletion]
	endpointCreate[ProductionStageCompletion]
	endpointCreateUpdateMany[ProductionStageCompletion]
	endpointDeleteMany[ProductionStageCompletion]
	endpointDelete
	endpointGetByID[ProductionStageCompletion]
	endpointUpdate[ProductionStageCompletion]
}

func NewProductionStageCompletionService(client *Client) ProductionStageCompletionService {
	e := NewEndpoint(client, "entity/productionstagecompletion")
	return &productionStageCompletionService{
		Endpoint:                 e,
		endpointGetList:          endpointGetList[ProductionStageCompletion]{e},
		endpointCreate:           endpointCreate[ProductionStageCompletion]{e},
		endpointCreateUpdateMany: endpointCreateUpdateMany[ProductionStageCompletion]{e},
		endpointDeleteMany:       endpointDeleteMany[ProductionStageCompletion]{e},
		endpointDelete:           endpointDelete{e},
		endpointGetByID:          endpointGetByID[ProductionStageCompletion]{e},
		endpointUpdate:           endpointUpdate[ProductionStageCompletion]{e},
	}
}

// GetMaterials Получить Материалы выполнения этапа производства.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vypolnenie-atapa-proizwodstwa-poluchit-materialy-wypolneniq-atapa-proizwodstwa
func (service *productionStageCompletionService) GetMaterials(ctx context.Context, id uuid.UUID, params ...*Params) (*MetaArray[ProductionStageCompletionMaterial], *resty.Response, error) {
	path := fmt.Sprintf("%s/materials", id)
	return NewRequestBuilder[MetaArray[ProductionStageCompletionMaterial]](service.client, path).SetParams(params...).Get(ctx)
}

// CreateMaterial Добавить Материал выполнения этапа производства.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vypolnenie-atapa-proizwodstwa-dobawit-material-wypolneniq-atapa-proizwodstwa
func (service *productionStageCompletionService) CreateMaterial(ctx context.Context, id uuid.UUID, productionStageCompletionMaterial *ProductionStageCompletionMaterial, params ...*Params) (*ProductionStageCompletionMaterial, *resty.Response, error) {
	path := fmt.Sprintf("%s/materials", id)
	return NewRequestBuilder[ProductionStageCompletionMaterial](service.client, path).SetParams(params...).Post(ctx, productionStageCompletionMaterial)
}

// UpdateMaterial Изменить Материал выполнения этапа производства.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vypolnenie-atapa-proizwodstwa-izmenit-material-wypolneniq-atapa-proizwodstwa
func (service *productionStageCompletionService) UpdateMaterial(ctx context.Context, id uuid.UUID, materialID uuid.UUID, productionStageCompletionMaterial *ProductionStageCompletionMaterial, params ...*Params) (*ProductionStageCompletionMaterial, *resty.Response, error) {
	path := fmt.Sprintf("%s/materials/%s", id, materialID)
	return NewRequestBuilder[ProductionStageCompletionMaterial](service.client, path).SetParams(params...).Put(ctx, productionStageCompletionMaterial)
}

// GetProducts Получить Продукты выполнения этапа производства.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vypolnenie-atapa-proizwodstwa-poluchit-produkty-wypolneniq-atapa-proizwodstwa
func (service *productionStageCompletionService) GetProducts(ctx context.Context, id uuid.UUID, params ...*Params) (*MetaArray[ProductionStageCompletionResult], *resty.Response, error) {
	path := fmt.Sprintf("%s/products", id)
	return NewRequestBuilder[MetaArray[ProductionStageCompletionResult]](service.client, path).SetParams(params...).Get(ctx)
}

// UpdateProduct Изменить Продукт выполнения этапа производства.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vypolnenie-atapa-proizwodstwa-izmenit-produkt-wypolneniq-atapa-proizwodstwa
func (service *productionStageCompletionService) UpdateProduct(ctx context.Context, id uuid.UUID, productID uuid.UUID, productionStageCompletionResult *ProductionStageCompletionResult, params ...*Params) (*ProductionStageCompletionResult, *resty.Response, error) {
	path := fmt.Sprintf("%s/products/%s", id, productID)
	return NewRequestBuilder[ProductionStageCompletionResult](service.client, path).SetParams(params...).Put(ctx, productionStageCompletionResult)
}
