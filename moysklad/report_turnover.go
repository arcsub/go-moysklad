package moysklad

import "github.com/shopspring/decimal"

// Turnover Атрибуты объекта отчета.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-oboroty-oboroty-po-towaram-atributy-ob-ekta-otcheta
type Turnover struct {
	OnPeriodStart TurnoverIncomeOutcome `json:"onPeriodStart"` // Показатели на начало периода
	OnPeriodEnd   TurnoverIncomeOutcome `json:"onPeriodEnd"`   // Показатели на конец периода
	Income        TurnoverIncomeOutcome `json:"income"`        // Показатели прихода в течение периода отчета
	Outcome       TurnoverIncomeOutcome `json:"outcome"`       // Показатели расхода в течение периода отчета
}

// TurnoverIncomeOutcome Структура объекта показатели (onPeriodStart, onPeriodEnd, income, outcome).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-oboroty-oboroty-po-towaram-struktura-ob-ekta-pokazateli-onperiodstart-onperiodend-income-outcome
type TurnoverIncomeOutcome struct {
	Sum      decimal.Decimal `json:"sum"`      // Сумма себестоимости
	Quantity float64         `json:"quantity"` // Количество единиц товара
}

// TurnoverAll Обороты по товарам.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-oboroty-poluchit-oboroty-po-towaram
type TurnoverAll struct {
	Assortment MetaName `json:"assortment"` // Краткое представление Товара или Модификации в отчете
	Turnover
}

// TurnoverAssortment Структура объекта assortment.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-oboroty-oboroty-po-towaru-s-detalizaciej-po-skladam-struktura-ob-ekta-assortment
type TurnoverAssortment struct {
	Article       string        `json:"article"`       // Артикул Товара
	Code          string        `json:"code"`          // Код Товара
	Image         Image         `json:"image"`         // Первое изображение Товара или Модификации
	Meta          Meta          `json:"meta"`          // Метаданные Товара или Модификации
	Name          string        `json:"name"`          // Наименование Товара или Модификации
	ProductFolder ProductFolder `json:"productFolder"` // Группа Товара или Модификации
	Uom           MetaName      `json:"uom"`           // Единица измерения
}

// TurnoverByOperation Обороты по товару с детализацией по документам.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-oboroty-oboroty-po-towaru-s-detalizaciej-po-dokumentam
type TurnoverByOperation struct {
	Assortment TurnoverAssortment `json:"assortment"`
	Operation  TurnoverOperation  `json:"operation"`
	Store      MetaName           `json:"store"`
	Cost       decimal.Decimal    `json:"cost"`
	Sum        decimal.Decimal    `json:"sum"`
	Quantity   float64            `json:"quantity"`
}

// TurnoverOperation Структура объекта operation.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-oboroty-oboroty-po-towaru-s-detalizaciej-po-dokumentam-struktura-ob-ekta-operation
type TurnoverOperation struct {
	Meta        Meta      `json:"meta"`        // Метаданные документа
	Name        string    `json:"name"`        // Номер документа
	Description string    `json:"description"` // Комментарий к документу
	Moment      Timestamp `json:"moment"`      // Дата проведения документа
	Agent       MetaName  `json:"agent"`       // Контрагент документа
}

// TurnoverByStore Обороты по товару с детализацией по складам.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-oboroty-oboroty-po-towaru-s-detalizaciej-po-skladam
type TurnoverByStore struct {
	Assortment   TurnoverAssortment       `json:"assortment"`   // Краткое представление Товара или Модификации в отчете
	StockByStore []TurnoverByStoreElement `json:"stockByStore"` // Детализация оборотов по складам
}

// TurnoverByStoreElement Структура объекта детализация оборотов по складам.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-oboroty-oboroty-po-towaru-s-detalizaciej-po-skladam-struktura-ob-ekta-detalizaciq-oborotow-po-skladam
type TurnoverByStoreElement struct {
	Store MetaName `json:"store"` // Склад
	Turnover
}
