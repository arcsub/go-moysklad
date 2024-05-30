package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
)

// Dashboard Структура объекта показателей
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-struktura-ob-ekta-pokazatelej
type Dashboard struct {
	Money  DashboardMoney
	Sales  DashboardSalesOrders
	Orders DashboardSalesOrders
}

// DashboardMoney Деньги за период.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-struktura-ob-ekta-pokazatelej-den-gi-za-period
type DashboardMoney struct {
	Income        Decimal `json:"income"`        // Доходы за период
	Outcome       Decimal `json:"outcome"`       // Расходы за период
	Balance       Decimal `json:"balance"`       // Текущий баланс
	TodayMovement Decimal `json:"todayMovement"` // Дельта за сегодня
	Movement      Decimal `json:"movement"`      // Дельта за период
}

// DashboardSalesOrders Продажи/Заказы за период.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-struktura-ob-ekta-pokazatelej-prodazhi-za-period
type DashboardSalesOrders struct {
	Count          float64 `json:"count"`          // Количество продаж/заказов
	Amount         float64 `json:"amount"`         // Прибыль
	MovementAmount float64 `json:"movementAmount"` // Дельта по сравнению с прошлым аналогичным периодом
}

type ReportDashboardService interface {
	GetByDay(ctx context.Context) (*Dashboard, *resty.Response, error)
	GetByWeek(ctx context.Context) (*Dashboard, *resty.Response, error)
	GetByMonth(ctx context.Context) (*Dashboard, *resty.Response, error)
	GetByDayAsync(ctx context.Context) (AsyncResultService[Dashboard], *resty.Response, error)
	GetByWeekAsync(ctx context.Context) (AsyncResultService[Dashboard], *resty.Response, error)
	GetByMonthAsync(ctx context.Context) (AsyncResultService[Dashboard], *resty.Response, error)
}

type reportDashboardService struct {
	Endpoint
}

func NewReportDashboardService(client *Client) ReportDashboardService {
	e := NewEndpoint(client, "report/dashboard")
	return &reportDashboardService{e}
}

// GetByDay Запрос на получение показателей за день.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-poluchit-pokazateli-za-den
func (s *reportDashboardService) GetByDay(ctx context.Context) (*Dashboard, *resty.Response, error) {
	path := "report/dashboard/day"
	return NewRequestBuilder[Dashboard](s.client, path).Get(ctx)
}

// GetByWeek Запрос на получение показателей за неделю.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-poluchit-pokazateli-za-nedelu
func (s *reportDashboardService) GetByWeek(ctx context.Context) (*Dashboard, *resty.Response, error) {
	path := "report/dashboard/week"
	return NewRequestBuilder[Dashboard](s.client, path).Get(ctx)
}

// GetByMonth Запрос на получение показателей за месяц.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-poluchit-pokazateli-za-mesqc
func (s *reportDashboardService) GetByMonth(ctx context.Context) (*Dashboard, *resty.Response, error) {
	path := "report/dashboard/month"
	return NewRequestBuilder[Dashboard](s.client, path).Get(ctx)
}

// GetByDayAsync Запрос на получение показателей за день (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-poluchit-pokazateli-za-den
func (s *reportDashboardService) GetByDayAsync(ctx context.Context) (AsyncResultService[Dashboard], *resty.Response, error) {
	path := "report/dashboard/day"
	return NewRequestBuilder[Dashboard](s.client, path).Async(ctx)
}

// GetByWeekAsync Запрос на получение показателей за неделю (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-poluchit-pokazateli-za-den
func (s *reportDashboardService) GetByWeekAsync(ctx context.Context) (AsyncResultService[Dashboard], *resty.Response, error) {
	path := "report/dashboard/week"
	return NewRequestBuilder[Dashboard](s.client, path).Async(ctx)
}

// GetByMonthAsync Запрос на получение показателей за месяц (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-poluchit-pokazateli-za-mesqc
func (s *reportDashboardService) GetByMonthAsync(ctx context.Context) (AsyncResultService[Dashboard], *resty.Response, error) {
	path := "report/dashboard/month"
	return NewRequestBuilder[Dashboard](s.client, path).Async(ctx)
}
