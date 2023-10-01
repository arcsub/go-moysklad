package moysklad

import (
	"github.com/google/uuid"
)

// NamedFilter Сохраненный фильтр.
// Ключевое слово: namedfilter
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sohranennye-fil-try
type NamedFilter struct {
	AccountId *uuid.UUID `json:"accountId,omitempty"` // ID учетной записи
	Id        *uuid.UUID `json:"id,omitempty"`        // ID сущности
	Meta      *Meta      `json:"meta,omitempty"`      // Метаданные
	Name      *string    `json:"name,omitempty"`      // Наименование
	Owner     *Employee  `json:"owner,omitempty"`     // Владелец (Сотрудник)
}

func (n NamedFilter) String() string {
	return Stringify(n)
}

func (n NamedFilter) MetaType() MetaType {
	return MetaTypeNamedFilter
}
