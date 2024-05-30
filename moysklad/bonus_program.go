package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
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

// BonusProgramService
// Сервис для работы с бонусными программами.
type BonusProgramService interface {
	GetList(ctx context.Context, params *Params) (*List[BonusProgram], *resty.Response, error)
	Create(ctx context.Context, bonusProgram *BonusProgram, params *Params) (*BonusProgram, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, bonusProgram *BonusProgram, params *Params) (*BonusProgram, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*BonusProgram, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	DeleteMany(ctx context.Context, bonusProgramList []*BonusProgram) (*DeleteManyResponse, *resty.Response, error)
}

func NewBonusProgramService(client *Client) BonusProgramService {
	e := NewEndpoint(client, "entity/bonusprogram")
	return newMainService[BonusProgram, any, any, any](e)
}
