package moysklad

// Profit общие поля для структур отчёта "Прибыльность"
type Profit struct {
	Margin         Decimal `json:"margin"`
	Profit         Decimal `json:"profit"`
	ReturnAvgCheck Decimal `json:"returnAvgCheck"`
	ReturnCostSum  Decimal `json:"returnCostSum"`
	ReturnSum      Decimal `json:"returnSum"`
	SalesAvgCheck  Decimal `json:"salesAvgCheck"`
	SellCostSum    Decimal `json:"sellCostSum"`
	SellSum        Decimal `json:"sellSum"`
	ReturnCount    float64 `json:"returnCount"`
	SalesCount     float64 `json:"salesCount"`
}

// ProfitByAssortment Прибыльность по товарам
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-poluchit-pribyl-nost-po-towaram
type ProfitByAssortment struct {
	Assortment     ProfitReportAssortment `json:"assortment"`
	SellCostSum    Decimal                `json:"sellCostSum"`
	Profit         Decimal                `json:"profit"`
	ReturnCost     Decimal                `json:"returnCost"`
	ReturnCostSum  Decimal                `json:"returnCostSum"`
	ReturnPrice    Decimal                `json:"returnPrice"`
	ReturnSum      Decimal                `json:"returnSum"`
	SellCost       Decimal                `json:"sellCost"`
	Margin         Decimal                `json:"margin"`
	SellPrice      Decimal                `json:"SellPrice"`
	SellSum        Decimal                `json:"sellSum"`
	ReturnQuantity float64                `json:"returnQuantity"`
	SellQuantity   float64                `json:"sellQuantity"`
}

// ProfitReportAssortment Структура объекта assortment
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pribyl-nost-struktura-ob-ekta-assortment
type ProfitReportAssortment struct {
	Image Meta `json:"image"`
	MetaName
	Uom     MetaName `json:"uom,omitempty"`
	Code    string   `json:"code"`
	Article string   `json:"article"`
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
