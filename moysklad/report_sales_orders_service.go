package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
)

// ReportSalesService
// Сервис для работы с отчётом "Показатели продаж".
type ReportSalesService struct {
	Endpoint
}

func NewReportSalesService(client *Client) *ReportSalesService {
	e := NewEndpoint(client, "report/sales")
	return &ReportSalesService{e}
}

// GetPlotSeries Запрос на получение показателей продаж.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-prodazh-i-zakazow-pokazateli-zakazow
func (s *ReportSalesService) GetPlotSeries(ctx context.Context, params *Params) (*SalesPlotSeries, *resty.Response, error) {
	path := "report/sales/plotseries"
	return NewRequestBuilder[SalesPlotSeries](s.client, path).SetParams(params).Get(ctx)
}

// GetPlotSeriesAsync Запрос на получение показателей продаж (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-prodazh-i-zakazow-pokazateli-zakazow
func (s *ReportSalesService) GetPlotSeriesAsync(ctx context.Context, params *Params) (*AsyncResultService[SalesPlotSeries], *resty.Response, error) {
	path := "report/sales/plotseries"
	return NewRequestBuilder[SalesPlotSeries](s.client, path).SetParams(params).Async(ctx)
}

// ReportOrdersService
// Сервис для работы с отчётом "Показатели заказов".
type ReportOrdersService struct {
	Endpoint
}

func NewReportOrdersService(client *Client) *ReportOrdersService {
	e := NewEndpoint(client, "report/orders")
	return &ReportOrdersService{e}
}

// GetPlotSeries Запрос на получение показателей заказов.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-prodazh-i-zakazow-pokazateli-zakazow
func (s *ReportOrdersService) GetPlotSeries(ctx context.Context, params *Params) (*OrdersPlotSeries, *resty.Response, error) {
	path := "report/orders/plotseries"
	return NewRequestBuilder[OrdersPlotSeries](s.client, path).SetParams(params).Get(ctx)
}

// GetPlotSeriesAsync Запрос на получение показателей заказов (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-prodazh-i-zakazow-pokazateli-zakazow
func (s *ReportOrdersService) GetPlotSeriesAsync(ctx context.Context, params *Params) (*AsyncResultService[OrdersPlotSeries], *resty.Response, error) {
	path := "report/orders/plotseries"
	return NewRequestBuilder[OrdersPlotSeries](s.client, path).SetParams(params).Async(ctx)
}
