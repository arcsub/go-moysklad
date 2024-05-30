package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// CustomEntityService
// Сервис для работы с пользовательскими справочниками.
type CustomEntityService interface {
	Create(ctx context.Context, customEntity *CustomEntity, params *Params) (*CustomEntity, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, customEntity *CustomEntity, params *Params) (*CustomEntity, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetElements(ctx context.Context, id *uuid.UUID) (*List[CustomEntityElement], *resty.Response, error)
	CreateElement(ctx context.Context, id *uuid.UUID, element *CustomEntityElement) (*CustomEntityElement, *resty.Response, error)
	DeleteElement(ctx context.Context, id, elementID *uuid.UUID) (bool, *resty.Response, error)
	GetElementById(ctx context.Context, id, elementID *uuid.UUID) (*CustomEntityElement, *resty.Response, error)
	UpdateElement(ctx context.Context, id, elementID *uuid.UUID, element *CustomEntityElement) (*CustomEntityElement, *resty.Response, error)
}

type customEntityService struct {
	Endpoint
	endpointCreate[CustomEntity]
	endpointUpdate[CustomEntity]
	endpointDelete
}

func NewCustomEntityService(client *Client) CustomEntityService {
	e := NewEndpoint(client, "entity/customentity")
	return &customEntityService{
		Endpoint:       e,
		endpointCreate: endpointCreate[CustomEntity]{e},
		endpointUpdate: endpointUpdate[CustomEntity]{e},
		endpointDelete: endpointDelete{e},
	}
}

// GetElements Получить элементы справочника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-pol-zowatel-skij-sprawochnik-poluchit-alementy-sprawochnika
func (s *customEntityService) GetElements(ctx context.Context, id *uuid.UUID) (*List[CustomEntityElement], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s", s.uri, id)
	return NewRequestBuilder[List[CustomEntityElement]](s.client, path).Get(ctx)
}

// CreateElement Создать элемент справочника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-pol-zowatel-skij-sprawochnik-sozdat-alement-sprawochnika
func (s *customEntityService) CreateElement(ctx context.Context, id *uuid.UUID, element *CustomEntityElement) (*CustomEntityElement, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s", s.uri, id)
	return NewRequestBuilder[CustomEntityElement](s.client, path).Post(ctx, element)
}

// DeleteElement Удалить элемент справочника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-pol-zowatel-skij-sprawochnik-udalit-alement-sprawochnika
func (s *customEntityService) DeleteElement(ctx context.Context, id, elementId *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/%s", s.uri, id, elementId)
	return NewRequestBuilder[any](s.client, path).Delete(ctx)
}

// GetElementById Получить отдельный элементы справочника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-pol-zowatel-skij-sprawochnik-poluchit-alement
func (s *customEntityService) GetElementById(ctx context.Context, id, elementId *uuid.UUID) (*CustomEntityElement, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/%s", s.uri, id, elementId)
	return NewRequestBuilder[CustomEntityElement](s.client, path).Get(ctx)
}

// UpdateElement Изменить элемент справочника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-pol-zowatel-skij-sprawochnik-izmenit-alement
func (s *customEntityService) UpdateElement(ctx context.Context, id, elementId *uuid.UUID, element *CustomEntityElement) (*CustomEntityElement, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/%s", s.uri, id, elementId)
	return NewRequestBuilder[CustomEntityElement](s.client, path).Put(ctx, element)
}
