package moysklad

import (
	"github.com/google/uuid"
)

// Role Пользовательская роль.
// Ключевое слово: role
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-pol-zowatel-skie-roli
type Role struct {
	Id          *uuid.UUID           `json:"id,omitempty"`          // ID пользовательской роли
	Meta        *Meta                `json:"meta,omitempty"`        // Метаданные пользовательской роли
	Name        *string              `json:"name,omitempty"`        // Наименование пользовательской роли
	Permissions *EmployeePermissions `json:"permissions,omitempty"` // Список пермиссий
}

func (r Role) String() string {
	return Stringify(r)
}

func (r Role) MetaType() MetaType {
	return MetaTypeRole
}

type PermissionValue string

const (
	PermissionNo             PermissionValue = "NO"               // Ни на кого
	PermissionOwn            PermissionValue = "OWN"              // Только свои
	PermissionOwnShared      PermissionValue = "OWN_SHARED"       // Свои и общие
	PermissionOwnGroup       PermissionValue = "OWN_GROUP"        // Свои и отдела
	PermissionOwnGroupShared PermissionValue = "OWN_GROUP_SHARED" // Свои, отдела и общие
	PermissionAll            PermissionValue = "ALL"              // Все
)

type ScriptPermissionValue string

const (
	ScriptPermissionValueNo                ScriptPermissionValue = "NO"                 // Нет прав ни на какие задачи
	ScriptPermissionValueAuthor            ScriptPermissionValue = "AUTHOR"             // Созданные пользователем
	ScriptPermissionValueAssignee          ScriptPermissionValue = "ASSIGNEE"           // Назначенные
	ScriptPermissionValueOAuthorOrAssignee ScriptPermissionValue = "AUTHOR_OR_ASSIGNEE" // Созданные пользователем и назначенные ему
	ScriptPermissionValueAll               ScriptPermissionValue = "ALL"                // Возможность совершать действие над любыми задачами
)

type ViewCreateDeletePermission struct {
	View   PermissionValue `json:"view"`
	Create PermissionValue `json:"create"`
	Delete PermissionValue `json:"delete"`
}

type ViewPrintPermission struct {
	View  PermissionValue `json:"view"`
	Print PermissionValue `json:"print"`
}

type BasePermission struct {
	View   PermissionValue `json:"view"`
	Create PermissionValue `json:"create"`
	Update PermissionValue `json:"update"`
	Delete PermissionValue `json:"delete"`
}

type DictionaryPermission struct {
	View   PermissionValue `json:"view"`
	Create PermissionValue `json:"create"`
	Update PermissionValue `json:"update"`
	Delete PermissionValue `json:"delete"`
	Print  PermissionValue `json:"print"`
}

type OperationPermission struct {
	View    PermissionValue `json:"view"`
	Create  PermissionValue `json:"create"`
	Update  PermissionValue `json:"update"`
	Delete  PermissionValue `json:"delete"`
	Approve PermissionValue `json:"approve"`
	Print   PermissionValue `json:"print"`
}

type ScriptPermission struct {
	View   ScriptPermissionValue `json:"view"`
	Create ScriptPermissionValue `json:"create"`
	Update ScriptPermissionValue `json:"update"`
	Delete ScriptPermissionValue `json:"delete"`
	Done   ScriptPermissionValue `json:"done"`
}

type EmployeePermissions struct {
	ApiRequest                      bool                       `json:"apiRequest"`                      // Доступ по АПИ
	DeleteFromRecycleBin            bool                       `json:"deleteFromRecycleBin"`            // Очищать корзину
	EditCurrencyRateOfDocument      bool                       `json:"editCurrencyRateOfDocument"`      // Редактировать курс валюты документа
	EditDocumentTemplates           bool                       `json:"editDocumentTemplates"`           // Редактировать шаблоны документов и отчетов
	EditDocumentsOfRestrictedPeriod bool                       `json:"editDocumentsOfRestrictedPeriod"` // Редактировать документы закрытого периода
	ExportData                      bool                       `json:"exportData"`                      // Экспортировать данные
	ImportData                      bool                       `json:"importData"`                      // Импортировать данные
	ListenCalls                     bool                       `json:"listenCalls"`                     // Прослушивание звонков
	OnlineShops                     bool                       `json:"onlineShops"`                     // Интернет магазины
	PurchaseControl                 bool                       `json:"purchaseControl"`                 // Управление закупками
	RestoreFromRecycleBin           bool                       `json:"restoreFromRecycleBin"`           // Восстанавливать документы
	SendEmail                       bool                       `json:"sendEmail"`                       // Отправлять почту
	SubscriptionControl             bool                       `json:"subscriptionControl"`             // Управление подпиской
	ViewAudit                       bool                       `json:"viewAudit"`                       // Просматривать аудит
	ViewCashFlow                    bool                       `json:"viewCashFlow"`                    // Просматривать движение денежных средств
	ViewCommissionGoods             bool                       `json:"viewCommissionGoods"`             // Просматривать товары на реализации
	ViewCompanyCRM                  bool                       `json:"viewCompanyCRM"`                  // Просматривать показатели
	ViewCustomerBalanceList         bool                       `json:"viewCustomerBalanceList"`         // Просматривать взаиморасчеты
	ViewDashboard                   bool                       `json:"viewDashboard"`                   // Просматривать показатели
	ViewMoneyDashboard              bool                       `json:"viewMoneyDashboard"`              // Видеть остатки денег
	ViewProductCostAndProfit        bool                       `json:"viewProductCostAndProfit"`        // Видеть себестоимость, цену закупки и прибыль товаров
	ViewProfitAndLoss               bool                       `json:"viewProfitAndLoss"`               // Просматривать прибыль и убытки
	ViewPurchaseFunnel              bool                       `json:"viewPurchaseFunnel"`              // Просматривать воронку продаж
	ViewRecycleBin                  bool                       `json:"viewRecycleBin"`                  // Просматривать корзину
	ViewSaleProfit                  bool                       `json:"viewSaleProfit"`                  // Просматривать прибыльность
	ViewSerialNumbers               bool                       `json:"viewSerialNumbers"`               // Просматривать серийные номера
	ViewStockReport                 bool                       `json:"viewStockReport"`                 // Просматривать остатки по товарам
	ViewTurnover                    bool                       `json:"viewTurnover"`                    // Просматривать обороты
	GTINList                        ViewCreateDeletePermission `json:"GTINList"`                        // Список GTIN
	AccountAdjustment               DictionaryPermission       `json:"accountAdjustment"`               // Корректировка остатков на счете
	BonusTransaction                OperationPermission        `json:"bonusTransaction"`                // Бонусные баллы
	CashIn                          OperationPermission        `json:"cashIn"`                          // Приходной ордер
	CashOut                         OperationPermission        `json:"cashOut"`                         // Расходной ордер
	CashBoxAdjustment               DictionaryPermission       `json:"cashboxAdjustment"`               // Корректировка остатков в кассе
	CommissionReportIn              OperationPermission        `json:"commissionReportIn"`              // Полученный отчет комиссионера
	CommissionReportOut             OperationPermission        `json:"commissionReportOut"`             // Выданный отчет комиссионер
	Company                         DictionaryPermission       `json:"company"`                         // Контрагенты
	Contract                        DictionaryPermission       `json:"contract"`                        // Договоры
	CounterpartyAdjustment          DictionaryPermission       `json:"counterpartyAdjustment"`          // Корректировка баланса контрагента
	Country                         BasePermission             `json:"country"`                         // Страны
	CRPTCancellation                DictionaryPermission       `json:"crptCancellation"`                // Списание кодов маркировки
	CRPTPackageCreation             DictionaryPermission       `json:"crptPackageCreation"`             // Формирование упаковки
	CRPTPackageDisaggregation       DictionaryPermission       `json:"crptPackageDisaggregation"`       // Расформирование упаковки
	CRPTPackageItemRemoval          DictionaryPermission       `json:"crptPackageItemRemoval"`          // Изъятие из упаковки
	Currency                        BasePermission             `json:"currency"`                        // Валюты
	CustomEntity                    BasePermission             `json:"customEntity"`                    // Элементы пользовательских справочников
	CustomerOrder                   OperationPermission        `json:"customerOrder"`                   // Заказ покупателям
	Demand                          OperationPermission        `json:"demand"`                          // Отгрузка
	EmissionOrder                   DictionaryPermission       `json:"emissionOrder"`                   // Заказ кодов маркировки
	UtilizationReport               DictionaryPermission       `json:"utilizationReport"`               // Отчет об использовании
	AtkAggregation                  DictionaryPermission       `json:"atkAggregation"`                  // Формирование АТК
	Employee                        BasePermission             `json:"employee"`                        // Сотрудники
	EnrollOrder                     DictionaryPermission       `json:"enrollOrder"`                     // Ввод в оборот кодов маркировки
	Enter                           OperationPermission        `json:"enter"`                           // Оприходование
	FactureIn                       OperationPermission        `json:"factureIn"`                       // Счета-фактуры полученные
	FactureOut                      OperationPermission        `json:"factureOut"`                      // Счета-фактуры выданные
	Good                            DictionaryPermission       `json:"good"`                            // Товары и Услуги
	InternalOrder                   OperationPermission        `json:"internalOrder"`                   // Внутренние заказы
	Inventory                       DictionaryPermission       `json:"inventory"`                       // Инвентаризация
	InvoiceIn                       OperationPermission        `json:"invoiceIn"`                       // Счет поставщика
	InvoiceOut                      OperationPermission        `json:"invoiceOut"`                      // Счет покупателю
	Loss                            OperationPermission        `json:"loss"`                            // Списание
	Move                            OperationPermission        `json:"move"`                            // Перемещение
	MyCompany                       BasePermission             `json:"myCompany"`                       // Юр. Лица
	PaymentIn                       OperationPermission        `json:"paymentIn"`                       // Входящий платеж
	PaymentOut                      OperationPermission        `json:"paymentOut"`                      // Исходящий платеж
	Prepayment                      OperationPermission        `json:"prepayment"`                      // Предоплаты
	PrepaymentReturn                OperationPermission        `json:"prepaymentReturn"`                // Возврат предоплаты
	PriceList                       OperationPermission        `json:"priceList"`                       // Прайс-лист
	Processing                      BasePermission             `json:"processing"`                      // Тех. операции
	ProcessingOrder                 OperationPermission        `json:"processingOrder"`                 // Заказ на производство
	ProcessingPlan                  BasePermission             `json:"processingPlan"`                  // Тех. Карты
	Project                         BasePermission             `json:"project"`                         // Проекты
	PurchaseOrder                   OperationPermission        `json:"purchaseOrder"`                   // Заказ поставщикам
	PurchaseReturn                  OperationPermission        `json:"purchaseReturn"`                  // Возврат поставщику
	RemainsOrder                    DictionaryPermission       `json:"remainsOrder"`                    // Описание остатков
	RemarkingOrder                  DictionaryPermission       `json:"remarkingOrder"`                  // Перемаркировка
	RetailDemand                    OperationPermission        `json:"retailDemand"`                    // Продажи
	RetailDrawerCashIn              OperationPermission        `json:"retailDrawerCashIn"`              // Внесения
	RetailDrawerCashOut             OperationPermission        `json:"retailDrawerCashOut"`             // Выплаты
	RetailSalesReturn               OperationPermission        `json:"retailSalesReturn"`               // Возвраты
	RetailShift                     DictionaryPermission       `json:"retailShift"`                     // Смены
	RetailStore                     BasePermission             `json:"retailStore"`                     // Точка продаж
	RetireOrder                     DictionaryPermission       `json:"retireOrder"`                     // Возврат в оборот
	SalesReturn                     OperationPermission        `json:"salesReturn"`                     // Возврат покупателя
	Supply                          OperationPermission        `json:"supply"`                          // Приемки
	TrackingCodeList                ViewPrintPermission        `json:"trackingCodeList"`                // Коды маркировки
	Uom                             BasePermission             `json:"uom"`                             // Единицы измерения
	Warehouse                       BasePermission             `json:"warehouse"`                       // Склады
	Script                          ScriptPermission           `json:"script"`
}
