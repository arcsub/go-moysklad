package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// CounterPartyAdjustment Корректировка баланса контрагента.
// Ключевое слово: counterpartyadjustment
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-korrektirowka-balansa-kontragenta
type CounterPartyAdjustment struct {
	ExternalCode *string       `json:"externalCode,omitempty"`
	Printed      *bool         `json:"printed,omitempty"`
	AccountID    *uuid.UUID    `json:"accountId,omitempty"`
	Group        *Group        `json:"group,omitempty"`
	Files        *Files        `json:"files,omitempty"`
	Applicable   *bool         `json:"applicable,omitempty"`
	Updated      *Timestamp    `json:"updated,omitempty"`
	Created      *Timestamp    `json:"created,omitempty"`
	Deleted      *Timestamp    `json:"deleted,omitempty"`
	Description  *string       `json:"description,omitempty"`
	Name         *string       `json:"name,omitempty"`
	Agent        *Counterparty `json:"agent,omitempty"`
	Meta         *Meta         `json:"meta,omitempty"`
	Moment       *Timestamp    `json:"moment,omitempty"`
	Organization *Organization `json:"organization,omitempty"`
	Owner        *Employee     `json:"owner,omitempty"`
	ID           *uuid.UUID    `json:"id,omitempty"`
	Published    *bool         `json:"published,omitempty"`
	Shared       *bool         `json:"shared,omitempty"`
	Sum          *float64      `json:"sum,omitempty"`
	Attributes   Attributes    `json:"attributes,omitempty"`
}

func (c CounterPartyAdjustment) String() string {
	return Stringify(c)
}

func (c CounterPartyAdjustment) MetaType() MetaType {
	return MetaTypeCounterPartyAdjustment
}

// CounterPartyAdjustmentService
// Сервис для работы с корректировками баланса контрагента.
type CounterPartyAdjustmentService interface {
	GetList(ctx context.Context, params *Params) (*List[CounterPartyAdjustment], *resty.Response, error)
	Create(ctx context.Context, counterPartyAdjustment *CounterPartyAdjustment, params *Params) (*CounterPartyAdjustment, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, counterPartyAdjustmentList []*CounterPartyAdjustment, params *Params) (*[]CounterPartyAdjustment, *resty.Response, error)
	DeleteMany(ctx context.Context, counterPartyAdjustmentList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*CounterPartyAdjustment, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, counterPartyAdjustment *CounterPartyAdjustment, params *Params) (*CounterPartyAdjustment, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
	MoveToTrash(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewCounterPartyAdjustmentService(client *Client) CounterPartyAdjustmentService {
	e := NewEndpoint(client, "entity/counterpartyadjustment")
	return newMainService[CounterPartyAdjustment, any, MetadataAttributeSharedStates, any](e)
}
