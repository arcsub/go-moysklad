package moysklad

// GTD Грузовая таможенная декларация (ГТД).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-gruzowaq-tamozhennaq-deklaraciq-gtd
type GTD struct {
	Name *string `json:"name"` // Номер ГТД
}

func (gtd GTD) GetName() string {
	return Deref(gtd.Name)
}

func (gtd *GTD) SetName(name string) *GTD {
	gtd.Name = &name
	return gtd
}

func (gtd GTD) String() string {
	return Stringify(gtd)
}
