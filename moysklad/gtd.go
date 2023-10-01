package moysklad

// GTD Грузовая таможенная декларация (ГТД).
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-gruzowaq-tamozhennaq-deklaraciq-gtd
type GTD struct {
	Name *string `json:"name"` // Номер ГТД
}

func (g GTD) String() string {
	return Stringify(g)
}
