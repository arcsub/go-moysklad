package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// FactureIn Счет-фактура полученный
// Ключевое слово: facturein
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-schet-faktura-poluchennyj
type FactureIn struct {
	Moment         *Timestamp    `json:"moment,omitempty"`
	Applicable     *bool         `json:"applicable,omitempty"`
	Name           *string       `json:"name,omitempty"`
	AccountID      *uuid.UUID    `json:"accountId,omitempty"`
	Code           *string       `json:"code,omitempty"`
	Contract       *Contract     `json:"contract,omitempty"`
	Created        *Timestamp    `json:"created,omitempty"`
	Deleted        *Timestamp    `json:"deleted,omitempty"`
	Description    *string       `json:"description,omitempty"`
	ExternalCode   *string       `json:"externalCode,omitempty"`
	Files          *Files        `json:"files,omitempty"`
	Group          *Group        `json:"group,omitempty"`
	ID             *uuid.UUID    `json:"id,omitempty"`
	Meta           *Meta         `json:"meta,omitempty"`
	IncomingDate   *Timestamp    `json:"incomingDate,omitempty"`
	Agent          *Counterparty `json:"agent,omitempty"`
	Organization   *Organization `json:"organization,omitempty"`
	Owner          *Employee     `json:"owner,omitempty"`
	Printed        *bool         `json:"printed,omitempty"`
	Published      *bool         `json:"published,omitempty"`
	Rate           *Rate         `json:"rate,omitempty"`
	Shared         *bool         `json:"shared,omitempty"`
	State          *State        `json:"state,omitempty"`
	Sum            *Decimal      `json:"sum,omitempty"`
	SyncID         *uuid.UUID    `json:"syncId,omitempty"`
	Updated        *Timestamp    `json:"updated,omitempty"`
	Supplies       *Supplies     `json:"supplies,omitempty"`
	Payments       *Payments     `json:"payments,omitempty"`
	IncomingNumber *string       `json:"incomingNumber,omitempty"`
	Attributes     Attributes    `json:"attributes,omitempty"`
}

func (f FactureIn) String() string {
	return Stringify(f)
}

func (f FactureIn) MetaType() MetaType {
	return MetaTypeFactureIn
}

// FactureInService
// Сервис для работы со счетами-фактурами полученными.
type FactureInService interface {
	GetList(ctx context.Context, params *Params) (*List[FactureIn], *resty.Response, error)
	Create(ctx context.Context, factureIn *FactureIn, params *Params) (*FactureIn, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, factureInList []*FactureIn, params *Params) (*[]FactureIn, *resty.Response, error)
	DeleteMany(ctx context.Context, factureInList []*FactureIn) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*FactureIn, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, factureIn *FactureIn, params *Params) (*FactureIn, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id *uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error)
	UpdateAttribute(ctx context.Context, id *uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList []*Attribute) (*DeleteManyResponse, *resty.Response, error)
	//endpointTemplate[FactureIn]
	GetPublications(ctx context.Context, id *uuid.UUID) (*MetaArray[Publication], *resty.Response, error)
	GetPublicationByID(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (*Publication, *resty.Response, error)
	Publish(ctx context.Context, id *uuid.UUID, template *Templater) (*Publication, *resty.Response, error)
	DeletePublication(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (bool, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*FactureIn, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewFactureInService(client *Client) FactureInService {
	e := NewEndpoint(client, "entity/facturein")
	return newMainService[FactureIn, any, MetadataAttributeSharedStates, any](e)
}
