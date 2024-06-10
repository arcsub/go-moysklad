package moysklad

import (
	"github.com/google/uuid"
)

// InvoicePosition общие поля для счетов.
// Ключевое слово: invoicein
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-schet-postawschika-scheta-postawschikow-pozicii-scheta-postawschika
type InvoicePosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	Discount   *float64            `json:"discount,omitempty"`   // Процент скидки или наценки. Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID сущности
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Meta       *Meta               `json:"meta,omitempty"`       // Метаданные
	Price      *float64            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Vat        *int                `json:"vat,omitempty"`        // НДС, которым облагается текущая позиция
	VatEnabled *bool               `json:"vatEnabled,omitempty"` // Включен ли НДС для позиции. С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
	Stock      *Stock              `json:"stock,omitempty"`      // Остатки и себестоимость `?fields=stock&expand=positions`
}

func (invoicePosition InvoicePosition) GetAccountID() uuid.UUID {
	return Deref(invoicePosition.AccountID)
}

func (invoicePosition InvoicePosition) GetAssortment() AssortmentPosition {
	return Deref(invoicePosition.Assortment)
}

func (invoicePosition InvoicePosition) GetDiscount() float64 {
	return Deref(invoicePosition.Discount)
}

func (invoicePosition InvoicePosition) GetID() uuid.UUID {
	return Deref(invoicePosition.ID)
}

func (invoicePosition InvoicePosition) GetPack() Pack {
	return Deref(invoicePosition.Pack)
}

func (invoicePosition InvoicePosition) GetMeta() Meta {
	return Deref(invoicePosition.Meta)
}

func (invoicePosition InvoicePosition) GetPrice() float64 {
	return Deref(invoicePosition.Price)
}

func (invoicePosition InvoicePosition) GetQuantity() float64 {
	return Deref(invoicePosition.Quantity)
}

func (invoicePosition InvoicePosition) GetVat() int {
	return Deref(invoicePosition.Vat)
}

func (invoicePosition InvoicePosition) GetVatEnabled() bool {
	return Deref(invoicePosition.VatEnabled)
}

func (invoicePosition InvoicePosition) GetStock() Stock {
	return Deref(invoicePosition.Stock)
}

func (invoicePosition *InvoicePosition) SetAssortment(assortment AsAssortment) *InvoicePosition {
	invoicePosition.Assortment = assortment.AsAssortment()
	return invoicePosition
}

func (invoicePosition *InvoicePosition) SetDiscount(discount float64) *InvoicePosition {
	invoicePosition.Discount = &discount
	return invoicePosition
}

func (invoicePosition *InvoicePosition) SetPack(pack *Pack) *InvoicePosition {
	invoicePosition.Pack = pack
	return invoicePosition
}

func (invoicePosition *InvoicePosition) SetMeta(meta *Meta) *InvoicePosition {
	invoicePosition.Meta = meta
	return invoicePosition
}

func (invoicePosition *InvoicePosition) SetPrice(price float64) *InvoicePosition {
	invoicePosition.Price = &price
	return invoicePosition
}

func (invoicePosition *InvoicePosition) SetQuantity(quantity float64) *InvoicePosition {
	invoicePosition.Quantity = &quantity
	return invoicePosition
}

func (invoicePosition *InvoicePosition) SetVat(vat int) *InvoicePosition {
	invoicePosition.Vat = &vat
	return invoicePosition
}

func (invoicePosition *InvoicePosition) SetVatEnabled(vatEnabled bool) *InvoicePosition {
	invoicePosition.VatEnabled = &vatEnabled
	return invoicePosition
}

func (invoicePosition InvoicePosition) String() string {
	return Stringify(invoicePosition)
}

func (invoicePosition InvoicePosition) MetaType() MetaType {
	return MetaTypeInvoicePosition
}
