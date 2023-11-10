package moysklad

// Metadata Глобальные метаданные.
// Ключевое слово: metadata
type Metadata struct {
	Inventory                 AttributesCreateSharedWrapper       `json:"inventory"`
	Prepayment                AttributesCreateSharedWrapper       `json:"prepayment"`
	ProductionsTageCompletion AttributesCreateSharedWrapper       `json:"productionstagecompletion"`
	Enter                     AttributesCreateSharedWrapper       `json:"enter"`
	RetailSalesReturn         AttributesCreateSharedWrapper       `json:"retailsalesreturn"`
	Store                     AttributesCreateSharedWrapper       `json:"store"`
	BonusTransaction          AttributesCreateSharedWrapper       `json:"bonustransaction"`
	PriceList                 AttributesCreateSharedWrapper       `json:"pricelist"`
	PaymentIn                 AttributesCreateSharedWrapper       `json:"paymentin"`
	RetailDrawerCashIn        AttributesCreateSharedWrapper       `json:"retaildrawercashin"`
	Supply                    AttributesCreateSharedWrapper       `json:"supply"`
	CommissionReportIn        AttributesCreateSharedWrapper       `json:"commissionreportin"`
	CommissionReportOut       AttributesCreateSharedWrapper       `json:"commissionreportout"`
	PrepaymentReturn          AttributesCreateSharedWrapper       `json:"prepaymentreturn"`
	RetailShift               AttributesCreateSharedWrapper       `json:"retailshift"`
	Loss                      AttributesCreateSharedWrapper       `json:"loss"`
	FactureOut                AttributesCreateSharedWrapper       `json:"factureout"`
	PurchaseOrder             AttributesCreateSharedWrapper       `json:"purchaseorder"`
	CrptOrder                 AttributesCreateSharedWrapper       `json:"crptorder"`
	Move                      AttributesCreateSharedWrapper       `json:"move"`
	Employee                  AttributesCreateSharedWrapper       `json:"employee"`
	InvoiceOut                AttributesCreateSharedWrapper       `json:"invoiceout"`
	RetailDemand              AttributesCreateSharedWrapper       `json:"retaildemand"`
	SalesReturn               AttributesCreateSharedWrapper       `json:"salesreturn"`
	InternalOrder             AttributesCreateSharedWrapper       `json:"internalorder"`
	Organization              AttributesCreateSharedWrapper       `json:"organization"`
	InvoiceIn                 AttributesCreateSharedWrapper       `json:"invoicein"`
	Demand                    AttributesCreateSharedWrapper       `json:"demand"`
	PaymentOut                AttributesCreateSharedWrapper       `json:"paymentout"`
	CounterpartyAdjustment    AttributesCreateSharedWrapper       `json:"counterpartyadjustment"`
	ProductionTask            AttributesCreateSharedWrapper       `json:"productiontask"`
	Processing                AttributesCreateSharedWrapper       `json:"processing"`
	CashOut                   AttributesCreateSharedWrapper       `json:"cashout"`
	PurchaseReturn            AttributesCreateSharedWrapper       `json:"purchasereturn"`
	RetailDrawerCashOut       AttributesCreateSharedWrapper       `json:"retaildrawercashout"`
	ProcessingOrder           AttributesCreateSharedWrapper       `json:"processingorder"`
	Project                   AttributesCreateSharedWrapper       `json:"project"`
	FactureOn                 AttributesCreateSharedWrapper       `json:"facturein"`
	CashIn                    AttributesCreateSharedWrapper       `json:"cashin"`
	Contract                  AttributesCreateSharedWrapper       `json:"contract"`
	BonusProgram              AttributesWrapper                   `json:"bonusprogram"`
	ProcessingPlan            AttributesWrapper                   `json:"processingplan"`
	Consignment               AttributesWrapper                   `json:"consignment"`
	Application               MetaWrapper                         `json:"application"`
	ProcessingPlanFolder      MetaWrapper                         `json:"processingplanfolder"`
	ProductFolder             MetaWrapper                         `json:"productfolder"`
	Assortment                Meta                                `json:"assortment"`
	Counterparty              CounterPartyOption                  `json:"counterparty"`
	CompanySettings           MetadataCompanySettings             `json:"companysettings"`
	CustomerOrder             AttributesStatesCreateSharedWrapper `json:"customerorder"`
	Variant                   MetadataVariant                     `json:"variant"`
	Bundle                    MetadataAttributeShared             `json:"bundle"`
	Product                   MetadataAttributeShared             `json:"product"`
	Service                   MetadataAttributeShared             `json:"service"`
}

type AttributesWrapper struct {
	Meta       Meta        `json:"meta"`
	Attributes MetaWrapper `json:"attributes"`
}

type AttributesCreateSharedWrapper struct {
	AttributesWrapper
	Meta         Meta `json:"meta"`
	CreateShared bool `json:"createShared"`
}

type AttributesStatesCreateSharedWrapper struct {
	AttributesCreateSharedWrapper
	States []StatesElement `json:"states"`
}

type CounterPartyOption struct {
	AttributesStatesCreateSharedWrapper
	Tags []string `json:"tags"`
}

type StatesElement struct {
	Meta       Meta   `json:"meta"`
	EntityType string `json:"entityType"`
	StateType  string `json:"stateType"`
	Color      int    `json:"color"`
}

func (m Metadata) MetaType() MetaType {
	return MetaTypeMetadata
}

// Метаданные сущности
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/workbook/#workbook-metadannye-metadannye-suschnosti

type MetadataAttribute struct {
	Meta       Meta                 `json:"meta"`
	Attributes MetaArray[Attribute] `json:"attributes"`
}

type MetadataAttributeShared struct {
	MetadataAttribute
	CreateShared bool `json:"createShared"`
}

type MetadataAttributeSharedStates struct {
	States []State `json:"states"`
	MetadataAttributeShared
}

type MetadataVariant struct {
	Meta            Meta             `json:"meta"`            // Метаданные
	Characteristics []Characteristic `json:"characteristics"` // Коллекция всех созданных характеристик Модификаций
}

type MetadataCounterparty struct {
	Tags []string `json:"tags"`
	MetadataAttributeShared
}

type MetadataAttributeSharedPriceTypes struct {
	PriceTypes []struct {
		Name string `json:"name,omitempty"`
	} `json:"priceTypes"`
	MetadataAttributeShared // Наименование
}

type MetadataCompanySettings struct {
	CustomEntities []CustomEntityElement `json:"customEntities"`
	MetadataAttribute
}
