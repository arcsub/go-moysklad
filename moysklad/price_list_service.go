package moysklad

// PriceListService
// Сервис для работы с прайс-листами.
type PriceListService struct {
	endpointGetList[PriceList]
	endpointCreate[PriceList]
	endpointCreateUpdateDeleteMany[PriceList]
	endpointDelete
	endpointGetById[PriceList]
	endpointUpdate[PriceList]
	endpointMetadata[MetadataAttributeSharedStates]
	endpointPositions[PriceListPosition]
	endpointAttributes
	endpointSyncId[PriceList]
	endpointRemove
}

func NewPriceListService(client *Client) *PriceListService {
	e := NewEndpoint(client, "entity/pricelist")
	return &PriceListService{
		endpointGetList:                endpointGetList[PriceList]{e},
		endpointCreate:                 endpointCreate[PriceList]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[PriceList]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[PriceList]{e},
		endpointUpdate:                 endpointUpdate[PriceList]{e},
		endpointMetadata:               endpointMetadata[MetadataAttributeSharedStates]{e},
		endpointPositions:              endpointPositions[PriceListPosition]{e},
		endpointAttributes:             endpointAttributes{e},
		endpointSyncId:                 endpointSyncId[PriceList]{e},
		endpointRemove:                 endpointRemove{e},
	}
}
