package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Product Товар.
//
// Код сущности: product
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-towar
type Product struct {
	Supplier            *NullValue[Counterparty]  `json:"supplier,omitempty"`            // Метаданные контрагента-поставщика
	OnTap               *bool                     `json:"onTap,omitempty"`               // Признак разливного товара. Если его значение false - поле не отображается. Данная отметка не сочетается с признаками weighed, isSerialTrackable, ppeType. Маркировка разливных товаров поддерживается только для типов MILK, PERFUMERY.
	Weighted            *bool                     `json:"weighted,omitempty"`            // Признак весового товара. Если его значение false - поле не отображается. Данная отметка не сочетается с признаками onTap, isSerialTrackable, ppeType, alcoholic. Маркировка весовых товаров поддерживается только для типа MILK.
	Code                *string                   `json:"code,omitempty"`                // Код товара
	Description         *string                   `json:"description,omitempty"`         // Описание Товара
	ExternalCode        *string                   `json:"externalCode,omitempty"`        // Внешний код Товара
	ID                  *uuid.UUID                `json:"id,omitempty"`                  // ID Товара
	Meta                *Meta                     `json:"meta,omitempty"`                // Метаданные Товара
	Name                *string                   `json:"name,omitempty"`                // Наименование Товара
	Alcoholic           *NullValue[Alcoholic]     `json:"alcoholic,omitempty"`           // Объект, содержащий поля алкогольной продукции
	Archived            *bool                     `json:"archived,omitempty"`            // Добавлен ли Товар в архив
	Article             *string                   `json:"article,omitempty"`             // Артикул
	PaymentItemType     PaymentItem               `json:"paymentItemType,omitempty"`     // Признак предмета расчета
	BuyPrice            *BuyPrice                 `json:"buyPrice,omitempty"`            // Закупочная цена
	Country             *NullValue[Country]       `json:"country,omitempty"`             // Метаданные Страны
	DiscountProhibited  *bool                     `json:"discountProhibited,omitempty"`  // Признак запрета скидок
	EffectiveVat        *int                      `json:"effectiveVat,omitempty"`        // Реальный НДС %
	EffectiveVatEnabled *bool                     `json:"effectiveVatEnabled,omitempty"` // Дополнительный признак для определения разграничения реального НДС = 0 или "без НДС". (effectiveVat = 0, effectiveVatEnabled = false) -> "без НДС", (effectiveVat = 0, effectiveVatEnabled = true) -> 0%.
	Files               *MetaArray[File]          `json:"files,omitempty"`               // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group               *Group                    `json:"group,omitempty"`               // Отдел сотрудника
	Images              *MetaArray[Image]         `json:"images,omitempty"`              // Массив метаданных Изображений (Максимальное количество изображений - 10)
	IsSerialTrackable   *bool                     `json:"isSerialTrackable,omitempty"`   // Учет по серийным номерам. Данная отметка не сочетается с признаками weighed, alcoholic, ppeType, trackingType, onTap.
	MinPrice            *NullValue[MinPrice]      `json:"minPrice,omitempty"`            // Минимальная цена
	TaxSystem           TaxSystem                 `json:"taxSystem,omitempty"`           // Код системы налогообложения
	UseParentVat        *bool                     `json:"useParentVat,omitempty"`        // Используется ли ставка НДС родительской группы. Если «true» для единицы ассортимента будет применена ставка, установленная для родительской группы
	Owner               *Employee                 `json:"owner,omitempty"`               // Метаданные владельца (Сотрудника)
	Packs               Slice[Pack]               `json:"packs,omitempty"`               // Упаковки Товара
	PartialDisposal     *bool                     `json:"partialDisposal,omitempty"`     // Управление состоянием частичного выбытия маркированного товара. «true» - возможность включена.
	PathName            *string                   `json:"pathName,omitempty"`            // Наименование группы, в которую входит Товар. Атрибут pathName сам по себе является атрибутом только для чтения, однако его можно изменить с помощью обновления атрибута productFolder
	Weight              *float64                  `json:"weight,omitempty"`              // Вес
	PpeType             *string                   `json:"ppeType,omitempty"`             // Код вида номенклатурной классификации медицинских средств индивидуальной защиты (EAN-13). Данная отметка не сочетается с признаками weighed, isSerialTrackable, alcoholic, trackingType, onTap.
	ProductFolder       *NullValue[ProductFolder] `json:"productFolder,omitempty"`       // Метаданные группы Товара
	SalePrices          Slice[SalePrice]          `json:"salePrices,omitempty"`          // Цены продажи
	Shared              *bool                     `json:"shared,omitempty"`              // Общий доступ
	MinimumBalance      *float64                  `json:"minimumBalance,omitempty"`      // Неснижаемый остаток
	SyncID              *uuid.UUID                `json:"syncId,omitempty"`              // ID синхронизации
	AccountID           *uuid.UUID                `json:"accountId,omitempty"`           // ID учётной записи
	Things              Slice[string]             `json:"things,omitempty"`              // Серийные номера
	TNVED               *string                   `json:"tnved,omitempty"`               // Код ТН ВЭД
	VatEnabled          *bool                     `json:"vatEnabled,omitempty"`          // Включен ли НДС для товара. С помощью этого флага для товара можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
	Uom                 *NullValue[Uom]           `json:"uom,omitempty"`                 // Единица измерения
	Updated             *Timestamp                `json:"updated,omitempty"`             // Момент последнего обновления товара
	Barcodes            Slice[Barcode]            `json:"barcodes,omitempty"`            // Штрихкоды
	VariantsCount       *int                      `json:"variantsCount,omitempty"`       // Количество модификаций у данного товара
	Vat                 *int                      `json:"vat,omitempty"`                 // НДС %
	TrackingType        TrackingType              `json:"trackingType,omitempty"`        // Тип маркируемой продукции
	Volume              *float64                  `json:"volume,omitempty"`              // Объем
	Attributes          Slice[Attribute]          `json:"attributes,omitempty"`          // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (product Product) Clean() *Product {
	if product.Meta == nil {
		return nil
	}
	return &Product{Meta: product.Meta}
}

// NewProductFromAssortment пытается привести переданный в качестве аргумента [AssortmentPosition] к типу [Product].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [Product] или nil в случае неудачи.
func NewProductFromAssortment(assortmentPosition *AssortmentPosition) *Product {
	return UnmarshalAsType[Product](assortmentPosition)
}

// FromAssortment пытается привести переданный в качестве аргумента [AssortmentPosition] к типу [Product].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [Product] или nil в случае неудачи.
func (product Product) FromAssortment(assortmentPosition *AssortmentPosition) *Product {
	return UnmarshalAsType[Product](assortmentPosition)
}

// AsAssortment реализует интерфейс [AssortmentConverter].
func (product Product) AsAssortment() *AssortmentPosition {
	return &AssortmentPosition{Meta: product.GetMeta()}
}

// GetSupplier возвращает Метаданные контрагента-поставщика.
func (product Product) GetSupplier() Counterparty {
	return Deref(product.Supplier).GetValue()
}

// GetOnTap возвращает Признак разливного товара.
//
// Если его значение false - поле не отображается.
//
// Данная отметка не сочетается с признаками weighed, isSerialTrackable, ppeType.
//
// Маркировка разливных товаров поддерживается только для типов MILK, PERFUMERY.
func (product Product) GetOnTap() bool {
	return Deref(product.OnTap)
}

// GetWeighted возвращает Признак весового товара.
//
// Если его значение false - поле не отображается.
//
// Данная отметка не сочетается с признаками onTap, isSerialTrackable, ppeType, alcoholic.
//
// Маркировка весовых товаров поддерживается только для типа MILK.
func (product Product) GetWeighted() bool {
	return Deref(product.Weighted)
}

// GetCode возвращает Код товара.
func (product Product) GetCode() string {
	return Deref(product.Code)
}

// GetDescription возвращает Описание Товара.
func (product Product) GetDescription() string {
	return Deref(product.Description)
}

// GetExternalCode возвращает Внешний код Товара.
func (product Product) GetExternalCode() string {
	return Deref(product.ExternalCode)
}

// GetID возвращает ID Товара.
func (product Product) GetID() uuid.UUID {
	return Deref(product.ID)
}

// GetMeta возвращает Метаданные Товара.
func (product Product) GetMeta() Meta {
	return Deref(product.Meta)
}

// GetName возвращает Наименование Товара.
func (product Product) GetName() string {
	return Deref(product.Name)
}

// GetAlcoholic возвращает Объект, содержащий поля алкогольной продукции.
func (product Product) GetAlcoholic() Alcoholic {
	return Deref(product.Alcoholic).GetValue()
}

// GetArchived возвращает true, если Товар находится в архиве.
func (product Product) GetArchived() bool {
	return Deref(product.Archived)
}

// GetArticle возвращает Артикул Товара.
func (product Product) GetArticle() string {
	return Deref(product.Article)
}

// GetPaymentItemType возвращает Признак предмета расчета.
func (product Product) GetPaymentItemType() PaymentItem {
	return product.PaymentItemType
}

// GetBuyPrice возвращает Закупочную цену.
func (product Product) GetBuyPrice() BuyPrice {
	return Deref(product.BuyPrice)
}

// GetCountry возвращает Метаданные Страны.
func (product Product) GetCountry() Country {
	return Deref(product.Country).GetValue()
}

// GetDiscountProhibited возвращает Признак запрета скидок.
func (product Product) GetDiscountProhibited() bool {
	return Deref(product.DiscountProhibited)
}

// GetEffectiveVat возвращает Реальный НДС %.
func (product Product) GetEffectiveVat() int {
	return Deref(product.EffectiveVat)
}

// GetEffectiveVatEnabled возвращает Дополнительный признак для определения разграничения реального НДС = 0 или "без НДС".
//
// (effectiveVat = 0, effectiveVatEnabled = false) -> "без НДС",
//
// (effectiveVat = 0, effectiveVatEnabled = true) -> 0%.
func (product Product) GetEffectiveVatEnabled() bool {
	return Deref(product.EffectiveVatEnabled)
}

// GetFiles возвращает Метаданные массива Файлов.
func (product Product) GetFiles() MetaArray[File] {
	return Deref(product.Files)
}

// GetGroup возвращает Отдел сотрудника.
func (product Product) GetGroup() Group {
	return Deref(product.Group)
}

// GetImages возвращает Массив метаданных Изображений.
func (product Product) GetImages() MetaArray[Image] {
	return Deref(product.Images)
}

// GetIsSerialTrackable возвращает Учет по серийным номерам.
//
// Данная отметка не сочетается с признаками weighed, alcoholic, ppeType, trackingType, onTap.
func (product Product) GetIsSerialTrackable() bool {
	return Deref(product.IsSerialTrackable)
}

// GetMinPrice возвращает Минимальную цену.
func (product Product) GetMinPrice() MinPrice {
	return Deref(product.MinPrice).GetValue()
}

// GetTaxSystem возвращает Код системы налогообложения.
func (product Product) GetTaxSystem() TaxSystem {
	return product.TaxSystem
}

// GetUseParentVat возвращает флаг использования ставки НДС родительской группы.
//
// Если «true» для единицы ассортимента будет применена ставка, установленная для родительской группы.
func (product Product) GetUseParentVat() bool {
	return Deref(product.UseParentVat)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (product Product) GetOwner() Employee {
	return Deref(product.Owner)
}

// GetPacks возвращает Упаковки Товара.
func (product Product) GetPacks() Slice[Pack] {
	return product.Packs
}

// GetPartialDisposal возвращает Управление состоянием частичного выбытия маркированного комплекта.
//
// Если «true» - возможность включена.
func (product Product) GetPartialDisposal() bool {
	return Deref(product.PartialDisposal)
}

// GetPathName возвращает Наименование группы, в которую входит Товар.
func (product Product) GetPathName() string {
	return Deref(product.PathName)
}

// GetWeight возвращает Вес.
func (product Product) GetWeight() float64 {
	return Deref(product.Weight)
}

// GetPpeType возвращает Код вида номенклатурной классификации медицинских средств индивидуальной защиты (EAN-13).
//
// Данная отметка не сочетается с признаками weighed, isSerialTrackable, alcoholic, trackingType, onTap.
func (product Product) GetPpeType() string {
	return Deref(product.PpeType)
}

// GetProductFolder возвращает Метаданные группы Товара.
func (product Product) GetProductFolder() ProductFolder {
	return Deref(product.ProductFolder).GetValue()
}

// GetSalePrices возвращает Цены продажи.
func (product Product) GetSalePrices() Slice[SalePrice] {
	return product.SalePrices
}

// GetShared возвращает флаг Общего доступа.
func (product Product) GetShared() bool {
	return Deref(product.Shared)
}

// GetMinimumBalance возвращает Неснижаемый остаток.
func (product Product) GetMinimumBalance() float64 {
	return Deref(product.MinimumBalance)
}

// GetSyncID возвращает ID синхронизации.
func (product Product) GetSyncID() uuid.UUID {
	return Deref(product.SyncID)
}

// GetAccountID возвращает ID учётной записи.
func (product Product) GetAccountID() uuid.UUID {
	return Deref(product.AccountID)
}

// GetThings возвращает Серийные номера.
func (product Product) GetThings() Slice[string] {
	return product.Things
}

// GetTNVED возвращает Код ТН ВЭД.
func (product Product) GetTNVED() string {
	return Deref(product.TNVED)
}

// GetVatEnabled возвращает true, если учитывается НДС.
//
// С помощью этого флага для товара можно выставлять НДС = 0 или НДС = "без НДС".
//
// (vat = 0, vatEnabled = false) -> vat = "без НДС",
//
// (vat = 0, vatEnabled = true) -> vat = 0%.
func (product Product) GetVatEnabled() bool {
	return Deref(product.VatEnabled)
}

// GetUom возвращает Единицу измерения.
func (product Product) GetUom() Uom {
	return Deref(product.Uom).GetValue()
}

// GetUpdated возвращает Момент последнего обновления товара.
func (product Product) GetUpdated() Timestamp {
	return Deref(product.Updated)
}

// GetBarcodes возвращает Штрихкоды.
func (product Product) GetBarcodes() Slice[Barcode] {
	return product.Barcodes
}

// GetVariantsCount возвращает Количество модификаций у данного товара.
func (product Product) GetVariantsCount() int {
	return Deref(product.VariantsCount)
}

// GetVat возвращает НДС %.
func (product Product) GetVat() int {
	return Deref(product.Vat)
}

// GetTrackingType возвращает Тип маркируемой продукции.
func (product Product) GetTrackingType() TrackingType {
	return product.TrackingType
}

// GetVolume возвращает Объем.
func (product Product) GetVolume() float64 {
	return Deref(product.Volume)
}

// GetAttributes возвращает Список метаданных доп. полей.
func (product Product) GetAttributes() Slice[Attribute] {
	return product.Attributes
}

// SetSupplier устанавливает Метаданные контрагента-поставщика.
func (product *Product) SetSupplier(supplier *Counterparty) *Product {
	product.Supplier = NewNullValue(supplier.Clean())
	return product
}

// SetOnTap устанавливает Признак разливного товара.
//
//	Если его значение false - поле не отображается.
//
// Данная отметка не сочетается с признаками weighed, isSerialTrackable, ppeType.
//
// Маркировка разливных товаров поддерживается только для типов MILK, PERFUMERY.
func (product *Product) SetOnTap(onTap bool) *Product {
	product.OnTap = &onTap
	return product
}

// SetWeighted устанавливает весового товара.
//
// Если его значение false - поле не отображается.
//
// Данная отметка не сочетается с признаками onTap, isSerialTrackable, ppeType, alcoholic.
//
// Маркировка весовых товаров поддерживается только для типа MILK.
func (product *Product) SetWeighted(weighted bool) *Product {
	product.Weighted = &weighted
	return product
}

// SetCode устанавливает Код товара.
func (product *Product) SetCode(code string) *Product {
	product.Code = &code
	return product
}

// SetDescription устанавливает Описание товара.
func (product *Product) SetDescription(description string) *Product {
	product.Description = &description
	return product
}

// SetExternalCode устанавливает Внешний код товара.
func (product *Product) SetExternalCode(externalCode string) *Product {
	product.ExternalCode = &externalCode
	return product
}

// SetMeta устанавливает Метаданные товара.
func (product *Product) SetMeta(meta *Meta) *Product {
	product.Meta = meta
	return product
}

// SetName устанавливает Наименование Товара.
func (product *Product) SetName(name string) *Product {
	product.Name = &name
	return product
}

// SetAlcoholic устанавливает Объект, содержащий поля алкогольной продукции.
//
// Передача nil передаёт сброс значения (null).
func (product *Product) SetAlcoholic(alcoholic *Alcoholic) *Product {
	product.Alcoholic = NewNullValue(alcoholic)
	return product
}

// SetArchived устанавливает флаг нахождения товара в архиве.
func (product *Product) SetArchived(archived bool) *Product {
	product.Archived = &archived
	return product
}

// SetArticle устанавливает Артикул.
func (product *Product) SetArticle(article string) *Product {
	product.Article = &article
	return product
}

// SetPaymentItemType устанавливает Признак предмета расчета.
func (product *Product) SetPaymentItemType(paymentItem PaymentItem) *Product {
	product.PaymentItemType = paymentItem
	return product
}

// SetBuyPrice устанавливает Закупочную цену.
func (product *Product) SetBuyPrice(buyPrice *BuyPrice) *Product {
	product.BuyPrice = buyPrice
	return product
}

// SetCountry устанавливает Метаданные Страны.
//
// Передача nil передаёт сброс значения (null).
func (product *Product) SetCountry(country *Country) *Product {
	product.Country = NewNullValue(country)
	return product
}

// SetDiscountProhibited устанавливает Признак запрета скидок.
func (product *Product) SetDiscountProhibited(discountProhibited bool) *Product {
	product.DiscountProhibited = &discountProhibited
	return product
}

// SetFiles устанавливает Метаданные массива Файлов.
//
// Принимает множество объектов [File].
func (product *Product) SetFiles(files ...*File) *Product {
	product.Files = NewMetaArrayFrom(files)
	return product
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (product *Product) SetGroup(group *Group) *Product {
	if group != nil {
		product.Group = group.Clean()
	}
	return product
}

// SetImages устанавливает Массив метаданных Изображений.
//
// Принимает множество объектов [Image].
func (product *Product) SetImages(images ...*Image) *Product {
	product.Images = NewMetaArrayFrom(images)
	return product
}

// SetSerialTrackable устанавливает Учет по серийным номерам.
func (product *Product) SetSerialTrackable(isSerialTrackable bool) *Product {
	product.IsSerialTrackable = &isSerialTrackable
	return product
}

// SetMinPrice устанавливает Минимальную цену.
//
// Передача nil передаёт сброс значения (null).
func (product *Product) SetMinPrice(minPrice *MinPrice) *Product {
	product.MinPrice = NewNullValue(minPrice)
	return product
}

// SetTaxSystem устанавливает Код системы налогообложения.
func (product *Product) SetTaxSystem(taxSystem TaxSystem) *Product {
	product.TaxSystem = taxSystem
	return product
}

// SetUseParentVat устанавливает флаг использования ставки НДС родительской группы.
func (product *Product) SetUseParentVat(useParentVat bool) *Product {
	product.UseParentVat = &useParentVat
	return product
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (product *Product) SetOwner(owner *Employee) *Product {
	if owner != nil {
		product.Owner = owner.Clean()
	}
	return product
}

// SetPacks устанавливает Упаковки Товара.
//
// Принимает множество объектов [Pack].
func (product *Product) SetPacks(packs ...*Pack) *Product {
	product.Packs.Push(packs...)
	return product
}

// SetPartialDisposal устанавливает Управление состоянием частичного выбытия маркированного товара.
//
// Если «true» - возможность включена.
func (product *Product) SetPartialDisposal(partialDisposal bool) *Product {
	product.PartialDisposal = &partialDisposal
	return product
}

// SetWeight устанавливает Вес.
func (product *Product) SetWeight(weight float64) *Product {
	product.Weight = &weight
	return product
}

// SetPpeType устанавливает Код вида номенклатурной классификации медицинских средств индивидуальной защиты (EAN-13).
//
// Данная отметка не сочетается с признаками weighed, isSerialTrackable, alcoholic, trackingType, onTap.
func (product *Product) SetPpeType(ppeType string) *Product {
	product.PpeType = &ppeType
	return product
}

// SetProductFolder устанавливает Метаданные группы Товара.
//
// Передача nil передаёт сброс значения (null).
func (product *Product) SetProductFolder(productFolder *ProductFolder) *Product {
	product.ProductFolder = NewNullValue(productFolder)
	return product
}

// SetSalePrices устанавливает Цены продажи.
//
// Принимает множество объектов [SalePrice].
func (product *Product) SetSalePrices(salePrices ...*SalePrice) *Product {
	product.SalePrices.Push(salePrices...)
	return product
}

// SetShared устанавливает флаг общего доступа.
func (product *Product) SetShared(shared bool) *Product {
	product.Shared = &shared
	return product
}

// SetMinimumBalance устанавливает Неснижаемый остаток.
func (product *Product) SetMinimumBalance(minimumBalance float64) *Product {
	product.MinimumBalance = &minimumBalance
	return product
}

// SetSyncID устанавливает ID синхронизации.
func (product *Product) SetSyncID(syncID uuid.UUID) *Product {
	product.SyncID = &syncID
	return product
}

// SetTNVED устанавливает Код ТН ВЭД.
func (product *Product) SetTNVED(tnved string) *Product {
	product.TNVED = &tnved
	return product
}

// SetVatEnabled устанавливает значение, учитывающее НДС для товара.
//
// С помощью этого флага для товара можно выставлять НДС = 0 или НДС = "без НДС".
//
// (vat = 0, vatEnabled = false) -> vat = "без НДС",
//
// (vat = 0, vatEnabled = true) -> vat = 0%.
func (product *Product) SetVatEnabled(vatEnabled bool) *Product {
	product.VatEnabled = &vatEnabled
	return product
}

// SetUom устанавливает Единицу измерения.
func (product *Product) SetUom(uom *Uom) *Product {
	product.Uom = NewNullValue(uom)
	return product
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
func (product *Product) SetBarcodes(barcodes ...*Barcode) *Product {
	product.Barcodes = barcodes
	return product
}

// SetVat устанавливает НДС %.
func (product *Product) SetVat(vat int) *Product {
	product.Vat = &vat
	return product
}

// SetTrackingType устанавливает Тип маркируемой продукции.
func (product *Product) SetTrackingType(trackingType TrackingType) *Product {
	product.TrackingType = trackingType
	return product
}

// SetVolume устанавливает Объем.
func (product *Product) SetVolume(volume float64) *Product {
	product.Volume = &volume
	return product
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (product *Product) SetAttributes(attributes ...*Attribute) *Product {
	product.Attributes.Push(attributes...)
	return product
}

// String реализует интерфейс [fmt.Stringer].
func (product Product) String() string {
	return Stringify(product)
}

// MetaType возвращает код сущности.
func (Product) MetaType() MetaType {
	return MetaTypeProduct
}

// Create shortcut
func (product Product) Create(ctx context.Context, client *Client, params ...*Params) (*Product, *resty.Response, error) {
	return NewProductService(client).Create(ctx, &product, params...)
}

// Update shortcut
func (product Product) Update(ctx context.Context, client *Client, params ...*Params) (*Product, *resty.Response, error) {
	return NewProductService(client).Update(ctx, product.GetID(), &product, params...)
}

// Delete shortcut
func (product Product) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewProductService(client).Delete(ctx, product.GetID())
}

// Alcoholic Объект, содержащий поля алкогольной продукции.
//
// Данный объект не сочетается с признаками weighed, isSerialTrackable, ppeType, trackingType, если он не промаркирован как BEER_ALCOHOL или NOT_TRACKED.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-towar-towary-atributy-wlozhennyh-suschnostej-ob-ekt-soderzhaschij-polq-alkogol-noj-produkcii
type Alcoholic struct {
	Excise   *bool    `json:"excise,omitempty"`   // Содержит акцизную марку
	Type     *int     `json:"type,omitempty"`     // Код вида продукции
	Strength *float64 `json:"strength,omitempty"` // Крепость
	Volume   *float64 `json:"volume,omitempty"`   // Объём тары
}

// GetExcise возвращает признак содержания акцизной марки.
func (alcoholic Alcoholic) GetExcise() bool {
	return Deref(alcoholic.Excise)
}

// GetType возвращает Код вида продукции.
func (alcoholic Alcoholic) GetType() int {
	return Deref(alcoholic.Type)
}

// GetStrength возвращает Крепость.
func (alcoholic Alcoholic) GetStrength() float64 {
	return Deref(alcoholic.Strength)
}

// GetVolume возвращает Объём тары.
func (alcoholic Alcoholic) GetVolume() float64 {
	return Deref(alcoholic.Volume)
}

// SetExcise устанавливает признак содержания акцизной марки.
func (alcoholic *Alcoholic) SetExcise(excise bool) *Alcoholic {
	alcoholic.Excise = &excise
	return alcoholic
}

// SetType устанавливает Код вида продукции.
func (alcoholic *Alcoholic) SetType(value int) *Alcoholic {
	alcoholic.Type = &value
	return alcoholic
}

// SetStrength устанавливает Крепость.
func (alcoholic *Alcoholic) SetStrength(strength float64) *Alcoholic {
	alcoholic.Strength = &strength
	return alcoholic
}

// SetVolume устанавливает Объём тары.
func (alcoholic *Alcoholic) SetVolume(volume float64) *Alcoholic {
	alcoholic.Volume = &volume
	return alcoholic
}

// String реализует интерфейс [fmt.Stringer].
func (alcoholic Alcoholic) String() string {
	return Stringify(alcoholic)
}

// ProductService описывает методы сервиса для работы с товарами.
type ProductService interface {
	// GetList выполняет запрос на получение списка товаров.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[Product], *resty.Response, error)

	// Create выполняет запрос на создание товара.
	// Обязательные поля для заполнения:
	//	- name (Наименование товара)
	// Принимает контекст, товар и опционально объект параметров запроса Params.
	// Возвращает созданный товар.
	Create(ctx context.Context, product *Product, params ...*Params) (*Product, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение товаров.
	// Изменяемые товары должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список товаров и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых товаров.
	CreateUpdateMany(ctx context.Context, productList Slice[Product], params ...*Params) (*Slice[Product], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление товаров.
	// Принимает контекст и множество товаров.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*Product) (*DeleteManyResponse, *resty.Response, error)

	// Delete выполняет запрос на удаление товара.
	// Принимает контекст и ID товара.
	// Возвращает «true» в случае успешного удаления товара.
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельного товара по ID.
	// Принимает контекст, ID товара и опционально объект параметров запроса Params.
	// Возвращает найденный товар.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Product, *resty.Response, error)

	// Update выполняет запрос на изменение товара.
	// Принимает контекст, товар и опционально объект параметров запроса Params.
	// Возвращает изменённый товар.
	Update(ctx context.Context, id uuid.UUID, product *Product, params ...*Params) (*Product, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных товаров.
	// Принимает контекст.
	// Возвращает объект метаданных MetaAttributesSharedWrapper.
	GetMetadata(ctx context.Context) (*MetaAttributesSharedWrapper, *resty.Response, error)

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

	// GetImageList выполняет запрос на получение изображений товара в виде списка.
	// Принимает контекст и ID товара.
	// Возвращает объект List.
	GetImageList(ctx context.Context, id uuid.UUID) (*List[Image], *resty.Response, error)

	// CreateImage выполняет запрос на добавление изображения.
	// Принимает контекст, ID товара и изображение.
	// Возвращает список изображений.
	CreateImage(ctx context.Context, id uuid.UUID, image *Image) (*Slice[Image], *resty.Response, error)

	// UpdateImageMany выполняет запрос на обновления изображений.
	// Принимает контекст, ID товара и изображение.
	// Если необходимо оставить некоторые Изображения, то необходимо передать эти изображения.
	// Возвращает список изображений.
	UpdateImageMany(ctx context.Context, id uuid.UUID, images ...*Image) (*Slice[Image], *resty.Response, error)

	// DeleteImage выполняет запрос на удаление изображения товара.
	// Принимает контекст, ID товара и ID изображения.
	// Возвращает «true» в случае успешного удаления изображения товара.
	DeleteImage(ctx context.Context, id uuid.UUID, imageID uuid.UUID) (bool, *resty.Response, error)

	// DeleteImageMany выполняет запрос на массовое удаление изображений товара.
	// Принимает контекст, ID товара и множество файлов.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteImageMany(ctx context.Context, id uuid.UUID, images ...*Image) (*DeleteManyResponse, *resty.Response, error)

	// GetBySyncID выполняет запрос на получение отдельного документа по syncID.
	// Принимает контекст и syncID документа.
	// Возвращает найденный документ.
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*Product, *resty.Response, error)

	// DeleteBySyncID выполняет запрос на удаление документа по syncID.
	// Принимает контекст и syncID документа.
	// Возвращает «true» в случае успешного удаления документа.
	DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error)

	// GetNamedFilterList выполняет запрос на получение списка фильтров.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetNamedFilterList(ctx context.Context, params ...*Params) (*List[NamedFilter], *resty.Response, error)

	// GetNamedFilterByID выполняет запрос на получение отдельного фильтра по ID.
	// Принимает контекст и ID фильтра.
	// Возвращает найденный фильтр.
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)

	// GetAudit выполняет запрос на получения событий товара.
	// Принимает контекст, ID товара и опционально объект параметров запроса Params.
	// Возвращает объект List
	GetAudit(ctx context.Context, id uuid.UUID, params ...*Params) (*List[AuditEvent], *resty.Response, error)

	// PrintLabel выполняет запрос на печать этикеток и ценников.
	// Принимает контекст, ID товара и объект PrintLabelArg.
	// Возвращает объект PrintFile.
	PrintLabel(ctx context.Context, id uuid.UUID, PrintLabelArg *PrintLabelArg) (*PrintFile, *resty.Response, error)

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
}

// NewProductService принимает [Client] и возвращает сервис для работы с товарами.
func NewProductService(client *Client) ProductService {
	return newMainService[Product, any, MetaAttributesSharedWrapper, any](NewEndpoint(client, "entity/product"))
}
