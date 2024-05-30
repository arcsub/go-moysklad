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
	AccountID            *uuid.UUID                  `json:"accountId,omitempty"`            // ID учетной записи
	Agent                *Counterparty               `json:"agent,omitempty"`                // Метаданные контрагента
	AgentAccount         *AgentAccount               `json:"agentAccount,omitempty"`         // Метаданные счета контрагента
	Applicable           *bool                       `json:"applicable,omitempty"`           // Отметка о проведении
	Attributes           *Attributes                 `json:"attributes,omitempty"`           // Коллекция метаданных доп. полей
	Code                 *string                     `json:"code,omitempty"`                 // Код
	Contract             *Contract                   `json:"contract,omitempty"`             // Метаданные договора
	Created              *Timestamp                  `json:"created,omitempty"`              // Дата создания
	Deleted              *Timestamp                  `json:"deleted,omitempty"`              // Момент последнего удаления
	Description          *string                     `json:"description,omitempty"`          // Комментарий
	ExternalCode         *string                     `json:"externalCode,omitempty"`         // Внешний код
	Files                *Files                      `json:"files,omitempty"`                // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group                *Group                      `json:"group,omitempty"`                // Отдел сотрудника
	ID                   *uuid.UUID                  `json:"id,omitempty"`                   // ID сущности
	IncomingDate         *Timestamp                  `json:"incomingDate,omitempty"`         // Входящая дата
	IncomingNumber       *string                     `json:"incomingNumber,omitempty"`       // Входящий номер
	Meta                 *Meta                       `json:"meta,omitempty"`                 // Метаданные
	Moment               *Timestamp                  `json:"moment,omitempty"`               // Дата документа
	Name                 *string                     `json:"name,omitempty"`                 // Наименование
	Organization         *Organization               `json:"organization,omitempty"`         // Метаданные юрлица
	OrganizationAccount  *AgentAccount               `json:"organizationAccount,omitempty"`  // Метаданные счета юрлица
	Owner                *Employee                   `json:"owner,omitempty"`                // Владелец (Сотрудник)
	PayedSum             *Decimal                    `json:"payedSum,omitempty"`             // Сумма входящих платежей по Счету
	PaymentPlannedMoment *Timestamp                  `json:"paymentPlannedMoment,omitempty"` // Планируемая дата оплаты
	Positions            *Positions[InvoicePosition] `json:"positions,omitempty"`            // Метаданные позиций
	Printed              *bool                       `json:"printed,omitempty"`              // Напечатан ли документ
	Project              *Project                    `json:"project,omitempty"`              // Проект
	Published            *bool                       `json:"published,omitempty"`            // Опубликован ли документ
	Rate                 *Rate                       `json:"rate,omitempty"`                 // Валюта
	Shared               *bool                       `json:"shared,omitempty"`               // Общий доступ
	ShippedSum           *Decimal                    `json:"shippedSum,omitempty"`           // Сумма отгруженного
	State                *State                      `json:"state,omitempty"`                // Метаданные статуса
	Store                *Store                      `json:"store,omitempty"`                // Метаданные склада
	Sum                  *Decimal                    `json:"sum,omitempty"`                  // Сумма
	SyncID               *uuid.UUID                  `json:"syncId,omitempty"`               // ID синхронизации. После заполнения недоступен для изменения
	Updated              *Timestamp                  `json:"updated,omitempty"`              // Момент последнего обновления
	VatEnabled           *bool                       `json:"vatEnabled,omitempty"`           // Учитывается ли НДС
	VatIncluded          *bool                       `json:"vatIncluded,omitempty"`          // Включен ли НДС в цену
	VatSum               *Decimal                    `json:"vatSum,omitempty"`               // Сумма включая НДС
	Payments             *Payments                   `json:"payments,omitempty"`             // Массив ссылок на связанные операции в формате Метаданных
	PurchaseOrder        *PurchaseOrder              `json:"purchaseOrder,omitempty"`        // Ссылка на связанный заказ поставщику в формате Метаданных
	Supplies             *Supplies                   `json:"supplies,omitempty"`             // Ссылки на связанные приемки в формате Метаданных
}

func (i InvoiceIn) String() string {
	return Stringify(i)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (i InvoiceIn) GetMeta() *Meta {
	return i.Meta
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
	Create(ctx context.Context, entity *InvoiceIn, params *Params) (*InvoiceIn, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, entities []*InvoiceIn, params *Params) (*[]InvoiceIn, *resty.Response, error)
	DeleteMany(ctx context.Context, entities []*InvoiceIn) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*InvoiceIn, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, entity *InvoiceIn, params *Params) (*InvoiceIn, *resty.Response, error)
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
	DeleteAttributes(ctx context.Context, attributeList []*Attribute) (*DeleteManyResponse, *resty.Response, error)
	GetPublications(ctx context.Context, id *uuid.UUID) (*MetaArray[Publication], *resty.Response, error)
	GetPublicationByID(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (*Publication, *resty.Response, error)
	Publish(ctx context.Context, id *uuid.UUID, template *Templater) (*Publication, *resty.Response, error)
	DeletePublication(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (bool, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*InvoiceIn, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewInvoiceInService(client *Client) InvoiceInService {
	e := NewEndpoint(client, "entity/invoicein")
	return newMainService[InvoiceIn, InvoiceInPosition, MetadataAttributeSharedStates, any](e)
}
