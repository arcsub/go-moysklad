package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Group Отдел.
// Ключевое слово: group
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-otdel
type Group struct {
	AccountID *uuid.UUID `json:"accountId,omitempty"` // ID учетной записи
	ID        *uuid.UUID `json:"id,omitempty"`        // ID Отдела
	Index     *int       `json:"index,omitempty"`     // Порядковый номер в списке отделов
	Meta      *Meta      `json:"meta,omitempty"`      // Метаданные Отдела
	Name      *string    `json:"name,omitempty"`      // Наименование Отдела
}

// Clean возвращает сущность с единственным заполненным полем Meta
func (group Group) Clean() *Group {
	return &Group{Meta: group.Meta}
}

func (group Group) GetAccountID() uuid.UUID {
	return Deref(group.AccountID)
}

func (group Group) GetID() uuid.UUID {
	return Deref(group.ID)
}

func (group Group) GetIndex() int {
	return Deref(group.Index)
}

func (group Group) GetMeta() Meta {
	return Deref(group.Meta)
}

func (group Group) GetName() string {
	return Deref(group.Name)
}

func (group *Group) SetIndex(index int) *Group {
	group.Index = &index
	return group
}

func (group *Group) SetMeta(meta *Meta) *Group {
	group.Meta = meta
	return group
}

func (group *Group) SetName(name string) *Group {
	group.Name = &name
	return group
}

func (group Group) String() string {
	return Stringify(group)
}

// MetaType возвращает тип сущности.
func (Group) MetaType() MetaType {
	return MetaTypeGroup
}

// Update shortcut
func (group Group) Update(ctx context.Context, client *Client, params ...*Params) (*Group, *resty.Response, error) {
	return client.Entity().Group().Update(ctx, group.GetID(), &group, params...)
}

// Create shortcut
func (group Group) Create(ctx context.Context, client *Client, params ...*Params) (*Group, *resty.Response, error) {
	return client.Entity().Group().Create(ctx, &group, params...)
}

// Delete shortcut
func (group Group) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return client.Entity().Group().Delete(ctx, group.GetID())
}

// GroupService
// Сервис для работы с отделами.
type GroupService interface {
	GetList(ctx context.Context, params ...*Params) (*List[Group], *resty.Response, error)
	Create(ctx context.Context, group *Group, params ...*Params) (*Group, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Group, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, group *Group, params ...*Params) (*Group, *resty.Response, error)
}

func NewGroupService(client *Client) GroupService {
	e := NewEndpoint(client, "entity/group")
	return newMainService[Group, any, any, any](e)
}
