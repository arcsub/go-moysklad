package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"time"
)

// Subscription Подписка компании.
//
// Код сущности: subscription
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-podpiska-kompanii
type Subscription struct {
	Role                          string `json:"role,omitempty"`                          // Роль авторизованного пользователя (USER/ADMIN)
	Tariff                        Tariff `json:"tariff,omitempty"`                        // Действующий тариф Аккаунта
	SubscriptionEndDate           int64  `json:"subscriptionEndDate,omitempty"`           // Дата (в миллисекундах) окончания действия текущего тарифа, если тариф отличается от “Пробный” и “Бесплатный”
	IsSubscriptionChangeAvailable bool   `json:"isSubscriptionChangeAvailable,omitempty"` // Доступность изменения подписки
}

// GetSubscriptionEndDateAsTime возвращает дату окончания действия текущего тарифа, если тариф отличается от “Пробный” и “Бесплатный”.
func (subscription Subscription) GetSubscriptionEndDateAsTime() time.Time {
	return time.Unix(subscription.SubscriptionEndDate, 0)
}

// String реализует интерфейс [fmt.Stringer].
func (subscription Subscription) String() string {
	return Stringify(subscription)
}

// MetaType возвращает код сущности.
func (Subscription) MetaType() MetaType {
	return MetaTypeSubscription
}

// Tariff Действующий тариф аккаунта.
//
// Возможные значения:
//   - TariffBasic        – Тариф "Базовый"
//   - TariffCorporate    – Тариф "Корпоративный"
//   - TariffFree         – Тариф "Бесплатный 2014"
//   - TariffMinimal      – Тариф "Индивидуальный"
//   - TariffProfessional – Тариф "Профессиональный"
//   - TariffRetail       – Тариф "Бесплатный"
//   - TariffStart        – Тариф "Старт"
//   - TariffTrial        – Тариф "Пробный"
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-podpiska-kompanii-dejstwuuschij-tarif-akkaunta
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

const (
	EndpointSubscription = EndpointAccountSettings + string(MetaTypeSubscription)
)

type subscriptionService struct {
	Endpoint
}

// Get выполняет запрос на получение подписки компании.
// Принимает контекст.
// Возвращает подписку компании.
func (service *subscriptionService) Get(ctx context.Context) (*Subscription, *resty.Response, error) {
	return NewRequestBuilder[Subscription](service.client, service.uri).Get(ctx)
}

// SubscriptionService описывает методы сервиса для работы с подпиской компании.
type SubscriptionService interface {
	Get(ctx context.Context) (*Subscription, *resty.Response, error)
}

// NewSubscriptionService принимает [Client] и возвращает сервис для работы с подпиской компании.
func NewSubscriptionService(client *Client) SubscriptionService {
	return &subscriptionService{NewEndpoint(client, EndpointSubscription)}
}
