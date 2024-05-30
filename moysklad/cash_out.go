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
	AccountID      *uuid.UUID    `json:"accountId,omitempty"`      // ID учетной записи
	Agent          *Counterparty `json:"agent,omitempty"`          // Метаданные контрагента
	Applicable     *bool         `json:"applicable,omitempty"`     // Отметка о проведении
	Attributes     *Attributes   `json:"attributes,omitempty"`     // Коллекция метаданных доп. полей. Поля объекта
	Code           *string       `json:"code,omitempty"`           // Код
	Contract       *Contract     `json:"contract,omitempty"`       // Метаданные договора
	Created        *Timestamp    `json:"created,omitempty"`        // Дата создания
	Deleted        *Timestamp    `json:"deleted,omitempty"`        // Момент последнего удаления
	Description    *string       `json:"description,omitempty"`    // Комментарий
	ExpenseItem    *ExpenseItem  `json:"expenseItem,omitempty"`    // Метаданные Статьи расходов
	ExternalCode   *string       `json:"externalCode,omitempty"`   // Внешний код
	Files          *Files        `json:"files,omitempty"`          // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group          *Group        `json:"group,omitempty"`          // Отдел сотрудника
	ID             *uuid.UUID    `json:"id,omitempty"`             // ID сущности
	Meta           *Meta         `json:"meta,omitempty"`           // Метаданные
	Moment         *Timestamp    `json:"moment,omitempty"`         // Дата документа
	Name           *string       `json:"name,omitempty"`           // Наименование
	Organization   *Organization `json:"organization,omitempty"`   // Метаданные юрлица
	Owner          *Employee     `json:"owner,omitempty"`          // Владелец (Сотрудник)
	PaymentPurpose *string       `json:"paymentPurpose,omitempty"` // Основание
	Printed        *bool         `json:"printed,omitempty"`        // Напечатан ли документ
	Project        *Project      `json:"project,omitempty"`        // Метаданные проекта
	Published      *bool         `json:"published,omitempty"`      // Опубликован ли документ
	Rate           *Rate         `json:"rate,omitempty"`           // Валюта
	SalesChannel   *SalesChannel `json:"salesChannel,omitempty"`   // Метаданные канала продаж
	Shared         *bool         `json:"shared,omitempty"`         // Общий доступ
	State          *State        `json:"state,omitempty"`          // Метаданные статуса
	Sum            *Decimal      `json:"sum,omitempty"`            // Сумма
	SyncID         *uuid.UUID    `json:"syncId,omitempty"`         // ID синхронизации. После заполнения недоступен для изменения
	Updated        *Timestamp    `json:"updated,omitempty"`        // Момент последнего обновления
	VatSum         *Decimal      `json:"vatSum,omitempty"`         // Сумма включая НДС
	FactureOut     *FactureOut   `json:"factureOut,omitempty"`     // Ссылка на Счет-фактуру выданный, с которым связан этот платеж в формате Метаданных
	Operations     *Operations   `json:"operations,omitempty"`     // Массив ссылок на связанные операции в формате Метаданных
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
	DeleteMany(ctx context.Context, cashOutList []*CashOut) (*DeleteManyResponse, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id *uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error)
	UpdateAttribute(ctx context.Context, id *uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList []*Attribute) (*DeleteManyResponse, *resty.Response, error)
	//Template(ctx context.Context) (*CashOut, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*CashOut, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, cashOut *CashOut, params *Params) (*CashOut, *resty.Response, error)
	GetPublications(ctx context.Context, id *uuid.UUID) (*MetaArray[Publication], *resty.Response, error)
	GetPublicationByID(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (*Publication, *resty.Response, error)
	Publish(ctx context.Context, id *uuid.UUID, template *Templater) (*Publication, *resty.Response, error)
	DeletePublication(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (bool, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*CashOut, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewCashOutService(client *Client) CashOutService {
	e := NewEndpoint(client, "entity/cashout")
	return newMainService[CashOut, any, MetadataAttributeSharedStates, any](e)
}
