package moysklad

// LossService
// Сервис для работы со списаниями.
type LossService struct {
	endpointGetList[Loss]
	endpointCreate[Loss]
	endpointCreateUpdateDeleteMany[Loss]
	endpointDelete
	endpointGetById[Loss]
	endpointUpdate[Loss]
	endpointTemplate[Loss]
	endpointTemplateBasedOn[Loss, LossTemplateArg]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointPositions[LossPosition]
	endpointAttributes
	endpointPublication
	endpointSyncID[Loss]
	endpointRemove
}

func NewLossService(client *Client) *LossService {
	e := NewEndpoint(client, "entity/loss")
	return &LossService{
		endpointGetList:                endpointGetList[Loss]{e},
		endpointCreate:                 endpointCreate[Loss]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[Loss]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[Loss]{e},
		endpointUpdate:                 endpointUpdate[Loss]{e},
		endpointTemplate:               endpointTemplate[Loss]{e},
		endpointTemplateBasedOn:        endpointTemplateBasedOn[Loss, LossTemplateArg]{e},
		endpointMetadata:               endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointPositions:              endpointPositions[LossPosition]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointPublication:            endpointPublication{e},
		endpointSyncID:                 endpointSyncID[Loss]{e},
		endpointRemove:                 endpointRemove{e},
	}
}
