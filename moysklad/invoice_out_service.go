package moysklad

// InvoiceOutService
// Сервис для работы со счетами покупателей.
type InvoiceOutService struct {
	endpointGetList[InvoiceOut]
	endpointCreate[InvoiceOut]
	endpointCreateUpdateDeleteMany[InvoiceOut]
	endpointDelete
	endpointGetById[InvoiceOut]
	endpointUpdate[InvoiceOut]
	endpointTemplate[InvoiceOut]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointPositions[InvoiceOutPosition]
	endpointAttributes
	endpointPublication
	endpointSyncId[InvoiceOut]
	endpointRemove
}

func NewInvoiceOutService(client *Client) *InvoiceOutService {
	e := NewEndpoint(client, "entity/invoiceout")
	return &InvoiceOutService{
		endpointGetList:                endpointGetList[InvoiceOut]{e},
		endpointCreate:                 endpointCreate[InvoiceOut]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[InvoiceOut]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[InvoiceOut]{e},
		endpointUpdate:                 endpointUpdate[InvoiceOut]{e},
		endpointTemplate:               endpointTemplate[InvoiceOut]{e},
		endpointMetadata:               endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointPositions:              endpointPositions[InvoiceOutPosition]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointPublication:            endpointPublication{e},
		endpointSyncId:                 endpointSyncId[InvoiceOut]{e},
		endpointRemove:                 endpointRemove{e},
	}
}
