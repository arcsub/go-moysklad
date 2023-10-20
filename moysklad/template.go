package moysklad

import (
	"bytes"
	"github.com/google/uuid"
	"io"
	"os"
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

type Templater interface {
	HasMeta
	Type() TemplateType
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

type PrintPriceRequest struct {
	Organization MetaWrapper                       `json:"organization,omitempty"` // Метаданные Юрлица
	Count        int                               `json:"count,omitempty"`        // Количество ценников/термоэтикеток
	SalePrice    PrintPriceRequestPriceTypeWrapper `json:"salePrice,omitempty"`    // Цена продажи
	Template     MetaWrapper                       `json:"template,omitempty"`     // Метаданные Шаблона печати
}

type PrintPriceRequestPriceTypeWrapper struct {
	PriceType MetaWrapper `json:"priceType,omitempty"` // Метаданные типа цены
}

type PrintFile struct {
	*bytes.Buffer
	FileName string
}

func (f *PrintFile) SaveToFile(filename string) error {
	if len(filename) == 0 {
		filename = f.FileName
	}

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
