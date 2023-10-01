package moysklad

import (
	"github.com/google/uuid"
)

// Product Товар.
// Ключевое слово: product
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-towar
type Product struct {
	AccountId           *uuid.UUID     `json:"accountId,omitempty"`           // ID учетной записи
	Barcodes            *Barcodes      `json:"barcodes,omitempty"`            // Штрихкоды
	Code                *string        `json:"code,omitempty"`                // Код
	Description         *string        `json:"description,omitempty"`         // Описание
	ExternalCode        *string        `json:"externalCode,omitempty"`        // Внешний код
	Id                  *uuid.UUID     `json:"id,omitempty"`                  // ID сущности
	Meta                *Meta          `json:"meta,omitempty"`                // Метаданные
	Name                *string        `json:"name,omitempty"`                // Наименование
	Alcoholic           *Alcoholic     `json:"alcoholic,omitempty"`           // Объект, содержащий поля алкогольной продукции
	Archived            *bool          `json:"archived,omitempty"`            // Добавлен ли Товар в архив
	Article             *string        `json:"article,omitempty"`             // Артикул
	Attributes          *Attributes    `json:"attributes,omitempty"`          // Коллекция доп. полей
	BuyPrice            *BuyPrice      `json:"buyPrice,omitempty"`            // Закупочная цена
	Country             *Country       `json:"country,omitempty"`             // Метаданные Страны
	DiscountProhibited  *bool          `json:"discountProhibited,omitempty"`  // Признак запрета скидок
	EffectiveVat        *int           `json:"effectiveVat,omitempty"`        // Реальный НДС %
	EffectiveVatEnabled *bool          `json:"effectiveVatEnabled,omitempty"` // Дополнительный признак для определения разграничения реального НДС = 0 или "без НДС". (effectiveVat = 0, effectiveVatEnabled = false) -> "без НДС", (effectiveVat = 0, effectiveVatEnabled = true) -> 0%.
	Files               *Files         `json:"files,omitempty"`               // Метаданные массива Файлов
	Group               *Group         `json:"group,omitempty"`               // Метаданные отдела сотрудника
	Images              *Images        `json:"images,omitempty"`              // Массив метаданных Изображений (Максимальное количество изображений - 10)
	IsSerialTrackable   *bool          `json:"isSerialTrackable,omitempty"`   // Учет по серийным номерам. Данная отметка не сочетается с признаками weighed, alcoholic, ppeType, trackingType, onTap.
	MinPrice            *MinPrice      `json:"minPrice,omitempty"`            // Минимальная цена
	MinimumBalance      *float64       `json:"minimumBalance,omitempty"`      // Неснижаемый остаток
	OnTap               *bool          `json:"onTap,omitempty"`               // Флаг разливного товара
	Owner               *Employee      `json:"owner,omitempty"`               // Метаданные владельца (Сотрудника)
	Packs               *Packs         `json:"packs,omitempty"`               // Упаковки Товара
	PartialDisposal     *bool          `json:"partialDisposal,omitempty"`     // Управление состоянием частичного выбытия маркированного товара. «true» - возможность включена.
	PathName            *string        `json:"pathName,omitempty"`            // Наименование группы, в которую входит Товар
	PaymentItemType     PaymentItem    `json:"paymentItemType,omitempty"`     // Признак предмета расчета
	PpeType             *string        `json:"ppeType,omitempty"`             // Код вида номенклатурной классификации медицинских средств индивидуальной защиты (EAN-13)
	ProductFolder       *ProductFolder `json:"productFolder,omitempty"`       // Метаданные группы
	SalePrices          *SalePrices    `json:"salePrices,omitempty"`          // Цены продажи
	Shared              *bool          `json:"shared,omitempty"`              // Общий доступ
	Supplier            *Counterparty  `json:"supplier,omitempty"`            // Метаданные контрагента-поставщика
	SyncId              *uuid.UUID     `json:"syncId,omitempty"`              // ID синхронизации
	TaxSystem           GoodTaxSystem  `json:"taxSystem,omitempty"`           // Код системы налогообложения
	Things              *Things        `json:"things,omitempty"`              // Серийные номера
	Tnved               *string        `json:"tnved,omitempty"`               // Код ТН ВЭД
	TrackingType        TrackingType   `json:"trackingType,omitempty"`        // Тип маркируемой продукции
	Uom                 *Uom           `json:"uom,omitempty"`                 // Единицы измерения
	Updated             *Timestamp     `json:"updated,omitempty"`             // Момент последнего обновления сущности
	UseParentVat        *bool          `json:"useParentVat,omitempty"`        // Используется ли ставка НДС родительской группы. Если true для единицы ассортимента будет применена ставка, установленная для родительской группы.
	VariantsCount       *int           `json:"variantsCount,omitempty"`       // Количество модификаций у данного товара
	Vat                 *int           `json:"vat,omitempty"`                 // НДС %
	VatEnabled          *bool          `json:"vatEnabled,omitempty"`          // Включен ли НДС для товара. С помощью этого флага для товара можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
	Volume              *float64       `json:"volume,omitempty"`              // Объём
	Weight              *float64       `json:"weight,omitempty"`              // Вес
}

func (p Product) String() string {
	return Stringify(p)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (p Product) GetMeta() *Meta {
	return p.Meta
}

func (p Product) MetaType() MetaType {
	return MetaTypeProduct
}

func (p Product) ConvertToAssortmentPosition() (*AssortmentPosition, error) {
	return convertToAssortmentPosition(p)
}

// Alcoholic Объект, содержащий поля алкогольной продукции
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-towar-towary-atributy-wlozhennyh-suschnostej-ob-ekt-soderzhaschij-polq-alkogol-noj-produkcii
type Alcoholic struct {
	Excise   *bool    `json:"excise,omitempty"`   // Содержит акцизную марку
	Type     *int     `json:"type,omitempty"`     // Код вида продукции
	Strength *float64 `json:"strength,omitempty"` // Крепость
	Volume   *float64 `json:"volume,omitempty"`   // Объём тары
}

func (a Alcoholic) String() string {
	return Stringify(a)
}
