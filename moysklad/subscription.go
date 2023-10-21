package moysklad

// Subscription Подписка компании.
// Ключевое слово: subscription
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-podpiska-kompanii
type Subscription struct {
	Role                          *string `json:"role,omitempty"`
	IsSubscriptionChangeAvailable *bool   `json:"isSubscriptionChangeAvailable,omitempty"`
	SubscriptionEndDate           *int64  `json:"subscriptionEndDate,omitempty"`
	Tariff                        Tariff  `json:"tariff,omitempty"`
}

func (s Subscription) String() string {
	return Stringify(s)
}

func (s Subscription) MetaType() MetaType {
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
