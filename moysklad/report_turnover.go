package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
)

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
	Sum      float64 `json:"sum"`      // Сумма себестоимости
	Quantity float64 `json:"quantity"` // Количество единиц товара
}

// TurnoverAll Обороты по товарам.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-oboroty-poluchit-oboroty-po-towaram
type TurnoverAll struct {
	Assortment MetaNameWrapper `json:"assortment"` // Краткое представление Товара или Модификации в отчете
	Turnover
}

// TurnoverAssortment Структура объекта assortment.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-oboroty-oboroty-po-towaru-s-detalizaciej-po-skladam-struktura-ob-ekta-assortment
type TurnoverAssortment struct {
	Article       string          `json:"article"`       // Артикул Товара
	Code          string          `json:"code"`          // Код Товара
	Image         Image           `json:"image"`         // Первое изображение Товара или Модификации
	Meta          Meta            `json:"meta"`          // Метаданные Товара или Модификации
	Name          string          `json:"name"`          // Наименование Товара или Модификации
	ProductFolder ProductFolder   `json:"productFolder"` // Группа Товара или Модификации
	Uom           MetaNameWrapper `json:"uom"`           // Единица измерения
}

// TurnoverByOperation Обороты по товару с детализацией по документам.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-oboroty-oboroty-po-towaru-s-detalizaciej-po-dokumentam
type TurnoverByOperation struct {
	Assortment TurnoverAssortment `json:"assortment"`
	Operation  TurnoverOperation  `json:"operation"`
	Store      MetaNameWrapper    `json:"store"`
	Cost       float64            `json:"cost"`
	Sum        float64            `json:"sum"`
	Quantity   float64            `json:"quantity"`
}

// TurnoverOperation Структура объекта operation.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-oboroty-oboroty-po-towaru-s-detalizaciej-po-dokumentam-struktura-ob-ekta-operation
type TurnoverOperation struct {
	Meta        Meta            `json:"meta"`        // Метаданные документа
	Name        string          `json:"name"`        // Номер документа
	Description string          `json:"description"` // Комментарий к документу
	Moment      Timestamp       `json:"moment"`      // Дата проведения документа
	Agent       MetaNameWrapper `json:"agent"`       // Контрагент документа
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
	Store MetaNameWrapper `json:"store"` // Склад
	Turnover
}

// ReportTurnoverService
// Сервис для работы с отчётом "Обороты".
type ReportTurnoverService interface {
	GetAll(ctx context.Context) (*List[TurnoverAll], *resty.Response, error)
	GetByStoreWithProduct(ctx context.Context, product *Product) (*List[TurnoverByOperation], *resty.Response, error)
	GetByStoreWithVariant(ctx context.Context, variant *Variant) (*List[TurnoverByOperation], *resty.Response, error)
	GetByOperationsWithProduct(ctx context.Context, product *Product) (*List[TurnoverByOperation], *resty.Response, error)
	GetByOperationsWithVariant(ctx context.Context, variant *Variant) (*List[TurnoverByOperation], *resty.Response, error)
}

type reportTurnoverService struct {
	Endpoint
}

func NewReportTurnoverService(client *Client) ReportTurnoverService {
	e := NewEndpoint(client, "report/turnover")
	return &reportTurnoverService{e}
}

// GetAll Запрос на получение отчета "Обороты по товарам".
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-oboroty-poluchit-oboroty-po-towaram
func (service *reportTurnoverService) GetAll(ctx context.Context) (*List[TurnoverAll], *resty.Response, error) {
	path := fmt.Sprintf("%s/all", service.uri)
	return NewRequestBuilder[List[TurnoverAll]](service.client, path).Get(ctx)
}

// GetByStoreWithProduct Отчет обороты по товару и его модификациям с детализацией по складам.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-oboroty-oboroty-po-towaru-s-detalizaciej-po-skladam
func (service *reportTurnoverService) GetByStoreWithProduct(ctx context.Context, product *Product) (*List[TurnoverByOperation], *resty.Response, error) {
	path := fmt.Sprintf("%s/bystore", service.uri)
	params := NewParams().WithFilterEquals("product", *product.Meta.Href)
	return NewRequestBuilder[List[TurnoverByOperation]](service.client, path).SetParams(params).Get(ctx)
}

// GetByStoreWithVariant Отчет обороты по модификации с детализацией по складам.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-oboroty-oboroty-po-towaru-s-detalizaciej-po-skladam
func (service *reportTurnoverService) GetByStoreWithVariant(ctx context.Context, variant *Variant) (*List[TurnoverByOperation], *resty.Response, error) {
	path := fmt.Sprintf("%s/bystore", service.uri)
	params := NewParams().WithFilterEquals("variant", *variant.Meta.Href)
	return NewRequestBuilder[List[TurnoverByOperation]](service.client, path).SetParams(params).Get(ctx)
}

// GetByOperationsWithProduct Запрос на получение отчета Обороты по товару с детализацией по документам.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-oboroty-poluchit-oboroty-po-towaru-s-detalizaciej-po-dokumentam
func (service *reportTurnoverService) GetByOperationsWithProduct(ctx context.Context, product *Product) (*List[TurnoverByOperation], *resty.Response, error) {
	path := fmt.Sprintf("%s/byoperations", service.uri)
	params := NewParams().WithFilterEquals("product", *product.Meta.Href)
	return NewRequestBuilder[List[TurnoverByOperation]](service.client, path).SetParams(params).Get(ctx)
}

// GetByOperationsWithVariant Запрос на получение отчета Обороты по модификации с детализацией по документам.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-oboroty-poluchit-oboroty-po-towaru-s-detalizaciej-po-dokumentam
func (service *reportTurnoverService) GetByOperationsWithVariant(ctx context.Context, variant *Variant) (*List[TurnoverByOperation], *resty.Response, error) {
	path := fmt.Sprintf("%s/byoperations", service.uri)
	params := NewParams().WithFilterEquals("variant", *variant.Meta.Href)
	return NewRequestBuilder[List[TurnoverByOperation]](service.client, path).SetParams(params).Get(ctx)
}
