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
	AllProducts               *bool              `json:"allProducts,omitempty"`
	Active                    *bool              `json:"active,omitempty"`
	WelcomeBonusesValue       *int               `json:"welcomeBonusesValue,omitempty"`
	AllAgents                 *bool              `json:"allAgents,omitempty"`
	MaxPaidRatePercents       *int               `json:"maxPaidRatePercents,omitempty"`
	EarnRateRoublesToPoint    *int               `json:"earnRateRoublesToPoint,omitempty"`
	EarnWhileRedeeming        *bool              `json:"earnWhileRedeeming,omitempty"`
	ID                        *uuid.UUID         `json:"id,omitempty"`
	AccountID                 *uuid.UUID         `json:"accountId,omitempty"`
	Meta                      *Meta              `json:"meta,omitempty"`
	Name                      *string            `json:"name,omitempty"`
	PostponedBonusesDelayDays *int               `json:"postponedBonusesDelayDays,omitempty"`
	SpendRatePointsToRouble   *int               `json:"spendRatePointsToRouble,omitempty"`
	WelcomeBonusesEnabled     *bool              `json:"welcomeBonusesEnabled,omitempty"`
	WelcomeBonusesMode        WelcomeBonusesMode `json:"welcomeBonusesMode,omitempty"`
	AgentTags                 Tags               `json:"agentTags,omitempty"`
}

func (bonusProgram BonusProgram) String() string {
	return Stringify(bonusProgram)
}

func (bonusProgram BonusProgram) MetaType() MetaType {
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
	DeleteMany(ctx context.Context, bonusProgramList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
}

func NewBonusProgramService(client *Client) BonusProgramService {
	e := NewEndpoint(client, "entity/bonusprogram")
	return newMainService[BonusProgram, any, any, any](e)
}
