package moysklad

// GoodTaxSystem Код системы налогообложения.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-towar-towary-atributy-suschnosti-kod-sistemy-nalogooblozheniq
type GoodTaxSystem string

const (
	GoodTaxSystemGeneralTaxSystem                 GoodTaxSystem = "GENERAL_TAX_SYSTEM"                   // ОСН
	GoodTaxSystemSimplifiedTaxSystemIncome        GoodTaxSystem = "SIMPLIFIED_TAX_SYSTEM_INCOME"         // УСН. Доход
	GoodTaxSystemSimplifiedTaxSystemIncomeOutcome GoodTaxSystem = "SIMPLIFIED_TAX_SYSTEM_INCOME_OUTCOME" // УСН. Доход-Расход
	GoodTaxSystemUnifiedAgriculturalTax           GoodTaxSystem = "UNIFIED_AGRICULTURAL_TAX"             // ЕСХН
	GoodTaxSystemPresumptiveTaxSystem             GoodTaxSystem = "PRESUMPTIVE_TAX_SYSTEM"               // ЕНВД
	GoodTaxSystemPatentBased                      GoodTaxSystem = "PATENT_BASED"                         // Патент
	GoodTaxSystemSameAsGroup                      GoodTaxSystem = "TAX_SYSTEM_SAME_AS_GROUP"             // Совпадает с группой
)

// TaxSystem Код системы налогообложения по умолчанию.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tochka-prodazh-tochki-prodazh-atributy-suschnosti-kod-sistemy-nalogooblozheniq-po-umolchaniu
type TaxSystem string

const (
	GeneralTaxSystem                 TaxSystem = "GENERAL_TAX_SYSTEM"                   // ОСН
	SimplifiedTaxSystemIncome        TaxSystem = "SIMPLIFIED_TAX_SYSTEM_INCOME"         // УСН. Доход
	SimplifiedTaxSystemIncomeOutcome TaxSystem = "SIMPLIFIED_TAX_SYSTEM_INCOME_OUTCOME" // УСН. Доход-Расход
	UnifiedAgriculturalTax           TaxSystem = "UNIFIED_AGRICULTURAL_TAX"             // ЕСХН
	PresumptiveTaxSystem             TaxSystem = "PRESUMPTIVE_TAX_SYSTEM"               // ЕНВД
	PatentBased                      TaxSystem = "PATENT_BASED"                         // Патент
)
