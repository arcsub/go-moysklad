package moysklad

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
)

const (
	Version          = "v0.0.1"
	baseApiURL       = "https://api.moysklad.ru/api/remap/1.2/"
	defaultUserAgent = "go-moysklad/" + Version

	headerRateLimit         = "X-RateLimit-Limit"           // Количество запросов, которые равномерно можно сделать в течение интервала до появления 429 ошибки.
	headerRateRemaining     = "X-RateLimit-Remaining"       // Число запросов, которые можно отправить до получения 429 ошибки.
	headerRetryTimeInterval = "X-Lognex-Retry-TimeInterval" // Интервал в миллисекундах, в течение которого можно сделать эти запросы
	headerRateReset         = "X-Lognex-Reset"              // Время до сброса ограничения в миллисекундах. Равно нулю, если ограничение не установлено.
	headerRetryAfter        = "X-Lognex-Retry-After"        // Время до сброса ограничения в миллисекундах.
	headerWebHookDisable    = "X-Lognex-WebHook-Disable"    // Заголовок временного отключения уведомлений через API.
	headerGetContent        = "X-Lognex-Get-Content"        // Заголовок для получения файла напрямую.

	headerContentDisposition = "Content-Disposition" // Заголовок содержит название файла при `X-Lognex-Get-Content: true`

	MaxFiles      = 100  // Максимальное количество файлов
	MaxImages     = 10   // Максимальное количество изображений
	MaxPositions  = 1000 // Максимальное число объектов, передаваемых в одном массиве в запросе
	MaxPrintCount = 1000 // TODO: Максимальное количество ценников/термоэтикеток
)

// ClientExtended расширенный клиент для взаимодействия с API МойСклад.
// Отличается от базового наличием готовых к работе сервисов.
type ClientExtended struct {
	*Client
	Async    *AsyncService
	Audit    *AuditService
	Context  *ContextService
	Entity   *EntityService
	Report   *ReportService
	Security *SecurityService
}

// Client базовый клиент для взаимодействия с API МойСклад.
type Client struct {
	clientMu              sync.Mutex
	client                *RetryableClient
	disableWebhookContent bool // Отключить уведомления вебхуков в контексте данного клиента
}

// EntityService
// Сервис для работы с сущностями и документами.
type EntityService struct {
	Application            *ApplicationService
	Assortment             *AssortmentService
	BonusProgram           *BonusProgramService
	BonusTransaction       *BonusTransactionService
	Bundle                 *BundleService
	CashIn                 *CashInService
	CashOut                *CashOutService
	CommissionReportIn     *CommissionReportInService
	CommissionReportOut    *CommissionReportOutService
	Consignment            *ConsignmentService
	Contract               *ContractService
	Counterparty           *CounterpartyService
	CounterPartyAdjustment *CounterPartyAdjustmentService
	Country                *CountryService
	Currency               *CurrencyService
	CustomEntity           *CustomEntityService
	CustomerOrder          *CustomerOrderService
	Demand                 *DemandService
	Discount               *DiscountService
	Employee               *EmployeeService
	Enter                  *EnterService
	ExpenseItem            *ExpenseItemService
	FactureIn              *FactureInService
	FactureOut             *FactureOutService
	Group                  *GroupService
	InternalOrder          *InternalOrderService
	Inventory              *InventoryService
	InvoiceIn              *InvoiceInService
	InvoiceOut             *InvoiceOutService
	Loss                   *LossService
	Metadata               *MetadataService
	Move                   *MoveService
	Organization           *OrganizationService
	PaymentIn              *PaymentInService
	PaymentOut             *PaymentOutService
	Prepayment             *PrepaymentService
	PrepaymentReturn       *PrepaymentReturnService
	PriceList              *PriceListService
	Processing             *ProcessingService
	ProcessingOrder        *ProcessingOrderService
	ProcessingPlan         *ProcessingPlanService
	ProcessingPlanFolder   *ProcessingPlanFolderService
	ProcessingProcess      *ProcessingProcessService
	ProcessingStage        *ProcessingStageService
	Product                *ProductService
	ProductFolder          *ProductFolderService
	Project                *ProjectService
	PurchaseOrder          *PurchaseOrderService
	PurchaseReturn         *PurchaseReturnService
	Region                 *RegionService
	RetailDemand           *RetailDemandService
	RetailDrawerCashIn     *RetailDrawerCashInService
	RetailDrawerCashOut    *RetailDrawerCashOutService
	RetailSalesReturn      *RetailSalesReturnService
	RetailShift            *RetailShiftService
	RetailStore            *RetailStoreService
	Role                   *RoleService
	SalesChannel           *SalesChannelService
	SalesReturn            *SalesReturnService
	Service                *ServiceService
	Store                  *StoreService
	Subscription           *SubscriptionService
	Supply                 *SupplyService
	Task                   *TaskService
	TaxRate                *TaxRateService
	Uom                    *UomService
	Variant                *VariantService
	Webhook                *WebhookService
	WebhookStock           *WebhookStockService
}

func NewEntityService(client *Client) *EntityService {
	return &EntityService{
		Application:            NewApplicationService(client),
		Assortment:             NewAssortmentService(client),
		BonusProgram:           NewBonusProgramService(client),
		BonusTransaction:       NewBonusTransactionService(client),
		Bundle:                 NewBundleService(client),
		CashIn:                 NewCashInService(client),
		CashOut:                NewCashOutService(client),
		CommissionReportIn:     NewCommissionReportInService(client),
		CommissionReportOut:    NewCommissionReportOutService(client),
		Consignment:            NewConsignmentService(client),
		Contract:               NewContractService(client),
		Counterparty:           NewCounterpartyService(client),
		CounterPartyAdjustment: NewCounterPartyAdjustmentService(client),
		Country:                NewCountryService(client),
		Currency:               NewCurrencyService(client),
		CustomEntity:           NewCustomEntityService(client),
		CustomerOrder:          NewCustomerOrderService(client),
		Demand:                 NewDemandService(client),
		Discount:               NewDiscountService(client),
		Employee:               NewEmployeeService(client),
		Enter:                  NewEnterService(client),
		ExpenseItem:            NewExpenseItemService(client),
		FactureIn:              NewFactureInService(client),
		FactureOut:             NewFactureOutService(client),
		Group:                  NewGroupService(client),
		InternalOrder:          NewInternalOrderService(client),
		Inventory:              NewInventoryService(client),
		InvoiceIn:              NewInvoiceInService(client),
		InvoiceOut:             NewInvoiceOutService(client),
		Loss:                   NewLossService(client),
		Metadata:               NewMetadataService(client),
		Move:                   NewMoveService(client),
		Organization:           NewOrganizationService(client),
		PaymentIn:              NewPaymentInService(client),
		PaymentOut:             NewPaymentOutService(client),
		Prepayment:             NewPrepaymentService(client),
		PrepaymentReturn:       NewPrepaymentReturnService(client),
		PriceList:              NewPriceListService(client),
		Processing:             NewProcessingService(client),
		ProcessingOrder:        NewProcessingOrderService(client),
		ProcessingPlan:         NewProcessingPlanService(client),
		ProcessingPlanFolder:   NewProcessingPlanFolderService(client),
		ProcessingProcess:      NewProcessingProcessService(client),
		ProcessingStage:        NewProcessingStageService(client),
		Product:                NewProductService(client),
		ProductFolder:          NewProductFolderService(client),
		Project:                NewProjectService(client),
		PurchaseOrder:          NewPurchaseOrderService(client),
		PurchaseReturn:         NewPurchaseReturnService(client),
		Region:                 NewRegionService(client),
		RetailDemand:           NewRetailDemandService(client),
		RetailDrawerCashIn:     NewRetailDrawerCashInService(client),
		RetailDrawerCashOut:    NewRetailDrawerCashOutService(client),
		RetailSalesReturn:      NewRetailSalesReturnService(client),
		RetailShift:            NewRetailShiftService(client),
		RetailStore:            NewRetailStoreService(client),
		Role:                   NewRoleService(client),
		SalesChannel:           NewSalesChannelService(client),
		SalesReturn:            NewSalesReturnService(client),
		Service:                NewServiceService(client),
		Store:                  NewStoreService(client),
		Subscription:           NewSubscriptionService(client),
		Supply:                 NewSupplyService(client),
		Task:                   NewTaskService(client),
		TaxRate:                NewTaxRateService(client),
		Uom:                    NewUomService(client),
		Variant:                NewVariantService(client),
		Webhook:                NewWebhookService(client),
		WebhookStock:           NewWebhookStockService(client),
	}
}

// ContextService
// Сервис для работы с контекстом.
type ContextService struct {
	CompanySettings *ContextCompanySettingsService
	Employee        *ContextEmployeeService
	UserSettings    *UserSettingsService
}

func NewContextService(client *Client) *ContextService {
	return &ContextService{
		CompanySettings: NewContextCompanySettingsService(client),
		Employee:        NewContextEmployeeService(client),
		UserSettings:    NewContextUserSettingsService(client),
	}
}

// ReportService
// Сервис для работы с отчётами.
type ReportService struct {
	Counterparty *ReportCounterpartyService
	Dashboard    *ReportDashboardService
	Money        *ReportMoneyService
	Profit       *ReportProfitService
	Sales        *ReportSalesService
	Orders       *ReportOrdersService
	Stock        *ReportStockService
	Turnover     *ReportTurnoverService
}

func NewReportService(client *Client) *ReportService {
	return &ReportService{
		Counterparty: NewReportCounterpartyService(client),
		Dashboard:    NewReportDashboardService(client),
		Money:        NewReportMoneyService(client),
		Profit:       NewReportProfitService(client),
		Sales:        NewReportSalesService(client),
		Orders:       NewReportOrdersService(client),
		Stock:        NewReportStockService(client),
		Turnover:     NewReportTurnoverService(client),
	}
}

type SecurityService struct {
	Token *SecurityTokenService
}

func NewSecurityService(client *Client) *SecurityService {
	return &SecurityService{
		Token: NewSecurityTokenService(client),
	}
}

// NewClient возвращает новый клиент для работы с API МойСклад.
// Данный клиент не имеет встроенных сервисов.
// Его необходимо передавать при создании каждого нового экземпляра сервиса.
func NewClient() *Client {
	rc := newRetryableClient()
	return &Client{client: rc}
}

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (fn roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return fn(r)
}

// Extend инициализирует сервисы и возвращает "расширенный" клиент.
func (c *Client) Extend() *ClientExtended {
	c2 := &ClientExtended{Client: c.copy()}
	c2.initializeServices()
	return c2
}

// WithTokenAuth возвращает клиент с авторизацией через токен.
func (c *Client) WithTokenAuth(token string) *Client {
	c2 := c.copy()
	transport := c2.client.HTTPClient.Transport
	if transport == nil {
		transport = http.DefaultTransport
	}
	c2.client.HTTPClient.Transport = roundTripperFunc(
		func(req *http.Request) (*http.Response, error) {
			req = req.Clone(req.Context())
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
			return transport.RoundTrip(req)
		},
	)
	return c2
}

// WithBasicAuth возвращает клиент с базовой авторизацией логин/пароль.
func (c *Client) WithBasicAuth(username, password string) *Client {
	c2 := c.copy()
	transport := c2.client.HTTPClient.Transport
	if transport == nil {
		transport = http.DefaultTransport
	}
	c2.client.HTTPClient.Transport = roundTripperFunc(
		func(req *http.Request) (*http.Response, error) {
			req = req.Clone(req.Context())
			req.SetBasicAuth(username, password)
			return transport.RoundTrip(req)
		},
	)
	return c2
}

// WithMaxRetries устанавливает максимальное кол-во попыток для одного запроса.
func (c *Client) WithMaxRetries(retries int) *Client {
	c.client.RetryMax = retries
	return c
}

// WithDisabledWebhookContent устанавливает флаг, который отвечает
// за формирование заголовка временного отключения уведомления вебхуков через API (X-Lognex-WebHook-Disable).
// Подробнее: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-vebhuki-primer-webhuka-zagolowok-wremennogo-otklucheniq-cherez-api
func (c *Client) WithDisabledWebhookContent(value bool) *Client {
	c.disableWebhookContent = value
	return c
}

// copy возвращает копию клиента.
func (c *Client) copy() *Client {
	c.clientMu.Lock()
	defer c.clientMu.Unlock()
	clone := Client{
		client:                c.client,
		disableWebhookContent: c.disableWebhookContent,
	}
	return &clone
}

// initializeServices инициализирует сервисы клиента.
func (c *ClientExtended) initializeServices() {
	c.Async = NewAsyncService(c.Client)
	c.Audit = NewAuditService(c.Client)
	c.Context = NewContextService(c.Client)
	c.Entity = NewEntityService(c.Client)
	c.Report = NewReportService(c.Client)
	c.Security = NewSecurityService(c.Client)
}

// RateLimit содержит значения ограничений API.
//type RateLimit struct {
//	Limit             int   // Количество запросов, которые равномерно можно сделать в течение интервала до появления 429 ошибки
//	RetryTimeInterval int64 // Интервал в миллисекундах, в течение которого можно сделать эти запросы
//	Remaining         int   // Число запросов, которые можно отправить до получения 429 ошибки
//	Reset             int64 // Время до сброса ограничения в миллисекундах. Равно нулю, если ограничение не установлено
//	RetryAfter        int64 // Время до сброса ограничения в миллисекундах.
//}

// Response Ответ от API МойСклад.
// Оборачивает стандартный http.Response, полученный от API МойСклад.
type Response struct {
	*http.Response
}

// newResponse создаёт экземпляр Response с http.Response.
func newResponse(r *http.Response) (*Response, error) {
	resp := &Response{r}

	// заголовок Content-Encoding (01.12.2023)
	if !resp.Uncompressed && resp.Header.Get("Content-Encoding") == "gzip" {
		body, err := gzip.NewReader(resp.Body)
		if err != nil {
			return resp, err
		}
		resp.Body = body
	}

	// проверяем ответ от сервиса на наличие ошибок API.
	if code := r.StatusCode; code <= http.StatusFound {
		return resp, nil
	}

	apiErrs := &ApiErrors{}
	data, err := io.ReadAll(r.Body)
	if err == nil && data != nil {
		_ = json.Unmarshal(data, apiErrs)
	}
	return resp, apiErrs
}

// parseRateLimits парсит заголовки ответа, связанные с ограничениями
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-obschie-swedeniq-ogranicheniq
//func parseRateLimits(r *http.Response) RateLimit {
//	var rl RateLimit
//	if limit := r.Header.Get(headerRateLimit); limit != "" {
//		rl.Limit, _ = strconv.Atoi(limit)
//	}
//	if retryTimeInterval := r.Header.Get(headerRetryTimeInterval); retryTimeInterval != "" {
//		if retryTimeIntervalValue, err := strconv.ParseInt(retryTimeInterval, 10, 64); err == nil {
//			rl.RetryTimeInterval = retryTimeIntervalValue
//		}
//	}
//	if remaining := r.Header.Get(headerRateRemaining); remaining != "" {
//		rl.Remaining, _ = strconv.Atoi(remaining)
//	}
//	if resetHeader := r.Header.Get(headerRateReset); resetHeader != "" {
//		if resetValue, err := strconv.ParseInt(resetHeader, 10, 64); err == nil {
//			rl.Reset = resetValue
//		}
//	}
//	if retryAfterHeader := r.Header.Get(headerRetryAfter); retryAfterHeader != "" {
//		if retryAfterValue, err := strconv.ParseInt(retryAfterHeader, 10, 64); err == nil {
//			rl.RetryAfter = retryAfterValue
//		}
//	}
//	return rl
//}

// Заголовок Content-Encoding и декомпрессия данных (01.12.2023)
func decompressResponse(r *http.Response) (err error) {
	if !r.Uncompressed && r.Header.Get("Content-Encoding") == "gzip" {
		if r.Body, err = gzip.NewReader(r.Body); err != nil {
			return
		}
	}
	return
}

// Do отправляет запрос в API и обрабатывает ответ.
// В случае получения в ответе ошибки API возвращаемая ошибка будет содержать дополнительную информацию.
func (c *Client) Do(ctx context.Context, req *http.Request) (*Response, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	req.WithContext(ctx)

	// Convert the request to be retryable.
	retryableReq, err := FromRequest(req)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(retryableReq)

	defer func() {
		_ = resp.Body.Close()
	}()

	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		if e, ok := err.(*url.Error); ok {
			if u, err := url.Parse(e.URL); err == nil {
				e.URL = u.String()
				return nil, e
			}
		}
		return nil, err
	}

	return newResponse(resp)
}
