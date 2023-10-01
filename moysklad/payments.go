package moysklad

import (
	"encoding/json"
)

// PaymentInDocument Входящий платеж, Приходный ордер
type PaymentInDocument struct {
	Meta *Meta `json:"meta"`
	raw  json.RawMessage
}

// UnmarshalJSON анмаршалит Входящий платеж, Приходный ордер, при expand=payments
func (p *PaymentInDocument) UnmarshalJSON(data []byte) (err error) {
	type alias PaymentInDocument
	var t alias

	if err = json.Unmarshal(data, &t); err != nil {
		return err
	}

	t.raw = data
	*p = PaymentInDocument(t)

	return nil
}

// PaymentsIn Входящие платежи, Приходные ордеры
type PaymentsIn []PaymentInDocument

// PaymentOutDocument Исходящий платеж, Расходный ордер
type PaymentOutDocument struct {
	Meta *Meta `json:"meta"`
	raw  json.RawMessage
}

// UnmarshalJSON анмаршалит Исходящий платеж, Расходный ордер, при expand=payments
func (p *PaymentOutDocument) UnmarshalJSON(data []byte) (err error) {
	type alias PaymentOutDocument
	var t alias

	if err = json.Unmarshal(data, &t); err != nil {
		return err
	}

	t.raw = data
	*p = PaymentOutDocument(t)

	return nil
}

// PaymentsOut Исходящие платежи, Расходные ордеры
type PaymentsOut []PaymentOutDocument

// Payments срез всех типов платежей
type Payments struct {
	PaymentsIn
	PaymentsOut
}

//
//func (p PaymentInDocument) Data() json.RawMessage {
//	return p.w.raw
//}
//
//func (p PaymentInDocument) Meta() Meta {
//	return utils.Deref[Meta](p.w.Meta)
//}
//
//func (p PaymentInDocument) CashIn() (CashIn, bool) {
//	return ElementAsType[CashIn](p)
//}
//
//func (p PaymentInDocument) PaymentIn() (PaymentIn, bool) {
//	return ElementAsType[PaymentIn](p)
//}
//
////func (p PaymentsIn) GetCashesIn() Slice[CashIn] {
////	return getElementsByType[CashIn](p)
////}
////
////func (p PaymentsIn) GetPaymentsIn() Slice[PaymentIn] {
////	return getElementsByType[PaymentIn](p)
////}
//
//func (p PaymentOutDocument) Data() json.RawMessage {
//	return p.w.raw
//}
//
//func (p PaymentOutDocument) Meta() Meta {
//	return utils.Deref[Meta](p.w.Meta)
//}
//
//func (p PaymentOutDocument) CashOut() (CashOut, bool) {
//	return ElementAsType[CashOut](p)
//}
//
//func (p PaymentOutDocument) PaymentOut() (PaymentOut, bool) {
//	return ElementAsType[PaymentOut](p)
//}
//
////func (p PaymentsOut) GetCashesOut() Slice[CashOut] {
////	return getElementsByType[CashOut](p)
////}
////
////func (p PaymentsOut) GetPaymentsOut() Slice[PaymentOut] {
////	return getElementsByType[PaymentOut](p)
////}
