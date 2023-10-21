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
	endpointTemplateBasedOn[InvoiceOut, InvoiceOutTemplateArg]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointPositions[InvoiceOutPosition]
	endpointAttributes
	endpointPublication
	endpointSyncID[InvoiceOut]
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
		endpointTemplateBasedOn:        endpointTemplateBasedOn[InvoiceOut, InvoiceOutTemplateArg]{e},
		endpointMetadata:               endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointPositions:              endpointPositions[InvoiceOutPosition]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointPublication:            endpointPublication{e},
		endpointSyncID:                 endpointSyncID[InvoiceOut]{e},
		endpointRemove:                 endpointRemove{e},
	}
}
