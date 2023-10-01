package moysklad

// OrganizationService
// Сервис для работы с юридическими лицами.
type OrganizationService struct {
	endpointGetList[Organization]
	endpointCreate[Organization]
	endpointCreateUpdateDeleteMany[Organization]
	endpointDelete
	endpointGetById[Organization]
	endpointUpdate[Organization]
	endpointMetadata[MetadataAttributeShared]
	endpointAttributes
	endpointAccounts
	endpointSyncId[Organization]
}

func NewOrganizationService(client *Client) *OrganizationService {
	e := NewEndpoint(client, "entity/organization")
	return &OrganizationService{
		endpointGetList:                endpointGetList[Organization]{e},
		endpointCreate:                 endpointCreate[Organization]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[Organization]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[Organization]{e},
		endpointUpdate:                 endpointUpdate[Organization]{e},
		endpointMetadata:               endpointMetadata[MetadataAttributeShared]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointAccounts:               endpointAccounts{e},
		endpointSyncId:                 endpointSyncId[Organization]{e},
	}
}
