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
	CustomEntityMeta *Meta         `json:"customEntityMeta,omitempty"`
	Description      *string       `json:"description,omitempty"`
	ID               *uuid.UUID    `json:"id,omitempty"`
	Meta             *Meta         `json:"meta,omitempty"`
	Name             *string       `json:"name,omitempty"`
	Required         *bool         `json:"required,omitempty"`
	Show             *bool         `json:"show,omitempty"`
	Type             AttributeType `json:"type,omitempty"`
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

func (attribute Attribute) String() string {
	return Stringify(attribute)
}

func (attribute Attribute) MetaType() MetaType {
	return MetaTypeAttribute
}

// attributeValueFile содержит имя файла и файл, закодированный в base64
type attributeValueFile struct {
	Filename string `json:"filename"` // Имя файла
	Content  string `json:"content"`  // Байты файла, закодированные в base64
}

// AttributeValue Доп. поле со значением.
// Ключевое слово: attributemetadata
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-obschie-swedeniq-rabota-s-dopolnitel-nymi-polqmi-atributy-dop-polq-so-znacheniem
type AttributeValue struct {
	Value    any                 `json:"value,omitempty"`
	Meta     *Meta               `json:"meta,omitempty"`
	ID       *uuid.UUID          `json:"id,omitempty"`
	Name     *string             `json:"name,omitempty"`
	File     *attributeValueFile `json:"file,omitempty"`
	Download *Meta               `json:"download,omitempty"`
	Type     AttributeType       `json:"type,omitempty"`
}

func NewAttributeValueFromFile(filePath string) (*AttributeValue, error) {
	b, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	fileName := filepath.Base(filePath)
	content := base64.StdEncoding.EncodeToString(b)
	file := &attributeValueFile{fileName, content}
	attributeValue := &AttributeValue{File: file}

	return attributeValue, nil

}

func (attributeValue AttributeValue) GetMeta() Meta {
	return Deref(attributeValue.Meta)
}

func (attributeValue AttributeValue) GetID() uuid.UUID {
	return Deref(attributeValue.ID)
}

func (attributeValue AttributeValue) GetName() string {
	return Deref(attributeValue.Name)
}

func (attributeValue AttributeValue) GetType() AttributeType {
	return attributeValue.Type
}

func (attributeValue AttributeValue) GetValue() any {
	return attributeValue.Value
}

func (attributeValue AttributeValue) GetDownload() Meta {
	return Deref(attributeValue.Meta)
}

func (attributeValue AttributeValue) String() string {
	return Stringify(attributeValue)
}

func (attributeValue *AttributeValue) SetMeta(meta *Meta) *AttributeValue {
	attributeValue.Meta = meta
	return attributeValue
}

func (attributeValue *AttributeValue) SetName(name string) *AttributeValue {
	attributeValue.Name = &name
	return attributeValue
}

func (attributeValue *AttributeValue) SetType(attributeType AttributeType) *AttributeValue {
	attributeValue.Type = attributeType
	return attributeValue
}

func (attributeValue *AttributeValue) SetValue(value any) *AttributeValue {
	attributeValue.Value = value
	return attributeValue
}

type Attributes = Slice[AttributeValue]

func NewAttributes() Attributes {
	return make(Attributes, 0)
}

// AttributeValueType Тип значения атрибута
type AttributeValueType interface {
	bool | // Флажок AttributeTypeBoolean (boolean)
		float64 | // Число дробное AttributeTypeDouble (double)
		int | // Число целое AttributeTypeLong (long)
		string | // Ссылка AttributeTypeLink (link), Строка AttributeTypeString (string), Текст AttributeTypeText (text), Файл AttributeTypeFile (file)
		Timestamp | // Дата AttributeTypeTime (time)
		MetaName // Справочник AttributeTypeDictionary*
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
