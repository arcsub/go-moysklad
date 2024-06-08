package moysklad

import "encoding/json"

type RetailOperation struct {
	Meta Meta `json:"meta"`
	data json.RawMessage
}

func (retailOperation RetailOperation) String() string {
	return Stringify(retailOperation.Meta)
}

// MetaType удовлетворяет интерфейсу MetaTyper
func (retailOperation RetailOperation) MetaType() MetaType {
	return retailOperation.Meta.GetType()
}

// Raw удовлетворяет интерфейсу RawMetaTyper
func (retailOperation RetailOperation) Raw() json.RawMessage {
	return retailOperation.data
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (retailOperation *RetailOperation) UnmarshalJSON(data []byte) error {
	type alias RetailOperation
	var t alias

	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	t.data = data

	*retailOperation = RetailOperation(t)
	return nil
}

// AsRetailDemand десериализует сырые данные в тип *RetailDemand
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (retailOperation *RetailOperation) AsRetailDemand() *RetailDemand {
	return unmarshalAsType[RetailDemand](retailOperation)
}

// AsRetailDrawerCashIn десериализует сырые данные в тип *RetailDrawerCashIn
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (retailOperation *RetailOperation) AsRetailDrawerCashIn() *RetailDrawerCashIn {
	return unmarshalAsType[RetailDrawerCashIn](retailOperation)
}

// AsRetailDrawerCashOut десериализует сырые данные в тип *RetailDrawerCashOut
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (retailOperation *RetailOperation) AsRetailDrawerCashOut() *RetailDrawerCashOut {
	return unmarshalAsType[RetailDrawerCashOut](retailOperation)
}

// AsRetailSalesReturn десериализует сырые данные в тип *RetailSalesReturn
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (retailOperation *RetailOperation) AsRetailSalesReturn() *RetailSalesReturn {
	return unmarshalAsType[RetailSalesReturn](retailOperation)
}

// AsPrepayment десериализует сырые данные в тип *Prepayment
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (retailOperation *RetailOperation) AsPrepayment() *Prepayment {
	return unmarshalAsType[Prepayment](retailOperation)
}

// AsPrepaymentReturn десериализует сырые данные в тип *PrepaymentReturn
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (retailOperation *RetailOperation) AsPrepaymentReturn() *PrepaymentReturn {
	return unmarshalAsType[PrepaymentReturn](retailOperation)
}

type RetailOperations = Slice[RetailOperation]
