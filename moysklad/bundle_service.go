package moysklad

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

// BundleService
// Сервис для работы с комплектами.
type BundleService struct {
	Endpoint
	endpointGetList[Bundle]
	endpointCreate[Bundle]
	endpointCreateUpdateDeleteMany[Bundle]
	endpointGetById[Bundle]
	endpointUpdate[Bundle]
	endpointDelete
	endpointDeleteMany[Bundle]
}

func NewBundleService(client *Client) *BundleService {
	e := NewEndpoint(client, "entity/bundle")
	return &BundleService{
		Endpoint:                       e,
		endpointGetList:                endpointGetList[Bundle]{e},
		endpointCreate:                 endpointCreate[Bundle]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[Bundle]{e},
		endpointGetById:                endpointGetById[Bundle]{e},
		endpointUpdate:                 endpointUpdate[Bundle]{e},
		endpointDelete:                 endpointDelete{e},
		endpointDeleteMany:             endpointDeleteMany[Bundle]{e},
	}
}

// GetComponents Получить компоненты Комплекта.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-komplekt-poluchit-komponenty-komplekta
func (s *BundleService) GetComponents(ctx context.Context, id *uuid.UUID) (*List[BundleComponent], *Response, error) {
	path := fmt.Sprintf("entity/bundle/%s/components", id)
	return NewRequestBuilder[List[BundleComponent]](s.Endpoint, ctx).WithPath(path).Get()
}

// CreateComponent Добавить компонент Комплекта.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-komplekt-dobawit-komponent-komplekta
func (s *BundleService) CreateComponent(ctx context.Context, id *uuid.UUID, bundleComponent *BundleComponent) (*BundleComponent, *Response, error) {
	path := fmt.Sprintf("entity/bundle/%s/components", id)
	return NewRequestBuilder[BundleComponent](s.Endpoint, ctx).WithPath(path).WithBody(bundleComponent).Post()
}

// GetComponentById Получить компонент.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-komplekt-poluchit-komponent
func (s *BundleService) GetComponentById(ctx context.Context, id, componentId uuid.UUID) (*BundleComponent, *Response, error) {
	path := fmt.Sprintf("entity/bundle/%s/components/%s", id, componentId)
	return NewRequestBuilder[BundleComponent](s.Endpoint, ctx).WithPath(path).Get()
}

// UpdateComponent Изменить компонент.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-komplekt-izmenit-komponent
func (s *BundleService) UpdateComponent(ctx context.Context, id, componentId uuid.UUID, bundleComponent *BundleComponent) (*BundleComponent, *Response, error) {
	path := fmt.Sprintf("entity/bundle/%s/components/%s", id, componentId)
	return NewRequestBuilder[BundleComponent](s.Endpoint, ctx).WithPath(path).WithBody(bundleComponent).Put()
}

// DeleteComponent Удалить компонент.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-komplekt-udalit-komponent
func (s *BundleService) DeleteComponent(ctx context.Context, id, componentId uuid.UUID) (bool, *Response, error) {
	path := fmt.Sprintf("entity/bundle/%s/components/%s", id, componentId)
	return NewRequestBuilder[any](s.Endpoint, ctx).WithPath(path).WithBody(componentId).Delete()
}
