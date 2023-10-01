package moysklad

// EmbeddedTemplate Стандартный шаблон
// Ключевое слово: embeddedtemplate
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-shablon-pechatnoj-formy-atributy-suschnosti
type EmbeddedTemplate struct {
	Template
}

func (e EmbeddedTemplate) String() string {
	return Stringify(e)
}

func (e EmbeddedTemplate) MetaType() MetaType {
	return MetaTypeEmbeddedTemplate
}
