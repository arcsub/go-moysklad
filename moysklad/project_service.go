package moysklad

// ProjectService
// Сервис для работы с проектами.
type ProjectService struct {
	endpointGetList[Project]
	endpointCreate[Project]
	endpointCreateUpdateDeleteMany[Project]
	endpointDelete
	endpointGetById[Project]
	endpointUpdate[Project]
	endpointMetadata[MetadataAttributeShared]
	endpointAttributes
	endpointNamedFilter
}

func NewProjectService(client *Client) *ProjectService {
	e := NewEndpoint(client, "entity/project")
	return &ProjectService{
		endpointGetList:                endpointGetList[Project]{e},
		endpointCreate:                 endpointCreate[Project]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[Project]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[Project]{e},
		endpointUpdate:                 endpointUpdate[Project]{e},
		endpointMetadata:               endpointMetadata[MetadataAttributeShared]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointNamedFilter:            endpointNamedFilter{e},
	}
}
