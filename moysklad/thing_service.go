package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// ThingService
// Сервис для работы с серийными номерами.
type ThingService interface {
	GetList(ctx context.Context, params *Params) (*List[Thing], *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Thing, *resty.Response, error)
}

func NewThingService(client *Client) ThingService {
	e := NewEndpoint(client, "entity/thing")
	return newMainService[Thing, any, any, any](e)
}
