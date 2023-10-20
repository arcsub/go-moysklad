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
	Version          = "v0.0.7"
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

// Client базовый клиент для взаимодействия с API МойСклад.
type Client struct {
	clientMu              sync.Mutex
	client                *RetryableClient
	disableWebhookContent bool // Отключить уведомления вебхуков в контексте данного клиента
}

// NewClient возвращает новый клиент для работы с API МойСклад.
// Данный клиент не имеет встроенных сервисов.
// Его необходимо передавать при создании каждого нового экземпляра сервиса.
func NewClient() *Client {
	rc := newRetryableClient()
	return &Client{client: rc}
}

func (c *Client) Async() *AsyncService {
	return NewAsyncService(c)
}

func (c *Client) Audit() *AuditService {
	return NewAuditService(c)
}

func (c *Client) Context() *ContextService {
	return &ContextService{c}
}

func (c *Client) Entity() *EntityService {
	return &EntityService{c}
}

func (c *Client) Report() *ReportService {
	return &ReportService{c}
}

func (c *Client) Security() *SecurityTokenService {
	return NewSecurityTokenService(c)
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
		if req.Header.Get(headerGetContent) != "true" {
			_ = resp.Body.Close()
		}
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

// EntityService
// Сервис для работы с сущностями и документами.
type EntityService struct {
	client *Client
}

func (s *EntityService) Application() *ApplicationService {
	return NewApplicationService(s.client)
}

func (s *EntityService) Assortment() *AssortmentService {
	return NewAssortmentService(s.client)
}

func (s *EntityService) BonusProgram() *BonusProgramService {
	return NewBonusProgramService(s.client)
}

func (s *EntityService) BonusTransaction() *BonusTransactionService {
	return NewBonusTransactionService(s.client)
}

func (s *EntityService) Bundle() *BundleService {
	return NewBundleService(s.client)
}

func (s *EntityService) CashIn() *CashInService {
	return NewCashInService(s.client)
}

func (s *EntityService) CashOut() *CashOutService {
	return NewCashOutService(s.client)
}

func (s *EntityService) CommissionReportIn() *CommissionReportInService {
	return NewCommissionReportInService(s.client)
}

func (s *EntityService) CommissionReportOut() *CommissionReportOutService {
	return NewCommissionReportOutService(s.client)
}

func (s *EntityService) Consignment() *ConsignmentService {
	return NewConsignmentService(s.client)
}

func (s *EntityService) Contract() *ContractService {
	return NewContractService(s.client)
}

func (s *EntityService) Counterparty() *CounterpartyService {
	return NewCounterpartyService(s.client)
}

func (s *EntityService) CounterPartyAdjustment() *CounterPartyAdjustmentService {
	return NewCounterPartyAdjustmentService(s.client)
}

func (s *EntityService) Country() *CountryService {
	return NewCountryService(s.client)
}

func (s *EntityService) Currency() *CurrencyService {
	return NewCurrencyService(s.client)
}

func (s *EntityService) CustomEntity() *CustomEntityService {
	return NewCustomEntityService(s.client)
}

func (s *EntityService) CustomerOrder() *CustomerOrderService {
	return NewCustomerOrderService(s.client)
}

func (s *EntityService) Demand() *DemandService {
	return NewDemandService(s.client)
}

func (s *EntityService) Discount() *DiscountService {
	return NewDiscountService(s.client)
}

func (s *EntityService) Employee() *EmployeeService {
	return NewEmployeeService(s.client)
}

func (s *EntityService) Enter() *EnterService {
	return NewEnterService(s.client)
}

func (s *EntityService) ExpenseItem() *ExpenseItemService {
	return NewExpenseItemService(s.client)
}

func (s *EntityService) FactureIn() *FactureInService {
	return NewFactureInService(s.client)
}

func (s *EntityService) FactureOut() *FactureOutService {
	return NewFactureOutService(s.client)
}

func (s *EntityService) Group() *GroupService {
	return NewGroupService(s.client)
}

func (s *EntityService) InternalOrder() *InternalOrderService {
	return NewInternalOrderService(s.client)
}

func (s *EntityService) Inventory() *InventoryService {
	return NewInventoryService(s.client)
}

func (s *EntityService) InvoiceIn() *InvoiceInService {
	return NewInvoiceInService(s.client)
}

func (s *EntityService) InvoiceOut() *InvoiceOutService {
	return NewInvoiceOutService(s.client)
}

func (s *EntityService) Loss() *LossService {
	return NewLossService(s.client)
}

func (s *EntityService) Metadata() *MetadataService {
	return NewMetadataService(s.client)
}

func (s *EntityService) Move() *MoveService {
	return NewMoveService(s.client)
}

func (s *EntityService) Organization() *OrganizationService {
	return NewOrganizationService(s.client)
}

func (s *EntityService) PaymentIn() *PaymentInService {
	return NewPaymentInService(s.client)
}

func (s *EntityService) PaymentOut() *PaymentOutService {
	return NewPaymentOutService(s.client)
}

func (s *EntityService) Prepayment() *PrepaymentService {
	return NewPrepaymentService(s.client)
}

func (s *EntityService) PrepaymentReturn() *PrepaymentReturnService {
	return NewPrepaymentReturnService(s.client)
}

func (s *EntityService) PriceList() *PriceListService {
	return NewPriceListService(s.client)
}

func (s *EntityService) Processing() *ProcessingService {
	return NewProcessingService(s.client)
}

func (s *EntityService) ProcessingOrder() *ProcessingOrderService {
	return NewProcessingOrderService(s.client)
}

func (s *EntityService) ProcessingPlan() *ProcessingPlanService {
	return NewProcessingPlanService(s.client)
}

func (s *EntityService) ProcessingPlanFolder() *ProcessingPlanFolderService {
	return NewProcessingPlanFolderService(s.client)
}

func (s *EntityService) ProcessingProcess() *ProcessingProcessService {
	return NewProcessingProcessService(s.client)
}

func (s *EntityService) ProcessingStage() *ProcessingStageService {
	return NewProcessingStageService(s.client)
}

func (s *EntityService) Product() *ProductService {
	return NewProductService(s.client)
}

func (s *EntityService) ProductFolder() *ProductFolderService {
	return NewProductFolderService(s.client)
}

func (s *EntityService) Project() *ProjectService {
	return NewProjectService(s.client)
}

func (s *EntityService) PurchaseOrder() *PurchaseOrderService {
	return NewPurchaseOrderService(s.client)
}

func (s *EntityService) PurchaseReturn() *PurchaseReturnService {
	return NewPurchaseReturnService(s.client)
}

func (s *EntityService) Region() *RegionService {
	return NewRegionService(s.client)
}

func (s *EntityService) RetailDemand() *RetailDemandService {
	return NewRetailDemandService(s.client)
}

func (s *EntityService) RetailDrawerCashIn() *RetailDrawerCashInService {
	return NewRetailDrawerCashInService(s.client)
}

func (s *EntityService) RetailDrawerCashOut() *RetailDrawerCashOutService {
	return NewRetailDrawerCashOutService(s.client)
}

func (s *EntityService) RetailSalesReturn() *RetailSalesReturnService {
	return NewRetailSalesReturnService(s.client)
}

func (s *EntityService) RetailShift() *RetailShiftService {
	return NewRetailShiftService(s.client)
}

func (s *EntityService) RetailStore() *RetailStoreService {
	return NewRetailStoreService(s.client)
}

func (s *EntityService) Role() *RoleService {
	return NewRoleService(s.client)
}

func (s *EntityService) SalesChannel() *SalesChannelService {
	return NewSalesChannelService(s.client)
}

func (s *EntityService) SalesReturn() *SalesReturnService {
	return NewSalesReturnService(s.client)
}

func (s *EntityService) Service() *ServiceService {
	return NewServiceService(s.client)
}

func (s *EntityService) Store() *StoreService {
	return NewStoreService(s.client)
}

func (s *EntityService) Subscription() *SubscriptionService {
	return NewSubscriptionService(s.client)
}

func (s *EntityService) Supply() *SupplyService {
	return NewSupplyService(s.client)
}

func (s *EntityService) Task() *TaskService {
	return NewTaskService(s.client)
}

func (s *EntityService) TaxRate() *TaxRateService {
	return NewTaxRateService(s.client)
}

func (s *EntityService) Uom() *UomService {
	return NewUomService(s.client)
}

func (s *EntityService) Variant() *VariantService {
	return NewVariantService(s.client)
}

func (s *EntityService) Webhook() *WebhookService {
	return NewWebhookService(s.client)
}

func (s *EntityService) WebhookStock() *WebhookStockService {
	return NewWebhookStockService(s.client)
}

// ContextService
// Сервис для работы с контекстом.
type ContextService struct {
	client *Client
}

func (s *ContextService) CompanySettings() *ContextCompanySettingsService {
	return NewContextCompanySettingsService(s.client)
}

func (s *ContextService) Employee() *ContextEmployeeService {
	return NewContextEmployeeService(s.client)
}

func (s *ContextService) UserSettings() *UserSettingsService {
	return NewContextUserSettingsService(s.client)
}

// ReportService
// Сервис для работы с отчётами.
type ReportService struct {
	client *Client
}

func (s *ReportService) Counterparty() *ReportCounterpartyService {
	return NewReportCounterpartyService(s.client)
}

func (s *ReportService) Dashboard() *ReportDashboardService {
	return NewReportDashboardService(s.client)
}

func (s *ReportService) Money() *ReportMoneyService {
	return NewReportMoneyService(s.client)
}

func (s *ReportService) Profit() *ReportProfitService {
	return NewReportProfitService(s.client)
}

func (s *ReportService) Sales() *ReportSalesService {
	return NewReportSalesService(s.client)
}

func (s *ReportService) Orders() *ReportOrdersService {
	return NewReportOrdersService(s.client)
}

func (s *ReportService) Stock() *ReportStockService {
	return NewReportStockService(s.client)
}

func (s *ReportService) Turnover() *ReportTurnoverService {
	return NewReportTurnoverService(s.client)
}

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (fn roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return fn(r)
}
