package moysklad

import "context"

// ReportTurnoverService
// Сервис для работы с отчётом "Обороты".
type ReportTurnoverService struct {
	Endpoint
}

func NewReportTurnoverService(client *Client) *ReportTurnoverService {
	e := NewEndpoint(client, "report/turnover")
	return &ReportTurnoverService{e}
}

// GetAll Запрос на получение отчета "Обороты по товарам".
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-oboroty-poluchit-oboroty-po-towaram
func (s *ReportTurnoverService) GetAll(ctx context.Context) (*List[TurnoverAll], *Response, error) {
	path := "all"
	return NewRequestBuilder[List[TurnoverAll]](s.Endpoint, ctx).WithPath(path).Get()
}

// GetByStoreWithProduct Отчет обороты по товару и его модификациям с детализацией по складам.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-oboroty-oboroty-po-towaru-s-detalizaciej-po-skladam
func (s *ReportTurnoverService) GetByStoreWithProduct(ctx context.Context, product *Product) (*List[TurnoverByOperation], *Response, error) {
	path := "bystore"
	params := new(Params).WithFilterEquals("product", *product.Meta.Href)
	return NewRequestBuilder[List[TurnoverByOperation]](s.Endpoint, ctx).WithPath(path).WithParams(params).Get()
}

// GetByStoreWithVariant Отчет обороты по модификации с детализацией по складам.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-oboroty-oboroty-po-towaru-s-detalizaciej-po-skladam
func (s *ReportTurnoverService) GetByStoreWithVariant(ctx context.Context, variant *Variant) (*List[TurnoverByOperation], *Response, error) {
	path := "bystore"
	params := new(Params).WithFilterEquals("variant", *variant.Meta.Href)
	return NewRequestBuilder[List[TurnoverByOperation]](s.Endpoint, ctx).WithPath(path).WithParams(params).Get()
}

// GetByOperationsWithProduct Запрос на получение отчета Обороты по товару с детализацией по документам.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-oboroty-poluchit-oboroty-po-towaru-s-detalizaciej-po-dokumentam
func (s *ReportTurnoverService) GetByOperationsWithProduct(ctx context.Context, product *Product) (*List[TurnoverByOperation], *Response, error) {
	path := "byoperations"
	params := new(Params).WithFilterEquals("product", *product.Meta.Href)
	return NewRequestBuilder[List[TurnoverByOperation]](s.Endpoint, ctx).WithPath(path).WithParams(params).Get()
}

// GetByOperationsWithVariant Запрос на получение отчета Обороты по модификации с детализацией по документам.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-oboroty-poluchit-oboroty-po-towaru-s-detalizaciej-po-dokumentam
func (s *ReportTurnoverService) GetByOperationsWithVariant(ctx context.Context, variant *Variant) (*List[TurnoverByOperation], *Response, error) {
	path := "byoperations"
	params := new(Params).WithFilterEquals("variant", *variant.Meta.Href)
	return NewRequestBuilder[List[TurnoverByOperation]](s.Endpoint, ctx).WithPath(path).WithParams(params).Get()
}
