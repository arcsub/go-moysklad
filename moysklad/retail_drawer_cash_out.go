package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// RetailDrawerCashOut Выплата денег.
// Ключевое слово: retaildrawercashout
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vyplata-deneg
type RetailDrawerCashOut struct {
	Meta         *Meta             `json:"meta,omitempty"`
	Applicable   *bool             `json:"applicable,omitempty"`
	Moment       *Timestamp        `json:"moment,omitempty"`
	Name         *string           `json:"name,omitempty"`
	Code         *string           `json:"code,omitempty"`
	Created      *Timestamp        `json:"created,omitempty"`
	Deleted      *Timestamp        `json:"deleted,omitempty"`
	Description  *string           `json:"description,omitempty"`
	ExternalCode *string           `json:"externalCode,omitempty"`
	Files        *MetaArray[File]  `json:"files,omitempty"`
	Group        *Group            `json:"group,omitempty"`
	ID           *uuid.UUID        `json:"id,omitempty"`
	RetailShift  *RetailShift      `json:"retailShift,omitempty"`
	Agent        *Counterparty     `json:"agent,omitempty"`
	AccountID    *uuid.UUID        `json:"accountId,omitempty"`
	Organization *Organization     `json:"organization,omitempty"`
	Owner        *Employee         `json:"owner,omitempty"`
	Printed      *bool             `json:"printed,omitempty"`
	Published    *bool             `json:"published,omitempty"`
	Rate         *NullValue[Rate]  `json:"rate,omitempty"`
	Shared       *bool             `json:"shared,omitempty"`
	State        *NullValue[State] `json:"state,omitempty"`
	Sum          *float64          `json:"sum,omitempty"`
	SyncID       *uuid.UUID        `json:"syncId,omitempty"`
	Updated      *Timestamp        `json:"updated,omitempty"`
	Attributes   Slice[Attribute]  `json:"attributes,omitempty"`
}

// Clean возвращает сущность с единственным заполненным полем Meta
func (retailDrawerCashOut RetailDrawerCashOut) Clean() *RetailDrawerCashOut {
	return &RetailDrawerCashOut{Meta: retailDrawerCashOut.Meta}
}

// AsTaskOperation реализует интерфейс AsTaskOperationInterface
func (retailDrawerCashOut RetailDrawerCashOut) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: retailDrawerCashOut.Meta}
}

func (retailDrawerCashOut RetailDrawerCashOut) GetMeta() Meta {
	return Deref(retailDrawerCashOut.Meta)
}

func (retailDrawerCashOut RetailDrawerCashOut) GetApplicable() bool {
	return Deref(retailDrawerCashOut.Applicable)
}

func (retailDrawerCashOut RetailDrawerCashOut) GetMoment() Timestamp {
	return Deref(retailDrawerCashOut.Moment)
}

func (retailDrawerCashOut RetailDrawerCashOut) GetName() string {
	return Deref(retailDrawerCashOut.Name)
}

func (retailDrawerCashOut RetailDrawerCashOut) GetCode() string {
	return Deref(retailDrawerCashOut.Code)
}

func (retailDrawerCashOut RetailDrawerCashOut) GetCreated() Timestamp {
	return Deref(retailDrawerCashOut.Created)
}

func (retailDrawerCashOut RetailDrawerCashOut) GetDeleted() Timestamp {
	return Deref(retailDrawerCashOut.Deleted)
}

func (retailDrawerCashOut RetailDrawerCashOut) GetDescription() string {
	return Deref(retailDrawerCashOut.Description)
}

func (retailDrawerCashOut RetailDrawerCashOut) GetExternalCode() string {
	return Deref(retailDrawerCashOut.ExternalCode)
}

func (retailDrawerCashOut RetailDrawerCashOut) GetFiles() MetaArray[File] {
	return Deref(retailDrawerCashOut.Files)
}

func (retailDrawerCashOut RetailDrawerCashOut) GetGroup() Group {
	return Deref(retailDrawerCashOut.Group)
}

func (retailDrawerCashOut RetailDrawerCashOut) GetID() uuid.UUID {
	return Deref(retailDrawerCashOut.ID)
}

func (retailDrawerCashOut RetailDrawerCashOut) GetRetailShift() RetailShift {
	return Deref(retailDrawerCashOut.RetailShift)
}

func (retailDrawerCashOut RetailDrawerCashOut) GetAgent() Counterparty {
	return Deref(retailDrawerCashOut.Agent)
}

func (retailDrawerCashOut RetailDrawerCashOut) GetAccountID() uuid.UUID {
	return Deref(retailDrawerCashOut.AccountID)
}

func (retailDrawerCashOut RetailDrawerCashOut) GetOrganization() Organization {
	return Deref(retailDrawerCashOut.Organization)
}

func (retailDrawerCashOut RetailDrawerCashOut) GetOwner() Employee {
	return Deref(retailDrawerCashOut.Owner)
}

func (retailDrawerCashOut RetailDrawerCashOut) GetPrinted() bool {
	return Deref(retailDrawerCashOut.Printed)
}

func (retailDrawerCashOut RetailDrawerCashOut) GetPublished() bool {
	return Deref(retailDrawerCashOut.Published)
}

func (retailDrawerCashOut RetailDrawerCashOut) GetRate() Rate {
	return retailDrawerCashOut.Rate.Get()
}

func (retailDrawerCashOut RetailDrawerCashOut) GetShared() bool {
	return Deref(retailDrawerCashOut.Shared)
}

func (retailDrawerCashOut RetailDrawerCashOut) GetState() State {
	return retailDrawerCashOut.State.Get()
}

func (retailDrawerCashOut RetailDrawerCashOut) GetSum() float64 {
	return Deref(retailDrawerCashOut.Sum)
}

func (retailDrawerCashOut RetailDrawerCashOut) GetSyncID() uuid.UUID {
	return Deref(retailDrawerCashOut.SyncID)
}

func (retailDrawerCashOut RetailDrawerCashOut) GetUpdated() Timestamp {
	return Deref(retailDrawerCashOut.Updated)
}

func (retailDrawerCashOut RetailDrawerCashOut) GetAttributes() Slice[Attribute] {
	return retailDrawerCashOut.Attributes
}

func (retailDrawerCashOut *RetailDrawerCashOut) SetMeta(meta *Meta) *RetailDrawerCashOut {
	retailDrawerCashOut.Meta = meta
	return retailDrawerCashOut
}

func (retailDrawerCashOut *RetailDrawerCashOut) SetApplicable(applicable bool) *RetailDrawerCashOut {
	retailDrawerCashOut.Applicable = &applicable
	return retailDrawerCashOut
}

func (retailDrawerCashOut *RetailDrawerCashOut) SetMoment(moment *Timestamp) *RetailDrawerCashOut {
	retailDrawerCashOut.Moment = moment
	return retailDrawerCashOut
}

func (retailDrawerCashOut *RetailDrawerCashOut) SetName(name string) *RetailDrawerCashOut {
	retailDrawerCashOut.Name = &name
	return retailDrawerCashOut
}

func (retailDrawerCashOut *RetailDrawerCashOut) SetCode(code string) *RetailDrawerCashOut {
	retailDrawerCashOut.Code = &code
	return retailDrawerCashOut
}

func (retailDrawerCashOut *RetailDrawerCashOut) SetDescription(description string) *RetailDrawerCashOut {
	retailDrawerCashOut.Description = &description
	return retailDrawerCashOut
}

func (retailDrawerCashOut *RetailDrawerCashOut) SetExternalCode(externalCode string) *RetailDrawerCashOut {
	retailDrawerCashOut.ExternalCode = &externalCode
	return retailDrawerCashOut
}

func (retailDrawerCashOut *RetailDrawerCashOut) SetFiles(files ...*File) *RetailDrawerCashOut {
	retailDrawerCashOut.Files = NewMetaArrayFrom(files)
	return retailDrawerCashOut
}

func (retailDrawerCashOut *RetailDrawerCashOut) SetGroup(group *Group) *RetailDrawerCashOut {
	retailDrawerCashOut.Group = group.Clean()
	return retailDrawerCashOut
}

func (retailDrawerCashOut *RetailDrawerCashOut) SetRetailShift(retailShift *RetailShift) *RetailDrawerCashOut {
	retailDrawerCashOut.RetailShift = retailShift.Clean()
	return retailDrawerCashOut
}

func (retailDrawerCashOut *RetailDrawerCashOut) SetAgent(agent *Counterparty) *RetailDrawerCashOut {
	retailDrawerCashOut.Agent = agent.Clean()
	return retailDrawerCashOut
}

func (retailDrawerCashOut *RetailDrawerCashOut) SetOrganization(organization *Organization) *RetailDrawerCashOut {
	retailDrawerCashOut.Organization = organization.Clean()
	return retailDrawerCashOut
}

func (retailDrawerCashOut *RetailDrawerCashOut) SetOwner(owner *Employee) *RetailDrawerCashOut {
	retailDrawerCashOut.Owner = owner.Clean()
	return retailDrawerCashOut
}

func (retailDrawerCashOut *RetailDrawerCashOut) SetRate(rate *Rate) *RetailDrawerCashOut {
	retailDrawerCashOut.Rate = NewNullValueFrom(rate)
	return retailDrawerCashOut
}

func (retailDrawerCashOut *RetailDrawerCashOut) SetNullRate() *RetailDrawerCashOut {
	retailDrawerCashOut.Rate = NewNullValue[Rate]()
	return retailDrawerCashOut
}

func (retailDrawerCashOut *RetailDrawerCashOut) SetShared(shared bool) *RetailDrawerCashOut {
	retailDrawerCashOut.Shared = &shared
	return retailDrawerCashOut
}

func (retailDrawerCashOut *RetailDrawerCashOut) SetState(state *State) *RetailDrawerCashOut {
	retailDrawerCashOut.State = NewNullValueFrom(state.Clean())
	return retailDrawerCashOut
}

func (retailDrawerCashOut *RetailDrawerCashOut) SetNullState() *RetailDrawerCashOut {
	retailDrawerCashOut.State = NewNullValue[State]()
	return retailDrawerCashOut
}

func (retailDrawerCashOut *RetailDrawerCashOut) SetSyncID(syncID uuid.UUID) *RetailDrawerCashOut {
	retailDrawerCashOut.SyncID = &syncID
	return retailDrawerCashOut
}

func (retailDrawerCashOut *RetailDrawerCashOut) SetAttributes(attributes ...*Attribute) *RetailDrawerCashOut {
	retailDrawerCashOut.Attributes = attributes
	return retailDrawerCashOut
}

func (retailDrawerCashOut RetailDrawerCashOut) String() string {
	return Stringify(retailDrawerCashOut)
}

// MetaType возвращает тип сущности.
func (RetailDrawerCashOut) MetaType() MetaType {
	return MetaTypeRetailDrawerCashOut
}

// Update shortcut
func (retailDrawerCashOut RetailDrawerCashOut) Update(ctx context.Context, client *Client, params ...*Params) (*RetailDrawerCashOut, *resty.Response, error) {
	return client.Entity().RetailDrawerCashOut().Update(ctx, retailDrawerCashOut.GetID(), &retailDrawerCashOut, params...)
}

// Create shortcut
func (retailDrawerCashOut RetailDrawerCashOut) Create(ctx context.Context, client *Client, params ...*Params) (*RetailDrawerCashOut, *resty.Response, error) {
	return client.Entity().RetailDrawerCashOut().Create(ctx, &retailDrawerCashOut, params...)
}

// Delete shortcut
func (retailDrawerCashOut RetailDrawerCashOut) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return client.Entity().RetailDrawerCashOut().Delete(ctx, retailDrawerCashOut.GetID())
}

// RetailDrawerCashOutService
// Сервис для работы с выплатами денег.
type RetailDrawerCashOutService interface {
	GetList(ctx context.Context, params ...*Params) (*List[RetailDrawerCashOut], *resty.Response, error)
	Create(ctx context.Context, retailDrawerCashOut *RetailDrawerCashOut, params ...*Params) (*RetailDrawerCashOut, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, retailDrawerCashOutList Slice[RetailDrawerCashOut], params ...*Params) (*Slice[RetailDrawerCashOut], *resty.Response, error)
	DeleteMany(ctx context.Context, entities ...*RetailDrawerCashOut) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*RetailDrawerCashOut, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, retailDrawerCashOut *RetailDrawerCashOut, params ...*Params) (*RetailDrawerCashOut, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetaAttributesSharedStatesWrapper, *resty.Response, error)
	Template(ctx context.Context) (*RetailDrawerCashOut, *resty.Response, error)
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
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*RetailDrawerCashOut, *resty.Response, error)
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
}

func NewRetailDrawerCashOutService(client *Client) RetailDrawerCashOutService {
	e := NewEndpoint(client, "entity/retaildrawercashout")
	return newMainService[RetailDrawerCashOut, any, MetaAttributesSharedStatesWrapper, any](e)
}
