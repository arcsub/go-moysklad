package moysklad

// UomService
// Сервис для работы с единицами измерения.
type UomService struct {
	endpointGetList[Uom]
	endpointCreate[Uom]
	endpointCreateUpdateDeleteMany[Uom]
	endpointDelete
	endpointGetById[Uom]
	endpointUpdate[Uom]
}

func NewUomService(client *Client) *UomService {
	e := NewEndpoint(client, "entity/uom")
	return &UomService{
		endpointGetList:                endpointGetList[Uom]{e},
		endpointCreate:                 endpointCreate[Uom]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[Uom]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[Uom]{e},
		endpointUpdate:                 endpointUpdate[Uom]{e},
	}
}
