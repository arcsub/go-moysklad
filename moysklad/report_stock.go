package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// StockAll Расширенный отчёт об остатках.
//
// Код сущности: stock
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-rasshirennyj-otchet-ob-ostatkah
type StockAll struct {
	Image        Meta            `json:"image"`        // Метаданные изображения Товара/Модификации/Серии
	Meta         Meta            `json:"meta"`         // Метаданные Товара/Модификации/Серии по которой выдается остаток
	Folder       StockFolder     `json:"folder"`       // Группа Товара/Модификации/Серии
	Uom          MetaNameWrapper `json:"uom"`          // Единица измерения
	Article      string          `json:"article"`      // Артикул
	ExternalCode string          `json:"externalCode"` // Внешний код сущности, по которой выводится остаток
	Code         string          `json:"code"`         // Код
	Name         string          `json:"name"`         // Наименование
	InTransit    float64         `json:"inTransit"`    // Ожидание
	Price        float64         `json:"price"`        // Себестоимость в копейках
	Quantity     float64         `json:"quantity"`     // Доступно
	Reserve      float64         `json:"reserve"`      // Резерв
	SalePrice    float64         `json:"salePrice"`    // Цена продажи
	Stock        float64         `json:"stock"`        // Остаток
	StockDays    float64         `json:"stockDays"`    // Количество дней на складе
}

// MetaType возвращает код сущности.
func (StockAll) MetaType() MetaType {
	return MetaTypeReportStock
}

// StockByOperation Остатки по документам.
//
// Код сущности: stockbyoperation
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-ostatki-po-dokumentam
type StockByOperation struct {
	Meta      Meta                       `json:"meta"`      // Метаданные, представляющие собой ссылку на документ, по которому выдаются Остатки
	Positions []StockByOperationPosition `json:"positions"` // Массив объектов, представляющий собой Остаток по каждой из позиций
}

// MetaType возвращает код сущности.
func (StockByOperation) MetaType() MetaType {
	return MetaTypeReportStockByOperation
}

// StockByOperationPosition Остатки по документам (позиция)
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-ostatki-po-dokumentam-atributy-pozicii
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
//
// Код сущности: stockbystore
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-ostatki-po-skladam
type StockByStore struct {
	Meta         Meta                   `json:"meta"`         // Метаданные позиции, по которой выдается Остаток
	StockByStore []StockByStorePosition `json:"stockByStore"` // Остатки по складам
}

// MetaType возвращает код сущности.
func (StockByStore) MetaType() MetaType {
	return MetaTypeReportStockByStore
}

// StockByStorePosition Остатки по складам (позиция)
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-ostatki-po-skladam-ostatki-po-skladam
type StockByStorePosition struct {
	Meta      Meta    `json:"meta"`      // Метаданные склада, по которому выводится Остаток
	Name      string  `json:"name"`      // Наименование склада
	Stock     float64 `json:"stock"`     // Остаток
	InTransit float64 `json:"inTransit"` // Ожидание
	Reserve   float64 `json:"reserve"`   // Резерв
}

// StockCurrentAll Краткий отчёт об остатках
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-poluchit-kratkij-otchet-ob-ostatkah
type StockCurrentAll struct {
	AssortmentID string  `json:"assortmentId"` // ID Товара/Модификации/Серии
	Stock        float64 `json:"stock"`        // Физический остаток на складах, без учёта резерва и ожидания
	FreeStock    float64 `json:"freeStock"`    // Остаток на складах за вычетом резерва
	Quantity     float64 `json:"quantity"`     // Доступно. Учитывает резерв и ожидания
	Reserve      float64 `json:"reserve"`      // Резерв
	InTransit    float64 `json:"inTransit"`    // Ожидание
}

// StockCurrentByStore Краткий отчёт об остатках
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-poluchit-kratkij-otchet-ob-ostatkah
type StockCurrentByStore struct {
	AssortmentID string  `json:"assortmentId"` // Выдать в отчёте только указанные товары, модификации и серии
	StoreID      string  `json:"storeId"`      // ID склада
	Stock        float64 `json:"stock"`        // Физический остаток на складах, без учёта резерва и ожидания
	FreeStock    float64 `json:"freeStock"`    // Остаток на складах за вычетом резерва
	Quantity     float64 `json:"quantity"`     // Доступно. Учитывает резерв и ожидания
	Reserve      float64 `json:"reserve"`      // Резерв
	InTransit    float64 `json:"inTransit"`    // Ожидание
}

// StockFolder Группа
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-rasshirennyj-otchet-ob-ostatkah-gruppa
type StockFolder struct {
	Meta     Meta   `json:"meta"`     // Метаданные группы товара
	Name     string `json:"name"`     // Наименование группы
	PathName string `json:"pathName"` // Наименование родительской группы
}

// ReportStockService описывает методы сервиса для работы с отчётом Остатки.
type ReportStockService interface {
	// GetAll выполняет запрос на получение Расширенного отчёта об остатках.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetAll(ctx context.Context, params ...*Params) (*List[StockAll], *resty.Response, error)

	// GetByStore выполняет запрос на получение отчёта "Остатки по складам".
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetByStore(ctx context.Context, params ...*Params) (*List[StockByStore], *resty.Response, error)

	// GetCurrentAll выполняет запрос на получение текущих остатков без разбиения по складам.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает список остатков.
	GetCurrentAll(ctx context.Context, params ...*Params) (*Slice[StockCurrentAll], *resty.Response, error)

	// GetCurrentByStore выполняет запрос на получение текущих остатков с разбиением по складам.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает список остатков.
	GetCurrentByStore(ctx context.Context, params ...*Params) (*Slice[StockCurrentByStore], *resty.Response, error)

	// GetByOperationID выполняет запрос на получение отчёта "Остатки по документу".
	// Принимает контекст, ID документа и опционально объект параметров запроса Params.
	// Данный запрос работает со следующими типами документов:
	//	- Отгрузка
	//	- Заказ покупателя
	//	- Розничная продажа
	//	- Счет поставщика
	//	- Розничная продажа
	//	- Заказ поставщику
	//	- Приемка
	//	- Розничный возврат
	//	- Возврат поставщику
	//	- Возврат покупателя
	// Возвращает объект List.
	GetByOperationID(ctx context.Context, operationID uuid.UUID, params ...*Params) (*List[StockByOperation], *resty.Response, error)

	// GetAllAsync выполняет запрос на получение Расширенного отчёта об остатках (асинхронно).
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает сервис для работы с контекстом асинхронного запроса.
	GetAllAsync(ctx context.Context, params ...*Params) (AsyncResultService[List[StockAll]], *resty.Response, error)

	// GetByStoreAsync выполняет запрос на получение отчёта "Остатки по складам" (асинхронно).
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает сервис для работы с контекстом асинхронного запроса.
	GetByStoreAsync(ctx context.Context, params ...*Params) (AsyncResultService[List[StockByStore]], *resty.Response, error)
}

type reportStockService struct {
	Endpoint
}

// NewReportStockService принимает [Client] и возвращает сервис для работы с отчётом Остатки.
func NewReportStockService(client *Client) ReportStockService {
	return &reportStockService{NewEndpoint(client, "report/stock")}
}

func (service *reportStockService) GetAll(ctx context.Context, params ...*Params) (*List[StockAll], *resty.Response, error) {
	path := fmt.Sprintf("%s/all", service.uri)
	return NewRequestBuilder[List[StockAll]](service.client, path).SetParams(params...).Get(ctx)
}

func (service *reportStockService) GetByStore(ctx context.Context, params ...*Params) (*List[StockByStore], *resty.Response, error) {
	path := fmt.Sprintf("%s/bystore", service.uri)
	return NewRequestBuilder[List[StockByStore]](service.client, path).SetParams(params...).Get(ctx)
}

func (service *reportStockService) GetCurrentAll(ctx context.Context, params ...*Params) (*Slice[StockCurrentAll], *resty.Response, error) {
	path := fmt.Sprintf("%s/all/current", service.uri)
	return NewRequestBuilder[Slice[StockCurrentAll]](service.client, path).SetParams(params...).Get(ctx)
}

func (service *reportStockService) GetCurrentByStore(ctx context.Context, params ...*Params) (*Slice[StockCurrentByStore], *resty.Response, error) {
	path := fmt.Sprintf("%s/bystore/current", service.uri)
	return NewRequestBuilder[Slice[StockCurrentByStore]](service.client, path).SetParams(params...).Get(ctx)
}

func (service *reportStockService) GetByOperationID(ctx context.Context, operationID uuid.UUID, params ...*Params) (*List[StockByOperation], *resty.Response, error) {
	path := fmt.Sprintf("%s/byoperation?operation.id=%s", service.uri, operationID)
	return NewRequestBuilder[List[StockByOperation]](service.client, path).SetParams(params...).Get(ctx)
}

func (service *reportStockService) GetAllAsync(ctx context.Context, params ...*Params) (AsyncResultService[List[StockAll]], *resty.Response, error) {
	path := fmt.Sprintf("%s/all", service.uri)
	return NewRequestBuilder[List[StockAll]](service.client, path).SetParams(params...).Async(ctx)
}

func (service *reportStockService) GetByStoreAsync(ctx context.Context, params ...*Params) (AsyncResultService[List[StockByStore]], *resty.Response, error) {
	path := fmt.Sprintf("%s/bystore", service.uri)
	return NewRequestBuilder[List[StockByStore]](service.client, path).SetParams(params...).Async(ctx)
}
