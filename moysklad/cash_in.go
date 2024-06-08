package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// CashIn Приходный ордер.
// Ключевое слово: cashin
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-prihodnyj-order
type CashIn struct {
	Organization   *Organization `json:"organization,omitempty"`
	VatSum         *float64      `json:"vatSum,omitempty"`
	Applicable     *bool         `json:"applicable,omitempty"`
	Moment         *Timestamp    `json:"moment,omitempty"`
	Code           *string       `json:"code,omitempty"`
	Contract       *Contract     `json:"contract,omitempty"`
	AccountID      *uuid.UUID    `json:"accountId,omitempty"`
	Deleted        *Timestamp    `json:"deleted,omitempty"`
	Description    *string       `json:"description,omitempty"`
	ExternalCode   *string       `json:"externalCode,omitempty"`
	Files          *Files        `json:"files,omitempty"`
	Group          *Group        `json:"group,omitempty"`
	ID             *uuid.UUID    `json:"id,omitempty"`
	Meta           *Meta         `json:"meta,omitempty"`
	Operations     Operations    `json:"operations,omitempty"`
	Agent          *Counterparty `json:"agent,omitempty"`
	Created        *Timestamp    `json:"created,omitempty"`
	Owner          *Employee     `json:"owner,omitempty"`
	PaymentPurpose *string       `json:"paymentPurpose,omitempty"`
	Printed        *bool         `json:"printed,omitempty"`
	Project        *Project      `json:"project,omitempty"`
	Published      *bool         `json:"published,omitempty"`
	Rate           *Rate         `json:"rate,omitempty"`
	SalesChannel   *SalesChannel `json:"salesChannel,omitempty"`
	Shared         *bool         `json:"shared,omitempty"`
	State          *State        `json:"state,omitempty"`
	Sum            *float64      `json:"sum,omitempty"`
	SyncID         *uuid.UUID    `json:"syncId,omitempty"`
	Updated        *Timestamp    `json:"updated,omitempty"`
	Name           *string       `json:"name,omitempty"`
	FactureIn      *FactureIn    `json:"factureIn,omitempty"`
	Attributes     Attributes    `json:"attributes,omitempty"`
}

func (cashIn CashIn) GetOrganization() Organization {
	return Deref(cashIn.Organization)
}

func (cashIn CashIn) GetVatSum() float64 {
	return Deref(cashIn.VatSum)
}

func (cashIn CashIn) GetApplicable() bool {
	return Deref(cashIn.Applicable)
}

func (cashIn CashIn) GetMoment() Timestamp {
	return Deref(cashIn.Moment)
}

func (cashIn CashIn) GetCode() string {
	return Deref(cashIn.Code)
}

func (cashIn CashIn) GetContract() Contract {
	return Deref(cashIn.Contract)
}

func (cashIn CashIn) GetAccountID() uuid.UUID {
	return Deref(cashIn.AccountID)
}

func (cashIn CashIn) GetDeleted() Timestamp {
	return Deref(cashIn.Deleted)
}

func (cashIn CashIn) GetDescription() string {
	return Deref(cashIn.Description)
}

func (cashIn CashIn) GetExternalCode() string {
	return Deref(cashIn.ExternalCode)
}

func (cashIn CashIn) GetFiles() Files {
	return Deref(cashIn.Files)
}

func (cashIn CashIn) GetGroup() Group {
	return Deref(cashIn.Group)
}

func (cashIn CashIn) GetID() uuid.UUID {
	return Deref(cashIn.ID)
}

func (cashIn CashIn) GetMeta() Meta {
	return Deref(cashIn.Meta)
}

func (cashIn CashIn) GetOperations() Operations {
	return cashIn.Operations
}

func (cashIn CashIn) GetAgent() Counterparty {
	return Deref(cashIn.Agent)
}

func (cashIn CashIn) GetCreated() Timestamp {
	return Deref(cashIn.Created)
}

func (cashIn CashIn) GetOwner() Employee {
	return Deref(cashIn.Owner)
}

func (cashIn CashIn) GetPaymentPurpose() string {
	return Deref(cashIn.PaymentPurpose)
}

func (cashIn CashIn) GetPrinted() bool {
	return Deref(cashIn.Printed)
}

func (cashIn CashIn) GetProject() Project {
	return Deref(cashIn.Project)
}

func (cashIn CashIn) GetPublished() bool {
	return Deref(cashIn.Published)
}

func (cashIn CashIn) GetRate() Rate {
	return Deref(cashIn.Rate)
}

func (cashIn CashIn) GetSalesChannel() SalesChannel {
	return Deref(cashIn.SalesChannel)
}

func (cashIn CashIn) GetShared() bool {
	return Deref(cashIn.Shared)
}

func (cashIn CashIn) GetState() State {
	return Deref(cashIn.State)
}

func (cashIn CashIn) GetSum() float64 {
	return Deref(cashIn.Sum)
}

func (cashIn CashIn) GetSyncID() uuid.UUID {
	return Deref(cashIn.SyncID)
}

func (cashIn CashIn) GetUpdated() Timestamp {
	return Deref(cashIn.Updated)
}

func (cashIn CashIn) GetName() string {
	return Deref(cashIn.Name)
}

func (cashIn CashIn) GetFactureIn() FactureIn {
	return Deref(cashIn.FactureIn)
}

func (cashIn CashIn) GetAttributes() Attributes {
	return cashIn.Attributes
}

func (cashIn *CashIn) SetOrganization(organization *Organization) *CashIn {
	cashIn.Organization = organization
	return cashIn
}

func (cashIn *CashIn) SetVatSum(vatSum *float64) *CashIn {
	cashIn.VatSum = vatSum
	return cashIn
}

func (cashIn *CashIn) SetApplicable(applicable bool) *CashIn {
	cashIn.Applicable = &applicable
	return cashIn
}

func (cashIn *CashIn) SetMoment(moment *Timestamp) *CashIn {
	cashIn.Moment = moment
	return cashIn
}

func (cashIn *CashIn) SetCode(code string) *CashIn {
	cashIn.Code = &code
	return cashIn
}

func (cashIn *CashIn) SetContract(contract *Contract) *CashIn {
	cashIn.Contract = contract
	return cashIn
}

func (cashIn *CashIn) SetDescription(description string) *CashIn {
	cashIn.Description = &description
	return cashIn
}

func (cashIn *CashIn) SetExternalCode(externalCode string) *CashIn {
	cashIn.ExternalCode = &externalCode
	return cashIn
}

func (cashIn *CashIn) SetFiles(files *Files) *CashIn {
	cashIn.Files = files
	return cashIn
}

func (cashIn *CashIn) SetGroup(group *Group) *CashIn {
	cashIn.Group = group
	return cashIn
}

func (cashIn *CashIn) SetMeta(meta *Meta) *CashIn {
	cashIn.Meta = meta
	return cashIn
}

func (cashIn *CashIn) SetOperations(operations Operations) *CashIn {
	cashIn.Operations = operations
	return cashIn
}

func (cashIn *CashIn) SetAgent(agent *Counterparty) *CashIn {
	cashIn.Agent = agent
	return cashIn
}

func (cashIn *CashIn) SetOwner(owner *Employee) *CashIn {
	cashIn.Owner = owner
	return cashIn
}

func (cashIn *CashIn) SetPaymentPurpose(paymentPurpose string) *CashIn {
	cashIn.PaymentPurpose = &paymentPurpose
	return cashIn
}

func (cashIn *CashIn) SetProject(project *Project) *CashIn {
	cashIn.Project = project
	return cashIn
}

func (cashIn *CashIn) SetRate(rate *Rate) *CashIn {
	cashIn.Rate = rate
	return cashIn
}

func (cashIn *CashIn) SetSalesChannel(salesChannel *SalesChannel) *CashIn {
	cashIn.SalesChannel = salesChannel
	return cashIn
}

func (cashIn *CashIn) SetShared(shared bool) *CashIn {
	cashIn.Shared = &shared
	return cashIn
}

func (cashIn *CashIn) SetState(state *State) *CashIn {
	cashIn.State = state
	return cashIn
}

func (cashIn *CashIn) SetSum(sum *float64) *CashIn {
	cashIn.Sum = sum
	return cashIn
}

func (cashIn *CashIn) SetSyncID(syncID *uuid.UUID) *CashIn {
	cashIn.SyncID = syncID
	return cashIn
}

func (cashIn *CashIn) SetName(name string) *CashIn {
	cashIn.Name = &name
	return cashIn
}

func (cashIn *CashIn) SetFactureIn(factureIn *FactureIn) *CashIn {
	cashIn.FactureIn = factureIn
	return cashIn
}

func (cashIn *CashIn) SetAttributes(attributes Attributes) *CashIn {
	cashIn.Attributes = attributes
	return cashIn
}

func (cashIn CashIn) String() string {
	return Stringify(cashIn)
}

func (cashIn CashIn) MetaType() MetaType {
	return MetaTypeCashIn
}

// CashInTemplateArg
// Документ: Приходный ордер (cashin)
// Основание, на котором он может быть создан:
// - Заказ покупателя (customerorder)
// - Возврат поставщику (purchasereturn)
// - Отгрузка (demand)
// - Счет покупателю (invoiceout)
// - Полученный отчет комиссионера (commissionreportin)
type CashInTemplateArg struct {
	CustomerOrder      *MetaWrapper `json:"customerOrder,omitempty"`
	PurchaseReturn     *MetaWrapper `json:"purchaseReturn,omitempty"`
	Demand             *MetaWrapper `json:"demand,omitempty"`
	InvoiceOut         *MetaWrapper `json:"invoiceOut,omitempty"`
	CommissionReportIn *MetaWrapper `json:"commissionReportIn,omitempty"`
}

// CashInService
// Сервис для работы с приходными ордерами.
type CashInService interface {
	GetList(ctx context.Context, params *Params) (*List[CashIn], *resty.Response, error)
	Create(ctx context.Context, cashIn *CashIn, params *Params) (*CashIn, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, cashInList []*CashIn, params *Params) (*[]CashIn, *resty.Response, error)
	DeleteMany(ctx context.Context, cashInList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetaAttributesSharedStatesWrapper, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id *uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error)
	UpdateAttribute(ctx context.Context, id *uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	//Template(ctx context.Context) (*CashIn, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*CashIn, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, cashIn *CashIn, params *Params) (*CashIn, *resty.Response, error)
	GetPublications(ctx context.Context, id *uuid.UUID) (*MetaArray[Publication], *resty.Response, error)
	GetPublicationByID(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (*Publication, *resty.Response, error)
	Publish(ctx context.Context, id *uuid.UUID, template *Templater) (*Publication, *resty.Response, error)
	DeletePublication(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (bool, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*CashIn, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	MoveToTrash(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewCashInService(client *Client) CashInService {
	e := NewEndpoint(client, "entity/cashin")
	return newMainService[CashIn, any, MetaAttributesSharedStatesWrapper, any](e)
}
