package moysklad

// ProcessingPlanFolderService
// Сервис для работы с группами техкарт.
type ProcessingPlanFolderService struct {
	endpointGetList[ProcessingPlanFolder]
	endpointCreate[ProcessingPlanFolder]
	endpointDelete
	endpointGetById[ProcessingPlanFolder]
	endpointUpdate[ProcessingPlanFolder]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointNamedFilter
	endpointRemove
}

func NewProcessingPlanFolderService(client *Client) *ProcessingPlanFolderService {
	e := NewEndpoint(client, "entity/processingplanfolder")
	return &ProcessingPlanFolderService{
		endpointGetList:     endpointGetList[ProcessingPlanFolder]{e},
		endpointCreate:      endpointCreate[ProcessingPlanFolder]{e},
		endpointDelete:      endpointDelete{e},
		endpointGetById:     endpointGetById[ProcessingPlanFolder]{e},
		endpointUpdate:      endpointUpdate[ProcessingPlanFolder]{e},
		endpointMetadata:    endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointNamedFilter: endpointNamedFilter{e},
		endpointRemove:      endpointRemove{e},
	}
}
