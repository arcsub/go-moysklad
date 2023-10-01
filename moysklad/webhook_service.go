package moysklad

// WebhookService
// Сервис для работы с вебхуками.
type WebhookService struct {
	endpointGetList[Webhook]
	endpointCreate[Webhook]
	endpointCreateUpdateDeleteMany[Webhook]
	endpointDelete
	endpointGetById[Webhook]
	endpointUpdate[Webhook]
}

func NewWebhookService(client *Client) *WebhookService {
	e := NewEndpoint(client, "entity/webhook")
	return &WebhookService{
		endpointGetList:                endpointGetList[Webhook]{e},
		endpointCreate:                 endpointCreate[Webhook]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[Webhook]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[Webhook]{e},
		endpointUpdate:                 endpointUpdate[Webhook]{e},
	}
}
