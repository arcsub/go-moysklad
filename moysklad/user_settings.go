package moysklad

// UserSettings Настройки пользователя.
// Ключевое слово: usersettings
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-nastrojki-pol-zowatelq
type UserSettings struct {
	AutoShowReports             *bool         `json:"autoShowReports,omitempty"`             // Строить ли отчеты автоматически при переходе на вкладку с отчетом
	DefaultCompany              *Meta         `json:"defaultCompany,omitempty"`              // Метаданные Организации, которая будет использоваться по умолчанию в документах
	DefaultCustomerCounterparty *Meta         `json:"defaultCustomerCounterparty,omitempty"` // Метаданные Покупателя, который будет использоваться по умолчанию в документах раздела "Продажи"
	DefaultPlace                *Meta         `json:"defaultPlace,omitempty"`                // Метаданные Склада, который будет использоваться по умолчанию в документах
	DefaultProject              *Meta         `json:"defaultProject,omitempty"`              // Метаданные Проекта, который будет использоваться по умолчанию в документах
	DefaultPurchaseCounterparty *Meta         `json:"defaultPurchaseCounterparty,omitempty"` // Метаданные Поставщика, который будет использоваться по умолчанию в документах раздела "Закупки"
	DefaultScreen               DefaultScreen `json:"defaultScreen,omitempty"`               // Страница, которая открывается у пользователя при логине
	FieldsPerRow                *int          `json:"fieldsPerRow,omitempty"`                // Количество столбцов, в которых будут располагаться дополнительные поля в документах
	Locale                      Locale        `json:"locale,omitempty"`                      // Язык системы. Допустимые значения "ru_RU" и "en_US"
	MailFooter                  *string       `json:"mailFooter,omitempty"`                  // Подставляется в подпись в письмах, отправляемых из МС
	Meta                        *Meta         `json:"meta,omitempty"`                        // Метаданные настроек
	PrintFormat                 PrintFormat   `json:"printFormat,omitempty"`                 // Правила печати документов
}

func (u UserSettings) String() string {
	return Stringify(u)
}

func (u UserSettings) MetaType() MetaType {
	return MetaTypeUserSettings
}

// DefaultScreen Стартовый экран.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-nastrojki-pol-zowatelq-startowyj-akran
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
	DefaultScreenCommissionReportOut       DefaultScreen = "commissionreportout"       // Выданные отчеты комиссионера
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
	DefaultScreenCommissionReport          DefaultScreen = "commissionreport"          // Отчеты комиссионера
	DefaultScreenFiscalEvent               DefaultScreen = "fiscalevent"               // Очередь облачных чеков
	DefaultScreenFiscalQueue               DefaultScreen = "fiscalqueue"               // Очередь облачных чеков
	DefaultScreenRemarkingOrder            DefaultScreen = "remarkingorder"            // Перемаркировка
	DefaultScreenMove                      DefaultScreen = "move"                      // Перемещения
	DefaultScreenFinance                   DefaultScreen = "finance"                   // Платежи
	DefaultScreenPayments                  DefaultScreen = "payments"                  // Подписка
	DefaultScreenDashboard                 DefaultScreen = "dashboard"                 // Показатели
	DefaultScreenCommissionReportIn        DefaultScreen = "commissionreportin"        // Полученные отчеты комиссионера
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
	DefaultScreenFactureIn                 DefaultScreen = "facturein"                 // Счета-фактуры полученные
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
	DefaultScreenHomePage                  DefaultScreen = "homepage"                  // Юр. лица
)

type Locale string

const (
	LocaleRu Locale = "ru_RU"
	LocaleEn Locale = "en_US"
)

type PrintFormat string

const (
	PrintFormatPDF           PrintFormat = "pdf"        // Скачать в формате PDF
	PrintFormatXLS           PrintFormat = "xls"        // Скачать в формате Excel
	PrintFormatODS           PrintFormat = "ods"        // Скачать в формате Open Office Calc
	PrintFormatDefault       PrintFormat = ""           // Предлагать выбор
	PrintFormatOpenInBrowser PrintFormat = "individual" // Открыть в браузере
)
