package moysklad

// RetailDrawerCashOutService
// Сервис для работы с выплатами денег.
type RetailDrawerCashOutService struct {
	endpointGetList[RetailDrawerCashOut]
	endpointCreate[RetailDrawerCashOut]
	endpointCreateUpdateDeleteMany[RetailDrawerCashOut]
	endpointDelete
	endpointGetById[RetailDrawerCashOut]
	endpointUpdate[RetailDrawerCashOut]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointTemplate[RetailDrawerCashOut]
	endpointAttributes
	endpointPublication
	endpointSyncID[RetailDrawerCashOut]
	endpointNamedFilter
	endpointRemove
}

func NewRetailDrawerCashOutService(client *Client) *RetailDrawerCashOutService {
	e := NewEndpoint(client, "entity/retaildrawercashout")
	return &RetailDrawerCashOutService{
		endpointGetList:                endpointGetList[RetailDrawerCashOut]{e},
		endpointCreate:                 endpointCreate[RetailDrawerCashOut]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[RetailDrawerCashOut]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[RetailDrawerCashOut]{e},
		endpointUpdate:                 endpointUpdate[RetailDrawerCashOut]{e},
		endpointMetadata:               endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointTemplate:               endpointTemplate[RetailDrawerCashOut]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointPublication:            endpointPublication{e},
		endpointSyncID:                 endpointSyncID[RetailDrawerCashOut]{e},
		endpointNamedFilter:            endpointNamedFilter{e},
		endpointRemove:                 endpointRemove{e},
	}
}
