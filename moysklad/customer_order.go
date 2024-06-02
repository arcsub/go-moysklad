package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// CustomerOrder Заказ покупателя.
// Ключевое слово: customerorder
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-zakaz-pokupatelq
type CustomerOrder struct {
	OrganizationAccount   *AgentAccount                     `json:"organizationAccount,omitempty"`
	Project               *Project                          `json:"project,omitempty"`
	AgentAccount          *AgentAccount                     `json:"agentAccount,omitempty"`
	Applicable            *bool                             `json:"applicable,omitempty"`
	Moves                 *Moves                            `json:"moves,omitempty"`
	Code                  *string                           `json:"code,omitempty"`
	Agent                 *Counterparty                     `json:"agent,omitempty"`
	Created               *Timestamp                        `json:"created,omitempty"`
	Deleted               *Timestamp                        `json:"deleted,omitempty"`
	DeliveryPlannedMoment *Timestamp                        `json:"deliveryPlannedMoment,omitempty"`
	Description           *string                           `json:"description,omitempty"`
	ExternalCode          *string                           `json:"externalCode,omitempty"`
	Files                 *Files                            `json:"files,omitempty"`
	Group                 *Group                            `json:"group,omitempty"`
	ID                    *uuid.UUID                        `json:"id,omitempty"`
	InvoicedSum           *Decimal                          `json:"invoicedSum,omitempty"`
	Meta                  *Meta                             `json:"meta,omitempty"`
	Name                  *string                           `json:"name,omitempty"`
	Moment                *Timestamp                        `json:"moment,omitempty"`
	Organization          *Organization                     `json:"organization,omitempty"`
	Printed               *bool                             `json:"printed,omitempty"`
	Owner                 *Employee                         `json:"owner,omitempty"`
	PayedSum              *Decimal                          `json:"payedSum,omitempty"`
	Positions             *Positions[CustomerOrderPosition] `json:"positions,omitempty"`
	AccountID             *uuid.UUID                        `json:"accountId,omitempty"`
	Contract              *Contract                         `json:"contract,omitempty"`
	Published             *bool                             `json:"published,omitempty"`
	Rate                  *Rate                             `json:"rate,omitempty"`
	ReservedSum           *Decimal                          `json:"reservedSum,omitempty"`
	SalesChannel          *SalesChannel                     `json:"salesChannel,omitempty"`
	Shared                *bool                             `json:"shared,omitempty"`
	ShipmentAddress       *string                           `json:"shipmentAddress,omitempty"`
	ShipmentAddressFull   *Address                          `json:"shipmentAddressFull,omitempty"`
	ShippedSum            *Decimal                          `json:"shippedSum,omitempty"`
	State                 *State                            `json:"state,omitempty"`
	Store                 *Store                            `json:"store,omitempty"`
	Sum                   *Decimal                          `json:"sum,omitempty"`
	SyncID                *uuid.UUID                        `json:"syncId,omitempty"`
	Prepayments           *Prepayments                      `json:"prepayments,omitempty"`
	Updated               *Timestamp                        `json:"updated,omitempty"`
	VatEnabled            *bool                             `json:"vatEnabled,omitempty"`
	VatIncluded           *bool                             `json:"vatIncluded,omitempty"`
	VatSum                *Decimal                          `json:"vatSum,omitempty"`
	PurchaseOrders        *PurchaseOrders                   `json:"purchaseOrders,omitempty"`
	Demands               *Demands                          `json:"demands,omitempty"`
	Payments              *Payments                         `json:"payments,omitempty"`
	InvoicesOut           *InvoicesOut                      `json:"invoicesOut,omitempty"`
	TaxSystem             TaxSystem                         `json:"taxSystem,omitempty"`
	Attributes            Attributes                        `json:"attributes,omitempty"`
}

func (c CustomerOrder) String() string {
	return Stringify(c)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (c CustomerOrder) GetMeta() Meta {
	return Deref(c.Meta)
}

func (c CustomerOrder) MetaType() MetaType {
	return MetaTypeCustomerOrder
}

type CustomerOrders = Slice[CustomerOrder]

// CustomerOrderPosition Позиция Заказа покупателя.
// Ключевое слово: customerorderposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-zakaz-pokupatelq-zakazy-pokupatelej-pozicii-zakaza-pokupatelq
type CustomerOrderPosition struct {
	Quantity   *float64            `json:"quantity,omitempty"`
	Assortment *AssortmentPosition `json:"assortment,omitempty"`
	Discount   *Decimal            `json:"discount,omitempty"`
	ID         *uuid.UUID          `json:"id,omitempty"`
	Pack       *Pack               `json:"pack,omitempty"`
	Price      *Decimal            `json:"price,omitempty"`
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`
	Reserve    *Decimal            `json:"reserve,omitempty"`
	Shipped    *Decimal            `json:"shipped,omitempty"`
	Vat        *int                `json:"vat,omitempty"`
	VatEnabled *bool               `json:"vatEnabled,omitempty"`
	Stock      *Stock              `json:"stock,omitempty"`
	TaxSystem  GoodTaxSystem       `json:"taxSystem,omitempty"`
}

func (c CustomerOrderPosition) String() string {
	return Stringify(c)
}

func (c CustomerOrderPosition) MetaType() MetaType {
	return MetaTypeCustomerOrderPosition
}

// CustomerOrderService
// Сервис для работы с заказами покупателя.
type CustomerOrderService interface {
	GetList(ctx context.Context, params *Params) (*List[CustomerOrder], *resty.Response, error)
	Create(ctx context.Context, customerOrder *CustomerOrder, params *Params) (*CustomerOrder, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, customerOrderList []*CustomerOrder, params *Params) (*[]CustomerOrder, *resty.Response, error)
	DeleteMany(ctx context.Context, customerOrderList []*CustomerOrder) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*CustomerOrder, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, customerOrder *CustomerOrder, params *Params) (*CustomerOrder, *resty.Response, error)
	//endpointTemplate[CustomerOrder]
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	GetPositions(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[CustomerOrderPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, params *Params) (*CustomerOrderPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, position *CustomerOrderPosition, params *Params) (*CustomerOrderPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id *uuid.UUID, position *CustomerOrderPosition) (*CustomerOrderPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id *uuid.UUID, positions []*CustomerOrderPosition) (*[]CustomerOrderPosition, *resty.Response, error)
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
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*CustomerOrder, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
	Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetEmbeddedTemplates(ctx context.Context) (*List[EmbeddedTemplate], *resty.Response, error)
	GetEmbeddedTemplateByID(ctx context.Context, id *uuid.UUID) (*EmbeddedTemplate, *resty.Response, error)
	GetCustomTemplates(ctx context.Context) (*List[CustomTemplate], *resty.Response, error)
	GetCustomTemplateByID(ctx context.Context, id *uuid.UUID) (*CustomTemplate, *resty.Response, error)
}

func NewCustomerOrderService(client *Client) CustomerOrderService {
	e := NewEndpoint(client, "entity/customerorder")
	return newMainService[CustomerOrder, CustomerOrderPosition, MetadataAttributeSharedStates, any](e)
}
