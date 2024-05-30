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
	AccountID      *uuid.UUID    `json:"accountId,omitempty"`      // ID учетной записи
	Agent          *Counterparty `json:"agent,omitempty"`          // Метаданные контрагента
	Applicable     *bool         `json:"applicable,omitempty"`     // Отметка о проведении
	Attributes     *Attributes   `json:"attributes,omitempty"`     // Коллекция метаданных доп. полей. Поля объекта
	Code           *string       `json:"code,omitempty"`           // Код
	Contract       *Contract     `json:"contract,omitempty"`       // Метаданные договора
	Created        *Timestamp    `json:"created,omitempty"`        // Дата создания
	Deleted        *Timestamp    `json:"deleted,omitempty"`        // Момент последнего удаления
	Description    *string       `json:"description,omitempty"`    // Комментарий
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
	FactureIn      *FactureIn    `json:"factureIn,omitempty"`      // Ссылка на Счет-фактуру полученный, с которым связан этот платеж в формате Метаданных
	Operations     *Operations   `json:"operations,omitempty"`     // Массив ссылок на связанные операции в формате Метаданных
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
