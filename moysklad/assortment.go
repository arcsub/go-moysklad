package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
)

// Assortment Ассортимент.
//
// Код сущности: assortment
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-assortiment
type Assortment Slice[AssortmentPosition]

// Push добавляет элементы в конец среза.
func (assortment *Assortment) Push(elements ...AssortmentConverter) *Assortment {
	for _, element := range elements {
		if element != nil {
			*assortment = append(*assortment, element.AsAssortment())
		}
	}
	return assortment
}

// MetaType возвращает код сущности.
func (Assortment) MetaType() MetaType {
	return MetaTypeAssortment
}

// AssortmentPosition представляет позицию ассортимента.
//
// Позицией ассортимента могут быть:
//   - Product (Товар)
//   - Variant (Модификация)
//   - Bundle (Комплект)
//   - Service (Услуга)
//   - Consignment (Серия)
//
// Создать позицию можно с помощью [NewAssortmentPosition], передав в качестве аргумента объект,
// удовлетворяющий интерфейсу [AssortmentConverter].
type AssortmentPosition struct {
	Meta         Meta           `json:"meta"`                   // Метаданные сущности
	Code         string         `json:"code,omitempty"`         // Код сущности
	Description  string         `json:"description,omitempty"`  // Комментарий сущности
	ExternalCode string         `json:"externalCode,omitempty"` // Внешний код сущности
	Name         string         `json:"name,omitempty"`         // Наименование сущности
	Barcodes     Slice[Barcode] `json:"barcodes,omitempty"`     // Штрихкоды
	raw          []byte         // сырые данные для последующей конвертации в нужный тип
	AccountID    uuid.UUID      `json:"accountId,omitempty"` // ID учётной записи
	ID           uuid.UUID      `json:"id,omitempty"`        // ID сущности
}

// AssortmentConverter описывает метод, возвращающий [AssortmentPosition].
type AssortmentConverter interface {
	AsAssortment() *AssortmentPosition
}

// NewAssortmentPosition принимает в качестве аргумента объект, удовлетворяющий интерфейсу [AssortmentConverter].
//
// Возвращает [AssortmentPosition] с заполненным полем Meta.
func NewAssortmentPosition[T AssortmentConverter](entity T) *AssortmentPosition {
	return entity.AsAssortment()
}

// String реализует интерфейс [fmt.Stringer].
func (assortmentPosition *AssortmentPosition) String() string {
	return Stringify(assortmentPosition)
}

// MetaType возвращает код сущности.
func (assortmentPosition AssortmentPosition) MetaType() MetaType {
	return assortmentPosition.Meta.GetType()
}

// GetMeta возвращает Метаданные сущности.
func (assortmentPosition AssortmentPosition) GetMeta() Meta {
	return assortmentPosition.Meta
}

// Raw реализует интерфейс [RawMetaTyper].
func (assortmentPosition AssortmentPosition) Raw() []byte {
	return assortmentPosition.raw
}

// UnmarshalJSON реализует интерфейс [json.Unmarshaler].
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

// IsProduct возвращает true, если объект имеет код сущности [MetaTypeProduct].
func (assortmentPosition *AssortmentPosition) IsProduct() bool {
	return CheckType(assortmentPosition, MetaTypeProduct)
}

// IsVariant возвращает true, если объект имеет код сущности [MetaTypeVariant].
func (assortmentPosition *AssortmentPosition) IsVariant() bool {
	return CheckType(assortmentPosition, MetaTypeVariant)
}

// IsBundle возвращает true, если объект имеет код сущности [MetaTypeBundle].
func (assortmentPosition *AssortmentPosition) IsBundle() bool {
	return CheckType(assortmentPosition, MetaTypeBundle)
}

// IsService возвращает true, если объект имеет код сущности [MetaTypeService].
func (assortmentPosition *AssortmentPosition) IsService() bool {
	return CheckType(assortmentPosition, MetaTypeService)
}

// IsConsignment возвращает true, если объект имеет код сущности [MetaTypeConsignment].
func (assortmentPosition *AssortmentPosition) IsConsignment() bool {
	return CheckType(assortmentPosition, MetaTypeConsignment)
}

// AsProduct пытается привести объект к типу [Product].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [Product] или nil в случае неудачи.
func (assortmentPosition *AssortmentPosition) AsProduct() *Product {
	return UnmarshalAsType[Product](assortmentPosition)
}

// AsVariant пытается привести объект к типу [Variant].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [Variant] или nil в случае неудачи.
func (assortmentPosition *AssortmentPosition) AsVariant() *Variant {
	return UnmarshalAsType[Variant](assortmentPosition)
}

// AsBundle пытается привести объект к типу [Bundle].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [Bundle] или nil в случае неудачи.
func (assortmentPosition *AssortmentPosition) AsBundle() *Bundle {
	return UnmarshalAsType[Bundle](assortmentPosition)
}

// AsService пытается привести объект к типу [Service].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [Service] или nil в случае неудачи.
func (assortmentPosition *AssortmentPosition) AsService() *Service {
	return UnmarshalAsType[Service](assortmentPosition)
}

// AsConsignment пытается привести объект к типу [Consignment].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает: [Consignment] или nil в случае неудачи.
func (assortmentPosition *AssortmentPosition) AsConsignment() *Consignment {
	return UnmarshalAsType[Consignment](assortmentPosition)
}

// FilterBundle фильтрует позиции по типу [Bundle] (Комплект).
func (assortment Assortment) FilterBundle() Slice[Bundle] {
	return filterType[Bundle](assortment)
}

// FilterProduct фильтрует позиции по типу [Product] (Товар).
func (assortment Assortment) FilterProduct() Slice[Product] {
	return filterType[Product](assortment)
}

// FilterVariant фильтрует позиции по типу [Variant] (Модификация).
func (assortment Assortment) FilterVariant() Slice[Variant] {
	return filterType[Variant](assortment)
}

// FilterConsignment фильтрует позиции по типу [Consignment] (Серия).
func (assortment Assortment) FilterConsignment() Slice[Consignment] {
	return filterType[Consignment](assortment)
}

// FilterService фильтрует позиции по типу [Service] (Услуга).
func (assortment Assortment) FilterService() Slice[Service] {
	return filterType[Service](assortment)
}

// AssortmentSettings Настройки справочника.
//
// Код сущности: assortmentsettings
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-assortiment-nastrojki-sprawochnika
type AssortmentSettings struct {
	Meta            *Meta            `json:"meta,omitempty"`            // Метаданные Настроек справочника
	BarcodeRules    *BarcodeRules    `json:"barcodeRules,omitempty"`    // Настройки правил штрихкодов для сущностей справочника
	UniqueCodeRules *UniqueCodeRules `json:"uniqueCodeRules,omitempty"` // Настройки уникальности кода для сущностей справочника
	CreatedShared   *bool            `json:"createdShared,omitempty"`   // Создавать новые документы с меткой «Общий»
}

// GetMeta возвращает Метаданные Настроек справочника.
func (assortmentSettings AssortmentSettings) GetMeta() Meta {
	return Deref(assortmentSettings.Meta)
}

// GetBarcodeRules возвращает Настройки правил штрихкодов для сущностей справочника.
func (assortmentSettings AssortmentSettings) GetBarcodeRules() BarcodeRules {
	return Deref(assortmentSettings.BarcodeRules)
}

// GetUniqueCodeRules возвращает Настройки уникальности кода для сущностей справочника.
func (assortmentSettings AssortmentSettings) GetUniqueCodeRules() UniqueCodeRules {
	return Deref(assortmentSettings.UniqueCodeRules)
}

// GetCreatedShared возвращает true, если новые документы создаются с пометкой «Общий».
func (assortmentSettings AssortmentSettings) GetCreatedShared() bool {
	return Deref(assortmentSettings.CreatedShared)
}

// SetBarcodeRules устанавливает Настройки правил штрихкодов для сущностей справочника.
func (assortmentSettings *AssortmentSettings) SetBarcodeRules(barcodeRules *BarcodeRules) *AssortmentSettings {
	assortmentSettings.BarcodeRules = barcodeRules
	return assortmentSettings
}

// SetUniqueCodeRules устанавливает Настройки уникальности кода для сущностей справочника.
func (assortmentSettings *AssortmentSettings) SetUniqueCodeRules(uniqueCodeRules *UniqueCodeRules) *AssortmentSettings {
	assortmentSettings.UniqueCodeRules = uniqueCodeRules
	return assortmentSettings
}

// SetCreatedShared устанавливает значение создания новых документов с пометкой «Общий».
func (assortmentSettings *AssortmentSettings) SetCreatedShared(createdShared bool) *AssortmentSettings {
	assortmentSettings.CreatedShared = &createdShared
	return assortmentSettings
}

// String реализует интерфейс [fmt.Stringer].
func (assortmentSettings AssortmentSettings) String() string {
	return Stringify(assortmentSettings)
}

// MetaType возвращает код сущности.
func (AssortmentSettings) MetaType() MetaType {
	return MetaTypeAssortmentSettings
}

// BarcodeRules Настройки правил штрихкодов для сущностей справочника.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-assortiment-atributy-wlozhennyh-suschnostej-nastrojki-prawil-shtrihkodow-dlq-suschnostej-sprawochnika
type BarcodeRules struct {
	FillEAN13Barcode    *bool `json:"fillEAN13Barcode,omitempty"`    // Автоматически создавать штрихкод EAN13 для новых товаров, комплектов, модификаций и услуг
	WeightBarcode       *bool `json:"weightBarcode,omitempty"`       // Использовать префиксы штрихкодов для весовых товаров
	WeightBarcodePrefix *int  `json:"weightBarcodePrefix,omitempty"` // Префикс штрихкодов для весовых товаров. Возможные значения: число формата X или XX
}

// GetFillEAN13Barcode возвращает true, если штрихкод EAN13 для новых товаров, комплектов, модификаций и услуг создаётся автоматически.
func (barcodeRules BarcodeRules) GetFillEAN13Barcode() bool {
	return Deref(barcodeRules.FillEAN13Barcode)
}

// GetWeightBarcode возвращает true, если используются префиксы штрихкодов для весовых товаров.
func (barcodeRules BarcodeRules) GetWeightBarcode() bool {
	return Deref(barcodeRules.WeightBarcode)
}

// GetWeightBarcodePrefix возвращает Префикс штрихкодов для весовых товаров. Возможные значения: число формата X или XX.
func (barcodeRules BarcodeRules) GetWeightBarcodePrefix() int {
	return Deref(barcodeRules.WeightBarcodePrefix)
}

// SetFillEAN13Barcode устанавливает значение автоматического создания штрихкода EAN13 для новых товаров, комплектов, модификаций и услуг.
func (barcodeRules *BarcodeRules) SetFillEAN13Barcode(fillEAN13Barcode bool) *BarcodeRules {
	barcodeRules.FillEAN13Barcode = &fillEAN13Barcode
	return barcodeRules
}

// SetWeightBarcode устанавливает значение использования префиксов штрихкодов для весовых товаров.
func (barcodeRules *BarcodeRules) SetWeightBarcode(weightBarcode bool) *BarcodeRules {
	barcodeRules.WeightBarcode = &weightBarcode
	return barcodeRules
}

// SetWeightBarcodePrefix устанавливает Префикс штрихкодов для весовых товаров.
func (barcodeRules *BarcodeRules) SetWeightBarcodePrefix(weightBarcodePrefix int) *BarcodeRules {
	barcodeRules.WeightBarcodePrefix = &weightBarcodePrefix
	return barcodeRules
}

// String реализует интерфейс [fmt.Stringer].
func (barcodeRules BarcodeRules) String() string {
	return Stringify(barcodeRules)
}

// AssortmentResponse объект ответа на запрос получения ассортимента.
type AssortmentResponse struct {
	Context Context        `json:"context,omitempty"` // Информация о сотруднике, выполнившем запрос
	Rows    Assortment     `json:"rows,omitempty"`    // Список товаров, услуг, комплектов, модификаций и серий
	Meta    MetaCollection `json:"meta,omitempty"`    // Информация о контексте запроса
}

// String реализует интерфейс [fmt.Stringer].
func (assortmentResponse AssortmentResponse) String() string {
	return Stringify(assortmentResponse)
}

// BuyPrice Закупочная цена.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-towar-towary-atributy-wlozhennyh-suschnostej-zakupochnaq-cena
type BuyPrice struct {
	Value    *float64  `json:"value,omitempty"`    // Значение цены
	Currency *Currency `json:"currency,omitempty"` // Метаданные валюты
}

// GetValue возвращает Значение цены.
func (buyPrice BuyPrice) GetValue() float64 {
	return Deref(buyPrice.Value)
}

// GetCurrency возвращает Метаданные валюты.
func (buyPrice BuyPrice) GetCurrency() Currency {
	return Deref(buyPrice.Currency)
}

// SetValue устанавливает Значение цены.
func (buyPrice *BuyPrice) SetValue(value *float64) *BuyPrice {
	buyPrice.Value = value
	return buyPrice
}

// SetCurrency устанавливает Метаданные валюты.
func (buyPrice *BuyPrice) SetCurrency(currency *Currency) *BuyPrice {
	buyPrice.Currency = currency
	return buyPrice
}

// String реализует интерфейс [fmt.Stringer].
func (buyPrice BuyPrice) String() string {
	return Stringify(buyPrice)
}

// MinPrice Минимальная цена.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-towar-towary-atributy-wlozhennyh-suschnostej-minimal-naq-cena
type MinPrice struct {
	Value    *float64  `json:"value,omitempty"`    // Значение цены
	Currency *Currency `json:"currency,omitempty"` // Ссылка на валюту в формате Метаданных
}

// GetValue возвращает Значение цены.
func (minPrice MinPrice) GetValue() float64 {
	return Deref(minPrice.Value)
}

// GetCurrency возвращает Ссылку на валюту в формате Метаданных.
func (minPrice MinPrice) GetCurrency() Currency {
	return Deref(minPrice.Currency)
}

// SetValue устанавливает Значение цены.
func (minPrice *MinPrice) SetValue(value float64) *MinPrice {
	minPrice.Value = &value
	return minPrice
}

// SetCurrency устанавливает Ссылку на валюту в формате Метаданных.
func (minPrice *MinPrice) SetCurrency(currency *Currency) *MinPrice {
	if currency != nil {
		minPrice.Currency = currency.Clean()
	}
	return minPrice
}

// String реализует интерфейс [fmt.Stringer].
func (minPrice MinPrice) String() string {
	return Stringify(minPrice)
}

// SalePrice Цена продажи
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-towar-towary-atributy-wlozhennyh-suschnostej-ceny-prodazhi
type SalePrice struct {
	Value     *float64   `json:"value,omitempty"`     // Значение цены
	Currency  *Currency  `json:"currency,omitempty"`  // Ссылка на валюту в формате Метаданных
	PriceType *PriceType `json:"priceType,omitempty"` // Тип цены
}

// GetValue возвращает Значение цены.
func (salePrice SalePrice) GetValue() float64 {
	return Deref(salePrice.Value)
}

// GetCurrency возвращает Ссылку на валюту в формате Метаданных.
func (salePrice SalePrice) GetCurrency() Currency {
	return Deref(salePrice.Currency)
}

// GetPriceType возвращает Тип цены.
func (salePrice SalePrice) GetPriceType() PriceType {
	return Deref(salePrice.PriceType)
}

// SetValue устанавливает Значение цены.
func (salePrice *SalePrice) SetValue(value float64) *SalePrice {
	salePrice.Value = &value
	return salePrice
}

// SetCurrency устанавливает Ссылку на валюту в формате Метаданных.
func (salePrice *SalePrice) SetCurrency(currency *Currency) *SalePrice {
	if currency != nil {
		salePrice.Currency = currency.Clean()
	}
	return salePrice
}

// SetPriceType устанавливает Тип цены.
func (salePrice *SalePrice) SetPriceType(priceType *PriceType) *SalePrice {
	if priceType != nil {
		salePrice.PriceType = priceType.Clean()
	}
	return salePrice
}

// String реализует интерфейс [fmt.Stringer].
func (salePrice SalePrice) String() string {
	return Stringify(salePrice)
}

// Stock Остатки и себестоимость в позициях документов.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-obschie-swedeniq-ostatki-i-sebestoimost-w-poziciqh-dokumentow
type Stock struct {
	Cost      float64 `json:"cost"`      // Себестоимость
	Quantity  float64 `json:"quantity"`  // Количество
	Reserve   float64 `json:"reserve"`   // Резерв
	InTransit float64 `json:"intransit"` // Ожидание
	Available float64 `json:"available"` // Доступно
}

// String реализует интерфейс [fmt.Stringer].
func (stock Stock) String() string {
	return Stringify(stock)
}

// PaymentItem Признак предмета расчета.
//
// Возможные варианты:
//   - PaymentItemGood                – Товар
//   - PaymentItemExcisableGood       – Подакцизный товар
//   - PaymentItemCompoundPaymentItem – Составной предмет расчета
//   - PaymentItemAnotherPaymentItem  – Иной предмет расчета
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-towar-towary-atributy-suschnosti-priznak-predmeta-rascheta
type PaymentItem string

const (
	PaymentItemGood                PaymentItem = "GOOD"                  // Товар
	PaymentItemExcisableGood       PaymentItem = "EXCISABLE_GOOD"        // Подакцизный товар
	PaymentItemCompoundPaymentItem PaymentItem = "COMPOUND_PAYMENT_ITEM" // Составной предмет расчета
	PaymentItemAnotherPaymentItem  PaymentItem = "ANOTHER_PAYMENT_ITEM"  // Иной предмет расчета
)

// TrackingType Тип маркируемой продукции.
//
// Возможные варианты:
//   - TrackingTypeBeerAlcohol 		– Пиво и слабоалкогольная продукция
//   - TrackingTypeElectronics 		– Фотокамеры и лампы-вспышки
//   - TrackingTypeFoodSupplement 	– Биологически активные добавки к пище
//   - TrackingTypeClothes 			– Тип маркировки "Одежда"
//   - TrackingTypeLinens 			– Тип маркировки "Постельное белье"
//   - TrackingTypeMedicalDevices 	– Медизделия и кресла-коляски
//   - TrackingTypeMilk 			– Молочная продукция
//   - TrackingTypeNcp 				– Никотиносодержащая продукция
//   - TrackingTypeNotTracked 		– Без маркировки
//   - TrackingTypeOtp 				– Альтернативная табачная продукция
//   - TrackingTypePerfumery 		– Духи и туалетная вода
//   - TrackingTypeSanitizer 		– Антисептики
//   - TrackingTypeShoes 			– Тип маркировки "Обувь"
//   - TrackingTypeTires 			– Шины и покрышки
//   - TrackingTypeTobacco 			– Тип маркировки "Табак"
//   - TrackingTypeWater 			– Упакованная вода
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-towar-towary-atributy-suschnosti-tip-markiruemoj-produkcii
type TrackingType string

const (
	TrackingTypeBeerAlcohol    TrackingType = "BEER_ALCOHOL"    // Пиво и слабоалкогольная продукция
	TrackingTypeElectronics    TrackingType = "ELECTRONICS"     // Фотокамеры и лампы-вспышки
	TrackingTypeFoodSupplement TrackingType = "FOOD_SUPPLEMENT" // Биологически активные добавки к пище
	TrackingTypeClothes        TrackingType = "LP_CLOTHES"      // Тип маркировки "Одежда"
	TrackingTypeLinens         TrackingType = "LP_LINENS"       // Тип маркировки "Постельное белье"
	TrackingTypeMedicalDevices TrackingType = "MEDICAL_DEVICES" // Медизделия и кресла-коляски
	TrackingTypeMilk           TrackingType = "MILK"            // Молочная продукция
	TrackingTypeNcp            TrackingType = "NCP"             // Никотиносодержащая продукция
	TrackingTypeNotTracked     TrackingType = "NOT_TRACKED"     // Без маркировки
	TrackingTypeOtp            TrackingType = "OTP"             // Альтернативная табачная продукция
	TrackingTypePerfumery      TrackingType = "PERFUMERY"       // Духи и туалетная вода
	TrackingTypeSanitizer      TrackingType = "SANITIZER"       // Антисептики
	TrackingTypeShoes          TrackingType = "SHOES"           // Тип маркировки "Обувь"
	TrackingTypeTires          TrackingType = "TIRES"           // Шины и покрышки
	TrackingTypeTobacco        TrackingType = "TOBACCO"         // Тип маркировки "Табак"
	TrackingTypeWater          TrackingType = "WATER"           // Упакованная вода
)

// TaxSystem Код системы налогообложения.
//
// Возможные варианты:
//   - TaxSystemGeneral                 – ОСН
//   - TaxSystemSimplifiedIncome        – УСН. Доход
//   - TaxSystemSimplifiedIncomeOutcome – УСН. Доход-Расход
//   - TaxSystemUnifiedAgricultural     – ЕСХН
//   - TaxSystemPresumptive             – ЕНВД
//   - TaxSystemPatentBased             – Патент
//   - TaxSystemSameAsGroup             – Совпадает с группой
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-towar-towary-atributy-suschnosti-kod-sistemy-nalogooblozheniq
type TaxSystem string

const (
	TaxSystemGeneral                 TaxSystem = "GENERAL_TAX_SYSTEM"                   // ОСН
	TaxSystemSimplifiedIncome        TaxSystem = "SIMPLIFIED_TAX_SYSTEM_INCOME"         // УСН. Доход
	TaxSystemSimplifiedIncomeOutcome TaxSystem = "SIMPLIFIED_TAX_SYSTEM_INCOME_OUTCOME" // УСН. Доход-Расход
	TaxSystemUnifiedAgricultural     TaxSystem = "UNIFIED_AGRICULTURAL_TAX"             // ЕСХН
	TaxSystemPresumptive             TaxSystem = "PRESUMPTIVE_TAX_SYSTEM"               // ЕНВД
	TaxSystemPatentBased             TaxSystem = "PATENT_BASED"                         // Патент
	TaxSystemSameAsGroup             TaxSystem = "TAX_SYSTEM_SAME_AS_GROUP"             // Совпадает с группой
)

// AssortmentService методы сервиса для работы с ассортиментом.
type AssortmentService interface {
	// Get выполняет запрос на получение всех товаров, услуг, комплектов, модификаций и серий в виде списка.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект AssortmentResponse.
	Get(ctx context.Context, params ...*Params) (*AssortmentResponse, *resty.Response, error)

	// GetListAsync выполняет асинхронный запрос на получение всех товаров, услуг, комплектов, модификаций и серий в виде списка.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает готовый сервис AsyncResultService для обработки данного запроса.
	GetListAsync(ctx context.Context, params ...*Params) (AsyncResultService[AssortmentResponse], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление позиций в Ассортименте.
	// Принимает контекст и множество объектов, реализующих интерфейс AssortmentConverter.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...AssortmentConverter) (*DeleteManyResponse, *resty.Response, error)

	// GetSettings выполняет запрос на получение настроек справочника ассортимента.
	// Принимает контекст.
	// Возвращает объект AssortmentSettings.
	GetSettings(ctx context.Context) (*AssortmentSettings, *resty.Response, error)

	// UpdateSettings выполняет запрос на изменение метаданных справочника ассортимента.
	// Принимает контекст и объект AssortmentSettings.
	// Возвращает изменённый объект AssortmentSettings.
	UpdateSettings(ctx context.Context, settings *AssortmentSettings) (*AssortmentSettings, *resty.Response, error)
}

type assortmentService struct {
	Endpoint
}

const (
	EndpointAssortment         = EndpointEntity + string(MetaTypeAssortment)
	EndpointAssortmentSettings = EndpointAssortment + "/settings"
)

func (service *assortmentService) Get(ctx context.Context, params ...*Params) (*AssortmentResponse, *resty.Response, error) {
	return NewRequestBuilder[AssortmentResponse](service.client, service.uri).SetParams(params...).Get(ctx)
}

func (service *assortmentService) GetListAsync(ctx context.Context, params ...*Params) (AsyncResultService[AssortmentResponse], *resty.Response, error) {
	p := NewParams()
	if len(params) > 0 {
		p = params[0]
	}
	p.WithAsync()
	_, resp, err := NewRequestBuilder[any](service.client, service.uri).SetParams(p).Get(ctx)
	if err != nil {
		return nil, resp, nil
	}
	async := NewAsyncResultService[AssortmentResponse](service.client, resp)
	return async, resp, err
}

func (service *assortmentService) DeleteMany(ctx context.Context, entities ...AssortmentConverter) (*DeleteManyResponse, *resty.Response, error) {
	var mw = make([]MetaWrapper, 0, len(entities))
	for _, entity := range entities {
		if entity != nil {
			mw = append(mw, entity.AsAssortment().GetMeta().Wrap())
		}
	}
	return NewRequestBuilder[DeleteManyResponse](service.client, service.uri).Post(ctx, mw)
}

func (service *assortmentService) GetSettings(ctx context.Context) (*AssortmentSettings, *resty.Response, error) {
	return NewRequestBuilder[AssortmentSettings](service.client, EndpointAssortmentSettings).Get(ctx)
}

func (service *assortmentService) UpdateSettings(ctx context.Context, settings *AssortmentSettings) (*AssortmentSettings, *resty.Response, error) {
	return NewRequestBuilder[AssortmentSettings](service.client, EndpointAssortmentSettings).Put(ctx, settings)
}

// NewAssortmentService принимает [Client] и возвращает сервис для работы с ассортиментом.
func NewAssortmentService(client *Client) AssortmentService {
	return &assortmentService{NewEndpoint(client, EndpointAssortment)}
}
