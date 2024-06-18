package moysklad

// Publication Публикация документов.
// Ключевое слово: operationpublication
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-publikaciq-dokumentow
type Publication struct {
	Meta     *Meta     `json:"meta,omitempty"`     // Метаданные Публикации
	Template *Template `json:"template,omitempty"` // Метаданные Шаблона печати
	Href     *string   `json:"href,omitempty"`     // Ссылка на страницу Публикации
}

func (publication Publication) GetMeta() Meta {
	return Deref(publication.Meta)
}

func (publication Publication) GetTemplate() Template {
	return Deref(publication.Template)
}

func (publication Publication) GetHref() string {
	return Deref(publication.Href)
}

func (publication *Publication) SetTemplate(template TemplateInterface) *Publication {
	meta := template.GetMeta()
	publication.Template = &Template{Meta: &meta}
	return publication
}

func (publication Publication) String() string {
	return Stringify(publication)
}

func (publication Publication) MetaType() MetaType {
	return MetaTypePublication
}
