package moysklad

import (
	"github.com/google/uuid"
)

// Contract Договор.
// Ключевое слово: contract
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-dogowor
type Contract struct {
	AccountID           *uuid.UUID    `json:"accountId,omitempty"`           // ID учетной записи
	Agent               *Counterparty `json:"agent,omitempty"`               // Метаданные Контрагента
	AgentAccount        *AgentAccount `json:"agentAccount,omitempty"`        // Метаданные счета контрагента
	Archived            *bool         `json:"archived,omitempty"`            // Добавлен ли Договор в архив
	Attributes          *Attributes   `json:"attributes,omitempty"`          // Коллекция доп. полей
	Code                *string       `json:"code,omitempty"`                // Код Договора
	ContractType        ContractType  `json:"contractType,omitempty"`        // Тип Договора. Возможные значения: Договор комиссии, Договор купли-продажи
	Description         *string       `json:"description,omitempty"`         // Описание Договора
	ExternalCode        *string       `json:"externalCode,omitempty"`        // Внешний код Договора
	Group               *Group        `json:"group,omitempty"`               // Метаданные отдела сотрудника
	ID                  *uuid.UUID    `json:"id,omitempty"`                  // ID сущности
	Meta                *Meta         `json:"meta,omitempty"`                // Метаданные
	Moment              *Timestamp    `json:"moment,omitempty"`              // Дата Договора
	Name                *string       `json:"name,omitempty"`                // Наименование
	OrganizationAccount *AgentAccount `json:"organizationAccount,omitempty"` // Метаданные счета вашего юрлица
	OwnAgent            *Organization `json:"ownAgent,omitempty"`            // Метаданные вашего юрлица
	Owner               *Employee     `json:"owner,omitempty"`               // Метаданные владельца (Сотрудника)
	Rate                *Rate         `json:"rate,omitempty"`                // Метаданные валюты
	RewardPercent       *int          `json:"rewardPercent,omitempty"`       // Вознаграждение в процентах (от 0 до 100)
	RewardType          RewardType    `json:"rewardType,omitempty"`          // Тип Вознаграждения. Возможные значения: Процент от суммы продажи, Не рассчитывать
	Shared              *bool         `json:"shared,omitempty"`              // Общий доступ
	State               *State        `json:"state,omitempty"`               // Метаданные статуса договора
	Sum                 *float64      `json:"sum,omitempty"`                 // Сумма Договора
	SyncID              *uuid.UUID    `json:"syncId,omitempty"`
	Printed             *bool         `json:"printed,omitempty"`   // Напечатан ли документ
	Published           *bool         `json:"published,omitempty"` // Опубликован ли документ
	Updated             *Timestamp    `json:"updated,omitempty"`   // Момент последнего обновления сущности
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
