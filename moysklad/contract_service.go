package moysklad

// ContractService
// Сервис для работы с договорами.
type ContractService struct {
	endpointGetList[Contract]
	endpointCreate[Contract]
	endpointCreateUpdateDeleteMany[Contract]
	endpointDelete
	endpointMetadata[MetadataAttributeSharedStates]
	endpointAttributes
	endpointGetById[Contract]
	endpointUpdate[Contract]
	endpointPublication
	endpointNamedFilter
	endpointRemove
}

func NewContractService(client *Client) *ContractService {
	e := NewEndpoint(client, "entity/contract")
	return &ContractService{
		endpointGetList:                endpointGetList[Contract]{e},
		endpointCreate:                 endpointCreate[Contract]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[Contract]{e},
		endpointDelete:                 endpointDelete{e},
		endpointMetadata:               endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointGetById:                endpointGetById[Contract]{e},
		endpointUpdate:                 endpointUpdate[Contract]{e},
		endpointPublication:            endpointPublication{e},
		endpointNamedFilter:            endpointNamedFilter{e},
		endpointRemove:                 endpointRemove{e},
	}
}
