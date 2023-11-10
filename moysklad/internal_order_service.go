package moysklad

// InternalOrderService
// Сервис для работы с внутренними заказами.
type InternalOrderService struct {
	endpointGetList[InternalOrder]
	endpointCreate[InternalOrder]
	endpointCreateUpdateDeleteMany[InternalOrder]
	endpointDelete
	endpointGetById[InternalOrder]
	endpointUpdate[InternalOrder]
	endpointTemplate[InternalOrder]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointPositions[InternalOrderPosition]
	endpointAttributes
	endpointPublication
	endpointSyncID[InternalOrder]
	endpointRemove
}

func NewInternalOrderService(client *Client) *InternalOrderService {
	e := NewEndpoint(client, "entity/internalorder")
	return &InternalOrderService{
		endpointGetList:                endpointGetList[InternalOrder]{e},
		endpointCreate:                 endpointCreate[InternalOrder]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[InternalOrder]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[InternalOrder]{e},
		endpointUpdate:                 endpointUpdate[InternalOrder]{e},
		endpointTemplate:               endpointTemplate[InternalOrder]{e},
		endpointMetadata:               endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointPositions:              endpointPositions[InternalOrderPosition]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointPublication:            endpointPublication{e},
		endpointSyncID:                 endpointSyncID[InternalOrder]{e},
		endpointRemove:                 endpointRemove{e},
	}
}
