package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
)

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

// ReportMoneyService
// Сервис для работы с отчётом "Деньги".
type ReportMoneyService interface {
	GetPlotSeries(ctx context.Context, params *Params) (*MoneyPlotSeries, *resty.Response, error)
	GetMoney(ctx context.Context) (*List[Money], *resty.Response, error)
	GetPlotSeriesAsync(ctx context.Context, params *Params) (AsyncResultService[MoneyPlotSeries], *resty.Response, error)
	GetMoneyReportAsync(ctx context.Context) (AsyncResultService[List[Money]], *resty.Response, error)
}
type reportMoneyService struct {
	Endpoint
}

func NewReportMoneyService(client *Client) ReportMoneyService {
	e := NewEndpoint(client, "report/money")
	return &reportMoneyService{e}
}

// GetPlotSeries Запрос на получение графика движения денежных средств.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-den-gi-dwizhenie-denezhnyh-sredstw
func (s *reportMoneyService) GetPlotSeries(ctx context.Context, params *Params) (*MoneyPlotSeries, *resty.Response, error) {
	path := "report/money/plotseries"
	return NewRequestBuilder[MoneyPlotSeries](s.client, path).SetParams(params).Get(ctx)
}

// GetMoney Запрос на получение остатков денежных средств по кассам и счетам.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-den-gi-ostatki-denezhnyh-sredstw
func (s *reportMoneyService) GetMoney(ctx context.Context) (*List[Money], *resty.Response, error) {
	path := "report/money/byaccount"
	return NewRequestBuilder[List[Money]](s.client, path).Get(ctx)
}

// GetPlotSeriesAsync Запрос на получение графика движения денежных средств (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-den-gi-dwizhenie-denezhnyh-sredstw
func (s *reportMoneyService) GetPlotSeriesAsync(ctx context.Context, params *Params) (AsyncResultService[MoneyPlotSeries], *resty.Response, error) {
	path := "report/money/plotseries"
	params.withAsync()
	return NewRequestBuilder[MoneyPlotSeries](s.client, path).SetParams(params).Async(ctx)
}

// GetMoneyReportAsync Запрос на получение остатков денежных средств по кассам и счетам (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-den-gi-ostatki-denezhnyh-sredstw
func (s *reportMoneyService) GetMoneyReportAsync(ctx context.Context) (AsyncResultService[List[Money]], *resty.Response, error) {
	path := "report/money/byaccount"
	params := new(Params).withAsync()
	return NewRequestBuilder[List[Money]](s.client, path).SetParams(params).Async(ctx)
}
