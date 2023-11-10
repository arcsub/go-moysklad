package moysklad

// PaymentInService
// Сервис для работы с входящими платежами.
type PaymentInService struct {
	endpointGetList[PaymentIn]
	endpointCreate[PaymentIn]
	endpointCreateUpdateDeleteMany[PaymentIn]
	endpointDelete
	endpointGetById[PaymentIn]
	endpointUpdate[PaymentIn]
	endpointTemplate[PaymentIn]
	endpointTemplateBasedOn[PaymentIn, PaymentInTemplateArg]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointAttributes
	endpointPublication
	endpointSyncID[PaymentIn]
	endpointRemove
}

func NewPaymentInService(client *Client) *PaymentInService {
	e := NewEndpoint(client, "entity/paymentin")
	return &PaymentInService{
		endpointGetList:                endpointGetList[PaymentIn]{e},
		endpointCreate:                 endpointCreate[PaymentIn]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[PaymentIn]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[PaymentIn]{e},
		endpointUpdate:                 endpointUpdate[PaymentIn]{e},
		endpointTemplate:               endpointTemplate[PaymentIn]{e},
		endpointTemplateBasedOn:        endpointTemplateBasedOn[PaymentIn, PaymentInTemplateArg]{e},
		endpointMetadata:               endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointPublication:            endpointPublication{e},
		endpointSyncID:                 endpointSyncID[PaymentIn]{e},
		endpointRemove:                 endpointRemove{e},
	}
}
