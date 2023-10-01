package moysklad

// TaxRateService
// Сервис для работы со ставками НДС.
type TaxRateService struct {
	endpointGetList[TaxRate]
	endpointCreate[TaxRate]
	endpointCreateUpdateDeleteMany[TaxRate]
	endpointDelete
	endpointGetById[TaxRate]
	endpointUpdate[TaxRate]
}

func NewTaxRateService(client *Client) *TaxRateService {
	e := NewEndpoint(client, "entity/taxrate")
	return &TaxRateService{
		endpointGetList:                endpointGetList[TaxRate]{e},
		endpointCreate:                 endpointCreate[TaxRate]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[TaxRate]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[TaxRate]{e},
		endpointUpdate:                 endpointUpdate[TaxRate]{e},
	}
}
