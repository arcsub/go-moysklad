package moysklad

// PrepaymentService
// Сервис для работы с предоплатами.
type PrepaymentService struct {
	endpointGetList[Prepayment]
	endpointDelete
	endpointGetById[Prepayment]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointPositions[PrepaymentPosition]
	endpointAttributes
	endpointSyncID[Prepayment]
	endpointRemove
}

func NewPrepaymentService(client *Client) *PrepaymentService {
	e := NewEndpoint(client, "entity/prepayment")
	return &PrepaymentService{
		endpointGetList:    endpointGetList[Prepayment]{e},
		endpointDelete:     endpointDelete{e},
		endpointGetById:    endpointGetById[Prepayment]{e},
		endpointMetadata:   endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointPositions:  endpointPositions[PrepaymentPosition]{e},
		endpointAttributes: endpointAttributes{e},
		endpointSyncID:     endpointSyncID[Prepayment]{e},
		endpointRemove:     endpointRemove{e},
	}
}
