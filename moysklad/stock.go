package moysklad

import "github.com/shopspring/decimal"

// Stock Остатки и себестоимость в позициях документов
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-obschie-swedeniq-ostatki-i-sebestoimost-w-poziciqh-dokumentow
type Stock struct {
	Cost      decimal.Decimal `json:"cost"`
	Quantity  float64         `json:"quantity"`
	Reserve   float64         `json:"reserve"`
	InTransit float64         `json:"intransit"`
	Available float64         `json:"available"`
}

func (s Stock) String() string {
	return Stringify(s)
}
