package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// ReportCounterpartyService
// Сервис для работы с показателями контрагентов.
type ReportCounterpartyService struct {
	Endpoint
}

func NewReportCounterpartyService(client *Client) *ReportCounterpartyService {
	e := NewEndpoint(client, "report/counterparty")
	return &ReportCounterpartyService{e}
}

// GetList Запрос на получение отчета по контрагентам
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pokazateli-kontragentow-poluchit-pokazateli-kontragentow
func (s *ReportCounterpartyService) GetList(ctx context.Context, params *Params) (*List[ReportCounterparty], *resty.Response, error) {
	return NewRequestBuilder[List[ReportCounterparty]](s.client, s.uri).SetParams(params).Get(ctx)
}

// GetListAsync Запрос на получение отчета по контрагентам (асинхронно)
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pokazateli-kontragentow-poluchit-pokazateli-kontragentow
func (s *ReportCounterpartyService) GetListAsync(ctx context.Context) (*AsyncResultService[List[ReportCounterparty]], *resty.Response, error) {
	return NewRequestBuilder[List[ReportCounterparty]](s.client, s.uri).Async(ctx)
}

// GetByCounterparties Пример запроса отчетов для нескольких контрагентов
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pokazateli-kontragentow-vyborochnye-pokazateli-kontragentow
func (s *ReportCounterpartyService) GetByCounterparties(ctx context.Context, data *CounterpartiesMeta) (*List[ReportCounterparty], *resty.Response, error) {
	return NewRequestBuilder[List[ReportCounterparty]](s.client, s.uri).Post(ctx, data)
}

// GetByCounterpartyId Запрос на получение отчета по контрагенту с указанным id
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pokazateli-kontragentow-pokazateli-kontragenta
func (s *ReportCounterpartyService) GetByCounterpartyId(ctx context.Context, id *uuid.UUID) (*ReportCounterparty, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s", s.uri, id)
	return NewRequestBuilder[ReportCounterparty](s.client, path).Get(ctx)
}
