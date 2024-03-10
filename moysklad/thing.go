package moysklad

import "github.com/google/uuid"

// Thing Серийный номер
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-serijnyj-nomer
type Thing struct {
	AccountId   *uuid.UUID `json:"accountId,omitempty"`   // ID учетной записи
	Description *string    `json:"description,omitempty"` // Описание Серийного номера
	ID          *uuid.UUID `json:"id,omitempty"`          // ID Серийного номера
	Meta        *Meta      `json:"meta,omitempty"`        // Метаданные о Серийном номере
	Name        *string    `json:"name,omitempty"`        // Наименование Серийного номера
}

func (t Thing) String() string {
	return Stringify(t)
}

func (t Thing) MetaType() MetaType {
	return MetaTypeThing
}

// Things серийные номера.
type Things = []string
