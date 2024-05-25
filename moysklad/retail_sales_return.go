package moysklad

import (
	"github.com/google/uuid"
)

// RetailSalesReturn Розничный возврат.
// Ключевое слово: retailsalesreturn
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-roznichnyj-wozwrat
type RetailSalesReturn struct {
	OrganizationAccount *AgentAccount              `json:"organizationAccount,omitempty"`
	SyncID              *uuid.UUID                 `json:"syncId,omitempty"`
	AgentAccount        *AgentAccount              `json:"agentAccount,omitempty"`
	Applicable          *bool                      `json:"applicable,omitempty"`
	Attributes          *Attributes                `json:"attributes,omitempty"`
	CashSum             *Decimal                   `json:"cashSum,omitempty"`
	NoCashSum           *Decimal                   `json:"noCashSum,omitempty"`
	Contract            *Contract                  `json:"contract,omitempty"`
	Created             *Timestamp                 `json:"created,omitempty"`
	Deleted             *Timestamp                 `json:"deleted,omitempty"`
	Demand              *RetailDemand              `json:"demand,omitempty"`
	Description         *string                    `json:"description,omitempty"`
	ExternalCode        *string                    `json:"externalCode,omitempty"`
	Group               *Group                     `json:"group,omitempty"`
	ID                  *uuid.UUID                 `json:"id,omitempty"`
	Meta                *Meta                      `json:"meta,omitempty"`
	Moment              *Timestamp                 `json:"moment,omitempty"`
	AccountID           *uuid.UUID                 `json:"accountId,omitempty"`
	Code                *string                    `json:"code,omitempty"`
	Agent               *Counterparty              `json:"agent,omitempty"`
	Name                *string                    `json:"name,omitempty"`
	Owner               *Employee                  `json:"owner,omitempty"`
	Positions           *Positions[RetailPosition] `json:"positions,omitempty"`
	Printed             *bool                      `json:"printed,omitempty"`
	Project             *Project                   `json:"project,omitempty"`
	Published           *bool                      `json:"published,omitempty"`
	QrSum               *Decimal                   `json:"qrSum,omitempty"`
	Rate                *Rate                      `json:"rate,omitempty"`
	RetailShift         *RetailShift               `json:"retailShift,omitempty"`
	RetailStore         *RetailStore               `json:"retailStore,omitempty"`
	Shared              *bool                      `json:"shared,omitempty"`
	State               *State                     `json:"state,omitempty"`
	Store               *Store                     `json:"store,omitempty"`
	Sum                 *Decimal                   `json:"sum,omitempty"`
	Organization        *Organization              `json:"organization,omitempty"`
	VatSum              *Decimal                   `json:"vatSum,omitempty"`
	Updated             *Timestamp                 `json:"updated,omitempty"`
	VatEnabled          *bool                      `json:"vatEnabled,omitempty"`
	VatIncluded         *bool                      `json:"vatIncluded,omitempty"`
	TaxSystem           TaxSystem                  `json:"taxSystem,omitempty"`
}

func (r RetailSalesReturn) String() string {
	return Stringify(r)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (r RetailSalesReturn) GetMeta() *Meta {
	return r.Meta
}

func (r RetailSalesReturn) MetaType() MetaType {
	return MetaTypeRetailSalesReturn
}

// RetailSalesReturnPosition позиция розничного возврата.
// Ключевое слово: salesreturnposition
type RetailSalesReturnPosition struct {
	RetailPosition
}

func (r RetailSalesReturnPosition) MetaType() MetaType {
	return MetaTypeRetailSalesReturnPosition
}
