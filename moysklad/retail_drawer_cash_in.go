package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// RetailDrawerCashIn Внесение денег.
// Ключевое слово: retaildrawercashin
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vnesenie-deneg
type RetailDrawerCashIn struct {
	Moment       *Timestamp            `json:"moment,omitempty"`
	Created      *Timestamp            `json:"created,omitempty"`
	AccountID    *uuid.UUID            `json:"accountId,omitempty"`
	RetailShift  *RetailShift          `json:"retailShift,omitempty"`
	Name         *string               `json:"name,omitempty"`
	Deleted      *Timestamp            `json:"deleted,omitempty"`
	Description  *string               `json:"description,omitempty"`
	Organization *Organization         `json:"organization,omitempty"`
	Files        *MetaArray[File]      `json:"files,omitempty"`
	Group        *Group                `json:"group,omitempty"`
	ID           *uuid.UUID            `json:"id,omitempty"`
	Meta         *Meta                 `json:"meta,omitempty"`
	Applicable   *bool                 `json:"applicable,omitempty"`
	Agent        *Counterparty         `json:"agent,omitempty"`
	ExternalCode *string               `json:"externalCode,omitempty"`
	Owner        *Employee             `json:"owner,omitempty"`
	Printed      *bool                 `json:"printed,omitempty"`
	Published    *bool                 `json:"published,omitempty"`
	Rate         *Rate                 `json:"rate,omitempty"`
	Shared       *bool                 `json:"shared,omitempty"`
	State        *State                `json:"state,omitempty"`
	Sum          *float64              `json:"sum,omitempty"`
	SyncID       *uuid.UUID            `json:"syncId,omitempty"`
	Updated      *Timestamp            `json:"updated,omitempty"`
	Attributes   Slice[AttributeValue] `json:"attributes,omitempty"`
}

func (retailDrawerCashIn RetailDrawerCashIn) Clean() *RetailDrawerCashIn {
	return &RetailDrawerCashIn{Meta: retailDrawerCashIn.Meta}
}

func (retailDrawerCashIn RetailDrawerCashIn) GetMoment() Timestamp {
	return Deref(retailDrawerCashIn.Moment)
}

func (retailDrawerCashIn RetailDrawerCashIn) GetCreated() Timestamp {
	return Deref(retailDrawerCashIn.Created)
}

func (retailDrawerCashIn RetailDrawerCashIn) GetAccountID() uuid.UUID {
	return Deref(retailDrawerCashIn.AccountID)
}

func (retailDrawerCashIn RetailDrawerCashIn) GetRetailShift() RetailShift {
	return Deref(retailDrawerCashIn.RetailShift)
}

func (retailDrawerCashIn RetailDrawerCashIn) GetName() string {
	return Deref(retailDrawerCashIn.Name)
}

func (retailDrawerCashIn RetailDrawerCashIn) GetDeleted() Timestamp {
	return Deref(retailDrawerCashIn.Deleted)
}

func (retailDrawerCashIn RetailDrawerCashIn) GetDescription() string {
	return Deref(retailDrawerCashIn.Description)
}

func (retailDrawerCashIn RetailDrawerCashIn) GetOrganization() Organization {
	return Deref(retailDrawerCashIn.Organization)
}

func (retailDrawerCashIn RetailDrawerCashIn) GetFiles() MetaArray[File] {
	return Deref(retailDrawerCashIn.Files)
}

func (retailDrawerCashIn RetailDrawerCashIn) GetGroup() Group {
	return Deref(retailDrawerCashIn.Group)
}

func (retailDrawerCashIn RetailDrawerCashIn) GetID() uuid.UUID {
	return Deref(retailDrawerCashIn.ID)
}

func (retailDrawerCashIn RetailDrawerCashIn) GetMeta() Meta {
	return Deref(retailDrawerCashIn.Meta)
}

func (retailDrawerCashIn RetailDrawerCashIn) GetApplicable() bool {
	return Deref(retailDrawerCashIn.Applicable)
}

func (retailDrawerCashIn RetailDrawerCashIn) GetAgent() Counterparty {
	return Deref(retailDrawerCashIn.Agent)
}

func (retailDrawerCashIn RetailDrawerCashIn) GetExternalCode() string {
	return Deref(retailDrawerCashIn.ExternalCode)
}

func (retailDrawerCashIn RetailDrawerCashIn) GetOwner() Employee {
	return Deref(retailDrawerCashIn.Owner)
}

func (retailDrawerCashIn RetailDrawerCashIn) GetPrinted() bool {
	return Deref(retailDrawerCashIn.Printed)
}

func (retailDrawerCashIn RetailDrawerCashIn) GetPublished() bool {
	return Deref(retailDrawerCashIn.Published)
}

func (retailDrawerCashIn RetailDrawerCashIn) GetRate() Rate {
	return Deref(retailDrawerCashIn.Rate)
}

func (retailDrawerCashIn RetailDrawerCashIn) GetShared() bool {
	return Deref(retailDrawerCashIn.Shared)
}

func (retailDrawerCashIn RetailDrawerCashIn) GetState() State {
	return Deref(retailDrawerCashIn.State)
}

func (retailDrawerCashIn RetailDrawerCashIn) GetSum() float64 {
	return Deref(retailDrawerCashIn.Sum)
}

func (retailDrawerCashIn RetailDrawerCashIn) GetSyncID() uuid.UUID {
	return Deref(retailDrawerCashIn.SyncID)
}

func (retailDrawerCashIn RetailDrawerCashIn) GetUpdated() Timestamp {
	return Deref(retailDrawerCashIn.Updated)
}

func (retailDrawerCashIn RetailDrawerCashIn) GetAttributes() Slice[AttributeValue] {
	return retailDrawerCashIn.Attributes
}

func (retailDrawerCashIn *RetailDrawerCashIn) SetMoment(moment *Timestamp) *RetailDrawerCashIn {
	retailDrawerCashIn.Moment = moment
	return retailDrawerCashIn
}

func (retailDrawerCashIn *RetailDrawerCashIn) SetRetailShift(retailShift *RetailShift) *RetailDrawerCashIn {
	retailDrawerCashIn.RetailShift = retailShift.Clean()
	return retailDrawerCashIn
}

func (retailDrawerCashIn *RetailDrawerCashIn) SetName(name string) *RetailDrawerCashIn {
	retailDrawerCashIn.Name = &name
	return retailDrawerCashIn
}

func (retailDrawerCashIn *RetailDrawerCashIn) SetDescription(description string) *RetailDrawerCashIn {
	retailDrawerCashIn.Description = &description
	return retailDrawerCashIn
}

func (retailDrawerCashIn *RetailDrawerCashIn) SetOrganization(organization *Organization) *RetailDrawerCashIn {
	retailDrawerCashIn.Organization = organization.Clean()
	return retailDrawerCashIn
}

func (retailDrawerCashIn *RetailDrawerCashIn) SetFiles(files Slice[File]) *RetailDrawerCashIn {
	retailDrawerCashIn.Files = NewMetaArrayRows(files)
	return retailDrawerCashIn
}

func (retailDrawerCashIn *RetailDrawerCashIn) SetGroup(group *Group) *RetailDrawerCashIn {
	retailDrawerCashIn.Group = group.Clean()
	return retailDrawerCashIn
}

func (retailDrawerCashIn *RetailDrawerCashIn) SetMeta(meta *Meta) *RetailDrawerCashIn {
	retailDrawerCashIn.Meta = meta
	return retailDrawerCashIn
}

func (retailDrawerCashIn *RetailDrawerCashIn) SetApplicable(applicable bool) *RetailDrawerCashIn {
	retailDrawerCashIn.Applicable = &applicable
	return retailDrawerCashIn
}

func (retailDrawerCashIn *RetailDrawerCashIn) SetAgent(agent *Counterparty) *RetailDrawerCashIn {
	retailDrawerCashIn.Agent = agent.Clean()
	return retailDrawerCashIn
}

func (retailDrawerCashIn *RetailDrawerCashIn) SetExternalCode(externalCode string) *RetailDrawerCashIn {
	retailDrawerCashIn.ExternalCode = &externalCode
	return retailDrawerCashIn
}

func (retailDrawerCashIn *RetailDrawerCashIn) SetOwner(owner *Employee) *RetailDrawerCashIn {
	retailDrawerCashIn.Owner = owner.Clean()
	return retailDrawerCashIn
}

func (retailDrawerCashIn *RetailDrawerCashIn) SetRate(rate *Rate) *RetailDrawerCashIn {
	retailDrawerCashIn.Rate = rate
	return retailDrawerCashIn
}

func (retailDrawerCashIn *RetailDrawerCashIn) SetShared(shared bool) *RetailDrawerCashIn {
	retailDrawerCashIn.Shared = &shared
	return retailDrawerCashIn
}

func (retailDrawerCashIn *RetailDrawerCashIn) SetState(state *State) *RetailDrawerCashIn {
	retailDrawerCashIn.State = state.Clean()
	return retailDrawerCashIn
}

func (retailDrawerCashIn *RetailDrawerCashIn) SetSyncID(syncID uuid.UUID) *RetailDrawerCashIn {
	retailDrawerCashIn.SyncID = &syncID
	return retailDrawerCashIn
}

func (retailDrawerCashIn *RetailDrawerCashIn) SetAttributes(attributes Slice[AttributeValue]) *RetailDrawerCashIn {
	retailDrawerCashIn.Attributes = attributes
	return retailDrawerCashIn
}

func (retailDrawerCashIn RetailDrawerCashIn) String() string {
	return Stringify(retailDrawerCashIn)
}

func (retailDrawerCashIn RetailDrawerCashIn) MetaType() MetaType {
	return MetaTypeRetailDrawerCashIn
}

// RetailDrawerCashInService
// Сервис для работы с внесениями денег.
type RetailDrawerCashInService interface {
	GetList(ctx context.Context, params *Params) (*List[RetailDrawerCashIn], *resty.Response, error)
	Create(ctx context.Context, retailDrawerCashIn *RetailDrawerCashIn, params *Params) (*RetailDrawerCashIn, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, retailDrawerCashInList Slice[RetailDrawerCashIn], params *Params) (*Slice[RetailDrawerCashIn], *resty.Response, error)
	DeleteMany(ctx context.Context, retailDrawerCashInList []MetaWrapper) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params *Params) (*RetailDrawerCashIn, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, retailDrawerCashIn *RetailDrawerCashIn, params *Params) (*RetailDrawerCashIn, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetaAttributesSharedStatesWrapper, *resty.Response, error)
	//endpointTemplate[RetailDrawerCashIn]
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
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*RetailDrawerCashIn, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)
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

func NewRetailDrawerCashInService(client *Client) RetailDrawerCashInService {
	e := NewEndpoint(client, "entity/retaildrawercashin")
	return newMainService[RetailDrawerCashIn, any, MetaAttributesSharedStatesWrapper, any](e)
}
