package moysklad

import "context"

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
func (s *RoleService) GetAdminRole(ctx context.Context) (*AdminRole, *Response, error) {
	path := "admin"
	return NewRequestBuilder[AdminRole](s.Endpoint, ctx).WithPath(path).Get()
}

// GetIndividualRole Запрос на получение индивидуальной роли.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sotrudnik-zapros-na-poluchenie-indiwidual-noj-roli
func (s *RoleService) GetIndividualRole(ctx context.Context) (*IndividualRole, *Response, error) {
	path := "individual"
	return NewRequestBuilder[IndividualRole](s.Endpoint, ctx).WithPath(path).Get()
}

// GetCashierRole Запрос на получение роли кассира.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sotrudnik-zapros-na-poluchenie-indiwidual-noj-roli
func (s *RoleService) GetCashierRole(ctx context.Context) (*CashierRole, *Response, error) {
	path := "cashier"
	return NewRequestBuilder[CashierRole](s.Endpoint, ctx).WithPath(path).Get()
}
