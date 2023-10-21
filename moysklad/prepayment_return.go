package moysklad

import (
	"github.com/google/uuid"
)

// PrepaymentReturn Возврат предоплаты.
// Ключевое слово: prepaymentreturn
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vozwrat-predoplaty
type PrepaymentReturn struct {
	Owner        *Employee                            `json:"owner,omitempty"`
	Attributes   *Attributes                          `json:"attributes,omitempty"`
	NoCashSum    *float64                             `json:"noCashSum,omitempty"`
	Organization *Organization                        `json:"organization,omitempty"`
	CashSum      *float64                             `json:"cashSum,omitempty"`
	Code         *string                              `json:"code,omitempty"`
	Created      *Timestamp                           `json:"created,omitempty"`
	Deleted      *Timestamp                           `json:"deleted,omitempty"`
	Description  *string                              `json:"description,omitempty"`
	ExternalCode *string                              `json:"externalCode,omitempty"`
	Files        *Files                               `json:"files,omitempty"`
	Group        *Group                               `json:"group,omitempty"`
	ID           *uuid.UUID                           `json:"id,omitempty"`
	Meta         *Meta                                `json:"meta,omitempty"`
	Moment       *Timestamp                           `json:"moment,omitempty"`
	Name         *string                              `json:"name,omitempty"`
	Applicable   *bool                                `json:"applicable,omitempty"`
	AccountID    *uuid.UUID                           `json:"accountId,omitempty"`
	Agent        *Counterparty                        `json:"agent,omitempty"`
	Positions    *Positions[PrepaymentReturnPosition] `json:"positions,omitempty"`
	Prepayment   *Prepayment                          `json:"prepayment,omitempty"`
	Printed      *bool                                `json:"printed,omitempty"`
	Published    *bool                                `json:"published,omitempty"`
	QRSum        *float64                             `json:"qrSum,omitempty"`
	Rate         *Rate                                `json:"rate,omitempty"`
	RetailShift  *RetailShift                         `json:"retailShift,omitempty"`
	RetailStore  *RetailStore                         `json:"retailStore,omitempty"`
	Shared       *bool                                `json:"shared,omitempty"`
	State        *State                               `json:"state,omitempty"`
	Sum          *float64                             `json:"sum,omitempty"`
	SyncID       *uuid.UUID                           `json:"syncId,omitempty"`
	VatSum       *float64                             `json:"vatSum,omitempty"`
	Updated      *Timestamp                           `json:"updated,omitempty"`
	VatEnabled   *bool                                `json:"vatEnabled,omitempty"`
	VatIncluded  *bool                                `json:"vatIncluded,omitempty"`
	TaxSystem    TaxSystem                            `json:"taxSystem,omitempty"`
}

func (p PrepaymentReturn) String() string {
	return Stringify(p)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (p PrepaymentReturn) GetMeta() *Meta {
	return p.Meta
}

func (p PrepaymentReturn) MetaType() MetaType {
	return MetaTypePrepaymentReturn
}

type PrepaymentReturns = Slice[PrepaymentReturn]

// PrepaymentReturnPosition Позиция Возврата предоплаты.
// Ключевое слово: prepaymentreturnposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vozwrat-predoplaty-atributy-suschnosti-pozicii-vozwrata-predoplaty
type PrepaymentReturnPosition struct {
	PrepaymentPosition
}

func (p PrepaymentReturnPosition) String() string {
	return Stringify(p)
}

func (p PrepaymentReturnPosition) MetaType() MetaType {
	return MetaTypePrepaymentReturnPosition
}
