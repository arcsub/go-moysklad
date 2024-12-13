package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
)

// ProfitByAssortment Прибыльность по товарам
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-towaram
type ProfitByAssortment struct {
	Assortment     ReportProfitAssortment `json:"assortment"`     // Краткое представление Модификации, Услуги или Комплекта в отчете
	SellCostSum    float64                `json:"sellCostSum"`    // Сумма себестоимостей продаж в копейках
	Profit         float64                `json:"profit"`         // Прибыль
	ReturnCost     float64                `json:"returnCost"`     // Себестоимость возвратов в копейках
	ReturnCostSum  float64                `json:"returnCostSum"`  // Сумма себестоимостей возвратов в копейках
	ReturnPrice    float64                `json:"returnPrice"`    // Цена возвратов
	ReturnSum      float64                `json:"returnSum"`      // Сумма возвратов
	SellCost       float64                `json:"sellCost"`       // Себестоимость в копейках
	Margin         float64                `json:"margin"`         // Рентабельность
	SalesMargin    float64                `json:"salesMargin"`    // Рентабельность продаж
	SellPrice      float64                `json:"sellPrice"`      // Цена продаж (средняя)
	SellSum        float64                `json:"sellSum"`        // Сумма продаж
	ReturnQuantity float64                `json:"returnQuantity"` // Возвращенное количество
	SellQuantity   float64                `json:"sellQuantity"`   // Проданное количество
}

// ReportProfitAssortment Структура объекта assortment
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-struktura-ob-ekta-assortment
type ReportProfitAssortment struct {
	Image   Meta            `json:"image"`         // Изображение товара
	Meta    Meta            `json:"meta"`          // Метаданные Товара или Услуги
	Name    string          `json:"name"`          // Наименование Товара или Услуги
	Uom     MetaNameWrapper `json:"uom,omitempty"` // Единица измерения
	Code    string          `json:"code"`          // Код товара или услуги
	Article string          `json:"article"`       // Артикул товара
}

type ReportProfitSalesChannel struct {
	Meta Meta             `json:"meta"` // Метаданные Канала продаж
	Name string           `json:"name"` // Наименование Канала продаж
	Type SalesChannelType `json:"type"` // Тип Канала продаж
}

// ProfitByCounterparty Прибыльность по покупателям
//
// Код сущности: salesbyCounterparty
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-pokupatelqm
type ProfitByCounterparty struct {
	Counterparty   MetaNameWrapper `json:"counterparty"`
	Margin         float64         `json:"margin"`         // Рентабельность
	Profit         float64         `json:"profit"`         // Прибыль
	ReturnAvgCheck float64         `json:"returnAvgCheck"` // Средний чек возврата
	ReturnCostSum  float64         `json:"returnCostSum"`  // Сумма себестоимостей возвратов в копейках
	ReturnSum      float64         `json:"returnSum"`      // Сумма возвратов
	SalesAvgCheck  float64         `json:"salesAvgCheck"`  // Средний чек продаж
	SellCostSum    float64         `json:"sellCostSum"`    // Сумма себестоимостей продаж в копейках
	SellSum        float64         `json:"sellSum"`        // Сумма продаж
	ReturnCount    float64         `json:"returnCount"`    // Количество возвратов
	SalesCount     float64         `json:"salesCount"`     // Количество продаж
	SalesMargin    float64         `json:"salesMargin"`    // Рентабельность продаж
}

// MetaType возвращает код сущности.
func (ProfitByCounterparty) MetaType() MetaType {
	return MetaTypeReportProfitByCounterparty
}

// ProfitByEmployee Прибыльность по сотрудникам
//
// Код сущности: salesbyemployee
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-sotrudnikam
type ProfitByEmployee struct {
	Employee       MetaNameWrapper `json:"employee"`       // Краткое представление Сотрудника в отчете
	Margin         float64         `json:"margin"`         // Рентабельность
	Profit         float64         `json:"profit"`         // Прибыль
	ReturnAvgCheck float64         `json:"returnAvgCheck"` // Средний чек возврата
	ReturnCostSum  float64         `json:"returnCostSum"`  // Сумма себестоимостей возвратов в копейках
	ReturnSum      float64         `json:"returnSum"`      // Сумма возвратов
	SalesAvgCheck  float64         `json:"salesAvgCheck"`  // Средний чек продаж
	SellCostSum    float64         `json:"sellCostSum"`    // Сумма себестоимостей продаж в копейках
	SellSum        float64         `json:"sellSum"`        // Сумма продаж
	ReturnCount    float64         `json:"returnCount"`    // Количество возвратов
	SalesCount     float64         `json:"salesCount"`     // Количество продаж
	SalesMargin    float64         `json:"salesMargin"`    // Рентабельность продаж
}

// MetaType возвращает код сущности.
func (ProfitByEmployee) MetaType() MetaType {
	return MetaTypeReportProfitByEmployee
}

// ProfitByProduct Прибыльность по товарам
//
// Код сущности: salesbyproduct
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-towaram
type ProfitByProduct struct {
	ProfitByAssortment
}

// MetaType возвращает код сущности.
func (ProfitByProduct) MetaType() MetaType {
	return MetaTypeReportProfitByProduct
}

// ProfitByVariant Прибыльность по модификациям
//
// Код сущности: salesbyvariant
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-modifikaciqm
type ProfitByVariant struct {
	ProfitByAssortment
}

// MetaType возвращает код сущности.
func (ProfitByVariant) MetaType() MetaType {
	return MetaTypeReportProfitByVariant
}

// ProfitBySalesChannel Прибыльность по каналам продаж
//
// Код сущности: salesbysaleschannel
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-kanalam-prodazh
type ProfitBySalesChannel struct {
	SalesChannel   ReportProfitSalesChannel `json:"salesChannel"`   // Краткое представление Канала продаж в отчете
	Margin         float64                  `json:"margin"`         // Рентабельность
	Profit         float64                  `json:"profit"`         // Прибыль
	ReturnAvgCheck float64                  `json:"returnAvgCheck"` // Средний чек возврата
	ReturnCostSum  float64                  `json:"returnCostSum"`  // Сумма себестоимостей возвратов в копейках
	ReturnSum      float64                  `json:"returnSum"`      // Сумма возвратов
	SalesAvgCheck  float64                  `json:"salesAvgCheck"`  // Средний чек продаж
	SellCostSum    float64                  `json:"sellCostSum"`    // Сумма себестоимостей продаж в копейках
	SellSum        float64                  `json:"sellSum"`        // Сумма продаж
	ReturnCount    float64                  `json:"returnCount"`    // Количество возвратов
	SalesCount     float64                  `json:"salesCount"`     // Количество продаж
	SalesMargin    float64                  `json:"salesMargin"`    // Рентабельность продаж
}

// MetaType возвращает код сущности.
func (ProfitBySalesChannel) MetaType() MetaType {
	return MetaTypeReportProfitBySalesChannel
}

// ReportProfitService описывает методы сервиса для работы с отчётом Прибыльность.
type ReportProfitService interface {
	// GetByProduct выполняет запрос на получение отчёта "Прибыльность по товарам".
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetByProduct(ctx context.Context, params ...func(*Params)) (*List[ProfitByProduct], *resty.Response, error)

	// GetByVariant выполняет запрос на получение отчёта "Прибыльность по модификациям".
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetByVariant(ctx context.Context, params ...func(*Params)) (*List[ProfitByVariant], *resty.Response, error)

	// GetByEmployee выполняет запрос на получение отчёта "Прибыльность по сотрудникам".
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetByEmployee(ctx context.Context, params ...func(*Params)) (*List[ProfitByEmployee], *resty.Response, error)

	// GetByCounterparty выполняет запрос на получение отчёта "Прибыльность по покупателям".
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetByCounterparty(ctx context.Context, params ...func(*Params)) (*List[ProfitByCounterparty], *resty.Response, error)

	// GetBySalesChannel выполняет запрос на получение отчёта "Прибыльность по каналам продаж".
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetBySalesChannel(ctx context.Context, params ...func(*Params)) (*List[ProfitBySalesChannel], *resty.Response, error)

	// GetByProductAsync выполняет запрос на получение отчёта "Прибыльность по товарам" (асинхронно).
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает сервис для работы с контекстом асинхронного запроса.
	GetByProductAsync(ctx context.Context, params ...func(*Params)) (AsyncResultService[List[ProfitByProduct]], *resty.Response, error)

	// GetByVariantAsync выполняет запрос на получение отчёта "Прибыльность по модификациям".
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает сервис для работы с контекстом асинхронного запроса.
	GetByVariantAsync(ctx context.Context, params ...func(*Params)) (AsyncResultService[List[ProfitByVariant]], *resty.Response, error)

	// GetByEmployeeAsync выполняет запрос на получение отчёта "Прибыльность по сотрудникам".
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает сервис для работы с контекстом асинхронного запроса.
	GetByEmployeeAsync(ctx context.Context, params ...func(*Params)) (AsyncResultService[List[ProfitByEmployee]], *resty.Response, error)

	// GetByCounterpartyAsync выполняет запрос на получение отчёта "Прибыльность по покупателям".
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает сервис для работы с контекстом асинхронного запроса.
	GetByCounterpartyAsync(ctx context.Context, params ...func(*Params)) (AsyncResultService[List[ProfitByCounterparty]], *resty.Response, error)

	// GetBySalesChannelAsync выполняет запрос на получение отчёта "Прибыльность по каналам продаж".
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает сервис для работы с контекстом асинхронного запроса.
	GetBySalesChannelAsync(ctx context.Context, params ...func(*Params)) (AsyncResultService[List[ProfitBySalesChannel]], *resty.Response, error)
}

const (
	EndpointReportProfit               = EndpointReport + "profit"
	EndpointReportProfitByProduct      = EndpointReportProfit + "/byproduct"
	EndpointReportProfitByVariant      = EndpointReportProfit + "/byvariant"
	EndpointReportProfitByEmployee     = EndpointReportProfit + "/byemployee"
	EndpointReportProfitByCounterparty = EndpointReportProfit + "/bycounterparty"
	EndpointReportProfitBySalesChannel = EndpointReportProfit + "/bysaleschannel"
)

type reportProfitService struct {
	Endpoint
}

func (service *reportProfitService) GetByProduct(ctx context.Context, params ...func(*Params)) (*List[ProfitByProduct], *resty.Response, error) {
	return NewRequestBuilder[List[ProfitByProduct]](service.client, EndpointReportProfitByProduct).SetParams(params).Get(ctx)
}

func (service *reportProfitService) GetByVariant(ctx context.Context, params ...func(*Params)) (*List[ProfitByVariant], *resty.Response, error) {
	return NewRequestBuilder[List[ProfitByVariant]](service.client, EndpointReportProfitByVariant).SetParams(params).Get(ctx)
}

func (service *reportProfitService) GetByEmployee(ctx context.Context, params ...func(*Params)) (*List[ProfitByEmployee], *resty.Response, error) {
	return NewRequestBuilder[List[ProfitByEmployee]](service.client, EndpointReportProfitByEmployee).SetParams(params).Get(ctx)
}

func (service *reportProfitService) GetByCounterparty(ctx context.Context, params ...func(*Params)) (*List[ProfitByCounterparty], *resty.Response, error) {
	return NewRequestBuilder[List[ProfitByCounterparty]](service.client, EndpointReportProfitByCounterparty).SetParams(params).Get(ctx)
}

func (service *reportProfitService) GetBySalesChannel(ctx context.Context, params ...func(*Params)) (*List[ProfitBySalesChannel], *resty.Response, error) {
	return NewRequestBuilder[List[ProfitBySalesChannel]](service.client, EndpointReportProfitBySalesChannel).SetParams(params).Get(ctx)
}

func (service *reportProfitService) GetByProductAsync(ctx context.Context, params ...func(*Params)) (AsyncResultService[List[ProfitByProduct]], *resty.Response, error) {
	return NewRequestBuilder[List[ProfitByProduct]](service.client, EndpointReportProfitByProduct).SetParams(params).Async(ctx)
}

func (service *reportProfitService) GetByVariantAsync(ctx context.Context, params ...func(*Params)) (AsyncResultService[List[ProfitByVariant]], *resty.Response, error) {
	return NewRequestBuilder[List[ProfitByVariant]](service.client, EndpointReportProfitByVariant).SetParams(params).Async(ctx)
}

func (service *reportProfitService) GetByEmployeeAsync(ctx context.Context, params ...func(*Params)) (AsyncResultService[List[ProfitByEmployee]], *resty.Response, error) {
	return NewRequestBuilder[List[ProfitByEmployee]](service.client, EndpointReportProfitByEmployee).SetParams(params).Async(ctx)
}

func (service *reportProfitService) GetByCounterpartyAsync(ctx context.Context, params ...func(*Params)) (AsyncResultService[List[ProfitByCounterparty]], *resty.Response, error) {
	return NewRequestBuilder[List[ProfitByCounterparty]](service.client, EndpointReportProfitByCounterparty).SetParams(params).Async(ctx)
}

func (service *reportProfitService) GetBySalesChannelAsync(ctx context.Context, params ...func(*Params)) (AsyncResultService[List[ProfitBySalesChannel]], *resty.Response, error) {
	return NewRequestBuilder[List[ProfitBySalesChannel]](service.client, EndpointReportProfitBySalesChannel).SetParams(params).Async(ctx)
}

// NewReportProfitService принимает [Client] и возвращает сервис для работы с отчётом Прибыльность.
func NewReportProfitService(client *Client) ReportProfitService {
	return &reportProfitService{NewEndpoint(client, EndpointReportProfit)}
}
