package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
)

// SeriesElement Показатели (series).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-prodazh-i-zakazow-pokazateli-series
type SeriesElement struct {
	Date     Timestamp `json:"date"`
	Sum      Decimal   `json:"sum"`
	Quantity float64   `json:"quantity"`
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

// ReportSalesService
// Сервис для работы с отчётом "Показатели продаж".
type ReportSalesService interface {
	GetPlotSeries(ctx context.Context, params *Params) (*SalesPlotSeries, *resty.Response, error)
	GetPlotSeriesAsync(ctx context.Context, params *Params) (AsyncResultService[SalesPlotSeries], *resty.Response, error)
}

type reportSalesService struct {
	Endpoint
}

func NewReportSalesService(client *Client) ReportSalesService {
	e := NewEndpoint(client, "report/sales")
	return &reportSalesService{e}
}

// GetPlotSeries Запрос на получение показателей продаж.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-prodazh-i-zakazow-pokazateli-zakazow
func (s *reportSalesService) GetPlotSeries(ctx context.Context, params *Params) (*SalesPlotSeries, *resty.Response, error) {
	path := "report/sales/plotseries"
	return NewRequestBuilder[SalesPlotSeries](s.client, path).SetParams(params).Get(ctx)
}

// GetPlotSeriesAsync Запрос на получение показателей продаж (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-prodazh-i-zakazow-pokazateli-zakazow
func (s *reportSalesService) GetPlotSeriesAsync(ctx context.Context, params *Params) (AsyncResultService[SalesPlotSeries], *resty.Response, error) {
	path := "report/sales/plotseries"
	return NewRequestBuilder[SalesPlotSeries](s.client, path).SetParams(params).Async(ctx)
}

// ReportOrdersService
// Сервис для работы с отчётом "Показатели заказов".
type ReportOrdersService interface {
	GetPlotSeries(ctx context.Context, params *Params) (*OrdersPlotSeries, *resty.Response, error)
	GetPlotSeriesAsync(ctx context.Context, params *Params) (AsyncResultService[OrdersPlotSeries], *resty.Response, error)
}

type reportOrdersService struct {
	Endpoint
}

func NewReportOrdersService(client *Client) ReportOrdersService {
	e := NewEndpoint(client, "report/orders")
	return &reportOrdersService{e}
}

// GetPlotSeries Запрос на получение показателей заказов.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-prodazh-i-zakazow-pokazateli-zakazow
func (s *reportOrdersService) GetPlotSeries(ctx context.Context, params *Params) (*OrdersPlotSeries, *resty.Response, error) {
	path := "report/orders/plotseries"
	return NewRequestBuilder[OrdersPlotSeries](s.client, path).SetParams(params).Get(ctx)
}

// GetPlotSeriesAsync Запрос на получение показателей заказов (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-prodazh-i-zakazow-pokazateli-zakazow
func (s *reportOrdersService) GetPlotSeriesAsync(ctx context.Context, params *Params) (AsyncResultService[OrdersPlotSeries], *resty.Response, error) {
	path := "report/orders/plotseries"
	return NewRequestBuilder[OrdersPlotSeries](s.client, path).SetParams(params).Async(ctx)
}
