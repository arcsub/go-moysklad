package moysklad

import (
	"github.com/google/uuid"
)

// Attribute Дополнительное поле.
//
// Код сущности: attributemetadata
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-obschie-swedeniq-rabota-s-dopolnitel-nymi-polqmi
type Attribute struct {
	CustomEntityMeta *Meta                     `json:"customEntityMeta,omitempty"` // Метаданные пользовательского справочника
	Description      *string                   `json:"description,omitempty"`      // Описание доп. поля
	ID               *uuid.UUID                `json:"id,omitempty"`               // ID доп. поля [Обязательное при ответе] [Только для чтения]
	Meta             *Meta                     `json:"meta,omitempty"`             // Метаданные доп. поля [Обязательное при ответе]
	Name             *string                   `json:"name,omitempty"`             // Наименование доп. поля [Обязательное при ответе] [Необходимо при создании]
	Required         *bool                     `json:"required,omitempty"`         // Является ли доп. поле обязательным [Обязательное при ответе]
	Show             *bool                     `json:"show,omitempty"`             // Показывать ли доп. поле на UI. Не может быть скрытым и обязательным одновременно. Только для операций
	Value            *NullValueAny             `json:"value,omitempty"`            // Значение, указанное в доп. поле
	File             *NullValue[AttributeFile] `json:"file,omitempty"`             // Описание файла и контент (поле доступно только для доп.поля типа Файл)
	Download         *Meta                     `json:"download,omitempty"`         // Метаданные, содержащие ссылку на скачивание файла. Поле отображается только для типа доп поля AttributeTypeFile (Файл)
	Type             AttributeType             `json:"type,omitempty"`             // Тип доп. поля [Обязательное при ответе] [Необходимо при создании] [После заполнения недоступно для изменения]
}

// GetCustomEntityMeta возвращает Метаданные пользовательского справочника.
func (attribute Attribute) GetCustomEntityMeta() Meta {
	return Deref(attribute.CustomEntityMeta)
}

// GetDescription возвращает Описание доп. поля.
func (attribute Attribute) GetDescription() string {
	return Deref(attribute.Description)
}

// GetID возвращает ID доп. поля.
func (attribute Attribute) GetID() uuid.UUID {
	return Deref(attribute.ID)
}

// GetMeta возвращает Метаданные доп. поля.
func (attribute Attribute) GetMeta() Meta {
	return Deref(attribute.Meta)
}

// GetName возвращает Наименование доп. поля.
func (attribute Attribute) GetName() string {
	return Deref(attribute.Name)
}

// GetRequired возвращает true, если доп. поле является обязательным.
func (attribute Attribute) GetRequired() bool {
	return Deref(attribute.Required)
}

// GetShow возвращает true, если доп. поле отображается в UI.
func (attribute Attribute) GetShow() bool {
	return Deref(attribute.Show)
}

// GetType возвращает Тип доп. поля.
func (attribute Attribute) GetType() AttributeType {
	return attribute.Type
}

// GetValue возвращает Значение, указанное в доп. поле.
func (attribute Attribute) GetValue() any {
	return attribute.Value.Get()
}

// GetDownload возвращает Метаданные, содержащие ссылку на скачивание файла.
func (attribute Attribute) GetDownload() Meta {
	return Deref(attribute.Meta)
}

// GetFile возвращает Описание файла и контент.
func (attribute Attribute) GetFile() AttributeFile {
	return Deref(attribute.File).GetValue()
}

// SetMeta устанавливает Метаданные доп. поля.
func (attribute *Attribute) SetMeta(meta *Meta) *Attribute {
	attribute.Meta = meta
	return attribute
}

// SetName устанавливает Наименование доп. поля.
func (attribute *Attribute) SetName(name string) *Attribute {
	attribute.Name = &name
	return attribute
}

// SetType устанавливает Тип доп. поля.
func (attribute *Attribute) SetType(attributeType AttributeType) *Attribute {
	attribute.Type = attributeType
	return attribute
}

// SetValue устанавливает Значение доп. поля.
//
// Передача nil передаёт сброс значения (null).
func (attribute *Attribute) SetValue(value any) *Attribute {
	if value == nil {
		attribute.Value = NewNullValueAny()
	} else {
		attribute.Value = NewNullValueAnyFrom(value)
	}
	return attribute
}

// SetFile устанавливает Описание файла и контент.
//
// Передача nil передаёт сброс значения (null).
func (attribute *Attribute) SetFile(file *AttributeFile) *Attribute {
	attribute.File = NewNullValue(file)
	return attribute
}

// SetFileFromPath устанавливает Описание файла и контент.
//
// Принимает путь до файла.
//
// Возвращает [Attribute] и ошибку, если она произошла при чтении файла.
func (attribute *Attribute) SetFileFromPath(filePath string) (*Attribute, error) {
	fileName, content, err := getFilenameContent(filePath)
	if err != nil {
		return nil, err
	}
	file := &AttributeFile{fileName, content}
	attribute.File = NewNullValue(file)
	return attribute, nil
}

// String реализует интерфейс [fmt.Stringer].
func (attribute Attribute) String() string {
	return Stringify(attribute)
}

// MetaType возвращает код сущности.
func (Attribute) MetaType() MetaType {
	return MetaTypeAttribute
}

// AttributeFile Дополнительное поле типа AttributeTypeFile (файл).
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/workbook/#workbook-rabota-s-dopolnitel-nymi-polqmi-cherez-json-api-dopolnitel-noe-pole-tipa-fajl
type AttributeFile struct {
	Filename string `json:"filename"` // Имя файла [Обязательное при ответе] [Необходимо при создании]
	Content  string `json:"content"`  // Байты файла, закодированные в base64 [Обязательное при ответе] [Необходимо при создании]
}

// NewAttributeFromFile возвращает предзаполненный [Attribute] типа [AttributeTypeFile] (Файл) и ошибку,
// если она произошла при чтении файла.
//
// Принимает путь до файла.
func NewAttributeFromFile(filePath string) (*Attribute, error) {
	fileName, content, err := getFilenameContent(filePath)
	if err != nil {
		return nil, err
	}
	file := AttributeFile{fileName, content}
	attribute := &Attribute{File: &NullValue[AttributeFile]{value: &file}, Type: AttributeTypeFile}
	return attribute, nil

}

// NewAttributes возвращает пустой срез для удобной работы с множеством объектов типа [Attribute].
func NewAttributes() Slice[Attribute] {
	return NewSlice[Attribute]()
}

// AttributeType Тип доп. поля.
//
// Дополнительные поля типов Файл и Флажок не могут быть обязательными (поле required не может быть true).
//
// Возможные варианты:
//   - AttributeTypeTime                   – Дата
//   - AttributeTypeLink                   – Ссылка
//   - AttributeTypeString                 – Строка
//   - AttributeTypeText                   – Текст
//   - AttributeTypeFile                   – Файл
//   - AttributeTypeBoolean                – Флажок
//   - AttributeTypeDouble                 – Число дробное
//   - AttributeTypeLong                   – Число целое
//   - AttributeTypeDictionaryContract     – Справочник типа Договор [Contract]
//   - AttributeTypeDictionaryCounterParty – Справочник типа Контрагент [Counterparty]
//   - AttributeTypeDictionaryProject      – Справочник типа Проект [Project]
//   - AttributeTypeDictionaryStore        – Справочник типа Склад [Store]
//   - AttributeTypeDictionaryEmployee     – Справочник типа Сотрудник [Employee]
//   - AttributeTypeDictionaryProduct      – Справочник типа Товар [Product]
//   - AttributeTypeDictionaryCustom       – Пользовательский справочник [CustomEntity]
type AttributeType string

const (
	AttributeTypeTime                   AttributeType = "time"         // Дата
	AttributeTypeLink                   AttributeType = "link"         // Ссылка
	AttributeTypeString                 AttributeType = "string"       // Строка
	AttributeTypeText                   AttributeType = "text"         // Текст
	AttributeTypeFile                   AttributeType = "file"         // Файл
	AttributeTypeBoolean                AttributeType = "boolean"      // Флажок
	AttributeTypeDouble                 AttributeType = "double"       // Число дробное
	AttributeTypeLong                   AttributeType = "long"         // Число целое
	AttributeTypeDictionaryContract     AttributeType = "contract"     // Справочник типа Договор
	AttributeTypeDictionaryCounterParty AttributeType = "counterparty" // Справочник типа Контрагент
	AttributeTypeDictionaryProject      AttributeType = "project"      // Справочник типа Проект
	AttributeTypeDictionaryStore        AttributeType = "store"        // Справочник типа Склад
	AttributeTypeDictionaryEmployee     AttributeType = "employee"     // Справочник типа Сотрудник
	AttributeTypeDictionaryProduct      AttributeType = "product"      // Справочник типа Товар
	AttributeTypeDictionaryCustom       AttributeType = "customentity" // Пользовательский справочник
)
