package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// UserSettings Настройки пользователя.
//
// Код сущности: usersettings
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-nastrojki-pol-zowatelq
type UserSettings struct {
	AutoShowReports             *bool         `json:"autoShowReports,omitempty"`             // Строить ли отчёты автоматически при переходе на вкладку с отчётом
	DefaultCompany              *Meta         `json:"defaultCompany,omitempty"`              // Метаданные Организации, которая будет использоваться по умолчанию в документах
	DefaultCustomerCounterparty *Meta         `json:"defaultCustomerCounterparty,omitempty"` // Метаданные Покупателя, который будет использоваться по умолчанию в документах раздела "Продажи"
	DefaultPlace                *Meta         `json:"defaultPlace,omitempty"`                // Метаданные Склада, который будет использоваться по умолчанию в документах
	DefaultProject              *Meta         `json:"defaultProject,omitempty"`              // Метаданные Проекта, который будет использоваться по умолчанию в документах
	DefaultPurchaseCounterparty *Meta         `json:"defaultPurchaseCounterparty,omitempty"` // Метаданные Поставщика, который будет использоваться по умолчанию в документах раздела "Закупки"
	DefaultScreen               DefaultScreen `json:"defaultScreen,omitempty"`               // Страница, которая открывается у пользователя при логине
	FieldsPerRow                *int          `json:"fieldsPerRow,omitempty"`                // Количество столбцов, в которых будут располагаться дополнительные поля в документах
	Locale                      Locale        `json:"locale,omitempty"`                      // Язык системы
	MailFooter                  *string       `json:"mailFooter,omitempty"`                  // Подставляется в подпись в письмах, отправляемых из МС
	Meta                        *Meta         `json:"meta,omitempty"`                        // Метаданные настроек
	PrintFormat                 PrintFormat   `json:"printFormat,omitempty"`                 // Правила печати документов
}

// GetAutoShowReports возвращает флаг построения отчётов автоматически при переходе на вкладку с отчётом.
func (userSettings UserSettings) GetAutoShowReports() bool {
	return Deref(userSettings.AutoShowReports)
}

// GetDefaultCompany возвращает Метаданные Организации, которая будет использоваться по умолчанию в документах.
func (userSettings UserSettings) GetDefaultCompany() Meta {
	return Deref(userSettings.DefaultCompany)
}

// GetDefaultCustomerCounterparty возвращает Метаданные Покупателя, который будет использоваться по умолчанию в документах раздела "Продажи".
func (userSettings UserSettings) GetDefaultCustomerCounterparty() Meta {
	return Deref(userSettings.DefaultCustomerCounterparty)
}

// GetDefaultPlace возвращает Метаданные Склада, который будет использоваться по умолчанию в документах.
func (userSettings UserSettings) GetDefaultPlace() Meta {
	return Deref(userSettings.DefaultPlace)
}

// GetDefaultProject возвращает Метаданные Проекта, который будет использоваться по умолчанию в документах.
func (userSettings UserSettings) GetDefaultProject() Meta {
	return Deref(userSettings.DefaultProject)
}

// GetDefaultPurchaseCounterparty возвращает Метаданные Поставщика, который будет использоваться по умолчанию в документах раздела "Закупки".
func (userSettings UserSettings) GetDefaultPurchaseCounterparty() Meta {
	return Deref(userSettings.DefaultPurchaseCounterparty)
}

// GetDefaultScreen возвращает Страницу, которая открывается у пользователя при логине.
func (userSettings UserSettings) GetDefaultScreen() DefaultScreen {
	return userSettings.DefaultScreen
}

// GetFieldsPerRow возвращает Количество столбцов, в которых будут располагаться дополнительные поля в документах.
func (userSettings UserSettings) GetFieldsPerRow() int {
	return Deref(userSettings.FieldsPerRow)
}

// GetLocale возвращает Язык системы.
func (userSettings UserSettings) GetLocale() Locale {
	return userSettings.Locale
}

// GetMailFooter возвращает подпись к письмам, отправляемых из МС.
func (userSettings UserSettings) GetMailFooter() string {
	return Deref(userSettings.MailFooter)
}

// GetMeta возвращает Метаданные настроек.
func (userSettings UserSettings) GetMeta() Meta {
	return Deref(userSettings.Meta)
}

// GetPrintFormat возвращает Правила печати документов.
func (userSettings UserSettings) GetPrintFormat() PrintFormat {
	return userSettings.PrintFormat
}

// SetAutoShowReports устанавливает флаг построения отчётов автоматически при переходе на вкладку с отчётом.
func (userSettings *UserSettings) SetAutoShowReports(autoShowReports bool) *UserSettings {
	userSettings.AutoShowReports = &autoShowReports
	return userSettings
}

// SetDefaultCompany устанавливает Метаданные Организации, которая будет использоваться по умолчанию в документах.
func (userSettings *UserSettings) SetDefaultCompany(defaultCompany *Organization) *UserSettings {
	if defaultCompany != nil {
		userSettings.DefaultCompany = defaultCompany.Meta
	}
	return userSettings
}

// SetDefaultCustomerCounterparty устанавливает Метаданные Покупателя, который будет использоваться по умолчанию в документах раздела "Продажи".
//
// Принимает [Counterparty] или [Organization].
func (userSettings *UserSettings) SetDefaultCustomerCounterparty(defaultCustomerCounterparty AgentOrganizationConverter) *UserSettings {
	if defaultCustomerCounterparty != nil {
		userSettings.DefaultCustomerCounterparty = defaultCustomerCounterparty.AsOrganizationAgent().Meta
	}
	return userSettings
}

// SetDefaultPlace устанавливает Метаданные Склада, который будет использоваться по умолчанию в документах.
func (userSettings *UserSettings) SetDefaultPlace(defaultPlace *Store) *UserSettings {
	if defaultPlace != nil {
		userSettings.DefaultPlace = defaultPlace.Meta
	}
	return userSettings
}

// SetDefaultProject устанавливает Метаданные Проекта, который будет использоваться по умолчанию в документа.
func (userSettings *UserSettings) SetDefaultProject(defaultProject *Project) *UserSettings {
	if defaultProject != nil {
		userSettings.DefaultProject = defaultProject.Meta
	}
	return userSettings
}

// SetDefaultPurchaseCounterparty устанавливает Метаданные Поставщика, который будет использоваться по умолчанию в документах раздела "Закупки".
//
// Принимает [Counterparty] или [Organization].
func (userSettings *UserSettings) SetDefaultPurchaseCounterparty(defaultPurchaseCounterparty AgentOrganizationConverter) *UserSettings {
	if defaultPurchaseCounterparty != nil {
		userSettings.DefaultPurchaseCounterparty = defaultPurchaseCounterparty.AsOrganizationAgent().Meta
	}
	return userSettings
}

// SetDefaultScreen устанавливает Страницу, которая открывается у пользователя при логине.
func (userSettings *UserSettings) SetDefaultScreen(defaultScreen DefaultScreen) *UserSettings {
	userSettings.DefaultScreen = defaultScreen
	return userSettings
}

// SetFieldsPerRow устанавливает Количество столбцов, в которых будут располагаться дополнительные поля в документах.
func (userSettings *UserSettings) SetFieldsPerRow(fieldsPerRow int) *UserSettings {
	userSettings.FieldsPerRow = &fieldsPerRow
	return userSettings
}

// SetLocale устанавливает Язык системы.
func (userSettings *UserSettings) SetLocale(locale Locale) *UserSettings {
	userSettings.Locale = locale
	return userSettings
}

// SetLocaleRU устанавливает Язык системы в значение [LocaleRU].
func (userSettings *UserSettings) SetLocaleRU() *UserSettings {
	userSettings.Locale = LocaleRU
	return userSettings
}

// SetLocaleEN устанавливает Язык системы в значение [LocaleEN].
func (userSettings *UserSettings) SetLocaleEN() *UserSettings {
	userSettings.Locale = LocaleEN
	return userSettings
}

// SetMailFooter устанавливает подпись к письмам, отправляемых из МС.
func (userSettings *UserSettings) SetMailFooter(mailFooter string) *UserSettings {
	userSettings.MailFooter = &mailFooter
	return userSettings
}

// SetPrintFormat устанавливает Правила печати документов.
func (userSettings *UserSettings) SetPrintFormat(printFormat PrintFormat) *UserSettings {
	userSettings.PrintFormat = printFormat
	return userSettings
}

// SetPrintFormatPDF устанавливает Правила печати документов в значение [PrintFormatPDF].
func (userSettings *UserSettings) SetPrintFormatPDF() *UserSettings {
	userSettings.PrintFormat = PrintFormatPDF
	return userSettings
}

// SetPrintFormatXLS устанавливает Правила печати документов в значение [PrintFormatXLS].
func (userSettings *UserSettings) SetPrintFormatXLS() *UserSettings {
	userSettings.PrintFormat = PrintFormatXLS
	return userSettings
}

// SetPrintFormatODS устанавливает Правила печати документов в значение [PrintFormatODS].
func (userSettings *UserSettings) SetPrintFormatODS() *UserSettings {
	userSettings.PrintFormat = PrintFormatODS
	return userSettings
}

// SetPrintFormatDefault устанавливает Правила печати документов в значение [PrintFormatDefault].
func (userSettings *UserSettings) SetPrintFormatDefault() *UserSettings {
	userSettings.PrintFormat = PrintFormatDefault
	return userSettings
}

// SetPrintFormatOpenInBrowser устанавливает Правила печати документов в значение [PrintFormatOpenInBrowser].
func (userSettings *UserSettings) SetPrintFormatOpenInBrowser() *UserSettings {
	userSettings.PrintFormat = PrintFormatOpenInBrowser
	return userSettings
}

// String реализует интерфейс [fmt.Stringer].
func (userSettings UserSettings) String() string {
	return Stringify(userSettings)
}

// MetaType возвращает код сущности.
func (UserSettings) MetaType() MetaType {
	return MetaTypeUserSettings
}

// DefaultScreen Стартовый экран.
//
// Возможные значения:
//   - DefaultScreenAudit                     – Аудит
//   - DefaultScreenCurrency                  – Валюты
//   - DefaultScreenEnrollOrder               – Ввод в оборот кодов маркировки
//   - DefaultScreenCustomersBalanceList      – Взаиморасчеты
//   - DefaultScreenRetailDrawerCashIn        – Внесения
//   - DefaultScreenInternalOrder             – Внутренние заказы
//   - DefaultScreenEnrollReturn              – Возврат в оборот
//   - DefaultScreenRetailSalesReturn         – Возвраты
//   - DefaultScreenSalesReturn               – Возвраты покупателей
//   - DefaultScreenPurchaseReturn            – Возвраты поставщикам
//   - DefaultScreenPrepaymentReturn          – Возвраты предоплат
//   - DefaultScreenPurchaseFunnel            – Воронка продаж
//   - DefaultScreenRetireOrder               – Вывод из оборота
//   - DefaultScreenCommissionReportOut       – Выданные отчёты комиссионера
//   - DefaultScreenRetailDrawerCashOut       – Выплаты
//   - DefaultScreenCashFlow                  – Движение денежных средств
//   - DefaultScreenContract                  – Договоры
//   - DefaultScreenOperation                 – Документы
//   - DefaultScreenUom                       – Единицы измерения
//   - DefaultScreenCRPTLog                   – Журнал запросов в ИС МП
//   - DefaultScreenLoyaltyLog                – Журнал запросов в систему лояльности
//   - DefaultScreenPurpose                   – Задачи
//   - DefaultScreenCRPTDemand                – Заказ кодов маркировки
//   - DefaultScreenProcessingOrder           – Заказы на производство
//   - DefaultScreenCustomerOrder             – Заказы покупателей
//   - DefaultScreenPurchaseOrder             – Заказы поставщикам
//   - DefaultScreenEvotorRequest             – Запросы
//   - DefaultScreenPhoneCall                 – Звонки
//   - DefaultScreenCRPTPackageItemRemoval    – Изъятие из упаковки
//   - DefaultScreenImport                    – Импорт
//   - DefaultScreenImportGoods               – Импорт из Excel
//   - DefaultScreenImportEdo                 – Импорт приемки
//   - DefaultScreenImportCustom              – Импорт справочника
//   - DefaultScreenInventory                 – Инвентаризации
//   - DefaultScreenCompany                   – Контрагенты
//   - DefaultScreenRecycleBin                – Корзина
//   - DefaultScreenAdjustment                – Корректировки
//   - DefaultScreenBulkEdit                  – Массовое редактирование
//   - DefaultScreenEvotorMapping             – Настройка обмена с Эвотор
//   - DefaultScreenCompanySettings           – Настройки
//   - DefaultScreenFeed                      – Новости
//   - DefaultScreenTurnover                  – Обороты
//   - DefaultScreenBonusTransaction          – Операции с баллами
//   - DefaultScreenRemainsOrder              – Описание остатков
//   - DefaultScreenEnter                     – Оприходования
//   - DefaultScreenStockReport               – Остатки
//   - DefaultScreenDemand                    – Отгрузки
//   - DefaultScreenCommissionReport          – отчёты комиссионера
//   - DefaultScreenFiscalEvent               – Очередь облачных чеков
//   - DefaultScreenFiscalQueue               – Очередь облачных чеков
//   - DefaultScreenRemarkingOrder            – Перемаркировка
//   - DefaultScreenMove                      – Перемещения
//   - DefaultScreenFinance                   – Платежи
//   - DefaultScreenPayments                  – Подписка
//   - DefaultScreenDashboard                 – Показатели
//   - DefaultScreenCommissionReportIn        – Полученные отчёты комиссионера
//   - DefaultScreenPriceList                 – Прайс-листы
//   - DefaultScreenPrepayment                – Предоплаты
//   - DefaultScreenPnl3                      – Прибыли и убытки
//   - DefaultScreenPnl                       – Прибыльность
//   - DefaultScreenSupply                    – Приемки
//   - DefaultScreenApps                      – Приложения
//   - DefaultScreenEmbedApps                 – Приложения
//   - DefaultScreenCheckEquipment            – Проверка комплектации
//   - DefaultScreenRetailDemand              – Продажи
//   - DefaultScreenProject                   – Проекты
//   - DefaultScreenTrackingIdentify          – Просмотр информации о КМ или ТУ
//   - DefaultScreenCRPTPackageDisaggregation – Расформирование упаковки
//   - DefaultScreenOrderAssembly             – Сбор заказа
//   - DefaultScreenSerialNumbers             – Сер. номера
//   - DefaultScreenConnectorSettings         – Синхронизация
//   - DefaultScreenDiscount                  – Скидки
//   - DefaultScreenWarehouse                 – Склады
//   - DefaultScreenRetailShift               – Смены
//   - DefaultScreenEvotorEvent               – События обмена с Эвотор
//   - DefaultScreenEmployee                  – Сотрудники
//   - DefaultScreenSpecialOffers             – Спецпредложения
//   - DefaultScreenCRPTCancellation          – Списание кодов маркировки
//   - DefaultScreenLoss                      – Списания
//   - DefaultScreenCountry                   – Страны
//   - DefaultScreenScriptTemplate            – Сценарии
//   - DefaultScreenInvoiceOut                – Счета покупателям
//   - DefaultScreenInvoiceIn                 – Счета поставщиков
//   - DefaultScreenFactureOut                – Счета-фактуры выданные
//   - DefaultScreenFactureIn                 – полученные счета-фактуры
//   - DefaultScreenProcessingPlan            – Тех. карты
//   - DefaultScreenProcessing                – Тех. операции
//   - DefaultScreenGood                      – Товары и услуги
//   - DefaultScreenCommissionGoods           – Товары на реализации
//   - DefaultScreenRetailStore               – Точки продаж
//   - DefaultScreenNotifications             – Уведомления
//   - DefaultScreenPurchaseControl           – Управление закупками
//   - DefaultScreenAccount                   – Учетная запись
//   - DefaultScreenCRPTPackageCreation       – Формирование упаковки
//   - DefaultScreenFeature                   – Характеристика
//   - DefaultScreenExport                    – Экспорт
//   - DefaultScreenMyCompany                 – Юр. лица
//   - DefaultScreenHomePage                  – Начало работы
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-nastrojki-pol-zowatelq-startowyj-akran
type DefaultScreen string

const (
	DefaultScreenAudit                     DefaultScreen = "audit"                     // Аудит
	DefaultScreenCurrency                  DefaultScreen = "currency"                  // Валюты
	DefaultScreenEnrollOrder               DefaultScreen = "enrollorder"               // Ввод в оборот кодов маркировки
	DefaultScreenCustomersBalanceList      DefaultScreen = "customersbalancelist"      // Взаиморасчеты
	DefaultScreenRetailDrawerCashIn        DefaultScreen = "retaildrawercashin"        // Внесения
	DefaultScreenInternalOrder             DefaultScreen = "internalorder"             // Внутренние заказы
	DefaultScreenEnrollReturn              DefaultScreen = "enrollreturn"              // Возврат в оборот
	DefaultScreenRetailSalesReturn         DefaultScreen = "retailsalesreturn"         // Возвраты
	DefaultScreenSalesReturn               DefaultScreen = "salesreturn"               // Возвраты покупателей
	DefaultScreenPurchaseReturn            DefaultScreen = "purchasereturn"            // Возвраты поставщикам
	DefaultScreenPrepaymentReturn          DefaultScreen = "prepaymentreturn"          // Возвраты предоплат
	DefaultScreenPurchaseFunnel            DefaultScreen = "purchasefunnel"            // Воронка продаж
	DefaultScreenRetireOrder               DefaultScreen = "retireorder"               // Вывод из оборота
	DefaultScreenCommissionReportOut       DefaultScreen = "commissionreportout"       // Выданные отчёты комиссионера
	DefaultScreenRetailDrawerCashOut       DefaultScreen = "retaildrawercashout"       // Выплаты
	DefaultScreenCashFlow                  DefaultScreen = "cashflow"                  // Движение денежных средств
	DefaultScreenContract                  DefaultScreen = "contract"                  // Договоры
	DefaultScreenOperation                 DefaultScreen = "operation"                 // Документы
	DefaultScreenUom                       DefaultScreen = "uom"                       // Единицы измерения
	DefaultScreenCRPTLog                   DefaultScreen = "crptlog"                   // Журнал запросов в ИС МП
	DefaultScreenLoyaltyLog                DefaultScreen = "loyaltylog"                // Журнал запросов в систему лояльности
	DefaultScreenPurpose                   DefaultScreen = "purpose"                   // Задачи
	DefaultScreenCRPTDemand                DefaultScreen = "crptdemand"                // Заказ кодов маркировки
	DefaultScreenProcessingOrder           DefaultScreen = "processingorder"           // Заказы на производство
	DefaultScreenCustomerOrder             DefaultScreen = "customerorder"             // Заказы покупателей
	DefaultScreenPurchaseOrder             DefaultScreen = "purchaseorder"             // Заказы поставщикам
	DefaultScreenEvotorRequest             DefaultScreen = "evotorrequest"             // Запросы
	DefaultScreenPhoneCall                 DefaultScreen = "phonecall"                 // Звонки
	DefaultScreenCRPTPackageItemRemoval    DefaultScreen = "crptpackageitemremoval"    // Изъятие из упаковки
	DefaultScreenImport                    DefaultScreen = "import"                    // Импорт
	DefaultScreenImportGoods               DefaultScreen = "importgoods"               // Импорт из Excel
	DefaultScreenImportEdo                 DefaultScreen = "importedo"                 // Импорт приемки
	DefaultScreenImportCustom              DefaultScreen = "importcustom"              // Импорт справочника
	DefaultScreenInventory                 DefaultScreen = "inventory"                 // Инвентаризации
	DefaultScreenCompany                   DefaultScreen = "company"                   // Контрагенты
	DefaultScreenRecycleBin                DefaultScreen = "recyclebin"                // Корзина
	DefaultScreenAdjustment                DefaultScreen = "adjustment"                // Корректировки
	DefaultScreenBulkEdit                  DefaultScreen = "bulkEdit"                  // Массовое редактирование
	DefaultScreenEvotorMapping             DefaultScreen = "evotormapping"             // Настройка обмена с Эвотор
	DefaultScreenCompanySettings           DefaultScreen = "companysettings"           // Настройки
	DefaultScreenFeed                      DefaultScreen = "feed"                      // Новости
	DefaultScreenTurnover                  DefaultScreen = "turnover"                  // Обороты
	DefaultScreenBonusTransaction          DefaultScreen = "bonustransaction"          // Операции с баллами
	DefaultScreenRemainsOrder              DefaultScreen = "remainsorder"              // Описание остатков
	DefaultScreenEnter                     DefaultScreen = "enter"                     // Оприходования
	DefaultScreenStockReport               DefaultScreen = "stockreport"               // Остатки
	DefaultScreenDemand                    DefaultScreen = "demand"                    // Отгрузки
	DefaultScreenCommissionReport          DefaultScreen = "commissionreport"          // отчёты комиссионера
	DefaultScreenFiscalEvent               DefaultScreen = "fiscalevent"               // Очередь облачных чеков
	DefaultScreenFiscalQueue               DefaultScreen = "fiscalqueue"               // Очередь облачных чеков
	DefaultScreenRemarkingOrder            DefaultScreen = "remarkingorder"            // Перемаркировка
	DefaultScreenMove                      DefaultScreen = "move"                      // Перемещения
	DefaultScreenFinance                   DefaultScreen = "finance"                   // Платежи
	DefaultScreenPayments                  DefaultScreen = "payments"                  // Подписка
	DefaultScreenDashboard                 DefaultScreen = "dashboard"                 // Показатели
	DefaultScreenCommissionReportIn        DefaultScreen = "commissionreportin"        // Полученные отчёты комиссионера
	DefaultScreenPriceList                 DefaultScreen = "pricelist"                 // Прайс-листы
	DefaultScreenPrepayment                DefaultScreen = "prepayment"                // Предоплаты
	DefaultScreenPnl3                      DefaultScreen = "pnl3"                      // Прибыли и убытки
	DefaultScreenPnl                       DefaultScreen = "pnl"                       // Прибыльность
	DefaultScreenSupply                    DefaultScreen = "supply"                    // Приемки
	DefaultScreenApps                      DefaultScreen = "apps"                      // Приложения
	DefaultScreenEmbedApps                 DefaultScreen = "embed-apps"                // Приложения
	DefaultScreenCheckEquipment            DefaultScreen = "checkequipment"            // Проверка комплектации
	DefaultScreenRetailDemand              DefaultScreen = "retaildemand"              // Продажи
	DefaultScreenProject                   DefaultScreen = "project"                   // Проекты
	DefaultScreenTrackingIdentify          DefaultScreen = "trackingidentify"          // Просмотр информации о КМ или ТУ
	DefaultScreenCRPTPackageDisaggregation DefaultScreen = "crptpackagedisaggregation" // Расформирование упаковки
	DefaultScreenOrderAssembly             DefaultScreen = "orderassembly"             // Сбор заказа
	DefaultScreenSerialNumbers             DefaultScreen = "serialnumbers"             // Сер. номера
	DefaultScreenConnectorSettings         DefaultScreen = "connectorsettings"         // Синхронизация
	DefaultScreenDiscount                  DefaultScreen = "discount"                  // Скидки
	DefaultScreenWarehouse                 DefaultScreen = "warehouse"                 // Склады
	DefaultScreenRetailShift               DefaultScreen = "retailshift"               // Смены
	DefaultScreenEvotorEvent               DefaultScreen = "evotorevent"               // События обмена с Эвотор
	DefaultScreenEmployee                  DefaultScreen = "employee"                  // Сотрудники
	DefaultScreenSpecialOffers             DefaultScreen = "specialoffers"             // Спецпредложения
	DefaultScreenCRPTCancellation          DefaultScreen = "crptcancellation"          // Списание кодов маркировки
	DefaultScreenLoss                      DefaultScreen = "loss"                      // Списания
	DefaultScreenCountry                   DefaultScreen = "country"                   // Страны
	DefaultScreenScriptTemplate            DefaultScreen = "scripttemplate"            // Сценарии
	DefaultScreenInvoiceOut                DefaultScreen = "invoiceout"                // Счета покупателям
	DefaultScreenInvoiceIn                 DefaultScreen = "invoicein"                 // Счета поставщиков
	DefaultScreenFactureOut                DefaultScreen = "factureout"                // Счета-фактуры выданные
	DefaultScreenFactureIn                 DefaultScreen = "facturein"                 // полученные счета-фактуры
	DefaultScreenProcessingPlan            DefaultScreen = "processingplan"            // Тех. карты
	DefaultScreenProcessing                DefaultScreen = "processing"                // Тех. операции
	DefaultScreenGood                      DefaultScreen = "good"                      // Товары и услуги
	DefaultScreenCommissionGoods           DefaultScreen = "commissiongoods"           // Товары на реализации
	DefaultScreenRetailStore               DefaultScreen = "retailstore"               // Точки продаж
	DefaultScreenNotifications             DefaultScreen = "notifications"             // Уведомления
	DefaultScreenPurchaseControl           DefaultScreen = "purchasecontrol"           // Управление закупками
	DefaultScreenAccount                   DefaultScreen = "account"                   // Учетная запись
	DefaultScreenCRPTPackageCreation       DefaultScreen = "crptpackagecreation"       // Формирование упаковки
	DefaultScreenFeature                   DefaultScreen = "feature"                   // Характеристика
	DefaultScreenExport                    DefaultScreen = "export"                    // Экспорт
	DefaultScreenMyCompany                 DefaultScreen = "mycompany"                 // Юр. лица
	DefaultScreenHomePage                  DefaultScreen = "homepage"                  // Начало работы
)

// Locale Язык системы.
//
// Возможные значения:
//   - LocaleRU – русский
//   - LocaleEN – английский
type Locale string

const (
	LocaleRU Locale = "ru_RU" // русский
	LocaleEN Locale = "en_US" // английский
)

// PrintFormat Правила печати документов.
//
// Возможные значения:
//   - PrintFormatPDF           – Скачать в формате pdf
//   - PrintFormatXLS           – Скачать в формате excel
//   - PrintFormatODS           – Скачать в формате Open Office Calc
//   - PrintFormatDefault       – Предлагать выбор
//   - PrintFormatOpenInBrowser – Открыть в браузере
type PrintFormat string

const (
	PrintFormatPDF           PrintFormat = "pdf"        // Скачать в формате pdf
	PrintFormatXLS           PrintFormat = "xls"        // Скачать в формате excel
	PrintFormatODS           PrintFormat = "ods"        // Скачать в формате Open Office Calc
	PrintFormatDefault       PrintFormat = ""           // Предлагать выбор
	PrintFormatOpenInBrowser PrintFormat = "individual" // Открыть в браузере
)

// UserSettingsService описывает методы сервиса для работы с настройками пользователей.
type UserSettingsService interface {
	// Get выполняет запрос на получение настроек пользователя.
	// Принимает контекст.
	// Возвращает настройки пользователя.
	Get(ctx context.Context) (*UserSettings, *resty.Response, error)

	// Update выполняет запрос на изменение настроек пользователя.
	// Принимает контекст и настройки пользователя.
	// Возвращает изменённые настройки пользователя.
	Update(ctx context.Context, id uuid.UUID, userSettings *UserSettings) (*UserSettings, *resty.Response, error)
}

const (
	EndpointUserSettings = EndpointContext + string(MetaTypeUserSettings)
)

type userSettingsService struct {
	Endpoint
}

func (service *userSettingsService) Get(ctx context.Context) (*UserSettings, *resty.Response, error) {
	return NewRequestBuilder[UserSettings](service.client, service.uri).Get(ctx)
}

func (service *userSettingsService) Update(ctx context.Context, id uuid.UUID, userSettings *UserSettings) (*UserSettings, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s", service.uri, id)
	return NewRequestBuilder[UserSettings](service.client, path).Put(ctx, userSettings)
}

// NewContextUserSettingsService принимает [Client] и возвращает сервис для работы с настройками пользователей.
func NewContextUserSettingsService(client *Client) UserSettingsService {
	return &userSettingsService{NewEndpoint(client, EndpointUserSettings)}
}
