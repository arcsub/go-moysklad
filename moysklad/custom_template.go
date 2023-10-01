package moysklad

// CustomTemplate Пользовательский Шаблон.
// Ключевое слово: customtemplate
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-shablon-pechatnoj-formy-atributy-suschnosti
type CustomTemplate struct {
	Template
}

func (c CustomTemplate) String() string {
	return Stringify(c)
}

func (c CustomTemplate) MetaType() MetaType {
	return MetaTypeCustomTemplate
}
