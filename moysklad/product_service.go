package moysklad

// ProductService
// Сервис для работы с товарами.
type ProductService struct {
	endpointGetList[Product]
	endpointCreate[Product]
	endpointCreateUpdateDeleteMany[Product]
	endpointDelete
	endpointMetadata[MetadataAttributeShared]
	endpointAttributes
	endpointGetById[Product]
	endpointUpdate[Product]
	endpointImages
	endpointSyncID[Product]
	endpointNamedFilter
	endpointAudit
	endpointPrintLabel
}

func NewProductService(client *Client) *ProductService {
	e := NewEndpoint(client, "entity/product")
	return &ProductService{
		endpointGetList:                endpointGetList[Product]{e},
		endpointCreate:                 endpointCreate[Product]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[Product]{e},
		endpointDelete:                 endpointDelete{e},
		endpointMetadata:               endpointMetadata[MetadataAttributeShared]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointGetById:                endpointGetById[Product]{e},
		endpointUpdate:                 endpointUpdate[Product]{e},
		endpointImages:                 endpointImages{e},
		endpointSyncID:                 endpointSyncID[Product]{e},
		endpointNamedFilter:            endpointNamedFilter{e},
		endpointAudit:                  endpointAudit{e},
		endpointPrintLabel:             endpointPrintLabel{e},
	}
}
