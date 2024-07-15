package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"time"
)

// CustomEntity Пользовательский справочник.
//
// Код сущности: customentity
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-pol-zowatel-skij-sprawochnik
type CustomEntity struct {
	ID   *uuid.UUID `json:"id,omitempty"`   // ID Пользовательского справочника
	Meta *Meta      `json:"meta,omitempty"` // Метаданные Пользовательского справочника
	Name *string    `json:"name,omitempty"` // Наименование Пользовательского справочника
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (customEntity CustomEntity) Clean() *CustomEntity {
	if customEntity.Meta == nil {
		return nil
	}
	return &CustomEntity{Meta: customEntity.Meta}
}

// GetID возвращает ID Пользовательского справочника.
func (customEntity CustomEntity) GetID() uuid.UUID {
	return Deref(customEntity.ID)
}

// GetMeta возвращает Метаданные Пользовательского справочника.
func (customEntity CustomEntity) GetMeta() Meta {
	return Deref(customEntity.Meta)
}

// GetName возвращает Наименование Пользовательского справочника.
func (customEntity CustomEntity) GetName() string {
	return Deref(customEntity.Name)
}

// SetMeta устанавливает Метаданные Пользовательского справочника.
func (customEntity *CustomEntity) SetMeta(meta *Meta) *CustomEntity {
	customEntity.Meta = meta
	return customEntity
}

// SetName устанавливает Наименование Пользовательского справочника.
func (customEntity *CustomEntity) SetName(name string) *CustomEntity {
	customEntity.Name = &name
	return customEntity
}

// String реализует интерфейс [fmt.Stringer].
func (customEntity CustomEntity) String() string {
	return Stringify(customEntity)
}

// MetaType возвращает код сущности.
func (CustomEntity) MetaType() MetaType {
	return MetaTypeCustomEntity
}

// Update shortcut
func (customEntity *CustomEntity) Update(ctx context.Context, client *Client, params ...*Params) (*CustomEntity, *resty.Response, error) {
	return NewCustomEntityService(client).Update(ctx, customEntity.GetID(), customEntity, params...)
}

// Create shortcut
func (customEntity *CustomEntity) Create(ctx context.Context, client *Client, params ...*Params) (*CustomEntity, *resty.Response, error) {
	return NewCustomEntityService(client).Create(ctx, customEntity, params...)
}

// Delete shortcut
func (customEntity *CustomEntity) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewCustomEntityService(client).Delete(ctx, customEntity)
}

// CustomEntityElement Элемент Пользовательского справочника.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-pol-zowatel-skij-sprawochnik-jelementy-pol-zowatel-skogo-sprawochnika
type CustomEntityElement struct {
	AccountID    *uuid.UUID `json:"accountId,omitempty"`    // ID учётной записи
	Code         *string    `json:"code,omitempty"`         // Код элемента Пользовательского справочника
	Description  *string    `json:"description,omitempty"`  // Описание элемента Пользовательского справочника
	ExternalCode *string    `json:"externalCode,omitempty"` // Внешний код элемента Пользовательского справочника
	ID           *uuid.UUID `json:"id,omitempty"`           // ID элемента Пользовательского справочника
	Meta         *Meta      `json:"meta,omitempty"`         // Метаданные элемента Пользовательского справочника
	Name         *string    `json:"name,omitempty"`         // Наименование элемента Пользовательского справочника
	Updated      *Timestamp `json:"updated,omitempty"`      // Момент последнего обновления элементе Пользовательского справочника
	Group        *Group     `json:"group,omitempty"`        // Отдел сотрудника
	Owner        *Employee  `json:"owner,omitempty"`        // Метаданные владельца (Сотрудника)
	Shared       *bool      `json:"shared,omitempty"`       // Общий доступ
}

// GetAccountID возвращает ID учётной записи.
func (customEntityElement CustomEntityElement) GetAccountID() uuid.UUID {
	return Deref(customEntityElement.AccountID)
}

// GetCode возвращает Код элемента Пользовательского справочника.
func (customEntityElement CustomEntityElement) GetCode() string {
	return Deref(customEntityElement.Code)
}

// GetDescription возвращает Описание элемента Пользовательского справочника.
func (customEntityElement CustomEntityElement) GetDescription() string {
	return Deref(customEntityElement.Description)
}

// GetExternalCode возвращает Внешний код элемента Пользовательского справочника.
func (customEntityElement CustomEntityElement) GetExternalCode() string {
	return Deref(customEntityElement.ExternalCode)
}

// GetID возвращает ID элемента Пользовательского справочника.
func (customEntityElement CustomEntityElement) GetID() uuid.UUID {
	return Deref(customEntityElement.ID)
}

// GetMeta возвращает Метаданные элемента Пользовательского справочника.
func (customEntityElement CustomEntityElement) GetMeta() Meta {
	return Deref(customEntityElement.Meta)
}

// GetName возвращает Наименование элемента Пользовательского справочника.
func (customEntityElement CustomEntityElement) GetName() string {
	return Deref(customEntityElement.Name)
}

// GetUpdated возвращает Момент последнего обновления элементе Пользовательского справочника.
func (customEntityElement CustomEntityElement) GetUpdated() time.Time {
	return Deref(customEntityElement.Updated).Time()
}

// GetGroup возвращает Отдел сотрудника.
func (customEntityElement CustomEntityElement) GetGroup() Group {
	return Deref(customEntityElement.Group)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (customEntityElement CustomEntityElement) GetOwner() Employee {
	return Deref(customEntityElement.Owner)
}

// GetShared возвращает флаг Общего доступа.
func (customEntityElement CustomEntityElement) GetShared() bool {
	return Deref(customEntityElement.Shared)
}

// SetCode устанавливает Код элемента Пользовательского справочника.
func (customEntityElement *CustomEntityElement) SetCode(code string) *CustomEntityElement {
	customEntityElement.Code = &code
	return customEntityElement
}

// SetDescription устанавливает Описание элемента Пользовательского справочника.
func (customEntityElement *CustomEntityElement) SetDescription(description string) *CustomEntityElement {
	customEntityElement.Description = &description
	return customEntityElement
}

// SetExternalCode устанавливает Внешний код элемента Пользовательского справочника.
func (customEntityElement *CustomEntityElement) SetExternalCode(externalCode string) *CustomEntityElement {
	customEntityElement.ExternalCode = &externalCode
	return customEntityElement
}

// SetMeta устанавливает Метаданные элемента Пользовательского справочника.
func (customEntityElement *CustomEntityElement) SetMeta(meta *Meta) *CustomEntityElement {
	customEntityElement.Meta = meta
	return customEntityElement
}

// SetName устанавливает Наименование элемента Пользовательского справочника.
func (customEntityElement *CustomEntityElement) SetName(name string) *CustomEntityElement {
	customEntityElement.Name = &name
	return customEntityElement
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (customEntityElement *CustomEntityElement) SetGroup(group *Group) *CustomEntityElement {
	if group != nil {
		customEntityElement.Group = group.Clean()
	}
	return customEntityElement
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (customEntityElement *CustomEntityElement) SetOwner(owner *Employee) *CustomEntityElement {
	if owner != nil {
		customEntityElement.Owner = owner.Clean()
	}
	return customEntityElement
}

// SetShared устанавливает флаг общего доступа.
func (customEntityElement *CustomEntityElement) SetShared(shared bool) *CustomEntityElement {
	customEntityElement.Shared = &shared
	return customEntityElement
}

// String реализует интерфейс [fmt.Stringer].
func (customEntityElement CustomEntityElement) String() string {
	return Stringify(customEntityElement)
}

// CustomEntityService описывает методы сервиса для работы с пользовательскими справочниками.
type CustomEntityService interface {
	// Create выполняет запрос на создание пользовательского справочника.
	// Обязательные поля для заполнения:
	//	- name (Наименование Пользовательского справочника)
	// Принимает контекст, пользовательский справочник и опционально объект параметров запроса Params.
	// Возвращает созданный пользовательский справочник.
	Create(ctx context.Context, customEntity *CustomEntity, params ...*Params) (*CustomEntity, *resty.Response, error)

	// Update выполняет запрос на изменение пользовательского справочника.
	// Принимает контекст, пользовательский справочник и опционально объект параметров запроса Params.
	// Возвращает изменённый пользовательский справочник.
	Update(ctx context.Context, id uuid.UUID, customEntity *CustomEntity, params ...*Params) (*CustomEntity, *resty.Response, error)

	// DeleteByID выполняет запрос на удаление пользовательского справочника по ID.
	// Принимает контекст и ID пользовательского справочника.
	// Возвращает «true» в случае успешного удаления пользовательского справочника.
	DeleteByID(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// Delete выполняет запрос на удаление пользовательского справочника.
	// Принимает контекст и пользовательский справочник.
	// Возвращает «true» в случае успешного удаления пользовательского справочника.
	Delete(ctx context.Context, entity *CustomEntity) (bool, *resty.Response, error)

	// GetElementList выполняет запрос на получение списка элементов пользовательского справочника.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetElementList(ctx context.Context, id uuid.UUID) (*List[CustomEntityElement], *resty.Response, error)

	// CreateElement выполняет запрос на создание элемента пользовательского справочника.
	// Обязательные поля для заполнения:
	//	- name (Наименование элемента пользовательского справочника)
	// Принимает контекст, элемент пользовательского справочника и опционально объект параметров запроса Params.
	// Возвращает созданный элемент пользовательского справочника.
	CreateElement(ctx context.Context, id uuid.UUID, element *CustomEntityElement) (*CustomEntityElement, *resty.Response, error)

	// DeleteElement выполняет запрос на удаление элемента пользовательского справочника.
	// Принимает контекст и ID элемента пользовательского справочника.
	// Возвращает «true» в случае успешного удаления элемента пользовательского справочника.
	DeleteElement(ctx context.Context, id, elementID uuid.UUID) (bool, *resty.Response, error)

	// GetElementByID выполняет запрос на получение отдельного элемента пользовательского справочника по ID.
	// Принимает контекст, ID элемента пользовательского справочника и опционально объект параметров запроса Params.
	// Возвращает найденный элемент пользовательского справочника.
	GetElementByID(ctx context.Context, id, elementID uuid.UUID) (*CustomEntityElement, *resty.Response, error)

	// UpdateElement выполняет запрос на изменение элемента пользовательского справочника.
	// Принимает контекст, элемент пользовательского справочника и опционально объект параметров запроса Params.
	// Возвращает изменённый элемент пользовательского справочника.
	UpdateElement(ctx context.Context, id, elementID uuid.UUID, element *CustomEntityElement) (*CustomEntityElement, *resty.Response, error)
}

const (
	EndpointCustomEntity     = EndpointEntity + string(MetaTypeCustomEntity)
	EndpointCustomEntityID   = EndpointCustomEntity + "/%s"
	EndpointCustomEntityIDID = EndpointCustomEntityID + "/%s"
)

type customEntityService struct {
	Endpoint
	endpointCreate[CustomEntity]
	endpointUpdate[CustomEntity]
	endpointDeleteByID
	endpointDelete[CustomEntity]
}

func (service *customEntityService) GetElementList(ctx context.Context, id uuid.UUID) (*List[CustomEntityElement], *resty.Response, error) {
	path := fmt.Sprintf(EndpointCustomEntityID, id)
	return NewRequestBuilder[List[CustomEntityElement]](service.client, path).Get(ctx)
}

func (service *customEntityService) CreateElement(ctx context.Context, id uuid.UUID, element *CustomEntityElement) (*CustomEntityElement, *resty.Response, error) {
	path := fmt.Sprintf(EndpointCustomEntityID, id)
	return NewRequestBuilder[CustomEntityElement](service.client, path).Post(ctx, element)
}

func (service *customEntityService) DeleteElement(ctx context.Context, id, elementID uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf(EndpointCustomEntityIDID, id, elementID)
	return NewRequestBuilder[any](service.client, path).Delete(ctx)
}

func (service *customEntityService) GetElementByID(ctx context.Context, id, elementID uuid.UUID) (*CustomEntityElement, *resty.Response, error) {
	path := fmt.Sprintf(EndpointCustomEntityIDID, id, elementID)
	return NewRequestBuilder[CustomEntityElement](service.client, path).Get(ctx)
}

func (service *customEntityService) UpdateElement(ctx context.Context, id, elementID uuid.UUID, element *CustomEntityElement) (*CustomEntityElement, *resty.Response, error) {
	path := fmt.Sprintf(EndpointCustomEntityIDID, id, elementID)
	return NewRequestBuilder[CustomEntityElement](service.client, path).Put(ctx, element)
}

// NewCustomEntityService принимает [Client] и возвращает сервис для работы с пользовательскими справочниками.
func NewCustomEntityService(client *Client) CustomEntityService {
	e := NewEndpoint(client, EndpointCustomEntity)
	return &customEntityService{
		Endpoint:           e,
		endpointCreate:     endpointCreate[CustomEntity]{e},
		endpointUpdate:     endpointUpdate[CustomEntity]{e},
		endpointDeleteByID: endpointDeleteByID{e},
		endpointDelete:     endpointDelete[CustomEntity]{e},
	}
}
