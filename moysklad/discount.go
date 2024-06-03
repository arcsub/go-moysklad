package moysklad

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
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

// UnmarshalJSON implements the json.Unmarshaler interface.
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

// Raw удовлетворяет интерфейсу RawMetaTyper
func (d *Discount) Raw() json.RawMessage {
	return d.data
}

// AccumulationDiscount десериализует сырые данные в тип *AccumulationDiscount
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (d *Discount) AccumulationDiscount() *AccumulationDiscount {
	return unmarshalAsType[AccumulationDiscount](d)
}

// PersonalDiscount десериализует сырые данные в тип *PersonalDiscount
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (d *Discount) PersonalDiscount() *PersonalDiscount {
	return unmarshalAsType[PersonalDiscount](d)
}

// SpecialPriceDiscount десериализует сырые данные в тип *SpecialPriceDiscount
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (d *Discount) SpecialPriceDiscount() *SpecialPriceDiscount {
	return unmarshalAsType[SpecialPriceDiscount](d)
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

// DiscountService
// Сервис для работы со скидками.
type DiscountService interface {
	GetList(ctx context.Context, params *Params) (*List[Discount], *resty.Response, error)
	UpdateRoundOffDiscount(ctx context.Context, id *uuid.UUID, entity *RoundOffDiscount) (*RoundOffDiscount, *resty.Response, error)
	GetAccumulationDiscounts(ctx context.Context, params *Params) (*List[AccumulationDiscount], *resty.Response, error)
	CreateAccumulationDiscount(ctx context.Context, accumulationDiscount *AccumulationDiscount) (*AccumulationDiscount, *resty.Response, error)
	GetByIdAccumulationDiscount(ctx context.Context, id *uuid.UUID, params *Params) (*AccumulationDiscount, *resty.Response, error)
	UpdateAccumulationDiscount(ctx context.Context, id *uuid.UUID, accumulationDiscount *AccumulationDiscount) (*AccumulationDiscount, *resty.Response, error)
	DeleteAccumulationDiscount(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetPersonalDiscounts(ctx context.Context, params *Params) (*List[PersonalDiscount], *resty.Response, error)
	CreatePersonalDiscount(ctx context.Context, personalDiscount *PersonalDiscount) (*PersonalDiscount, *resty.Response, error)
	GetByIdPersonalDiscount(ctx context.Context, id *uuid.UUID, params *Params) (*PersonalDiscount, *resty.Response, error)
	UpdatePersonalDiscount(ctx context.Context, id *uuid.UUID, personalDiscount *PersonalDiscount) (*PersonalDiscount, *resty.Response, error)
	DeletePersonalDiscount(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetSpecialPriceDiscounts(ctx context.Context, params *Params) (*List[SpecialPriceDiscount], *resty.Response, error)
	CreateSpecialPriceDiscount(ctx context.Context, specialPriceDiscount *SpecialPriceDiscount) (*SpecialPriceDiscount, *resty.Response, error)
	GetByIdSpecialPriceDiscount(ctx context.Context, id *uuid.UUID, params *Params) (*SpecialPriceDiscount, *resty.Response, error)
	UpdateSpecialPriceDiscount(ctx context.Context, id *uuid.UUID, specialPriceDiscount *SpecialPriceDiscount) (*SpecialPriceDiscount, *resty.Response, error)
	DeleteSpecialPriceDiscount(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

type discountService struct {
	Endpoint
	endpointGetList[Discount]
}

func NewDiscountService(client *Client) DiscountService {
	e := NewEndpoint(client, "entity/discount")
	return &discountService{
		Endpoint:        e,
		endpointGetList: endpointGetList[Discount]{e},
	}
}

// UpdateRoundOffDiscount Изменить округление копеек.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-izmenit-okruglenie-kopeek
func (s *discountService) UpdateRoundOffDiscount(ctx context.Context, id *uuid.UUID, entity *RoundOffDiscount) (*RoundOffDiscount, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s", s.uri, id)
	return NewRequestBuilder[RoundOffDiscount](s.client, path).Put(ctx, entity)
}

// GetAccumulationDiscounts Получить все накопительные скидки.
func (s *discountService) GetAccumulationDiscounts(ctx context.Context, params *Params) (*List[AccumulationDiscount], *resty.Response, error) {
	path := "entity/accumulationdiscount"
	return NewRequestBuilder[List[AccumulationDiscount]](s.client, path).SetParams(params).Get(ctx)
}

// CreateAccumulationDiscount Создать накопительную скидку.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-sozdat-nakopitel-nuu-skidku
func (s *discountService) CreateAccumulationDiscount(ctx context.Context, entity *AccumulationDiscount) (*AccumulationDiscount, *resty.Response, error) {
	path := "entity/accumulationdiscount"
	return NewRequestBuilder[AccumulationDiscount](s.client, path).Post(ctx, entity)
}

// GetByIdAccumulationDiscount Получить накопительную скидку.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-poluchit-nakopitel-nuu-skidku
func (s *discountService) GetByIdAccumulationDiscount(ctx context.Context, id *uuid.UUID, params *Params) (*AccumulationDiscount, *resty.Response, error) {
	path := fmt.Sprintf("entity/accumulationdiscount/%s", id)
	return NewRequestBuilder[AccumulationDiscount](s.client, path).SetParams(params).Get(ctx)
}

// UpdateAccumulationDiscount Изменить накопительную скидку.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-izmenit-nakopitel-nuu-skidku
func (s *discountService) UpdateAccumulationDiscount(ctx context.Context, id *uuid.UUID, entity *AccumulationDiscount) (*AccumulationDiscount, *resty.Response, error) {
	path := fmt.Sprintf("entity/accumulationdiscount/%s", id)
	return NewRequestBuilder[AccumulationDiscount](s.client, path).Put(ctx, entity)
}

// DeleteAccumulationDiscount Удалить накопительную скидку.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-udalit-nakopitel-nuu-skidku
func (s *discountService) DeleteAccumulationDiscount(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("entity/accumulationdiscount/%s", id)
	return NewRequestBuilder[any](s.client, path).Delete(ctx)
}

// GetPersonalDiscounts Получить все персональные скидки.
func (s *discountService) GetPersonalDiscounts(ctx context.Context, params *Params) (*List[PersonalDiscount], *resty.Response, error) {
	path := "entity/personaldiscount"
	return NewRequestBuilder[List[PersonalDiscount]](s.client, path).SetParams(params).Get(ctx)
}

// CreatePersonalDiscount Создать персональную скидку.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-sozdat-personal-nuu-skidku
func (s *discountService) CreatePersonalDiscount(ctx context.Context, entity *PersonalDiscount) (*PersonalDiscount, *resty.Response, error) {
	path := "entity/personaldiscount"
	return NewRequestBuilder[PersonalDiscount](s.client, path).Post(ctx, entity)
}

// GetByIdPersonalDiscount Получить персональную скидку.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-poluchit-personal-nuu-skidku
func (s *discountService) GetByIdPersonalDiscount(ctx context.Context, id *uuid.UUID, params *Params) (*PersonalDiscount, *resty.Response, error) {
	path := fmt.Sprintf("entity/personaldiscount/%s", id)
	return NewRequestBuilder[PersonalDiscount](s.client, path).SetParams(params).Get(ctx)
}

// UpdatePersonalDiscount Изменить персональную скидку.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-izmenit-personal-nuu-skidku
func (s *discountService) UpdatePersonalDiscount(ctx context.Context, id *uuid.UUID, entity *PersonalDiscount) (*PersonalDiscount, *resty.Response, error) {
	path := fmt.Sprintf("entity/personaldiscount/%s", id)
	return NewRequestBuilder[PersonalDiscount](s.client, path).Put(ctx, entity)
}

// DeletePersonalDiscount Удалить персональную скидку.
func (s *discountService) DeletePersonalDiscount(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("entity/personaldiscount/%s", id)
	return NewRequestBuilder[any](s.client, path).Delete(ctx)
}

// GetSpecialPriceDiscounts Получить все специальные цены.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-udalit-personal-nuu-skidku
func (s *discountService) GetSpecialPriceDiscounts(ctx context.Context, params *Params) (*List[SpecialPriceDiscount], *resty.Response, error) {
	path := "entity/specialpricediscount"
	return NewRequestBuilder[List[SpecialPriceDiscount]](s.client, path).SetParams(params).Get(ctx)
}

// CreateSpecialPriceDiscount Создать специальную цену.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-sozdat-special-nuu-cenu
func (s *discountService) CreateSpecialPriceDiscount(ctx context.Context, entity *SpecialPriceDiscount) (*SpecialPriceDiscount, *resty.Response, error) {
	path := "entity/specialpricediscount"
	return NewRequestBuilder[SpecialPriceDiscount](s.client, path).Post(ctx, entity)
}

// GetByIdSpecialPriceDiscount Получить специальную цену.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-poluchit-special-nuu-cenu
func (s *discountService) GetByIdSpecialPriceDiscount(ctx context.Context, id *uuid.UUID, params *Params) (*SpecialPriceDiscount, *resty.Response, error) {
	path := fmt.Sprintf("entity/specialpricediscount/%s", id)
	return NewRequestBuilder[SpecialPriceDiscount](s.client, path).SetParams(params).Get(ctx)
}

// UpdateSpecialPriceDiscount Изменить специальную цену.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-izmenit-special-nuu-cenu
func (s *discountService) UpdateSpecialPriceDiscount(ctx context.Context, id *uuid.UUID, entity *SpecialPriceDiscount) (*SpecialPriceDiscount, *resty.Response, error) {
	path := fmt.Sprintf("entity/specialpricediscount/%s", id)
	return NewRequestBuilder[SpecialPriceDiscount](s.client, path).Put(ctx, entity)
}

// DeleteSpecialPriceDiscount Удалить специальную цену.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-udalit-special-nuu-cenu
func (s *discountService) DeleteSpecialPriceDiscount(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("entity/specialpricediscount/%s", id)
	return NewRequestBuilder[any](s.client, path).Delete(ctx)
}
