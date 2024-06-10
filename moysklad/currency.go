package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Currency Валюта.
// Ключевое слово: currency
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-valuta
type Currency struct {
	Margin         *float64       `json:"margin,omitempty"`
	ISOCode        *string        `json:"isoCode,omitempty"`
	Default        *bool          `json:"default,omitempty"`
	FullName       *string        `json:"fullName,omitempty"`
	Archived       *bool          `json:"archived,omitempty"`
	Indirect       *bool          `json:"indirect,omitempty"`
	Code           *string        `json:"code,omitempty"`
	MajorUnit      *CurrencyUnit  `json:"majorUnit,omitempty"`
	ID             *uuid.UUID     `json:"id,omitempty"`
	Meta           *Meta          `json:"meta,omitempty"`
	MinorUnit      *CurrencyUnit  `json:"minorUnit,omitempty"`
	Multiplicity   *int           `json:"multiplicity,omitempty"`
	Name           *string        `json:"name,omitempty"`
	Rate           *float64       `json:"rate,omitempty"`
	System         *bool          `json:"system,omitempty"`
	RateUpdateType RateUpdateType `json:"rateUpdateType,omitempty"`
}

func (currency Currency) Clean() *Currency {
	return &Currency{Meta: currency.Meta}
}

func (currency Currency) GetMargin() float64 {
	return Deref(currency.Margin)
}

func (currency Currency) GetISOCode() string {
	return Deref(currency.ISOCode)
}

func (currency Currency) GetDefault() bool {
	return Deref(currency.Default)
}

func (currency Currency) GetFullName() string {
	return Deref(currency.FullName)
}

func (currency Currency) GetArchived() bool {
	return Deref(currency.Archived)
}

func (currency Currency) GetIndirect() bool {
	return Deref(currency.Indirect)
}

func (currency Currency) GetCode() string {
	return Deref(currency.Code)
}

func (currency Currency) GetMajorUnit() CurrencyUnit {
	return Deref(currency.MajorUnit)
}

func (currency Currency) GetID() uuid.UUID {
	return Deref(currency.ID)
}

func (currency Currency) GetMeta() Meta {
	return Deref(currency.Meta)
}

func (currency Currency) GetMinorUnit() CurrencyUnit {
	return Deref(currency.MinorUnit)
}

func (currency Currency) GetMultiplicity() int {
	return Deref(currency.Multiplicity)
}

func (currency Currency) GetName() string {
	return Deref(currency.Name)
}

func (currency Currency) GetRate() float64 {
	return Deref(currency.Rate)
}

func (currency Currency) GetSystem() bool {
	return Deref(currency.System)
}

func (currency Currency) GetRateUpdateType() RateUpdateType {
	return currency.RateUpdateType
}

func (currency *Currency) SetMargin(margin float64) *Currency {
	currency.Margin = &margin
	return currency
}

func (currency *Currency) SetISOCode(isoCode string) *Currency {
	currency.ISOCode = &isoCode
	return currency
}

func (currency *Currency) SetFullName(fullName string) *Currency {
	currency.FullName = &fullName
	return currency
}

func (currency *Currency) SetArchived(archived bool) *Currency {
	currency.Archived = &archived
	return currency
}

func (currency *Currency) SetIndirect(indirect bool) *Currency {
	currency.Indirect = &indirect
	return currency
}

func (currency *Currency) SetCode(code string) *Currency {
	currency.Code = &code
	return currency
}

func (currency *Currency) SetMajorUnit(majorUnit *CurrencyUnit) *Currency {
	currency.MajorUnit = majorUnit
	return currency
}

func (currency *Currency) SetMeta(meta *Meta) *Currency {
	currency.Meta = meta
	return currency
}

func (currency *Currency) SetMinorUnit(minorUnit *CurrencyUnit) *Currency {
	currency.MinorUnit = minorUnit
	return currency
}

func (currency *Currency) SetMultiplicity(multiplicity int) *Currency {
	currency.Multiplicity = &multiplicity
	return currency
}

func (currency *Currency) SetName(name string) *Currency {
	currency.Name = &name
	return currency
}

func (currency *Currency) SetRate(rate float64) *Currency {
	currency.Rate = &rate
	return currency
}

func (currency Currency) String() string {
	return Stringify(currency)
}

func (currency Currency) MetaType() MetaType {
	return MetaTypeCurrency
}

// RateUpdateType Способ обновления курса.
type RateUpdateType string

const (
	RateUpdateTypeAuto   RateUpdateType = "auto"   // Автоматический
	RateUpdateTypeManual RateUpdateType = "manual" // Ручной
)

// CurrencyUnit Формы единиц.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-valuta-formy-edinic
type CurrencyUnit struct {
	S1     *string    `json:"s1,omitempty'"`
	S2     *string    `json:"s2,omitempty"`
	S5     *string    `json:"s5,omitempty"`
	Gender UnitGender `json:"gender,omitempty"`
}

func (currencyUnit CurrencyUnit) GetGender() UnitGender {
	return currencyUnit.Gender
}

func (currencyUnit CurrencyUnit) GetS1() string {
	return Deref(currencyUnit.S1)
}

func (currencyUnit CurrencyUnit) GetS2() string {
	return Deref(currencyUnit.S2)
}

func (currencyUnit CurrencyUnit) GetS5() string {
	return Deref(currencyUnit.S5)
}

func (currencyUnit *CurrencyUnit) SetGender(gender UnitGender) *CurrencyUnit {
	currencyUnit.Gender = gender
	return currencyUnit
}

func (currencyUnit *CurrencyUnit) SetS1(s1 string) *CurrencyUnit {
	currencyUnit.S1 = &s1
	return currencyUnit
}

func (currencyUnit *CurrencyUnit) SetS2(s2 string) *CurrencyUnit {
	currencyUnit.S2 = &s2
	return currencyUnit
}

func (currencyUnit *CurrencyUnit) SetS5(s3 string) *CurrencyUnit {
	currencyUnit.S5 = &s3
	return currencyUnit
}

func (currencyUnit CurrencyUnit) String() string {
	return Stringify(currencyUnit)
}

// UnitGender Грамматический род единицы валюты.
type UnitGender string

const (
	UnitGenderMasculine UnitGender = "masculine" // мужской
	UnitGenderFeminine  UnitGender = "feminine"  // женский
)

// CurrencyService
// Сервис для работы с валютами.
type CurrencyService interface {
	GetList(ctx context.Context, params *Params) (*List[Currency], *resty.Response, error)
	Create(ctx context.Context, currency *Currency, params *Params) (*Currency, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, currencyList []*Currency, params *Params) (*[]Currency, *resty.Response, error)
	DeleteMany(ctx context.Context, currencyList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params *Params) (*Currency, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, currency *Currency, params *Params) (*Currency, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)
}

func NewCurrencyService(client *Client) CurrencyService {
	e := NewEndpoint(client, "entity/currency")
	return newMainService[Currency, any, any, any](e)
}
