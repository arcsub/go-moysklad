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

func (e ExpenseItem) String() string {
	return Stringify(e)
}

func (e ExpenseItem) MetaType() MetaType {
	return MetaTypeExpenseItem
}

// ExpenseItemService
// Сервис для работы со статьями расходов.
type ExpenseItemService interface {
	GetList(ctx context.Context, params *Params) (*List[ExpenseItem], *resty.Response, error)
	Create(ctx context.Context, expenseItem *ExpenseItem, params *Params) (*ExpenseItem, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, expenseItemList []*ExpenseItem, params *Params) (*[]ExpenseItem, *resty.Response, error)
	DeleteMany(ctx context.Context, expenseItemList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*ExpenseItem, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, expenseItem *ExpenseItem, params *Params) (*ExpenseItem, *resty.Response, error)
	MoveToTrash(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewExpenseItemService(client *Client) ExpenseItemService {
	e := NewEndpoint(client, "entity/expenseitem")
	return newMainService[ExpenseItem, any, any, any](e)
}
