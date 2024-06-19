package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// CounterpartyAdjustment Корректировка взаиморасчетов.
// Ключевое слово: counterpartyadjustment
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-korrektirowka-wzaimoraschetow
type CounterpartyAdjustment struct {
	ExternalCode *string          `json:"externalCode,omitempty"`
	Printed      *bool            `json:"printed,omitempty"`
	AccountID    *uuid.UUID       `json:"accountId,omitempty"`
	Group        *Group           `json:"group,omitempty"`
	Files        *MetaArray[File] `json:"files,omitempty"`
	Applicable   *bool            `json:"applicable,omitempty"`
	Updated      *Timestamp       `json:"updated,omitempty"`
	Created      *Timestamp       `json:"created,omitempty"`
	Deleted      *Timestamp       `json:"deleted,omitempty"`
	Description  *string          `json:"description,omitempty"`
	Name         *string          `json:"name,omitempty"`
	Agent        *Counterparty    `json:"agent,omitempty"`
	Meta         *Meta            `json:"meta,omitempty"`
	Moment       *Timestamp       `json:"moment,omitempty"`
	Organization *Organization    `json:"organization,omitempty"`
	Owner        *Employee        `json:"owner,omitempty"`
	ID           *uuid.UUID       `json:"id,omitempty"`
	Published    *bool            `json:"published,omitempty"`
	Shared       *bool            `json:"shared,omitempty"`
	Sum          *float64         `json:"sum,omitempty"`
	Attributes   Slice[Attribute] `json:"attributes,omitempty"`
}

// Clean возвращает сущность с единственным заполненным полем Meta
func (counterPartyAdjustment CounterpartyAdjustment) Clean() *CounterpartyAdjustment {
	return &CounterpartyAdjustment{Meta: counterPartyAdjustment.Meta}
}

// AsTaskOperation реализует интерфейс AsTaskOperationInterface
func (counterPartyAdjustment CounterpartyAdjustment) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: counterPartyAdjustment.Meta}
}

func (counterPartyAdjustment CounterpartyAdjustment) GetExternalCode() string {
	return Deref(counterPartyAdjustment.ExternalCode)
}

func (counterPartyAdjustment CounterpartyAdjustment) GetPrinted() bool {
	return Deref(counterPartyAdjustment.Printed)
}

func (counterPartyAdjustment CounterpartyAdjustment) GetAccountID() uuid.UUID {
	return Deref(counterPartyAdjustment.AccountID)
}

func (counterPartyAdjustment CounterpartyAdjustment) GetGroup() Group {
	return Deref(counterPartyAdjustment.Group)
}

func (counterPartyAdjustment CounterpartyAdjustment) GetFiles() MetaArray[File] {
	return Deref(counterPartyAdjustment.Files)
}

func (counterPartyAdjustment CounterpartyAdjustment) GetApplicable() bool {
	return Deref(counterPartyAdjustment.Applicable)
}

func (counterPartyAdjustment CounterpartyAdjustment) GetUpdated() Timestamp {
	return Deref(counterPartyAdjustment.Updated)
}

func (counterPartyAdjustment CounterpartyAdjustment) GetCreated() Timestamp {
	return Deref(counterPartyAdjustment.Created)
}

func (counterPartyAdjustment CounterpartyAdjustment) GetDeleted() Timestamp {
	return Deref(counterPartyAdjustment.Deleted)
}

func (counterPartyAdjustment CounterpartyAdjustment) GetDescription() string {
	return Deref(counterPartyAdjustment.Description)
}

func (counterPartyAdjustment CounterpartyAdjustment) GetName() string {
	return Deref(counterPartyAdjustment.Name)
}

func (counterPartyAdjustment CounterpartyAdjustment) GetAgent() Counterparty {
	return Deref(counterPartyAdjustment.Agent)
}

func (counterPartyAdjustment CounterpartyAdjustment) GetMeta() Meta {
	return Deref(counterPartyAdjustment.Meta)
}

func (counterPartyAdjustment CounterpartyAdjustment) GetMoment() Timestamp {
	return Deref(counterPartyAdjustment.Moment)
}

func (counterPartyAdjustment CounterpartyAdjustment) GetOrganization() Organization {
	return Deref(counterPartyAdjustment.Organization)
}

func (counterPartyAdjustment CounterpartyAdjustment) GetOwner() Employee {
	return Deref(counterPartyAdjustment.Owner)
}

func (counterPartyAdjustment CounterpartyAdjustment) GetID() uuid.UUID {
	return Deref(counterPartyAdjustment.ID)
}

func (counterPartyAdjustment CounterpartyAdjustment) GetPublished() bool {
	return Deref(counterPartyAdjustment.Published)
}

func (counterPartyAdjustment CounterpartyAdjustment) GetShared() bool {
	return Deref(counterPartyAdjustment.Shared)
}

func (counterPartyAdjustment CounterpartyAdjustment) GetSum() float64 {
	return Deref(counterPartyAdjustment.Sum)
}

func (counterPartyAdjustment CounterpartyAdjustment) GetAttributes() Slice[Attribute] {
	return counterPartyAdjustment.Attributes
}

func (counterPartyAdjustment *CounterpartyAdjustment) SetExternalCode(externalCode string) *CounterpartyAdjustment {
	counterPartyAdjustment.ExternalCode = &externalCode
	return counterPartyAdjustment
}

func (counterPartyAdjustment *CounterpartyAdjustment) SetGroup(group *Group) *CounterpartyAdjustment {
	counterPartyAdjustment.Group = group.Clean()
	return counterPartyAdjustment
}

func (counterPartyAdjustment *CounterpartyAdjustment) SetFiles(files ...*File) *CounterpartyAdjustment {
	counterPartyAdjustment.Files = NewMetaArrayFrom(files)
	return counterPartyAdjustment
}

func (counterPartyAdjustment *CounterpartyAdjustment) SetApplicable(applicable bool) *CounterpartyAdjustment {
	counterPartyAdjustment.Applicable = &applicable
	return counterPartyAdjustment
}

func (counterPartyAdjustment *CounterpartyAdjustment) SetDescription(description string) *CounterpartyAdjustment {
	counterPartyAdjustment.Description = &description
	return counterPartyAdjustment
}

func (counterPartyAdjustment *CounterpartyAdjustment) SetName(name string) *CounterpartyAdjustment {
	counterPartyAdjustment.Name = &name
	return counterPartyAdjustment
}

func (counterPartyAdjustment *CounterpartyAdjustment) SetAgent(agent *Counterparty) *CounterpartyAdjustment {
	counterPartyAdjustment.Agent = agent.Clean()
	return counterPartyAdjustment
}

func (counterPartyAdjustment *CounterpartyAdjustment) SetMeta(meta *Meta) *CounterpartyAdjustment {
	counterPartyAdjustment.Meta = meta
	return counterPartyAdjustment
}

func (counterPartyAdjustment *CounterpartyAdjustment) SetMoment(moment *Timestamp) *CounterpartyAdjustment {
	counterPartyAdjustment.Moment = moment
	return counterPartyAdjustment
}

func (counterPartyAdjustment *CounterpartyAdjustment) SetOrganization(organization *Organization) *CounterpartyAdjustment {
	counterPartyAdjustment.Organization = organization.Clean()
	return counterPartyAdjustment
}

func (counterPartyAdjustment *CounterpartyAdjustment) SetOwner(owner *Employee) *CounterpartyAdjustment {
	counterPartyAdjustment.Owner = owner.Clean()
	return counterPartyAdjustment
}

func (counterPartyAdjustment *CounterpartyAdjustment) SetShared(shared bool) *CounterpartyAdjustment {
	counterPartyAdjustment.Shared = &shared
	return counterPartyAdjustment
}

func (counterPartyAdjustment *CounterpartyAdjustment) SetAttributes(attributes ...*Attribute) *CounterpartyAdjustment {
	counterPartyAdjustment.Attributes = attributes
	return counterPartyAdjustment
}

func (counterPartyAdjustment CounterpartyAdjustment) String() string {
	return Stringify(counterPartyAdjustment)
}

// MetaType возвращает тип сущности.
func (CounterpartyAdjustment) MetaType() MetaType {
	return MetaTypeCounterPartyAdjustment
}

// Update shortcut
func (counterPartyAdjustment CounterpartyAdjustment) Update(ctx context.Context, client *Client, params ...*Params) (*CounterpartyAdjustment, *resty.Response, error) {
	return client.Entity().CounterPartyAdjustment().Update(ctx, counterPartyAdjustment.GetID(), &counterPartyAdjustment, params...)
}

// Create shortcut
func (counterPartyAdjustment CounterpartyAdjustment) Create(ctx context.Context, client *Client, params ...*Params) (*CounterpartyAdjustment, *resty.Response, error) {
	return client.Entity().CounterPartyAdjustment().Create(ctx, &counterPartyAdjustment, params...)
}

// Delete shortcut
func (counterPartyAdjustment CounterpartyAdjustment) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return client.Entity().CounterPartyAdjustment().Delete(ctx, counterPartyAdjustment.GetID())
}

// CounterPartyAdjustmentService
// Сервис для работы с корректировками баланса контрагента.
type CounterPartyAdjustmentService interface {
	GetList(ctx context.Context, params ...*Params) (*List[CounterpartyAdjustment], *resty.Response, error)
	Create(ctx context.Context, counterPartyAdjustment *CounterpartyAdjustment, params ...*Params) (*CounterpartyAdjustment, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, counterPartyAdjustmentList Slice[CounterpartyAdjustment], params ...*Params) (*Slice[CounterpartyAdjustment], *resty.Response, error)
	DeleteMany(ctx context.Context, entities ...*CounterpartyAdjustment) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*CounterpartyAdjustment, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, counterPartyAdjustment *CounterpartyAdjustment, params ...*Params) (*CounterpartyAdjustment, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetaAttributesSharedStatesWrapper, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params ...*Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)
	MoveToTrash(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetFiles(ctx context.Context, id uuid.UUID) (*MetaArray[File], *resty.Response, error)
	CreateFile(ctx context.Context, id uuid.UUID, file *File) (*Slice[File], *resty.Response, error)
	UpdateFileMany(ctx context.Context, id uuid.UUID, files ...*File) (*Slice[File], *resty.Response, error)
	DeleteFile(ctx context.Context, id uuid.UUID, fileID uuid.UUID) (bool, *resty.Response, error)
	DeleteFileMany(ctx context.Context, id uuid.UUID, files ...*File) (*DeleteManyResponse, *resty.Response, error)
}

func NewCounterPartyAdjustmentService(client *Client) CounterPartyAdjustmentService {
	e := NewEndpoint(client, "entity/counterpartyadjustment")
	return newMainService[CounterpartyAdjustment, any, MetaAttributesSharedStatesWrapper, any](e)
}
