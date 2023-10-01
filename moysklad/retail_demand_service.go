package moysklad

// RetailDemandService
// Сервис для работы с розничными продажами.
type RetailDemandService struct {
	endpointGetList[RetailDemand]
	endpointCreate[RetailDemand]
	endpointCreateUpdateDeleteMany[RetailDemand]
	endpointDelete
	endpointGetById[RetailDemand]
	endpointUpdate[RetailDemand]
	endpointTemplate[RetailDemand]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointPositions[RetailPosition]
	endpointAttributes
	endpointPublication
	endpointSyncId[RetailDemand]
	endpointNamedFilter
	endpointRemove
}

func NewRetailDemandService(client *Client) *RetailDemandService {
	e := NewEndpoint(client, "entity/retaildemand")
	return &RetailDemandService{
		endpointGetList:                endpointGetList[RetailDemand]{e},
		endpointCreate:                 endpointCreate[RetailDemand]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[RetailDemand]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[RetailDemand]{e},
		endpointUpdate:                 endpointUpdate[RetailDemand]{e},
		endpointTemplate:               endpointTemplate[RetailDemand]{e},
		endpointMetadata:               endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointPositions:              endpointPositions[RetailPosition]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointPublication:            endpointPublication{e},
		endpointSyncId:                 endpointSyncId[RetailDemand]{e},
		endpointNamedFilter:            endpointNamedFilter{e},
		endpointRemove:                 endpointRemove{e},
	}
}
