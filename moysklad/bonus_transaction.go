package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// BonusTransaction Бонусная операция.
//
// Ключевое слово: bonustransaction
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-bonusnaq-operaciq
type BonusTransaction struct {
	Organization      *Organization            `json:"organization,omitempty"`      // Метаданные юрлица
	Group             *Group                   `json:"group,omitempty"`             // Метаданные Отдела сотрудника
	BonusProgram      *BonusProgram            `json:"bonusProgram,omitempty"`      // Метаданные бонусной программы
	BonusValue        *int                     `json:"bonusValue,omitempty"`        // Количество бонусных баллов
	ID                *uuid.UUID               `json:"id,omitempty"`                // ID Бонусной операции
	Code              *string                  `json:"code,omitempty"`              // Код Бонусной операции
	Created           *Timestamp               `json:"created,omitempty"`           // Момент создания Бонусной операции
	ExecutionDate     *Timestamp               `json:"executionDate,omitempty"`     // Дата начисления бонусной операции
	Meta              *Meta                    `json:"meta,omitempty"`              // Метаданные Бонусной операции
	Applicable        *bool                    `json:"applicable,omitempty"`        // Отметка о проведении
	Agent             *Counterparty            `json:"agent,omitempty"`             // Метаданные Контрагента, связанного с бонусной операцией
	ExternalCode      *string                  `json:"externalCode,omitempty"`      // Внешний код Бонусной операции
	Moment            *Timestamp               `json:"moment,omitempty"`            // Время проведения бонусной операции
	Name              *string                  `json:"name,omitempty"`              // Наименование Бонусной операции
	AccountID         *uuid.UUID               `json:"accountId,omitempty"`         // ID учетной записи
	Owner             *Employee                `json:"owner,omitempty"`             // Метаданные владельца (Сотрудника)
	ParentDocument    *BonusTransaction        `json:"parentDocument,omitempty"`    // Метаданные связанного документа бонусной операции
	Shared            *bool                    `json:"shared,omitempty"`            // Общий доступ
	Updated           *Timestamp               `json:"updated,omitempty"`           // Момент последнего обновления Бонусной операции
	TransactionType   BonusTransactionType     `json:"transactionType,omitempty"`   // Тип бонусной операции. Возможные значения: EARNING, SPENDING
	TransactionStatus BonusTransactionStatus   `json:"transactionStatus,omitempty"` // Статус бонусной операции. Возможные значения: WAIT_PROCESSING, COMPLETED, CANCELED
	CategoryType      BonusTransactionCategory `json:"categoryType,omitempty"`      // Категория бонусной операции. Возможные значения: REGULAR, WELCOME
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (bonusTransaction BonusTransaction) Clean() *BonusTransaction {
	return &BonusTransaction{Meta: bonusTransaction.Meta}
}

// AsTaskOperation реализует интерфейс AsTaskOperationInterface
func (bonusTransaction BonusTransaction) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: bonusTransaction.Meta}
}

// GetOrganization возвращает Метаданные юрлица.
func (bonusTransaction BonusTransaction) GetOrganization() Organization {
	return Deref(bonusTransaction.Organization)
}

// GetGroup возвращает Метаданные Отдела сотрудника.
func (bonusTransaction BonusTransaction) GetGroup() Group {
	return Deref(bonusTransaction.Group)
}

// GetBonusProgram возвращает Метаданные бонусной программы.
func (bonusTransaction BonusTransaction) GetBonusProgram() BonusProgram {
	return Deref(bonusTransaction.BonusProgram)
}

// GetBonusValue возвращает Количество бонусных баллов.
func (bonusTransaction BonusTransaction) GetBonusValue() int {
	return Deref(bonusTransaction.BonusValue)
}

// GetID возвращает ID Бонусной операции.
func (bonusTransaction BonusTransaction) GetID() uuid.UUID {
	return Deref(bonusTransaction.ID)
}

// GetCode возвращает Код Бонусной операции.
func (bonusTransaction BonusTransaction) GetCode() string {
	return Deref(bonusTransaction.Code)
}

// GetCreated возвращает Момент создания Бонусной операции.
func (bonusTransaction BonusTransaction) GetCreated() Timestamp {
	return Deref(bonusTransaction.Created)
}

// GetExecutionDate возвращает Дату начисления бонусной операции.
func (bonusTransaction BonusTransaction) GetExecutionDate() Timestamp {
	return Deref(bonusTransaction.ExecutionDate)
}

// GetMeta возвращает Метаданные Бонусной операции.
func (bonusTransaction BonusTransaction) GetMeta() Meta {
	return Deref(bonusTransaction.Meta)
}

// GetApplicable возвращает Отметку о проведении.
func (bonusTransaction BonusTransaction) GetApplicable() bool {
	return Deref(bonusTransaction.Applicable)
}

// GetAgent возвращает Метаданные Контрагента, связанного с бонусной операцией.
func (bonusTransaction BonusTransaction) GetAgent() Counterparty {
	return Deref(bonusTransaction.Agent)
}

// GetExternalCode возвращает Внешний код Бонусной операции.
func (bonusTransaction BonusTransaction) GetExternalCode() string {
	return Deref(bonusTransaction.ExternalCode)
}

// GetMoment возвращает Время проведения бонусной операции.
func (bonusTransaction BonusTransaction) GetMoment() Timestamp {
	return Deref(bonusTransaction.Moment)
}

// GetName возвращает Наименование Бонусной операции.
func (bonusTransaction BonusTransaction) GetName() string {
	return Deref(bonusTransaction.Name)
}

// GetAccountID возвращает ID учетной записи.
func (bonusTransaction BonusTransaction) GetAccountID() uuid.UUID {
	return Deref(bonusTransaction.AccountID)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (bonusTransaction BonusTransaction) GetOwner() Employee {
	return Deref(bonusTransaction.Owner)
}

// GetParentDocument возвращает Метаданные связанного документа бонусной операции.
func (bonusTransaction BonusTransaction) GetParentDocument() BonusTransaction {
	return Deref(bonusTransaction.ParentDocument)
}

// GetShared возвращает флаг общего доступа.
func (bonusTransaction BonusTransaction) GetShared() bool {
	return Deref(bonusTransaction.Shared)
}

// GetUpdated возвращает Момент последнего обновления Бонусной операции.
func (bonusTransaction BonusTransaction) GetUpdated() Timestamp {
	return Deref(bonusTransaction.Updated)
}

// GetTransactionType возвращает Тип бонусной операции.
func (bonusTransaction BonusTransaction) GetTransactionType() BonusTransactionType {
	return bonusTransaction.TransactionType
}

// GetTransactionStatus возвращает Статус бонусной операции.
func (bonusTransaction BonusTransaction) GetTransactionStatus() BonusTransactionStatus {
	return bonusTransaction.TransactionStatus
}

// GetCategoryType возвращает Категорию бонусной операции.
func (bonusTransaction BonusTransaction) GetCategoryType() BonusTransactionCategory {
	return bonusTransaction.CategoryType
}

// SetOrganization устанавливает Метаданные юрлица.
func (bonusTransaction *BonusTransaction) SetOrganization(organization *Organization) *BonusTransaction {
	bonusTransaction.Organization = organization.Clean()
	return bonusTransaction
}

// SetGroup устанавливает Метаданные Отдела сотрудника.
func (bonusTransaction *BonusTransaction) SetGroup(group *Group) *BonusTransaction {
	bonusTransaction.Group = group.Clean()
	return bonusTransaction
}

// SetBonusProgram устанавливает Метаданные бонусной программы.
func (bonusTransaction *BonusTransaction) SetBonusProgram(bonusProgram *BonusProgram) *BonusTransaction {
	bonusTransaction.BonusProgram = bonusProgram.Clean()
	return bonusTransaction
}

// SetBonusValue устанавливает Количество бонусных баллов.
func (bonusTransaction *BonusTransaction) SetBonusValue(bonusValue int) *BonusTransaction {
	bonusTransaction.BonusValue = &bonusValue
	return bonusTransaction
}

// SetCode устанавливает Код Бонусной операции.
func (bonusTransaction *BonusTransaction) SetCode(code string) *BonusTransaction {
	bonusTransaction.Code = &code
	return bonusTransaction
}

// SetMeta устанавливает Метаданные Бонусной операции.
func (bonusTransaction *BonusTransaction) SetMeta(meta *Meta) *BonusTransaction {
	bonusTransaction.Meta = meta
	return bonusTransaction
}

// SetApplicable устанавливает флаг о проведении.
func (bonusTransaction *BonusTransaction) SetApplicable(applicable bool) *BonusTransaction {
	bonusTransaction.Applicable = &applicable
	return bonusTransaction
}

// SetAgent устанавливает Метаданные Контрагента, связанного с бонусной операцией.
func (bonusTransaction *BonusTransaction) SetAgent(agent *Counterparty) *BonusTransaction {
	bonusTransaction.Agent = agent.Clean()
	return bonusTransaction
}

// SetExternalCode устанавливает Внешний код Бонусной операции.
func (bonusTransaction *BonusTransaction) SetExternalCode(externalCode string) *BonusTransaction {
	bonusTransaction.ExternalCode = &externalCode
	return bonusTransaction
}

// SetMoment устанавливает Время проведения бонусной операции.
func (bonusTransaction *BonusTransaction) SetMoment(moment *Timestamp) *BonusTransaction {
	bonusTransaction.Moment = moment
	return bonusTransaction
}

// SetName устанавливает Наименование Бонусной операции.
func (bonusTransaction *BonusTransaction) SetName(name string) *BonusTransaction {
	bonusTransaction.Name = &name
	return bonusTransaction
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (bonusTransaction *BonusTransaction) SetOwner(owner *Employee) *BonusTransaction {
	bonusTransaction.Owner = owner.Clean()
	return bonusTransaction
}

// SetParentDocument устанавливает Метаданные связанного документа бонусной операции.
func (bonusTransaction *BonusTransaction) SetParentDocument(parentDocument *BonusTransaction) *BonusTransaction {
	bonusTransaction.ParentDocument = parentDocument.Clean()
	return bonusTransaction
}

// SetShared устанавливает флаг общего доступа.
func (bonusTransaction *BonusTransaction) SetShared(shared bool) *BonusTransaction {
	bonusTransaction.Shared = &shared
	return bonusTransaction
}

// SetTransactionType устанавливает Тип бонусной операции.
func (bonusTransaction *BonusTransaction) SetTransactionType(transactionType BonusTransactionType) *BonusTransaction {
	bonusTransaction.TransactionType = transactionType
	return bonusTransaction
}

// SetTransactionTypeEarning устанавливает Тип бонусной операции [Earning] (Начисление).
func (bonusTransaction *BonusTransaction) SetTransactionTypeEarning() *BonusTransaction {
	bonusTransaction.TransactionType = Earning
	return bonusTransaction
}

// SetTransactionTypeSpending устанавливает Тип бонусной операции [Spending] (Списание).
func (bonusTransaction *BonusTransaction) SetTransactionTypeSpending() *BonusTransaction {
	bonusTransaction.TransactionType = Spending
	return bonusTransaction
}

// String реализует интерфейс [fmt.Stringer].
func (bonusTransaction BonusTransaction) String() string {
	return Stringify(bonusTransaction)
}

// MetaType возвращает тип сущности.
func (BonusTransaction) MetaType() MetaType {
	return MetaTypeBonusTransaction
}

// Update shortcut
func (bonusTransaction BonusTransaction) Update(ctx context.Context, client *Client, params ...*Params) (*BonusTransaction, *resty.Response, error) {
	return NewBonusTransactionService(client).Update(ctx, bonusTransaction.GetID(), &bonusTransaction, params...)
}

// Create shortcut
func (bonusTransaction BonusTransaction) Create(ctx context.Context, client *Client, params ...*Params) (*BonusTransaction, *resty.Response, error) {
	return NewBonusTransactionService(client).Create(ctx, &bonusTransaction, params...)
}

// Delete shortcut
func (bonusTransaction BonusTransaction) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewBonusTransactionService(client).Delete(ctx, bonusTransaction.GetID())
}

// BonusTransactionCategory Категория бонусной операции.
//
// Возможные значения:
//   - BonusTransactionCategoryTypeRegular – Стандартная операция
//   - BonusTransactionCategoryTypeWelcome – Начисление приветственных баллов
type BonusTransactionCategory string

const (
	BonusTransactionCategoryTypeRegular BonusTransactionCategory = "REGULAR" // Стандартная операция
	BonusTransactionCategoryTypeWelcome BonusTransactionCategory = "WELCOME" // Начисление приветственных баллов
)

// BonusTransactionStatus Статус бонусной операции.
//
// Возможные значения:
//   - BonusTransactionStatusWaitProcessing – В процессе ожидания начисления
//   - BonusTransactionStatusCompleted      – Завершена
//   - BonusTransactionStatusCanceled       – Отменена
type BonusTransactionStatus string

const (
	BonusTransactionStatusWaitProcessing BonusTransactionStatus = "WAIT_PROCESSING" // В процессе ожидания начисления
	BonusTransactionStatusCompleted      BonusTransactionStatus = "COMPLETED"       // Завершена
	BonusTransactionStatusCanceled       BonusTransactionStatus = "CANCELED"        // Отменена
)

// BonusTransactionType Тип бонусной операции.
//
// Возможные значения:
//   - Earning  – Начисление баллов
//   - Spending – Списание баллов
type BonusTransactionType string

const (
	Earning  BonusTransactionType = "EARNING"  // Начисление баллов
	Spending BonusTransactionType = "SPENDING" // Списание баллов
)

// BonusTransactionService Сервис для работы с бонусными операциями.
type BonusTransactionService interface {
	// GetList выполняет запрос на получение списка бонусных операций.
	// Принимает контекст context.Context и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[BonusTransaction], *resty.Response, error)

	// Create выполняет запрос на создание бонусной операции.
	// Обязательные поля для заполнения:
	//	- agent (Метаданные Контрагента, связанного с бонусной операцией)
	//	- bonusProgram (Метаданные Бонусной программы)
	//	- transactionType (Тип бонусной операции)
	// Принимает контекст context.Context, бонусную операцию и опционально объект параметров запроса Params.
	// Возвращает созданную бонусную операцию.
	Create(ctx context.Context, bonusTransaction *BonusTransaction, params ...*Params) (*BonusTransaction, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и обновление бонусных операций.
	// Обновляемые Бонусные операции должны содержать идентификатор в виде метаданных.
	// Принимает контекст context.Context, множество бонусных операций и опционально объект параметров запроса Params.
	// Возвращает множество созданных и/или обновлённых бонусных операций.
	CreateUpdateMany(ctx context.Context, bonusTransactionList Slice[BonusTransaction], params ...*Params) (*Slice[BonusTransaction], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление бонусных операций.
	// Принимает контекст context.Context и множество бонусных операций.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*BonusTransaction) (*DeleteManyResponse, *resty.Response, error)

	// Delete выполняет запрос на удаление бонусной операции.
	// Принимает контекст context.Context и ID бонусной операции.
	// Возвращает true в случае успешного удаления бонусной операции.
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение бонусной операции.
	// Принимает контекст context.Context, ID бонусной операции и опционально объект параметров запроса Params.
	// Возвращает бонусную операцию.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*BonusTransaction, *resty.Response, error)

	// Update выполняет запрос на обновление бонусной операции.
	// Принимает контекст context.Context, бонусную операцию и опционально объект параметров запроса Params.
	// Возвращает обновлённую бонусную операцию.
	Update(ctx context.Context, id uuid.UUID, bonusTransaction *BonusTransaction, params ...*Params) (*BonusTransaction, *resty.Response, error)
}

// NewBonusTransactionService возвращает сервис для работы с бонусными операциями.
func NewBonusTransactionService(client *Client) BonusTransactionService {
	return newMainService[BonusTransaction, any, any, any](NewEndpoint(client, "entity/bonustransaction"))
}
