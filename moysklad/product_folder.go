package moysklad

import (
	"github.com/google/uuid"
)

// ProductFolder Группа товаров.
// Ключевое слово: productfolder
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-gruppa-towarow
type ProductFolder struct {
	Name                *string        `json:"name,omitempty"`
	UseParentVat        *bool          `json:"useParentVat,omitempty"`
	Code                *string        `json:"code,omitempty"`
	Description         *string        `json:"description,omitempty"`
	EffectiveVat        *int           `json:"effectiveVat,omitempty"`
	EffectiveVatEnabled *bool          `json:"effectiveVatEnabled,omitempty"`
	ExternalCode        *string        `json:"externalCode,omitempty"`
	AccountID           *uuid.UUID     `json:"accountId,omitempty"`
	VatEnabled          *bool          `json:"vatEnabled,omitempty"`
	Archived            *bool          `json:"archived,omitempty"`
	Group               *Group         `json:"group,omitempty"`
	Owner               *Employee      `json:"owner,omitempty"`
	PathName            *string        `json:"pathName,omitempty"`
	ProductFolder       *ProductFolder `json:"productFolder,omitempty"`
	Shared              *bool          `json:"shared,omitempty"`
	ID                  *uuid.UUID     `json:"id,omitempty"`
	Updated             *Timestamp     `json:"updated,omitempty"`
	Meta                *Meta          `json:"meta,omitempty"`
	Vat                 *int           `json:"vat,omitempty"`
	TaxSystem           GoodTaxSystem  `json:"taxSystem,omitempty"`
}

func (p ProductFolder) String() string {
	return Stringify(p)
}

func (p ProductFolder) MetaType() MetaType {
	return MetaTypeProductFolder
}

type ProductFolders MetaArray[ProductFolder]
