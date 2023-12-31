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
	endpointTemplateBasedOn[CashOut, CashOutTemplateArg]
	endpointGetById[CashOut]
	endpointUpdate[CashOut]
	endpointPublication
	endpointSyncID[CashOut]
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
		endpointTemplateBasedOn:        endpointTemplateBasedOn[CashOut, CashOutTemplateArg]{e},
		endpointGetById:                endpointGetById[CashOut]{e},
		endpointUpdate:                 endpointUpdate[CashOut]{e},
		endpointPublication:            endpointPublication{e},
		endpointSyncID:                 endpointSyncID[CashOut]{e},
		endpointRemove:                 endpointRemove{e},
	}
}
