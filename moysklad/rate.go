package moysklad

// Rate Валюта в документах
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-obschie-swedeniq-valuta-w-dokumentah
type Rate struct {
	Currency *Currency `json:"currency,omitempty"` // Метаданные валюты
	Value    *Decimal  `json:"value,omitempty"`    // Курс валюты в этом документе (содержится в ответе, если значение курса отлично от 1)
}

func (r Rate) String() string {
	return Stringify(r)
}
