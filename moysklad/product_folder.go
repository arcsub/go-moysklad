package moysklad

import (
	"github.com/google/uuid"
)

// ProductFolder Группа товаров.
// Ключевое слово: productfolder
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-gruppa-towarow
type ProductFolder struct {
	AccountID           *uuid.UUID     `json:"accountId,omitempty"`           // ID учетной записи
	Archived            *bool          `json:"archived,omitempty"`            // Добавлена ли Группа товаров в архив
	Code                *string        `json:"code,omitempty"`                // Код Группы товаров
	Description         *string        `json:"description,omitempty"`         // Описание Группы товаров
	EffectiveVat        *int           `json:"effectiveVat,omitempty"`        // Реальный НДС %
	EffectiveVatEnabled *bool          `json:"effectiveVatEnabled,omitempty"` // Дополнительный признак для определения разграничения реального НДС = 0 или "без НДС". (effectiveVat = 0, effectiveVatEnabled = false) -> "без НДС", (effectiveVat = 0, effectiveVatEnabled = true) -> 0%.
	ExternalCode        *string        `json:"externalCode,omitempty"`        // Внешний код Группы товаров
	Group               *Group         `json:"group,omitempty"`               // Метаданные отдела сотрудника
	ID                  *uuid.UUID     `json:"id,omitempty"`                  // ID сущности
	Meta                *Meta          `json:"meta,omitempty"`                // Метаданные
	Name                *string        `json:"name,omitempty"`                // Название
	Owner               *Employee      `json:"owner,omitempty"`               // Метаданные владельца (Сотрудника)
	PathName            *string        `json:"pathName,omitempty"`            // Наименование Группы товаров, в которую входит данная Группа товаров
	ProductFolder       *ProductFolder `json:"productFolder,omitempty"`       // Ссылка на Группу товаров, в которую входит данная Группа товаров, в формате Метаданных
	Shared              *bool          `json:"shared,omitempty"`              // Общий доступ
	TaxSystem           GoodTaxSystem  `json:"taxSystem,omitempty"`           // Код системы налогообложения
	Updated             *Timestamp     `json:"updated,omitempty"`             // Момент последнего обновления сущности
	UseParentVat        *bool          `json:"useParentVat,omitempty"`        // Используется ли ставка НДС родительской группы. Если true для единицы ассортимента будет применена ставка, установленная для родительской группы.
	Vat                 *int           `json:"vat,omitempty"`                 // НДС %
	VatEnabled          *bool          `json:"vatEnabled,omitempty"`          // Включен ли НДС для группы. С помощью этого флага для группы можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
}

func (p ProductFolder) String() string {
	return Stringify(p)
}

func (p ProductFolder) MetaType() MetaType {
	return MetaTypeProductFolder
}

type ProductFolders MetaArray[ProductFolder]
