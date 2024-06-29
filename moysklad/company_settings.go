package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// CompanySettings Настройки компании.
//
// Код сущности: companysettings
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-nastrojki-kompanii
type CompanySettings struct {
	Meta                     *Meta            `json:"meta,omitempty"`                     // Метаданные Настроек компании
	CheckMinPrice            *bool            `json:"checkMinPrice,omitempty"`            // Автоматически устанавливать минимальную цену. Если включено, при сохранении документов продажи с ценами меньше минимальных цен (указанных в карточках товара) цены будут автоматически увеличены до минимальных.
	CheckShippingStock       *bool            `json:"checkShippingStock,omitempty"`       // Запретить отгрузку отсутствующих товаров. Если запрет установлен (true значение), пользователи не смогут провести отгрузку со склада отсутствующих товаров.
	CompanyAddress           *string          `json:"companyAddress,omitempty"`           // Адрес компании для электронных писем
	Currency                 *Currency        `json:"currency,omitempty"`                 // Стандартная валюта
	GlobalOperationNumbering *bool            `json:"globalOperationNumbering,omitempty"` // Использовать сквозную нумерацию документов. Если проставлен true, будет установлена сквозная нумерация за всю историю, иначе нумерация документов будет начинаться заново каждый календарный год.
	UseCompanyAddress        *bool            `json:"useCompanyAddress,omitempty"`        // Использовать адрес компании для электронных писем. Если включено, письма будут отправляться с адреса, указанного в companyAddress, иначе письма будут отправляться с адреса пользователя
	UseRecycleBin            *bool            `json:"useRecycleBin,omitempty"`            // Использовать корзину. Если включено, то все документы при удалении будут помещаться в корзину. Также появится возможность восстанавливать ошибочно удаленные документы.
	DiscountStrategy         DiscountStrategy `json:"discountStrategy,omitempty"`         // Совместное применение скидок
	AccountCountry           AccountCountry   `json:"accountCountry,omitempty"`           // Передается для информации о том, какая страновая конфигурация активна на аккаунте пользователя
	PriceTypes               Slice[PriceType] `json:"priceTypes,omitempty"`               // Коллекция всех существующих типов цен
}

// GetMeta возвращает Метаданные Настроек компании.
func (companySettings CompanySettings) GetMeta() Meta {
	return Deref(companySettings.Meta)
}

// GetCheckMinPrice возвращает флаг Автоматической установки минимальной цены.
func (companySettings CompanySettings) GetCheckMinPrice() bool {
	return Deref(companySettings.CheckShippingStock)
}

// GetCheckShippingStock возвращает флаг Запрета отгрузки отсутствующих товаров.
func (companySettings CompanySettings) GetCheckShippingStock() bool {
	return Deref(companySettings.CheckShippingStock)
}

// GetCompanyAddress возвращает Адрес компании для электронных писем.
func (companySettings CompanySettings) GetCompanyAddress() string {
	return Deref(companySettings.CompanyAddress)
}

// GetCurrency возвращает Стандартную валюту.
func (companySettings CompanySettings) GetCurrency() Currency {
	return Deref(companySettings.Currency)
}

// GetGlobalOperationNumbering возвращает флаг Использования сквозной нумерации документов.
func (companySettings CompanySettings) GetGlobalOperationNumbering() bool {
	return Deref(companySettings.GlobalOperationNumbering)
}

// GetPriceTypes возвращает Коллекцию всех существующих типов цен.
func (companySettings CompanySettings) GetPriceTypes() Slice[PriceType] {
	return companySettings.PriceTypes
}

// GetUseCompanyAddress возвращает флаг Использования адреса компании для электронных писем.
func (companySettings CompanySettings) GetUseCompanyAddress() bool {
	return Deref(companySettings.UseCompanyAddress)
}

// GetUseRecycleBin возвращает флаг Использования корзины.
func (companySettings CompanySettings) GetUseRecycleBin() bool {
	return Deref(companySettings.UseRecycleBin)
}

// GetDiscountStrategy возвращает Совместное применение скидок.
func (companySettings CompanySettings) GetDiscountStrategy() DiscountStrategy {
	return companySettings.DiscountStrategy
}

// GetAccountCountry возвращает конфигурацию страны.
func (companySettings CompanySettings) GetAccountCountry() AccountCountry {
	return companySettings.AccountCountry
}

// SetCheckMinPrice устанавливает флаг Автоматической установки минимальной цены.
func (companySettings *CompanySettings) SetCheckMinPrice(checkMinPrice bool) *CompanySettings {
	companySettings.CheckMinPrice = &checkMinPrice
	return companySettings
}

// SetCheckShippingStock устанавливает флаг Запрета отгрузки отсутствующих товаров.
func (companySettings *CompanySettings) SetCheckShippingStock(checkShippingStock bool) *CompanySettings {
	companySettings.CheckShippingStock = &checkShippingStock
	return companySettings
}

// SetCompanyAddress устанавливает Адрес компании для электронных писем.
func (companySettings *CompanySettings) SetCompanyAddress(companyAddress string) *CompanySettings {
	companySettings.CompanyAddress = &companyAddress
	return companySettings
}

// SetGlobalOperationNumbering устанавливает флаг Использования сквозной нумерации документов.
func (companySettings *CompanySettings) SetGlobalOperationNumbering(globalOperationNumbering bool) *CompanySettings {
	companySettings.GlobalOperationNumbering = &globalOperationNumbering
	return companySettings
}

// SetUseCompanyAddress устанавливает флаг Использования адреса компании для электронных писем.
func (companySettings *CompanySettings) SetUseCompanyAddress(useCompanyAddress bool) *CompanySettings {
	companySettings.UseCompanyAddress = &useCompanyAddress
	return companySettings
}

// SetUseRecycleBin устанавливает флаг Использования корзины.
func (companySettings *CompanySettings) SetUseRecycleBin(useRecycleBin bool) *CompanySettings {
	companySettings.UseRecycleBin = &useRecycleBin
	return companySettings
}

// SetDiscountStrategy устанавливает Совместное применение скидок
func (companySettings *CompanySettings) SetDiscountStrategy(discountStrategy DiscountStrategy) *CompanySettings {
	companySettings.DiscountStrategy = discountStrategy
	return companySettings
}

// SetDiscountStrategyBySum устанавливает Совместное применение скидок по Сумме скидок.
//
// Должна действовать сумма скидок.
func (companySettings *CompanySettings) SetDiscountStrategyBySum() *CompanySettings {
	companySettings.DiscountStrategy = DiscountStrategyBySum
	return companySettings
}

// SetDiscountStrategyByPriority устанавливает Совместное применение скидок по Приоритету.
//
// Должна действовать одна, наиболее выгодная для покупателя скидка.
func (companySettings *CompanySettings) SetDiscountStrategyByPriority() *CompanySettings {
	companySettings.DiscountStrategy = DiscountStrategyByPriority
	return companySettings
}

// String реализует интерфейс [fmt.Stringer].
func (companySettings CompanySettings) String() string {
	return Stringify(companySettings)
}

// MetaType возвращает код сущности.
func (CompanySettings) MetaType() MetaType {
	return MetaTypeCompanySettings
}

// DiscountStrategy Совместное применение скидок.
//
// Возможные значения:
//   - DiscountStrategyBySum      – Сумма скидок (должна действовать сумма скидок)
//   - DiscountStrategyByPriority – Приоритетная (должна действовать одна, наиболее выгодная для покупателя скидка)
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-nastrojki-kompanii-sowmestnoe-primenenie-skidok
type DiscountStrategy string

const (
	DiscountStrategyBySum      DiscountStrategy = "bySum"      // Сумма скидок (должна действовать сумма скидок)
	DiscountStrategyByPriority DiscountStrategy = "byPriority" // Приоритетная (должна действовать одна, наиболее выгодная для покупателя скидка)
)

// AccountCountry конфигурация страны.
//
// Возможные значения:
//   - AccountCountryRU – RU
//   - AccountCountryBY – BY
//   - AccountCountryKZ – KZ
type AccountCountry string

const (
	AccountCountryRU AccountCountry = "RU"
	AccountCountryBY AccountCountry = "BY"
	AccountCountryKZ AccountCountry = "KZ"
)

// ContextCompanySettingsService описывает методы сервиса для работы с настройками компании.
type ContextCompanySettingsService interface {
	// Get выполняет запрос на получение настроек компании.
	// Принимает контекст.
	// Возвращает настройки компании.
	Get(ctx context.Context) (*CompanySettings, *resty.Response, error)

	// Update выполняет запрос на изменение настроек компании.
	// Принимает контекст и настройки компании.
	// Возвращает изменённые настройки компании.
	Update(ctx context.Context, settings *CompanySettings) (*CompanySettings, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных настроек компании.
	// Принимает контекст.
	// Возвращает метаданные настроек компании.
	GetMetadata(ctx context.Context) (*MetadataCompanySettings, *resty.Response, error)

	// GetPriceTypes выполняет запрос на получение типов цен в виде списка.
	// Принимает контекст.
	// Возвращает список всех типов цен.
	GetPriceTypes(ctx context.Context) (*Slice[PriceType], *resty.Response, error)

	// CreatePriceType выполняет запрос на создания типа цен.
	// Принимает контекст и тип цен.
	// Возвращает список типов цен.
	CreatePriceType(ctx context.Context, priceType *PriceType) (*Slice[PriceType], *resty.Response, error)

	// UpdatePriceTypeMany выполняет запрос на изменение списка типов цен.
	// Принимает контекст и типы цен в виде списка.
	// Возвращает изменённый список типов цен.
	UpdatePriceTypeMany(ctx context.Context, priceTypes Slice[PriceType]) (*Slice[PriceType], *resty.Response, error)

	// GetPriceTypeByID выполняет запрос на получение отдельного типа цен по ID.
	// Принимает контекст и ID типа цен.
	// Возвращает найденный тип цен.
	GetPriceTypeByID(ctx context.Context, id uuid.UUID) (*PriceType, *resty.Response, error)

	// GetPriceTypeDefault выполняет запрос на получение типа цен по умолчанию.
	// Принимает контекст.
	// Возвращает тип цен по умолчанию.
	GetPriceTypeDefault(ctx context.Context) (*PriceType, *resty.Response, error)
}

type contextCompanySettingsService struct {
	Endpoint
	endpointUpdate[CompanySettings]
	endpointMetadata[MetadataCompanySettings]
}

// NewContextCompanySettingsService принимает [Client] и возвращает сервис для работы с настройками компании.
func NewContextCompanySettingsService(client *Client) ContextCompanySettingsService {
	e := NewEndpoint(client, "context/companysettings")
	return &contextCompanySettingsService{
		Endpoint:         e,
		endpointUpdate:   endpointUpdate[CompanySettings]{e},
		endpointMetadata: endpointMetadata[MetadataCompanySettings]{e},
	}
}

func (service *contextCompanySettingsService) Get(ctx context.Context) (*CompanySettings, *resty.Response, error) {
	return NewRequestBuilder[CompanySettings](service.client, service.uri).Get(ctx)
}

func (service *contextCompanySettingsService) Update(ctx context.Context, settings *CompanySettings) (*CompanySettings, *resty.Response, error) {
	return NewRequestBuilder[CompanySettings](service.client, service.uri).Put(ctx, settings)
}

func (service *contextCompanySettingsService) GetPriceTypes(ctx context.Context) (*Slice[PriceType], *resty.Response, error) {
	path := "context/companysettings/pricetype"
	return NewRequestBuilder[Slice[PriceType]](service.client, path).Get(ctx)
}

func (service *contextCompanySettingsService) CreatePriceType(ctx context.Context, priceType *PriceType) (*Slice[PriceType], *resty.Response, error) {
	priceTypes, resp, err := service.GetPriceTypes(ctx) // получаем список всех типов цен
	if err != nil {
		return nil, resp, err
	}
	priceTypes.Push(priceType) // добавляем новый тип цен в конец списка
	path := "context/companysettings/pricetype"
	return NewRequestBuilder[Slice[PriceType]](service.client, path).Post(ctx, priceTypes)
}

func (service *contextCompanySettingsService) UpdatePriceTypeMany(ctx context.Context, priceTypes Slice[PriceType]) (*Slice[PriceType], *resty.Response, error) {
	path := "context/companysettings/pricetype"
	return NewRequestBuilder[Slice[PriceType]](service.client, path).Post(ctx, priceTypes)
}

func (service *contextCompanySettingsService) GetPriceTypeByID(ctx context.Context, id uuid.UUID) (*PriceType, *resty.Response, error) {
	path := fmt.Sprintf("context/companysettings/pricetype/%s", id)
	return NewRequestBuilder[PriceType](service.client, path).Get(ctx)
}

func (service *contextCompanySettingsService) GetPriceTypeDefault(ctx context.Context) (*PriceType, *resty.Response, error) {
	path := "context/companysettings/pricetype/default"
	return NewRequestBuilder[PriceType](service.client, path).Get(ctx)
}
