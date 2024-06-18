package moysklad

import (
	"encoding/base64"
	"github.com/google/uuid"
	"os"
	"path/filepath"
)

// Attribute описание Доп поле.
// Ключевое слово: attributemetadata
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/workbook/#workbook-rabota-s-dopolnitel-nymi-polqmi-cherez-json-api
type Attribute struct {
	CustomEntityMeta *Meta                     `json:"customEntityMeta,omitempty"`
	Description      *string                   `json:"description,omitempty"`
	ID               *uuid.UUID                `json:"id,omitempty"`
	Meta             *Meta                     `json:"meta,omitempty"`
	Name             *string                   `json:"name,omitempty"`
	Required         *bool                     `json:"required,omitempty"`
	Show             *bool                     `json:"show,omitempty"`
	Value            *NullValueAny             `json:"value,omitempty"`
	File             *NullValue[AttributeFile] `json:"file,omitempty"`
	Download         *Meta                     `json:"download,omitempty"`
	Type             AttributeType             `json:"type,omitempty"`
}

func (attribute Attribute) GetCustomEntityMeta() Meta {
	return Deref(attribute.CustomEntityMeta)
}

func (attribute Attribute) GetDescription() string {
	return Deref(attribute.Description)
}

func (attribute Attribute) GetID() uuid.UUID {
	return Deref(attribute.ID)
}

func (attribute Attribute) GetMeta() Meta {
	return Deref(attribute.Meta)
}

func (attribute Attribute) GetName() string {
	return Deref(attribute.Name)
}

func (attribute Attribute) GetRequired() bool {
	return Deref(attribute.Required)
}

func (attribute Attribute) GetShow() bool {
	return Deref(attribute.Show)
}

func (attribute Attribute) GetType() AttributeType {
	return attribute.Type
}

func (attribute Attribute) GetValue() any {
	return attribute.Value.Get()
}

func (attribute Attribute) GetDownload() Meta {
	return Deref(attribute.Meta)
}

func (attribute *Attribute) SetMeta(meta *Meta) *Attribute {
	attribute.Meta = meta
	return attribute
}

func (attribute *Attribute) SetName(name string) *Attribute {
	attribute.Name = &name
	return attribute
}

func (attribute *Attribute) SetType(attributeType AttributeType) *Attribute {
	attribute.Type = attributeType
	return attribute
}

func (attribute *Attribute) SetValue(value any) *Attribute {
	attribute.Value = NewNullValueAnyFrom(value)
	return attribute
}

func (attribute *Attribute) SetNullValue() *Attribute {
	attribute.Value = NewNullValueAny()
	return attribute
}

func (attribute *Attribute) SetNullFile() *Attribute {
	attribute.File = &NullValue[AttributeFile]{null: true}
	return attribute
}

func (attribute Attribute) String() string {
	return Stringify(attribute)
}

func (attribute Attribute) MetaType() MetaType {
	return MetaTypeAttribute
}

// AttributeFile содержит имя файла и файл, закодированный в base64
type AttributeFile struct {
	Filename string `json:"filename"` // Имя файла
	Content  string `json:"content"`  // Байты файла, закодированные в base64
}

func NewAttributeFromFile(filePath string) (*Attribute, error) {
	b, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	fileName := filepath.Base(filePath)
	content := base64.StdEncoding.EncodeToString(b)
	file := AttributeFile{fileName, content}
	Attribute := &Attribute{File: &NullValue[AttributeFile]{value: file}}

	return Attribute, nil

}

func NewAttributes() Slice[Attribute] {
	return NewSlice[Attribute]()
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
