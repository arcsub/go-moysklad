package moysklad

// RetailShiftService
// Сервис для работы с розничными сменами.
type RetailShiftService struct {
	endpointGetList[RetailShift]
	endpointCreate[RetailShift]
	endpointDelete
	endpointGetById[RetailShift]
	endpointUpdate[RetailShift]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointAttributes
	endpointSyncId[RetailShift]
	endpointNamedFilter
	endpointRemove
}

func NewRetailShiftService(client *Client) *RetailShiftService {
	e := NewEndpoint(client, "entity/retailshift")
	return &RetailShiftService{
		endpointGetList:     endpointGetList[RetailShift]{e},
		endpointCreate:      endpointCreate[RetailShift]{e},
		endpointDelete:      endpointDelete{e},
		endpointGetById:     endpointGetById[RetailShift]{e},
		endpointUpdate:      endpointUpdate[RetailShift]{e},
		endpointMetadata:    endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointAttributes:  endpointAttributes{e},
		endpointSyncId:      endpointSyncId[RetailShift]{e},
		endpointNamedFilter: endpointNamedFilter{e},
		endpointRemove:      endpointRemove{e},
	}
}
