package moysklad

// ProcessingOrderService
// Сервис для работы с заказами на производство.
type ProcessingOrderService struct {
	endpointGetList[ProcessingOrder]
	endpointCreate[ProcessingOrder]
	endpointCreateUpdateDeleteMany[ProcessingOrder]
	endpointDelete
	endpointGetById[ProcessingOrder]
	endpointUpdate[ProcessingOrder]
	endpointTemplate[ProcessingOrder]
	endpointTemplateBasedOn[ProcessingOrder, ProcessingOrderTemplateArg]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointPositions[ProcessingOrderPosition]
	endpointAttributes
	endpointSyncID[ProcessingOrder]
	endpointRemove
}

func NewProcessingOrderService(client *Client) *ProcessingOrderService {
	e := NewEndpoint(client, "entity/processingorder")
	return &ProcessingOrderService{
		endpointGetList:                endpointGetList[ProcessingOrder]{e},
		endpointCreate:                 endpointCreate[ProcessingOrder]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[ProcessingOrder]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[ProcessingOrder]{e},
		endpointUpdate:                 endpointUpdate[ProcessingOrder]{e},
		endpointTemplate:               endpointTemplate[ProcessingOrder]{e},
		endpointTemplateBasedOn:        endpointTemplateBasedOn[ProcessingOrder, ProcessingOrderTemplateArg]{e},
		endpointMetadata:               endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointPositions:              endpointPositions[ProcessingOrderPosition]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointSyncID:                 endpointSyncID[ProcessingOrder]{e},
		endpointRemove:                 endpointRemove{e},
	}
}
