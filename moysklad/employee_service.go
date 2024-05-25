package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"net/http"
)

// EmployeeService
// Сервис для работы с сотрудниками.
type EmployeeService struct {
	Endpoint
	endpointGetList[Employee]
	endpointCreate[Employee]
	endpointCreateUpdateDeleteMany[Employee]
	endpointDelete
	endpointMetadata[MetadataAttributeShared]
	endpointAttributes
	endpointGetById[Employee]
	endpointUpdate[Employee]
}

func NewEmployeeService(client *Client) *EmployeeService {
	e := NewEndpoint(client, "entity/employee")
	return &EmployeeService{
		Endpoint:                       e,
		endpointGetList:                endpointGetList[Employee]{e},
		endpointCreate:                 endpointCreate[Employee]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[Employee]{e},
		endpointDelete:                 endpointDelete{e},
		endpointMetadata:               endpointMetadata[MetadataAttributeShared]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointGetById:                endpointGetById[Employee]{e},
		endpointUpdate:                 endpointUpdate[Employee]{e},
	}
}

// GetPermissions Запрос на получение информации о правах Сотрудника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sotrudnik-poluchit-informaciu-o-prawah-sotrudnika
func (s *EmployeeService) GetPermissions(ctx context.Context, id *uuid.UUID) (*EmployeePermission, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/security", s.uri, id)
	return NewRequestBuilder[EmployeePermission](s.client, path).Get(ctx)
}

// UpdatePermissions Запрос на изменение информации о правах Сотрудника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sotrudnik-izmenit-informaciu-o-prawah-sotrudnika
func (s *EmployeeService) UpdatePermissions(ctx context.Context, id *uuid.UUID, permissions *EmployeePermission) (*EmployeePermission, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/security", s.uri, id)
	return NewRequestBuilder[EmployeePermission](s.client, path).Put(ctx, permissions)
}

// Activate Активация Сотрудника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sotrudnik-aktiwaciq-sotrudnika
func (s *EmployeeService) Activate(ctx context.Context, id *uuid.UUID, permissions *EmployeePermission) (*MailActivationRequired, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/access/activate", s.uri, id)
	return NewRequestBuilder[MailActivationRequired](s.client, path).Put(ctx, permissions)
}

// Deactivate Деактивация Сотрудника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sotrudnik-deaktiwaciq-sotrudnika
func (s *EmployeeService) Deactivate(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/access/deactivate", s.uri, id)
	response, err := s.client.R().SetContext(ctx).Put(path)
	if err != nil {
		return false, response, err
	}
	ok := response.StatusCode() == http.StatusNoContent
	return ok, response, nil
}

// ResetPassword Сброс пароля Сотрудника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sotrudnik-sbros-parolq-sotrudnika
func (s *EmployeeService) ResetPassword(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/access/resetpassword", id)
	response, err := s.client.R().SetContext(ctx).Put(path)
	if err != nil {
		return false, response, err
	}
	ok := response.StatusCode() == http.StatusNoContent
	return ok, response, nil
}
