package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// CurrencyService
// Сервис для работы с валютами.
type CurrencyService interface {
	GetList(ctx context.Context, params *Params) (*List[Currency], *resty.Response, error)
	Create(ctx context.Context, currency *Currency, params *Params) (*Currency, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, currencyList []*Currency, params *Params) (*[]Currency, *resty.Response, error)
	DeleteMany(ctx context.Context, currencyList []*Currency) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Currency, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, currency *Currency, params *Params) (*Currency, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
}

func NewCurrencyService(client *Client) CurrencyService {
	e := NewEndpoint(client, "entity/currency")
	return newMainService[Currency, any, any, any](e)
}
