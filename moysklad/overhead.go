package moysklad

import "github.com/shopspring/decimal"

// Overhead Накладные расходы.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-priemka-priemki-nakladnye-rashody
type Overhead struct {
	Sum          *decimal.Decimal `json:"sum,omitempty"`          // Сумма в копейках
	Distribution Distribution     `json:"distribution,omitempty"` // Распределение накладных расходов
}

func (o Overhead) String() string {
	return Stringify(o)
}

// Distribution Тип Распределения накладных расходов.
type Distribution string

const (
	DistributionWeight Distribution = "weight" // по весу
	DistributionVolume Distribution = "volume" // по объему
	DistributionPrice  Distribution = "price"  // по цене
)
