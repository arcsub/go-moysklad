package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Currency Валюта.
//
// Код сущности: currency
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-valuta
type Currency struct {
	Margin         *float64       `json:"margin,omitempty"`         // Наценка при автоматическом обновлении курса
	ISOCode        *string        `json:"isoCode,omitempty"`        // Буквенный код Валюты
	Default        *bool          `json:"default,omitempty"`        // Является ли валюта валютой учета
	FullName       *string        `json:"fullName,omitempty"`       // Полное наименование Валюты
	Archived       *bool          `json:"archived,omitempty"`       // Добавлена ли Валюта в архив
	Indirect       *bool          `json:"indirect,omitempty"`       // Признак обратного курса Валюты
	Code           *string        `json:"code,omitempty"`           // Цифровой код Валюты
	MajorUnit      *CurrencyUnit  `json:"majorUnit,omitempty"`      // Формы единиц целой части Валюты
	ID             *uuid.UUID     `json:"id,omitempty"`             // ID Валюты
	Meta           *Meta          `json:"meta,omitempty"`           // Метаданные Валюты
	MinorUnit      *CurrencyUnit  `json:"minorUnit,omitempty"`      // Формы единиц дробной части Валюты
	Multiplicity   *int           `json:"multiplicity,omitempty"`   // Кратность курса Валюты
	Name           *string        `json:"name,omitempty"`           // Краткое наименование Валюты
	Rate           *float64       `json:"rate,omitempty"`           // Курс Валюты
	System         *bool          `json:"system,omitempty"`         // Основана ли валюта на валюте из системного справочника
	RateUpdateType RateUpdateType `json:"rateUpdateType,omitempty"` // Способ обновления курса Валюты
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (currency Currency) Clean() *Currency {
	if currency.Meta == nil {
		return nil
	}
	return &Currency{Meta: currency.Meta}
}

// GetMargin возвращает Наценку при автоматическом обновлении курса.
func (currency Currency) GetMargin() float64 {
	return Deref(currency.Margin)
}

// GetISOCode возвращает Буквенный код Валюты.
func (currency Currency) GetISOCode() string {
	return Deref(currency.ISOCode)
}

// GetDefault возвращает true, если валюта является валютой по умолчанию.
func (currency Currency) GetDefault() bool {
	return Deref(currency.Default)
}

// GetFullName возвращает Полное наименование Валюты.
func (currency Currency) GetFullName() string {
	return Deref(currency.FullName)
}

// GetArchived возвращает true, если валюта находится в архиве.
func (currency Currency) GetArchived() bool {
	return Deref(currency.Archived)
}

// GetIndirect возвращает Признак обратного курса Валюты.
func (currency Currency) GetIndirect() bool {
	return Deref(currency.Indirect)
}

// GetCode возвращает Цифровой код Валюты.
func (currency Currency) GetCode() string {
	return Deref(currency.Code)
}

// GetMajorUnit возвращает Форму единиц целой части Валюты.
func (currency Currency) GetMajorUnit() CurrencyUnit {
	return Deref(currency.MajorUnit)
}

// GetID возвращает ID Валюты.
func (currency Currency) GetID() uuid.UUID {
	return Deref(currency.ID)
}

// GetMeta возвращает Метаданные Валюты.
func (currency Currency) GetMeta() Meta {
	return Deref(currency.Meta)
}

// GetMinorUnit возвращает Форму единиц дробной части Валюты.
func (currency Currency) GetMinorUnit() CurrencyUnit {
	return Deref(currency.MinorUnit)
}

// GetMultiplicity возвращает Кратность курса Валюты.
func (currency Currency) GetMultiplicity() int {
	return Deref(currency.Multiplicity)
}

// GetName возвращает Краткое наименование Валюты.
func (currency Currency) GetName() string {
	return Deref(currency.Name)
}

// GetRate возвращает Курс Валюты.
func (currency Currency) GetRate() float64 {
	return Deref(currency.Rate)
}

// GetSystem возвращает true, если Валюта основана на валюте из системного справочника.
func (currency Currency) GetSystem() bool {
	return Deref(currency.System)
}

// GetRateUpdateType возвращает Способ обновления курса Валюты.
func (currency Currency) GetRateUpdateType() RateUpdateType {
	return currency.RateUpdateType
}

// SetMargin устанавливает Наценку при автоматическом обновлении курса.
func (currency *Currency) SetMargin(margin float64) *Currency {
	currency.Margin = &margin
	return currency
}

// SetISOCode устанавливает Буквенный код Валюты.
func (currency *Currency) SetISOCode(isoCode string) *Currency {
	currency.ISOCode = &isoCode
	return currency
}

// SetFullName устанавливает Полное наименование Валюты.
func (currency *Currency) SetFullName(fullName string) *Currency {
	currency.FullName = &fullName
	return currency
}

// SetArchived устанавливает флаг нахождения валюты в архиве.
func (currency *Currency) SetArchived(archived bool) *Currency {
	currency.Archived = &archived
	return currency
}

// SetIndirect устанавливает признак обратного курса Валюты.
func (currency *Currency) SetIndirect(indirect bool) *Currency {
	currency.Indirect = &indirect
	return currency
}

// SetCode устанавливает Цифровой код Валюты.
func (currency *Currency) SetCode(code string) *Currency {
	currency.Code = &code
	return currency
}

// SetMajorUnit устанавливает Форму единиц целой части Валюты.
func (currency *Currency) SetMajorUnit(majorUnit *CurrencyUnit) *Currency {
	currency.MajorUnit = majorUnit
	return currency
}

// SetMeta устанавливает Метаданные Валюты.
func (currency *Currency) SetMeta(meta *Meta) *Currency {
	currency.Meta = meta
	return currency
}

// SetMinorUnit устанавливает Форму единиц дробной части Валюты.
func (currency *Currency) SetMinorUnit(minorUnit *CurrencyUnit) *Currency {
	currency.MinorUnit = minorUnit
	return currency
}

// SetMultiplicity устанавливает Кратность курса Валюты.
func (currency *Currency) SetMultiplicity(multiplicity int) *Currency {
	currency.Multiplicity = &multiplicity
	return currency
}

// SetName устанавливает Краткое наименование Валюты.
func (currency *Currency) SetName(name string) *Currency {
	currency.Name = &name
	return currency
}

// SetRate устанавливает Курс Валюты.
func (currency *Currency) SetRate(rate float64) *Currency {
	currency.Rate = &rate
	return currency
}

// String реализует интерфейс [fmt.Stringer].
func (currency Currency) String() string {
	return Stringify(currency)
}

// MetaType возвращает код сущности.
func (Currency) MetaType() MetaType {
	return MetaTypeCurrency
}

// Update shortcut
func (currency Currency) Update(ctx context.Context, client *Client, params ...*Params) (*Currency, *resty.Response, error) {
	return client.Entity().Currency().Update(ctx, currency.GetID(), &currency, params...)
}

// Create shortcut
func (currency Currency) Create(ctx context.Context, client *Client, params ...*Params) (*Currency, *resty.Response, error) {
	return client.Entity().Currency().Create(ctx, &currency, params...)
}

// Delete shortcut
func (currency Currency) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return client.Entity().Currency().Delete(ctx, currency.GetID())
}

// RateUpdateType Способ обновления курса.
//
// Возможные значения:
//   - RateUpdateTypeAuto   – Автоматический
//   - RateUpdateTypeManual – Ручной
type RateUpdateType string

const (
	RateUpdateTypeAuto   RateUpdateType = "auto"   // Автоматический
	RateUpdateTypeManual RateUpdateType = "manual" // Ручной
)

// CurrencyUnit Формы единиц.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-valuta-formy-edinic
type CurrencyUnit struct {
	S1     *string    `json:"s1,omitempty'"`    // Форма единицы, используемая при числительном 1
	S2     *string    `json:"s2,omitempty"`     // Форма единицы, используемая при числительном 2
	S5     *string    `json:"s5,omitempty"`     // Форма единицы, используемая при числительном 5
	Gender UnitGender `json:"gender,omitempty"` // Грамматический род единицы валюты
}

// GetGender возвращает Грамматический род единицы валюты.
func (currencyUnit CurrencyUnit) GetGender() UnitGender {
	return currencyUnit.Gender
}

// GetS1 возвращает Форму единицы, используемую при числительном 1.
func (currencyUnit CurrencyUnit) GetS1() string {
	return Deref(currencyUnit.S1)
}

// GetS2 возвращает Форму единицы, используемую при числительном 2.
func (currencyUnit CurrencyUnit) GetS2() string {
	return Deref(currencyUnit.S2)
}

// GetS5 возвращает Форму единицы, используемую при числительном 5.
func (currencyUnit CurrencyUnit) GetS5() string {
	return Deref(currencyUnit.S5)
}

// SetGender устанавливает Грамматический род единицы валюты.
func (currencyUnit *CurrencyUnit) SetGender(gender UnitGender) *CurrencyUnit {
	currencyUnit.Gender = gender
	return currencyUnit
}

// SetS1 устанавливает Форму единицы, используемую при числительном 1.
func (currencyUnit *CurrencyUnit) SetS1(s1 string) *CurrencyUnit {
	currencyUnit.S1 = &s1
	return currencyUnit
}

// SetS2 устанавливает Форму единицы, используемую при числительном 2.
func (currencyUnit *CurrencyUnit) SetS2(s2 string) *CurrencyUnit {
	currencyUnit.S2 = &s2
	return currencyUnit
}

// SetS5 устанавливает Форму единицы, используемую при числительном 5.
func (currencyUnit *CurrencyUnit) SetS5(s3 string) *CurrencyUnit {
	currencyUnit.S5 = &s3
	return currencyUnit
}

// String реализует интерфейс [fmt.Stringer].
func (currencyUnit CurrencyUnit) String() string {
	return Stringify(currencyUnit)
}

// UnitGender Грамматический род единицы валюты.
//
// Возможные значения:
//   - UnitGenderMasculine – мужской
//   - UnitGenderFeminine  – женский
type UnitGender string

const (
	UnitGenderMasculine UnitGender = "masculine" // мужской
	UnitGenderFeminine  UnitGender = "feminine"  // женский
)

// CurrencyService описывает методы сервиса для работы с валютами.
type CurrencyService interface {
	// GetList выполняет запрос на получение списка валют
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[Currency], *resty.Response, error)

	// Create выполняет запрос на создание валюты.
	// Обязательные поля для заполнения:
	//	- name (Краткое наименование Валюты)
	//	- code (Цифровой код Валюты)
	//	- isoCode (Буквенный код Валюты)
	// Принимает контекст, валюту и опционально объект параметров запроса Params.
	// Возвращает созданную валюту.
	Create(ctx context.Context, currency *Currency, params ...*Params) (*Currency, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение валют.
	// Изменяемые валюты должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список валют и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых валют.
	CreateUpdateMany(ctx context.Context, currencyList Slice[Currency], params ...*Params) (*Slice[Currency], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление валют.
	// Принимает контекст и множество валют.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*Currency) (*DeleteManyResponse, *resty.Response, error)

	// Delete выполняет запрос на удаление валюты.
	// Принимает контекст и ID валюты.
	// Возвращает «true» в случае успешного удаления валюты.
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельной валюты по ID.
	// Принимает контекст, ID валюты взаиморасчётов и опционально объект параметров запроса Params.
	// Возвращает найденную валюту.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Currency, *resty.Response, error)

	// Update выполняет запрос на изменение валюты.
	// Принимает контекст, валюту и опционально объект параметров запроса Params.
	// Возвращает изменённую валюту.
	Update(ctx context.Context, id uuid.UUID, currency *Currency, params ...*Params) (*Currency, *resty.Response, error)

	// GetNamedFilterList выполняет запрос на получение списка фильтров.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetNamedFilterList(ctx context.Context, params ...*Params) (*List[NamedFilter], *resty.Response, error)

	// GetNamedFilterByID выполняет запрос на получение отдельного фильтра по ID.
	// Принимает контекст и ID фильтра.
	// Возвращает найденный фильтр.
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)
}

const (
	EndpointCurrency = EndpointEntity + string(MetaTypeCurrency)
)

// NewCurrencyService принимает [Client] и возвращает сервис для работы с валютами.
func NewCurrencyService(client *Client) CurrencyService {
	return newMainService[Currency, any, any, any](client, EndpointCurrency)
}
