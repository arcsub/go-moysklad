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
	AccountID               *uuid.UUID                 `json:"accountId,omitempty"`
	Agent                   *Counterparty              `json:"agent,omitempty"`
	AgentAccount            *AgentAccount              `json:"agentAccount,omitempty"`
	Applicable              *bool                      `json:"applicable,omitempty"`
	Code                    *string                    `json:"code,omitempty"`
	Contract                *Contract                  `json:"contract,omitempty"`
	Created                 *Timestamp                 `json:"created,omitempty"`
	Deleted                 *Timestamp                 `json:"deleted,omitempty"`
	Description             *string                    `json:"description,omitempty"`
	ExternalCode            *string                    `json:"externalCode,omitempty"`
	Files                   *Files                     `json:"files,omitempty"`
	Group                   *Group                     `json:"group,omitempty"`
	ID                      *uuid.UUID                 `json:"id,omitempty"`
	Meta                    *Meta                      `json:"meta,omitempty"`
	Moment                  *Timestamp                 `json:"moment,omitempty"`
	Name                    *string                    `json:"name,omitempty"`
	Organization            *Organization              `json:"organization,omitempty"`
	OrganizationAccount     *AgentAccount              `json:"organizationAccount,omitempty"`
	Overhead                *Overhead                  `json:"overhead,omitempty"`
	Owner                   *Employee                  `json:"owner,omitempty"`
	PayedSum                *Decimal                   `json:"payedSum,omitempty"`
	Positions               *Positions[DemandPosition] `json:"positions,omitempty"`
	Printed                 *bool                      `json:"printed,omitempty"`
	Project                 *Project                   `json:"project,omitempty"`
	Published               *bool                      `json:"published,omitempty"`
	Rate                    *Rate                      `json:"rate,omitempty"`
	SalesChannel            *SalesChannel              `json:"salesChannel,omitempty"`
	Shared                  *bool                      `json:"shared,omitempty"`
	ShipmentAddress         *string                    `json:"shipmentAddress,omitempty"`
	ShipmentAddressFull     *Address                   `json:"shipmentAddressFull,omitempty"`
	State                   *State                     `json:"state,omitempty"`
	Store                   *Store                     `json:"store,omitempty"`
	Sum                     *Decimal                   `json:"sum,omitempty"`
	SyncID                  *uuid.UUID                 `json:"syncId,omitempty"`
	Updated                 *Timestamp                 `json:"updated,omitempty"`
	VatEnabled              *bool                      `json:"vatEnabled,omitempty"`
	VatIncluded             *bool                      `json:"vatIncluded,omitempty"`
	VatSum                  *Decimal                   `json:"vatSum,omitempty"`
	CustomerOrder           *CustomerOrder             `json:"customerOrder,omitempty"`
	FactureOut              *FactureOut                `json:"factureOut,omitempty"`
	Returns                 *SalesReturns              `json:"returns,omitempty"`
	Payments                *Payments                  `json:"payments,omitempty"`
	InvoicesOut             *InvoicesOut               `json:"invoicesOut,omitempty"`
	CargoName               *string                    `json:"cargoName,omitempty"`
	Carrier                 *Counterparty              `json:"carrier,omitempty"`
	Consignee               *Counterparty              `json:"consignee,omitempty"`
	GoodPackQuantity        *int                       `json:"goodPackQuantity,omitempty"`
	ShippingInstructions    *string                    `json:"shippingInstructions,omitempty"`
	StateContractId         *string                    `json:"stateContractId,omitempty"`
	TransportFacility       *string                    `json:"transportFacility,omitempty"`
	TransportFacilityNumber *string                    `json:"transportFacilityNumber,omitempty"`
	Attributes              Attributes                 `json:"attributes,omitempty"`
}

func (d Demand) String() string {
	return Stringify(d)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (d Demand) GetMeta() Meta {
	return Deref(d.Meta)
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
