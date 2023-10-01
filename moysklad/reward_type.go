package moysklad

// RewardType Тип Вознаграждения.
type RewardType string

const (
	RewardTypePercentOfSales RewardType = "PercentOfSales" // Процент от суммы продажи
	RewardTypeNone           RewardType = "None"           // Не рассчитывать
)
