package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// BonusProgram Бонусная программа.
//
// Код сущности: bonusprogram
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-bonusnaq-programma
type BonusProgram struct {
	AllProducts               *bool              `json:"allProducts,omitempty"`               // Индикатор, действует ли бонусная программа на все товары
	Active                    *bool              `json:"active,omitempty"`                    // Индикатор, является ли бонусная программа активной на данный момент
	WelcomeBonusesValue       *uint              `json:"welcomeBonusesValue,omitempty"`       // Количество приветственных баллов, начисляемых участникам бонусной программы. Не может быть отрицательным. Не может быть пустым, если welcomeBonusesEnabled = true
	AllAgents                 *bool              `json:"allAgents,omitempty"`                 // Индикатор, действует ли бонусная программа на всех контрагентов
	MaxPaidRatePercents       *int               `json:"maxPaidRatePercents,omitempty"`       // Максимальный процент оплаты баллами
	EarnRateRoublesToPoint    *int               `json:"earnRateRoublesToPoint,omitempty"`    // Курс начисления
	EarnWhileRedeeming        *bool              `json:"earnWhileRedeeming,omitempty"`        // Разрешить одновременное начисление и списание бонусов. Если true - бонусы будут начислены на денежную часть покупки, даже при частичной оплате покупки баллами.
	ID                        *uuid.UUID         `json:"id,omitempty"`                        // ID Бонусной программы
	AccountID                 *uuid.UUID         `json:"accountId,omitempty"`                 // ID учётной записи
	Meta                      *Meta              `json:"meta,omitempty"`                      // Метаданные Бонусной программы
	Name                      *string            `json:"name,omitempty"`                      // Наименование Бонусной программы
	PostponedBonusesDelayDays *int               `json:"postponedBonusesDelayDays,omitempty"` // Баллы начисляются через [N] дней [Тарифная опция «Расширенная бонусная программа»]
	SpendRatePointsToRouble   *int               `json:"spendRatePointsToRouble,omitempty"`   // Курс списания
	WelcomeBonusesEnabled     *bool              `json:"welcomeBonusesEnabled,omitempty"`     // Возможность начисления приветственных баллов
	WelcomeBonusesMode        WelcomeBonusesMode `json:"welcomeBonusesMode,omitempty"`        // Условие начисления приветственных баллов. Не может быть пустым, если welcomeBonusesEnabled = true.
	AgentTags                 Slice[string]      `json:"agentTags,omitempty"`                 // Теги контрагентов, к которым применяется бонусная программа. В случае пустого значения контрагентов в результате выводится пустой массив.
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (bonusProgram BonusProgram) Clean() *BonusProgram {
	return &BonusProgram{Meta: bonusProgram.Meta}
}

// GetAllProducts возвращает флаг действия бонусной программы на все товары.
func (bonusProgram BonusProgram) GetAllProducts() bool {
	return Deref(bonusProgram.AllProducts)
}

// GetActive возвращает флаг активной бонусной программы.
func (bonusProgram BonusProgram) GetActive() bool {
	return Deref(bonusProgram.Active)
}

// GetWelcomeBonusesValue возвращает Количество приветственных баллов, начисляемых участникам бонусной программы.
func (bonusProgram BonusProgram) GetWelcomeBonusesValue() uint {
	return Deref(bonusProgram.WelcomeBonusesValue)
}

// GetAllAgents возвращает флаг действия бонусной программы на всех контрагентов.
func (bonusProgram BonusProgram) GetAllAgents() bool {
	return Deref(bonusProgram.AllAgents)
}

// GetMaxPaidRatePercents возвращает Максимальный процент оплаты баллами.
func (bonusProgram BonusProgram) GetMaxPaidRatePercents() int {
	return Deref(bonusProgram.MaxPaidRatePercents)
}

// GetEarnRateRoublesToPoint возвращает Курс начисления.
func (bonusProgram BonusProgram) GetEarnRateRoublesToPoint() int {
	return Deref(bonusProgram.EarnRateRoublesToPoint)
}

// GetEarnWhileRedeeming возвращает флаг одновременного начисления и списания бонусов.
func (bonusProgram BonusProgram) GetEarnWhileRedeeming() bool {
	return Deref(bonusProgram.EarnWhileRedeeming)
}

// GetID возвращает ID Бонусной программы.
func (bonusProgram BonusProgram) GetID() uuid.UUID {
	return Deref(bonusProgram.ID)
}

// GetAccountID возвращает ID учётной записи.
func (bonusProgram BonusProgram) GetAccountID() uuid.UUID {
	return Deref(bonusProgram.AccountID)
}

// GetMeta возвращает Метаданные Бонусной программы.
func (bonusProgram BonusProgram) GetMeta() Meta {
	return Deref(bonusProgram.Meta)
}

// GetName возвращает Наименование Бонусной программы.
func (bonusProgram BonusProgram) GetName() string {
	return Deref(bonusProgram.Name)
}

// GetPostponedBonusesDelayDays возвращает количество дней, через которое начисляются баллы. Тарифная опция «Расширенная бонусная программа».
func (bonusProgram BonusProgram) GetPostponedBonusesDelayDays() int {
	return Deref(bonusProgram.PostponedBonusesDelayDays)
}

// GetSpendRatePointsToRouble возвращает Курс списания.
func (bonusProgram BonusProgram) GetSpendRatePointsToRouble() int {
	return Deref(bonusProgram.SpendRatePointsToRouble)
}

// GetWelcomeBonusesEnabled возвращает возможность начисления приветственных баллов.
func (bonusProgram BonusProgram) GetWelcomeBonusesEnabled() bool {
	return Deref(bonusProgram.WelcomeBonusesEnabled)
}

// GetWelcomeBonusesMode возвращает Условие начисления приветственных баллов.
func (bonusProgram BonusProgram) GetWelcomeBonusesMode() WelcomeBonusesMode {
	return bonusProgram.WelcomeBonusesMode
}

// GetAgentTags возвращает Теги контрагентов, к которым применяется бонусная программа.
func (bonusProgram BonusProgram) GetAgentTags() Slice[string] {
	return bonusProgram.AgentTags
}

// SetAllProducts устанавливает Индикатор действия бонусной программы на все товары.
func (bonusProgram *BonusProgram) SetAllProducts(allProducts bool) *BonusProgram {
	bonusProgram.AllProducts = &allProducts
	return bonusProgram
}

// SetActive устанавливает флаг активности бонусной программы.
func (bonusProgram *BonusProgram) SetActive(active bool) *BonusProgram {
	bonusProgram.Active = &active
	return bonusProgram
}

// SetWelcomeBonusesValue устанавливает Количество приветственных баллов, начисляемых участникам бонусной программы. Не может быть отрицательным.
func (bonusProgram *BonusProgram) SetWelcomeBonusesValue(welcomeBonusesValue uint) *BonusProgram {
	bonusProgram.WelcomeBonusesValue = &welcomeBonusesValue
	return bonusProgram
}

// SetAllAgents устанавливает флаг действия бонусной программы на всех контрагентов.
func (bonusProgram *BonusProgram) SetAllAgents(allAgents bool) *BonusProgram {
	bonusProgram.AllAgents = &allAgents
	return bonusProgram
}

// SetMaxPaidRatePercents устанавливает Максимальный процент оплаты баллами.
func (bonusProgram *BonusProgram) SetMaxPaidRatePercents(maxPaidRatePercents int) *BonusProgram {
	bonusProgram.MaxPaidRatePercents = &maxPaidRatePercents
	return bonusProgram
}

// SetEarnRateRoublesToPoint устанавливает Курс начисления.
func (bonusProgram *BonusProgram) SetEarnRateRoublesToPoint(earnRateRoublesToPoint int) *BonusProgram {
	bonusProgram.EarnRateRoublesToPoint = &earnRateRoublesToPoint
	return bonusProgram
}

// SetEarnWhileRedeeming устанавливает флаг одновременного начисления и списания бонусов.
func (bonusProgram *BonusProgram) SetEarnWhileRedeeming(earnWhileRedeeming bool) *BonusProgram {
	bonusProgram.EarnWhileRedeeming = &earnWhileRedeeming
	return bonusProgram
}

// SetMeta устанавливает Метаданные Бонусной программы.
func (bonusProgram *BonusProgram) SetMeta(meta *Meta) *BonusProgram {
	bonusProgram.Meta = meta
	return bonusProgram
}

// SetName устанавливает Наименование Бонусной программы.
func (bonusProgram *BonusProgram) SetName(name string) *BonusProgram {
	bonusProgram.Name = &name
	return bonusProgram
}

// SetPostponedBonusesDelayDays устанавливает количество дней, через которое начисляются баллы. Тарифная опция «Расширенная бонусная программа».
func (bonusProgram *BonusProgram) SetPostponedBonusesDelayDays(postponedBonusesDelayDays int) *BonusProgram {
	bonusProgram.PostponedBonusesDelayDays = &postponedBonusesDelayDays
	return bonusProgram
}

// SetSpendRatePointsToRouble устанавливает Курс списания.
func (bonusProgram *BonusProgram) SetSpendRatePointsToRouble(spendRatePointsToRouble int) *BonusProgram {
	bonusProgram.SpendRatePointsToRouble = &spendRatePointsToRouble
	return bonusProgram
}

// SetWelcomeBonusesEnabled устанавливает флаг начисления приветственных баллов.
func (bonusProgram *BonusProgram) SetWelcomeBonusesEnabled(welcomeBonusesEnabled bool) *BonusProgram {
	bonusProgram.WelcomeBonusesEnabled = &welcomeBonusesEnabled
	return bonusProgram
}

// SetWelcomeBonusesMode устанавливает Условие начисления приветственных баллов. Не может быть пустым, если welcomeBonusesEnabled = true.
func (bonusProgram *BonusProgram) SetWelcomeBonusesMode(welcomeBonusesMode WelcomeBonusesMode) *BonusProgram {
	bonusProgram.WelcomeBonusesMode = welcomeBonusesMode
	return bonusProgram
}

// SetWelcomeBonusesModeRegistration устанавливает Условие начисления баллов участникам после регистрации в бонусной программе. Не может быть пустым, если welcomeBonusesEnabled = true.
func (bonusProgram *BonusProgram) SetWelcomeBonusesModeRegistration() *BonusProgram {
	bonusProgram.WelcomeBonusesMode = WelcomeBonusesRegistration
	return bonusProgram
}

// SetWelcomeBonusesModeFirstPurchase устанавливает Условие начисления баллов участникам бонусной программы после совершения первой покупки. Не может быть пустым, если welcomeBonusesEnabled = true.
func (bonusProgram *BonusProgram) SetWelcomeBonusesModeFirstPurchase() *BonusProgram {
	bonusProgram.WelcomeBonusesMode = WelcomeBonusesFirstPurchase
	return bonusProgram
}

// SetAgentTags устанавливает Теги контрагентов, к которым применяется бонусная программа.
func (bonusProgram *BonusProgram) SetAgentTags(agentTags ...string) *BonusProgram {
	bonusProgram.AgentTags = NewSliceFrom(agentTags)
	return bonusProgram
}

// String реализует интерфейс [fmt.Stringer].
func (bonusProgram BonusProgram) String() string {
	return Stringify(bonusProgram)
}

// MetaType возвращает тип сущности.
func (BonusProgram) MetaType() MetaType {
	return MetaTypeBonusProgram
}

// Update shortcut
func (bonusProgram BonusProgram) Update(ctx context.Context, client *Client, params ...*Params) (*BonusProgram, *resty.Response, error) {
	return NewBonusProgramService(client).Update(ctx, bonusProgram.GetID(), &bonusProgram, params...)
}

// Create shortcut
func (bonusProgram BonusProgram) Create(ctx context.Context, client *Client, params ...*Params) (*BonusProgram, *resty.Response, error) {
	return NewBonusProgramService(client).Create(ctx, &bonusProgram, params...)
}

// Delete shortcut
func (bonusProgram BonusProgram) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewBonusProgramService(client).Delete(ctx, bonusProgram.GetID())
}

// WelcomeBonusesMode Условия бонусных баллов.
//
// Возможные значения:
//   - WelcomeBonusesRegistration  – Приветственные баллы начисляются участникам после регистрации в бонусной программе.
//   - WelcomeBonusesFirstPurchase – Приветственные баллы начисляются участникам бонусной программы после совершения первой покупки.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-bonusnaq-programma-bonusnye-programmy-atributy-suschnosti-uslowiq-bonusnyh-ballow
type WelcomeBonusesMode string

const (
	WelcomeBonusesRegistration  WelcomeBonusesMode = "REGISTRATION"   // Приветственные баллы начисляются участникам после регистрации в бонусной программе.
	WelcomeBonusesFirstPurchase WelcomeBonusesMode = "FIRST_PURCHASE" // Приветственные баллы начисляются участникам бонусной программы после совершения первой покупки.
)

// BonusProgramService методы сервиса для работы с бонусными программами.
type BonusProgramService interface {
	// GetList выполняет запрос на получение списка бонусных программ.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[BonusProgram], *resty.Response, error)

	// Create выполняет запрос на создание бонусной программы.
	// Обязательные поля для заполнения:
	//	- name (имя бонусной программы)
	//	- active (активна ли бонусная программа)
	//	- allProducts (действует ли бонусная программа на все товары)
	//	- allAgents (действует ли бонусная программа на всех контрагентов)
	//	- earnRateRoublesToPoint (курс начисления)
	//	- spendRatePointsToRouble (курс списания)
	//	- maxPaidRatePercents (максимальный процент оплаты баллами)
	// Принимает контекст, бонусную программу и опционально объект параметров запроса Params.
	// Возвращает созданную бонусную программу.
	Create(ctx context.Context, bonusProgram *BonusProgram, params ...*Params) (*BonusProgram, *resty.Response, error)

	// Update выполняет запрос на изменение бонусной программы.
	// Принимает контекст, бонусную программу и опционально объект параметров запроса Params.
	// Возвращает изменённую бонусную программу.
	Update(ctx context.Context, id uuid.UUID, bonusProgram *BonusProgram, params ...*Params) (*BonusProgram, *resty.Response, error)

	// GetByID выполняет запрос на получение бонусной программы.
	// Принимает контекст, ID бонусной программы и опционально объект параметров запроса Params.
	// Возвращает бонусную программу.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*BonusProgram, *resty.Response, error)

	// Delete выполняет запрос на удаление бонусной программы.
	// Принимает контекст и ID бонусной программы.
	// Возвращает true в случае успешного удаления бонусной программы.
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление бонусных программ.
	// Принимает контекст и множество бонусных программ.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*BonusProgram) (*DeleteManyResponse, *resty.Response, error)
}

// NewBonusProgramService принимает [Client] и возвращает сервис для работы с бонусными программами.
func NewBonusProgramService(client *Client) BonusProgramService {
	return newMainService[BonusProgram, any, any, any](NewEndpoint(client, "entity/bonusprogram"))
}
