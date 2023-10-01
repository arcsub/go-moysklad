package moysklad

// Publication Публикация документов.
// Ключевое слово: operationpublication
// Публикации доступны только для следующих типов: Заказ покупателя, Счет покупателю, Отгрузка,
// Заказ поставщику, Счет поставщика, Приемка, Входящий платеж, Приходный ордер, Исходящий платеж,
// Расходный ордер, Внутренний заказ, Перемещение, Оприходование, Списание, Счет-фактура выданный,
// Счет-фактура полученный, Возврат поставщику, Возврат покупателя, Выплата денег, Внесение денег,
// Розничный возврат, Розничная продажа, Договор, Розничная смена, Заказ на производство,
// Полученный отчет комиссионера, Выданный отчет комиссионера, Инвентаризация, Техоперация.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-publikaciq-dokumentow
type Publication struct {
	Meta     *Meta        `json:"meta,omitempty"`     // Метаданные Публикации
	Template *MetaWrapper `json:"template,omitempty"` // Метаданные Шаблона печати
	Href     *string      `json:"href,omitempty"`     // Ссылка на страницу Публикации
}

func (p Publication) String() string {
	return Stringify(p)
}

func (p Publication) MetaType() MetaType {
	return MetaTypePublication
}

func (p *Publication) SetTemplate(template *Templater) *Publication {
	p.Template = &MetaWrapper{Meta: *(*template).GetMeta()}
	return p
}
