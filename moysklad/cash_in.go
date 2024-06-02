package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// CashIn Приходный ордер.
// Ключевое слово: cashin
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-prihodnyj-order
type CashIn struct {
	Organization   *Organization `json:"organization,omitempty"`
	VatSum         *Decimal      `json:"vatSum,omitempty"`
	Applicable     *bool         `json:"applicable,omitempty"`
	Moment         *Timestamp    `json:"moment,omitempty"`
	Code           *string       `json:"code,omitempty"`
	Contract       *Contract     `json:"contract,omitempty"`
	AccountID      *uuid.UUID    `json:"accountId,omitempty"`
	Deleted        *Timestamp    `json:"deleted,omitempty"`
	Description    *string       `json:"description,omitempty"`
	ExternalCode   *string       `json:"externalCode,omitempty"`
	Files          *Files        `json:"files,omitempty"`
	Group          *Group        `json:"group,omitempty"`
	ID             *uuid.UUID    `json:"id,omitempty"`
	Meta           *Meta         `json:"meta,omitempty"`
	Operations     *Operations   `json:"operations,omitempty"`
	Agent          *Counterparty `json:"agent,omitempty"`
	Created        *Timestamp    `json:"created,omitempty"`
	Owner          *Employee     `json:"owner,omitempty"`
	PaymentPurpose *string       `json:"paymentPurpose,omitempty"`
	Printed        *bool         `json:"printed,omitempty"`
	Project        *Project      `json:"project,omitempty"`
	Published      *bool         `json:"published,omitempty"`
	Rate           *Rate         `json:"rate,omitempty"`
	SalesChannel   *SalesChannel `json:"salesChannel,omitempty"`
	Shared         *bool         `json:"shared,omitempty"`
	State          *State        `json:"state,omitempty"`
	Sum            *Decimal      `json:"sum,omitempty"`
	SyncID         *uuid.UUID    `json:"syncId,omitempty"`
	Updated        *Timestamp    `json:"updated,omitempty"`
	Name           *string       `json:"name,omitempty"`
	FactureIn      *FactureIn    `json:"factureIn,omitempty"`
	Attributes     Attributes    `json:"attributes,omitempty"`
}

func (c CashIn) String() string {
	return Stringify(c)
}

func (c CashIn) MetaType() MetaType {
	return MetaTypeCashIn
}

// BindDocuments Привязка платежей к документам.
// Необходимо передать *Meta документов, к которым необходимо привязать платёж.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-obschie-swedeniq-priwqzka-platezhej-k-dokumentam
func (c *CashIn) BindDocuments(documentsMeta ...*Meta) *CashIn {
	if c.Operations == nil {
		c.Operations = &Operations{}
	}

	for _, meta := range documentsMeta {
		*c.Operations = append(*c.Operations, Operation{Meta: Deref(meta)})
	}

	return c
}

// CashInTemplateArg
// Документ: Приходный ордер (cashin)
// Основание, на котором он может быть создан:
// - Заказ покупателя (customerorder)
// - Возврат поставщику (purchasereturn)
// - Отгрузка (demand)
// - Счет покупателю (invoiceout)
// - Полученный отчет комиссионера (commissionreportin)
type CashInTemplateArg struct {
	CustomerOrder      *MetaWrapper `json:"customerOrder,omitempty"`
	PurchaseReturn     *MetaWrapper `json:"purchaseReturn,omitempty"`
	Demand             *MetaWrapper `json:"demand,omitempty"`
	InvoiceOut         *MetaWrapper `json:"invoiceOut,omitempty"`
	CommissionReportIn *MetaWrapper `json:"commissionReportIn,omitempty"`
}

// CashInService
// Сервис для работы с приходными ордерами.
type CashInService interface {
	GetList(ctx context.Context, params *Params) (*List[CashIn], *resty.Response, error)
	Create(ctx context.Context, cashIn *CashIn, params *Params) (*CashIn, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, cashInList []*CashIn, params *Params) (*[]CashIn, *resty.Response, error)
	DeleteMany(ctx context.Context, cashInList []*CashIn) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id *uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error)
	UpdateAttribute(ctx context.Context, id *uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList []*Attribute) (*DeleteManyResponse, *resty.Response, error)
	//Template(ctx context.Context) (*CashIn, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*CashIn, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, cashIn *CashIn, params *Params) (*CashIn, *resty.Response, error)
	GetPublications(ctx context.Context, id *uuid.UUID) (*MetaArray[Publication], *resty.Response, error)
	GetPublicationByID(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (*Publication, *resty.Response, error)
	Publish(ctx context.Context, id *uuid.UUID, template *Templater) (*Publication, *resty.Response, error)
	DeletePublication(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (bool, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*CashIn, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewCashInService(client *Client) CashInService {
	e := NewEndpoint(client, "entity/cashin")
	return newMainService[CashIn, any, MetadataAttributeSharedStates, any](e)
}
