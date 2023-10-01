package moysklad

// Dashboard Структура объекта показателей
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-struktura-ob-ekta-pokazatelej
type Dashboard struct {
	Sales  DashboardSalesOrders // Вложенный объект, представляющий собой информацию о продажах за указанный период
	Orders DashboardSalesOrders // Вложенный объект, представляющий собой информацию о заказах за указанный период
	Money  DashboardMoney       // Вложенный объект, представляющий собой информацию о деньгах за указанный период
}

// DashboardMoney Деньги за период.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-struktura-ob-ekta-pokazatelej-den-gi-za-period
type DashboardMoney struct {
	Income        float64 `json:"income"`        // Доходы за период
	Outcome       float64 `json:"outcome"`       // Расходы за период
	Balance       float64 `json:"balance"`       // Текущий баланс
	TodayMovement float64 `json:"todayMovement"` // Дельта за сегодня
	Movement      float64 `json:"movement"`      // Дельта за период
}

// DashboardSalesOrders Продажи/Заказы за период.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-pokazateli-struktura-ob-ekta-pokazatelej-prodazhi-za-period
type DashboardSalesOrders struct {
	Count          float64 `json:"count"`          // Количество продаж/заказов
	Amount         float64 `json:"amount"`         // Прибыль
	MovementAmount float64 `json:"movementAmount"` // Дельта по сравнению с прошлым аналогичным периодом
}
