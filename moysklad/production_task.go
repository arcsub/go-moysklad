package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"time"
)

// ProductionTask Производственное задание.
//
// Код сущности: productiontask
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie
type ProductionTask struct {
	Moment                *Timestamp                       `json:"moment,omitempty"` // Дата документа
	Created               *Timestamp                       `json:"created,omitempty"`
	AccountID             *uuid.UUID                       `json:"accountId,omitempty"` // ID учётной записи
	Code                  *string                          `json:"code,omitempty"`
	Name                  *string                          `json:"name,omitempty"`
	Deleted               *Timestamp                       `json:"deleted,omitempty"`
	DeliveryPlannedMoment *Timestamp                       `json:"deliveryPlannedMoment,omitempty"`
	Description           *string                          `json:"description,omitempty"`
	ExternalCode          *string                          `json:"externalCode,omitempty"`
	Files                 *MetaArray[File]                 `json:"files,omitempty"` // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group                 *Group                           `json:"group,omitempty"` // Отдел сотрудника
	Organization          *Organization                    `json:"organization,omitempty"`
	MaterialsStore        *Store                           `json:"materialsStore,omitempty"`
	Meta                  *Meta                            `json:"meta,omitempty"`
	Updated               *Timestamp                       `json:"updated,omitempty"`
	Applicable            *bool                            `json:"applicable,omitempty"` // Отметка о проведении
	ID                    *uuid.UUID                       `json:"id,omitempty"`
	Owner                 *Employee                        `json:"owner,omitempty"`   // Метаданные владельца (Сотрудника)
	Printed               *bool                            `json:"printed,omitempty"` // Напечатан ли документ
	ProductionRows        *MetaArray[ProductionRow]        `json:"productionRows,omitempty"`
	ProductionEnd         *Timestamp                       `json:"productionEnd,omitempty"`
	ProductionStart       *Timestamp                       `json:"productionStart,omitempty"`
	Products              *MetaArray[ProductionTaskResult] `json:"products,omitempty"`
	ProductsStore         *Store                           `json:"productsStore,omitempty"`
	Published             *bool                            `json:"published,omitempty"` // Опубликован ли документ
	Reserve               *bool                            `json:"reserve,omitempty"`
	Shared                *bool                            `json:"shared,omitempty"` // Общий доступ
	State                 *NullValue[State]                `json:"state,omitempty"`
	Attributes            Slice[Attribute]                 `json:"attributes,omitempty"` // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (productionTask ProductionTask) Clean() *ProductionTask {
	if productionTask.Meta == nil {
		return nil
	}
	return &ProductionTask{Meta: productionTask.Meta}
}

// AsTaskOperation реализует интерфейс [TaskOperationConverter].
func (productionTask ProductionTask) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: productionTask.Meta}
}

func (productionTask ProductionTask) GetMoment() time.Time {
	return Deref(productionTask.Moment).Time()
}

func (productionTask ProductionTask) GetCreated() time.Time {
	return Deref(productionTask.Created).Time()
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

func (productionTask ProductionTask) GetDeleted() time.Time {
	return Deref(productionTask.Deleted).Time()
}

func (productionTask ProductionTask) GetDeliveryPlannedMoment() time.Time {
	return Deref(productionTask.DeliveryPlannedMoment).Time()
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

func (productionTask ProductionTask) GetUpdated() time.Time {
	return Deref(productionTask.Updated).Time()
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

func (productionTask ProductionTask) GetProductionRows() MetaArray[ProductionRow] {
	return Deref(productionTask.ProductionRows)
}

func (productionTask ProductionTask) GetProductionEnd() time.Time {
	return Deref(productionTask.ProductionEnd).Time()
}

func (productionTask ProductionTask) GetProductionStart() time.Time {
	return Deref(productionTask.ProductionStart).Time()
}

func (productionTask ProductionTask) GetProducts() MetaArray[ProductionTaskResult] {
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
	return Deref(productionTask.State).getValue()
}

func (productionTask ProductionTask) GetAttributes() Slice[Attribute] {
	return productionTask.Attributes
}

func (productionTask *ProductionTask) SetMoment(moment time.Time) *ProductionTask {
	productionTask.Moment = NewTimestamp(moment)
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

func (productionTask *ProductionTask) SetDeliveryPlannedMoment(deliveryPlannedMoment time.Time) *ProductionTask {
	productionTask.DeliveryPlannedMoment = NewTimestamp(deliveryPlannedMoment)
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
	productionTask.ProductionRows = NewMetaArrayFrom(productionRows)
	return productionTask
}

func (productionTask *ProductionTask) SetProductionStart(productionStart time.Time) *ProductionTask {
	productionTask.ProductionStart = NewTimestamp(productionStart)
	return productionTask
}

func (productionTask *ProductionTask) SetProducts(products ...*ProductionTaskResult) *ProductionTask {
	productionTask.Products = NewMetaArrayFrom(products)
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
	productionTask.State = NewNullValue(state.Clean())
	return productionTask
}

func (productionTask *ProductionTask) SetAttributes(attributes ...*Attribute) *ProductionTask {
	productionTask.Attributes.Push(attributes...)
	return productionTask
}

func (productionTask ProductionTask) String() string {
	return Stringify(productionTask)
}

// MetaType возвращает код сущности.
func (ProductionTask) MetaType() MetaType {
	return MetaTypeProductionTask
}

// Update shortcut
func (productionTask *ProductionTask) Update(ctx context.Context, client *Client, params ...*Params) (*ProductionTask, *resty.Response, error) {
	return NewProductionTaskService(client).Update(ctx, productionTask.GetID(), productionTask, params...)
}

// Create shortcut
func (productionTask *ProductionTask) Create(ctx context.Context, client *Client, params ...*Params) (*ProductionTask, *resty.Response, error) {
	return NewProductionTaskService(client).Create(ctx, productionTask, params...)
}

// Delete shortcut
func (productionTask *ProductionTask) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewProductionTaskService(client).Delete(ctx, productionTask)
}

// ProductionRow Позиция производственного задания
//
// Код сущности: productionrow
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-proizwodstwennye-zadaniq-pozicii-proizwodstwennogo-zadaniq
type ProductionRow struct {
	AccountID        *uuid.UUID      `json:"accountId,omitempty"`        // ID учётной записи
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

func (productionRow ProductionRow) GetUpdated() time.Time {
	return Deref(productionRow.Updated).Time()
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

// MetaType возвращает код сущности.
func (ProductionRow) MetaType() MetaType {
	return MetaTypeProductionRow
}

// ProductionTaskResult Продукт производственного задания
//
// Код сущности: productiontaskresult
// https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-proizwodstwennye-zadaniq-produkty-proizwodstwennogo-zadaniq
type ProductionTaskResult struct {
	AccountID     *uuid.UUID          `json:"accountId,omitempty"`     // ID учётной записи
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

// Принимает объект, реализующий интерфейс [AssortmentConverter].
func (productionTaskResult *ProductionTaskResult) SetAssortment(assortment AssortmentConverter) *ProductionTaskResult {
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

// MetaType возвращает код сущности.
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
	GetMetadata(ctx context.Context) (*MetaAttributesStatesSharedWrapper, *resty.Response, error)

	// GetAttributeList выполняет запрос на получение списка доп полей.
	// Принимает контекст.
	// Возвращает объект List.
	GetAttributeList(ctx context.Context) (*List[Attribute], *resty.Response, error)

	// GetAttributeByID выполняет запрос на получение отдельного доп поля по ID.
	// Принимает контекст и ID доп поля.
	// Возвращает найденное доп поле.
	GetAttributeByID(ctx context.Context, id uuid.UUID) (*Attribute, *resty.Response, error)

	// CreateAttribute выполняет запрос на создание доп поля.
	// Принимает контекст и доп поле.
	// Возвращает созданное доп поле.
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)

	// CreateUpdateAttributeMany выполняет запрос на массовое создание и/или изменение доп полей.
	// Изменяемые доп поля должны содержать идентификатор в виде метаданных.
	// Принимает контекст и множество доп полей.
	// Возвращает список созданных и/или изменённых доп полей.
	CreateUpdateAttributeMany(ctx context.Context, attributes ...*Attribute) (*Slice[Attribute], *resty.Response, error)

	// UpdateAttribute выполняет запрос на изменения доп поля.
	// Принимает контекст, ID доп поля и доп поле.
	// Возвращает изменённое доп поле.
	UpdateAttribute(ctx context.Context, id uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)

	// DeleteAttribute выполняет запрос на удаление доп поля.
	// Принимает контекст и ID доп поля.
	// Возвращает «true» в случае успешного удаления доп поля.
	DeleteAttribute(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// DeleteAttributeMany выполняет запрос на массовое удаление доп полей.
	// Принимает контекст и множество доп полей.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteAttributeMany(ctx context.Context, attributes ...*Attribute) (*DeleteManyResponse, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*ProductionTask, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, productionTask *ProductionTask, params ...*Params) (*ProductionTask, *resty.Response, error)
	DeleteByID(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// Delete выполняет запрос на удаление производственного задания.
	// Принимает контекст и производственное задание.
	// Возвращает «true» в случае успешного удаления производственного задания.
	Delete(ctx context.Context, entity *ProductionTask) (bool, *resty.Response, error)

	// GetPositionList выполняет запрос на получение списка позиций документа.
	// Принимает контекст, ID документа и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetPositionList(ctx context.Context, id uuid.UUID, params ...*Params) (*List[ProductionRow], *resty.Response, error)

	GetPositionListAll(ctx context.Context, id uuid.UUID, params ...*Params) (*Slice[ProductionRow], *resty.Response, error)

	// GetPositionByID выполняет запрос на получение отдельной позиции документа по ID.
	// Принимает контекст, ID документа, ID позиции и опционально объект параметров запроса Params.
	// Возвращает найденную позицию.
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params ...*Params) (*ProductionRow, *resty.Response, error)

	// UpdatePosition выполняет запрос на изменение позиции документа.
	// Принимает контекст, ID документа, ID позиции, позицию документа и опционально объект параметров запроса Params.
	// Возвращает изменённую позицию.
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *ProductionRow, params ...*Params) (*ProductionRow, *resty.Response, error)
	CreatePosition(ctx context.Context, id uuid.UUID, position *ProductionRow, params ...*Params) (*ProductionRow, *resty.Response, error)

	// CreatePositionMany выполняет запрос на массовое добавление позиций документа.
	// Принимает контекст, ID документа и множество позиций.
	// Возвращает список добавленных позиций.
	CreatePositionMany(ctx context.Context, id uuid.UUID, positions ...*ProductionRow) (*Slice[ProductionRow], *resty.Response, error)

	// DeletePosition выполняет запрос на удаление позиции документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает «true» в случае успешного удаления позиции.
	DeletePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (bool, *resty.Response, error)

	// DeletePositionMany выполняет запрос на массовое удаление позиций документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeletePositionMany(ctx context.Context, id uuid.UUID, positions ...*ProductionRow) (*DeleteManyResponse, *resty.Response, error)

	// GetPositionTrackingCodeList выполняет запрос на получение кодов маркировки позиции документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает объект List.
	GetPositionTrackingCodeList(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (*List[TrackingCode], *resty.Response, error)

	// CreateUpdatePositionTrackingCodeMany выполняет запрос на массовое создание/изменение кодов маркировки позиции документа.
	// Принимает контекст, ID документа, ID позиции и множество кодов маркировки.
	// Возвращает список созданных и/или изменённых кодов маркировки позиции документа.
	CreateUpdatePositionTrackingCodeMany(ctx context.Context, id uuid.UUID, positionID uuid.UUID, trackingCodes ...*TrackingCode) (*Slice[TrackingCode], *resty.Response, error)

	// DeletePositionTrackingCodeMany выполняет запрос на массовое удаление кодов маркировки позиции документа.
	// Принимает контекст, ID документа, ID позиции и множество кодов маркировки.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeletePositionTrackingCodeMany(ctx context.Context, id uuid.UUID, positionID uuid.UUID, trackingCodes ...*TrackingCode) (*DeleteManyResponse, *resty.Response, error)
	GetProducts(ctx context.Context, id uuid.UUID, params ...*Params) (*MetaArray[ProductionTaskResult], *resty.Response, error)
	GetProductByID(ctx context.Context, id uuid.UUID, productID uuid.UUID, params ...*Params) (*ProductionTaskResult, *resty.Response, error)
	CreateProduct(ctx context.Context, id uuid.UUID, productionTaskResult *ProductionTaskResult, params ...*Params) (*ProductionTaskResult, *resty.Response, error)
	UpdateProduct(ctx context.Context, id uuid.UUID, productID uuid.UUID, productionTaskResult *ProductionTaskResult, params ...*Params) (*ProductionTaskResult, *resty.Response, error)
	DeleteProduct(ctx context.Context, id uuid.UUID, productID uuid.UUID) (bool, *resty.Response, error)
	DeleteProductMany(ctx context.Context, id uuid.UUID) (*DeleteManyResponse, *resty.Response, error)

	// GetFileList выполняет запрос на получение файлов в виде списка.
	// Принимает контекст и ID сущности/документа.
	// Возвращает объект List.
	GetFileList(ctx context.Context, id uuid.UUID) (*List[File], *resty.Response, error)

	// CreateFile выполняет запрос на добавление файла.
	// Принимает контекст, ID сущности/документа и файл.
	// Возвращает список файлов.
	CreateFile(ctx context.Context, id uuid.UUID, file *File) (*Slice[File], *resty.Response, error)

	// UpdateFileMany выполняет запрос на массовое создание и/или изменение файлов сущности/документа.
	// Принимает контекст, ID сущности/документа и множество файлов.
	// Возвращает созданных и/или изменённых файлов.
	UpdateFileMany(ctx context.Context, id uuid.UUID, files ...*File) (*Slice[File], *resty.Response, error)

	// DeleteFile выполняет запрос на удаление файла сущности/документа.
	// Принимает контекст, ID сущности/документа и ID файла.
	// Возвращает «true» в случае успешного удаления файла.
	DeleteFile(ctx context.Context, id uuid.UUID, fileID uuid.UUID) (bool, *resty.Response, error)

	// DeleteFileMany выполняет запрос на массовое удаление файлов сущности/документа.
	// Принимает контекст, ID сущности/документа и множество файлов.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteFileMany(ctx context.Context, id uuid.UUID, files ...*File) (*DeleteManyResponse, *resty.Response, error)
}

const (
	EndpointProductionTask               = EndpointEntity + string(MetaTypeProductionTask)
	EndpointProductionTaskProducts       = EndpointProductionTask + "/%s/products"
	EndpointProductionTaskProductsID     = EndpointProductionTaskProducts + "/%s"
	EndpointProductionTaskProductsDelete = EndpointProductionTaskProducts + EndpointDelete
)

type productionTaskService struct {
	Endpoint
	endpointGetList[ProductionTask]
	endpointCreate[ProductionTask]
	endpointCreateUpdateMany[ProductionTask]
	endpointDeleteMany[ProductionTask]
	endpointMetadata[MetaAttributesStatesSharedWrapper]
	endpointAttributes
	endpointGetByID[ProductionTask]
	endpointUpdate[ProductionTask]
	endpointDeleteByID
	endpointDelete[ProductionTask]
	endpointPositions[ProductionRow]
	endpointStates
	endpointFiles
}

// GetProducts Получить Продукты производственного задания.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-poluchit-produkty-proizwodstwennogo-zadaniq
func (service *productionTaskService) GetProducts(ctx context.Context, id uuid.UUID, params ...*Params) (*MetaArray[ProductionTaskResult], *resty.Response, error) {
	path := fmt.Sprintf(EndpointProductionTaskProducts, id)
	return NewRequestBuilder[MetaArray[ProductionTaskResult]](service.client, path).SetParams(params...).Get(ctx)
}

// GetProductByID Получить продукт производственного задания.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-produkt-proizwodstwennogo-zadaniq
func (service *productionTaskService) GetProductByID(ctx context.Context, id uuid.UUID, productID uuid.UUID, params ...*Params) (*ProductionTaskResult, *resty.Response, error) {
	path := fmt.Sprintf(EndpointProductionTaskProductsID, id, productID)
	return NewRequestBuilder[ProductionTaskResult](service.client, path).SetParams(params...).Get(ctx)
}

// CreateProduct Создать продукт.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-sozdat-produkt
func (service *productionTaskService) CreateProduct(ctx context.Context, id uuid.UUID, productionTaskResult *ProductionTaskResult, params ...*Params) (*ProductionTaskResult, *resty.Response, error) {
	path := fmt.Sprintf(EndpointProductionTaskProducts, id) // fixme:в документации указан endpoint без 's' на конце, вероятно ошибка
	return NewRequestBuilder[ProductionTaskResult](service.client, path).SetParams(params...).Post(ctx, productionTaskResult)
}

// UpdateProduct Изменить продукт.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-izmenit-produkt
func (service *productionTaskService) UpdateProduct(ctx context.Context, id uuid.UUID, productID uuid.UUID, productionTaskResult *ProductionTaskResult, params ...*Params) (*ProductionTaskResult, *resty.Response, error) {
	path := fmt.Sprintf(EndpointProductionTaskProductsID, id, productID)
	return NewRequestBuilder[ProductionTaskResult](service.client, path).SetParams(params...).Put(ctx, productionTaskResult)
}

// DeleteProduct Удалить продукт.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-udalit-produkt
func (service *productionTaskService) DeleteProduct(ctx context.Context, id uuid.UUID, productID uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf(EndpointProductionTaskProductsID, id, productID)
	return NewRequestBuilder[any](service.client, path).Delete(ctx)
}

// DeleteProductMany Массовое удаление продуктов Производственного задания.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-udalit-produkt
func (service *productionTaskService) DeleteProductMany(ctx context.Context, id uuid.UUID) (*DeleteManyResponse, *resty.Response, error) {
	path := fmt.Sprintf(EndpointProductionTaskProductsDelete, id)
	return NewRequestBuilder[DeleteManyResponse](service.client, path).Post(ctx, nil)
}

func NewProductionTaskService(client *Client) ProductionTaskService {
	e := NewEndpoint(client, EndpointProductionTask)
	return &productionTaskService{
		Endpoint:                 e,
		endpointGetList:          endpointGetList[ProductionTask]{e},
		endpointCreate:           endpointCreate[ProductionTask]{e},
		endpointCreateUpdateMany: endpointCreateUpdateMany[ProductionTask]{e},
		endpointDeleteMany:       endpointDeleteMany[ProductionTask]{e},
		endpointMetadata:         endpointMetadata[MetaAttributesStatesSharedWrapper]{e},
		endpointGetByID:          endpointGetByID[ProductionTask]{e},
		endpointUpdate:           endpointUpdate[ProductionTask]{e},
		endpointPositions:        endpointPositions[ProductionRow]{e},
		endpointDeleteByID:       endpointDeleteByID{e},
		endpointDelete:           endpointDelete[ProductionTask]{e},
		endpointAttributes:       endpointAttributes{e},
		endpointStates:           endpointStates{e},
		endpointFiles:            endpointFiles{e},
	}
}
