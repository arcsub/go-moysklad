package moysklad

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// CustomerOrder Заказ покупателя.
// Ключевое слово: customerorder
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-zakaz-pokupatelq
type CustomerOrder struct {
	Printed               *bool                             `json:"printed,omitempty"`
	Contract              *Contract                         `json:"contract,omitempty"`
	AgentAccount          *AgentAccount                     `json:"agentAccount,omitempty"`
	Applicable            *bool                             `json:"applicable,omitempty"`
	Attributes            *Attributes                       `json:"attributes,omitempty"`
	Code                  *string                           `json:"code,omitempty"`
	Project               *Project                          `json:"project,omitempty"`
	Created               *Timestamp                        `json:"created,omitempty"`
	Deleted               *Timestamp                        `json:"deleted,omitempty"`
	DeliveryPlannedMoment *Timestamp                        `json:"deliveryPlannedMoment,omitempty"`
	Description           *string                           `json:"description,omitempty"`
	Published             *bool                             `json:"published,omitempty"`
	Files                 *Files                            `json:"files,omitempty"`
	Group                 *Group                            `json:"group,omitempty"`
	ID                    *uuid.UUID                        `json:"id,omitempty"`
	InvoicedSum           *decimal.Decimal                  `json:"invoicedSum,omitempty"`
	Meta                  *Meta                             `json:"meta,omitempty"`
	Name                  *string                           `json:"name,omitempty"`
	Moment                *Timestamp                        `json:"moment,omitempty"`
	Organization          *Organization                     `json:"organization,omitempty"`
	AccountID             *uuid.UUID                        `json:"accountId,omitempty"`
	Owner                 *Employee                         `json:"owner,omitempty"`
	PayedSum              *decimal.Decimal                  `json:"payedSum,omitempty"`
	Positions             *Positions[CustomerOrderPosition] `json:"positions,omitempty"`
	OrganizationAccount   *AgentAccount                     `json:"organizationAccount,omitempty"`
	Agent                 *Counterparty                     `json:"agent,omitempty"`
	ExternalCode          *string                           `json:"externalCode,omitempty"`
	Rate                  *Rate                             `json:"rate,omitempty"`
	ReservedSum           *decimal.Decimal                  `json:"reservedSum,omitempty"`
	SalesChannel          *SalesChannel                     `json:"salesChannel,omitempty"`
	Shared                *bool                             `json:"shared,omitempty"`
	ShipmentAddress       *string                           `json:"shipmentAddress,omitempty"`
	ShipmentAddressFull   *Address                          `json:"shipmentAddressFull,omitempty"`
	ShippedSum            *decimal.Decimal                  `json:"shippedSum,omitempty"`
	State                 *State                            `json:"state,omitempty"`
	Store                 *Store                            `json:"store,omitempty"`
	Sum                   *decimal.Decimal                  `json:"sum,omitempty"`
	SyncID                *uuid.UUID                        `json:"syncId,omitempty"`
	Prepayments           *Prepayments                      `json:"prepayments,omitempty"`
	Updated               *Timestamp                        `json:"updated,omitempty"`
	VatEnabled            *bool                             `json:"vatEnabled,omitempty"`
	VatIncluded           *bool                             `json:"vatIncluded,omitempty"`
	VatSum                *decimal.Decimal                  `json:"vatSum,omitempty"`
	PurchaseOrders        *PurchaseOrders                   `json:"purchaseOrders,omitempty"`
	Demands               *Demands                          `json:"demands,omitempty"`
	Payments              *Payments                         `json:"payments,omitempty"`
	InvoicesOut           *InvoicesOut                      `json:"invoicesOut,omitempty"`
	Moves                 *Moves                            `json:"moves,omitempty"`
	TaxSystem             TaxSystem                         `json:"taxSystem,omitempty"`
}

func (c CustomerOrder) String() string {
	return Stringify(c)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (c CustomerOrder) GetMeta() *Meta {
	return c.Meta
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
	Discount   *float64            `json:"discount,omitempty"`
	ID         *uuid.UUID          `json:"id,omitempty"`
	Pack       *Pack               `json:"pack,omitempty"`
	Price      *decimal.Decimal    `json:"price,omitempty"`
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`
	Reserve    *decimal.Decimal    `json:"reserve,omitempty"`
	Shipped    *decimal.Decimal    `json:"shipped,omitempty"`
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
