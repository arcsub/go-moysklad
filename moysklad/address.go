package moysklad

// Address Адрес с детализацией по отдельным полям.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-sklady-attributy-suschnosti-adres
type Address struct {
	AddInfo    *string  `json:"addInfo,omitempty"`    // Другое
	Apartment  *string  `json:"apartment,omitempty"`  // Квартира
	City       *string  `json:"city,omitempty"`       // Город
	Comment    *string  `json:"comment,omitempty"`    // Комментарий
	Country    *Country `json:"country,omitempty"`    // Метаданные страны
	House      *string  `json:"house,omitempty"`      // Дом
	PostalCode *string  `json:"postalCode,omitempty"` // Почтовый индекс
	Region     *Region  `json:"region,omitempty"`     // Метаданные региона
	Street     *string  `json:"street,omitempty"`     // Улица
}

func (a Address) String() string {
	return Stringify(a)
}
