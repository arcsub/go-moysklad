package moysklad

// RegionService
// Сервис для работы с регионами.
type RegionService struct {
	endpointGetList[Region]
	endpointGetById[Region]
}

func NewRegionService(client *Client) *RegionService {
	e := NewEndpoint(client, "entity/region")
	return &RegionService{
		endpointGetList: endpointGetList[Region]{e},
		endpointGetById: endpointGetById[Region]{e},
	}
}
