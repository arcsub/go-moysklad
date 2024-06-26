package moysklad

import (
	"github.com/google/uuid"
)

// AgentAccount Счёт Контрагента.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-kontragenty-attributy-suschnosti-adres-scheta-kontragentow
type AgentAccount struct {
	AccountID            *uuid.UUID `json:"accountId,omitempty"`            // ID учётной записи
	AccountNumber        *string    `json:"accountNumber,omitempty"`        // Номер счета
	BankLocation         *string    `json:"bankLocation,omitempty"`         // Адрес банка
	BankName             *string    `json:"bankName,omitempty"`             // Наименование банка
	BIC                  *string    `json:"bic,omitempty"`                  // БИК
	CorrespondentAccount *string    `json:"correspondentAccount,omitempty"` // Корр счет
	ID                   *uuid.UUID `json:"id,omitempty"`                   // ID счета
	IsDefault            *bool      `json:"isDefault,omitempty"`            // Является ли счет основным счетом Контрагента
	Meta                 *Meta      `json:"meta,omitempty"`                 // Метаданные Счета Контрагента
	Updated              *Timestamp `json:"updated,omitempty"`              // Момент последнего обновления
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (agentAccount AgentAccount) Clean() *AgentAccount {
	return &AgentAccount{Meta: agentAccount.Meta}
}

// GetAccountID возвращает ID учётной записи.
func (agentAccount AgentAccount) GetAccountID() uuid.UUID {
	return Deref(agentAccount.AccountID)
}

// GetAccountNumber возвращает Номер счета.
func (agentAccount AgentAccount) GetAccountNumber() string {
	return Deref(agentAccount.AccountNumber)
}

// GetBankLocation возвращает Адрес банка.
func (agentAccount AgentAccount) GetBankLocation() string {
	return Deref(agentAccount.BankLocation)
}

// GetBankName возвращает Наименование банка.
func (agentAccount AgentAccount) GetBankName() string {
	return Deref(agentAccount.BankName)
}

// GetBic возвращает БИК.
func (agentAccount AgentAccount) GetBic() string {
	return Deref(agentAccount.BIC)
}

// GetCorrespondentAccount возвращает Корр счет.
func (agentAccount AgentAccount) GetCorrespondentAccount() string {
	return Deref(agentAccount.CorrespondentAccount)
}

// GetID возвращает ID счета.
func (agentAccount AgentAccount) GetID() uuid.UUID {
	return Deref(agentAccount.ID)
}

// GetIsDefault возвращает true, если Счет является основным счетом Контрагента.
func (agentAccount AgentAccount) GetIsDefault() bool {
	return Deref(agentAccount.IsDefault)
}

// GetMeta возвращает Метаданные Счета Контрагента.
func (agentAccount AgentAccount) GetMeta() Meta {
	return Deref(agentAccount.Meta)
}

// GetUpdated возвращает Момент последнего обновления.
func (agentAccount AgentAccount) GetUpdated() Timestamp {
	return Deref(agentAccount.Updated)
}

// SetAccountNumber устанавливает Номер счета.
func (agentAccount *AgentAccount) SetAccountNumber(accountNumber string) *AgentAccount {
	agentAccount.AccountNumber = &accountNumber
	return agentAccount
}

// SetBankLocation устанавливает Адрес банка.
func (agentAccount *AgentAccount) SetBankLocation(bankLocation string) *AgentAccount {
	agentAccount.BankLocation = &bankLocation
	return agentAccount
}

// SetBankName устанавливает Наименование банка.
func (agentAccount *AgentAccount) SetBankName(bankName string) *AgentAccount {
	agentAccount.BankName = &bankName
	return agentAccount
}

// SetBic устанавливает БИК.
func (agentAccount *AgentAccount) SetBic(bic string) *AgentAccount {
	agentAccount.BIC = &bic
	return agentAccount
}

// SetCorrespondentAccount устанавливает Корр счет.
func (agentAccount *AgentAccount) SetCorrespondentAccount(correspondentAccount string) *AgentAccount {
	agentAccount.CorrespondentAccount = &correspondentAccount
	return agentAccount
}

// SetDefault устанавливает счет основным.
func (agentAccount *AgentAccount) SetDefault(value bool) *AgentAccount {
	agentAccount.IsDefault = &value
	return agentAccount
}

// SetMeta устанавливает Метаданные Счета Контрагента.
func (agentAccount *AgentAccount) SetMeta(meta *Meta) *AgentAccount {
	agentAccount.Meta = meta
	return agentAccount
}

// String реализует интерфейс [fmt.Stringer].
func (agentAccount AgentAccount) String() string {
	return Stringify(agentAccount)
}
