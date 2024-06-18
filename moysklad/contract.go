package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Contract Договор.
// Ключевое слово: contract
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-dogowor
type Contract struct {
	AgentAccount        *AgentAccount         `json:"agentAccount,omitempty"`
	Published           *bool                 `json:"published,omitempty"`
	RewardPercent       *float64              `json:"rewardPercent,omitempty"`
	Archived            *bool                 `json:"archived,omitempty"`
	Agent               *Counterparty         `json:"agent,omitempty"`
	Code                *string               `json:"code,omitempty"`
	Name                *string               `json:"name,omitempty"`
	Description         *string               `json:"description,omitempty"`
	ExternalCode        *string               `json:"externalCode,omitempty"`
	Group               *Group                `json:"group,omitempty"`
	ID                  *uuid.UUID            `json:"id,omitempty"`
	Meta                *Meta                 `json:"meta,omitempty"`
	Moment              *Timestamp            `json:"moment,omitempty"`
	Printed             *bool                 `json:"printed,omitempty"`
	OrganizationAccount *AgentAccount         `json:"organizationAccount,omitempty"`
	OwnAgent            *Organization         `json:"ownAgent,omitempty"`
	Owner               *Employee             `json:"owner,omitempty"`
	Rate                *NullValue[Rate]      `json:"rate,omitempty"`
	AccountID           *uuid.UUID            `json:"accountId,omitempty"`
	Updated             *Timestamp            `json:"updated,omitempty"`
	Shared              *bool                 `json:"shared,omitempty"`
	State               *NullValue[State]     `json:"state,omitempty"`
	Sum                 *float64              `json:"sum,omitempty"`
	SyncID              *uuid.UUID            `json:"syncId,omitempty"`
	ContractType        ContractType          `json:"contractType,omitempty"`
	RewardType          RewardType            `json:"rewardType,omitempty"`
	Attributes          Slice[AttributeValue] `json:"attributes,omitempty"`
}

// Clean возвращает сущность с единственным заполненным полем Meta
func (contract Contract) Clean() *Contract {
	return &Contract{Meta: contract.Meta}
}

func (contract Contract) GetAgentAccount() AgentAccount {
	return Deref(contract.AgentAccount)
}

func (contract Contract) GetPublished() bool {
	return Deref(contract.Published)
}

func (contract Contract) GetRewardPercent() float64 {
	return Deref(contract.RewardPercent)
}

func (contract Contract) GetArchived() bool {
	return Deref(contract.Archived)
}

func (contract Contract) GetAgent() Counterparty {
	return Deref(contract.Agent)
}

func (contract Contract) GetCode() string {
	return Deref(contract.Code)
}

func (contract Contract) GetName() string {
	return Deref(contract.Name)
}

func (contract Contract) GetDescription() string {
	return Deref(contract.Description)
}

func (contract Contract) GetExternalCode() string {
	return Deref(contract.ExternalCode)
}

func (contract Contract) GetGroup() Group {
	return Deref(contract.Group)
}

func (contract Contract) GetID() uuid.UUID {
	return Deref(contract.ID)
}

func (contract Contract) GetMeta() Meta {
	return Deref(contract.Meta)
}

func (contract Contract) GetMoment() Timestamp {
	return Deref(contract.Moment)
}

func (contract Contract) GetPrinted() bool {
	return Deref(contract.Printed)
}

func (contract Contract) GetOrganizationAccount() AgentAccount {
	return Deref(contract.OrganizationAccount)
}

func (contract Contract) GetOwnAgent() Organization {
	return Deref(contract.OwnAgent)
}

func (contract Contract) GetOwner() Employee {
	return Deref(contract.Owner)
}

func (contract Contract) GetRate() Rate {
	return contract.Rate.Get()
}

func (contract Contract) GetAccountID() uuid.UUID {
	return Deref(contract.AccountID)
}

func (contract Contract) GetUpdated() Timestamp {
	return Deref(contract.Updated)
}

func (contract Contract) GetShared() bool {
	return Deref(contract.Shared)
}

func (contract Contract) GetState() State {
	return contract.State.Get()
}

func (contract Contract) GetSum() float64 {
	return Deref(contract.Sum)
}

func (contract Contract) GetSyncID() uuid.UUID {
	return Deref(contract.SyncID)
}

func (contract Contract) GetContractType() ContractType {
	return contract.ContractType
}

func (contract Contract) GetRewardType() RewardType {
	return contract.RewardType
}

func (contract Contract) GetAttributes() Slice[AttributeValue] {
	return contract.Attributes
}

func (contract *Contract) SetAgentAccount(agentAccount *AgentAccount) *Contract {
	contract.AgentAccount = agentAccount.Clean()
	return contract
}

func (contract *Contract) SetRewardPercent(rewardPercent float64) *Contract {
	contract.RewardPercent = &rewardPercent
	return contract
}

func (contract *Contract) SetArchived(archived bool) *Contract {
	contract.Archived = &archived
	return contract
}

func (contract *Contract) SetAgent(agent *Counterparty) *Contract {
	contract.Agent = agent.Clean()
	return contract
}

func (contract *Contract) SetCode(code string) *Contract {
	contract.Code = &code
	return contract
}

func (contract *Contract) SetName(name string) *Contract {
	contract.Name = &name
	return contract
}

func (contract *Contract) SetDescription(description string) *Contract {
	contract.Description = &description
	return contract
}

func (contract *Contract) SetExternalCode(externalCode string) *Contract {
	contract.ExternalCode = &externalCode
	return contract
}

func (contract *Contract) SetGroup(group *Group) *Contract {
	contract.Group = group.Clean()
	return contract
}

func (contract *Contract) SetMeta(meta *Meta) *Contract {
	contract.Meta = meta
	return contract
}

func (contract *Contract) SetMoment(moment *Timestamp) *Contract {
	contract.Moment = moment
	return contract
}

func (contract *Contract) SetOrganizationAccount(organizationAccount *AgentAccount) *Contract {
	contract.OrganizationAccount = organizationAccount.Clean()
	return contract
}

func (contract *Contract) SetOwnAgent(ownAgent *Organization) *Contract {
	contract.OwnAgent = ownAgent.Clean()
	return contract
}

func (contract *Contract) SetOwner(owner *Employee) *Contract {
	contract.Owner = owner.Clean()
	return contract
}

func (contract *Contract) SetRate(rate *Rate) *Contract {
	contract.Rate = NewNullValueWith(rate)
	return contract
}

func (contract *Contract) SetNullRate() *Contract {
	contract.Rate = NewNullValue[Rate]()
	return contract
}

func (contract *Contract) SetShared(shared bool) *Contract {
	contract.Shared = &shared
	return contract
}

func (contract *Contract) SetState(state *State) *Contract {
	contract.State = NewNullValueWith(state.Clean())
	return contract
}

func (contract *Contract) SetNullState() *Contract {
	contract.State = NewNullValue[State]()
	return contract
}

func (contract *Contract) SetSum(sum *float64) *Contract {
	contract.Sum = sum
	return contract
}

func (contract *Contract) SetSyncID(syncID uuid.UUID) *Contract {
	contract.SyncID = &syncID
	return contract
}

func (contract *Contract) SetContractType(contractType ContractType) *Contract {
	contract.ContractType = contractType
	return contract
}

func (contract *Contract) SetRewardType(rewardType RewardType) *Contract {
	contract.RewardType = rewardType
	return contract
}

func (contract *Contract) SetAttributes(attributes Slice[AttributeValue]) *Contract {
	contract.Attributes = attributes
	return contract
}

func (contract Contract) String() string {
	return Stringify(contract)
}

func (contract Contract) MetaType() MetaType {
	return MetaTypeContract
}

// ContractType Тип Договора.
type ContractType string

const (
	ContractTypeCommission ContractType = "Commission" // Договор комиссии
	ContractTypeSales      ContractType = "Sales"      // Договор купли-продажи
)

// RewardType Тип Вознаграждения.
type RewardType string

const (
	RewardTypePercentOfSales RewardType = "PercentOfSales" // Процент от суммы продажи
	RewardTypeNone           RewardType = "None"           // Не рассчитывать
)

// ContractService
// Сервис для работы с договорами.
type ContractService interface {
	GetList(ctx context.Context, params ...*Params) (*List[Contract], *resty.Response, error)
	Create(ctx context.Context, contract *Contract, params ...*Params) (*Contract, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, contractList Slice[Contract], params ...*Params) (*Slice[Contract], *resty.Response, error)
	DeleteMany(ctx context.Context, entities ...Contract) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetaAttributesSharedStatesWrapper, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList Slice[Attribute]) (*Slice[Attribute], *resty.Response, error)
	UpdateAttribute(ctx context.Context, id uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList []MetaWrapper) (*DeleteManyResponse, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Contract, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, contract *Contract, params ...*Params) (*Contract, *resty.Response, error)
	GetPublications(ctx context.Context, id uuid.UUID) (*MetaArray[Publication], *resty.Response, error)
	GetPublicationByID(ctx context.Context, id uuid.UUID, publicationID uuid.UUID) (*Publication, *resty.Response, error)
	Publish(ctx context.Context, id uuid.UUID, template Templater) (*Publication, *resty.Response, error)
	DeletePublication(ctx context.Context, id uuid.UUID, publicationID uuid.UUID) (bool, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params ...*Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)
	MoveToTrash(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetStateByID(ctx context.Context, id uuid.UUID) (*State, *resty.Response, error)
	CreateState(ctx context.Context, state *State) (*State, *resty.Response, error)
	UpdateState(ctx context.Context, id uuid.UUID, state *State) (*State, *resty.Response, error)
	CreateOrUpdateStates(ctx context.Context, states Slice[State]) (*Slice[State], *resty.Response, error)
	DeleteState(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
}

func NewContractService(client *Client) ContractService {
	e := NewEndpoint(client, "entity/contract")
	return newMainService[Contract, any, MetaAttributesSharedStatesWrapper, any](e)
}
