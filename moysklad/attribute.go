package moysklad

import (
	"github.com/google/uuid"
)

// Attribute Доп поле.
// Ключевое слово: attributemetadata
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/workbook/#workbook-rabota-s-dopolnitel-nymi-polqmi-cherez-json-api
type Attribute struct {
	Value            any           `json:"value,omitempty"`
	CustomEntityMeta *Meta         `json:"customEntityMeta,omitempty"`
	Description      *string       `json:"description,omitempty"`
	ID               *uuid.UUID    `json:"id,omitempty"`
	Meta             *Meta         `json:"meta,omitempty"`
	Name             *string       `json:"name,omitempty"`
	Required         *bool         `json:"required,omitempty"`
	Show             *bool         `json:"show,omitempty"`
	File             *File         `json:"file,omitempty"`
	Type             AttributeType `json:"type,omitempty"`
}

func (a Attribute) String() string {
	return Stringify(a)
}

func (a Attribute) MetaType() MetaType {
	return MetaTypeAttribute
}

type Attributes []*Attribute

// Push добавляет элементы в срез.
func (a *Attributes) Push(elements ...*Attribute) *Attributes {
	*a = append(*a, elements...)
	return a
}

// AttributeType Тип доп. поля
type AttributeType string

const (
	AttributeTypeTime                   AttributeType = "time"
	AttributeTypeLink                   AttributeType = "link"
	AttributeTypeString                 AttributeType = "string"
	AttributeTypeText                   AttributeType = "text"
	AttributeTypeFile                   AttributeType = "file"
	AttributeTypeBoolean                AttributeType = "boolean"
	AttributeTypeDouble                 AttributeType = "double"
	AttributeTypeLong                   AttributeType = "long"
	AttributeTypeDictionaryContract     AttributeType = "contract"
	AttributeTypeDictionaryCounterParty AttributeType = "counterparty"
	AttributeTypeDictionaryProject      AttributeType = "project"
	AttributeTypeDictionaryStore        AttributeType = "store"
	AttributeTypeDictionaryEmployee     AttributeType = "employee"
	AttributeTypeDictionaryProduct      AttributeType = "product"
	AttributeTypeDictionaryCustom       AttributeType = "customentity"
)
