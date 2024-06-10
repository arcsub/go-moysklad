package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// CounterPartyAdjustment Корректировка взаиморасчетов.
// Ключевое слово: counterpartyadjustment
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-korrektirowka-wzaimoraschetow
type CounterPartyAdjustment struct {
	ExternalCode *string               `json:"externalCode,omitempty"`
	Printed      *bool                 `json:"printed,omitempty"`
	AccountID    *uuid.UUID            `json:"accountId,omitempty"`
	Group        *Group                `json:"group,omitempty"`
	Files        *MetaArray[File]      `json:"files,omitempty"`
	Applicable   *bool                 `json:"applicable,omitempty"`
	Updated      *Timestamp            `json:"updated,omitempty"`
	Created      *Timestamp            `json:"created,omitempty"`
	Deleted      *Timestamp            `json:"deleted,omitempty"`
	Description  *string               `json:"description,omitempty"`
	Name         *string               `json:"name,omitempty"`
	Agent        *Counterparty         `json:"agent,omitempty"`
	Meta         *Meta                 `json:"meta,omitempty"`
	Moment       *Timestamp            `json:"moment,omitempty"`
	Organization *Organization         `json:"organization,omitempty"`
	Owner        *Employee             `json:"owner,omitempty"`
	ID           *uuid.UUID            `json:"id,omitempty"`
	Published    *bool                 `json:"published,omitempty"`
	Shared       *bool                 `json:"shared,omitempty"`
	Sum          *float64              `json:"sum,omitempty"`
	Attributes   Slice[AttributeValue] `json:"attributes,omitempty"`
}

func (counterPartyAdjustment CounterPartyAdjustment) Clean() *CounterPartyAdjustment {
	return &CounterPartyAdjustment{Meta: counterPartyAdjustment.Meta}
}

func (counterPartyAdjustment CounterPartyAdjustment) GetExternalCode() string {
	return Deref(counterPartyAdjustment.ExternalCode)
}

func (counterPartyAdjustment CounterPartyAdjustment) GetPrinted() bool {
	return Deref(counterPartyAdjustment.Printed)
}

func (counterPartyAdjustment CounterPartyAdjustment) GetAccountID() uuid.UUID {
	return Deref(counterPartyAdjustment.AccountID)
}

func (counterPartyAdjustment CounterPartyAdjustment) GetGroup() Group {
	return Deref(counterPartyAdjustment.Group)
}

func (counterPartyAdjustment CounterPartyAdjustment) GetFiles() MetaArray[File] {
	return Deref(counterPartyAdjustment.Files)
}

func (counterPartyAdjustment CounterPartyAdjustment) GetApplicable() bool {
	return Deref(counterPartyAdjustment.Applicable)
}

func (counterPartyAdjustment CounterPartyAdjustment) GetUpdated() Timestamp {
	return Deref(counterPartyAdjustment.Updated)
}

func (counterPartyAdjustment CounterPartyAdjustment) GetCreated() Timestamp {
	return Deref(counterPartyAdjustment.Created)
}

func (counterPartyAdjustment CounterPartyAdjustment) GetDeleted() Timestamp {
	return Deref(counterPartyAdjustment.Deleted)
}

func (counterPartyAdjustment CounterPartyAdjustment) GetDescription() string {
	return Deref(counterPartyAdjustment.Description)
}

func (counterPartyAdjustment CounterPartyAdjustment) GetName() string {
	return Deref(counterPartyAdjustment.Name)
}

func (counterPartyAdjustment CounterPartyAdjustment) GetAgent() Counterparty {
	return Deref(counterPartyAdjustment.Agent)
}

func (counterPartyAdjustment CounterPartyAdjustment) GetMeta() Meta {
	return Deref(counterPartyAdjustment.Meta)
}

func (counterPartyAdjustment CounterPartyAdjustment) GetMoment() Timestamp {
	return Deref(counterPartyAdjustment.Moment)
}

func (counterPartyAdjustment CounterPartyAdjustment) GetOrganization() Organization {
	return Deref(counterPartyAdjustment.Organization)
}

func (counterPartyAdjustment CounterPartyAdjustment) GetOwner() Employee {
	return Deref(counterPartyAdjustment.Owner)
}

func (counterPartyAdjustment CounterPartyAdjustment) GetID() uuid.UUID {
	return Deref(counterPartyAdjustment.ID)
}

func (counterPartyAdjustment CounterPartyAdjustment) GetPublished() bool {
	return Deref(counterPartyAdjustment.Published)
}

func (counterPartyAdjustment CounterPartyAdjustment) GetShared() bool {
	return Deref(counterPartyAdjustment.Shared)
}

func (counterPartyAdjustment CounterPartyAdjustment) GetSum() float64 {
	return Deref(counterPartyAdjustment.Sum)
}

func (counterPartyAdjustment CounterPartyAdjustment) GetAttributes() Slice[AttributeValue] {
	return counterPartyAdjustment.Attributes
}

func (counterPartyAdjustment *CounterPartyAdjustment) SetExternalCode(externalCode string) *CounterPartyAdjustment {
	counterPartyAdjustment.ExternalCode = &externalCode
	return counterPartyAdjustment
}

func (counterPartyAdjustment *CounterPartyAdjustment) SetGroup(group *Group) *CounterPartyAdjustment {
	counterPartyAdjustment.Group = group.Clean()
	return counterPartyAdjustment
}

func (counterPartyAdjustment *CounterPartyAdjustment) SetFiles(files Slice[File]) *CounterPartyAdjustment {
	counterPartyAdjustment.Files = NewMetaArrayRows(files)
	return counterPartyAdjustment
}

func (counterPartyAdjustment *CounterPartyAdjustment) SetApplicable(applicable bool) *CounterPartyAdjustment {
	counterPartyAdjustment.Applicable = &applicable
	return counterPartyAdjustment
}

func (counterPartyAdjustment *CounterPartyAdjustment) SetDescription(description string) *CounterPartyAdjustment {
	counterPartyAdjustment.Description = &description
	return counterPartyAdjustment
}

func (counterPartyAdjustment *CounterPartyAdjustment) SetName(name string) *CounterPartyAdjustment {
	counterPartyAdjustment.Name = &name
	return counterPartyAdjustment
}

func (counterPartyAdjustment *CounterPartyAdjustment) SetAgent(agent *Counterparty) *CounterPartyAdjustment {
	counterPartyAdjustment.Agent = agent.Clean()
	return counterPartyAdjustment
}

func (counterPartyAdjustment *CounterPartyAdjustment) SetMeta(meta *Meta) *CounterPartyAdjustment {
	counterPartyAdjustment.Meta = meta
	return counterPartyAdjustment
}

func (counterPartyAdjustment *CounterPartyAdjustment) SetMoment(moment *Timestamp) *CounterPartyAdjustment {
	counterPartyAdjustment.Moment = moment
	return counterPartyAdjustment
}

func (counterPartyAdjustment *CounterPartyAdjustment) SetOrganization(organization *Organization) *CounterPartyAdjustment {
	counterPartyAdjustment.Organization = organization.Clean()
	return counterPartyAdjustment
}

func (counterPartyAdjustment *CounterPartyAdjustment) SetOwner(owner *Employee) *CounterPartyAdjustment {
	counterPartyAdjustment.Owner = owner.Clean()
	return counterPartyAdjustment
}

func (counterPartyAdjustment *CounterPartyAdjustment) SetShared(shared bool) *CounterPartyAdjustment {
	counterPartyAdjustment.Shared = &shared
	return counterPartyAdjustment
}

func (counterPartyAdjustment *CounterPartyAdjustment) SetAttributes(attributes Slice[AttributeValue]) *CounterPartyAdjustment {
	counterPartyAdjustment.Attributes = attributes
	return counterPartyAdjustment
}

func (counterPartyAdjustment CounterPartyAdjustment) String() string {
	return Stringify(counterPartyAdjustment)
}

func (counterPartyAdjustment CounterPartyAdjustment) MetaType() MetaType {
	return MetaTypeCounterPartyAdjustment
}

// CounterPartyAdjustmentService
// Сервис для работы с корректировками баланса контрагента.
type CounterPartyAdjustmentService interface {
	GetList(ctx context.Context, params *Params) (*List[CounterPartyAdjustment], *resty.Response, error)
	Create(ctx context.Context, counterPartyAdjustment *CounterPartyAdjustment, params *Params) (*CounterPartyAdjustment, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, counterPartyAdjustmentList []*CounterPartyAdjustment, params *Params) (*[]CounterPartyAdjustment, *resty.Response, error)
	DeleteMany(ctx context.Context, counterPartyAdjustmentList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params *Params) (*CounterPartyAdjustment, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, counterPartyAdjustment *CounterPartyAdjustment, params *Params) (*CounterPartyAdjustment, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetaAttributesSharedStatesWrapper, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)
	MoveToTrash(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
}

func NewCounterPartyAdjustmentService(client *Client) CounterPartyAdjustmentService {
	e := NewEndpoint(client, "entity/counterpartyadjustment")
	return newMainService[CounterPartyAdjustment, any, MetaAttributesSharedStatesWrapper, any](e)
}
