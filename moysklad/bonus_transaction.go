package moysklad

import (
	"github.com/google/uuid"
)

// BonusTransaction Бонусная операция.
// Ключевое слово: bonustransaction
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-bonusnaq-operaciq
type BonusTransaction struct {
	Organization      *Organization            `json:"organization,omitempty"`
	UpdatedBy         *string                  `json:"updatedBy,omitempty"`
	Group             *Group                   `json:"group,omitempty"`
	BonusProgram      *BonusProgram            `json:"bonusProgram,omitempty"`
	BonusValue        *int                     `json:"bonusValue,omitempty"`
	ID                *uuid.UUID               `json:"id,omitempty"`
	Code              *string                  `json:"code,omitempty"`
	Created           *Timestamp               `json:"created,omitempty"`
	ExecutionDate     *Timestamp               `json:"executionDate,omitempty"`
	Meta              *Meta                    `json:"meta,omitempty"`
	Applicable        *bool                    `json:"applicable,omitempty"`
	Agent             *Counterparty            `json:"agent,omitempty"`
	ExternalCode      *string                  `json:"externalCode,omitempty"`
	Moment            *Timestamp               `json:"moment,omitempty"`
	Name              *string                  `json:"name,omitempty"`
	AccountID         *uuid.UUID               `json:"accountId,omitempty"`
	Owner             *Employee                `json:"owner,omitempty"`
	ParentDocument    *BonusTransaction        `json:"parentDocument,omitempty"`
	Shared            *bool                    `json:"shared,omitempty"`
	Updated           *Timestamp               `json:"updated,omitempty"`
	TransactionType   BonusTransactionType     `json:"transactionType,omitempty"`
	TransactionStatus BonusTransactionStatus   `json:"transactionStatus,omitempty"`
	CategoryType      BonusTransactionCategory `json:"categoryType,omitempty"`
}

func (b BonusTransaction) String() string {
	return Stringify(b)
}

func (b BonusTransaction) MetaType() MetaType {
	return MetaTypeBonusTransaction
}

// BonusTransactionCategory Категория бонусной операции
type BonusTransactionCategory string

const (
	BonusTransactionCategoryTypeRegular BonusTransactionCategory = "REGULAR"
	BonusTransactionCategoryTypeWelcome BonusTransactionCategory = "WELCOME"
)

// BonusTransactionStatus Статус бонусной операции
type BonusTransactionStatus string

const (
	BonusTransactionStatusWaitProcessing BonusTransactionStatus = "WAIT_PROCESSING"
	BonusTransactionStatusCompleted      BonusTransactionStatus = "COMPLETED"
	BonusTransactionStatusCanceled       BonusTransactionStatus = "CANCELED"
)

// BonusTransactionType Тип бонусной операции
type BonusTransactionType string

const (
	Earning  BonusTransactionType = "EARNING"
	Spending BonusTransactionType = "SPENDING"
)
