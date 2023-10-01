package moysklad

import (
	"github.com/google/uuid"
)

// Service Услуга.
// Ключевое слово: service
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-usluga
type Service struct {
	AccountId           *uuid.UUID     `json:"accountId,omitempty"`           // ID учетной записи
	Barcodes            *Barcodes      `json:"barcodes,omitempty"`            // Штрихкоды
	Code                *string        `json:"code,omitempty"`                // Код
	Description         *string        `json:"description,omitempty"`         // Описание
	ExternalCode        *string        `json:"externalCode,omitempty"`        // Внешний код
	Id                  *uuid.UUID     `json:"id,omitempty"`                  // ID сущности
	Meta                *Meta          `json:"meta,omitempty"`                // Метаданные
	Name                *string        `json:"name,omitempty"`                // Наименование
	Archived            *bool          `json:"archived,omitempty"`            // Добавлена ли Услуга в архив
	Attributes          *Attributes    `json:"attributes,omitempty"`          // Коллекция доп. полей
	BuyPrice            *BuyPrice      `json:"buyPrice,omitempty"`            // Закупочная цена
	DiscountProhibited  *bool          `json:"discountProhibited,omitempty"`  // Признак запрета скидок
	EffectiveVat        *int           `json:"effectiveVat,omitempty"`        // Реальный НДС %
	EffectiveVatEnabled *bool          `json:"effectiveVatEnabled,omitempty"` // Дополнительный признак для определения разграничения реального НДС = 0 или "без НДС". (effectiveVat = 0, effectiveVatEnabled = false) -> "без НДС", (effectiveVat = 0, effectiveVatEnabled = true) -> 0%.
	Files               *Files         `json:"files,omitempty"`               // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group               *Group         `json:"group,omitempty"`               // Метаданные отдела сотрудника
	MinPrice            *MinPrice      `json:"minPrice,omitempty"`            // Минимальная цена
	Owner               *Employee      `json:"owner,omitempty"`               // Метаданные владельца (Сотрудника)
	PathName            *string        `json:"pathName,omitempty"`            // Наименование группы, в которую входит Услуга
	PaymentItemType     PaymentItem    `json:"paymentItemType,omitempty"`     // Признак предмета расчета
	ProductFolder       *ProductFolder `json:"productFolder,omitempty"`       // Метаданные группы
	SalePrices          *SalePrices    `json:"salePrices,omitempty"`          // Цены продажи
	Shared              *bool          `json:"shared,omitempty"`              // Общий доступ
	SyncId              *uuid.UUID     `json:"syncId,omitempty"`              // ID синхронизации
	TaxSystem           TaxSystem      `json:"taxSystem,omitempty"`           // Код системы налогообложения
	Uom                 *Uom           `json:"uom,omitempty"`                 // Единицы измерения
	Updated             *Timestamp     `json:"updated,omitempty"`             // Момент последнего обновления сущности
	UseParentVat        *bool          `json:"useParentVat,omitempty"`        // Используется ли ставка НДС родительской группы. Если true для единицы ассортимента будет применена ставка, установленная для родительской группы.
	Vat                 *int           `json:"vat,omitempty"`                 // НДС %
	VatEnabled          *bool          `json:"vatEnabled,omitempty"`          // Включен ли НДС для услуги. С помощью этого флага для услуги можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
}

func (s Service) String() string {
	return Stringify(s)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (s Service) GetMeta() *Meta {
	return s.Meta
}

func (s Service) MetaType() MetaType {
	return MetaTypeService
}

func (s Service) ConvertToAssortmentPosition() (*AssortmentPosition, error) {
	return convertToAssortmentPosition(s)
}
