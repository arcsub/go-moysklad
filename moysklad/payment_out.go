package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// PaymentOut Исходящий платеж.
// Ключевое слово: paymentout
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-ishodqschij-platezh
type PaymentOut struct {
	Moment              *Timestamp               `json:"moment,omitempty"`
	Applicable          *bool                    `json:"applicable,omitempty"`
	AgentAccount        *AgentAccount            `json:"agentAccount,omitempty"`
	Name                *string                  `json:"name,omitempty"`
	Operations          Operations               `json:"operations,omitempty"`
	Organization        *Organization            `json:"organization,omitempty"`
	Contract            *NullValue[Contract]     `json:"contract,omitempty"`
	Created             *Timestamp               `json:"created,omitempty"`
	Deleted             *Timestamp               `json:"deleted,omitempty"`
	Description         *string                  `json:"description,omitempty"`
	ExpenseItem         *ExpenseItem             `json:"expenseItem,omitempty"`
	ExternalCode        *string                  `json:"externalCode,omitempty"`
	Files               *MetaArray[File]         `json:"files,omitempty"`
	Group               *Group                   `json:"group,omitempty"`
	ID                  *uuid.UUID               `json:"id,omitempty"`
	Meta                *Meta                    `json:"meta,omitempty"`
	FactureIn           *FactureIn               `json:"factureIn,omitempty"`
	Agent               *Counterparty            `json:"agent,omitempty"`
	Code                *string                  `json:"code,omitempty"`
	OrganizationAccount *AgentAccount            `json:"organizationAccount,omitempty"`
	Owner               *Employee                `json:"owner,omitempty"`
	PaymentPurpose      *string                  `json:"paymentPurpose,omitempty"`
	Printed             *bool                    `json:"printed,omitempty"`
	Project             *NullValue[Project]      `json:"project,omitempty"`
	Published           *bool                    `json:"published,omitempty"`
	Rate                *Rate                    `json:"rate,omitempty"`
	SalesChannel        *NullValue[SalesChannel] `json:"salesChannel,omitempty"`
	Shared              *bool                    `json:"shared,omitempty"`
	State               *NullValue[State]        `json:"state,omitempty"`
	Sum                 *float64                 `json:"sum,omitempty"`
	SyncID              *uuid.UUID               `json:"syncId,omitempty"`
	Updated             *Timestamp               `json:"updated,omitempty"`
	VatSum              *float64                 `json:"vatSum,omitempty"`
	AccountID           *uuid.UUID               `json:"accountId,omitempty"`
	Attributes          Slice[AttributeValue]    `json:"attributes,omitempty"`
}

// Clean возвращает сущность с единственным заполненным полем Meta
func (paymentOut PaymentOut) Clean() *PaymentOut {
	return &PaymentOut{Meta: paymentOut.Meta}
}

func (paymentOut PaymentOut) GetMoment() Timestamp {
	return Deref(paymentOut.Moment)
}

func (paymentOut PaymentOut) GetApplicable() bool {
	return Deref(paymentOut.Applicable)
}

func (paymentOut PaymentOut) GetAgentAccount() AgentAccount {
	return Deref(paymentOut.AgentAccount)
}

func (paymentOut PaymentOut) GetName() string {
	return Deref(paymentOut.Name)
}

func (paymentOut PaymentOut) GetOperations() Operations {
	return paymentOut.Operations
}

func (paymentOut PaymentOut) GetOrganization() Organization {
	return Deref(paymentOut.Organization)
}

func (paymentOut PaymentOut) GetContract() Contract {
	return paymentOut.Contract.Get()
}

func (paymentOut PaymentOut) GetCreated() Timestamp {
	return Deref(paymentOut.Created)
}

func (paymentOut PaymentOut) GetDeleted() Timestamp {
	return Deref(paymentOut.Deleted)
}

func (paymentOut PaymentOut) GetDescription() string {
	return Deref(paymentOut.Description)
}

func (paymentOut PaymentOut) GetExpenseItem() ExpenseItem {
	return Deref(paymentOut.ExpenseItem)
}

func (paymentOut PaymentOut) GetExternalCode() string {
	return Deref(paymentOut.ExternalCode)
}

func (paymentOut PaymentOut) GetFiles() MetaArray[File] {
	return Deref(paymentOut.Files)
}

func (paymentOut PaymentOut) GetGroup() Group {
	return Deref(paymentOut.Group)
}

func (paymentOut PaymentOut) GetID() uuid.UUID {
	return Deref(paymentOut.ID)
}

func (paymentOut PaymentOut) GetMeta() Meta {
	return Deref(paymentOut.Meta)
}

func (paymentOut PaymentOut) GetFactureIn() FactureIn {
	return Deref(paymentOut.FactureIn)
}

func (paymentOut PaymentOut) GetAgent() Counterparty {
	return Deref(paymentOut.Agent)
}

func (paymentOut PaymentOut) GetCode() string {
	return Deref(paymentOut.Code)
}

func (paymentOut PaymentOut) GetOrganizationAccount() AgentAccount {
	return Deref(paymentOut.OrganizationAccount)
}

func (paymentOut PaymentOut) GetOwner() Employee {
	return Deref(paymentOut.Owner)
}

func (paymentOut PaymentOut) GetPaymentPurpose() string {
	return Deref(paymentOut.PaymentPurpose)
}

func (paymentOut PaymentOut) GetPrinted() bool {
	return Deref(paymentOut.Printed)
}

func (paymentOut PaymentOut) GetProject() Project {
	return paymentOut.Project.Get()
}

func (paymentOut PaymentOut) GetPublished() bool {
	return Deref(paymentOut.Published)
}

func (paymentOut PaymentOut) GetRate() Rate {
	return Deref(paymentOut.Rate)
}

func (paymentOut PaymentOut) GetSalesChannel() SalesChannel {
	return paymentOut.SalesChannel.Get()
}

func (paymentOut PaymentOut) GetShared() bool {
	return Deref(paymentOut.Shared)
}

func (paymentOut PaymentOut) GetState() State {
	return paymentOut.State.Get()
}

func (paymentOut PaymentOut) GetSum() float64 {
	return Deref(paymentOut.Sum)
}

func (paymentOut PaymentOut) GetSyncID() uuid.UUID {
	return Deref(paymentOut.SyncID)
}

func (paymentOut PaymentOut) GetUpdated() Timestamp {
	return Deref(paymentOut.Updated)
}

func (paymentOut PaymentOut) GetVatSum() float64 {
	return Deref(paymentOut.VatSum)
}

func (paymentOut PaymentOut) GetAccountID() uuid.UUID {
	return Deref(paymentOut.AccountID)
}

func (paymentOut PaymentOut) GetAttributes() Slice[AttributeValue] {
	return paymentOut.Attributes
}

func (paymentOut *PaymentOut) SetMoment(moment *Timestamp) *PaymentOut {
	paymentOut.Moment = moment
	return paymentOut
}

func (paymentOut *PaymentOut) SetApplicable(applicable bool) *PaymentOut {
	paymentOut.Applicable = &applicable
	return paymentOut
}

func (paymentOut *PaymentOut) SetAgentAccount(agentAccount *AgentAccount) *PaymentOut {
	paymentOut.AgentAccount = agentAccount.Clean()
	return paymentOut
}

func (paymentOut *PaymentOut) SetName(name string) *PaymentOut {
	paymentOut.Name = &name
	return paymentOut
}

func (paymentOut *PaymentOut) SetOperations(operations Operations) *PaymentOut {
	paymentOut.Operations = operations
	return paymentOut
}

func (paymentOut *PaymentOut) SetOrganization(organization *Organization) *PaymentOut {
	paymentOut.Organization = organization.Clean()
	return paymentOut
}

func (paymentOut *PaymentOut) SetContract(contract *Contract) *PaymentOut {
	paymentOut.Contract = NewNullValueWith(contract.Clean())
	return paymentOut
}

func (paymentOut *PaymentOut) SetNullContract() *PaymentOut {
	paymentOut.Contract = NewNullValue[Contract]()
	return paymentOut
}

func (paymentOut *PaymentOut) SetDescription(description string) *PaymentOut {
	paymentOut.Description = &description
	return paymentOut
}

func (paymentOut *PaymentOut) SetExpenseItem(expenseItem *ExpenseItem) *PaymentOut {
	paymentOut.ExpenseItem = expenseItem.Clean()
	return paymentOut
}

func (paymentOut *PaymentOut) SetExternalCode(externalCode string) *PaymentOut {
	paymentOut.ExternalCode = &externalCode
	return paymentOut
}

func (paymentOut *PaymentOut) SetFiles(files Slice[File]) *PaymentOut {
	paymentOut.Files = NewMetaArrayRows(files)
	return paymentOut
}

func (paymentOut *PaymentOut) SetGroup(group *Group) *PaymentOut {
	paymentOut.Group = group.Clean()
	return paymentOut
}

func (paymentOut *PaymentOut) SetMeta(meta *Meta) *PaymentOut {
	paymentOut.Meta = meta
	return paymentOut
}

func (paymentOut *PaymentOut) SetFactureIn(factureIn *FactureIn) *PaymentOut {
	paymentOut.FactureIn = factureIn.Clean()
	return paymentOut
}

func (paymentOut *PaymentOut) SetAgent(agent *Counterparty) *PaymentOut {
	paymentOut.Agent = agent.Clean()
	return paymentOut
}

func (paymentOut *PaymentOut) SetCode(code string) *PaymentOut {
	paymentOut.Code = &code
	return paymentOut
}

func (paymentOut *PaymentOut) SetOrganizationAccount(organizationAccount *AgentAccount) *PaymentOut {
	paymentOut.OrganizationAccount = organizationAccount.Clean()
	return paymentOut
}

func (paymentOut *PaymentOut) SetOwner(owner *Employee) *PaymentOut {
	paymentOut.Owner = owner.Clean()
	return paymentOut
}

func (paymentOut *PaymentOut) SetPaymentPurpose(paymentPurpose string) *PaymentOut {
	paymentOut.PaymentPurpose = &paymentPurpose
	return paymentOut
}

func (paymentOut *PaymentOut) SetProject(project *Project) *PaymentOut {
	paymentOut.Project = NewNullValueWith(project.Clean())
	return paymentOut
}

func (paymentOut *PaymentOut) SetNullProject() *PaymentOut {
	paymentOut.Project = NewNullValue[Project]()
	return paymentOut
}

func (paymentOut *PaymentOut) SetRate(rate *Rate) *PaymentOut {
	paymentOut.Rate = rate
	return paymentOut
}

func (paymentOut *PaymentOut) SetSalesChannel(salesChannel *SalesChannel) *PaymentOut {
	paymentOut.SalesChannel = NewNullValueWith(salesChannel.Clean())
	return paymentOut
}

func (paymentOut *PaymentOut) SetNullSalesChannel() *PaymentOut {
	paymentOut.SalesChannel = NewNullValue[SalesChannel]()
	return paymentOut
}

func (paymentOut *PaymentOut) SetShared(shared bool) *PaymentOut {
	paymentOut.Shared = &shared
	return paymentOut
}

func (paymentOut *PaymentOut) SetState(state *State) *PaymentOut {
	paymentOut.State = NewNullValueWith(state.Clean())
	return paymentOut
}

func (paymentOut *PaymentOut) SetNullState() *PaymentOut {
	paymentOut.State = NewNullValue[State]()
	return paymentOut
}

func (paymentOut *PaymentOut) SetSum(sum float64) *PaymentOut {
	paymentOut.Sum = &sum
	return paymentOut
}

func (paymentOut *PaymentOut) SetSyncID(syncID uuid.UUID) *PaymentOut {
	paymentOut.SyncID = &syncID
	return paymentOut
}

func (paymentOut *PaymentOut) SetVatSum(vatSum float64) *PaymentOut {
	paymentOut.VatSum = &vatSum
	return paymentOut
}

func (paymentOut *PaymentOut) SetAttributes(attributes Slice[AttributeValue]) *PaymentOut {
	paymentOut.Attributes = attributes
	return paymentOut
}

func (paymentOut PaymentOut) String() string {
	return Stringify(paymentOut)
}

func (paymentOut PaymentOut) MetaType() MetaType {
	return MetaTypePaymentOut
}

func (paymentOut PaymentOut) AsPayment() *Payment {
	return &Payment{Meta: paymentOut.GetMeta()}
}

// PaymentOutTemplateArg
// Документ: Исходящий платеж (paymentout)
// Основание, на котором он может быть создан:
// - Возврат покупателя (salesreturn)
// - Приемка (supply)
// - Счет поставщика (invoicein)
// - Заказ поставщику (purchaseorder)
// - Выданный отчет комиссионера (commissionreportout)
//type PaymentOutTemplateArg struct {
//	SalesReturn         *MetaWrapper `json:"salesReturn,omitempty"`
//	Supply              *MetaWrapper `json:"supply,omitempty"`
//	InvoiceIn           *MetaWrapper `json:"invoiceIn,omitempty"`
//	PurchaseOrder       *MetaWrapper `json:"purchaseOrder,omitempty"`
//	CommissionReportOut *MetaWrapper `json:"commissionReportOut,omitempty"`
//}

// PaymentOutService
// Сервис для работы с исходящими платежами.
type PaymentOutService interface {
	GetList(ctx context.Context, params ...*Params) (*List[PaymentOut], *resty.Response, error)
	Create(ctx context.Context, paymentOut *PaymentOut, params ...*Params) (*PaymentOut, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, paymentOutList Slice[PaymentOut], params ...*Params) (*Slice[PaymentOut], *resty.Response, error)
	DeleteMany(ctx context.Context, paymentOutList []MetaWrapper) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*PaymentOut, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, paymentOut *PaymentOut, params ...*Params) (*PaymentOut, *resty.Response, error)
	//endpointTemplate[PaymentOut]
	//endpointTemplateBasedOn[PaymentOut, PaymentOutTemplateArg]
	GetMetadata(ctx context.Context) (*MetaAttributesSharedStatesWrapper, *resty.Response, error)
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
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*PaymentOut, *resty.Response, error)
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

func NewPaymentOutService(client *Client) PaymentOutService {
	e := NewEndpoint(client, "entity/paymentout")
	return newMainService[PaymentOut, any, MetaAttributesSharedStatesWrapper, any](e)
}
