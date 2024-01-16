package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
)

type ReportDashboardService struct {
	Endpoint
}

func NewReportDashboardService(client *Client) *ReportDashboardService {
	e := NewEndpoint(client, "report/dashboard")
	return &ReportDashboardService{e}
}

// GetByDay Запрос на получение показателей за день.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-poluchit-pokazateli-za-den
func (s *ReportDashboardService) GetByDay(ctx context.Context) (*Dashboard, *resty.Response, error) {
	path := "report/dashboard/day"
	return NewRequestBuilder[Dashboard](s.client, path).Get(ctx)
}

// GetByWeek Запрос на получение показателей за неделю.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-poluchit-pokazateli-za-nedelu
func (s *ReportDashboardService) GetByWeek(ctx context.Context) (*Dashboard, *resty.Response, error) {
	path := "report/dashboard/week"
	return NewRequestBuilder[Dashboard](s.client, path).Get(ctx)
}

// GetByMonth Запрос на получение показателей за месяц.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-poluchit-pokazateli-za-mesqc
func (s *ReportDashboardService) GetByMonth(ctx context.Context) (*Dashboard, *resty.Response, error) {
	path := "report/dashboard/month"
	return NewRequestBuilder[Dashboard](s.client, path).Get(ctx)
}

// GetByDayAsync Запрос на получение показателей за день (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-poluchit-pokazateli-za-den
func (s *ReportDashboardService) GetByDayAsync(ctx context.Context) (*AsyncResultService[Dashboard], *resty.Response, error) {
	path := "report/dashboard/day"
	return NewRequestBuilder[Dashboard](s.client, path).Async(ctx)
}

// GetByWeekAsync Запрос на получение показателей за неделю (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-poluchit-pokazateli-za-den
func (s *ReportDashboardService) GetByWeekAsync(ctx context.Context) (*AsyncResultService[Dashboard], *resty.Response, error) {
	path := "report/dashboard/week"
	return NewRequestBuilder[Dashboard](s.client, path).Async(ctx)
}

// GetByMonthAsync Запрос на получение показателей за месяц (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-poluchit-pokazateli-za-mesqc
func (s *ReportDashboardService) GetByMonthAsync(ctx context.Context) (*AsyncResultService[Dashboard], *resty.Response, error) {
	path := "report/dashboard/month"
	return NewRequestBuilder[Dashboard](s.client, path).Async(ctx)
}
