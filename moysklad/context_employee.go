package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
)

// ContextEmployee Контекст запроса сотрудника.
//
// Код сущности: employee
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-kontext-zaprosa-sotrudnika-poluchit-kontext-sotrudnika
type ContextEmployee struct {
	Image        Image              `json:"image,omitempty"`        // Фотография сотрудника
	Meta         Meta               `json:"meta,omitempty"`         // Метаданные Сотрудника
	Group        Group              `json:"group,omitempty"`        // Отдел сотрудника
	Created      Timestamp          `json:"created,omitempty"`      // Момент создания Сотрудника
	Updated      Timestamp          `json:"updated,omitempty"`      // Момент последнего обновления Сотрудника
	Salary       Salary             `json:"salary,omitempty"`       // Оклад сотрудника
	Permissions  Permissions        `json:"permissions,omitempty"`  // Перечисление пермиссий сотрудника
	Position     string             `json:"position,omitempty"`     // Должность сотрудника
	Code         string             `json:"code,omitempty"`         // Код Сотрудника
	FirstName    string             `json:"firstName,omitempty"`    // Имя
	FullName     string             `json:"fullName,omitempty"`     // Имя Отчество Фамилия
	Email        string             `json:"email,omitempty"`        // Электронная почта сотрудника
	Description  string             `json:"description,omitempty"`  // Комментарий к Сотруднику
	ExternalCode string             `json:"externalCode,omitempty"` // Внешний код Сотрудника
	UID          string             `json:"uid,omitempty"`          // Логин Сотрудника
	LastName     string             `json:"lastName,omitempty"`     // Фамилия
	ShortFio     string             `json:"shortFio,omitempty"`     // Краткое ФИО
	MiddleName   string             `json:"middleName,omitempty"`   // Отчество
	Name         string             `json:"name,omitempty"`         // Наименование Сотрудника
	INN          string             `json:"inn,omitempty"`          // ИНН сотрудника (в формате ИНН физического лица)
	Phone        string             `json:"phone,omitempty"`        // Телефон сотрудника
	Owner        Employee           `json:"owner,omitempty"`        // Метаданные владельца (Сотрудника)
	Attributes   Slice[Attribute]   `json:"attributes,omitempty"`   // Дополнительные поля Сотрудника
	Cashiers     MetaArray[Cashier] `json:"cashiers,omitempty"`     // Массив кассиров
	AccountID    string             `json:"accountId,omitempty"`    // ID учётной записи
	ID           string             `json:"id,omitempty"`           // ID Сотрудника
	Shared       bool               `json:"shared,omitempty"`       // Общий доступ
	Archived     bool               `json:"archived,omitempty"`     // Добавлен ли Сотрудник в архив
}

// GetMeta возвращает Метаданные Сотрудника.
func (contextEmployee ContextEmployee) GetMeta() Meta {
	return contextEmployee.Meta
}

// String реализует интерфейс [fmt.Stringer].
func (contextEmployee ContextEmployee) String() string {
	return Stringify(contextEmployee)
}

// MetaType возвращает код сущности.
func (ContextEmployee) MetaType() MetaType { return MetaTypeEmployeeContext }

// Permissions пермиссии сотрудника.
type Permissions struct {
	Currency                  BasePermission       `json:"currency"`                  // Валюты
	Uom                       BasePermission       `json:"uom"`                       // Единицы измерения
	ProductFolder             DictionaryPermission `json:"productfolder"`             // Группы товаров
	Product                   DictionaryPermission `json:"product"`                   // Товары
	Bundle                    DictionaryPermission `json:"bundle"`                    // Комплекты
	Service                   DictionaryPermission `json:"service"`                   // Услуги
	Consignment               DictionaryPermission `json:"consignment"`               // Серии
	Variant                   DictionaryPermission `json:"variant"`                   // Модификации
	Store                     BasePermission       `json:"store"`                     // Склады
	Counterparty              DictionaryPermission `json:"counterparty"`              // Контрагенты
	Organization              BasePermission       `json:"organization"`              // Юрлица
	Employee                  BasePermission       `json:"employee"`                  // Сотрудники
	Settings                  DictionaryPermission `json:"settings"`                  // Настройки
	Contract                  DictionaryPermission `json:"contract"`                  // Договоры
	Project                   BasePermission       `json:"project"`                   // Проекты
	SalesChannel              DictionaryPermission `json:"saleschannel"`              // Каналы продаж
	Country                   BasePermission       `json:"country"`                   // Страны
	CustomEntity              BasePermission       `json:"customentity"`              // Элементы пользовательских справочников
	Demand                    OperationPermission  `json:"demand"`                    // Отгрузки
	CustomerOrder             OperationPermission  `json:"customerorder"`             // Заказы покупателей
	InternalOrder             OperationPermission  `json:"internalorder"`             // Внутренние заказы
	InvoiceOut                OperationPermission  `json:"invoiceout"`                // Счет покупателям
	InvoiceIn                 OperationPermission  `json:"invoicein"`                 // Счета поставщиков
	PaymentIn                 OperationPermission  `json:"paymentin"`                 // Входящие платежи
	PaymentOut                OperationPermission  `json:"paymentout"`                // Исходящие платежи
	CashIn                    OperationPermission  `json:"cashin"`                    // Приходной ордер
	CashOut                   OperationPermission  `json:"cashout"`                   // Расходной ордер
	Supply                    OperationPermission  `json:"supply"`                    // Приемки
	SalesReturn               OperationPermission  `json:"salesreturn"`               // Возвраты покупателей
	PurchaseReturn            OperationPermission  `json:"purchasereturn"`            // Возвраты поставщикам
	RetailStore               BasePermission       `json:"retailstore"`               // Точки продаж
	ReceiptTemplate           BasePermission       `json:"receipttemplate"`           // Шаблоны
	RetailStoreStatus         BasePermission       `json:"retailstorestatus"`         // Статусы точек продаж
	RetailShift               DictionaryPermission `json:"retailshift"`               // Смены
	RetailDemand              OperationPermission  `json:"retaildemand"`              // Продажи
	RetailSalesReturn         OperationPermission  `json:"retailsalesreturn"`         // Возвраты
	RetailDrawerCashIn        OperationPermission  `json:"retaildrawercashin"`        // Внесения
	RetailDrawerCashOut       OperationPermission  `json:"retaildrawercashout"`       // Выплаты
	Prepayment                OperationPermission  `json:"prepayment"`                // Предоплаты
	PrepaymentReturn          OperationPermission  `json:"prepaymentreturn"`          // Возвраты предоплат
	PurchaseOrder             OperationPermission  `json:"purchaseorder"`             // Заказы поставщикам
	Move                      OperationPermission  `json:"move"`                      // Перемещения
	Enter                     OperationPermission  `json:"enter"`                     // Оприходования
	Loss                      OperationPermission  `json:"loss"`                      // Списания
	FactureIn                 OperationPermission  `json:"facturein"`                 // полученные счета-фактуры
	FactureOut                OperationPermission  `json:"factureout"`                // Счета-фактуры выданные
	CommissionReportIn        OperationPermission  `json:"commissionreportin"`        // Полученный отчет комиссионера
	CommissionReportOut       OperationPermission  `json:"commissionreportout"`       // Выданный отчет комиссионер
	PriceList                 OperationPermission  `json:"pricelist"`                 // Прайс-листы
	ProcessingPlanFolder      BasePermission       `json:"processingplanfolder"`      // Группы техкарт
	ProcessingPlan            BasePermission       `json:"processingplan"`            // Техкарты
	ProcessingStage           BasePermission       `json:"processingstage"`           // Этапы производства
	Processing                OperationPermission  `json:"processing"`                // Техоперации
	ProcessingOrder           OperationPermission  `json:"processingorder"`           // Заказы на производство
	ProcessingProcess         OperationPermission  `json:"processingprocess"`         // Техпроцессы
	CounterpartyAdjustment    DictionaryPermission `json:"counterpartyadjustment"`    // Корректировка взаиморасчетов
	Assortment                DictionaryPermission `json:"assortment"`                // Товары и Услуги
	Inventory                 DictionaryPermission `json:"inventory"`                 // Инвентаризации
	BonusTransaction          OperationPermission  `json:"bonustransaction"`          // Бонусные баллы
	CRPTOrder                 DictionaryPermission `json:"crptorder"`                 // Заказ маркировок
	ProductionTask            OperationPermission  `json:"productiontask"`            // Производственные задания
	ProductionStageCompletion DictionaryPermission `json:"productionstagecompletion"` // Выполнения этапов производства
	Payroll                   OperationPermission  `json:"payroll"`                   // Зарплата
	TaxRate                   BasePermission       `json:"taxrate"`                   // Ставки НДС
	Webhook                   BasePermission       `json:"webhook"`                   // Вебхуки
	Task                      ScriptPermission     `json:"task"`                      // Задачи
	Dashboard                 ViewPermission       `json:"dashboard"`                 // Просматривать показатели
	Stock                     ViewPermission       `json:"stock"`                     // Просматривать остатки по товарам
	CustomAttributes          ViewPermission       `json:"customAttributes"`          // Работа с доп. полями
	Profit                    ViewPermission       `json:"pnl"`                       // Просматривать прибыльность
	CompanyCrm                ViewPermission       `json:"company_crm"`               // Просматривать показатели
	TariffCrm                 ViewPermission       `json:"tariff_crm"`                // Присутствует ли опция CRM на аккаунте
	AuditDashboard            ViewPermission       `json:"audit_dashboard"`           // Просматривать аудит
	Admin                     ViewPermission       `json:"admin"`                     // Является ли сотрудник админом
	DashboardMoney            ViewPermission       `json:"dashboardMoney"`            // Видеть остатки денег
	ViewCashFlow              ViewPermission       `json:"viewCashFlow"`              // Видеть движение денежных средств
}

// ContextEmployeeService описывает методы сервиса для работы с контекстом сотрудника.
type ContextEmployeeService interface {
	// Get выполняет запрос на получение контекста запроса сотрудника.
	// Принимает контекст.
	// Возвращает контекст запроса сотрудника.
	Get(ctx context.Context) (*ContextEmployee, *resty.Response, error)
}

const (
	EndpointContextEmployee = EndpointContext + string(MetaTypeEmployee)
)

type contextEmployeeService struct {
	Endpoint
}

func (service *contextEmployeeService) Get(ctx context.Context) (*ContextEmployee, *resty.Response, error) {
	return NewRequestBuilder[ContextEmployee](service.client, service.uri).Get(ctx)
}

// NewContextEmployeeService принимает [Client] и возвращает сервис для работы с контекстом сотрудника.
func NewContextEmployeeService(client *Client) ContextEmployeeService {
	return &contextEmployeeService{NewEndpoint(client, EndpointContextEmployee)}
}
