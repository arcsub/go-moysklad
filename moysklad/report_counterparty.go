package moysklad

// ReportCounterparty Показатели контрагентов.
// Ключевое слово: counterparty
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/reports/#otchety-otchet-pokazateli-kontragentow-pokazateli-kontragentow
type ReportCounterparty struct {
	AverageReceipt  float64          `json:"averageReceipt"`  // Средний чек
	Balance         float64          `json:"balance"`         // Баланс
	BonusBalance    float64          `json:"bonusBalance"`    // Баллы
	Counterparty    CounterpartyData `json:"counterparty"`    // Контрагент
	DemandsCount    int              `json:"demandsCount"`    // Количество продаж
	DemandsSum      float64          `json:"demandsSum"`      // Сумма продаж
	DiscountsSum    float64          `json:"discountsSum"`    // Сумма скидок
	FirstDemandDate Timestamp        `json:"firstDemandDate"` // Дата первой продажи
	LastDemandDate  Timestamp        `json:"lastDemandDate"`  // Дата последней продажи
	LastEventDate   Timestamp        `json:"lastEventDate"`   // Дата последнего события
	LastEventText   string           `json:"lastEventText"`   // Текст последнего события
	Meta            Meta             `json:"meta"`            // Метаданные Отчета по данному контрагенту
	Profit          float64          `json:"profit"`          // Прибыль
	ReturnsCount    int              `json:"returnsCount"`    // Количество возвратов
	ReturnsSum      float64          `json:"returnsSum"`      // Сумма возвратов
	Updated         Timestamp        `json:"updated"`         // Момент последнего изменения контрагента
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
		ce := CounterpartyElement{
			Counterparty: MetaWrapper{
				Meta: *element.GetMeta(),
			},
		}
		c.Counterparties = append(c.Counterparties, ce)
	}
}
