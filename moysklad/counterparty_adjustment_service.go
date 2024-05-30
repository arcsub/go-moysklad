package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// CounterPartyAdjustmentService
// Сервис для работы с корректировками баланса контрагента.
type CounterPartyAdjustmentService interface {
	GetList(ctx context.Context, params *Params) (*List[CounterPartyAdjustment], *resty.Response, error)
	Create(ctx context.Context, counterPartyAdjustment *CounterPartyAdjustment, params *Params) (*CounterPartyAdjustment, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, counterPartyAdjustmentList []*CounterPartyAdjustment, params *Params) (*[]CounterPartyAdjustment, *resty.Response, error)
	DeleteMany(ctx context.Context, counterPartyAdjustmentList []*CounterPartyAdjustment) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*CounterPartyAdjustment, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, counterPartyAdjustment *CounterPartyAdjustment, params *Params) (*CounterPartyAdjustment, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
	Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewCounterPartyAdjustmentService(client *Client) CounterPartyAdjustmentService {
	e := NewEndpoint(client, "entity/counterpartyadjustment")
	return newMainService[CounterPartyAdjustment, any, MetadataAttributeSharedStates, any](e)
}
