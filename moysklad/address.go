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

func (address *Address) SetAddInfo(addInfo string) *Address {
	address.AddInfo = &addInfo
	return address
}

func (address *Address) SetApartment(apartment string) *Address {
	address.Apartment = &apartment
	return address
}

func (address *Address) SetCity(city string) *Address {
	address.City = &city
	return address
}

func (address *Address) SetComment(comment string) *Address {
	address.Comment = &comment
	return address
}

func (address *Address) SetCountry(country *Country) *Address {
	address.Country = &Country{Meta: country.Meta}
	return address
}

func (address *Address) SetHouse(house string) *Address {
	address.House = &house
	return address
}

func (address *Address) SetPostalCode(postalCode string) *Address {
	address.PostalCode = &postalCode
	return address
}

func (address *Address) SetRegion(region *Region) *Address {
	address.Region = &Region{Meta: region.Meta}
	return address
}

func (address *Address) SetStreet(street string) *Address {
	address.Street = &street
	return address
}

func (address Address) GetAddInfo() string {
	return Deref(address.AddInfo)
}

func (address Address) GetApartment() string {
	return Deref(address.Apartment)
}

func (address Address) GetCity() string {
	return Deref(address.City)
}

func (address Address) GetComment() string {
	return Deref(address.Comment)
}

func (address Address) GetCountry() Country {
	return Deref(address.Country)
}

func (address Address) GetHouse() string {
	return Deref(address.House)
}

func (address Address) GetPostalCode() string {
	return Deref(address.PostalCode)
}

func (address Address) GetRegion() Region {
	return Deref(address.Region)
}

func (address Address) GetStreet() string {
	return Deref(address.Street)
}

func (address Address) String() string {
	return Stringify(address)
}
