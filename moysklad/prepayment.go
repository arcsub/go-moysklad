package moysklad

import (
	"github.com/google/uuid"
)

// Prepayment Предоплата.
// Ключевое слово: prepayment
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-predoplata
type Prepayment struct {
	Organization  *Organization                  `json:"organization,omitempty"`
	VatIncluded   *bool                          `json:"vatIncluded,omitempty"`
	Applicable    *bool                          `json:"applicable,omitempty"`
	AccountID     *uuid.UUID                     `json:"accountId,omitempty"`
	CashSum       *float64                       `json:"cashSum,omitempty"`
	Code          *string                        `json:"code,omitempty"`
	Created       *Timestamp                     `json:"created,omitempty"`
	CustomerOrder *CustomerOrder                 `json:"customerOrder,omitempty"`
	Deleted       *Timestamp                     `json:"deleted,omitempty"`
	Description   *string                        `json:"description,omitempty"`
	ExternalCode  *string                        `json:"externalCode,omitempty"`
	Files         *Files                         `json:"files,omitempty"`
	Group         *Group                         `json:"group,omitempty"`
	ID            *uuid.UUID                     `json:"id,omitempty"`
	Meta          *Meta                          `json:"meta,omitempty"`
	Moment        *Timestamp                     `json:"moment,omitempty"`
	Name          *string                        `json:"name,omitempty"`
	NoCashSum     *float64                       `json:"noCashSum,omitempty"`
	Attributes    *Attributes                    `json:"attributes,omitempty"`
	Agent         *Counterparty                  `json:"agent,omitempty"`
	Printed       *bool                          `json:"printed,omitempty"`
	Positions     *Positions[PrepaymentPosition] `json:"positions,omitempty"`
	Published     *bool                          `json:"published,omitempty"`
	QRSum         *float64                       `json:"qrSum,omitempty"`
	Rate          *Rate                          `json:"rate,omitempty"`
	RetailShift   *RetailShift                   `json:"retailShift,omitempty"`
	RetailStore   *RetailStore                   `json:"retailStore,omitempty"`
	Returns       *PrepaymentReturns             `json:"returns,omitempty"`
	Shared        *bool                          `json:"shared,omitempty"`
	State         *State                         `json:"state,omitempty"`
	Sum           *float64                       `json:"sum,omitempty"`
	SyncID        *uuid.UUID                     `json:"syncId,omitempty"`
	VatSum        *float64                       `json:"vatSum,omitempty"`
	Updated       *Timestamp                     `json:"updated,omitempty"`
	VatEnabled    *bool                          `json:"vatEnabled,omitempty"`
	Owner         *Employee                      `json:"owner,omitempty"`
	TaxSystem     TaxSystem                      `json:"taxSystem,omitempty"`
}

func (p Prepayment) String() string {
	return Stringify(p)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (p Prepayment) GetMeta() *Meta {
	return p.Meta
}

func (p Prepayment) MetaType() MetaType {
	return MetaTypePrepayment
}

type Prepayments = Slice[Prepayment]

// PrepaymentPosition Позиция Предоплаты.
// Ключевое слово: prepaymentposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-predoplata-predoplaty-pozicii-predoplaty
type PrepaymentPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	Discount   *float64            `json:"discount,omitempty"`   // Процент скидки или наценки. Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Price      *float64            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Vat        *int                `json:"vat,omitempty"`        // НДС, которым облагается текущая позиция
	VatEnabled *bool               `json:"vatEnabled,omitempty"` // Включен ли НДС для позиции. С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
}

func (p PrepaymentPosition) String() string {
	return Stringify(p)
}

func (p PrepaymentPosition) MetaType() MetaType {
	return MetaTypePrepaymentPosition
}
