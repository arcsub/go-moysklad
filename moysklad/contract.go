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
	RewardPercent       *float64      `json:"rewardPercent,omitempty"`
	Updated             *Timestamp    `json:"updated,omitempty"`
	Moment              *Timestamp    `json:"moment,omitempty"`
	Archived            *bool         `json:"archived,omitempty"`
	Attributes          *Attributes   `json:"attributes,omitempty"`
	Code                *string       `json:"code,omitempty"`
	Name                *string       `json:"name,omitempty"`
	Description         *string       `json:"description,omitempty"`
	ExternalCode        *string       `json:"externalCode,omitempty"`
	Group               *Group        `json:"group,omitempty"`
	ID                  *uuid.UUID    `json:"id,omitempty"`
	Meta                *Meta         `json:"meta,omitempty"`
	AgentAccount        *AgentAccount `json:"agentAccount,omitempty"`
	Agent               *Counterparty `json:"agent,omitempty"`
	State               *State        `json:"state,omitempty"`
	OwnAgent            *Organization `json:"ownAgent,omitempty"`
	Owner               *Employee     `json:"owner,omitempty"`
	Rate                *Rate         `json:"rate,omitempty"`
	AccountID           *uuid.UUID    `json:"accountId,omitempty"`
	Published           *bool         `json:"published,omitempty"`
	Shared              *bool         `json:"shared,omitempty"`
	OrganizationAccount *AgentAccount `json:"organizationAccount,omitempty"`
	Sum                 *Decimal      `json:"sum,omitempty"`
	SyncID              *uuid.UUID    `json:"syncId,omitempty"`
	Printed             *bool         `json:"printed,omitempty"`
	RewardType          RewardType    `json:"rewardType,omitempty"`
	ContractType        ContractType  `json:"contractType,omitempty"`
}

func (c Contract) String() string {
	return Stringify(c)
}

func (c Contract) MetaType() MetaType {
	return MetaTypeContract
}

// ContractType Тип Договора.
type ContractType string

const (
	ContractTypeCommission ContractType = "Commission" // Договор комиссии
	ContractTypeSales      ContractType = "Sales"      // Договор купли-продажи
)

// ContractService
// Сервис для работы с договорами.
type ContractService interface {
	GetList(ctx context.Context, params *Params) (*List[Contract], *resty.Response, error)
	Create(ctx context.Context, contract *Contract, params *Params) (*Contract, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, contractList []*Contract, params *Params) (*[]Contract, *resty.Response, error)
	DeleteMany(ctx context.Context, contractList []*Contract) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id *uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error)
	UpdateAttribute(ctx context.Context, id *uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList []*Attribute) (*DeleteManyResponse, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Contract, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, contract *Contract, params *Params) (*Contract, *resty.Response, error)
	GetPublications(ctx context.Context, id *uuid.UUID) (*MetaArray[Publication], *resty.Response, error)
	GetPublicationByID(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (*Publication, *resty.Response, error)
	Publish(ctx context.Context, id *uuid.UUID, template *Templater) (*Publication, *resty.Response, error)
	DeletePublication(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (bool, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
	Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewContractService(client *Client) ContractService {
	e := NewEndpoint(client, "entity/contract")
	return newMainService[Contract, any, MetadataAttributeSharedStates, any](e)
}
