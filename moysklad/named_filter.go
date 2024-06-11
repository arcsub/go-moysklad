package moysklad

import (
	"github.com/google/uuid"
)

// NamedFilter Сохраненный фильтр.
// Ключевое слово: namedfilter
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sohranennye-fil-try
type NamedFilter struct {
	AccountID *uuid.UUID `json:"accountId,omitempty"` // ID учетной записи
	ID        *uuid.UUID `json:"id,omitempty"`        // ID сущности
	Meta      *Meta      `json:"meta,omitempty"`      // Метаданные
	Name      *string    `json:"name,omitempty"`      // Наименование
	Owner     *Employee  `json:"owner,omitempty"`     // Владелец (Сотрудник)
}

func (namedFilter NamedFilter) Clean() *NamedFilter {
	return &NamedFilter{Meta: namedFilter.Meta}
}

func (namedFilter NamedFilter) GetAccountID() uuid.UUID {
	return Deref(namedFilter.AccountID)
}

func (namedFilter NamedFilter) GetID() uuid.UUID {
	return Deref(namedFilter.ID)
}

func (namedFilter NamedFilter) GetMeta() Meta {
	return Deref(namedFilter.Meta)
}

func (namedFilter NamedFilter) GetName() string {
	return Deref(namedFilter.Name)
}

func (namedFilter NamedFilter) GetOwner() Employee {
	return Deref(namedFilter.Owner)
}

func (namedFilter NamedFilter) String() string {
	return Stringify(namedFilter)
}

func (namedFilter NamedFilter) MetaType() MetaType {
	return MetaTypeNamedFilter
}
