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
	AgentTags                 Slice[string]      `json:"agentTags,omitempty"`
}

func (bonusProgram BonusProgram) Clean() *BonusProgram {
	return &BonusProgram{Meta: bonusProgram.Meta}
}

func (bonusProgram BonusProgram) GetAllProducts() bool {
	return Deref(bonusProgram.AllProducts)
}

func (bonusProgram BonusProgram) GetActive() bool {
	return Deref(bonusProgram.Active)
}

func (bonusProgram BonusProgram) GetWelcomeBonusesValue() int {
	return Deref(bonusProgram.WelcomeBonusesValue)
}

func (bonusProgram BonusProgram) GetAllAgents() bool {
	return Deref(bonusProgram.AllAgents)
}

func (bonusProgram BonusProgram) GetMaxPaidRatePercents() int {
	return Deref(bonusProgram.MaxPaidRatePercents)
}

func (bonusProgram BonusProgram) GetEarnRateRoublesToPoint() int {
	return Deref(bonusProgram.EarnRateRoublesToPoint)
}

func (bonusProgram BonusProgram) GetEarnWhileRedeeming() bool {
	return Deref(bonusProgram.EarnWhileRedeeming)
}

func (bonusProgram BonusProgram) GetID() uuid.UUID {
	return Deref(bonusProgram.ID)
}

func (bonusProgram BonusProgram) GetAccountID() uuid.UUID {
	return Deref(bonusProgram.AccountID)
}

func (bonusProgram BonusProgram) GetMeta() Meta {
	return Deref(bonusProgram.Meta)
}

func (bonusProgram BonusProgram) GetName() string {
	return Deref(bonusProgram.Name)
}

func (bonusProgram BonusProgram) GetPostponedBonusesDelayDays() int {
	return Deref(bonusProgram.PostponedBonusesDelayDays)
}

func (bonusProgram BonusProgram) GetSpendRatePointsToRouble() int {
	return Deref(bonusProgram.SpendRatePointsToRouble)
}

func (bonusProgram BonusProgram) GetWelcomeBonusesEnabled() bool {
	return Deref(bonusProgram.WelcomeBonusesEnabled)
}

func (bonusProgram BonusProgram) GetWelcomeBonusesMode() WelcomeBonusesMode {
	return bonusProgram.WelcomeBonusesMode
}

func (bonusProgram BonusProgram) GetAgentTags() Slice[string] {
	return bonusProgram.AgentTags
}

func (bonusProgram *BonusProgram) SetAllProducts(allProducts bool) *BonusProgram {
	bonusProgram.AllProducts = &allProducts
	return bonusProgram
}

func (bonusProgram *BonusProgram) SetActive(active bool) *BonusProgram {
	bonusProgram.Active = &active
	return bonusProgram
}

func (bonusProgram *BonusProgram) SetWelcomeBonusesValue(welcomeBonusesValue int) *BonusProgram {
	bonusProgram.WelcomeBonusesValue = &welcomeBonusesValue
	return bonusProgram
}

func (bonusProgram *BonusProgram) SetAllAgents(allAgents bool) *BonusProgram {
	bonusProgram.AllAgents = &allAgents
	return bonusProgram
}

func (bonusProgram *BonusProgram) SetMaxPaidRatePercents(maxPaidRatePercents int) *BonusProgram {
	bonusProgram.MaxPaidRatePercents = &maxPaidRatePercents
	return bonusProgram
}

func (bonusProgram *BonusProgram) SetEarnRateRoublesToPoint(earnRateRoublesToPoint int) *BonusProgram {
	bonusProgram.EarnRateRoublesToPoint = &earnRateRoublesToPoint
	return bonusProgram
}

func (bonusProgram *BonusProgram) SetEarnWhileRedeeming(earnWhileRedeeming bool) *BonusProgram {
	bonusProgram.EarnWhileRedeeming = &earnWhileRedeeming
	return bonusProgram
}

func (bonusProgram *BonusProgram) SetMeta(meta *Meta) *BonusProgram {
	bonusProgram.Meta = meta
	return bonusProgram
}

func (bonusProgram *BonusProgram) SetName(name string) *BonusProgram {
	bonusProgram.Name = &name
	return bonusProgram
}

func (bonusProgram *BonusProgram) SetPostponedBonusesDelayDays(postponedBonusesDelayDays int) *BonusProgram {
	bonusProgram.PostponedBonusesDelayDays = &postponedBonusesDelayDays
	return bonusProgram
}

func (bonusProgram *BonusProgram) SetSpendRatePointsToRouble(spendRatePointsToRouble int) *BonusProgram {
	bonusProgram.SpendRatePointsToRouble = &spendRatePointsToRouble
	return bonusProgram
}

func (bonusProgram *BonusProgram) SetWelcomeBonusesEnabled(welcomeBonusesEnabled bool) *BonusProgram {
	bonusProgram.WelcomeBonusesEnabled = &welcomeBonusesEnabled
	return bonusProgram
}

func (bonusProgram *BonusProgram) SetWelcomeBonusesMode(welcomeBonusesMode WelcomeBonusesMode) *BonusProgram {
	bonusProgram.WelcomeBonusesMode = welcomeBonusesMode
	return bonusProgram
}

func (bonusProgram *BonusProgram) SetAgentTags(agentTags Slice[string]) *BonusProgram {
	bonusProgram.AgentTags = agentTags
	return bonusProgram
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
	Update(ctx context.Context, id uuid.UUID, bonusProgram *BonusProgram, params *Params) (*BonusProgram, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params *Params) (*BonusProgram, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	DeleteMany(ctx context.Context, bonusProgramList []MetaWrapper) (*DeleteManyResponse, *resty.Response, error)
}

func NewBonusProgramService(client *Client) BonusProgramService {
	e := NewEndpoint(client, "entity/bonusprogram")
	return newMainService[BonusProgram, any, any, any](e)
}
