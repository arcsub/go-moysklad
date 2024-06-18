package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Processing Техоперация.
// Ключевое слово: processing
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-tehoperaciq
type Processing struct {
	Organization        *Organization                     `json:"organization,omitempty"`
	SyncID              *uuid.UUID                        `json:"syncId,omitempty"`
	Code                *string                           `json:"code,omitempty"`
	Created             *Timestamp                        `json:"created,omitempty"`
	Deleted             *Timestamp                        `json:"deleted,omitempty"`
	AccountID           *uuid.UUID                        `json:"accountId,omitempty"`
	ExternalCode        *string                           `json:"externalCode,omitempty"`
	Files               *MetaArray[File]                  `json:"files,omitempty"`
	Group               *Group                            `json:"group,omitempty"`
	ID                  *uuid.UUID                        `json:"id,omitempty"`
	Moment              *Timestamp                        `json:"moment,omitempty"`
	MaterialsStore      *Store                            `json:"materialsStore,omitempty"`
	Meta                *Meta                             `json:"meta,omitempty"`
	ProcessingOrder     *ProcessingOrder                  `json:"processingOrder,omitempty"`
	Applicable          *bool                             `json:"applicable,omitempty"`
	Description         *string                           `json:"description,omitempty"`
	OrganizationAccount *AgentAccount                     `json:"organizationAccount,omitempty"`
	Owner               *Employee                         `json:"owner,omitempty"`
	Printed             *bool                             `json:"printed,omitempty"`
	ProcessingPlan      *ProcessingPlan                   `json:"processingPlan,omitempty"`
	ProcessingSum       *float64                          `json:"processingSum,omitempty"`
	Updated             *Timestamp                        `json:"updated,omitempty"`
	ProductsStore       *Store                            `json:"productsStore,omitempty"`
	Project             *NullValue[Project]               `json:"project,omitempty"`
	Published           *bool                             `json:"published,omitempty"`
	Quantity            *float64                          `json:"quantity,omitempty"`
	Shared              *bool                             `json:"shared,omitempty"`
	State               *State                            `json:"state,omitempty"`
	Name                *string                           `json:"name,omitempty"`
	Products            Slice[ProcessingPositionProduct]  `json:"products,omitempty"`
	Materials           Slice[ProcessingPositionMaterial] `json:"materials,omitempty"`
	Attributes          Slice[Attribute]                  `json:"attributes,omitempty"`
}

// Clean возвращает сущность с единственным заполненным полем Meta
func (processing Processing) Clean() *Processing {
	return &Processing{Meta: processing.Meta}
}

// AsTaskOperation реализует интерфейс AsTaskOperationInterface
func (processing Processing) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: processing.Meta}
}

func (processing Processing) GetOrganization() Organization {
	return Deref(processing.Organization)
}

func (processing Processing) GetSyncID() uuid.UUID {
	return Deref(processing.SyncID)
}

func (processing Processing) GetCode() string {
	return Deref(processing.Code)
}

func (processing Processing) GetCreated() Timestamp {
	return Deref(processing.Created)
}

func (processing Processing) GetDeleted() Timestamp {
	return Deref(processing.Deleted)
}

func (processing Processing) GetAccountID() uuid.UUID {
	return Deref(processing.AccountID)
}

func (processing Processing) GetExternalCode() string {
	return Deref(processing.ExternalCode)
}

func (processing Processing) GetFiles() MetaArray[File] {
	return Deref(processing.Files)
}

func (processing Processing) GetGroup() Group {
	return Deref(processing.Group)
}

func (processing Processing) GetID() uuid.UUID {
	return Deref(processing.ID)
}

func (processing Processing) GetMoment() Timestamp {
	return Deref(processing.Moment)
}

func (processing Processing) GetMaterialsStore() Store {
	return Deref(processing.MaterialsStore)
}

func (processing Processing) GetMeta() Meta {
	return Deref(processing.Meta)
}

func (processing Processing) GetProcessingOrder() ProcessingOrder {
	return Deref(processing.ProcessingOrder)
}

func (processing Processing) GetApplicable() bool {
	return Deref(processing.Applicable)
}

func (processing Processing) GetDescription() string {
	return Deref(processing.Description)
}

func (processing Processing) GetOrganizationAccount() AgentAccount {
	return Deref(processing.OrganizationAccount)
}

func (processing Processing) GetOwner() Employee {
	return Deref(processing.Owner)
}

func (processing Processing) GetPrinted() bool {
	return Deref(processing.Printed)
}

func (processing Processing) GetProcessingPlan() ProcessingPlan {
	return Deref(processing.ProcessingPlan)
}

func (processing Processing) GetProcessingSum() float64 {
	return Deref(processing.ProcessingSum)
}

func (processing Processing) GetUpdated() Timestamp {
	return Deref(processing.Updated)
}

func (processing Processing) GetProductsStore() Store {
	return Deref(processing.ProductsStore)
}

func (processing Processing) GetProject() Project {
	return processing.Project.Get()
}

func (processing Processing) GetPublished() bool {
	return Deref(processing.Published)
}

func (processing Processing) GetQuantity() float64 {
	return Deref(processing.Quantity)
}

func (processing Processing) GetShared() bool {
	return Deref(processing.Shared)
}

func (processing Processing) GetState() State {
	return Deref(processing.State)
}

func (processing Processing) GetName() string {
	return Deref(processing.Name)
}

func (processing Processing) GetProducts() Slice[ProcessingPositionProduct] {
	return processing.Products
}

func (processing Processing) GetMaterials() Slice[ProcessingPositionMaterial] {
	return processing.Materials
}

func (processing Processing) GetAttributes() Slice[Attribute] {
	return processing.Attributes
}

func (processing *Processing) SetOrganization(organization *Organization) *Processing {
	processing.Organization = organization.Clean()
	return processing
}

func (processing *Processing) SetSyncID(syncID uuid.UUID) *Processing {
	processing.SyncID = &syncID
	return processing
}

func (processing *Processing) SetCode(code string) *Processing {
	processing.Code = &code
	return processing
}

func (processing *Processing) SetExternalCode(externalCode string) *Processing {
	processing.ExternalCode = &externalCode
	return processing
}

func (processing *Processing) SetFiles(files ...*File) *Processing {
	processing.Files = NewMetaArrayFrom(files)
	return processing
}

func (processing *Processing) SetGroup(group *Group) *Processing {
	processing.Group = group.Clean()
	return processing
}

func (processing *Processing) SetMoment(moment *Timestamp) *Processing {
	processing.Moment = moment
	return processing
}

func (processing *Processing) SetMaterialsStore(materialsStore *Store) *Processing {
	processing.MaterialsStore = materialsStore.Clean()
	return processing
}

func (processing *Processing) SetMeta(meta *Meta) *Processing {
	processing.Meta = meta
	return processing
}

func (processing *Processing) SetProcessingOrder(processingOrder *ProcessingOrder) *Processing {
	processing.ProcessingOrder = processingOrder.Clean()
	return processing
}

func (processing *Processing) SetApplicable(applicable bool) *Processing {
	processing.Applicable = &applicable
	return processing
}

func (processing *Processing) SetDescription(description string) *Processing {
	processing.Description = &description
	return processing
}

func (processing *Processing) SetOrganizationAccount(organizationAccount *AgentAccount) *Processing {
	processing.OrganizationAccount = organizationAccount.Clean()
	return processing
}

func (processing *Processing) SetOwner(owner *Employee) *Processing {
	processing.Owner = owner.Clean()
	return processing
}

func (processing *Processing) SetProcessingPlan(processingPlan *ProcessingPlan) *Processing {
	processing.ProcessingPlan = processingPlan.Clean()
	return processing
}

func (processing *Processing) SetProcessingSum(processingSum float64) *Processing {
	processing.ProcessingSum = &processingSum
	return processing
}

func (processing *Processing) SetProductsStore(productsStore *Store) *Processing {
	processing.ProductsStore = productsStore.Clean()
	return processing
}

func (processing *Processing) SetProject(project *Project) *Processing {
	processing.Project = NewNullValueFrom(project.Clean())
	return processing
}

func (processing *Processing) SetNullProject() *Processing {
	processing.Project = NewNullValue[Project]()
	return processing
}

func (processing *Processing) SetQuantity(quantity float64) *Processing {
	processing.Quantity = &quantity
	return processing
}

func (processing *Processing) SetShared(shared bool) *Processing {
	processing.Shared = &shared
	return processing
}

func (processing *Processing) SetState(state *State) *Processing {
	processing.State = state.Clean()
	return processing
}

func (processing *Processing) SetName(name string) *Processing {
	processing.Name = &name
	return processing
}

func (processing *Processing) SetProducts(products ...*ProcessingPositionProduct) *Processing {
	processing.Products = products
	return processing
}

func (processing *Processing) SetMaterials(materials ...*ProcessingPositionMaterial) *Processing {
	processing.Materials = materials
	return processing
}

func (processing *Processing) SetAttributes(attributes ...*Attribute) *Processing {
	processing.Attributes = attributes
	return processing
}

func (processing Processing) String() string {
	return Stringify(processing)
}

func (processing Processing) MetaType() MetaType {
	return MetaTypeProcessing
}

// Update shortcut
func (processing Processing) Update(ctx context.Context, client *Client, params ...*Params) (*Processing, *resty.Response, error) {
	return client.Entity().Processing().Update(ctx, processing.GetID(), &processing, params...)
}

// Create shortcut
func (processing Processing) Create(ctx context.Context, client *Client, params ...*Params) (*Processing, *resty.Response, error) {
	return client.Entity().Processing().Create(ctx, &processing, params...)
}

// Delete shortcut
func (processing Processing) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return client.Entity().Processing().Delete(ctx, processing.GetID())
}

// ProcessingPositionMaterial Материал Техоперации.
// Ключевое слово: processingpositionmaterial
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-tehoperaciq-tehoperacii-materialy-tehoperacii
type ProcessingPositionMaterial struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/серии/модификации, которую представляет собой позиция
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров данного вида в позиции
}

func (processingPositionMaterial ProcessingPositionMaterial) GetAccountID() uuid.UUID {
	return Deref(processingPositionMaterial.AccountID)
}

func (processingPositionMaterial ProcessingPositionMaterial) GetAssortment() AssortmentPosition {
	return Deref(processingPositionMaterial.Assortment)
}

func (processingPositionMaterial ProcessingPositionMaterial) GetID() uuid.UUID {
	return Deref(processingPositionMaterial.ID)
}

func (processingPositionMaterial ProcessingPositionMaterial) GetQuantity() float64 {
	return Deref(processingPositionMaterial.Quantity)
}

func (processingPositionMaterial *ProcessingPositionMaterial) SetAssortment(assortment AsAssortment) *ProcessingPositionMaterial {
	processingPositionMaterial.Assortment = assortment.AsAssortment()
	return processingPositionMaterial
}

func (processingPositionMaterial *ProcessingPositionMaterial) SetQuantity(quantity float64) *ProcessingPositionMaterial {
	processingPositionMaterial.Quantity = &quantity
	return processingPositionMaterial
}

func (processingPositionMaterial ProcessingPositionMaterial) String() string {
	return Stringify(processingPositionMaterial)
}

func (processingPositionMaterial ProcessingPositionMaterial) MetaType() MetaType {
	return MetaTypeProcessingPositionMaterial
}

// ProcessingPositionProduct Продукт Техоперации.
// Ключевое слово: processingpositionresult
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-tehoperaciq-tehoperacii-produkty-tehoperacii
type ProcessingPositionProduct struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/серии/модификации, которую представляет собой позиция
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров данного вида в позиции
}

func (processingPositionProduct ProcessingPositionProduct) GetAccountID() uuid.UUID {
	return Deref(processingPositionProduct.AccountID)
}

func (processingPositionProduct ProcessingPositionProduct) GetAssortment() AssortmentPosition {
	return Deref(processingPositionProduct.Assortment)
}

func (processingPositionProduct ProcessingPositionProduct) GetID() uuid.UUID {
	return Deref(processingPositionProduct.ID)
}

func (processingPositionProduct ProcessingPositionProduct) GetQuantity() float64 {
	return Deref(processingPositionProduct.Quantity)
}

func (processingPositionProduct *ProcessingPositionProduct) SetAssortment(assortment AsAssortment) *ProcessingPositionProduct {
	processingPositionProduct.Assortment = assortment.AsAssortment()
	return processingPositionProduct
}

func (processingPositionProduct *ProcessingPositionProduct) SetQuantity(quantity float64) *ProcessingPositionProduct {
	processingPositionProduct.Quantity = &quantity
	return processingPositionProduct
}

func (processingPositionProduct ProcessingPositionProduct) String() string {
	return Stringify(processingPositionProduct)
}

func (processingPositionProduct ProcessingPositionProduct) MetaType() MetaType {
	return MetaTypeProcessingPositionProduct
}

// ProcessingService
// Сервис для работы с Техоперациями.
type ProcessingService interface {
	GetList(ctx context.Context, params ...*Params) (*List[Processing], *resty.Response, error)
	Create(ctx context.Context, processing *Processing, params ...*Params) (*Processing, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, processingList Slice[Processing], params ...*Params) (*Slice[Processing], *resty.Response, error)
	DeleteMany(ctx context.Context, entities ...*Processing) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Processing, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, processing *Processing, params ...*Params) (*Processing, *resty.Response, error)
	Template(ctx context.Context) (*Processing, *resty.Response, error)
	TemplateBased(ctx context.Context, basedOn ...MetaOwner) (*Processing, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetaAttributesSharedStatesWrapper, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributeMany(ctx context.Context, attributes ...*Attribute) (*Slice[Attribute], *resty.Response, error)
	UpdateAttribute(ctx context.Context, id uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributeMany(ctx context.Context, attributes ...*Attribute) (*DeleteManyResponse, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*Processing, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error)
	MoveToTrash(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetMaterials(ctx context.Context, id uuid.UUID) (*List[ProcessingPlanMaterial], *resty.Response, error)
	CreateMaterial(ctx context.Context, id uuid.UUID, material *ProcessingPlanMaterial) (*ProcessingPlanMaterial, *resty.Response, error)
	CreateMaterialMany(ctx context.Context, id uuid.UUID, materials ...*ProcessingPlanMaterial) (*Slice[ProcessingPlanMaterial], *resty.Response, error)
	GetMaterialByID(ctx context.Context, id, materialID uuid.UUID) (*ProcessingPlanMaterial, *resty.Response, error)
	UpdateMaterial(ctx context.Context, id, materialID uuid.UUID, material *ProcessingPlanMaterial) (*ProcessingPlanMaterial, *resty.Response, error)
	DeleteMaterial(ctx context.Context, id, materialID uuid.UUID) (bool, *resty.Response, error)
	GetProducts(ctx context.Context, id uuid.UUID) (*List[ProcessingPlanProduct], *resty.Response, error)
	CreateProduct(ctx context.Context, id uuid.UUID, product *ProcessingPlanProduct) (*ProcessingPlanProduct, *resty.Response, error)
	CreateProductMany(ctx context.Context, id uuid.UUID, products ...*ProcessingPlanProduct) (*Slice[ProcessingPlanProduct], *resty.Response, error)
	GetProductByID(ctx context.Context, id, productID uuid.UUID) (*ProcessingPlanProduct, *resty.Response, error)
	UpdateProduct(ctx context.Context, id, productID uuid.UUID, product *ProcessingPlanProduct) (*ProcessingPlanProduct, *resty.Response, error)
	DeleteProduct(ctx context.Context, id, productID uuid.UUID) (bool, *resty.Response, error)
	GetFiles(ctx context.Context, id uuid.UUID) (*MetaArray[File], *resty.Response, error)
	CreateFile(ctx context.Context, id uuid.UUID, file *File) (*Slice[File], *resty.Response, error)
	UpdateFileMany(ctx context.Context, id uuid.UUID, files ...*File) (*Slice[File], *resty.Response, error)
	DeleteFile(ctx context.Context, id uuid.UUID, fileID uuid.UUID) (bool, *resty.Response, error)
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
	endpointMetadata[MetaAttributesSharedStatesWrapper]
	endpointAttributes
	endpointSyncID[Processing]
	endpointTrash
	endpointStates
	endpointFiles
}

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
		endpointMetadata:         endpointMetadata[MetaAttributesSharedStatesWrapper]{e},
		endpointAttributes:       endpointAttributes{e},
		endpointSyncID:           endpointSyncID[Processing]{e},
		endpointTrash:            endpointTrash{e},
		endpointStates:           endpointStates{e},
		endpointFiles:            endpointFiles{e},
		endpointTemplate:         endpointTemplate[Processing]{e},
		endpointTemplateBased:    endpointTemplateBased[Processing]{e},
	}
}

// GetMaterials Получить материалы Тех. карты.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-poluchit-materialy-teh-karty
func (service *processingService) GetMaterials(ctx context.Context, id uuid.UUID) (*List[ProcessingPlanMaterial], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/materials", service.uri, id)
	return NewRequestBuilder[List[ProcessingPlanMaterial]](service.client, path).Get(ctx)
}

// CreateMaterial Создать материал Тех. карты.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-sozdat-material-teh-karty
func (service *processingService) CreateMaterial(ctx context.Context, id uuid.UUID, material *ProcessingPlanMaterial) (*ProcessingPlanMaterial, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/materials", service.uri, id)
	return NewRequestBuilder[ProcessingPlanMaterial](service.client, path).Post(ctx, material)
}

// CreateMaterialMany Создать несколько материалов Тех. карты.
func (service *processingService) CreateMaterialMany(ctx context.Context, id uuid.UUID, materials ...*ProcessingPlanMaterial) (*Slice[ProcessingPlanMaterial], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/materials", service.uri, id)
	return NewRequestBuilder[Slice[ProcessingPlanMaterial]](service.client, path).Post(ctx, materials)
}

// GetMaterialByID Получить материал.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-poluchit-material
func (service *processingService) GetMaterialByID(ctx context.Context, id, materialID uuid.UUID) (*ProcessingPlanMaterial, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/materials/%s", service.uri, id, materialID)
	return NewRequestBuilder[ProcessingPlanMaterial](service.client, path).Get(ctx)
}

// UpdateMaterial Изменить материал.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-izmenit-material
func (service *processingService) UpdateMaterial(ctx context.Context, id, materialID uuid.UUID, material *ProcessingPlanMaterial) (*ProcessingPlanMaterial, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/materials/%s", service.uri, id, materialID)
	return NewRequestBuilder[ProcessingPlanMaterial](service.client, path).Put(ctx, material)
}

// DeleteMaterial Удалить материал.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-udalit-material
func (service *processingService) DeleteMaterial(ctx context.Context, id, materialID uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/materials/%s", service.uri, id, materialID)
	return NewRequestBuilder[any](service.client, path).Delete(ctx)
}

// GetProducts Получить продукты Тех. карты.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-poluchit-produkty-teh-karty
func (service *processingService) GetProducts(ctx context.Context, id uuid.UUID) (*List[ProcessingPlanProduct], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/products", service.uri, id)
	return NewRequestBuilder[List[ProcessingPlanProduct]](service.client, path).Get(ctx)
}

// CreateProduct Создать продукт Тех. карты.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-sozdat-produkt-teh-karty
func (service *processingService) CreateProduct(ctx context.Context, id uuid.UUID, product *ProcessingPlanProduct) (*ProcessingPlanProduct, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/products", service.uri, id)
	return NewRequestBuilder[ProcessingPlanProduct](service.client, path).Post(ctx, product)
}

// CreateProductMany Создать несколько продуктов Тех. карты.
func (service *processingService) CreateProductMany(ctx context.Context, id uuid.UUID, products ...*ProcessingPlanProduct) (*Slice[ProcessingPlanProduct], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/products", service.uri, id)
	return NewRequestBuilder[Slice[ProcessingPlanProduct]](service.client, path).Post(ctx, products)
}

// GetProductByID Получить продукт.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-poluchit-produkt
func (service *processingService) GetProductByID(ctx context.Context, id, productID uuid.UUID) (*ProcessingPlanProduct, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/products/%s", service.uri, id, productID)
	return NewRequestBuilder[ProcessingPlanProduct](service.client, path).Get(ctx)
}

// UpdateProduct Изменить продукт.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-izmenit-produkt
func (service *processingService) UpdateProduct(ctx context.Context, id, productID uuid.UUID, product *ProcessingPlanProduct) (*ProcessingPlanProduct, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/products/%s", service.uri, id, productID)
	return NewRequestBuilder[ProcessingPlanProduct](service.client, path).Put(ctx, product)
}

// DeleteProduct Удалить продукт.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-karta-udalit-produkt
func (service *processingService) DeleteProduct(ctx context.Context, id, productID uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/products/%s", service.uri, id, productID)
	return NewRequestBuilder[any](service.client, path).Delete(ctx)
}
