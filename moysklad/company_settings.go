package moysklad

// CompanySettings Настройки компании.
// Ключевое слово: companysettings
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-nastrojki-kompanii
type CompanySettings struct {
	Meta                     *Meta            `json:"meta,omitempty"`                     // Метаданные
	CheckMinPrice            *bool            `json:"checkMinPrice,omitempty"`            // Автоматически устанавливать минимальную цену. Если включено, при сохранении документов продажи с ценами меньше минимальных цен (указанных в карточках товара) цены будут автоматически увеличены до минимальных.
	CheckShippingStock       *bool            `json:"checkShippingStock,omitempty"`       // Запретить отгрузку отсутствующих товаров. Если запрет установлен (true значение), пользователи не смогут провести отгрузку со склада отсутствующих товаров.
	CompanyAddress           *string          `json:"companyAddress,omitempty"`           // Адрес компании для электронных писем
	Currency                 *Currency        `json:"currency,omitempty"`                 // Метаданные стандартной валюты
	DiscountStrategy         DiscountStrategy `json:"discountStrategy,omitempty"`         // Совместное применение скидок
	GlobalOperationNumbering *bool            `json:"globalOperationNumbering,omitempty"` // Использовать сквозную нумерацию документов. Если проставлен true, будет установлена сквозная нумерация за всю историю, иначе нумерация документов будет начинаться заново каждый календарный год.
	PriceTypes               *PriceTypes      `json:"priceTypes,omitempty"`               // Коллекция всех существующих типов цен
	UseCompanyAddress        *bool            `json:"useCompanyAddress,omitempty"`        // Использовать адрес компании для электронных писем. Если включено, письма будут отправляться с адреса, указанного в companyAddress, иначе письма будут отправляться с адреса пользователя.
	UseRecycleBin            *bool            `json:"useRecycleBin,omitempty"`            // Использовать корзину. Если включено, то все документы при удалении будут помещаться в корзину. Также появится возможность восстанавливать ошибочно удаленные документы.
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
