package moysklad

// ProductService
// Сервис для работы с товарами.
type ProductService struct {
	endpointGetList[Product]
	endpointCreate[Product]
	endpointCreateUpdateDeleteMany[Product]
	endpointDelete
	endpointDeleteMany[Product]
	endpointMetadata[MetadataAttributeShared]
	endpointAttributes
	endpointGetById[Product]
	endpointUpdate[Product]
	endpointImages
	endpointSyncId[Product]
	endpointNamedFilter
	endpointAudit
	endpointPrintPrice
}

func NewProductService(client *Client) *ProductService {
	e := NewEndpoint(client, "entity/product")
	return &ProductService{
		endpointGetList:                endpointGetList[Product]{e},
		endpointCreate:                 endpointCreate[Product]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[Product]{e},
		endpointDelete:                 endpointDelete{e},
		endpointDeleteMany:             endpointDeleteMany[Product]{e},
		endpointMetadata:               endpointMetadata[MetadataAttributeShared]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointGetById:                endpointGetById[Product]{e},
		endpointUpdate:                 endpointUpdate[Product]{e},
		endpointImages:                 endpointImages{e},
		endpointSyncId:                 endpointSyncId[Product]{e},
		endpointNamedFilter:            endpointNamedFilter{e},
		endpointAudit:                  endpointAudit{e},
		endpointPrintPrice:             endpointPrintPrice{e},
	}
}
