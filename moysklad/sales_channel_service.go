package moysklad

// SalesChannelService
// Сервис для работы с каналами продаж.
type SalesChannelService struct {
	endpointGetList[SalesChannel]
	endpointCreate[SalesChannel]
	endpointCreateUpdateDeleteMany[SalesChannel]
	endpointDelete
	endpointGetById[SalesChannel]
	endpointUpdate[SalesChannel]
}

func NewSalesChannelService(client *Client) *SalesChannelService {
	e := NewEndpoint(client, "entity/saleschannel")
	return &SalesChannelService{
		endpointGetList:                endpointGetList[SalesChannel]{e},
		endpointCreate:                 endpointCreate[SalesChannel]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[SalesChannel]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[SalesChannel]{e},
		endpointUpdate:                 endpointUpdate[SalesChannel]{e},
	}
}
