package moysklad

//
//import (
//	"encoding/json"
//	"github.com/arcsub/moysklad/utils"
//)
//
//// OperationRetail розничная операция
//type OperationRetail struct {
//	wrapper[operationRetail]
//}
//
//type operationRetail struct {
//	Meta      *Meta    `json:"meta,omitempty"`      // Метаданные
//	LinkedSum *float64 `json:"linkedSum,omitempty"` // Сумма, оплаченная по данному документу из этого платежа
//	raw       json.RawMessage
//}
//
//func (o *OperationRetail) UnmarshalJSON(data []byte) (err error) {
//	type alias OperationRetail
//	var t alias
//
//	if err = json.Unmarshal(data, &t); err != nil {
//		return err
//	}
//
//	t.w.raw = data
//	*o = OperationRetail(t)
//
//	return nil
//}
//
//func (o OperationRetail) Data() json.RawMessage {
//	return o.w.raw
//}
//
//func (o OperationRetail) LinkedSum() float64 {
//	return utils.Deref[float64](o.w.LinkedSum)
//}
//
//func (o OperationRetail) Meta() Meta {
//	return utils.Deref[Meta](o.w.Meta)
//}
//
//func (o OperationRetail) RetailDemand() (RetailDemand, bool) {
//	return ElementAsType[RetailDemand](o)
//}
//
//func (o OperationRetail) RetailDrawerCashIn() (RetailDrawerCashIn, bool) {
//	return ElementAsType[RetailDrawerCashIn](o)
//}
//
//func (o OperationRetail) RetailDrawerCashOut() (RetailDrawerCashOut, bool) {
//	return ElementAsType[RetailDrawerCashOut](o)
//}
//
//func (o OperationRetail) RetailSalesReturn() (RetailSalesReturn, bool) {
//	return ElementAsType[RetailSalesReturn](o)
//}
//
//func (o OperationRetail) Prepayment() (Prepayment, bool) {
//	return ElementAsType[Prepayment](o)
//}
//
//func (o OperationRetail) PrepaymentReturn() (PrepaymentReturn, bool) {
//	return ElementAsType[PrepaymentReturn](o)
//}
