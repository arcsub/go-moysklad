package moysklad

// CompanyType Тип Контрагента
type CompanyType string

const (
	CompanyLegal        CompanyType = "legal"        // Юридическое лицо
	CompanyEntrepreneur CompanyType = "entrepreneur" // Индивидуальный предприниматель
	CompanyIndividual   CompanyType = "individual"   // Физическое лицо
)
