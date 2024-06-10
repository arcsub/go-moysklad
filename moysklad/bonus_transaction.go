package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
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

func (bonusTransaction BonusTransaction) Clean() *BonusTransaction {
	return &BonusTransaction{Meta: bonusTransaction.Meta}
}

func (bonusTransaction BonusTransaction) GetOrganization() Organization {
	return Deref(bonusTransaction.Organization)
}

func (bonusTransaction BonusTransaction) GetUpdatedBy() string {
	return Deref(bonusTransaction.UpdatedBy)
}

func (bonusTransaction BonusTransaction) GetGroup() Group {
	return Deref(bonusTransaction.Group)
}

func (bonusTransaction BonusTransaction) GetBonusProgram() BonusProgram {
	return Deref(bonusTransaction.BonusProgram)
}

func (bonusTransaction BonusTransaction) GetBonusValue() int {
	return Deref(bonusTransaction.BonusValue)
}

func (bonusTransaction BonusTransaction) GetID() uuid.UUID {
	return Deref(bonusTransaction.ID)
}

func (bonusTransaction BonusTransaction) GetCode() string {
	return Deref(bonusTransaction.Code)
}

func (bonusTransaction BonusTransaction) GetCreated() Timestamp {
	return Deref(bonusTransaction.Created)
}

func (bonusTransaction BonusTransaction) GetExecutionDate() Timestamp {
	return Deref(bonusTransaction.ExecutionDate)
}

func (bonusTransaction BonusTransaction) GetMeta() Meta {
	return Deref(bonusTransaction.Meta)
}

func (bonusTransaction BonusTransaction) GetApplicable() bool {
	return Deref(bonusTransaction.Applicable)
}

func (bonusTransaction BonusTransaction) GetAgent() Counterparty {
	return Deref(bonusTransaction.Agent)
}

func (bonusTransaction BonusTransaction) GetExternalCode() string {
	return Deref(bonusTransaction.ExternalCode)
}

func (bonusTransaction BonusTransaction) GetMoment() Timestamp {
	return Deref(bonusTransaction.Moment)
}

func (bonusTransaction BonusTransaction) GetName() string {
	return Deref(bonusTransaction.Name)
}

func (bonusTransaction BonusTransaction) GetAccountID() uuid.UUID {
	return Deref(bonusTransaction.AccountID)
}

func (bonusTransaction BonusTransaction) GetOwner() Employee {
	return Deref(bonusTransaction.Owner)
}

func (bonusTransaction BonusTransaction) GetParentDocument() BonusTransaction {
	return Deref(bonusTransaction.ParentDocument)
}

func (bonusTransaction BonusTransaction) GetShared() bool {
	return Deref(bonusTransaction.Shared)
}

func (bonusTransaction BonusTransaction) GetUpdated() Timestamp {
	return Deref(bonusTransaction.Updated)
}

func (bonusTransaction BonusTransaction) GetTransactionType() BonusTransactionType {
	return bonusTransaction.TransactionType
}

func (bonusTransaction BonusTransaction) GetTransactionStatus() BonusTransactionStatus {
	return bonusTransaction.TransactionStatus
}

func (bonusTransaction BonusTransaction) GetCategoryType() BonusTransactionCategory {
	return bonusTransaction.CategoryType
}

func (bonusTransaction *BonusTransaction) SetOrganization(organization *Organization) *BonusTransaction {
	bonusTransaction.Organization = organization.Clean()
	return bonusTransaction
}

func (bonusTransaction *BonusTransaction) SetGroup(group *Group) *BonusTransaction {
	bonusTransaction.Group = group.Clean()
	return bonusTransaction
}

func (bonusTransaction *BonusTransaction) SetBonusProgram(bonusProgram *BonusProgram) *BonusTransaction {
	bonusTransaction.BonusProgram = bonusProgram.Clean()
	return bonusTransaction
}

func (bonusTransaction *BonusTransaction) SetBonusValue(bonusValue int) *BonusTransaction {
	bonusTransaction.BonusValue = &bonusValue
	return bonusTransaction
}

func (bonusTransaction *BonusTransaction) SetCode(code string) *BonusTransaction {
	bonusTransaction.Code = &code
	return bonusTransaction
}

func (bonusTransaction *BonusTransaction) SetCreated(created *Timestamp) *BonusTransaction {
	bonusTransaction.Created = created
	return bonusTransaction
}

func (bonusTransaction *BonusTransaction) SetMeta(meta *Meta) *BonusTransaction {
	bonusTransaction.Meta = meta
	return bonusTransaction
}

func (bonusTransaction *BonusTransaction) SetApplicable(applicable bool) *BonusTransaction {
	bonusTransaction.Applicable = &applicable
	return bonusTransaction
}

func (bonusTransaction *BonusTransaction) SetAgent(agent *Counterparty) *BonusTransaction {
	bonusTransaction.Agent = agent.Clean()
	return bonusTransaction
}

func (bonusTransaction *BonusTransaction) SetExternalCode(externalCode string) *BonusTransaction {
	bonusTransaction.ExternalCode = &externalCode
	return bonusTransaction
}

func (bonusTransaction *BonusTransaction) SetMoment(moment *Timestamp) *BonusTransaction {
	bonusTransaction.Moment = moment
	return bonusTransaction
}

func (bonusTransaction *BonusTransaction) SetName(name string) *BonusTransaction {
	bonusTransaction.Name = &name
	return bonusTransaction
}

func (bonusTransaction *BonusTransaction) SetOwner(owner *Employee) *BonusTransaction {
	bonusTransaction.Owner = owner.Clean()
	return bonusTransaction
}

func (bonusTransaction *BonusTransaction) SetParentDocument(parentDocument *BonusTransaction) *BonusTransaction {
	bonusTransaction.ParentDocument = parentDocument.Clean()
	return bonusTransaction
}

func (bonusTransaction *BonusTransaction) SetShared(shared bool) *BonusTransaction {
	bonusTransaction.Shared = &shared
	return bonusTransaction
}

func (bonusTransaction *BonusTransaction) SetTransactionType(transactionType BonusTransactionType) *BonusTransaction {
	bonusTransaction.TransactionType = transactionType
	return bonusTransaction
}

func (bonusTransaction BonusTransaction) String() string {
	return Stringify(bonusTransaction)
}

func (bonusTransaction BonusTransaction) MetaType() MetaType {
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

// BonusTransactionService
// Сервис для работы с бонусными операциями.
type BonusTransactionService interface {
	GetList(ctx context.Context, params *Params) (*List[BonusTransaction], *resty.Response, error)
	Create(ctx context.Context, bonusTransaction *BonusTransaction, params *Params) (*BonusTransaction, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, bonusTransactionList []*BonusTransaction, params *Params) (*[]BonusTransaction, *resty.Response, error)
	DeleteMany(ctx context.Context, bonusTransactionList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params *Params) (*BonusTransaction, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, bonusTransaction *BonusTransaction, params *Params) (*BonusTransaction, *resty.Response, error)
}

func NewBonusTransactionService(client *Client) BonusTransactionService {
	e := NewEndpoint(client, "entity/bonustransaction")
	return newMainService[BonusTransaction, any, any, any](e)
}
