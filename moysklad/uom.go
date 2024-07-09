package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Uom Единица измерения.
//
// Код сущности: uom
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-edinica-izmereniq
type Uom struct {
	AccountID    *uuid.UUID `json:"accountId,omitempty"`    // ID учётной записи
	Code         *string    `json:"code,omitempty"`         // Код Единицы измерения
	Description  *string    `json:"description,omitempty"`  // Описание Единицы измерения
	ExternalCode *string    `json:"externalCode,omitempty"` // Внешний код Единицы измерения
	Group        *Group     `json:"group,omitempty"`        // Отдел сотрудника
	ID           *uuid.UUID `json:"id,omitempty"`           // ID сущности
	Meta         *Meta      `json:"meta,omitempty"`         // Метаданные
	Name         *string    `json:"name,omitempty"`         // Наименование Единицы измерения
	Owner        *Employee  `json:"owner,omitempty"`        // Метаданные владельца (Сотрудника)
	Shared       *bool      `json:"shared,omitempty"`       // Общий доступ
	Updated      *Timestamp `json:"updated,omitempty"`      // Момент последнего обновления Единицы измерения
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (uom Uom) Clean() *Uom {
	if uom.Meta == nil {
		return nil
	}
	return &Uom{Meta: uom.Meta}
}

func (uom Uom) GetAccountID() uuid.UUID {
	return Deref(uom.AccountID)
}

func (uom Uom) GetCode() string {
	return Deref(uom.Code)
}

func (uom Uom) GetDescription() string {
	return Deref(uom.Description)
}

func (uom Uom) GetExternalCode() string {
	return Deref(uom.ExternalCode)
}

func (uom Uom) GetGroup() Group {
	return Deref(uom.Group)
}

func (uom Uom) GetID() uuid.UUID {
	return Deref(uom.ID)
}

func (uom Uom) GetMeta() Meta {
	return Deref(uom.Meta)
}

func (uom Uom) GetName() string {
	return Deref(uom.Name)
}

func (uom Uom) GetOwner() Employee {
	return Deref(uom.Owner)
}

func (uom Uom) GetShared() bool {
	return Deref(uom.Shared)
}

func (uom Uom) GetUpdated() Timestamp {
	return Deref(uom.Updated)
}

func (uom *Uom) SetCode(code string) *Uom {
	uom.Code = &code
	return uom
}

func (uom *Uom) SetDescription(detDescription string) *Uom {
	uom.Description = &detDescription
	return uom
}

func (uom *Uom) SetExternalCode(externalCode string) *Uom {
	uom.ExternalCode = &externalCode
	return uom
}

func (uom *Uom) SetGroup(group *Group) *Uom {
	uom.Group = group.Clean()
	return uom
}

func (uom *Uom) SetMeta(meta *Meta) *Uom {
	uom.Meta = meta
	return uom
}

func (uom *Uom) SetName(name string) *Uom {
	uom.Name = &name
	return uom
}

func (uom *Uom) SetOwner(owner *Employee) *Uom {
	uom.Owner = owner.Clean()
	return uom
}

func (uom *Uom) SetShared(shared bool) *Uom {
	uom.Shared = &shared
	return uom
}

// String реализует интерфейс [fmt.Stringer].
func (uom Uom) String() string {
	return Stringify(uom)
}

// MetaType возвращает код сущности.
func (Uom) MetaType() MetaType {
	return MetaTypeUom
}

// Update shortcut
func (uom Uom) Update(ctx context.Context, client *Client, params ...*Params) (*Uom, *resty.Response, error) {
	return NewUomService(client).Update(ctx, uom.GetID(), &uom, params...)
}

// Create shortcut
func (uom Uom) Create(ctx context.Context, client *Client, params ...*Params) (*Uom, *resty.Response, error) {
	return NewUomService(client).Create(ctx, &uom, params...)
}

// Delete shortcut
func (uom Uom) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewUomService(client).Delete(ctx, uom.GetID())
}

// UomService описывает методы сервиса для работы с единицами измерения.
type UomService interface {
	// GetList выполняет запрос на получение списка единиц измерения.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[Uom], *resty.Response, error)

	// Create выполняет запрос на создание единицы измерения.
	// Обязательные поля для заполнения:
	//	- name (Наименование единицы измерения)
	// Принимает контекст, единицу измерения и опционально объект параметров запроса Params.
	// Возвращает созданную единицу измерения.
	Create(ctx context.Context, uom *Uom, params ...*Params) (*Uom, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение единиц измерения.
	// Изменяемые единицы измерения должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список единиц измерения и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых единиц измерения.
	CreateUpdateMany(ctx context.Context, uomList Slice[Uom], params ...*Params) (*Slice[Uom], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление единиц измерения.
	// Принимает контекст и множество единиц измерения.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*Uom) (*DeleteManyResponse, *resty.Response, error)

	// Delete выполняет запрос на удаление единицы измерения.
	// Принимает контекст и ID единицы измерения.
	// Возвращает «true» в случае успешного удаления единицы измерения.
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельной единицы измерения по ID.
	// Принимает контекст, ID единицы измерения и опционально объект параметров запроса Params.
	// Возвращает найденную единицу измерения.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Uom, *resty.Response, error)

	// Update выполняет запрос на изменение единицы измерения.
	// Принимает контекст, единицу измерения и опционально объект параметров запроса Params.
	// Возвращает изменённую единицу измерения.
	Update(ctx context.Context, id uuid.UUID, uom *Uom, params ...*Params) (*Uom, *resty.Response, error)
}

const EntityEndpoint = "entity/"
const UomEndpoint = EntityEndpoint + "uom"

// NewUomService принимает [Client] и возвращает сервис для работы с единицами измерения.
func NewUomService(client *Client) UomService {
	return newMainService[Uom, any, any, any](NewEndpoint(client, UomEndpoint))
}
