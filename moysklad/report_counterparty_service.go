package moysklad

import (
	"context"
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
func (s *ReportCounterpartyService) GetList(ctx context.Context, params *Params) (*List[ReportCounterparty], *Response, error) {
	return NewRequestBuilder[List[ReportCounterparty]](s.Endpoint, ctx).WithParams(params).Get()
}

// GetListAsync Запрос на получение отчета по контрагентам (асинхронно)
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pokazateli-kontragentow-poluchit-pokazateli-kontragentow
func (s *ReportCounterpartyService) GetListAsync(ctx context.Context) (*AsyncResultService[List[ReportCounterparty]], *Response, error) {
	return NewRequestBuilder[List[ReportCounterparty]](s.Endpoint, ctx).Async()
}

// GetByCounterparties Пример запроса отчетов для нескольких контрагентов
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pokazateli-kontragentow-vyborochnye-pokazateli-kontragentow
func (s *ReportCounterpartyService) GetByCounterparties(ctx context.Context, data *CounterpartiesMeta) (*List[ReportCounterparty], *Response, error) {
	return NewRequestBuilder[List[ReportCounterparty]](s.Endpoint, ctx).WithBody(data).Post()
}

// GetByCounterpartyId Запрос на получение отчета по контрагенту с указанным id
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pokazateli-kontragentow-pokazateli-kontragenta
func (s *ReportCounterpartyService) GetByCounterpartyId(ctx context.Context, id *uuid.UUID) (*ReportCounterparty, *Response, error) {
	path := id.String()
	return NewRequestBuilder[ReportCounterparty](s.Endpoint, ctx).WithPath(path).Get()
}
