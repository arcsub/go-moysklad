package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Country Страна.
// Ключевое слово: country
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-strana
type Country struct {
	AccountID    *uuid.UUID `json:"accountId,omitempty"`    // ID учетной записи
	Code         *string    `json:"code,omitempty"`         // Код Страны
	Description  *string    `json:"description,omitempty"`  // Описание Страны
	ExternalCode *string    `json:"externalCode,omitempty"` // Внешний код Страны
	Group        *Group     `json:"group,omitempty"`        // Отдел-владелец
	ID           *uuid.UUID `json:"id,omitempty"`           // ID сущности
	Meta         *Meta      `json:"meta,omitempty"`         // Метаданные
	Name         *string    `json:"name,omitempty"`         // Наименование
	Owner        *Employee  `json:"owner,omitempty"`        // Сотрудник-владелец
	Shared       *bool      `json:"shared,omitempty"`       // Флаг Общий доступ
	Updated      *Timestamp `json:"updated,omitempty"`      // Момент последнего обновления сущности
}

func (c Country) String() string {
	return Stringify(c)
}

func (c Country) MetaType() MetaType {
	return MetaTypeCountry
}

// CountryService
// Сервис для работы со странами.
type CountryService interface {
	GetList(ctx context.Context, params *Params) (*List[Country], *resty.Response, error)
	Create(ctx context.Context, country *Country, params *Params) (*Country, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, countryList []*Country, params *Params) (*[]Country, *resty.Response, error)
	DeleteMany(ctx context.Context, countryList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Country, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, country *Country, params *Params) (*Country, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
}

func NewCountryService(client *Client) CountryService {
	e := NewEndpoint(client, "entity/country")
	return newMainService[Country, any, any, any](e)
}
