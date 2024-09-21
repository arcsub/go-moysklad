package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"time"
)

// Bundle Комплект.
//
// Код сущности: bundle
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-komplekt
type Bundle struct {
	Volume              *float64                    `json:"volume,omitempty"`              // Объем
	SyncID              *uuid.UUID                  `json:"syncId,omitempty"`              // ID синхронизации
	Code                *string                     `json:"code,omitempty"`                // Код Комплекта
	Description         *string                     `json:"description,omitempty"`         // Описание Комплекта
	ExternalCode        *string                     `json:"externalCode,omitempty"`        // Внешний код Комплекта
	ID                  *uuid.UUID                  `json:"id,omitempty"`                  // ID Комплекта
	Meta                *Meta                       `json:"meta,omitempty"`                // Метаданные Комплекта
	Name                *string                     `json:"name,omitempty"`                // Наименование Комплекта
	Archived            *bool                       `json:"archived,omitempty"`            // Добавлен ли Комплект в архив
	Article             *string                     `json:"article,omitempty"`             // Артикул
	Images              *MetaArray[Image]           `json:"images,omitempty"`              // Массив метаданных Изображений (Максимальное количество изображений - 10)
	Components          *MetaArray[BundleComponent] `json:"components,omitempty"`          // Массив компонентов Комплекта
	Country             *NullValue[Country]         `json:"country,omitempty"`             // Метаданные Страны
	DiscountProhibited  *bool                       `json:"discountProhibited"`            // Признак запрета скидок
	EffectiveVat        *int                        `json:"effectiveVat,omitempty"`        // Реальный НДС %
	EffectiveVatEnabled *bool                       `json:"effectiveVatEnabled,omitempty"` // Дополнительный признак для определения разграничения реального НДС = 0 или "без НДС". (effectiveVat = 0, effectiveVatEnabled = false) -> "без НДС", (effectiveVat = 0, effectiveVatEnabled = true) -> 0%.
	Files               *MetaArray[File]            `json:"files,omitempty"`               // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group               *Group                      `json:"group,omitempty"`               // Отдел сотрудника
	Vat                 *int                        `json:"vat,omitempty"`                 // НДС %
	MinPrice            *NullValue[MinPrice]        `json:"minPrice,omitempty"`            // Минимальная цена
	Overhead            *NullValue[BundleOverhead]  `json:"overhead,omitempty"`            // Дополнительные расходы
	Owner               *Employee                   `json:"owner,omitempty"`               // Метаданные владельца (Сотрудника)
	PartialDisposal     *bool                       `json:"partialDisposal,omitempty"`     // Управление состоянием частичного выбытия маркированного товара. «true» - возможность включена.
	PathName            *string                     `json:"pathName,omitempty"`            // Наименование группы, в которую входит Комплект
	Weight              *float64                    `json:"weight,omitempty"`              // Вес
	ProductFolder       *NullValue[ProductFolder]   `json:"productFolder,omitempty"`       // Метаданные группы Комплекта
	Shared              *bool                       `json:"shared,omitempty"`              // Общий доступ
	Updated             *Timestamp                  `json:"updated,omitempty"`             // Момент последнего обновления сущности
	AccountID           *uuid.UUID                  `json:"accountId,omitempty"`           // ID учётной записи
	TNVED               *string                     `json:"tnved,omitempty"`               // Код ТН ВЭД
	VatEnabled          *bool                       `json:"vatEnabled,omitempty"`          // Включен ли НДС для комплекта. С помощью этого флага для комплекта можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
	Uom                 *NullValue[Uom]             `json:"uom,omitempty"`                 // Единица измерения
	UseParentVat        *bool                       `json:"useParentVat,omitempty"`        // Используется ли ставка НДС родительской группы. Если «true» для единицы ассортимента будет применена ставка, установленная для родительской группы.
	TaxSystem           TaxSystem                   `json:"taxSystem,omitempty"`           // Код системы налогообложения
	TrackingType        TrackingType                `json:"trackingType,omitempty"`        // Тип маркируемой продукции
	PaymentItemType     PaymentItem                 `json:"paymentItemType,omitempty"`     // Признак предмета расчета
	Barcodes            Slice[Barcode]              `json:"barcodes,omitempty"`            // Штрихкоды
	SalePrices          Slice[SalePrice]            `json:"salePrices,omitempty"`          // Цены продажи
	Attributes          Slice[Attribute]            `json:"attributes,omitempty"`          // Список метаданных доп. полей
}

// NewBundleFromAssortment пытается привести переданный в качестве аргумента [AssortmentPosition] к типу [Bundle].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [Bundle] или nil в случае неудачи.
func NewBundleFromAssortment(assortmentPosition *AssortmentPosition) *Bundle {
	return UnmarshalAsType[Bundle](assortmentPosition)
}

// FromAssortment пытается привести переданный в качестве аргумента [AssortmentPosition] к типу [Bundle].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [Bundle] или nil в случае неудачи.
func (bundle Bundle) FromAssortment(assortmentPosition *AssortmentPosition) *Bundle {
	return UnmarshalAsType[Bundle](assortmentPosition)
}

// AsAssortment реализует интерфейс [AssortmentConverter].
func (bundle Bundle) AsAssortment() *AssortmentPosition {
	return &AssortmentPosition{Meta: bundle.GetMeta()}
}

// GetVolume возвращает Объем.
func (bundle Bundle) GetVolume() float64 {
	return Deref(bundle.Volume)
}

// GetSyncID возвращает ID синхронизации.
func (bundle Bundle) GetSyncID() uuid.UUID {
	return Deref(bundle.SyncID)
}

// GetCode возвращает Код Комплекта.
func (bundle Bundle) GetCode() string {
	return Deref(bundle.Code)
}

// GetDescription возвращает Описание Комплекта.
func (bundle Bundle) GetDescription() string {
	return Deref(bundle.Description)
}

// GetExternalCode возвращает Внешний код Комплекта.
func (bundle Bundle) GetExternalCode() string {
	return Deref(bundle.ExternalCode)
}

// GetID возвращает ID Комплекта.
func (bundle Bundle) GetID() uuid.UUID {
	return Deref(bundle.ID)
}

// GetMeta возвращает Метаданные Комплекта.
func (bundle Bundle) GetMeta() Meta {
	return Deref(bundle.Meta)
}

// GetName возвращает Наименование Комплекта.
func (bundle Bundle) GetName() string {
	return Deref(bundle.Name)
}

// GetArchived возвращает true, если комплект находится в архиве.
func (bundle Bundle) GetArchived() bool {
	return Deref(bundle.Archived)
}

// GetArticle возвращает Артикул Комплекта.
func (bundle Bundle) GetArticle() string {
	return Deref(bundle.Article)
}

// GetImages возвращает Массив метаданных Изображений.
func (bundle Bundle) GetImages() MetaArray[Image] {
	return Deref(bundle.Images)
}

// GetComponents возвращает Массив компонентов Комплекта.
func (bundle Bundle) GetComponents() MetaArray[BundleComponent] {
	return Deref(bundle.Components)
}

// GetCountry возвращает Метаданные Страны.
func (bundle Bundle) GetCountry() Country {
	return Deref(bundle.Country).getValue()
}

// GetDiscountProhibited возвращает Признак запрета скидок.
func (bundle Bundle) GetDiscountProhibited() bool {
	return Deref(bundle.DiscountProhibited)
}

// GetEffectiveVat возвращает Реальный НДС %.
func (bundle Bundle) GetEffectiveVat() int {
	return Deref(bundle.EffectiveVat)
}

// GetEffectiveVatEnabled возвращает Дополнительный признак для определения разграничения реального НДС = 0 или "без НДС".
func (bundle Bundle) GetEffectiveVatEnabled() bool {
	return Deref(bundle.EffectiveVatEnabled)
}

// GetFiles возвращает Метаданные массива Файлов.
func (bundle Bundle) GetFiles() MetaArray[File] {
	return Deref(bundle.Files)
}

// GetGroup возвращает Отдел сотрудника.
func (bundle Bundle) GetGroup() Group {
	return Deref(bundle.Group)
}

// GetVat возвращает НДС %.
func (bundle Bundle) GetVat() int {
	return Deref(bundle.Vat)
}

// GetMinPrice возвращает Минимальную цену.
func (bundle Bundle) GetMinPrice() MinPrice {
	return Deref(bundle.MinPrice).getValue()
}

// GetOverhead возвращает Дополнительные расходы.
func (bundle Bundle) GetOverhead() BundleOverhead {
	return Deref(bundle.Overhead).getValue()
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (bundle Bundle) GetOwner() Employee {
	return Deref(bundle.Owner)
}

// GetPartialDisposal возвращает Управление состоянием частичного выбытия маркированного комплекта. «true» - возможность включена.
func (bundle Bundle) GetPartialDisposal() bool {
	return Deref(bundle.PartialDisposal)
}

// GetPathName возвращает Наименование группы, в которую входит Комплект.
func (bundle Bundle) GetPathName() string {
	return Deref(bundle.PathName)
}

// GetWeight возвращает Вес.
func (bundle Bundle) GetWeight() float64 {
	return Deref(bundle.Weight)
}

// GetSalePrices возвращает Цены продажи.
func (bundle Bundle) GetSalePrices() Slice[SalePrice] {
	return bundle.SalePrices
}

// GetProductFolder возвращает Метаданные группы Комплекта.
func (bundle Bundle) GetProductFolder() ProductFolder {
	return Deref(bundle.ProductFolder).getValue()
}

// GetShared возвращает флаг общего доступа.
func (bundle Bundle) GetShared() bool {
	return Deref(bundle.Shared)
}

// GetUpdated возвращает Момент последнего обновления сущности.
func (bundle Bundle) GetUpdated() time.Time {
	return Deref(bundle.Updated).Time()
}

// GetAccountID возвращает ID учётной записи.
func (bundle Bundle) GetAccountID() uuid.UUID {
	return Deref(bundle.AccountID)
}

// GetTNVED возвращает Код ТН ВЭД.
func (bundle Bundle) GetTNVED() string {
	return Deref(bundle.TNVED)
}

// GetVatEnabled возвращает true, если учитывается НДС.
func (bundle Bundle) GetVatEnabled() bool {
	return Deref(bundle.VatEnabled)
}

// GetUom возвращает Единицу измерения.
func (bundle Bundle) GetUom() Uom {
	return Deref(bundle.Uom).getValue()
}

// GetBarcodes возвращает Штрихкоды.
func (bundle Bundle) GetBarcodes() Slice[Barcode] {
	return bundle.Barcodes
}

// GetUseParentVat возвращает флаг использования ставки НДС родительской группы.
func (bundle Bundle) GetUseParentVat() bool {
	return Deref(bundle.UseParentVat)
}

// GetTaxSystem возвращает Код системы налогообложения.
func (bundle Bundle) GetTaxSystem() TaxSystem {
	return bundle.TaxSystem
}

// GetTrackingType возвращает Тип маркируемой продукции.
func (bundle Bundle) GetTrackingType() TrackingType {
	return bundle.TrackingType
}

// GetPaymentItemType возвращает Признак предмета расчета.
func (bundle Bundle) GetPaymentItemType() PaymentItem {
	return bundle.PaymentItemType
}

// GetAttributes возвращает Список метаданных доп. полей.
func (bundle Bundle) GetAttributes() Slice[Attribute] {
	return bundle.Attributes
}

// SetVolume устанавливает Объем.
func (bundle *Bundle) SetVolume(volume float64) *Bundle {
	bundle.Volume = &volume
	return bundle
}

// SetSyncID устанавливает ID синхронизации.
func (bundle *Bundle) SetSyncID(syncID uuid.UUID) *Bundle {
	bundle.SyncID = &syncID
	return bundle
}

// SetCode устанавливает Код Комплекта.
func (bundle *Bundle) SetCode(code string) *Bundle {
	bundle.Code = &code
	return bundle
}

// SetDescription устанавливает Описание Комплекта.
func (bundle *Bundle) SetDescription(description string) *Bundle {
	bundle.Description = &description
	return bundle
}

// SetExternalCode устанавливает Внешний код Комплекта.
func (bundle *Bundle) SetExternalCode(externalCode string) *Bundle {
	bundle.ExternalCode = &externalCode
	return bundle
}

// SetMeta устанавливает Метаданные Комплекта.
func (bundle *Bundle) SetMeta(meta *Meta) *Bundle {
	bundle.Meta = meta
	return bundle
}

// SetName устанавливает Наименование Комплекта.
func (bundle *Bundle) SetName(name string) *Bundle {
	bundle.Name = &name
	return bundle
}

// SetArchived устанавливает флаг нахождения комплекта в архиве.
func (bundle *Bundle) SetArchived(archived bool) *Bundle {
	bundle.Archived = &archived
	return bundle
}

// SetArticle устанавливает Артикул.
func (bundle *Bundle) SetArticle(article string) *Bundle {
	bundle.Article = &article
	return bundle
}

// SetImages устанавливает Массив метаданных Изображений.
//
// Принимает множество объектов [Image].
func (bundle *Bundle) SetImages(images ...*Image) *Bundle {
	bundle.Images = NewMetaArrayFrom(images)
	return bundle
}

// SetComponents устанавливает Массив компонентов Комплекта.
//
// Принимает множество объектов [BundleComponent].
func (bundle *Bundle) SetComponents(components ...*BundleComponent) *Bundle {
	bundle.Components = NewMetaArrayFrom(components)
	return bundle
}

// SetCountry устанавливает Метаданные Страны.
//
// Передача nil передаёт сброс значения (null).
func (bundle *Bundle) SetCountry(country *Country) *Bundle {
	bundle.Country = NewNullValue(country)
	return bundle
}

// SetDiscountProhibited устанавливает Признак запрета скидок.
func (bundle *Bundle) SetDiscountProhibited(discountProhibited bool) *Bundle {
	bundle.DiscountProhibited = &discountProhibited
	return bundle
}

// SetFiles устанавливает Метаданные массива Файлов.
//
// Принимает множество объектов [File].
func (bundle *Bundle) SetFiles(files ...*File) *Bundle {
	bundle.Files = NewMetaArrayFrom(files)
	return bundle
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (bundle *Bundle) SetGroup(group *Group) *Bundle {
	if group != nil {
		bundle.Group = group.Clean()
	}
	return bundle
}

// SetVat устанавливает НДС %.
func (bundle *Bundle) SetVat(vat int) *Bundle {
	bundle.Vat = &vat
	return bundle
}

// SetMinPrice устанавливает Минимальную цену.
//
// Передача nil передаёт сброс значения (null).
func (bundle *Bundle) SetMinPrice(minPrice *MinPrice) *Bundle {
	bundle.MinPrice = NewNullValue(minPrice)
	return bundle
}

// SetOverhead устанавливает Дополнительные расходы.
//
// Передача nil передаёт сброс значения (null).
func (bundle *Bundle) SetOverhead(overhead *BundleOverhead) *Bundle {
	bundle.Overhead = NewNullValue(overhead)
	return bundle
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (bundle *Bundle) SetOwner(owner *Employee) *Bundle {
	if owner != nil {
		bundle.Owner = owner.Clean()
	}
	return bundle
}

// SetPartialDisposal устанавливает Управление состоянием частичного выбытия маркированного товара. «true» - возможность включена.
func (bundle *Bundle) SetPartialDisposal(partialDisposal bool) *Bundle {
	bundle.PartialDisposal = &partialDisposal
	return bundle
}

// SetWeight устанавливает Вес.
func (bundle *Bundle) SetWeight(weight float64) *Bundle {
	bundle.Weight = &weight
	return bundle
}

// SetSalePrices устанавливает Цены продажи.
//
// Принимает множество объектов [SalePrice].
func (bundle *Bundle) SetSalePrices(salePrices ...*SalePrice) *Bundle {
	bundle.SalePrices.Push(salePrices...)
	return bundle
}

// SetProductFolder устанавливает Метаданные группы Комплекта.
//
// Передача nil передаёт сброс значения (null).
func (bundle *Bundle) SetProductFolder(productFolder *ProductFolder) *Bundle {
	bundle.ProductFolder = NewNullValue(productFolder)
	return bundle
}

// SetShared устанавливает флаг общего доступа.
func (bundle *Bundle) SetShared(shared bool) *Bundle {
	bundle.Shared = &shared
	return bundle
}

// SetTNVED устанавливает Код ТН ВЭД.
func (bundle *Bundle) SetTNVED(tnved string) *Bundle {
	bundle.TNVED = &tnved
	return bundle
}

// SetVatEnabled устанавливает значение, учитывающее НДС для комплекта.
func (bundle *Bundle) SetVatEnabled(vatEnabled bool) *Bundle {
	bundle.VatEnabled = &vatEnabled
	return bundle
}

// SetUom устанавливает Единицу измерения.
//
// Передача nil передаёт сброс значения (null).
func (bundle *Bundle) SetUom(uom *Uom) *Bundle {
	bundle.Uom = NewNullValue(uom)
	return bundle
}

// SetBarcodes устанавливает Штрихкоды.
//
// Для обновления списка штрихкодов необходимо передавать их полный список, включающий как старые,
// так и новые значения. Отсутствующие значения штрихкодов при обновлении будут удалены.
// При обновлении списка штрихкодов валидируются только новые значения.
// Ранее сохраненные штрихкоды не валидируются.
// То есть, если один из старых штрихкодов не соответствует требованиям к валидации,
// то ошибки при обновлении списка не будет. Если на вход передан пустой список штрихкодов
// или список из пустых значений, то ранее созданные штрихкоды будут удалены.
//
// Принимает множество объектов [Barcode].
func (bundle *Bundle) SetBarcodes(barcodes ...*Barcode) *Bundle {
	bundle.Barcodes = barcodes
	return bundle
}

// SetUseParentVat устанавливает флаг использования ставки НДС родительской группы.
func (bundle *Bundle) SetUseParentVat(useParentVat bool) *Bundle {
	bundle.UseParentVat = &useParentVat
	return bundle
}

// SetTaxSystem устанавливает Код системы налогообложения.
func (bundle *Bundle) SetTaxSystem(taxSystem TaxSystem) *Bundle {
	bundle.TaxSystem = taxSystem
	return bundle
}

// SetTrackingType устанавливает Тип маркируемой продукции.
func (bundle *Bundle) SetTrackingType(trackingType TrackingType) *Bundle {
	bundle.TrackingType = trackingType
	return bundle
}

// SetPaymentItemType устанавливает Признак предмета расчета.
func (bundle *Bundle) SetPaymentItemType(paymentItemType PaymentItem) *Bundle {
	bundle.PaymentItemType = paymentItemType
	return bundle
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (bundle *Bundle) SetAttributes(attributes ...*Attribute) *Bundle {
	bundle.Attributes.Push(attributes...)
	return bundle
}

// String реализует интерфейс [fmt.Stringer].
func (bundle Bundle) String() string {
	return Stringify(bundle)
}

// MetaType возвращает код сущности.
func (Bundle) MetaType() MetaType {
	return MetaTypeBundle
}

// Update shortcut
func (bundle *Bundle) Update(ctx context.Context, client *Client, params ...*Params) (*Bundle, *resty.Response, error) {
	return NewBundleService(client).Update(ctx, bundle.GetID(), bundle, params...)
}

// Create shortcut
func (bundle *Bundle) Create(ctx context.Context, client *Client, params ...*Params) (*Bundle, *resty.Response, error) {
	return NewBundleService(client).Create(ctx, bundle, params...)
}

// Delete shortcut
func (bundle *Bundle) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewBundleService(client).Delete(ctx, bundle)
}

// BundleOverhead Дополнительные расходы
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-komplekt-komplekty-atributy-wlozhennyh-suschnostej-dopolnitel-nye-rashody
type BundleOverhead struct {
	Value    *float64  `json:"value,omitempty"`    // Значение цены
	Currency *Currency `json:"currency,omitempty"` // Метаданные валюты
}

// GetValue возвращает Значение цены.
func (bundleOverhead BundleOverhead) GetValue() float64 {
	return Deref(bundleOverhead.Value)
}

// GetCurrency возвращает Метаданные валюты.
func (bundleOverhead BundleOverhead) GetCurrency() Currency {
	return Deref(bundleOverhead.Currency)
}

// SetValue устанавливает Значение цены.
func (bundleOverhead *BundleOverhead) SetValue(value *float64) *BundleOverhead {
	bundleOverhead.Value = value
	return bundleOverhead
}

// SetCurrency устанавливает Метаданные валюты.
func (bundleOverhead *BundleOverhead) SetCurrency(currency *Currency) *BundleOverhead {
	if currency != nil {
		bundleOverhead.Currency = currency.Clean()
	}
	return bundleOverhead
}

// String реализует интерфейс [fmt.Stringer].
func (bundleOverhead BundleOverhead) String() string {
	return Stringify(bundleOverhead)
}

// BundleComponent Компонент комплекта.
//
// Код сущности: bundlecomponent
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-komplekt-komplekty-komponenty-komplekta
type BundleComponent struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учётной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги, которую представляет собой компонент
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID компонента
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в компоненте
}

// NewBundleComponent принимает объект, реализующий интерфейс [AssortmentConverter] и количество.
// Возвращает новый компонент комплекта.
func NewBundleComponent(assortment AssortmentConverter, quantity float64) *BundleComponent {
	return &BundleComponent{Assortment: assortment.AsAssortment(), Quantity: &quantity}
}

// GetAccountID возвращает ID учётной записи.
func (bundleComponent BundleComponent) GetAccountID() uuid.UUID {
	return Deref(bundleComponent.AccountID)
}

// GetAssortment возвращает Метаданные товара/услуги, которую представляет собой компонент.
func (bundleComponent BundleComponent) GetAssortment() AssortmentPosition {
	return Deref(bundleComponent.Assortment)
}

// GetID возвращает ID компонента.
func (bundleComponent BundleComponent) GetID() uuid.UUID {
	return Deref(bundleComponent.ID)
}

// GetQuantity возвращает Количество товаров/услуг данного вида в компоненте.
func (bundleComponent BundleComponent) GetQuantity() float64 {
	return Deref(bundleComponent.Quantity)
}

// SetAssortment устанавливает Метаданные товара/услуги, которую представляет собой компонент.
//
// Принимает объект, реализующий интерфейс [AssortmentConverter].
func (bundleComponent *BundleComponent) SetAssortment(assortment AssortmentConverter) *BundleComponent {
	if assortment != nil {
		bundleComponent.Assortment = assortment.AsAssortment()
	}
	return bundleComponent
}

// SetQuantity устанавливает Количество товаров/услуг данного вида в компоненте.
func (bundleComponent *BundleComponent) SetQuantity(quantity float64) *BundleComponent {
	bundleComponent.Quantity = &quantity
	return bundleComponent
}

// String реализует интерфейс [fmt.Stringer].
func (bundleComponent BundleComponent) String() string {
	return Stringify(bundleComponent)
}

// MetaType возвращает код сущности.
func (BundleComponent) MetaType() MetaType {
	return MetaTypeBundleComponent
}

// BundleService методы сервиса для работы с комплектами.
type BundleService interface {
	// GetList выполняет запрос на получение списка комплектов.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[Bundle], *resty.Response, error)

	// GetListAll выполняет запрос на получение всех комплектов в виде списка.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает список объектов.
	GetListAll(ctx context.Context, params ...*Params) (*Slice[Bundle], *resty.Response, error)

	// Create выполняет запрос на создание бонусной программы.
	// Обязательные поля для заполнения:
	//	- name (Наименование комплекта)
	//	- components (Компоненты комплекта)
	// Принимает контекст, комплект и опционально объект параметров запроса Params.
	// Возвращает созданный комплект.
	Create(ctx context.Context, bundle *Bundle, params ...*Params) (*Bundle, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение комплектов.
	// Изменяемые комплекты должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список комплектов и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых комплектов.
	CreateUpdateMany(ctx context.Context, bundleList Slice[Bundle], params ...*Params) (*Slice[Bundle], *resty.Response, error)

	// GetByID выполняет запрос на получение комплекта по ID.
	// Принимает контекст, ID комплекта и опционально объект параметров запроса Params.
	// Возвращает комплект.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Bundle, *resty.Response, error)

	// Update выполняет запрос на изменение комплекта.
	// Принимает контекст, комплект и опционально объект параметров запроса Params.
	// Возвращает изменённый комплект.
	Update(ctx context.Context, id uuid.UUID, bundle *Bundle, params ...*Params) (*Bundle, *resty.Response, error)

	// DeleteByID выполняет запрос на удаление комплекта по ID.
	// Принимает контекст и ID комплекта.
	// Возвращает «true» в случае успешного удаления комплекта.
	DeleteByID(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// Delete выполняет запрос на удаление комплекта.
	// Принимает контекст и комплект.
	// Возвращает «true» в случае успешного удаления комплекта.
	Delete(ctx context.Context, entity *Bundle) (bool, *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление комплектов.
	// Принимает контекст и множество комплектов.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*Bundle) (*DeleteManyResponse, *resty.Response, error)

	// GetComponentList выполняет запрос на получение компонентов комплекта в виде списка.
	// Принимает контекст и ID комплекта.
	// Возвращает объект List.
	GetComponentList(ctx context.Context, id uuid.UUID) (*List[BundleComponent], *resty.Response, error)

	// CreateComponent выполняет запрос на добавление компонента комплекта.
	// Принимает контекст, ID комплекта и компонент комплекта.
	// Возвращает добавленный компонент комплекта.
	CreateComponent(ctx context.Context, id uuid.UUID, bundleComponent *BundleComponent) (*BundleComponent, *resty.Response, error)

	// GetComponentByID выполняет запрос на получение компонента комплекта по ID.
	// Принимает контекст, ID комплекта и ID компонента комплекта.
	GetComponentByID(ctx context.Context, id, componentID uuid.UUID) (*BundleComponent, *resty.Response, error)

	// UpdateComponent выполняет запрос на изменение компонента комплекта.
	// Принимает контекст, ID комплекта, ID компонента комплекта и компонент комплекта.
	// Возвращает изменённый компонент комплекта.
	UpdateComponent(ctx context.Context, id, componentID uuid.UUID, bundleComponent *BundleComponent) (*BundleComponent, *resty.Response, error)

	// DeleteComponent выполняет запрос на удаление компонента комплекта по ID.
	// Принимает контекст, ID комплекта и ID компонента комплекта.
	// Возвращает «true» в случае успешного удаления компонента комплекта.
	DeleteComponent(ctx context.Context, id, componentID uuid.UUID) (bool, *resty.Response, error)

	// GetFileList выполняет запрос на получение файлов в виде списка.
	// Принимает контекст и ID сущности/документа.
	// Возвращает объект List.
	GetFileList(ctx context.Context, id uuid.UUID) (*List[File], *resty.Response, error)

	// CreateFile выполняет запрос на добавление файла.
	// Принимает контекст, ID сущности/документа и файл.
	// Возвращает список файлов.
	CreateFile(ctx context.Context, id uuid.UUID, file *File) (*Slice[File], *resty.Response, error)

	// UpdateFileMany выполняет запрос на массовое создание и/или изменение файлов сущности/документа.
	// Принимает контекст, ID сущности/документа и множество файлов.
	// Возвращает созданных и/или изменённых файлов.
	UpdateFileMany(ctx context.Context, id uuid.UUID, files ...*File) (*Slice[File], *resty.Response, error)

	// DeleteFile выполняет запрос на удаление файла сущности/документа.
	// Принимает контекст, ID сущности/документа и ID файла.
	// Возвращает «true» в случае успешного удаления файла.
	DeleteFile(ctx context.Context, id uuid.UUID, fileID uuid.UUID) (bool, *resty.Response, error)

	// DeleteFileMany выполняет запрос на массовое удаление файлов сущности/документа.
	// Принимает контекст, ID сущности/документа и множество файлов.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteFileMany(ctx context.Context, id uuid.UUID, files ...*File) (*DeleteManyResponse, *resty.Response, error)

	// GetImageList выполняет запрос на получение изображений комплекта в виде списка.
	// Принимает контекст и ID комплекта.
	// Возвращает объект List.
	GetImageList(ctx context.Context, id uuid.UUID) (*List[Image], *resty.Response, error)

	// CreateImage выполняет запрос на добавление изображения.
	// Принимает контекст, ID комплекта и изображение.
	// Возвращает список изображений.
	CreateImage(ctx context.Context, id uuid.UUID, image *Image) (*Slice[Image], *resty.Response, error)

	// UpdateImageMany выполняет запрос на обновления изображений.
	// Принимает контекст, ID комплекта и изображение.
	// Если необходимо оставить некоторые Изображения, то необходимо передать эти изображения.
	// Возвращает список изображений.
	UpdateImageMany(ctx context.Context, id uuid.UUID, images ...*Image) (*Slice[Image], *resty.Response, error)

	// DeleteImage выполняет запрос на удаление изображения комплекта.
	// Принимает контекст, ID комплекта и ID изображения.
	// Возвращает «true» в случае успешного удаления изображения комплекта.
	DeleteImage(ctx context.Context, id uuid.UUID, imageID uuid.UUID) (bool, *resty.Response, error)

	// DeleteImageMany выполняет запрос на массовое удаление изображений комплекта.
	// Принимает контекст, ID комплекта и множество файлов.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteImageMany(ctx context.Context, id uuid.UUID, images ...*Image) (*DeleteManyResponse, *resty.Response, error)

	// GetAttributeList выполняет запрос на получение списка доп полей.
	// Принимает контекст.
	// Возвращает объект List.
	GetAttributeList(ctx context.Context) (*List[Attribute], *resty.Response, error)

	// GetAttributeByID выполняет запрос на получение отдельного доп поля по ID.
	// Принимает контекст и ID доп поля.
	// Возвращает найденное доп поле.
	GetAttributeByID(ctx context.Context, id uuid.UUID) (*Attribute, *resty.Response, error)

	// CreateAttribute выполняет запрос на создание доп поля.
	// Принимает контекст и доп поле.
	// Возвращает созданное доп поле.
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)

	// CreateUpdateAttributeMany выполняет запрос на массовое создание и/или изменение доп полей.
	// Изменяемые доп поля должны содержать идентификатор в виде метаданных.
	// Принимает контекст и множество доп полей.
	// Возвращает список созданных и/или изменённых доп полей.
	CreateUpdateAttributeMany(ctx context.Context, attributes ...*Attribute) (*Slice[Attribute], *resty.Response, error)

	// UpdateAttribute выполняет запрос на изменения доп поля.
	// Принимает контекст, ID доп поля и доп поле.
	// Возвращает изменённое доп поле.
	UpdateAttribute(ctx context.Context, id uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)

	// DeleteAttribute выполняет запрос на удаление доп поля.
	// Принимает контекст и ID доп поля.
	// Возвращает «true» в случае успешного удаления доп поля.
	DeleteAttribute(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// DeleteAttributeMany выполняет запрос на массовое удаление доп полей.
	// Принимает контекст и множество доп полей.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteAttributeMany(ctx context.Context, attributes ...*Attribute) (*DeleteManyResponse, *resty.Response, error)

	// GetBySyncID выполняет запрос на получение отдельного документа по syncID.
	// Принимает контекст и syncID документа.
	// Возвращает найденный документ.
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*Bundle, *resty.Response, error)

	// DeleteBySyncID выполняет запрос на удаление документа по syncID.
	// Принимает контекст и syncID документа.
	// Возвращает «true» в случае успешного удаления документа.
	DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error)
}

const (
	EndpointBundle             = EndpointEntity + string(MetaTypeBundle)
	EndpointBundleComponents   = EndpointBundle + "/%s/components"
	EndpointBundleComponentsID = EndpointBundleComponents + "/%s"
)

type bundleService struct {
	Endpoint
	endpointGetList[Bundle]
	endpointCreate[Bundle]
	endpointCreateUpdateMany[Bundle]
	endpointGetByID[Bundle]
	endpointUpdate[Bundle]
	endpointDeleteByID
	endpointDelete[Bundle]
	endpointDeleteMany[Bundle]
	endpointFiles
	endpointImages
	endpointAttributes
	endpointSyncID[Bundle]
}

func (service *bundleService) GetComponentList(ctx context.Context, id uuid.UUID) (*List[BundleComponent], *resty.Response, error) {
	path := fmt.Sprintf(EndpointBundleComponents, id)
	return NewRequestBuilder[List[BundleComponent]](service.client, path).Get(ctx)
}

func (service *bundleService) CreateComponent(ctx context.Context, id uuid.UUID, bundleComponent *BundleComponent) (*BundleComponent, *resty.Response, error) {
	path := fmt.Sprintf(EndpointBundleComponents, id)
	return NewRequestBuilder[BundleComponent](service.client, path).Post(ctx, bundleComponent)
}

func (service *bundleService) GetComponentByID(ctx context.Context, id, componentID uuid.UUID) (*BundleComponent, *resty.Response, error) {
	path := fmt.Sprintf(EndpointBundleComponentsID, id, componentID)
	return NewRequestBuilder[BundleComponent](service.client, path).Get(ctx)
}

func (service *bundleService) UpdateComponent(ctx context.Context, id, componentID uuid.UUID, bundleComponent *BundleComponent) (*BundleComponent, *resty.Response, error) {
	path := fmt.Sprintf(EndpointBundleComponentsID, id, componentID)
	return NewRequestBuilder[BundleComponent](service.client, path).Put(ctx, bundleComponent)
}

func (service *bundleService) DeleteComponent(ctx context.Context, id, componentID uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf(EndpointBundleComponentsID, id, componentID)
	return NewRequestBuilder[any](service.client, path).Delete(ctx)
}

// NewBundleService принимает [Client] и возвращает сервис для работы с комплектами.
func NewBundleService(client *Client) BundleService {
	e := NewEndpoint(client, EndpointBundle)
	return &bundleService{
		Endpoint:                 e,
		endpointGetList:          endpointGetList[Bundle]{e},
		endpointCreate:           endpointCreate[Bundle]{e},
		endpointCreateUpdateMany: endpointCreateUpdateMany[Bundle]{e},
		endpointGetByID:          endpointGetByID[Bundle]{e},
		endpointUpdate:           endpointUpdate[Bundle]{e},
		endpointDeleteByID:       endpointDeleteByID{e},
		endpointDelete:           endpointDelete[Bundle]{e},
		endpointDeleteMany:       endpointDeleteMany[Bundle]{e},
		endpointFiles:            endpointFiles{e},
		endpointImages:           endpointImages{e},
		endpointAttributes:       endpointAttributes{e},
		endpointSyncID:           endpointSyncID[Bundle]{e},
	}
}
