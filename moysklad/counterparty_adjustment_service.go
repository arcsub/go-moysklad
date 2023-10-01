package moysklad

// CounterPartyAdjustmentService
// Сервис для работы с корректировками баланса контрагента.
type CounterPartyAdjustmentService struct {
	endpointGetList[CounterPartyAdjustment]
	endpointCreate[CounterPartyAdjustment]
	endpointCreateUpdateDeleteMany[CounterPartyAdjustment]
	endpointDelete
	endpointGetById[CounterPartyAdjustment]
	endpointUpdate[CounterPartyAdjustment]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointNamedFilter
	endpointRemove
}

func NewCounterPartyAdjustmentService(client *Client) *CounterPartyAdjustmentService {
	e := NewEndpoint(client, "entity/counterpartyadjustment")
	return &CounterPartyAdjustmentService{
		endpointGetList:                endpointGetList[CounterPartyAdjustment]{e},
		endpointCreate:                 endpointCreate[CounterPartyAdjustment]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[CounterPartyAdjustment]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[CounterPartyAdjustment]{e},
		endpointUpdate:                 endpointUpdate[CounterPartyAdjustment]{e},
		endpointMetadata:               endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointNamedFilter:            endpointNamedFilter{e},
		endpointRemove:                 endpointRemove{e},
	}
}
