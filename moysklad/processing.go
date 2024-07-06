package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"time"
)

// Processing Техоперация.
//
// Код сущности: processing
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-tehoperaciq
type Processing struct {
	Organization        *Organization                     `json:"organization,omitempty"`        // Метаданные юрлица
	SyncID              *uuid.UUID                        `json:"syncId,omitempty"`              // ID синхронизации
	Code                *string                           `json:"code,omitempty"`                // Код Техоперации
	Created             *Timestamp                        `json:"created,omitempty"`             // Дата создания
	Deleted             *Timestamp                        `json:"deleted,omitempty"`             // Момент последнего удаления Техоперации
	AccountID           *uuid.UUID                        `json:"accountId,omitempty"`           // ID учётной записи
	ExternalCode        *string                           `json:"externalCode,omitempty"`        // Внешний код Техоперации
	Files               *MetaArray[File]                  `json:"files,omitempty"`               // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group               *Group                            `json:"group,omitempty"`               // Отдел сотрудника
	ID                  *uuid.UUID                        `json:"id,omitempty"`                  // ID Техоперации
	Moment              *Timestamp                        `json:"moment,omitempty"`              // Дата документа
	MaterialsStore      *Store                            `json:"materialsStore,omitempty"`      // Метаданные склада для материалов
	Meta                *Meta                             `json:"meta,omitempty"`                // Метаданные Техоперации
	ProcessingOrder     *ProcessingOrder                  `json:"processingOrder,omitempty"`     // Ссылка на заказ на производство
	Applicable          *bool                             `json:"applicable,omitempty"`          // Отметка о проведении
	Description         *string                           `json:"description,omitempty"`         // Комментарий Техоперации
	OrganizationAccount *AgentAccount                     `json:"organizationAccount,omitempty"` // Метаданные счета юрлица
	Owner               *Employee                         `json:"owner,omitempty"`               // Метаданные владельца (Сотрудника)
	Printed             *bool                             `json:"printed,omitempty"`             // Напечатан ли документ
	ProcessingPlan      *ProcessingPlan                   `json:"processingPlan,omitempty"`      // Метаданные Техкарты
	ProcessingSum       *float64                          `json:"processingSum,omitempty"`       // Затраты на производство за единицу объема производства
	Updated             *Timestamp                        `json:"updated,omitempty"`             // Момент последнего обновления Техоперации
	ProductsStore       *Store                            `json:"productsStore,omitempty"`       // Метаданные склада для продукции
	Project             *NullValue[Project]               `json:"project,omitempty"`             // Метаданные проекта
	Published           *bool                             `json:"published,omitempty"`           // Опубликован ли документ
	Quantity            *float64                          `json:"quantity,omitempty"`            // Объем производства
	Shared              *bool                             `json:"shared,omitempty"`              // Общий доступ
	State               *NullValue[State]                 `json:"state,omitempty"`               // Метаданные статуса Техоперации
	Name                *string                           `json:"name,omitempty"`                // Наименование Техоперации
	Products            Slice[ProcessingPositionProduct]  `json:"products,omitempty"`            // Список Метаданных готовых продуктов Техоперации
	Materials           Slice[ProcessingPositionMaterial] `json:"materials,omitempty"`           // Список Метаданных материалов Техоперации
	Attributes          Slice[Attribute]                  `json:"attributes,omitempty"`          // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (processing Processing) Clean() *Processing {
	if processing.Meta == nil {
		return nil
	}
	return &Processing{Meta: processing.Meta}
}

// AsTaskOperation реализует интерфейс [TaskOperationInterface].
func (processing Processing) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: processing.Meta}
}

// GetOrganization возвращает Метаданные юрлица.
func (processing Processing) GetOrganization() Organization {
	return Deref(processing.Organization)
}

// GetSyncID возвращает ID синхронизации.
func (processing Processing) GetSyncID() uuid.UUID {
	return Deref(processing.SyncID)
}

// GetCode возвращает Код Техоперации.
func (processing Processing) GetCode() string {
	return Deref(processing.Code)
}

// GetCreated возвращает Дату создания.
func (processing Processing) GetCreated() Timestamp {
	return Deref(processing.Created)
}

// GetDeleted возвращает Момент последнего удаления Техоперации.
func (processing Processing) GetDeleted() Timestamp {
	return Deref(processing.Deleted)
}

// GetAccountID возвращает ID учётной записи.
func (processing Processing) GetAccountID() uuid.UUID {
	return Deref(processing.AccountID)
}

// GetExternalCode возвращает Внешний код Техоперации.
func (processing Processing) GetExternalCode() string {
	return Deref(processing.ExternalCode)
}

// GetFiles возвращает Метаданные массива Файлов.
func (processing Processing) GetFiles() MetaArray[File] {
	return Deref(processing.Files)
}

// GetGroup возвращает Отдел сотрудника.
func (processing Processing) GetGroup() Group {
	return Deref(processing.Group)
}

// GetID возвращает ID Техоперации.
func (processing Processing) GetID() uuid.UUID {
	return Deref(processing.ID)
}

// GetMoment возвращает Дату документа.
func (processing Processing) GetMoment() Timestamp {
	return Deref(processing.Moment)
}

// GetMaterialListStore возвращает Метаданные склада для материалов.
func (processing Processing) GetMaterialListStore() Store {
	return Deref(processing.MaterialsStore)
}

// GetMeta возвращает Метаданные Техоперации.
func (processing Processing) GetMeta() Meta {
	return Deref(processing.Meta)
}

// GetProcessingOrder возвращает Ссылку на заказ на производство.
func (processing Processing) GetProcessingOrder() ProcessingOrder {
	return Deref(processing.ProcessingOrder)
}

// GetApplicable возвращает Отметку о проведении.
func (processing Processing) GetApplicable() bool {
	return Deref(processing.Applicable)
}

// GetDescription возвращает Комментарий Техоперации.
func (processing Processing) GetDescription() string {
	return Deref(processing.Description)
}

// GetOrganizationAccount возвращает Метаданные счета юрлица.
func (processing Processing) GetOrganizationAccount() AgentAccount {
	return Deref(processing.OrganizationAccount)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (processing Processing) GetOwner() Employee {
	return Deref(processing.Owner)
}

// GetPrinted возвращает true, если документ напечатан.
func (processing Processing) GetPrinted() bool {
	return Deref(processing.Printed)
}

// GetProcessingPlan возвращает Метаданные Техкарты.
func (processing Processing) GetProcessingPlan() ProcessingPlan {
	return Deref(processing.ProcessingPlan)
}

// GetProcessingSum возвращает Затраты на производство за единицу объема производства.
func (processing Processing) GetProcessingSum() float64 {
	return Deref(processing.ProcessingSum)
}

// GetUpdated возвращает Момент последнего обновления Техоперации.
func (processing Processing) GetUpdated() Timestamp {
	return Deref(processing.Updated)
}

// GetProductListStore возвращает Метаданные склада для продукции.
func (processing Processing) GetProductListStore() Store {
	return Deref(processing.ProductsStore)
}

// GetProject возвращает Метаданные проекта.
func (processing Processing) GetProject() Project {
	return Deref(processing.Project).GetValue()
}

// GetPublished возвращает true, если документ опубликован.
func (processing Processing) GetPublished() bool {
	return Deref(processing.Published)
}

// GetQuantity возвращает Объем производства.
func (processing Processing) GetQuantity() float64 {
	return Deref(processing.Quantity)
}

// GetShared возвращает флаг Общего доступа.
func (processing Processing) GetShared() bool {
	return Deref(processing.Shared)
}

// GetState возвращает Метаданные статуса Техоперации.
func (processing Processing) GetState() State {
	return Deref(processing.State).GetValue()
}

// GetName возвращает Наименование Техоперации.
func (processing Processing) GetName() string {
	return Deref(processing.Name)
}

// GetProductList возвращает Список Метаданных готовых продуктов Техоперации.
func (processing Processing) GetProductList() Slice[ProcessingPositionProduct] {
	return processing.Products
}

// GetMaterialList возвращает Список Метаданных материалов Техоперации.
func (processing Processing) GetMaterialList() Slice[ProcessingPositionMaterial] {
	return processing.Materials
}

// GetAttributes возвращает Список метаданных доп. полей.
func (processing Processing) GetAttributes() Slice[Attribute] {
	return processing.Attributes
}

// SetOrganization устанавливает Метаданные юрлица.
func (processing *Processing) SetOrganization(organization *Organization) *Processing {
	if organization != nil {
		processing.Organization = organization.Clean()
	}
	return processing
}

// SetSyncID устанавливает ID синхронизации.
func (processing *Processing) SetSyncID(syncID uuid.UUID) *Processing {
	processing.SyncID = &syncID
	return processing
}

// SetCode устанавливает Код Техоперации.
func (processing *Processing) SetCode(code string) *Processing {
	processing.Code = &code
	return processing
}

// SetExternalCode устанавливает Внешний код Техоперации.
func (processing *Processing) SetExternalCode(externalCode string) *Processing {
	processing.ExternalCode = &externalCode
	return processing
}

// SetFiles устанавливает Метаданные массива Файлов.
//
// Принимает множество объектов [File].
func (processing *Processing) SetFiles(files ...*File) *Processing {
	processing.Files = NewMetaArrayFrom(files)
	return processing
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (processing *Processing) SetGroup(group *Group) *Processing {
	if group != nil {
		processing.Group = group.Clean()
	}
	return processing
}

// SetMoment устанавливает Дату документа.
func (processing *Processing) SetMoment(moment time.Time) *Processing {
	processing.Moment = NewTimestamp(moment)
	return processing
}

// SetMaterialsStore устанавливает Метаданные склада для материалов.
func (processing *Processing) SetMaterialsStore(materialsStore *Store) *Processing {
	if materialsStore != nil {
		processing.MaterialsStore = materialsStore.Clean()
	}
	return processing
}

// SetMeta устанавливает Метаданные Техоперации.
func (processing *Processing) SetMeta(meta *Meta) *Processing {
	processing.Meta = meta
	return processing
}

// SetProcessingOrder устанавливает Ссылку на заказ на производство.
func (processing *Processing) SetProcessingOrder(processingOrder *ProcessingOrder) *Processing {
	processing.ProcessingOrder = processingOrder.Clean()
	return processing
}

// SetApplicable устанавливает Отметку о проведении.
func (processing *Processing) SetApplicable(applicable bool) *Processing {
	processing.Applicable = &applicable
	return processing
}

// SetDescription устанавливает Комментарий Техоперации.
func (processing *Processing) SetDescription(description string) *Processing {
	processing.Description = &description
	return processing
}

// SetOrganizationAccount устанавливает Метаданные счета юрлица.
func (processing *Processing) SetOrganizationAccount(organizationAccount *AgentAccount) *Processing {
	if organizationAccount != nil {
		processing.OrganizationAccount = organizationAccount.Clean()
	}
	return processing
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (processing *Processing) SetOwner(owner *Employee) *Processing {
	if owner != nil {
		processing.Owner = owner.Clean()
	}
	return processing
}

// SetProcessingPlan устанавливает Метаданные Техкарты.
func (processing *Processing) SetProcessingPlan(processingPlan *ProcessingPlan) *Processing {
	if processingPlan != nil {
		processing.ProcessingPlan = processingPlan.Clean()
	}
	return processing
}

// SetProcessingSum устанавливает Затраты на производство за единицу объема производства.
func (processing *Processing) SetProcessingSum(processingSum float64) *Processing {
	processing.ProcessingSum = &processingSum
	return processing
}

// SetProductsStore устанавливает Метаданные склада для продукции.
func (processing *Processing) SetProductsStore(productsStore *Store) *Processing {
	if productsStore != nil {
		processing.ProductsStore = productsStore.Clean()
	}
	return processing
}

// SetProject устанавливает Метаданные проекта.
//
// Передача nil передаёт сброс значения (null).
func (processing *Processing) SetProject(project *Project) *Processing {
	processing.Project = NewNullValue(project)
	return processing
}

// SetQuantity устанавливает Объем производства.
func (processing *Processing) SetQuantity(quantity float64) *Processing {
	processing.Quantity = &quantity
	return processing
}

// SetShared устанавливает флаг общего доступа.
func (processing *Processing) SetShared(shared bool) *Processing {
	processing.Shared = &shared
	return processing
}

// SetState устанавливает Метаданные статуса Списания.
//
// Передача nil передаёт сброс значения (null).
func (processing *Processing) SetState(state *State) *Processing {
	processing.State = NewNullValue(state)
	return processing
}

// SetName устанавливает Наименование Техоперации.
func (processing *Processing) SetName(name string) *Processing {
	processing.Name = &name
	return processing
}

// SetProducts устанавливает Список Метаданных готовых продуктов Техоперации.
//
// Принимает множество объектов [ProcessingPositionProduct].
func (processing *Processing) SetProducts(products ...*ProcessingPositionProduct) *Processing {
	processing.Products = products
	return processing
}

// SetMaterials устанавливает Список Метаданных материалов Техоперации.
//
// Принимает множество объектов [ProcessingPositionMaterial].
func (processing *Processing) SetMaterials(materials ...*ProcessingPositionMaterial) *Processing {
	processing.Materials = materials
	return processing
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (processing *Processing) SetAttributes(attributes ...*Attribute) *Processing {
	processing.Attributes.Push(attributes...)
	return processing
}

// String реализует интерфейс [fmt.Stringer].
func (processing Processing) String() string {
	return Stringify(processing)
}

// MetaType возвращает код сущности.
func (Processing) MetaType() MetaType {
	return MetaTypeProcessing
}

// Update shortcut
func (processing Processing) Update(ctx context.Context, client *Client, params ...*Params) (*Processing, *resty.Response, error) {
	return NewProcessingService(client).Update(ctx, processing.GetID(), &processing, params...)
}

// Create shortcut
func (processing Processing) Create(ctx context.Context, client *Client, params ...*Params) (*Processing, *resty.Response, error) {
	return NewProcessingService(client).Create(ctx, &processing, params...)
}

// Delete shortcut
func (processing Processing) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewProcessingService(client).Delete(ctx, processing.GetID())
}

// ProcessingPositionMaterial Материал Техоперации.
//
// Код сущности: processingpositionmaterial
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-tehoperaciq-tehoperacii-materialy-tehoperacii
type ProcessingPositionMaterial struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учётной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/серии/модификации, которую представляет собой позиция
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров данного вида в позиции
}

// GetAccountID возвращает ID учётной записи.
func (processingPositionMaterial ProcessingPositionMaterial) GetAccountID() uuid.UUID {
	return Deref(processingPositionMaterial.AccountID)
}

// GetAssortment возвращает Метаданные товара/серии/модификации, которую представляет собой позиция.
func (processingPositionMaterial ProcessingPositionMaterial) GetAssortment() AssortmentPosition {
	return Deref(processingPositionMaterial.Assortment)
}

// GetID возвращает ID позиции.
func (processingPositionMaterial ProcessingPositionMaterial) GetID() uuid.UUID {
	return Deref(processingPositionMaterial.ID)
}

// GetQuantity возвращает Количество товаров данного вида в позиции
func (processingPositionMaterial ProcessingPositionMaterial) GetQuantity() float64 {
	return Deref(processingPositionMaterial.Quantity)
}

// SetAssortment устанавливает Метаданные товара/серии/модификации, которую представляет собой позиция.
//
// Принимает объект, реализующий интерфейс [AssortmentInterface].
func (processingPositionMaterial *ProcessingPositionMaterial) SetAssortment(assortment AssortmentInterface) *ProcessingPositionMaterial {
	if assortment != nil {
		processingPositionMaterial.Assortment = assortment.AsAssortment()
	}
	return processingPositionMaterial
}

// SetQuantity устанавливает Количество товаров данного вида в позиции.
func (processingPositionMaterial *ProcessingPositionMaterial) SetQuantity(quantity float64) *ProcessingPositionMaterial {
	processingPositionMaterial.Quantity = &quantity
	return processingPositionMaterial
}

// String реализует интерфейс [fmt.Stringer].
func (processingPositionMaterial ProcessingPositionMaterial) String() string {
	return Stringify(processingPositionMaterial)
}

// MetaType возвращает код сущности.
func (ProcessingPositionMaterial) MetaType() MetaType {
	return MetaTypeProcessingPositionMaterial
}

// ProcessingPositionProduct Продукт Техоперации.
//
// Код сущности: processingpositionresult
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-tehoperaciq-tehoperacii-produkty-tehoperacii
type ProcessingPositionProduct struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учётной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/серии/модификации, которую представляет собой позиция
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров данного вида в позиции
}

// GetAccountID возвращает ID учётной записи.
func (processingPositionProduct ProcessingPositionProduct) GetAccountID() uuid.UUID {
	return Deref(processingPositionProduct.AccountID)
}

// GetAssortment возвращает Метаданные товара/серии/модификации, которую представляет собой позиция.
func (processingPositionProduct ProcessingPositionProduct) GetAssortment() AssortmentPosition {
	return Deref(processingPositionProduct.Assortment)
}

// GetID возвращает ID позиции.
func (processingPositionProduct ProcessingPositionProduct) GetID() uuid.UUID {
	return Deref(processingPositionProduct.ID)
}

// GetQuantity возвращает Количество товаров данного вида в позиции.
func (processingPositionProduct ProcessingPositionProduct) GetQuantity() float64 {
	return Deref(processingPositionProduct.Quantity)
}

// SetAssortment устанавливает Метаданные товара/серии/модификации, которую представляет собой позиция.
//
// Принимает объект, реализующий интерфейс [AssortmentInterface].
func (processingPositionProduct *ProcessingPositionProduct) SetAssortment(assortment AssortmentInterface) *ProcessingPositionProduct {
	if assortment != nil {
		processingPositionProduct.Assortment = assortment.AsAssortment()
	}
	return processingPositionProduct
}

// SetQuantity устанавливает Количество товаров данного вида в позиции.
func (processingPositionProduct *ProcessingPositionProduct) SetQuantity(quantity float64) *ProcessingPositionProduct {
	processingPositionProduct.Quantity = &quantity
	return processingPositionProduct
}

// String реализует интерфейс [fmt.Stringer].
func (processingPositionProduct ProcessingPositionProduct) String() string {
	return Stringify(processingPositionProduct)
}

// MetaType возвращает код сущности.
func (ProcessingPositionProduct) MetaType() MetaType {
	return MetaTypeProcessingPositionProduct
}

// ProcessingService описывает методы сервиса для работы с техоперациями.
type ProcessingService interface {
	// GetList выполняет запрос на получение списка техопераций.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[Processing], *resty.Response, error)

	// Create выполняет запрос на создание техоперации.
	// Обязательные для создания поля с привязкой техкарты:
	//	- organization (Ссылка на ваше юрлицо)
	//	- productsStore (Ссылка на склад для продукции)
	//	- materialsStore (Ссылка на склад для материалов)
	//	- processingPlan (Ссылка на Техоперацию)
	//
	// Обязательные для создания поля без привязки техкарты:
	//	- organization (Ссылка на ваше юрлицо)
	//	- productsStore (Ссылка на склад для продукции)
	//	- materialsStore (Ссылка на склад для материалов)
	//	- products (Список готовых продуктов Техоперации)
	//	- processingSum (Затраты на производство за единицу объема производства)
	// Если не передается поле quantity, то будет выставлено дефолтное значение равное 1.
	//
	// Принимает контекст, техоперацию и опционально объект параметров запроса Params.
	// Возвращает созданную техоперацию.
	Create(ctx context.Context, processing *Processing, params ...*Params) (*Processing, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение техопераций.
	// Изменяемые техоперации должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список техопераций и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых техопераций.
	CreateUpdateMany(ctx context.Context, processingList Slice[Processing], params ...*Params) (*Slice[Processing], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление техопераций.
	// Принимает контекст и множество техопераций.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*Processing) (*DeleteManyResponse, *resty.Response, error)

	// Delete выполняет запрос на удаление техоперации.
	// Принимает контекст и ID техоперации.
	// Возвращает «true» в случае успешного удаления техоперации.
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельной техоперации по ID.
	// Принимает контекст, ID техоперации и опционально объект параметров запроса Params.
	// Возвращает найденную техоперацию.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Processing, *resty.Response, error)

	// Update выполняет запрос на изменение техоперации.
	// Принимает контекст, техоперацию и опционально объект параметров запроса Params.
	// Возвращает изменённую техоперацию.
	Update(ctx context.Context, id uuid.UUID, processing *Processing, params ...*Params) (*Processing, *resty.Response, error)

	// Template выполняет запрос на получение предзаполненной техоперации со стандартными полями без связи с какими-либо другими документами.
	// Принимает контекст.
	// Возвращает предзаполненную техоперацию.
	Template(ctx context.Context) (*Processing, *resty.Response, error)

	// TemplateBased выполняет запрос на получение шаблона техоперации на основе других документов.
	// Основание, на котором может быть создана:
	//	- Заказ на производство (ProcessingOrder)
	//	- Техкарта (ProcessingPlan)
	// Принимает контекст и множество документов из списка выше.
	// Возвращает предзаполненную техоперацию на основании переданных документов.
	TemplateBased(ctx context.Context, basedOn ...MetaOwner) (*Processing, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных техопераций.
	// Принимает контекст.
	// Возвращает объект метаданных MetaAttributesStatesSharedWrapper.
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

	// GetBySyncID выполняет запрос на получение отдельного документа по syncID.
	// Принимает контекст и syncID документа.
	// Возвращает найденный документ.
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*Processing, *resty.Response, error)

	// DeleteBySyncID выполняет запрос на удаление документа по syncID.
	// Принимает контекст и syncID документа.
	// Возвращает «true» в случае успешного удаления документа.
	DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error)

	// MoveToTrash выполняет запрос на перемещение документа с указанным ID в корзину.
	// Принимает контекст и ID документа.
	// Возвращает «true» в случае успешного перемещения в корзину.
	MoveToTrash(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetMaterialList выполняет запрос на получения списка материалов техоперации.
	// Принимает контекст и ID техоперации.
	// Возвращает объект List.
	GetMaterialList(ctx context.Context, id uuid.UUID) (*List[ProcessingPlanMaterial], *resty.Response, error)

	// CreateMaterial выполняет запрос на создание материала техоперации.
	// Принимает контекст, ID техоперации и материал техоперации.
	// Возвращает созданный материал техоперации.
	CreateMaterial(ctx context.Context, id uuid.UUID, material *ProcessingPlanMaterial) (*ProcessingPlanMaterial, *resty.Response, error)

	// CreateMaterialMany выполняет запрос на массовое создание материалов техоперации.
	// Принимает контекст, ID техоперации и множество материалов техоперации.
	// Возвращает список материалов техоперации.
	CreateMaterialMany(ctx context.Context, id uuid.UUID, materials ...*ProcessingPlanMaterial) (*Slice[ProcessingPlanMaterial], *resty.Response, error)

	// GetMaterialByID выполняет запрос на получение материала техоперации по ID.
	// Принимает контекст, ID техоперации и ID материала техоперации.
	// Возвращает найденный материал техоперации.
	GetMaterialByID(ctx context.Context, id, materialID uuid.UUID) (*ProcessingPlanMaterial, *resty.Response, error)

	// UpdateMaterial выполняет запрос на изменение материала техоперации.
	// Принимает контекст, ID техоперации, ID материала техоперации и материал техоперации.
	// Возвращает изменённый материал техоперации.
	UpdateMaterial(ctx context.Context, id, materialID uuid.UUID, material *ProcessingPlanMaterial) (*ProcessingPlanMaterial, *resty.Response, error)

	// DeleteMaterial выполняет запрос на удаление материала техоперации.
	// Принимает контекст, ID техоперации и ID материала техоперации.
	// Возвращает «true» в случае успешного удаления материала техоперации.
	DeleteMaterial(ctx context.Context, id, materialID uuid.UUID) (bool, *resty.Response, error)

	// GetProductList выполняет запрос на получения списка продуктов техоперации.
	// Принимает контекст и ID техоперации.
	// Возвращает объект List.
	GetProductList(ctx context.Context, id uuid.UUID) (*List[ProcessingPlanProduct], *resty.Response, error)

	// CreateProduct выполняет запрос на создание продукта техоперации.
	// Принимает контекст, ID техоперации и продукт техоперации.
	// Возвращает созданный продукт техоперации.
	CreateProduct(ctx context.Context, id uuid.UUID, product *ProcessingPlanProduct) (*ProcessingPlanProduct, *resty.Response, error)

	// CreateProductMany выполняет запрос на массовое создание продуктов техоперации.
	// Принимает контекст, ID техоперации и множество продуктов техоперации.
	// Возвращает список продуктов техоперации.
	CreateProductMany(ctx context.Context, id uuid.UUID, products ...*ProcessingPlanProduct) (*Slice[ProcessingPlanProduct], *resty.Response, error)

	// GetProductByID выполняет запрос на получение продукта техоперации по ID.
	// Принимает контекст, ID техоперации и ID продукта техоперации.
	// Возвращает найденный продукт техоперации.
	GetProductByID(ctx context.Context, id, productID uuid.UUID) (*ProcessingPlanProduct, *resty.Response, error)

	// UpdateProduct выполняет запрос на изменение продукта техоперации.
	// Принимает контекст, ID техоперации, ID продукта техоперации и продукт техоперации.
	// Возвращает изменённый продукт техоперации.
	UpdateProduct(ctx context.Context, id, productID uuid.UUID, product *ProcessingPlanProduct) (*ProcessingPlanProduct, *resty.Response, error)

	// DeleteProduct выполняет запрос на удаление продукта техоперации.
	// Принимает контекст, ID техоперации и ID продукта техоперации.
	// Возвращает «true» в случае успешного удаления продукта техоперации.
	DeleteProduct(ctx context.Context, id, productID uuid.UUID) (bool, *resty.Response, error)

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

type processingService struct {
	Endpoint
	endpointGetList[Processing]
	endpointCreate[Processing]
	endpointCreateUpdateMany[Processing]
	endpointDeleteMany[Processing]
	endpointDelete
	endpointGetByID[Processing]
	endpointUpdate[Processing]
	endpointTemplate[Processing]
	endpointTemplateBased[Processing]
	endpointMetadata[MetaAttributesStatesSharedWrapper]
	endpointAttributes
	endpointSyncID[Processing]
	endpointTrash
	endpointStates
	endpointFiles
}

func (service *processingService) GetMaterialList(ctx context.Context, id uuid.UUID) (*List[ProcessingPlanMaterial], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/materials", service.uri, id)
	return NewRequestBuilder[List[ProcessingPlanMaterial]](service.client, path).Get(ctx)
}

func (service *processingService) CreateMaterial(ctx context.Context, id uuid.UUID, material *ProcessingPlanMaterial) (*ProcessingPlanMaterial, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/materials", service.uri, id)
	return NewRequestBuilder[ProcessingPlanMaterial](service.client, path).Post(ctx, material)
}

func (service *processingService) CreateMaterialMany(ctx context.Context, id uuid.UUID, materials ...*ProcessingPlanMaterial) (*Slice[ProcessingPlanMaterial], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/materials", service.uri, id)
	return NewRequestBuilder[Slice[ProcessingPlanMaterial]](service.client, path).Post(ctx, materials)
}

func (service *processingService) GetMaterialByID(ctx context.Context, id, materialID uuid.UUID) (*ProcessingPlanMaterial, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/materials/%s", service.uri, id, materialID)
	return NewRequestBuilder[ProcessingPlanMaterial](service.client, path).Get(ctx)
}

func (service *processingService) UpdateMaterial(ctx context.Context, id, materialID uuid.UUID, material *ProcessingPlanMaterial) (*ProcessingPlanMaterial, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/materials/%s", service.uri, id, materialID)
	return NewRequestBuilder[ProcessingPlanMaterial](service.client, path).Put(ctx, material)
}

func (service *processingService) DeleteMaterial(ctx context.Context, id, materialID uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/materials/%s", service.uri, id, materialID)
	return NewRequestBuilder[any](service.client, path).Delete(ctx)
}

func (service *processingService) GetProductList(ctx context.Context, id uuid.UUID) (*List[ProcessingPlanProduct], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/products", service.uri, id)
	return NewRequestBuilder[List[ProcessingPlanProduct]](service.client, path).Get(ctx)
}

func (service *processingService) CreateProduct(ctx context.Context, id uuid.UUID, product *ProcessingPlanProduct) (*ProcessingPlanProduct, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/products", service.uri, id)
	return NewRequestBuilder[ProcessingPlanProduct](service.client, path).Post(ctx, product)
}

func (service *processingService) CreateProductMany(ctx context.Context, id uuid.UUID, products ...*ProcessingPlanProduct) (*Slice[ProcessingPlanProduct], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/products", service.uri, id)
	return NewRequestBuilder[Slice[ProcessingPlanProduct]](service.client, path).Post(ctx, products)
}

func (service *processingService) GetProductByID(ctx context.Context, id, productID uuid.UUID) (*ProcessingPlanProduct, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/products/%s", service.uri, id, productID)
	return NewRequestBuilder[ProcessingPlanProduct](service.client, path).Get(ctx)
}

func (service *processingService) UpdateProduct(ctx context.Context, id, productID uuid.UUID, product *ProcessingPlanProduct) (*ProcessingPlanProduct, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/products/%s", service.uri, id, productID)
	return NewRequestBuilder[ProcessingPlanProduct](service.client, path).Put(ctx, product)
}

func (service *processingService) DeleteProduct(ctx context.Context, id, productID uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/products/%s", service.uri, id, productID)
	return NewRequestBuilder[any](service.client, path).Delete(ctx)
}

// NewProcessingService принимает [Client] и возвращает сервис для работы с техоперациями.
func NewProcessingService(client *Client) ProcessingService {
	e := NewEndpoint(client, "entity/processing")
	return &processingService{
		Endpoint:                 e,
		endpointGetList:          endpointGetList[Processing]{e},
		endpointCreate:           endpointCreate[Processing]{e},
		endpointCreateUpdateMany: endpointCreateUpdateMany[Processing]{e},
		endpointDeleteMany:       endpointDeleteMany[Processing]{e},
		endpointDelete:           endpointDelete{e},
		endpointGetByID:          endpointGetByID[Processing]{e},
		endpointUpdate:           endpointUpdate[Processing]{e},
		endpointMetadata:         endpointMetadata[MetaAttributesStatesSharedWrapper]{e},
		endpointAttributes:       endpointAttributes{e},
		endpointSyncID:           endpointSyncID[Processing]{e},
		endpointTrash:            endpointTrash{e},
		endpointStates:           endpointStates{e},
		endpointFiles:            endpointFiles{e},
		endpointTemplate:         endpointTemplate[Processing]{e},
		endpointTemplateBased:    endpointTemplateBased[Processing]{e},
	}
}
