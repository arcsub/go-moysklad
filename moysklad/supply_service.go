package moysklad

// SupplyService
// Сервис для работы с приёмками.
type SupplyService struct {
	endpointGetList[Supply]
	endpointCreate[Supply]
	endpointCreateUpdateDeleteMany[Supply]
	endpointDelete
	endpointGetById[Supply]
	endpointUpdate[Supply]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointTemplate[Supply]
	endpointPositions[SupplyPosition]
	endpointAttributes
	endpointPublication
	endpointPrintDoc
	endpointFiles
	endpointNamedFilter
	endpointStates
	endpointSyncId[Supply]
	endpointRemove
}

func NewSupplyService(client *Client) *SupplyService {
	e := NewEndpoint(client, "entity/supply")
	return &SupplyService{
		endpointGetList:                endpointGetList[Supply]{e},
		endpointCreate:                 endpointCreate[Supply]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[Supply]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[Supply]{e},
		endpointUpdate:                 endpointUpdate[Supply]{e},
		endpointMetadata:               endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointTemplate:               endpointTemplate[Supply]{e},
		endpointPositions:              endpointPositions[SupplyPosition]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointPublication:            endpointPublication{e},
		endpointPrintDoc:               endpointPrintDoc{e},
		endpointFiles:                  endpointFiles{e},
		endpointNamedFilter:            endpointNamedFilter{e},
		endpointStates:                 endpointStates{e},
		endpointSyncId:                 endpointSyncId[Supply]{e},
		endpointRemove:                 endpointRemove{e},
	}
}
