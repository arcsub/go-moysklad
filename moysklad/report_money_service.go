package moysklad

import (
	"context"
)

// ReportMoneyService
// Сервис для работы с отчётом "Деньги".
type ReportMoneyService struct {
	Endpoint
}

func NewReportMoneyService(client *Client) *ReportMoneyService {
	e := NewEndpoint(client, "report/money")
	return &ReportMoneyService{e}
}

// GetPlotSeries Запрос на получение графика движения денежных средств.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-den-gi-dwizhenie-denezhnyh-sredstw
func (s *ReportMoneyService) GetPlotSeries(ctx context.Context, params *Params) (*MoneyPlotSeries, *Response, error) {
	path := "plotseries"
	return NewRequestBuilder[MoneyPlotSeries](s.Endpoint, ctx).WithPath(path).WithParams(params).Get()
}

// GetMoney Запрос на получение остатков денежных средств по кассам и счетам.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-den-gi-ostatki-denezhnyh-sredstw
func (s *ReportMoneyService) GetMoney(ctx context.Context) (*List[Money], *Response, error) {
	path := "byaccount"
	return NewRequestBuilder[List[Money]](s.Endpoint, ctx).WithPath(path).Get()
}

// GetPlotSeriesAsync Запрос на получение графика движения денежных средств (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-den-gi-dwizhenie-denezhnyh-sredstw
func (s *ReportMoneyService) GetPlotSeriesAsync(ctx context.Context, params *Params) (*AsyncResultService[MoneyPlotSeries], *Response, error) {
	path := "plotseries"
	params.withAsync()
	return NewRequestBuilder[MoneyPlotSeries](s.Endpoint, ctx).WithPath(path).WithParams(params).Async()
}

// GetMoneyReportAsync Запрос на получение остатков денежных средств по кассам и счетам (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-den-gi-ostatki-denezhnyh-sredstw
func (s *ReportMoneyService) GetMoneyReportAsync(ctx context.Context) (*AsyncResultService[List[Money]], *Response, error) {
	path := "byaccount"
	params := new(Params)
	params.withAsync()
	return NewRequestBuilder[List[Money]](s.Endpoint, ctx).WithPath(path).WithParams(params).Async()
}
