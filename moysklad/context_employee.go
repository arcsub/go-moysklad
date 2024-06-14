package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
)

// ContextEmployee Контекст сотрудника.
// Ключевое слово: employee
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/index.html#mojsklad-json-api-obschie-swedeniq-kontext-zaprosa-sotrudnika-poluchit-kontext-sotrudnika
type ContextEmployee struct {
	Owner       Employee  `json:"owner"`
	Image       Image     `json:"image"`
	Meta        Meta      `json:"meta"`
	Group       Group     `json:"group"`
	Updated     Timestamp `json:"updated"`
	Permissions struct {
		Currency struct {
			View   string `json:"view"`
			Create string `json:"create"`
			Update string `json:"update"`
			Delete string `json:"delete"`
		} `json:"currency"`
		Uom struct {
			View   string `json:"view"`
			Create string `json:"create"`
			Update string `json:"update"`
			Delete string `json:"delete"`
		} `json:"uom"`
		Productfolder struct {
			View   string `json:"view"`
			Create string `json:"create"`
			Update string `json:"update"`
			Delete string `json:"delete"`
			Print  string `json:"print"`
		} `json:"productfolder"`
		Product struct {
			View   string `json:"view"`
			Create string `json:"create"`
			Update string `json:"update"`
			Delete string `json:"delete"`
			Print  string `json:"print"`
		} `json:"product"`
		Bundle struct {
			View   string `json:"view"`
			Create string `json:"create"`
			Update string `json:"update"`
			Delete string `json:"delete"`
			Print  string `json:"print"`
		} `json:"bundle"`
		Service struct {
			View   string `json:"view"`
			Create string `json:"create"`
			Update string `json:"update"`
			Delete string `json:"delete"`
			Print  string `json:"print"`
		} `json:"service"`
		Consignment struct {
			View   string `json:"view"`
			Create string `json:"create"`
			Update string `json:"update"`
			Delete string `json:"delete"`
			Print  string `json:"print"`
		} `json:"consignment"`
		Variant struct {
			View   string `json:"view"`
			Create string `json:"create"`
			Update string `json:"update"`
			Delete string `json:"delete"`
			Print  string `json:"print"`
		} `json:"variant"`
		Store struct {
			View   string `json:"view"`
			Create string `json:"create"`
			Update string `json:"update"`
			Delete string `json:"delete"`
		} `json:"store"`
		Counterparty struct {
			View   string `json:"view"`
			Create string `json:"create"`
			Update string `json:"update"`
			Delete string `json:"delete"`
			Print  string `json:"print"`
		} `json:"counterparty"`
		Organization struct {
			View   string `json:"view"`
			Create string `json:"create"`
			Update string `json:"update"`
			Delete string `json:"delete"`
		} `json:"organization"`
		Employee struct {
			View   string `json:"view"`
			Create string `json:"create"`
			Update string `json:"update"`
			Delete string `json:"delete"`
		} `json:"employee"`
		Settings struct {
			View   string `json:"view"`
			Create string `json:"create"`
			Update string `json:"update"`
			Delete string `json:"delete"`
			Print  string `json:"print"`
		} `json:"settings"`
		Contract struct {
			View   string `json:"view"`
			Create string `json:"create"`
			Update string `json:"update"`
			Delete string `json:"delete"`
			Print  string `json:"print"`
		} `json:"contract"`
		Project struct {
			View   string `json:"view"`
			Create string `json:"create"`
			Update string `json:"update"`
			Delete string `json:"delete"`
		} `json:"project"`
		Saleschannel struct {
			View   string `json:"view"`
			Create string `json:"create"`
			Update string `json:"update"`
			Delete string `json:"delete"`
			Print  string `json:"print"`
		} `json:"saleschannel"`
		Country struct {
			View   string `json:"view"`
			Create string `json:"create"`
			Update string `json:"update"`
			Delete string `json:"delete"`
		} `json:"country"`
		Customentity struct {
			View   string `json:"view"`
			Create string `json:"create"`
			Update string `json:"update"`
			Delete string `json:"delete"`
		} `json:"customentity"`
		Demand struct {
			View    string `json:"view"`
			Create  string `json:"create"`
			Update  string `json:"update"`
			Delete  string `json:"delete"`
			Approve string `json:"approve"`
			Print   string `json:"print"`
		} `json:"demand"`
		Customerorder struct {
			View    string `json:"view"`
			Create  string `json:"create"`
			Update  string `json:"update"`
			Delete  string `json:"delete"`
			Approve string `json:"approve"`
			Print   string `json:"print"`
		} `json:"customerorder"`
		Internalorder struct {
			View    string `json:"view"`
			Create  string `json:"create"`
			Update  string `json:"update"`
			Delete  string `json:"delete"`
			Approve string `json:"approve"`
			Print   string `json:"print"`
		} `json:"internalorder"`
		Invoiceout struct {
			View    string `json:"view"`
			Create  string `json:"create"`
			Update  string `json:"update"`
			Delete  string `json:"delete"`
			Approve string `json:"approve"`
			Print   string `json:"print"`
		} `json:"invoiceout"`
		Invoicein struct {
			View    string `json:"view"`
			Create  string `json:"create"`
			Update  string `json:"update"`
			Delete  string `json:"delete"`
			Approve string `json:"approve"`
			Print   string `json:"print"`
		} `json:"invoicein"`
		Paymentin struct {
			View    string `json:"view"`
			Create  string `json:"create"`
			Update  string `json:"update"`
			Delete  string `json:"delete"`
			Approve string `json:"approve"`
			Print   string `json:"print"`
		} `json:"paymentin"`
		Paymentout struct {
			View    string `json:"view"`
			Create  string `json:"create"`
			Update  string `json:"update"`
			Delete  string `json:"delete"`
			Approve string `json:"approve"`
			Print   string `json:"print"`
		} `json:"paymentout"`
		Cashin struct {
			View    string `json:"view"`
			Create  string `json:"create"`
			Update  string `json:"update"`
			Delete  string `json:"delete"`
			Approve string `json:"approve"`
			Print   string `json:"print"`
		} `json:"cashin"`
		Cashout struct {
			View    string `json:"view"`
			Create  string `json:"create"`
			Update  string `json:"update"`
			Delete  string `json:"delete"`
			Approve string `json:"approve"`
			Print   string `json:"print"`
		} `json:"cashout"`
		Supply struct {
			View    string `json:"view"`
			Create  string `json:"create"`
			Update  string `json:"update"`
			Delete  string `json:"delete"`
			Approve string `json:"approve"`
			Print   string `json:"print"`
		} `json:"supply"`
		Salesreturn struct {
			View    string `json:"view"`
			Create  string `json:"create"`
			Update  string `json:"update"`
			Delete  string `json:"delete"`
			Approve string `json:"approve"`
			Print   string `json:"print"`
		} `json:"salesreturn"`
		Purchasereturn struct {
			View    string `json:"view"`
			Create  string `json:"create"`
			Update  string `json:"update"`
			Delete  string `json:"delete"`
			Approve string `json:"approve"`
			Print   string `json:"print"`
		} `json:"purchasereturn"`
		Retailstore struct {
			View   string `json:"view"`
			Create string `json:"create"`
			Update string `json:"update"`
			Delete string `json:"delete"`
		} `json:"retailstore"`
		Receipttemplate struct {
			View   string `json:"view"`
			Create string `json:"create"`
			Update string `json:"update"`
			Delete string `json:"delete"`
		} `json:"receipttemplate"`
		Retailstorestatus struct {
			View   string `json:"view"`
			Create string `json:"create"`
			Update string `json:"update"`
			Delete string `json:"delete"`
		} `json:"retailstorestatus"`
		Retailshift struct {
			View   string `json:"view"`
			Create string `json:"create"`
			Update string `json:"update"`
			Delete string `json:"delete"`
			Print  string `json:"print"`
		} `json:"retailshift"`
		Retaildemand struct {
			View    string `json:"view"`
			Create  string `json:"create"`
			Update  string `json:"update"`
			Delete  string `json:"delete"`
			Approve string `json:"approve"`
			Print   string `json:"print"`
		} `json:"retaildemand"`
		Retailsalesreturn struct {
			View    string `json:"view"`
			Create  string `json:"create"`
			Update  string `json:"update"`
			Delete  string `json:"delete"`
			Approve string `json:"approve"`
			Print   string `json:"print"`
		} `json:"retailsalesreturn"`
		Retaildrawercashin struct {
			View    string `json:"view"`
			Create  string `json:"create"`
			Update  string `json:"update"`
			Delete  string `json:"delete"`
			Approve string `json:"approve"`
			Print   string `json:"print"`
		} `json:"retaildrawercashin"`
		Retaildrawercashout struct {
			View    string `json:"view"`
			Create  string `json:"create"`
			Update  string `json:"update"`
			Delete  string `json:"delete"`
			Approve string `json:"approve"`
			Print   string `json:"print"`
		} `json:"retaildrawercashout"`
		Prepayment struct {
			View    string `json:"view"`
			Create  string `json:"create"`
			Update  string `json:"update"`
			Delete  string `json:"delete"`
			Approve string `json:"approve"`
			Print   string `json:"print"`
		} `json:"prepayment"`
		Prepaymentreturn struct {
			View    string `json:"view"`
			Create  string `json:"create"`
			Update  string `json:"update"`
			Delete  string `json:"delete"`
			Approve string `json:"approve"`
			Print   string `json:"print"`
		} `json:"prepaymentreturn"`
		Purchaseorder struct {
			View    string `json:"view"`
			Create  string `json:"create"`
			Update  string `json:"update"`
			Delete  string `json:"delete"`
			Approve string `json:"approve"`
			Print   string `json:"print"`
		} `json:"purchaseorder"`
		Move struct {
			View    string `json:"view"`
			Create  string `json:"create"`
			Update  string `json:"update"`
			Delete  string `json:"delete"`
			Approve string `json:"approve"`
			Print   string `json:"print"`
		} `json:"move"`
		Enter struct {
			View    string `json:"view"`
			Create  string `json:"create"`
			Update  string `json:"update"`
			Delete  string `json:"delete"`
			Approve string `json:"approve"`
			Print   string `json:"print"`
		} `json:"enter"`
		Loss struct {
			View    string `json:"view"`
			Create  string `json:"create"`
			Update  string `json:"update"`
			Delete  string `json:"delete"`
			Approve string `json:"approve"`
			Print   string `json:"print"`
		} `json:"loss"`
		Facturein struct {
			View    string `json:"view"`
			Create  string `json:"create"`
			Update  string `json:"update"`
			Delete  string `json:"delete"`
			Approve string `json:"approve"`
			Print   string `json:"print"`
		} `json:"facturein"`
		Factureout struct {
			View    string `json:"view"`
			Create  string `json:"create"`
			Update  string `json:"update"`
			Delete  string `json:"delete"`
			Approve string `json:"approve"`
			Print   string `json:"print"`
		} `json:"factureout"`
		Commissionreportin struct {
			View    string `json:"view"`
			Create  string `json:"create"`
			Update  string `json:"update"`
			Delete  string `json:"delete"`
			Approve string `json:"approve"`
			Print   string `json:"print"`
		} `json:"commissionreportin"`
		Commissionreportout struct {
			View    string `json:"view"`
			Create  string `json:"create"`
			Update  string `json:"update"`
			Delete  string `json:"delete"`
			Approve string `json:"approve"`
			Print   string `json:"print"`
		} `json:"commissionreportout"`
		Pricelist struct {
			View    string `json:"view"`
			Create  string `json:"create"`
			Update  string `json:"update"`
			Delete  string `json:"delete"`
			Approve string `json:"approve"`
			Print   string `json:"print"`
		} `json:"pricelist"`
		Processingplanfolder struct {
			View   string `json:"view"`
			Create string `json:"create"`
			Update string `json:"update"`
			Delete string `json:"delete"`
		} `json:"processingplanfolder"`
		Processingplan struct {
			View   string `json:"view"`
			Create string `json:"create"`
			Update string `json:"update"`
			Delete string `json:"delete"`
		} `json:"processingplan"`
		Processing struct {
			View    string `json:"view"`
			Create  string `json:"create"`
			Update  string `json:"update"`
			Delete  string `json:"delete"`
			Approve string `json:"approve"`
			Print   string `json:"print"`
		} `json:"processing"`
		Processingorder struct {
			View    string `json:"view"`
			Create  string `json:"create"`
			Update  string `json:"update"`
			Delete  string `json:"delete"`
			Approve string `json:"approve"`
			Print   string `json:"print"`
		} `json:"processingorder"`
		Counterpartyadjustment struct {
			View   string `json:"view"`
			Create string `json:"create"`
			Update string `json:"update"`
			Delete string `json:"delete"`
			Print  string `json:"print"`
		} `json:"counterpartyadjustment"`
		Assortment struct {
			View   string `json:"view"`
			Create string `json:"create"`
			Update string `json:"update"`
			Delete string `json:"delete"`
			Print  string `json:"print"`
		} `json:"assortment"`
		Inventory struct {
			View   string `json:"view"`
			Create string `json:"create"`
			Update string `json:"update"`
			Delete string `json:"delete"`
			Print  string `json:"print"`
		} `json:"inventory"`
		Bonustransaction struct {
			View    string `json:"view"`
			Create  string `json:"create"`
			Update  string `json:"update"`
			Delete  string `json:"delete"`
			Approve string `json:"approve"`
			Print   string `json:"print"`
		} `json:"bonustransaction"`
		Crptorder struct {
			View   string `json:"view"`
			Create string `json:"create"`
			Update string `json:"update"`
			Delete string `json:"delete"`
			Print  string `json:"print"`
		} `json:"crptorder"`
		Productiontask struct {
			View    string `json:"view"`
			Create  string `json:"create"`
			Update  string `json:"update"`
			Delete  string `json:"delete"`
			Approve string `json:"approve"`
			Print   string `json:"print"`
		} `json:"productiontask"`
		Productionstagecompletion struct {
			View   string `json:"view"`
			Create string `json:"create"`
			Update string `json:"update"`
			Delete string `json:"delete"`
			Print  string `json:"print"`
		} `json:"productionstagecompletion"`
		Taxrate struct {
			View   string `json:"view"`
			Create string `json:"create"`
			Update string `json:"update"`
			Delete string `json:"delete"`
		} `json:"taxrate"`
		Webhook struct {
			View   string `json:"view"`
			Create string `json:"create"`
			Update string `json:"update"`
			Delete string `json:"delete"`
		} `json:"webhook"`
		Task struct {
			View   string `json:"view"`
			Create string `json:"create"`
			Update string `json:"update"`
			Delete string `json:"delete"`
			Done   string `json:"done"`
		} `json:"task"`
		Dashboard struct {
			View string `json:"view"`
		} `json:"dashboard"`
		Stock struct {
			View string `json:"view"`
		} `json:"stock"`
		CustomAttributes struct {
			View string `json:"view"`
		} `json:"customAttributes"`
		Pnl struct {
			View string `json:"view"`
		} `json:"pnl"`
		CompanyCrm struct {
			View string `json:"view"`
		} `json:"company_crm"`
		TariffCrm struct {
			View string `json:"view"`
		} `json:"tariff_crm"`
		AuditDashboard struct {
			View string `json:"view"`
		} `json:"audit_dashboard"`
		Admin struct {
			View string `json:"view"`
		} `json:"admin"`
		DashboardMoney struct {
			View string `json:"view"`
		} `json:"dashboardMoney"`
		ViewCashFlow struct {
			View string `json:"view"`
		} `json:"viewCashFlow"`
	} `json:"permissions,omitempty"`
	Created      string             `json:"created"`
	MiddleName   string             `json:"middleName"`
	ExternalCode string             `json:"externalCode"`
	ID           string             `json:"id"`
	Position     string             `json:"position"`
	Uid          string             `json:"uid"`
	Email        string             `json:"email"`
	Phone        string             `json:"phone"`
	FirstName    string             `json:"firstName"`
	Name         string             `json:"name"`
	LastName     string             `json:"lastName"`
	FullName     string             `json:"fullName"`
	ShortFio     string             `json:"shortFio"`
	AccountID    string             `json:"accountId"`
	Cashiers     MetaArray[Cashier] `json:"cashiers"`
	Shared       bool               `json:"shared"`
	Archived     bool               `json:"archived"`
}

func (contextEmployee ContextEmployee) String() string {
	return Stringify(contextEmployee)
}

func (contextEmployee ContextEmployee) MetaType() MetaType {
	return MetaTypeEmployeeContext
}

// ContextEmployeeService
// Сервис для работы с контекстом сотрудника.
type ContextEmployeeService interface {
	Get(ctx context.Context, params ...*Params) (*ContextEmployee, *resty.Response, error)
}

func NewContextEmployeeService(client *Client) ContextEmployeeService {
	e := NewEndpoint(client, "context/employee")
	return newMainService[ContextEmployee, any, any, any](e)
}
