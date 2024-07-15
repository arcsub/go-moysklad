package moysklad

// Address Адрес с детализацией по отдельным полям.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-sklady-attributy-suschnosti-adres
type Address struct {
	AddInfo    *string      `json:"addInfo,omitempty"`    // Другое
	Apartment  *string      `json:"apartment,omitempty"`  // Квартира
	City       *string      `json:"city,omitempty"`       // Город
	Comment    *string      `json:"comment,omitempty"`    // Комментарий
	Country    *MetaWrapper `json:"country,omitempty"`    // Метаданные Страны
	House      *string      `json:"house,omitempty"`      // Дом
	PostalCode *string      `json:"postalCode,omitempty"` // Почтовый индекс
	Region     *MetaWrapper `json:"region,omitempty"`     // Метаданные региона
	Street     *string      `json:"street,omitempty"`     // Улица
}

// GetAddInfo возвращает Другое.
func (address Address) GetAddInfo() string {
	return Deref(address.AddInfo)
}

// GetApartment возвращает Квартиру.
func (address Address) GetApartment() string {
	return Deref(address.Apartment)
}

// GetCity возвращает Город.
func (address Address) GetCity() string {
	return Deref(address.City)
}

// GetComment возвращает Комментарий.
func (address Address) GetComment() string {
	return Deref(address.Comment)
}

// GetCountry возвращает Метаданные страны.
func (address Address) GetCountry() MetaWrapper {
	return Deref(address.Country)
}

// GetHouse возвращает Дом.
func (address Address) GetHouse() string {
	return Deref(address.House)
}

// GetPostalCode возвращает Почтовый индекс.
func (address Address) GetPostalCode() string {
	return Deref(address.PostalCode)
}

// GetRegion возвращает Метаданные региона.
func (address Address) GetRegion() MetaWrapper {
	return Deref(address.Region)
}

// GetStreet возвращает Улицу.
func (address Address) GetStreet() string {
	return Deref(address.Street)
}

// SetAddInfo устанавливает Другое.
func (address *Address) SetAddInfo(addInfo string) *Address {
	address.AddInfo = &addInfo
	return address
}

// SetApartment устанавливает Квартиру.
func (address *Address) SetApartment(apartment string) *Address {
	address.Apartment = &apartment
	return address
}

// SetCity устанавливает Город.
func (address *Address) SetCity(city string) *Address {
	address.City = &city
	return address
}

// SetComment устанавливает Комментарий.
func (address *Address) SetComment(comment string) *Address {
	address.Comment = &comment
	return address
}

// SetCountry устанавливает Метаданные страны.
func (address *Address) SetCountry(country *Country) *Address {
	if country != nil {
		mw := country.GetMeta().Wrap()
		address.Country = &mw
	}
	return address
}

// SetHouse устанавливает Дом.
func (address *Address) SetHouse(house string) *Address {
	address.House = &house
	return address
}

// SetPostalCode устанавливает Почтовый индекс.
func (address *Address) SetPostalCode(postalCode string) *Address {
	address.PostalCode = &postalCode
	return address
}

// SetRegion устанавливает Метаданные региона.
func (address *Address) SetRegion(region *Region) *Address {
	if region != nil {
		mw := region.GetMeta().Wrap()
		address.Region = &mw
	}
	return address
}

// SetStreet устанавливает Улицу.
func (address *Address) SetStreet(street string) *Address {
	address.Street = &street
	return address
}

// String реализует интерфейс [fmt.Stringer].
func (address Address) String() string {
	return Stringify(address)
}
