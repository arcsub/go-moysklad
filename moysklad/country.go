package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"time"
)

// Country Страна.
//
// Код сущности: country
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-strana
type Country struct {
	AccountID    *uuid.UUID `json:"accountId,omitempty"`    // ID учётной записи
	Code         *string    `json:"code,omitempty"`         // Код Страны
	Description  *string    `json:"description,omitempty"`  // Описание Страны
	ExternalCode *string    `json:"externalCode,omitempty"` // Внешний код Страны
	Group        *Group     `json:"group,omitempty"`        // Отдел сотрудника
	ID           *uuid.UUID `json:"id,omitempty"`           // ID Страны
	Meta         *Meta      `json:"meta,omitempty"`         // Метаданные Страны
	Name         *string    `json:"name,omitempty"`         // Наименование Страны
	Owner        *Employee  `json:"owner,omitempty"`        // Метаданные владельца (Сотрудника)
	Shared       *bool      `json:"shared,omitempty"`       // Общий доступ
	Updated      *Timestamp `json:"updated,omitempty"`      // Момент последнего обновления Страны
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (country Country) Clean() *Country {
	if country.Meta == nil {
		return nil
	}
	return &Country{Meta: country.Meta}
}

// GetAccountID возвращает ID учётной записи.
func (country Country) GetAccountID() uuid.UUID {
	return Deref(country.AccountID)
}

// GetCode возвращает Код Страны.
func (country Country) GetCode() string {
	return Deref(country.Code)
}

// GetDescription возвращает Описание Страны.
func (country Country) GetDescription() string {
	return Deref(country.Description)
}

// GetExternalCode возвращает Внешний код Страны.
func (country Country) GetExternalCode() string {
	return Deref(country.ExternalCode)
}

// GetGroup возвращает Отдел сотрудника.
func (country Country) GetGroup() Group {
	return Deref(country.Group)
}

// GetID возвращает ID Страны.
func (country Country) GetID() uuid.UUID {
	return Deref(country.ID)
}

// GetMeta возвращает Метаданные Страны.
func (country Country) GetMeta() Meta {
	return Deref(country.Meta)
}

// GetName возвращает Наименование Страны.
func (country Country) GetName() string {
	return Deref(country.Name)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (country Country) GetOwner() Employee {
	return Deref(country.Owner)
}

// GetShared возвращает флаг Общего доступа.
func (country Country) GetShared() bool {
	return Deref(country.Shared)
}

// GetUpdated возвращает Момент последнего обновления Страны.
func (country Country) GetUpdated() time.Time {
	return Deref(country.Updated).Time()
}

// SetCode устанавливает Код Страны.
func (country *Country) SetCode(code string) *Country {
	country.Code = &code
	return country
}

// SetDescription устанавливает Описание Страны.
func (country *Country) SetDescription(description string) *Country {
	country.Description = &description
	return country
}

// SetExternalCode устанавливает Внешний код Страны.
func (country *Country) SetExternalCode(externalCode string) *Country {
	country.ExternalCode = &externalCode
	return country
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (country *Country) SetGroup(group *Group) *Country {
	if group != nil {
		country.Group = group.Clean()
	}
	return country
}

// SetMeta устанавливает Метаданные Страны.
func (country *Country) SetMeta(meta *Meta) *Country {
	country.Meta = meta
	return country
}

// SetName устанавливает Наименование Страны.
func (country *Country) SetName(name string) *Country {
	country.Name = &name
	return country
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (country *Country) SetOwner(owner *Employee) *Country {
	if owner != nil {
		country.Owner = owner.Clean()
	}
	return country
}

// SetShared устанавливает флаг общего доступа.
func (country *Country) SetShared(shared bool) *Country {
	country.Shared = &shared
	return country
}

// String реализует интерфейс [fmt.Stringer].
func (country Country) String() string {
	return Stringify(country)
}

// MetaType возвращает код сущности.
func (Country) MetaType() MetaType {
	return MetaTypeCountry
}

// Update shortcut
func (country *Country) Update(ctx context.Context, client *Client, params ...*Params) (*Country, *resty.Response, error) {
	return NewCountryService(client).Update(ctx, country.GetID(), country, params...)
}

// Create shortcut
func (country *Country) Create(ctx context.Context, client *Client, params ...*Params) (*Country, *resty.Response, error) {
	return NewCountryService(client).Create(ctx, country, params...)
}

// Delete shortcut
func (country *Country) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewCountryService(client).Delete(ctx, country)
}

// CountryService описывает методы сервиса для работы со странами.
type CountryService interface {
	// GetList выполняет запрос на получение списка стран.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[Country], *resty.Response, error)

	// GetListAll выполняет запрос на получение всех стран в виде списка.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает список объектов.
	GetListAll(ctx context.Context, params ...*Params) (*Slice[Country], *resty.Response, error)

	// Create выполняет запрос на создание страны.
	// Обязательные поля для заполнения:
	//	- name (Наименование страны)
	// Принимает контекст, страну и опционально объект параметров запроса Params.
	// Возвращает созданную страну.
	Create(ctx context.Context, country *Country, params ...*Params) (*Country, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение стран.
	// Изменяемые страны должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список стран и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых стран.
	CreateUpdateMany(ctx context.Context, countryList Slice[Country], params ...*Params) (*Slice[Country], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление стран.
	// Принимает контекст и множество стран.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*Country) (*DeleteManyResponse, *resty.Response, error)

	// DeleteByID выполняет запрос на удаление страны по ID.
	// Принимает контекст и ID страны.
	// Возвращает «true» в случае успешного удаления страны.
	DeleteByID(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// Delete выполняет запрос на удаление страны.
	// Принимает контекст и страну.
	// Возвращает «true» в случае успешного удаления страны.
	Delete(ctx context.Context, entity *Country) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельной страны по ID.
	// Принимает контекст, ID страны и опционально объект параметров запроса Params.
	// Возвращает найденную страну.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Country, *resty.Response, error)

	// Update выполняет запрос на изменение страны.
	// Принимает контекст, страну и опционально объект параметров запроса Params.
	// Возвращает изменённую страну.
	Update(ctx context.Context, id uuid.UUID, country *Country, params ...*Params) (*Country, *resty.Response, error)

	// GetNamedFilterList выполняет запрос на получение списка фильтров.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetNamedFilterList(ctx context.Context, params ...*Params) (*List[NamedFilter], *resty.Response, error)

	// GetNamedFilterByID выполняет запрос на получение отдельного фильтра по ID.
	// Принимает контекст и ID фильтра.
	// Возвращает найденный фильтр.
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)
}

const (
	EndpointCountry = EndpointEntity + string(MetaTypeCountry)
)

// NewCountryService принимает [Client] и возвращает сервис для работы со странами.
func NewCountryService(client *Client) CountryService {
	return newMainService[Country, any, any, any](client, EndpointCountry)
}
