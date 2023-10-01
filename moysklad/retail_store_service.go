package moysklad

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

// RetailStoreService
// Сервис для работы с точками продаж.
type RetailStoreService struct {
	Endpoint
	endpointGetList[RetailStore]
	endpointCreate[RetailStore]
	endpointCreateUpdateDeleteMany[RetailStore]
	endpointDelete
	endpointGetById[RetailStore]
	endpointUpdate[RetailStore]
	endpointNamedFilter
}

func NewRetailStoreService(client *Client) *RetailStoreService {
	e := NewEndpoint(client, "entity/retailstore")
	return &RetailStoreService{
		Endpoint:                       e,
		endpointGetList:                endpointGetList[RetailStore]{e},
		endpointCreate:                 endpointCreate[RetailStore]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[RetailStore]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[RetailStore]{e},
		endpointUpdate:                 endpointUpdate[RetailStore]{e},
		endpointNamedFilter:            endpointNamedFilter{e},
	}
}

// GetCashiers Получить Кассиров.
// https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kassir-poluchit-kassirow
func (s *RetailStoreService) GetCashiers(ctx context.Context, id *uuid.UUID) (*MetaArray[Cashier], *Response, error) {
	path := fmt.Sprintf("%s/cashiers", id)
	return NewRequestBuilder[MetaArray[Cashier]](s.Endpoint, ctx).WithPath(path).Get()
}

// GetCashierById Получить Кассира.
// https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kassir-poluchit-kassira
func (s *RetailStoreService) GetCashierById(ctx context.Context, id, cashierId uuid.UUID) (*Cashier, *Response, error) {
	path := fmt.Sprintf("%s/cashiers/%s", id, cashierId)
	return NewRequestBuilder[Cashier](s.Endpoint, ctx).WithPath(path).Get()
}
