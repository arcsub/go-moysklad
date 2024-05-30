package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// WebhookService
// Сервис для работы с вебхуками.
type WebhookService interface {
	GetList(ctx context.Context, params *Params) (*List[Webhook], *resty.Response, error)
	Create(ctx context.Context, webhook *Webhook, params *Params) (*Webhook, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, webhookList []*Webhook, params *Params) (*[]Webhook, *resty.Response, error)
	DeleteMany(ctx context.Context, webhookList []*Webhook) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Webhook, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, webhook *Webhook, params *Params) (*Webhook, *resty.Response, error)
}

func NewWebhookService(client *Client) WebhookService {
	e := NewEndpoint(client, "entity/webhook")
	return newMainService[Webhook, any, any, any](e)
}
