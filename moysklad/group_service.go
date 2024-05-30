package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

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
