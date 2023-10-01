package moysklad

// ProcessingStageService
// Сервис для работы с этапами производства.
type ProcessingStageService struct {
	endpointGetList[ProcessingStage]
	endpointCreate[ProcessingStage]
	endpointCreateUpdateDeleteMany[ProcessingStage]
	endpointDelete
	endpointGetById[ProcessingStage]
	endpointUpdate[ProcessingStage]
	endpointNamedFilter
	endpointRemove
}

func NewProcessingStageService(client *Client) *ProcessingStageService {
	e := NewEndpoint(client, "entity/processingstage")
	return &ProcessingStageService{
		endpointGetList:                endpointGetList[ProcessingStage]{e},
		endpointCreate:                 endpointCreate[ProcessingStage]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[ProcessingStage]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[ProcessingStage]{e},
		endpointUpdate:                 endpointUpdate[ProcessingStage]{e},
		endpointNamedFilter:            endpointNamedFilter{e},
		endpointRemove:                 endpointRemove{e},
	}
}
