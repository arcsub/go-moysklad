package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// CashOut Расходный ордер.
// Ключевое слово: cashout
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-rashodnyj-order
type CashOut struct {
	Name           *string       `json:"name,omitempty"`
	Deleted        *Timestamp    `json:"deleted,omitempty"`
	Applicable     *bool         `json:"applicable,omitempty"`
	AccountID      *uuid.UUID    `json:"accountId,omitempty"`
	Code           *string       `json:"code,omitempty"`
	Contract       *Contract     `json:"contract,omitempty"`
	Created        *Timestamp    `json:"created,omitempty"`
	Organization   *Organization `json:"organization,omitempty"`
	Description    *string       `json:"description,omitempty"`
	ExpenseItem    *ExpenseItem  `json:"expenseItem,omitempty"`
	ExternalCode   *string       `json:"externalCode,omitempty"`
	Files          *Files        `json:"files,omitempty"`
	Group          *Group        `json:"group,omitempty"`
	Owner          *Employee     `json:"owner,omitempty"`
	Meta           *Meta         `json:"meta,omitempty"`
	Moment         *Timestamp    `json:"moment,omitempty"`
	Operations     *Operations   `json:"operations,omitempty"`
	Agent          *Counterparty `json:"agent,omitempty"`
	ID             *uuid.UUID    `json:"id,omitempty"`
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
	VatSum         *Decimal      `json:"vatSum,omitempty"`
	FactureOut     *FactureOut   `json:"factureOut,omitempty"`
	Attributes     Attributes    `json:"attributes,omitempty"`
}

func (c CashOut) String() string {
	return Stringify(c)
}

func (c CashOut) MetaType() MetaType {
	return MetaTypeCashOut
}

// BindDocuments Привязка платежей к документам.
// Необходимо передать *Meta документов, к которым необходимо привязать платёж.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-obschie-swedeniq-priwqzka-platezhej-k-dokumentam
func (c *CashOut) BindDocuments(documentsMeta ...*Meta) *CashOut {
	if c.Operations == nil {
		c.Operations = &Operations{}
	}

	for _, meta := range documentsMeta {
		*c.Operations = append(*c.Operations, Operation{Meta: Deref(meta)})
	}

	return c
}

// CashOutTemplateArg
// Документ: Расходный ордер (cashout)
// Основание, на котором он может быть создан:
// - Возврат покупателя (salesreturn)
// - Приемка (supply)
// - Счет поставщика (invoicein)
// - Заказ поставщику (purchaseorder)
// - Выданный отчет комиссионера (commissionreportout)
type CashOutTemplateArg struct {
	SalesReturn         *MetaWrapper `json:"salesReturn,omitempty"`
	Supply              *MetaWrapper `json:"supply,omitempty"`
	InvoiceIn           *MetaWrapper `json:"invoiceIn,omitempty"`
	PurchaseOrder       *MetaWrapper `json:"purchaseOrder,omitempty"`
	CommissionReportOut *MetaWrapper `json:"commissionReportOut,omitempty"`
}

// CashOutService cashout
// Сервис для работы с расходными ордерами.
type CashOutService interface {
	GetList(ctx context.Context, params *Params) (*List[CashOut], *resty.Response, error)
	Create(ctx context.Context, cashOut *CashOut, params *Params) (*CashOut, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, cashOutList []*CashOut, params *Params) (*[]CashOut, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	DeleteMany(ctx context.Context, cashOutList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id *uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error)
	UpdateAttribute(ctx context.Context, id *uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	//Template(ctx context.Context) (*CashOut, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*CashOut, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, cashOut *CashOut, params *Params) (*CashOut, *resty.Response, error)
	GetPublications(ctx context.Context, id *uuid.UUID) (*MetaArray[Publication], *resty.Response, error)
	GetPublicationByID(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (*Publication, *resty.Response, error)
	Publish(ctx context.Context, id *uuid.UUID, template *Templater) (*Publication, *resty.Response, error)
	DeletePublication(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (bool, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*CashOut, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	MoveToTrash(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewCashOutService(client *Client) CashOutService {
	e := NewEndpoint(client, "entity/cashout")
	return newMainService[CashOut, any, MetadataAttributeSharedStates, any](e)
}
