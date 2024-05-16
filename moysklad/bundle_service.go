package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
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
func (s *BundleService) GetComponents(ctx context.Context, id *uuid.UUID) (*List[BundleComponent], *resty.Response, error) {
	path := fmt.Sprintf("entity/bundle/%s/components", id)
	return NewRequestBuilder[List[BundleComponent]](s.client, path).Get(ctx)
}

// CreateComponent Добавить компонент Комплекта.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-komplekt-dobawit-komponent-komplekta
func (s *BundleService) CreateComponent(ctx context.Context, id *uuid.UUID, bundleComponent *BundleComponent) (*BundleComponent, *resty.Response, error) {
	path := fmt.Sprintf("entity/bundle/%s/components", id)
	return NewRequestBuilder[BundleComponent](s.client, path).Post(ctx, bundleComponent)
}

// GetComponentById Получить компонент.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-komplekt-poluchit-komponent
func (s *BundleService) GetComponentById(ctx context.Context, id, componentId *uuid.UUID) (*BundleComponent, *resty.Response, error) {
	path := fmt.Sprintf("entity/bundle/%s/components/%s", id, componentId)
	return NewRequestBuilder[BundleComponent](s.client, path).Get(ctx)
}

// UpdateComponent Изменить компонент.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-komplekt-izmenit-komponent
func (s *BundleService) UpdateComponent(ctx context.Context, id, componentId *uuid.UUID, bundleComponent *BundleComponent) (*BundleComponent, *resty.Response, error) {
	path := fmt.Sprintf("entity/bundle/%s/components/%s", id, componentId)
	return NewRequestBuilder[BundleComponent](s.client, path).Put(ctx, bundleComponent)
}

// DeleteComponent Удалить компонент.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-komplekt-udalit-komponent
func (s *BundleService) DeleteComponent(ctx context.Context, id, componentId *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("entity/bundle/%s/components/%s", id, componentId)
	return NewRequestBuilder[any](s.client, path).Delete(ctx)
}
