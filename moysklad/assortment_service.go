package moysklad

// AssortmentService
// Сервис для работы с ассортиментом.
type AssortmentService struct {
	endpointGetOne[AssortmentResult]
	endpointGetOneAsync[AssortmentResult]
	endpointDeleteMany[AssortmentPosition]
	endpointSettings[AssortmentSettings]
	endpointTemplates
}

func NewAssortmentService(client *Client) *AssortmentService {
	e := NewEndpoint(client, "entity/assortment")
	return &AssortmentService{
		endpointGetOne:      endpointGetOne[AssortmentResult]{e},
		endpointGetOneAsync: endpointGetOneAsync[AssortmentResult]{e},
		endpointDeleteMany:  endpointDeleteMany[AssortmentPosition]{e},
		endpointSettings:    endpointSettings[AssortmentSettings]{e},
		endpointTemplates:   endpointTemplates{e},
	}
}
