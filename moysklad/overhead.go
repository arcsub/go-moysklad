package moysklad

// Overhead Накладные расходы.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-priemka-priemki-nakladnye-rashody
type Overhead struct {
	Sum          *float64     `json:"sum,omitempty"`          // Сумма в копейках
	Distribution Distribution `json:"distribution,omitempty"` // Распределение накладных расходов
}

// GetSum возвращает Сумму в копейках.
func (overhead Overhead) GetSum() float64 {
	return Deref(overhead.Sum)
}

// GetDistribution возвращает Распределение накладных расходов.
func (overhead Overhead) GetDistribution() Distribution {
	return overhead.Distribution
}

// SetSum устанавливает Сумму в копейках.
func (overhead *Overhead) SetSum(sum float64) *Overhead {
	overhead.Sum = &sum
	return overhead
}

// SetDistribution устанавливает Распределение накладных расходов.
func (overhead *Overhead) SetDistribution(distribution Distribution) *Overhead {
	overhead.Distribution = distribution
	return overhead
}

// SetDistributionWeight устанавливает Распределение накладных расходов по весу.
func (overhead *Overhead) SetDistributionWeight() *Overhead {
	overhead.Distribution = DistributionWeight
	return overhead
}

// SetDistributionVolume устанавливает Распределение накладных расходов по объему.
func (overhead *Overhead) SetDistributionVolume() *Overhead {
	overhead.Distribution = DistributionVolume
	return overhead
}

// SetDistributionPrice устанавливает Распределение накладных расходов по цене.
func (overhead *Overhead) SetDistributionPrice() *Overhead {
	overhead.Distribution = DistributionPrice
	return overhead
}

// String реализует интерфейс [fmt.Stringer].
func (overhead Overhead) String() string {
	return Stringify(overhead)
}

// Distribution Тип Распределения накладных расходов.
//
// Возможные значения:
//   - DistributionWeight – по весу
//   - DistributionVolume – по объему
//   - DistributionPrice  – по цене
type Distribution string

const (
	DistributionWeight Distribution = "weight" // по весу
	DistributionVolume Distribution = "volume" // по объему
	DistributionPrice  Distribution = "price"  // по цене
)
