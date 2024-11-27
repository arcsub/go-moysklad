package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
)

// Group Отдел.
//
// Код сущности: group
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-otdel
type Group struct {
	AccountID *string `json:"accountId,omitempty"` // ID учётной записи
	ID        *string `json:"id,omitempty"`        // ID Отдела
	Index     *int    `json:"index,omitempty"`     // Порядковый номер в списке отделов
	Meta      *Meta   `json:"meta,omitempty"`      // Метаданные Отдела
	Name      *string `json:"name,omitempty"`      // Наименование Отдела
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (group Group) Clean() *Group {
	if group.Meta == nil {
		return nil
	}
	return &Group{Meta: group.Meta}
}

// GetAccountID возвращает ID учётной записи.
func (group Group) GetAccountID() string {
	return Deref(group.AccountID)
}

// GetID возвращает ID Отдела.
func (group Group) GetID() string {
	return Deref(group.ID)
}

// GetIndex возвращает Порядковый номер в списке отделов.
func (group Group) GetIndex() int {
	return Deref(group.Index)
}

// GetMeta возвращает Метаданные Отдела.
func (group Group) GetMeta() Meta {
	return Deref(group.Meta)
}

// GetName возвращает Наименование Отдела.
func (group Group) GetName() string {
	return Deref(group.Name)
}

// SetIndex устанавливает Порядковый номер в списке отделов.
func (group *Group) SetIndex(index int) *Group {
	group.Index = &index
	return group
}

// SetMeta устанавливает Метаданные Отдела.
func (group *Group) SetMeta(meta *Meta) *Group {
	group.Meta = meta
	return group
}

// SetName устанавливает Наименование Отдела.
func (group *Group) SetName(name string) *Group {
	group.Name = &name
	return group
}

// String реализует интерфейс [fmt.Stringer].
func (group Group) String() string {
	return Stringify(group)
}

// MetaType возвращает код сущности.
func (Group) MetaType() MetaType {
	return MetaTypeGroup
}

// Update shortcut
func (group *Group) Update(ctx context.Context, client *Client, params ...*Params) (*Group, *resty.Response, error) {
	return NewGroupService(client).Update(ctx, group.GetID(), group, params...)
}

// Create shortcut
func (group *Group) Create(ctx context.Context, client *Client, params ...*Params) (*Group, *resty.Response, error) {
	return NewGroupService(client).Create(ctx, group, params...)
}

// Delete shortcut
func (group *Group) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewGroupService(client).Delete(ctx, group)
}

// GroupService описывает методы сервиса для работы с отделами.
type GroupService interface {
	// GetList выполняет запрос на получение списка отделов.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[Group], *resty.Response, error)

	// GetListAll выполняет запрос на получение всех отделов в виде списка.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает список объектов.
	GetListAll(ctx context.Context, params ...*Params) (*Slice[Group], *resty.Response, error)

	// Create выполняет запрос на создание отдела.
	// Обязательные поля для заполнения:
	//	- name (Наименование отдела)
	// Принимает контекст, отдел и опционально объект параметров запроса Params.
	// Возвращает созданный отдел.
	Create(ctx context.Context, group *Group, params ...*Params) (*Group, *resty.Response, error)

	// DeleteByID выполняет запрос на удаление отдела по ID.
	// Принимает контекст и ID отдела.
	// Возвращает «true» в случае успешного удаления отдела.
	DeleteByID(ctx context.Context, id string) (bool, *resty.Response, error)

	// Delete выполняет запрос на удаление отдела.
	// Принимает контекст и отдел.
	// Возвращает «true» в случае успешного удаления отдела.
	Delete(ctx context.Context, entity *Group) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельного отдела по ID.
	// Принимает контекст, ID отдела и опционально объект параметров запроса Params.
	// Возвращает найденный отдел.
	GetByID(ctx context.Context, id string, params ...*Params) (*Group, *resty.Response, error)

	// Update выполняет запрос на изменение отдела.
	// Принимает контекст, отдел и опционально объект параметров запроса Params.
	// Возвращает изменённый отдел.
	Update(ctx context.Context, id string, group *Group, params ...*Params) (*Group, *resty.Response, error)
}

const (
	EndpointGroup = EndpointEntity + string(MetaTypeGroup)
)

// NewGroupService принимает [Client] и возвращает сервис для работы с отделами.
func NewGroupService(client *Client) GroupService {
	return newMainService[Group, any, any, any](client, EndpointGroup)
}
