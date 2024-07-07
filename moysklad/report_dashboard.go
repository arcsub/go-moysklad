package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
)

// Dashboard Структура объекта показателей.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-struktura-ob-ekta-pokazatelej
type Dashboard struct {
	Money  DashboardMoney       `json:"money"`  // Вложенный объект, представляющий собой информацию о деньгах за указанный период
	Sales  DashboardSalesOrders `json:"sales"`  // Вложенный объект, представляющий собой информацию о продажах за указанный период
	Orders DashboardSalesOrders `json:"orders"` // Вложенный объект, представляющий собой информацию о заказах за указанный период
}

// DashboardMoney Деньги за период.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-struktura-ob-ekta-pokazatelej-den-gi-za-period
type DashboardMoney struct {
	Income        float64 `json:"income"`        // Доходы за период
	Outcome       float64 `json:"outcome"`       // Расходы за период
	Balance       float64 `json:"balance"`       // Текущий баланс
	TodayMovement float64 `json:"todayMovement"` // Дельта за сегодня
	Movement      float64 `json:"movement"`      // Дельта за период
}

// DashboardSalesOrders Продажи/Заказы за период.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-struktura-ob-ekta-pokazatelej-prodazhi-za-period
type DashboardSalesOrders struct {
	Count          float64 `json:"count"`          // Количество продаж/заказов
	Amount         float64 `json:"amount"`         // Прибыль
	MovementAmount float64 `json:"movementAmount"` // Дельта по сравнению с прошлым аналогичным периодом
}

// ReportDashboardService описывает методы сервиса для работы с отчётом показатели.
type ReportDashboardService interface {
	// GetByDay выполняет запрос на получение показателей за день.
	// Принимает контекст.
	// Возвращает объект показателей за день.
	GetByDay(ctx context.Context) (*Dashboard, *resty.Response, error)

	// GetByWeek выполняет запрос на получение показателей за неделю.
	// Принимает контекст.
	// Возвращает объект показателей за неделю.
	GetByWeek(ctx context.Context) (*Dashboard, *resty.Response, error)

	// GetByMonth выполняет запрос на получение показателей за месяц.
	// Принимает контекст.
	// Возвращает объект показателей за месяц.
	GetByMonth(ctx context.Context) (*Dashboard, *resty.Response, error)

	// GetByDayAsync выполняет запрос на получение показателей за день (асинхронно).
	// Принимает контекст.
	// Возвращает сервис для работы с контекстом асинхронного запроса.
	GetByDayAsync(ctx context.Context) (AsyncResultService[Dashboard], *resty.Response, error)

	// GetByWeekAsync выполняет запрос на получение показателей за неделю (асинхронно).
	// Принимает контекст.
	// Возвращает сервис для работы с контекстом асинхронного запроса.
	GetByWeekAsync(ctx context.Context) (AsyncResultService[Dashboard], *resty.Response, error)

	// GetByMonthAsync выполняет запрос на получение показателей за месяц (асинхронно).
	// Принимает контекст.
	// Возвращает сервис для работы с контекстом асинхронного запроса.
	GetByMonthAsync(ctx context.Context) (AsyncResultService[Dashboard], *resty.Response, error)
}

type reportDashboardService struct {
	Endpoint
}

// NewReportDashboardService принимает [Client] и возвращает сервис для работы с отчётом показатели.
func NewReportDashboardService(client *Client) ReportDashboardService {
	return &reportDashboardService{NewEndpoint(client, "report/dashboard")}
}

func (service *reportDashboardService) GetByDay(ctx context.Context) (*Dashboard, *resty.Response, error) {
	return NewRequestBuilder[Dashboard](service.client, "report/dashboard/day").Get(ctx)
}

func (service *reportDashboardService) GetByWeek(ctx context.Context) (*Dashboard, *resty.Response, error) {
	return NewRequestBuilder[Dashboard](service.client, "report/dashboard/week").Get(ctx)
}

func (service *reportDashboardService) GetByMonth(ctx context.Context) (*Dashboard, *resty.Response, error) {
	return NewRequestBuilder[Dashboard](service.client, "report/dashboard/month").Get(ctx)
}

func (service *reportDashboardService) GetByDayAsync(ctx context.Context) (AsyncResultService[Dashboard], *resty.Response, error) {
	return NewRequestBuilder[Dashboard](service.client, "report/dashboard/day").Async(ctx)
}

func (service *reportDashboardService) GetByWeekAsync(ctx context.Context) (AsyncResultService[Dashboard], *resty.Response, error) {
	return NewRequestBuilder[Dashboard](service.client, "report/dashboard/week").Async(ctx)
}

func (service *reportDashboardService) GetByMonthAsync(ctx context.Context) (AsyncResultService[Dashboard], *resty.Response, error) {
	return NewRequestBuilder[Dashboard](service.client, "report/dashboard/month").Async(ctx)
}
