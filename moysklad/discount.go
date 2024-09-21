package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
)

// Discount Скидка.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki
type Discount struct {
	Meta        *Meta         `json:"meta,omitempty"`        // Метаданные Скидки
	ID          *uuid.UUID    `json:"id,omitempty"`          // ID Скидки
	AccountID   *uuid.UUID    `json:"accountId,omitempty"`   // ID учётной записи
	Name        *string       `json:"name,omitempty"`        // Наименование Скидки
	Active      *bool         `json:"active,omitempty"`      // Индикатор, является ли скидка активной на данный момент
	AllProducts *bool         `json:"allProducts,omitempty"` // Индикатор, действует ли скидка на все товары
	AllAgents   *bool         `json:"allAgents,omitempty"`   // Индикатор, действует ли скидка на всех контрагентов
	Assortment  Assortment    `json:"assortment,omitempty"`  // Массив метаданных Товаров и Услуг, которые были выбраны для применения скидки, если та применяется не ко всем товарам
	AgentTags   Slice[string] `json:"agentTags,omitempty"`   // Теги контрагентов, к которым применяется скидка, если применяется не ко всем контрагентам
	data        []byte        // сырые данные для последующей конвертации в нужный тип
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (discount Discount) Clean() *Discount {
	if discount.Meta == nil {
		return nil
	}
	return &Discount{Meta: discount.Meta}
}

// GetMeta возвращает Метаданные Скидки.
func (discount Discount) GetMeta() Meta {
	return Deref(discount.Meta)
}

// GetID возвращает ID Скидки.
func (discount Discount) GetID() uuid.UUID {
	return Deref(discount.ID)
}

// GetAccountID возвращает ID учётной записи.
func (discount Discount) GetAccountID() uuid.UUID {
	return Deref(discount.AccountID)
}

// GetName возвращает Наименование Скидки.
func (discount Discount) GetName() string {
	return Deref(discount.Name)
}

// GetActive возвращает Индикатор, является ли скидка активной на данный момент.
func (discount Discount) GetActive() bool {
	return Deref(discount.Active)
}

// GetAllAgents возвращает Индикатор, действует ли скидка на всех контрагентов.
func (discount Discount) GetAllAgents() bool {
	return Deref(discount.AllAgents)
}

// GetAgentTags возвращает Теги контрагентов, к которым применяется скидка, если применяется не ко всем контрагентам.
func (discount Discount) GetAgentTags() Slice[string] {
	return discount.AgentTags
}

// SetMeta устанавливает Метаданные Скидки.
func (discount *Discount) SetMeta(meta *Meta) *Discount {
	discount.Meta = meta
	return discount
}

// SetName устанавливает Наименование Скидки.
func (discount *Discount) SetName(name string) *Discount {
	discount.Name = &name
	return discount
}

// SetActive устанавливает Индикатор, является ли скидка активной на данный момент.
func (discount *Discount) SetActive(active bool) *Discount {
	discount.Active = &active
	return discount
}

// SetAllAgents устанавливает Индикатор, действует ли скидка на всех контрагентов.
func (discount *Discount) SetAllAgents(allAgents bool) *Discount {
	discount.AllAgents = &allAgents
	return discount
}

// SetAgentTags устанавливает Теги контрагентов, к которым применяется скидка, если применяется не ко всем контрагентам.
//
// Принимает множество string.
func (discount *Discount) SetAgentTags(agentTags ...string) *Discount {
	discount.AgentTags = NewSliceFrom(agentTags)
	return discount
}

// String реализует интерфейс [fmt.Stringer].
func (discount Discount) String() string {
	return Stringify(discount.Meta)
}

// MetaType возвращает код сущности.
func (discount Discount) MetaType() MetaType {
	return discount.Meta.GetType()
}

// Raw реализует интерфейс [RawMetaTyper].
func (discount *Discount) Raw() []byte {
	return discount.data
}

// UnmarshalJSON реализует интерфейс [json.Unmarshaler].
func (discount *Discount) UnmarshalJSON(data []byte) error {
	type alias Discount
	var t alias
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	t.data = data
	*discount = Discount(t)
	return nil
}

// IsBonusProgram возвращает true, если объект является типом [BonusProgram].
func (discount *Discount) IsBonusProgram() bool {
	return CheckType(discount, MetaTypeBonusProgram)
}

// IsAccumulationDiscount возвращает true, если объект является типом [AccumulationDiscount].
func (discount *Discount) IsAccumulationDiscount() bool {
	return CheckType(discount, MetaTypeAccumulationDiscount)
}

// IsPersonalDiscount возвращает true, если объект является типом [PersonalDiscount].
func (discount *Discount) IsPersonalDiscount() bool {
	return CheckType(discount, MetaTypePersonalDiscount)
}

// IsSpecialPriceDiscount возвращает true, если объект является типом [SpecialPriceDiscount].
func (discount *Discount) IsSpecialPriceDiscount() bool {
	return CheckType(discount, MetaTypeSpecialPriceDiscount)
}

// AsBonusProgram пытается привести объект к типу [BonusProgram].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [BonusProgram] или nil в случае неудачи.
func (discount *Discount) AsBonusProgram() *BonusProgram {
	return UnmarshalAsType[BonusProgram](discount)
}

// AsAccumulationDiscount пытается привести объект к типу [AccumulationDiscount].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [AccumulationDiscount] или nil в случае неудачи.
func (discount *Discount) AsAccumulationDiscount() *AccumulationDiscount {
	return UnmarshalAsType[AccumulationDiscount](discount)
}

// AsPersonalDiscount пытается привести объект к типу [PersonalDiscount].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [PersonalDiscount] или nil в случае неудачи.
func (discount *Discount) AsPersonalDiscount() *PersonalDiscount {
	return UnmarshalAsType[PersonalDiscount](discount)
}

// AsSpecialPriceDiscount пытается привести объект к типу [SpecialPriceDiscount].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [SpecialPriceDiscount] или nil в случае неудачи.
func (discount *Discount) AsSpecialPriceDiscount() *SpecialPriceDiscount {
	return UnmarshalAsType[SpecialPriceDiscount](discount)
}

// AccumulationDiscount Накопительная скидка.
//
// Код сущности: accumulationdiscount
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-polq-nakopitel-nyh-skidok
type AccumulationDiscount struct {
	AccountID      *uuid.UUID                `json:"accountId,omitempty"`      // ID учётной записи
	ID             *uuid.UUID                `json:"id,omitempty"`             // ID накопительной скидки
	Name           *string                   `json:"name,omitempty"`           // Наименование накопительной скидки
	Meta           *Meta                     `json:"meta,omitempty"`           // Метаданные накопительной скидки
	Active         *bool                     `json:"active,omitempty"`         // Индикатор, является ли скидка активной на данный момент
	AgentTags      Slice[string]             `json:"agentTags,omitempty"`      // Теги контрагентов, к которым применяется скидка, если применяется не ко всем контрагентам
	AllProducts    *bool                     `json:"allProducts,omitempty"`    // Индикатор, действует ли скидка на все товары
	AllAgents      *bool                     `json:"allAgents,omitempty"`      // Индикатор, действует ли скидка на всех агентов
	Assortment     Assortment                `json:"assortment,omitempty"`     // Массив метаданных Товаров и Услуг, которые были выбраны для применения скидки, если та применяется не ко всем товарам
	ProductFolders *MetaArray[ProductFolder] `json:"productFolders,omitempty"` // Группы товаров которые были выбраны для применения скидки (если применяется не ко всем товарам)
	Levels         Slice[AccumulationLevel]  `json:"levels,omitempty"`         // Проценты скидок при определенной сумме продаж
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (accumulationDiscount AccumulationDiscount) Clean() *AccumulationDiscount {
	if accumulationDiscount.Meta == nil {
		return nil
	}
	return &AccumulationDiscount{Meta: accumulationDiscount.Meta}
}

// GetAccountID возвращает ID учётной записи.
func (accumulationDiscount AccumulationDiscount) GetAccountID() uuid.UUID {
	return Deref(accumulationDiscount.AccountID)
}

// GetID возвращает ID накопительной скидки.
func (accumulationDiscount AccumulationDiscount) GetID() uuid.UUID {
	return Deref(accumulationDiscount.ID)
}

// GetName возвращает Наименование накопительной скидки.
func (accumulationDiscount AccumulationDiscount) GetName() string {
	return Deref(accumulationDiscount.Name)
}

// GetMeta возвращает Метаданные накопительной скидки.
func (accumulationDiscount AccumulationDiscount) GetMeta() Meta {
	return Deref(accumulationDiscount.Meta)
}

// GetActive возвращает Индикатор, является ли скидка активной на данный момент.
func (accumulationDiscount AccumulationDiscount) GetActive() bool {
	return Deref(accumulationDiscount.Active)
}

// GetAgentTags возвращает Теги контрагентов, к которым применяется скидка, если применяется не ко всем контрагентам.
func (accumulationDiscount AccumulationDiscount) GetAgentTags() Slice[string] {
	return accumulationDiscount.AgentTags
}

// GetAllProducts возвращает Индикатор, действует ли скидка на все товары.
func (accumulationDiscount AccumulationDiscount) GetAllProducts() bool {
	return Deref(accumulationDiscount.AllProducts)
}

// GetAllAgents возвращает Индикатор, действует ли скидка на всех контрагентов.
func (accumulationDiscount AccumulationDiscount) GetAllAgents() bool {
	return Deref(accumulationDiscount.AllAgents)
}

// GetAssortment возвращает Массив метаданных Товаров и Услуг, которые были выбраны для применения скидки, если та применяется не ко всем товарам.
func (accumulationDiscount AccumulationDiscount) GetAssortment() Assortment {
	return accumulationDiscount.Assortment
}

// GetProductFolders возвращает Группы товаров которые были выбраны для применения скидки (если применяется не ко всем товарам).
func (accumulationDiscount AccumulationDiscount) GetProductFolders() MetaArray[ProductFolder] {
	return Deref(accumulationDiscount.ProductFolders)
}

// GetLevels возвращает Проценты скидок при определенной сумме продаж.
func (accumulationDiscount AccumulationDiscount) GetLevels() Slice[AccumulationLevel] {
	return accumulationDiscount.Levels
}

// SetName устанавливает Наименование накопительной скидки.
func (accumulationDiscount *AccumulationDiscount) SetName(name string) *AccumulationDiscount {
	accumulationDiscount.Name = &name
	return accumulationDiscount
}

// SetMeta устанавливает Метаданные накопительной скидки.
func (accumulationDiscount *AccumulationDiscount) SetMeta(meta *Meta) *AccumulationDiscount {
	accumulationDiscount.Meta = meta
	return accumulationDiscount
}

// SetActive устанавливает Индикатор, является ли скидка активной на данный момент.
func (accumulationDiscount *AccumulationDiscount) SetActive(active bool) *AccumulationDiscount {
	accumulationDiscount.Active = &active
	return accumulationDiscount
}

// SetAgentTags устанавливает Теги контрагентов, к которым применяется скидка, если применяется не ко всем контрагентам.
//
// Принимает множество string.
func (accumulationDiscount *AccumulationDiscount) SetAgentTags(agentTags ...string) *AccumulationDiscount {
	accumulationDiscount.AgentTags = NewSliceFrom(agentTags)
	return accumulationDiscount
}

// SetAllProducts устанавливает Индикатор действия бонусной программы на все товары.
func (accumulationDiscount *AccumulationDiscount) SetAllProducts(allProducts bool) *AccumulationDiscount {
	accumulationDiscount.AllProducts = &allProducts
	return accumulationDiscount
}

// SetAllAgents устанавливает Индикатор, действует ли скидка на всех контрагентов.
func (accumulationDiscount *AccumulationDiscount) SetAllAgents(allAgents bool) *AccumulationDiscount {
	accumulationDiscount.AllAgents = &allAgents
	return accumulationDiscount
}

// SetAssortment устанавливает Метаданные товара/услуги, которую представляет собой компонент.
//
// Принимает множество объектов, реализующих интерфейс [AssortmentConverter].
func (accumulationDiscount *AccumulationDiscount) SetAssortment(entities ...AssortmentConverter) *AccumulationDiscount {
	accumulationDiscount.Assortment.Push(entities...)
	return accumulationDiscount
}

// SetProductFolders устанавливает Группы товаров которые были выбраны для применения скидки (если применяется не ко всем товарам).
//
// Принимает множество объектов [ProductFolder].
func (accumulationDiscount *AccumulationDiscount) SetProductFolders(productFolders ...*ProductFolder) *AccumulationDiscount {
	accumulationDiscount.ProductFolders = NewMetaArrayFrom(productFolders)
	return accumulationDiscount
}

// SetLevels устанавливает Проценты скидок при определенной сумме продаж.
//
// Принимает множество объектов [AccumulationLevel].
func (accumulationDiscount *AccumulationDiscount) SetLevels(levels ...*AccumulationLevel) *AccumulationDiscount {
	accumulationDiscount.Levels.Push(levels...)
	return accumulationDiscount
}

// String реализует интерфейс [fmt.Stringer].
func (accumulationDiscount AccumulationDiscount) String() string {
	return Stringify(accumulationDiscount)
}

// MetaType возвращает код сущности.
func (AccumulationDiscount) MetaType() MetaType {
	return MetaTypeAccumulationDiscount
}

// AccumulationLevel Проценты скидок при определенной сумме продаж.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-levels
type AccumulationLevel struct {
	Amount   *float64 `json:"amount,omitempty"`   // Сумма накоплений в копейках
	Discount *float64 `json:"discount,omitempty"` // Процент скидки, соответствующий данной сумме
}

// GetAmount возвращает Сумму накоплений в копейках.
func (accumulationLevel AccumulationLevel) GetAmount() float64 {
	return Deref(accumulationLevel.Amount)
}

// GetDiscount возвращает Процент скидки, соответствующий данной сумме.
func (accumulationLevel AccumulationLevel) GetDiscount() float64 {
	return Deref(accumulationLevel.Discount)
}

// SetAmount устанавливает Сумму накоплений в копейках.
func (accumulationLevel *AccumulationLevel) SetAmount(amount float64) *AccumulationLevel {
	accumulationLevel.Amount = &amount
	return accumulationLevel
}

// SetDiscount устанавливает Процент скидки, соответствующий данной сумме.
func (accumulationLevel *AccumulationLevel) SetDiscount(discount float64) *AccumulationLevel {
	accumulationLevel.Discount = &discount
	return accumulationLevel
}

// String реализует интерфейс [fmt.Stringer].
func (accumulationLevel *AccumulationLevel) String() string {
	return Stringify(accumulationLevel)
}

// PersonalDiscount Персональная скидка.
//
// Код сущности: personaldiscount
type PersonalDiscount struct {
	AccountID      *uuid.UUID                `json:"accountId,omitempty"`      // ID учётной записи
	ID             *uuid.UUID                `json:"id,omitempty"`             // ID персональной скидки
	Name           *string                   `json:"name,omitempty"`           // Наименование персональной скидки
	Meta           *Meta                     `json:"meta,omitempty"`           // Метаданные персональной скидки
	Active         *bool                     `json:"active,omitempty"`         // Индикатор, является ли скидка активной на данный момент
	AllProducts    *bool                     `json:"allProducts,omitempty"`    // Индикатор, действует ли скидка на все товары
	AllAgents      *bool                     `json:"allAgents,omitempty"`      // Индикатор, действует ли скидка на всех агентов
	ProductFolders *MetaArray[ProductFolder] `json:"productFolders,omitempty"` // Группы товаров которые были выбраны для применения скидки (если применяется не ко всем товарам)
	AgentTags      Slice[string]             `json:"agentTags,omitempty"`      // Теги контрагентов, к которым применяется скидка, если применяется не ко всем контрагентам
	Assortment     Assortment                `json:"assortment,omitempty"`     // Массив метаданных Товаров и Услуг, которые были выбраны для применения скидки, если та применяется не ко всем товарам
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (personalDiscount PersonalDiscount) Clean() *PersonalDiscount {
	if personalDiscount.Meta == nil {
		return nil
	}
	return &PersonalDiscount{Meta: personalDiscount.Meta}
}

// GetAccountID возвращает ID учётной записи.
func (personalDiscount PersonalDiscount) GetAccountID() uuid.UUID {
	return Deref(personalDiscount.AccountID)
}

// GetID возвращает ID персональной скидки.
func (personalDiscount PersonalDiscount) GetID() uuid.UUID {
	return Deref(personalDiscount.ID)
}

// GetName возвращает Наименование персональной скидки.
func (personalDiscount PersonalDiscount) GetName() string {
	return Deref(personalDiscount.Name)
}

// GetMeta возвращает Метаданные персональной скидки.
func (personalDiscount PersonalDiscount) GetMeta() Meta {
	return Deref(personalDiscount.Meta)
}

// GetActive возвращает Индикатор, является ли скидка активной на данный момент.
func (personalDiscount PersonalDiscount) GetActive() bool {
	return Deref(personalDiscount.Active)
}

// GetAgentTags возвращает Теги контрагентов, к которым применяется скидка, если применяется не ко всем контрагентам.
func (personalDiscount PersonalDiscount) GetAgentTags() Slice[string] {
	return personalDiscount.AgentTags
}

// GetAllProducts возвращает Индикатор, действует ли скидка на все товары.
func (personalDiscount PersonalDiscount) GetAllProducts() bool {
	return Deref(personalDiscount.AllProducts)
}

// GetAllAgents возвращает Индикатор, действует ли скидка на всех контрагентов.
func (personalDiscount PersonalDiscount) GetAllAgents() bool {
	return Deref(personalDiscount.AllAgents)
}

// GetAssortment возвращает Массив метаданных Товаров и Услуг, которые были выбраны для применения скидки, если та применяется не ко всем товарам.
func (personalDiscount PersonalDiscount) GetAssortment() Assortment {
	return personalDiscount.Assortment
}

// GetProductFolders возвращает Группы товаров которые были выбраны для применения скидки (если применяется не ко всем товарам).
func (personalDiscount PersonalDiscount) GetProductFolders() MetaArray[ProductFolder] {
	return Deref(personalDiscount.ProductFolders)
}

// SetName устанавливает Наименование персональной скидки.
func (personalDiscount *PersonalDiscount) SetName(name string) *PersonalDiscount {
	personalDiscount.Name = &name
	return personalDiscount
}

// SetMeta устанавливает Метаданные персональной скидки.
func (personalDiscount *PersonalDiscount) SetMeta(meta *Meta) *PersonalDiscount {
	personalDiscount.Meta = meta
	return personalDiscount
}

// SetActive устанавливает Индикатор, является ли скидка активной на данный момент.
func (personalDiscount *PersonalDiscount) SetActive(active bool) *PersonalDiscount {
	personalDiscount.Active = &active
	return personalDiscount
}

// SetAgentTags устанавливает Теги контрагентов, к которым применяется скидка, если применяется не ко всем контрагентам.
//
// Принимает множество string.
func (personalDiscount *PersonalDiscount) SetAgentTags(agentTags ...string) *PersonalDiscount {
	personalDiscount.AgentTags = NewSliceFrom(agentTags)
	return personalDiscount
}

// SetAllProducts устанавливает Индикатор действия бонусной программы на все товары.
func (personalDiscount *PersonalDiscount) SetAllProducts(allProducts bool) *PersonalDiscount {
	personalDiscount.AllProducts = &allProducts
	return personalDiscount
}

// SetAllAgents устанавливает Индикатор, действует ли скидка на всех контрагентов.
func (personalDiscount *PersonalDiscount) SetAllAgents(allAgents bool) *PersonalDiscount {
	personalDiscount.AllAgents = &allAgents
	return personalDiscount
}

// SetAssortment устанавливает Метаданные товара/услуги, которую представляет собой компонент.
//
// Принимает множество объектов, реализующих интерфейс [AssortmentConverter].
func (personalDiscount *PersonalDiscount) SetAssortment(entities ...AssortmentConverter) *PersonalDiscount {
	personalDiscount.Assortment.Push(entities...)
	return personalDiscount
}

// SetProductFolders устанавливает Группы товаров которые были выбраны для применения скидки (если применяется не ко всем товарам).
//
// Принимает множество объектов [ProductFolder].
func (personalDiscount *PersonalDiscount) SetProductFolders(productFolders ...*ProductFolder) *PersonalDiscount {
	personalDiscount.ProductFolders = NewMetaArrayFrom(productFolders)
	return personalDiscount
}

// String реализует интерфейс [fmt.Stringer].
func (personalDiscount PersonalDiscount) String() string {
	return Stringify(personalDiscount)
}

// MetaType возвращает код сущности.
func (PersonalDiscount) MetaType() MetaType {
	return MetaTypePersonalDiscount
}

// SpecialPriceDiscount Специальная цена.
//
// Код сущности: specialpricediscount
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-polq-spec-cen
type SpecialPriceDiscount struct {
	AllProducts    *bool                     `json:"allProducts,omitempty"`    // Индикатор, действует ли скидка на все товары
	ID             *uuid.UUID                `json:"id,omitempty"`             // ID специальной цены
	Name           *string                   `json:"name,omitempty"`           // Наименование специальной цены
	Meta           *Meta                     `json:"meta,omitempty"`           // Метаданные специальной цены
	Active         *bool                     `json:"active,omitempty"`         // Индикатор, является ли скидка активной на данный момент
	AccountID      *uuid.UUID                `json:"accountId,omitempty"`      // ID учётной записи
	AllAgents      *bool                     `json:"allAgents,omitempty"`      // Индикатор, действует ли скидка на всех агентов
	UsePriceType   *bool                     `json:"usePriceType,omitempty"`   // Индикатор, использовать ли специальную цену
	Assortment     Assortment                `json:"assortment,omitempty"`     // Массив метаданных Товаров и Услуг, которые были выбраны для применения скидки, если та применяется не ко всем товарам
	ProductFolders *MetaArray[ProductFolder] `json:"productFolders,omitempty"` // Группы товаров которые были выбраны для применения скидки (если применяется не ко всем товарам)
	Discount       *float64                  `json:"discount,omitempty"`       // Процент скидки если выбран фиксированный процент
	SpecialPrice   *SpecialPrice             `json:"specialPrice,omitempty"`   // Спец. цена (если выбран тип цен)
	AgentTags      Slice[string]             `json:"agentTags,omitempty"`      // Теги контрагентов, к которым применяется скидка, если применяется не ко всем контрагентам
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (specialPriceDiscount SpecialPriceDiscount) Clean() *SpecialPriceDiscount {
	if specialPriceDiscount.Meta == nil {
		return nil
	}
	return &SpecialPriceDiscount{Meta: specialPriceDiscount.Meta}
}

// GetAccountID возвращает ID учётной записи.
func (specialPriceDiscount SpecialPriceDiscount) GetAccountID() uuid.UUID {
	return Deref(specialPriceDiscount.AccountID)
}

// GetID возвращает ID специальной цены.
func (specialPriceDiscount SpecialPriceDiscount) GetID() uuid.UUID {
	return Deref(specialPriceDiscount.ID)
}

// GetName возвращает Наименование специальной цены.
func (specialPriceDiscount SpecialPriceDiscount) GetName() string {
	return Deref(specialPriceDiscount.Name)
}

// GetMeta возвращает Метаданные специальной цены.
func (specialPriceDiscount SpecialPriceDiscount) GetMeta() Meta {
	return Deref(specialPriceDiscount.Meta)
}

// GetActive возвращает Индикатор, является ли скидка активной на данный момент.
func (specialPriceDiscount SpecialPriceDiscount) GetActive() bool {
	return Deref(specialPriceDiscount.Active)
}

// GetAgentTags возвращает Теги контрагентов, к которым применяется скидка, если применяется не ко всем контрагентам.
func (specialPriceDiscount SpecialPriceDiscount) GetAgentTags() Slice[string] {
	return specialPriceDiscount.AgentTags
}

// GetAllProducts возвращает Индикатор, действует ли скидка на все товары.
func (specialPriceDiscount SpecialPriceDiscount) GetAllProducts() bool {
	return Deref(specialPriceDiscount.AllProducts)
}

// GetAllAgents возвращает Индикатор, действует ли скидка на всех контрагентов.
func (specialPriceDiscount SpecialPriceDiscount) GetAllAgents() bool {
	return Deref(specialPriceDiscount.AllAgents)
}

// GetUsePriceType возвращает Индикатор, использовать ли специальную цену.
func (specialPriceDiscount SpecialPriceDiscount) GetUsePriceType() bool {
	return Deref(specialPriceDiscount.UsePriceType)
}

// GetAssortment возвращает Массив метаданных Товаров и Услуг, которые были выбраны для применения скидки, если та применяется не ко всем товарам.
func (specialPriceDiscount SpecialPriceDiscount) GetAssortment() Assortment {
	return specialPriceDiscount.Assortment
}

// GetProductFolders возвращает Группы товаров которые были выбраны для применения скидки (если применяется не ко всем товарам).
func (specialPriceDiscount SpecialPriceDiscount) GetProductFolders() MetaArray[ProductFolder] {
	return Deref(specialPriceDiscount.ProductFolders)
}

// GetDiscount возвращает Процент скидки если выбран фиксированный процент.
func (specialPriceDiscount SpecialPriceDiscount) GetDiscount() float64 {
	return Deref(specialPriceDiscount.Discount)
}

// GetSpecialPrice возвращает Спец. цену (если выбран тип цен).
func (specialPriceDiscount SpecialPriceDiscount) GetSpecialPrice() SpecialPrice {
	return Deref(specialPriceDiscount.SpecialPrice)
}

// SetName устанавливает Наименование специальной цены.
func (specialPriceDiscount *SpecialPriceDiscount) SetName(name string) *SpecialPriceDiscount {
	specialPriceDiscount.Name = &name
	return specialPriceDiscount
}

// SetMeta устанавливает Метаданные специальной цены.
func (specialPriceDiscount *SpecialPriceDiscount) SetMeta(meta *Meta) *SpecialPriceDiscount {
	specialPriceDiscount.Meta = meta
	return specialPriceDiscount
}

// SetActive устанавливает Индикатор, является ли скидка активной на данный момент.
func (specialPriceDiscount *SpecialPriceDiscount) SetActive(active bool) *SpecialPriceDiscount {
	specialPriceDiscount.Active = &active
	return specialPriceDiscount
}

// SetAgentTags устанавливает Теги контрагентов, к которым применяется скидка, если применяется не ко всем контрагентам.
//
// Принимает множество string.
func (specialPriceDiscount *SpecialPriceDiscount) SetAgentTags(agentTags ...string) *SpecialPriceDiscount {
	specialPriceDiscount.AgentTags = NewSliceFrom(agentTags)
	return specialPriceDiscount
}

// SetAllProducts устанавливает Индикатор действия бонусной программы на все товары.
func (specialPriceDiscount *SpecialPriceDiscount) SetAllProducts(allProducts bool) *SpecialPriceDiscount {
	specialPriceDiscount.AllProducts = &allProducts
	return specialPriceDiscount
}

// SetAllAgents устанавливает Индикатор, действует ли скидка на всех контрагентов.
func (specialPriceDiscount *SpecialPriceDiscount) SetAllAgents(allAgents bool) *SpecialPriceDiscount {
	specialPriceDiscount.AllAgents = &allAgents
	return specialPriceDiscount
}

// SetUsePriceType устанавливает Индикатор, использовать ли специальную цену.
func (specialPriceDiscount *SpecialPriceDiscount) SetUsePriceType(usePriceType bool) *SpecialPriceDiscount {
	specialPriceDiscount.UsePriceType = &usePriceType
	return specialPriceDiscount
}

// SetAssortment устанавливает Метаданные товара/услуги, которую представляет собой компонент.
//
// Принимает множество объектов, реализующих интерфейс [AssortmentConverter].
func (specialPriceDiscount *SpecialPriceDiscount) SetAssortment(entities ...AssortmentConverter) *SpecialPriceDiscount {
	specialPriceDiscount.Assortment.Push(entities...)
	return specialPriceDiscount
}

// SetProductFolders устанавливает Группы товаров которые были выбраны для применения скидки (если применяется не ко всем товарам).
//
// Принимает множество объектов [ProductFolder].
func (specialPriceDiscount *SpecialPriceDiscount) SetProductFolders(productFolders ...*ProductFolder) *SpecialPriceDiscount {
	specialPriceDiscount.ProductFolders = NewMetaArrayFrom(productFolders)
	return specialPriceDiscount
}

// SetDiscount устанавливает Процент скидки, если выбран фиксированный процент.
func (specialPriceDiscount *SpecialPriceDiscount) SetDiscount(discount float64) *SpecialPriceDiscount {
	specialPriceDiscount.Discount = &discount
	return specialPriceDiscount
}

// SetSpecialPrice устанавливает Спец. цену (если выбран тип цен).
func (specialPriceDiscount *SpecialPriceDiscount) SetSpecialPrice(specialPrice *SpecialPrice) *SpecialPriceDiscount {
	specialPriceDiscount.SpecialPrice = specialPrice
	return specialPriceDiscount
}

// String реализует интерфейс [fmt.Stringer].
func (specialPriceDiscount SpecialPriceDiscount) String() string {
	return Stringify(specialPriceDiscount)
}

// MetaType возвращает код сущности.
func (SpecialPriceDiscount) MetaType() MetaType {
	return MetaTypeSpecialPriceDiscount
}

// SpecialPrice Спец. цена
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-specialprice
type SpecialPrice struct {
	PriceType *PriceType `json:"priceType,omitempty"` // Тип цены
	Value     *int       `json:"value,omitempty"`     // Значение цены, если выбрано фиксированное значение
}

// GetPriceType возвращает Тип цены.
func (specialPrice SpecialPrice) GetPriceType() PriceType {
	return Deref(specialPrice.PriceType)
}

// GetValue возвращает Значение цены, если выбрано фиксированное значение.
func (specialPrice SpecialPrice) GetValue() int {
	return Deref(specialPrice.Value)
}

// SetPriceType устанавливает Тип цены.
func (specialPrice *SpecialPrice) SetPriceType(priceType *PriceType) *SpecialPrice {
	if priceType != nil {
		specialPrice.PriceType = priceType.Clean()
	}
	return specialPrice
}

// SetValue устанавливает Значение цены, если выбрано фиксированное значение.
func (specialPrice *SpecialPrice) SetValue(value int) *SpecialPrice {
	specialPrice.Value = &value
	return specialPrice
}

// String реализует интерфейс [fmt.Stringer].
func (specialPrice SpecialPrice) String() string {
	return Stringify(specialPrice)
}

// DiscountService описывает методы сервиса для работы со скидками.
type DiscountService interface {
	// GetList выполняет запрос на получение списка всех скидок.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[Discount], *resty.Response, error)

	// GetListAll выполняет запрос на получение всех скидок в виде списка.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает список объектов.
	GetListAll(ctx context.Context, params ...*Params) (*Slice[Discount], *resty.Response, error)

	// UpdateRoundOffDiscount выполняет запрос на изменение округления копеек.
	// Принимает контекст, ID округления копеек и скидку.
	// Возвращает скидку.
	UpdateRoundOffDiscount(ctx context.Context, id uuid.UUID, entity *Discount) (*Discount, *resty.Response, error)

	// GetAccumulationDiscountList выполняет запрос на получение списка накопительных скидок.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetAccumulationDiscountList(ctx context.Context, params ...*Params) (*List[AccumulationDiscount], *resty.Response, error)

	// CreateAccumulationDiscount выполняет запрос на создание накопительной скидки.
	// Обязательные поля для заполнения:
	//	- name (Наименование скидки)
	//	- active (Активна ли скидка)
	//	- allProducts (Действует ли скидка на все товары)
	//	- allAgents (Действует ли скидка на всех контрагентов)
	// Принимает контекст, накопительную скидку и опционально объект параметров запроса Params.
	// Возвращает созданную накопительную скидку.
	CreateAccumulationDiscount(ctx context.Context, accumulationDiscount *AccumulationDiscount) (*AccumulationDiscount, *resty.Response, error)

	// GetAccumulationDiscountByID выполняет запрос на получение накопительной скидки по ID.
	// Принимает контекст, ID накопительной скидки и опционально объект параметров запроса Params.
	// Возвращает накопительную скидку.
	GetAccumulationDiscountByID(ctx context.Context, id uuid.UUID, params ...*Params) (*AccumulationDiscount, *resty.Response, error)

	// UpdateAccumulationDiscount выполняет запрос на изменение накопительной скидки.
	// Принимает контекст, накопительную скидку и опционально объект параметров запроса Params.
	// Возвращает изменённую накопительную скидку.
	UpdateAccumulationDiscount(ctx context.Context, id uuid.UUID, accumulationDiscount *AccumulationDiscount) (*AccumulationDiscount, *resty.Response, error)

	// DeleteAccumulationDiscount выполняет запрос на удаление накопительной скидки.
	// Принимает контекст и ID накопительной скидки.
	// Возвращает «true» в случае успешного удаления накопительной скидки.
	DeleteAccumulationDiscount(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetPersonalDiscountList выполняет запрос на получение списка персональных скидок.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetPersonalDiscountList(ctx context.Context, params ...*Params) (*List[PersonalDiscount], *resty.Response, error)

	// CreatePersonalDiscount выполняет запрос на создание персональной скидки.
	// Обязательные поля для заполнения:
	//	- name (Наименование скидки)
	//	- active (Активна ли скидка)
	//	- allProducts (Действует ли скидка на все товары)
	//	- allAgents (Действует ли скидка на всех контрагентов)
	// Принимает контекст, персональную скидку и опционально объект параметров запроса Params.
	// Возвращает созданную персональную скидку.
	CreatePersonalDiscount(ctx context.Context, personalDiscount *PersonalDiscount) (*PersonalDiscount, *resty.Response, error)

	// GetPersonalDiscountByID выполняет запрос на получение персональной скидки по ID.
	// Принимает контекст, ID персональной скидки и опционально объект параметров запроса Params.
	// Возвращает персональную скидку.
	GetPersonalDiscountByID(ctx context.Context, id uuid.UUID, params ...*Params) (*PersonalDiscount, *resty.Response, error)

	// UpdatePersonalDiscount выполняет запрос на изменение персональной скидки.
	// Принимает контекст, персональную скидку и опционально объект параметров запроса Params.
	// Возвращает изменённую персональную скидку.
	UpdatePersonalDiscount(ctx context.Context, id uuid.UUID, personalDiscount *PersonalDiscount) (*PersonalDiscount, *resty.Response, error)

	// DeletePersonalDiscount выполняет запрос на удаление персональной скидки.
	// Принимает контекст и ID персональной скидки.
	// Возвращает «true» в случае успешного удаления персональной скидки.
	DeletePersonalDiscount(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetSpecialPriceDiscountList выполняет запрос на получение списка специальных цен.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetSpecialPriceDiscountList(ctx context.Context, params ...*Params) (*List[SpecialPriceDiscount], *resty.Response, error)

	// CreateSpecialPriceDiscount выполняет запрос на создание специальной цены.
	// Обязательные поля для заполнения:
	//	- name (Наименование скидки)
	//	- active (Активна ли скидка)
	//	- allProducts (Действует ли скидка на все товары)
	//	- allAgents (Действует ли скидка на всех контрагентов)
	//	- usePriceType (Использовать ли специальную цену)
	// Принимает контекст, специальную цену и опционально объект параметров запроса Params.
	// Возвращает созданную специальную цену.
	CreateSpecialPriceDiscount(ctx context.Context, specialPriceDiscount *SpecialPriceDiscount) (*SpecialPriceDiscount, *resty.Response, error)

	// GetSpecialPriceDiscountByID выполняет запрос на получение специальной цены по ID.
	// Принимает контекст, ID специальной цены и опционально объект параметров запроса Params.
	// Возвращает специальную цену.
	GetSpecialPriceDiscountByID(ctx context.Context, id uuid.UUID, params ...*Params) (*SpecialPriceDiscount, *resty.Response, error)

	// UpdateSpecialPriceDiscount выполняет запрос на изменение специальной цены.
	// Принимает контекст, специальную цену и опционально объект параметров запроса Params.
	// Возвращает изменённую специальную цену.
	UpdateSpecialPriceDiscount(ctx context.Context, id uuid.UUID, specialPriceDiscount *SpecialPriceDiscount) (*SpecialPriceDiscount, *resty.Response, error)

	// DeleteSpecialPriceDiscount выполняет запрос на удаление специальной цены.
	// Принимает контекст и ID специальной цены.
	// Возвращает «true» в случае успешного удаления специальной цены.
	DeleteSpecialPriceDiscount(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
}

const (
	EndpointDiscount               = EndpointEntity + string(MetaTypeDiscount)
	EndpointAccumulationDiscount   = EndpointEntity + string(MetaTypeAccumulationDiscount)
	EndpointAccumulationDiscountID = EndpointAccumulationDiscount + "/%s"
	EndpointPersonalDiscount       = EndpointEntity + string(MetaTypePersonalDiscount)
	EndpointPersonalDiscountID     = EndpointPersonalDiscount + "/%s"
	EndpointSpecialPriceDiscount   = EndpointEntity + string(MetaTypeSpecialPriceDiscount)
	EndpointSpecialPriceDiscountID = EndpointSpecialPriceDiscount + "/%s"
)

type discountService struct {
	Endpoint
	endpointGetList[Discount]
}

func (service *discountService) UpdateRoundOffDiscount(ctx context.Context, id uuid.UUID, entity *Discount) (*Discount, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s", service.uri, id)
	return NewRequestBuilder[Discount](service.client, path).Put(ctx, entity)
}

func (service *discountService) GetAccumulationDiscountList(ctx context.Context, params ...*Params) (*List[AccumulationDiscount], *resty.Response, error) {
	return NewRequestBuilder[List[AccumulationDiscount]](service.client, EndpointAccumulationDiscount).SetParams(params...).Get(ctx)
}

func (service *discountService) CreateAccumulationDiscount(ctx context.Context, entity *AccumulationDiscount) (*AccumulationDiscount, *resty.Response, error) {
	return NewRequestBuilder[AccumulationDiscount](service.client, EndpointAccumulationDiscount).Post(ctx, entity)
}

func (service *discountService) GetAccumulationDiscountByID(ctx context.Context, id uuid.UUID, params ...*Params) (*AccumulationDiscount, *resty.Response, error) {
	path := fmt.Sprintf(EndpointAccumulationDiscountID, id)
	return NewRequestBuilder[AccumulationDiscount](service.client, path).SetParams(params...).Get(ctx)
}

func (service *discountService) UpdateAccumulationDiscount(ctx context.Context, id uuid.UUID, entity *AccumulationDiscount) (*AccumulationDiscount, *resty.Response, error) {
	path := fmt.Sprintf(EndpointAccumulationDiscountID, id)
	return NewRequestBuilder[AccumulationDiscount](service.client, path).Put(ctx, entity)
}

func (service *discountService) DeleteAccumulationDiscount(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf(EndpointAccumulationDiscountID, id)
	return NewRequestBuilder[any](service.client, path).Delete(ctx)
}

func (service *discountService) GetPersonalDiscountList(ctx context.Context, params ...*Params) (*List[PersonalDiscount], *resty.Response, error) {
	return NewRequestBuilder[List[PersonalDiscount]](service.client, EndpointPersonalDiscount).SetParams(params...).Get(ctx)
}

func (service *discountService) CreatePersonalDiscount(ctx context.Context, entity *PersonalDiscount) (*PersonalDiscount, *resty.Response, error) {
	return NewRequestBuilder[PersonalDiscount](service.client, EndpointPersonalDiscount).Post(ctx, entity)
}

func (service *discountService) GetPersonalDiscountByID(ctx context.Context, id uuid.UUID, params ...*Params) (*PersonalDiscount, *resty.Response, error) {
	path := fmt.Sprintf(EndpointPersonalDiscountID, id)
	return NewRequestBuilder[PersonalDiscount](service.client, path).SetParams(params...).Get(ctx)
}

func (service *discountService) UpdatePersonalDiscount(ctx context.Context, id uuid.UUID, entity *PersonalDiscount) (*PersonalDiscount, *resty.Response, error) {
	path := fmt.Sprintf(EndpointPersonalDiscountID, id)
	return NewRequestBuilder[PersonalDiscount](service.client, path).Put(ctx, entity)
}

func (service *discountService) DeletePersonalDiscount(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf(EndpointPersonalDiscountID, id)
	return NewRequestBuilder[any](service.client, path).Delete(ctx)
}

func (service *discountService) GetSpecialPriceDiscountList(ctx context.Context, params ...*Params) (*List[SpecialPriceDiscount], *resty.Response, error) {
	return NewRequestBuilder[List[SpecialPriceDiscount]](service.client, EndpointSpecialPriceDiscount).SetParams(params...).Get(ctx)
}

func (service *discountService) CreateSpecialPriceDiscount(ctx context.Context, entity *SpecialPriceDiscount) (*SpecialPriceDiscount, *resty.Response, error) {
	return NewRequestBuilder[SpecialPriceDiscount](service.client, EndpointSpecialPriceDiscount).Post(ctx, entity)
}

func (service *discountService) GetSpecialPriceDiscountByID(ctx context.Context, id uuid.UUID, params ...*Params) (*SpecialPriceDiscount, *resty.Response, error) {
	path := fmt.Sprintf(EndpointSpecialPriceDiscountID, id)
	return NewRequestBuilder[SpecialPriceDiscount](service.client, path).SetParams(params...).Get(ctx)
}

func (service *discountService) UpdateSpecialPriceDiscount(ctx context.Context, id uuid.UUID, entity *SpecialPriceDiscount) (*SpecialPriceDiscount, *resty.Response, error) {
	path := fmt.Sprintf(EndpointSpecialPriceDiscountID, id)
	return NewRequestBuilder[SpecialPriceDiscount](service.client, path).Put(ctx, entity)
}

func (service *discountService) DeleteSpecialPriceDiscount(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf(EndpointSpecialPriceDiscountID, id)
	return NewRequestBuilder[any](service.client, path).Delete(ctx)
}

// NewDiscountService принимает [Client] и возвращает сервис для работы со скидками.
func NewDiscountService(client *Client) DiscountService {
	e := NewEndpoint(client, EndpointDiscount)
	return &discountService{
		Endpoint:        e,
		endpointGetList: endpointGetList[Discount]{e},
	}
}
