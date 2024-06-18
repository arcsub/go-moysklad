package moysklad

import (
	"bytes"
	"github.com/google/uuid"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Template Общие поля для шаблонов.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-shablon-pechatnoj-formy
type Template struct {
	Content *string      `json:"content,omitempty"` // Ссылка на скачивание
	ID      *uuid.UUID   `json:"id,omitempty"`      // ID сущности
	Meta    *Meta        `json:"meta,omitempty"`    // Метаданные Стандартного шаблона
	Name    *string      `json:"name,omitempty"`    // Наименование шаблона
	Type    TemplateType `json:"type,omitempty"`    // Тип шаблона (entity - документ, mxtemplate - новый тип шаблона для ценников и этикеток)
}

func (template Template) GetContent() string {
	return Deref(template.Content)
}

func (template Template) GetID() uuid.UUID {
	return Deref(template.ID)
}

func (template Template) GetMeta() Meta {
	return Deref(template.Meta)
}

func (template Template) GetName() string {
	return Deref(template.Name)
}

func (template Template) GetType() TemplateType {
	return template.Type
}

func (template *Template) SetContent(content string) *Template {
	template.Content = &content
	return template
}

func (template *Template) SetMeta(meta *Meta) *Template {
	template.Meta = meta
	return template
}

func (template *Template) SetName(name string) *Template {
	template.Name = &name
	return template
}

func (template *Template) SetType(templateType TemplateType) *Template {
	template.Type = templateType
	return template
}

func (template Template) String() string {
	return Stringify(template)
}

// CustomTemplate Пользовательский Шаблон.
// Ключевое слово: customtemplate
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-shablon-pechatnoj-formy-atributy-suschnosti
type CustomTemplate struct {
	Template
}

func (customTemplate CustomTemplate) String() string {
	return Stringify(customTemplate)
}

func (customTemplate CustomTemplate) MetaType() MetaType {
	return MetaTypeCustomTemplate
}

// EmbeddedTemplate Стандартный шаблон
// Ключевое слово: embeddedtemplate
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-shablon-pechatnoj-formy-atributy-suschnosti
type EmbeddedTemplate struct {
	Template
}

func (embeddedTemplate EmbeddedTemplate) String() string {
	return Stringify(embeddedTemplate)
}

func (embeddedTemplate EmbeddedTemplate) MetaType() MetaType {
	return MetaTypeEmbeddedTemplate
}

type TemplateInterface interface {
	HasMeta
	GetType() TemplateType
}

type TemplateType string

const (
	TemplateTypeEntity     TemplateType = "entity"     // Документ
	TemplateTypePriceType  TemplateType = "pricetype"  // Ценник/этикетка
	TemplateTypeMXTemplate TemplateType = "mxtemplate" // Ценник/этикетка нового формата
)

type Extension string

const (
	ExtensionXLS  Extension = "xls"
	ExtensionPDF  Extension = "pdf"
	ExtensionHTML Extension = "html"
	ExtensionODS  Extension = "ods"
)

// Документы

type PrintDocumentArg struct {
	Template  MetaWrapper                   `json:"template,omitempty"`
	Extension Extension                     `json:"extension,omitempty"`
	Templates Slice[PrintDocArgManyElement] `json:"templates,omitempty"`
}

type PrintDocArgManyElement struct {
	Template MetaWrapper `json:"template,omitempty"`
	Count    int         `json:"count,omitempty"`
}

// NewPrintDocArgOne создаёт и возвращает заполненный объект для запроса печати документов.
// Применяется только для запроса 1 документа.
// Возможные расширения:
// – ExtensionXLS для документа с расширением .xls
// – ExtensionPDF для документа с расширением .pdf
// – ExtensionHTML для документа с расширением .html
// – ExtensionODS для документа с расширением .ods
func NewPrintDocArgOne(template TemplateInterface, ext Extension) *PrintDocumentArg {
	return &PrintDocumentArg{
		Template:  template.GetMeta().Wrap(),
		Extension: ext,
	}
}

// NewPrintDocArgManyElement создаёт и возвращает *PrintDocArgManyElement,
// который служит аргументом для метода NewPrintDocArgMany
func NewPrintDocArgManyElement(template TemplateInterface, count int) *PrintDocArgManyElement {
	return &PrintDocArgManyElement{
		Template: template.GetMeta().Wrap(),
		Count:    count,
	}
}

// NewPrintDocArgMany создаёт и возвращает заполненный объект для запроса печати документов.
// Применяется для запроса комплекта документов.
// Каждый аргумент создаётся с помощью метода NewPrintDocArgManyElement
func NewPrintDocArgMany(templates ...*PrintDocArgManyElement) *PrintDocumentArg {
	return &PrintDocumentArg{
		Templates: templates,
	}
}

// Ценники

type PrintLabelArg struct {
	Organization MetaWrapper            `json:"organization,omitempty"`
	SalePrice    PrintLabelArgSalePrice `json:"salePrice,omitempty"`
	Template     MetaWrapper            `json:"template,omitempty"`
	Count        int                    `json:"count,omitempty"`
}

type PrintLabelArgSalePrice struct {
	PriceType MetaWrapper `json:"priceType,omitempty"` // Метаданные типа цены
}

// NewPrintLabelArg создаёт и возвращает заполненный объект для запроса печати ценников.
// Аргументы: организация, тип цен, шаблон и кол-во ценников
func NewPrintLabelArg(organization *Organization, priceType *PriceType, template TemplateInterface, count int) *PrintLabelArg {
	return &PrintLabelArg{
		Organization: organization.GetMeta().Wrap(),
		SalePrice:    PrintLabelArgSalePrice{PriceType: priceType.GetMeta().Wrap()},
		Template:     template.GetMeta().Wrap(),
		Count:        count,
	}
}

type PrintFile struct {
	*bytes.Buffer
	FileName string
}

// Save сохраняет полученный файл в указанную папку.
// Аргументом является директория, в которую необходимо сохранить файл.
func (f *PrintFile) Save(path string) error {
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}

	if err := os.MkdirAll(filepath.Dir(path), 0770); err != nil {
		return err
	}

	filename := path + f.FileName
	fo, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer func(fo *os.File) {
		_ = fo.Close()
	}(fo)

	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}

		if n == 0 {
			break
		}

		if _, err = fo.Write(buf[:n]); err != nil {
			return err
		}
	}
	return nil
}
