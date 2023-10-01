package moysklad

// ApplicationService
// Сервис для работы с приложениями.
type ApplicationService struct {
	endpointGetList[Application] // Получение списка сущностей установленных приложений
	endpointGetById[Application] // Получение сущности установленного приложения
}

func NewApplicationService(client *Client) *ApplicationService {
	e := NewEndpoint(client, "entity/application")
	return &ApplicationService{
		endpointGetList: endpointGetList[Application]{e},
		endpointGetById: endpointGetById[Application]{e},
	}
}
