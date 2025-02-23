package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"

	"time"
)

// BonusTransaction Бонусная операция.
//
// Код сущности: bonustransaction
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-bonusnaq-operaciq
type BonusTransaction struct {
	Organization      *Organization            `json:"organization,omitempty"`      // Метаданные юрлица
	Group             *Group                   `json:"group,omitempty"`             // Отдел сотрудника
	BonusProgram      *BonusProgram            `json:"bonusProgram,omitempty"`      // Метаданные бонусной программы
	BonusValue        *int                     `json:"bonusValue,omitempty"`        // Количество бонусных баллов
	ID                *string                  `json:"id,omitempty"`                // ID Бонусной операции
	Code              *string                  `json:"code,omitempty"`              // Код Бонусной операции
	Created           *Timestamp               `json:"created,omitempty"`           // Момент создания Бонусной операции
	ExecutionDate     *Timestamp               `json:"executionDate,omitempty"`     // Дата начисления бонусной операции
	Meta              *Meta                    `json:"meta,omitempty"`              // Метаданные Бонусной операции
	Applicable        *bool                    `json:"applicable,omitempty"`        // Отметка о проведении
	Agent             *Agent                   `json:"agent,omitempty"`             // Метаданные Контрагента, связанного с бонусной операцией
	ExternalCode      *string                  `json:"externalCode,omitempty"`      // Внешний код Бонусной операции
	Moment            *Timestamp               `json:"moment,omitempty"`            // Время проведения бонусной операции
	Name              *string                  `json:"name,omitempty"`              // Наименование Бонусной операции
	AccountID         *string                  `json:"accountId,omitempty"`         // ID учётной записи
	Owner             *Employee                `json:"owner,omitempty"`             // Метаданные владельца (Сотрудника)
	ParentDocument    *BonusTransaction        `json:"parentDocument,omitempty"`    // Метаданные связанного документа бонусной операции
	Shared            *bool                    `json:"shared,omitempty"`            // Общий доступ
	Updated           *Timestamp               `json:"updated,omitempty"`           // Момент последнего обновления Бонусной операции
	TransactionType   BonusTransactionType     `json:"transactionType,omitempty"`   // Тип бонусной операции. Возможные значения: EARNING, SPENDING
	TransactionStatus BonusTransactionStatus   `json:"transactionStatus,omitempty"` // Статус бонусной операции. Возможные значения: WAIT_PROCESSING, COMPLETED, CANCELED
	CategoryType      BonusTransactionCategory `json:"categoryType,omitempty"`      // Категория бонусной операции. Возможные значения: REGULAR, WELCOME
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (bonusTransaction BonusTransaction) Clean() *BonusTransaction {
	if bonusTransaction.Meta == nil {
		return nil
	}
	return &BonusTransaction{Meta: bonusTransaction.Meta}
}

// AsTaskOperation реализует интерфейс [TaskOperationConverter].
func (bonusTransaction BonusTransaction) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: bonusTransaction.Meta}
}

// GetOrganization возвращает Метаданные юрлица.
func (bonusTransaction BonusTransaction) GetOrganization() Organization {
	return Deref(bonusTransaction.Organization)
}

// GetGroup возвращает Отдел сотрудника.
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
func (bonusTransaction BonusTransaction) GetID() string {
	return Deref(bonusTransaction.ID)
}

// GetCode возвращает Код Бонусной операции.
func (bonusTransaction BonusTransaction) GetCode() string {
	return Deref(bonusTransaction.Code)
}

// GetCreated возвращает Дату создания.
func (bonusTransaction BonusTransaction) GetCreated() time.Time {
	return Deref(bonusTransaction.Created).Time()
}

// GetExecutionDate возвращает Дату начисления бонусной операции.
func (bonusTransaction BonusTransaction) GetExecutionDate() time.Time {
	return Deref(bonusTransaction.ExecutionDate).Time()
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
func (bonusTransaction BonusTransaction) GetAgent() Agent {
	return Deref(bonusTransaction.Agent)
}

// GetExternalCode возвращает Внешний код Бонусной операции.
func (bonusTransaction BonusTransaction) GetExternalCode() string {
	return Deref(bonusTransaction.ExternalCode)
}

// GetMoment возвращает Время проведения бонусной операции.
func (bonusTransaction BonusTransaction) GetMoment() time.Time {
	return Deref(bonusTransaction.Moment).Time()
}

// GetName возвращает Наименование Бонусной операции.
func (bonusTransaction BonusTransaction) GetName() string {
	return Deref(bonusTransaction.Name)
}

// GetAccountID возвращает ID учётной записи.
func (bonusTransaction BonusTransaction) GetAccountID() string {
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
func (bonusTransaction BonusTransaction) GetUpdated() time.Time {
	return Deref(bonusTransaction.Updated).Time()
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
	if organization != nil {
		bonusTransaction.Organization = organization.Clean()
	}
	return bonusTransaction
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (bonusTransaction *BonusTransaction) SetGroup(group *Group) *BonusTransaction {
	if group != nil {
		bonusTransaction.Group = group.Clean()
	}
	return bonusTransaction
}

// SetBonusProgram устанавливает Метаданные бонусной программы.
func (bonusTransaction *BonusTransaction) SetBonusProgram(bonusProgram *BonusProgram) *BonusTransaction {
	if bonusProgram != nil {
		bonusTransaction.BonusProgram = bonusProgram.Clean()
	}
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

// SetApplicable устанавливает Отметку о проведении.
func (bonusTransaction *BonusTransaction) SetApplicable(applicable bool) *BonusTransaction {
	bonusTransaction.Applicable = &applicable
	return bonusTransaction
}

// SetAgent устанавливает Метаданные Контрагента, связанного с бонусной операцией.
//
// Принимает [Counterparty] или [Organization].
func (bonusTransaction *BonusTransaction) SetAgent(agent AgentOrganizationConverter) *BonusTransaction {
	if agent != nil {
		bonusTransaction.Agent = agent.AsOrganizationAgent()
	}
	return bonusTransaction
}

// SetExternalCode устанавливает Внешний код Бонусной операции.
func (bonusTransaction *BonusTransaction) SetExternalCode(externalCode string) *BonusTransaction {
	bonusTransaction.ExternalCode = &externalCode
	return bonusTransaction
}

// SetMoment устанавливает Время проведения бонусной операции.
func (bonusTransaction *BonusTransaction) SetMoment(moment time.Time) *BonusTransaction {
	bonusTransaction.Moment = NewTimestamp(moment)
	return bonusTransaction
}

// SetName устанавливает Наименование Бонусной операции.
func (bonusTransaction *BonusTransaction) SetName(name string) *BonusTransaction {
	bonusTransaction.Name = &name
	return bonusTransaction
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (bonusTransaction *BonusTransaction) SetOwner(owner *Employee) *BonusTransaction {
	if owner != nil {
		bonusTransaction.Owner = owner.Clean()
	}
	return bonusTransaction
}

// SetParentDocument устанавливает Метаданные связанного документа бонусной операции.
func (bonusTransaction *BonusTransaction) SetParentDocument(parentDocument *BonusTransaction) *BonusTransaction {
	if parentDocument != nil {
		bonusTransaction.ParentDocument = parentDocument.Clean()
	}
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

// MetaType возвращает код сущности.
func (BonusTransaction) MetaType() MetaType {
	return MetaTypeBonusTransaction
}

// Update shortcut
func (bonusTransaction *BonusTransaction) Update(ctx context.Context, client *Client, params ...func(*Params)) (*BonusTransaction, *resty.Response, error) {
	return NewBonusTransactionService(client).Update(ctx, bonusTransaction.GetID(), bonusTransaction, params...)
}

// Create shortcut
func (bonusTransaction *BonusTransaction) Create(ctx context.Context, client *Client, params ...func(*Params)) (*BonusTransaction, *resty.Response, error) {
	return NewBonusTransactionService(client).Create(ctx, bonusTransaction, params...)
}

// Delete shortcut
func (bonusTransaction *BonusTransaction) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewBonusTransactionService(client).Delete(ctx, bonusTransaction)
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

// BonusTransactionService методы сервиса для работы с бонусными операциями.
type BonusTransactionService interface {
	// GetList выполняет запрос на получение списка бонусных операций.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...func(*Params)) (*List[BonusTransaction], *resty.Response, error)

	// GetListAll выполняет запрос на получение всех бонусных операций в виде списка.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает список объектов.
	GetListAll(ctx context.Context, params ...func(*Params)) (*Slice[BonusTransaction], *resty.Response, error)

	// Create выполняет запрос на создание бонусной операции.
	// Обязательные поля для заполнения:
	//	- agent (Метаданные Контрагента, связанного с бонусной операцией)
	//	- bonusProgram (Метаданные Бонусной программы)
	//	- transactionType (Тип бонусной операции)
	// Принимает контекст, бонусную операцию и опционально объект параметров запроса Params.
	// Возвращает созданную бонусную операцию.
	Create(ctx context.Context, bonusTransaction *BonusTransaction, params ...func(*Params)) (*BonusTransaction, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение бонусных операций.
	// Изменяемые Бонусные операции должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список бонусных операций и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых бонусных операций.
	CreateUpdateMany(ctx context.Context, bonusTransactionList Slice[BonusTransaction], params ...func(*Params)) (*Slice[BonusTransaction], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление бонусных операций.
	// Принимает контекст и множество бонусных операций.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*BonusTransaction) (*DeleteManyResponse, *resty.Response, error)

	// DeleteByID выполняет запрос на удаление бонусной операции по ID.
	// Принимает контекст и ID бонусной операции.
	// Возвращает «true» в случае успешного удаления бонусной операции.
	DeleteByID(ctx context.Context, id string) (bool, *resty.Response, error)

	// Delete выполняет запрос на удаление бонусной операции.
	// Принимает контекст и бонусную операцию.
	// Возвращает «true» в случае успешного удаления бонусной операции.
	Delete(ctx context.Context, entity *BonusTransaction) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение бонусной операции.
	// Принимает контекст, ID бонусной операции и опционально объект параметров запроса Params.
	// Возвращает бонусную операцию.
	GetByID(ctx context.Context, id string, params ...func(*Params)) (*BonusTransaction, *resty.Response, error)

	// Update выполняет запрос на изменение бонусной операции.
	// Принимает контекст, бонусную операцию и опционально объект параметров запроса Params.
	// Возвращает изменённую бонусную операцию.
	Update(ctx context.Context, id string, bonusTransaction *BonusTransaction, params ...func(*Params)) (*BonusTransaction, *resty.Response, error)
}

const (
	EndpointBonusTransaction = EndpointEntity + string(MetaTypeBonusTransaction)
)

// NewBonusTransactionService принимает [Client] и возвращает сервис для работы с бонусными операциями.
func NewBonusTransactionService(client *Client) BonusTransactionService {
	return newMainService[BonusTransaction, any, any, any](client, EndpointBonusTransaction)
}
