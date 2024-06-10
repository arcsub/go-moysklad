package moysklad

import (
	"context"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Assortment Ассортимент.
// Ключевое слово: assortment
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-assortiment
type Assortment Slice[AssortmentPosition]

func (assortment Assortment) MetaType() MetaType {
	return MetaTypeAssortment
}

// AssortmentPosition представляет позицию ассортимента.
//
// Создать позицию можно с помощью NewAssortmentPosition, передав в качестве аргумента объект,
// удовлетворяющий интерфейсу AssortmentType.
type AssortmentPosition struct {
	Meta         Meta            `json:"meta"`
	Code         string          `json:"code,omitempty"`
	Description  string          `json:"description,omitempty"`
	ExternalCode string          `json:"externalCode,omitempty"`
	Name         string          `json:"name,omitempty"`
	Barcodes     Slice[Barcode]  `json:"barcodes,omitempty"`
	raw          json.RawMessage // сырые данные для последующей десериализации в нужный тип
	AccountID    uuid.UUID       `json:"accountId,omitempty"`
	ID           uuid.UUID       `json:"id,omitempty"`
}

// AssortmentType описывает типы, которые входят в состав ассортимента.
type AssortmentType interface {
	Product | Variant | Bundle | Service | Consignment
	MetaOwner
}

// NewAssortmentPosition принимает в качестве аргумента объект, удовлетворяющий интерфейсу AssortmentType.
//
// Возвращает позицию ассортимента с заполненным полем Meta.
func NewAssortmentPosition[T AssortmentType](entity T) *AssortmentPosition {
	return &AssortmentPosition{Meta: entity.GetMeta()}
}

func (assortmentPosition *AssortmentPosition) String() string {
	return Stringify(assortmentPosition.Meta)
}

// MetaType удовлетворяет интерфейсу MetaTyper
func (assortmentPosition AssortmentPosition) MetaType() MetaType {
	return assortmentPosition.Meta.GetType()
}

// Raw удовлетворяет интерфейсу RawMetaTyper
func (assortmentPosition AssortmentPosition) Raw() json.RawMessage {
	return assortmentPosition.raw
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (assortmentPosition *AssortmentPosition) UnmarshalJSON(data []byte) error {
	type alias AssortmentPosition
	var t alias
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	t.raw = data
	*assortmentPosition = AssortmentPosition(t)
	return nil
}

// AsProduct десериализует сырые данные в тип *Product
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (assortmentPosition *AssortmentPosition) AsProduct() *Product {
	return unmarshalAsType[Product](assortmentPosition)
}

// AsVariant десериализует сырые данные в тип *Variant
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (assortmentPosition *AssortmentPosition) AsVariant() *Variant {
	return unmarshalAsType[Variant](assortmentPosition)
}

// AsBundle десериализует сырые данные в тип *Bundle
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (assortmentPosition *AssortmentPosition) AsBundle() *Bundle {
	return unmarshalAsType[Bundle](assortmentPosition)
}

// AsService десериализует сырые данные в тип *Service
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (assortmentPosition *AssortmentPosition) AsService() *Service {
	return unmarshalAsType[Service](assortmentPosition)
}

// AsConsignment десериализует сырые данные в тип *Consignment
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (assortmentPosition *AssortmentPosition) AsConsignment() *Consignment {
	return unmarshalAsType[Consignment](assortmentPosition)
}

// FilterBundle фильтрует позиции по типу Bundle (Комплект)
func (assortment Assortment) FilterBundle() Slice[Bundle] {
	return filterType[Bundle](assortment)
}

// FilterProduct фильтрует позиции по типу Product (Товар)
func (assortment Assortment) FilterProduct() Slice[Product] {
	return filterType[Product](assortment)
}

// FilterVariant фильтрует позиции по типу Variant (Модификация)
func (assortment Assortment) FilterVariant() Slice[Variant] {
	return filterType[Variant](assortment)
}

// FilterConsignment фильтрует позиции по типу Consignment (Серия)
func (assortment Assortment) FilterConsignment() Slice[Consignment] {
	return filterType[Consignment](assortment)
}

// FilterService фильтрует позиции по типу Service (Услуга)
func (assortment Assortment) FilterService() Slice[Service] {
	return filterType[Service](assortment)
}

// AssortmentSettings Настройки справочника.
// Ключевое слово: assortmentsettings
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-assortiment-nastrojki-sprawochnika
type AssortmentSettings struct {
	Meta            *Meta            `json:"meta,omitempty"`            // Метаданные Настроек справочника
	BarcodeRules    *BarcodeRules    `json:"barcodeRules,omitempty"`    // Настройки правил штрихкодов для сущностей справочника
	UniqueCodeRules *UniqueCodeRules `json:"uniqueCodeRules,omitempty"` // Настройки уникальности кода для сущностей справочника
	CreatedShared   *bool            `json:"createdShared,omitempty"`   // Создавать новые документы с меткой «Общий»
}

func (assortmentSettings AssortmentSettings) GetMeta() Meta {
	return Deref(assortmentSettings.Meta)
}

func (assortmentSettings AssortmentSettings) GetBarcodeRules() BarcodeRules {
	return Deref(assortmentSettings.BarcodeRules)
}

func (assortmentSettings AssortmentSettings) GetUniqueCodeRules() UniqueCodeRules {
	return Deref(assortmentSettings.UniqueCodeRules)
}

func (assortmentSettings AssortmentSettings) GetCreatedShared() bool {
	return Deref(assortmentSettings.CreatedShared)
}

func (assortmentSettings *AssortmentSettings) SetBarcodeRules(barcodeRules *BarcodeRules) *AssortmentSettings {
	assortmentSettings.BarcodeRules = barcodeRules
	return assortmentSettings
}

func (assortmentSettings *AssortmentSettings) SetUniqueCodeRules(uniqueCodeRules *UniqueCodeRules) *AssortmentSettings {
	assortmentSettings.UniqueCodeRules = uniqueCodeRules
	return assortmentSettings
}

func (assortmentSettings *AssortmentSettings) SetCreatedShared(createdShared bool) *AssortmentSettings {
	assortmentSettings.CreatedShared = &createdShared
	return assortmentSettings
}

func (assortmentSettings AssortmentSettings) String() string {
	return Stringify(assortmentSettings)
}

func (assortmentSettings AssortmentSettings) MetaType() MetaType {
	return MetaTypeAssortmentSettings
}

// BarcodeRules Настройки правил штрихкодов для сущностей справочника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-assortiment-atributy-wlozhennyh-suschnostej-nastrojki-prawil-shtrihkodow-dlq-suschnostej-sprawochnika
type BarcodeRules struct {
	FillEAN13Barcode    *bool `json:"fillEAN13Barcode,omitempty"`    // Автоматически создавать штрихкод EAN13 для новых товаров, комплектов, модификаций и услуг
	WeightBarcode       *bool `json:"weightBarcode,omitempty"`       // Использовать префиксы штрихкодов для весовых товаров
	WeightBarcodePrefix *int  `json:"weightBarcodePrefix,omitempty"` // Префикс штрихкодов для весовых товаров. Возможные значения: число формата X или XX
}

func (barcodeRules BarcodeRules) GetFillEAN13Barcode() bool {
	return Deref(barcodeRules.FillEAN13Barcode)
}

func (barcodeRules BarcodeRules) GetWeightBarcode() bool {
	return Deref(barcodeRules.WeightBarcode)
}

func (barcodeRules BarcodeRules) GetWeightBarcodePrefix() int {
	return Deref(barcodeRules.WeightBarcodePrefix)
}

func (barcodeRules *BarcodeRules) SetFillEAN13Barcode(fillEAN13Barcode bool) *BarcodeRules {
	barcodeRules.FillEAN13Barcode = &fillEAN13Barcode
	return barcodeRules
}

func (barcodeRules *BarcodeRules) SetWeightBarcode(weightBarcode bool) *BarcodeRules {
	barcodeRules.WeightBarcode = &weightBarcode
	return barcodeRules
}

func (barcodeRules *BarcodeRules) SetWeightBarcodePrefix(weightBarcodePrefix int) *BarcodeRules {
	barcodeRules.WeightBarcodePrefix = &weightBarcodePrefix
	return barcodeRules
}

func (barcodeRules BarcodeRules) String() string {
	return Stringify(barcodeRules)
}

type AssortmentResponse struct {
	Context Context        `json:"context,omitempty"`
	Rows    Assortment     `json:"rows,omitempty"`
	Meta    MetaCollection `json:"meta,omitempty"`
}

func (assortmentResponse AssortmentResponse) String() string {
	return Stringify(assortmentResponse)
}

// PaymentItem Признак предмета расчета.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-towar-towary-atributy-suschnosti-priznak-predmeta-rascheta
type PaymentItem string

const (
	PaymentItemGood                PaymentItem = "GOOD"                  // Товар
	PaymentItemExcisableGood       PaymentItem = "EXCISABLE_GOOD"        // Подакцизный товар
	PaymentItemCompoundPaymentItem PaymentItem = "COMPOUND_PAYMENT_ITEM" // Составной предмет расчета
	PaymentItemAnotherPaymentItem  PaymentItem = "ANOTHER_PAYMENT_ITEM"  // Иной предмет расчета
)

// TrackingType Тип маркируемой продукции.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-towar-towary-atributy-suschnosti-tip-markiruemoj-produkcii
type TrackingType string

const (
	TrackingTypeFoodSupplement TrackingType = "FOOD_SUPPLEMENT" // [22-12-2023]
	TrackingTypeBeerAlcohol    TrackingType = "BEER_ALCOHOL"    // [26-10-2023]
	TrackingTypeElectronics    TrackingType = "ELECTRONICS"
	TrackingTypeClothes        TrackingType = "LP_CLOTHES"
	TrackingTypeLinens         TrackingType = "LP_LINENS"
	TrackingTypeMilk           TrackingType = "MILK"
	TrackingTypeNcp            TrackingType = "NCP"
	TrackingTypeNotTracked     TrackingType = "NOT_TRACKED"
	TrackingTypeOtp            TrackingType = "OTP"
	TrackingTypePerfumery      TrackingType = "PERFUMERY"
	TrackingTypeShoes          TrackingType = "SHOES"
	TrackingTypeTires          TrackingType = "TIRES"
	TrackingTypeTobacco        TrackingType = "TOBACCO"
	TrackingTypeWater          TrackingType = "WATER"
)

// GoodTaxSystem Код системы налогообложения.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-towar-towary-atributy-suschnosti-kod-sistemy-nalogooblozheniq
type GoodTaxSystem string

const (
	GoodTaxSystemGeneralTaxSystem                 GoodTaxSystem = "GENERAL_TAX_SYSTEM"                   // ОСН
	GoodTaxSystemSimplifiedTaxSystemIncome        GoodTaxSystem = "SIMPLIFIED_TAX_SYSTEM_INCOME"         // УСН. Доход
	GoodTaxSystemSimplifiedTaxSystemIncomeOutcome GoodTaxSystem = "SIMPLIFIED_TAX_SYSTEM_INCOME_OUTCOME" // УСН. Доход-Расход
	GoodTaxSystemUnifiedAgriculturalTax           GoodTaxSystem = "UNIFIED_AGRICULTURAL_TAX"             // ЕСХН
	GoodTaxSystemPresumptiveTaxSystem             GoodTaxSystem = "PRESUMPTIVE_TAX_SYSTEM"               // ЕНВД
	GoodTaxSystemPatentBased                      GoodTaxSystem = "PATENT_BASED"                         // Патент
	GoodTaxSystemSameAsGroup                      GoodTaxSystem = "TAX_SYSTEM_SAME_AS_GROUP"             // Совпадает с группой
)

// TaxSystem Код системы налогообложения по умолчанию.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-kod-sistemy-nalogooblozheniq-po-umolchaniu
type TaxSystem string

const (
	GeneralTaxSystem                 TaxSystem = "GENERAL_TAX_SYSTEM"                   // ОСН
	SimplifiedTaxSystemIncome        TaxSystem = "SIMPLIFIED_TAX_SYSTEM_INCOME"         // УСН. Доход
	SimplifiedTaxSystemIncomeOutcome TaxSystem = "SIMPLIFIED_TAX_SYSTEM_INCOME_OUTCOME" // УСН. Доход-Расход
	UnifiedAgriculturalTax           TaxSystem = "UNIFIED_AGRICULTURAL_TAX"             // ЕСХН
	PresumptiveTaxSystem             TaxSystem = "PRESUMPTIVE_TAX_SYSTEM"               // ЕНВД
	PatentBased                      TaxSystem = "PATENT_BASED"                         // Патент
)

// AssortmentService
// Сервис для работы с ассортиментом.
type AssortmentService interface {
	Get(ctx context.Context, params *Params) (*AssortmentResponse, *resty.Response, error)
	GetAsync(ctx context.Context) (AsyncResultService[AssortmentResponse], *resty.Response, error)
	DeleteMany(ctx context.Context, entities *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	GetSettings(ctx context.Context) (*AssortmentSettings, *resty.Response, error)
	UpdateSettings(ctx context.Context, settings *AssortmentSettings) (*AssortmentSettings, *resty.Response, error)
	GetEmbeddedTemplates(ctx context.Context) (*List[EmbeddedTemplate], *resty.Response, error)
	GetEmbeddedTemplateByID(ctx context.Context, id uuid.UUID) (*EmbeddedTemplate, *resty.Response, error)
	GetCustomTemplates(ctx context.Context) (*List[CustomTemplate], *resty.Response, error)
	GetCustomTemplateByID(ctx context.Context, id uuid.UUID) (*CustomTemplate, *resty.Response, error)
}

func NewAssortmentService(client *Client) AssortmentService {
	e := NewEndpoint(client, "entity/assortment")
	return newMainService[AssortmentResponse, any, any, AssortmentSettings](e)
}
