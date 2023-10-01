package moysklad

// ProcessingProcessService
// Сервис для работы с тех процессами.
type ProcessingProcessService struct {
	endpointGetList[ProcessingProcess]
	endpointCreate[ProcessingProcess]
	endpointCreateUpdateDeleteMany[ProcessingProcess]
	endpointDelete
	endpointGetById[ProcessingProcess]
	endpointUpdate[ProcessingProcess]
	endpointPositions[ProcessingProcessPosition]
	endpointNamedFilter
	endpointRemove
}

func NewProcessingProcessService(client *Client) *ProcessingProcessService {
	e := NewEndpoint(client, "entity/processingprocess")
	return &ProcessingProcessService{
		endpointGetList:                endpointGetList[ProcessingProcess]{e},
		endpointCreate:                 endpointCreate[ProcessingProcess]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[ProcessingProcess]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[ProcessingProcess]{e},
		endpointUpdate:                 endpointUpdate[ProcessingProcess]{e},
		endpointPositions:              endpointPositions[ProcessingProcessPosition]{e},
		endpointNamedFilter:            endpointNamedFilter{e},
		endpointRemove:                 endpointRemove{e},
	}
}
