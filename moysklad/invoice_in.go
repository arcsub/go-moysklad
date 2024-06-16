package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// InvoiceIn Счет поставщика.
// Ключевое слово: invoicein
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-schet-postawschika
type InvoiceIn struct {
	OrganizationAccount  *AgentAccount                 `json:"organizationAccount,omitempty"`
	Created              *Timestamp                    `json:"created,omitempty"`
	PayedSum             *float64                      `json:"payedSum,omitempty"`
	Applicable           *bool                         `json:"applicable,omitempty"`
	Supplies             Slice[Supply]                 `json:"supplies,omitempty"`
	Code                 *string                       `json:"code,omitempty"`
	Contract             *NullValue[Contract]          `json:"contract,omitempty"`
	Owner                *Employee                     `json:"owner,omitempty"`
	Deleted              *Timestamp                    `json:"deleted,omitempty"`
	Description          *string                       `json:"description,omitempty"`
	ExternalCode         *string                       `json:"externalCode,omitempty"`
	Files                *MetaArray[File]              `json:"files,omitempty"`
	AccountID            *uuid.UUID                    `json:"accountId,omitempty"`
	ID                   *uuid.UUID                    `json:"id,omitempty"`
	IncomingDate         *Timestamp                    `json:"incomingDate,omitempty"`
	IncomingNumber       *string                       `json:"incomingNumber,omitempty"`
	Meta                 *Meta                         `json:"meta,omitempty"`
	Moment               *Timestamp                    `json:"moment,omitempty"`
	Name                 *string                       `json:"name,omitempty"`
	Organization         *Organization                 `json:"organization,omitempty"`
	Group                *Group                        `json:"group,omitempty"`
	Agent                *Counterparty                 `json:"agent,omitempty"`
	AgentAccount         *AgentAccount                 `json:"agentAccount,omitempty"`
	PaymentPlannedMoment *Timestamp                    `json:"paymentPlannedMoment,omitempty"`
	Positions            *Positions[InvoiceInPosition] `json:"positions,omitempty"`
	Printed              *bool                         `json:"printed,omitempty"`
	Project              *NullValue[Project]           `json:"project,omitempty"`
	Published            *bool                         `json:"published,omitempty"`
	Rate                 *NullValue[Rate]              `json:"rate,omitempty"`
	Shared               *bool                         `json:"shared,omitempty"`
	ShippedSum           *float64                      `json:"shippedSum,omitempty"`
	State                *NullValue[State]             `json:"state,omitempty"`
	Store                *NullValue[Store]             `json:"store,omitempty"`
	Sum                  *float64                      `json:"sum,omitempty"`
	SyncID               *uuid.UUID                    `json:"syncId,omitempty"`
	Updated              *Timestamp                    `json:"updated,omitempty"`
	VatEnabled           *bool                         `json:"vatEnabled,omitempty"`
	VatIncluded          *bool                         `json:"vatIncluded,omitempty"`
	VatSum               *float64                      `json:"vatSum,omitempty"`
	Payments             Slice[Payment]                `json:"payments,omitempty"`
	PurchaseOrder        *PurchaseOrder                `json:"purchaseOrder,omitempty"`
	Attributes           Slice[AttributeValue]         `json:"attributes,omitempty"`
}

// Clean возвращает сущность с единственным заполненным полем Meta
func (invoiceIn InvoiceIn) Clean() *InvoiceIn {
	return &InvoiceIn{Meta: invoiceIn.Meta}
}

// AsOperation возвращает объект Operation c полем Meta сущности
func (invoiceIn InvoiceIn) AsOperation() *Operation {
	return &Operation{Meta: invoiceIn.GetMeta()}
}

func (invoiceIn InvoiceIn) GetOrganizationAccount() AgentAccount {
	return Deref(invoiceIn.OrganizationAccount)
}

func (invoiceIn InvoiceIn) GetCreated() Timestamp {
	return Deref(invoiceIn.Created)
}

func (invoiceIn InvoiceIn) GetPayedSum() float64 {
	return Deref(invoiceIn.PayedSum)
}

func (invoiceIn InvoiceIn) GetApplicable() bool {
	return Deref(invoiceIn.Applicable)
}

func (invoiceIn InvoiceIn) GetSupplies() Slice[Supply] {
	return invoiceIn.Supplies
}

func (invoiceIn InvoiceIn) GetCode() string {
	return Deref(invoiceIn.Code)
}

func (invoiceIn InvoiceIn) GetContract() Contract {
	return invoiceIn.Contract.Get()
}

func (invoiceIn InvoiceIn) GetOwner() Employee {
	return Deref(invoiceIn.Owner)
}

func (invoiceIn InvoiceIn) GetDeleted() Timestamp {
	return Deref(invoiceIn.Deleted)
}

func (invoiceIn InvoiceIn) GetDescription() string {
	return Deref(invoiceIn.Description)
}

func (invoiceIn InvoiceIn) GetExternalCode() string {
	return Deref(invoiceIn.ExternalCode)
}

func (invoiceIn InvoiceIn) GetFiles() MetaArray[File] {
	return Deref(invoiceIn.Files)
}

func (invoiceIn InvoiceIn) GetAccountID() uuid.UUID {
	return Deref(invoiceIn.AccountID)
}

func (invoiceIn InvoiceIn) GetID() uuid.UUID {
	return Deref(invoiceIn.ID)
}

func (invoiceIn InvoiceIn) GetIncomingDate() Timestamp {
	return Deref(invoiceIn.IncomingDate)
}

func (invoiceIn InvoiceIn) GetIncomingNumber() string {
	return Deref(invoiceIn.IncomingNumber)
}

func (invoiceIn InvoiceIn) GetMeta() Meta {
	return Deref(invoiceIn.Meta)
}

func (invoiceIn InvoiceIn) GetMoment() Timestamp {
	return Deref(invoiceIn.Moment)
}

func (invoiceIn InvoiceIn) GetName() string {
	return Deref(invoiceIn.Name)
}

func (invoiceIn InvoiceIn) GetOrganization() Organization {
	return Deref(invoiceIn.Organization)
}

func (invoiceIn InvoiceIn) GetGroup() Group {
	return Deref(invoiceIn.Group)
}

func (invoiceIn InvoiceIn) GetAgent() Counterparty {
	return Deref(invoiceIn.Agent)
}

func (invoiceIn InvoiceIn) GetAgentAccount() AgentAccount {
	return Deref(invoiceIn.AgentAccount)
}

func (invoiceIn InvoiceIn) GetPaymentPlannedMoment() Timestamp {
	return Deref(invoiceIn.PaymentPlannedMoment)
}

func (invoiceIn InvoiceIn) GetPositions() Positions[InvoiceInPosition] {
	return Deref(invoiceIn.Positions)
}

func (invoiceIn InvoiceIn) GetPrinted() bool {
	return Deref(invoiceIn.Printed)
}

func (invoiceIn InvoiceIn) GetProject() Project {
	return invoiceIn.Project.Get()
}

func (invoiceIn InvoiceIn) GetPublished() bool {
	return Deref(invoiceIn.Published)
}

func (invoiceIn InvoiceIn) GetRate() Rate {
	return invoiceIn.Rate.Get()
}

func (invoiceIn InvoiceIn) GetShared() bool {
	return Deref(invoiceIn.Shared)
}

func (invoiceIn InvoiceIn) GetShippedSum() float64 {
	return Deref(invoiceIn.ShippedSum)
}

func (invoiceIn InvoiceIn) GetState() State {
	return invoiceIn.State.Get()
}

func (invoiceIn InvoiceIn) GetStore() Store {
	return invoiceIn.Store.Get()
}

func (invoiceIn InvoiceIn) GetSum() float64 {
	return Deref(invoiceIn.Sum)
}

func (invoiceIn InvoiceIn) GetSyncID() uuid.UUID {
	return Deref(invoiceIn.SyncID)
}

func (invoiceIn InvoiceIn) GetUpdated() Timestamp {
	return Deref(invoiceIn.Updated)
}

func (invoiceIn InvoiceIn) GetVatEnabled() bool {
	return Deref(invoiceIn.VatEnabled)
}

func (invoiceIn InvoiceIn) GetVatIncluded() bool {
	return Deref(invoiceIn.VatIncluded)
}

func (invoiceIn InvoiceIn) GetVatSum() float64 {
	return Deref(invoiceIn.VatSum)
}

func (invoiceIn InvoiceIn) GetPayments() Slice[Payment] {
	return invoiceIn.Payments
}

func (invoiceIn InvoiceIn) GetPurchaseOrder() PurchaseOrder {
	return Deref(invoiceIn.PurchaseOrder)
}

func (invoiceIn InvoiceIn) GetAttributes() Slice[AttributeValue] {
	return invoiceIn.Attributes
}

func (invoiceIn *InvoiceIn) SetOrganizationAccount(organizationAccount *AgentAccount) *InvoiceIn {
	invoiceIn.OrganizationAccount = organizationAccount.Clean()
	return invoiceIn
}

func (invoiceIn *InvoiceIn) SetApplicable(applicable bool) *InvoiceIn {
	invoiceIn.Applicable = &applicable
	return invoiceIn
}

func (invoiceIn *InvoiceIn) SetSupplies(supplies Slice[Supply]) *InvoiceIn {
	invoiceIn.Supplies = supplies
	return invoiceIn
}

func (invoiceIn *InvoiceIn) SetCode(code string) *InvoiceIn {
	invoiceIn.Code = &code
	return invoiceIn
}

func (invoiceIn *InvoiceIn) SetContract(contract *Contract) *InvoiceIn {
	invoiceIn.Contract = NewNullValueWith(contract.Clean())
	return invoiceIn
}

func (invoiceIn *InvoiceIn) SetNullContract() *InvoiceIn {
	invoiceIn.Contract = NewNullValue[Contract]()
	return invoiceIn
}

func (invoiceIn *InvoiceIn) SetOwner(owner *Employee) *InvoiceIn {
	invoiceIn.Owner = owner.Clean()
	return invoiceIn
}

func (invoiceIn *InvoiceIn) SetDescription(description string) *InvoiceIn {
	invoiceIn.Description = &description
	return invoiceIn
}

func (invoiceIn *InvoiceIn) SetExternalCode(externalCode string) *InvoiceIn {
	invoiceIn.ExternalCode = &externalCode
	return invoiceIn
}

func (invoiceIn *InvoiceIn) SetFiles(files Slice[File]) *InvoiceIn {
	invoiceIn.Files = NewMetaArrayRows(files)
	return invoiceIn
}

func (invoiceIn *InvoiceIn) SetIncomingDate(incomingDate *Timestamp) *InvoiceIn {
	invoiceIn.IncomingDate = incomingDate
	return invoiceIn
}

func (invoiceIn *InvoiceIn) SetIncomingNumber(incomingNumber string) *InvoiceIn {
	invoiceIn.IncomingNumber = &incomingNumber
	return invoiceIn
}

func (invoiceIn *InvoiceIn) SetMeta(meta *Meta) *InvoiceIn {
	invoiceIn.Meta = meta
	return invoiceIn
}

func (invoiceIn *InvoiceIn) SetMoment(moment *Timestamp) *InvoiceIn {
	invoiceIn.Moment = moment
	return invoiceIn
}

func (invoiceIn *InvoiceIn) SetName(name string) *InvoiceIn {
	invoiceIn.Name = &name
	return invoiceIn
}

func (invoiceIn *InvoiceIn) SetOrganization(organization *Organization) *InvoiceIn {
	invoiceIn.Organization = organization.Clean()
	return invoiceIn
}

func (invoiceIn *InvoiceIn) SetGroup(group *Group) *InvoiceIn {
	invoiceIn.Group = group.Clean()
	return invoiceIn
}

func (invoiceIn *InvoiceIn) SetAgent(agent *Counterparty) *InvoiceIn {
	invoiceIn.Agent = agent.Clean()
	return invoiceIn
}

func (invoiceIn *InvoiceIn) SetAgentAccount(agentAccount *AgentAccount) *InvoiceIn {
	invoiceIn.AgentAccount = agentAccount.Clean()
	return invoiceIn
}

func (invoiceIn *InvoiceIn) SetPaymentPlannedMoment(paymentPlannedMoment *Timestamp) *InvoiceIn {
	invoiceIn.PaymentPlannedMoment = paymentPlannedMoment
	return invoiceIn
}

func (invoiceIn *InvoiceIn) SetPositions(positions *Positions[InvoiceInPosition]) *InvoiceIn {
	invoiceIn.Positions = positions
	return invoiceIn
}

func (invoiceIn *InvoiceIn) SetProject(project *Project) *InvoiceIn {
	invoiceIn.Project = NewNullValueWith(project.Clean())
	return invoiceIn
}

func (invoiceIn *InvoiceIn) SetNullProject() *InvoiceIn {
	invoiceIn.Project = NewNullValue[Project]()
	return invoiceIn
}

func (invoiceIn *InvoiceIn) SetRate(rate *Rate) *InvoiceIn {
	invoiceIn.Rate = NewNullValueWith(rate)
	return invoiceIn
}

func (invoiceIn *InvoiceIn) SetNullRate() *InvoiceIn {
	invoiceIn.Rate = NewNullValue[Rate]()
	return invoiceIn
}

func (invoiceIn *InvoiceIn) SetShared(shared bool) *InvoiceIn {
	invoiceIn.Shared = &shared
	return invoiceIn
}

func (invoiceIn *InvoiceIn) SetState(state *State) *InvoiceIn {
	invoiceIn.State = NewNullValueWith(state.Clean())
	return invoiceIn
}

func (invoiceIn *InvoiceIn) SetNullState() *InvoiceIn {
	invoiceIn.State = NewNullValue[State]()
	return invoiceIn
}

func (invoiceIn *InvoiceIn) SetStore(store *Store) *InvoiceIn {
	invoiceIn.Store = NewNullValueWith(store.Clean())
	return invoiceIn
}

func (invoiceIn *InvoiceIn) SetNullStore() *InvoiceIn {
	invoiceIn.Store = NewNullValue[Store]()
	return invoiceIn
}

func (invoiceIn *InvoiceIn) SetSyncID(syncID uuid.UUID) *InvoiceIn {
	invoiceIn.SyncID = &syncID
	return invoiceIn
}

func (invoiceIn *InvoiceIn) SetVatEnabled(vatEnabled bool) *InvoiceIn {
	invoiceIn.VatEnabled = &vatEnabled
	return invoiceIn
}

func (invoiceIn *InvoiceIn) SetVatIncluded(vatIncluded bool) *InvoiceIn {
	invoiceIn.VatIncluded = &vatIncluded
	return invoiceIn
}

func (invoiceIn *InvoiceIn) SetPayments(payments Slice[Payment]) *InvoiceIn {
	invoiceIn.Payments = payments
	return invoiceIn
}

func (invoiceIn *InvoiceIn) SetPurchaseOrder(purchaseOrder *PurchaseOrder) *InvoiceIn {
	invoiceIn.PurchaseOrder = purchaseOrder.Clean()
	return invoiceIn
}

func (invoiceIn *InvoiceIn) SetAttributes(attributes Slice[AttributeValue]) *InvoiceIn {
	invoiceIn.Attributes = attributes
	return invoiceIn
}

func (invoiceIn InvoiceIn) String() string {
	return Stringify(invoiceIn)
}

func (invoiceIn InvoiceIn) MetaType() MetaType {
	return MetaTypeInvoiceIn
}

// InvoiceInPosition Позиция Счета поставщика.
// Ключевое слово: invoiceposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-schet-postawschika-scheta-postawschikow-pozicii-scheta-postawschika
type InvoiceInPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`
	Assortment *AssortmentPosition `json:"assortment,omitempty"`
	Discount   *float64            `json:"discount,omitempty"`
	ID         *uuid.UUID          `json:"id,omitempty"`
	Pack       *Pack               `json:"pack,omitempty"`
	Meta       *Meta               `json:"meta,omitempty"`
	Price      *float64            `json:"price,omitempty"`
	Quantity   *float64            `json:"quantity,omitempty"`
	Vat        *int                `json:"vat,omitempty"`
	VatEnabled *bool               `json:"vatEnabled,omitempty"`
	Stock      *Stock              `json:"stock,omitempty"`
}

func (invoiceInPosition InvoiceInPosition) GetAccountID() uuid.UUID {
	return Deref(invoiceInPosition.AccountID)
}

func (invoiceInPosition InvoiceInPosition) GetAssortment() AssortmentPosition {
	return Deref(invoiceInPosition.Assortment)
}

func (invoiceInPosition InvoiceInPosition) GetDiscount() float64 {
	return Deref(invoiceInPosition.Discount)
}

func (invoiceInPosition InvoiceInPosition) GetID() uuid.UUID {
	return Deref(invoiceInPosition.ID)
}

func (invoiceInPosition InvoiceInPosition) GetPack() Pack {
	return Deref(invoiceInPosition.Pack)
}

func (invoiceInPosition InvoiceInPosition) GetMeta() Meta {
	return Deref(invoiceInPosition.Meta)
}

func (invoiceInPosition InvoiceInPosition) GetPrice() float64 {
	return Deref(invoiceInPosition.Price)
}

func (invoiceInPosition InvoiceInPosition) GetQuantity() float64 {
	return Deref(invoiceInPosition.Quantity)
}

func (invoiceInPosition InvoiceInPosition) GetVat() int {
	return Deref(invoiceInPosition.Vat)
}

func (invoiceInPosition InvoiceInPosition) GetVatEnabled() bool {
	return Deref(invoiceInPosition.VatEnabled)
}

func (invoiceInPosition InvoiceInPosition) GetStock() Stock {
	return Deref(invoiceInPosition.Stock)
}

func (invoiceInPosition *InvoiceInPosition) SetAssortment(assortment AsAssortment) *InvoiceInPosition {
	invoiceInPosition.Assortment = assortment.AsAssortment()
	return invoiceInPosition
}

func (invoiceInPosition *InvoiceInPosition) SetDiscount(discount float64) *InvoiceInPosition {
	invoiceInPosition.Discount = &discount
	return invoiceInPosition
}

func (invoiceInPosition *InvoiceInPosition) SetPack(pack *Pack) *InvoiceInPosition {
	invoiceInPosition.Pack = pack
	return invoiceInPosition
}

func (invoiceInPosition *InvoiceInPosition) SetMeta(meta *Meta) *InvoiceInPosition {
	invoiceInPosition.Meta = meta
	return invoiceInPosition
}

func (invoiceInPosition *InvoiceInPosition) SetPrice(price float64) *InvoiceInPosition {
	invoiceInPosition.Price = &price
	return invoiceInPosition
}

func (invoiceInPosition *InvoiceInPosition) SetQuantity(quantity float64) *InvoiceInPosition {
	invoiceInPosition.Quantity = &quantity
	return invoiceInPosition
}

func (invoiceInPosition *InvoiceInPosition) SetVat(vat int) *InvoiceInPosition {
	invoiceInPosition.Vat = &vat
	return invoiceInPosition
}

func (invoiceInPosition *InvoiceInPosition) SetVatEnabled(vatEnabled bool) *InvoiceInPosition {
	invoiceInPosition.VatEnabled = &vatEnabled
	return invoiceInPosition
}

func (invoiceInPosition InvoiceInPosition) String() string {
	return Stringify(invoiceInPosition)
}

func (invoiceInPosition InvoiceInPosition) MetaType() MetaType {
	return MetaTypeInvoicePosition
}

// InvoiceInService
// Сервис для работы со счетами поставщиков.
type InvoiceInService interface {
	GetList(ctx context.Context, params ...*Params) (*List[InvoiceIn], *resty.Response, error)
	Create(ctx context.Context, invoiceIn *InvoiceIn, params ...*Params) (*InvoiceIn, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, invoiceInList Slice[InvoiceIn], params ...*Params) (*Slice[InvoiceIn], *resty.Response, error)
	DeleteMany(ctx context.Context, invoiceInList []MetaWrapper) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*InvoiceIn, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, invoiceIn *InvoiceIn, params ...*Params) (*InvoiceIn, *resty.Response, error)
	Template(ctx context.Context) (*InvoiceIn, *resty.Response, error)
	TemplateBased(ctx context.Context, basedOn ...MetaOwner) (*InvoiceIn, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetaAttributesSharedStatesWrapper, *resty.Response, error)
	GetPositions(ctx context.Context, id uuid.UUID, params ...*Params) (*MetaArray[InvoiceInPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params ...*Params) (*InvoiceInPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *InvoiceInPosition, params ...*Params) (*InvoiceInPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id uuid.UUID, position *InvoiceInPosition) (*InvoiceInPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id uuid.UUID, positions Slice[InvoiceInPosition]) (*Slice[InvoiceInPosition], *resty.Response, error)
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
	DeleteAttributes(ctx context.Context, attributeList []MetaWrapper) (*DeleteManyResponse, *resty.Response, error)
	GetPublications(ctx context.Context, id uuid.UUID) (*MetaArray[Publication], *resty.Response, error)
	GetPublicationByID(ctx context.Context, id uuid.UUID, publicationID uuid.UUID) (*Publication, *resty.Response, error)
	Publish(ctx context.Context, id uuid.UUID, template Templater) (*Publication, *resty.Response, error)
	DeletePublication(ctx context.Context, id uuid.UUID, publicationID uuid.UUID) (bool, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*InvoiceIn, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error)
	MoveToTrash(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetStateByID(ctx context.Context, id uuid.UUID) (*State, *resty.Response, error)
	CreateState(ctx context.Context, state *State) (*State, *resty.Response, error)
	UpdateState(ctx context.Context, id uuid.UUID, state *State) (*State, *resty.Response, error)
	CreateOrUpdateStates(ctx context.Context, states Slice[State]) (*Slice[State], *resty.Response, error)
	DeleteState(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetFiles(ctx context.Context, id uuid.UUID) (*MetaArray[File], *resty.Response, error)
	CreateFile(ctx context.Context, id uuid.UUID, file *File) (*Slice[File], *resty.Response, error)
	UpdateFiles(ctx context.Context, id uuid.UUID, files Slice[File]) (*Slice[File], *resty.Response, error)
	DeleteFile(ctx context.Context, id uuid.UUID, fileID uuid.UUID) (bool, *resty.Response, error)
	DeleteFiles(ctx context.Context, id uuid.UUID, files []MetaWrapper) (*DeleteManyResponse, *resty.Response, error)
	Evaluate(ctx context.Context, entity *InvoiceIn, evaluate ...Evaluate) (*InvoiceIn, *resty.Response, error)
}

func NewInvoiceInService(client *Client) InvoiceInService {
	e := NewEndpoint(client, "entity/invoicein")
	return newMainService[InvoiceIn, InvoiceInPosition, MetaAttributesSharedStatesWrapper, any](e)
}
