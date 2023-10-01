package moysklad

// CustomerOrderService
// Сервис для работы с заказами покупателя.
type CustomerOrderService struct {
	endpointGetList[CustomerOrder]
	endpointCreate[CustomerOrder]
	endpointCreateUpdateDeleteMany[CustomerOrder]
	endpointDelete
	endpointGetById[CustomerOrder]
	endpointUpdate[CustomerOrder]
	endpointTemplate[CustomerOrder]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointPositions[CustomerOrderPosition]
	endpointAttributes
	endpointPublication
	endpointSyncId[CustomerOrder]
	endpointNamedFilter
	endpointRemove
}

func NewCustomerOrderService(client *Client) *CustomerOrderService {
	e := NewEndpoint(client, "entity/customerorder")
	return &CustomerOrderService{
		endpointGetList:                endpointGetList[CustomerOrder]{e},
		endpointCreate:                 endpointCreate[CustomerOrder]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[CustomerOrder]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[CustomerOrder]{e},
		endpointUpdate:                 endpointUpdate[CustomerOrder]{e},
		endpointTemplate:               endpointTemplate[CustomerOrder]{e},
		endpointMetadata:               endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointPositions:              endpointPositions[CustomerOrderPosition]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointPublication:            endpointPublication{e},
		endpointSyncId:                 endpointSyncId[CustomerOrder]{e},
		endpointNamedFilter:            endpointNamedFilter{e},
		endpointRemove:                 endpointRemove{e},
	}
}
