package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// WebhookStockService
// Сервис для работы с вебхуками на изменение остатков.
type WebhookStockService interface {
	GetList(ctx context.Context, params *Params) (*List[WebhookStock], *resty.Response, error)
	Create(ctx context.Context, webhookStock *WebhookStock, params *Params) (*WebhookStock, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, webhookStockList []*WebhookStock, params *Params) (*[]WebhookStock, *resty.Response, error)
	DeleteMany(ctx context.Context, webhookStockList []*WebhookStock) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*WebhookStock, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, webhookStock *WebhookStock, params *Params) (*WebhookStock, *resty.Response, error)
}

func NewWebhookStockService(client *Client) WebhookStockService {
	e := NewEndpoint(client, "entity/webhookstock")
	return newMainService[WebhookStock, any, any, any](e)
}
