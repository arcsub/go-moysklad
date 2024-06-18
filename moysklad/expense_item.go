package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// ExpenseItem Статья расходов.
// Ключевое слово: expenseitem
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-stat-q-rashodow
type ExpenseItem struct {
	AccountID    *uuid.UUID `json:"accountId,omitempty"`    // ID учетной записи
	Code         *string    `json:"code,omitempty"`         // Код Статьи расходов
	Description  *string    `json:"description,omitempty"`  // Описание Статьи расходов
	ExternalCode *string    `json:"externalCode,omitempty"` // Внешний код Статьи расходов
	ID           *uuid.UUID `json:"id,omitempty"`           // ID Статьи расходов
	Meta         *Meta      `json:"meta,omitempty"`         // Метаданные о Статье расходов
	Name         *string    `json:"name,omitempty"`         // Наименование Статьи расходов
	Updated      *Timestamp `json:"updated,omitempty"`      // Момент последнего обновления сущности
}

// Clean возвращает сущность с единственным заполненным полем Meta
func (expenseItem ExpenseItem) Clean() *ExpenseItem {
	return &ExpenseItem{Meta: expenseItem.Meta}
}

func (expenseItem ExpenseItem) GetAccountID() uuid.UUID {
	return Deref(expenseItem.AccountID)
}

func (expenseItem ExpenseItem) GetCode() string {
	return Deref(expenseItem.Code)
}

func (expenseItem ExpenseItem) GetDescription() string {
	return Deref(expenseItem.Description)
}

func (expenseItem ExpenseItem) GetExternalCode() string {
	return Deref(expenseItem.ExternalCode)
}

func (expenseItem ExpenseItem) GetID() uuid.UUID {
	return Deref(expenseItem.ID)
}

func (expenseItem ExpenseItem) GetMeta() Meta {
	return Deref(expenseItem.Meta)
}

func (expenseItem ExpenseItem) GetName() string {
	return Deref(expenseItem.Name)
}

func (expenseItem *ExpenseItem) SetCode(code string) *ExpenseItem {
	expenseItem.Code = &code
	return expenseItem
}

func (expenseItem *ExpenseItem) SetDescription(description string) *ExpenseItem {
	expenseItem.Description = &description
	return expenseItem
}

func (expenseItem *ExpenseItem) SetExternalCode(externalCode string) *ExpenseItem {
	expenseItem.ExternalCode = &externalCode
	return expenseItem
}

func (expenseItem *ExpenseItem) SetMeta(meta *Meta) *ExpenseItem {
	expenseItem.Meta = meta
	return expenseItem
}

func (expenseItem *ExpenseItem) SetName(name string) *ExpenseItem {
	expenseItem.Name = &name
	return expenseItem
}

func (expenseItem ExpenseItem) GetUpdated() Timestamp {
	return Deref(expenseItem.Updated)
}

func (expenseItem ExpenseItem) String() string {
	return Stringify(expenseItem)
}

func (expenseItem ExpenseItem) MetaType() MetaType {
	return MetaTypeExpenseItem
}

// Update shortcut
func (expenseItem ExpenseItem) Update(ctx context.Context, client *Client, params ...*Params) (*ExpenseItem, *resty.Response, error) {
	return client.Entity().ExpenseItem().Update(ctx, expenseItem.GetID(), &expenseItem, params...)
}

// Create shortcut
func (expenseItem ExpenseItem) Create(ctx context.Context, client *Client, params ...*Params) (*ExpenseItem, *resty.Response, error) {
	return client.Entity().ExpenseItem().Create(ctx, &expenseItem, params...)
}

// Delete shortcut
func (expenseItem ExpenseItem) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return client.Entity().ExpenseItem().Delete(ctx, expenseItem.GetID())
}

// ExpenseItemService
// Сервис для работы со статьями расходов.
type ExpenseItemService interface {
	GetList(ctx context.Context, params ...*Params) (*List[ExpenseItem], *resty.Response, error)
	Create(ctx context.Context, expenseItem *ExpenseItem, params ...*Params) (*ExpenseItem, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, expenseItemList Slice[ExpenseItem], params ...*Params) (*Slice[ExpenseItem], *resty.Response, error)
	DeleteMany(ctx context.Context, entities ...*ExpenseItem) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*ExpenseItem, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, expenseItem *ExpenseItem, params ...*Params) (*ExpenseItem, *resty.Response, error)
	MoveToTrash(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
}

func NewExpenseItemService(client *Client) ExpenseItemService {
	e := NewEndpoint(client, "entity/expenseitem")
	return newMainService[ExpenseItem, any, any, any](e)
}
