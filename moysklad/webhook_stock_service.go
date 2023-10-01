package moysklad

// WebhookStockService
// Сервис для работы с вебхуками на изменение остатков.
type WebhookStockService struct {
	endpointGetList[WebhookStock]
	endpointCreate[WebhookStock]
	endpointCreateUpdateDeleteMany[WebhookStock]
	endpointDelete
	endpointGetById[WebhookStock]
	endpointUpdate[WebhookStock]
}

func NewWebhookStockService(client *Client) *WebhookStockService {
	e := NewEndpoint(client, "entity/webhookstock")
	return &WebhookStockService{
		endpointGetList:                endpointGetList[WebhookStock]{e},
		endpointCreate:                 endpointCreate[WebhookStock]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[WebhookStock]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[WebhookStock]{e},
		endpointUpdate:                 endpointUpdate[WebhookStock]{e},
	}
}
