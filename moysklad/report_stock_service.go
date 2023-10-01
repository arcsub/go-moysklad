package moysklad

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

// ReportStockService
// Сервис для работы с отчётом "Остатки".
type ReportStockService struct {
	Endpoint
}

func NewReportStockService(client *Client) *ReportStockService {
	e := NewEndpoint(client, "report/stock")
	return &ReportStockService{e}
}

// GetAll Запрос на получение Расширенного отчета об остатках.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-poluchit-rasshirennyj-otchet-ob-ostatkah
func (s *ReportStockService) GetAll(ctx context.Context) (*List[StockAll], *Response, error) {
	path := "all"
	return NewRequestBuilder[List[StockAll]](s.Endpoint, ctx).WithPath(path).Get()
}

// GetByStore Запрос на получение отчета "Остатки по складам".
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-poluchit-ostatki-po-skladam
func (s *ReportStockService) GetByStore(ctx context.Context) (*List[StockByStore], *Response, error) {
	path := "bystore"
	return NewRequestBuilder[List[StockByStore]](s.Endpoint, ctx).WithPath(path).Get()
}

// GetCurrentAll Запрос на получение текущих остатков без разбиения по складам.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-poluchit-kratkij-otchet-ob-ostatkah
func (s *ReportStockService) GetCurrentAll(ctx context.Context) (*Slice[StockCurrentAll], *Response, error) {
	path := "all/current"
	return NewRequestBuilder[Slice[StockCurrentAll]](s.Endpoint, ctx).WithPath(path).Get()
}

// GetCurrentByStore Запрос на получение текущих остатков без разбиения по складам.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-poluchit-kratkij-otchet-ob-ostatkah
func (s *ReportStockService) GetCurrentByStore(ctx context.Context) (*Slice[StockCurrentByStore], *Response, error) {
	path := "bystore/current"
	return NewRequestBuilder[Slice[StockCurrentByStore]](s.Endpoint, ctx).WithPath(path).Get()
}

// GetByOperationId
// Запрос на получение отчёта "Остатки по документу"
// Данный запрос работает со следующими типами документов:
// – Отгрузка
// – Заказ покупателя
// – Розничная продажа
// – Счет поставщика
// – Розничная продажа
// – Заказ поставщику
// – Приемка
// – Розничный возврат
// – Возврат поставщику
// – Возврат покупателя
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-poluchit-ostatki-po-dokumentu
func (s *ReportStockService) GetByOperationId(ctx context.Context, operationId uuid.UUID) (*List[StockByOperation], *Response, error) {
	path := fmt.Sprintf("byoperation?operation.id=%s", operationId)
	return NewRequestBuilder[List[StockByOperation]](s.Endpoint, ctx).WithPath(path).Get()
}

// GetAllAsync Запрос на получение Расширенного отчета об остатках (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-poluchit-rasshirennyj-otchet-ob-ostatkah
func (s *ReportStockService) GetAllAsync(ctx context.Context) (*AsyncResultService[List[StockAll]], *Response, error) {
	path := "all"
	return NewRequestBuilder[List[StockAll]](s.Endpoint, ctx).WithPath(path).Async()
}

// GetByStoreAsync Запрос на получение отчета "Остатки по складам" (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-poluchit-ostatki-po-skladam
func (s *ReportStockService) GetByStoreAsync(ctx context.Context) (*AsyncResultService[List[StockByStore]], *Response, error) {
	path := "bystore"
	return NewRequestBuilder[List[StockByStore]](s.Endpoint, ctx).WithPath(path).Async()
}
