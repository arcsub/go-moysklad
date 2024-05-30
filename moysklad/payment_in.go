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
	AccountID           *uuid.UUID    `json:"accountId,omitempty"`           // ID учетной записи
	Agent               *Counterparty `json:"agent,omitempty"`               // Метаданные контрагента
	AgentAccount        *AgentAccount `json:"agentAccount,omitempty"`        // Метаданные счета контрагента
	Applicable          *bool         `json:"applicable,omitempty"`          // Отметка о проведении
	Attributes          *Attributes   `json:"attributes,omitempty"`          // Коллекция метаданных доп. полей. Поля объекта
	Code                *string       `json:"code,omitempty"`                // Код
	Contract            *Contract     `json:"contract,omitempty"`            // Метаданные договора
	Created             *Timestamp    `json:"created,omitempty"`             // Дата создания
	Deleted             *Timestamp    `json:"deleted,omitempty"`             // Момент последнего удаления
	Description         *string       `json:"description,omitempty"`         // Комментарий
	ExternalCode        *string       `json:"externalCode,omitempty"`        // Внешний код
	Files               *Files        `json:"files,omitempty"`               // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group               *Group        `json:"group,omitempty"`               // Отдел сотрудника
	ID                  *uuid.UUID    `json:"id,omitempty"`                  // ID сущности
	IncomingDate        *Timestamp    `json:"incomingDate,omitempty"`        // Входящая дата
	IncomingNumber      *string       `json:"incomingNumber,omitempty"`      // Входящий номер
	Meta                *Meta         `json:"meta,omitempty"`                // Метаданные
	Moment              *Timestamp    `json:"moment,omitempty"`              // Дата документа
	Name                *string       `json:"name,omitempty"`                // Наименование
	Organization        *Organization `json:"organization,omitempty"`        // Метаданные юрлица
	OrganizationAccount *AgentAccount `json:"organizationAccount,omitempty"` // Метаданные счета юрлица
	Owner               *Employee     `json:"owner,omitempty"`               // Владелец (Сотрудник)
	PaymentPurpose      *string       `json:"paymentPurpose,omitempty"`      // Назначение платежа
	Printed             *bool         `json:"printed,omitempty"`             // Напечатан ли документ
	Project             *Project      `json:"project,omitempty"`             // Проект
	Published           *bool         `json:"published,omitempty"`           // Опубликован ли документ
	Rate                *Rate         `json:"rate,omitempty"`                // Валюта
	Shared              *bool         `json:"shared,omitempty"`              // Общий доступ
	SalesChannel        *SalesChannel `json:"salesChannel,omitempty"`        // Метаданные канала продаж
	State               *State        `json:"state,omitempty"`               // Метаданные статуса
	Sum                 *Decimal      `json:"sum,omitempty"`                 // Сумма
	SyncID              *uuid.UUID    `json:"syncId,omitempty"`              // ID синхронизации. После заполнения недоступен для изменения
	Updated             *Timestamp    `json:"updated,omitempty"`             // Момент последнего обновления
	FactureOut          *FactureOut   `json:"factureOut,omitempty"`          // Ссылка на Счет-фактуру выданный, с которым связан этот платеж в формате Метаданных
	Operations          *Operations   `json:"operations,omitempty"`          // Массив ссылок на связанные операции в формате Метаданных
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
	DeleteMany(ctx context.Context, paymentInList []*PaymentIn) (*DeleteManyResponse, *resty.Response, error)
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
	DeleteAttributes(ctx context.Context, attributeList []*Attribute) (*DeleteManyResponse, *resty.Response, error)
	GetPublications(ctx context.Context, id *uuid.UUID) (*MetaArray[Publication], *resty.Response, error)
	GetPublicationByID(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (*Publication, *resty.Response, error)
	Publish(ctx context.Context, id *uuid.UUID, template *Templater) (*Publication, *resty.Response, error)
	DeletePublication(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (bool, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*PaymentIn, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewPaymentInService(client *Client) PaymentInService {
	e := NewEndpoint(client, "entity/paymentin")
	return newMainService[PaymentIn, any, MetadataAttributeSharedStates, any](e)
}
