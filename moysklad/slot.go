package moysklad

import (
	"github.com/google/uuid"
)

// Slot Ячейка склада.
// Ключевое слово: slot
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-yachejki-sklada
type Slot struct {
	AccountId    *uuid.UUID `json:"accountId,omitempty"`    // ID учетной записи
	ExternalCode *string    `json:"externalCode,omitempty"` // Внешний код Ячейки
	Id           *uuid.UUID `json:"id,omitempty"`           // ID Ячейки
	Meta         *Meta      `json:"meta,omitempty"`         // Метаданные Ячейки
	Name         *string    `json:"name,omitempty"`         // Наименование Ячейки
	Updated      *Timestamp `json:"updated,omitempty"`      // Момент последнего обновления Ячейки
	Zone         *Zone      `json:"zone,omitempty"`         // Зона ячейки
}

func (s Slot) String() string {
	return Stringify(s)
}

func (s Slot) MetaType() MetaType {
	return MetaTypeSlot
}
