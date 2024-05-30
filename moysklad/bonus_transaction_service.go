package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// BonusTransactionService
// Сервис для работы с бонусными операциями.
type BonusTransactionService interface {
	GetList(ctx context.Context, params *Params) (*List[BonusTransaction], *resty.Response, error)
	Create(ctx context.Context, bonusTransaction *BonusTransaction, params *Params) (*BonusTransaction, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, bonusTransactionList []*BonusTransaction, params *Params) (*[]BonusTransaction, *resty.Response, error)
	DeleteMany(ctx context.Context, bonusTransactionList []*BonusTransaction) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*BonusTransaction, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, bonusTransaction *BonusTransaction, params *Params) (*BonusTransaction, *resty.Response, error)
}

func NewBonusTransactionService(client *Client) BonusTransactionService {
	e := NewEndpoint(client, "entity/bonustransaction")
	return newMainService[BonusTransaction, any, any, any](e)
}
