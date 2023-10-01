package moysklad

import (
	"context"
	"fmt"
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
func (s *EmployeeService) GetPermissions(ctx context.Context, id *uuid.UUID) (*EmployeePermission, *Response, error) {
	path := fmt.Sprintf("%s/security", id)
	return NewRequestBuilder[EmployeePermission](s.Endpoint, ctx).WithPath(path).Get()
}

// UpdatePermissions Запрос на изменение информации о правах Сотрудника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sotrudnik-izmenit-informaciu-o-prawah-sotrudnika
func (s *EmployeeService) UpdatePermissions(ctx context.Context, id *uuid.UUID, permissions *EmployeePermission) (*EmployeePermission, *Response, error) {
	path := fmt.Sprintf("%s/security", id)
	return NewRequestBuilder[EmployeePermission](s.Endpoint, ctx).WithPath(path).WithBody(permissions).Put()
}

// Activate Активация Сотрудника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sotrudnik-aktiwaciq-sotrudnika
func (s *EmployeeService) Activate(ctx context.Context, id *uuid.UUID, permissions *EmployeePermission) (*MailActivationRequired, *Response, error) {
	path := fmt.Sprintf("%s/access/activate", id)
	return NewRequestBuilder[MailActivationRequired](s.Endpoint, ctx).WithPath(path).WithBody(permissions).Put()
}

// Deactivate Деактивация Сотрудника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sotrudnik-deaktiwaciq-sotrudnika
func (s *EmployeeService) Deactivate(ctx context.Context, id *uuid.UUID) (bool, *Response, error) {
	path := fmt.Sprintf("%s/access/deactivate", id)
	rb := NewRequestBuilder[any](s.Endpoint, ctx).WithPath(path)
	response, err := rb.do(http.MethodPut)
	if err != nil {
		return false, response, err
	}
	ok := response.StatusCode == http.StatusNoContent
	return ok, response, nil
}

// ResetPassword Сброс пароля Сотрудника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sotrudnik-sbros-parolq-sotrudnika
func (s *EmployeeService) ResetPassword(ctx context.Context, id *uuid.UUID) (bool, *Response, error) {
	path := fmt.Sprintf("%s/access/resetpassword", id)
	rb := NewRequestBuilder[any](s.Endpoint, ctx).WithPath(path)
	response, err := rb.do(http.MethodPut)
	if err != nil {
		return false, response, err
	}
	ok := response.StatusCode == http.StatusNoContent
	return ok, response, nil
}
