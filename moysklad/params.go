package moysklad

import (
	"fmt"
	"github.com/google/go-querystring/query"
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

func GetParamsFromSliceOrNew(params []*Params) *Params {
	if len(params) == 0 {
		return NewParams()
	}

	if params[0] == nil {
		return NewParams()
	}

	return params[0]
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

// NewParams возвращает новый объект параметров запроса.
func NewParams() *Params {
	return &Params{}
}

// Clone копирует параметры запроса в новый указатель.
func (params *Params) Clone() *Params {
	clone := NewParams()
	*clone = *params
	return clone
}

// WithMomentFrom Начало периода.
//
// momentFrom=value
func (params *Params) WithMomentFrom(momentFrom time.Time) *Params {
	params.MomentFrom = momentFrom.Format(time.DateTime)
	return params
}

// WithMomentTo Конец периода.
//
// momentTo=value
func (params *Params) WithMomentTo(momentTo time.Time) *Params {
	params.MomentTo = momentTo.Format(time.DateTime)
	return params
}

// WithInterval Интервал, с которым будет построен отчёт.
//
// interval=value
func (params *Params) WithInterval(interval Interval) *Params {
	params.Interval = interval
	return params
}

// WithIntervalHour Интервал, с которым будет построен отчёт (час).
//
// interval=hour
func (params *Params) WithIntervalHour() *Params {
	params.Interval = IntervalHour
	return params
}

// WithIntervalDay Интервал, с которым будет построен отчёт (день).
//
// interval=day
func (params *Params) WithIntervalDay() *Params {
	params.Interval = IntervalDay
	return params
}

// WithIntervalMonth Интервал, с которым будет построен отчёт (месяц).
//
// interval=month
func (params *Params) WithIntervalMonth() *Params {
	params.Interval = IntervalMonth
	return params
}

// WithAsync Запрос на создание асинхронной задачи.
//
// async=true
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-asinhronnyj-obmen-vypolnenie-zaprosa-w-asinhronnom-rezhime
func (params *Params) WithAsync() *Params {
	params.Async = true
	return params
}

// WithStockFiled Остатки и себестоимость в позициях документов.
//
// fields=stock&limit=100&expand=positions
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-obschie-swedeniq-ostatki-i-sebestoimost-w-poziciqh-dokumentow
func (params *Params) WithStockFiled() *Params {
	params.Fields = "stock"
	params.WithLimit(100)
	params.WithExpand("positions")
	return params
}

// WithExpand Замена ссылок объектами.
//
// expand=fieldName
//
// [Документация МойСклад #1]
//
// [Документация МойСклад #2]
//
// [Документация МойСклад #1]: https://dev.moysklad.ru/doc/api/remap/1.2/workbook/#workbook-chto-takoe-expand
// [Документация МойСклад #2]: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-zamena-ssylok-ob-ektami-s-pomosch-u-expand
func (params *Params) WithExpand(fieldName string) *Params {
	params.Expand = append(params.Expand, fieldName)
	return params
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
func (params *Params) WithFilterObject(object MetaOwner) *Params {
	metaType := object.GetMeta().GetType().String()
	href := object.GetMeta().GetHref()

	if metaType != "" && href != "" {
		params.Filter = append(params.Filter, newFilter(metaType, href, FilterEquals))
	}

	return params
}

// WithFilterEquals Фильтрация по значению.
//
// key=value
func (params *Params) WithFilterEquals(key, value string) *Params {
	params.Filter = append(params.Filter, newFilter(key, value, FilterEquals))
	return params
}

// WithFilterGreater Больше.
//
// key>value
func (params *Params) WithFilterGreater(key, value string) *Params {
	params.Filter = append(params.Filter, newFilter(key, value, FilterGreater))
	return params
}

// WithFilterLesser Меньше.
//
// key<value
func (params *Params) WithFilterLesser(key, value string) *Params {
	params.Filter = append(params.Filter, newFilter(key, value, FilterLesser))
	return params
}

// WithFilterGreaterOrEquals Больше или равно.
//
// key>=value
func (params *Params) WithFilterGreaterOrEquals(key, value string) *Params {
	params.Filter = append(params.Filter, newFilter(key, value, FilterGreaterOrEquals))
	return params
}

// WithFilterLesserOrEquals Меньше или равно.
//
// key<=value
func (params *Params) WithFilterLesserOrEquals(key, value string) *Params {
	params.Filter = append(params.Filter, newFilter(key, value, FilterLesserOrEquals))
	return params
}

// WithFilterNotEquals Не равно.
//
// key!=value
func (params *Params) WithFilterNotEquals(key, value string) *Params {
	params.Filter = append(params.Filter, newFilter(key, value, FilterNotEquals))
	return params
}

// WithFilterEquivalence Частичное совпадение.
//
// key~value
func (params *Params) WithFilterEquivalence(key, value string) *Params {
	params.Filter = append(params.Filter, newFilter(key, value, FilterEquivalence))
	return params
}

// WithFilterEquivalenceLeft Полное совпадение в начале значения.
//
// key~=value
func (params *Params) WithFilterEquivalenceLeft(key, value string) *Params {
	params.Filter = append(params.Filter, newFilter(key, value, FilterEquivalenceLeft))
	return params
}

// WithFilterEquivalenceRight Полное совпадение в конце значения.
//
// key=~value
func (params *Params) WithFilterEquivalenceRight(key, value string) *Params {
	params.Filter = append(params.Filter, newFilter(key, value, FilterEquivalenceRight))
	return params
}

// WithFilterNotEquivalence Частичное совпадение не выводится.
//
// key!~value
func (params *Params) WithFilterNotEquivalence(key, value string) *Params {
	params.Filter = append(params.Filter, newFilter(key, value, FilterNotEquivalence))
	return params
}

// WithFilterDeleted Фильтрация по удалённым документам.
//
// isDeleted=true
func (params *Params) WithFilterDeleted(value bool) *Params {
	params.Filter = append(params.Filter, newFilter("isDeleted", strconv.FormatBool(value), FilterEquals))
	return params
}

// WithFilterPrinted Фильтрация по напечатанным документам.
//
// printed=true
func (params *Params) WithFilterPrinted(value bool) *Params {
	params.Filter = append(params.Filter, newFilter("printed", strconv.FormatBool(value), FilterEquals))
	return params
}

// WithFilterPublished Фильтрация по опубликованным документам.
//
// published=true
func (params *Params) WithFilterPublished(value bool) *Params {
	params.Filter = append(params.Filter, newFilter("published", strconv.FormatBool(value), FilterEquals))
	return params
}

// WithFilterArchived Фильтрация по архивным сущностям.
//
// archived=true
func (params *Params) WithFilterArchived(value bool) *Params {
	params.Filter = append(params.Filter, newFilter("archived", strconv.FormatBool(value), FilterEquals))
	return params
}

// WithGroupBy Группировка выдачи.
//
// groupBy=value
func (params *Params) WithGroupBy(value GroupBy) *Params {
	params.GroupBy = value
	return params
}

// WithGroupByProduct Группировка выдачи (только товары).
//
// groupBy=product
func (params *Params) WithGroupByProduct() *Params {
	params.GroupBy = GroupByProduct
	return params
}

// WithGroupByVariant Группировка выдачи (товары и модификации).
//
// groupBy=variant
func (params *Params) WithGroupByVariant() *Params {
	params.GroupBy = GroupByVariant
	return params
}

// WithGroupByConsignment Группировка выдачи (товары, модификации и серии).
//
// groupBy=consignment
func (params *Params) WithGroupByConsignment() *Params {
	params.GroupBy = GroupByConsignment
	return params
}

// WithLimit Количество элементов на странице.
//
// Диапазон значения: от 1 до 1000.
//
// limit=value
func (params *Params) WithLimit(limit int) *Params {
	params.Limit = Clamp(limit, 1, MaxPositions)
	return params
}

// WithOffset Смещение от первого элемента.
//
// offset=value
func (params *Params) WithOffset(offset int) *Params {
	params.Offset = offset
	return params
}

// WithNamedFilter позволяет использовать сохранённый фильтр в качестве параметра.
//
// namedfilter=https://api.moysklad.ru/api/remap/1.2/entity/product/namedfilter/b5863410-ca86-11eb-ac12-000d00000019
func (params *Params) WithNamedFilter(filter *NamedFilter) *Params {
	if filter.Meta == nil {
		return params
	}
	params.NamedFilter = filter.GetMeta().GetHref()
	return params
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
func (params *Params) WithOrder(fieldName string) *Params {
	params.Order = append(params.Order, newOrderParam(fieldName, OrderDirectionDefault))
	return params
}

// WithOrderAsc сортирует объекты по полю fieldName, по возрастанию (asc).
//
// order=fieldName,asc
func (params *Params) WithOrderAsc(fieldName string) *Params {
	params.Order = append(params.Order, newOrderParam(fieldName, OrderDirectionAsc))
	return params
}

// WithOrderDesc сортирует объекты по полю fieldName, по убыванию (desc).
//
// order=fieldName,desc
func (params *Params) WithOrderDesc(fieldName string) *Params {
	params.Order = append(params.Order, newOrderParam(fieldName, OrderDirectionDesc))
	return params
}

// WithSearch Контекстный поиск.
//
// search=value
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/workbook/#workbook-fil-traciq-listanie-poisk-i-sortirowka-poisk
func (params *Params) WithSearch(search string) *Params {
	params.Search = search
	return params
}

// WithStockType Параметром stockType выбирается тип остатка, который необходимо рассчитать.
//
// stockType=value
func (params *Params) WithStockType(stockType StockType) *Params {
	params.StockType = stockType
	return params
}

// WithStockDefault Физический остаток на складах, без учёта резерва и ожидания.
//
// stockType=stock
func (params *Params) WithStockDefault() *Params {
	params.StockType = StockDefault
	return params
}

// WithStockFree Остаток на складах за вычетом резерва.
//
// stockType=freeStock
func (params *Params) WithStockFree() *Params {
	params.StockType = StockFreeStock
	return params
}

// WithStockQuantity Доступно. Учитывает резерв и ожидания.
//
// stockType=quantity
func (params *Params) WithStockQuantity() *Params {
	params.StockType = StockQuantity
	return params
}

// WithStockReserve Резерв.
//
// stockType=reserve
func (params *Params) WithStockReserve() *Params {
	params.StockType = StockReserve
	return params
}

// WithStockInTransit Ожидание.
//
// stockType=inTransit
func (params *Params) WithStockInTransit() *Params {
	params.StockType = StockInTransit
	return params
}

// QueryString конкатенирует поля и возвращает URI валидную строку.
func (params *Params) QueryString() string {
	v, _ := query.Values(params)
	return v.Encode()
}

// WithEvaluate автозаполнение.
//
// action=value1,value2...
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-awtozapolnenie
func (params *Params) WithEvaluate(evaluate ...Evaluate) *Params {
	params.Action = append(params.Action, evaluate...)
	return params
}
