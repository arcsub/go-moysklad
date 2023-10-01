package moysklad

// Metadata Глобальные метаданные.
// Ключевое слово: metadata
type Metadata struct {
	CustomerOrder             AttributesStatesCreateSharedWrapper `json:"customerorder"`
	InvoiceIn                 AttributesCreateSharedWrapper       `json:"invoicein"`
	InvoiceOut                AttributesCreateSharedWrapper       `json:"invoiceout"`
	Inventory                 AttributesCreateSharedWrapper       `json:"inventory"`
	Prepayment                AttributesCreateSharedWrapper       `json:"prepayment"`
	RetailSalesReturn         AttributesCreateSharedWrapper       `json:"retailsalesreturn"`
	Consignment               AttributesWrapper                   `json:"consignment"`
	Bundle                    MetadataAttributeShared             `json:"bundle"`
	Product                   MetadataAttributeShared             `json:"product"`
	Service                   MetadataAttributeShared             `json:"service"`
	Store                     AttributesCreateSharedWrapper       `json:"store"`
	BonusTransaction          AttributesCreateSharedWrapper       `json:"bonustransaction"`
	PriceList                 AttributesCreateSharedWrapper       `json:"pricelist"`
	PaymentIn                 AttributesCreateSharedWrapper       `json:"paymentin"`
	Supply                    AttributesCreateSharedWrapper       `json:"supply"`
	Enter                     AttributesCreateSharedWrapper       `json:"enter"`
	Counterparty              CounterPartyOption                  `json:"counterparty"`
	Variant                   MetadataVariant                     `json:"variant"`
	RetailDemand              AttributesCreateSharedWrapper       `json:"retaildemand"`
	FactureOut                AttributesCreateSharedWrapper       `json:"factureout"`
	PurchaseOrder             AttributesCreateSharedWrapper       `json:"purchaseorder"`
	CrptOrder                 AttributesCreateSharedWrapper       `json:"crptorder"`
	RetailDrawerCashIn        AttributesCreateSharedWrapper       `json:"retaildrawercashin"`
	Assortment                Meta                                `json:"assortment"`
	ProductFolder             MetaWrapper                         `json:"productfolder"`
	Loss                      AttributesCreateSharedWrapper       `json:"loss"`
	SalesReturn               AttributesCreateSharedWrapper       `json:"salesreturn"`
	Application               MetaWrapper                         `json:"application"`
	Organization              AttributesCreateSharedWrapper       `json:"organization"`
	CompanySettings           MetadataCompanySettings             `json:"companysettings"`
	Demand                    AttributesCreateSharedWrapper       `json:"demand"`
	PaymentOut                AttributesCreateSharedWrapper       `json:"paymentout"`
	CounterpartyAdjustment    AttributesCreateSharedWrapper       `json:"counterpartyadjustment"`
	ProductionTask            AttributesCreateSharedWrapper       `json:"productiontask"`
	Processing                AttributesCreateSharedWrapper       `json:"processing"`
	CashOut                   AttributesCreateSharedWrapper       `json:"cashout"`
	PurchaseReturn            AttributesCreateSharedWrapper       `json:"purchasereturn"`
	ProcessingPlan            AttributesWrapper                   `json:"processingplan"`
	ProcessingOrder           AttributesCreateSharedWrapper       `json:"processingorder"`
	Project                   AttributesCreateSharedWrapper       `json:"project"`
	BonusProgram              AttributesWrapper                   `json:"bonusprogram"`
	CashIn                    AttributesCreateSharedWrapper       `json:"cashin"`
	Contract                  AttributesCreateSharedWrapper       `json:"contract"`
	FactureOn                 AttributesCreateSharedWrapper       `json:"facturein"`
	ProcessingPlanFolder      MetaWrapper                         `json:"processingplanfolder"`
	Move                      AttributesCreateSharedWrapper       `json:"move"`
	RetailDrawerCashOut       AttributesCreateSharedWrapper       `json:"retaildrawercashout"`
	Employee                  AttributesCreateSharedWrapper       `json:"employee"`
	CommissionReportIn        AttributesCreateSharedWrapper       `json:"commissionreportin"`
	ProductionsTageCompletion AttributesCreateSharedWrapper       `json:"productionstagecompletion"`
	CommissionReportOut       AttributesCreateSharedWrapper       `json:"commissionreportout"`
	InternalOrder             AttributesCreateSharedWrapper       `json:"internalorder"`
	RetailShift               AttributesCreateSharedWrapper       `json:"retailshift"`
	PrepaymentReturn          AttributesCreateSharedWrapper       `json:"prepaymentreturn"`
}

type AttributesWrapper struct {
	Meta       Meta        `json:"meta"`
	Attributes MetaWrapper `json:"attributes"`
}

type AttributesCreateSharedWrapper struct {
	Meta         Meta `json:"meta"`
	CreateShared bool `json:"createShared"`
	AttributesWrapper
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
	Color      int    `json:"color"`
	EntityType string `json:"entityType"`
	StateType  string `json:"stateType"`
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
	MetadataAttributeShared
	States []State `json:"states"`
}

type MetadataVariant struct {
	Meta            Meta             `json:"meta"`            // Метаданные
	Characteristics []Characteristic `json:"characteristics"` // Коллекция всех созданных характеристик Модификаций
}

type MetadataCounterparty struct {
	MetadataAttributeShared
	Tags []string `json:"tags"`
}

type MetadataAttributeSharedPriceTypes struct {
	MetadataAttributeShared
	PriceTypes []struct {
		Name string `json:"name,omitempty"` // Наименование
	} `json:"priceTypes"`
}

type MetadataCompanySettings struct {
	MetadataAttribute
	CustomEntities []CustomEntityElement `json:"customEntities"`
}
