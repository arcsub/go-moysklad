package moysklad

// EntityService
// Сервис для работы с сущностями и документами.
type EntityService interface {
	Application() ApplicationService
	Assortment() AssortmentService
	BonusProgram() BonusProgramService
	BonusTransaction() BonusTransactionService
	Bundle() BundleService
	CashIn() CashInService
	CashOut() CashOutService
	CommissionReportIn() CommissionReportInService
	CommissionReportOut() CommissionReportOutService
	Consignment() ConsignmentService
	Contract() ContractService
	Counterparty() CounterpartyService
	CounterPartyAdjustment() CounterPartyAdjustmentService
	Country() CountryService
	Currency() CurrencyService
	CustomEntity() CustomEntityService
	CustomerOrder() CustomerOrderService
	Demand() DemandService
	Discount() DiscountService
	Employee() EmployeeService
	Enter() EnterService
	ExpenseItem() ExpenseItemService
	FactureIn() FactureInService
	FactureOut() FactureOutService
	Group() GroupService
	InternalOrder() InternalOrderService
	Inventory() InventoryService
	InvoiceIn() InvoiceInService
	InvoiceOut() InvoiceOutService
	Loss() LossService
	Metadata() MetadataService
	Move() MoveService
	Organization() OrganizationService
	PaymentIn() PaymentInService
	PaymentOut() PaymentOutService
	Prepayment() PrepaymentService
	PrepaymentReturn() PrepaymentReturnService
	PriceList() PriceListService
	Processing() ProcessingService
	ProcessingOrder() ProcessingOrderService
	ProcessingPlan() ProcessingPlanService
	ProcessingPlanFolder() ProcessingPlanFolderService
	ProcessingProcess() ProcessingProcessService
	ProcessingStage() ProcessingStageService
	Product() ProductService
	ProductFolder() ProductFolderService
	ProductionTask() ProductionTaskService
	ProductionStage() ProductionStageService
	ProductionStageCompletion() ProductionStageCompletionService
	Project() ProjectService
	PurchaseOrder() PurchaseOrderService
	PurchaseReturn() PurchaseReturnService
	Region() RegionService
	RetailDemand() RetailDemandService
	RetailDrawerCashIn() RetailDrawerCashInService
	RetailDrawerCashOut() RetailDrawerCashOutService
	RetailSalesReturn() RetailSalesReturnService
	RetailShift() RetailShiftService
	RetailStore() RetailStoreService
	Role() RoleService
	SalesChannel() SalesChannelService
	SalesReturn() SalesReturnService
	Service() ServiceService
	Store() StoreService
	Subscription() SubscriptionService
	Supply() SupplyService
	Task() TaskService
	TaxRate() TaxRateService
	Thing() ThingService
	Uom() UomService
	Variant() VariantService
	Webhook() WebhookService
	WebhookStock() WebhookStockService
}

// ContextService
// Сервис для работы с контекстом.
type ContextService interface {
	CompanySettings() ContextCompanySettingsService
	Employee() ContextEmployeeService
	UserSettings() UserSettingsService
}

// ReportService
// Сервис для работы с отчётами.
type ReportService interface {
	Counterparty() ReportCounterpartyService
	Dashboard() ReportDashboardService
	Money() ReportMoneyService
	Profit() ReportProfitService
	Sales() ReportSalesService
	Orders() ReportOrdersService
	Stock() ReportStockService
	Turnover() ReportTurnoverService
}

func (client *Client) Async() AsyncService {
	return NewAsyncService(client)
}

func (client *Client) Audit() AuditService {
	return NewAuditService(client)
}

func (client *Client) Context() ContextService {
	return &contextService{client}
}

func (client *Client) Entity() EntityService {
	return &entityService{client}
}

func (client *Client) Report() ReportService {
	return &reportService{client}
}

func (client *Client) Security() SecurityTokenService {
	return NewSecurityTokenService(client)
}

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
