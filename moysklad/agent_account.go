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

func (a AgentAccount) String() string {
	return Stringify(a)
}

type AgentAccounts = Slice[AgentAccount]
