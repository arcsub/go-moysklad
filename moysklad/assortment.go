package moysklad

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Assortment Ассортимент.
// Ключевое слово: assortment
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-assortiment
type Assortment Slice[AssortmentPosition]

func (a Assortment) MetaType() MetaType {
	return MetaTypeAssortment
}

// AssortmentPosition позиция ассортимента.
// Представляет из себя структуру из полей:
// Meta для хранения метаданных сущности
// data для хранения сырых данных
// Product | Variant | Bundle | Service | Consignment
type AssortmentPosition struct {
	Meta         Meta     `json:"meta"`
	Code         string   `json:"code,omitempty"`
	Description  string   `json:"description,omitempty"`
	ExternalCode string   `json:"externalCode,omitempty"`
	Name         string   `json:"name,omitempty"`
	Barcodes     Barcodes `json:"barcodes,omitempty"`
	data         json.RawMessage
	AccountID    uuid.UUID `json:"accountId,omitempty"`
	ID           uuid.UUID `json:"id,omitempty"`
}

type AssortmentPositionTypes interface {
	Product | Variant | Bundle | Service | Consignment
	HasMeta
}

func (a *AssortmentPosition) String() string {
	return Stringify(a.Meta)
}

// MetaType удовлетворяет интерфейсу MetaTyper
func (a AssortmentPosition) MetaType() MetaType {
	return a.Meta.Type
}

// Data удовлетворяет интерфейсу DataMetaTyper
func (a AssortmentPosition) Data() json.RawMessage {
	return a.data
}

func (a *AssortmentPosition) UnmarshalJSON(data []byte) error {
	type alias AssortmentPosition
	var t alias
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	t.data = data
	*a = AssortmentPosition(t)
	return nil
}

// ConvertToProduct структурирует сущность в *Product
// Возвращает ошибку в случае неудачи
func (a *AssortmentPosition) ConvertToProduct() (*Product, error) {
	return unmarshalTo[Product](a)
}

// ConvertToVariant структурирует позицию в *Variant
// Возвращает ошибку в случае неудачи
func (a *AssortmentPosition) ConvertToVariant() (*Variant, error) {
	return unmarshalTo[Variant](a)
}

// ConvertToBundle структурирует позицию в *Bundle
// Возвращает ошибку в случае неудачи
func (a *AssortmentPosition) ConvertToBundle() (*Bundle, error) {
	return unmarshalTo[Bundle](a)
}

// ConvertToService структурирует позицию в *Service
// Возвращает ошибку в случае неудачи
func (a *AssortmentPosition) ConvertToService() (*Service, error) {
	return unmarshalTo[Service](a)
}

// ConvertToConsignment структурирует позицию в *Consignment
// Возвращает ошибку в случае неудачи
func (a *AssortmentPosition) ConvertToConsignment() (*Consignment, error) {
	return unmarshalTo[Consignment](a)
}

func convertToAssortmentPosition[E AssortmentPositionTypes](element E) (*AssortmentPosition, error) {
	meta := element.GetMeta()
	if meta == nil {
		return nil, errors.New("meta is nil")
	}
	data, err := json.Marshal(element)
	if err != nil {
		return nil, err
	}
	position := &AssortmentPosition{Meta: *meta, data: data}
	return position, nil
}

// FilterBundle фильтрует позиции по типу Bundle (Комплект)
func (a Assortment) FilterBundle() Slice[Bundle] {
	return filterEntity[Bundle](a)
}

// FilterProduct фильтрует позиции по типу Product (Товар)
func (a Assortment) FilterProduct() Slice[Product] {
	return filterEntity[Product](a)
}

// FilterVariant фильтрует позиции по типу Variant (Модификация)
func (a Assortment) FilterVariant() Slice[Variant] {
	return filterEntity[Variant](a)
}

// FilterConsignment фильтрует позиции по типу Consignment (Серия)
func (a Assortment) FilterConsignment() Slice[Consignment] {
	return filterEntity[Consignment](a)
}

// FilterService фильтрует позиции по типу Service (Услуга)
func (a Assortment) FilterService() Slice[Service] {
	return filterEntity[Service](a)
}

func filterEntity[E MetaTyper, A DataMetaTyper](elements []A) Slice[E] {
	var n Slice[E]
	for _, el := range elements {
		if e, err := unmarshalTo[E](el); err == nil {
			n = append(n, e)
		}
	}
	return n
}

// AssortmentSettings Настройки справочника.
// Ключевое слово: assortmentsettings
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-assortiment-nastrojki-sprawochnika
type AssortmentSettings struct {
	Meta            *Meta            `json:"meta,omitempty"`            // Метаданные Настроек справочника
	BarcodeRules    *BarcodeRules    `json:"barcodeRules,omitempty"`    // Настройки правил штрихкодов для сущностей справочника
	UniqueCodeRules *UniqueCodeRules `json:"uniqueCodeRules,omitempty"` // Настройки уникальности кода для сущностей справочника
	CreatedShared   *bool            `json:"createdShared,omitempty"`   // Создавать новые документы с меткой «Общий»
}

func (a AssortmentSettings) String() string {
	return Stringify(a)
}

func (a AssortmentSettings) MetaType() MetaType {
	return MetaTypeAssortmentSettings
}

// BarcodeRules Настройки правил штрихкодов для сущностей справочника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-assortiment-atributy-wlozhennyh-suschnostej-nastrojki-prawil-shtrihkodow-dlq-suschnostej-sprawochnika
type BarcodeRules struct {
	FillEAN13Barcode    *bool `json:"fillEAN13Barcode,omitempty"`    // Автоматически создавать штрихкод EAN13 для новых товаров, комплектов, модификаций и услуг
	WeightBarcode       *bool `json:"weightBarcode,omitempty"`       // Использовать префиксы штрихкодов для весовых товаров
	WeightBarcodePrefix *int  `json:"weightBarcodePrefix,omitempty"` // Префикс штрихкодов для весовых товаров. Возможные значения: число формата X или XX
}

func (b BarcodeRules) String() string {
	return Stringify(b)
}

// AssortmentService
// Сервис для работы с ассортиментом.
type AssortmentService interface {
	Get(ctx context.Context, params *Params) (*AssortmentResult, *resty.Response, error)
	GetAsync(ctx context.Context) (AsyncResultService[AssortmentResult], *resty.Response, error)
	DeleteMany(ctx context.Context, entities []*AssortmentPosition) (*DeleteManyResponse, *resty.Response, error)
	GetSettings(ctx context.Context) (*AssortmentSettings, *resty.Response, error)
	UpdateSettings(ctx context.Context, settings *AssortmentSettings) (*AssortmentSettings, *resty.Response, error)
	GetEmbeddedTemplates(ctx context.Context) (*List[EmbeddedTemplate], *resty.Response, error)
	GetEmbeddedTemplateByID(ctx context.Context, id *uuid.UUID) (*EmbeddedTemplate, *resty.Response, error)
	GetCustomTemplates(ctx context.Context) (*List[CustomTemplate], *resty.Response, error)
	GetCustomTemplateByID(ctx context.Context, id *uuid.UUID) (*CustomTemplate, *resty.Response, error)
}

type assortmentService struct {
	endpointGetOne[AssortmentResult]
	endpointGetOneAsync[AssortmentResult]
	endpointDeleteMany[AssortmentPosition]
	endpointSettings[AssortmentSettings]
	endpointPrintTemplates
}

func NewAssortmentService(client *Client) AssortmentService {
	e := NewEndpoint(client, "entity/assortment")
	return &assortmentService{
		endpointGetOne:         endpointGetOne[AssortmentResult]{e},
		endpointGetOneAsync:    endpointGetOneAsync[AssortmentResult]{e},
		endpointDeleteMany:     endpointDeleteMany[AssortmentPosition]{e},
		endpointSettings:       endpointSettings[AssortmentSettings]{e},
		endpointPrintTemplates: endpointPrintTemplates{e},
	}
}
