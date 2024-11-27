package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
)

// ReportCounterparty Показатели контрагентов.
//
// Код сущности: counterparty
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pokazateli-kontragentow-pokazateli-kontragentow
type ReportCounterparty struct {
	Updated         Timestamp              `json:"updated"`         // Момент последнего изменения контрагента
	LastEventDate   Timestamp              `json:"lastEventDate"`   // Дата последнего события
	LastDemandDate  Timestamp              `json:"lastDemandDate"`  // Дата последней продажи
	FirstDemandDate Timestamp              `json:"firstDemandDate"` // Дата первой продажи
	Counterparty    ReportCounterpartyInfo `json:"counterparty"`    // Контрагент
	Meta            Meta                   `json:"meta"`            // Метаданные Отчета по данному контрагенту
	LastEventText   string                 `json:"lastEventText"`   // Текст последнего события
	DemandsSum      float64                `json:"demandsSum"`      // Сумма продаж
	DiscountsSum    float64                `json:"discountsSum"`    // Сумма скидок
	AverageReceipt  float64                `json:"averageReceipt"`  // Средний чек
	BonusBalance    float64                `json:"bonusBalance"`    // Баллы
	Profit          float64                `json:"profit"`          // Прибыль
	ReturnsSum      float64                `json:"returnsSum"`      // Сумма возвратов
	Balance         float64                `json:"balance"`         // Баланс
	DemandsCount    int                    `json:"demandsCount"`    // Количество продаж
	ReturnsCount    int                    `json:"returnsCount"`    // Количество возвратов
}

// ReportCounterpartyInfo Краткая информация о контрагенте.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pokazateli-kontragentow-dopolnitel-nye-atributy-dostupnye-dlq-fil-tracii-kontragent
type ReportCounterpartyInfo struct {
	CompanyType  CompanyType `json:"companyType"`  // Тип контрагента
	ExternalCode string      `json:"externalCode"` // Внешний код контрагента
	ID           string      `json:"id"`           // ID Контрагента
	Meta         Meta        `json:"meta"`         // Метаданные Контрагента
	Name         string      `json:"name"`         // Наименование Контрагента
	INN          string      `json:"inn"`          // ИНН
}

// MetaType возвращает код сущности.
func (ReportCounterparty) MetaType() MetaType {
	return MetaTypeReportCounterparty
}

type CounterpartyOwner struct {
	Counterparty MetaWrapper `json:"counterparty"`
}

type CounterpartiesMeta struct {
	Counterparties Slice[CounterpartyOwner] `json:"counterparties"`
}

// ReportCounterpartyService описывает методы сервиса для работы с показателями контрагентов.
type ReportCounterpartyService interface {
	// GetList выполняет запрос на получение отчёта по контрагентам.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[ReportCounterparty], *resty.Response, error)

	// GetListAsync выполняет запрос на получение отчёта по контрагентам (асинхронно).
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает сервис для работы с контекстом асинхронного запроса.
	GetListAsync(ctx context.Context) (AsyncResultService[List[ReportCounterparty]], *resty.Response, error)

	// GetByCounterparties выполняет запрос на получение выборочных показателей контрагентов.
	// Принимает контекст и множество контрагентов.
	// Возвращает объект List.
	GetByCounterparties(ctx context.Context, counterparties ...*Counterparty) (*List[ReportCounterparty], *resty.Response, error)

	// GetByCounterpartyID выполняет запрос на получение отчёта по контрагенту с указанным ID.
	// Принимает контекст и ID контрагента.
	// Возвращает отчёт по конкретному контрагенту.
	GetByCounterpartyID(ctx context.Context, id string) (*ReportCounterparty, *resty.Response, error)
}

const (
	EndpointReportCounterparty = EndpointReport + string(MetaTypeCounterparty)
)

type reportCounterpartyService struct {
	Endpoint
}

func (service *reportCounterpartyService) GetList(ctx context.Context, params ...*Params) (*List[ReportCounterparty], *resty.Response, error) {
	return NewRequestBuilder[List[ReportCounterparty]](service.client, service.uri).SetParams(params...).Get(ctx)
}

func (service *reportCounterpartyService) GetListAsync(ctx context.Context) (AsyncResultService[List[ReportCounterparty]], *resty.Response, error) {
	return NewRequestBuilder[List[ReportCounterparty]](service.client, service.uri).Async(ctx)
}

func (service *reportCounterpartyService) GetByCounterparties(ctx context.Context, counterparties ...*Counterparty) (*List[ReportCounterparty], *resty.Response, error) {
	var data CounterpartiesMeta
	for _, element := range counterparties {
		data.Counterparties.Push(&CounterpartyOwner{element.GetMeta().Wrap()})
	}
	return NewRequestBuilder[List[ReportCounterparty]](service.client, service.uri).Post(ctx, data)
}

func (service *reportCounterpartyService) GetByCounterpartyID(ctx context.Context, id string) (*ReportCounterparty, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s", service.uri, id)
	return NewRequestBuilder[ReportCounterparty](service.client, path).Get(ctx)
}

// NewReportCounterpartyService принимает [Client] и возвращает сервис для работы с показателями контрагентов.
func NewReportCounterpartyService(client *Client) ReportCounterpartyService {
	return &reportCounterpartyService{NewEndpoint(client, EndpointReportCounterparty)}
}
