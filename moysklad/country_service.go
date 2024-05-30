package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// CountryService
// Сервис для работы со странами.
type CountryService interface {
	GetList(ctx context.Context, params *Params) (*List[Country], *resty.Response, error)
	Create(ctx context.Context, country *Country, params *Params) (*Country, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, countryList []*Country, params *Params) (*[]Country, *resty.Response, error)
	DeleteMany(ctx context.Context, countryList []*Country) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Country, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, country *Country, params *Params) (*Country, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
}

func NewCountryService(client *Client) CountryService {
	e := NewEndpoint(client, "entity/country")
	return newMainService[Country, any, any, any](e)
}
