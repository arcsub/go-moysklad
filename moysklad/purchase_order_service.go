package moysklad

// PurchaseOrderService
// Сервис для работы с заказами поставщикам.
type PurchaseOrderService struct {
	endpointGetList[PurchaseOrder]
	endpointCreate[PurchaseOrder]
	endpointCreateUpdateDeleteMany[PurchaseOrder]
	endpointDelete
	endpointGetById[PurchaseOrder]
	endpointUpdate[PurchaseOrder]
	endpointTemplate[PurchaseOrder]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointPositions[PurchaseOrderPosition]
	endpointAttributes
	endpointPublication
	endpointSyncId[PurchaseOrder]
	endpointNamedFilter
	endpointRemove
}

func NewPurchaseOrderService(client *Client) *PurchaseOrderService {
	e := NewEndpoint(client, "entity/purchaseorder")
	return &PurchaseOrderService{
		endpointGetList:                endpointGetList[PurchaseOrder]{e},
		endpointCreate:                 endpointCreate[PurchaseOrder]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[PurchaseOrder]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[PurchaseOrder]{e},
		endpointUpdate:                 endpointUpdate[PurchaseOrder]{e},
		endpointTemplate:               endpointTemplate[PurchaseOrder]{e},
		endpointMetadata:               endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointPositions:              endpointPositions[PurchaseOrderPosition]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointPublication:            endpointPublication{e},
		endpointSyncId:                 endpointSyncId[PurchaseOrder]{e},
		endpointNamedFilter:            endpointNamedFilter{e},
		endpointRemove:                 endpointRemove{e},
	}
}
