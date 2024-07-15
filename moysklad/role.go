package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Role Пользовательская роль.
//
// Код сущности: role
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-pol-zowatel-skie-roli
type Role struct {
	ID          *uuid.UUID           `json:"id,omitempty"`          // ID пользовательской роли
	Meta        *Meta                `json:"meta,omitempty"`        // Метаданные пользовательской роли
	Name        *string              `json:"name,omitempty"`        // Наименование пользовательской роли
	Permissions *EmployeePermissions `json:"permissions,omitempty"` // Список пермиссий
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (role Role) Clean() *Role {
	if role.Meta == nil {
		return nil
	}
	return &Role{Meta: role.Meta}
}

// GetID возвращает ID пользовательской роли.
func (role Role) GetID() uuid.UUID {
	return Deref(role.ID)
}

// GetMeta возвращает Метаданные пользовательской роли.
func (role Role) GetMeta() Meta {
	return Deref(role.Meta)
}

// GetName возвращает Наименование пользовательской роли.
func (role Role) GetName() string {
	return Deref(role.Name)
}

// GetPermissions возвращает Список пермиссий.
func (role Role) GetPermissions() EmployeePermissions {
	return Deref(role.Permissions)
}

// SetMeta устанавливает Метаданные пользовательской роли.
func (role *Role) SetMeta(meta *Meta) *Role {
	role.Meta = meta
	return role
}

// SetName устанавливает Наименование пользовательской роли.
func (role *Role) SetName(name string) *Role {
	role.Name = &name
	return role
}

// SetPermissions устанавливает Список пермиссий.
func (role *Role) SetPermissions(permissions *EmployeePermissions) *Role {
	role.Permissions = permissions
	return role
}

// String реализует интерфейс [fmt.Stringer].
func (role Role) String() string { return Stringify(role) }

// MetaType возвращает код сущности.
func (Role) MetaType() MetaType { return MetaTypeRole }

// Update shortcut
func (role *Role) Update(ctx context.Context, client *Client, params ...*Params) (*Role, *resty.Response, error) {
	return NewRoleService(client).Update(ctx, role.GetID(), role, params...)
}

// Create shortcut
func (role *Role) Create(ctx context.Context, client *Client, params ...*Params) (*Role, *resty.Response, error) {
	return NewRoleService(client).Create(ctx, role, params...)
}

// Delete shortcut
func (role *Role) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewRoleService(client).Delete(ctx, role)
}

// AdminRole Роль администратора.
type AdminRole struct {
	Meta Meta `json:"meta,omitempty"` // метаданные роли
}

// String реализует интерфейс [fmt.Stringer].
func (adminRole AdminRole) String() string {
	return Stringify(adminRole)
}

// MetaType возвращает код сущности.
func (AdminRole) MetaType() MetaType {
	return MetaTypeSystemRole
}

// IndividualRole Индивидуальная роль.
type IndividualRole struct {
	Meta Meta `json:"meta,omitempty"` // метаданные роли
}

// String реализует интерфейс [fmt.Stringer].
func (individualRole IndividualRole) String() string {
	return Stringify(individualRole)
}

// MetaType возвращает код сущности.
func (IndividualRole) MetaType() MetaType {
	return MetaTypeIndividualRole
}

// CashierRole Роль кассира.
type CashierRole struct {
	Meta Meta `json:"meta,omitempty"` // метаданные роли
}

// String реализует интерфейс [fmt.Stringer].
func (cashierRole CashierRole) String() string {
	return Stringify(cashierRole)
}

// MetaType возвращает код сущности.
func (CashierRole) MetaType() MetaType {
	return MetaTypeSystemRole
}

// WorkerRole Роль сотрудника производства.
type WorkerRole struct {
	Meta Meta `json:"meta,omitempty"` // метаданные роли
}

// String реализует интерфейс [fmt.Stringer].
func (workerRole WorkerRole) String() string {
	return Stringify(workerRole)
}

// MetaType возвращает код сущности.
func (WorkerRole) MetaType() MetaType {
	return MetaTypeSystemRole
}

type PermissionValue string

const (
	PermissionOwn            PermissionValue = "OWN"              // Только свои
	PermissionOwnShared      PermissionValue = "OWN_SHARED"       // Свои и общие
	PermissionOwnGroup       PermissionValue = "OWN_GROUP"        // Свои и отдела
	PermissionOwnGroupShared PermissionValue = "OWN_GROUP_SHARED" // Свои, отдела и общие
	PermissionAll            PermissionValue = "ALL"              // Все
	PermissionNo             PermissionValue = "NO"               // Ни на кого
	PermissionNone           PermissionValue = ""                 // Отсутствие прав
)

type ScriptPermissionValue string

const (
	ScriptPermissionValueOAuthorOrAssignee ScriptPermissionValue = "AUTHOR_OR_ASSIGNEE" // Созданные пользователем и назначенные ему
	ScriptPermissionValueAssignee          ScriptPermissionValue = "ASSIGNEE"           // Назначенные
	ScriptPermissionValueAuthor            ScriptPermissionValue = "AUTHOR"             // Созданные пользователем
	ScriptPermissionValueAll               ScriptPermissionValue = "ALL"                // Возможность совершать действие над любыми задачами
	ScriptPermissionValueNo                ScriptPermissionValue = "NO"                 // Нет прав ни на какие задачи
	ScriptPermissionValueNone              PermissionValue       = ""                   // Отсутствие прав
)

type ViewPermission struct {
	View PermissionValue `json:"view"` // Смотреть
}

type ViewCreateDeletePermission struct {
	View   PermissionValue `json:"view"`   // Смотреть
	Create PermissionValue `json:"create"` // Создавать
	Delete PermissionValue `json:"delete"` // Удалять
}

type ViewPrintPermission struct {
	View  PermissionValue `json:"view"`  // Смотреть
	Print PermissionValue `json:"print"` // Печатать
}

type BasePermission struct {
	View   PermissionValue `json:"view"`   // Смотреть
	Create PermissionValue `json:"create"` // Создавать
	Update PermissionValue `json:"update"` // Редактировать
	Delete PermissionValue `json:"delete"` // Удалять
}

type DictionaryPermission struct {
	View   PermissionValue `json:"view"`   // Смотреть
	Create PermissionValue `json:"create"` // Создавать
	Update PermissionValue `json:"update"` // Редактировать
	Delete PermissionValue `json:"delete"` // Удалять
	Print  PermissionValue `json:"print"`  // Печатать
}

type OperationPermission struct {
	View    PermissionValue `json:"view"`    // Смотреть
	Create  PermissionValue `json:"create"`  // Создавать
	Update  PermissionValue `json:"update"`  // Редактировать
	Delete  PermissionValue `json:"delete"`  // Удалять
	Approve PermissionValue `json:"approve"` // Проводить
	Print   PermissionValue `json:"print"`   // Печатать
}

type ScriptPermission struct {
	View   ScriptPermissionValue `json:"view"`   // Смотреть
	Create ScriptPermissionValue `json:"create"` // Создавать
	Update ScriptPermissionValue `json:"update"` // Редактировать
	Delete ScriptPermissionValue `json:"delete"` // Удалять
	Done   ScriptPermissionValue `json:"done"`   // Выполнять
}

type EmployeePermissions struct {
	ProcessingOrder                 OperationPermission        `json:"processingOrder"`                 // Заказ на производство
	CustomerOrder                   OperationPermission        `json:"customerOrder"`                   // Заказ покупателя
	BonusTransaction                OperationPermission        `json:"bonusTransaction"`                // Бонусные баллы
	SalesReturn                     OperationPermission        `json:"salesReturn"`                     // Возврат покупателя
	PriceList                       OperationPermission        `json:"priceList"`                       // Прайс-лист
	Enter                           OperationPermission        `json:"enter"`                           // Оприходование
	RetailSalesReturn               OperationPermission        `json:"retailSalesReturn"`               // Возвраты
	RetailDrawerCashOut             OperationPermission        `json:"retailDrawerCashOut"`             // Выплаты
	RetailDrawerCashIn              OperationPermission        `json:"retailDrawerCashIn"`              // Внесения
	RetailDemand                    OperationPermission        `json:"retailDemand"`                    // Продажи
	CommissionReportOut             OperationPermission        `json:"commissionReportOut"`             // Выданный отчет комиссионера
	PrepaymentReturn                OperationPermission        `json:"prepaymentReturn"`                // Возврат предоплаты
	PurchaseReturn                  OperationPermission        `json:"purchaseReturn"`                  // Возврат поставщику
	PurchaseOrder                   OperationPermission        `json:"purchaseOrder"`                   // Заказ поставщику
	Supply                          OperationPermission        `json:"supply"`                          // Приемки
	Demand                          OperationPermission        `json:"demand"`                          // Отгрузка
	CommissionReportIn              OperationPermission        `json:"commissionReportIn"`              // Полученный отчет комиссионера
	Prepayment                      OperationPermission        `json:"prepayment"`                      // Предоплаты
	PaymentOut                      OperationPermission        `json:"paymentOut"`                      // Исходящий платеж
	PaymentIn                       OperationPermission        `json:"paymentIn"`                       // Входящий платеж
	Move                            OperationPermission        `json:"move"`                            // Перемещение
	Loss                            OperationPermission        `json:"loss"`                            // Списание
	InvoiceOut                      OperationPermission        `json:"invoiceOut"`                      // Счет покупателю
	InvoiceIn                       OperationPermission        `json:"invoiceIn"`                       // Счет поставщика
	CashOut                         OperationPermission        `json:"cashOut"`                         // Расходной ордер
	InternalOrder                   OperationPermission        `json:"internalOrder"`                   // Внутренние заказы
	CashIn                          OperationPermission        `json:"cashIn"`                          // Приходной ордер
	FactureOut                      OperationPermission        `json:"factureOut"`                      // Счета-фактуры выданные
	FactureIn                       OperationPermission        `json:"factureIn"`                       // полученные счета-фактуры
	ProductionTask                  OperationPermission        `json:"productionTask"`                  // Производственные задания
	Inventory                       DictionaryPermission       `json:"inventory"`                       // Инвентаризация
	AccountAdjustment               DictionaryPermission       `json:"accountAdjustment"`               // Корректировка остатков на счете
	Good                            DictionaryPermission       `json:"good"`                            // Товары и Услуги
	UtilizationReport               DictionaryPermission       `json:"utilizationReport"`               // Отчет об использовании
	CashBoxAdjustment               DictionaryPermission       `json:"cashboxAdjustment"`               // Корректировка остатков в кассе
	RemainsOrder                    DictionaryPermission       `json:"remainsOrder"`                    // Описание остатков
	RemarkingOrder                  DictionaryPermission       `json:"remarkingOrder"`                  // Перемаркировка
	Company                         DictionaryPermission       `json:"company"`                         // Контрагенты
	Contract                        DictionaryPermission       `json:"contract"`                        // Договоры
	CounterpartyAdjustment          DictionaryPermission       `json:"counterpartyAdjustment"`          // Корректировка взаиморасчетов
	RetailShift                     DictionaryPermission       `json:"retailShift"`                     // Смены
	CRPTCancellation                DictionaryPermission       `json:"crptCancellation"`                // Списание кодов маркировки
	CRPTPackageCreation             DictionaryPermission       `json:"crptPackageCreation"`             // Формирование упаковки
	CRPTPackageDisaggregation       DictionaryPermission       `json:"crptPackageDisaggregation"`       // Расформирование упаковки
	CRPTPackageItemRemoval          DictionaryPermission       `json:"crptPackageItemRemoval"`          // Изъятие из упаковки
	EnrollOrder                     DictionaryPermission       `json:"enrollOrder"`                     // Ввод в оборот кодов маркировки
	AtkAggregation                  DictionaryPermission       `json:"atkAggregation"`                  // Формирование АТК
	RetireOrderOSU                  DictionaryPermission       `json:"retireOrderOSU"`                  // Вывод из оборота ОСУ
	RetireOrder                     DictionaryPermission       `json:"retireOrder"`                     // Вывод из оборота
	ProductionStageCompletion       DictionaryPermission       `json:"productionStageCompletion"`       // Выполнение этапов
	Script                          ScriptPermission           `json:"script"`                          // ?
	EmissionOrder                   DictionaryPermission       `json:"emissionOrder"`                   // Заказ кодов маркировки
	Currency                        BasePermission             `json:"currency"`                        // Валюты
	MyCompany                       BasePermission             `json:"myCompany"`                       // Юр. Лица
	Employee                        BasePermission             `json:"employee"`                        // Сотрудники
	Warehouse                       BasePermission             `json:"warehouse"`                       // Склады
	Country                         BasePermission             `json:"country"`                         // Страны
	Uom                             BasePermission             `json:"uom"`                             // Единицы измерения
	RetailStore                     BasePermission             `json:"retailStore"`                     // Точка продаж
	Project                         BasePermission             `json:"project"`                         // Проекты
	ProcessingPlan                  BasePermission             `json:"processingPlan"`                  // Техкарты
	CustomEntity                    BasePermission             `json:"customEntity"`                    // Элементы пользовательских справочников
	Processing                      BasePermission             `json:"processing"`                      // Техоперации
	ProcessingStage                 BasePermission             `json:"processingStage"`                 // Этапы производства
	ProcessingProcess               BasePermission             `json:"processingProcess"`               // Техпроцессы
	TaxRate                         BasePermission             `json:"taxrate"`                         // Ставки НДС
	GTINList                        ViewCreateDeletePermission `json:"GTINList"`                        // Список GTIN
	TrackingCodeList                ViewPrintPermission        `json:"trackingCodeList"`                // Коды маркировки
	ViewCashFlow                    bool                       `json:"viewCashFlow"`                    // Просматривать движение денежных средств
	SendEmail                       bool                       `json:"sendEmail"`                       // Отправлять почту
	ViewMoneyDashboard              bool                       `json:"viewMoneyDashboard"`              // Видеть остатки денег
	ViewDashboard                   bool                       `json:"viewDashboard"`                   // Просматривать показатели
	ViewCustomerBalanceList         bool                       `json:"viewCustomerBalanceList"`         // Просматривать взаиморасчеты
	ViewCompanyCRM                  bool                       `json:"viewCompanyCRM"`                  // Просматривать показатели
	ViewCommissionGoods             bool                       `json:"viewCommissionGoods"`             // Просматривать товары на реализации
	ViewProfitAndLoss               bool                       `json:"viewProfitAndLoss"`               // Просматривать прибыль и убытки
	ViewPurchaseFunnel              bool                       `json:"viewPurchaseFunnel"`              // Просматривать воронку продаж
	ViewRecycleBin                  bool                       `json:"viewRecycleBin"`                  // Просматривать корзину
	ViewSaleProfit                  bool                       `json:"viewSaleProfit"`                  // Просматривать прибыльность
	ViewAudit                       bool                       `json:"viewAudit"`                       // Просматривать аудит
	SubscriptionControl             bool                       `json:"subscriptionControl"`             // Управление подпиской
	ViewProductCostAndProfit        bool                       `json:"viewProductCostAndProfit"`        // Видеть себестоимость, цену закупки и прибыль товаров
	RestoreFromRecycleBin           bool                       `json:"restoreFromRecycleBin"`           // Восстанавливать документы
	PurchaseControl                 bool                       `json:"purchaseControl"`                 // Управление закупками
	OnlineShops                     bool                       `json:"onlineShops"`                     // Интернет магазины
	ListenCalls                     bool                       `json:"listenCalls"`                     // Прослушивание звонков
	ImportData                      bool                       `json:"importData"`                      // Импортировать данные
	ExportData                      bool                       `json:"exportData"`                      // Экспортировать данные
	ViewSerialNumbers               bool                       `json:"viewSerialNumbers"`               // Просматривать серийные номера
	EditDocumentsOfRestrictedPeriod bool                       `json:"editDocumentsOfRestrictedPeriod"` // Редактировать документы закрытого периода
	EditDocumentTemplates           bool                       `json:"editDocumentTemplates"`           // Редактировать шаблоны документов и отчетов
	EditCurrencyRateOfDocument      bool                       `json:"editCurrencyRateOfDocument"`      // Редактировать курс валюты документа
	ViewStockReport                 bool                       `json:"viewStockReport"`                 // Просматривать остатки по товарам
	ViewTurnover                    bool                       `json:"viewTurnover"`                    // Просматривать обороты
	ApiRequest                      bool                       `json:"apiRequest"`                      // Доступ по АПИ
	DeleteFromRecycleBin            bool                       `json:"deleteFromRecycleBin"`            // Очищать корзину
}

// RoleService описывает методы сервиса для работы с ролями и правами сотрудников.
type RoleService interface {
	// GetList выполняет запрос на получение списка пользовательских ролей.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[Role], *resty.Response, error)

	// Create выполняет запрос на создание пользовательской роли.
	// Обязательные поля для заполнения:
	//	- name (Наименование пользовательской роли)
	//	- permissions (Список пермиссий)
	// Принимает контекст, пользовательскую роль и опционально объект параметров запроса Params.
	// Возвращает созданную пользовательскую роль.
	Create(ctx context.Context, role *Role, params ...*Params) (*Role, *resty.Response, error)

	// DeleteByID выполняет запрос на массовое удаление пользовательской роли по ID.
	// Принимает контекст и ID пользовательской роли.
	// Возвращает «true» в случае успешного удаления пользовательской роли.
	DeleteByID(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// Delete выполняет запрос на удаление пользовательской роли.
	// Принимает контекст и пользовательскую роль.
	// Возвращает «true» в случае успешного удаления пользовательской роли.
	Delete(ctx context.Context, entity *Role) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельной пользовательской роли по ID.
	// Принимает контекст, ID пользовательской роли и опционально объект параметров запроса Params.
	// Возвращает пользовательскую роль.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Role, *resty.Response, error)

	// Update выполняет запрос на изменение пользовательской роли.
	// Принимает контекст, пользовательскую роль и опционально объект параметров запроса Params.
	// Возвращает изменённую пользовательскую роль.
	Update(ctx context.Context, id uuid.UUID, role *Role, params ...*Params) (*Role, *resty.Response, error)

	// GetAdminRole выполняет запрос на получение роли администратора.
	// Принимает контекст.
	// Возвращает роль администратора.
	GetAdminRole(ctx context.Context) (*AdminRole, *resty.Response, error)

	// GetIndividualRole выполняет запрос на получение индивидуальной роли.
	// Принимает контекст.
	// Возвращает индивидуальную роль.
	GetIndividualRole(ctx context.Context) (*IndividualRole, *resty.Response, error)

	// GetCashierRole выполняет запрос на получение роли кассира.
	// Принимает контекст.
	// Возвращает роль кассира.
	GetCashierRole(ctx context.Context) (*CashierRole, *resty.Response, error)

	// GetWorkerRole выполняет запрос на получение роли сотрудника производства.
	// Принимает контекст.
	// Возвращает роль сотрудника производства.
	GetWorkerRole(ctx context.Context) (*WorkerRole, *resty.Response, error)
}

const (
	EndpointRole           = EndpointEntity + string(MetaTypeRole)
	EndpointRoleAdmin      = EndpointRole + "/admin"
	EndpointRoleIndividual = EndpointRole + "/individual"
	EndpointRoleCashier    = EndpointRole + "/cashier"
	EndpointRoleWorker     = EndpointRole + "/worker"
)

type roleService struct {
	Endpoint
	endpointGetList[Role]
	endpointCreate[Role]
	endpointDeleteByID
	endpointDelete[Role]
	endpointGetByID[Role]
	endpointUpdate[Role]
}

func (service *roleService) GetAdminRole(ctx context.Context) (*AdminRole, *resty.Response, error) {
	return NewRequestBuilder[AdminRole](service.client, EndpointRoleAdmin).Get(ctx)
}

func (service *roleService) GetIndividualRole(ctx context.Context) (*IndividualRole, *resty.Response, error) {
	return NewRequestBuilder[IndividualRole](service.client, EndpointRoleIndividual).Get(ctx)
}

func (service *roleService) GetCashierRole(ctx context.Context) (*CashierRole, *resty.Response, error) {
	return NewRequestBuilder[CashierRole](service.client, EndpointRoleCashier).Get(ctx)
}

func (service *roleService) GetWorkerRole(ctx context.Context) (*WorkerRole, *resty.Response, error) {
	return NewRequestBuilder[WorkerRole](service.client, EndpointRoleWorker).Get(ctx)
}

// NewRoleService принимает [Client] и возвращает сервис для работы с ролями и правами сотрудников.
func NewRoleService(client *Client) RoleService {
	e := NewEndpoint(client, EndpointRole)
	return &roleService{
		Endpoint:           e,
		endpointGetList:    endpointGetList[Role]{e},
		endpointCreate:     endpointCreate[Role]{e},
		endpointDeleteByID: endpointDeleteByID{e},
		endpointDelete:     endpointDelete[Role]{e},
		endpointGetByID:    endpointGetByID[Role]{e},
		endpointUpdate:     endpointUpdate[Role]{e},
	}
}
