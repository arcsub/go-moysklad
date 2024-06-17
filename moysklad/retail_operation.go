package moysklad

import "github.com/goccy/go-json"

type RetailOperation struct {
	Meta Meta `json:"meta"`
	data []byte
}

func (retailOperation RetailOperation) String() string {
	return Stringify(retailOperation.Meta)
}

// MetaType реализует интерфейс MetaTyper
func (retailOperation RetailOperation) MetaType() MetaType {
	return retailOperation.Meta.GetType()
}

// Raw реализует интерфейс RawMetaTyper
func (retailOperation RetailOperation) Raw() []byte {
	return retailOperation.data
}

// UnmarshalJSON реализует интерфейс json.Unmarshaler
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

// AsRetailDemand десериализует объект в тип *RetailDemand
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (retailOperation *RetailOperation) AsRetailDemand() *RetailDemand {
	return UnmarshalAsType[RetailDemand](retailOperation)
}

// AsRetailDrawerCashIn десериализует объект в тип *RetailDrawerCashIn
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (retailOperation *RetailOperation) AsRetailDrawerCashIn() *RetailDrawerCashIn {
	return UnmarshalAsType[RetailDrawerCashIn](retailOperation)
}

// AsRetailDrawerCashOut десериализует объект в тип *RetailDrawerCashOut
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (retailOperation *RetailOperation) AsRetailDrawerCashOut() *RetailDrawerCashOut {
	return UnmarshalAsType[RetailDrawerCashOut](retailOperation)
}

// AsRetailSalesReturn десериализует объект в тип *RetailSalesReturn
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (retailOperation *RetailOperation) AsRetailSalesReturn() *RetailSalesReturn {
	return UnmarshalAsType[RetailSalesReturn](retailOperation)
}

// AsPrepayment десериализует объект в тип *Prepayment
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (retailOperation *RetailOperation) AsPrepayment() *Prepayment {
	return UnmarshalAsType[Prepayment](retailOperation)
}

// AsPrepaymentReturn десериализует объект в тип *PrepaymentReturn
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (retailOperation *RetailOperation) AsPrepaymentReturn() *PrepaymentReturn {
	return UnmarshalAsType[PrepaymentReturn](retailOperation)
}
