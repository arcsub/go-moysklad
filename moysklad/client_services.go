package moysklad

// EntityService описывает методы сервиса для работы с сущностями и документами.
type EntityService interface {
	// Application возвращает сервис для работы с серверными приложениями.
	Application() ApplicationService

	// Assortment возвращает сервис для работы с ассортиментом.
	Assortment() AssortmentService

	// BonusProgram возвращает сервис для работы с бонусными программами.
	BonusProgram() BonusProgramService

	// BonusTransaction возвращает сервис для работы с бонусными операциями.
	BonusTransaction() BonusTransactionService

	// Bundle возвращает сервис для работы с комплектами.
	Bundle() BundleService

	// CashIn возвращает сервис для работы с приходными ордерами.
	CashIn() CashInService

	// CashOut возвращает сервис для работы с расходными ордерами.
	CashOut() CashOutService

	// CommissionReportIn возвращает сервис для работы с полученными отчетами комиссионера.
	CommissionReportIn() CommissionReportInService

	// CommissionReportOut возвращает сервис для работы с выданными отчетами комиссионера.
	CommissionReportOut() CommissionReportOutService

	// Consignment возвращает сервис для работы с сериями.
	Consignment() ConsignmentService

	// Contract возвращает сервис для работы с договорами.
	Contract() ContractService

	// Counterparty возвращает сервис для работы с контрагентами.
	Counterparty() CounterpartyService

	// CounterPartyAdjustment возвращает сервис для работы с корректировками взаиморасчетов.
	CounterPartyAdjustment() CounterPartyAdjustmentService

	// Country возвращает сервис для работы со странами.
	Country() CountryService

	// Currency возвращает сервис для работы с валютами.
	Currency() CurrencyService

	// CustomEntity возвращает сервис для работы с пользовательскими справочниками.
	CustomEntity() CustomEntityService

	// CustomerOrder возвращает сервис для работы с заказами покупателей.
	CustomerOrder() CustomerOrderService

	// Demand возвращает сервис для работы с отгрузками.
	Demand() DemandService

	// Discount возвращает сервис для работы со скидками.
	Discount() DiscountService

	// Employee возвращает сервис для работы с сотрудниками.
	Employee() EmployeeService

	// Enter возвращает сервис для работы с оприходованиями.
	Enter() EnterService

	// ExpenseItem возвращает сервис для работы со статьями расходов.
	ExpenseItem() ExpenseItemService

	// FactureIn возвращает сервис для работы со счетами-фактурами полученными.
	FactureIn() FactureInService

	// FactureOut возвращает сервис для работы со счетами-фактурами выданными.
	FactureOut() FactureOutService

	// Group возвращает сервис для работы с отделами.
	Group() GroupService

	// InternalOrder возвращает сервис для работы с внутренними заказами.
	InternalOrder() InternalOrderService

	// Inventory возвращает сервис для работы с инвентаризациями.
	Inventory() InventoryService

	// InvoiceIn возвращает сервис для работы со счетами поставщиков.
	InvoiceIn() InvoiceInService

	// InvoiceOut возвращает сервис для работы со счетами покупателям.
	InvoiceOut() InvoiceOutService

	// Loss возвращает сервис для работы со списаниями.
	Loss() LossService

	// Metadata возвращает сервис для работы с глобальными метаданными.
	Metadata() MetadataService

	// Move возвращает сервис для работы с перемещениями.
	Move() MoveService

	// Organization возвращает сервис для работы с юр лицами.
	Organization() OrganizationService

	// PaymentIn возвращает сервис для работы со входящими платежами.
	PaymentIn() PaymentInService

	// PaymentOut возвращает сервис для работы с исходящими платежами.
	PaymentOut() PaymentOutService

	// Prepayment возвращает сервис для работы с предоплатами.
	Prepayment() PrepaymentService

	// PrepaymentReturn возвращает сервис для работы с возвратами предоплат.
	PrepaymentReturn() PrepaymentReturnService

	// PriceList возвращает сервис для работы с прайс-листами.
	PriceList() PriceListService

	// Processing возвращает сервис для работы с техоперациями.
	Processing() ProcessingService

	// ProcessingOrder возвращает сервис для работы с заказами на производство.
	ProcessingOrder() ProcessingOrderService

	// ProcessingPlan возвращает сервис для работы с техкартами.
	ProcessingPlan() ProcessingPlanService

	// ProcessingPlanFolder возвращает сервис для работы с группами техкарт.
	ProcessingPlanFolder() ProcessingPlanFolderService

	// ProcessingProcess возвращает сервис для работы с техпроцессами.
	ProcessingProcess() ProcessingProcessService

	// ProcessingStage возвращает сервис для работы с этапами производства.
	ProcessingStage() ProcessingStageService

	// Product возвращает сервис для работы с товарами.
	Product() ProductService

	// ProductFolder возвращает сервис для работы с группами товаров.
	ProductFolder() ProductFolderService

	// ProductionTask возвращает сервис для работы с производственными заданиями.
	ProductionTask() ProductionTaskService

	// ProductionStage возвращает сервис для работы с производственными этапами.
	ProductionStage() ProductionStageService

	// ProductionStageCompletion возвращает сервис для работы с выполнениями этапов производства.
	ProductionStageCompletion() ProductionStageCompletionService

	// Project возвращает сервис для работы с проектами.
	Project() ProjectService

	// PurchaseOrder возвращает сервис для работы с заказами поставщикам.
	PurchaseOrder() PurchaseOrderService

	// PurchaseReturn возвращает сервис для работы с возвратами поставщикам.
	PurchaseReturn() PurchaseReturnService

	// Region возвращает сервис для работы с регионами.
	Region() RegionService

	// RetailDemand возвращает сервис для работы с розничными продажами.
	RetailDemand() RetailDemandService

	// RetailDrawerCashIn возвращает сервис для работы с внесениями.
	RetailDrawerCashIn() RetailDrawerCashInService

	// RetailDrawerCashOut возвращает сервис для работы с выплатами.
	RetailDrawerCashOut() RetailDrawerCashOutService

	// RetailSalesReturn возвращает сервис для работы с розничными возвратами.
	RetailSalesReturn() RetailSalesReturnService

	// RetailShift возвращает сервис для работы с розничными сменами.
	RetailShift() RetailShiftService

	// RetailStore возвращает сервис для работы с точками продаж.
	RetailStore() RetailStoreService

	// Role возвращает сервис для работы с ролями.
	Role() RoleService

	// SalesChannel возвращает сервис для работы с каналами продаж.
	SalesChannel() SalesChannelService

	// SalesReturn возвращает сервис для работы с возвратами покупателей.
	SalesReturn() SalesReturnService

	// Service возвращает сервис для работы с услугами.
	Service() ServiceService

	// Store возвращает сервис для работы со складами.
	Store() StoreService

	// Subscription возвращает сервис для работы с подпиской.
	Subscription() SubscriptionService

	// Supply возвращает сервис для работы с приёмками.
	Supply() SupplyService

	// Task возвращает сервис для работы с задачами.
	Task() TaskService

	// TaxRate возвращает сервис для работы со ставками НДС.
	TaxRate() TaxRateService

	// Thing возвращает сервис для работы с серийными номерами.
	Thing() ThingService

	// Uom возвращает сервис для работы с единицами измерения.
	Uom() UomService

	// Variant возвращает сервис для работы с модификациями.
	Variant() VariantService

	// Webhook возвращает сервис для работы с вебхуками.
	Webhook() WebhookService

	// WebhookStock возвращает сервис для работы с вебхуками на изменение остатков.
	WebhookStock() WebhookStockService
}

// ContextService описывает методы сервиса для работы с контекстом.
type ContextService interface {
	// CompanySettings возвращает сервис для работы с настройками компании.
	CompanySettings() ContextCompanySettingsService

	// Employee возвращает сервис для работы с контекстом сотрудника.
	Employee() ContextEmployeeService

	// UserSettings возвращает сервис для работы с настройками пользователя.
	UserSettings() UserSettingsService
}

// ReportService описывает методы сервиса для работы с отчётами.
type ReportService interface {
	// Counterparty возвращает сервис для работы с отчётом Показатели контрагентов.
	Counterparty() ReportCounterpartyService

	// Dashboard возвращает сервис для работы с отчётом Показатели.
	Dashboard() ReportDashboardService

	// Money возвращает сервис для работы с отчётом Деньги.
	Money() ReportMoneyService

	// Profit возвращает сервис для работы с отчётом Прибыльность.
	Profit() ReportProfitService

	// Sales возвращает сервис для работы с отчётом Показатели продаж.
	Sales() ReportSalesService

	// Orders возвращает сервис для работы с отчётом Показатели заказов.
	Orders() ReportOrdersService

	// Stock возвращает сервис для работы с отчётом Остатки.
	Stock() ReportStockService

	// Turnover возвращает сервис для работы с отчётом Обороты.
	Turnover() ReportTurnoverService
}

// Async возвращает сервис для работы с асинхронными задачами.
func (client *Client) Async() AsyncService {
	return NewAsyncService(client)
}

// Audit возвращает сервис для работы с аудитом.
func (client *Client) Audit() AuditService {
	return NewAuditService(client)
}

// Context возвращает сервис для работы с контекстом.
func (client *Client) Context() ContextService {
	return &contextService{client}
}

// Entity возвращает сервис для работы с сущностями и документами.
func (client *Client) Entity() EntityService {
	return &entityService{client}
}

// Report возвращает сервис для работы с отчётами.
func (client *Client) Report() ReportService {
	return &reportService{client}
}

// Security возвращает сервис для получения нового токена.
func (client *Client) Security() SecurityTokenService {
	return NewSecurityTokenService(client)
}

// Notification возвращает сервис для работы с уведомлениями.
func (client *Client) Notification() NotificationService {
	return NewNotificationService(client)
}

type entityService struct{ client *Client }

func (service *entityService) Application() ApplicationService {
	return NewApplicationService(service.client)
}

func (service *entityService) Assortment() AssortmentService {
	return NewAssortmentService(service.client)
}

func (service *entityService) BonusProgram() BonusProgramService {
	return NewBonusProgramService(service.client)
}

func (service *entityService) BonusTransaction() BonusTransactionService {
	return NewBonusTransactionService(service.client)
}

func (service *entityService) Bundle() BundleService {
	return NewBundleService(service.client)
}

func (service *entityService) CashIn() CashInService {
	return NewCashInService(service.client)
}

func (service *entityService) CashOut() CashOutService {
	return NewCashOutService(service.client)
}

func (service *entityService) CommissionReportIn() CommissionReportInService {
	return NewCommissionReportInService(service.client)
}

func (service *entityService) CommissionReportOut() CommissionReportOutService {
	return NewCommissionReportOutService(service.client)
}

func (service *entityService) Consignment() ConsignmentService {
	return NewConsignmentService(service.client)
}

func (service *entityService) Contract() ContractService {
	return NewContractService(service.client)
}

func (service *entityService) Counterparty() CounterpartyService {
	return NewCounterpartyService(service.client)
}

func (service *entityService) CounterPartyAdjustment() CounterPartyAdjustmentService {
	return NewCounterPartyAdjustmentService(service.client)
}

func (service *entityService) Country() CountryService {
	return NewCountryService(service.client)
}

func (service *entityService) Currency() CurrencyService {
	return NewCurrencyService(service.client)
}

func (service *entityService) CustomEntity() CustomEntityService {
	return NewCustomEntityService(service.client)
}

func (service *entityService) CustomerOrder() CustomerOrderService {
	return NewCustomerOrderService(service.client)
}

func (service *entityService) Demand() DemandService {
	return NewDemandService(service.client)
}

func (service *entityService) Discount() DiscountService {
	return NewDiscountService(service.client)
}

func (service *entityService) Employee() EmployeeService {
	return NewEmployeeService(service.client)
}

func (service *entityService) Enter() EnterService {
	return NewEnterService(service.client)
}

func (service *entityService) ExpenseItem() ExpenseItemService {
	return NewExpenseItemService(service.client)
}

func (service *entityService) FactureIn() FactureInService {
	return NewFactureInService(service.client)
}

func (service *entityService) FactureOut() FactureOutService {
	return NewFactureOutService(service.client)
}

func (service *entityService) Group() GroupService {
	return NewGroupService(service.client)
}

func (service *entityService) InternalOrder() InternalOrderService {
	return NewInternalOrderService(service.client)
}

func (service *entityService) Inventory() InventoryService {
	return NewInventoryService(service.client)
}

func (service *entityService) InvoiceIn() InvoiceInService {
	return NewInvoiceInService(service.client)
}

func (service *entityService) InvoiceOut() InvoiceOutService {
	return NewInvoiceOutService(service.client)
}

func (service *entityService) Loss() LossService {
	return NewLossService(service.client)
}

func (service *entityService) Metadata() MetadataService {
	return NewMetadataService(service.client)
}

func (service *entityService) Move() MoveService {
	return NewMoveService(service.client)
}

func (service *entityService) Organization() OrganizationService {
	return NewOrganizationService(service.client)
}

func (service *entityService) PaymentIn() PaymentInService {
	return NewPaymentInService(service.client)
}

func (service *entityService) PaymentOut() PaymentOutService {
	return NewPaymentOutService(service.client)
}

func (service *entityService) Prepayment() PrepaymentService {
	return NewPrepaymentService(service.client)
}

func (service *entityService) PrepaymentReturn() PrepaymentReturnService {
	return NewPrepaymentReturnService(service.client)
}

func (service *entityService) PriceList() PriceListService {
	return NewPriceListService(service.client)
}

func (service *entityService) Processing() ProcessingService {
	return NewProcessingService(service.client)
}

func (service *entityService) ProcessingOrder() ProcessingOrderService {
	return NewProcessingOrderService(service.client)
}

func (service *entityService) ProcessingPlan() ProcessingPlanService {
	return NewProcessingPlanService(service.client)
}

func (service *entityService) ProcessingPlanFolder() ProcessingPlanFolderService {
	return NewProcessingPlanFolderService(service.client)
}

func (service *entityService) ProcessingProcess() ProcessingProcessService {
	return NewProcessingProcessService(service.client)
}

func (service *entityService) ProcessingStage() ProcessingStageService {
	return NewProcessingStageService(service.client)
}

func (service *entityService) Product() ProductService {
	return NewProductService(service.client)
}

func (service *entityService) ProductFolder() ProductFolderService {
	return NewProductFolderService(service.client)
}

func (service *entityService) ProductionTask() ProductionTaskService {
	return NewProductionTaskService(service.client)
}

func (service *entityService) ProductionStage() ProductionStageService {
	return NewProductionStageService(service.client)
}

func (service *entityService) ProductionStageCompletion() ProductionStageCompletionService {
	return NewProductionStageCompletionService(service.client)
}

func (service *entityService) Project() ProjectService {
	return NewProjectService(service.client)
}

func (service *entityService) PurchaseOrder() PurchaseOrderService {
	return NewPurchaseOrderService(service.client)
}

func (service *entityService) PurchaseReturn() PurchaseReturnService {
	return NewPurchaseReturnService(service.client)
}

func (service *entityService) Region() RegionService {
	return NewRegionService(service.client)
}

func (service *entityService) RetailDemand() RetailDemandService {
	return NewRetailDemandService(service.client)
}

func (service *entityService) RetailDrawerCashIn() RetailDrawerCashInService {
	return NewRetailDrawerCashInService(service.client)
}

func (service *entityService) RetailDrawerCashOut() RetailDrawerCashOutService {
	return NewRetailDrawerCashOutService(service.client)
}

func (service *entityService) RetailSalesReturn() RetailSalesReturnService {
	return NewRetailSalesReturnService(service.client)
}

func (service *entityService) RetailShift() RetailShiftService {
	return NewRetailShiftService(service.client)
}

func (service *entityService) RetailStore() RetailStoreService {
	return NewRetailStoreService(service.client)
}

func (service *entityService) Role() RoleService {
	return NewRoleService(service.client)
}

func (service *entityService) SalesChannel() SalesChannelService {
	return NewSalesChannelService(service.client)
}

func (service *entityService) SalesReturn() SalesReturnService {
	return NewSalesReturnService(service.client)
}

func (service *entityService) Service() ServiceService {
	return NewServiceService(service.client)
}

func (service *entityService) Store() StoreService {
	return NewStoreService(service.client)
}

func (service *entityService) Subscription() SubscriptionService {
	return NewSubscriptionService(service.client)
}

func (service *entityService) Supply() SupplyService {
	return NewSupplyService(service.client)
}

func (service *entityService) Task() TaskService {
	return NewTaskService(service.client)
}

func (service *entityService) TaxRate() TaxRateService {
	return NewTaxRateService(service.client)
}

func (service *entityService) Thing() ThingService {
	return NewThingService(service.client)
}

func (service *entityService) Uom() UomService {
	return NewUomService(service.client)
}

func (service *entityService) Variant() VariantService {
	return NewVariantService(service.client)
}

func (service *entityService) Webhook() WebhookService {
	return NewWebhookService(service.client)
}

func (service *entityService) WebhookStock() WebhookStockService {
	return NewWebhookStockService(service.client)
}

type contextService struct{ client *Client }

func (service *contextService) CompanySettings() ContextCompanySettingsService {
	return NewContextCompanySettingsService(service.client)
}

func (service *contextService) Employee() ContextEmployeeService {
	return NewContextEmployeeService(service.client)
}

func (service *contextService) UserSettings() UserSettingsService {
	return NewContextUserSettingsService(service.client)
}

type reportService struct{ client *Client }

func (service *reportService) Counterparty() ReportCounterpartyService {
	return NewReportCounterpartyService(service.client)
}

func (service *reportService) Dashboard() ReportDashboardService {
	return NewReportDashboardService(service.client)
}

func (service *reportService) Money() ReportMoneyService {
	return NewReportMoneyService(service.client)
}

func (service *reportService) Profit() ReportProfitService {
	return NewReportProfitService(service.client)
}

func (service *reportService) Sales() ReportSalesService {
	return NewReportSalesService(service.client)
}

func (service *reportService) Orders() ReportOrdersService {
	return NewReportOrdersService(service.client)
}

func (service *reportService) Stock() ReportStockService {
	return NewReportStockService(service.client)
}

func (service *reportService) Turnover() ReportTurnoverService {
	return NewReportTurnoverService(service.client)
}
