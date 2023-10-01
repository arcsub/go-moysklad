package moysklad

// RoleService
// Сервис для работы с ролями и правами сотрудников.
type RoleService struct {
	endpointGetList[Role]
	endpointCreate[Role]
	endpointDelete
	endpointGetById[Role]
	endpointUpdate[Role]
}

// TODO: запросы на получение роли
func NewRoleService(client *Client) *RoleService {
	e := NewEndpoint(client, "entity/role")
	return &RoleService{
		endpointGetList: endpointGetList[Role]{e},
		endpointCreate:  endpointCreate[Role]{e},
		endpointDelete:  endpointDelete{e},
		endpointGetById: endpointGetById[Role]{e},
		endpointUpdate:  endpointUpdate[Role]{e},
	}
}
