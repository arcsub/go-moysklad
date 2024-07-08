package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
)

// TurnoverIncomeOutcome Структура объекта показатели (onPeriodStart, onPeriodEnd, income, outcome).
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-oboroty-oboroty-po-towaram-struktura-ob-ekta-pokazateli-onperiodstart-onperiodend-income-outcome
type TurnoverIncomeOutcome struct {
	Sum      float64 `json:"sum"`      // Сумма себестоимости
	Quantity float64 `json:"quantity"` // Количество единиц товара
}

// TurnoverAll Обороты по товарам.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-oboroty-poluchit-oboroty-po-towaram
type TurnoverAll struct {
	Assortment    MetaNameWrapper       `json:"assortment"`    // Краткое представление Товара или Модификации в отчёте
	OnPeriodStart TurnoverIncomeOutcome `json:"onPeriodStart"` // Показатели на начало периода
	OnPeriodEnd   TurnoverIncomeOutcome `json:"onPeriodEnd"`   // Показатели на конец периода
	Income        TurnoverIncomeOutcome `json:"income"`        // Показатели прихода в течение периода отчёта
	Outcome       TurnoverIncomeOutcome `json:"outcome"`       // Показатели расхода в течение периода отчёта
}

// TurnoverAssortment Структура объекта assortment.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-oboroty-oboroty-po-towaru-s-detalizaciej-po-skladam-struktura-ob-ekta-assortment
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
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-oboroty-oboroty-po-towaru-s-detalizaciej-po-dokumentam
type TurnoverByOperation struct {
	Assortment TurnoverAssortment `json:"assortment"` // Краткое представление Товара или Модификации в отчете
	Operation  TurnoverOperation  `json:"operation"`  // Документ, связанный с Товаром
	Store      MetaNameWrapper    `json:"store"`      // Склад
	Cost       float64            `json:"cost"`       // Себестоимость товара в копейках в документе
	Sum        float64            `json:"sum"`        // Сумма себестоимостей в копейках
	Quantity   float64            `json:"quantity"`   // Количество товара в документе
}

// TurnoverOperation Структура объекта operation.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-oboroty-oboroty-po-towaru-s-detalizaciej-po-dokumentam-struktura-ob-ekta-operation
type TurnoverOperation struct {
	Meta        Meta            `json:"meta"`        // Метаданные документа
	Name        string          `json:"name"`        // Номер документа
	Description string          `json:"description"` // Комментарий к документу
	Moment      Timestamp       `json:"moment"`      // Дата проведения документа
	Agent       MetaNameWrapper `json:"agent"`       // Контрагент документа
}

// TurnoverByStore Обороты по товару с детализацией по складам.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-oboroty-oboroty-po-towaru-s-detalizaciej-po-skladam
type TurnoverByStore struct {
	Assortment   TurnoverAssortment       `json:"assortment"`   // Краткое представление Товара или Модификации в отчёте
	StockByStore []TurnoverByStoreElement `json:"stockByStore"` // Детализация оборотов по складам
}

// TurnoverByStoreElement Структура объекта детализация оборотов по складам.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-oboroty-oboroty-po-towaru-s-detalizaciej-po-skladam-struktura-ob-ekta-detalizaciq-oborotow-po-skladam
type TurnoverByStoreElement struct {
	Store         MetaNameWrapper       `json:"store"`         // Склад
	OnPeriodStart TurnoverIncomeOutcome `json:"onPeriodStart"` // Показатели на начало периода
	OnPeriodEnd   TurnoverIncomeOutcome `json:"onPeriodEnd"`   // Показатели на конец периода
	Income        TurnoverIncomeOutcome `json:"income"`        // Показатели прихода в течение периода отчёта
	Outcome       TurnoverIncomeOutcome `json:"outcome"`       // Показатели расхода в течение периода отчёта
}

// ReportTurnoverService описывает методы сервиса для работы с отчётом Обороты.
type ReportTurnoverService interface {
	// GetAll выполняет запрос на получение отчёта "Обороты по товарам".
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetAll(ctx context.Context, params ...*Params) (*List[TurnoverAll], *resty.Response, error)

	// GetByStoreWithProduct выполняет запрос на получение отчёта обороты по товару и его модификациям с детализацией по складам.
	// Принимает контекст и товар.
	// Возвращает объект List.
	GetByStoreWithProduct(ctx context.Context, product *Product) (*List[TurnoverByOperation], *resty.Response, error)

	// GetByStoreWithVariant выполняет запрос на получение отчёта обороты по модификации с детализацией по складам.
	// Принимает контекст и модификацию.
	// Возвращает объект List.
	GetByStoreWithVariant(ctx context.Context, variant *Variant) (*List[TurnoverByOperation], *resty.Response, error)

	// GetByOperationsWithProduct выполняет запрос на получение отчёта обороты по товару с детализацией по документам.
	// Принимает контекст и модификацию.
	// Возвращает объект List.
	GetByOperationsWithProduct(ctx context.Context, product *Product) (*List[TurnoverByOperation], *resty.Response, error)

	// GetByOperationsWithVariant выполняет запрос на получение отчёта обороты по модификации с детализацией по документам.
	// Принимает контекст и модификацию.
	// Возвращает объект List.
	GetByOperationsWithVariant(ctx context.Context, variant *Variant) (*List[TurnoverByOperation], *resty.Response, error)
}

type reportTurnoverService struct {
	Endpoint
}

// NewReportTurnoverService принимает [Client] и возвращает сервис для работы с отчётом Обороты.
func NewReportTurnoverService(client *Client) ReportTurnoverService {
	return &reportTurnoverService{NewEndpoint(client, "report/turnover")}
}

func (service *reportTurnoverService) GetAll(ctx context.Context, params ...*Params) (*List[TurnoverAll], *resty.Response, error) {
	path := fmt.Sprintf("%s/all", service.uri)
	return NewRequestBuilder[List[TurnoverAll]](service.client, path).SetParams(params...).Get(ctx)
}

func (service *reportTurnoverService) GetByStoreWithProduct(ctx context.Context, product *Product) (*List[TurnoverByOperation], *resty.Response, error) {
	path := fmt.Sprintf("%s/bystore", service.uri)
	params := NewParams().WithFilterObject(product)
	return NewRequestBuilder[List[TurnoverByOperation]](service.client, path).SetParams(params).Get(ctx)
}

func (service *reportTurnoverService) GetByStoreWithVariant(ctx context.Context, variant *Variant) (*List[TurnoverByOperation], *resty.Response, error) {
	path := fmt.Sprintf("%s/bystore", service.uri)
	params := NewParams().WithFilterObject(variant)
	return NewRequestBuilder[List[TurnoverByOperation]](service.client, path).SetParams(params).Get(ctx)
}

func (service *reportTurnoverService) GetByOperationsWithProduct(ctx context.Context, product *Product) (*List[TurnoverByOperation], *resty.Response, error) {
	path := fmt.Sprintf("%s/byoperations", service.uri)
	params := NewParams().WithFilterObject(product)
	return NewRequestBuilder[List[TurnoverByOperation]](service.client, path).SetParams(params).Get(ctx)
}

func (service *reportTurnoverService) GetByOperationsWithVariant(ctx context.Context, variant *Variant) (*List[TurnoverByOperation], *resty.Response, error) {
	path := fmt.Sprintf("%s/byoperations", service.uri)
	params := NewParams().WithFilterObject(variant)
	return NewRequestBuilder[List[TurnoverByOperation]](service.client, path).SetParams(params).Get(ctx)
}
