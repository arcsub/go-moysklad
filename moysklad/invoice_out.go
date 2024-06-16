package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// InvoiceOut Счет покупателю.
// Ключевое слово: invoiceout
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-schet-pokupatelu
type InvoiceOut struct {
	PayedSum             *float64                       `json:"payedSum,omitempty"`
	VatEnabled           *bool                          `json:"vatEnabled,omitempty"`
	AgentAccount         *AgentAccount                  `json:"agentAccount,omitempty"`
	Applicable           *bool                          `json:"applicable,omitempty"`
	Demands              Slice[Demand]                  `json:"demands,omitempty"`
	Code                 *string                        `json:"code,omitempty"`
	OrganizationAccount  *AgentAccount                  `json:"organizationAccount,omitempty"`
	Created              *Timestamp                     `json:"created,omitempty"`
	Deleted              *Timestamp                     `json:"deleted,omitempty"`
	Description          *string                        `json:"description,omitempty"`
	ExternalCode         *string                        `json:"externalCode,omitempty"`
	Files                *MetaArray[File]               `json:"files,omitempty"`
	Group                *Group                         `json:"group,omitempty"`
	ID                   *uuid.UUID                     `json:"id,omitempty"`
	Meta                 *Meta                          `json:"meta,omitempty"`
	Moment               *Timestamp                     `json:"moment,omitempty"`
	Name                 *string                        `json:"name,omitempty"`
	AccountID            *uuid.UUID                     `json:"accountId,omitempty"`
	Contract             *NullValue[Contract]           `json:"contract,omitempty"`
	Agent                *Counterparty                  `json:"agent,omitempty"`
	Organization         *Organization                  `json:"organization,omitempty"`
	PaymentPlannedMoment *Timestamp                     `json:"paymentPlannedMoment,omitempty"`
	Positions            *Positions[InvoiceOutPosition] `json:"positions,omitempty"`
	Printed              *bool                          `json:"printed,omitempty"`
	Project              *NullValue[Project]            `json:"project,omitempty"`
	Published            *bool                          `json:"published,omitempty"`
	Rate                 *Rate                          `json:"rate,omitempty"`
	Shared               *bool                          `json:"shared,omitempty"`
	ShippedSum           *float64                       `json:"shippedSum,omitempty"`
	State                *NullValue[State]              `json:"state,omitempty"`
	Store                *NullValue[Store]              `json:"store,omitempty"`
	Sum                  *float64                       `json:"sum,omitempty"`
	SyncID               *uuid.UUID                     `json:"syncId,omitempty"`
	Updated              *Timestamp                     `json:"updated,omitempty"`
	Owner                *Employee                      `json:"owner,omitempty"`
	VatIncluded          *bool                          `json:"vatIncluded,omitempty"`
	VatSum               *float64                       `json:"vatSum,omitempty"`
	CustomerOrder        *CustomerOrder                 `json:"customerOrder,omitempty"`
	Payments             Slice[Payment]                 `json:"payments,omitempty"`
	Attributes           Slice[AttributeValue]          `json:"attributes,omitempty"`
}

// Clean возвращает сущность с единственным заполненным полем Meta
func (invoiceOut InvoiceOut) Clean() *InvoiceOut {
	return &InvoiceOut{Meta: invoiceOut.Meta}
}

// AsOperation возвращает объект Operation c полем Meta сущности
func (invoiceOut InvoiceOut) AsOperation() *Operation {
	return &Operation{Meta: invoiceOut.GetMeta()}
}

func (invoiceOut InvoiceOut) GetPayedSum() float64 {
	return Deref(invoiceOut.PayedSum)
}

func (invoiceOut InvoiceOut) GetVatEnabled() bool {
	return Deref(invoiceOut.VatEnabled)
}

func (invoiceOut InvoiceOut) GetAgentAccount() AgentAccount {
	return Deref(invoiceOut.AgentAccount)
}

func (invoiceOut InvoiceOut) GetApplicable() bool {
	return Deref(invoiceOut.Applicable)
}

func (invoiceOut InvoiceOut) GetDemands() Slice[Demand] {
	return invoiceOut.Demands
}

func (invoiceOut InvoiceOut) GetCode() string {
	return Deref(invoiceOut.Code)
}

func (invoiceOut InvoiceOut) GetOrganizationAccount() AgentAccount {
	return Deref(invoiceOut.OrganizationAccount)
}

func (invoiceOut InvoiceOut) GetCreated() Timestamp {
	return Deref(invoiceOut.Created)
}

func (invoiceOut InvoiceOut) GetDeleted() Timestamp {
	return Deref(invoiceOut.Deleted)
}

func (invoiceOut InvoiceOut) GetDescription() string {
	return Deref(invoiceOut.Description)
}

func (invoiceOut InvoiceOut) GetExternalCode() string {
	return Deref(invoiceOut.ExternalCode)
}

func (invoiceOut InvoiceOut) GetFiles() MetaArray[File] {
	return Deref(invoiceOut.Files)
}

func (invoiceOut InvoiceOut) GetGroup() Group {
	return Deref(invoiceOut.Group)
}

func (invoiceOut InvoiceOut) GetID() uuid.UUID {
	return Deref(invoiceOut.ID)
}

func (invoiceOut InvoiceOut) GetMeta() Meta {
	return Deref(invoiceOut.Meta)
}

func (invoiceOut InvoiceOut) GetMoment() Timestamp {
	return Deref(invoiceOut.Moment)
}

func (invoiceOut InvoiceOut) GetName() string {
	return Deref(invoiceOut.Name)
}

func (invoiceOut InvoiceOut) GetAccountID() uuid.UUID {
	return Deref(invoiceOut.AccountID)
}

func (invoiceOut InvoiceOut) GetContract() Contract {
	return invoiceOut.Contract.Get()
}

func (invoiceOut InvoiceOut) GetAgent() Counterparty {
	return Deref(invoiceOut.Agent)
}

func (invoiceOut InvoiceOut) GetOrganization() Organization {
	return Deref(invoiceOut.Organization)
}

func (invoiceOut InvoiceOut) GetPaymentPlannedMoment() Timestamp {
	return Deref(invoiceOut.PaymentPlannedMoment)
}

func (invoiceOut InvoiceOut) GetPositions() Positions[InvoiceOutPosition] {
	return Deref(invoiceOut.Positions)
}

func (invoiceOut InvoiceOut) GetPrinted() bool {
	return Deref(invoiceOut.Printed)
}

func (invoiceOut InvoiceOut) GetProject() Project {
	return invoiceOut.Project.Get()
}

func (invoiceOut InvoiceOut) GetPublished() bool {
	return Deref(invoiceOut.Published)
}

func (invoiceOut InvoiceOut) GetRate() Rate {
	return Deref(invoiceOut.Rate)
}

func (invoiceOut InvoiceOut) GetShared() bool {
	return Deref(invoiceOut.Shared)
}

func (invoiceOut InvoiceOut) GetShippedSum() float64 {
	return Deref(invoiceOut.ShippedSum)
}

func (invoiceOut InvoiceOut) GetState() State {
	return invoiceOut.State.Get()
}

func (invoiceOut InvoiceOut) GetStore() Store {
	return invoiceOut.Store.Get()
}

func (invoiceOut InvoiceOut) GetSum() float64 {
	return Deref(invoiceOut.Sum)
}

func (invoiceOut InvoiceOut) GetSyncID() uuid.UUID {
	return Deref(invoiceOut.SyncID)
}

func (invoiceOut InvoiceOut) GetUpdated() Timestamp {
	return Deref(invoiceOut.Updated)
}

func (invoiceOut InvoiceOut) GetOwner() Employee {
	return Deref(invoiceOut.Owner)
}

func (invoiceOut InvoiceOut) GetVatIncluded() bool {
	return Deref(invoiceOut.VatIncluded)
}

func (invoiceOut InvoiceOut) GetVatSum() float64 {
	return Deref(invoiceOut.VatSum)
}

func (invoiceOut InvoiceOut) GetCustomerOrder() CustomerOrder {
	return Deref(invoiceOut.CustomerOrder)
}

func (invoiceOut InvoiceOut) GetPayments() Slice[Payment] {
	return invoiceOut.Payments
}

func (invoiceOut InvoiceOut) GetAttributes() Slice[AttributeValue] {
	return invoiceOut.Attributes
}

func (invoiceOut *InvoiceOut) SetVatEnabled(vatEnabled bool) *InvoiceOut {
	invoiceOut.VatEnabled = &vatEnabled
	return invoiceOut
}

func (invoiceOut *InvoiceOut) SetAgentAccount(agentAccount *AgentAccount) *InvoiceOut {
	invoiceOut.AgentAccount = agentAccount.Clean()
	return invoiceOut
}

func (invoiceOut *InvoiceOut) SetApplicable(applicable bool) *InvoiceOut {
	invoiceOut.Applicable = &applicable
	return invoiceOut
}

func (invoiceOut *InvoiceOut) SetDemands(demands Slice[Demand]) *InvoiceOut {
	invoiceOut.Demands = demands
	return invoiceOut
}

func (invoiceOut *InvoiceOut) SetCode(code string) *InvoiceOut {
	invoiceOut.Code = &code
	return invoiceOut
}

func (invoiceOut *InvoiceOut) SetOrganizationAccount(organizationAccount *AgentAccount) *InvoiceOut {
	invoiceOut.OrganizationAccount = organizationAccount.Clean()
	return invoiceOut
}

func (invoiceOut *InvoiceOut) SetDescription(description string) *InvoiceOut {
	invoiceOut.Description = &description
	return invoiceOut
}

func (invoiceOut *InvoiceOut) SetExternalCode(externalCode string) *InvoiceOut {
	invoiceOut.ExternalCode = &externalCode
	return invoiceOut
}

func (invoiceOut *InvoiceOut) SetFiles(files Slice[File]) *InvoiceOut {
	invoiceOut.Files = NewMetaArrayRows(files)
	return invoiceOut
}

func (invoiceOut *InvoiceOut) SetGroup(group *Group) *InvoiceOut {
	invoiceOut.Group = group.Clean()
	return invoiceOut
}

func (invoiceOut *InvoiceOut) SetMeta(meta *Meta) *InvoiceOut {
	invoiceOut.Meta = meta
	return invoiceOut
}

func (invoiceOut *InvoiceOut) SetMoment(moment *Timestamp) *InvoiceOut {
	invoiceOut.Moment = moment
	return invoiceOut
}

func (invoiceOut *InvoiceOut) SetName(name string) *InvoiceOut {
	invoiceOut.Name = &name
	return invoiceOut
}

func (invoiceOut *InvoiceOut) SetContract(contract *Contract) *InvoiceOut {
	invoiceOut.Contract = NewNullValueWith(contract.Clean())
	return invoiceOut
}

func (invoiceOut *InvoiceOut) SetNullContract() *InvoiceOut {
	invoiceOut.Contract = NewNullValue[Contract]()
	return invoiceOut
}

func (invoiceOut *InvoiceOut) SetAgent(agent *Counterparty) *InvoiceOut {
	invoiceOut.Agent = agent.Clean()
	return invoiceOut
}

func (invoiceOut *InvoiceOut) SetOrganization(organization *Organization) *InvoiceOut {
	invoiceOut.Organization = organization.Clean()
	return invoiceOut
}

func (invoiceOut *InvoiceOut) SetPaymentPlannedMoment(paymentPlannedMoment *Timestamp) *InvoiceOut {
	invoiceOut.PaymentPlannedMoment = paymentPlannedMoment
	return invoiceOut
}

func (invoiceOut *InvoiceOut) SetPositions(positions *Positions[InvoiceOutPosition]) *InvoiceOut {
	invoiceOut.Positions = positions
	return invoiceOut
}

func (invoiceOut *InvoiceOut) SetProject(project *Project) *InvoiceOut {
	invoiceOut.Project = NewNullValueWith(project.Clean())
	return invoiceOut
}

func (invoiceOut *InvoiceOut) SetNullProject() *InvoiceOut {
	invoiceOut.Project = NewNullValue[Project]()
	return invoiceOut
}

func (invoiceOut *InvoiceOut) SetRate(rate *Rate) *InvoiceOut {
	invoiceOut.Rate = rate
	return invoiceOut
}

func (invoiceOut *InvoiceOut) SetShared(shared bool) *InvoiceOut {
	invoiceOut.Shared = &shared
	return invoiceOut
}

func (invoiceOut *InvoiceOut) SetState(state *State) *InvoiceOut {
	invoiceOut.State = NewNullValueWith(state.Clean())
	return invoiceOut
}

func (invoiceOut *InvoiceOut) SetNullState() *InvoiceOut {
	invoiceOut.State = NewNullValue[State]()
	return invoiceOut
}

func (invoiceOut *InvoiceOut) SetStore(store *Store) *InvoiceOut {
	invoiceOut.Store = NewNullValueWith(store.Clean())
	return invoiceOut
}

func (invoiceOut *InvoiceOut) SetNullStore() *InvoiceOut {
	invoiceOut.Store = NewNullValue[Store]()
	return invoiceOut
}

func (invoiceOut *InvoiceOut) SetSyncID(syncID uuid.UUID) *InvoiceOut {
	invoiceOut.SyncID = &syncID
	return invoiceOut
}

func (invoiceOut *InvoiceOut) SetOwner(owner *Employee) *InvoiceOut {
	invoiceOut.Owner = owner.Clean()
	return invoiceOut
}

func (invoiceOut *InvoiceOut) SetVatIncluded(vatIncluded bool) *InvoiceOut {
	invoiceOut.VatIncluded = &vatIncluded
	return invoiceOut
}

func (invoiceOut *InvoiceOut) SetCustomerOrder(customerOrder *CustomerOrder) *InvoiceOut {
	invoiceOut.CustomerOrder = customerOrder.Clean()
	return invoiceOut
}

func (invoiceOut *InvoiceOut) SetPayments(payments Slice[Payment]) *InvoiceOut {
	invoiceOut.Payments = payments
	return invoiceOut
}

func (invoiceOut *InvoiceOut) SetAttributes(attributes Slice[AttributeValue]) *InvoiceOut {
	invoiceOut.Attributes = attributes
	return invoiceOut
}

func (invoiceOut InvoiceOut) String() string {
	return Stringify(invoiceOut)
}

func (invoiceOut InvoiceOut) MetaType() MetaType {
	return MetaTypeInvoiceOut
}

// InvoiceOutPosition Позиция Счета покупателю.
// Ключевое слово: invoiceposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-schet-pokupatelu-scheta-pokupatelqm-pozicii-scheta-pokupatelu
type InvoiceOutPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	Discount   *float64            `json:"discount,omitempty"`   // Процент скидки или наценки. Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID сущности
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Meta       *Meta               `json:"meta,omitempty"`       // Метаданные
	Price      *float64            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Vat        *int                `json:"vat,omitempty"`        // НДС, которым облагается текущая позиция
	VatEnabled *bool               `json:"vatEnabled,omitempty"` // Включен ли НДС для позиции. С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
	Stock      *Stock              `json:"stock,omitempty"`      // Остатки и себестоимость `?fields=stock&expand=positions`
}

func (invoiceOutPosition InvoiceOutPosition) GetAccountID() uuid.UUID {
	return Deref(invoiceOutPosition.AccountID)
}

func (invoiceOutPosition InvoiceOutPosition) GetAssortment() AssortmentPosition {
	return Deref(invoiceOutPosition.Assortment)
}

func (invoiceOutPosition InvoiceOutPosition) GetDiscount() float64 {
	return Deref(invoiceOutPosition.Discount)
}

func (invoiceOutPosition InvoiceOutPosition) GetID() uuid.UUID {
	return Deref(invoiceOutPosition.ID)
}

func (invoiceOutPosition InvoiceOutPosition) GetPack() Pack {
	return Deref(invoiceOutPosition.Pack)
}

func (invoiceOutPosition InvoiceOutPosition) GetMeta() Meta {
	return Deref(invoiceOutPosition.Meta)
}

func (invoiceOutPosition InvoiceOutPosition) GetPrice() float64 {
	return Deref(invoiceOutPosition.Price)
}

func (invoiceOutPosition InvoiceOutPosition) GetQuantity() float64 {
	return Deref(invoiceOutPosition.Quantity)
}

func (invoiceOutPosition InvoiceOutPosition) GetVat() int {
	return Deref(invoiceOutPosition.Vat)
}

func (invoiceOutPosition InvoiceOutPosition) GetVatEnabled() bool {
	return Deref(invoiceOutPosition.VatEnabled)
}

func (invoiceOutPosition InvoiceOutPosition) GetStock() Stock {
	return Deref(invoiceOutPosition.Stock)
}

func (invoiceOutPosition *InvoiceOutPosition) SetAssortment(assortment AsAssortment) *InvoiceOutPosition {
	invoiceOutPosition.Assortment = assortment.AsAssortment()
	return invoiceOutPosition
}

func (invoiceOutPosition *InvoiceOutPosition) SetDiscount(discount float64) *InvoiceOutPosition {
	invoiceOutPosition.Discount = &discount
	return invoiceOutPosition
}

func (invoiceOutPosition *InvoiceOutPosition) SetPack(pack *Pack) *InvoiceOutPosition {
	invoiceOutPosition.Pack = pack
	return invoiceOutPosition
}

func (invoiceOutPosition *InvoiceOutPosition) SetMeta(meta *Meta) *InvoiceOutPosition {
	invoiceOutPosition.Meta = meta
	return invoiceOutPosition
}

func (invoiceOutPosition *InvoiceOutPosition) SetPrice(price float64) *InvoiceOutPosition {
	invoiceOutPosition.Price = &price
	return invoiceOutPosition
}

func (invoiceOutPosition *InvoiceOutPosition) SetQuantity(quantity float64) *InvoiceOutPosition {
	invoiceOutPosition.Quantity = &quantity
	return invoiceOutPosition
}

func (invoiceOutPosition *InvoiceOutPosition) SetVat(vat int) *InvoiceOutPosition {
	invoiceOutPosition.Vat = &vat
	return invoiceOutPosition
}

func (invoiceOutPosition *InvoiceOutPosition) SetVatEnabled(vatEnabled bool) *InvoiceOutPosition {
	invoiceOutPosition.VatEnabled = &vatEnabled
	return invoiceOutPosition
}

func (invoiceOutPosition InvoiceOutPosition) String() string {
	return Stringify(invoiceOutPosition)
}

func (invoiceOutPosition InvoiceOutPosition) MetaType() MetaType {
	return MetaTypeInvoicePosition
}

// InvoiceOutService
// Сервис для работы со счетами покупателей.
type InvoiceOutService interface {
	GetList(ctx context.Context, params ...*Params) (*List[InvoiceOut], *resty.Response, error)
	Create(ctx context.Context, invoiceOut *InvoiceOut, params ...*Params) (*InvoiceOut, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, invoiceOutList Slice[InvoiceOut], params ...*Params) (*Slice[InvoiceOut], *resty.Response, error)
	DeleteMany(ctx context.Context, invoiceOutList []MetaWrapper) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*InvoiceOut, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, invoiceOut *InvoiceOut, params ...*Params) (*InvoiceOut, *resty.Response, error)
	Template(ctx context.Context) (*InvoiceOut, *resty.Response, error)
	TemplateBased(ctx context.Context, basedOn ...MetaOwner) (*InvoiceOut, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetaAttributesSharedStatesWrapper, *resty.Response, error)
	GetPositions(ctx context.Context, id uuid.UUID, params ...*Params) (*MetaArray[InvoiceOutPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params ...*Params) (*InvoiceOutPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *InvoiceOutPosition, params ...*Params) (*InvoiceOutPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id uuid.UUID, position *InvoiceOutPosition) (*InvoiceOutPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id uuid.UUID, positions Slice[InvoiceOutPosition]) (*Slice[InvoiceOutPosition], *resty.Response, error)
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
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*InvoiceOut, *resty.Response, error)
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
}

func NewInvoiceOutService(client *Client) InvoiceOutService {
	e := NewEndpoint(client, "entity/invoiceout")
	return newMainService[InvoiceOut, InvoiceOutPosition, MetaAttributesSharedStatesWrapper, any](e)
}
