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
	Country    *MetaWrapper `json:"country,omitempty"`    // Метаданные страны
	House      *string      `json:"house,omitempty"`      // Дом
	PostalCode *string      `json:"postalCode,omitempty"` // Почтовый индекс
	Region     *MetaWrapper `json:"region,omitempty"`     // Метаданные региона
	Street     *string      `json:"street,omitempty"`     // Улица
}

// GetAddInfo возвращает значение поля AddInfo (Другое).
func (address Address) GetAddInfo() string {
	return Deref(address.AddInfo)
}

// GetApartment возвращает значение поля Apartment (Квартира).
func (address Address) GetApartment() string {
	return Deref(address.Apartment)
}

// GetCity возвращает значение поля City (Город).
func (address Address) GetCity() string {
	return Deref(address.City)
}

// GetComment возвращает значение поля Comment (Комментарий).
func (address Address) GetComment() string {
	return Deref(address.Comment)
}

// GetCountry возвращает значение поля Country (Метаданные страны).
func (address Address) GetCountry() MetaWrapper {
	return Deref(address.Country)
}

// GetHouse возвращает значение поля House (Дом).
func (address Address) GetHouse() string {
	return Deref(address.House)
}

// GetPostalCode возвращает значение поля PostalCode (Почтовый индекс).
func (address Address) GetPostalCode() string {
	return Deref(address.PostalCode)
}

// GetRegion возвращает значение поля Region (Метаданные региона).
func (address Address) GetRegion() MetaWrapper {
	return Deref(address.Region)
}

// GetStreet возвращает значение поля Street (Улица).
func (address Address) GetStreet() string {
	return Deref(address.Street)
}

// SetAddInfo устанавливает значение поля AddInfo (Другое) и возвращает указатель на [Address].
func (address *Address) SetAddInfo(addInfo string) *Address {
	address.AddInfo = &addInfo
	return address
}

// SetApartment устанавливает значение поля Apartment (Квартира) и возвращает указатель на [Address].
func (address *Address) SetApartment(apartment string) *Address {
	address.Apartment = &apartment
	return address
}

// SetCity устанавливает значение поля City (Город) и возвращает указатель на [Address].
func (address *Address) SetCity(city string) *Address {
	address.City = &city
	return address
}

// SetComment устанавливает значение поля Comment (Комментарий) и возвращает указатель на [Address].
func (address *Address) SetComment(comment string) *Address {
	address.Comment = &comment
	return address
}

// SetCountry устанавливает значение поля Country (Метаданные страны) и возвращает указатель на [Address].
func (address *Address) SetCountry(country *Country) *Address {
	mw := country.GetMeta().Wrap()
	address.Country = &mw
	return address
}

// SetHouse устанавливает значение поля House (Дом) и возвращает указатель на [Address].
func (address *Address) SetHouse(house string) *Address {
	address.House = &house
	return address
}

// SetPostalCode устанавливает значение поля PostalCode (Почтовый индекс) и возвращает указатель на [Address].
func (address *Address) SetPostalCode(postalCode string) *Address {
	address.PostalCode = &postalCode
	return address
}

// SetRegion устанавливает значение поля Region (Метаданные региона) и возвращает указатель на [Address].
func (address *Address) SetRegion(region *Region) *Address {
	mw := region.GetMeta().Wrap()
	address.Region = &mw
	return address
}

// SetStreet устанавливает значение поля Street (Улица) и возвращает указатель на [Address].
func (address *Address) SetStreet(street string) *Address {
	address.Street = &street
	return address
}

// String реализует интерфейс [fmt.Stringer].
func (address Address) String() string {
	return Stringify(address)
}
