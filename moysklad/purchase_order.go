package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// PurchaseOrder Заказ поставщику.
// Ключевое слово: purchaseorder
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-zakaz-postawschiku
type PurchaseOrder struct {
	PayedSum              *float64                          `json:"payedSum,omitempty"`
	Applicable            *bool                             `json:"applicable,omitempty"`
	AgentAccount          *AgentAccount                     `json:"agentAccount,omitempty"`
	Owner                 *Employee                         `json:"owner,omitempty"`
	InternalOrder         *InternalOrder                    `json:"internalOrder,omitempty"`
	Code                  *string                           `json:"code,omitempty"`
	Contract              *NullValue[Contract]              `json:"contract,omitempty"`
	Created               *Timestamp                        `json:"created,omitempty"`
	Deleted               *Timestamp                        `json:"deleted,omitempty"`
	DeliveryPlannedMoment *Timestamp                        `json:"deliveryPlannedMoment,omitempty"`
	OrganizationAccount   *AgentAccount                     `json:"organizationAccount,omitempty"`
	ExternalCode          *string                           `json:"externalCode,omitempty"`
	AccountID             *uuid.UUID                        `json:"accountId,omitempty"`
	Group                 *Group                            `json:"group,omitempty"`
	ID                    *uuid.UUID                        `json:"id,omitempty"`
	InvoicedSum           *float64                          `json:"invoicedSum,omitempty"`
	Meta                  *Meta                             `json:"meta,omitempty"`
	Moment                *Timestamp                        `json:"moment,omitempty"`
	Name                  *string                           `json:"name,omitempty"`
	Organization          *Organization                     `json:"organization,omitempty"`
	Description           *string                           `json:"description,omitempty"`
	Agent                 *Counterparty                     `json:"agent,omitempty"`
	Files                 *MetaArray[File]                  `json:"files,omitempty"`
	Positions             *Positions[PurchaseOrderPosition] `json:"positions,omitempty"`
	Printed               *bool                             `json:"printed,omitempty"`
	Project               *NullValue[Project]               `json:"project,omitempty"`
	Published             *bool                             `json:"published,omitempty"`
	Rate                  *NullValue[Rate]                  `json:"rate,omitempty"`
	Shared                *bool                             `json:"shared,omitempty"`
	ShippedSum            *float64                          `json:"shippedSum,omitempty"`
	State                 *NullValue[State]                 `json:"state,omitempty"`
	Store                 *Store                            `json:"store,omitempty"`
	Sum                   *float64                          `json:"sum,omitempty"`
	SyncID                *uuid.UUID                        `json:"syncId,omitempty"`
	Updated               *Timestamp                        `json:"updated,omitempty"`
	VatEnabled            *bool                             `json:"vatEnabled,omitempty"`
	VatIncluded           *bool                             `json:"vatIncluded,omitempty"`
	VatSum                *float64                          `json:"vatSum,omitempty"`
	WaitSum               *float64                          `json:"waitSum,omitempty"`
	CustomerOrders        Slice[CustomerOrder]              `json:"customerOrders,omitempty"`
	InvoicesIn            Slice[InvoiceIn]                  `json:"invoicesIn,omitempty"`
	Payments              Slice[Payment]                    `json:"payments,omitempty"`
	Supplies              Slice[Supply]                     `json:"supplies,omitempty"`
	Attributes            Slice[Attribute]                  `json:"attributes,omitempty"`
}

// Clean возвращает сущность с единственным заполненным полем Meta
func (purchaseOrder PurchaseOrder) Clean() *PurchaseOrder {
	return &PurchaseOrder{Meta: purchaseOrder.Meta}
}

// AsOperation возвращает объект Operation c полем Meta сущности
func (purchaseOrder PurchaseOrder) AsOperation() *Operation {
	return &Operation{Meta: purchaseOrder.GetMeta()}
}

// AsTaskOperation реализует интерфейс AsTaskOperationInterface
func (purchaseOrder PurchaseOrder) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: purchaseOrder.Meta}
}

func (purchaseOrder PurchaseOrder) GetPayedSum() float64 {
	return Deref(purchaseOrder.PayedSum)
}

func (purchaseOrder PurchaseOrder) GetApplicable() bool {
	return Deref(purchaseOrder.Applicable)
}

func (purchaseOrder PurchaseOrder) GetAgentAccount() AgentAccount {
	return Deref(purchaseOrder.AgentAccount)
}

func (purchaseOrder PurchaseOrder) GetOwner() Employee {
	return Deref(purchaseOrder.Owner)
}

func (purchaseOrder PurchaseOrder) GetInternalOrder() InternalOrder {
	return Deref(purchaseOrder.InternalOrder)
}

func (purchaseOrder PurchaseOrder) GetCode() string {
	return Deref(purchaseOrder.Code)
}

func (purchaseOrder PurchaseOrder) GetContract() Contract {
	return purchaseOrder.Contract.Get()
}

func (purchaseOrder PurchaseOrder) GetCreated() Timestamp {
	return Deref(purchaseOrder.Created)
}

func (purchaseOrder PurchaseOrder) GetDeleted() Timestamp {
	return Deref(purchaseOrder.Deleted)
}

func (purchaseOrder PurchaseOrder) GetDeliveryPlannedMoment() Timestamp {
	return Deref(purchaseOrder.DeliveryPlannedMoment)
}

func (purchaseOrder PurchaseOrder) GetOrganizationAccount() AgentAccount {
	return Deref(purchaseOrder.OrganizationAccount)
}

func (purchaseOrder PurchaseOrder) GetExternalCode() string {
	return Deref(purchaseOrder.ExternalCode)
}

func (purchaseOrder PurchaseOrder) GetAccountID() uuid.UUID {
	return Deref(purchaseOrder.AccountID)
}

func (purchaseOrder PurchaseOrder) GetGroup() Group {
	return Deref(purchaseOrder.Group)
}

func (purchaseOrder PurchaseOrder) GetID() uuid.UUID {
	return Deref(purchaseOrder.ID)
}

func (purchaseOrder PurchaseOrder) GetInvoicedSum() float64 {
	return Deref(purchaseOrder.InvoicedSum)
}

func (purchaseOrder PurchaseOrder) GetMeta() Meta {
	return Deref(purchaseOrder.Meta)
}

func (purchaseOrder PurchaseOrder) GetMoment() Timestamp {
	return Deref(purchaseOrder.Moment)
}

func (purchaseOrder PurchaseOrder) GetName() string {
	return Deref(purchaseOrder.Name)
}

func (purchaseOrder PurchaseOrder) GetOrganization() Organization {
	return Deref(purchaseOrder.Organization)
}

func (purchaseOrder PurchaseOrder) GetDescription() string {
	return Deref(purchaseOrder.Description)
}

func (purchaseOrder PurchaseOrder) GetAgent() Counterparty {
	return Deref(purchaseOrder.Agent)
}

func (purchaseOrder PurchaseOrder) GetFiles() MetaArray[File] {
	return Deref(purchaseOrder.Files)
}

func (purchaseOrder PurchaseOrder) GetPositions() Positions[PurchaseOrderPosition] {
	return Deref(purchaseOrder.Positions)
}

func (purchaseOrder PurchaseOrder) GetPrinted() bool {
	return Deref(purchaseOrder.Printed)
}

func (purchaseOrder PurchaseOrder) GetProject() Project {
	return purchaseOrder.Project.Get()
}

func (purchaseOrder PurchaseOrder) GetPublished() bool {
	return Deref(purchaseOrder.Published)
}

func (purchaseOrder PurchaseOrder) GetRate() Rate {
	return purchaseOrder.Rate.Get()
}

func (purchaseOrder PurchaseOrder) GetShared() bool {
	return Deref(purchaseOrder.Shared)
}

func (purchaseOrder PurchaseOrder) GetShippedSum() float64 {
	return Deref(purchaseOrder.ShippedSum)
}

func (purchaseOrder PurchaseOrder) GetState() State {
	return purchaseOrder.State.Get()
}

func (purchaseOrder PurchaseOrder) GetStore() Store {
	return Deref(purchaseOrder.Store)
}

func (purchaseOrder PurchaseOrder) GetSum() float64 {
	return Deref(purchaseOrder.Sum)
}

func (purchaseOrder PurchaseOrder) GetSyncID() uuid.UUID {
	return Deref(purchaseOrder.SyncID)
}

func (purchaseOrder PurchaseOrder) GetUpdated() Timestamp {
	return Deref(purchaseOrder.Updated)
}

func (purchaseOrder PurchaseOrder) GetVatEnabled() bool {
	return Deref(purchaseOrder.VatEnabled)
}

func (purchaseOrder PurchaseOrder) GetVatIncluded() bool {
	return Deref(purchaseOrder.VatIncluded)
}

func (purchaseOrder PurchaseOrder) GetVatSum() float64 {
	return Deref(purchaseOrder.VatSum)
}

func (purchaseOrder PurchaseOrder) GetWaitSum() float64 {
	return Deref(purchaseOrder.WaitSum)
}

func (purchaseOrder PurchaseOrder) GetCustomerOrders() Slice[CustomerOrder] {
	return purchaseOrder.CustomerOrders
}

func (purchaseOrder PurchaseOrder) GetInvoicesIn() Slice[InvoiceIn] {
	return purchaseOrder.InvoicesIn
}

func (purchaseOrder PurchaseOrder) GetPayments() Slice[Payment] {
	return purchaseOrder.Payments
}

func (purchaseOrder PurchaseOrder) GetSupplies() Slice[Supply] {
	return purchaseOrder.Supplies
}

func (purchaseOrder PurchaseOrder) GetAttributes() Slice[Attribute] {
	return purchaseOrder.Attributes
}

func (purchaseOrder *PurchaseOrder) SetApplicable(applicable bool) *PurchaseOrder {
	purchaseOrder.Applicable = &applicable
	return purchaseOrder
}

func (purchaseOrder *PurchaseOrder) SetAgentAccount(agentAccount *AgentAccount) *PurchaseOrder {
	purchaseOrder.AgentAccount = agentAccount.Clean()
	return purchaseOrder
}

func (purchaseOrder *PurchaseOrder) SetOwner(owner *Employee) *PurchaseOrder {
	purchaseOrder.Owner = owner.Clean()
	return purchaseOrder
}

func (purchaseOrder *PurchaseOrder) SetInternalOrder(internalOrder *InternalOrder) *PurchaseOrder {
	purchaseOrder.InternalOrder = internalOrder.Clean()
	return purchaseOrder
}

func (purchaseOrder *PurchaseOrder) SetCode(code string) *PurchaseOrder {
	purchaseOrder.Code = &code
	return purchaseOrder
}

func (purchaseOrder *PurchaseOrder) SetContract(contract *Contract) *PurchaseOrder {
	purchaseOrder.Contract = NewNullValueFrom(contract.Clean())
	return purchaseOrder
}

func (purchaseOrder *PurchaseOrder) SetDeliveryPlannedMoment(deliveryPlannedMoment *Timestamp) *PurchaseOrder {
	purchaseOrder.DeliveryPlannedMoment = deliveryPlannedMoment
	return purchaseOrder
}

func (purchaseOrder *PurchaseOrder) SetOrganizationAccount(organizationAccount *AgentAccount) *PurchaseOrder {
	purchaseOrder.OrganizationAccount = organizationAccount.Clean()
	return purchaseOrder
}

func (purchaseOrder *PurchaseOrder) SetExternalCode(externalCode string) *PurchaseOrder {
	purchaseOrder.ExternalCode = &externalCode
	return purchaseOrder
}

func (purchaseOrder *PurchaseOrder) SetGroup(group *Group) *PurchaseOrder {
	purchaseOrder.Group = group.Clean()
	return purchaseOrder
}

func (purchaseOrder *PurchaseOrder) SetMeta(meta *Meta) *PurchaseOrder {
	purchaseOrder.Meta = meta
	return purchaseOrder
}

func (purchaseOrder *PurchaseOrder) SetMoment(moment *Timestamp) *PurchaseOrder {
	purchaseOrder.Moment = moment
	return purchaseOrder
}

func (purchaseOrder *PurchaseOrder) SetName(name string) *PurchaseOrder {
	purchaseOrder.Name = &name
	return purchaseOrder
}

func (purchaseOrder *PurchaseOrder) SetOrganization(organization *Organization) *PurchaseOrder {
	purchaseOrder.Organization = organization
	return purchaseOrder
}

func (purchaseOrder *PurchaseOrder) SetDescription(description string) *PurchaseOrder {
	purchaseOrder.Description = &description
	return purchaseOrder
}

func (purchaseOrder *PurchaseOrder) SetAgent(agent *Counterparty) *PurchaseOrder {
	purchaseOrder.Agent = agent.Clean()
	return purchaseOrder
}

func (purchaseOrder *PurchaseOrder) SetFiles(files ...*File) *PurchaseOrder {
	purchaseOrder.Files = NewMetaArrayFrom(files)
	return purchaseOrder
}

func (purchaseOrder *PurchaseOrder) SetPositions(positions ...*PurchaseOrderPosition) *PurchaseOrder {
	purchaseOrder.Positions = NewPositionsFrom(positions)
	return purchaseOrder
}

func (purchaseOrder *PurchaseOrder) SetProject(project *Project) *PurchaseOrder {
	purchaseOrder.Project = NewNullValueFrom(project.Clean())
	return purchaseOrder
}

func (purchaseOrder *PurchaseOrder) SetNullProject() *PurchaseOrder {
	purchaseOrder.Project = NewNullValue[Project]()
	return purchaseOrder
}

func (purchaseOrder *PurchaseOrder) SetRate(rate *Rate) *PurchaseOrder {
	purchaseOrder.Rate = NewNullValueFrom(rate)
	return purchaseOrder
}

func (purchaseOrder *PurchaseOrder) SetNullRate() *PurchaseOrder {
	purchaseOrder.Rate = NewNullValue[Rate]()
	return purchaseOrder
}

func (purchaseOrder *PurchaseOrder) SetShared(shared bool) *PurchaseOrder {
	purchaseOrder.Shared = &shared
	return purchaseOrder
}

func (purchaseOrder *PurchaseOrder) SetState(state *State) *PurchaseOrder {
	purchaseOrder.State = NewNullValueFrom(state.Clean())
	return purchaseOrder
}

func (purchaseOrder *PurchaseOrder) SetNullState() *PurchaseOrder {
	purchaseOrder.State = NewNullValue[State]()
	return purchaseOrder
}

func (purchaseOrder *PurchaseOrder) SetStore(store *Store) *PurchaseOrder {
	purchaseOrder.Store = store.Clean()
	return purchaseOrder
}

func (purchaseOrder *PurchaseOrder) SetSyncID(syncID uuid.UUID) *PurchaseOrder {
	purchaseOrder.SyncID = &syncID
	return purchaseOrder
}

func (purchaseOrder *PurchaseOrder) SetVatEnabled(vatEnabled bool) *PurchaseOrder {
	purchaseOrder.VatEnabled = &vatEnabled
	return purchaseOrder
}

func (purchaseOrder *PurchaseOrder) SetVatIncluded(vatIncluded bool) *PurchaseOrder {
	purchaseOrder.VatIncluded = &vatIncluded
	return purchaseOrder
}

func (purchaseOrder *PurchaseOrder) SetCustomerOrders(customerOrders ...*CustomerOrder) *PurchaseOrder {
	purchaseOrder.CustomerOrders = customerOrders
	return purchaseOrder
}

func (purchaseOrder *PurchaseOrder) SetInvoicesIn(invoicesIn ...*InvoiceIn) *PurchaseOrder {
	purchaseOrder.InvoicesIn = invoicesIn
	return purchaseOrder
}

func (purchaseOrder *PurchaseOrder) SetPayments(payments ...*Payment) *PurchaseOrder {
	purchaseOrder.Payments = payments
	return purchaseOrder
}

func (purchaseOrder *PurchaseOrder) SetSupplies(supplies ...*Supply) *PurchaseOrder {
	purchaseOrder.Supplies = supplies
	return purchaseOrder
}

func (purchaseOrder *PurchaseOrder) SetAttributes(attributes ...*Attribute) *PurchaseOrder {
	purchaseOrder.Attributes = attributes
	return purchaseOrder
}

func (purchaseOrder PurchaseOrder) String() string {
	return Stringify(purchaseOrder)
}

// MetaType возвращает тип сущности.
func (PurchaseOrder) MetaType() MetaType {
	return MetaTypePurchaseOrder
}

// Update shortcut
func (purchaseOrder PurchaseOrder) Update(ctx context.Context, client *Client, params ...*Params) (*PurchaseOrder, *resty.Response, error) {
	return client.Entity().PurchaseOrder().Update(ctx, purchaseOrder.GetID(), &purchaseOrder, params...)
}

// Create shortcut
func (purchaseOrder PurchaseOrder) Create(ctx context.Context, client *Client, params ...*Params) (*PurchaseOrder, *resty.Response, error) {
	return client.Entity().PurchaseOrder().Create(ctx, &purchaseOrder, params...)
}

// Delete shortcut
func (purchaseOrder PurchaseOrder) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return client.Entity().PurchaseOrder().Delete(ctx, purchaseOrder.GetID())
}

// PurchaseOrderPosition Позиция Заказа поставщику.
// Ключевое слово: purchaseorderposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-zakaz-postawschiku-zakazy-postawschikam-pozicii-zakaza-postawschiku
type PurchaseOrderPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	Discount   *float64            `json:"discount,omitempty"`   // Процент скидки или наценки. Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Price      *float64            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Shipped    *float64            `json:"shipped,omitempty"`    // Принято
	InTransit  *float64            `json:"inTransit,omitempty"`  // Ожидание
	Vat        *float64            `json:"vat,omitempty"`        // НДС, которым облагается текущая позиция
	VatEnabled *bool               `json:"vatEnabled,omitempty"` // Включен ли НДС для позиции. С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
	Wait       *bool               `json:"wait,omitempty"`       // Ожидается данной позиции
	Stock      *Stock              `json:"stock,omitempty"`      // Остатки и себестоимость `?fields=stock&expand=positions`
}

func (purchaseOrderPosition PurchaseOrderPosition) GetAccountID() uuid.UUID {
	return Deref(purchaseOrderPosition.AccountID)
}

func (purchaseOrderPosition PurchaseOrderPosition) GetAssortment() AssortmentPosition {
	return Deref(purchaseOrderPosition.Assortment)
}

func (purchaseOrderPosition PurchaseOrderPosition) GetDiscount() float64 {
	return Deref(purchaseOrderPosition.Discount)
}

func (purchaseOrderPosition PurchaseOrderPosition) GetID() uuid.UUID {
	return Deref(purchaseOrderPosition.ID)
}

func (purchaseOrderPosition PurchaseOrderPosition) GetPack() Pack {
	return Deref(purchaseOrderPosition.Pack)
}

func (purchaseOrderPosition PurchaseOrderPosition) GetPrice() float64 {
	return Deref(purchaseOrderPosition.Price)
}

func (purchaseOrderPosition PurchaseOrderPosition) GetQuantity() float64 {
	return Deref(purchaseOrderPosition.Quantity)
}

func (purchaseOrderPosition PurchaseOrderPosition) GetShipped() float64 {
	return Deref(purchaseOrderPosition.Shipped)
}

func (purchaseOrderPosition PurchaseOrderPosition) GetInTransit() float64 {
	return Deref(purchaseOrderPosition.InTransit)
}

func (purchaseOrderPosition PurchaseOrderPosition) GetVat() float64 {
	return Deref(purchaseOrderPosition.Vat)
}

func (purchaseOrderPosition PurchaseOrderPosition) GetVatEnabled() bool {
	return Deref(purchaseOrderPosition.VatEnabled)
}

func (purchaseOrderPosition PurchaseOrderPosition) GetWait() bool {
	return Deref(purchaseOrderPosition.Wait)
}

func (purchaseOrderPosition PurchaseOrderPosition) GetStock() Stock {
	return Deref(purchaseOrderPosition.Stock)
}

func (purchaseOrderPosition *PurchaseOrderPosition) SetAssortment(assortment AsAssortment) *PurchaseOrderPosition {
	purchaseOrderPosition.Assortment = assortment.AsAssortment()
	return purchaseOrderPosition
}

func (purchaseOrderPosition *PurchaseOrderPosition) SetDiscount(discount float64) *PurchaseOrderPosition {
	purchaseOrderPosition.Discount = &discount
	return purchaseOrderPosition
}

func (purchaseOrderPosition *PurchaseOrderPosition) SetPack(pack *Pack) *PurchaseOrderPosition {
	purchaseOrderPosition.Pack = pack
	return purchaseOrderPosition
}

func (purchaseOrderPosition *PurchaseOrderPosition) SetPrice(price float64) *PurchaseOrderPosition {
	purchaseOrderPosition.Price = &price
	return purchaseOrderPosition
}

func (purchaseOrderPosition *PurchaseOrderPosition) SetQuantity(quantity float64) *PurchaseOrderPosition {
	purchaseOrderPosition.Quantity = &quantity
	return purchaseOrderPosition
}

func (purchaseOrderPosition *PurchaseOrderPosition) SetShipped(shipped float64) *PurchaseOrderPosition {
	purchaseOrderPosition.Shipped = &shipped
	return purchaseOrderPosition
}

func (purchaseOrderPosition *PurchaseOrderPosition) SetInTransit(inTransit float64) *PurchaseOrderPosition {
	purchaseOrderPosition.InTransit = &inTransit
	return purchaseOrderPosition
}

func (purchaseOrderPosition *PurchaseOrderPosition) SetVat(vat float64) *PurchaseOrderPosition {
	purchaseOrderPosition.Vat = &vat
	return purchaseOrderPosition
}

func (purchaseOrderPosition *PurchaseOrderPosition) SetVatEnabled(vatEnabled bool) *PurchaseOrderPosition {
	purchaseOrderPosition.VatEnabled = &vatEnabled
	return purchaseOrderPosition
}

func (purchaseOrderPosition *PurchaseOrderPosition) SetWait(wait bool) *PurchaseOrderPosition {
	purchaseOrderPosition.Wait = &wait
	return purchaseOrderPosition
}

func (purchaseOrderPosition PurchaseOrderPosition) String() string {
	return Stringify(purchaseOrderPosition)
}

// MetaType возвращает тип сущности.
func (PurchaseOrderPosition) MetaType() MetaType {
	return MetaTypePurchaseOrderPosition
}

// PurchaseOrderService
// Сервис для работы с заказами поставщикам.
type PurchaseOrderService interface {
	GetList(ctx context.Context, params ...*Params) (*List[PurchaseOrder], *resty.Response, error)
	Create(ctx context.Context, purchaseOrder *PurchaseOrder, params ...*Params) (*PurchaseOrder, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, purchaseOrderList Slice[PurchaseOrder], params ...*Params) (*Slice[PurchaseOrder], *resty.Response, error)
	DeleteMany(ctx context.Context, entities ...*PurchaseOrder) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*PurchaseOrder, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, purchaseOrder *PurchaseOrder, params ...*Params) (*PurchaseOrder, *resty.Response, error)
	Template(ctx context.Context) (*PurchaseOrder, *resty.Response, error)
	TemplateBased(ctx context.Context, basedOn ...MetaOwner) (*PurchaseOrder, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetaAttributesSharedStatesWrapper, *resty.Response, error)
	GetPositions(ctx context.Context, id uuid.UUID, params ...*Params) (*MetaArray[PurchaseOrderPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params ...*Params) (*PurchaseOrderPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *PurchaseOrderPosition, params ...*Params) (*PurchaseOrderPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id uuid.UUID, position *PurchaseOrderPosition) (*PurchaseOrderPosition, *resty.Response, error)
	CreatePositionMany(ctx context.Context, id uuid.UUID, positions ...*PurchaseOrderPosition) (*Slice[PurchaseOrderPosition], *resty.Response, error)
	DeletePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (bool, *resty.Response, error)
	DeletePositionMany(ctx context.Context, id uuid.UUID, entities ...*PurchaseOrderPosition) (*DeleteManyResponse, *resty.Response, error)
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
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*PurchaseOrder, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params ...*Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)
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
	Evaluate(ctx context.Context, entity *PurchaseOrder, evaluate ...Evaluate) (*PurchaseOrder, *resty.Response, error)
}

func NewPurchaseOrderService(client *Client) PurchaseOrderService {
	e := NewEndpoint(client, "entity/purchaseorder")
	return newMainService[PurchaseOrder, PurchaseOrderPosition, MetaAttributesSharedStatesWrapper, any](e)
}
