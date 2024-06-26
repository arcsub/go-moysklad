package moysklad

// ReceiptTemplate Шаблон печати кассовых чеков.
// TODO: сервис не распознает тип сущности 'receipttemplate'
type ReceiptTemplate struct {
	//AccountID *uuid.UUID `json:"accountId,omitempty"` // ID учётной записи
	//ID        *uuid.UUID `json:"id,omitempty"`        // ID сущности
	//Name   *string `json:"name,omitempty"`   // Наименование
	Meta   *Meta   `json:"meta,omitempty"`   // Метаданные
	Header *string `json:"header,omitempty"` // Верхний блок
	Footer *string `json:"footer,omitempty"` // Нижний блок
}

// String реализует интерфейс [fmt.Stringer].
func (receiptTemplate ReceiptTemplate) String() string {
	return Stringify(receiptTemplate)
}
