package moysklad

// MinPrice Минимальная цена.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-towar-towary-atributy-wlozhennyh-suschnostej-minimal-naq-cena
type MinPrice struct {
	Value    *float64  `json:"value,omitempty"`    // Значение цены
	Currency *Currency `json:"currency,omitempty"` // Ссылка на валюту в формате Метаданных
}

func (minPrice MinPrice) GetValue() float64 {
	return Deref(minPrice.Value)
}

func (minPrice MinPrice) GetCurrency() Currency {
	return Deref(minPrice.Currency)
}

func (minPrice *MinPrice) SetValue(value float64) *MinPrice {
	minPrice.Value = &value
	return minPrice
}

func (minPrice *MinPrice) SetCurrency(currency *Currency) *MinPrice {
	minPrice.Currency = currency
	return minPrice
}

func (minPrice MinPrice) String() string {
	return Stringify(minPrice)
}
