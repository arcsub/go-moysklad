package moysklad

// ServiceService
// Сервис для работы с услугами.
type ServiceService struct {
	endpointGetList[Service]
	endpointCreate[Service]
	endpointCreateUpdateDeleteMany[Service]
	endpointDelete
	endpointGetById[Service]
	endpointUpdate[Service]
	endpointSyncID[Service]
	endpointNamedFilter
}

func NewServiceService(client *Client) *ServiceService {
	e := NewEndpoint(client, "entity/service")
	return &ServiceService{
		endpointGetList:                endpointGetList[Service]{e},
		endpointCreate:                 endpointCreate[Service]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[Service]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[Service]{e},
		endpointUpdate:                 endpointUpdate[Service]{e},
		endpointSyncID:                 endpointSyncID[Service]{e},
		endpointNamedFilter:            endpointNamedFilter{e},
	}
}
