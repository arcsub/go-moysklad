package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// RoleService
// Сервис для работы с ролями и правами сотрудников.
type RoleService interface {
	GetList(ctx context.Context, params *Params) (*List[Role], *resty.Response, error)
	Create(ctx context.Context, role *Role, params *Params) (*Role, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Role, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, role *Role, params *Params) (*Role, *resty.Response, error)
	GetAdminRole(ctx context.Context) (*AdminRole, *resty.Response, error)
	GetIndividualRole(ctx context.Context) (*IndividualRole, *resty.Response, error)
	GetCashierRole(ctx context.Context) (*CashierRole, *resty.Response, error)
	GetWorkerRole(ctx context.Context) (*CashierRole, *resty.Response, error)
}

type roleService struct {
	Endpoint
	endpointGetList[Role]
	endpointCreate[Role]
	endpointDelete
	endpointGetById[Role]
	endpointUpdate[Role]
}

func NewRoleService(client *Client) RoleService {
	e := NewEndpoint(client, "entity/role")
	return &roleService{
		Endpoint:        e,
		endpointGetList: endpointGetList[Role]{e},
		endpointCreate:  endpointCreate[Role]{e},
		endpointDelete:  endpointDelete{e},
		endpointGetById: endpointGetById[Role]{e},
		endpointUpdate:  endpointUpdate[Role]{e},
	}
}

// GetAdminRole Запрос на получение роли админа.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sotrudnik-zapros-na-poluchenie-roli-admina
func (s *roleService) GetAdminRole(ctx context.Context) (*AdminRole, *resty.Response, error) {
	path := "entity/role/admin"
	return NewRequestBuilder[AdminRole](s.client, path).Get(ctx)
}

// GetIndividualRole Запрос на получение индивидуальной роли.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sotrudnik-zapros-na-poluchenie-indiwidual-noj-roli
func (s *roleService) GetIndividualRole(ctx context.Context) (*IndividualRole, *resty.Response, error) {
	path := "entity/role/individual"
	return NewRequestBuilder[IndividualRole](s.client, path).Get(ctx)
}

// GetCashierRole Запрос на получение роли кассира.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sotrudnik-zapros-na-poluchenie-indiwidual-noj-roli
func (s *roleService) GetCashierRole(ctx context.Context) (*CashierRole, *resty.Response, error) {
	path := "entity/role/cashier"
	return NewRequestBuilder[CashierRole](s.client, path).Get(ctx)
}

// GetWorkerRole Запрос на получение роли кассира.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sotrudnik-zapros-na-poluchenie-roli-sotrudnika-proizwodstwa
func (s *roleService) GetWorkerRole(ctx context.Context) (*CashierRole, *resty.Response, error) {
	path := "entity/role/worker"
	return NewRequestBuilder[CashierRole](s.client, path).Get(ctx)
}
