package moysklad

// ThingService
// Сервис для работы с серийными номерами.
type ThingService struct {
	endpointGetList[Thing]
	endpointGetById[Thing]
}

func NewThingService(client *Client) *ThingService {
	e := NewEndpoint(client, "entity/thing")
	return &ThingService{
		endpointGetList: endpointGetList[Thing]{e},
		endpointGetById: endpointGetById[Thing]{e},
	}
}
