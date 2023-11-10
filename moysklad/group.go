package moysklad

import (
	"github.com/google/uuid"
)

// Group Отдел.
// Ключевое слово: group
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-otdel
type Group struct {
	AccountID *uuid.UUID `json:"accountId,omitempty"` // ID учетной записи
	ID        *uuid.UUID `json:"id,omitempty"`        // ID Отдела
	Index     *int       `json:"index,omitempty"`     // Порядковый номер в списке отделов
	Meta      *Meta      `json:"meta,omitempty"`      // Метаданные Отдела
	Name      *string    `json:"name,omitempty"`      // Наименование Отдела
}

func (g Group) String() string {
	return Stringify(g)
}

func (g Group) MetaType() MetaType {
	return MetaTypeGroup
}
