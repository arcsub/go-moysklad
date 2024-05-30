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
	PriceTypes               *PriceTypes      `json:"priceTypes,omitempty"`
	UseCompanyAddress        *bool            `json:"useCompanyAddress,omitempty"`
	UseRecycleBin            *bool            `json:"useRecycleBin,omitempty"`
	DiscountStrategy         DiscountStrategy `json:"discountStrategy,omitempty"`
}

func (c CompanySettings) String() string {
	return Stringify(c)
}

func (c CompanySettings) MetaType() MetaType {
	return MetaTypeCompanySettings
}

// DiscountStrategy Совместное применение скидок
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-nastrojki-kompanii-sowmestnoe-primenenie-skidok
type DiscountStrategy string

const (
	DiscountStrategyBySum      DiscountStrategy = "bySum"      // Сумма скидок (должна действовать сумма скидок)
	DiscountStrategyByPriority DiscountStrategy = "byPriority" // Приоритетная (должна действовать одна, наиболее выгодная для покупателя скидка)
)

// ContextCompanySettingsService
// Сервис для работы с настройками компании.
type ContextCompanySettingsService interface {
	Get(ctx context.Context, params *Params) (*CompanySettings, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, settings *CompanySettings, params *Params) (*CompanySettings, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetadataCompanySettings, *resty.Response, error)
	GetPriceTypes(ctx context.Context) (*[]PriceType, *resty.Response, error)
	CreatePriceType(ctx context.Context, priceType *PriceType) (*[]PriceType, *resty.Response, error)
	UpdatePriceTypes(ctx context.Context, priceTypeList []*PriceType) (*[]PriceType, *resty.Response, error)
	GetPriceTypeById(ctx context.Context, id *uuid.UUID) (*PriceType, *resty.Response, error)
	GetPriceTypeDefault(ctx context.Context) (*PriceType, *resty.Response, error)
}

type contextCompanySettingsService struct {
	Endpoint
	endpointGetOne[CompanySettings]
	endpointUpdate[CompanySettings]
	endpointMetadata[MetadataCompanySettings]
}

func NewContextCompanySettingsService(client *Client) ContextCompanySettingsService {
	e := NewEndpoint(client, "context/companysettings")
	return &contextCompanySettingsService{
		Endpoint:         e,
		endpointGetOne:   endpointGetOne[CompanySettings]{e},
		endpointUpdate:   endpointUpdate[CompanySettings]{e},
		endpointMetadata: endpointMetadata[MetadataCompanySettings]{e},
	}
}

// GetPriceTypes Получить список всех типов цен.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tipy-cen-poluchit-spisok-wseh-tipow-cen
func (s *contextCompanySettingsService) GetPriceTypes(ctx context.Context) (*[]PriceType, *resty.Response, error) {
	path := "context/companysettings/pricetype"
	return NewRequestBuilder[[]PriceType](s.client, path).Get(ctx)
}

// CreatePriceType Создать тип цен.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tipy-cen-redaktirowanie-spiska-tipow-cen
func (s *contextCompanySettingsService) CreatePriceType(ctx context.Context, priceType *PriceType) (*[]PriceType, *resty.Response, error) {
	path := "context/companysettings/pricetype"
	return NewRequestBuilder[[]PriceType](s.client, path).Post(ctx, priceType)
}

// UpdatePriceTypes Редактирование списка типов цен.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tipy-cen-redaktirowanie-spiska-tipow-cen
func (s *contextCompanySettingsService) UpdatePriceTypes(ctx context.Context, priceTypeList []*PriceType) (*[]PriceType, *resty.Response, error) {
	path := "context/companysettings/pricetype"
	return NewRequestBuilder[[]PriceType](s.client, path).Post(ctx, priceTypeList)
}

// GetPriceTypeById Получить тип цены по ID.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tipy-cen-poluchit-tip-ceny-po-id
func (s *contextCompanySettingsService) GetPriceTypeById(ctx context.Context, id *uuid.UUID) (*PriceType, *resty.Response, error) {
	path := fmt.Sprintf("context/companysettings/pricetype/%s", id)
	return NewRequestBuilder[PriceType](s.client, path).Get(ctx)
}

// GetPriceTypeDefault Получить тип цены по умолчанию.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tipy-cen-poluchit-tip-ceny-po-umolchaniu
func (s *contextCompanySettingsService) GetPriceTypeDefault(ctx context.Context) (*PriceType, *resty.Response, error) {
	path := "context/companysettings/pricetype/default"
	return NewRequestBuilder[PriceType](s.client, path).Get(ctx)
}
