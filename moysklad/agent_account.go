package moysklad

import (
	"github.com/google/uuid"
)

// AgentAccount Счёт Контрагента.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-kontragenty-attributy-suschnosti-adres-scheta-kontragentow
type AgentAccount struct {
	AccountID            *uuid.UUID `json:"accountId,omitempty"`            // ID учетной записи
	AccountNumber        *string    `json:"accountNumber,omitempty"`        // Номер счета
	BankLocation         *string    `json:"bankLocation,omitempty"`         // Адрес банка
	BankName             *string    `json:"bankName,omitempty"`             // Наименование банка
	Bic                  *string    `json:"bic,omitempty"`                  // БИК
	CorrespondentAccount *string    `json:"correspondentAccount,omitempty"` // Корр счет
	ID                   *uuid.UUID `json:"id,omitempty"`                   // ID сущности
	IsDefault            *bool      `json:"isDefault,omitempty"`            // Является ли счет основным счетом Контрагента
	Meta                 *Meta      `json:"meta,omitempty"`                 // Метаданные
	Updated              *Timestamp `json:"updated,omitempty"`              // Момент последнего обновления
}

// Clean возвращает сущность с единственным заполненным полем Meta
func (agentAccount AgentAccount) Clean() *AgentAccount {
	return &AgentAccount{Meta: agentAccount.Meta}
}
func (agentAccount AgentAccount) GetAccountID() uuid.UUID {
	return Deref(agentAccount.AccountID)
}

func (agentAccount AgentAccount) GetAccountNumber() string {
	return Deref(agentAccount.AccountNumber)
}

func (agentAccount AgentAccount) GetBankLocation() string {
	return Deref(agentAccount.BankLocation)
}

func (agentAccount AgentAccount) GetBankName() string {
	return Deref(agentAccount.BankName)
}

func (agentAccount AgentAccount) GetBic() string {
	return Deref(agentAccount.Bic)
}

func (agentAccount AgentAccount) GetCorrespondentAccount() string {
	return Deref(agentAccount.CorrespondentAccount)
}

func (agentAccount AgentAccount) GetID() uuid.UUID {
	return Deref(agentAccount.ID)
}

func (agentAccount AgentAccount) GetIsDefault() bool {
	return Deref(agentAccount.IsDefault)
}

func (agentAccount AgentAccount) GetMeta() Meta {
	return Deref(agentAccount.Meta)
}

func (agentAccount AgentAccount) GetUpdated() Timestamp {
	return Deref(agentAccount.Updated)
}

func (agentAccount *AgentAccount) SetAccountNumber(accountNumber string) *AgentAccount {
	agentAccount.AccountNumber = &accountNumber
	return agentAccount
}

func (agentAccount *AgentAccount) SetBankLocation(bankLocation string) *AgentAccount {
	agentAccount.BankLocation = &bankLocation
	return agentAccount
}

func (agentAccount *AgentAccount) SetBankName(bankName string) *AgentAccount {
	agentAccount.BankName = &bankName
	return agentAccount
}

func (agentAccount *AgentAccount) SetBic(bic string) *AgentAccount {
	agentAccount.Bic = &bic
	return agentAccount
}

func (agentAccount *AgentAccount) SetCorrespondentAccount(correspondentAccount string) *AgentAccount {
	agentAccount.CorrespondentAccount = &correspondentAccount
	return agentAccount
}

func (agentAccount *AgentAccount) SetIsDefault(isDefault bool) *AgentAccount {
	agentAccount.IsDefault = &isDefault
	return agentAccount
}

func (agentAccount *AgentAccount) SetMeta(meta *Meta) *AgentAccount {
	agentAccount.Meta = meta
	return agentAccount
}

func (agentAccount AgentAccount) String() string {
	return Stringify(agentAccount)
}
