package moysklad

// BuyPrice Закупочная цена.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-towar-towary-atributy-wlozhennyh-suschnostej-zakupochnaq-cena
type BuyPrice struct {
	Value    *float64  `json:"value,omitempty"`    // Значение цены
	Currency *Currency `json:"currency,omitempty"` // Ссылка на валюту в формате Метаданных
}

func (buyPrice BuyPrice) GetValue() float64 {
	return Deref(buyPrice.Value)
}

func (buyPrice BuyPrice) GetCurrency() Currency {
	return Deref(buyPrice.Currency)
}

func (buyPrice *BuyPrice) SetValue(value *float64) *BuyPrice {
	buyPrice.Value = value
	return buyPrice
}

func (buyPrice *BuyPrice) SetCurrency(currency *Currency) *BuyPrice {
	buyPrice.Currency = currency
	return buyPrice
}

func (buyPrice BuyPrice) String() string {
	return Stringify(buyPrice)
}
