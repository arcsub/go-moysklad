package moysklad

// ProductFolderService
// Сервис для работы с группами товаров.
type ProductFolderService struct {
	endpointGetList[ProductFolder]
	endpointCreate[ProductFolder]
	endpointCreateUpdateDeleteMany[ProductFolder]
	endpointDelete
	endpointMetadata[MetadataAttribute]
	endpointAttributes
	endpointGetById[ProductFolder]
	endpointUpdate[ProductFolder]
}

func NewProductFolderService(client *Client) *ProductFolderService {
	e := NewEndpoint(client, "entity/productfolder")
	return &ProductFolderService{
		endpointGetList:                endpointGetList[ProductFolder]{e},
		endpointCreate:                 endpointCreate[ProductFolder]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[ProductFolder]{e},
		endpointDelete:                 endpointDelete{e},
		endpointMetadata:               endpointMetadata[MetadataAttribute]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointGetById:                endpointGetById[ProductFolder]{e},
		endpointUpdate:                 endpointUpdate[ProductFolder]{e},
	}
}
