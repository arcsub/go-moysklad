package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// ProcessingOrder Заказ на производство.
// Ключевое слово: processingorder
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-zakaz-na-proizwodstwo
type ProcessingOrder struct {
	Name                  *string                             `json:"name,omitempty"`
	Published             *bool                               `json:"published,omitempty"`
	Organization          *Organization                       `json:"organization,omitempty"`
	Code                  *string                             `json:"code,omitempty"`
	Created               *Timestamp                          `json:"created,omitempty"`
	Deleted               *Timestamp                          `json:"deleted,omitempty"`
	DeliveryPlannedMoment *Timestamp                          `json:"deliveryPlannedMoment,omitempty"`
	Description           *string                             `json:"description,omitempty"`
	ExternalCode          *string                             `json:"externalCode,omitempty"`
	Files                 *MetaArray[File]                    `json:"files,omitempty"`
	Group                 *Group                              `json:"group,omitempty"`
	ID                    *uuid.UUID                          `json:"id,omitempty"`
	Meta                  *Meta                               `json:"meta,omitempty"`
	Moment                *Timestamp                          `json:"moment,omitempty"`
	AccountID             *uuid.UUID                          `json:"accountId,omitempty"`
	OrganizationAccount   *AgentAccount                       `json:"organizationAccount,omitempty"`
	Owner                 *Employee                           `json:"owner,omitempty"`
	Positions             *Positions[ProcessingOrderPosition] `json:"positions,omitempty"`
	Printed               *bool                               `json:"printed,omitempty"`
	ProcessingPlan        *ProcessingPlan                     `json:"processingPlan,omitempty"`
	Project               *Project                            `json:"project,omitempty"`
	Applicable            *bool                               `json:"applicable,omitempty"`
	Quantity              *float64                            `json:"quantity,omitempty"`
	Shared                *bool                               `json:"shared,omitempty"`
	State                 *State                              `json:"state,omitempty"`
	Store                 *Store                              `json:"store,omitempty"`
	SyncID                *uuid.UUID                          `json:"syncId,omitempty"`
	Updated               *Timestamp                          `json:"updated,omitempty"`
	Processings           Slice[Processing]                   `json:"processings,omitempty"`
	Attributes            Slice[AttributeValue]               `json:"attributes,omitempty"`
}

func (processingOrder ProcessingOrder) Clean() *ProcessingOrder {
	return &ProcessingOrder{Meta: processingOrder.Meta}
}

func (processingOrder ProcessingOrder) GetName() string {
	return Deref(processingOrder.Name)
}

func (processingOrder ProcessingOrder) GetPublished() bool {
	return Deref(processingOrder.Published)
}

func (processingOrder ProcessingOrder) GetOrganization() Organization {
	return Deref(processingOrder.Organization)
}

func (processingOrder ProcessingOrder) GetCode() string {
	return Deref(processingOrder.Code)
}

func (processingOrder ProcessingOrder) GetCreated() Timestamp {
	return Deref(processingOrder.Created)
}

func (processingOrder ProcessingOrder) GetDeleted() Timestamp {
	return Deref(processingOrder.Deleted)
}

func (processingOrder ProcessingOrder) GetDeliveryPlannedMoment() Timestamp {
	return Deref(processingOrder.DeliveryPlannedMoment)
}

func (processingOrder ProcessingOrder) GetDescription() string {
	return Deref(processingOrder.Description)
}

func (processingOrder ProcessingOrder) GetExternalCode() string {
	return Deref(processingOrder.ExternalCode)
}

func (processingOrder ProcessingOrder) GetFiles() MetaArray[File] {
	return Deref(processingOrder.Files)
}

func (processingOrder ProcessingOrder) GetGroup() Group {
	return Deref(processingOrder.Group)
}

func (processingOrder ProcessingOrder) GetID() uuid.UUID {
	return Deref(processingOrder.ID)
}

func (processingOrder ProcessingOrder) GetMeta() Meta {
	return Deref(processingOrder.Meta)
}

func (processingOrder ProcessingOrder) GetMoment() Timestamp {
	return Deref(processingOrder.Moment)
}

func (processingOrder ProcessingOrder) GetAccountID() uuid.UUID {
	return Deref(processingOrder.AccountID)
}

func (processingOrder ProcessingOrder) GetOrganizationAccount() AgentAccount {
	return Deref(processingOrder.OrganizationAccount)
}

func (processingOrder ProcessingOrder) GetOwner() Employee {
	return Deref(processingOrder.Owner)
}

func (processingOrder ProcessingOrder) GetPositions() Positions[ProcessingOrderPosition] {
	return Deref(processingOrder.Positions)
}

func (processingOrder ProcessingOrder) GetPrinted() bool {
	return Deref(processingOrder.Printed)
}

func (processingOrder ProcessingOrder) GetProcessingPlan() ProcessingPlan {
	return Deref(processingOrder.ProcessingPlan)
}

func (processingOrder ProcessingOrder) GetProject() Project {
	return Deref(processingOrder.Project)
}

func (processingOrder ProcessingOrder) GetApplicable() bool {
	return Deref(processingOrder.Applicable)
}

func (processingOrder ProcessingOrder) GetQuantity() float64 {
	return Deref(processingOrder.Quantity)
}

func (processingOrder ProcessingOrder) GetShared() bool {
	return Deref(processingOrder.Shared)
}

func (processingOrder ProcessingOrder) GetState() State {
	return Deref(processingOrder.State)
}

func (processingOrder ProcessingOrder) GetStore() Store {
	return Deref(processingOrder.Store)
}

func (processingOrder ProcessingOrder) GetSyncID() uuid.UUID {
	return Deref(processingOrder.SyncID)
}

func (processingOrder ProcessingOrder) GetUpdated() Timestamp {
	return Deref(processingOrder.Updated)
}

func (processingOrder ProcessingOrder) GetProcessings() Slice[Processing] {
	return processingOrder.Processings
}

func (processingOrder ProcessingOrder) GetAttributes() Slice[AttributeValue] {
	return processingOrder.Attributes
}

func (processingOrder *ProcessingOrder) SetName(name string) *ProcessingOrder {
	processingOrder.Name = &name
	return processingOrder
}

func (processingOrder *ProcessingOrder) SetOrganization(organization *Organization) *ProcessingOrder {
	processingOrder.Organization = organization.Clean()
	return processingOrder
}

func (processingOrder *ProcessingOrder) SetCode(code string) *ProcessingOrder {
	processingOrder.Code = &code
	return processingOrder
}

func (processingOrder *ProcessingOrder) SetDeliveryPlannedMoment(deliveryPlannedMoment *Timestamp) *ProcessingOrder {
	processingOrder.DeliveryPlannedMoment = deliveryPlannedMoment
	return processingOrder
}

func (processingOrder *ProcessingOrder) SetDescription(description string) *ProcessingOrder {
	processingOrder.Description = &description
	return processingOrder
}

func (processingOrder *ProcessingOrder) SetExternalCode(externalCode string) *ProcessingOrder {
	processingOrder.ExternalCode = &externalCode
	return processingOrder
}

func (processingOrder *ProcessingOrder) SetFiles(files Slice[File]) *ProcessingOrder {
	processingOrder.Files = NewMetaArrayRows(files)
	return processingOrder
}

func (processingOrder *ProcessingOrder) SetGroup(group *Group) *ProcessingOrder {
	processingOrder.Group = group.Clean()
	return processingOrder
}

func (processingOrder *ProcessingOrder) SetMeta(meta *Meta) *ProcessingOrder {
	processingOrder.Meta = meta
	return processingOrder
}

func (processingOrder *ProcessingOrder) SetMoment(moment *Timestamp) *ProcessingOrder {
	processingOrder.Moment = moment
	return processingOrder
}

func (processingOrder *ProcessingOrder) SetOrganizationAccount(organizationAccount *AgentAccount) *ProcessingOrder {
	processingOrder.OrganizationAccount = organizationAccount.Clean()
	return processingOrder
}

func (processingOrder *ProcessingOrder) SetOwner(owner *Employee) *ProcessingOrder {
	processingOrder.Owner = owner.Clean()
	return processingOrder
}

func (processingOrder *ProcessingOrder) SetPositions(positions *Positions[ProcessingOrderPosition]) *ProcessingOrder {
	processingOrder.Positions = positions
	return processingOrder
}

func (processingOrder *ProcessingOrder) SetProcessingPlan(processingPlan *ProcessingPlan) *ProcessingOrder {
	processingOrder.ProcessingPlan = processingPlan.Clean()
	return processingOrder
}

func (processingOrder *ProcessingOrder) SetProject(project *Project) *ProcessingOrder {
	processingOrder.Project = project.Clean()
	return processingOrder
}

func (processingOrder *ProcessingOrder) SetApplicable(applicable bool) *ProcessingOrder {
	processingOrder.Applicable = &applicable
	return processingOrder
}

func (processingOrder *ProcessingOrder) SetQuantity(quantity float64) *ProcessingOrder {
	processingOrder.Quantity = &quantity
	return processingOrder
}

func (processingOrder *ProcessingOrder) SetShared(shared bool) *ProcessingOrder {
	processingOrder.Shared = &shared
	return processingOrder
}

func (processingOrder *ProcessingOrder) SetState(state *State) *ProcessingOrder {
	processingOrder.State = state.Clean()
	return processingOrder
}

func (processingOrder *ProcessingOrder) SetStore(store *Store) *ProcessingOrder {
	processingOrder.Store = store.Clean()
	return processingOrder
}

func (processingOrder *ProcessingOrder) SetSyncID(syncID uuid.UUID) *ProcessingOrder {
	processingOrder.SyncID = &syncID
	return processingOrder
}

func (processingOrder *ProcessingOrder) SetProcessings(processings Slice[Processing]) *ProcessingOrder {
	processingOrder.Processings = processings
	return processingOrder
}

func (processingOrder *ProcessingOrder) SetAttributes(attributes Slice[AttributeValue]) *ProcessingOrder {
	processingOrder.Attributes = attributes
	return processingOrder
}

func (processingOrder ProcessingOrder) String() string {
	return Stringify(processingOrder)
}

func (processingOrder ProcessingOrder) MetaType() MetaType {
	return MetaTypeProcessingOrder
}

// ProcessingOrderPosition Позиция Заказа на производство.
// Ключевое слово: processingorderposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-zakaz-na-proizwodstwo-zakazy-na-proizwodstwo-pozicii-zakaza-na-proizwodstwo
type ProcessingOrderPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Reserve    *float64            `json:"reserve,omitempty"`    // Резерв данной позиции
}

func (processingOrderPosition ProcessingOrderPosition) GetAccountID() uuid.UUID {
	return Deref(processingOrderPosition.AccountID)
}

func (processingOrderPosition ProcessingOrderPosition) GetAssortment() AssortmentPosition {
	return Deref(processingOrderPosition.Assortment)
}

func (processingOrderPosition ProcessingOrderPosition) GetID() uuid.UUID {
	return Deref(processingOrderPosition.ID)
}

func (processingOrderPosition ProcessingOrderPosition) GetPack() Pack {
	return Deref(processingOrderPosition.Pack)
}

func (processingOrderPosition ProcessingOrderPosition) GetQuantity() float64 {
	return Deref(processingOrderPosition.Quantity)
}

func (processingOrderPosition ProcessingOrderPosition) GetReserve() float64 {
	return Deref(processingOrderPosition.Reserve)
}

func (processingOrderPosition *ProcessingOrderPosition) SetAssortment(assortment AsAssortment) *ProcessingOrderPosition {
	processingOrderPosition.Assortment = assortment.AsAssortment()
	return processingOrderPosition
}

func (processingOrderPosition *ProcessingOrderPosition) SetPack(pack *Pack) *ProcessingOrderPosition {
	processingOrderPosition.Pack = pack
	return processingOrderPosition
}

func (processingOrderPosition *ProcessingOrderPosition) SetQuantity(quantity float64) *ProcessingOrderPosition {
	processingOrderPosition.Quantity = &quantity
	return processingOrderPosition
}

func (processingOrderPosition *ProcessingOrderPosition) SetReserve(reserve float64) *ProcessingOrderPosition {
	processingOrderPosition.Reserve = &reserve
	return processingOrderPosition
}

func (processingOrderPosition ProcessingOrderPosition) String() string {
	return Stringify(processingOrderPosition)
}

func (processingOrderPosition ProcessingOrderPosition) MetaType() MetaType {
	return MetaTypeProcessingOrderPosition
}

// ProcessingOrderTemplateArg
// Документ: Заказ на производство (processingorder)
// Основание, на котором он может быть создан:
// - Техкарта (processingplan)
//type ProcessingOrderTemplateArg struct {
//	ProcessingPlan *MetaWrapper `json:"processingPlan,omitempty"`
//}

// ProcessingOrderService
// Сервис для работы с заказами на производство.
type ProcessingOrderService interface {
	GetList(ctx context.Context, params *Params) (*List[ProcessingOrder], *resty.Response, error)
	Create(ctx context.Context, processingOrder *ProcessingOrder, params *Params) (*ProcessingOrder, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, processingOrderList Slice[ProcessingOrder], params *Params) (*Slice[ProcessingOrder], *resty.Response, error)
	DeleteMany(ctx context.Context, processingOrderList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params *Params) (*ProcessingOrder, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, processingOrder *ProcessingOrder, params *Params) (*ProcessingOrder, *resty.Response, error)
	//endpointTemplate[ProcessingOrder]
	//endpointTemplateBasedOn[ProcessingOrder, ProcessingOrderTemplateArg]
	GetMetadata(ctx context.Context) (*MetaAttributesSharedStatesWrapper, *resty.Response, error)
	GetPositions(ctx context.Context, id uuid.UUID, params *Params) (*MetaArray[ProcessingOrderPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params *Params) (*ProcessingOrderPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *ProcessingOrderPosition, params *Params) (*ProcessingOrderPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id uuid.UUID, position *ProcessingOrderPosition) (*ProcessingOrderPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id uuid.UUID, positions Slice[ProcessingOrderPosition]) (*Slice[ProcessingOrderPosition], *resty.Response, error)
	DeletePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (bool, *resty.Response, error)
	GetPositionTrackingCodes(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (*MetaArray[TrackingCode], *resty.Response, error)
	CreateOrUpdatePositionTrackingCodes(ctx context.Context, id uuid.UUID, positionID uuid.UUID, trackingCodes Slice[TrackingCode]) (*Slice[TrackingCode], *resty.Response, error)
	DeletePositionTrackingCodes(ctx context.Context, id uuid.UUID, positionID uuid.UUID, trackingCodes Slice[TrackingCode]) (*DeleteManyResponse, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList Slice[Attribute]) (*Slice[Attribute], *resty.Response, error)
	UpdateAttribute(ctx context.Context, id uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*ProcessingOrder, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error)
	MoveToTrash(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetStateByID(ctx context.Context, id uuid.UUID) (*State, *resty.Response, error)
	CreateState(ctx context.Context, state *State) (*State, *resty.Response, error)
	UpdateState(ctx context.Context, id uuid.UUID, state *State) (*State, *resty.Response, error)
	CreateOrUpdateStates(ctx context.Context, states Slice[State]) (*Slice[State], *resty.Response, error)
	DeleteState(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
}

func NewProcessingOrderService(client *Client) ProcessingOrderService {
	e := NewEndpoint(client, "entity/processingorder")
	return newMainService[ProcessingOrder, ProcessingOrderPosition, MetaAttributesSharedStatesWrapper, any](e)
}
