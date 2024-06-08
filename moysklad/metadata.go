package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
)

// Metadata Глобальные метаданные.
// Ключевое слово: metadata
type Metadata struct {
	CompanySettings           MetadataCompanySettings             `json:"companysettings"`
	BonusProgram              MetaAttributesWrapper               `json:"bonusprogram"`
	Consignment               MetaAttributesWrapper               `json:"consignment"`
	ProcessingPlan            MetaAttributesWrapper               `json:"processingplan"`
	Assortment                Meta                                `json:"assortment"`
	ProductFolder             MetaWrapper                         `json:"productfolder"`
	ProcessingPlanFolder      MetaWrapper                         `json:"processingplanfolder"`
	Application               MetaWrapper                         `json:"application"`
	InvoiceIn                 AttributesCreateSharedWrapper       `json:"invoicein"`
	Processing                AttributesCreateSharedWrapper       `json:"processing"`
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
	Inventory                 AttributesCreateSharedWrapper       `json:"inventory"`
	Demand                    AttributesCreateSharedWrapper       `json:"demand"`
	PaymentOut                AttributesCreateSharedWrapper       `json:"paymentout"`
	CounterpartyAdjustment    AttributesCreateSharedWrapper       `json:"counterpartyadjustment"`
	ProductionTask            AttributesCreateSharedWrapper       `json:"productiontask"`
	RetailDrawerCashIn        AttributesCreateSharedWrapper       `json:"retaildrawercashin"`
	CashOut                   AttributesCreateSharedWrapper       `json:"cashout"`
	PurchaseReturn            AttributesCreateSharedWrapper       `json:"purchasereturn"`
	RetailDrawerCashOut       AttributesCreateSharedWrapper       `json:"retaildrawercashout"`
	ProcessingOrder           AttributesCreateSharedWrapper       `json:"processingorder"`
	Project                   AttributesCreateSharedWrapper       `json:"project"`
	FactureOn                 AttributesCreateSharedWrapper       `json:"facturein"`
	CashIn                    AttributesCreateSharedWrapper       `json:"cashin"`
	Contract                  AttributesCreateSharedWrapper       `json:"contract"`
	PaymentIn                 AttributesCreateSharedWrapper       `json:"paymentin"`
	PriceList                 AttributesCreateSharedWrapper       `json:"pricelist"`
	BonusTransaction          AttributesCreateSharedWrapper       `json:"bonustransaction"`
	Store                     AttributesCreateSharedWrapper       `json:"store"`
	RetailSalesReturn         AttributesCreateSharedWrapper       `json:"retailsalesreturn"`
	Enter                     AttributesCreateSharedWrapper       `json:"enter"`
	ProductionsTageCompletion AttributesCreateSharedWrapper       `json:"productionstagecompletion"`
	Service                   MetaAttributesSharedWrapper         `json:"service"`
	Prepayment                AttributesCreateSharedWrapper       `json:"prepayment"`
	Product                   MetaAttributesSharedWrapper         `json:"product"`
	Bundle                    MetaAttributesSharedWrapper         `json:"bundle"`
	Counterparty              CounterpartyOption                  `json:"counterparty"`
	CustomerOrder             AttributesStatesCreateSharedWrapper `json:"customerorder"`
	Variant                   MetaCharacteristicsWrapper          `json:"variant"`
}

type MetaAttributesWrapper struct {
	Meta       Meta        `json:"meta"`
	Attributes MetaWrapper `json:"attributes"`
}

type AttributesCreateSharedWrapper struct {
	MetaAttributesWrapper
	CreateShared bool `json:"createShared"`
}

type AttributesStatesCreateSharedWrapper struct {
	AttributesCreateSharedWrapper
	States Slice[StatesElement] `json:"states"`
}

type CounterpartyOption struct {
	AttributesStatesCreateSharedWrapper
	Tags Tags `json:"tags"`
}

type StatesElement struct {
	Meta       Meta   `json:"meta"`
	EntityType string `json:"entityType"`
	StateType  string `json:"stateType"`
	Color      int    `json:"color"`
}

func (metadata Metadata) MetaType() MetaType {
	return MetaTypeMetadata
}

// Метаданные сущности
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/workbook/#workbook-metadannye-metadannye-suschnosti

type MetaAttributesSharedWrapper struct {
	MetaAttributesWrapper
	CreateShared bool `json:"createShared"`
}

type MetaAttributesSharedStatesWrapper struct {
	MetaAttributesSharedWrapper
	States Slice[State] `json:"states"`
}

type MetaCharacteristicsWrapper struct {
	Meta            Meta                  `json:"meta"`            // Метаданные
	Characteristics Slice[Characteristic] `json:"characteristics"` // Коллекция всех созданных характеристик Модификаций
}

type MetaTagsWrapper struct {
	MetaAttributesSharedWrapper
	Tags Tags `json:"tags"`
}

//type MetadataAttributeSharedPriceTypes struct {
//	PriceTypes []struct {
//		Name string `json:"name,omitempty"`
//	} `json:"priceTypes"`
//	MetaAttributesSharedWrapper // Наименование
//}

type MetaNameShared struct {
	Meta         Meta   `json:"meta,omitempty"`
	Name         string `json:"name,omitempty"`
	CreateShared bool   `json:"createShared,omitempty"`
}

type MetadataCompanySettings struct {
	MetaAttributesWrapper
	CustomEntities Slice[MetaNameShared] `json:"customEntities"`
}

// MetadataService
// Сервис для работы с метаданными.
type MetadataService interface {
	Get(ctx context.Context, params *Params) (*Metadata, *resty.Response, error)
}

func NewMetadataService(client *Client) MetadataService {
	e := NewEndpoint(client, "entity/metadata")
	return newMainService[Metadata, any, any, any](e)
}
