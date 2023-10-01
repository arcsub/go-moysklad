package moysklad

// CurrencyService
// Сервис для работы с валютами.
type CurrencyService struct {
	endpointGetList[Currency]
	endpointCreate[Currency]
	endpointCreateUpdateDeleteMany[Currency]
	endpointDelete
	endpointGetById[Currency]
	endpointUpdate[Currency]
	endpointNamedFilter
}

func NewCurrencyService(client *Client) *CurrencyService {
	e := NewEndpoint(client, "entity/currency")
	return &CurrencyService{
		endpointGetList:                endpointGetList[Currency]{e},
		endpointCreate:                 endpointCreate[Currency]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[Currency]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[Currency]{e},
		endpointUpdate:                 endpointUpdate[Currency]{e},
		endpointNamedFilter:            endpointNamedFilter{e},
	}
}
