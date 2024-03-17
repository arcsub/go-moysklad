package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// ContextCompanySettingsService
// Сервис для работы с настройками компании.
type ContextCompanySettingsService struct {
	Endpoint
	endpointGetOne[CompanySettings]
	endpointUpdate[CompanySettings]
	endpointMetadata[MetadataCompanySettings]
}

func NewContextCompanySettingsService(client *Client) *ContextCompanySettingsService {
	e := NewEndpoint(client, "context/companysettings")
	return &ContextCompanySettingsService{
		Endpoint:         e,
		endpointGetOne:   endpointGetOne[CompanySettings]{e},
		endpointUpdate:   endpointUpdate[CompanySettings]{e},
		endpointMetadata: endpointMetadata[MetadataCompanySettings]{e},
	}
}

// GetPriceTypes Получить список всех типов цен.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tipy-cen-poluchit-spisok-wseh-tipow-cen
func (s *ContextCompanySettingsService) GetPriceTypes(ctx context.Context) (*[]PriceType, *resty.Response, error) {
	path := "context/companysettings/pricetype"
	return NewRequestBuilder[[]PriceType](s.client, path).Get(ctx)
}

// CreatePriceType Создать тип цен.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tipy-cen-redaktirowanie-spiska-tipow-cen
func (s *ContextCompanySettingsService) CreatePriceType(ctx context.Context, priceType *PriceType) (*[]PriceType, *resty.Response, error) {
	path := "context/companysettings/pricetype"
	return NewRequestBuilder[[]PriceType](s.client, path).Post(ctx, priceType)
}

// UpdatePriceTypes Редактирование списка типов цен.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tipy-cen-redaktirowanie-spiska-tipow-cen
func (s *ContextCompanySettingsService) UpdatePriceTypes(ctx context.Context, priceTypes []*PriceType) (*[]PriceType, *resty.Response, error) {
	path := "context/companysettings/pricetype"
	return NewRequestBuilder[[]PriceType](s.client, path).Post(ctx, priceTypes)
}

// GetPriceTypeById Получить тип цены по ID.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tipy-cen-poluchit-tip-ceny-po-id
func (s *ContextCompanySettingsService) GetPriceTypeById(ctx context.Context, id *uuid.UUID) (*PriceType, *resty.Response, error) {
	path := fmt.Sprintf("context/companysettings/pricetype/%s", id)
	return NewRequestBuilder[PriceType](s.client, path).Get(ctx)
}

// GetPriceTypeDefault Получить тип цены по умолчанию.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tipy-cen-poluchit-tip-ceny-po-umolchaniu
func (s *ContextCompanySettingsService) GetPriceTypeDefault(ctx context.Context) (*PriceType, *resty.Response, error) {
	path := "context/companysettings/pricetype/default"
	return NewRequestBuilder[PriceType](s.client, path).Get(ctx)
}
