package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// InvoiceIn Счет поставщика.
// Ключевое слово: invoicein
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-schet-postawschika
type InvoiceIn struct {
	OrganizationAccount  *AgentAccount               `json:"organizationAccount,omitempty"`
	Created              *Timestamp                  `json:"created,omitempty"`
	PayedSum             *float64                    `json:"payedSum,omitempty"`
	Applicable           *bool                       `json:"applicable,omitempty"`
	Supplies             *Supplies                   `json:"supplies,omitempty"`
	Code                 *string                     `json:"code,omitempty"`
	Contract             *Contract                   `json:"contract,omitempty"`
	Owner                *Employee                   `json:"owner,omitempty"`
	Deleted              *Timestamp                  `json:"deleted,omitempty"`
	Description          *string                     `json:"description,omitempty"`
	ExternalCode         *string                     `json:"externalCode,omitempty"`
	Files                *Files                      `json:"files,omitempty"`
	AccountID            *uuid.UUID                  `json:"accountId,omitempty"`
	ID                   *uuid.UUID                  `json:"id,omitempty"`
	IncomingDate         *Timestamp                  `json:"incomingDate,omitempty"`
	IncomingNumber       *string                     `json:"incomingNumber,omitempty"`
	Meta                 *Meta                       `json:"meta,omitempty"`
	Moment               *Timestamp                  `json:"moment,omitempty"`
	Name                 *string                     `json:"name,omitempty"`
	Organization         *Organization               `json:"organization,omitempty"`
	Group                *Group                      `json:"group,omitempty"`
	Agent                *Counterparty               `json:"agent,omitempty"`
	AgentAccount         *AgentAccount               `json:"agentAccount,omitempty"`
	PaymentPlannedMoment *Timestamp                  `json:"paymentPlannedMoment,omitempty"`
	Positions            *Positions[InvoicePosition] `json:"positions,omitempty"`
	Printed              *bool                       `json:"printed,omitempty"`
	Project              *Project                    `json:"project,omitempty"`
	Published            *bool                       `json:"published,omitempty"`
	Rate                 *Rate                       `json:"rate,omitempty"`
	Shared               *bool                       `json:"shared,omitempty"`
	ShippedSum           *float64                    `json:"shippedSum,omitempty"`
	State                *State                      `json:"state,omitempty"`
	Store                *Store                      `json:"store,omitempty"`
	Sum                  *float64                    `json:"sum,omitempty"`
	SyncID               *uuid.UUID                  `json:"syncId,omitempty"`
	Updated              *Timestamp                  `json:"updated,omitempty"`
	VatEnabled           *bool                       `json:"vatEnabled,omitempty"`
	VatIncluded          *bool                       `json:"vatIncluded,omitempty"`
	VatSum               *float64                    `json:"vatSum,omitempty"`
	Payments             *Payments                   `json:"payments,omitempty"`
	PurchaseOrder        *PurchaseOrder              `json:"purchaseOrder,omitempty"`
	Attributes           Attributes                  `json:"attributes,omitempty"`
}

func (i InvoiceIn) String() string {
	return Stringify(i)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (i InvoiceIn) GetMeta() Meta {
	return Deref(i.Meta)
}

func (i InvoiceIn) MetaType() MetaType {
	return MetaTypeInvoiceIn
}

type InvoicesIn = Slice[InvoiceIn]

// InvoiceInPosition Позиция Счета поставщика.
// Ключевое слово: invoiceposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-schet-postawschika-scheta-postawschikow-pozicii-scheta-postawschika
type InvoiceInPosition struct {
	InvoicePosition
}

func (i InvoiceInPosition) String() string {
	return Stringify(i)
}

func (i InvoiceInPosition) MetaType() MetaType {
	return MetaTypeInvoicePosition
}

// InvoiceInTemplateArg
// Документ: Счет поставщика (invoicein)
// Основание, на котором он может быть создан:
// - Заказ поставщику (purchaseorder)
type InvoiceInTemplateArg struct {
	PurchaseOrder *MetaWrapper `json:"purchaseOrder,omitempty"`
}

// InvoiceInService
// Сервис для работы со счетами поставщиков.
type InvoiceInService interface {
	GetList(ctx context.Context, params *Params) (*List[InvoiceIn], *resty.Response, error)
	Create(ctx context.Context, invoiceIn *InvoiceIn, params *Params) (*InvoiceIn, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, invoiceInList []*InvoiceIn, params *Params) (*[]InvoiceIn, *resty.Response, error)
	DeleteMany(ctx context.Context, invoiceInList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*InvoiceIn, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, invoiceIn *InvoiceIn, params *Params) (*InvoiceIn, *resty.Response, error)
	//endpointTemplate[InvoiceIn]
	//endpointTemplateBasedOn[InvoiceIn, InvoiceInTemplateArg]
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	GetPositions(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[InvoiceInPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, params *Params) (*InvoiceInPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, position *InvoiceInPosition, params *Params) (*InvoiceInPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id *uuid.UUID, position *InvoiceInPosition) (*InvoiceInPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id *uuid.UUID, positions []*InvoiceInPosition) (*[]InvoiceInPosition, *resty.Response, error)
	DeletePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID) (bool, *resty.Response, error)
	GetPositionTrackingCodes(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID) (*MetaArray[TrackingCode], *resty.Response, error)
	CreateOrUpdatePositionTrackingCodes(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, trackingCodes TrackingCodes) (*[]TrackingCode, *resty.Response, error)
	DeletePositionTrackingCodes(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, trackingCodes TrackingCodes) (*DeleteManyResponse, *resty.Response, error)
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
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*InvoiceIn, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	MoveToTrash(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewInvoiceInService(client *Client) InvoiceInService {
	e := NewEndpoint(client, "entity/invoicein")
	return newMainService[InvoiceIn, InvoiceInPosition, MetadataAttributeSharedStates, any](e)
}
