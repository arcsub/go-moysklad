package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
)

// Profit общие поля для структур отчёта "Прибыльность"
type Profit struct {
	Margin         float64 `json:"margin"`
	Profit         float64 `json:"profit"`
	ReturnAvgCheck float64 `json:"returnAvgCheck"`
	ReturnCostSum  float64 `json:"returnCostSum"`
	ReturnSum      float64 `json:"returnSum"`
	SalesAvgCheck  float64 `json:"salesAvgCheck"`
	SellCostSum    float64 `json:"sellCostSum"`
	SellSum        float64 `json:"sellSum"`
	ReturnCount    float64 `json:"returnCount"`
	SalesCount     float64 `json:"salesCount"`
}

// ProfitByAssortment Прибыльность по товарам
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-towaram
type ProfitByAssortment struct {
	Assortment     ProfitReportAssortment `json:"assortment"`
	SellCostSum    float64                `json:"sellCostSum"`
	Profit         float64                `json:"profit"`
	ReturnCost     float64                `json:"returnCost"`
	ReturnCostSum  float64                `json:"returnCostSum"`
	ReturnPrice    float64                `json:"returnPrice"`
	ReturnSum      float64                `json:"returnSum"`
	SellCost       float64                `json:"sellCost"`
	Margin         float64                `json:"margin"`
	SellPrice      float64                `json:"SellPrice"`
	SellSum        float64                `json:"sellSum"`
	ReturnQuantity float64                `json:"returnQuantity"`
	SellQuantity   float64                `json:"sellQuantity"`
}

// ProfitReportAssortment Структура объекта assortment
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-struktura-ob-ekta-assortment
type ProfitReportAssortment struct {
	Image Meta `json:"image"`
	MetaNameWrapper
	Uom     MetaNameWrapper `json:"uom,omitempty"`
	Code    string          `json:"code"`
	Article string          `json:"article"`
}

// ProfitByCounterparty Прибыльность по покупателям
// Ключевое слово: salesbyCounterparty
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-pokupatelqm
type ProfitByCounterparty struct {
	Counterparty MetaNameWrapper `json:"counterparty"`
	Profit
}

func (profitByCounterparty ProfitByCounterparty) MetaType() MetaType {
	return MetaTypeReportProfitByCounterparty
}

// ProfitByEmployee Прибыльность по сотрудникам
// Ключевое слово: salesbyemployee
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-sotrudnikam
type ProfitByEmployee struct {
	Employee MetaNameWrapper `json:"employee"`
	Profit
}

func (profitByEmployee ProfitByEmployee) MetaType() MetaType {
	return MetaTypeReportProfitByEmployee
}

// ProfitByProduct Прибыльность по товарам
// Ключевое слово: salesbyproduct
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-towaram
type ProfitByProduct struct {
	ProfitByAssortment
}

func (profitByProduct ProfitByProduct) MetaType() MetaType {
	return MetaTypeReportProfitByProduct
}

// ProfitBySalesChannel Прибыльность по каналам продаж
// Ключевое слово: salesbysaleschannel
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-kanalam-prodazh
type ProfitBySalesChannel struct {
	SalesChannel struct {
		Meta Meta             `json:"meta"`
		Name string           `json:"name"`
		Type SalesChannelType `json:"type"`
	} `json:"salesChannel"`
	Profit
}

func (profitBySalesChannel ProfitBySalesChannel) MetaType() MetaType {
	return MetaTypeReportProfitBySalesChannel
}

// ProfitByVariant Прибыльность по модификациям
// Ключевое слово: salesbyvariant
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-modifikaciqm
type ProfitByVariant struct {
	ProfitByAssortment
}

func (profitByVariant ProfitByVariant) MetaType() MetaType {
	return MetaTypeReportProfitByVariant
}

// ReportProfitService
// Сервис для работы с отчётом "Прибыльность".
type ReportProfitService interface {
	GetByProduct(ctx context.Context, params ...*Params) (*List[ProfitByProduct], *resty.Response, error)
	GetByVariant(ctx context.Context, params ...*Params) (*List[ProfitByVariant], *resty.Response, error)
	GetByEmployee(ctx context.Context, params ...*Params) (*List[ProfitByEmployee], *resty.Response, error)
	GetByCounterparty(ctx context.Context, params ...*Params) (*List[ProfitByCounterparty], *resty.Response, error)
	GetBySalesChannel(ctx context.Context, params ...*Params) (*List[ProfitBySalesChannel], *resty.Response, error)
	GetByProductAsync(ctx context.Context, params ...*Params) (AsyncResultService[List[ProfitByProduct]], *resty.Response, error)
	GetByVariantAsync(ctx context.Context, params ...*Params) (AsyncResultService[List[ProfitByVariant]], *resty.Response, error)
	GetByEmployeeAsync(ctx context.Context, params ...*Params) (AsyncResultService[List[ProfitByEmployee]], *resty.Response, error)
	GetByCounterpartyAsync(ctx context.Context, params ...*Params) (AsyncResultService[List[ProfitByCounterparty]], *resty.Response, error)
	GetBySalesChannelAsync(ctx context.Context, params ...*Params) (AsyncResultService[List[ProfitBySalesChannel]], *resty.Response, error)
}

type reportProfitService struct {
	Endpoint
}

func NewReportProfitService(client *Client) ReportProfitService {
	e := NewEndpoint(client, "report/profit")
	return &reportProfitService{e}
}

// GetByProduct  Запрос на получение отчета "Прибыльность по товарам".
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-towaram
func (service *reportProfitService) GetByProduct(ctx context.Context, params ...*Params) (*List[ProfitByProduct], *resty.Response, error) {
	path := "report/profit/byproduct"
	return NewRequestBuilder[List[ProfitByProduct]](service.client, path).SetParams(params...).Get(ctx)
}

// GetByVariant Запрос на получение отчета "Прибыльность по модификациям".
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-modifikaciqm
func (service *reportProfitService) GetByVariant(ctx context.Context, params ...*Params) (*List[ProfitByVariant], *resty.Response, error) {
	path := "report/profit/byvariant"
	return NewRequestBuilder[List[ProfitByVariant]](service.client, path).SetParams(params...).Get(ctx)
}

// GetByEmployee Запрос на получение отчета "Прибыльность по сотрудникам".
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-sotrudnikam
func (service *reportProfitService) GetByEmployee(ctx context.Context, params ...*Params) (*List[ProfitByEmployee], *resty.Response, error) {
	path := "report/profit/byemployee"
	return NewRequestBuilder[List[ProfitByEmployee]](service.client, path).SetParams(params...).Get(ctx)
}

// GetByCounterparty Запрос на получение отчета "Прибыльность по покупателям".
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-pokupatelqm
func (service *reportProfitService) GetByCounterparty(ctx context.Context, params ...*Params) (*List[ProfitByCounterparty], *resty.Response, error) {
	path := "report/profit/bycounterparty"
	return NewRequestBuilder[List[ProfitByCounterparty]](service.client, path).SetParams(params...).Get(ctx)
}

// GetBySalesChannel Запрос на получение отчета "Прибыльность по каналам продаж".
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-kanalam-prodazh
func (service *reportProfitService) GetBySalesChannel(ctx context.Context, params ...*Params) (*List[ProfitBySalesChannel], *resty.Response, error) {
	path := "report/profit/bysaleschannel"
	return NewRequestBuilder[List[ProfitBySalesChannel]](service.client, path).SetParams(params...).Get(ctx)
}

// GetByProductAsync Запрос на получение отчета "Прибыльность по товарам" (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-towaram
func (service *reportProfitService) GetByProductAsync(ctx context.Context, params ...*Params) (AsyncResultService[List[ProfitByProduct]], *resty.Response, error) {
	path := "report/profit/byproduct"
	return NewRequestBuilder[List[ProfitByProduct]](service.client, path).SetParams(params...).Async(ctx)
}

// GetByVariantAsync Запрос на получение отчета "Прибыльность по модификациям" (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-modifikaciqm
func (service *reportProfitService) GetByVariantAsync(ctx context.Context, params ...*Params) (AsyncResultService[List[ProfitByVariant]], *resty.Response, error) {
	path := "report/profit/byvariant"
	return NewRequestBuilder[List[ProfitByVariant]](service.client, path).SetParams(params...).Async(ctx)
}

// GetByEmployeeAsync Запрос на получение отчета "Прибыльность по сотрудникам" (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-sotrudnikam
func (service *reportProfitService) GetByEmployeeAsync(ctx context.Context, params ...*Params) (AsyncResultService[List[ProfitByEmployee]], *resty.Response, error) {
	path := "report/profit/byemployee"
	return NewRequestBuilder[List[ProfitByEmployee]](service.client, path).SetParams(params...).Async(ctx)
}

// GetByCounterpartyAsync Запрос на получение отчета "Прибыльность по покупателям" (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-pokupatelqm
func (service *reportProfitService) GetByCounterpartyAsync(ctx context.Context, params ...*Params) (AsyncResultService[List[ProfitByCounterparty]], *resty.Response, error) {
	path := "report/profit/bycounterparty"
	return NewRequestBuilder[List[ProfitByCounterparty]](service.client, path).SetParams(params...).Async(ctx)
}

// GetBySalesChannelAsync Запрос на получение отчета "Прибыльность по каналам продаж" (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-kanalam-prodazh
func (service *reportProfitService) GetBySalesChannelAsync(ctx context.Context, params ...*Params) (AsyncResultService[List[ProfitBySalesChannel]], *resty.Response, error) {
	path := "report/profit/bysaleschannel"
	return NewRequestBuilder[List[ProfitBySalesChannel]](service.client, path).SetParams(params...).Async(ctx)
}
