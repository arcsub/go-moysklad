package moysklad

// StockAll Расширенный отчет об остатках.
// Ключевое слово: stock
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-rasshirennyj-otchet-ob-ostatkah
type StockAll struct {
	Article      string      `json:"article"`      // Артикул
	Code         string      `json:"code"`         // Код
	ExternalCode string      `json:"externalCode"` // Внешний код сущности, по которой выводится остаток
	Folder       StockFolder `json:"folder"`       // Группа Товара/Модификации/Серии
	Image        Meta        `json:"image"`        // Метаданные изображения Товара/Модификации/Серии
	InTransit    float64     `json:"inTransit"`    // Ожидание
	Meta         Meta        `json:"meta"`         // Метаданные Товара/Модификации/Серии по которой выдается остаток
	Name         string      `json:"name"`         // Наименование
	Price        float64     `json:"price"`        // Себестоимость
	Quantity     float64     `json:"quantity"`     // Доступно
	Reserve      float64     `json:"reserve"`      // Резерв
	SalePrice    float64     `json:"salePrice"`    // Цена продажи
	Stock        float64     `json:"stock"`        // Остаток
	StockDays    int         `json:"stockDays"`    // Количество дней на складе
	Uom          MetaName    `json:"uom"`          // Единица измерения
}

func (s StockAll) MetaType() MetaType {
	return MetaTypeReportStock
}

// StockByOperation Остатки по документам.
// Ключевое слово: stockbyoperation
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-ostatki-po-dokumentam
type StockByOperation struct {
	Meta      Meta                       `json:"meta"`      // Метаданные, представляющие собой ссылку на документ, по которому выдаются Остатки
	Positions []StockByOperationPosition `json:"positions"` // Массив объектов, представляющий собой Остаток по каждой из позиций
}

func (s StockByOperation) MetaType() MetaType {
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

func (s StockByStore) MetaType() MetaType {
	return MetaTypeReportStockByStore
}

// StockByStorePosition Остатки по складам (позиция)
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-ostatki-ostatki-po-skladam-ostatki-po-skladam
type StockByStorePosition struct {
	Meta      Meta    `json:"meta"`      // Метаданные склада, по которому выводится Остаток
	Stock     float64 `json:"stock"`     // Остаток
	InTransit float64 `json:"inTransit"` // Ожидание
	Reserve   float64 `json:"reserve"`   // Резерв
	Name      string  `json:"name"`      // Наименование склада
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
