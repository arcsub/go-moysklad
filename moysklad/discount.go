package moysklad

import (
	"encoding/json"
	"github.com/google/uuid"
)

// Discount Скидка.
// Представляет из себя структуру из полей:
// `Meta` для определения типа сущности
// `data` для хранения сырых данных
// AccumulationDiscount | PersonalDiscount | SpecialPriceDiscount
//
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki
type Discount struct {
	Meta *Meta `json:"meta,omitempty"`
	data json.RawMessage
}

func (d Discount) String() string {
	return Stringify(d.Meta)
}

type Discounts = Slice[Discount]

// RoundOffDiscount Округление копеек.
type RoundOffDiscount struct {
	Meta      *Meta      `json:"meta,omitempty"`
	ID        *uuid.UUID `json:"id,omitempty"`
	AccountID *uuid.UUID `json:"accountId,omitempty"`
	Name      *string    `json:"name,omitempty"`
	Active    *bool      `json:"active,omitempty"`
	AllAgents *bool      `json:"allAgents,omitempty"`
	AgentTags *Tags      `json:"agentTags,omitempty"`
}

func (r RoundOffDiscount) String() string {
	return Stringify(r)
}

func (r RoundOffDiscount) MetaType() MetaType {
	return MetaTypeRoundOffDiscount
}

func (d *Discount) UnmarshalJSON(data []byte) error {
	type alias Discount
	var t alias
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	t.data = data
	*d = Discount(t)
	return nil
}

// MetaType удовлетворяет интерфейсу MetaTyper
func (d *Discount) MetaType() MetaType {
	return d.Meta.Type
}

// Data удовлетворяет интерфейсу DataMetaTyper
func (d *Discount) Data() json.RawMessage {
	return d.data
}

// AccumulationDiscount структурирует сущность в *AccumulationDiscount
// Возвращает ошибку в случае неудачи
func (d *Discount) AccumulationDiscount() (*AccumulationDiscount, error) {
	return unmarshalTo[AccumulationDiscount](d)
}

// PersonalDiscount структурирует сущность в *PersonalDiscount
// Возвращает ошибку в случае неудачи
func (d *Discount) PersonalDiscount() (*PersonalDiscount, error) {
	return unmarshalTo[PersonalDiscount](d)
}

// SpecialPriceDiscount структурирует сущность в *SpecialPriceDiscount
// Возвращает ошибку в случае неудачи
func (d *Discount) SpecialPriceDiscount() (*SpecialPriceDiscount, error) {
	return unmarshalTo[SpecialPriceDiscount](d)
}

// AccumulationDiscount Накопительная скидка.
// Ключевое слово: accumulationdiscount
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-polq-nakopitel-nyh-skidok
type AccumulationDiscount struct {
	AccountID      *uuid.UUID      `json:"accountId,omitempty"`      // ID учетной записи
	ID             *uuid.UUID      `json:"id,omitempty"`             // ID сущности
	Name           *string         `json:"name,omitempty"`           // Наименование Скидки
	Meta           *Meta           `json:"meta,omitempty"`           // Метаданные
	Active         *bool           `json:"active,omitempty"`         // Индикатор, является ли скидка активной на данный момент
	AgentTags      *Tags           `json:"agentTags,omitempty"`      // Тэги контрагентов, к которым применяется скидка, если применяется не ко всем контрагентам
	AllProducts    *bool           `json:"allProducts,omitempty"`    // Индикатор, действует ли скидка на все товары
	AllAgents      *bool           `json:"allAgents,omitempty"`      // Индикатор, действует ли скидка на всех агентов
	Assortment     *Assortment     `json:"assortment,omitempty"`     // Массив метаданных Товаров и Услуг, которые были выбраны для применения скидки, если та применяется не ко всем товарам
	ProductFolders *ProductFolders `json:"productFolders,omitempty"` // Группы товаров которые были выбраны для применения скидки (если применяется не ко всем товарам)
	Levels         *Levels         `json:"levels,omitempty"`         // Проценты скидок при определенной сумме продаж
}

func (a AccumulationDiscount) String() string {
	return Stringify(a)
}

func (a AccumulationDiscount) MetaType() MetaType {
	return MetaTypeAccumulationDiscount
}

// AccumulationLevel Проценты скидок при определенной сумме продаж.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-levels
type AccumulationLevel struct {
	Amount   *Decimal `json:"amount,omitempty"`   // Сумма накоплений в копейках
	Discount *Decimal `json:"discount,omitempty"` // Процент скидки, соответствующий данной сумме
}

type Levels = Slice[AccumulationLevel]

// PersonalDiscount Персональная скидка.
// Ключевое слово: personaldiscount
type PersonalDiscount struct {
	AccountID      *uuid.UUID      `json:"accountId,omitempty"`      // ID учетной записи
	ID             *uuid.UUID      `json:"id,omitempty"`             // ID сущности
	Name           *string         `json:"name,omitempty"`           // Наименование Скидки
	Meta           *Meta           `json:"meta,omitempty"`           // Метаданные
	Active         *bool           `json:"active,omitempty"`         // Индикатор, является ли скидка активной на данный момент
	AgentTags      *Tags           `json:"agentTags,omitempty"`      // Тэги контрагентов, к которым применяется скидка, если применяется не ко всем контрагентам
	AllProducts    *bool           `json:"allProducts,omitempty"`    // Индикатор, действует ли скидка на все товары
	AllAgents      *bool           `json:"allAgents,omitempty"`      // Индикатор, действует ли скидка на всех агентов
	Assortment     *Assortment     `json:"assortment,omitempty"`     // Массив метаданных Товаров и Услуг, которые были выбраны для применения скидки, если та применяется не ко всем товарам
	ProductFolders *ProductFolders `json:"productFolders,omitempty"` // Группы товаров которые были выбраны для применения скидки (если применяется не ко всем товарам)
}

func (p PersonalDiscount) String() string {
	return Stringify(p)
}

func (p PersonalDiscount) MetaType() MetaType {
	return MetaTypePersonalDiscount
}

// SpecialPriceDiscount Специальная цена.
// Ключевое слово: specialpricediscount
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-polq-spec-cen
type SpecialPriceDiscount struct {
	AccountID      *uuid.UUID      `json:"accountId,omitempty"`      // ID учетной записи
	ID             *uuid.UUID      `json:"id,omitempty"`             // ID сущности
	Name           *string         `json:"name,omitempty"`           // Наименование Скидки
	Meta           *Meta           `json:"meta,omitempty"`           // Метаданные
	Active         *bool           `json:"active,omitempty"`         // Индикатор, является ли скидка активной на данный момент
	AgentTags      *Tags           `json:"agentTags,omitempty"`      // Тэги контрагентов, к которым применяется скидка, если применяется не ко всем контрагентам
	AllProducts    *bool           `json:"allProducts,omitempty"`    // Индикатор, действует ли скидка на все товары
	AllAgents      *bool           `json:"allAgents,omitempty"`      // Индикатор, действует ли скидка на всех агентов
	UsePriceType   *bool           `json:"usePriceType,omitempty"`   // Использовать специальную цену
	Assortment     *Assortment     `json:"assortment,omitempty"`     // Массив метаданных Товаров и Услуг, которые были выбраны для применения скидки, если та применяется не ко всем товарам
	ProductFolders *ProductFolders `json:"productFolders,omitempty"` // Группы товаров которые были выбраны для применения скидки (если применяется не ко всем товарам)
	Discount       *Decimal        `json:"discount,omitempty"`       // Процент скидки если выбран фиксированный процент
	SpecialPrice   *SpecialPrice   `json:"specialPrice,omitempty"`   // Спец. цена (если выбран тип цен)
}

func (s SpecialPriceDiscount) String() string {
	return Stringify(s)
}

func (s SpecialPriceDiscount) MetaType() MetaType {
	return MetaTypeSpecialPriceDiscount
}

// SpecialPrice Спец. цена
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-specialprice
type SpecialPrice struct {
	PriceType *PriceType `json:"priceType,omitempty"` // Тип цены
	Value     *int       `json:"value,omitempty"`     // Значение цены, если выбрано фиксированное значение
}

func (s SpecialPrice) String() string {
	return Stringify(s)
}
