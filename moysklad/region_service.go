package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// RegionService
// Сервис для работы с регионами.
type RegionService interface {
	GetList(ctx context.Context, params *Params) (*List[Region], *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Region, *resty.Response, error)
}

func NewRegionService(client *Client) RegionService {
	e := NewEndpoint(client, "entity/region")
	return newMainService[Region, any, any, any](e)
}
