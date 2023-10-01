package moysklad

import (
	"github.com/google/uuid"
)

// Characteristic Характеристика
// Ключевое слово: attributemetadata
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-modifikaciq-modifikacii-atributy-wlozhennyh-suschnostej-metadannye-modifikacij-harakteristiki-modifikacii
type Characteristic struct {
	Id       *uuid.UUID `json:"id,omitempty"`       // ID соответствующей характеристики
	Meta     *Meta      `json:"meta,omitempty"`     // Метаданные характеристики
	Name     *string    `json:"name,omitempty"`     // Наименование характеристики
	Required *bool      `json:"required,omitempty"` // Флаг о том, является ли характеристика обязательной
	Type     *string    `json:"type,omitempty"`     // Тип значения характеристики (значение всегда "string")
	Value    *string    `json:"value,omitempty"`    // Значение характеристики
}

func (c Characteristic) String() string {
	return Stringify(c)
}

func (c Characteristic) MetaType() MetaType {
	return MetaTypeCharacteristic
}

type Characteristics = Iterator[Characteristic]
