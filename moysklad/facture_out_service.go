package moysklad

// FactureOutService
// Сервис для работы со счетами-фактурами выданными.
type FactureOutService struct {
	endpointGetList[FactureOut]
	endpointCreate[FactureOut]
	endpointCreateUpdateDeleteMany[FactureOut]
	endpointDelete
	endpointGetById[FactureOut]
	endpointUpdate[FactureOut]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointAttributes
	endpointPublication
	endpointSyncID[FactureOut]
	endpointRemove
}

func NewFactureOutService(client *Client) *FactureOutService {
	e := NewEndpoint(client, "entity/factureout")
	return &FactureOutService{
		endpointGetList:                endpointGetList[FactureOut]{e},
		endpointCreate:                 endpointCreate[FactureOut]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[FactureOut]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[FactureOut]{e},
		endpointUpdate:                 endpointUpdate[FactureOut]{e},
		endpointMetadata:               endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointPublication:            endpointPublication{e},
		endpointSyncID:                 endpointSyncID[FactureOut]{e},
		endpointRemove:                 endpointRemove{e},
	}
}
