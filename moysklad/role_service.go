package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
)

// RoleService
// Сервис для работы с ролями и правами сотрудников.
type RoleService struct {
	Endpoint
	endpointGetList[Role]
	endpointCreate[Role]
	endpointDelete
	endpointGetById[Role]
	endpointUpdate[Role]
}

func NewRoleService(client *Client) *RoleService {
	e := NewEndpoint(client, "entity/role")
	return &RoleService{
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
func (s *RoleService) GetAdminRole(ctx context.Context) (*AdminRole, *resty.Response, error) {
	path := "entity/role/admin"
	return NewRequestBuilder[AdminRole](s.client, path).Get(ctx)
}

// GetIndividualRole Запрос на получение индивидуальной роли.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sotrudnik-zapros-na-poluchenie-indiwidual-noj-roli
func (s *RoleService) GetIndividualRole(ctx context.Context) (*IndividualRole, *resty.Response, error) {
	path := "entity/role/individual"
	return NewRequestBuilder[IndividualRole](s.client, path).Get(ctx)
}

// GetCashierRole Запрос на получение роли кассира.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sotrudnik-zapros-na-poluchenie-indiwidual-noj-roli
func (s *RoleService) GetCashierRole(ctx context.Context) (*CashierRole, *resty.Response, error) {
	path := "entity/role/cashier"
	return NewRequestBuilder[CashierRole](s.client, path).Get(ctx)
}
