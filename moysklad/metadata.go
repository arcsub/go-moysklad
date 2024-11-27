package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
)

// Metadata Глобальные метаданные.
//
// Код сущности: metadata
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-metadannye-metadannye-suschnosti
type Metadata struct {
	CompanySettings           MetadataCompanySettings               `json:"companysettings"`
	BonusProgram              MetaAttributesWrapper                 `json:"bonusprogram"`
	Consignment               MetaAttributesWrapper                 `json:"consignment"`
	ProcessingPlan            MetaAttributesWrapper                 `json:"processingplan"`
	Assortment                MetaWrapper                           `json:"assortment"`
	ProductFolder             MetaWrapper                           `json:"productfolder"`
	ProcessingPlanFolder      MetaWrapper                           `json:"processingplanfolder"`
	Application               MetaWrapper                           `json:"application"`
	InvoiceIn                 MetaAttributesSharedWrapper           `json:"invoicein"`
	Processing                MetaAttributesSharedWrapper           `json:"processing"`
	Supply                    MetaAttributesSharedWrapper           `json:"supply"`
	CommissionReportIn        MetaAttributesSharedWrapper           `json:"commissionreportin"`
	CommissionReportOut       MetaAttributesSharedWrapper           `json:"commissionreportout"`
	PrepaymentReturn          MetaAttributesSharedWrapper           `json:"prepaymentreturn"`
	RetailShift               MetaAttributesSharedWrapper           `json:"retailshift"`
	Loss                      MetaAttributesSharedWrapper           `json:"loss"`
	FactureOut                MetaAttributesSharedWrapper           `json:"factureout"`
	PurchaseOrder             MetaAttributesSharedWrapper           `json:"purchaseorder"`
	CrptOrder                 MetaAttributesSharedWrapper           `json:"crptorder"`
	Move                      MetaAttributesSharedWrapper           `json:"move"`
	Employee                  MetaAttributesSharedWrapper           `json:"employee"`
	InvoiceOut                MetaAttributesSharedWrapper           `json:"invoiceout"`
	RetailDemand              MetaAttributesSharedWrapper           `json:"retaildemand"`
	SalesReturn               MetaAttributesSharedWrapper           `json:"salesreturn"`
	InternalOrder             MetaAttributesSharedWrapper           `json:"internalorder"`
	Organization              MetaAttributesSharedWrapper           `json:"organization"`
	Inventory                 MetaAttributesSharedWrapper           `json:"inventory"`
	Demand                    MetaAttributesSharedWrapper           `json:"demand"`
	PaymentOut                MetaAttributesSharedWrapper           `json:"paymentout"`
	CounterpartyAdjustment    MetaAttributesSharedWrapper           `json:"counterpartyadjustment"`
	ProductionTask            MetaAttributesSharedWrapper           `json:"productiontask"`
	RetailDrawerCashIn        MetaAttributesSharedWrapper           `json:"retaildrawercashin"`
	CashOut                   MetaAttributesSharedWrapper           `json:"cashout"`
	PurchaseReturn            MetaAttributesSharedWrapper           `json:"purchasereturn"`
	RetailDrawerCashOut       MetaAttributesSharedWrapper           `json:"retaildrawercashout"`
	ProcessingOrder           MetaAttributesSharedWrapper           `json:"processingorder"`
	Project                   MetaAttributesSharedWrapper           `json:"project"`
	FactureOn                 MetaAttributesSharedWrapper           `json:"facturein"`
	CashIn                    MetaAttributesSharedWrapper           `json:"cashin"`
	Contract                  MetaAttributesSharedWrapper           `json:"contract"`
	PaymentIn                 MetaAttributesSharedWrapper           `json:"paymentin"`
	PriceList                 MetaAttributesSharedWrapper           `json:"pricelist"`
	BonusTransaction          MetaAttributesSharedWrapper           `json:"bonustransaction"`
	Store                     MetaAttributesSharedWrapper           `json:"store"`
	RetailSalesReturn         MetaAttributesSharedWrapper           `json:"retailsalesreturn"`
	Enter                     MetaAttributesSharedWrapper           `json:"enter"`
	ProductionsTageCompletion MetaAttributesSharedWrapper           `json:"productionstagecompletion"`
	Service                   MetaAttributesSharedWrapper           `json:"service"`
	Prepayment                MetaAttributesSharedWrapper           `json:"prepayment"`
	Product                   MetaAttributesSharedWrapper           `json:"product"`
	Bundle                    MetaAttributesSharedWrapper           `json:"bundle"`
	Counterparty              MetaAttributesStatesSharedTagsWrapper `json:"counterparty"`
	CustomerOrder             MetaAttributesStatesSharedWrapper     `json:"customerorder"`
	Variant                   MetaCharacteristicsWrapper            `json:"variant"`
}

// MetaType возвращает код сущности.
func (Metadata) MetaType() MetaType {
	return MetaTypeMetadata
}

// MetaAttributesWrapper содержит [Meta] и коллекцию доп. полей.
type MetaAttributesWrapper struct {
	Meta       Meta        `json:"meta"`       // Метаданные
	Attributes MetaWrapper `json:"attributes"` // Коллекция доп. полей
}

// MetaAttributesSharedWrapper содержит [Meta], коллекцию доп. полей и флаг создания новых сущностей с меткой "Общий".
type MetaAttributesSharedWrapper struct {
	MetaAttributesWrapper
	CreateShared bool `json:"createShared"` // Создавать новые сущности с меткой "Общий"
}

// MetaAttributesStatesSharedWrapper содержит [Meta], коллекцию доп. полей, коллекцию статусов и флаг создания новых сущностей с меткой "Общий".
type MetaAttributesStatesSharedWrapper struct {
	MetaAttributesSharedWrapper
	States Slice[StatesElement] `json:"states"` // Массив статусов
}

// MetaAttributesStatesSharedTagsWrapper содержит [Meta], коллекцию доп. полей, коллекцию статусов, группы контрагентов и флаг создания новых сущностей с меткой "Общий".
type MetaAttributesStatesSharedTagsWrapper struct {
	MetaAttributesSharedWrapper
	Tags Slice[string] `json:"tags"`
}

// StatesElement элемент списка массива статусов.
type StatesElement struct {
	Meta       Meta     `json:"meta"`       // Метаданные статуса
	EntityType MetaType `json:"entityType"` // Код сущности, к которой относится Статус (ключевое слово в рамках JSON API)
	StateType  string   `json:"stateType"`  // Тип Статуса
	Color      int      `json:"color"`      // Цвет Статуса
}

// MetaCharacteristicsWrapper содержит метаданные и коллекцию всех созданных характеристик Модификаций
type MetaCharacteristicsWrapper struct {
	Meta            Meta                  `json:"meta"`            // Метаданные
	Characteristics Slice[Characteristic] `json:"characteristics"` // Коллекция всех созданных характеристик Модификаций
}

type MetaNameShared struct {
	Meta         Meta   `json:"meta,omitempty"`
	Name         string `json:"name,omitempty"`
	CreateShared bool   `json:"createShared,omitempty"`
}

// MetadataCompanySettings объект метаданных компании.
type MetadataCompanySettings struct {
	MetaAttributesWrapper
	CustomEntities Slice[MetaNameShared] `json:"customEntities"`
}

// MetadataService описывает метод сервиса для работы с глобальными метаданными.
type MetadataService interface {
	// Get выполняет запрос на получение метаданных.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект Metadata.
	Get(ctx context.Context, params ...func(*Params)) (*Metadata, *resty.Response, error)
}

const (
	EndpointMetadata = EndpointEntity + string(MetaTypeMetadata)
)

type metadataService struct {
	Endpoint
}

func (service *metadataService) Get(ctx context.Context, params ...func(*Params)) (*Metadata, *resty.Response, error) {
	return NewRequestBuilder[Metadata](service.client, service.uri).SetParams(params).Get(ctx)
}

// NewMetadataService принимает [Client] и возвращает сервис для работы с глобальными метаданными.
func NewMetadataService(client *Client) MetadataService {
	return &metadataService{NewEndpoint(client, EndpointMetadata)}
}
