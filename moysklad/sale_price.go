package moysklad

// SalePrice Цена продажи
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-towar-towary-atributy-wlozhennyh-suschnostej-ceny-prodazhi
type SalePrice struct {
	Value     *float64   `json:"value,omitempty"`     // Значение цены
	Currency  *Currency  `json:"currency,omitempty"`  // Ссылка на валюту в формате Метаданных
	PriceType *PriceType `json:"priceType,omitempty"` // Тип цены
}

func (salePrice SalePrice) GetValue() float64 {
	return Deref(salePrice.Value)
}

func (salePrice SalePrice) GetCurrency() Currency {
	return Deref(salePrice.Currency)
}

func (salePrice SalePrice) GetPriceType() PriceType {
	return Deref(salePrice.PriceType)
}

func (salePrice *SalePrice) SetValue(value float64) *SalePrice {
	salePrice.Value = &value
	return salePrice
}

func (salePrice *SalePrice) SetCurrency(currency *Currency) *SalePrice {
	salePrice.Currency = currency
	return salePrice
}

func (salePrice *SalePrice) SetPriceType(priceType *PriceType) *SalePrice {
	salePrice.PriceType = priceType
	return salePrice
}

func (salePrice SalePrice) String() string {
	return Stringify(salePrice)
}
