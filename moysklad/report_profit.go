package moysklad

// Profit общие поля для структур отчёта "Прибыльность"
type Profit struct {
	Margin         float64 `json:"margin"`         // Рентабельность
	Profit         float64 `json:"profit"`         // Прибыль
	ReturnAvgCheck float64 `json:"returnAvgCheck"` // Средний чек возврата
	ReturnCostSum  float64 `json:"returnCostSum"`  // Сумма себестоимостей возвратов
	ReturnCount    float64 `json:"returnCount"`    // Количество возвратов
	ReturnSum      float64 `json:"returnSum"`      // Сумма возвратов
	SalesAvgCheck  float64 `json:"salesAvgCheck"`  // Средний чек продаж
	SalesCount     float64 `json:"salesCount"`     // Количество продаж
	SellCostSum    float64 `json:"sellCostSum"`    // Сумма себестоимостей продаж
	SellSum        float64 `json:"sellSum"`        // Сумма продаж
}

// ProfitByAssortment Прибыльность по товарам
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-towaram
type ProfitByAssortment struct {
	Assortment     ProfitReportAssortment `json:"assortment"`     // Краткое представление Товара или Услуги в отчете
	Margin         float64                `json:"margin"`         // Рентабельность
	Profit         float64                `json:"profit"`         // Прибыль
	ReturnCost     float64                `json:"returnCost"`     // Себестоимость возвратов
	ReturnCostSum  float64                `json:"returnCostSum"`  // Сумма себестоимостей возвратов
	ReturnPrice    float64                `json:"returnPrice"`    // Цена возвратов
	ReturnQuantity float64                `json:"returnQuantity"` // Количество возвратов
	ReturnSum      float64                `json:"returnSum"`      // Сумма возвратов
	SellCost       float64                `json:"sellCost"`       // Себестоимость
	SellCostSum    float64                `json:"sellCostSum"`    // Сумма себестоимостей продаж
	SellPrice      float64                `json:"SellPrice"`      // Цена продаж (средняя)
	SellQuantity   float64                `json:"sellQuantity"`   // Проданное количество
	SellSum        float64                `json:"sellSum"`        // Сумма продаж
}

// ProfitReportAssortment Структура объекта assortment
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-struktura-ob-ekta-assortment
type ProfitReportAssortment struct {
	MetaName          // Метаданные/Наименование Товара или Услуги
	Code     string   `json:"code"`          // Код товара или услуги
	Uom      MetaName `json:"uom,omitempty"` // Единица измерения
	Article  string   `json:"article"`       // Артикул товара
	Image    Meta     `json:"image"`         // Изображение товара
}

// ProfitByCounterparty Прибыльность по покупателям
// Ключевое слово: salesbyCounterparty
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-pokupatelqm
type ProfitByCounterparty struct {
	Counterparty MetaName `json:"counterparty"`
	Profit
}

func (r ProfitByCounterparty) MetaType() MetaType {
	return MetaTypeReportProfitByCounterparty
}

// ProfitByEmployee Прибыльность по сотрудникам
// Ключевое слово: salesbyemployee
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-sotrudnikam
type ProfitByEmployee struct {
	Employee MetaName `json:"employee"`
	Profit
}

func (r ProfitByEmployee) MetaType() MetaType {
	return MetaTypeReportProfitByEmployee
}

// ProfitByProduct Прибыльность по товарам
// Ключевое слово: salesbyproduct
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-towaram
type ProfitByProduct struct {
	ProfitByAssortment
}

func (r ProfitByProduct) MetaType() MetaType {
	return MetaTypeReportProfitByProduct
}

// ProfitBySalesChannel Прибыльность по каналам продаж
// Ключевое слово: salesbysaleschannel
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-kanalam-prodazh
type ProfitBySalesChannel struct {
	SalesChannel struct {
		Meta Meta             `json:"meta"`
		Name string           `json:"name"`
		Type SalesChannelType `json:"type"`
	} `json:"salesChannel"`
	Profit
}

func (r ProfitBySalesChannel) MetaType() MetaType {
	return MetaTypeReportProfitBySalesChannel
}

// ProfitByVariant Прибыльность по модификациям
// Ключевое слово: salesbyvariant
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-modifikaciqm
type ProfitByVariant struct {
	ProfitByAssortment
}

func (r ProfitByVariant) MetaType() MetaType {
	return MetaTypeReportProfitByVariant
}
