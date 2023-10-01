package moysklad

import "context"

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
func (s *ReportProfitService) GetByProduct(ctx context.Context) (*List[ProfitByProduct], *Response, error) {
	path := "byproduct"
	return NewRequestBuilder[List[ProfitByProduct]](s.Endpoint, ctx).WithPath(path).Get()
}

// GetByVariant Запрос на получение отчета "Прибыльность по модификациям".
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-modifikaciqm
func (s *ReportProfitService) GetByVariant(ctx context.Context) (*List[ProfitByVariant], *Response, error) {
	path := "byvariant"
	return NewRequestBuilder[List[ProfitByVariant]](s.Endpoint, ctx).WithPath(path).Get()
}

// GetByEmployee Запрос на получение отчета "Прибыльность по сотрудникам".
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-sotrudnikam
func (s *ReportProfitService) GetByEmployee(ctx context.Context) (*List[ProfitByEmployee], *Response, error) {
	path := "byemployee"
	return NewRequestBuilder[List[ProfitByEmployee]](s.Endpoint, ctx).WithPath(path).Get()
}

// GetByCounterparty Запрос на получение отчета "Прибыльность по покупателям".
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-pokupatelqm
func (s *ReportProfitService) GetByCounterparty(ctx context.Context) (*List[ProfitByCounterparty], *Response, error) {
	path := "bycounterparty"
	return NewRequestBuilder[List[ProfitByCounterparty]](s.Endpoint, ctx).WithPath(path).Get()
}

// GetBySalesChannel Запрос на получение отчета "Прибыльность по каналам продаж".
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-kanalam-prodazh
func (s *ReportProfitService) GetBySalesChannel(ctx context.Context) (*List[ProfitBySalesChannel], *Response, error) {
	path := "bysaleschannel"
	return NewRequestBuilder[List[ProfitBySalesChannel]](s.Endpoint, ctx).WithPath(path).Get()
}

// GetByProductAsync Запрос на получение отчета "Прибыльность по товарам" (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-towaram
func (s *ReportProfitService) GetByProductAsync(ctx context.Context) (*AsyncResultService[List[ProfitByProduct]], *Response, error) {
	path := "byproduct"
	return NewRequestBuilder[List[ProfitByProduct]](s.Endpoint, ctx).WithPath(path).Async()
}

// GetByVariantAsync Запрос на получение отчета "Прибыльность по модификациям" (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-modifikaciqm
func (s *ReportProfitService) GetByVariantAsync(ctx context.Context) (*AsyncResultService[List[ProfitByVariant]], *Response, error) {
	path := "byvariant"
	return NewRequestBuilder[List[ProfitByVariant]](s.Endpoint, ctx).WithPath(path).Async()
}

// GetByEmployeeAsync Запрос на получение отчета "Прибыльность по сотрудникам" (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-sotrudnikam
func (s *ReportProfitService) GetByEmployeeAsync(ctx context.Context) (*AsyncResultService[List[ProfitByEmployee]], *Response, error) {
	path := "byemployee"
	return NewRequestBuilder[List[ProfitByEmployee]](s.Endpoint, ctx).WithPath(path).Async()
}

// GetByCounterpartyAsync Запрос на получение отчета "Прибыльность по покупателям" (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-pokupatelqm
func (s *ReportProfitService) GetByCounterpartyAsync(ctx context.Context) (*AsyncResultService[List[ProfitByCounterparty]], *Response, error) {
	path := "bycounterparty"
	return NewRequestBuilder[List[ProfitByCounterparty]](s.Endpoint, ctx).WithPath(path).Async()
}

// GetBySalesChannelAsync Запрос на получение отчета "Прибыльность по каналам продаж" (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-kanalam-prodazh
func (s *ReportProfitService) GetBySalesChannelAsync(ctx context.Context) (*AsyncResultService[List[ProfitBySalesChannel]], *Response, error) {
	path := "bysaleschannel"
	return NewRequestBuilder[List[ProfitBySalesChannel]](s.Endpoint, ctx).WithPath(path).Async()
}
