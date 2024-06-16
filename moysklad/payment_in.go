package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// PaymentIn Входящий платеж.
// Ключевое слово: paymentin
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vhodqschij-platezh
type PaymentIn struct {
	Meta                *Meta                    `json:"meta,omitempty"`
	Applicable          *bool                    `json:"applicable,omitempty"`
	AgentAccount        *AgentAccount            `json:"agentAccount,omitempty"`
	Moment              *Timestamp               `json:"moment,omitempty"`
	Operations          Operations               `json:"operations,omitempty"`
	Name                *string                  `json:"name,omitempty"`
	Contract            *NullValue[Contract]     `json:"contract,omitempty"`
	Created             *Timestamp               `json:"created,omitempty"`
	Deleted             *Timestamp               `json:"deleted,omitempty"`
	Description         *string                  `json:"description,omitempty"`
	ExternalCode        *string                  `json:"externalCode,omitempty"`
	Files               *MetaArray[File]         `json:"files,omitempty"`
	Group               *Group                   `json:"group,omitempty"`
	ID                  *uuid.UUID               `json:"id,omitempty"`
	IncomingDate        *Timestamp               `json:"incomingDate,omitempty"`
	IncomingNumber      *string                  `json:"incomingNumber,omitempty"`
	FactureOut          *FactureOut              `json:"factureOut,omitempty"`
	Agent               *Counterparty            `json:"agent,omitempty"`
	Code                *string                  `json:"code,omitempty"`
	Organization        *Organization            `json:"organization,omitempty"`
	OrganizationAccount *AgentAccount            `json:"organizationAccount,omitempty"`
	Owner               *Employee                `json:"owner,omitempty"`
	PaymentPurpose      *string                  `json:"paymentPurpose,omitempty"`
	Printed             *bool                    `json:"printed,omitempty"`
	Project             *NullValue[Project]      `json:"project,omitempty"`
	Published           *bool                    `json:"published,omitempty"`
	Rate                *Rate                    `json:"rate,omitempty"`
	Shared              *bool                    `json:"shared,omitempty"`
	SalesChannel        *NullValue[SalesChannel] `json:"salesChannel,omitempty"`
	State               *NullValue[State]        `json:"state,omitempty"`
	Sum                 *float64                 `json:"sum,omitempty"`
	SyncID              *uuid.UUID               `json:"syncId,omitempty"`
	Updated             *Timestamp               `json:"updated,omitempty"`
	AccountID           *uuid.UUID               `json:"accountId,omitempty"`
	Attributes          Slice[AttributeValue]    `json:"attributes,omitempty"`
}

// Clean возвращает сущность с единственным заполненным полем Meta
func (paymentIn PaymentIn) Clean() *PaymentIn {
	return &PaymentIn{Meta: paymentIn.Meta}
}

func (paymentIn PaymentIn) GetMeta() Meta {
	return Deref(paymentIn.Meta)
}

func (paymentIn PaymentIn) GetApplicable() bool {
	return Deref(paymentIn.Applicable)
}

func (paymentIn PaymentIn) GetAgentAccount() AgentAccount {
	return Deref(paymentIn.AgentAccount)
}

func (paymentIn PaymentIn) GetMoment() Timestamp {
	return Deref(paymentIn.Moment)
}

func (paymentIn PaymentIn) GetOperations() Operations {
	return paymentIn.Operations
}

func (paymentIn PaymentIn) GetName() string {
	return Deref(paymentIn.Name)
}

func (paymentIn PaymentIn) GetContract() Contract {
	return paymentIn.Contract.Get()
}

func (paymentIn PaymentIn) GetCreated() Timestamp {
	return Deref(paymentIn.Created)
}

func (paymentIn PaymentIn) GetDeleted() Timestamp {
	return Deref(paymentIn.Deleted)
}

func (paymentIn PaymentIn) GetDescription() string {
	return Deref(paymentIn.Description)
}

func (paymentIn PaymentIn) GetExternalCode() string {
	return Deref(paymentIn.ExternalCode)
}

func (paymentIn PaymentIn) GetFiles() MetaArray[File] {
	return Deref(paymentIn.Files)
}

func (paymentIn PaymentIn) GetGroup() Group {
	return Deref(paymentIn.Group)
}

func (paymentIn PaymentIn) GetID() uuid.UUID {
	return Deref(paymentIn.ID)
}

func (paymentIn PaymentIn) GetIncomingDate() Timestamp {
	return Deref(paymentIn.IncomingDate)
}

func (paymentIn PaymentIn) GetIncomingNumber() string {
	return Deref(paymentIn.IncomingNumber)
}

func (paymentIn PaymentIn) GetFactureOut() FactureOut {
	return Deref(paymentIn.FactureOut)
}

func (paymentIn PaymentIn) GetAgent() Counterparty {
	return Deref(paymentIn.Agent)
}

func (paymentIn PaymentIn) GetCode() string {
	return Deref(paymentIn.Code)
}

func (paymentIn PaymentIn) GetOrganization() Organization {
	return Deref(paymentIn.Organization)
}

func (paymentIn PaymentIn) GetOrganizationAccount() AgentAccount {
	return Deref(paymentIn.OrganizationAccount)
}

func (paymentIn PaymentIn) GetOwner() Employee {
	return Deref(paymentIn.Owner)
}

func (paymentIn PaymentIn) GetPaymentPurpose() string {
	return Deref(paymentIn.PaymentPurpose)
}

func (paymentIn PaymentIn) GetPrinted() bool {
	return Deref(paymentIn.Printed)
}

func (paymentIn PaymentIn) GetProject() Project {
	return paymentIn.Project.Get()
}

func (paymentIn PaymentIn) GetPublished() bool {
	return Deref(paymentIn.Published)
}

func (paymentIn PaymentIn) GetRate() Rate {
	return Deref(paymentIn.Rate)
}

func (paymentIn PaymentIn) GetShared() bool {
	return Deref(paymentIn.Shared)
}

func (paymentIn PaymentIn) GetSalesChannel() SalesChannel {
	return paymentIn.SalesChannel.Get()
}

func (paymentIn PaymentIn) GetState() State {
	return paymentIn.State.Get()
}

func (paymentIn PaymentIn) GetSum() float64 {
	return Deref(paymentIn.Sum)
}

func (paymentIn PaymentIn) GetSyncID() uuid.UUID {
	return Deref(paymentIn.SyncID)
}

func (paymentIn PaymentIn) GetUpdated() Timestamp {
	return Deref(paymentIn.Updated)
}

func (paymentIn PaymentIn) GetAccountID() uuid.UUID {
	return Deref(paymentIn.AccountID)
}

func (paymentIn PaymentIn) GetAttributes() Slice[AttributeValue] {
	return paymentIn.Attributes
}

func (paymentIn *PaymentIn) SetMeta(meta *Meta) *PaymentIn {
	paymentIn.Meta = meta
	return paymentIn
}

func (paymentIn *PaymentIn) SetApplicable(applicable bool) *PaymentIn {
	paymentIn.Applicable = &applicable
	return paymentIn
}

func (paymentIn *PaymentIn) SetAgentAccount(agentAccount *AgentAccount) *PaymentIn {
	paymentIn.AgentAccount = agentAccount.Clean()
	return paymentIn
}

func (paymentIn *PaymentIn) SetMoment(moment *Timestamp) *PaymentIn {
	paymentIn.Moment = moment
	return paymentIn
}

func (paymentIn *PaymentIn) SetOperations(operations Operations) *PaymentIn {
	paymentIn.Operations = operations
	return paymentIn
}

func (paymentIn *PaymentIn) SetName(name string) *PaymentIn {
	paymentIn.Name = &name
	return paymentIn
}

func (paymentIn *PaymentIn) SetContract(contract *Contract) *PaymentIn {
	paymentIn.Contract = NewNullValueWith(contract.Clean())
	return paymentIn
}

func (paymentIn *PaymentIn) SetNullContract() *PaymentIn {
	paymentIn.Contract = NewNullValue[Contract]()
	return paymentIn
}

func (paymentIn *PaymentIn) SetDescription(description string) *PaymentIn {
	paymentIn.Description = &description
	return paymentIn
}

func (paymentIn *PaymentIn) SetExternalCode(externalCode string) *PaymentIn {
	paymentIn.ExternalCode = &externalCode
	return paymentIn
}

func (paymentIn *PaymentIn) SetFiles(files Slice[File]) *PaymentIn {
	paymentIn.Files = NewMetaArrayRows(files)
	return paymentIn
}

func (paymentIn *PaymentIn) SetGroup(group *Group) *PaymentIn {
	paymentIn.Group = group.Clean()
	return paymentIn
}

func (paymentIn *PaymentIn) SetIncomingDate(incomingDate *Timestamp) *PaymentIn {
	paymentIn.IncomingDate = incomingDate
	return paymentIn
}

func (paymentIn *PaymentIn) SetIncomingNumber(incomingNumber string) *PaymentIn {
	paymentIn.IncomingNumber = &incomingNumber
	return paymentIn
}

func (paymentIn *PaymentIn) SetFactureOut(factureOut *FactureOut) *PaymentIn {
	paymentIn.FactureOut = factureOut.Clean()
	return paymentIn
}

func (paymentIn *PaymentIn) SetAgent(agent *Counterparty) *PaymentIn {
	paymentIn.Agent = agent.Clean()
	return paymentIn
}

func (paymentIn *PaymentIn) SetCode(code string) *PaymentIn {
	paymentIn.Code = &code
	return paymentIn
}

func (paymentIn *PaymentIn) SetOrganization(organization *Organization) *PaymentIn {
	paymentIn.Organization = organization.Clean()
	return paymentIn
}

func (paymentIn *PaymentIn) SetOrganizationAccount(organizationAccount *AgentAccount) *PaymentIn {
	paymentIn.OrganizationAccount = organizationAccount.Clean()
	return paymentIn
}

func (paymentIn *PaymentIn) SetOwner(owner *Employee) *PaymentIn {
	paymentIn.Owner = owner.Clean()
	return paymentIn
}

func (paymentIn *PaymentIn) SetPaymentPurpose(paymentPurpose string) *PaymentIn {
	paymentIn.PaymentPurpose = &paymentPurpose
	return paymentIn
}

func (paymentIn *PaymentIn) SetProject(project *Project) *PaymentIn {
	paymentIn.Project = NewNullValueWith(project.Clean())
	return paymentIn
}

func (paymentIn *PaymentIn) SetNullProject() *PaymentIn {
	paymentIn.Project = NewNullValue[Project]()
	return paymentIn
}

func (paymentIn *PaymentIn) SetRate(rate *Rate) *PaymentIn {
	paymentIn.Rate = rate
	return paymentIn
}

func (paymentIn *PaymentIn) SetShared(shared bool) *PaymentIn {
	paymentIn.Shared = &shared
	return paymentIn
}

func (paymentIn *PaymentIn) SetSalesChannel(salesChannel *SalesChannel) *PaymentIn {
	paymentIn.SalesChannel = NewNullValueWith(salesChannel.Clean())
	return paymentIn
}

func (paymentIn *PaymentIn) SetNullSalesChannel() *PaymentIn {
	paymentIn.SalesChannel = NewNullValue[SalesChannel]()
	return paymentIn
}

func (paymentIn *PaymentIn) SetState(state *State) *PaymentIn {
	paymentIn.State = NewNullValueWith(state.Clean())
	return paymentIn
}

func (paymentIn *PaymentIn) SetNullState() *PaymentIn {
	paymentIn.State = NewNullValue[State]()
	return paymentIn
}

func (paymentIn *PaymentIn) SetSum(sum float64) *PaymentIn {
	paymentIn.Sum = &sum
	return paymentIn
}

func (paymentIn *PaymentIn) SetSyncID(syncID uuid.UUID) *PaymentIn {
	paymentIn.SyncID = &syncID
	return paymentIn
}

func (paymentIn *PaymentIn) SetAttributes(attributes Slice[AttributeValue]) *PaymentIn {
	paymentIn.Attributes = attributes
	return paymentIn
}

func (paymentIn PaymentIn) String() string {
	return Stringify(paymentIn)
}

func (paymentIn PaymentIn) MetaType() MetaType {
	return MetaTypePaymentIn
}

// AsOperation возвращает объект Operation c полем Meta сущности
func (paymentIn PaymentIn) AsOperation() *Operation {
	return &Operation{Meta: paymentIn.GetMeta()}
}

// PaymentInService
// Сервис для работы с входящими платежами.
type PaymentInService interface {
	GetList(ctx context.Context, params ...*Params) (*List[PaymentIn], *resty.Response, error)
	Create(ctx context.Context, paymentIn *PaymentIn, params ...*Params) (*PaymentIn, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, paymentInList Slice[PaymentIn], params ...*Params) (*Slice[PaymentIn], *resty.Response, error)
	DeleteMany(ctx context.Context, paymentInList []MetaWrapper) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*PaymentIn, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, paymentIn *PaymentIn, params ...*Params) (*PaymentIn, *resty.Response, error)
	Template(ctx context.Context) (*PaymentIn, *resty.Response, error)
	TemplateBased(ctx context.Context, basedOn ...MetaOwner) (*PaymentIn, *resty.Response, error)
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
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*PaymentIn, *resty.Response, error)
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

func NewPaymentInService(client *Client) PaymentInService {
	e := NewEndpoint(client, "entity/paymentin")
	return newMainService[PaymentIn, any, MetaAttributesSharedStatesWrapper, any](e)
}
