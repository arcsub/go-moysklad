package moysklad

import (
	"context"
	"fmt"
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
func (s *ContextCompanySettingsService) GetPriceTypes(ctx context.Context) (*Slice[PriceType], *Response, error) {
	path := "pricetype"
	return NewRequestBuilder[Slice[PriceType]](s.Endpoint, ctx).WithPath(path).Get()
}

// CreatePriceType Создать тип цен.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tipy-cen-redaktirowanie-spiska-tipow-cen
func (s *ContextCompanySettingsService) CreatePriceType(ctx context.Context, priceType *PriceType) (*Slice[PriceType], *Response, error) {
	path := "pricetype"
	return NewRequestBuilder[Slice[PriceType]](s.Endpoint, ctx).WithPath(path).WithBody(priceType).Post()
}

// UpdatePriceTypes Редактирование списка типов цен.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tipy-cen-redaktirowanie-spiska-tipow-cen
func (s *ContextCompanySettingsService) UpdatePriceTypes(ctx context.Context, priceTypes []*PriceType) (*Slice[PriceType], *Response, error) {
	path := "pricetype"
	return NewRequestBuilder[Slice[PriceType]](s.Endpoint, ctx).WithPath(path).WithBody(priceTypes).Post()
}

// GetPriceTypeById Получить тип цены по ID.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tipy-cen-poluchit-tip-ceny-po-id
func (s *ContextCompanySettingsService) GetPriceTypeById(ctx context.Context, id *uuid.UUID) (*PriceType, *Response, error) {
	path := fmt.Sprintf("pricetype/%s", id)
	return NewRequestBuilder[PriceType](s.Endpoint, ctx).WithPath(path).Get()
}

// GetPriceTypeDefault Получить тип цены по умолчанию.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tipy-cen-poluchit-tip-ceny-po-umolchaniu
func (s *ContextCompanySettingsService) GetPriceTypeDefault(ctx context.Context) (*PriceType, *Response, error) {
	path := "pricetype/default"
	return NewRequestBuilder[PriceType](s.Endpoint, ctx).WithPath(path).Get()
}
