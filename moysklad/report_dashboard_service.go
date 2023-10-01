package moysklad

import "context"

type ReportDashboardService struct {
	Endpoint
}

func NewReportDashboardService(client *Client) *ReportDashboardService {
	e := NewEndpoint(client, "report/dashboard")
	return &ReportDashboardService{e}
}

// GetByDay Запрос на получение показателей за день.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-poluchit-pokazateli-za-den
func (s *ReportDashboardService) GetByDay(ctx context.Context) (*Dashboard, *Response, error) {
	path := "day"
	return NewRequestBuilder[Dashboard](s.Endpoint, ctx).WithPath(path).Get()
}

// GetByWeek Запрос на получение показателей за неделю.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-poluchit-pokazateli-za-nedelu
func (s *ReportDashboardService) GetByWeek(ctx context.Context) (*Dashboard, *Response, error) {
	path := "week"
	return NewRequestBuilder[Dashboard](s.Endpoint, ctx).WithPath(path).Get()
}

// GetByMonth Запрос на получение показателей за месяц.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-poluchit-pokazateli-za-mesqc
func (s *ReportDashboardService) GetByMonth(ctx context.Context) (*Dashboard, *Response, error) {
	path := "month"
	return NewRequestBuilder[Dashboard](s.Endpoint, ctx).WithPath(path).Get()
}

// GetByDayAsync Запрос на получение показателей за день (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-poluchit-pokazateli-za-den
func (s *ReportDashboardService) GetByDayAsync(ctx context.Context) (*AsyncResultService[Dashboard], *Response, error) {
	path := "day"
	return NewRequestBuilder[Dashboard](s.Endpoint, ctx).WithPath(path).Async()
}

// GetByWeekAsync Запрос на получение показателей за неделю (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-poluchit-pokazateli-za-den
func (s *ReportDashboardService) GetByWeekAsync(ctx context.Context) (*AsyncResultService[Dashboard], *Response, error) {
	path := "week"
	return NewRequestBuilder[Dashboard](s.Endpoint, ctx).WithPath(path).Async()
}

// GetByMonthAsync Запрос на получение показателей за месяц (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-poluchit-pokazateli-za-mesqc
func (s *ReportDashboardService) GetByMonthAsync(ctx context.Context) (*AsyncResultService[Dashboard], *Response, error) {
	path := "month"
	return NewRequestBuilder[Dashboard](s.Endpoint, ctx).WithPath(path).Async()
}
