package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Demand Отгрузка.
// Ключевое слово: demand
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-otgruzka
type Demand struct {
	AccountID               *uuid.UUID                 `json:"accountId,omitempty"`               // ID учетной записи
	Agent                   *Counterparty              `json:"agent,omitempty"`                   // Метаданные контрагента
	AgentAccount            *AgentAccount              `json:"agentAccount,omitempty"`            // Метаданные счета контрагента
	Applicable              *bool                      `json:"applicable,omitempty"`              // Отметка о проведении
	Attributes              *Attributes                `json:"attributes,omitempty"`              // Коллекция метаданных доп. полей. Поля объекта
	Code                    *string                    `json:"code,omitempty"`                    // Код Отгрузки
	Contract                *Contract                  `json:"contract,omitempty"`                // Метаданные договора
	Created                 *Timestamp                 `json:"created,omitempty"`                 // Дата создания
	Deleted                 *Timestamp                 `json:"deleted,omitempty"`                 // Момент последнего удаления Отгрузки
	Description             *string                    `json:"description,omitempty"`             // Комментарий Отгрузки
	ExternalCode            *string                    `json:"externalCode,omitempty"`            // Внешний код Отгрузки
	Files                   *Files                     `json:"files,omitempty"`                   // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group                   *Group                     `json:"group,omitempty"`                   // Отдел сотрудника
	ID                      *uuid.UUID                 `json:"id,omitempty"`                      // ID сущности
	Meta                    *Meta                      `json:"meta,omitempty"`                    // Метаданные
	Moment                  *Timestamp                 `json:"moment,omitempty"`                  // Дата документа
	Name                    *string                    `json:"name,omitempty"`                    // Наименование
	Organization            *Organization              `json:"organization,omitempty"`            // Метаданные юрлица
	OrganizationAccount     *AgentAccount              `json:"organizationAccount,omitempty"`     // Метаданные счета юрлица
	Overhead                *Overhead                  `json:"overhead,omitempty"`                // Накладные расходы. Если Позиции Отгрузки не заданы, то накладные расходы нельзя задать
	Owner                   *Employee                  `json:"owner,omitempty"`                   // Владелец (Сотрудник)
	PayedSum                *Decimal                   `json:"payedSum,omitempty"`                // Сумма входящих платежей по Отгрузке
	Positions               *Positions[DemandPosition] `json:"positions,omitempty"`               // Метаданные позиций Отгрузки
	Printed                 *bool                      `json:"printed,omitempty"`                 // Напечатан ли документ
	Project                 *Project                   `json:"project,omitempty"`                 // Метаданные проекта
	Published               *bool                      `json:"published,omitempty"`               // Опубликован ли документ
	Rate                    *Rate                      `json:"rate,omitempty"`                    // Валюта
	SalesChannel            *SalesChannel              `json:"salesChannel,omitempty"`            // Метаданные канала продаж
	Shared                  *bool                      `json:"shared,omitempty"`                  // Общий доступ
	ShipmentAddress         *string                    `json:"shipmentAddress,omitempty"`         // Адрес доставки Отгрузки
	ShipmentAddressFull     *Address                   `json:"shipmentAddressFull,omitempty"`     // Адрес доставки Отгрузки с детализацией по отдельным полям
	State                   *State                     `json:"state,omitempty"`                   // Метаданные статуса Отгрузки
	Store                   *Store                     `json:"store,omitempty"`                   // Метаданные склада
	Sum                     *Decimal                   `json:"sum,omitempty"`                     // Сумма
	SyncID                  *uuid.UUID                 `json:"syncId,omitempty"`                  // ID синхронизации. После заполнения недоступен для изменения
	Updated                 *Timestamp                 `json:"updated,omitempty"`                 // Момент последнего обновления
	VatEnabled              *bool                      `json:"vatEnabled,omitempty"`              // Учитывается ли НДС
	VatIncluded             *bool                      `json:"vatIncluded,omitempty"`             // Включен ли НДС в цену
	VatSum                  *Decimal                   `json:"vatSum,omitempty"`                  // Сумма включая НДС
	CustomerOrder           *CustomerOrder             `json:"customerOrder,omitempty"`           // Ссылка на Заказ Покупателя, с которым связана эта Отгрузка в формате Метаданных
	FactureOut              *FactureOut                `json:"factureOut,omitempty"`              // Ссылка на Счет-фактуру выданный, с которым связан этот платеж в формате Метаданных
	Returns                 *SalesReturns              `json:"returns,omitempty"`                 // Массив ссылок на связанные возвраты в формате Метаданных
	Payments                *Payments                  `json:"payments,omitempty"`                // Массив ссылок на связанные платежи в формате Метаданных
	InvoicesOut             *InvoicesOut               `json:"invoicesOut,omitempty"`             // Массив ссылок на связанные счета покупателям в формате Метаданных
	CargoName               *string                    `json:"cargoName,omitempty"`               // Наименование груза
	Carrier                 *Counterparty              `json:"carrier,omitempty"`                 // Метаданные перевозчика (контрагент или юрлицо)
	Consignee               *Counterparty              `json:"consignee,omitempty"`               // Метаданные грузополучателя (контрагент или юрлицо)
	GoodPackQuantity        *int                       `json:"goodPackQuantity,omitempty"`        // Всего мест
	ShippingInstructions    *string                    `json:"shippingInstructions,omitempty"`    // Указания грузоотправителя
	StateContractId         *string                    `json:"stateContractId,omitempty"`         // Идентификатор государственного контракта, договора (соглашения)
	TransportFacility       *string                    `json:"transportFacility,omitempty"`       // Транспортное средство
	TransportFacilityNumber *string                    `json:"transportFacilityNumber,omitempty"` // Номер автомобиля
}

func (d Demand) String() string {
	return Stringify(d)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (d Demand) GetMeta() *Meta {
	return d.Meta
}

func (d Demand) MetaType() MetaType {
	return MetaTypeDemand
}

// type = Demands Slice[Demand] // go.dev/issue/50729

type Demands Slice[Demand]

// DemandPosition Позиция Отгрузки
// Ключевое слово: demandposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-otgruzka-otgruzki-pozicii-otgruzki
type DemandPosition struct {
	AccountID         *uuid.UUID          `json:"accountId,omitempty"`          // ID учетной записи
	Assortment        *AssortmentPosition `json:"assortment,omitempty"`         // Метаданные товара/услуги/серии/модификации/комплекта, которую представляет собой позиция
	Cost              *int                `json:"cost,omitempty"`               // Себестоимость (только для услуг)
	Discount          *Decimal            `json:"discount,omitempty"`           // Процент скидки или наценки. Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%
	ID                *uuid.UUID          `json:"id,omitempty"`                 // ID сущности
	Pack              *Pack               `json:"pack,omitempty"`               // Упаковка Товара
	Price             *Decimal            `json:"price,omitempty"`              // Цена товара/услуги в копейках
	Quantity          *float64            `json:"quantity,omitempty"`           // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Slot              *Slot               `json:"slot,omitempty"`               // Ячейка на складе
	Things            *Things             `json:"things,omitempty"`             // Серийные номера. Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете. В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута
	TrackingCodes     *TrackingCodes      `json:"trackingCodes,omitempty"`      // Коды маркировки товаров и транспортных упаковок
	TrackingCodes1162 *TrackingCodes      `json:"trackingCodes_1162,omitempty"` // Коды маркировки товаров в формате тега 1162
	Overhead          *Decimal            `json:"overhead,omitempty"`           // Накладные расходы. Если Позиции Отгрузки не заданы, то накладные расходы нельзя задать
	Vat               *int                `json:"vat,omitempty"`                // НДС, которым облагается текущая позиция
	VatEnabled        *bool               `json:"vatEnabled,omitempty"`         // Включен ли НДС для позиции. С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
	Stock             *Stock              `json:"stock,omitempty"`              // Остатки и себестоимость `?fields=stock&expand=positions`
}

func (d DemandPosition) String() string {
	return Stringify(d)
}

func (d DemandPosition) MetaType() MetaType {
	return MetaTypeDemandPosition
}

// DemandTemplateArg
// Документ: Отгрузка (demand)
// Основание, на котором он может быть создан:
// - Заказ покупателя (customerorder)
type DemandTemplateArg struct {
	CustomerOrder *MetaWrapper `json:"customerOrder,omitempty"`
}

// DemandService
// Сервис для работы с отгрузками.
type DemandService interface {
	GetList(ctx context.Context, params *Params) (*List[Demand], *resty.Response, error)
	Create(ctx context.Context, entity *Demand, params *Params) (*Demand, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, entities []*Demand, params *Params) (*[]Demand, *resty.Response, error)
	DeleteMany(ctx context.Context, entities []*Demand) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Demand, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, entity *Demand, params *Params) (*Demand, *resty.Response, error)
	//endpointTemplate[Demand]
	//endpointTemplateBasedOn[Demand, DemandTemplateArg]
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	GetPositions(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[DemandPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, params *Params) (*DemandPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, position *DemandPosition, params *Params) (*DemandPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id *uuid.UUID, position *DemandPosition) (*DemandPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id *uuid.UUID, positions []*DemandPosition) (*[]DemandPosition, *resty.Response, error)
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
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*Demand, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
	Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetEmbeddedTemplates(ctx context.Context) (*List[EmbeddedTemplate], *resty.Response, error)
	GetEmbeddedTemplateByID(ctx context.Context, id *uuid.UUID) (*EmbeddedTemplate, *resty.Response, error)
	GetCustomTemplates(ctx context.Context) (*List[CustomTemplate], *resty.Response, error)
	GetCustomTemplateByID(ctx context.Context, id *uuid.UUID) (*CustomTemplate, *resty.Response, error)
	PrintDocument(ctx context.Context, id *uuid.UUID, PrintDocumentArg *PrintDocumentArg) (*PrintFile, *resty.Response, error)
}

func NewDemandService(client *Client) DemandService {
	e := NewEndpoint(client, "entity/demand")
	return newMainService[Demand, DemandPosition, MetadataAttributeSharedStates, any](e)
}
