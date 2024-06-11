package moysklad

// Overhead Накладные расходы.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-priemka-priemki-nakladnye-rashody
type Overhead struct {
	Sum          *float64     `json:"sum,omitempty"`          // Сумма в копейках
	Distribution Distribution `json:"distribution,omitempty"` // Распределение накладных расходов
}

func (overhead Overhead) GetSum() float64 {
	return Deref(overhead.Sum)
}

func (overhead Overhead) GetDistribution() Distribution {
	return overhead.Distribution
}

func (overhead *Overhead) SetSum(sum float64) *Overhead {
	overhead.Sum = &sum
	return overhead
}

func (overhead *Overhead) SetDistribution(distribution Distribution) *Overhead {
	overhead.Distribution = distribution
	return overhead
}

func (overhead Overhead) String() string {
	return Stringify(overhead)
}

// Distribution Тип Распределения накладных расходов.
type Distribution string

const (
	DistributionWeight Distribution = "weight" // по весу
	DistributionVolume Distribution = "volume" // по объему
	DistributionPrice  Distribution = "price"  // по цене
)
