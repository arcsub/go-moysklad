package moysklad

// InvoiceInService
// Сервис для работы со счетами поставщиков.
type InvoiceInService struct {
	endpointGetList[InvoiceIn]
	endpointCreate[InvoiceIn]
	endpointCreateUpdateDeleteMany[InvoiceIn]
	endpointDelete
	endpointGetById[InvoiceIn]
	endpointUpdate[InvoiceIn]
	endpointTemplate[InvoiceIn]
	endpointTemplateBasedOn[InvoiceIn, InvoiceInTemplateArg]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointPositions[InvoiceInPosition]
	endpointAttributes
	endpointPublication
	endpointSyncID[InvoiceIn]
	endpointRemove
}

func NewInvoiceInService(client *Client) *InvoiceInService {
	e := NewEndpoint(client, "entity/invoicein")
	return &InvoiceInService{
		endpointGetList:                endpointGetList[InvoiceIn]{e},
		endpointCreate:                 endpointCreate[InvoiceIn]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[InvoiceIn]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[InvoiceIn]{e},
		endpointUpdate:                 endpointUpdate[InvoiceIn]{e},
		endpointTemplate:               endpointTemplate[InvoiceIn]{e},
		endpointTemplateBasedOn:        endpointTemplateBasedOn[InvoiceIn, InvoiceInTemplateArg]{e},
		endpointMetadata:               endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointPositions:              endpointPositions[InvoiceInPosition]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointPublication:            endpointPublication{e},
		endpointSyncID:                 endpointSyncID[InvoiceIn]{e},
		endpointRemove:                 endpointRemove{e},
	}
}
