package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
)

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
func (s *ReportTurnoverService) GetAll(ctx context.Context) (*List[TurnoverAll], *resty.Response, error) {
	path := fmt.Sprintf("%s/all", s.uri)
	return NewRequestBuilder[List[TurnoverAll]](s.client, path).Get(ctx)
}

// GetByStoreWithProduct Отчет обороты по товару и его модификациям с детализацией по складам.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-oboroty-oboroty-po-towaru-s-detalizaciej-po-skladam
func (s *ReportTurnoverService) GetByStoreWithProduct(ctx context.Context, product *Product) (*List[TurnoverByOperation], *resty.Response, error) {
	path := fmt.Sprintf("%s/bystore", s.uri)
	params := new(Params).WithFilterEquals("product", *product.Meta.Href)
	return NewRequestBuilder[List[TurnoverByOperation]](s.client, path).SetParams(params).Get(ctx)
}

// GetByStoreWithVariant Отчет обороты по модификации с детализацией по складам.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-oboroty-oboroty-po-towaru-s-detalizaciej-po-skladam
func (s *ReportTurnoverService) GetByStoreWithVariant(ctx context.Context, variant *Variant) (*List[TurnoverByOperation], *resty.Response, error) {
	path := fmt.Sprintf("%s/bystore", s.uri)
	params := new(Params).WithFilterEquals("variant", *variant.Meta.Href)
	return NewRequestBuilder[List[TurnoverByOperation]](s.client, path).SetParams(params).Get(ctx)
}

// GetByOperationsWithProduct Запрос на получение отчета Обороты по товару с детализацией по документам.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-oboroty-poluchit-oboroty-po-towaru-s-detalizaciej-po-dokumentam
func (s *ReportTurnoverService) GetByOperationsWithProduct(ctx context.Context, product *Product) (*List[TurnoverByOperation], *resty.Response, error) {
	path := fmt.Sprintf("%s/byoperations", s.uri)
	params := new(Params).WithFilterEquals("product", *product.Meta.Href)
	return NewRequestBuilder[List[TurnoverByOperation]](s.client, path).SetParams(params).Get(ctx)
}

// GetByOperationsWithVariant Запрос на получение отчета Обороты по модификации с детализацией по документам.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-oboroty-poluchit-oboroty-po-towaru-s-detalizaciej-po-dokumentam
func (s *ReportTurnoverService) GetByOperationsWithVariant(ctx context.Context, variant *Variant) (*List[TurnoverByOperation], *resty.Response, error) {
	path := fmt.Sprintf("%s/byoperations", s.uri)
	params := new(Params).WithFilterEquals("variant", *variant.Meta.Href)
	return NewRequestBuilder[List[TurnoverByOperation]](s.client, path).SetParams(params).Get(ctx)
}
