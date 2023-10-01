package moysklad

import (
	"github.com/google/uuid"
)

// Bundle Комплект.
// Ключевое слово: bundle
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-komplekt
type Bundle struct {
	AccountId           *uuid.UUID                  `json:"accountId,omitempty"`           // ID учетной записи
	Barcodes            *Barcodes                   `json:"barcodes,omitempty"`            // Штрихкоды
	Code                *string                     `json:"code,omitempty"`                // Код
	Description         *string                     `json:"description,omitempty"`         // Описание
	ExternalCode        *string                     `json:"externalCode,omitempty"`        // Внешний код
	Id                  *uuid.UUID                  `json:"id,omitempty"`                  // ID сущности
	Meta                *Meta                       `json:"meta,omitempty"`                // Метаданные
	Name                *string                     `json:"name,omitempty"`                // Наименование
	Archived            *bool                       `json:"archived,omitempty"`            // Добавлен ли Комплект в архив
	Article             *string                     `json:"article,omitempty"`             // Артикул
	Attributes          *Attributes                 `json:"attributes,omitempty"`          // Коллекция доп. полей
	Components          *Positions[BundleComponent] `json:"components,omitempty"`          // Массив компонентов Комплекта
	Country             *Country                    `json:"country,omitempty"`             // Метаданные Страны
	DiscountProhibited  *bool                       `json:"discountProhibited"`            // Признак запрета скидок
	EffectiveVat        *int                        `json:"effectiveVat,omitempty"`        // Реальный НДС %
	EffectiveVatEnabled *bool                       `json:"effectiveVatEnabled,omitempty"` // Дополнительный признак для определения разграничения реального НДС = 0 или "без НДС". (effectiveVat = 0, effectiveVatEnabled = false) -> "без НДС", (effectiveVat = 0, effectiveVatEnabled = true) -> 0%.
	Files               *Files                      `json:"files,omitempty"`               // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group               *Group                      `json:"group,omitempty"`               // Метаданные отдела сотрудника
	Images              *Images                     `json:"images,omitempty"`              // Массив метаданных Изображений (Максимальное количество изображений - 10)
	MinPrice            *MinPrice                   `json:"minPrice,omitempty"`            // Минимальная цена
	Overhead            *BundleOverhead             `json:"overhead,omitempty"`            // Дополнительные расходы
	Owner               *Employee                   `json:"owner,omitempty"`               // Метаданные владельца (Сотрудника)
	PartialDisposal     *bool                       `json:"partialDisposal,omitempty"`     // Управление состоянием частичного выбытия маркированного товара. «true» - возможность включена.
	PathName            *string                     `json:"pathName,omitempty"`            // Наименование группы, в которую входит Комплект
	PaymentItemType     PaymentItem                 `json:"paymentItemType,omitempty"`     // Признак предмета расчета
	SalePrices          *SalePrices                 `json:"salePrices,omitempty"`          // Цены продажи
	ProductFolder       *ProductFolder              `json:"productFolder,omitempty"`       // Метаданные группы
	Shared              *bool                       `json:"shared,omitempty"`              // Общий доступ
	SyncId              *uuid.UUID                  `json:"syncId,omitempty"`              // Общий доступ
	TaxSystem           GoodTaxSystem               `json:"taxSystem,omitempty"`           // Код системы налогообложения
	Tnved               *string                     `json:"tnved,omitempty"`               // Код ТН ВЭД
	TrackingType        TrackingType                `json:"trackingType,omitempty"`        // Тип маркируемой продукции
	Uom                 *Uom                        `json:"uom,omitempty"`                 // Единицы измерения
	Updated             *Timestamp                  `json:"updated,omitempty"`             // Момент последнего обновления сущности
	UseParentVat        *bool                       `json:"useParentVat,omitempty"`        // Используется ли ставка НДС родительской группы. Если true для единицы ассортимента будет применена ставка, установленная для родительской группы.
	Vat                 *int                        `json:"vat,omitempty"`                 // НДС %
	VatEnabled          *bool                       `json:"vatEnabled,omitempty"`          // Включен ли НДС для товара. С помощью этого флага для товара можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
	Volume              *float64                    `json:"volume,omitempty"`              // Объем
	Weight              *float64                    `json:"weight,omitempty"`              // Вес
}

func (b Bundle) String() string {
	return Stringify(b)
}

func (b Bundle) MetaType() MetaType {
	return MetaTypeBundle
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (b Bundle) GetMeta() *Meta {
	return b.Meta
}

func (b Bundle) ConvertToAssortmentPosition() (*AssortmentPosition, error) {
	return convertToAssortmentPosition(b)
}

// BundleOverhead Дополнительные расходы
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-komplekt-komplekty-atributy-wlozhennyh-suschnostej-dopolnitel-nye-rashody
type BundleOverhead struct {
	Value    *float64  `json:"value,omitempty"`    // Значение цены
	Currency *Currency `json:"currency,omitempty"` // Ссылка на валюту в формате Метаданных
}

func (b BundleOverhead) String() string {
	return Stringify(b)
}

// BundleComponent Компонент комплекта.
// Ключевое слово: bundlecomponent
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-komplekt-komplekty-komponenty-komplekta
type BundleComponent struct {
	AccountId  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги, которую представляет собой компонент
	Id         *uuid.UUID          `json:"id,omitempty"`         // ID сущности
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в компоненте
}

func (b BundleComponent) String() string {
	return Stringify(b)
}

func (b BundleComponent) MetaType() MetaType {
	return MetaTypeBundleComponent
}
