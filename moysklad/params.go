package moysklad

import (
	"fmt"
	"github.com/google/go-querystring/query"
	"strconv"
	"time"
)

// Params структура параметров запроса.
type Params struct {
	Limit       int         `url:"limit,omitempty"`
	Offset      int         `url:"offset,omitempty"`
	Filter      []string    `url:"filter,omitempty" del:";"`
	Expand      []string    `url:"expand,omitempty" del:","`
	Search      string      `url:"search,omitempty"`
	Order       []string    `url:"order,omitempty" del:";"`
	GroupBy     GroupByType `url:"groupBy,omitempty"`
	StockType   StockType   `url:"stockType,omitempty"`
	NamedFilter string      `url:"namedfilter,omitempty"`
	MomentFrom  string      `url:"momentFrom,omitempty"`
	MomentTo    string      `url:"momentTo,omitempty"`
	Interval    Interval    `url:"interval,omitempty"`
	Async       bool        `url:"async,omitempty"`
	Fields      string      `url:"fields,omitempty"`
}

// WithMomentFrom Начало периода.
func (p *Params) WithMomentFrom(momentFrom time.Time) *Params {
	p.MomentFrom = momentFrom.Format(time.DateTime)
	return p
}

// WithMomentTo Конец периода.
func (p *Params) WithMomentTo(momentTo time.Time) *Params {
	p.MomentTo = momentTo.Format(time.DateTime)
	return p
}

// WithInterval Интервал, с которым будет построен отчет.
func (p *Params) WithInterval(interval Interval) *Params {
	p.Interval = interval
	return p
}

// withAsync Запрос на создание асинхронной задачи.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-asinhronnyj-obmen-vypolnenie-zaprosa-w-asinhronnom-rezhime
func (p *Params) withAsync() *Params {
	p.Async = true
	return p
}

// WithStockFiled Остатки и себестоимость в позициях документов.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-obschie-swedeniq-ostatki-i-sebestoimost-w-poziciqh-dokumentow
func (p *Params) WithStockFiled() *Params {
	p.Fields = "stock"
	p.WithLimit(100)
	p.WithExpand("positions")
	return p
}

// WithExpand Замена ссылок объектами.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/workbook/#workbook-chto-takoe-expand
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-zamena-ssylok-ob-ektami-s-pomosch-u-expand
func (p *Params) WithExpand(fieldName string) *Params {
	p.Expand = append(p.Expand, fieldName)
	return p
}

// FilterType Фильтрация выборки с помощью параметра filter.
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
func (p *Params) WithFilterEquals(key, value string) *Params {
	p.Filter = append(p.Filter, newFilterParam(key, value, FilterEquals))
	return p
}

// WithFilterGreater Больше.
// key>value
func (p *Params) WithFilterGreater(key, value string) *Params {
	p.Filter = append(p.Filter, newFilterParam(key, value, FilterGreater))
	return p
}

// WithFilterLesser Меньше.
// key<value
func (p *Params) WithFilterLesser(key, value string) *Params {
	p.Filter = append(p.Filter, newFilterParam(key, value, FilterLesser))
	return p
}

// WithFilterGreaterOrEquals Больше или равно.
// key=>value
func (p *Params) WithFilterGreaterOrEquals(key, value string) *Params {
	p.Filter = append(p.Filter, newFilterParam(key, value, FilterGreaterOrEquals))
	return p
}

// WithFilterLesserOrEquals Меньше или равно.
// key<=value
func (p *Params) WithFilterLesserOrEquals(key, value string) *Params {
	p.Filter = append(p.Filter, newFilterParam(key, value, FilterLesserOrEquals))
	return p
}

// WithFilterNotEquals Не равно.
// key!=value
func (p *Params) WithFilterNotEquals(key, value string) *Params {
	p.Filter = append(p.Filter, newFilterParam(key, value, FilterNotEquals))
	return p
}

// WithFilterEquivalence Частичное совпадение.
// key~value
func (p *Params) WithFilterEquivalence(key, value string) *Params {
	p.Filter = append(p.Filter, newFilterParam(key, value, FilterEquivalence))
	return p
}

// WithFilterEquivalenceLeft Полное совпадение в начале значения.
// key~=value
func (p *Params) WithFilterEquivalenceLeft(key, value string) *Params {
	p.Filter = append(p.Filter, newFilterParam(key, value, FilterEquivalenceLeft))
	return p
}

// WithFilterEquivalenceRight Полное совпадение в конце значения.
// key=~value
func (p *Params) WithFilterEquivalenceRight(key, value string) *Params {
	p.Filter = append(p.Filter, newFilterParam(key, value, FilterEquivalenceRight))
	return p
}

// WithFilterNotEquivalence Частичное совпадение не выводится.
// key!~value
func (p *Params) WithFilterNotEquivalence(key, value string) *Params {
	p.Filter = append(p.Filter, newFilterParam(key, value, FilterNotEquivalence))
	return p
}

// WithFilterDeleted Фильтрация по удалённым документам.
func (p *Params) WithFilterDeleted(value bool) *Params {
	p.Filter = append(p.Filter, newFilterParam("isDeleted", strconv.FormatBool(value), FilterEquals))
	return p
}

// WithFilterPrinted Фильтрация по напечатанным документам.
// printed=true
func (p *Params) WithFilterPrinted(value bool) *Params {
	p.Filter = append(p.Filter, newFilterParam("printed", strconv.FormatBool(value), FilterEquals))
	return p
}

// WithFilterPublished Фильтрация по опубликованным документам.
// published=true
func (p *Params) WithFilterPublished(value bool) *Params {
	p.Filter = append(p.Filter, newFilterParam("published", strconv.FormatBool(value), FilterEquals))
	return p
}

// WithFilterArchived Фильтрация по архивным сущностям.
// archived=true
func (p *Params) WithFilterArchived(value bool) *Params {
	p.Filter = append(p.Filter, newFilterParam("archived", strconv.FormatBool(value), FilterEquals))
	return p
}

// WithGroupBy Группировка выдачи.
func (p *Params) WithGroupBy(value GroupByType) *Params {
	p.GroupBy = value
	return p
}

// GroupByType Тип группировки выдачи.
type GroupByType string

const (
	GroupByProduct     GroupByType = "product"     // Выдает только товары
	GroupByVariant     GroupByType = "variant"     // Выдает товары и модификации
	GroupByConsignment GroupByType = "consignment" // Выдает товары, модификации, серии
)

func (s GroupByType) String() string {
	return string(s)
}

// Листание.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/workbook/#workbook-fil-traciq-listanie-poisk-i-sortirowka-listanie

// WithLimit Количество элементов на странице.
// От 1 до 1000
func (p *Params) WithLimit(limit int) *Params {
	p.Limit = Clamp(limit, 1, MaxPositions)
	return p
}

// WithOffset Смещение от первого элемента.
func (p *Params) WithOffset(offset int) *Params {
	p.Offset = offset
	return p
}

// WithNamedFilter позволяет использовать сохранённый фильтр в качестве параметра.
func (p *Params) WithNamedFilter(filter *NamedFilter) *Params {
	if filter.Meta == nil {
		return p
	}
	if filter.Meta.Href == nil {
		return p
	}
	p.NamedFilter = *filter.Meta.Href
	return p
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
func (p *Params) WithOrder(fieldName string) *Params {
	p.Order = append(p.Order, newOrderParam(fieldName, OrderDirectionDefault))
	return p
}

// WithOrderAsc сортирует объекты по полю fieldName, по возрастанию (asc).
func (p *Params) WithOrderAsc(fieldName string) *Params {
	p.Order = append(p.Order, newOrderParam(fieldName, OrderDirectionAsc))
	return p
}

// WithOrderDesc сортирует объекты по полю fieldName, по убыванию (desc).
func (p *Params) WithOrderDesc(fieldName string) *Params {
	p.Order = append(p.Order, newOrderParam(fieldName, OrderDirectionDesc))
	return p
}

// WithSearch Контекстный поиск.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/workbook/#workbook-fil-traciq-listanie-poisk-i-sortirowka-poisk
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-kontextnyj-poisk
func (p *Params) WithSearch(search string) *Params {
	p.Search = search
	return p
}

// WithStockType Параметром stockType выбирается тип остатка, который необходимо рассчитать.
func (p *Params) WithStockType(stockType StockType) *Params {
	p.StockType = stockType
	return p
}

func (p *Params) QueryString() string {
	v, _ := query.Values(p)
	return v.Encode()

}

// StockType Параметр stockType.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-poluchit-kratkij-otchet-ob-ostatkah-parametr-stocktype
type StockType string

const (
	StockDefault   StockType = "stock"     // Физический остаток на складах, без учёта резерва и ожидания
	StockFreeStock StockType = "freeStock" // Остаток на складах за вычетом резерва
	StockQuantity  StockType = "quantity"  // Доступно. Учитывает резерв и ожидания
)

func (s StockType) String() string {
	return string(s)
}
