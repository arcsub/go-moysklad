package moysklad

import (
	"github.com/google/uuid"
)

// Consignment Серия.
// Ключевое слово: consignment
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-seriq
type Consignment struct {
	AccountID    *uuid.UUID          `json:"accountId,omitempty"`    // ID учетной записи
	Barcodes     *Barcodes           `json:"barcodes,omitempty"`     // Штрихкоды
	Code         *string             `json:"code,omitempty"`         // Код
	Description  *string             `json:"description,omitempty"`  // Описание
	ExternalCode *string             `json:"externalCode,omitempty"` // Внешний код
	ID           *uuid.UUID          `json:"id,omitempty"`           // ID сущности
	Meta         *Meta               `json:"meta,omitempty"`         // Метаданные
	Name         *string             `json:"name,omitempty"`         // Наименование
	Assortment   *AssortmentPosition `json:"assortment,omitempty"`   //
	Attributes   *Attributes         `json:"attributes,omitempty"`   // Метаданные ссылки или модификации
	Image        *Image              `json:"image,omitempty"`        // Изображение товара, к которому относится данная серия
	Label        *string             `json:"label,omitempty"`        // Метка Серии
	Updated      *Timestamp          `json:"updated,omitempty"`      // Момент последнего обновления сущности
}

func (c Consignment) String() string {
	return Stringify(c)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (c Consignment) GetMeta() *Meta {
	return c.Meta
}

func (c Consignment) MetaType() MetaType {
	return MetaTypeConsignment
}

func (c Consignment) ConvertToAssortmentPosition() (*AssortmentPosition, error) {
	return convertToAssortmentPosition(c)
}
