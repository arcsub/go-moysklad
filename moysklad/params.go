package moysklad

import (
	"fmt"
	"github.com/google/go-querystring/query"
	"strconv"
	"time"
)

// Params структура параметров запроса.
type Params struct {
	MomentFrom  string      `url:"momentFrom,omitempty"`
	GroupBy     GroupByType `url:"groupBy,omitempty"`
	Fields      string      `url:"fields,omitempty"`
	Interval    Interval    `url:"interval,omitempty"`
	Search      string      `url:"search,omitempty"`
	NamedFilter string      `url:"namedfilter,omitempty"`
	MomentTo    string      `url:"momentTo,omitempty"`
	StockType   StockType   `url:"stockType,omitempty"`
	Order       []string    `url:"order,omitempty" del:";"`
	Expand      []string    `url:"expand,omitempty" del:","`
	Filter      []string    `url:"filter,omitempty" del:";"`
	Offset      int         `url:"offset,omitempty"`
	Limit       int         `url:"limit,omitempty"`
	Async       bool        `url:"async,omitempty"`
}

func NewParams() *Params {
	return &Params{}
}

// WithMomentFrom Начало периода.
func (params *Params) WithMomentFrom(momentFrom time.Time) *Params {
	params.MomentFrom = momentFrom.Format(time.DateTime)
	return params
}

// WithMomentTo Конец периода.
func (params *Params) WithMomentTo(momentTo time.Time) *Params {
	params.MomentTo = momentTo.Format(time.DateTime)
	return params
}

// WithInterval Интервал, с которым будет построен отчет.
func (params *Params) WithInterval(interval Interval) *Params {
	params.Interval = interval
	return params
}

// withAsync Запрос на создание асинхронной задачи.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-asinhronnyj-obmen-vypolnenie-zaprosa-w-asinhronnom-rezhime
func (params *Params) withAsync() *Params {
	params.Async = true
	return params
}

// WithStockFiled Остатки и себестоимость в позициях документов.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-obschie-swedeniq-ostatki-i-sebestoimost-w-poziciqh-dokumentow
func (params *Params) WithStockFiled() *Params {
	params.Fields = "stock"
	params.WithLimit(100)
	params.WithExpand("positions")
	return params
}

// WithExpand Замена ссылок объектами.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/workbook/#workbook-chto-takoe-expand
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-zamena-ssylok-ob-ektami-s-pomosch-u-expand
func (params *Params) WithExpand(fieldName string) *Params {
	params.Expand = append(params.Expand, fieldName)
	return params
}

// FilterType Фильтрация выборки с помощью параметра Filter.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-fil-traciq-wyborki-s-pomosch-u-parametra-filter
type FilterType string

const (
	FilterEquals           FilterType = "="  // Фильтрация по значению
	FilterGreater          FilterType = ">"  // Больше
	FilterLesser           FilterType = "<"  // Меньше
	FilterGreaterOrEquals  FilterType = "=>" // Больше или равно
	FilterLesserOrEquals   FilterType = "<=" // Меньше или равно
	FilterNotEquals        FilterType = "!=" // Не равно
	FilterEquivalence      FilterType = "~"  // Подобие
	FilterEquivalenceLeft  FilterType = "~=" // Полное совпадение в начале значения
	FilterEquivalenceRight FilterType = "=~" // Полное совпадение в конце значения
	FilterNotEquivalence   FilterType = "!~" // Частичное совпадение не выводится
)

func newFilterParam(key, value string, filterType FilterType) string {
	return fmt.Sprintf("%s%s%s", key, filterType, value)
}

// WithFilterEquals Фильтрация по значению.
// key=value
func (params *Params) WithFilterEquals(key, value string) *Params {
	params.Filter = append(params.Filter, newFilterParam(key, value, FilterEquals))
	return params
}

// WithFilterGreater Больше.
// key>value
func (params *Params) WithFilterGreater(key, value string) *Params {
	params.Filter = append(params.Filter, newFilterParam(key, value, FilterGreater))
	return params
}

// WithFilterLesser Меньше.
// key<value
func (params *Params) WithFilterLesser(key, value string) *Params {
	params.Filter = append(params.Filter, newFilterParam(key, value, FilterLesser))
	return params
}

// WithFilterGreaterOrEquals Больше или равно.
// key=>value
func (params *Params) WithFilterGreaterOrEquals(key, value string) *Params {
	params.Filter = append(params.Filter, newFilterParam(key, value, FilterGreaterOrEquals))
	return params
}

// WithFilterLesserOrEquals Меньше или равно.
// key<=value
func (params *Params) WithFilterLesserOrEquals(key, value string) *Params {
	params.Filter = append(params.Filter, newFilterParam(key, value, FilterLesserOrEquals))
	return params
}

// WithFilterNotEquals Не равно.
// key!=value
func (params *Params) WithFilterNotEquals(key, value string) *Params {
	params.Filter = append(params.Filter, newFilterParam(key, value, FilterNotEquals))
	return params
}

// WithFilterEquivalence Частичное совпадение.
// key~value
func (params *Params) WithFilterEquivalence(key, value string) *Params {
	params.Filter = append(params.Filter, newFilterParam(key, value, FilterEquivalence))
	return params
}

// WithFilterEquivalenceLeft Полное совпадение в начале значения.
// key~=value
func (params *Params) WithFilterEquivalenceLeft(key, value string) *Params {
	params.Filter = append(params.Filter, newFilterParam(key, value, FilterEquivalenceLeft))
	return params
}

// WithFilterEquivalenceRight Полное совпадение в конце значения.
// key=~value
func (params *Params) WithFilterEquivalenceRight(key, value string) *Params {
	params.Filter = append(params.Filter, newFilterParam(key, value, FilterEquivalenceRight))
	return params
}

// WithFilterNotEquivalence Частичное совпадение не выводится.
// key!~value
func (params *Params) WithFilterNotEquivalence(key, value string) *Params {
	params.Filter = append(params.Filter, newFilterParam(key, value, FilterNotEquivalence))
	return params
}

// WithFilterDeleted Фильтрация по удалённым документам.
func (params *Params) WithFilterDeleted(value bool) *Params {
	params.Filter = append(params.Filter, newFilterParam("isDeleted", strconv.FormatBool(value), FilterEquals))
	return params
}

// WithFilterPrinted Фильтрация по напечатанным документам.
// printed=true
func (params *Params) WithFilterPrinted(value bool) *Params {
	params.Filter = append(params.Filter, newFilterParam("printed", strconv.FormatBool(value), FilterEquals))
	return params
}

// WithFilterPublished Фильтрация по опубликованным документам.
// published=true
func (params *Params) WithFilterPublished(value bool) *Params {
	params.Filter = append(params.Filter, newFilterParam("published", strconv.FormatBool(value), FilterEquals))
	return params
}

// WithFilterArchived Фильтрация по архивным сущностям.
// archived=true
func (params *Params) WithFilterArchived(value bool) *Params {
	params.Filter = append(params.Filter, newFilterParam("archived", strconv.FormatBool(value), FilterEquals))
	return params
}

// WithGroupBy Группировка выдачи.
func (params *Params) WithGroupBy(value GroupByType) *Params {
	params.GroupBy = value
	return params
}

// GroupByType Тип группировки выдачи.
type GroupByType string

const (
	GroupByProduct     GroupByType = "product"     // Выдает только товары
	GroupByVariant     GroupByType = "variant"     // Выдает товары и модификации
	GroupByConsignment GroupByType = "consignment" // Выдает товары, модификации, серии
)

func (groupByType GroupByType) String() string {
	return string(groupByType)
}

// Листание.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/workbook/#workbook-fil-traciq-listanie-poisk-i-sortirowka-listanie

// WithLimit Количество элементов на странице.
// От 1 до 1000
func (params *Params) WithLimit(limit int) *Params {
	params.Limit = Clamp(limit, 1, MaxPositions)
	return params
}

// WithOffset Смещение от первого элемента.
func (params *Params) WithOffset(offset int) *Params {
	params.Offset = offset
	return params
}

// WithNamedFilter позволяет использовать сохранённый фильтр в качестве параметра.
func (params *Params) WithNamedFilter(filter *NamedFilter) *Params {
	if filter.Meta == nil {
		return params
	}
	if filter.Meta.Href == nil {
		return params
	}
	params.NamedFilter = *filter.Meta.Href
	return params
}

// Сортировка объектов.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/workbook/#workbook-fil-traciq-listanie-poisk-i-sortirowka-sortirowka
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-sortirowka-ob-ektow

// OrderDirection Направление сортировки.
type OrderDirection string

const (
	OrderDirectionDefault OrderDirection = ""     // По умолчанию
	OrderDirectionAsc     OrderDirection = "asc"  // По возрастанию (Значение по умолчанию)
	OrderDirectionDesc    OrderDirection = "desc" // По убыванию
)

func newOrderParam(fieldName string, dir OrderDirection) string {
	if dir == OrderDirectionDefault {
		return fieldName
	}
	return fmt.Sprintf("%s,%s", fieldName, dir)
}

// WithOrder сортирует объекты по полю fieldName, по возрастанию (asc).
func (params *Params) WithOrder(fieldName string) *Params {
	params.Order = append(params.Order, newOrderParam(fieldName, OrderDirectionDefault))
	return params
}

// WithOrderAsc сортирует объекты по полю fieldName, по возрастанию (asc).
func (params *Params) WithOrderAsc(fieldName string) *Params {
	params.Order = append(params.Order, newOrderParam(fieldName, OrderDirectionAsc))
	return params
}

// WithOrderDesc сортирует объекты по полю fieldName, по убыванию (desc).
func (params *Params) WithOrderDesc(fieldName string) *Params {
	params.Order = append(params.Order, newOrderParam(fieldName, OrderDirectionDesc))
	return params
}

// WithSearch Контекстный поиск.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/workbook/#workbook-fil-traciq-listanie-poisk-i-sortirowka-poisk
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-kontextnyj-poisk
func (params *Params) WithSearch(search string) *Params {
	params.Search = search
	return params
}

// WithStockType Параметром stockType выбирается тип остатка, который необходимо рассчитать.
func (params *Params) WithStockType(stockType StockType) *Params {
	params.StockType = stockType
	return params
}

func (params *Params) QueryString() string {
	v, _ := query.Values(params)
	return v.Encode()

}

// StockType Параметр stockType.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-poluchit-kratkij-otchet-ob-ostatkah-parametr-stocktype
type StockType string

const (
	StockDefault   StockType = "stock"     // Физический остаток на складах, без учёта резерва и ожидания
	StockFreeStock StockType = "freeStock" // Остаток на складах за вычетом резерва
	StockQuantity  StockType = "quantity"  // Доступно. Учитывает резерв и ожидания
	StockReserve   StockType = "reserve"   // Резерв [05-10-2023]
	StockInTransit StockType = "inTransit" // Ожидание [05-10-2023]
)

func (stockType StockType) String() string {
	return string(stockType)
}
