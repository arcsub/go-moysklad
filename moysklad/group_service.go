package moysklad

// GroupService
// Сервис для работы с отделами.
type GroupService struct {
	endpointGetList[Group]
	endpointCreate[Group]
	endpointDelete
	endpointGetById[Group]
	endpointUpdate[Group]
}

func NewGroupService(client *Client) *GroupService {
	e := NewEndpoint(client, "entity/group")
	return &GroupService{
		endpointGetList: endpointGetList[Group]{e},
		endpointCreate:  endpointCreate[Group]{e},
		endpointDelete:  endpointDelete{e},
		endpointGetById: endpointGetById[Group]{e},
		endpointUpdate:  endpointUpdate[Group]{e},
	}
}
