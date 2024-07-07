package moysklad

// Publication Публикация документов.
//
// Код сущности: operationpublication
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-publikaciq-dokumentow
type Publication struct {
	Meta     *Meta     `json:"meta,omitempty"`     // Метаданные Публикации
	Template *Template `json:"template,omitempty"` // Метаданные Шаблона печати
	Href     *string   `json:"href,omitempty"`     // Ссылка на страницу Публикации
}

// GetMeta возвращает Метаданные Публикации.
func (publication Publication) GetMeta() Meta {
	return Deref(publication.Meta)
}

// GetTemplate возвращает Метаданные Шаблона печати.
func (publication Publication) GetTemplate() Template {
	return Deref(publication.Template)
}

// GetHref возвращает Ссылка на страницу Публикации.
func (publication Publication) GetHref() string {
	return Deref(publication.Href)
}

// String реализует интерфейс [fmt.Stringer].
func (publication Publication) String() string {
	return Stringify(publication)
}

// MetaType возвращает код сущности.
func (Publication) MetaType() MetaType {
	return MetaTypePublication
}
