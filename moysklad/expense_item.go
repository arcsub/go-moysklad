package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"time"
)

// ExpenseItem Статья расходов.
//
// Код сущности: expenseitem
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-stat-q-rashodow
type ExpenseItem struct {
	AccountID    *uuid.UUID `json:"accountId,omitempty"`    // ID учётной записи
	Code         *string    `json:"code,omitempty"`         // Код Статьи расходов
	Description  *string    `json:"description,omitempty"`  // Описание Статьи расходов
	ExternalCode *string    `json:"externalCode,omitempty"` // Внешний код Статьи расходов
	ID           *uuid.UUID `json:"id,omitempty"`           // ID Статьи расходов
	Meta         *Meta      `json:"meta,omitempty"`         // Метаданные о Статье расходов
	Name         *string    `json:"name,omitempty"`         // Наименование Статьи расходов
	Updated      *Timestamp `json:"updated,omitempty"`      // Момент последнего обновления сущности
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (expenseItem ExpenseItem) Clean() *ExpenseItem {
	if expenseItem.Meta == nil {
		return nil
	}
	return &ExpenseItem{Meta: expenseItem.Meta}
}

// GetAccountID возвращает ID учётной записи.
func (expenseItem ExpenseItem) GetAccountID() uuid.UUID {
	return Deref(expenseItem.AccountID)
}

// GetCode возвращает Код Статьи расходов.
func (expenseItem ExpenseItem) GetCode() string {
	return Deref(expenseItem.Code)
}

// GetDescription возвращает Описание Статьи расходов.
func (expenseItem ExpenseItem) GetDescription() string {
	return Deref(expenseItem.Description)
}

// GetExternalCode возвращает Внешний код Статьи расходов.
func (expenseItem ExpenseItem) GetExternalCode() string {
	return Deref(expenseItem.ExternalCode)
}

// GetID возвращает ID Статьи расходов.
func (expenseItem ExpenseItem) GetID() uuid.UUID {
	return Deref(expenseItem.ID)
}

// GetMeta возвращает Метаданные Статьи расходов.
func (expenseItem ExpenseItem) GetMeta() Meta {
	return Deref(expenseItem.Meta)
}

// GetName возвращает Наименование Статьи расходов.
func (expenseItem ExpenseItem) GetName() string {
	return Deref(expenseItem.Name)
}

// GetUpdated возвращает Момент последнего обновления сущности.
func (expenseItem ExpenseItem) GetUpdated() time.Time {
	return Deref(expenseItem.Updated).Time()
}

// SetCode устанавливает Код Статьи расходов.
func (expenseItem *ExpenseItem) SetCode(code string) *ExpenseItem {
	expenseItem.Code = &code
	return expenseItem
}

// SetDescription устанавливает Описание Статьи расходов.
func (expenseItem *ExpenseItem) SetDescription(description string) *ExpenseItem {
	expenseItem.Description = &description
	return expenseItem
}

// SetExternalCode устанавливает Внешний код Статьи расходов.
func (expenseItem *ExpenseItem) SetExternalCode(externalCode string) *ExpenseItem {
	expenseItem.ExternalCode = &externalCode
	return expenseItem
}

// SetMeta устанавливает Метаданные Статьи расходов.
func (expenseItem *ExpenseItem) SetMeta(meta *Meta) *ExpenseItem {
	expenseItem.Meta = meta
	return expenseItem
}

// SetName устанавливает Наименование Статьи расходов.
func (expenseItem *ExpenseItem) SetName(name string) *ExpenseItem {
	expenseItem.Name = &name
	return expenseItem
}

// String реализует интерфейс [fmt.Stringer].
func (expenseItem ExpenseItem) String() string {
	return Stringify(expenseItem)
}

// MetaType возвращает код сущности.
func (ExpenseItem) MetaType() MetaType {
	return MetaTypeExpenseItem
}

// Update shortcut
func (expenseItem *ExpenseItem) Update(ctx context.Context, client *Client, params ...*Params) (*ExpenseItem, *resty.Response, error) {
	return NewExpenseItemService(client).Update(ctx, expenseItem.GetID(), expenseItem, params...)
}

// Create shortcut
func (expenseItem *ExpenseItem) Create(ctx context.Context, client *Client, params ...*Params) (*ExpenseItem, *resty.Response, error) {
	return NewExpenseItemService(client).Create(ctx, expenseItem, params...)
}

// Delete shortcut
func (expenseItem *ExpenseItem) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewExpenseItemService(client).Delete(ctx, expenseItem)
}

// ExpenseItemService описывает методы сервиса для работы со статьями расходов.
type ExpenseItemService interface {
	// GetList выполняет запрос на получение списка статей расходов.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[ExpenseItem], *resty.Response, error)

	// GetListAll выполняет запрос на получение всех статей расходов в виде списка.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает список объектов.
	GetListAll(ctx context.Context, params ...*Params) (*Slice[ExpenseItem], *resty.Response, error)

	// Create выполняет запрос на создание статьи расходов.
	// Обязательные поля для заполнения:
	//	- name (Наименование Статьи расходов)
	// Принимает контекст, статью расходов и опционально объект параметров запроса Params.
	// Возвращает созданную статью расходов.
	Create(ctx context.Context, expenseItem *ExpenseItem, params ...*Params) (*ExpenseItem, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или статей расходов.
	// Изменяемые статьи расходов должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список статей расходов и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых статей расходов.
	CreateUpdateMany(ctx context.Context, expenseItemList Slice[ExpenseItem], params ...*Params) (*Slice[ExpenseItem], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление статей расходов.
	// Принимает контекст и множество статей расходов.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*ExpenseItem) (*DeleteManyResponse, *resty.Response, error)

	// DeleteByID выполняет запрос на удаление статьи расходов по ID.
	// Принимает контекст и ID статьи расходов.
	// Возвращает «true» в случае успешного удаления статьи расходов.
	DeleteByID(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// Delete выполняет запрос на удаление статьи расходов.
	// Принимает контекст и статью расходов.
	// Возвращает «true» в случае успешного удаления статьи расходов.
	Delete(ctx context.Context, entity *ExpenseItem) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельной статьи расходов по ID.
	// Принимает контекст, ID статьи расходов взаиморасчётов и опционально объект параметров запроса Params.
	// Возвращает найденную статью расходов.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*ExpenseItem, *resty.Response, error)

	// Update выполняет запрос на изменение статьи расходов.
	// Принимает контекст, статью расходов и опционально объект параметров запроса Params.
	// Возвращает изменённую статью расходов.
	Update(ctx context.Context, id uuid.UUID, expenseItem *ExpenseItem, params ...*Params) (*ExpenseItem, *resty.Response, error)

	// MoveToTrash выполняет запрос на перемещение документа с указанным ID в корзину.
	// Принимает контекст и ID документа.
	// Возвращает «true» в случае успешного перемещения в корзину.
	MoveToTrash(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
}

const (
	EndpointExpenseItem = EndpointEntity + string(MetaTypeExpenseItem)
)

// NewExpenseItemService принимает [Client] и возвращает сервис для работы со статьями расходов.
func NewExpenseItemService(client *Client) ExpenseItemService {
	return newMainService[ExpenseItem, any, any, any](client, EndpointExpenseItem)
}
