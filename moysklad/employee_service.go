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
type EmployeeService interface {
	GetList(ctx context.Context, params *Params) (*List[Employee], *resty.Response, error)
	Create(ctx context.Context, employee *Employee, params *Params) (*Employee, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, employeeList []*Employee, params *Params) (*[]Employee, *resty.Response, error)
	DeleteMany(ctx context.Context, employeeList []*Employee) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetadataAttributeShared, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id *uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error)
	UpdateAttribute(ctx context.Context, id *uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList []*Attribute) (*DeleteManyResponse, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Employee, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, employee *Employee, params *Params) (*Employee, *resty.Response, error)
	GetPermissions(ctx context.Context, id *uuid.UUID) (*EmployeePermission, *resty.Response, error)
	UpdatePermissions(ctx context.Context, id *uuid.UUID, permissions *EmployeePermission) (*EmployeePermission, *resty.Response, error)
	Activate(ctx context.Context, id *uuid.UUID, permissions *EmployeePermission) (*MailActivationRequired, *resty.Response, error)
	Deactivate(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	ResetPassword(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

type employeeService struct {
	Endpoint
	endpointGetList[Employee]
	endpointCreate[Employee]
	endpointCreateUpdateMany[Employee]
	endpointDeleteMany[Employee]
	endpointDelete
	endpointMetadata[MetadataAttributeShared]
	endpointAttributes
	endpointGetById[Employee]
	endpointUpdate[Employee]
}

func NewEmployeeService(client *Client) EmployeeService {
	e := NewEndpoint(client, "entity/employee")
	return &employeeService{
		Endpoint:                 e,
		endpointGetList:          endpointGetList[Employee]{e},
		endpointCreate:           endpointCreate[Employee]{e},
		endpointCreateUpdateMany: endpointCreateUpdateMany[Employee]{e},
		endpointDeleteMany:       endpointDeleteMany[Employee]{e},
		endpointDelete:           endpointDelete{e},
		endpointMetadata:         endpointMetadata[MetadataAttributeShared]{e},
		endpointAttributes:       endpointAttributes{e},
		endpointGetById:          endpointGetById[Employee]{e},
		endpointUpdate:           endpointUpdate[Employee]{e},
	}
}

// GetPermissions Запрос на получение информации о правах Сотрудника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sotrudnik-poluchit-informaciu-o-prawah-sotrudnika
func (s *employeeService) GetPermissions(ctx context.Context, id *uuid.UUID) (*EmployeePermission, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/security", s.uri, id)
	return NewRequestBuilder[EmployeePermission](s.client, path).Get(ctx)
}

// UpdatePermissions Запрос на изменение информации о правах Сотрудника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sotrudnik-izmenit-informaciu-o-prawah-sotrudnika
func (s *employeeService) UpdatePermissions(ctx context.Context, id *uuid.UUID, permissions *EmployeePermission) (*EmployeePermission, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/security", s.uri, id)
	return NewRequestBuilder[EmployeePermission](s.client, path).Put(ctx, permissions)
}

// Activate Активация Сотрудника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sotrudnik-aktiwaciq-sotrudnika
func (s *employeeService) Activate(ctx context.Context, id *uuid.UUID, permissions *EmployeePermission) (*MailActivationRequired, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/access/activate", s.uri, id)
	return NewRequestBuilder[MailActivationRequired](s.client, path).Put(ctx, permissions)
}

// Deactivate Деактивация Сотрудника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sotrudnik-deaktiwaciq-sotrudnika
func (s *employeeService) Deactivate(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/access/deactivate", s.uri, id)
	_, resp, err := NewRequestBuilder[any](s.client, path).Put(ctx, nil)
	if err != nil {
		return false, resp, err
	}
	return resp.StatusCode() == http.StatusNoContent, resp, nil
}

// ResetPassword Сброс пароля Сотрудника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sotrudnik-sbros-parolq-sotrudnika
func (s *employeeService) ResetPassword(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/access/resetpassword", id)
	_, resp, err := NewRequestBuilder[any](s.client, path).Put(ctx, nil)
	if err != nil {
		return false, resp, err
	}
	return resp.StatusCode() == http.StatusNoContent, resp, nil
}
