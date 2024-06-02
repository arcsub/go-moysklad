package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// RetailDrawerCashIn Внесение денег.
// Ключевое слово: retaildrawercashin
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vnesenie-deneg
type RetailDrawerCashIn struct {
	Moment       *Timestamp    `json:"moment,omitempty"`
	Created      *Timestamp    `json:"created,omitempty"`
	AccountID    *uuid.UUID    `json:"accountId,omitempty"`
	RetailShift  *RetailShift  `json:"retailShift,omitempty"`
	Name         *string       `json:"name,omitempty"`
	Deleted      *Timestamp    `json:"deleted,omitempty"`
	Description  *string       `json:"description,omitempty"`
	Organization *Organization `json:"organization,omitempty"`
	Files        *Files        `json:"files,omitempty"`
	Group        *Group        `json:"group,omitempty"`
	ID           *uuid.UUID    `json:"id,omitempty"`
	Meta         *Meta         `json:"meta,omitempty"`
	Applicable   *bool         `json:"applicable,omitempty"`
	Agent        *Counterparty `json:"agent,omitempty"`
	ExternalCode *string       `json:"externalCode,omitempty"`
	Owner        *Employee     `json:"owner,omitempty"`
	Printed      *bool         `json:"printed,omitempty"`
	Published    *bool         `json:"published,omitempty"`
	Rate         *Rate         `json:"rate,omitempty"`
	Shared       *bool         `json:"shared,omitempty"`
	State        *State        `json:"state,omitempty"`
	Sum          *Decimal      `json:"sum,omitempty"`
	SyncID       *uuid.UUID    `json:"syncId,omitempty"`
	Updated      *Timestamp    `json:"updated,omitempty"`
	Attributes   Attributes    `json:"attributes,omitempty"`
}

func (r RetailDrawerCashIn) String() string {
	return Stringify(r)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (r RetailDrawerCashIn) GetMeta() Meta {
	return Deref(r.Meta)
}

func (r RetailDrawerCashIn) MetaType() MetaType {
	return MetaTypeRetailDrawerCashIn
}

// RetailDrawerCashInService
// Сервис для работы с внесениями денег.
type RetailDrawerCashInService interface {
	GetList(ctx context.Context, params *Params) (*List[RetailDrawerCashIn], *resty.Response, error)
	Create(ctx context.Context, retailDrawerCashIn *RetailDrawerCashIn, params *Params) (*RetailDrawerCashIn, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, retailDrawerCashInList []*RetailDrawerCashIn, params *Params) (*[]RetailDrawerCashIn, *resty.Response, error)
	DeleteMany(ctx context.Context, retailDrawerCashInList []*RetailDrawerCashIn) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*RetailDrawerCashIn, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, retailDrawerCashIn *RetailDrawerCashIn, params *Params) (*RetailDrawerCashIn, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	//endpointTemplate[RetailDrawerCashIn]
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id *uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error)
	UpdateAttribute(ctx context.Context, id *uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList []*Attribute) (*DeleteManyResponse, *resty.Response, error)
	GetPublications(ctx context.Context, id *uuid.UUID) (*MetaArray[Publication], *resty.Response, error)
	GetPublicationByID(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (*Publication, *resty.Response, error)
	Publish(ctx context.Context, id *uuid.UUID, template *Templater) (*Publication, *resty.Response, error)
	DeletePublication(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (bool, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*RetailDrawerCashIn, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
	Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewRetailDrawerCashInService(client *Client) RetailDrawerCashInService {
	e := NewEndpoint(client, "entity/retaildrawercashin")
	return newMainService[RetailDrawerCashIn, any, MetadataAttributeSharedStates, any](e)
}
