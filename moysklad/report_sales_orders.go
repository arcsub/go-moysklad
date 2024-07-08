package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
)

// SeriesElement Показатели (series).
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-prodazh-i-zakazow-pokazateli-series
type SeriesElement struct {
	Date     Timestamp `json:"date"`     // Дата
	Sum      float64   `json:"sum"`      // Количество
	Quantity float64   `json:"quantity"` // Сумма
}

// SalesPlotSeries Показатели продаж.
//
// Код сущности: salesplotseries
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-prodazh-i-zakazow-pokazateli-prodazh
type SalesPlotSeries struct {
	Context Context         `json:"context"` // Метаданные о выполнившем запрос сотруднике
	Meta    Meta            `json:"meta"`    // Метаданные запроса
	Series  []SeriesElement `json:"series"`  // Массив показателей
}

// MetaType возвращает код сущности.
func (SalesPlotSeries) MetaType() MetaType {
	return MetaTypeReportSales
}

// OrdersPlotSeries Показатели заказов.
//
// Код сущности: ordersplotseries
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-prodazh-i-zakazow-pokazateli-zakazow
type OrdersPlotSeries struct {
	Context Context         `json:"context"` // Метаданные о выполнившем запрос сотруднике
	Meta    Meta            `json:"meta"`    // Метаданные запроса
	Series  []SeriesElement `json:"series"`  // Массив показателей
}

// MetaType возвращает код сущности.
func (OrdersPlotSeries) MetaType() MetaType {
	return MetaTypeReportOrders
}

// ReportSalesService описывает методы сервиса для работы с отчётом Показатели продаж.
type ReportSalesService interface {
	// GetPlotSeries выполняет запрос на получение показателей продаж.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает показатели продаж.
	GetPlotSeries(ctx context.Context, params ...*Params) (*SalesPlotSeries, *resty.Response, error)

	// GetPlotSeriesAsync выполняет запрос на получение показателей продаж (асинхронно).
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает сервис для работы с контекстом асинхронного запроса.
	GetPlotSeriesAsync(ctx context.Context, params ...*Params) (AsyncResultService[SalesPlotSeries], *resty.Response, error)
}

type reportSalesService struct {
	Endpoint
}

// NewReportSalesService принимает [Client] и возвращает сервис для работы с отчётом Показатели продаж.
func NewReportSalesService(client *Client) ReportSalesService {
	return &reportSalesService{NewEndpoint(client, "report/sales")}
}

func (service *reportSalesService) GetPlotSeries(ctx context.Context, params ...*Params) (*SalesPlotSeries, *resty.Response, error) {
	return NewRequestBuilder[SalesPlotSeries](service.client, "report/sales/plotseries").SetParams(params...).Get(ctx)
}

func (service *reportSalesService) GetPlotSeriesAsync(ctx context.Context, params ...*Params) (AsyncResultService[SalesPlotSeries], *resty.Response, error) {
	return NewRequestBuilder[SalesPlotSeries](service.client, "report/sales/plotseries").SetParams(params...).Async(ctx)
}

// ReportOrdersService описывает методы сервиса для работы с отчётом Показатели заказов.
type ReportOrdersService interface {
	// GetPlotSeries выполняет запрос на получение показателей заказов.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает показатели заказов.
	GetPlotSeries(ctx context.Context, params ...*Params) (*OrdersPlotSeries, *resty.Response, error)

	// GetPlotSeriesAsync выполняет запрос на получение показателей заказов (асинхронно).
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает сервис для работы с контекстом асинхронного запроса.
	GetPlotSeriesAsync(ctx context.Context, params ...*Params) (AsyncResultService[OrdersPlotSeries], *resty.Response, error)
}

type reportOrdersService struct {
	Endpoint
}

// NewReportOrdersService принимает [Client] и возвращает сервис для работы с отчётом Показатели заказов.
func NewReportOrdersService(client *Client) ReportOrdersService {
	return &reportOrdersService{NewEndpoint(client, "report/orders")}
}

func (service *reportOrdersService) GetPlotSeries(ctx context.Context, params ...*Params) (*OrdersPlotSeries, *resty.Response, error) {
	return NewRequestBuilder[OrdersPlotSeries](service.client, "report/orders/plotseries").SetParams(params...).Get(ctx)
}

func (service *reportOrdersService) GetPlotSeriesAsync(ctx context.Context, params ...*Params) (AsyncResultService[OrdersPlotSeries], *resty.Response, error) {
	return NewRequestBuilder[OrdersPlotSeries](service.client, "report/orders/plotseries").SetParams(params...).Async(ctx)
}
