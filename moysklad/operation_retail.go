package moysklad

import "encoding/json"

type RetailOperation struct {
	Meta Meta `json:"meta"`
	data json.RawMessage
}

func (o RetailOperation) String() string {
	return Stringify(o.Meta)
}

// MetaType удовлетворяет интерфейсу MetaTyper
func (o RetailOperation) MetaType() MetaType {
	return o.Meta.Type
}

// Raw удовлетворяет интерфейсу RawMetaTyper
func (o RetailOperation) Raw() json.RawMessage {
	return o.data
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (o *RetailOperation) UnmarshalJSON(data []byte) error {
	type alias RetailOperation
	var t alias

	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	t.data = data

	*o = RetailOperation(t)
	return nil
}

// RetailDemand десериализует сырые данные в тип *RetailDemand
func (o *RetailOperation) RetailDemand() *RetailDemand {
	return unmarshalAsType[RetailDemand](o)
}

// RetailDrawerCashIn десериализует сырые данные в тип *RetailDrawerCashIn
func (o *RetailOperation) RetailDrawerCashIn() *RetailDrawerCashIn {
	return unmarshalAsType[RetailDrawerCashIn](o)
}

// RetailDrawerCashOut десериализует сырые данные в тип *RetailDrawerCashOut
func (o *RetailOperation) RetailDrawerCashOut() *RetailDrawerCashOut {
	return unmarshalAsType[RetailDrawerCashOut](o)
}

// RetailSalesReturn десериализует сырые данные в тип *RetailSalesReturn
func (o *RetailOperation) RetailSalesReturn() *RetailSalesReturn {
	return unmarshalAsType[RetailSalesReturn](o)
}

// Prepayment десериализует сырые данные в тип *Prepayment
func (o *RetailOperation) Prepayment() *Prepayment {
	return unmarshalAsType[Prepayment](o)
}

// PrepaymentReturn десериализует сырые данные в тип *PrepaymentReturn
func (o *RetailOperation) PrepaymentReturn() *PrepaymentReturn {
	return unmarshalAsType[PrepaymentReturn](o)
}

type RetailOperations []RetailOperation
