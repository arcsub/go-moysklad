package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// FactureIn Счет-фактура полученный.
// Ключевое слово: facturein
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-schet-faktura-poluchennyj
type FactureIn struct {
	Moment         *Timestamp           `json:"moment,omitempty"`
	Applicable     *bool                `json:"applicable,omitempty"`
	Name           *string              `json:"name,omitempty"`
	AccountID      *uuid.UUID           `json:"accountId,omitempty"`
	Code           *string              `json:"code,omitempty"`
	Contract       *NullValue[Contract] `json:"contract,omitempty"`
	Created        *Timestamp           `json:"created,omitempty"`
	Deleted        *Timestamp           `json:"deleted,omitempty"`
	Description    *string              `json:"description,omitempty"`
	ExternalCode   *string              `json:"externalCode,omitempty"`
	Files          *MetaArray[File]     `json:"files,omitempty"`
	Group          *Group               `json:"group,omitempty"`
	ID             *uuid.UUID           `json:"id,omitempty"`
	Meta           *Meta                `json:"meta,omitempty"`
	IncomingDate   *Timestamp           `json:"incomingDate,omitempty"`
	Agent          *Counterparty        `json:"agent,omitempty"`
	Organization   *Organization        `json:"organization,omitempty"`
	Owner          *Employee            `json:"owner,omitempty"`
	Printed        *bool                `json:"printed,omitempty"`
	Published      *bool                `json:"published,omitempty"`
	Rate           *NullValue[Rate]     `json:"rate,omitempty"`
	Shared         *bool                `json:"shared,omitempty"`
	State          *NullValue[State]    `json:"state,omitempty"`
	Sum            *float64             `json:"sum,omitempty"`
	SyncID         *uuid.UUID           `json:"syncId,omitempty"`
	Updated        *Timestamp           `json:"updated,omitempty"`
	Supplies       Slice[Supply]        `json:"supplies,omitempty"`
	Payments       Slice[Payment]       `json:"payments,omitempty"`
	IncomingNumber *string              `json:"incomingNumber,omitempty"`
	Attributes     Slice[Attribute]     `json:"attributes,omitempty"`
}

// Clean возвращает сущность с единственным заполненным полем Meta
func (factureIn FactureIn) Clean() *FactureIn {
	return &FactureIn{Meta: factureIn.Meta}
}

// AsOperation возвращает объект Operation c полем Meta сущности
func (factureIn FactureIn) AsOperation() *Operation {
	return &Operation{Meta: factureIn.GetMeta()}
}

// AsTaskOperation реализует интерфейс AsTaskOperationInterface
func (factureIn FactureIn) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: factureIn.Meta}
}

func (factureIn FactureIn) GetMoment() Timestamp {
	return Deref(factureIn.Moment)
}

func (factureIn FactureIn) GetApplicable() bool {
	return Deref(factureIn.Applicable)
}

func (factureIn FactureIn) GetName() string {
	return Deref(factureIn.Name)
}

func (factureIn FactureIn) GetAccountID() uuid.UUID {
	return Deref(factureIn.AccountID)
}

func (factureIn FactureIn) GetCode() string {
	return Deref(factureIn.Code)
}

func (factureIn FactureIn) GetContract() Contract {
	return factureIn.Contract.Get()
}

func (factureIn FactureIn) GetCreated() Timestamp {
	return Deref(factureIn.Created)
}

func (factureIn FactureIn) GetDeleted() Timestamp {
	return Deref(factureIn.Deleted)
}

func (factureIn FactureIn) GetDescription() string {
	return Deref(factureIn.Description)
}

func (factureIn FactureIn) GetExternalCode() string {
	return Deref(factureIn.ExternalCode)
}

func (factureIn FactureIn) GetFiles() MetaArray[File] {
	return Deref(factureIn.Files)
}

func (factureIn FactureIn) GetGroup() Group {
	return Deref(factureIn.Group)
}

func (factureIn FactureIn) GetID() uuid.UUID {
	return Deref(factureIn.ID)
}

func (factureIn FactureIn) GetMeta() Meta {
	return Deref(factureIn.Meta)
}

func (factureIn FactureIn) GetIncomingDate() Timestamp {
	return Deref(factureIn.IncomingDate)
}

func (factureIn FactureIn) GetAgent() Counterparty {
	return Deref(factureIn.Agent)
}

func (factureIn FactureIn) GetOrganization() Organization {
	return Deref(factureIn.Organization)
}

func (factureIn FactureIn) GetOwner() Employee {
	return Deref(factureIn.Owner)
}

func (factureIn FactureIn) GetPrinted() bool {
	return Deref(factureIn.Printed)
}

func (factureIn FactureIn) GetPublished() bool {
	return Deref(factureIn.Published)
}

func (factureIn FactureIn) GetRate() Rate {
	return factureIn.Rate.Get()
}

func (factureIn FactureIn) GetShared() bool {
	return Deref(factureIn.Shared)
}

func (factureIn FactureIn) GetState() State {
	return factureIn.State.Get()
}

func (factureIn FactureIn) GetSum() float64 {
	return Deref(factureIn.Sum)
}

func (factureIn FactureIn) GetSyncID() uuid.UUID {
	return Deref(factureIn.SyncID)
}

func (factureIn FactureIn) GetUpdated() Timestamp {
	return Deref(factureIn.Updated)
}

func (factureIn FactureIn) GetSupplies() Slice[Supply] {
	return factureIn.Supplies
}

func (factureIn FactureIn) GetPayments() Slice[Payment] {
	return factureIn.Payments
}

func (factureIn FactureIn) GetIncomingNumber() string {
	return Deref(factureIn.IncomingNumber)
}

func (factureIn FactureIn) GetAttributes() Slice[Attribute] {
	return factureIn.Attributes
}

func (factureIn *FactureIn) SetMoment(moment *Timestamp) *FactureIn {
	factureIn.Moment = moment
	return factureIn
}

func (factureIn *FactureIn) SetApplicable(applicable bool) *FactureIn {
	factureIn.Applicable = &applicable
	return factureIn
}

func (factureIn *FactureIn) SetName(name string) *FactureIn {
	factureIn.Name = &name
	return factureIn
}

func (factureIn *FactureIn) SetCode(code string) *FactureIn {
	factureIn.Code = &code
	return factureIn
}

func (factureIn *FactureIn) SetContract(contract *Contract) *FactureIn {
	factureIn.Contract = NewNullValueFrom(contract.Clean())
	return factureIn
}

func (factureIn *FactureIn) SetNullContract() *FactureIn {
	factureIn.Contract = NewNullValue[Contract]()
	return factureIn
}

func (factureIn *FactureIn) SetDescription(description string) *FactureIn {
	factureIn.Description = &description
	return factureIn
}

func (factureIn *FactureIn) SetExternalCode(externalCode string) *FactureIn {
	factureIn.ExternalCode = &externalCode
	return factureIn
}

func (factureIn *FactureIn) SetFiles(files ...*File) *FactureIn {
	factureIn.Files = NewMetaArrayFrom(files)
	return factureIn
}

func (factureIn *FactureIn) SetGroup(group *Group) *FactureIn {
	factureIn.Group = group.Clean()
	return factureIn
}

func (factureIn *FactureIn) SetMeta(meta *Meta) *FactureIn {
	factureIn.Meta = meta
	return factureIn
}

func (factureIn *FactureIn) SetIncomingDate(incomingDate *Timestamp) *FactureIn {
	factureIn.IncomingDate = incomingDate
	return factureIn
}

func (factureIn *FactureIn) SetAgent(agent *Counterparty) *FactureIn {
	factureIn.Agent = agent.Clean()
	return factureIn
}

func (factureIn *FactureIn) SetOrganization(organization *Organization) *FactureIn {
	factureIn.Organization = organization.Clean()
	return factureIn
}

func (factureIn *FactureIn) SetOwner(owner *Employee) *FactureIn {
	factureIn.Owner = owner.Clean()
	return factureIn
}

func (factureIn *FactureIn) SetRate(rate *Rate) *FactureIn {
	factureIn.Rate = NewNullValueFrom(rate)
	return factureIn
}

func (factureIn *FactureIn) SetNullRate() *FactureIn {
	factureIn.Rate = NewNullValue[Rate]()
	return factureIn
}

func (factureIn *FactureIn) SetShared(shared bool) *FactureIn {
	factureIn.Shared = &shared
	return factureIn
}

func (factureIn *FactureIn) SetState(state *State) *FactureIn {
	factureIn.State = NewNullValueFrom(state.Clean())
	return factureIn
}

func (factureIn *FactureIn) SetNullState() *FactureIn {
	factureIn.State = NewNullValue[State]()
	return factureIn
}

func (factureIn *FactureIn) SetSyncID(syncID uuid.UUID) *FactureIn {
	factureIn.SyncID = &syncID
	return factureIn
}

func (factureIn *FactureIn) SetSupplies(supplies ...*Supply) *FactureIn {
	factureIn.Supplies = supplies
	return factureIn
}

func (factureIn *FactureIn) SetPayments(payments ...*Payment) *FactureIn {
	factureIn.Payments = payments
	return factureIn
}

func (factureIn *FactureIn) SetIncomingNumber(incomingNumber string) *FactureIn {
	factureIn.IncomingNumber = &incomingNumber
	return factureIn
}

func (factureIn *FactureIn) SetAttributes(attributes ...*Attribute) *FactureIn {
	factureIn.Attributes = attributes
	return factureIn
}

func (factureIn FactureIn) String() string {
	return Stringify(factureIn)
}

// MetaType возвращает тип сущности.
func (FactureIn) MetaType() MetaType {
	return MetaTypeFactureIn
}

// Update shortcut
func (factureIn FactureIn) Update(ctx context.Context, client *Client, params ...*Params) (*FactureIn, *resty.Response, error) {
	return client.Entity().FactureIn().Update(ctx, factureIn.GetID(), &factureIn, params...)
}

// Create shortcut
func (factureIn FactureIn) Create(ctx context.Context, client *Client, params ...*Params) (*FactureIn, *resty.Response, error) {
	return client.Entity().FactureIn().Create(ctx, &factureIn, params...)
}

// Delete shortcut
func (factureIn FactureIn) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return client.Entity().FactureIn().Delete(ctx, factureIn.GetID())
}

// FactureInService
// Сервис для работы со счетами-фактурами полученными.
type FactureInService interface {
	GetList(ctx context.Context, params ...*Params) (*List[FactureIn], *resty.Response, error)
	Create(ctx context.Context, factureIn *FactureIn, params ...*Params) (*FactureIn, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, factureInList Slice[FactureIn], params ...*Params) (*Slice[FactureIn], *resty.Response, error)
	DeleteMany(ctx context.Context, entities ...*FactureIn) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*FactureIn, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, factureIn *FactureIn, params ...*Params) (*FactureIn, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetaAttributesSharedStatesWrapper, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributeMany(ctx context.Context, attributes ...*Attribute) (*Slice[Attribute], *resty.Response, error)
	UpdateAttribute(ctx context.Context, id uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributeMany(ctx context.Context, attributes ...*Attribute) (*DeleteManyResponse, *resty.Response, error)
	Template(ctx context.Context) (*FactureIn, *resty.Response, error)
	GetPublications(ctx context.Context, id uuid.UUID) (*MetaArray[Publication], *resty.Response, error)
	GetPublicationByID(ctx context.Context, id uuid.UUID, publicationID uuid.UUID) (*Publication, *resty.Response, error)
	Publish(ctx context.Context, id uuid.UUID, template TemplateInterface) (*Publication, *resty.Response, error)
	DeletePublication(ctx context.Context, id uuid.UUID, publicationID uuid.UUID) (bool, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*FactureIn, *resty.Response, error)
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
}

func NewFactureInService(client *Client) FactureInService {
	e := NewEndpoint(client, "entity/facturein")
	return newMainService[FactureIn, any, MetaAttributesSharedStatesWrapper, any](e)
}
