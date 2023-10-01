package moysklad

import (
	"encoding/json"
	"errors"
)

type operation struct {
	Meta      *Meta    `json:"meta,omitempty"`      // Метаданные
	LinkedSum *float64 `json:"linkedSum,omitempty"` // Оплаченная Сумма
	data      json.RawMessage
}

type OperationTypes interface {
	OperationIn | OperationOut
	HasMeta
	setMeta(*Meta)
	setLinkedSum(*float64)
	setData([]byte)
}

// Operation операция, к которой привязан платёж
// Представляет из себя структуру из полей:
// `Meta` для определения типа сущности
// `LinkedSum` для хранения суммы по операции
// `data` для хранения сырых данных
type Operation struct {
	Meta      *Meta    `json:"meta,omitempty"`      // Метаданные
	LinkedSum *float64 `json:"linkedSum,omitempty"` // Сумма, оплаченная по данному документу
	data      json.RawMessage
}

func (o Operation) String() string {
	return Stringify(o.Meta)
}

func (o *Operation) UnmarshalJSON(data []byte) error {
	type alias Operation
	var t alias
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	t.data = data
	*o = Operation(t)
	return nil
}

// MetaType удовлетворяет интерфейсу MetaTyper
func (o *Operation) MetaType() MetaType {
	return o.Meta.Type
}

// Data удовлетворяет интерфейсу DataMetaTyper
func (o *Operation) Data() json.RawMessage {
	return o.data
}

// CustomerOrder структурирует сущность в *CustomerOrder
// Возвращает ошибку в случае неудачи
func (o *Operation) CustomerOrder() (*CustomerOrder, error) {
	return unmarshalTo[CustomerOrder](o)
}

// PurchaseReturn структурирует сущность в *PurchaseReturn
// Возвращает ошибку в случае неудачи
func (o *Operation) PurchaseReturn() (*PurchaseReturn, error) {
	return unmarshalTo[PurchaseReturn](o)
}

// Demand структурирует сущность в *Demand
// Возвращает ошибку в случае неудачи
func (o *Operation) Demand() (*Demand, error) {
	return unmarshalTo[Demand](o)
}

// InvoiceOut структурирует сущность в *InvoiceOut
// Возвращает ошибку в случае неудачи
func (o *Operation) InvoiceOut() (*InvoiceOut, error) {
	return unmarshalTo[InvoiceOut](o)
}

// CommissionReportIn структурирует сущность в *CommissionReportIn
// Возвращает ошибку в случае неудачи
func (o *Operation) CommissionReportIn() (*CommissionReportIn, error) {
	return unmarshalTo[CommissionReportIn](o)
}

// RetailShift структурирует сущность в *RetailShift
// Возвращает ошибку в случае неудачи
func (o *Operation) RetailShift() (*RetailShift, error) {
	return unmarshalTo[RetailShift](o)
}

// SalesReturn структурирует сущность в *SalesReturn
// Возвращает ошибку в случае неудачи
func (o *Operation) SalesReturn() (*SalesReturn, error) {
	return unmarshalTo[SalesReturn](o)
}

// Supply структурирует сущность в *Supply
// Возвращает ошибку в случае неудачи
func (o *Operation) Supply() (*Supply, error) {
	return unmarshalTo[Supply](o)
}

// InvoiceIn структурирует сущность в *InvoiceIn
// Возвращает ошибку в случае неудачи
func (o *Operation) InvoiceIn() (*InvoiceIn, error) {
	return unmarshalTo[InvoiceIn](o)
}

// PurchaseOrder структурирует сущность в *PurchaseOrder
// Возвращает ошибку в случае неудачи
func (o *Operation) PurchaseOrder() (*PurchaseOrder, error) {
	return unmarshalTo[PurchaseOrder](o)
}

// CommissionReportOut структурирует сущность в *CommissionReportOut
// Возвращает ошибку в случае неудачи
func (o *Operation) CommissionReportOut() (*CommissionReportOut, error) {
	return unmarshalTo[CommissionReportOut](o)
}

// RetailDemand структурирует сущность в *RetailDemand
// Возвращает ошибку в случае неудачи
func (o *Operation) RetailDemand() (*RetailDemand, error) {
	return unmarshalTo[RetailDemand](o)
}

// RetailDrawerCashIn структурирует сущность в *RetailDrawerCashIn
// Возвращает ошибку в случае неудачи
func (o *Operation) RetailDrawerCashIn() (*RetailDrawerCashIn, error) {
	return unmarshalTo[RetailDrawerCashIn](o)
}

// RetailDrawerCashOut структурирует сущность в *RetailDrawerCashOut
// Возвращает ошибку в случае неудачи
func (o *Operation) RetailDrawerCashOut() (*RetailDrawerCashOut, error) {
	return unmarshalTo[RetailDrawerCashOut](o)
}

// RetailSalesReturn структурирует сущность в *RetailSalesReturn
// Возвращает ошибку в случае неудачи
func (o *Operation) RetailSalesReturn() (*RetailSalesReturn, error) {
	return unmarshalTo[RetailSalesReturn](o)
}

// Prepayment структурирует сущность в *Prepayment
// Возвращает ошибку в случае неудачи
func (o *Operation) Prepayment() (*Prepayment, error) {
	return unmarshalTo[Prepayment](o)
}

// PrepaymentReturn структурирует сущность в *PrepaymentReturn
// Возвращает ошибку в случае неудачи
func (o *Operation) PrepaymentReturn() (*PrepaymentReturn, error) {
	return unmarshalTo[PrepaymentReturn](o)
}

// Counterparty структурирует сущность в *Counterparty
// Возвращает ошибку в случае неудачи
func (o *Operation) Counterparty() (*Counterparty, error) {
	return unmarshalTo[Counterparty](o)
}

// Organization структурирует сущность в *Organization
// Возвращает ошибку в случае неудачи
func (o *Operation) Organization() (*Organization, error) {
	return unmarshalTo[Organization](o)
}

// type Operations Slice[Operation]
//
//	func (o Operations) FilterCustomerOrder() Slice[CustomerOrder] {
//		return filterEntity[CustomerOrder](o)
//	}
//
//	func (o Operations) FilterPurchaseReturn() Slice[PurchaseReturn] {
//		return filterEntity[PurchaseReturn](o)
//	}
//
//	func (o Operations) FilterDemand() Slice[Demand] {
//		return filterEntity[Demand](o)
//	}
//
//	func (o Operations) FilterInvoiceOut() Slice[InvoiceOut] {
//		return filterEntity[InvoiceOut](o)
//	}
//
//	func (o Operations) FilterCommissionReportIn() Slice[CommissionReportIn] {
//		return filterEntity[CommissionReportIn](o)
//	}
//
//	func (o Operations) FilterRetailShift() Slice[RetailShift] {
//		return filterEntity[RetailShift](o)
//	}
//
//	func (o Operations) FilterSalesReturn() Slice[SalesReturn] {
//		return filterEntity[SalesReturn](o)
//	}
//
//	func (o Operations) FilterSupply() Slice[Supply] {
//		return filterEntity[Supply](o)
//	}
//
//	func (o Operations) FilterInvoiceIn() Slice[InvoiceIn] {
//		return filterEntity[InvoiceIn](o)
//	}
//
//	func (o Operations) FilterPurchaseOrder() Slice[PurchaseOrder] {
//		return filterEntity[PurchaseOrder](o)
//	}
//
//	func (o Operations) FilterCommissionReportOut() Slice[CommissionReportOut] {
//		return filterEntity[CommissionReportOut](o)
//	}
//
//	func (o Operations) FilterRetailDemand() Slice[RetailDemand] {
//		return filterEntity[RetailDemand](o)
//	}
//
//	func (o Operations) FilterRetailDrawerCashIn() Slice[RetailDrawerCashIn] {
//		return filterEntity[RetailDrawerCashIn](o)
//	}
//
//	func (o Operations) FilterRetailDrawerCashOut() Slice[RetailDrawerCashOut] {
//		return filterEntity[RetailDrawerCashOut](o)
//	}
//
//	func (o Operations) FilterRetailSalesReturn() Slice[RetailSalesReturn] {
//		return filterEntity[RetailSalesReturn](o)
//	}
//
//	func (o Operations) FilterPrepayment() Slice[Prepayment] {
//		return filterEntity[Prepayment](o)
//	}
//
//	func (o Operations) FilterPrepaymentReturn() Slice[PrepaymentReturn] {
//		return filterEntity[PrepaymentReturn](o)
//	}
//
//	func (o Operations) FilterCounterparty() Slice[Counterparty] {
//		return filterEntity[Counterparty](o)
//	}
//
//	func (o Operations) FilterOrganization() Slice[Organization] {
//		return filterEntity[Organization](o)
//	}
func convertToOperation[O OperationTypes, E HasMeta](element E, linkedSum *float64) (*O, error) {
	meta := element.GetMeta()
	if meta == nil {
		return nil, errors.New("meta is nil")
	}
	data, err := json.Marshal(element)
	if err != nil {
		return nil, err
	}
	position := *new(O)
	position.setMeta(meta)
	position.setLinkedSum(linkedSum)
	position.setData(data)
	return &position, nil
}

//func OperationFromEntity[O OperationTypes](entity O, linkedSum *float64) (*Operation, error) {
//	return convertToOperation(entity, linkedSum)
//}
