package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
)

// ReportProfitService
// Сервис для работы с отчётом "Прибыльность".
type ReportProfitService struct {
	Endpoint
}

func NewReportProfitService(client *Client) *ReportProfitService {
	e := NewEndpoint(client, "report/profit")
	return &ReportProfitService{e}
}

// GetByProduct  Запрос на получение отчета "Прибыльность по товарам".
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-towaram
func (s *ReportProfitService) GetByProduct(ctx context.Context) (*List[ProfitByProduct], *resty.Response, error) {
	path := "report/profit/byproduct"
	return NewRequestBuilder[List[ProfitByProduct]](s.client, path).Get(ctx)
}

// GetByVariant Запрос на получение отчета "Прибыльность по модификациям".
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-modifikaciqm
func (s *ReportProfitService) GetByVariant(ctx context.Context) (*List[ProfitByVariant], *resty.Response, error) {
	path := "report/profit/byvariant"
	return NewRequestBuilder[List[ProfitByVariant]](s.client, path).Get(ctx)
}

// GetByEmployee Запрос на получение отчета "Прибыльность по сотрудникам".
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-sotrudnikam
func (s *ReportProfitService) GetByEmployee(ctx context.Context) (*List[ProfitByEmployee], *resty.Response, error) {
	path := "report/profit/byemployee"
	return NewRequestBuilder[List[ProfitByEmployee]](s.client, path).Get(ctx)
}

// GetByCounterparty Запрос на получение отчета "Прибыльность по покупателям".
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-pokupatelqm
func (s *ReportProfitService) GetByCounterparty(ctx context.Context) (*List[ProfitByCounterparty], *resty.Response, error) {
	path := "report/profit/bycounterparty"
	return NewRequestBuilder[List[ProfitByCounterparty]](s.client, path).Get(ctx)
}

// GetBySalesChannel Запрос на получение отчета "Прибыльность по каналам продаж".
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-kanalam-prodazh
func (s *ReportProfitService) GetBySalesChannel(ctx context.Context) (*List[ProfitBySalesChannel], *resty.Response, error) {
	path := "report/profit/bysaleschannel"
	return NewRequestBuilder[List[ProfitBySalesChannel]](s.client, path).Get(ctx)
}

// GetByProductAsync Запрос на получение отчета "Прибыльность по товарам" (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-towaram
func (s *ReportProfitService) GetByProductAsync(ctx context.Context) (*AsyncResultService[List[ProfitByProduct]], *resty.Response, error) {
	path := "report/profit/byproduct"
	return NewRequestBuilder[List[ProfitByProduct]](s.client, path).Async(ctx)
}

// GetByVariantAsync Запрос на получение отчета "Прибыльность по модификациям" (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-modifikaciqm
func (s *ReportProfitService) GetByVariantAsync(ctx context.Context) (*AsyncResultService[List[ProfitByVariant]], *resty.Response, error) {
	path := "report/profit/byvariant"
	return NewRequestBuilder[List[ProfitByVariant]](s.client, path).Async(ctx)
}

// GetByEmployeeAsync Запрос на получение отчета "Прибыльность по сотрудникам" (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-sotrudnikam
func (s *ReportProfitService) GetByEmployeeAsync(ctx context.Context) (*AsyncResultService[List[ProfitByEmployee]], *resty.Response, error) {
	path := "report/profit/byemployee"
	return NewRequestBuilder[List[ProfitByEmployee]](s.client, path).Async(ctx)
}

// GetByCounterpartyAsync Запрос на получение отчета "Прибыльность по покупателям" (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-pokupatelqm
func (s *ReportProfitService) GetByCounterpartyAsync(ctx context.Context) (*AsyncResultService[List[ProfitByCounterparty]], *resty.Response, error) {
	path := "report/profit/bycounterparty"
	return NewRequestBuilder[List[ProfitByCounterparty]](s.client, path).Async(ctx)
}

// GetBySalesChannelAsync Запрос на получение отчета "Прибыльность по каналам продаж" (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-kanalam-prodazh
func (s *ReportProfitService) GetBySalesChannelAsync(ctx context.Context) (*AsyncResultService[List[ProfitBySalesChannel]], *resty.Response, error) {
	path := "report/profit/bysaleschannel"
	return NewRequestBuilder[List[ProfitBySalesChannel]](s.client, path).Async(ctx)
}
