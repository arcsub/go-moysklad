package moysklad

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

// CustomEntityService
// Сервис для работы с пользовательскими справочниками.
type CustomEntityService struct {
	Endpoint
	endpointCreate[CustomEntity]
	endpointUpdate[CustomEntity]
	endpointDelete
}

func NewCustomEntityService(client *Client) *CustomEntityService {
	e := NewEndpoint(client, "entity/customentity")
	return &CustomEntityService{
		Endpoint:       e,
		endpointCreate: endpointCreate[CustomEntity]{e},
		endpointUpdate: endpointUpdate[CustomEntity]{e},
		endpointDelete: endpointDelete{e},
	}
}

// GetElements Получить элементы справочника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-pol-zowatel-skij-sprawochnik-poluchit-alementy-sprawochnika
func (s *CustomEntityService) GetElements(ctx context.Context, id *uuid.UUID) (*List[CustomEntityElement], *Response, error) {
	path := id.String()
	return NewRequestBuilder[List[CustomEntityElement]](s.Endpoint, ctx).WithPath(path).Get()
}

// CreateElement Создать элемент справочника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-pol-zowatel-skij-sprawochnik-sozdat-alement-sprawochnika
func (s *CustomEntityService) CreateElement(ctx context.Context, id *uuid.UUID, element *CustomEntityElement) (*CustomEntityElement, *Response, error) {
	path := id.String()
	return NewRequestBuilder[CustomEntityElement](s.Endpoint, ctx).WithPath(path).WithBody(element).Post()
}

// DeleteElement Удалить элемент справочника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-pol-zowatel-skij-sprawochnik-udalit-alement-sprawochnika
func (s *CustomEntityService) DeleteElement(ctx context.Context, id, elementId uuid.UUID) (bool, *Response, error) {
	path := fmt.Sprintf("%s/%s", id, elementId)
	return NewRequestBuilder[any](s.Endpoint, ctx).WithPath(path).Delete()
}

// GetElementById Получить отдельный элементы справочника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-pol-zowatel-skij-sprawochnik-poluchit-alement
func (s *CustomEntityService) GetElementById(ctx context.Context, id, elementId uuid.UUID) (*CustomEntityElement, *Response, error) {
	path := fmt.Sprintf("%s/%s", id, elementId)
	return NewRequestBuilder[CustomEntityElement](s.Endpoint, ctx).WithPath(path).Get()
}

// UpdateElement Изменить элемент справочника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-pol-zowatel-skij-sprawochnik-izmenit-alement
func (s *CustomEntityService) UpdateElement(ctx context.Context, id, elementId uuid.UUID, element *CustomEntityElement) (*CustomEntityElement, *Response, error) {
	path := fmt.Sprintf("%s/%s", id, elementId)
	return NewRequestBuilder[CustomEntityElement](s.Endpoint, ctx).WithPath(path).WithBody(element).Put()
}
