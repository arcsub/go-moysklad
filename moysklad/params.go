package moysklad

import (
	"fmt"
	"github.com/google/go-querystring/query"
	"net/url"
	"strconv"
	"time"
)

// Params объект параметров запроса.
type Params struct {
	MomentFrom  string     `url:"momentFrom,omitempty"`     // Параметр выборки "От даты"
	GroupBy     GroupBy    `url:"groupBy,omitempty"`        // Тип, по которому нужно сгруппировать выдачу
	Fields      string     `url:"fields,omitempty"`         // Получать остатки и себестоимость позиций этих документов
	Interval    Interval   `url:"interval,omitempty"`       // Интервал, с которым будет построен отчет
	Search      string     `url:"search,omitempty"`         // Контекстный поиск
	NamedFilter string     `url:"namedfilter,omitempty"`    // Применение сохраненного фильтра
	MomentTo    string     `url:"momentTo,omitempty"`       // Параметр выборки "До даты"
	StockType   StockType  `url:"stockType,omitempty"`      // тип остатка, резерва, ожидания, которые необходимо рассчитать
	Order       []string   `url:"order,omitempty" del:";"`  // Сортировка списка объектов
	Expand      []string   `url:"expand,omitempty" del:","` // Замена ссылок объектами
	Filter      []string   `url:"filter,omitempty" del:";"` // Фильтрация
	Action      []Evaluate `url:"action,omitempty" del:","` // Параметры автозаполнения
	Offset      int        `url:"offset,omitempty"`         // Смещение от первого элемента (считается с нуля)
	Limit       int        `url:"limit,omitempty"`          // Количество элементов на странице (по умолчанию 1000, максимум 1000)
	Async       bool       `url:"async,omitempty"`          // Параметр создания асинхронной задачи
}

// String реализует интерфейс [fmt.Stringer].
func (params *Params) String() string {
	return params.Values().Encode()
}

func (params *Params) Values() url.Values {
	v, _ := query.Values(params)

	return v
}

func ApplyParams(params []func(*Params)) *Params {
	p := &Params{}

	for _, o := range params {
		o(p)
	}

	return p
}

// Interval Интервал, с которым будет построен отчёт.
//
// Возможные значения:
//   - IntervalHour  – час
//   - IntervalDay   – день
//   - IntervalMonth – месяц
type Interval string

func (interval Interval) String() string {
	return string(interval)
}

const (
	IntervalHour  Interval = "hour"
	IntervalDay   Interval = "day"
	IntervalMonth Interval = "month"
)

// FilterType Фильтрация выборки с помощью параметра Filter.
//
// Возможные значения:
//   - FilterEquals           – Фильтрация по значению ("=")
//   - FilterGreater          – Больше (">")
//   - FilterLesser           – Меньше ("<")
//   - FilterGreaterOrEquals  – Больше или равно (">=")
//   - FilterLesserOrEquals   – Меньше или равно ("<=")
//   - FilterNotEquals        – Не равно ("!=")
//   - FilterEquivalence      – Подобие ("~")
//   - FilterEquivalenceLeft  – Полное совпадение в начале значения ("~=")
//   - FilterEquivalenceRight – Полное совпадение в конце значения ("=~")
//   - FilterNotEquivalence   – Частичное совпадение не выводится ("!~")
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-fil-traciq-wyborki-s-pomosch-u-parametra-filter
type FilterType string

const (
	FilterEquals           FilterType = "="  // Фильтрация по значению
	FilterGreater          FilterType = ">"  // Больше
	FilterLesser           FilterType = "<"  // Меньше
	FilterGreaterOrEquals  FilterType = ">=" // Больше или равно
	FilterLesserOrEquals   FilterType = "<=" // Меньше или равно
	FilterNotEquals        FilterType = "!=" // Не равно
	FilterEquivalence      FilterType = "~"  // Подобие
	FilterEquivalenceLeft  FilterType = "~=" // Полное совпадение в начале значения
	FilterEquivalenceRight FilterType = "=~" // Полное совпадение в конце значения
	FilterNotEquivalence   FilterType = "!~" // Частичное совпадение не выводится
)

// Evaluate определяет какую информацию нужно заполнить.
//
// Возможные значения:
//   - EvaluateDiscount – скидки
//   - EvaluatePrice    – цены
//   - EvaluateVat      – ндс
//   - EvaluateCost     – себестоимость
type Evaluate string

const (
	EvaluateDiscount = "evaluate_discount" // скидки
	EvaluatePrice    = "evaluate_price"    // цены
	EvaluateVat      = "evaluate_vat"      // ндс
	EvaluateCost     = "evaluate_cost"     // себестоимость
)

// String реализует интерфейс [fmt.Stringer].
func (evaluate Evaluate) String() string {
	return string(evaluate)
}

// StockType Параметр stockType.
//
// Возможные значения:
//   - StockDefault   – Физический остаток на складах, без учёта резерва и ожидания
//   - StockFreeStock – Остаток на складах за вычетом резерва
//   - StockQuantity  – Доступно. Учитывает резерв и ожидания
//   - StockReserve   – Резерв
//   - StockInTransit – Ожидание
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-poluchit-kratkij-otchet-ob-ostatkah-parametr-stocktype
type StockType string

const (
	StockDefault   StockType = "stock"     // Физический остаток на складах, без учёта резерва и ожидания
	StockFreeStock StockType = "freeStock" // Остаток на складах за вычетом резерва
	StockQuantity  StockType = "quantity"  // Доступно. Учитывает резерв и ожидания
	StockReserve   StockType = "reserve"   // Резерв
	StockInTransit StockType = "inTransit" // Ожидание
)

// String реализует интерфейс [fmt.Stringer].
func (stockType StockType) String() string {
	return string(stockType)
}

// OrderDirection направление сортировки.
//
// Возможные значения:
//   - OrderDirectionDefault – По умолчанию (по возрастанию)
//   - OrderDirectionAsc     – По возрастанию (по умолчанию)
//   - OrderDirectionDesc    – По убыванию
type OrderDirection string

const (
	OrderDirectionDefault OrderDirection = ""     // По умолчанию
	OrderDirectionAsc     OrderDirection = "asc"  // По возрастанию (Значение по умолчанию)
	OrderDirectionDesc    OrderDirection = "desc" // По убыванию
)

// GroupBy Тип группировки выдачи.
//
// Возможные значения:
//   - GroupByProduct     – Выдает только товары
//   - GroupByVariant     – Выдает товары и модификации
//   - GroupByConsignment – Выдает товары, модификации, серии
type GroupBy string

const (
	GroupByProduct     GroupBy = "product"     // Выдает только товары
	GroupByVariant     GroupBy = "variant"     // Выдает товары и модификации
	GroupByConsignment GroupBy = "consignment" // Выдает товары, модификации, серии
)

// String реализует интерфейс [fmt.Stringer].
func (groupByType GroupBy) String() string {
	return string(groupByType)
}

// WithMomentFrom Начало периода.
//
// momentFrom=value
func WithMomentFrom(momentFrom time.Time) func(*Params) {
	return func(params *Params) {
		params.MomentFrom = momentFrom.Format(time.DateTime)
	}
}

// WithMomentTo Конец периода.
//
// momentTo=value
func WithMomentTo(momentTo time.Time) func(*Params) {
	return func(params *Params) {
		params.MomentFrom = momentTo.Format(time.DateTime)
	}
}

// WithInterval Интервал, с которым будет построен отчёт.
//
// interval=value
func WithInterval(interval Interval) func(*Params) {
	return func(params *Params) {
		params.Interval = interval
	}
}

// WithIntervalHour Интервал, с которым будет построен отчёт (час).
//
// interval=hour
func WithIntervalHour() func(*Params) {
	return func(params *Params) {
		params.Interval = IntervalHour
	}
}

// WithIntervalDay Интервал, с которым будет построен отчёт (день).
//
// interval=day
func WithIntervalDay() func(*Params) {
	return func(params *Params) {
		params.Interval = IntervalDay
	}
}

// WithIntervalMonth Интервал, с которым будет построен отчёт (месяц).
//
// interval=month
func WithIntervalMonth() func(*Params) {
	return func(params *Params) {
		params.Interval = IntervalMonth
	}
}

// WithAsync Запрос на создание асинхронной задачи.
//
// async=true
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-asinhronnyj-obmen-vypolnenie-zaprosa-w-asinhronnom-rezhime
func WithAsync() func(*Params) {
	return func(params *Params) {
		params.Async = true
	}
}

// WithStockFiled Остатки и себестоимость в позициях документов.
//
// fields=stock&limit=100&expand=positions
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-obschie-swedeniq-ostatki-i-sebestoimost-w-poziciqh-dokumentow
func WithStockFiled() func(*Params) {
	return func(params *Params) {
		params.Fields = "stock"
		params.Limit = 100
		params.Expand = append(params.Expand, "positions")
	}
}

// WithExpand Замена ссылок объектами.
//
// expand=fieldName1,fieldName2,...
//
// [Документация МойСклад #1]
//
// [Документация МойСклад #2]
//
// [Документация МойСклад #1]: https://dev.moysklad.ru/doc/api/remap/1.2/workbook/#workbook-chto-takoe-expand
// [Документация МойСклад #2]: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-zamena-ssylok-ob-ektami-s-pomosch-u-expand
func WithExpand(fields ...string) func(*Params) {
	return func(params *Params) {
		params.Expand = append(params.Expand, fields...)
	}
}

func newFilter(key, value string, filterType FilterType) string {
	return fmt.Sprintf("%s%s%s", key, filterType, value)
}

// WithFilterObject принимает объект, реализующий интерфейс [MetaOwner] и передаёт его ссылку в качестве фильтрации.
//
// Например:
//
// `params.WithFilterObject(store)`
//
// где store - указатель на [Store], который содержит заполненный объект [Meta].
//
// ?filter=store=https://api.moysklad.ru/api/remap/1.2/entity/store/dbc4fab9-2226-11ee-0a80-112e00094247
func WithFilterObject(object MetaOwner) func(*Params) {
	return func(params *Params) {
		metaType := object.GetMeta().GetType().String()
		href := object.GetMeta().GetHref()

		if metaType != "" && href != "" {
			params.Filter = append(params.Filter, newFilter(metaType, href, FilterEquals))
		}
	}
}

// WithFilterEquals Фильтрация по значению.
//
// key=value
func WithFilterEquals(key, value string) func(*Params) {
	return func(params *Params) {
		params.Filter = append(params.Filter, newFilter(key, value, FilterEquals))
	}
}

// WithFilterGreater Больше.
//
// key>value
func WithFilterGreater(key, value string) func(*Params) {
	return func(params *Params) {
		params.Filter = append(params.Filter, newFilter(key, value, FilterGreater))
	}
}

// WithFilterLesser Меньше.
//
// key<value
func WithFilterLesser(key, value string) func(*Params) {
	return func(params *Params) {
		params.Filter = append(params.Filter, newFilter(key, value, FilterLesser))
	}
}

// WithFilterGreaterOrEquals Больше или равно.
//
// key>=value
func WithFilterGreaterOrEquals(key, value string) func(*Params) {
	return func(params *Params) {
		params.Filter = append(params.Filter, newFilter(key, value, FilterGreaterOrEquals))
	}
}

// WithFilterLesserOrEquals Меньше или равно.
//
// key<=value
func WithFilterLesserOrEquals(key, value string) func(*Params) {
	return func(params *Params) {
		params.Filter = append(params.Filter, newFilter(key, value, FilterLesserOrEquals))
	}
}

// WithFilterNotEquals Не равно.
//
// key!=value
func WithFilterNotEquals(key, value string) func(*Params) {
	return func(params *Params) {
		params.Filter = append(params.Filter, newFilter(key, value, FilterNotEquals))
	}
}

// WithFilterEquivalence Частичное совпадение.
//
// key~value
func WithFilterEquivalence(key, value string) func(*Params) {
	return func(params *Params) {
		params.Filter = append(params.Filter, newFilter(key, value, FilterEquivalence))
	}
}

// WithFilterEquivalenceLeft Полное совпадение в начале значения.
//
// key~=value
func WithFilterEquivalenceLeft(key, value string) func(*Params) {
	return func(params *Params) {
		params.Filter = append(params.Filter, newFilter(key, value, FilterEquivalenceLeft))
	}
}

// WithFilterEquivalenceRight Полное совпадение в конце значения.
//
// key=~value
func WithFilterEquivalenceRight(key, value string) func(*Params) {
	return func(params *Params) {
		params.Filter = append(params.Filter, newFilter(key, value, FilterEquivalenceRight))
	}
}

// WithFilterNotEquivalence Частичное совпадение не выводится.
//
// key!~value
func WithFilterNotEquivalence(key, value string) func(*Params) {
	return func(params *Params) {
		params.Filter = append(params.Filter, newFilter(key, value, FilterNotEquivalence))
	}
}

// WithFilterDeleted Фильтрация по удалённым документам.
//
// isDeleted=true
func WithFilterDeleted(value bool) func(*Params) {
	return func(params *Params) {
		params.Filter = append(params.Filter, newFilter("isDeleted", strconv.FormatBool(value), FilterEquals))
	}
}

// WithFilterPrinted Фильтрация по напечатанным документам.
//
// printed=true
func WithFilterPrinted(value bool) func(*Params) {
	return func(params *Params) {
		params.Filter = append(params.Filter, newFilter("printed", strconv.FormatBool(value), FilterEquals))
	}
}

// WithFilterPublished Фильтрация по опубликованным документам.
//
// published=true
func WithFilterPublished(value bool) func(*Params) {
	return func(params *Params) {
		params.Filter = append(params.Filter, newFilter("published", strconv.FormatBool(value), FilterEquals))
	}
}

// WithFilterArchived Фильтрация по архивным сущностям.
//
// archived=true
func WithFilterArchived(value bool) func(*Params) {
	return func(params *Params) {
		params.Filter = append(params.Filter, newFilter("archived", strconv.FormatBool(value), FilterEquals))
	}
}

// WithGroupBy Группировка выдачи.
//
// groupBy=value
func WithGroupBy(value GroupBy) func(*Params) {
	return func(params *Params) {
		params.GroupBy = value
	}
}

// WithGroupByProduct Группировка выдачи (только товары).
//
// groupBy=product
func WithGroupByProduct() func(*Params) {
	return func(params *Params) {
		params.GroupBy = GroupByProduct
	}
}

// WithGroupByVariant Группировка выдачи (товары и модификации).
//
// groupBy=variant
func WithGroupByVariant() func(*Params) {
	return func(params *Params) {
		params.GroupBy = GroupByVariant
	}
}

// WithGroupByConsignment Группировка выдачи (товары, модификации и серии).
//
// groupBy=consignment
func WithGroupByConsignment() func(*Params) {
	return func(params *Params) {
		params.GroupBy = GroupByConsignment
	}
}

// WithLimit Количество элементов на странице.
//
// Диапазон значения: от 1 до 1000.
//
// limit=value
func WithLimit(limit int) func(*Params) {
	return func(params *Params) {
		params.Limit = Clamp(limit, 1, MaxPositions)
	}
}

// WithOffset Смещение от первого элемента.
//
// offset=value
func WithOffset(offset int) func(*Params) {
	return func(params *Params) {
		params.Offset = offset
	}
}

// WithNamedFilter позволяет использовать сохранённый фильтр в качестве параметра.
//
// namedfilter=https://api.moysklad.ru/api/remap/1.2/entity/product/namedfilter/b5863410-ca86-11eb-ac12-000d00000019
func WithNamedFilter(filter *NamedFilter) func(*Params) {
	return func(params *Params) {
		if filter.Meta != nil {
			params.NamedFilter = filter.GetMeta().GetHref()
		}
	}
}

func newOrderParam(fieldName string, dir OrderDirection) string {
	if dir == OrderDirectionDefault {
		return fieldName
	}
	return fmt.Sprintf("%s,%s", fieldName, dir)
}

// WithOrder сортирует объекты по полю fieldName, по умолчанию (asc).
//
// order=fieldName
func WithOrder(fields ...string) func(*Params) {
	return func(params *Params) {
		for _, fieldName := range fields {
			params.Order = append(params.Order, newOrderParam(fieldName, OrderDirectionDefault))
		}
	}
}

// WithOrderAsc сортирует объекты по полю fieldName, по возрастанию (asc).
//
// order=fieldName,asc
func WithOrderAsc(fields ...string) func(*Params) {
	return func(params *Params) {
		for _, fieldName := range fields {
			params.Order = append(params.Order, newOrderParam(fieldName, OrderDirectionAsc))
		}
	}
}

// WithOrderDesc сортирует объекты по полю fieldName, по убыванию (desc).
//
// order=fieldName,desc
func WithOrderDesc(fields ...string) func(*Params) {
	return func(params *Params) {
		for _, fieldName := range fields {
			params.Order = append(params.Order, newOrderParam(fieldName, OrderDirectionDesc))
		}
	}
}

// WithSearch Контекстный поиск.
//
// search=value
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/workbook/#workbook-fil-traciq-listanie-poisk-i-sortirowka-poisk
func WithSearch(search string) func(*Params) {
	return func(params *Params) {
		params.Search = search
	}
}

// WithStockType Параметром stockType выбирается тип остатка, который необходимо рассчитать.
//
// stockType=value
func WithStockType(stockType StockType) func(*Params) {
	return func(params *Params) {
		params.StockType = stockType
	}
}

// WithStockDefault Физический остаток на складах, без учёта резерва и ожидания.
//
// stockType=stock
func WithStockDefault() func(*Params) {
	return func(params *Params) {
		params.StockType = StockDefault
	}
}

// WithStockFree Остаток на складах за вычетом резерва.
//
// stockType=freeStock
func WithStockFree() func(*Params) {
	return func(params *Params) {
		params.StockType = StockFreeStock
	}
}

// WithStockQuantity Доступно. Учитывает резерв и ожидания.
//
// stockType=quantity
func WithStockQuantity() func(*Params) {
	return func(params *Params) {
		params.StockType = StockQuantity
	}
}

// WithStockReserve Резерв.
//
// stockType=reserve
func WithStockReserve() func(*Params) {
	return func(params *Params) {
		params.StockType = StockReserve
	}
}

// WithStockInTransit Ожидание.
//
// stockType=inTransit
func WithStockInTransit() func(*Params) {
	return func(params *Params) {
		params.StockType = StockInTransit
	}
}

// WithEvaluate автозаполнение.
//
// action=value1,value2...
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-awtozapolnenie
func WithEvaluate(evaluate ...Evaluate) func(*Params) {
	return func(params *Params) {
		params.Action = append(params.Action, evaluate...)
	}
}
