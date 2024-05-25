package moysklad

import (
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
