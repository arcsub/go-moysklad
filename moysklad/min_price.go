package moysklad

// MinPrice Минимальная цена.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-towar-towary-atributy-wlozhennyh-suschnostej-minimal-naq-cena
type MinPrice struct {
	Value    *Decimal  `json:"value,omitempty"`    // Значение цены
	Currency *Currency `json:"currency,omitempty"` // Ссылка на валюту в формате Метаданных
}

func (m MinPrice) String() string {
	return Stringify(m)
}
