package moysklad

// RetailDrawerCashInService
// Сервис для работы с внесениями денег.
type RetailDrawerCashInService struct {
	endpointGetList[RetailDrawerCashIn]
	endpointCreate[RetailDrawerCashIn]
	endpointCreateUpdateDeleteMany[RetailDrawerCashIn]
	endpointDelete
	endpointGetById[RetailDrawerCashIn]
	endpointUpdate[RetailDrawerCashIn]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointTemplate[RetailDrawerCashIn]
	endpointAttributes
	endpointPublication
	endpointSyncId[RetailDrawerCashIn]
	endpointNamedFilter
	endpointRemove
}

func NewRetailDrawerCashInService(client *Client) *RetailDrawerCashInService {
	e := NewEndpoint(client, "entity/retaildrawercashin")
	return &RetailDrawerCashInService{
		endpointGetList:                endpointGetList[RetailDrawerCashIn]{e},
		endpointCreate:                 endpointCreate[RetailDrawerCashIn]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[RetailDrawerCashIn]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[RetailDrawerCashIn]{e},
		endpointUpdate:                 endpointUpdate[RetailDrawerCashIn]{e},
		endpointMetadata:               endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointTemplate:               endpointTemplate[RetailDrawerCashIn]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointPublication:            endpointPublication{e},
		endpointSyncId:                 endpointSyncId[RetailDrawerCashIn]{e},
		endpointNamedFilter:            endpointNamedFilter{e},
		endpointRemove:                 endpointRemove{e},
	}
}
