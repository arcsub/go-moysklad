package moysklad

import "encoding/json"

// RetailOperation розничная операция.
type RetailOperation struct {
	Meta *Meta  `json:"meta"` // Метаданные розничной операции
	raw  []byte // сырые данные для последующей конвертации в нужный тип
}

// GetMeta возвращает Метаданные розничной операции.
func (retailOperation RetailOperation) GetMeta() Meta {
	return Deref(retailOperation.Meta)
}

// String реализует интерфейс [fmt.Stringer].
func (retailOperation RetailOperation) String() string {
	return Stringify(retailOperation.Meta)
}

// MetaType возвращает код сущности.
func (retailOperation RetailOperation) MetaType() MetaType {
	return retailOperation.Meta.GetType()
}

// Raw реализует интерфейс [RawMetaTyper].
func (retailOperation RetailOperation) Raw() []byte {
	return retailOperation.raw
}

// UnmarshalJSON реализует интерфейс [json.Unmarshaler].
func (retailOperation *RetailOperation) UnmarshalJSON(data []byte) error {
	type alias RetailOperation
	var t alias

	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	t.raw = data

	*retailOperation = RetailOperation(t)
	return nil
}

// IsRetailDemand возвращает true, если объект является контрагентом [RetailDemand].
func (retailOperation RetailOperation) IsRetailDemand() bool {
	return CheckType(retailOperation, MetaTypeRetailDemand)
}

// IsRetailDrawerCashIn возвращает true, если объект является контрагентом [RetailDrawerCashIn].
func (retailOperation RetailOperation) IsRetailDrawerCashIn() bool {
	return CheckType(retailOperation, MetaTypeRetailDrawerCashIn)
}

// IsRetailDrawerCashOut возвращает true, если объект является контрагентом [RetailDrawerCashOut].
func (retailOperation RetailOperation) IsRetailDrawerCashOut() bool {
	return CheckType(retailOperation, MetaTypeRetailDrawerCashOut)
}

// IsRetailSalesReturn возвращает true, если объект является контрагентом [RetailSalesReturn].
func (retailOperation RetailOperation) IsRetailSalesReturn() bool {
	return CheckType(retailOperation, MetaTypeRetailSalesReturn)
}

// IsPrepayment возвращает true, если объект является контрагентом [Prepayment].
func (retailOperation RetailOperation) IsPrepayment() bool {
	return CheckType(retailOperation, MetaTypePrepayment)
}

// IsPrepaymentReturn возвращает true, если объект является контрагентом [PrepaymentReturn].
func (retailOperation RetailOperation) IsPrepaymentReturn() bool {
	return CheckType(retailOperation, MetaTypePrepaymentReturn)
}

// AsRetailDemand пытается привести объект к типу [RetailDemand].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [RetailDemand] или nil в случае неудачи.
func (retailOperation *RetailOperation) AsRetailDemand() *RetailDemand {
	return UnmarshalAsType[RetailDemand](retailOperation)
}

// AsRetailDrawerCashIn пытается привести объект к типу [RetailDrawerCashIn].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [RetailDrawerCashIn] или nil в случае неудачи.
func (retailOperation *RetailOperation) AsRetailDrawerCashIn() *RetailDrawerCashIn {
	return UnmarshalAsType[RetailDrawerCashIn](retailOperation)
}

// AsRetailDrawerCashOut пытается привести объект к типу [RetailDrawerCashOut].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [RetailDrawerCashOut] или nil в случае неудачи.
func (retailOperation *RetailOperation) AsRetailDrawerCashOut() *RetailDrawerCashOut {
	return UnmarshalAsType[RetailDrawerCashOut](retailOperation)
}

// AsRetailSalesReturn пытается привести объект к типу [RetailSalesReturn].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [RetailSalesReturn] или nil в случае неудачи.
func (retailOperation *RetailOperation) AsRetailSalesReturn() *RetailSalesReturn {
	return UnmarshalAsType[RetailSalesReturn](retailOperation)
}

// AsPrepayment пытается привести объект к типу [Prepayment].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [Prepayment] или nil в случае неудачи.
func (retailOperation *RetailOperation) AsPrepayment() *Prepayment {
	return UnmarshalAsType[Prepayment](retailOperation)
}

// AsPrepaymentReturn пытается привести объект к типу [PrepaymentReturn].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [PrepaymentReturn] или nil в случае неудачи.
func (retailOperation *RetailOperation) AsPrepaymentReturn() *PrepaymentReturn {
	return UnmarshalAsType[PrepaymentReturn](retailOperation)
}
