package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// ProductionTask Производственное задание
// Ключевое слово: productiontask
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie
type ProductionTask struct {
	Moment                *Timestamp                       `json:"moment,omitempty"`
	Created               *Timestamp                       `json:"created,omitempty"`
	AccountID             *uuid.UUID                       `json:"accountId,omitempty"`
	Code                  *string                          `json:"code,omitempty"`
	Name                  *string                          `json:"name,omitempty"`
	Deleted               *Timestamp                       `json:"deleted,omitempty"`
	DeliveryPlannedMoment *Timestamp                       `json:"deliveryPlannedMoment,omitempty"`
	Description           *string                          `json:"description,omitempty"`
	ExternalCode          *string                          `json:"externalCode,omitempty"`
	Files                 *MetaArray[File]                 `json:"files,omitempty"`
	Group                 *Group                           `json:"group,omitempty"`
	Organization          *Organization                    `json:"organization,omitempty"`
	MaterialsStore        *Store                           `json:"materialsStore,omitempty"`
	Meta                  *Meta                            `json:"meta,omitempty"`
	Updated               *Timestamp                       `json:"updated,omitempty"`
	Applicable            *bool                            `json:"applicable,omitempty"`
	ID                    *uuid.UUID                       `json:"id,omitempty"`
	Owner                 *Employee                        `json:"owner,omitempty"`
	Printed               *bool                            `json:"printed,omitempty"`
	ProductionRows        *Positions[ProductionRow]        `json:"productionRows,omitempty"`
	ProductionEnd         *Timestamp                       `json:"productionEnd,omitempty"`
	ProductionStart       *Timestamp                       `json:"productionStart,omitempty"`
	Products              *Positions[ProductionTaskResult] `json:"products,omitempty"`
	ProductsStore         *Store                           `json:"productsStore,omitempty"`
	Published             *bool                            `json:"published,omitempty"`
	Reserve               *bool                            `json:"reserve,omitempty"`
	Shared                *bool                            `json:"shared,omitempty"`
	State                 *NullValue[State]                `json:"state,omitempty"`
	Attributes            Slice[Attribute]                 `json:"attributes,omitempty"`
}

// Clean возвращает сущность с единственным заполненным полем Meta
func (productionTask ProductionTask) Clean() *ProductionTask {
	return &ProductionTask{Meta: productionTask.Meta}
}

// AsTaskOperation реализует интерфейс AsTaskOperationInterface
func (productionTask ProductionTask) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: productionTask.Meta}
}

func (productionTask ProductionTask) GetMoment() Timestamp {
	return Deref(productionTask.Moment)
}

func (productionTask ProductionTask) GetCreated() Timestamp {
	return Deref(productionTask.Created)
}

func (productionTask ProductionTask) GetAccountID() uuid.UUID {
	return Deref(productionTask.AccountID)
}

func (productionTask ProductionTask) GetCode() string {
	return Deref(productionTask.Code)
}

func (productionTask ProductionTask) GetName() string {
	return Deref(productionTask.Name)
}

func (productionTask ProductionTask) GetDeleted() Timestamp {
	return Deref(productionTask.Deleted)
}

func (productionTask ProductionTask) GetDeliveryPlannedMoment() Timestamp {
	return Deref(productionTask.DeliveryPlannedMoment)
}

func (productionTask ProductionTask) GetDescription() string {
	return Deref(productionTask.Description)
}

func (productionTask ProductionTask) GetExternalCode() string {
	return Deref(productionTask.ExternalCode)
}

func (productionTask ProductionTask) GetFiles() MetaArray[File] {
	return Deref(productionTask.Files)
}

func (productionTask ProductionTask) GetGroup() Group {
	return Deref(productionTask.Group)
}

func (productionTask ProductionTask) GetOrganization() Organization {
	return Deref(productionTask.Organization)
}

func (productionTask ProductionTask) GetMaterialsStore() Store {
	return Deref(productionTask.MaterialsStore)
}

func (productionTask ProductionTask) GetMeta() Meta {
	return Deref(productionTask.Meta)
}

func (productionTask ProductionTask) GetUpdated() Timestamp {
	return Deref(productionTask.Updated)
}

func (productionTask ProductionTask) GetApplicable() bool {
	return Deref(productionTask.Applicable)
}

func (productionTask ProductionTask) GetID() uuid.UUID {
	return Deref(productionTask.ID)
}

func (productionTask ProductionTask) GetOwner() Employee {
	return Deref(productionTask.Owner)
}

func (productionTask ProductionTask) GetPrinted() bool {
	return Deref(productionTask.Printed)
}

func (productionTask ProductionTask) GetProductionRows() Positions[ProductionRow] {
	return Deref(productionTask.ProductionRows)
}

func (productionTask ProductionTask) GetProductionEnd() Timestamp {
	return Deref(productionTask.ProductionEnd)
}

func (productionTask ProductionTask) GetProductionStart() Timestamp {
	return Deref(productionTask.ProductionStart)
}

func (productionTask ProductionTask) GetProducts() Positions[ProductionTaskResult] {
	return Deref(productionTask.Products)
}

func (productionTask ProductionTask) GetProductsStore() Store {
	return Deref(productionTask.ProductsStore)
}

func (productionTask ProductionTask) GetPublished() bool {
	return Deref(productionTask.Published)
}

func (productionTask ProductionTask) GetReserve() bool {
	return Deref(productionTask.Reserve)
}

func (productionTask ProductionTask) GetShared() bool {
	return Deref(productionTask.Shared)
}

func (productionTask ProductionTask) GetState() State {
	return productionTask.State.Get()
}

func (productionTask ProductionTask) GetAttributes() Slice[Attribute] {
	return productionTask.Attributes
}

func (productionTask *ProductionTask) SetMoment(moment *Timestamp) *ProductionTask {
	productionTask.Moment = moment
	return productionTask
}

func (productionTask *ProductionTask) SetCode(code string) *ProductionTask {
	productionTask.Code = &code
	return productionTask
}

func (productionTask *ProductionTask) SetName(name string) *ProductionTask {
	productionTask.Name = &name
	return productionTask
}

func (productionTask *ProductionTask) SetDeliveryPlannedMoment(deliveryPlannedMoment *Timestamp) *ProductionTask {
	productionTask.DeliveryPlannedMoment = deliveryPlannedMoment
	return productionTask
}

func (productionTask *ProductionTask) SetDescription(description string) *ProductionTask {
	productionTask.Description = &description
	return productionTask
}

func (productionTask *ProductionTask) SetExternalCode(externalCode string) *ProductionTask {
	productionTask.ExternalCode = &externalCode
	return productionTask
}

func (productionTask *ProductionTask) SetFiles(files ...*File) *ProductionTask {
	productionTask.Files = NewMetaArrayFrom(files)
	return productionTask
}

func (productionTask *ProductionTask) SetGroup(group *Group) *ProductionTask {
	productionTask.Group = group.Clean()
	return productionTask
}

func (productionTask *ProductionTask) SetOrganization(organization *Organization) *ProductionTask {
	productionTask.Organization = organization.Clean()
	return productionTask
}

func (productionTask *ProductionTask) SetMaterialsStore(materialsStore *Store) *ProductionTask {
	productionTask.MaterialsStore = materialsStore.Clean()
	return productionTask
}

func (productionTask *ProductionTask) SetMeta(meta *Meta) *ProductionTask {
	productionTask.Meta = meta
	return productionTask
}

func (productionTask *ProductionTask) SetApplicable(applicable bool) *ProductionTask {
	productionTask.Applicable = &applicable
	return productionTask
}

func (productionTask *ProductionTask) SetOwner(owner *Employee) *ProductionTask {
	productionTask.Owner = owner.Clean()
	return productionTask
}

func (productionTask *ProductionTask) SetProductionRows(productionRows ...*ProductionRow) *ProductionTask {
	productionTask.ProductionRows = NewPositionsFrom(productionRows)
	return productionTask
}

func (productionTask *ProductionTask) SetProductionStart(productionStart *Timestamp) *ProductionTask {
	productionTask.ProductionStart = productionStart
	return productionTask
}

func (productionTask *ProductionTask) SetProducts(products ...*ProductionTaskResult) *ProductionTask {
	productionTask.Products = NewPositionsFrom(products)
	return productionTask
}

func (productionTask *ProductionTask) SetProductsStore(productsStore *Store) *ProductionTask {
	productionTask.ProductsStore = productsStore.Clean()
	return productionTask
}

func (productionTask *ProductionTask) SetReserve(reserve bool) *ProductionTask {
	productionTask.Reserve = &reserve
	return productionTask
}

func (productionTask *ProductionTask) SetShared(shared bool) *ProductionTask {
	productionTask.Shared = &shared
	return productionTask
}

func (productionTask *ProductionTask) SetState(state *State) *ProductionTask {
	productionTask.State = NewNullValueFrom(state.Clean())
	return productionTask
}

func (productionTask *ProductionTask) SetNullState() *ProductionTask {
	productionTask.State = NewNullValue[State]()
	return productionTask
}

func (productionTask *ProductionTask) SetAttributes(attributes ...*Attribute) *ProductionTask {
	productionTask.Attributes = attributes
	return productionTask
}

func (productionTask ProductionTask) String() string {
	return Stringify(productionTask)
}

// MetaType возвращает тип сущности.
func (ProductionTask) MetaType() MetaType {
	return MetaTypeProductionTask
}

// Update shortcut
func (productionTask ProductionTask) Update(ctx context.Context, client *Client, params ...*Params) (*ProductionTask, *resty.Response, error) {
	return client.Entity().ProductionTask().Update(ctx, productionTask.GetID(), &productionTask, params...)
}

// Create shortcut
func (productionTask ProductionTask) Create(ctx context.Context, client *Client, params ...*Params) (*ProductionTask, *resty.Response, error) {
	return client.Entity().ProductionTask().Create(ctx, &productionTask, params...)
}

// Delete shortcut
func (productionTask ProductionTask) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return client.Entity().ProductionTask().Delete(ctx, productionTask.GetID())
}

// ProductionRow Позиция производственного задания
// Ключевое слово: productionrow
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-proizwodstwennye-zadaniq-pozicii-proizwodstwennogo-zadaniq
type ProductionRow struct {
	AccountID        *uuid.UUID      `json:"accountId,omitempty"`        // ID учетной записи
	ExternalCode     *string         `json:"externalCode,omitempty"`     // Внешний код
	ID               *uuid.UUID      `json:"id,omitempty"`               // ID позиции
	Name             *string         `json:"name,omitempty"`             // Наименование
	ProcessingPlan   *ProcessingPlan `json:"processingPlan,omitempty"`   // Метаданные Техкарты
	ProductionVolume *float64        `json:"productionVolume,omitempty"` // Объем производства.
	Updated          *Timestamp      `json:"updated,omitempty"`          // Момент последнего обновления Производственного задания
}

func (productionRow ProductionRow) GetAccountID() uuid.UUID {
	return Deref(productionRow.AccountID)
}

func (productionRow ProductionRow) GetExternalCode() string {
	return Deref(productionRow.ExternalCode)
}

func (productionRow ProductionRow) GetID() uuid.UUID {
	return Deref(productionRow.ID)
}

func (productionRow ProductionRow) GetName() string {
	return Deref(productionRow.Name)
}

func (productionRow ProductionRow) GetProcessingPlan() ProcessingPlan {
	return Deref(productionRow.ProcessingPlan)
}

func (productionRow ProductionRow) GetProductionVolume() float64 {
	return Deref(productionRow.ProductionVolume)
}

func (productionRow ProductionRow) GetUpdated() Timestamp {
	return Deref(productionRow.Updated)
}

func (productionRow *ProductionRow) SetExternalCode(externalCode string) *ProductionRow {
	productionRow.ExternalCode = &externalCode
	return productionRow
}

func (productionRow *ProductionRow) SetName(name string) *ProductionRow {
	productionRow.Name = &name
	return productionRow
}

func (productionRow *ProductionRow) SetProcessingPlan(processingPlan *ProcessingPlan) *ProductionRow {
	productionRow.ProcessingPlan = processingPlan.Clean()
	return productionRow
}

func (productionRow *ProductionRow) SetProductionVolume(productionVolume float64) *ProductionRow {
	productionRow.ProductionVolume = &productionVolume
	return productionRow
}

func (productionRow ProductionRow) String() string {
	return Stringify(productionRow)
}

// MetaType возвращает тип сущности.
func (ProductionRow) MetaType() MetaType {
	return MetaTypeProductionRow
}

// ProductionTaskResult Продукт производственного задания
// Ключевое слово: productiontaskresult
// https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-proizwodstwennye-zadaniq-produkty-proizwodstwennogo-zadaniq
type ProductionTaskResult struct {
	AccountID     *uuid.UUID          `json:"accountId,omitempty"`     // ID учетной записи
	Assortment    *AssortmentPosition `json:"assortment,omitempty"`    // Ссылка на товар/серию/модификацию, которую представляет собой позиция.
	ID            *uuid.UUID          `json:"id,omitempty"`            // ID позиции
	PlanQuantity  *float64            `json:"planQuantity,omitempty"`  // Запланированное для производства количество продукта
	ProductionRow *ProductionRow      `json:"productionRow,omitempty"` // Метаданные Позиции производственного задания
}

func (productionTaskResult ProductionTaskResult) GetAccountID() uuid.UUID {
	return Deref(productionTaskResult.AccountID)
}

func (productionTaskResult ProductionTaskResult) GetAssortment() AssortmentPosition {
	return Deref(productionTaskResult.Assortment)
}

func (productionTaskResult ProductionTaskResult) GetID() uuid.UUID {
	return Deref(productionTaskResult.ID)
}

func (productionTaskResult ProductionTaskResult) GetPlanQuantity() float64 {
	return Deref(productionTaskResult.PlanQuantity)
}

func (productionTaskResult ProductionTaskResult) GetProductionRow() ProductionRow {
	return Deref(productionTaskResult.ProductionRow)
}

func (productionTaskResult *ProductionTaskResult) SetAssortment(assortment AsAssortment) *ProductionTaskResult {
	productionTaskResult.Assortment = assortment.AsAssortment()
	return productionTaskResult
}

func (productionTaskResult *ProductionTaskResult) SetPlanQuantity(planQuantity float64) *ProductionTaskResult {
	productionTaskResult.PlanQuantity = &planQuantity
	return productionTaskResult
}

func (productionTaskResult ProductionTaskResult) String() string {
	return Stringify(productionTaskResult)
}

// MetaType возвращает тип сущности.
func (ProductionTaskResult) MetaType() MetaType {
	return MetaTypeProductionTaskResult
}

// ProductionTaskService
// Сервис для работы с производственными заданиями
type ProductionTaskService interface {
	GetList(ctx context.Context, params ...*Params) (*List[ProductionTask], *resty.Response, error)
	Create(ctx context.Context, productionTask *ProductionTask, params ...*Params) (*ProductionTask, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, productionTaskList Slice[ProductionTask], params ...*Params) (*Slice[ProductionTask], *resty.Response, error)
	DeleteMany(ctx context.Context, entities ...*ProductionTask) (*DeleteManyResponse, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetaAttributesSharedStatesWrapper, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributeMany(ctx context.Context, attributes ...*Attribute) (*Slice[Attribute], *resty.Response, error)
	UpdateAttribute(ctx context.Context, id uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributeMany(ctx context.Context, attributes ...*Attribute) (*DeleteManyResponse, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*ProductionTask, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, productionTask *ProductionTask, params ...*Params) (*ProductionTask, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetPositions(ctx context.Context, id uuid.UUID, params ...*Params) (*MetaArray[ProductionRow], *resty.Response, error)
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params ...*Params) (*ProductionRow, *resty.Response, error)
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *ProductionRow, params ...*Params) (*ProductionRow, *resty.Response, error)
	CreatePosition(ctx context.Context, id uuid.UUID, position *ProductionRow) (*ProductionRow, *resty.Response, error)
	CreatePositionMany(ctx context.Context, id uuid.UUID, positions ...*ProductionRow) (*Slice[ProductionRow], *resty.Response, error)
	DeletePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (bool, *resty.Response, error)
	DeletePositionMany(ctx context.Context, id uuid.UUID, entities ...*ProductionRow) (*DeleteManyResponse, *resty.Response, error)
	GetPositionTrackingCodes(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (*MetaArray[TrackingCode], *resty.Response, error)
	CreateUpdatePositionTrackingCodeMany(ctx context.Context, id uuid.UUID, positionID uuid.UUID, trackingCodes ...*TrackingCode) (*Slice[TrackingCode], *resty.Response, error)
	DeletePositionTrackingCodeMany(ctx context.Context, id uuid.UUID, positionID uuid.UUID, trackingCodes ...*TrackingCode) (*DeleteManyResponse, *resty.Response, error)
	GetProducts(ctx context.Context, id uuid.UUID, params ...*Params) (*MetaArray[ProductionTaskResult], *resty.Response, error)
	GetProductByID(ctx context.Context, id uuid.UUID, productID uuid.UUID, params ...*Params) (*ProductionTaskResult, *resty.Response, error)
	CreateProduct(ctx context.Context, id uuid.UUID, productionTaskResult *ProductionTaskResult, params ...*Params) (*ProductionTaskResult, *resty.Response, error)
	UpdateProduct(ctx context.Context, id uuid.UUID, productID uuid.UUID, productionTaskResult *ProductionTaskResult, params ...*Params) (*ProductionTaskResult, *resty.Response, error)
	DeleteProduct(ctx context.Context, id uuid.UUID, productID uuid.UUID) (bool, *resty.Response, error)
	DeleteProductMany(ctx context.Context, id uuid.UUID) (*DeleteManyResponse, *resty.Response, error)
	GetFiles(ctx context.Context, id uuid.UUID) (*MetaArray[File], *resty.Response, error)
	CreateFile(ctx context.Context, id uuid.UUID, file *File) (*Slice[File], *resty.Response, error)
	UpdateFileMany(ctx context.Context, id uuid.UUID, files ...*File) (*Slice[File], *resty.Response, error)
	DeleteFile(ctx context.Context, id uuid.UUID, fileID uuid.UUID) (bool, *resty.Response, error)
	DeleteFileMany(ctx context.Context, id uuid.UUID, files ...*File) (*DeleteManyResponse, *resty.Response, error)
}

type productionTaskService struct {
	Endpoint
	endpointGetList[ProductionTask]
	endpointCreate[ProductionTask]
	endpointCreateUpdateMany[ProductionTask]
	endpointDeleteMany[ProductionTask]
	endpointMetadata[MetaAttributesSharedStatesWrapper]
	endpointAttributes
	endpointGetByID[ProductionTask]
	endpointUpdate[ProductionTask]
	endpointDelete
	endpointPositions[ProductionRow]
	endpointStates
	endpointFiles
}

func NewProductionTaskService(client *Client) ProductionTaskService {
	e := NewEndpoint(client, "entity/productiontask")
	return &productionTaskService{
		Endpoint:                 e,
		endpointGetList:          endpointGetList[ProductionTask]{e},
		endpointCreate:           endpointCreate[ProductionTask]{e},
		endpointCreateUpdateMany: endpointCreateUpdateMany[ProductionTask]{e},
		endpointDeleteMany:       endpointDeleteMany[ProductionTask]{e},
		endpointMetadata:         endpointMetadata[MetaAttributesSharedStatesWrapper]{e},
		endpointGetByID:          endpointGetByID[ProductionTask]{e},
		endpointUpdate:           endpointUpdate[ProductionTask]{e},
		endpointPositions:        endpointPositions[ProductionRow]{e},
		endpointDelete:           endpointDelete{e},
		endpointAttributes:       endpointAttributes{e},
		endpointStates:           endpointStates{e},
		endpointFiles:            endpointFiles{e},
	}
}

// GetProducts Получить Продукты производственного задания.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-poluchit-produkty-proizwodstwennogo-zadaniq
func (service *productionTaskService) GetProducts(ctx context.Context, id uuid.UUID, params ...*Params) (*MetaArray[ProductionTaskResult], *resty.Response, error) {
	path := fmt.Sprintf("%s/products", id)
	return NewRequestBuilder[MetaArray[ProductionTaskResult]](service.client, path).SetParams(params...).Get(ctx)
}

// GetProductByID Получить продукт производственного задания.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-produkt-proizwodstwennogo-zadaniq
func (service *productionTaskService) GetProductByID(ctx context.Context, id uuid.UUID, productID uuid.UUID, params ...*Params) (*ProductionTaskResult, *resty.Response, error) {
	path := fmt.Sprintf("%s/products/%s", id, productID)
	return NewRequestBuilder[ProductionTaskResult](service.client, path).SetParams(params...).Get(ctx)
}

// CreateProduct Создать продукт.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-sozdat-produkt
func (service *productionTaskService) CreateProduct(ctx context.Context, id uuid.UUID, productionTaskResult *ProductionTaskResult, params ...*Params) (*ProductionTaskResult, *resty.Response, error) {
	path := fmt.Sprintf("%s/products", id) // fixme:в документации указан endpoint без 's' на конце, вероятно ошибка
	return NewRequestBuilder[ProductionTaskResult](service.client, path).SetParams(params...).Post(ctx, productionTaskResult)
}

// UpdateProduct Изменить продукт.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-izmenit-produkt
func (service *productionTaskService) UpdateProduct(ctx context.Context, id uuid.UUID, productID uuid.UUID, productionTaskResult *ProductionTaskResult, params ...*Params) (*ProductionTaskResult, *resty.Response, error) {
	path := fmt.Sprintf("%s/products/%s", id, productID)
	return NewRequestBuilder[ProductionTaskResult](service.client, path).SetParams(params...).Put(ctx, productionTaskResult)
}

// DeleteProduct Удалить продукт.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-udalit-produkt
func (service *productionTaskService) DeleteProduct(ctx context.Context, id uuid.UUID, productID uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/products/%s", id, productID)
	return NewRequestBuilder[any](service.client, path).Delete(ctx)
}

// DeleteProductMany Массовое удаление продуктов Производственного задания.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-udalit-produkt
func (service *productionTaskService) DeleteProductMany(ctx context.Context, id uuid.UUID) (*DeleteManyResponse, *resty.Response, error) {
	path := fmt.Sprintf("%s/products/delete", id)
	return NewRequestBuilder[DeleteManyResponse](service.client, path).Post(ctx, nil)
}
