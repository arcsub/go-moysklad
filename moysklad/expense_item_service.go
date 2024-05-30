package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// ExpenseItemService
// Сервис для работы со статьями расходов.
type ExpenseItemService interface {
	GetList(ctx context.Context, params *Params) (*List[ExpenseItem], *resty.Response, error)
	Create(ctx context.Context, expenseItem *ExpenseItem, params *Params) (*ExpenseItem, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, expenseItemList []*ExpenseItem, params *Params) (*[]ExpenseItem, *resty.Response, error)
	DeleteMany(ctx context.Context, expenseItemList []*ExpenseItem) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*ExpenseItem, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, expenseItem *ExpenseItem, params *Params) (*ExpenseItem, *resty.Response, error)
	Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewExpenseItemService(client *Client) ExpenseItemService {
	e := NewEndpoint(client, "entity/expenseitem")
	return newMainService[ExpenseItem, any, any, any](e)
}
