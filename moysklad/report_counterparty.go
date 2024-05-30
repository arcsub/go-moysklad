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
	FirstDemandDate Timestamp        `json:"firstDemandDate"`
	Updated         Timestamp        `json:"updated"`
	LastEventDate   Timestamp        `json:"lastEventDate"`
	LastDemandDate  Timestamp        `json:"lastDemandDate"`
	Counterparty    CounterpartyData `json:"counterparty"`
	Meta            Meta             `json:"meta"`
	DiscountsSum    Decimal          `json:"discountsSum"`
	LastEventText   string           `json:"lastEventText"`
	DemandsSum      Decimal          `json:"demandsSum"`
	AverageReceipt  Decimal          `json:"averageReceipt"`
	BonusBalance    Decimal          `json:"bonusBalance"`
	Profit          Decimal          `json:"profit"`
	ReturnsSum      Decimal          `json:"returnsSum"`
	Balance         Decimal          `json:"balance"`
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

func (r ReportCounterparty) MetaType() MetaType {
	return MetaTypeReportCounterparty
}

type CounterpartyElement struct {
	Counterparty MetaWrapper `json:"counterparty"`
}

type CounterpartiesMeta struct {
	Counterparties Slice[CounterpartyElement] `json:"counterparties"`
}

func (c *CounterpartiesMeta) Push(elements ...*Counterparty) {
	for _, element := range elements {
		ce := &CounterpartyElement{
			Counterparty: MetaWrapper{
				Meta: *element.GetMeta(),
			},
		}
		c.Counterparties = append(c.Counterparties, ce)
	}
}

// ReportCounterpartyService
// Сервис для работы с показателями контрагентов.
type ReportCounterpartyService interface {
	GetList(ctx context.Context, params *Params) (*List[ReportCounterparty], *resty.Response, error)
	GetListAsync(ctx context.Context) (AsyncResultService[List[ReportCounterparty]], *resty.Response, error)
	GetByCounterparties(ctx context.Context, data *CounterpartiesMeta) (*List[ReportCounterparty], *resty.Response, error)
	GetByCounterpartyId(ctx context.Context, id *uuid.UUID) (*ReportCounterparty, *resty.Response, error)
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
func (s *reportCounterpartyService) GetList(ctx context.Context, params *Params) (*List[ReportCounterparty], *resty.Response, error) {
	return NewRequestBuilder[List[ReportCounterparty]](s.client, s.uri).SetParams(params).Get(ctx)
}

// GetListAsync Запрос на получение отчета по контрагентам (асинхронно)
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pokazateli-kontragentow-poluchit-pokazateli-kontragentow
func (s *reportCounterpartyService) GetListAsync(ctx context.Context) (AsyncResultService[List[ReportCounterparty]], *resty.Response, error) {
	return NewRequestBuilder[List[ReportCounterparty]](s.client, s.uri).Async(ctx)
}

// GetByCounterparties Пример запроса отчетов для нескольких контрагентов
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pokazateli-kontragentow-vyborochnye-pokazateli-kontragentow
func (s *reportCounterpartyService) GetByCounterparties(ctx context.Context, data *CounterpartiesMeta) (*List[ReportCounterparty], *resty.Response, error) {
	return NewRequestBuilder[List[ReportCounterparty]](s.client, s.uri).Post(ctx, data)
}

// GetByCounterpartyId Запрос на получение отчета по контрагенту с указанным id
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pokazateli-kontragentow-pokazateli-kontragenta
func (s *reportCounterpartyService) GetByCounterpartyId(ctx context.Context, id *uuid.UUID) (*ReportCounterparty, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s", s.uri, id)
	return NewRequestBuilder[ReportCounterparty](s.client, path).Get(ctx)
}
