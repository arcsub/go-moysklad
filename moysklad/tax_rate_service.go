package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// TaxRateService
// Сервис для работы со ставками НДС.
type TaxRateService interface {
	GetList(ctx context.Context, params *Params) (*List[TaxRate], *resty.Response, error)
	Create(ctx context.Context, taxRate *TaxRate, params *Params) (*TaxRate, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, taxRateList []*TaxRate, params *Params) (*[]TaxRate, *resty.Response, error)
	DeleteMany(ctx context.Context, taxRateList []*TaxRate) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*TaxRate, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, taxRate *TaxRate, params *Params) (*TaxRate, *resty.Response, error)
}

func NewTaxRateService(client *Client) TaxRateService {
	e := NewEndpoint(client, "entity/taxrate")
	return newMainService[TaxRate, any, any, any](e)
}
