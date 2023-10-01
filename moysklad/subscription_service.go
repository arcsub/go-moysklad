package moysklad

// SubscriptionService
// Сервис для работы с подпиской компании.
type SubscriptionService struct {
	endpointGetOne[Subscription]
}

func NewSubscriptionService(client *Client) *SubscriptionService {
	e := NewEndpoint(client, "entity/subscription")
	return &SubscriptionService{
		endpointGetOne: endpointGetOne[Subscription]{e},
	}
}
