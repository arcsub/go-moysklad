package moysklad

import "github.com/shopspring/decimal"

// ReportCounterparty Показатели контрагентов.
// Ключевое слово: counterparty
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pokazateli-kontragentow-pokazateli-kontragentow
type ReportCounterparty struct {
	FirstDemandDate Timestamp        `json:"firstDemandDate"`
	Updated         Timestamp        `json:"updated"`
	LastEventDate   Timestamp        `json:"lastEventDate"`
	LastDemandDate  Timestamp        `json:"lastDemandDate"`
	Counterparty    CounterpartyData `json:"counterparty"`
	Meta            Meta             `json:"meta"`
	DiscountsSum    decimal.Decimal  `json:"discountsSum"`
	LastEventText   string           `json:"lastEventText"`
	DemandsSum      decimal.Decimal  `json:"demandsSum"`
	AverageReceipt  decimal.Decimal  `json:"averageReceipt"`
	BonusBalance    decimal.Decimal  `json:"bonusBalance"`
	Profit          decimal.Decimal  `json:"profit"`
	ReturnsSum      decimal.Decimal  `json:"returnsSum"`
	Balance         decimal.Decimal  `json:"balance"`
	DemandsCount    int              `json:"demandsCount"`
	ReturnsCount    int              `json:"returnsCount"`
}

// CounterpartyData Контрагент
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pokazateli-kontragentow-dopolnitel-nye-atributy-dostupnye-dlq-fil-tracii-kontragent
type CounterpartyData struct {
	CompanyType  CompanyType `json:"companyType"`  // Тип контрагента
	ExternalCode string      `json:"externalCode"` // Внешний код контрагента
	ID           string      `json:"id"`           // ID Контрагента
	Meta         Meta        `json:"meta"`         // Метаданные Контрагента
	Name         string      `json:"name"`         // Наименование Контрагента
}

func (r ReportCounterparty) MetaType() MetaType {
	return MetaTypeReportCounterparty
}

type CounterpartyElement struct {
	Counterparty MetaWrapper `json:"counterparty"`
}

type CounterpartiesMeta struct {
	Counterparties Slice[CounterpartyElement] `json:"counterparties"`
}

func (c *CounterpartiesMeta) Push(elements ...*Counterparty) {
	for _, element := range elements {
		ce := &CounterpartyElement{
			Counterparty: MetaWrapper{
				Meta: *element.GetMeta(),
			},
		}
		c.Counterparties = append(c.Counterparties, ce)
	}
}
