package moysklad

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// DiscountType описывает типы, которые могут быть скидкой.
type DiscountType interface {
	AccumulationDiscount | PersonalDiscount | SpecialPriceDiscount | BonusProgram
}

// NewDiscount принимает в качестве аргумента объект, удовлетворяющий интерфейсу DiscountType.
//
// Возвращает скидку с общими полями.
func NewDiscount[T DiscountType](entity T) *Discount {
	var discount Discount
	b, _ := json.Marshal(entity)
	_ = json.Unmarshal(b, &discount)
	return &discount
}

// Discount Скидка.
//
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki
type Discount struct {
	Meta      *Meta           `json:"meta,omitempty"`
	ID        *uuid.UUID      `json:"id,omitempty"`
	AccountID *uuid.UUID      `json:"accountId,omitempty"`
	Name      *string         `json:"name,omitempty"`
	Active    *bool           `json:"active,omitempty"`
	AllAgents *bool           `json:"allAgents,omitempty"`
	AgentTags Slice[string]   `json:"agentTags,omitempty"`
	data      json.RawMessage // сырые данные для последующей десериализации в нужный тип
}

func (discount Discount) Clean() *Discount {
	return &Discount{Meta: discount.Meta}
}

func (discount Discount) GetMeta() Meta {
	return Deref(discount.Meta)
}

func (discount Discount) GetID() uuid.UUID {
	return Deref(discount.ID)
}

func (discount Discount) GetAccountID() uuid.UUID {
	return Deref(discount.AccountID)
}

func (discount Discount) GetName() string {
	return Deref(discount.Name)
}

func (discount Discount) GetActive() bool {
	return Deref(discount.Active)
}

func (discount Discount) GetAllAgents() bool {
	return Deref(discount.AllAgents)
}

func (discount Discount) GetAgentTags() Slice[string] {
	return discount.AgentTags
}

func (discount *Discount) SetMeta(meta *Meta) *Discount {
	discount.Meta = meta
	return discount
}

func (discount *Discount) SetName(name string) *Discount {
	discount.Name = &name
	return discount
}

func (discount *Discount) SetActive(active bool) *Discount {
	discount.Active = &active
	return discount
}

func (discount *Discount) SetAllAgents(allAgents bool) *Discount {
	discount.AllAgents = &allAgents
	return discount
}

func (discount *Discount) SetAgentTags(agentTags Slice[string]) *Discount {
	discount.AgentTags = agentTags
	return discount
}

func (discount Discount) String() string {
	return Stringify(discount.Meta)
}

// MetaType удовлетворяет интерфейсу MetaTyper
func (discount *Discount) MetaType() MetaType {
	return discount.Meta.GetType()
}

// Raw удовлетворяет интерфейсу RawMetaTyper
func (discount *Discount) Raw() json.RawMessage {
	return discount.data
}

// UnmarshalJSON implements the json.Unmarshaler interface.
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

// AsBonusProgram десериализует сырые данные в тип *BonusProgram
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (discount *Discount) AsBonusProgram() *BonusProgram {
	return unmarshalAsType[BonusProgram](discount)
}

// AsAccumulationDiscount десериализует сырые данные в тип *AccumulationDiscount
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (discount *Discount) AsAccumulationDiscount() *AccumulationDiscount {
	return unmarshalAsType[AccumulationDiscount](discount)
}

// AsPersonalDiscount десериализует сырые данные в тип *PersonalDiscount
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (discount *Discount) AsPersonalDiscount() *PersonalDiscount {
	return unmarshalAsType[PersonalDiscount](discount)
}

// AsSpecialPriceDiscount десериализует сырые данные в тип *SpecialPriceDiscount
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (discount *Discount) AsSpecialPriceDiscount() *SpecialPriceDiscount {
	return unmarshalAsType[SpecialPriceDiscount](discount)
}

// AccumulationDiscount Накопительная скидка.
// Ключевое слово: accumulationdiscount
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-polq-nakopitel-nyh-skidok
type AccumulationDiscount struct {
	AccountID      *uuid.UUID                `json:"accountId,omitempty"`      // ID учетной записи
	ID             *uuid.UUID                `json:"id,omitempty"`             // ID сущности
	Name           *string                   `json:"name,omitempty"`           // Наименование Скидки
	Meta           *Meta                     `json:"meta,omitempty"`           // Метаданные
	Active         *bool                     `json:"active,omitempty"`         // Индикатор, является ли скидка активной на данный момент
	AgentTags      Slice[string]             `json:"agentTags,omitempty"`      // Тэги контрагентов, к которым применяется скидка, если применяется не ко всем контрагентам
	AllProducts    *bool                     `json:"allProducts,omitempty"`    // Индикатор, действует ли скидка на все товары
	AllAgents      *bool                     `json:"allAgents,omitempty"`      // Индикатор, действует ли скидка на всех агентов
	Assortment     Assortment                `json:"assortment,omitempty"`     // Массив метаданных Товаров и Услуг, которые были выбраны для применения скидки, если та применяется не ко всем товарам
	ProductFolders *MetaArray[ProductFolder] `json:"productFolders,omitempty"` // Группы товаров которые были выбраны для применения скидки (если применяется не ко всем товарам)
	Levels         Slice[AccumulationLevel]  `json:"levels,omitempty"`         // Проценты скидок при определенной сумме продаж
}

func (accumulationDiscount AccumulationDiscount) Clean() *AccumulationDiscount {
	return &AccumulationDiscount{Meta: accumulationDiscount.Meta}
}

func (accumulationDiscount AccumulationDiscount) GetAccountID() uuid.UUID {
	return Deref(accumulationDiscount.AccountID)
}

func (accumulationDiscount AccumulationDiscount) GetID() uuid.UUID {
	return Deref(accumulationDiscount.ID)
}

func (accumulationDiscount AccumulationDiscount) GetName() string {
	return Deref(accumulationDiscount.Name)
}

func (accumulationDiscount AccumulationDiscount) GetMeta() Meta {
	return Deref(accumulationDiscount.Meta)
}

func (accumulationDiscount AccumulationDiscount) GetActive() bool {
	return Deref(accumulationDiscount.Active)
}

func (accumulationDiscount AccumulationDiscount) GetAgentTags() Slice[string] {
	return accumulationDiscount.AgentTags
}

func (accumulationDiscount AccumulationDiscount) GetAllProducts() bool {
	return Deref(accumulationDiscount.AllProducts)
}

func (accumulationDiscount AccumulationDiscount) GetAllAgents() bool {
	return Deref(accumulationDiscount.AllAgents)
}

func (accumulationDiscount AccumulationDiscount) GetAssortment() Assortment {
	return accumulationDiscount.Assortment
}

func (accumulationDiscount AccumulationDiscount) GetProductFolders() MetaArray[ProductFolder] {
	return Deref(accumulationDiscount.ProductFolders)
}

func (accumulationDiscount AccumulationDiscount) GetLevels() Slice[AccumulationLevel] {
	return accumulationDiscount.Levels
}

func (accumulationDiscount *AccumulationDiscount) SetName(name string) *AccumulationDiscount {
	accumulationDiscount.Name = &name
	return accumulationDiscount
}

func (accumulationDiscount *AccumulationDiscount) SetMeta(meta *Meta) *AccumulationDiscount {
	accumulationDiscount.Meta = meta
	return accumulationDiscount
}

func (accumulationDiscount *AccumulationDiscount) SetActive(active bool) *AccumulationDiscount {
	accumulationDiscount.Active = &active
	return accumulationDiscount
}

func (accumulationDiscount *AccumulationDiscount) SetAgentTags(agentTags Slice[string]) *AccumulationDiscount {
	accumulationDiscount.AgentTags = agentTags
	return accumulationDiscount
}

func (accumulationDiscount *AccumulationDiscount) SetAllProducts(allProducts bool) *AccumulationDiscount {
	accumulationDiscount.AllProducts = &allProducts
	return accumulationDiscount
}

func (accumulationDiscount *AccumulationDiscount) SetAllAgents(allAgents bool) *AccumulationDiscount {
	accumulationDiscount.AllAgents = &allAgents
	return accumulationDiscount
}

func (accumulationDiscount *AccumulationDiscount) SetAssortment(assortment Assortment) *AccumulationDiscount {
	accumulationDiscount.Assortment = assortment
	return accumulationDiscount
}

func (accumulationDiscount *AccumulationDiscount) SetProductFolders(productFolders Slice[ProductFolder]) *AccumulationDiscount {
	accumulationDiscount.ProductFolders = NewMetaArrayRows(productFolders)
	return accumulationDiscount
}

func (accumulationDiscount *AccumulationDiscount) SetLevels(levels Slice[AccumulationLevel]) *AccumulationDiscount {
	accumulationDiscount.Levels = levels
	return accumulationDiscount
}

func (accumulationDiscount AccumulationDiscount) String() string {
	return Stringify(accumulationDiscount)
}

func (accumulationDiscount AccumulationDiscount) MetaType() MetaType {
	return MetaTypeAccumulationDiscount
}

// AccumulationLevel Проценты скидок при определенной сумме продаж.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-levels
type AccumulationLevel struct {
	Amount   *float64 `json:"amount,omitempty"`   // Сумма накоплений в копейках
	Discount *float64 `json:"discount,omitempty"` // Процент скидки, соответствующий данной сумме
}

func (accumulationLevel AccumulationLevel) GetAmount() float64 {
	return Deref(accumulationLevel.Amount)
}

func (accumulationLevel AccumulationLevel) GetDiscount() float64 {
	return Deref(accumulationLevel.Discount)
}

func (accumulationLevel *AccumulationLevel) SetAmount(amount float64) *AccumulationLevel {
	accumulationLevel.Amount = &amount
	return accumulationLevel
}

func (accumulationLevel *AccumulationLevel) SetDiscount(discount float64) *AccumulationLevel {
	accumulationLevel.Discount = &discount
	return accumulationLevel
}

// PersonalDiscount Персональная скидка.
// Ключевое слово: personaldiscount
type PersonalDiscount struct {
	AccountID      *uuid.UUID                `json:"accountId,omitempty"`
	ID             *uuid.UUID                `json:"id,omitempty"`
	Name           *string                   `json:"name,omitempty"`
	Meta           *Meta                     `json:"meta,omitempty"`
	Active         *bool                     `json:"active,omitempty"`
	AllProducts    *bool                     `json:"allProducts,omitempty"`
	AllAgents      *bool                     `json:"allAgents,omitempty"`
	ProductFolders *MetaArray[ProductFolder] `json:"productFolders,omitempty"`
	AgentTags      Slice[string]             `json:"agentTags,omitempty"`
	Assortment     Assortment                `json:"assortment,omitempty"`
}

func (personalDiscount PersonalDiscount) Clean() *PersonalDiscount {
	return &PersonalDiscount{Meta: personalDiscount.Meta}
}

func (personalDiscount PersonalDiscount) GetAccountID() uuid.UUID {
	return Deref(personalDiscount.AccountID)
}

func (personalDiscount PersonalDiscount) GetID() uuid.UUID {
	return Deref(personalDiscount.ID)
}

func (personalDiscount PersonalDiscount) GetName() string {
	return Deref(personalDiscount.Name)
}

func (personalDiscount PersonalDiscount) GetMeta() Meta {
	return Deref(personalDiscount.Meta)
}

func (personalDiscount PersonalDiscount) GetActive() bool {
	return Deref(personalDiscount.Active)
}

func (personalDiscount PersonalDiscount) GetAgentTags() Slice[string] {
	return personalDiscount.AgentTags
}

func (personalDiscount PersonalDiscount) GetAllProducts() bool {
	return Deref(personalDiscount.AllProducts)
}

func (personalDiscount PersonalDiscount) GetAllAgents() bool {
	return Deref(personalDiscount.AllAgents)
}

func (personalDiscount PersonalDiscount) GetAssortment() Assortment {
	return personalDiscount.Assortment
}

func (personalDiscount PersonalDiscount) GetProductFolders() MetaArray[ProductFolder] {
	return Deref(personalDiscount.ProductFolders)
}

func (personalDiscount *PersonalDiscount) SetName(name string) *PersonalDiscount {
	personalDiscount.Name = &name
	return personalDiscount
}

func (personalDiscount *PersonalDiscount) SetMeta(meta *Meta) *PersonalDiscount {
	personalDiscount.Meta = meta
	return personalDiscount
}

func (personalDiscount *PersonalDiscount) SetActive(active bool) *PersonalDiscount {
	personalDiscount.Active = &active
	return personalDiscount
}

func (personalDiscount *PersonalDiscount) SetAgentTags(agentTags Slice[string]) *PersonalDiscount {
	personalDiscount.AgentTags = agentTags
	return personalDiscount
}

func (personalDiscount *PersonalDiscount) SetAllProducts(allProducts bool) *PersonalDiscount {
	personalDiscount.AllProducts = &allProducts
	return personalDiscount
}

func (personalDiscount *PersonalDiscount) SetAllAgents(allAgents bool) *PersonalDiscount {
	personalDiscount.AllAgents = &allAgents
	return personalDiscount
}

func (personalDiscount *PersonalDiscount) SetAssortment(assortment Assortment) *PersonalDiscount {
	personalDiscount.Assortment = assortment
	return personalDiscount
}

func (personalDiscount *PersonalDiscount) SetProductFolders(productFolders Slice[ProductFolder]) *PersonalDiscount {
	personalDiscount.ProductFolders = NewMetaArrayRows(productFolders)
	return personalDiscount
}

func (personalDiscount PersonalDiscount) String() string {
	return Stringify(personalDiscount)
}

func (personalDiscount PersonalDiscount) MetaType() MetaType {
	return MetaTypePersonalDiscount
}

// SpecialPriceDiscount Специальная цена.
// Ключевое слово: specialpricediscount
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-polq-spec-cen
type SpecialPriceDiscount struct {
	AllProducts    *bool                     `json:"allProducts,omitempty"`
	ID             *uuid.UUID                `json:"id,omitempty"`
	Name           *string                   `json:"name,omitempty"`
	Meta           *Meta                     `json:"meta,omitempty"`
	Active         *bool                     `json:"active,omitempty"`
	AccountID      *uuid.UUID                `json:"accountId,omitempty"`
	AllAgents      *bool                     `json:"allAgents,omitempty"`
	UsePriceType   *bool                     `json:"usePriceType,omitempty"`
	Assortment     Assortment                `json:"assortment,omitempty"`
	ProductFolders *MetaArray[ProductFolder] `json:"productFolders,omitempty"`
	Discount       *float64                  `json:"discount,omitempty"`
	SpecialPrice   *SpecialPrice             `json:"specialPrice,omitempty"`
	AgentTags      Slice[string]             `json:"agentTags,omitempty"`
}

func (specialPriceDiscount SpecialPriceDiscount) Clean() *SpecialPriceDiscount {
	return &SpecialPriceDiscount{Meta: specialPriceDiscount.Meta}
}

func (specialPriceDiscount SpecialPriceDiscount) GetAccountID() uuid.UUID {
	return Deref(specialPriceDiscount.AccountID)
}

func (specialPriceDiscount SpecialPriceDiscount) GetID() uuid.UUID {
	return Deref(specialPriceDiscount.ID)
}

func (specialPriceDiscount SpecialPriceDiscount) GetName() string {
	return Deref(specialPriceDiscount.Name)
}

func (specialPriceDiscount SpecialPriceDiscount) GetMeta() Meta {
	return Deref(specialPriceDiscount.Meta)
}

func (specialPriceDiscount SpecialPriceDiscount) GetActive() bool {
	return Deref(specialPriceDiscount.Active)
}

func (specialPriceDiscount SpecialPriceDiscount) GetAgentTags() Slice[string] {
	return specialPriceDiscount.AgentTags
}

func (specialPriceDiscount SpecialPriceDiscount) GetAllProducts() bool {
	return Deref(specialPriceDiscount.AllProducts)
}

func (specialPriceDiscount SpecialPriceDiscount) GetAllAgents() bool {
	return Deref(specialPriceDiscount.AllAgents)
}

func (specialPriceDiscount SpecialPriceDiscount) GetUsePriceType() bool {
	return Deref(specialPriceDiscount.UsePriceType)
}

func (specialPriceDiscount SpecialPriceDiscount) GetAssortment() Assortment {
	return specialPriceDiscount.Assortment
}

func (specialPriceDiscount SpecialPriceDiscount) GetProductFolders() MetaArray[ProductFolder] {
	return Deref(specialPriceDiscount.ProductFolders)
}

func (specialPriceDiscount SpecialPriceDiscount) GetDiscount() float64 {
	return Deref(specialPriceDiscount.Discount)
}

func (specialPriceDiscount SpecialPriceDiscount) GetSpecialPrice() SpecialPrice {
	return Deref(specialPriceDiscount.SpecialPrice)
}

func (specialPriceDiscount *SpecialPriceDiscount) SetName(name string) *SpecialPriceDiscount {
	specialPriceDiscount.Name = &name
	return specialPriceDiscount
}

func (specialPriceDiscount *SpecialPriceDiscount) SetMeta(meta *Meta) *SpecialPriceDiscount {
	specialPriceDiscount.Meta = meta
	return specialPriceDiscount
}

func (specialPriceDiscount *SpecialPriceDiscount) SetActive(active bool) *SpecialPriceDiscount {
	specialPriceDiscount.Active = &active
	return specialPriceDiscount
}

func (specialPriceDiscount *SpecialPriceDiscount) SetAgentTags(agentTags Slice[string]) *SpecialPriceDiscount {
	specialPriceDiscount.AgentTags = agentTags
	return specialPriceDiscount
}

func (specialPriceDiscount *SpecialPriceDiscount) SetAllProducts(allProducts bool) *SpecialPriceDiscount {
	specialPriceDiscount.AllProducts = &allProducts
	return specialPriceDiscount
}

func (specialPriceDiscount *SpecialPriceDiscount) SetAllAgents(allAgents bool) *SpecialPriceDiscount {
	specialPriceDiscount.AllAgents = &allAgents
	return specialPriceDiscount
}

func (specialPriceDiscount *SpecialPriceDiscount) SetUsePriceType(usePriceType bool) *SpecialPriceDiscount {
	specialPriceDiscount.UsePriceType = &usePriceType
	return specialPriceDiscount
}

func (specialPriceDiscount *SpecialPriceDiscount) SetAssortment(assortment Assortment) *SpecialPriceDiscount {
	specialPriceDiscount.Assortment = assortment
	return specialPriceDiscount
}

func (specialPriceDiscount *SpecialPriceDiscount) SetProductFolders(productFolders Slice[ProductFolder]) *SpecialPriceDiscount {
	specialPriceDiscount.ProductFolders = NewMetaArrayRows(productFolders)
	return specialPriceDiscount
}

func (specialPriceDiscount *SpecialPriceDiscount) SetDiscount(discount float64) *SpecialPriceDiscount {
	specialPriceDiscount.Discount = &discount
	return specialPriceDiscount
}

func (specialPriceDiscount *SpecialPriceDiscount) SetSpecialPrice(specialPrice *SpecialPrice) *SpecialPriceDiscount {
	specialPriceDiscount.SpecialPrice = specialPrice
	return specialPriceDiscount
}

func (specialPriceDiscount SpecialPriceDiscount) String() string {
	return Stringify(specialPriceDiscount)
}

func (specialPriceDiscount SpecialPriceDiscount) MetaType() MetaType {
	return MetaTypeSpecialPriceDiscount
}

// SpecialPrice Спец. цена
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-specialprice
type SpecialPrice struct {
	PriceType *PriceType `json:"priceType,omitempty"` // Тип цены
	Value     *int       `json:"value,omitempty"`     // Значение цены, если выбрано фиксированное значение
}

func (specialPrice SpecialPrice) GetPriceType() PriceType {
	return Deref(specialPrice.PriceType)
}

func (specialPrice SpecialPrice) GetValue() int {
	return Deref(specialPrice.Value)
}

func (specialPrice *SpecialPrice) SetPriceType(priceType *PriceType) *SpecialPrice {
	specialPrice.PriceType = priceType
	return specialPrice
}

func (specialPrice *SpecialPrice) SetValue(value int) *SpecialPrice {
	specialPrice.Value = &value
	return specialPrice
}

func (specialPrice SpecialPrice) String() string {
	return Stringify(specialPrice)
}

// DiscountService
// Сервис для работы со скидками.
type DiscountService interface {
	GetList(ctx context.Context, params ...*Params) (*List[Discount], *resty.Response, error)
	UpdateRoundOffDiscount(ctx context.Context, id uuid.UUID, entity *Discount) (*Discount, *resty.Response, error)
	GetAccumulationDiscounts(ctx context.Context, params ...*Params) (*List[AccumulationDiscount], *resty.Response, error)
	CreateAccumulationDiscount(ctx context.Context, accumulationDiscount *AccumulationDiscount) (*AccumulationDiscount, *resty.Response, error)
	GetAccumulationDiscountByID(ctx context.Context, id uuid.UUID, params ...*Params) (*AccumulationDiscount, *resty.Response, error)
	UpdateAccumulationDiscount(ctx context.Context, id uuid.UUID, accumulationDiscount *AccumulationDiscount) (*AccumulationDiscount, *resty.Response, error)
	DeleteAccumulationDiscount(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetPersonalDiscounts(ctx context.Context, params ...*Params) (*List[PersonalDiscount], *resty.Response, error)
	CreatePersonalDiscount(ctx context.Context, personalDiscount *PersonalDiscount) (*PersonalDiscount, *resty.Response, error)
	GetPersonalDiscountByID(ctx context.Context, id uuid.UUID, params ...*Params) (*PersonalDiscount, *resty.Response, error)
	UpdatePersonalDiscount(ctx context.Context, id uuid.UUID, personalDiscount *PersonalDiscount) (*PersonalDiscount, *resty.Response, error)
	DeletePersonalDiscount(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetSpecialPriceDiscounts(ctx context.Context, params ...*Params) (*List[SpecialPriceDiscount], *resty.Response, error)
	CreateSpecialPriceDiscount(ctx context.Context, specialPriceDiscount *SpecialPriceDiscount) (*SpecialPriceDiscount, *resty.Response, error)
	GetSpecialPriceDiscountByID(ctx context.Context, id uuid.UUID, params ...*Params) (*SpecialPriceDiscount, *resty.Response, error)
	UpdateSpecialPriceDiscount(ctx context.Context, id uuid.UUID, specialPriceDiscount *SpecialPriceDiscount) (*SpecialPriceDiscount, *resty.Response, error)
	DeleteSpecialPriceDiscount(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
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
func (service *discountService) UpdateRoundOffDiscount(ctx context.Context, id uuid.UUID, entity *Discount) (*Discount, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s", service.uri, id)
	return NewRequestBuilder[Discount](service.client, path).Put(ctx, entity)
}

// GetAccumulationDiscounts Получить все накопительные скидки.
func (service *discountService) GetAccumulationDiscounts(ctx context.Context, params ...*Params) (*List[AccumulationDiscount], *resty.Response, error) {
	path := "entity/accumulationdiscount"
	return NewRequestBuilder[List[AccumulationDiscount]](service.client, path).SetParams(params...).Get(ctx)
}

// CreateAccumulationDiscount Создать накопительную скидку.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-sozdat-nakopitel-nuu-skidku
func (service *discountService) CreateAccumulationDiscount(ctx context.Context, entity *AccumulationDiscount) (*AccumulationDiscount, *resty.Response, error) {
	path := "entity/accumulationdiscount"
	return NewRequestBuilder[AccumulationDiscount](service.client, path).Post(ctx, entity)
}

// GetAccumulationDiscountByID Получить накопительную скидку.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-poluchit-nakopitel-nuu-skidku
func (service *discountService) GetAccumulationDiscountByID(ctx context.Context, id uuid.UUID, params ...*Params) (*AccumulationDiscount, *resty.Response, error) {
	path := fmt.Sprintf("entity/accumulationdiscount/%s", id)
	return NewRequestBuilder[AccumulationDiscount](service.client, path).SetParams(params...).Get(ctx)
}

// UpdateAccumulationDiscount Изменить накопительную скидку.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-izmenit-nakopitel-nuu-skidku
func (service *discountService) UpdateAccumulationDiscount(ctx context.Context, id uuid.UUID, entity *AccumulationDiscount) (*AccumulationDiscount, *resty.Response, error) {
	path := fmt.Sprintf("entity/accumulationdiscount/%s", id)
	return NewRequestBuilder[AccumulationDiscount](service.client, path).Put(ctx, entity)
}

// DeleteAccumulationDiscount Удалить накопительную скидку.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-udalit-nakopitel-nuu-skidku
func (service *discountService) DeleteAccumulationDiscount(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("entity/accumulationdiscount/%s", id)
	return NewRequestBuilder[any](service.client, path).Delete(ctx)
}

// GetPersonalDiscounts Получить все персональные скидки.
func (service *discountService) GetPersonalDiscounts(ctx context.Context, params ...*Params) (*List[PersonalDiscount], *resty.Response, error) {
	path := "entity/personaldiscount"
	return NewRequestBuilder[List[PersonalDiscount]](service.client, path).SetParams(params...).Get(ctx)
}

// CreatePersonalDiscount Создать персональную скидку.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-sozdat-personal-nuu-skidku
func (service *discountService) CreatePersonalDiscount(ctx context.Context, entity *PersonalDiscount) (*PersonalDiscount, *resty.Response, error) {
	path := "entity/personaldiscount"
	return NewRequestBuilder[PersonalDiscount](service.client, path).Post(ctx, entity)
}

// GetPersonalDiscountByID Получить персональную скидку.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-poluchit-personal-nuu-skidku
func (service *discountService) GetPersonalDiscountByID(ctx context.Context, id uuid.UUID, params ...*Params) (*PersonalDiscount, *resty.Response, error) {
	path := fmt.Sprintf("entity/personaldiscount/%s", id)
	return NewRequestBuilder[PersonalDiscount](service.client, path).SetParams(params...).Get(ctx)
}

// UpdatePersonalDiscount Изменить персональную скидку.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-izmenit-personal-nuu-skidku
func (service *discountService) UpdatePersonalDiscount(ctx context.Context, id uuid.UUID, entity *PersonalDiscount) (*PersonalDiscount, *resty.Response, error) {
	path := fmt.Sprintf("entity/personaldiscount/%s", id)
	return NewRequestBuilder[PersonalDiscount](service.client, path).Put(ctx, entity)
}

// DeletePersonalDiscount Удалить персональную скидку.
func (service *discountService) DeletePersonalDiscount(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("entity/personaldiscount/%s", id)
	return NewRequestBuilder[any](service.client, path).Delete(ctx)
}

// GetSpecialPriceDiscounts Получить все специальные цены.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-udalit-personal-nuu-skidku
func (service *discountService) GetSpecialPriceDiscounts(ctx context.Context, params ...*Params) (*List[SpecialPriceDiscount], *resty.Response, error) {
	path := "entity/specialpricediscount"
	return NewRequestBuilder[List[SpecialPriceDiscount]](service.client, path).SetParams(params...).Get(ctx)
}

// CreateSpecialPriceDiscount Создать специальную цену.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-sozdat-special-nuu-cenu
func (service *discountService) CreateSpecialPriceDiscount(ctx context.Context, entity *SpecialPriceDiscount) (*SpecialPriceDiscount, *resty.Response, error) {
	path := "entity/specialpricediscount"
	return NewRequestBuilder[SpecialPriceDiscount](service.client, path).Post(ctx, entity)
}

// GetSpecialPriceDiscountByID Получить специальную цену.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-poluchit-special-nuu-cenu
func (service *discountService) GetSpecialPriceDiscountByID(ctx context.Context, id uuid.UUID, params ...*Params) (*SpecialPriceDiscount, *resty.Response, error) {
	path := fmt.Sprintf("entity/specialpricediscount/%s", id)
	return NewRequestBuilder[SpecialPriceDiscount](service.client, path).SetParams(params...).Get(ctx)
}

// UpdateSpecialPriceDiscount Изменить специальную цену.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-izmenit-special-nuu-cenu
func (service *discountService) UpdateSpecialPriceDiscount(ctx context.Context, id uuid.UUID, entity *SpecialPriceDiscount) (*SpecialPriceDiscount, *resty.Response, error) {
	path := fmt.Sprintf("entity/specialpricediscount/%s", id)
	return NewRequestBuilder[SpecialPriceDiscount](service.client, path).Put(ctx, entity)
}

// DeleteSpecialPriceDiscount Удалить специальную цену.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-skidki-udalit-special-nuu-cenu
func (service *discountService) DeleteSpecialPriceDiscount(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("entity/specialpricediscount/%s", id)
	return NewRequestBuilder[any](service.client, path).Delete(ctx)
}
