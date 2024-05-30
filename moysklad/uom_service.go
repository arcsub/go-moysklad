package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// UomService
// Сервис для работы с единицами измерения.
type UomService interface {
	GetList(ctx context.Context, params *Params) (*List[Uom], *resty.Response, error)
	Create(ctx context.Context, uom *Uom, params *Params) (*Uom, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, uomList []*Uom, params *Params) (*[]Uom, *resty.Response, error)
	DeleteMany(ctx context.Context, uomList []*Uom) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Uom, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, uom *Uom, params *Params) (*Uom, *resty.Response, error)
}

func NewUomService(client *Client) UomService {
	e := NewEndpoint(client, "entity/uom")
	return newMainService[Uom, any, any, any](e)
}
