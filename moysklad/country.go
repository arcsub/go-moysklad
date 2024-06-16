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

// Clean возвращает сущность с единственным заполненным полем Meta
func (country Country) Clean() *Country {
	return &Country{Meta: country.Meta}
}

func (country Country) GetAccountID() uuid.UUID {
	return Deref(country.AccountID)
}

func (country Country) GetCode() string {
	return Deref(country.Code)
}

func (country Country) GetDescription() string {
	return Deref(country.Description)
}

func (country Country) GetExternalCode() string {
	return Deref(country.ExternalCode)
}

func (country Country) GetGroup() Group {
	return Deref(country.Group)
}

func (country Country) GetID() uuid.UUID {
	return Deref(country.ID)
}

func (country Country) GetMeta() Meta {
	return Deref(country.Meta)
}

func (country Country) GetName() string {
	return Deref(country.Name)
}

func (country Country) GetOwner() Employee {
	return Deref(country.Owner)
}

func (country Country) GetShared() bool {
	return Deref(country.Shared)
}

func (country Country) GetUpdated() Timestamp {
	return Deref(country.Updated)
}

func (country *Country) SetCode(code string) *Country {
	country.Code = &code
	return country
}

func (country *Country) SetDescription(description string) *Country {
	country.Description = &description
	return country
}

func (country *Country) SetExternalCode(externalCode string) *Country {
	country.ExternalCode = &externalCode
	return country
}

func (country *Country) SetGroup(group *Group) *Country {
	country.Group = group.Clean()
	return country
}

func (country *Country) SetMeta(meta *Meta) *Country {
	country.Meta = meta
	return country
}

func (country *Country) SetName(name string) *Country {
	country.Name = &name
	return country
}

func (country *Country) SetOwner(owner *Employee) *Country {
	country.Owner = owner.Clean()
	return country
}

func (country *Country) SetShared(shared bool) *Country {
	country.Shared = &shared
	return country
}

func (country Country) String() string {
	return Stringify(country)
}

func (country Country) MetaType() MetaType {
	return MetaTypeCountry
}

// Update shortcut
func (country Country) Update(ctx context.Context, client *Client, params ...*Params) (*Country, *resty.Response, error) {
	return client.Entity().Country().Update(ctx, country.GetID(), &country, params...)
}

// Create shortcut
func (country Country) Create(ctx context.Context, client *Client, params ...*Params) (*Country, *resty.Response, error) {
	return client.Entity().Country().Create(ctx, &country, params...)
}

// Delete shortcut
func (country Country) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return client.Entity().Country().Delete(ctx, country.GetID())
}

// CountryService
// Сервис для работы со странами.
type CountryService interface {
	GetList(ctx context.Context, params ...*Params) (*List[Country], *resty.Response, error)
	Create(ctx context.Context, country *Country, params ...*Params) (*Country, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, countryList Slice[Country], params ...*Params) (*Slice[Country], *resty.Response, error)
	DeleteMany(ctx context.Context, countryList []MetaWrapper) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Country, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, country *Country, params ...*Params) (*Country, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params ...*Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)
}

func NewCountryService(client *Client) CountryService {
	e := NewEndpoint(client, "entity/country")
	return newMainService[Country, any, any, any](e)
}
