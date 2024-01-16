package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
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
func (s *ReportMoneyService) GetPlotSeries(ctx context.Context, params *Params) (*MoneyPlotSeries, *resty.Response, error) {
	path := "report/money/plotseries"
	return NewRequestBuilder[MoneyPlotSeries](s.client, path).SetParams(params).Get(ctx)
}

// GetMoney Запрос на получение остатков денежных средств по кассам и счетам.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-den-gi-ostatki-denezhnyh-sredstw
func (s *ReportMoneyService) GetMoney(ctx context.Context) (*List[Money], *resty.Response, error) {
	path := "report/money/byaccount"
	return NewRequestBuilder[List[Money]](s.client, path).Get(ctx)
}

// GetPlotSeriesAsync Запрос на получение графика движения денежных средств (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-den-gi-dwizhenie-denezhnyh-sredstw
func (s *ReportMoneyService) GetPlotSeriesAsync(ctx context.Context, params *Params) (*AsyncResultService[MoneyPlotSeries], *resty.Response, error) {
	path := "report/money/plotseries"
	params.withAsync()
	return NewRequestBuilder[MoneyPlotSeries](s.client, path).SetParams(params).Async(ctx)
}

// GetMoneyReportAsync Запрос на получение остатков денежных средств по кассам и счетам (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-den-gi-ostatki-denezhnyh-sredstw
func (s *ReportMoneyService) GetMoneyReportAsync(ctx context.Context) (*AsyncResultService[List[Money]], *resty.Response, error) {
	path := "report/money/byaccount"
	params := new(Params)
	params.withAsync()
	return NewRequestBuilder[List[Money]](s.client, path).SetParams(params).Async(ctx)
}
