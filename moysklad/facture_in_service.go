package moysklad

// FactureInService
// Сервис для работы со счетами-фактурами полученными.
type FactureInService struct {
	endpointGetList[FactureIn]
	endpointCreate[FactureIn]
	endpointCreateUpdateDeleteMany[FactureIn]
	endpointDelete
	endpointGetById[FactureIn]
	endpointUpdate[FactureIn]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointAttributes
	endpointTemplate[FactureIn]
	endpointPublication
	endpointSyncId[FactureIn]
	endpointRemove
}

func NewFactureInService(client *Client) *FactureInService {
	e := NewEndpoint(client, "entity/facturein")
	return &FactureInService{
		endpointGetList:                endpointGetList[FactureIn]{e},
		endpointCreate:                 endpointCreate[FactureIn]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[FactureIn]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[FactureIn]{e},
		endpointUpdate:                 endpointUpdate[FactureIn]{e},
		endpointMetadata:               endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointTemplate:               endpointTemplate[FactureIn]{e},
		endpointPublication:            endpointPublication{e},
		endpointSyncId:                 endpointSyncId[FactureIn]{e},
		endpointRemove:                 endpointRemove{e},
	}
}
