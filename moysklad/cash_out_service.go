package moysklad

// CashOutService cashout
// Сервис для работы с расходными ордерами.
type CashOutService struct {
	endpointGetList[CashOut]
	endpointCreate[CashOut]
	endpointCreateUpdateDeleteMany[CashOut]
	endpointDelete
	endpointDeleteMany[CashOut]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointAttributes
	endpointTemplate[CashOut]
	endpointGetById[CashOut]
	endpointUpdate[CashOut]
	endpointPublication
	endpointSyncId[CashOut]
	endpointRemove
}

func NewCashOutService(client *Client) *CashOutService {
	e := NewEndpoint(client, "entity/cashout")
	return &CashOutService{
		endpointGetList:                endpointGetList[CashOut]{e},
		endpointCreate:                 endpointCreate[CashOut]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[CashOut]{e},
		endpointDelete:                 endpointDelete{e},
		endpointDeleteMany:             endpointDeleteMany[CashOut]{e},
		endpointMetadata:               endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointTemplate:               endpointTemplate[CashOut]{e},
		endpointGetById:                endpointGetById[CashOut]{e},
		endpointUpdate:                 endpointUpdate[CashOut]{e},
		endpointPublication:            endpointPublication{e},
		endpointSyncId:                 endpointSyncId[CashOut]{e},
		endpointRemove:                 endpointRemove{e},
	}
}
