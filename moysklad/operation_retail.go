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

// Data удовлетворяет интерфейсу DataMetaTyper
func (o RetailOperation) Data() json.RawMessage {
	return o.data
}

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

// RetailDemand структурирует сущность в *RetailDemand
func (o *RetailOperation) RetailDemand() (*RetailDemand, error) {
	return unmarshalTo[RetailDemand](o)
}

// RetailDrawerCashIn структурирует сущность в *RetailDrawerCashIn
func (o *RetailOperation) RetailDrawerCashIn() (*RetailDrawerCashIn, error) {
	return unmarshalTo[RetailDrawerCashIn](o)
}

// RetailDrawerCashOut структурирует сущность в *RetailDrawerCashOut
func (o *RetailOperation) RetailDrawerCashOut() (*RetailDrawerCashOut, error) {
	return unmarshalTo[RetailDrawerCashOut](o)
}

// RetailSalesReturn структурирует сущность в *RetailSalesReturn
func (o *RetailOperation) RetailSalesReturn() (*RetailSalesReturn, error) {
	return unmarshalTo[RetailSalesReturn](o)
}

// Prepayment структурирует сущность в *Prepayment
func (o *RetailOperation) Prepayment() (*Prepayment, error) {
	return unmarshalTo[Prepayment](o)
}

// PrepaymentReturn структурирует сущность в *PrepaymentReturn
func (o *RetailOperation) PrepaymentReturn() (*PrepaymentReturn, error) {
	return unmarshalTo[PrepaymentReturn](o)
}

type RetailOperations []RetailOperation
