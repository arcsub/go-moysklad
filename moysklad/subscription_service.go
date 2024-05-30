package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
)

// SubscriptionService
// Сервис для работы с подпиской компании.
type SubscriptionService interface {
	Get(ctx context.Context, params *Params) (*Subscription, *resty.Response, error)
}

func NewSubscriptionService(client *Client) SubscriptionService {
	e := NewEndpoint(client, "entity/subscription")
	return newMainService[Subscription, any, any, any](e)
}
