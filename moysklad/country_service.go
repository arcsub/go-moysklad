package moysklad

// CountryService
// Сервис для работы со странами.
type CountryService struct {
	endpointGetList[Country]
	endpointCreate[Country]
	endpointCreateUpdateDeleteMany[Country]
	endpointDelete
	endpointGetById[Country]
	endpointUpdate[Country]
	endpointNamedFilter
}

func NewCountryService(client *Client) *CountryService {
	e := NewEndpoint(client, "entity/country")
	return &CountryService{
		endpointGetList:                endpointGetList[Country]{e},
		endpointCreate:                 endpointCreate[Country]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[Country]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[Country]{e},
		endpointUpdate:                 endpointUpdate[Country]{e},
		endpointNamedFilter:            endpointNamedFilter{e},
	}
}
