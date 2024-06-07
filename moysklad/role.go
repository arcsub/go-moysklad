package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Role Пользовательская роль.
// Ключевое слово: role
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-pol-zowatel-skie-roli
type Role struct {
	ID          *uuid.UUID           `json:"id,omitempty"`          // ID пользовательской роли
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
	ProcessingOrder                 OperationPermission        `json:"processingOrder"`
	CustomerOrder                   OperationPermission        `json:"customerOrder"`
	BonusTransaction                OperationPermission        `json:"bonusTransaction"`
	SalesReturn                     OperationPermission        `json:"salesReturn"`
	PriceList                       OperationPermission        `json:"priceList"`
	Enter                           OperationPermission        `json:"enter"`
	RetailSalesReturn               OperationPermission        `json:"retailSalesReturn"`
	RetailDrawerCashOut             OperationPermission        `json:"retailDrawerCashOut"`
	RetailDrawerCashIn              OperationPermission        `json:"retailDrawerCashIn"`
	RetailDemand                    OperationPermission        `json:"retailDemand"`
	CommissionReportOut             OperationPermission        `json:"commissionReportOut"`
	PrepaymentReturn                OperationPermission        `json:"prepaymentReturn"`
	PurchaseReturn                  OperationPermission        `json:"purchaseReturn"`
	PurchaseOrder                   OperationPermission        `json:"purchaseOrder"`
	Supply                          OperationPermission        `json:"supply"`
	Demand                          OperationPermission        `json:"demand"`
	CommissionReportIn              OperationPermission        `json:"commissionReportIn"`
	Prepayment                      OperationPermission        `json:"prepayment"`
	PaymentOut                      OperationPermission        `json:"paymentOut"`
	PaymentIn                       OperationPermission        `json:"paymentIn"`
	Move                            OperationPermission        `json:"move"`
	Loss                            OperationPermission        `json:"loss"`
	InvoiceOut                      OperationPermission        `json:"invoiceOut"`
	InvoiceIn                       OperationPermission        `json:"invoiceIn"`
	CashOut                         OperationPermission        `json:"cashOut"`
	InternalOrder                   OperationPermission        `json:"internalOrder"`
	CashIn                          OperationPermission        `json:"cashIn"`
	FactureOut                      OperationPermission        `json:"factureOut"`
	FactureIn                       OperationPermission        `json:"factureIn"`
	Inventory                       DictionaryPermission       `json:"inventory"`
	AccountAdjustment               DictionaryPermission       `json:"accountAdjustment"`
	Good                            DictionaryPermission       `json:"good"`
	UtilizationReport               DictionaryPermission       `json:"utilizationReport"`
	CashBoxAdjustment               DictionaryPermission       `json:"cashboxAdjustment"`
	RemainsOrder                    DictionaryPermission       `json:"remainsOrder"`
	RemarkingOrder                  DictionaryPermission       `json:"remarkingOrder"`
	Company                         DictionaryPermission       `json:"company"`
	Contract                        DictionaryPermission       `json:"contract"`
	CounterpartyAdjustment          DictionaryPermission       `json:"counterpartyAdjustment"`
	RetailShift                     DictionaryPermission       `json:"retailShift"`
	CRPTCancellation                DictionaryPermission       `json:"crptCancellation"`
	CRPTPackageCreation             DictionaryPermission       `json:"crptPackageCreation"`
	CRPTPackageDisaggregation       DictionaryPermission       `json:"crptPackageDisaggregation"`
	CRPTPackageItemRemoval          DictionaryPermission       `json:"crptPackageItemRemoval"`
	EnrollOrder                     DictionaryPermission       `json:"enrollOrder"`
	AtkAggregation                  DictionaryPermission       `json:"atkAggregation"`
	RetireOrder                     DictionaryPermission       `json:"retireOrder"`
	Script                          ScriptPermission           `json:"script"`
	EmissionOrder                   DictionaryPermission       `json:"emissionOrder"`
	Currency                        BasePermission             `json:"currency"`
	MyCompany                       BasePermission             `json:"myCompany"`
	Employee                        BasePermission             `json:"employee"`
	Warehouse                       BasePermission             `json:"warehouse"`
	Country                         BasePermission             `json:"country"`
	Uom                             BasePermission             `json:"uom"`
	RetailStore                     BasePermission             `json:"retailStore"`
	Project                         BasePermission             `json:"project"`
	ProcessingPlan                  BasePermission             `json:"processingPlan"`
	CustomEntity                    BasePermission             `json:"customEntity"`
	Processing                      BasePermission             `json:"processing"`
	GTINList                        ViewCreateDeletePermission `json:"GTINList"`
	TrackingCodeList                ViewPrintPermission        `json:"trackingCodeList"`
	ViewCashFlow                    bool                       `json:"viewCashFlow"`
	SendEmail                       bool                       `json:"sendEmail"`
	ViewMoneyDashboard              bool                       `json:"viewMoneyDashboard"`
	ViewDashboard                   bool                       `json:"viewDashboard"`
	ViewCustomerBalanceList         bool                       `json:"viewCustomerBalanceList"`
	ViewCompanyCRM                  bool                       `json:"viewCompanyCRM"`
	ViewCommissionGoods             bool                       `json:"viewCommissionGoods"`
	ViewProfitAndLoss               bool                       `json:"viewProfitAndLoss"`
	ViewPurchaseFunnel              bool                       `json:"viewPurchaseFunnel"`
	ViewRecycleBin                  bool                       `json:"viewRecycleBin"`
	ViewSaleProfit                  bool                       `json:"viewSaleProfit"`
	ViewAudit                       bool                       `json:"viewAudit"`
	SubscriptionControl             bool                       `json:"subscriptionControl"`
	ViewProductCostAndProfit        bool                       `json:"viewProductCostAndProfit"`
	RestoreFromRecycleBin           bool                       `json:"restoreFromRecycleBin"`
	PurchaseControl                 bool                       `json:"purchaseControl"`
	OnlineShops                     bool                       `json:"onlineShops"`
	ListenCalls                     bool                       `json:"listenCalls"`
	ImportData                      bool                       `json:"importData"`
	ExportData                      bool                       `json:"exportData"`
	ViewSerialNumbers               bool                       `json:"viewSerialNumbers"`
	EditDocumentsOfRestrictedPeriod bool                       `json:"editDocumentsOfRestrictedPeriod"`
	EditDocumentTemplates           bool                       `json:"editDocumentTemplates"`
	EditCurrencyRateOfDocument      bool                       `json:"editCurrencyRateOfDocument"`
	ViewStockReport                 bool                       `json:"viewStockReport"`
	ViewTurnover                    bool                       `json:"viewTurnover"`
	ApiRequest                      bool                       `json:"apiRequest"`
	DeleteFromRecycleBin            bool                       `json:"deleteFromRecycleBin"`
}

// AdminRole Роль администратора
type AdminRole struct {
	Meta Meta `json:"meta,omitempty"`
}

func (r AdminRole) String() string {
	return Stringify(r)
}

func (r AdminRole) MetaType() MetaType {
	return MetaTypeSystemRole
}

// IndividualRole Индивидуальная роль
type IndividualRole struct {
	Meta Meta `json:"meta,omitempty"`
}

func (r IndividualRole) String() string {
	return Stringify(r)
}

func (r IndividualRole) MetaType() MetaType {
	return MetaTypeIndividualRole
}

// CashierRole Роль кассира
type CashierRole struct {
	Meta Meta `json:"meta,omitempty"`
}

func (r CashierRole) String() string {
	return Stringify(r)
}

func (r CashierRole) MetaType() MetaType {
	return MetaTypeSystemRole
}

// RoleService
// Сервис для работы с ролями и правами сотрудников.
type RoleService interface {
	GetList(ctx context.Context, params *Params) (*List[Role], *resty.Response, error)
	Create(ctx context.Context, role *Role, params *Params) (*Role, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Role, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, role *Role, params *Params) (*Role, *resty.Response, error)
	GetAdminRole(ctx context.Context) (*AdminRole, *resty.Response, error)
	GetIndividualRole(ctx context.Context) (*IndividualRole, *resty.Response, error)
	GetCashierRole(ctx context.Context) (*CashierRole, *resty.Response, error)
	GetWorkerRole(ctx context.Context) (*CashierRole, *resty.Response, error)
}

type roleService struct {
	Endpoint
	endpointGetList[Role]
	endpointCreate[Role]
	endpointDelete
	endpointGetByID[Role]
	endpointUpdate[Role]
}

func NewRoleService(client *Client) RoleService {
	e := NewEndpoint(client, "entity/role")
	return &roleService{
		Endpoint:        e,
		endpointGetList: endpointGetList[Role]{e},
		endpointCreate:  endpointCreate[Role]{e},
		endpointDelete:  endpointDelete{e},
		endpointGetByID: endpointGetByID[Role]{e},
		endpointUpdate:  endpointUpdate[Role]{e},
	}
}

// GetAdminRole Запрос на получение роли админа.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sotrudnik-zapros-na-poluchenie-roli-admina
func (s *roleService) GetAdminRole(ctx context.Context) (*AdminRole, *resty.Response, error) {
	path := "entity/role/admin"
	return NewRequestBuilder[AdminRole](s.client, path).Get(ctx)
}

// GetIndividualRole Запрос на получение индивидуальной роли.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sotrudnik-zapros-na-poluchenie-indiwidual-noj-roli
func (s *roleService) GetIndividualRole(ctx context.Context) (*IndividualRole, *resty.Response, error) {
	path := "entity/role/individual"
	return NewRequestBuilder[IndividualRole](s.client, path).Get(ctx)
}

// GetCashierRole Запрос на получение роли кассира.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sotrudnik-zapros-na-poluchenie-indiwidual-noj-roli
func (s *roleService) GetCashierRole(ctx context.Context) (*CashierRole, *resty.Response, error) {
	path := "entity/role/cashier"
	return NewRequestBuilder[CashierRole](s.client, path).Get(ctx)
}

// GetWorkerRole Запрос на получение роли кассира.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sotrudnik-zapros-na-poluchenie-roli-sotrudnika-proizwodstwa
func (s *roleService) GetWorkerRole(ctx context.Context) (*CashierRole, *resty.Response, error) {
	path := "entity/role/worker"
	return NewRequestBuilder[CashierRole](s.client, path).Get(ctx)
}
