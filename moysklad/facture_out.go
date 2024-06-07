package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// FactureOut Счет-фактура выданный.
// Ключевое слово: factureout
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-schet-faktura-wydannyj
type FactureOut struct {
	Organization    *Organization    `json:"organization,omitempty"`
	Deleted         *Timestamp       `json:"deleted,omitempty"`
	Applicable      *bool            `json:"applicable,omitempty"`
	AccountID       *uuid.UUID       `json:"accountId,omitempty"`
	Code            *string          `json:"code,omitempty"`
	Contract        *Contract        `json:"contract,omitempty"`
	Created         *Timestamp       `json:"created,omitempty"`
	Owner           *Employee        `json:"owner,omitempty"`
	Description     *string          `json:"description,omitempty"`
	ExternalCode    *string          `json:"externalCode,omitempty"`
	Files           *Files           `json:"files,omitempty"`
	Group           *Group           `json:"group,omitempty"`
	ID              *uuid.UUID       `json:"id,omitempty"`
	Printed         *bool            `json:"printed,omitempty"`
	Moment          *Timestamp       `json:"moment,omitempty"`
	Name            *string          `json:"name,omitempty"`
	PaymentDate     *Timestamp       `json:"paymentDate,omitempty"`
	Agent           *Counterparty    `json:"agent,omitempty"`
	Meta            *Meta            `json:"meta,omitempty"`
	Published       *bool            `json:"published,omitempty"`
	Rate            *Rate            `json:"rate,omitempty"`
	Shared          *bool            `json:"shared,omitempty"`
	State           *State           `json:"state,omitempty"`
	StateContractId *string          `json:"stateContractId,omitempty"`
	Sum             *float64         `json:"sum,omitempty"`
	SyncID          *uuid.UUID       `json:"syncId,omitempty"`
	Updated         *Timestamp       `json:"updated,omitempty"`
	Demands         *Demands         `json:"demands,omitempty"`
	Payments        *Payments        `json:"payments,omitempty"`
	Returns         *PurchaseReturns `json:"returns,omitempty"`
	Consignee       *Counterparty    `json:"consignee,omitempty"`
	PaymentNumber   *string          `json:"paymentNumber,omitempty"`
	Attributes      Attributes       `json:"attributes,omitempty"`
}

func (f FactureOut) String() string {
	return Stringify(f)
}

func (f FactureOut) MetaType() MetaType {
	return MetaTypeFactureOut
}

// FactureOutService
// Сервис для работы со счетами-фактурами выданными.
type FactureOutService interface {
	GetList(ctx context.Context, params *Params) (*List[FactureOut], *resty.Response, error)
	Create(ctx context.Context, factureOut *FactureOut, params *Params) (*FactureOut, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, factureOutList []*FactureOut, params *Params) (*[]FactureOut, *resty.Response, error)
	DeleteMany(ctx context.Context, factureOutList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*FactureOut, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, factureOut *FactureOut, params *Params) (*FactureOut, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
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
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*FactureOut, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	MoveToTrash(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewFactureOutService(client *Client) FactureOutService {
	e := NewEndpoint(client, "entity/factureout")
	return newMainService[FactureOut, any, MetadataAttributeSharedStates, any](e)
}
