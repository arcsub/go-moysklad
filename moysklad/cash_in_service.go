package moysklad

// CashInService
// Сервис для работы с приходными ордерами.
type CashInService struct {
	endpointGetList[CashIn]
	endpointCreate[CashIn]
	endpointCreateUpdateDeleteMany[CashIn]
	endpointDelete
	endpointDeleteMany[CashIn]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointAttributes
	endpointTemplate[CashIn]
	endpointGetById[CashIn]
	endpointUpdate[CashIn]
	endpointPublication
	endpointSyncId[CashIn]
	endpointRemove
}

func NewCashInService(client *Client) *CashInService {
	e := NewEndpoint(client, "entity/cashin")
	return &CashInService{
		endpointGetList:                endpointGetList[CashIn]{e},
		endpointCreate:                 endpointCreate[CashIn]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[CashIn]{e},
		endpointDelete:                 endpointDelete{e},
		endpointDeleteMany:             endpointDeleteMany[CashIn]{e},
		endpointMetadata:               endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointTemplate:               endpointTemplate[CashIn]{e},
		endpointGetById:                endpointGetById[CashIn]{e},
		endpointUpdate:                 endpointUpdate[CashIn]{e},
		endpointPublication:            endpointPublication{e},
		endpointSyncId:                 endpointSyncId[CashIn]{e},
		endpointRemove:                 endpointRemove{e},
	}
}
