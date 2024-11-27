package moysklad

import (
	"bytes"

	"io"
	"os"
	"path/filepath"
	"strings"
)

// Template Общие поля для шаблонов.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-shablon-pechatnoj-formy
type Template struct {
	Content *string      `json:"content,omitempty"` // Ссылка на скачивание
	ID      *string      `json:"id,omitempty"`      // ID шаблона
	Meta    *Meta        `json:"meta,omitempty"`    // Метаданные шаблона
	Name    *string      `json:"name,omitempty"`    // Наименование шаблона
	Type    TemplateType `json:"type,omitempty"`    // Тип шаблона (entity - документ, mxtemplate - новый тип шаблона для ценников и этикеток)
}

// GetContent возвращает Ссылку на скачивание.
func (template Template) GetContent() string {
	return Deref(template.Content)
}

// GetID возвращает ID шаблона.
func (template Template) GetID() string {
	return Deref(template.ID)
}

// GetMeta возвращает Метаданные шаблона.
func (template Template) GetMeta() Meta {
	return Deref(template.Meta)
}

// GetName возвращает Наименование шаблона.
func (template Template) GetName() string {
	return Deref(template.Name)
}

// GetType возвращает Тип шаблона.
func (template Template) GetType() TemplateType {
	return template.Type
}

// String реализует интерфейс [fmt.Stringer].
func (template Template) String() string {
	return Stringify(template)
}

// CustomTemplate Пользовательский Шаблон.
//
// Код сущности: customtemplate
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-shablon-pechatnoj-formy-atributy-suschnosti
type CustomTemplate struct {
	Template
}

// AsTemplate реализует интерфейс [TemplateConverter].
func (customTemplate CustomTemplate) AsTemplate() *TemplateOwner {
	return newTemplateOwner(customTemplate.Meta)
}

// SetMeta устанавливает Метаданные шаблона.
func (customTemplate *CustomTemplate) SetMeta(meta *Meta) *CustomTemplate {
	customTemplate.Meta = meta
	return customTemplate
}

// SetName устанавливает Наименование шаблона.
func (customTemplate *CustomTemplate) SetName(name string) *CustomTemplate {
	customTemplate.Name = &name
	return customTemplate
}

// SetType устанавливает Тип шаблона.
func (customTemplate *CustomTemplate) SetType(templateType TemplateType) *CustomTemplate {
	customTemplate.Type = templateType
	return customTemplate
}

// SetTypeEntity устанавливает Тип шаблона в значение [TemplateTypeEntity].
func (customTemplate *CustomTemplate) SetTypeEntity() *CustomTemplate {
	customTemplate.Type = TemplateTypeEntity
	return customTemplate
}

// SetTypePriceType устанавливает Тип шаблона в значение [TemplateTypePriceType].
func (customTemplate *CustomTemplate) SetTypePriceType() *CustomTemplate {
	customTemplate.Type = TemplateTypePriceType
	return customTemplate
}

// SetTypeMXTemplate устанавливает Тип шаблона в значение [TemplateTypeMXTemplate].
func (customTemplate *CustomTemplate) SetTypeMXTemplate() *CustomTemplate {
	customTemplate.Type = TemplateTypeMXTemplate
	return customTemplate
}

// String реализует интерфейс [fmt.Stringer].
func (customTemplate CustomTemplate) String() string {
	return Stringify(customTemplate)
}

// MetaType возвращает код сущности.
func (CustomTemplate) MetaType() MetaType {
	return MetaTypeCustomTemplate
}

// EmbeddedTemplate Стандартный шаблон
//
// Код сущности: embeddedtemplate
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-shablon-pechatnoj-formy-atributy-suschnosti
type EmbeddedTemplate struct {
	Template
}

// AsTemplate реализует интерфейс [TemplateConverter].
func (embeddedTemplate EmbeddedTemplate) AsTemplate() *TemplateOwner {
	return newTemplateOwner(embeddedTemplate.Meta)
}

// String реализует интерфейс [fmt.Stringer].
func (embeddedTemplate EmbeddedTemplate) String() string {
	return Stringify(embeddedTemplate)
}

// MetaType возвращает код сущности.
func (EmbeddedTemplate) MetaType() MetaType {
	return MetaTypeEmbeddedTemplate
}

// TemplateConverter описывает метод, который возвращает [TemplateOwner].
//
// Интерфейс должны реализовывать: [CustomTemplate] и [EmbeddedTemplate].
type TemplateConverter interface {
	AsTemplate() *TemplateOwner
}

type TemplateOwner struct {
	Template Template `json:"template"`
}

func newTemplateOwner(meta *Meta) *TemplateOwner {
	return &TemplateOwner{Template: Template{Meta: meta}}
}

// TemplateType тип шаблона.
//
// Возможные значения:
//   - TemplateTypeEntity     – Документ
//   - TemplateTypePriceType  – Ценник/этикетка
//   - TemplateTypeMXTemplate – Ценник/этикетка нового формата
type TemplateType string

const (
	TemplateTypeEntity     TemplateType = "entity"     // Документ
	TemplateTypePriceType  TemplateType = "pricetype"  // Ценник/этикетка
	TemplateTypeMXTemplate TemplateType = "mxtemplate" // Ценник/этикетка нового формата
)

// Extension расширение документа.
//
// Возможные значения:
//   - XLS  – для документа с расширением .xls
//   - PDF  – для документа с расширением .pdf
//   - HTML – для документа с расширением .html
//   - ODS  – для документа с расширением .ods
type Extension string

const (
	XLS  Extension = "xls"  // для документа с расширением .xls
	PDF  Extension = "pdf"  // для документа с расширением .pdf
	HTML Extension = "html" // для документа с расширением .html
	ODS  Extension = "ods"  // для документа с расширением .ods
)

// Документы

type PrintDocumentArg struct {
	Template  *TemplateOwner          `json:"template,omitempty"`  // Метаданные Шаблона печати
	Extension Extension               `json:"extension,omitempty"` // Расширение, в котором нужно напечатать форму
	Templates Slice[PrintDocTemplate] `json:"templates,omitempty"` // Метаданные Шаблонов печати
}

type PrintDocTemplate struct {
	Template *TemplateOwner `json:"template,omitempty"`
	Count    int            `json:"count,omitempty"`
}

// NewPrintDocArgOne создаёт и возвращает заполненный объект для запроса печати документов.
//
// Применяется только для запроса на печать одного документа.
func NewPrintDocArgOne(template TemplateConverter, ext Extension) *PrintDocumentArg {
	printDocumentArg := &PrintDocumentArg{Extension: ext}

	if template != nil {
		printDocumentArg.Template = template.AsTemplate()
	}

	return printDocumentArg
}

// NewPrintDocTemplate создаёт и возвращает [PrintDocTemplate], который служит аргументом для функции [NewPrintDocArgMany].
func NewPrintDocTemplate(template TemplateConverter, count int) *PrintDocTemplate {
	printDocArgManyElement := &PrintDocTemplate{Count: count}

	if template != nil {
		printDocArgManyElement.Template = template.AsTemplate()
	}

	return printDocArgManyElement
}

// NewPrintDocArgMany создаёт и возвращает заполненный объект для запроса печати документов.
//
// Применяется для запроса комплекта документов.
//
// Каждый аргумент создаётся с помощью функции [NewPrintDocTemplate].
func NewPrintDocArgMany(templates ...*PrintDocTemplate) *PrintDocumentArg {
	printDocumentArg := new(PrintDocumentArg)
	printDocumentArg.Templates.Push(templates...)
	return printDocumentArg
}

// Ценники

type PrintLabelArg struct {
	Organization MetaWrapper    `json:"organization,omitempty"` // Метаданные Юрлица
	SalePrice    PriceTypeOwner `json:"salePrice,omitempty"`    // Цена продажи
	Template     *TemplateOwner `json:"template,omitempty"`     // Метаданные Шаблона печати
	Count        int            `json:"count,omitempty"`        // Количество ценников/термоэтикеток. Максимальное количество - 1000
}

type PriceTypeOwner struct {
	PriceType MetaWrapper `json:"priceType,omitempty"` // Метаданные типа цены
}

// NewPrintLabelArg создаёт и возвращает заполненный объект для запроса печати ценников.
//
// Аргументы: организация, тип цен, шаблон и кол-во ценников.
func NewPrintLabelArg(organization *Organization, priceType *PriceType, template TemplateConverter, count int) *PrintLabelArg {
	printLabelArg := &PrintLabelArg{Count: Clamp(count, 1, MaxPrintCount)}

	if organization != nil {
		printLabelArg.Organization = organization.GetMeta().Wrap()
	}

	if priceType != nil {
		printLabelArg.SalePrice = PriceTypeOwner{priceType.GetMeta().Wrap()}
	}

	if template != nil {
		printLabelArg.Template = template.AsTemplate()
	}

	return printLabelArg
}

type PrintFile struct {
	*bytes.Buffer
	FileName string
}

// Save сохраняет полученный файл в указанную папку.
//
// Принимает путь (директорию), по которому необходимо сохранить файл.
func (printFile *PrintFile) Save(path string) error {
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}

	if err := os.MkdirAll(filepath.Dir(path), 0770); err != nil {
		return err
	}

	filename := path + printFile.FileName
	fo, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer func(fo *os.File) {
		_ = fo.Close()
	}(fo)

	buf := make([]byte, 1024)
	for {
		n, err := printFile.Read(buf)
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
