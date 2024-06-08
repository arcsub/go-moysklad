package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// RetailDrawerCashOut Выплата денег.
// Ключевое слово: retaildrawercashout
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vyplata-deneg
type RetailDrawerCashOut struct {
	Meta         *Meta         `json:"meta,omitempty"`
	Applicable   *bool         `json:"applicable,omitempty"`
	Moment       *Timestamp    `json:"moment,omitempty"`
	Name         *string       `json:"name,omitempty"`
	Code         *string       `json:"code,omitempty"`
	Created      *Timestamp    `json:"created,omitempty"`
	Deleted      *Timestamp    `json:"deleted,omitempty"`
	Description  *string       `json:"description,omitempty"`
	ExternalCode *string       `json:"externalCode,omitempty"`
	Files        *Files        `json:"files,omitempty"`
	Group        *Group        `json:"group,omitempty"`
	ID           *uuid.UUID    `json:"id,omitempty"`
	RetailShift  *RetailShift  `json:"retailShift,omitempty"`
	Agent        *Counterparty `json:"agent,omitempty"`
	AccountID    *uuid.UUID    `json:"accountId,omitempty"`
	Organization *Organization `json:"organization,omitempty"`
	Owner        *Employee     `json:"owner,omitempty"`
	Printed      *bool         `json:"printed,omitempty"`
	Published    *bool         `json:"published,omitempty"`
	Rate         *Rate         `json:"rate,omitempty"`
	Shared       *bool         `json:"shared,omitempty"`
	State        *State        `json:"state,omitempty"`
	Sum          *float64      `json:"sum,omitempty"`
	SyncID       *uuid.UUID    `json:"syncId,omitempty"`
	Updated      *Timestamp    `json:"updated,omitempty"`
	Attributes   Attributes    `json:"attributes,omitempty"`
}

func (r RetailDrawerCashOut) String() string {
	return Stringify(r)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (r RetailDrawerCashOut) GetMeta() Meta {
	return Deref(r.Meta)
}

func (r RetailDrawerCashOut) MetaType() MetaType {
	return MetaTypeRetailDrawerCashOut
}

// RetailDrawerCashOutService
// Сервис для работы с выплатами денег.
type RetailDrawerCashOutService interface {
	GetList(ctx context.Context, params *Params) (*List[RetailDrawerCashOut], *resty.Response, error)
	Create(ctx context.Context, retailDrawerCashOut *RetailDrawerCashOut, params *Params) (*RetailDrawerCashOut, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, retailDrawerCashOutList []*RetailDrawerCashOut, params *Params) (*[]RetailDrawerCashOut, *resty.Response, error)
	DeleteMany(ctx context.Context, retailDrawerCashOutList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*RetailDrawerCashOut, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, retailDrawerCashOut *RetailDrawerCashOut, params *Params) (*RetailDrawerCashOut, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetaAttributesSharedStatesWrapper, *resty.Response, error)
	//endpointTemplate[RetailDrawerCashOut]
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id *uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error)
	UpdateAttribute(ctx context.Context, id *uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	GetPublications(ctx context.Context, id *uuid.UUID) (*MetaArray[Publication], *resty.Response, error)
	GetPublicationByID(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (*Publication, *resty.Response, error)
	Publish(ctx context.Context, id *uuid.UUID, template *Templater) (*Publication, *resty.Response, error)
	DeletePublication(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (bool, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*RetailDrawerCashOut, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
	MoveToTrash(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewRetailDrawerCashOutService(client *Client) RetailDrawerCashOutService {
	e := NewEndpoint(client, "entity/retaildrawercashout")
	return newMainService[RetailDrawerCashOut, any, MetaAttributesSharedStatesWrapper, any](e)
}
