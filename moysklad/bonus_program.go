package moysklad

import (
	"github.com/google/uuid"
)

// BonusProgram Бонусная программа.
// Ключевое слово: bonusprogram
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-bonusnaq-programma
type BonusProgram struct {
	AccountID                 *uuid.UUID         `json:"accountId,omitempty"`                 // ID учетной записи
	Active                    *bool              `json:"active,omitempty"`                    // Индикатор, является ли бонусная программа активной на данный момент
	AgentTags                 *Tags              `json:"agentTags,omitempty"`                 // Тэги контрагентов, к которым применяется бонусная программа. В случае пустого значения контрагентов в результате выводится пустой массив.
	AllAgents                 *bool              `json:"allAgents,omitempty"`                 // Индикатор, действует ли скидка на всех контрагентов (см. Скидки)
	AllProducts               *bool              `json:"allProducts,omitempty"`               // Индикатор, действует ли бонусная программа на все товары (всегда true, см. Скидки)
	EarnRateRoublesToPoint    *int               `json:"earnRateRoublesToPoint,omitempty"`    // Курс начисления
	EarnWhileRedeeming        *bool              `json:"earnWhileRedeeming,omitempty"`        // Разрешить одновременное начисление и списание бонусов. Если true - бонусы будут начислены на денежную часть покупки, даже при частичной оплате покупки баллами.
	ID                        *uuid.UUID         `json:"id,omitempty"`                        // ID сущности
	MaxPaidRatePercents       *int               `json:"maxPaidRatePercents,omitempty"`       // Максимальный процент оплаты баллами
	Meta                      *Meta              `json:"meta,omitempty"`                      // Метаданные
	Name                      *string            `json:"name,omitempty"`                      // Наименование Бонусной программы
	PostponedBonusesDelayDays *int               `json:"postponedBonusesDelayDays,omitempty"` // Баллы начисляются через [N] дней (Тарифная опция «Расширенная бонусная программа»)
	SpendRatePointsToRouble   *int               `json:"spendRatePointsToRouble,omitempty"`   // Курс списания
	WelcomeBonusesEnabled     *bool              `json:"welcomeBonusesEnabled,omitempty"`     // Возможность начисления приветственных баллов
	WelcomeBonusesMode        WelcomeBonusesMode `json:"welcomeBonusesMode,omitempty"`        // Условие начисления приветственных баллов. Не может быть пустым, если welcomeBonusesEnabled = true
	WelcomeBonusesValue       *int               `json:"welcomeBonusesValue,omitempty"`       //	Количество приветственных баллов, начисляемых участникам бонусной программы. Не может быть отрицательным. Не может быть пустым, если welcomeBonusesEnabled = true
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
