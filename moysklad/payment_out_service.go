package moysklad

// PaymentOutService
// Сервис для работы с исходящими платежами.
type PaymentOutService struct {
	endpointGetList[PaymentOut]
	endpointCreate[PaymentOut]
	endpointCreateUpdateDeleteMany[PaymentOut]
	endpointDelete
	endpointGetById[PaymentOut]
	endpointUpdate[PaymentOut]
	endpointTemplate[PaymentOut]
	endpointTemplateBasedOn[PaymentOut, PaymentOutTemplateArg]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointAttributes
	endpointPublication
	endpointSyncID[PaymentOut]
	endpointRemove
}

func NewPaymentOutService(client *Client) *PaymentOutService {
	e := NewEndpoint(client, "entity/paymentout")
	return &PaymentOutService{
		endpointGetList:                endpointGetList[PaymentOut]{e},
		endpointCreate:                 endpointCreate[PaymentOut]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[PaymentOut]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[PaymentOut]{e},
		endpointUpdate:                 endpointUpdate[PaymentOut]{e},
		endpointTemplate:               endpointTemplate[PaymentOut]{e},
		endpointTemplateBasedOn:        endpointTemplateBasedOn[PaymentOut, PaymentOutTemplateArg]{e},
		endpointMetadata:               endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointPublication:            endpointPublication{e},
		endpointSyncID:                 endpointSyncID[PaymentOut]{e},
		endpointRemove:                 endpointRemove{e},
	}
}
