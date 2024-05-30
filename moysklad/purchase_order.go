package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// PurchaseOrder Заказ поставщику.
// Ключевое слово: purchaseorder
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-zakaz-postawschiku
type PurchaseOrder struct {
	AccountID             *uuid.UUID                        `json:"accountId,omitempty"`             // ID учетной записи
	Agent                 *Counterparty                     `json:"agent,omitempty"`                 // Метаданные контрагента
	AgentAccount          *AgentAccount                     `json:"agentAccount,omitempty"`          // Метаданные счета контрагента
	Applicable            *bool                             `json:"applicable,omitempty"`            // Отметка о проведении
	Attributes            *Attributes                       `json:"attributes,omitempty"`            // Коллекция метаданных доп. полей. Поля объекта
	Code                  *string                           `json:"code,omitempty"`                  // Код
	Contract              *Contract                         `json:"contract,omitempty"`              // Метаданные договора
	Created               *Timestamp                        `json:"created,omitempty"`               // Дата создания
	Deleted               *Timestamp                        `json:"deleted,omitempty"`               // Момент последнего удаления
	DeliveryPlannedMoment *Timestamp                        `json:"deliveryPlannedMoment,omitempty"` // Планируемая дата отгрузки
	Description           *string                           `json:"description,omitempty"`           // Комментарий
	ExternalCode          *string                           `json:"externalCode,omitempty"`          // Внешний код
	Files                 *Files                            `json:"files,omitempty"`                 // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group                 *Group                            `json:"group,omitempty"`                 // Отдел сотрудника
	ID                    *uuid.UUID                        `json:"id,omitempty"`                    // ID сущности
	InvoicedSum           *Decimal                          `json:"invoicedSum,omitempty"`           // Сумма счетов поставщику
	Meta                  *Meta                             `json:"meta,omitempty"`                  // Метаданные
	Moment                *Timestamp                        `json:"moment,omitempty"`                // Дата документа
	Name                  *string                           `json:"name,omitempty"`                  // Наименование
	Organization          *Organization                     `json:"organization,omitempty"`          // Метаданные юрлица
	OrganizationAccount   *AgentAccount                     `json:"organizationAccount,omitempty"`   // Метаданные счета юрлица
	Owner                 *Employee                         `json:"owner,omitempty"`                 // Владелец (Сотрудник)
	PayedSum              *Decimal                          `json:"payedSum,omitempty"`              // Сумма входящих платежей по Счету поставщика
	Positions             *Positions[PurchaseOrderPosition] `json:"positions,omitempty"`             // Метаданные позиций Заказа поставщику
	Printed               *bool                             `json:"printed,omitempty"`               // Напечатан ли документ
	Project               *Project                          `json:"project,omitempty"`               // Проект
	Published             *bool                             `json:"published,omitempty"`             // Опубликован ли документ
	Rate                  *Rate                             `json:"rate,omitempty"`                  // Валюта
	Shared                *bool                             `json:"shared,omitempty"`                // Общий доступ
	ShippedSum            *Decimal                          `json:"shippedSum,omitempty"`            // Сумма отгруженного
	State                 *State                            `json:"state,omitempty"`                 // Метаданные статуса
	Store                 *Store                            `json:"store,omitempty"`                 // Метаданные склада
	Sum                   *Decimal                          `json:"sum,omitempty"`                   // Сумма
	SyncID                *uuid.UUID                        `json:"syncId,omitempty"`                // ID синхронизации. После заполнения недоступен для изменения
	Updated               *Timestamp                        `json:"updated,omitempty"`               // Момент последнего обновления
	VatEnabled            *bool                             `json:"vatEnabled,omitempty"`            // Учитывается ли НДС
	VatIncluded           *bool                             `json:"vatIncluded,omitempty"`           // Включен ли НДС в цену
	VatSum                *Decimal                          `json:"vatSum,omitempty"`                // Сумма включая НДС
	WaitSum               *Decimal                          `json:"waitSum,omitempty"`               // Сумма товаров в пути
	CustomerOrders        *CustomerOrders                   `json:"customerOrders,omitempty"`        // Массив ссылок на связанные заказы покупателей в формате Метаданных
	InvoicesIn            *InvoicesIn                       `json:"invoicesIn,omitempty"`            // Массив ссылок на связанные счета поставщиков в формате Метаданных
	Payments              *Payments                         `json:"payments,omitempty"`              // Массив ссылок на связанные платежи в формате Метаданных
	Supplies              *Supplies                         `json:"supplies,omitempty"`              // Массив ссылок на связанные приемки в формате Метаданных
	InternalOrder         *InternalOrder                    `json:"internalOrder,omitempty"`         // Внутренний заказ, связанный с заказом поставщику, в формате Метаданных
}

func (p PurchaseOrder) String() string {
	return Stringify(p)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (p PurchaseOrder) GetMeta() *Meta {
	return p.Meta
}

func (p PurchaseOrder) MetaType() MetaType {
	return MetaTypePurchaseOrder
}

type PurchaseOrders = Slice[PurchaseOrder]

// PurchaseOrderPosition Позиция Заказа поставщику.
// Ключевое слово: purchaseorderposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-zakaz-postawschiku-zakazy-postawschikam-pozicii-zakaza-postawschiku
type PurchaseOrderPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	Discount   *Decimal            `json:"discount,omitempty"`   // Процент скидки или наценки. Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Price      *Decimal            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Shipped    *Decimal            `json:"shipped,omitempty"`    // Принято
	InTransit  *Decimal            `json:"inTransit,omitempty"`  // Ожидание
	Vat        *float64            `json:"vat,omitempty"`        // НДС, которым облагается текущая позиция
	VatEnabled *bool               `json:"vatEnabled,omitempty"` // Включен ли НДС для позиции. С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
	Wait       *bool               `json:"wait,omitempty"`       // Ожидается данной позиции
	Stock      *Stock              `json:"stock,omitempty"`      // Остатки и себестоимость `?fields=stock&expand=positions`
}

func (p PurchaseOrderPosition) String() string {
	return Stringify(p)
}

func (p PurchaseOrderPosition) MetaType() MetaType {
	return MetaTypePurchaseOrderPosition
}

// PurchaseOrderTemplateArg
// Документ: Заказ поставщику (purchaseorder)
// Основание, на котором он может быть создан:
// - Внутренний заказ (internalorder)
type PurchaseOrderTemplateArg struct {
	InternalOrder *MetaWrapper `json:"internalOrder,omitempty"`
}

// PurchaseOrderService
// Сервис для работы с заказами поставщикам.
type PurchaseOrderService interface {
	GetList(ctx context.Context, params *Params) (*List[PurchaseOrder], *resty.Response, error)
	Create(ctx context.Context, purchaseOrder *PurchaseOrder, params *Params) (*PurchaseOrder, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, purchaseOrderList []*PurchaseOrder, params *Params) (*[]PurchaseOrder, *resty.Response, error)
	DeleteMany(ctx context.Context, purchaseOrderList []*PurchaseOrder) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*PurchaseOrder, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, purchaseOrder *PurchaseOrder, params *Params) (*PurchaseOrder, *resty.Response, error)
	//endpointTemplate[PurchaseOrder]
	//endpointTemplateBasedOn[PurchaseOrder, PurchaseOrderTemplateArg]
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	GetPositions(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[PurchaseOrderPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, params *Params) (*PurchaseOrderPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, position *PurchaseOrderPosition, params *Params) (*PurchaseOrderPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id *uuid.UUID, position *PurchaseOrderPosition) (*PurchaseOrderPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id *uuid.UUID, positions []*PurchaseOrderPosition) (*[]PurchaseOrderPosition, *resty.Response, error)
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
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*PurchaseOrder, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
	Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewPurchaseOrderService(client *Client) PurchaseOrderService {
	e := NewEndpoint(client, "entity/purchaseorder")
	return newMainService[PurchaseOrder, PurchaseOrderPosition, MetadataAttributeSharedStates, any](e)
}
