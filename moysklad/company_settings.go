package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// CompanySettings Настройки компании.
// Ключевое слово: companysettings
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-nastrojki-kompanii
type CompanySettings struct {
	Meta                     *Meta            `json:"meta,omitempty"`
	CheckMinPrice            *bool            `json:"checkMinPrice,omitempty"`
	CheckShippingStock       *bool            `json:"checkShippingStock,omitempty"`
	CompanyAddress           *string          `json:"companyAddress,omitempty"`
	Currency                 *Currency        `json:"currency,omitempty"`
	GlobalOperationNumbering *bool            `json:"globalOperationNumbering,omitempty"`
	UseCompanyAddress        *bool            `json:"useCompanyAddress,omitempty"`
	UseRecycleBin            *bool            `json:"useRecycleBin,omitempty"`
	DiscountStrategy         DiscountStrategy `json:"discountStrategy,omitempty"`
	AccountCountry           AccountCountry   `json:"accountCountry,omitempty"`
	PriceTypes               Slice[PriceType] `json:"priceTypes,omitempty"`
}

func (companySettings CompanySettings) GetMeta() Meta {
	return Deref(companySettings.Meta)
}

func (companySettings CompanySettings) GetCheckMinPrice() bool {
	return Deref(companySettings.CheckShippingStock)
}

func (companySettings CompanySettings) GetCheckShippingStock() bool {
	return Deref(companySettings.CheckShippingStock)
}

func (companySettings CompanySettings) GetCompanyAddress() string {
	return Deref(companySettings.CompanyAddress)
}

func (companySettings CompanySettings) GetCurrency() Currency {
	return Deref(companySettings.Currency)
}

func (companySettings CompanySettings) GetGlobalOperationNumbering() bool {
	return Deref(companySettings.GlobalOperationNumbering)
}

func (companySettings CompanySettings) GetPriceTypes() Slice[PriceType] {
	return companySettings.PriceTypes
}

func (companySettings CompanySettings) GetUseCompanyAddress() bool {
	return Deref(companySettings.UseCompanyAddress)
}

func (companySettings CompanySettings) GetUseRecycleBin() bool {
	return Deref(companySettings.UseRecycleBin)
}

func (companySettings CompanySettings) GetDiscountStrategy() DiscountStrategy {
	return companySettings.DiscountStrategy
}

func (companySettings CompanySettings) GetAccountCountry() AccountCountry {
	return companySettings.AccountCountry
}

func (companySettings *CompanySettings) SetMeta(meta *Meta) *CompanySettings {
	companySettings.Meta = meta
	return companySettings
}

func (companySettings *CompanySettings) SetCheckMinPrice(checkMinPrice bool) *CompanySettings {
	companySettings.CheckMinPrice = &checkMinPrice
	return companySettings
}

func (companySettings *CompanySettings) SetCheckShippingStock(checkShippingStock bool) *CompanySettings {
	companySettings.CheckShippingStock = &checkShippingStock
	return companySettings
}

func (companySettings *CompanySettings) SetCompanyAddress(companyAddress string) *CompanySettings {
	companySettings.CompanyAddress = &companyAddress
	return companySettings
}

func (companySettings *CompanySettings) SetCurrency(currency *Currency) *CompanySettings {
	companySettings.Currency = currency.Clean()
	return companySettings
}

func (companySettings *CompanySettings) SetGlobalOperationNumbering(globalOperationNumbering bool) *CompanySettings {
	companySettings.GlobalOperationNumbering = &globalOperationNumbering
	return companySettings
}

func (companySettings *CompanySettings) SetPriceTypes(priceTypes ...*PriceType) *CompanySettings {
	companySettings.PriceTypes = priceTypes
	return companySettings
}

func (companySettings *CompanySettings) SetUseCompanyAddress(useCompanyAddress bool) *CompanySettings {
	companySettings.UseCompanyAddress = &useCompanyAddress
	return companySettings
}

func (companySettings *CompanySettings) SetUseRecycleBin(useRecycleBin bool) *CompanySettings {
	companySettings.UseRecycleBin = &useRecycleBin
	return companySettings
}

func (companySettings *CompanySettings) SetDiscountStrategy(discountStrategy DiscountStrategy) *CompanySettings {
	companySettings.DiscountStrategy = discountStrategy
	return companySettings
}

func (companySettings CompanySettings) String() string {
	return Stringify(companySettings)
}

// MetaType возвращает тип сущности.
func (CompanySettings) MetaType() MetaType {
	return MetaTypeCompanySettings
}

// DiscountStrategy Совместное применение скидок
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-nastrojki-kompanii-sowmestnoe-primenenie-skidok
type DiscountStrategy string

const (
	DiscountStrategyBySum      DiscountStrategy = "bySum"      // Сумма скидок (должна действовать сумма скидок)
	DiscountStrategyByPriority DiscountStrategy = "byPriority" // Приоритетная (должна действовать одна, наиболее выгодная для покупателя скидка)
)

type AccountCountry string

const (
	AccountCountryRU AccountCountry = "RU"
	AccountCountryBY AccountCountry = "BY"
	AccountCountryKZ AccountCountry = "KZ"
)

// ContextCompanySettingsService
// Сервис для работы с настройками компании.
type ContextCompanySettingsService interface {
	Get(ctx context.Context) (*CompanySettings, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, settings *CompanySettings, params ...*Params) (*CompanySettings, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetadataCompanySettings, *resty.Response, error)
	GetPriceTypes(ctx context.Context) (*Slice[PriceType], *resty.Response, error)
	CreatePriceType(ctx context.Context, priceType *PriceType) (*Slice[PriceType], *resty.Response, error)
	UpdatePriceTypeMany(ctx context.Context, priceTypes ...*PriceType) (*Slice[PriceType], *resty.Response, error)
	GetPriceTypeByID(ctx context.Context, id uuid.UUID) (*PriceType, *resty.Response, error)
	GetPriceTypeDefault(ctx context.Context) (*PriceType, *resty.Response, error)
}

type contextCompanySettingsService struct {
	Endpoint
	endpointUpdate[CompanySettings]
	endpointMetadata[MetadataCompanySettings]
}

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

// GetPriceTypes Получить список всех типов цен.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tipy-cen-poluchit-spisok-wseh-tipow-cen
func (service *contextCompanySettingsService) GetPriceTypes(ctx context.Context) (*Slice[PriceType], *resty.Response, error) {
	path := "context/companysettings/pricetype"
	return NewRequestBuilder[Slice[PriceType]](service.client, path).Get(ctx)
}

// CreatePriceType Создать тип цен.
func (service *contextCompanySettingsService) CreatePriceType(ctx context.Context, priceType *PriceType) (*Slice[PriceType], *resty.Response, error) {
	priceTypes, resp, err := service.GetPriceTypes(ctx)
	if err != nil {
		return nil, resp, err
	}
	priceTypes.Push(priceType)
	path := "context/companysettings/pricetype"
	return NewRequestBuilder[Slice[PriceType]](service.client, path).Post(ctx, priceTypes)
}

// UpdatePriceTypeMany Редактирование списка типов цен.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tipy-cen-redaktirowanie-spiska-tipow-cen
func (service *contextCompanySettingsService) UpdatePriceTypeMany(ctx context.Context, priceTypes ...*PriceType) (*Slice[PriceType], *resty.Response, error) {
	path := "context/companysettings/pricetype"
	return NewRequestBuilder[Slice[PriceType]](service.client, path).Post(ctx, priceTypes)
}

// GetPriceTypeByID Получить тип цены по ID.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tipy-cen-poluchit-tip-ceny-po-id
func (service *contextCompanySettingsService) GetPriceTypeByID(ctx context.Context, id uuid.UUID) (*PriceType, *resty.Response, error) {
	path := fmt.Sprintf("context/companysettings/pricetype/%s", id)
	return NewRequestBuilder[PriceType](service.client, path).Get(ctx)
}

// GetPriceTypeDefault Получить тип цены по умолчанию.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tipy-cen-poluchit-tip-ceny-po-umolchaniu
func (service *contextCompanySettingsService) GetPriceTypeDefault(ctx context.Context) (*PriceType, *resty.Response, error) {
	path := "context/companysettings/pricetype/default"
	return NewRequestBuilder[PriceType](service.client, path).Get(ctx)
}
