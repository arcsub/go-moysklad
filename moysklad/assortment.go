package moysklad

import (
	"encoding/json"
	"errors"
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
	// Общие поля
	AccountId    uuid.UUID `json:"accountId,omitempty"`    // ID учетной записи
	Barcodes     Barcodes  `json:"barcodes,omitempty"`     // Штрихкоды
	Code         string    `json:"code,omitempty"`         // Код
	Description  string    `json:"description,omitempty"`  // Описание
	ExternalCode string    `json:"externalCode,omitempty"` // Внешний код
	Id           uuid.UUID `json:"id,omitempty"`           // ID сущности
	Meta         Meta      `json:"meta"`                   // Метаданные
	Name         string    `json:"name,omitempty"`         // Наименование

	// сырые данные
	data json.RawMessage
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
func (a Assortment) FilterBundle() []Bundle {
	return filterEntity[Bundle](a)
}

// FilterProduct фильтрует позиции по типу Product (Товар)
func (a Assortment) FilterProduct() []Product {
	return filterEntity[Product](a)
}

// FilterVariant фильтрует позиции по типу Variant (Модификация)
func (a Assortment) FilterVariant() []Variant {
	return filterEntity[Variant](a)
}

// FilterConsignment фильтрует позиции по типу Consignment (Серия)
func (a Assortment) FilterConsignment() []Consignment {
	return filterEntity[Consignment](a)
}

// FilterService фильтрует позиции по типу Service (Услуга)
func (a Assortment) FilterService() []Service {
	return filterEntity[Service](a)
}

func filterEntity[E MetaTyper, A DataMetaTyper](elements []A) []E {
	var n []E
	for _, el := range elements {
		if e, err := unmarshalTo[E](el); err == nil {
			n = append(n, *e)
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
