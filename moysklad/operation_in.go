package moysklad

import (
	"encoding/json"
)

type OperationsIn []*OperationIn

type OperationInType interface {
	HasMeta
	ConvertToOperation(*float64) (*OperationIn, error)
}

func (o *OperationsIn) Bind(element OperationInType, linkedSum *float64) error {
	op, err := element.ConvertToOperation(linkedSum)
	if err != nil {
		return err
	}

	*o = append(*o, op)
	return nil
}

// OperationIn операция входящего платежа/приходного ордера
type OperationIn struct {
	operation
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (o OperationIn) GetMeta() *Meta {
	return o.Meta
}

func (o OperationIn) setMeta(meta *Meta) {
	*o.Meta = *meta
}

func (o OperationIn) setLinkedSum(linkedSum *float64) {
	*o.LinkedSum = *linkedSum
}

func (o OperationIn) setData(data []byte) {
	o.data = data
}

func (o *OperationIn) UnmarshalJSON(data []byte) (err error) {
	type alias OperationIn
	var t alias

	if err = json.Unmarshal(data, &t); err != nil {
		return err
	}

	t.data = data
	*o = OperationIn(t)

	return nil
}

//func (o OperationIn) CustomerOrder() (*CustomerOrder, error) {
//	return unmarshalTo[CustomerOrder](o)
//}
//
//func (o OperationIn) PurchaseReturn() (*PurchaseReturn, error) {
//	return unmarshalTo[PurchaseReturn](o)
//}
//
//func (o OperationIn) Demand() (*Demand, error) {
//	return unmarshalTo[Demand](o)
//}
//
//func (o OperationIn) InvoiceOut() (*InvoiceOut, error) {
//	return unmarshalTo[InvoiceOut](o)
//}
//
//func (o OperationIn) RetailShift() (*RetailShift, error) {
//	return unmarshalTo[RetailShift](o)
//}
//
//func (o OperationIn) CommissionReportIn() (*CommissionReportIn, error) {
//	return unmarshalTo[CommissionReportIn](o)
//}
