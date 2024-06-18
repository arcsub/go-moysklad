package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
)

// Subscription Подписка компании.
// Ключевое слово: subscription
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-podpiska-kompanii
type Subscription struct {
	Role                          string `json:"role,omitempty"`
	Tariff                        Tariff `json:"tariff,omitempty"`
	SubscriptionEndDate           int64  `json:"subscriptionEndDate,omitempty"`
	IsSubscriptionChangeAvailable bool   `json:"isSubscriptionChangeAvailable,omitempty"`
}

func (subscription Subscription) String() string {
	return Stringify(subscription)
}

func (subscription Subscription) MetaType() MetaType {
	return MetaTypeSubscription
}

// Tariff Действующий тариф аккаунта.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-podpiska-kompanii-dejstwuuschij-tarif-akkaunta
type Tariff string

const (
	TariffBasic        Tariff = "BASIC"        // Тариф "Базовый"
	TariffCorporate    Tariff = "CORPORATE"    // Тариф "Корпоративный"
	TariffFree         Tariff = "FREE"         // Тариф "Бесплатный 2014"
	TariffMinimal      Tariff = "MINIMAL"      // Тариф "Индивидуальный"
	TariffProfessional Tariff = "PROFESSIONAL" // Тариф "Профессиональный"
	TariffRetail       Tariff = "RETAIL"       // Тариф "Бесплатный"
	TariffStart        Tariff = "START"        // Тариф "Старт"
	TariffTrial        Tariff = "TRIAL"        // Тариф "Пробный"
)

type subscriptionService struct {
	Endpoint
}

func (service *subscriptionService) Get(ctx context.Context) (*Subscription, *resty.Response, error) {
	return NewRequestBuilder[Subscription](service.client, service.uri).Get(ctx)
}

// SubscriptionService Сервис для работы с подпиской компании.
type SubscriptionService interface {
	Get(ctx context.Context) (*Subscription, *resty.Response, error)
}

func NewSubscriptionService(client *Client) SubscriptionService {
	e := NewEndpoint(client, "entity/subscription")
	return &subscriptionService{e}
}
