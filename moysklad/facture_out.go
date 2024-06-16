package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// FactureOut Счет-фактура выданный.
// Ключевое слово: factureout
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-schet-faktura-wydannyj
type FactureOut struct {
	Organization    *Organization         `json:"organization,omitempty"`
	Deleted         *Timestamp            `json:"deleted,omitempty"`
	Applicable      *bool                 `json:"applicable,omitempty"`
	AccountID       *uuid.UUID            `json:"accountId,omitempty"`
	Code            *string               `json:"code,omitempty"`
	Contract        *NullValue[Contract]  `json:"contract,omitempty"`
	Created         *Timestamp            `json:"created,omitempty"`
	Owner           *Employee             `json:"owner,omitempty"`
	Description     *string               `json:"description,omitempty"`
	ExternalCode    *string               `json:"externalCode,omitempty"`
	Files           *MetaArray[File]      `json:"files,omitempty"`
	Group           *Group                `json:"group,omitempty"`
	ID              *uuid.UUID            `json:"id,omitempty"`
	Printed         *bool                 `json:"printed,omitempty"`
	Moment          *Timestamp            `json:"moment,omitempty"`
	Name            *string               `json:"name,omitempty"`
	PaymentDate     *Timestamp            `json:"paymentDate,omitempty"`
	Agent           *Counterparty         `json:"agent,omitempty"`
	Meta            *Meta                 `json:"meta,omitempty"`
	Published       *bool                 `json:"published,omitempty"`
	Rate            *Rate                 `json:"rate,omitempty"`
	Shared          *bool                 `json:"shared,omitempty"`
	State           *NullValue[State]     `json:"state,omitempty"`
	StateContractID *string               `json:"stateContractId,omitempty"`
	Sum             *float64              `json:"sum,omitempty"`
	SyncID          *uuid.UUID            `json:"syncId,omitempty"`
	Updated         *Timestamp            `json:"updated,omitempty"`
	Demands         Slice[Demand]         `json:"demands,omitempty"`
	Payments        Slice[Payment]        `json:"payments,omitempty"`
	Returns         Slice[PurchaseReturn] `json:"returns,omitempty"`
	Consignee       *Counterparty         `json:"consignee,omitempty"`
	PaymentNumber   *string               `json:"paymentNumber,omitempty"`
	Attributes      Slice[AttributeValue] `json:"attributes,omitempty"`
}

// Clean возвращает сущность с единственным заполненным полем Meta
func (factureOut FactureOut) Clean() *FactureOut {
	return &FactureOut{Meta: factureOut.Meta}
}

func (factureOut FactureOut) GetOrganization() Organization {
	return Deref(factureOut.Organization)
}

func (factureOut FactureOut) GetDeleted() Timestamp {
	return Deref(factureOut.Deleted)
}

func (factureOut FactureOut) GetApplicable() bool {
	return Deref(factureOut.Applicable)
}

func (factureOut FactureOut) GetAccountID() uuid.UUID {
	return Deref(factureOut.AccountID)
}

func (factureOut FactureOut) GetCode() string {
	return Deref(factureOut.Code)
}

func (factureOut FactureOut) GetContract() Contract {
	return factureOut.Contract.Get()
}

func (factureOut FactureOut) GetCreated() Timestamp {
	return Deref(factureOut.Created)
}

func (factureOut FactureOut) GetOwner() Employee {
	return Deref(factureOut.Owner)
}

func (factureOut FactureOut) GetDescription() string {
	return Deref(factureOut.Description)
}

func (factureOut FactureOut) GetExternalCode() string {
	return Deref(factureOut.ExternalCode)
}

func (factureOut FactureOut) GetFiles() MetaArray[File] {
	return Deref(factureOut.Files)
}

func (factureOut FactureOut) GetGroup() Group {
	return Deref(factureOut.Group)
}

func (factureOut FactureOut) GetID() uuid.UUID {
	return Deref(factureOut.ID)
}

func (factureOut FactureOut) GetPrinted() bool {
	return Deref(factureOut.Printed)
}

func (factureOut FactureOut) GetMoment() Timestamp {
	return Deref(factureOut.Moment)
}

func (factureOut FactureOut) GetName() string {
	return Deref(factureOut.Name)
}

func (factureOut FactureOut) GetPaymentDate() Timestamp {
	return Deref(factureOut.PaymentDate)
}

func (factureOut FactureOut) GetAgent() Counterparty {
	return Deref(factureOut.Agent)
}

func (factureOut FactureOut) GetMeta() Meta {
	return Deref(factureOut.Meta)
}

func (factureOut FactureOut) GetPublished() bool {
	return Deref(factureOut.Published)
}

func (factureOut FactureOut) GetRate() Rate {
	return Deref(factureOut.Rate)
}

func (factureOut FactureOut) GetShared() bool {
	return Deref(factureOut.Shared)
}

func (factureOut FactureOut) GetState() State {
	return factureOut.State.Get()
}

func (factureOut FactureOut) GetStateContractID() string {
	return Deref(factureOut.StateContractID)
}

func (factureOut FactureOut) GetSum() float64 {
	return Deref(factureOut.Sum)
}

func (factureOut FactureOut) GetSyncID() uuid.UUID {
	return Deref(factureOut.SyncID)
}

func (factureOut FactureOut) GetUpdated() Timestamp {
	return Deref(factureOut.Updated)
}

func (factureOut FactureOut) GetDemands() Slice[Demand] {
	return factureOut.Demands
}

func (factureOut FactureOut) GetPayments() Slice[Payment] {
	return factureOut.Payments
}

func (factureOut FactureOut) GetReturns() Slice[PurchaseReturn] {
	return factureOut.Returns
}

func (factureOut FactureOut) GetConsignee() Counterparty {
	return Deref(factureOut.Consignee)
}

func (factureOut FactureOut) GetPaymentNumber() string {
	return Deref(factureOut.PaymentNumber)
}

func (factureOut FactureOut) GetAttributes() Slice[AttributeValue] {
	return factureOut.Attributes
}

func (factureOut *FactureOut) SetOrganization(organization *Organization) *FactureOut {
	factureOut.Organization = organization.Clean()
	return factureOut
}

func (factureOut *FactureOut) SetApplicable(applicable bool) *FactureOut {
	factureOut.Applicable = &applicable
	return factureOut
}

func (factureOut *FactureOut) SetCode(code string) *FactureOut {
	factureOut.Code = &code
	return factureOut
}

func (factureOut *FactureOut) SetContract(contract *Contract) *FactureOut {
	factureOut.Contract = NewNullValueWith(contract.Clean())
	return factureOut
}

func (factureOut *FactureOut) SetNullContract() *FactureOut {
	factureOut.Contract = NewNullValue[Contract]()
	return factureOut
}

func (factureOut *FactureOut) SetOwner(owner *Employee) *FactureOut {
	factureOut.Owner = owner.Clean()
	return factureOut
}

func (factureOut *FactureOut) SetDescription(description string) *FactureOut {
	factureOut.Description = &description
	return factureOut
}

func (factureOut *FactureOut) SetExternalCode(externalCode string) *FactureOut {
	factureOut.ExternalCode = &externalCode
	return factureOut
}

func (factureOut *FactureOut) SetFiles(files Slice[File]) *FactureOut {
	factureOut.Files = NewMetaArrayRows(files)
	return factureOut
}

func (factureOut *FactureOut) SetGroup(group *Group) *FactureOut {
	factureOut.Group = group.Clean()
	return factureOut
}

func (factureOut *FactureOut) SetMoment(moment *Timestamp) *FactureOut {
	factureOut.Moment = moment
	return factureOut
}

func (factureOut *FactureOut) SetName(name string) *FactureOut {
	factureOut.Name = &name
	return factureOut
}

func (factureOut *FactureOut) SetPaymentDate(paymentDate *Timestamp) *FactureOut {
	factureOut.PaymentDate = paymentDate
	return factureOut
}

func (factureOut *FactureOut) SetAgent(agent *Counterparty) *FactureOut {
	factureOut.Agent = agent.Clean()
	return factureOut
}

func (factureOut *FactureOut) SetMeta(meta *Meta) *FactureOut {
	factureOut.Meta = meta
	return factureOut
}

func (factureOut *FactureOut) SetRate(rate *Rate) *FactureOut {
	factureOut.Rate = rate
	return factureOut
}

func (factureOut *FactureOut) SetShared(shared bool) *FactureOut {
	factureOut.Shared = &shared
	return factureOut
}

func (factureOut *FactureOut) SetState(state *State) *FactureOut {
	factureOut.State = NewNullValueWith(state.Clean())
	return factureOut
}

func (factureOut *FactureOut) SetNullState() *FactureOut {
	factureOut.State = NewNullValue[State]()
	return factureOut
}

func (factureOut *FactureOut) SetStateContractID(stateContractID string) *FactureOut {
	factureOut.StateContractID = &stateContractID
	return factureOut
}

func (factureOut *FactureOut) SetSyncID(syncID uuid.UUID) *FactureOut {
	factureOut.SyncID = &syncID
	return factureOut
}

func (factureOut *FactureOut) SetDemands(demands Slice[Demand]) *FactureOut {
	factureOut.Demands = demands
	return factureOut
}

func (factureOut *FactureOut) SetPayments(payments Slice[Payment]) *FactureOut {
	factureOut.Payments = payments
	return factureOut
}

func (factureOut *FactureOut) SetReturns(returns Slice[PurchaseReturn]) *FactureOut {
	factureOut.Returns = returns
	return factureOut
}

func (factureOut *FactureOut) SetConsignee(consignee *Counterparty) *FactureOut {
	factureOut.Consignee = consignee.Clean()
	return factureOut
}

func (factureOut *FactureOut) SetPaymentNumber(paymentNumber string) *FactureOut {
	factureOut.PaymentNumber = &paymentNumber
	return factureOut
}

func (factureOut *FactureOut) SetAttributes(attributes Slice[AttributeValue]) *FactureOut {
	factureOut.Attributes = attributes
	return factureOut
}

func (factureOut FactureOut) String() string {
	return Stringify(factureOut)
}

func (factureOut FactureOut) MetaType() MetaType {
	return MetaTypeFactureOut
}

// FactureOutService
// Сервис для работы со счетами-фактурами выданными.
type FactureOutService interface {
	GetList(ctx context.Context, params ...*Params) (*List[FactureOut], *resty.Response, error)
	Create(ctx context.Context, factureOut *FactureOut, params ...*Params) (*FactureOut, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, factureOutList Slice[FactureOut], params ...*Params) (*Slice[FactureOut], *resty.Response, error)
	DeleteMany(ctx context.Context, factureOutList []MetaWrapper) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*FactureOut, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, factureOut *FactureOut, params ...*Params) (*FactureOut, *resty.Response, error)
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
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*FactureOut, *resty.Response, error)
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

func NewFactureOutService(client *Client) FactureOutService {
	e := NewEndpoint(client, "entity/factureout")
	return newMainService[FactureOut, any, MetaAttributesSharedStatesWrapper, any](e)
}
