package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// ServiceService
// Сервис для работы с услугами.
type ServiceService interface {
	GetList(ctx context.Context, params *Params) (*List[Service], *resty.Response, error)
	Create(ctx context.Context, service *Service, params *Params) (*Service, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, serviceList []*Service, params *Params) (*[]Service, *resty.Response, error)
	DeleteMany(ctx context.Context, serviceList []*Service) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Service, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, service *Service, params *Params) (*Service, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*Service, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
}

func NewServiceService(client *Client) ServiceService {
	e := NewEndpoint(client, "entity/service")
	return newMainService[Service, any, any, any](e)
}
