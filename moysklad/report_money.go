package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
)

// Money Остатки денежных средств.
//
// Код сущности: moneyreport
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-den-gi-ostatki-denezhnyh-sredstw
type Money struct {
	Account      MetaNameWrapper `json:"account"`      // Счет организации (не выводится для остатка кассы, так как касса одна на организацию)
	Organization MetaNameWrapper `json:"organization"` // Организация
	Balance      float64         `json:"balance"`      // Текущий остаток денежных средств
}

// MetaType возвращает код сущности.
func (Money) MetaType() MetaType {
	return MetaTypeReportMoney
}

// MoneyPlotSeries Движение денежных средств
//
// Код сущности: moneyplotseries
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-den-gi-dwizhenie-denezhnyh-sredstw
type MoneyPlotSeries struct {
	Context Context             `json:"context"` // Метаданные о выполнившем запрос сотруднике
	Meta    Meta                `json:"meta"`    // Метаданные запроса
	Series  []PlotSeriesElement `json:"series"`  // Массив показателей
	Credit  float64             `json:"credit"`  // Доход
	Debit   float64             `json:"debit"`   // Расход
}

// MetaType возвращает код сущности.
func (MoneyPlotSeries) MetaType() MetaType {
	return MetaTypeReportMoneyPlotSeries
}

// PlotSeriesElement Показатели (series)
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-den-gi-dwizhenie-denezhnyh-sredstw-pokazateli-series
type PlotSeriesElement struct {
	Date    string  `json:"date"`    // Дата
	Credit  float64 `json:"credit"`  // Доход за период
	Debit   float64 `json:"debit"`   // Расход за период
	Balance float64 `json:"balance"` // Баланс (доход-расход)
}

// ReportMoneyService описывает методы сервиса для работы с отчётом Деньги.
type ReportMoneyService interface {
	// GetPlotSeries выполняет запрос на получение графика движения денежных средств.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает движение денежных средств.
	GetPlotSeries(ctx context.Context, params ...*Params) (*MoneyPlotSeries, *resty.Response, error)

	// GetMoney выполняет запрос на получение остатков денежных средств по кассам и счетам.
	// Принимает контекст.
	// Возвращает объект List.
	GetMoney(ctx context.Context) (*List[Money], *resty.Response, error)

	// GetPlotSeriesAsync выполняет запрос на получение графика движения денежных средств (асинхронно).
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает сервис для работы с контекстом асинхронного запроса.
	GetPlotSeriesAsync(ctx context.Context, params ...*Params) (AsyncResultService[MoneyPlotSeries], *resty.Response, error)

	// GetMoneyReportAsync выполняет запрос на получение остатков денежных средств по кассам и счетам.
	// Принимает контекст.
	// Возвращает сервис для работы с контекстом асинхронного запроса.
	GetMoneyReportAsync(ctx context.Context) (AsyncResultService[List[Money]], *resty.Response, error)
}
type reportMoneyService struct {
	Endpoint
}

// NewReportMoneyService принимает [Client] и возвращает сервис для работы с отчётом Деньги.
func NewReportMoneyService(client *Client) ReportMoneyService {
	return &reportMoneyService{NewEndpoint(client, "report/money")}
}

func (service *reportMoneyService) GetPlotSeries(ctx context.Context, params ...*Params) (*MoneyPlotSeries, *resty.Response, error) {
	return NewRequestBuilder[MoneyPlotSeries](service.client, "report/money/plotseries").SetParams(params...).Get(ctx)
}

func (service *reportMoneyService) GetMoney(ctx context.Context) (*List[Money], *resty.Response, error) {
	return NewRequestBuilder[List[Money]](service.client, "report/money/byaccount").Get(ctx)
}

func (service *reportMoneyService) GetPlotSeriesAsync(ctx context.Context, params ...*Params) (AsyncResultService[MoneyPlotSeries], *resty.Response, error) {
	var param *Params
	if len(params) > 0 {
		param = params[0]
	} else {
		param = NewParams()
	}
	return NewRequestBuilder[MoneyPlotSeries](service.client, "report/money/plotseries").SetParams(param.WithAsync()).Async(ctx)
}

func (service *reportMoneyService) GetMoneyReportAsync(ctx context.Context) (AsyncResultService[List[Money]], *resty.Response, error) {
	return NewRequestBuilder[List[Money]](service.client, "report/money/byaccount").SetParams(NewParams().WithAsync()).Async(ctx)
}
