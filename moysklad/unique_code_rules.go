package moysklad

// UniqueCodeRules Настройки уникальности кода для сущностей справочника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-assortiment-atributy-wlozhennyh-suschnostej-nastrojki-unikal-nosti-koda-dlq-suschnostej-sprawochnika
type UniqueCodeRules struct {
	CheckUniqueCode *bool `json:"checkUniqueCode,omitempty"` // Проверка уникальности кода сущностей справочника товаров
	FillUniqueCode  *bool `json:"fillUniqueCode,omitempty"`  // Устанавливать уникальный код при создании создании сущностей справочника товаров
}

func (u UniqueCodeRules) String() string {
	return Stringify(u)
}
