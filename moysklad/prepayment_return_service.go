package moysklad

// PrepaymentReturnService
// Сервис для работы с возвратами предоплат.
type PrepaymentReturnService struct {
	endpointGetList[PrepaymentReturn]
	endpointGetById[PrepaymentReturn]
	endpointAttributes
	endpointMetadata[MetadataAttributeSharedStates]
	endpointPositions[PrepaymentReturnPosition]
	endpointSyncId[PrepaymentReturn]
	endpointRemove
}

func NewPrepaymentReturnService(client *Client) *PrepaymentReturnService {
	e := NewEndpoint(client, "entity/prepaymentreturn")
	return &PrepaymentReturnService{
		endpointGetList:    endpointGetList[PrepaymentReturn]{e},
		endpointGetById:    endpointGetById[PrepaymentReturn]{e},
		endpointAttributes: endpointAttributes{e},
		endpointMetadata:   endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointPositions:  endpointPositions[PrepaymentReturnPosition]{e},
		endpointSyncId:     endpointSyncId[PrepaymentReturn]{e},
		endpointRemove:     endpointRemove{e},
	}
}
