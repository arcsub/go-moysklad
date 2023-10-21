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
	endpointTemplateBasedOn[PurchaseOrder, PurchaseOrderTemplateArg]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointPositions[PurchaseOrderPosition]
	endpointAttributes
	endpointPublication
	endpointSyncID[PurchaseOrder]
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
		endpointTemplateBasedOn:        endpointTemplateBasedOn[PurchaseOrder, PurchaseOrderTemplateArg]{e},
		endpointMetadata:               endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointPositions:              endpointPositions[PurchaseOrderPosition]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointPublication:            endpointPublication{e},
		endpointSyncID:                 endpointSyncID[PurchaseOrder]{e},
		endpointNamedFilter:            endpointNamedFilter{e},
		endpointRemove:                 endpointRemove{e},
	}
}
