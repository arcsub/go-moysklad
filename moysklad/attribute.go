package moysklad

import (
	"github.com/google/uuid"
)

// Attribute Доп поле
// Ключевое слово: attributemetadata
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/workbook/#workbook-rabota-s-dopolnitel-nymi-polqmi-cherez-json-api
type Attribute struct {
	CustomEntityMeta *Meta         `json:"customEntityMeta,omitempty"` // Ссылка на метаданные пользовательского справочника
	Description      *string       `json:"description,omitempty"`      // Описание доп. поля
	ID               *uuid.UUID    `json:"id,omitempty"`               // ID доп. поля
	Meta             *Meta         `json:"meta,omitempty"`             // Ссылка на метаданные доп. поля
	Name             *string       `json:"name,omitempty"`             // Наименование доп. поля
	Required         *bool         `json:"required,omitempty"`         // Является ли доп. поле обязательным
	Show             *bool         `json:"show,omitempty"`             // Показывать ли доп. поле на UI. Не может быть скрытым и обязательным одновременно. Только для операций
	Type             AttributeType `json:"type,omitempty"`             // Тип доп. поля
	Value            any           `json:"value,omitempty"`            // Значение доп. поля
	File             *File         `json:"file,omitempty"`
}

func (a Attribute) String() string {
	return Stringify(a)
}

func (a Attribute) MetaType() MetaType {
	return MetaTypeAttribute
}

type Attributes = Iterator[Attribute]

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
