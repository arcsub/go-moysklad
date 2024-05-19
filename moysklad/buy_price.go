package moysklad

import "github.com/shopspring/decimal"

// BuyPrice Закупочная цена.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-towar-towary-atributy-wlozhennyh-suschnostej-zakupochnaq-cena
type BuyPrice struct {
	Value    *decimal.Decimal `json:"value,omitempty"`    // Значение цены
	Currency *Currency        `json:"currency,omitempty"` // Ссылка на валюту в формате Метаданных
}

func (b BuyPrice) String() string {
	return Stringify(b)
}
