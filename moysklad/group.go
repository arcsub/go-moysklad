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

func (g Group) String() string {
	return Stringify(g)
}

func (g Group) MetaType() MetaType {
	return MetaTypeGroup
}

// GroupService
// Сервис для работы с отделами.
type GroupService interface {
	GetList(ctx context.Context, params *Params) (*List[Group], *resty.Response, error)
	Create(ctx context.Context, group *Group, params *Params) (*Group, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Group, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, group *Group, params *Params) (*Group, *resty.Response, error)
}

func NewGroupService(client *Client) GroupService {
	e := NewEndpoint(client, "entity/group")
	return newMainService[Group, any, any, any](e)
}
