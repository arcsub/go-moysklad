package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// ReportCounterparty Показатели контрагентов.
// Ключевое слово: counterparty
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pokazateli-kontragentow-pokazateli-kontragentow
type ReportCounterparty struct {
	Updated         Timestamp        `json:"updated"`
	LastEventDate   Timestamp        `json:"lastEventDate"`
	LastDemandDate  Timestamp        `json:"lastDemandDate"`
	FirstDemandDate Timestamp        `json:"firstDemandDate"`
	Counterparty    CounterpartyData `json:"counterparty"`
	Meta            Meta             `json:"meta"`
	LastEventText   string           `json:"lastEventText"`
	DemandsSum      float64          `json:"demandsSum"`
	DiscountsSum    float64          `json:"discountsSum"`
	AverageReceipt  float64          `json:"averageReceipt"`
	BonusBalance    float64          `json:"bonusBalance"`
	Profit          float64          `json:"profit"`
	ReturnsSum      float64          `json:"returnsSum"`
	Balance         float64          `json:"balance"`
	DemandsCount    int              `json:"demandsCount"`
	ReturnsCount    int              `json:"returnsCount"`
}

// CounterpartyData Контрагент
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pokazateli-kontragentow-dopolnitel-nye-atributy-dostupnye-dlq-fil-tracii-kontragent
type CounterpartyData struct {
	CompanyType  CompanyType `json:"companyType"`  // Тип контрагента
	ExternalCode string      `json:"externalCode"` // Внешний код контрагента
	ID           string      `json:"id"`           // ID Контрагента
	Meta         Meta        `json:"meta"`         // Метаданные Контрагента
	Name         string      `json:"name"`         // Наименование Контрагента
}

// MetaType возвращает тип сущности.
func (ReportCounterparty) MetaType() MetaType {
	return MetaTypeReportCounterparty
}

type CounterpartyElement struct {
	Counterparty MetaWrapper `json:"counterparty"`
}

type CounterpartiesMeta struct {
	Counterparties Slice[CounterpartyElement] `json:"counterparties"`
}

func (counterpartiesMeta *CounterpartiesMeta) Push(elements ...*Counterparty) {
	for _, element := range elements {
		counterpartiesMeta.Counterparties.Push(&CounterpartyElement{element.GetMeta().Wrap()})
	}
}

// ReportCounterpartyService
// Сервис для работы с показателями контрагентов.
type ReportCounterpartyService interface {
	GetList(ctx context.Context, params ...*Params) (*List[ReportCounterparty], *resty.Response, error)
	GetListAsync(ctx context.Context) (AsyncResultService[List[ReportCounterparty]], *resty.Response, error)
	GetByCounterparties(ctx context.Context, data *CounterpartiesMeta) (*List[ReportCounterparty], *resty.Response, error)
	GetByCounterpartyID(ctx context.Context, id uuid.UUID) (*ReportCounterparty, *resty.Response, error)
}

type reportCounterpartyService struct {
	Endpoint
}

func NewReportCounterpartyService(client *Client) ReportCounterpartyService {
	e := NewEndpoint(client, "report/counterparty")
	return &reportCounterpartyService{e}
}

// GetList Запрос на получение отчета по контрагентам
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pokazateli-kontragentow-poluchit-pokazateli-kontragentow
func (service *reportCounterpartyService) GetList(ctx context.Context, params ...*Params) (*List[ReportCounterparty], *resty.Response, error) {
	return NewRequestBuilder[List[ReportCounterparty]](service.client, service.uri).SetParams(params...).Get(ctx)
}

// GetListAsync Запрос на получение отчета по контрагентам (асинхронно)
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pokazateli-kontragentow-poluchit-pokazateli-kontragentow
func (service *reportCounterpartyService) GetListAsync(ctx context.Context) (AsyncResultService[List[ReportCounterparty]], *resty.Response, error) {
	return NewRequestBuilder[List[ReportCounterparty]](service.client, service.uri).Async(ctx)
}

// GetByCounterparties Пример запроса отчетов для нескольких контрагентов
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pokazateli-kontragentow-vyborochnye-pokazateli-kontragentow
func (service *reportCounterpartyService) GetByCounterparties(ctx context.Context, data *CounterpartiesMeta) (*List[ReportCounterparty], *resty.Response, error) {
	return NewRequestBuilder[List[ReportCounterparty]](service.client, service.uri).Post(ctx, data)
}

// GetByCounterpartyID Запрос на получение отчета по контрагенту с указанным id
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pokazateli-kontragentow-pokazateli-kontragenta
func (service *reportCounterpartyService) GetByCounterpartyID(ctx context.Context, id uuid.UUID) (*ReportCounterparty, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s", service.uri, id)
	return NewRequestBuilder[ReportCounterparty](service.client, path).Get(ctx)
}
