package moysklad

// RetailSalesReturnService
// Сервис для работы с розничными возвратами.
type RetailSalesReturnService struct {
	endpointGetList[RetailSalesReturn]
	endpointCreate[RetailSalesReturn]
	endpointCreateUpdateDeleteMany[RetailSalesReturn]
	endpointDelete
	endpointGetById[RetailSalesReturn]
	endpointUpdate[RetailSalesReturn]
	endpointTemplate[RetailSalesReturn]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointPositions[RetailPosition]
	endpointAttributes
	endpointPublication
	endpointSyncId[RetailSalesReturn]
	endpointNamedFilter
	endpointRemove
}

func NewRetailSalesReturnService(client *Client) *RetailSalesReturnService {
	e := NewEndpoint(client, "entity/retailsalesreturn")
	return &RetailSalesReturnService{
		endpointGetList:                endpointGetList[RetailSalesReturn]{e},
		endpointCreate:                 endpointCreate[RetailSalesReturn]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[RetailSalesReturn]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[RetailSalesReturn]{e},
		endpointUpdate:                 endpointUpdate[RetailSalesReturn]{e},
		endpointTemplate:               endpointTemplate[RetailSalesReturn]{e},
		endpointMetadata:               endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointPositions:              endpointPositions[RetailPosition]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointPublication:            endpointPublication{e},
		endpointSyncId:                 endpointSyncId[RetailSalesReturn]{e},
		endpointNamedFilter:            endpointNamedFilter{e},
		endpointRemove:                 endpointRemove{e},
	}
}
