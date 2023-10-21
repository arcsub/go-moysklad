package moysklad

// CompanySettings Настройки компании.
// Ключевое слово: companysettings
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-nastrojki-kompanii
type CompanySettings struct {
	Meta                     *Meta            `json:"meta,omitempty"`
	CheckMinPrice            *bool            `json:"checkMinPrice,omitempty"`
	CheckShippingStock       *bool            `json:"checkShippingStock,omitempty"`
	CompanyAddress           *string          `json:"companyAddress,omitempty"`
	Currency                 *Currency        `json:"currency,omitempty"`
	GlobalOperationNumbering *bool            `json:"globalOperationNumbering,omitempty"`
	PriceTypes               *PriceTypes      `json:"priceTypes,omitempty"`
	UseCompanyAddress        *bool            `json:"useCompanyAddress,omitempty"`
	UseRecycleBin            *bool            `json:"useRecycleBin,omitempty"`
	DiscountStrategy         DiscountStrategy `json:"discountStrategy,omitempty"`
}

func (c CompanySettings) String() string {
	return Stringify(c)
}

func (c CompanySettings) MetaType() MetaType {
	return MetaTypeCompanySettings
}

// DiscountStrategy Совместное применение скидок
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-nastrojki-kompanii-sowmestnoe-primenenie-skidok
type DiscountStrategy string

const (
	DiscountStrategyBySum      DiscountStrategy = "bySum"      // Сумма скидок (должна действовать сумма скидок)
	DiscountStrategyByPriority DiscountStrategy = "byPriority" // Приоритетная (должна действовать одна, наиболее выгодная для покупателя скидка)
)
