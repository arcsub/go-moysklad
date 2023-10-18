package moysklad

// PurchaseReturnService
// Сервис для работы с возвратами поставщикам.
type PurchaseReturnService struct {
	endpointGetList[PurchaseReturn]
	endpointCreate[PurchaseReturn]
	endpointCreateUpdateDeleteMany[PurchaseReturn]
	endpointDelete
	endpointGetById[PurchaseReturn]
	endpointUpdate[PurchaseReturn]
	endpointTemplate[PurchaseReturn]
	endpointTemplateBasedOn[PurchaseReturn, PurchaseReturnTemplateArg]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointPositions[PurchaseReturnPosition]
	endpointAttributes
	endpointPublication
	endpointSyncId[PurchaseReturn]
	endpointNamedFilter
	endpointRemove
}

func NewPurchaseReturnService(client *Client) *PurchaseReturnService {
	e := NewEndpoint(client, "entity/purchasereturn")
	return &PurchaseReturnService{
		endpointGetList:                endpointGetList[PurchaseReturn]{e},
		endpointCreate:                 endpointCreate[PurchaseReturn]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[PurchaseReturn]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[PurchaseReturn]{e},
		endpointUpdate:                 endpointUpdate[PurchaseReturn]{e},
		endpointTemplate:               endpointTemplate[PurchaseReturn]{e},
		endpointTemplateBasedOn:        endpointTemplateBasedOn[PurchaseReturn, PurchaseReturnTemplateArg]{e},
		endpointMetadata:               endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointPositions:              endpointPositions[PurchaseReturnPosition]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointPublication:            endpointPublication{e},
		endpointSyncId:                 endpointSyncId[PurchaseReturn]{e},
		endpointNamedFilter:            endpointNamedFilter{e},
		endpointRemove:                 endpointRemove{e},
	}
}
