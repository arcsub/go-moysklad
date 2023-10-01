package moysklad

// ConsignmentService
// Сервис для работы с сериями.
type ConsignmentService struct {
	endpointGetList[Consignment]
	endpointCreate[Consignment]
	endpointCreateUpdateDeleteMany[Consignment]
	endpointDelete
	endpointGetById[Consignment]
	endpointUpdate[Consignment]
	endpointMetadata[MetadataAttribute]
	endpointAttributes
	endpointNamedFilter
}

func NewConsignmentService(client *Client) *ConsignmentService {
	e := NewEndpoint(client, "entity/consignment")
	return &ConsignmentService{
		endpointGetList:                endpointGetList[Consignment]{e},
		endpointCreate:                 endpointCreate[Consignment]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[Consignment]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[Consignment]{e},
		endpointUpdate:                 endpointUpdate[Consignment]{e},
		endpointMetadata:               endpointMetadata[MetadataAttribute]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointNamedFilter:            endpointNamedFilter{e},
	}
}
