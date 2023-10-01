package moysklad

// SeriesElement Показатели (series).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-prodazh-i-zakazow-pokazateli-series
type SeriesElement struct {
	Date     Timestamp `json:"date"`     // Дата
	Quantity float64   `json:"quantity"` // Количество
	Sum      float64   `json:"sum"`      // Сумма
}

// SalesPlotSeries Показатели продаж.
// Ключевое слово: salesplotseries
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-prodazh-i-zakazow-pokazateli-prodazh
type SalesPlotSeries struct {
	Context Context         `json:"context"`
	Meta    Meta            `json:"meta"`
	Series  []SeriesElement `json:"series"` // Массив показателей
}

func (s SalesPlotSeries) MetaType() MetaType {
	return MetaTypeReportSales
}

// OrdersPlotSeries Показатели заказов.
// Ключевое слово: ordersplotseries
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-prodazh-i-zakazow-pokazateli-zakazow
type OrdersPlotSeries struct {
	Context Context         `json:"context"`
	Meta    Meta            `json:"meta"`
	Series  []SeriesElement `json:"series"` // Массив показателей
}

func (o OrdersPlotSeries) MetaType() MetaType {
	return MetaTypeReportOrders
}
