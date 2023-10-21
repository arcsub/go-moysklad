package moysklad

// Money Остатки денежных средств.
// Ключевое слово: moneyreport
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-den-gi-ostatki-denezhnyh-sredstw
type Money struct {
	Account      MetaName `json:"account"`      // Счет организации (не выводится для остатка кассы, так как касса одна на организацию)
	Organization MetaName `json:"organization"` // Организация
	Balance      float64  `json:"balance"`      // Текущий остаток денежных средств
}

func (m Money) MetaType() MetaType {
	return MetaTypeReportMoney
}

// MoneyPlotSeries Движение денежных средств
// Ключевое слово: moneyplotseries
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-den-gi-dwizhenie-denezhnyh-sredstw
type MoneyPlotSeries struct {
	Context Context             `json:"context"`
	Meta    Meta                `json:"meta"`
	Series  []PlotSeriesElement `json:"series"`
	Credit  float64             `json:"credit"`
	Debit   float64             `json:"debit"`
}

func (m MoneyPlotSeries) MetaType() MetaType {
	return MetaTypeReportMoneyPlotSeries
}

// PlotSeriesElement Показатели (series)
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-den-gi-dwizhenie-denezhnyh-sredstw-pokazateli-series
type PlotSeriesElement struct {
	Date    string  `json:"date"`    // Дата
	Credit  float64 `json:"credit"`  // Доход за период
	Debit   float64 `json:"debit"`   // Расход за период
	Balance float64 `json:"balance"` // Баланс (доход-расход)
}
