package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// StockAll Расширенный отчет об остатках.
// Ключевое слово: stock
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-rasshirennyj-otchet-ob-ostatkah
type StockAll struct {
	Image        Meta            `json:"image"`
	Meta         Meta            `json:"meta"`
	Folder       StockFolder     `json:"folder"`
	Uom          MetaNameWrapper `json:"uom"`
	Article      string          `json:"article"`
	ExternalCode string          `json:"externalCode"`
	Code         string          `json:"code"`
	Name         string          `json:"name"`
	InTransit    float64         `json:"inTransit"`
	Price        float64         `json:"price"`
	Quantity     float64         `json:"quantity"`
	Reserve      float64         `json:"reserve"`
	SalePrice    float64         `json:"salePrice"`
	Stock        float64         `json:"stock"`
	StockDays    float64         `json:"stockDays"`
}

func (stockAll StockAll) MetaType() MetaType {
	return MetaTypeReportStock
}

// StockByOperation Остатки по документам.
// Ключевое слово: stockbyoperation
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-ostatki-po-dokumentam
type StockByOperation struct {
	Meta      Meta                       `json:"meta"`      // Метаданные, представляющие собой ссылку на документ, по которому выдаются Остатки
	Positions []StockByOperationPosition `json:"positions"` // Массив объектов, представляющий собой Остаток по каждой из позиций
}

func (stockByOperation StockByOperation) MetaType() MetaType {
	return MetaTypeReportStockByOperation
}

// StockByOperationPosition Остатки по документам (позиция)
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-ostatki-po-dokumentam-atributy-pozicii
type StockByOperationPosition struct {
	Meta      Meta    `json:"meta"`      // Метаданные склада, по которому выводится Остаток
	Name      string  `json:"name"`      // Наименование склада
	Stock     float64 `json:"stock"`     // Остаток
	Cost      float64 `json:"cost"`      // Себестоимость
	InTransit float64 `json:"inTransit"` // Ожидание
	Reserve   float64 `json:"reserve"`   // Резерв
	Quantity  float64 `json:"quantity"`  // Доступно
}

// StockByStore Остатки по складам.
// Ключевое слово: stockbystore
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-ostatki-po-skladam
type StockByStore struct {
	Meta         Meta                   `json:"meta"`         // Метаданные позиции, по которой выдается Остаток
	StockByStore []StockByStorePosition `json:"stockByStore"` // Остатки по складам
}

func (stockByStore StockByStore) MetaType() MetaType {
	return MetaTypeReportStockByStore
}

// StockByStorePosition Остатки по складам (позиция)
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-ostatki-po-skladam-ostatki-po-skladam
type StockByStorePosition struct {
	Meta      Meta    `json:"meta"`
	Name      string  `json:"name"`
	Stock     float64 `json:"stock"`
	InTransit float64 `json:"inTransit"`
	Reserve   float64 `json:"reserve"`
}

// StockCurrentAll Краткий отчет об остатках
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-poluchit-kratkij-otchet-ob-ostatkah
type StockCurrentAll struct {
	AssortmentId string `json:"assortmentId"`
	Stock        int    `json:"stock"`
	Quantity     int    `json:"quantity"`
}

// StockCurrentByStore Краткий отчет об остатках
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-poluchit-kratkij-otchet-ob-ostatkah
type StockCurrentByStore struct {
	AssortmentId string  `json:"assortmentId"` // Выдать в отчёте только указанные товары, модификации и серии
	StoreId      string  `json:"storeId"`      // ID склада
	Stock        float64 `json:"stock"`        // Физический остаток на складах, без учёта резерва и ожидания
	FreeStock    float64 `json:"freeStock"`    // Остаток на складах за вычетом резерва
	Quantity     float64 `json:"quantity"`     // Доступно. Учитывает резерв и ожидания
}

// StockFolder Группа
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-rasshirennyj-otchet-ob-ostatkah-gruppa
type StockFolder struct {
	Meta     Meta   `json:"meta"`     // Метаданные группы товара
	Name     string `json:"name"`     // Наименование группы
	PathName string `json:"pathName"` // Наименование родительской группы
}

// ReportStockService
// Сервис для работы с отчётом "Остатки".
type ReportStockService interface {
	GetAll(ctx context.Context, params *Params) (*List[StockAll], *resty.Response, error)
	GetByStore(ctx context.Context, params *Params) (*List[StockByStore], *resty.Response, error)
	GetCurrentAll(ctx context.Context, params *Params) (*[]StockCurrentAll, *resty.Response, error)
	GetCurrentByStore(ctx context.Context, params *Params) (*[]StockCurrentByStore, *resty.Response, error)
	GetByOperationID(ctx context.Context, operationID uuid.UUID, params *Params) (*List[StockByOperation], *resty.Response, error)
	GetAllAsync(ctx context.Context, params *Params) (AsyncResultService[List[StockAll]], *resty.Response, error)
	GetByStoreAsync(ctx context.Context, params *Params) (AsyncResultService[List[StockByStore]], *resty.Response, error)
}

type reportStockService struct {
	Endpoint
}

func NewReportStockService(client *Client) ReportStockService {
	e := NewEndpoint(client, "report/stock")
	return &reportStockService{e}
}

// GetAll Запрос на получение Расширенного отчета об остатках.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-poluchit-rasshirennyj-otchet-ob-ostatkah
func (service *reportStockService) GetAll(ctx context.Context, params *Params) (*List[StockAll], *resty.Response, error) {
	path := fmt.Sprintf("%s/all", service.uri)
	return NewRequestBuilder[List[StockAll]](service.client, path).SetParams(params).Get(ctx)
}

// GetByStore Запрос на получение отчета "Остатки по складам".
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-poluchit-ostatki-po-skladam
func (service *reportStockService) GetByStore(ctx context.Context, params *Params) (*List[StockByStore], *resty.Response, error) {
	path := fmt.Sprintf("%s/bystore", service.uri)
	return NewRequestBuilder[List[StockByStore]](service.client, path).SetParams(params).Get(ctx)
}

// GetCurrentAll Запрос на получение текущих остатков без разбиения по складам.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-poluchit-kratkij-otchet-ob-ostatkah
func (service *reportStockService) GetCurrentAll(ctx context.Context, params *Params) (*[]StockCurrentAll, *resty.Response, error) {
	path := fmt.Sprintf("%s/all/current", service.uri)
	return NewRequestBuilder[[]StockCurrentAll](service.client, path).SetParams(params).Get(ctx)
}

// GetCurrentByStore Запрос на получение текущих остатков без разбиения по складам.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-poluchit-kratkij-otchet-ob-ostatkah
func (service *reportStockService) GetCurrentByStore(ctx context.Context, params *Params) (*[]StockCurrentByStore, *resty.Response, error) {
	path := fmt.Sprintf("%s/bystore/current", service.uri)
	return NewRequestBuilder[[]StockCurrentByStore](service.client, path).SetParams(params).Get(ctx)
}

// GetByOperationID
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
func (service *reportStockService) GetByOperationID(ctx context.Context, operationID uuid.UUID, params *Params) (*List[StockByOperation], *resty.Response, error) {
	path := fmt.Sprintf("%s/byoperation?operation.id=%s", service.uri, operationID)
	return NewRequestBuilder[List[StockByOperation]](service.client, path).SetParams(params).Get(ctx)
}

// GetAllAsync Запрос на получение Расширенного отчета об остатках (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-poluchit-rasshirennyj-otchet-ob-ostatkah
func (service *reportStockService) GetAllAsync(ctx context.Context, params *Params) (AsyncResultService[List[StockAll]], *resty.Response, error) {
	path := fmt.Sprintf("%s/all", service.uri)
	return NewRequestBuilder[List[StockAll]](service.client, path).SetParams(params).Async(ctx)
}

// GetByStoreAsync Запрос на получение отчета "Остатки по складам" (асинхронно).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-poluchit-ostatki-po-skladam
func (service *reportStockService) GetByStoreAsync(ctx context.Context, params *Params) (AsyncResultService[List[StockByStore]], *resty.Response, error) {
	path := fmt.Sprintf("%s/bystore", service.uri)
	return NewRequestBuilder[List[StockByStore]](service.client, path).SetParams(params).Async(ctx)
}
