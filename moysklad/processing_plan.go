package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"

	"time"
)

// ProcessingPlan Техкарта.
//
// Код сущности: processingplan
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tehkarta-tehkarty
type ProcessingPlan struct {
	AccountID            *string                            `json:"accountId,omitempty"`            // ID учётной записи            // ID учётной записи
	Archived             *bool                              `json:"archived,omitempty"`             // Добавлена ли Тех. карта в архив
	Code                 *string                            `json:"code,omitempty"`                 // Код Тех. карты
	Cost                 *float64                           `json:"cost,omitempty"`                 // Стоимость производства
	ExternalCode         *string                            `json:"externalCode,omitempty"`         // Внешний код
	Group                *Group                             `json:"group,omitempty"`                // Отдел сотрудника                // Отдел сотрудника
	ID                   *string                            `json:"id,omitempty"`                   // ID сущности
	Stages               *MetaArray[ProcessingPlanStages]   `json:"stages,omitempty"`               // Коллекция метаданных этапов Тех. карты
	Materials            *MetaArray[ProcessingPlanMaterial] `json:"materials,omitempty"`            // Список Метаданных материалов Тех. операции
	Meta                 *Meta                              `json:"meta,omitempty"`                 // Метаданные
	Name                 *string                            `json:"name,omitempty"`                 // Наименование
	Owner                *Employee                          `json:"owner,omitempty"`                // Метаданные владельца (Сотрудника)                // Владелец (Сотрудник)
	Parent               *Group                             `json:"parent,omitempty"`               // Метаданные группы Тех. карты
	PathName             *string                            `json:"pathName,omitempty"`             // Наименование группы, в которую входит Тех. карта
	ProcessingProcess    *ProcessingProcess                 `json:"processingProcess,omitempty"`    // Метаданные Тех. процесса
	Products             *MetaArray[ProcessingPlanProduct]  `json:"products,omitempty"`             // Коллекция метаданных готовых продуктов Тех. карты
	Shared               *bool                              `json:"shared,omitempty"`               // Общий доступ               // Общий доступ
	Updated              *Timestamp                         `json:"updated,omitempty"`              // Момент последнего обновления
	CostDistributionType CostDistributionType               `json:"costDistributionType,omitempty"` // Тип распределения себестоимости. Возможные значения: BY_PRICE, BY_PRODUCTION
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (processingPlan ProcessingPlan) Clean() *ProcessingPlan {
	if processingPlan.Meta == nil {
		return nil
	}
	return &ProcessingPlan{Meta: processingPlan.Meta}
}

func (processingPlan ProcessingPlan) GetAccountID() string {
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

func (processingPlan ProcessingPlan) GetID() string {
	return Deref(processingPlan.ID)
}

func (processingPlan ProcessingPlan) GetStages() MetaArray[ProcessingPlanStages] {
	return Deref(processingPlan.Stages)
}

func (processingPlan ProcessingPlan) GetMaterials() MetaArray[ProcessingPlanMaterial] {
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

func (processingPlan ProcessingPlan) GetProducts() MetaArray[ProcessingPlanProduct] {
	return Deref(processingPlan.Products)
}

func (processingPlan ProcessingPlan) GetShared() bool {
	return Deref(processingPlan.Shared)
}

func (processingPlan ProcessingPlan) GetUpdated() time.Time {
	return Deref(processingPlan.Updated).Time()
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

func (processingPlan *ProcessingPlan) SetStages(stages ...*ProcessingPlanStages) *ProcessingPlan {
	processingPlan.Stages = NewMetaArrayFrom(stages)
	return processingPlan
}

func (processingPlan *ProcessingPlan) SetMaterials(materials ...*ProcessingPlanMaterial) *ProcessingPlan {
	processingPlan.Materials = NewMetaArrayFrom(materials)
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

func (processingPlan *ProcessingPlan) SetProducts(products ...*ProcessingPlanProduct) *ProcessingPlan {
	processingPlan.Products = NewMetaArrayFrom(products)
	return processingPlan
}

func (processingPlan *ProcessingPlan) SetShared(shared bool) *ProcessingPlan {
	processingPlan.Shared = &shared
	return processingPlan
}

func (processingPlan ProcessingPlan) String() string {
	return Stringify(processingPlan)
}

// MetaType возвращает код сущности.
func (ProcessingPlan) MetaType() MetaType {
	return MetaTypeProcessingPlan
}

// Update shortcut
func (processingPlan *ProcessingPlan) Update(ctx context.Context, client *Client, params ...func(*Params)) (*ProcessingPlan, *resty.Response, error) {
	return NewProcessingPlanService(client).Update(ctx, processingPlan.GetID(), processingPlan, params...)
}

// Create shortcut
func (processingPlan *ProcessingPlan) Create(ctx context.Context, client *Client, params ...func(*Params)) (*ProcessingPlan, *resty.Response, error) {
	return NewProcessingPlanService(client).Create(ctx, processingPlan, params...)
}

// Delete shortcut
func (processingPlan *ProcessingPlan) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewProcessingPlanService(client).Delete(ctx, processingPlan)
}

// ProcessingPlanStages Этапы Техкарты.
// // Код сущности: processingplanstages.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tehkarta-tehkarty-jetapy-tehkarty
type ProcessingPlanStages struct {
	AccountID                 *string  `json:"accountId,omitempty"`                 // ID учётной записи                 // ID учётной записи
	ID                        *string  `json:"id,omitempty"`                        // ID Материала
	Cost                      *float64 `json:"cost,omitempty"`                      // Стоимость производства, на определенном этапе
	LabourCost                *float64 `json:"labourCost,omitempty"`                // Оплата труда, на определенном этапе
	StandardHour              *float64 `json:"standardHour,omitempty"`              // Нормо-часы, на определенном этапе
	ProcessingProcessPosition *Meta    `json:"processingProcessPosition,omitempty"` // Метаданные позиции техпроцесса
}

func (processingPlanStages ProcessingPlanStages) GetAccountID() string {
	return Deref(processingPlanStages.AccountID)
}

func (processingPlanStages ProcessingPlanStages) GetID() string {
	return Deref(processingPlanStages.ID)
}

func (processingPlanStages ProcessingPlanStages) GetCost() float64 {
	return Deref(processingPlanStages.Cost)
}

func (processingPlanStages ProcessingPlanStages) GetLabourCost() float64 {
	return Deref(processingPlanStages.LabourCost)
}

func (processingPlanStages ProcessingPlanStages) GetStandardHour() float64 {
	return Deref(processingPlanStages.StandardHour)
}

func (processingPlanStages ProcessingPlanStages) GetProcessingProcessPosition() Meta {
	return Deref(processingPlanStages.ProcessingProcessPosition)
}

func (processingPlanStages *ProcessingPlanStages) SetCost(cost float64) *ProcessingPlanStages {
	processingPlanStages.Cost = &cost
	return processingPlanStages
}

func (processingPlanStages *ProcessingPlanStages) SetLabourCost(labourCost float64) *ProcessingPlanStages {
	processingPlanStages.LabourCost = &labourCost
	return processingPlanStages
}

func (processingPlanStages *ProcessingPlanStages) SetStandardHour(standardHour float64) *ProcessingPlanStages {
	processingPlanStages.StandardHour = &standardHour
	return processingPlanStages
}

func (processingPlanStages ProcessingPlanStages) String() string {
	return Stringify(processingPlanStages)
}

// MetaType возвращает код сущности.
func (ProcessingPlanStages) MetaType() MetaType {
	return MetaTypeProcessingPlanStages
}

// ProcessingPlanProduct Продукт Тех. карты.
//
// Код сущности: processingplanresult
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tehkarta-tehkarty-produkty-tehkarty
type ProcessingPlanProduct struct {
	AccountID  *string             `json:"accountId,omitempty"`  // ID учётной записи  // ID учётной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара или модификации позиции
	ID         *string             `json:"id,omitempty"`         // ID позиции
	Product    *Product            `json:"product,omitempty"`    // Метаданные товара позиции. В случае, если в поле assortment указана модификация, то это поле содержит товар, к которому относится эта модификация
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров данного вида в позиции
}

func (processingPlanProduct ProcessingPlanProduct) GetAccountID() string {
	return Deref(processingPlanProduct.AccountID)
}

func (processingPlanProduct ProcessingPlanProduct) GetAssortment() AssortmentPosition {
	return Deref(processingPlanProduct.Assortment)
}

func (processingPlanProduct ProcessingPlanProduct) GetID() string {
	return Deref(processingPlanProduct.ID)
}

func (processingPlanProduct ProcessingPlanProduct) GetProduct() Product {
	return Deref(processingPlanProduct.Product)
}

func (processingPlanProduct ProcessingPlanProduct) GetQuantity() float64 {
	return Deref(processingPlanProduct.Quantity)
}

// Принимает объект, реализующий интерфейс [AssortmentConverter].
func (processingPlanProduct *ProcessingPlanProduct) SetAssortment(assortment AssortmentConverter) *ProcessingPlanProduct {
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

// MetaType возвращает код сущности.
func (ProcessingPlanProduct) MetaType() MetaType {
	return MetaTypeProcessingPlanProduct
}

// ProcessingPlanMaterial Материал Тех. карты.
//
// Код сущности: processingplanmaterial
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tehkarta-tehkarty-materialy-tehkarty
type ProcessingPlanMaterial struct {
	AccountID                 *string             `json:"accountId,omitempty"`                 // ID учётной записи                 // ID учётной записи
	Assortment                *AssortmentPosition `json:"assortment,omitempty"`                // Метаданные товара или модификации позиции
	ID                        *string             `json:"id,omitempty"`                        // ID позиции
	Product                   *Product            `json:"product,omitempty"`                   // Метаданные товара позиции. В случае, если в поле assortment указана модификация, то это поле содержит товар, к которому относится эта модификация
	Quantity                  *float64            `json:"quantity,omitempty"`                  // Количество товаров данного вида в позиции
	ProcessingProcessPosition *Meta               `json:"processingProcessPosition,omitempty"` // Метаданные позиции Тех. процесса
	MaterialProcessingPlan    *Meta               `json:"materialProcessingPlan"`              // Метаданные техкарты материала [11-01-2024]
}

func (processingPlanMaterial ProcessingPlanMaterial) GetAccountID() string {
	return Deref(processingPlanMaterial.AccountID)
}

func (processingPlanMaterial ProcessingPlanMaterial) GetAssortment() AssortmentPosition {
	return Deref(processingPlanMaterial.Assortment)
}

func (processingPlanMaterial ProcessingPlanMaterial) GetID() string {
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

// Принимает объект, реализующий интерфейс [AssortmentConverter].
func (processingPlanMaterial *ProcessingPlanMaterial) SetAssortment(assortment AssortmentConverter) *ProcessingPlanMaterial {
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

// MetaType возвращает код сущности.
func (ProcessingPlanMaterial) MetaType() MetaType {
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
	GetList(ctx context.Context, params ...func(*Params)) (*List[ProcessingPlan], *resty.Response, error)
	Create(ctx context.Context, processingPlan *ProcessingPlan, params ...func(*Params)) (*ProcessingPlan, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, processingPlanList Slice[ProcessingPlan], params ...func(*Params)) (*Slice[ProcessingPlan], *resty.Response, error)
	DeleteMany(ctx context.Context, entities ...*ProcessingPlan) (*DeleteManyResponse, *resty.Response, error)
	DeleteByID(ctx context.Context, id string) (bool, *resty.Response, error)

	// Delete выполняет запрос на удаление техкарты.
	// Принимает контекст и техкарту.
	// Возвращает «true» в случае успешного удаления техкарты.
	Delete(ctx context.Context, entity *ProcessingPlan) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id string, params ...func(*Params)) (*ProcessingPlan, *resty.Response, error)
	Update(ctx context.Context, id string, processingPlan *ProcessingPlan, params ...func(*Params)) (*ProcessingPlan, *resty.Response, error)

	// GetPositionList выполняет запрос на получение списка позиций документа.
	// Принимает контекст, ID документа и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetPositionList(ctx context.Context, id string, params ...func(*Params)) (*List[ProcessingPlanProduct], *resty.Response, error)

	GetPositionListAll(ctx context.Context, id string, params ...func(*Params)) (*Slice[ProcessingPlanProduct], *resty.Response, error)

	// GetPositionByID выполняет запрос на получение отдельной позиции документа по ID.
	// Принимает контекст, ID документа, ID позиции и опционально объект параметров запроса Params.
	// Возвращает найденную позицию.
	GetPositionByID(ctx context.Context, id string, positionID string, params ...func(*Params)) (*ProcessingPlanProduct, *resty.Response, error)

	// UpdatePosition выполняет запрос на изменение позиции документа.
	// Принимает контекст, ID документа, ID позиции, позицию документа и опционально объект параметров запроса Params.
	// Возвращает изменённую позицию.
	UpdatePosition(ctx context.Context, id string, positionID string, position *ProcessingPlanProduct, params ...func(*Params)) (*ProcessingPlanProduct, *resty.Response, error)

	// CreatePosition выполняет запрос на добавление позиции документа.
	// Принимает контекст, ID документа, позицию документа и опционально объект параметров запроса Params.
	// Возвращает добавленную позицию.
	CreatePosition(ctx context.Context, id string, position *ProcessingPlanProduct, params ...func(*Params)) (*ProcessingPlanProduct, *resty.Response, error)

	// CreatePositionMany выполняет запрос на массовое добавление позиций документа.
	// Принимает контекст, ID документа и множество позиций.
	// Возвращает список добавленных позиций.
	CreatePositionMany(ctx context.Context, id string, positions ...*ProcessingPlanProduct) (*Slice[ProcessingPlanProduct], *resty.Response, error)

	// DeletePosition выполняет запрос на удаление позиции документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает «true» в случае успешного удаления позиции.
	DeletePosition(ctx context.Context, id string, positionID string) (bool, *resty.Response, error)

	// DeletePositionMany выполняет запрос на массовое удаление позиций документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeletePositionMany(ctx context.Context, id string, positions ...*ProcessingPlanProduct) (*DeleteManyResponse, *resty.Response, error)

	// GetPositionTrackingCodeList выполняет запрос на получение кодов маркировки позиции документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает объект List.
	GetPositionTrackingCodeList(ctx context.Context, id string, positionID string) (*List[TrackingCode], *resty.Response, error)

	// CreateUpdatePositionTrackingCodeMany выполняет запрос на массовое создание/изменение кодов маркировки позиции документа.
	// Принимает контекст, ID документа, ID позиции и множество кодов маркировки.
	// Возвращает список созданных и/или изменённых кодов маркировки позиции документа.
	CreateUpdatePositionTrackingCodeMany(ctx context.Context, id string, positionID string, trackingCodes ...*TrackingCode) (*Slice[TrackingCode], *resty.Response, error)

	// DeletePositionTrackingCodeMany выполняет запрос на массовое удаление кодов маркировки позиции документа.
	// Принимает контекст, ID документа, ID позиции и множество кодов маркировки.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeletePositionTrackingCodeMany(ctx context.Context, id string, positionID string, trackingCodes ...*TrackingCode) (*DeleteManyResponse, *resty.Response, error)

	// MoveToTrash выполняет запрос на перемещение документа с указанным ID в корзину.
	// Принимает контекст и ID документа.
	// Возвращает «true» в случае успешного перемещения в корзину.
	MoveToTrash(ctx context.Context, id string) (bool, *resty.Response, error)
	GetStages(ctx context.Context, id string, params ...func(*Params)) (*MetaArray[ProcessingStage], *resty.Response, error)
	GetStageByID(ctx context.Context, id, stageID string) (*ProcessingStage, *resty.Response, error)
	UpdateStage(ctx context.Context, id, stageID string, stage *ProcessingStage) (*ProcessingStage, *resty.Response, error)
	GetMaterials(ctx context.Context, id string) (*List[ProcessingPlanMaterial], *resty.Response, error)
	CreateMaterial(ctx context.Context, id string, material *ProcessingPlanMaterial) (*ProcessingPlanMaterial, *resty.Response, error)
	GetMaterialByID(ctx context.Context, id, materialID string) (*ProcessingPlanMaterial, *resty.Response, error)
	UpdateMaterial(ctx context.Context, id, materialID string, material *ProcessingPlanMaterial) (*ProcessingPlanMaterial, *resty.Response, error)
	DeleteMaterial(ctx context.Context, id, materialID string) (bool, *resty.Response, error)
	GetProducts(ctx context.Context, id string) (*List[ProcessingPlanProduct], *resty.Response, error)
	CreateProduct(ctx context.Context, id string, product *ProcessingPlanProduct) (*ProcessingPlanProduct, *resty.Response, error)
	GetProductByID(ctx context.Context, id, productID string) (*ProcessingPlanProduct, *resty.Response, error)
	UpdateProduct(ctx context.Context, id, productID string, product *ProcessingPlanProduct) (*ProcessingPlanProduct, *resty.Response, error)
	DeleteProduct(ctx context.Context, id, productID string) (bool, *resty.Response, error)
}

const (
	EndpointProcessingPlan            = EndpointEntity + string(MetaTypeProcessingPlan)
	EndpointProcessingPlanStages      = EndpointProcessingPlan + "/%s/stages"
	EndpointProcessingPlanStagesID    = EndpointProcessingPlanStages + "/%s"
	EndpointProcessingPlanMaterials   = EndpointProcessingPlan + "/%s/materials"
	EndpointProcessingPlanMaterialsID = EndpointProcessingPlanMaterials + "/%s"
	EndpointProcessingPlanProducts    = EndpointProcessingPlan + "/%s/products"
	EndpointProcessingPlanProductsID  = EndpointProcessingPlanProducts + "/%s"
)

type processingPlanService struct {
	Endpoint
	endpointGetList[ProcessingPlan]
	endpointCreate[ProcessingPlan]
	endpointCreateUpdateMany[ProcessingPlan]
	endpointDeleteMany[ProcessingPlan]
	endpointDeleteByID
	endpointDelete[ProcessingPlan]
	endpointGetByID[ProcessingPlan]
	endpointUpdate[ProcessingPlan]
	endpointPositions[ProcessingPlanProduct]
	endpointTrash
}

func (service *processingPlanService) GetStages(ctx context.Context, id string, params ...func(*Params)) (*MetaArray[ProcessingStage], *resty.Response, error) {
	path := fmt.Sprintf(EndpointProcessingPlanStages, id)
	return NewRequestBuilder[MetaArray[ProcessingStage]](service.client, path).SetParams(params).Get(ctx)
}

func (service *processingPlanService) GetStageByID(ctx context.Context, id, stageID string) (*ProcessingStage, *resty.Response, error) {
	path := fmt.Sprintf(EndpointProcessingPlanStagesID, id, stageID)
	return NewRequestBuilder[ProcessingStage](service.client, path).Get(ctx)
}

func (service *processingPlanService) UpdateStage(ctx context.Context, id, stageID string, stage *ProcessingStage) (*ProcessingStage, *resty.Response, error) {
	path := fmt.Sprintf(EndpointProcessingPlanStagesID, id, stageID)
	return NewRequestBuilder[ProcessingStage](service.client, path).Put(ctx, stage)
}

func (service *processingPlanService) GetMaterials(ctx context.Context, id string) (*List[ProcessingPlanMaterial], *resty.Response, error) {
	path := fmt.Sprintf(EndpointProcessingPlanMaterials, id)
	return NewRequestBuilder[List[ProcessingPlanMaterial]](service.client, path).Get(ctx)
}

func (service *processingPlanService) CreateMaterial(ctx context.Context, id string, material *ProcessingPlanMaterial) (*ProcessingPlanMaterial, *resty.Response, error) {
	path := fmt.Sprintf(EndpointProcessingPlanMaterials, id)
	return NewRequestBuilder[ProcessingPlanMaterial](service.client, path).Post(ctx, material)
}

func (service *processingPlanService) GetMaterialByID(ctx context.Context, id, materialID string) (*ProcessingPlanMaterial, *resty.Response, error) {
	path := fmt.Sprintf(EndpointProcessingPlanMaterialsID, id, materialID)
	return NewRequestBuilder[ProcessingPlanMaterial](service.client, path).Get(ctx)
}

func (service *processingPlanService) UpdateMaterial(ctx context.Context, id, materialID string, material *ProcessingPlanMaterial) (*ProcessingPlanMaterial, *resty.Response, error) {
	path := fmt.Sprintf(EndpointProcessingPlanMaterialsID, id, materialID)
	return NewRequestBuilder[ProcessingPlanMaterial](service.client, path).Put(ctx, material)
}

func (service *processingPlanService) DeleteMaterial(ctx context.Context, id, materialID string) (bool, *resty.Response, error) {
	path := fmt.Sprintf(EndpointProcessingPlanMaterialsID, id, materialID)
	return NewRequestBuilder[any](service.client, path).Delete(ctx)
}

func (service *processingPlanService) GetProducts(ctx context.Context, id string) (*List[ProcessingPlanProduct], *resty.Response, error) {
	path := fmt.Sprintf(EndpointProcessingPlanProducts, id)
	return NewRequestBuilder[List[ProcessingPlanProduct]](service.client, path).Get(ctx)
}

func (service *processingPlanService) CreateProduct(ctx context.Context, id string, product *ProcessingPlanProduct) (*ProcessingPlanProduct, *resty.Response, error) {
	path := fmt.Sprintf(EndpointProcessingPlanProducts, id)
	return NewRequestBuilder[ProcessingPlanProduct](service.client, path).Post(ctx, product)
}

func (service *processingPlanService) GetProductByID(ctx context.Context, id, productID string) (*ProcessingPlanProduct, *resty.Response, error) {
	path := fmt.Sprintf(EndpointProcessingPlanProductsID, id, productID)
	return NewRequestBuilder[ProcessingPlanProduct](service.client, path).Get(ctx)
}

func (service *processingPlanService) UpdateProduct(ctx context.Context, id, productID string, product *ProcessingPlanProduct) (*ProcessingPlanProduct, *resty.Response, error) {
	path := fmt.Sprintf(EndpointProcessingPlanProductsID, id, productID)
	return NewRequestBuilder[ProcessingPlanProduct](service.client, path).Put(ctx, product)
}

func (service *processingPlanService) DeleteProduct(ctx context.Context, id, productID string) (bool, *resty.Response, error) {
	path := fmt.Sprintf(EndpointProcessingPlanProductsID, id, productID)
	return NewRequestBuilder[any](service.client, path).Delete(ctx)
}

func NewProcessingPlanService(client *Client) ProcessingPlanService {
	e := NewEndpoint(client, EndpointProcessingPlan)
	return &processingPlanService{
		Endpoint:                 e,
		endpointGetList:          endpointGetList[ProcessingPlan]{e},
		endpointCreate:           endpointCreate[ProcessingPlan]{e},
		endpointCreateUpdateMany: endpointCreateUpdateMany[ProcessingPlan]{e},
		endpointDeleteMany:       endpointDeleteMany[ProcessingPlan]{e},
		endpointDeleteByID:       endpointDeleteByID{e},
		endpointDelete:           endpointDelete[ProcessingPlan]{e},
		endpointGetByID:          endpointGetByID[ProcessingPlan]{e},
		endpointUpdate:           endpointUpdate[ProcessingPlan]{e},
		endpointPositions:        endpointPositions[ProcessingPlanProduct]{e},
		endpointTrash:            endpointTrash{e},
	}
}
