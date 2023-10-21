package moysklad

import (
	"github.com/google/uuid"
)

// BonusProgram Бонусная программа.
// Ключевое слово: bonusprogram
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-bonusnaq-programma
type BonusProgram struct {
	MaxPaidRatePercents       *int               `json:"maxPaidRatePercents,omitempty"`
	EarnWhileRedeeming        *bool              `json:"earnWhileRedeeming,omitempty"`
	AgentTags                 *Tags              `json:"agentTags,omitempty"`
	AllAgents                 *bool              `json:"allAgents,omitempty"`
	AccountID                 *uuid.UUID         `json:"accountId,omitempty"`
	EarnRateRoublesToPoint    *int               `json:"earnRateRoublesToPoint,omitempty"`
	Active                    *bool              `json:"active,omitempty"`
	ID                        *uuid.UUID         `json:"id,omitempty"`
	AllProducts               *bool              `json:"allProducts,omitempty"`
	Meta                      *Meta              `json:"meta,omitempty"`
	Name                      *string            `json:"name,omitempty"`
	PostponedBonusesDelayDays *int               `json:"postponedBonusesDelayDays,omitempty"`
	SpendRatePointsToRouble   *int               `json:"spendRatePointsToRouble,omitempty"`
	WelcomeBonusesEnabled     *bool              `json:"welcomeBonusesEnabled,omitempty"`
	WelcomeBonusesValue       *int               `json:"welcomeBonusesValue,omitempty"`
	WelcomeBonusesMode        WelcomeBonusesMode `json:"welcomeBonusesMode,omitempty"`
}

func (b BonusProgram) String() string {
	return Stringify(b)
}

func (b BonusProgram) MetaType() MetaType {
	return MetaTypeBonusProgram
}

// WelcomeBonusesMode Условия бонусных баллов
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-bonusnaq-programma-bonusnye-programmy-atributy-suschnosti-uslowiq-bonusnyh-ballow
type WelcomeBonusesMode string

const (
	WelcomeBonusesModeRegistration  WelcomeBonusesMode = "REGISTRATION"   // Приветственные баллы начисляются участникам после регистрации в бонусной программе.
	WelcomeBonusesModeFirstPurchase WelcomeBonusesMode = "FIRST_PURCHASE" // Приветственные баллы начисляются участникам бонусной программы после совершения первой покупки.
)
