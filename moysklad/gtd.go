package moysklad

// GTD Грузовая таможенная декларация (ГТД).
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-gruzowaq-tamozhennaq-deklaraciq-gtd
type GTD struct {
	Name *string `json:"name"` // Номер ГТД
}

// GetName возвращает Номер ГТД.
func (gtd GTD) GetName() string {
	return Deref(gtd.Name)
}

// SetName устанавливает Номер ГТД.
func (gtd *GTD) SetName(name string) *GTD {
	gtd.Name = &name
	return gtd
}

// String реализует интерфейс [fmt.Stringer].
func (gtd GTD) String() string {
	return Stringify(gtd)
}
