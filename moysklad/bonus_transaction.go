package moysklad

import (
	"github.com/google/uuid"
)

// BonusTransaction Бонусная операция.
// Ключевое слово: bonustransaction
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-bonusnaq-operaciq
type BonusTransaction struct {
	AccountID         *uuid.UUID               `json:"accountId,omitempty"`         // ID учетной записи
	Agent             *Counterparty            `json:"agent,omitempty"`             // Метаданные Контрагента, связанного с бонусной операцией
	Applicable        *bool                    `json:"applicable,omitempty"`        // Отметка о проведении
	BonusProgram      *BonusProgram            `json:"bonusProgram,omitempty"`      // Метаданные бонусной программы
	BonusValue        *int                     `json:"bonusValue,omitempty"`        // Количество бонусных баллов
	CategoryType      BonusTransactionCategory `json:"categoryType,omitempty"`      // Категория бонусной операции
	Code              *string                  `json:"code,omitempty"`              // Код Бонусной операции
	Created           *Timestamp               `json:"created,omitempty"`           // Момент создания Бонусной операции
	ExecutionDate     *Timestamp               `json:"executionDate,omitempty"`     // Дата начисления бонусной операции.
	ExternalCode      *string                  `json:"externalCode,omitempty"`      // Внешний код Бонусной операции
	Group             *Group                   `json:"group,omitempty"`             // Отдел сотрудника
	ID                *uuid.UUID               `json:"id,omitempty"`                // ID сущности
	Meta              *Meta                    `json:"meta,omitempty"`              // Метаданные
	Moment            *Timestamp               `json:"moment,omitempty"`            // Время проведения бонусной операции
	Name              *string                  `json:"name,omitempty"`              // Наименование
	Organization      *Organization            `json:"organization,omitempty"`      // Метаданные юрлица
	Owner             *Employee                `json:"owner,omitempty"`             // Владелец (Сотрудник)
	ParentDocument    *BonusTransaction        `json:"parentDocument,omitempty"`    // Метаданные связанного документа бонусной операции
	Shared            *bool                    `json:"shared,omitempty"`            // Общий доступ
	TransactionStatus BonusTransactionStatus   `json:"transactionStatus,omitempty"` // Статус бонусной операции. Возможные значения: WAIT_PROCESSING, COMPLETED, CANCELED
	TransactionType   BonusTransactionType     `json:"transactionType,omitempty"`   // Тип бонусной операции. Возможные значения: EARNING, SPENDING
	Updated           *Timestamp               `json:"updated,omitempty"`           // Момент последнего обновления Бонусной операции
	UpdatedBy         *string                  `json:"updatedBy,omitempty"`         // Автор последнего обновления бонусной операции в формате uid (admin@admin) (Атрибут используется только для фильтрации)
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
