package moysklad

import (
	"encoding/json"
)

type OperationsOut []*OperationOut

type OperationOutTypes interface {
	SalesReturn | Supply | InvoiceIn | PurchaseOrder | CommissionReportOut
	HasMeta
}

// OperationOut операция исходящего платежа/расходного ордера
type OperationOut struct {
	operation
}

func (o *OperationOut) UnmarshalJSON(data []byte) (err error) {
	type alias OperationOut
	var t alias

	if err = json.Unmarshal(data, &t); err != nil {
		return err
	}

	t.data = data
	*o = OperationOut(t)

	return nil
}

//func (o OperationOut) SalesReturn() (*SalesReturn, error) {
//	return unmarshalTo[SalesReturn](o)
//}
//
//func (o OperationOut) Supply() (*Supply, error) {
//	return unmarshalTo[Supply](o)
//}
//
//func (o OperationOut) InvoiceIn() (*InvoiceIn, error) {
//	return unmarshalTo[InvoiceIn](o)
//}
//
//func (o OperationOut) PurchaseOrder() (*PurchaseOrder, error) {
//	return unmarshalTo[PurchaseOrder](o)
//}
//
//func (o OperationOut) CommissionReportOut() (*CommissionReportOut, error) {
//	return unmarshalTo[CommissionReportOut](o)
//}
