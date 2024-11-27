package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
)

// ReportByOperationsStock Отчет с остатками.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-po-dokumentam-nomenklatury-otchet-s-ostatkami
type ReportByOperationsStock struct {
	Assortment   MetaWrapper `json:"assortment"`   // Метаданные Товара/Модификации/Серии
	Operation    MetaWrapper `json:"operation"`    // Метаданные документа
	Store        MetaWrapper `json:"store"`        // Метаданные склада документа
	Moment       Timestamp   `json:"moment"`       // Дата документа
	AvgStockDays float64     `json:"avgStockDays"` // Количество дней на складе
	CostPerUnit  float64     `json:"costPerUnit"`  // Себестоимость за единицу
	Stock        float64     `json:"stock"`        // Остатки
	SumCost      float64     `json:"sumCost"`      // Сумма себестоимости
}

// ReportByOperationsReserve Отчет с резервами.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-po-dokumentam-nomenklatury-otchet-s-rezerwami
type ReportByOperationsReserve struct {
	Assortment MetaWrapper `json:"assortment"` // Метаданные Товара/Модификации/Серии
	Operation  MetaWrapper `json:"operation"`  // Метаданные документа
	Store      MetaWrapper `json:"store"`      // Метаданные склада документа
	Moment     Timestamp   `json:"moment"`     // Дата документа
	Reserve    float64     `json:"reserve"`    // Резерв
}

// ReportByOperationsTransit Отчет с ожиданием.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-po-dokumentam-nomenklatury-otchet-s-ozhidaniem
type ReportByOperationsTransit struct {
	Assortment MetaWrapper `json:"assortment"` // Метаданные Товара/Модификации/Серии
	Operation  MetaWrapper `json:"operation"`  // Метаданные документа
	Store      MetaWrapper `json:"store"`      // Метаданные склада документа
	Moment     Timestamp   `json:"moment"`     // Дата документа
	InTransit  float64     `json:"inTransit"`  // Ожидания
}

// ReportByOperationsService описывает методы сервиса для работы с отчётом по документам номенклатуры.
type ReportByOperationsService interface {
	// GetStock выполняет запрос на получение отчёта с остатками.
	// Принимает контекст и номенклатуру (товар/модификация/серия) и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetStock(ctx context.Context, assortment AssortmentConverter, params ...func(*Params)) (*List[ReportByOperationsStock], *resty.Response, error)

	// GetReserve выполняет запрос на получение отчёта с резервами.
	// Принимает контекст и номенклатуру (товар/модификация/серия) и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetReserve(ctx context.Context, assortment AssortmentConverter, params ...func(*Params)) (*List[ReportByOperationsReserve], *resty.Response, error)

	// GetTransit выполняет запрос на получение отчёта с ожиданием.
	// Принимает контекст и номенклатуру (товар/модификация/серия) и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetTransit(ctx context.Context, assortment AssortmentConverter, params ...func(*Params)) (*List[ReportByOperationsTransit], *resty.Response, error)
}

const (
	EndpointReportByOperationsStock   = EndpointReport + "byoperations/stock"
	EndpointReportByOperationsReserve = EndpointReport + "byoperations/reserve"
	EndpointReportByOperationsTransit = EndpointReport + "byoperations/intransit"
)

type reportByOperationsService struct {
	Endpoint
}

func makeParamByOperations(assortment AssortmentConverter, params []func(*Params)) []func(*Params) {
	return append(params, WithFilterEquals("assortment", assortment.AsAssortment().GetMeta().GetHref()))
}

func (service *reportByOperationsService) GetStock(ctx context.Context, assortment AssortmentConverter, params ...func(*Params)) (*List[ReportByOperationsStock], *resty.Response, error) {
	p := makeParamByOperations(assortment, params)
	return NewRequestBuilder[List[ReportByOperationsStock]](service.client, EndpointReportByOperationsStock).SetParams(p).Get(ctx)
}

func (service *reportByOperationsService) GetReserve(ctx context.Context, assortment AssortmentConverter, params ...func(*Params)) (*List[ReportByOperationsReserve], *resty.Response, error) {
	p := makeParamByOperations(assortment, params)
	return NewRequestBuilder[List[ReportByOperationsReserve]](service.client, EndpointReportByOperationsReserve).SetParams(p).Get(ctx)
}

func (service *reportByOperationsService) GetTransit(ctx context.Context, assortment AssortmentConverter, params ...func(*Params)) (*List[ReportByOperationsTransit], *resty.Response, error) {
	p := makeParamByOperations(assortment, params)
	return NewRequestBuilder[List[ReportByOperationsTransit]](service.client, EndpointReportByOperationsTransit).SetParams(p).Get(ctx)
}

// NewReportByOperationsService принимает [Client] и возвращает сервис для работы с отчётом по документам номенклатуры.
func NewReportByOperationsService(client *Client) ReportByOperationsService {
	return &reportByOperationsService{NewEndpoint(client, "")}
}
