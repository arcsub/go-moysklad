package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// InternalOrder Внутренний заказ.
// Ключевое слово: internalorder
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vnutrennij-zakaz
type InternalOrder struct {
	Organization          *Organization                     `json:"organization,omitempty"`
	Description           *string                           `json:"description,omitempty"`
	VatSum                *float64                          `json:"vatSum,omitempty"`
	AccountID             *uuid.UUID                        `json:"accountId,omitempty"`
	Created               *Timestamp                        `json:"created,omitempty"`
	Deleted               *Timestamp                        `json:"deleted,omitempty"`
	DeliveryPlannedMoment *Timestamp                        `json:"deliveryPlannedMoment,omitempty"`
	Owner                 *Employee                         `json:"owner,omitempty"`
	ExternalCode          *string                           `json:"externalCode,omitempty"`
	Files                 *MetaArray[File]                  `json:"files,omitempty"`
	Group                 *Group                            `json:"group,omitempty"`
	ID                    *uuid.UUID                        `json:"id,omitempty"`
	Meta                  *Meta                             `json:"meta,omitempty"`
	Positions             *Positions[InternalOrderPosition] `json:"positions,omitempty"`
	Moves                 Slice[Move]                       `json:"moves,omitempty"`
	Name                  *string                           `json:"name,omitempty"`
	Code                  *string                           `json:"code,omitempty"`
	Applicable            *bool                             `json:"applicable,omitempty"`
	Moment                *Timestamp                        `json:"moment,omitempty"`
	Printed               *bool                             `json:"printed,omitempty"`
	Project               *NullValue[Project]               `json:"project,omitempty"`
	Published             *bool                             `json:"published,omitempty"`
	PurchaseOrders        Slice[PurchaseOrder]              `json:"purchaseOrders,omitempty"`
	Rate                  *NullValue[Rate]                  `json:"rate,omitempty"`
	Shared                *bool                             `json:"shared,omitempty"`
	State                 *NullValue[State]                 `json:"state,omitempty"`
	Store                 *NullValue[Store]                 `json:"store,omitempty"`
	Sum                   *float64                          `json:"sum,omitempty"`
	SyncID                *uuid.UUID                        `json:"syncId,omitempty"`
	Updated               *Timestamp                        `json:"updated,omitempty"`
	VatEnabled            *bool                             `json:"vatEnabled,omitempty"`
	VatIncluded           *bool                             `json:"vatIncluded,omitempty"`
	Attributes            Slice[Attribute]                  `json:"attributes,omitempty"`
}

func (internalOrder InternalOrder) Clean() *InternalOrder {
	return &InternalOrder{Meta: internalOrder.Meta}
}

// AsTaskOperation реализует интерфейс AsTaskOperationInterface
func (internalOrder InternalOrder) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: internalOrder.Meta}
}

func (internalOrder InternalOrder) GetOrganization() Organization {
	return Deref(internalOrder.Organization)
}

func (internalOrder InternalOrder) GetDescription() string {
	return Deref(internalOrder.Description)
}

func (internalOrder InternalOrder) GetVatSum() float64 {
	return Deref(internalOrder.VatSum)
}

func (internalOrder InternalOrder) GetAccountID() uuid.UUID {
	return Deref(internalOrder.AccountID)
}

func (internalOrder InternalOrder) GetCreated() Timestamp {
	return Deref(internalOrder.Created)
}

func (internalOrder InternalOrder) GetDeleted() Timestamp {
	return Deref(internalOrder.Deleted)
}

func (internalOrder InternalOrder) GetDeliveryPlannedMoment() Timestamp {
	return Deref(internalOrder.DeliveryPlannedMoment)
}

func (internalOrder InternalOrder) GetOwner() Employee {
	return Deref(internalOrder.Owner)
}

func (internalOrder InternalOrder) GetExternalCode() string {
	return Deref(internalOrder.ExternalCode)
}

func (internalOrder InternalOrder) GetFiles() MetaArray[File] {
	return Deref(internalOrder.Files)
}

func (internalOrder InternalOrder) GetGroup() Group {
	return Deref(internalOrder.Group)
}

func (internalOrder InternalOrder) GetID() uuid.UUID {
	return Deref(internalOrder.ID)
}

func (internalOrder InternalOrder) GetMeta() Meta {
	return Deref(internalOrder.Meta)
}

func (internalOrder InternalOrder) GetPositions() Positions[InternalOrderPosition] {
	return Deref(internalOrder.Positions)
}

func (internalOrder InternalOrder) GetMoves() Slice[Move] {
	return internalOrder.Moves
}

func (internalOrder InternalOrder) GetName() string {
	return Deref(internalOrder.Name)
}

func (internalOrder InternalOrder) GetCode() string {
	return Deref(internalOrder.Code)
}

func (internalOrder InternalOrder) GetApplicable() bool {
	return Deref(internalOrder.Applicable)
}

func (internalOrder InternalOrder) GetMoment() Timestamp {
	return Deref(internalOrder.Moment)
}

func (internalOrder InternalOrder) GetPrinted() bool {
	return Deref(internalOrder.Printed)
}

func (internalOrder InternalOrder) GetProject() Project {
	return internalOrder.Project.Get()
}

func (internalOrder InternalOrder) GetPublished() bool {
	return Deref(internalOrder.Published)
}

func (internalOrder InternalOrder) GetPurchaseOrders() Slice[PurchaseOrder] {
	return internalOrder.PurchaseOrders
}

func (internalOrder InternalOrder) GetRate() Rate {
	return internalOrder.Rate.Get()
}

func (internalOrder InternalOrder) GetShared() bool {
	return Deref(internalOrder.Shared)
}

func (internalOrder InternalOrder) GetState() State {
	return internalOrder.State.Get()
}

func (internalOrder InternalOrder) GetStore() Store {
	return internalOrder.Store.Get()
}

func (internalOrder InternalOrder) GetSum() float64 {
	return Deref(internalOrder.Sum)
}

func (internalOrder InternalOrder) GetSyncID() uuid.UUID {
	return Deref(internalOrder.SyncID)
}

func (internalOrder InternalOrder) GetUpdated() Timestamp {
	return Deref(internalOrder.Updated)
}

func (internalOrder InternalOrder) GetVatEnabled() bool {
	return Deref(internalOrder.VatEnabled)
}

func (internalOrder InternalOrder) GetVatIncluded() bool {
	return Deref(internalOrder.VatIncluded)
}

func (internalOrder InternalOrder) GetAttributes() Slice[Attribute] {
	return internalOrder.Attributes
}

func (internalOrder *InternalOrder) SetOrganization(organization *Organization) *InternalOrder {
	internalOrder.Organization = organization.Clean()
	return internalOrder
}

func (internalOrder *InternalOrder) SetDescription(description string) *InternalOrder {
	internalOrder.Description = &description
	return internalOrder
}

func (internalOrder *InternalOrder) SetDeliveryPlannedMoment(deliveryPlannedMoment *Timestamp) *InternalOrder {
	internalOrder.DeliveryPlannedMoment = deliveryPlannedMoment
	return internalOrder
}

func (internalOrder *InternalOrder) SetOwner(owner *Employee) *InternalOrder {
	internalOrder.Owner = owner.Clean()
	return internalOrder
}

func (internalOrder *InternalOrder) SetExternalCode(externalCode string) *InternalOrder {
	internalOrder.ExternalCode = &externalCode
	return internalOrder
}

func (internalOrder *InternalOrder) SetFiles(files ...*File) *InternalOrder {
	internalOrder.Files = NewMetaArrayFrom(files)
	return internalOrder
}

func (internalOrder *InternalOrder) SetGroup(group *Group) *InternalOrder {
	internalOrder.Group = group.Clean()
	return internalOrder
}

func (internalOrder *InternalOrder) SetMeta(meta *Meta) *InternalOrder {
	internalOrder.Meta = meta
	return internalOrder
}

func (internalOrder *InternalOrder) SetPositions(positions ...*InternalOrderPosition) *InternalOrder {
	internalOrder.Positions = NewPositionsFrom(positions)
	return internalOrder
}

func (internalOrder *InternalOrder) SetMoves(moves ...*Move) *InternalOrder {
	internalOrder.Moves = moves
	return internalOrder
}

func (internalOrder *InternalOrder) SetName(name string) *InternalOrder {
	internalOrder.Name = &name
	return internalOrder
}

func (internalOrder *InternalOrder) SetCode(code string) *InternalOrder {
	internalOrder.Code = &code
	return internalOrder
}

func (internalOrder *InternalOrder) SetApplicable(applicable bool) *InternalOrder {
	internalOrder.Applicable = &applicable
	return internalOrder
}

func (internalOrder *InternalOrder) SetMoment(moment *Timestamp) *InternalOrder {
	internalOrder.Moment = moment
	return internalOrder
}

func (internalOrder *InternalOrder) SetProject(project *Project) *InternalOrder {
	internalOrder.Project = NewNullValueFrom(project.Clean())
	return internalOrder
}

func (internalOrder *InternalOrder) SetNullProject() *InternalOrder {
	internalOrder.Project = NewNullValue[Project]()
	return internalOrder
}

func (internalOrder *InternalOrder) SetPurchaseOrders(purchaseOrders ...*PurchaseOrder) *InternalOrder {
	internalOrder.PurchaseOrders = purchaseOrders
	return internalOrder
}

func (internalOrder *InternalOrder) SetRate(rate *Rate) *InternalOrder {
	internalOrder.Rate = NewNullValueFrom(rate)
	return internalOrder
}

func (internalOrder *InternalOrder) SetNullRate() *InternalOrder {
	internalOrder.Rate = NewNullValue[Rate]()
	return internalOrder
}

func (internalOrder *InternalOrder) SetShared(shared bool) *InternalOrder {
	internalOrder.Shared = &shared
	return internalOrder
}

func (internalOrder *InternalOrder) SetState(state *State) *InternalOrder {
	internalOrder.State = NewNullValueFrom(state.Clean())
	return internalOrder
}

func (internalOrder *InternalOrder) SetNullState() *InternalOrder {
	internalOrder.State = NewNullValue[State]()
	return internalOrder
}

func (internalOrder *InternalOrder) SetStore(store *Store) *InternalOrder {
	internalOrder.Store = NewNullValueFrom(store.Clean())
	return internalOrder
}

func (internalOrder *InternalOrder) SetNullStore() *InternalOrder {
	internalOrder.Store = NewNullValue[Store]()
	return internalOrder
}

func (internalOrder *InternalOrder) SetSyncID(syncID uuid.UUID) *InternalOrder {
	internalOrder.SyncID = &syncID
	return internalOrder
}

func (internalOrder *InternalOrder) SetVatEnabled(vatEnabled bool) *InternalOrder {
	internalOrder.VatEnabled = &vatEnabled
	return internalOrder
}

func (internalOrder *InternalOrder) SetVatIncluded(vatIncluded bool) *InternalOrder {
	internalOrder.VatIncluded = &vatIncluded
	return internalOrder
}

func (internalOrder *InternalOrder) SetAttributes(attributes ...*Attribute) *InternalOrder {
	internalOrder.Attributes = attributes
	return internalOrder
}

func (internalOrder InternalOrder) String() string {
	return Stringify(internalOrder)
}

// MetaType возвращает тип сущности.
func (InternalOrder) MetaType() MetaType {
	return MetaTypeInternalOrder
}

// Update shortcut
func (internalOrder InternalOrder) Update(ctx context.Context, client *Client, params ...*Params) (*InternalOrder, *resty.Response, error) {
	return client.Entity().InternalOrder().Update(ctx, internalOrder.GetID(), &internalOrder, params...)
}

// Create shortcut
func (internalOrder InternalOrder) Create(ctx context.Context, client *Client, params ...*Params) (*InternalOrder, *resty.Response, error) {
	return client.Entity().InternalOrder().Create(ctx, &internalOrder, params...)
}

// Delete shortcut
func (internalOrder InternalOrder) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return client.Entity().InternalOrder().Delete(ctx, internalOrder.GetID())
}

// InternalOrderPosition Позиция Внутреннего заказа.
// Ключевое слово: internalorderposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vnutrennij-zakaz-vnutrennie-zakazy-pozicii-vnutrennego-zakaza
type InternalOrderPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Price      *float64            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Vat        *int                `json:"vat,omitempty"`        // НДС, которым облагается текущая позиция
	VatEnabled *bool               `json:"vatEnabled,omitempty"` // Включен ли НДС для позиции. С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
}

func (internalOrderPosition InternalOrderPosition) GetAccountID() uuid.UUID {
	return Deref(internalOrderPosition.AccountID)
}

func (internalOrderPosition InternalOrderPosition) GetAssortment() AssortmentPosition {
	return Deref(internalOrderPosition.Assortment)
}

func (internalOrderPosition InternalOrderPosition) GetID() uuid.UUID {
	return Deref(internalOrderPosition.ID)
}

func (internalOrderPosition InternalOrderPosition) GetPack() Pack {
	return Deref(internalOrderPosition.Pack)
}

func (internalOrderPosition InternalOrderPosition) GetPrice() float64 {
	return Deref(internalOrderPosition.Price)
}

func (internalOrderPosition InternalOrderPosition) GetQuantity() float64 {
	return Deref(internalOrderPosition.Quantity)
}

func (internalOrderPosition InternalOrderPosition) GetVat() int {
	return Deref(internalOrderPosition.Vat)
}

func (internalOrderPosition InternalOrderPosition) GetVatEnabled() bool {
	return Deref(internalOrderPosition.VatEnabled)
}

func (internalOrderPosition *InternalOrderPosition) SetAssortment(assortment AsAssortment) *InternalOrderPosition {
	internalOrderPosition.Assortment = assortment.AsAssortment()
	return internalOrderPosition
}

func (internalOrderPosition *InternalOrderPosition) SetPack(pack *Pack) *InternalOrderPosition {
	internalOrderPosition.Pack = pack
	return internalOrderPosition
}

func (internalOrderPosition *InternalOrderPosition) SetPrice(price float64) *InternalOrderPosition {
	internalOrderPosition.Price = &price
	return internalOrderPosition
}

func (internalOrderPosition *InternalOrderPosition) SetQuantity(quantity float64) *InternalOrderPosition {
	internalOrderPosition.Quantity = &quantity
	return internalOrderPosition
}

func (internalOrderPosition *InternalOrderPosition) SetVat(vat int) *InternalOrderPosition {
	internalOrderPosition.Vat = &vat
	return internalOrderPosition
}

func (internalOrderPosition *InternalOrderPosition) SetVatEnabled(vatEnabled bool) *InternalOrderPosition {
	internalOrderPosition.VatEnabled = &vatEnabled
	return internalOrderPosition
}

func (internalOrderPosition InternalOrderPosition) String() string {
	return Stringify(internalOrderPosition)
}

// MetaType возвращает тип сущности.
func (InternalOrderPosition) MetaType() MetaType {
	return MetaTypeInternalOrderPosition
}

// InternalOrderService
// Сервис для работы с внутренними заказами.
type InternalOrderService interface {
	GetList(ctx context.Context, params ...*Params) (*List[InternalOrder], *resty.Response, error)
	Create(ctx context.Context, internalOrder *InternalOrder, params ...*Params) (*InternalOrder, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, internalOrderList Slice[InternalOrder], params ...*Params) (*Slice[InternalOrder], *resty.Response, error)
	DeleteMany(ctx context.Context, entities ...*InternalOrder) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*InternalOrder, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, internalOrder *InternalOrder, params ...*Params) (*InternalOrder, *resty.Response, error)
	Template(ctx context.Context) (*InternalOrder, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetaAttributesSharedStatesWrapper, *resty.Response, error)
	GetPositions(ctx context.Context, id uuid.UUID, params ...*Params) (*MetaArray[InternalOrderPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params ...*Params) (*InternalOrderPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *InternalOrderPosition, params ...*Params) (*InternalOrderPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id uuid.UUID, position *InternalOrderPosition) (*InternalOrderPosition, *resty.Response, error)
	CreatePositionMany(ctx context.Context, id uuid.UUID, positions ...*InternalOrderPosition) (*Slice[InternalOrderPosition], *resty.Response, error)
	DeletePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (bool, *resty.Response, error)
	DeletePositionMany(ctx context.Context, id uuid.UUID, entities ...*InternalOrderPosition) (*DeleteManyResponse, *resty.Response, error)
	GetPositionTrackingCodes(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (*MetaArray[TrackingCode], *resty.Response, error)
	CreateUpdatePositionTrackingCodeMany(ctx context.Context, id uuid.UUID, positionID uuid.UUID, trackingCodes ...*TrackingCode) (*Slice[TrackingCode], *resty.Response, error)
	DeletePositionTrackingCodeMany(ctx context.Context, id uuid.UUID, positionID uuid.UUID, trackingCodes ...*TrackingCode) (*DeleteManyResponse, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributeMany(ctx context.Context, attributes ...*Attribute) (*Slice[Attribute], *resty.Response, error)
	UpdateAttribute(ctx context.Context, id uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributeMany(ctx context.Context, attributes ...*Attribute) (*DeleteManyResponse, *resty.Response, error)
	GetPublications(ctx context.Context, id uuid.UUID) (*MetaArray[Publication], *resty.Response, error)
	GetPublicationByID(ctx context.Context, id uuid.UUID, publicationID uuid.UUID) (*Publication, *resty.Response, error)
	Publish(ctx context.Context, id uuid.UUID, template TemplateInterface) (*Publication, *resty.Response, error)
	DeletePublication(ctx context.Context, id uuid.UUID, publicationID uuid.UUID) (bool, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*InternalOrder, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error)
	MoveToTrash(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetStateByID(ctx context.Context, id uuid.UUID) (*State, *resty.Response, error)
	CreateState(ctx context.Context, state *State) (*State, *resty.Response, error)
	UpdateState(ctx context.Context, id uuid.UUID, state *State) (*State, *resty.Response, error)
	CreateUpdateStateMany(ctx context.Context, states ...*State) (*Slice[State], *resty.Response, error)
	DeleteState(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetFiles(ctx context.Context, id uuid.UUID) (*MetaArray[File], *resty.Response, error)
	CreateFile(ctx context.Context, id uuid.UUID, file *File) (*Slice[File], *resty.Response, error)
	UpdateFileMany(ctx context.Context, id uuid.UUID, files ...*File) (*Slice[File], *resty.Response, error)
	DeleteFile(ctx context.Context, id uuid.UUID, fileID uuid.UUID) (bool, *resty.Response, error)
	DeleteFileMany(ctx context.Context, id uuid.UUID, files ...*File) (*DeleteManyResponse, *resty.Response, error)
	Evaluate(ctx context.Context, entity *InternalOrder, evaluate ...Evaluate) (*InternalOrder, *resty.Response, error)
}

func NewInternalOrderService(client *Client) InternalOrderService {
	e := NewEndpoint(client, "entity/internalorder")
	return newMainService[InternalOrder, InternalOrderPosition, MetaAttributesSharedStatesWrapper, any](e)
}
