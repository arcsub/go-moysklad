package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// SalesChannelService
// Сервис для работы с каналами продаж.
type SalesChannelService interface {
	GetList(ctx context.Context, params *Params) (*List[SalesChannel], *resty.Response, error)
	Create(ctx context.Context, salesChannel *SalesChannel, params *Params) (*SalesChannel, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, salesChannelList []*SalesChannel, params *Params) (*[]SalesChannel, *resty.Response, error)
	DeleteMany(ctx context.Context, salesChannelList []*SalesChannel) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*SalesChannel, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, salesChannel *SalesChannel, params *Params) (*SalesChannel, *resty.Response, error)
}

func NewSalesChannelService(client *Client) SalesChannelService {
	e := NewEndpoint(client, "entity/saleschannel")
	return newMainService[SalesChannel, any, any, any](e)
}
