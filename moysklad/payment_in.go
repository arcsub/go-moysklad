package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// PaymentIn Входящий платеж.
// Ключевое слово: paymentin
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vhodqschij-platezh
type PaymentIn struct {
	Meta                *Meta         `json:"meta,omitempty"`
	Applicable          *bool         `json:"applicable,omitempty"`
	AgentAccount        *AgentAccount `json:"agentAccount,omitempty"`
	Moment              *Timestamp    `json:"moment,omitempty"`
	Operations          *Operations   `json:"operations,omitempty"`
	Name                *string       `json:"name,omitempty"`
	Contract            *Contract     `json:"contract,omitempty"`
	Created             *Timestamp    `json:"created,omitempty"`
	Deleted             *Timestamp    `json:"deleted,omitempty"`
	Description         *string       `json:"description,omitempty"`
	ExternalCode        *string       `json:"externalCode,omitempty"`
	Files               *Files        `json:"files,omitempty"`
	Group               *Group        `json:"group,omitempty"`
	ID                  *uuid.UUID    `json:"id,omitempty"`
	IncomingDate        *Timestamp    `json:"incomingDate,omitempty"`
	IncomingNumber      *string       `json:"incomingNumber,omitempty"`
	FactureOut          *FactureOut   `json:"factureOut,omitempty"`
	Agent               *Counterparty `json:"agent,omitempty"`
	Code                *string       `json:"code,omitempty"`
	Organization        *Organization `json:"organization,omitempty"`
	OrganizationAccount *AgentAccount `json:"organizationAccount,omitempty"`
	Owner               *Employee     `json:"owner,omitempty"`
	PaymentPurpose      *string       `json:"paymentPurpose,omitempty"`
	Printed             *bool         `json:"printed,omitempty"`
	Project             *Project      `json:"project,omitempty"`
	Published           *bool         `json:"published,omitempty"`
	Rate                *Rate         `json:"rate,omitempty"`
	Shared              *bool         `json:"shared,omitempty"`
	SalesChannel        *SalesChannel `json:"salesChannel,omitempty"`
	State               *State        `json:"state,omitempty"`
	Sum                 *Decimal      `json:"sum,omitempty"`
	SyncID              *uuid.UUID    `json:"syncId,omitempty"`
	Updated             *Timestamp    `json:"updated,omitempty"`
	AccountID           *uuid.UUID    `json:"accountId,omitempty"`
	Attributes          Attributes    `json:"attributes,omitempty"`
}

func (p PaymentIn) String() string {
	return Stringify(p)
}

func (p PaymentIn) MetaType() MetaType {
	return MetaTypePaymentIn
}

// BindDocuments Привязка платежей к документам.
// Необходимо передать *Meta документов, к которым необходимо привязать платёж.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-obschie-swedeniq-priwqzka-platezhej-k-dokumentam
func (p *PaymentIn) BindDocuments(documentsMeta ...*Meta) *PaymentIn {
	if p.Operations == nil {
		p.Operations = &Operations{}
	}

	for _, meta := range documentsMeta {
		*p.Operations = append(*p.Operations, Operation{Meta: Deref(meta)})
	}

	return p
}

// PaymentInTemplateArg
// Документ: Входящий платеж (paymentin)
// Основание, на котором он может быть создан:
// - Заказ покупателя (customerorder)
// - Возврат поставщику (purchasereturn)
// - Отгрузка (demand)
// - Счет покупателю (invoiceout)
// - Полученный отчет комиссионера (commissionreportin)
type PaymentInTemplateArg struct {
	CustomerOrder      *MetaWrapper `json:"customerOrder,omitempty"`
	PurchaseReturn     *MetaWrapper `json:"purchaseReturn,omitempty"`
	Demand             *MetaWrapper `json:"demand,omitempty"`
	InvoiceOut         *MetaWrapper `json:"invoiceOut,omitempty"`
	CommissionReportIn *MetaWrapper `json:"commissionReportIn,omitempty"`
}

// PaymentInService
// Сервис для работы с входящими платежами.
type PaymentInService interface {
	GetList(ctx context.Context, params *Params) (*List[PaymentIn], *resty.Response, error)
	Create(ctx context.Context, paymentIn *PaymentIn, params *Params) (*PaymentIn, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, paymentInList []*PaymentIn, params *Params) (*[]PaymentIn, *resty.Response, error)
	DeleteMany(ctx context.Context, paymentInList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*PaymentIn, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, paymentIn *PaymentIn, params *Params) (*PaymentIn, *resty.Response, error)
	//endpointTemplate[PaymentIn]
	//endpointTemplateBasedOn[PaymentIn, PaymentInTemplateArg]
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
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*PaymentIn, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	MoveToTrash(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewPaymentInService(client *Client) PaymentInService {
	e := NewEndpoint(client, "entity/paymentin")
	return newMainService[PaymentIn, any, MetadataAttributeSharedStates, any](e)
}
