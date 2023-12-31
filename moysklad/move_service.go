package moysklad

// MoveService
// Сервис для работы со перемещениями.
type MoveService struct {
	endpointGetList[Move]
	endpointCreate[Move]
	endpointCreateUpdateDeleteMany[Move]
	endpointDelete
	endpointGetById[Move]
	endpointUpdate[Move]
	endpointTemplate[Move]
	endpointTemplateBasedOn[Move, MoveTemplateArg]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointPositions[MovePosition]
	endpointAttributes
	endpointPublication
	endpointSyncID[Move]
	endpointRemove
}

func NewMoveService(client *Client) *MoveService {
	e := NewEndpoint(client, "entity/move")
	return &MoveService{
		endpointGetList:                endpointGetList[Move]{e},
		endpointCreate:                 endpointCreate[Move]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[Move]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[Move]{e},
		endpointUpdate:                 endpointUpdate[Move]{e},
		endpointTemplate:               endpointTemplate[Move]{e},
		endpointTemplateBasedOn:        endpointTemplateBasedOn[Move, MoveTemplateArg]{e},
		endpointMetadata:               endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointPositions:              endpointPositions[MovePosition]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointPublication:            endpointPublication{e},
		endpointSyncID:                 endpointSyncID[Move]{e},
		endpointRemove:                 endpointRemove{e},
	}
}
