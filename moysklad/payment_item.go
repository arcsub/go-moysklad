package moysklad

// PaymentItem Признак предмета расчета.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-towar-towary-atributy-suschnosti-priznak-predmeta-rascheta
type PaymentItem string

const (
	PaymentItemGood                PaymentItem = "GOOD"                  // Товар
	PaymentItemExcisableGood       PaymentItem = "EXCISABLE_GOOD"        // Подакцизный товар
	PaymentItemCompoundPaymentItem PaymentItem = "COMPOUND_PAYMENT_ITEM" // Составной предмет расчета
	PaymentItemAnotherPaymentItem  PaymentItem = "ANOTHER_PAYMENT_ITEM"  // Иной предмет расчета
)
