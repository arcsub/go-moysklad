package moysklad

// UniqueCodeRules Настройки уникальности кода для сущностей справочника.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-assortiment-atributy-wlozhennyh-suschnostej-nastrojki-unikal-nosti-koda-dlq-suschnostej-sprawochnika
type UniqueCodeRules struct {
	CheckUniqueCode *bool `json:"checkUniqueCode,omitempty"` // Проверка уникальности кода сущностей справочника товаров
	FillUniqueCode  *bool `json:"fillUniqueCode,omitempty"`  // Устанавливать уникальный код при создании создании сущностей справочника
}

// GetCheckUniqueCode возвращает флаг проверки уникальности кода сущностей справочника товаров.
func (uniqueCodeRules UniqueCodeRules) GetCheckUniqueCode() bool {
	return Deref(uniqueCodeRules.CheckUniqueCode)
}

// GetFillUniqueCode возвращает флаг установки уникального кода при создании создании сущностей справочника.
func (uniqueCodeRules UniqueCodeRules) GetFillUniqueCode() bool {
	return Deref(uniqueCodeRules.FillUniqueCode)
}

// SetCheckUniqueCode устанавливает флаг проверки уникальности кода сущностей справочника товаров.
func (uniqueCodeRules *UniqueCodeRules) SetCheckUniqueCode(checkUniqueCode bool) *UniqueCodeRules {
	uniqueCodeRules.CheckUniqueCode = &checkUniqueCode
	return uniqueCodeRules
}

// SetFillUniqueCode устанавливает флаг установки уникального кода при создании создании сущностей справочника.
func (uniqueCodeRules *UniqueCodeRules) SetFillUniqueCode(fillUniqueCode bool) *UniqueCodeRules {
	uniqueCodeRules.FillUniqueCode = &fillUniqueCode
	return uniqueCodeRules
}

// String реализует интерфейс [fmt.Stringer].
func (uniqueCodeRules UniqueCodeRules) String() string {
	return Stringify(uniqueCodeRules)
}
