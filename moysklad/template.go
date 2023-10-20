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

type PrintDocRequest struct {
	Template  *MetaWrapper             `json:"template,omitempty"`
	Templates *Slice[PrintDocTemplate] `json:"templates,omitempty"`
	Extension Extension                `json:"extension,omitempty"`
}

type PrintDocTemplate struct {
	Template *MetaWrapper `json:"template,omitempty"`
	Count    int          `json:"count,omitempty"`
}

// Ценники

type PrintPriceArg struct {
	Organization MetaWrapper            `json:"organization,omitempty"` // Метаданные Юрлица
	Count        int                    `json:"count,omitempty"`        // Количество ценников/термоэтикеток
	SalePrice    PrintPriceArgSalePrice `json:"salePrice,omitempty"`    // Цена продажи
	Template     MetaWrapper            `json:"template,omitempty"`     // Метаданные Шаблона печати
}

type PrintPriceArgSalePrice struct {
	PriceType MetaWrapper `json:"priceType,omitempty"` // Метаданные типа цены
}

// NewPrintPriceArg создаёт и возвращает заполненный объект для запроса печати ценников.
// Аргументы: организация, тип цен, шаблон и кол-во ценников
func NewPrintPriceArg(organization *Organization, priceType *PriceType, template Templater, count int) *PrintPriceArg {
	return &PrintPriceArg{
		Organization: MetaWrapper{Meta: Deref(organization.Meta)},
		SalePrice:    PrintPriceArgSalePrice{PriceType: MetaWrapper{Meta: Deref(priceType.Meta)}},
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

	if !strings.HasSuffix(f.FileName, ".pdf") {
		f.FileName += ".pdf"
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
