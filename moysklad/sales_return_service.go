package moysklad

// SalesReturnService
// Сервис для работы с возвратами покупателей.
type SalesReturnService struct {
	endpointGetList[SalesReturn]
	endpointCreate[SalesReturn]
	endpointCreateUpdateDeleteMany[SalesReturn]
	endpointDelete
	endpointGetById[SalesReturn]
	endpointUpdate[SalesReturn]
	endpointTemplate[SalesReturn]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointPositions[SalesReturnPosition]
	endpointAttributes
	endpointPublication
	endpointSyncId[SalesReturn]
	endpointNamedFilter
	endpointRemove
}

func NewSalesReturnService(client *Client) *SalesReturnService {
	e := NewEndpoint(client, "entity/salesreturn")
	return &SalesReturnService{
		endpointGetList:                endpointGetList[SalesReturn]{e},
		endpointCreate:                 endpointCreate[SalesReturn]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[SalesReturn]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[SalesReturn]{e},
		endpointUpdate:                 endpointUpdate[SalesReturn]{e},
		endpointTemplate:               endpointTemplate[SalesReturn]{e},
		endpointMetadata:               endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointPositions:              endpointPositions[SalesReturnPosition]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointPublication:            endpointPublication{e},
		endpointSyncId:                 endpointSyncId[SalesReturn]{e},
		endpointNamedFilter:            endpointNamedFilter{e},
		endpointRemove:                 endpointRemove{e},
	}
}
