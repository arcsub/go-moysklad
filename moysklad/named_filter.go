package moysklad

// NamedFilter Сохраненный фильтр.
//
// Код сущности: namedfilter
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sohranennye-fil-try
type NamedFilter struct {
	AccountID *string   `json:"accountId,omitempty"` // ID учётной записи
	ID        *string   `json:"id,omitempty"`        // ID фильтра
	Meta      *Meta     `json:"meta,omitempty"`      // Метаданные фильтра
	Name      *string   `json:"name,omitempty"`      // Наименование фильтра
	Owner     *Employee `json:"owner,omitempty"`     // Метаданные владельца (Сотрудника)
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (namedFilter NamedFilter) Clean() *NamedFilter {
	if namedFilter.Meta == nil {
		return nil
	}
	return &NamedFilter{Meta: namedFilter.Meta}
}

// GetAccountID возвращает ID учётной записи.
func (namedFilter NamedFilter) GetAccountID() string {
	return Deref(namedFilter.AccountID)
}

// GetID возвращает ID фильтра.
func (namedFilter NamedFilter) GetID() string {
	return Deref(namedFilter.ID)
}

// GetMeta возвращает Метаданные фильтра.
func (namedFilter NamedFilter) GetMeta() Meta {
	return Deref(namedFilter.Meta)
}

// GetName возвращает Наименование фильтра.
func (namedFilter NamedFilter) GetName() string {
	return Deref(namedFilter.Name)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (namedFilter NamedFilter) GetOwner() Employee {
	return Deref(namedFilter.Owner)
}

// String реализует интерфейс [fmt.Stringer].
func (namedFilter NamedFilter) String() string {
	return Stringify(namedFilter)
}

// MetaType возвращает код сущности.
func (NamedFilter) MetaType() MetaType {
	return MetaTypeNamedFilter
}
