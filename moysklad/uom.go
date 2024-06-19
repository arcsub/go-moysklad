package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Uom Единица измерения.
// Ключевое слово: uom
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-edinica-izmereniq
type Uom struct {
	AccountID    *uuid.UUID `json:"accountId,omitempty"`    // ID учетной записи
	Code         *string    `json:"code,omitempty"`         // Код Единицы измерения
	Description  *string    `json:"description,omitempty"`  // Описание Единциы измерения
	ExternalCode *string    `json:"externalCode,omitempty"` // Внешний код Единицы измерения
	Group        *Group     `json:"group,omitempty"`        // Отдел сотрудника
	ID           *uuid.UUID `json:"id,omitempty"`           // ID сущности
	Meta         *Meta      `json:"meta,omitempty"`         // Метаданные
	Name         *string    `json:"name,omitempty"`         // Наименование Единицы измерения
	Owner        *Employee  `json:"owner,omitempty"`        // Владелец (Сотрудник)
	Shared       *bool      `json:"shared,omitempty"`       // Общий доступ
	Updated      *Timestamp `json:"updated,omitempty"`      // Момент последнего обновления Единицы измерения
}

// Clean возвращает сущность с единственным заполненным полем Meta
func (uom Uom) Clean() *Uom {
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

func (uom Uom) String() string {
	return Stringify(uom)
}

// MetaType возвращает тип сущности.
func (Uom) MetaType() MetaType {
	return MetaTypeUom
}

// Update shortcut
func (uom Uom) Update(ctx context.Context, client *Client, params ...*Params) (*Uom, *resty.Response, error) {
	return client.Entity().Uom().Update(ctx, uom.GetID(), &uom, params...)
}

// Create shortcut
func (uom Uom) Create(ctx context.Context, client *Client, params ...*Params) (*Uom, *resty.Response, error) {
	return client.Entity().Uom().Create(ctx, &uom, params...)
}

// Delete shortcut
func (uom Uom) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return client.Entity().Uom().Delete(ctx, uom.GetID())
}

// UomService
// Сервис для работы с единицами измерения.
type UomService interface {
	GetList(ctx context.Context, params ...*Params) (*List[Uom], *resty.Response, error)
	Create(ctx context.Context, uom *Uom, params ...*Params) (*Uom, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, uomList Slice[Uom], params ...*Params) (*Slice[Uom], *resty.Response, error)
	DeleteMany(ctx context.Context, entities ...*Uom) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Uom, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, uom *Uom, params ...*Params) (*Uom, *resty.Response, error)
}

func NewUomService(client *Client) UomService {
	e := NewEndpoint(client, "entity/uom")
	return newMainService[Uom, any, any, any](e)
}
