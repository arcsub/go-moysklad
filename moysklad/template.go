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

func (t Template) String() string {
	return Stringify(t)
}

func (t Template) GetType() TemplateType {
	return t.Type
}

func (t Template) GetMeta() *Meta {
	return t.Meta
}

// CustomTemplate Пользовательский Шаблон.
// Ключевое слово: customtemplate
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-shablon-pechatnoj-formy-atributy-suschnosti
type CustomTemplate struct {
	Template
}

func (c CustomTemplate) String() string {
	return Stringify(c)
}

func (c CustomTemplate) MetaType() MetaType {
	return MetaTypeCustomTemplate
}

// EmbeddedTemplate Стандартный шаблон
// Ключевое слово: embeddedtemplate
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-shablon-pechatnoj-formy-atributy-suschnosti
type EmbeddedTemplate struct {
	Template
}

func (e EmbeddedTemplate) String() string {
	return Stringify(e)
}

func (e EmbeddedTemplate) MetaType() MetaType {
	return MetaTypeEmbeddedTemplate
}

type Templater interface {
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
	Template  *MetaWrapper                   `json:"template,omitempty"`
	Templates *Slice[PrintDocArgManyElement] `json:"templates,omitempty"`
	Extension Extension                      `json:"extension,omitempty"`
}

type PrintDocArgManyElement struct {
	Template *MetaWrapper `json:"template,omitempty"`
	Count    int          `json:"count,omitempty"`
}

// NewPrintDocArgOne создаёт и возвращает заполненный объект для запроса печати документов.
// Применяется только для запроса 1 документа.
// Возможные расширения:
// – ExtensionXLS для документа с расширением .xls
// – ExtensionPDF для документа с расширением .pdf
// – ExtensionHTML для документа с расширением .html
// – ExtensionODS для документа с расширением .ods
func NewPrintDocArgOne(template Templater, ext Extension) *PrintDocumentArg {
	return &PrintDocumentArg{
		Template:  &MetaWrapper{Meta: Deref(template.GetMeta())},
		Extension: ext,
	}
}

// NewPrintDocArgManyElement создаёт и возвращает *PrintDocArgManyElement,
// который служит аргументом для метода NewPrintDocArgMany
func NewPrintDocArgManyElement(template Templater, count int) *PrintDocArgManyElement {
	return &PrintDocArgManyElement{
		Template: &MetaWrapper{Deref(template.GetMeta())},
		Count:    count,
	}
}

// NewPrintDocArgMany создаёт и возвращает заполненный объект для запроса печати документов.
// Применяется для запроса комплекта документов.
// Каждый аргумент создаётся с помощью метода NewPrintDocArgManyElement
func NewPrintDocArgMany(templates ...*PrintDocArgManyElement) *PrintDocumentArg {
	return &PrintDocumentArg{
		Templates: (new(Slice[PrintDocArgManyElement])).Push(templates...),
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
func NewPrintLabelArg(organization *Organization, priceType *PriceType, template Templater, count int) *PrintLabelArg {
	return &PrintLabelArg{
		Organization: MetaWrapper{Meta: Deref(organization.Meta)},
		SalePrice:    PrintLabelArgSalePrice{PriceType: MetaWrapper{Meta: Deref(priceType.Meta)}},
		Template:     MetaWrapper{Meta: Deref(template.GetMeta())},
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
