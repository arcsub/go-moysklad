package moysklad

import "github.com/shopspring/decimal"

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
	Income        decimal.Decimal `json:"income"`        // Доходы за период
	Outcome       decimal.Decimal `json:"outcome"`       // Расходы за период
	Balance       decimal.Decimal `json:"balance"`       // Текущий баланс
	TodayMovement decimal.Decimal `json:"todayMovement"` // Дельта за сегодня
	Movement      decimal.Decimal `json:"movement"`      // Дельта за период
}

// DashboardSalesOrders Продажи/Заказы за период.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-struktura-ob-ekta-pokazatelej-prodazhi-za-period
type DashboardSalesOrders struct {
	Count          float64 `json:"count"`          // Количество продаж/заказов
	Amount         float64 `json:"amount"`         // Прибыль
	MovementAmount float64 `json:"movementAmount"` // Дельта по сравнению с прошлым аналогичным периодом
}
