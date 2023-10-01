package moysklad

//
//import (
//	"encoding/json"
//	"github.com/arcsub/moysklad/utils"
//)
//
//// TODO
//
//type OperationTask struct {
//	wrapper[operationTask]
//	OperationIn
//	OperationOut
//	OperationRetail
//}
//
//type operationTask struct {
//	Meta *Meta `json:"meta,omitempty"` // Метаданные
//	raw  json.RawMessage
//}
//
//func (o *OperationTask) UnmarshalJSON(data []byte) (err error) {
//	type alias OperationTask
//	var t alias
//
//	if err = json.Unmarshal(data, &t); err != nil {
//		return err
//	}
//
//	t.w.raw = data
//	*o = OperationTask(t)
//
//	return nil
//}
//
//func (o OperationTask) Data() json.RawMessage {
//	return o.w.raw
//}
//
//func (o OperationTask) Meta() Meta {
//	return utils.Deref[Meta](o.w.Meta)
//}
//
//func (o OperationTask) Counterparty() (Counterparty, bool) {
//	return ElementAsType[Counterparty](o)
//}
//
//func (o OperationTask) Organization() (Organization, bool) {
//	return ElementAsType[Organization](o)
//}
