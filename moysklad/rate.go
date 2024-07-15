package moysklad

// Rate Валюта в документах.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-obschie-swedeniq-valuta-w-dokumentah
type Rate struct {
	Currency *Currency `json:"currency,omitempty"` // Метаданные валюты
	Value    *float64  `json:"value,omitempty"`    // Курс валюты в этом документе (содержится в ответе, если значение курса отлично от 1)
}

// GetCurrency возвращает Метаданные валюты.
func (rate Rate) GetCurrency() Currency {
	return Deref(rate.Currency)
}

// GetValue возвращает Курс валюты.
func (rate Rate) GetValue() float64 {
	return Deref(rate.Value)
}

// SetCurrency устанавливает Метаданные валюты.
func (rate *Rate) SetCurrency(currency *Currency) *Rate {
	rate.Currency = currency
	return rate
}

// SetValue устанавливает Курс валюты.
func (rate *Rate) SetValue(value float64) *Rate {
	rate.Value = &value
	return rate
}

// String реализует интерфейс [fmt.Stringer].
func (rate Rate) String() string {
	return Stringify(rate)
}
