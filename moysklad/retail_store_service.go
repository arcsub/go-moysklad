package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// RetailStoreService
// Сервис для работы с точками продаж.
type RetailStoreService interface {
	GetList(ctx context.Context, params *Params) (*List[RetailStore], *resty.Response, error)
	Create(ctx context.Context, retailStore *RetailStore, params *Params) (*RetailStore, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, retailStore []*RetailStore, params *Params) (*[]RetailStore, *resty.Response, error)
	DeleteMany(ctx context.Context, retailStore []*RetailStore) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*RetailStore, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, entity *RetailStore, params *Params) (*RetailStore, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
	GetCashiers(ctx context.Context, id *uuid.UUID) (*MetaArray[Cashier], *resty.Response, error)
	GetCashierById(ctx context.Context, id, cashierId *uuid.UUID) (*Cashier, *resty.Response, error)
}

type retailStoreService struct {
	Endpoint
	endpointGetList[RetailStore]
	endpointCreate[RetailStore]
	endpointCreateUpdateMany[RetailStore]
	endpointDeleteMany[RetailStore]
	endpointDelete
	endpointGetById[RetailStore]
	endpointUpdate[RetailStore]
	endpointNamedFilter
}

func NewRetailStoreService(client *Client) RetailStoreService {
	e := NewEndpoint(client, "entity/retailstore")
	return &retailStoreService{
		Endpoint:                 e,
		endpointGetList:          endpointGetList[RetailStore]{e},
		endpointCreate:           endpointCreate[RetailStore]{e},
		endpointCreateUpdateMany: endpointCreateUpdateMany[RetailStore]{e},
		endpointDeleteMany:       endpointDeleteMany[RetailStore]{e},
		endpointDelete:           endpointDelete{e},
		endpointGetById:          endpointGetById[RetailStore]{e},
		endpointUpdate:           endpointUpdate[RetailStore]{e},
		endpointNamedFilter:      endpointNamedFilter{e},
	}
}

// GetCashiers Получить Кассиров.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kassir-poluchit-kassirow
func (s *retailStoreService) GetCashiers(ctx context.Context, id *uuid.UUID) (*MetaArray[Cashier], *resty.Response, error) {
	path := fmt.Sprintf("entity/retailstore/%s/cashiers", id)
	return NewRequestBuilder[MetaArray[Cashier]](s.client, path).Get(ctx)
}

// GetCashierById Получить Кассира.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kassir-poluchit-kassira
func (s *retailStoreService) GetCashierById(ctx context.Context, id, cashierId *uuid.UUID) (*Cashier, *resty.Response, error) {
	path := fmt.Sprintf("entity/retailstore/%s/cashiers/%s", id, cashierId)
	return NewRequestBuilder[Cashier](s.client, path).Get(ctx)
}
