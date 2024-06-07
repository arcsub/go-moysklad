package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// PaymentOut Исходящий платеж.
// Ключевое слово: paymentout
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-ishodqschij-platezh
type PaymentOut struct {
	Moment              *Timestamp    `json:"moment,omitempty"`
	Applicable          *bool         `json:"applicable,omitempty"`
	AgentAccount        *AgentAccount `json:"agentAccount,omitempty"`
	Name                *string       `json:"name,omitempty"`
	Operations          *Operations   `json:"operations,omitempty"`
	Organization        *Organization `json:"organization,omitempty"`
	Contract            *Contract     `json:"contract,omitempty"`
	Created             *Timestamp    `json:"created,omitempty"`
	Deleted             *Timestamp    `json:"deleted,omitempty"`
	Description         *string       `json:"description,omitempty"`
	ExpenseItem         *ExpenseItem  `json:"expenseItem,omitempty"`
	ExternalCode        *string       `json:"externalCode,omitempty"`
	Files               *Files        `json:"files,omitempty"`
	Group               *Group        `json:"group,omitempty"`
	ID                  *uuid.UUID    `json:"id,omitempty"`
	Meta                *Meta         `json:"meta,omitempty"`
	FactureIn           *FactureIn    `json:"factureIn,omitempty"`
	Agent               *Counterparty `json:"agent,omitempty"`
	Code                *string       `json:"code,omitempty"`
	OrganizationAccount *AgentAccount `json:"organizationAccount,omitempty"`
	Owner               *Employee     `json:"owner,omitempty"`
	PaymentPurpose      *string       `json:"paymentPurpose,omitempty"`
	Printed             *bool         `json:"printed,omitempty"`
	Project             *Project      `json:"project,omitempty"`
	Published           *bool         `json:"published,omitempty"`
	Rate                *Rate         `json:"rate,omitempty"`
	SalesChannel        *SalesChannel `json:"salesChannel,omitempty"`
	Shared              *bool         `json:"shared,omitempty"`
	State               *State        `json:"state,omitempty"`
	Sum                 *float64      `json:"sum,omitempty"`
	SyncID              *uuid.UUID    `json:"syncId,omitempty"`
	Updated             *Timestamp    `json:"updated,omitempty"`
	VatSum              *float64      `json:"vatSum,omitempty"`
	AccountID           *uuid.UUID    `json:"accountId,omitempty"`
	Attributes          Attributes    `json:"attributes,omitempty"`
}

func (p PaymentOut) String() string {
	return Stringify(p)
}

func (p PaymentOut) MetaType() MetaType {
	return MetaTypePaymentOut
}

// PaymentOutTemplateArg
// Документ: Исходящий платеж (paymentout)
// Основание, на котором он может быть создан:
// - Возврат покупателя (salesreturn)
// - Приемка (supply)
// - Счет поставщика (invoicein)
// - Заказ поставщику (purchaseorder)
// - Выданный отчет комиссионера (commissionreportout)
type PaymentOutTemplateArg struct {
	SalesReturn         *MetaWrapper `json:"salesReturn,omitempty"`
	Supply              *MetaWrapper `json:"supply,omitempty"`
	InvoiceIn           *MetaWrapper `json:"invoiceIn,omitempty"`
	PurchaseOrder       *MetaWrapper `json:"purchaseOrder,omitempty"`
	CommissionReportOut *MetaWrapper `json:"commissionReportOut,omitempty"`
}

// PaymentOutService
// Сервис для работы с исходящими платежами.
type PaymentOutService interface {
	GetList(ctx context.Context, params *Params) (*List[PaymentOut], *resty.Response, error)
	Create(ctx context.Context, paymentOut *PaymentOut, params *Params) (*PaymentOut, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, paymentOutList []*PaymentOut, params *Params) (*[]PaymentOut, *resty.Response, error)
	DeleteMany(ctx context.Context, paymentOutList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*PaymentOut, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, paymentOut *PaymentOut, params *Params) (*PaymentOut, *resty.Response, error)
	//endpointTemplate[PaymentOut]
	//endpointTemplateBasedOn[PaymentOut, PaymentOutTemplateArg]
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
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*PaymentOut, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	MoveToTrash(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewPaymentOutService(client *Client) PaymentOutService {
	e := NewEndpoint(client, "entity/paymentout")
	return newMainService[PaymentOut, any, MetadataAttributeSharedStates, any](e)
}
