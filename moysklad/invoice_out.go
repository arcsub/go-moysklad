package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"time"
)

// InvoiceOut Счет покупателю.
//
// Код сущности: invoiceout
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-schet-pokupatelu
type InvoiceOut struct {
	PayedSum             *float64                       `json:"payedSum,omitempty"`             // Сумма входящих платежей по Счету покупателю
	VatEnabled           *bool                          `json:"vatEnabled,omitempty"`           // Учитывается ли НДС
	AgentAccount         *AgentAccount                  `json:"agentAccount,omitempty"`         // Метаданные счета контрагента
	Applicable           *bool                          `json:"applicable,omitempty"`           // Отметка о проведении
	Demands              Slice[Demand]                  `json:"demands,omitempty"`              // Массив ссылок на связанные отгрузки
	Code                 *string                        `json:"code,omitempty"`                 // Код Счета покупателю
	OrganizationAccount  *AgentAccount                  `json:"organizationAccount,omitempty"`  // Метаданные счета юрлица
	Created              *Timestamp                     `json:"created,omitempty"`              // Дата создания
	Deleted              *Timestamp                     `json:"deleted,omitempty"`              // Момент последнего удаления Счета покупателю
	Description          *string                        `json:"description,omitempty"`          // Комментарий Счета покупателю
	ExternalCode         *string                        `json:"externalCode,omitempty"`         // Внешний код Счета покупателю
	Files                *MetaArray[File]               `json:"files,omitempty"`                // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group                *Group                         `json:"group,omitempty"`                // Отдел сотрудника
	ID                   *uuid.UUID                     `json:"id,omitempty"`                   // ID Счета покупателю
	Meta                 *Meta                          `json:"meta,omitempty"`                 // Метаданные Счета покупателю
	Moment               *Timestamp                     `json:"moment,omitempty"`               // Дата документа
	Name                 *string                        `json:"name,omitempty"`                 // Наименование Счета покупателю
	AccountID            *uuid.UUID                     `json:"accountId,omitempty"`            // ID учётной записи
	Contract             *NullValue[Contract]           `json:"contract,omitempty"`             // Метаданные договора
	Agent                *Agent                         `json:"agent,omitempty"`                // Метаданные контрагента
	Organization         *Organization                  `json:"organization,omitempty"`         // Метаданные юрлица
	PaymentPlannedMoment *Timestamp                     `json:"paymentPlannedMoment,omitempty"` // Планируемая дата оплаты
	Positions            *MetaArray[InvoiceOutPosition] `json:"positions,omitempty"`            // Метаданные позиций Счета покупателю
	Printed              *bool                          `json:"printed,omitempty"`              // Напечатан ли документ
	Project              *NullValue[Project]            `json:"project,omitempty"`              // Метаданные проекта
	Published            *bool                          `json:"published,omitempty"`            // Опубликован ли документ
	Rate                 *NullValue[Rate]               `json:"rate,omitempty"`                 // Валюта
	Shared               *bool                          `json:"shared,omitempty"`               // Общий доступ
	ShippedSum           *float64                       `json:"shippedSum,omitempty"`           // Сумма отгруженного
	State                *NullValue[State]              `json:"state,omitempty"`                // Метаданные статуса счета
	Store                *NullValue[Store]              `json:"store,omitempty"`                // Метаданные склада
	Sum                  *float64                       `json:"sum,omitempty"`                  // Сумма Счета в установленной валюте
	SyncID               *uuid.UUID                     `json:"syncId,omitempty"`               // ID синхронизации
	Updated              *Timestamp                     `json:"updated,omitempty"`              // Момент последнего обновления Счета покупателю
	Owner                *Employee                      `json:"owner,omitempty"`                // Метаданные владельца (Сотрудника)
	VatIncluded          *bool                          `json:"vatIncluded,omitempty"`          // Включен ли НДС в цену
	VatSum               *float64                       `json:"vatSum,omitempty"`               // Сумма НДС
	CustomerOrder        *CustomerOrder                 `json:"customerOrder,omitempty"`        // Ссылка на Заказ Покупателя, с которым связан этот Счет покупателю
	SalesChannel         *SalesChannel                  `json:"salesChannel,omitempty"`         // Метаданные канала продаж
	Payments             Slice[Payment]                 `json:"payments,omitempty"`             // Массив ссылок на связанные операции
	Attributes           Slice[Attribute]               `json:"attributes,omitempty"`           // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (invoiceOut InvoiceOut) Clean() *InvoiceOut {
	if invoiceOut.Meta == nil {
		return nil
	}
	return &InvoiceOut{Meta: invoiceOut.Meta}
}

// AsOperation реализует интерфейс [OperationConverter].
func (invoiceOut InvoiceOut) AsOperation() *Operation {
	return newOperation(invoiceOut)
}

// AsTaskOperation реализует интерфейс [TaskOperationConverter].
func (invoiceOut InvoiceOut) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: invoiceOut.Meta}
}

// AsOperationIn реализует интерфейс [OperationInConverter].
func (invoiceOut InvoiceOut) AsOperationIn() *Operation {
	return invoiceOut.AsOperation()
}

// GetPayedSum возвращает Сумму входящих платежей по Счету покупателю.
func (invoiceOut InvoiceOut) GetPayedSum() float64 {
	return Deref(invoiceOut.PayedSum)
}

// GetVatEnabled возвращает true, если учитывается НДС.
func (invoiceOut InvoiceOut) GetVatEnabled() bool {
	return Deref(invoiceOut.VatEnabled)
}

// GetAgentAccount возвращает Метаданные счета контрагента.
func (invoiceOut InvoiceOut) GetAgentAccount() AgentAccount {
	return Deref(invoiceOut.AgentAccount)
}

// GetApplicable возвращает Отметку о проведении.
func (invoiceOut InvoiceOut) GetApplicable() bool {
	return Deref(invoiceOut.Applicable)
}

// GetDemands возвращает Массив ссылок на связанные отгрузки.
func (invoiceOut InvoiceOut) GetDemands() Slice[Demand] {
	return invoiceOut.Demands
}

// GetCode возвращает Код Счета покупателю.
func (invoiceOut InvoiceOut) GetCode() string {
	return Deref(invoiceOut.Code)
}

// GetOrganizationAccount возвращает Метаданные счета юрлица.
func (invoiceOut InvoiceOut) GetOrganizationAccount() AgentAccount {
	return Deref(invoiceOut.OrganizationAccount)
}

// GetCreated возвращает Дату создания.
func (invoiceOut InvoiceOut) GetCreated() time.Time {
	return Deref(invoiceOut.Created).Time()
}

// GetDeleted возвращает Момент последнего удаления Счета покупателю.
func (invoiceOut InvoiceOut) GetDeleted() time.Time {
	return Deref(invoiceOut.Deleted).Time()
}

// GetDescription возвращает Комментарий Счета покупателю.
func (invoiceOut InvoiceOut) GetDescription() string {
	return Deref(invoiceOut.Description)
}

// GetExternalCode возвращает Внешний код Счета покупателю.
func (invoiceOut InvoiceOut) GetExternalCode() string {
	return Deref(invoiceOut.ExternalCode)
}

// GetFiles возвращает Метаданные массива Файлов.
func (invoiceOut InvoiceOut) GetFiles() MetaArray[File] {
	return Deref(invoiceOut.Files)
}

// GetGroup возвращает Отдел сотрудника.
func (invoiceOut InvoiceOut) GetGroup() Group {
	return Deref(invoiceOut.Group)
}

// GetID возвращает ID Счета покупателю.
func (invoiceOut InvoiceOut) GetID() uuid.UUID {
	return Deref(invoiceOut.ID)
}

// GetMeta возвращает Метаданные Счета покупателю.
func (invoiceOut InvoiceOut) GetMeta() Meta {
	return Deref(invoiceOut.Meta)
}

// GetMoment возвращает Дату документа.
func (invoiceOut InvoiceOut) GetMoment() time.Time {
	return Deref(invoiceOut.Moment).Time()
}

// GetName возвращает Наименование Счета покупателю.
func (invoiceOut InvoiceOut) GetName() string {
	return Deref(invoiceOut.Name)
}

// GetAccountID возвращает ID учётной записи.
func (invoiceOut InvoiceOut) GetAccountID() uuid.UUID {
	return Deref(invoiceOut.AccountID)
}

// GetContract возвращает Метаданные договора.
func (invoiceOut InvoiceOut) GetContract() Contract {
	return Deref(invoiceOut.Contract).getValue()
}

// GetAgent возвращает Метаданные контрагента.
func (invoiceOut InvoiceOut) GetAgent() Agent {
	return Deref(invoiceOut.Agent)
}

// GetOrganization возвращает Метаданные юрлица.
func (invoiceOut InvoiceOut) GetOrganization() Organization {
	return Deref(invoiceOut.Organization)
}

// GetPaymentPlannedMoment возвращает Планируемую дату оплаты.
func (invoiceOut InvoiceOut) GetPaymentPlannedMoment() time.Time {
	return Deref(invoiceOut.PaymentPlannedMoment).Time()
}

// GetPositions возвращает Метаданные позиций Счета покупателю.
func (invoiceOut InvoiceOut) GetPositions() MetaArray[InvoiceOutPosition] {
	return Deref(invoiceOut.Positions)
}

// GetPrinted возвращает true, если документ напечатан.
func (invoiceOut InvoiceOut) GetPrinted() bool {
	return Deref(invoiceOut.Printed)
}

// GetProject возвращает Метаданные проекта.
func (invoiceOut InvoiceOut) GetProject() Project {
	return Deref(invoiceOut.Project).getValue()
}

// GetPublished возвращает true, если документ опубликован.
func (invoiceOut InvoiceOut) GetPublished() bool {
	return Deref(invoiceOut.Published)
}

// GetRate возвращает Валюту.
func (invoiceOut InvoiceOut) GetRate() Rate {
	return Deref(invoiceOut.Rate).getValue()
}

// GetShared возвращает флаг Общего доступа.
func (invoiceOut InvoiceOut) GetShared() bool {
	return Deref(invoiceOut.Shared)
}

// GetShippedSum возвращает Сумму отгруженного.
func (invoiceOut InvoiceOut) GetShippedSum() float64 {
	return Deref(invoiceOut.ShippedSum)
}

// GetState возвращает Метаданные статуса счета.
func (invoiceOut InvoiceOut) GetState() State {
	return Deref(invoiceOut.State).getValue()
}

// GetStore возвращает Метаданные склада.
func (invoiceOut InvoiceOut) GetStore() Store {
	return Deref(invoiceOut.Store).getValue()
}

// GetSum возвращает Сумму Счета покупателю в установленной валюте.
func (invoiceOut InvoiceOut) GetSum() float64 {
	return Deref(invoiceOut.Sum)
}

// GetSyncID возвращает ID синхронизации.
func (invoiceOut InvoiceOut) GetSyncID() uuid.UUID {
	return Deref(invoiceOut.SyncID)
}

// GetUpdated возвращает Момент последнего обновления Счета покупателю.
func (invoiceOut InvoiceOut) GetUpdated() time.Time {
	return Deref(invoiceOut.Updated).Time()
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (invoiceOut InvoiceOut) GetOwner() Employee {
	return Deref(invoiceOut.Owner)
}

// GetVatIncluded возвращает true, если НДС включен в цену.
func (invoiceOut InvoiceOut) GetVatIncluded() bool {
	return Deref(invoiceOut.VatIncluded)
}

// GetVatSum возвращает Сумму НДС.
func (invoiceOut InvoiceOut) GetVatSum() float64 {
	return Deref(invoiceOut.VatSum)
}

// GetCustomerOrder возвращает Ссылку на Заказ Покупателя, с которым связан этот Счет покупателю.
func (invoiceOut InvoiceOut) GetCustomerOrder() CustomerOrder {
	return Deref(invoiceOut.CustomerOrder)
}

// GetPayments возвращает Массив ссылок на связанные платежи.
func (invoiceOut InvoiceOut) GetPayments() Slice[Payment] {
	return invoiceOut.Payments
}

// GetSalesChannel возвращает Метаданные канала продаж.
func (invoiceOut InvoiceOut) GetSalesChannel() SalesChannel {
	return Deref(invoiceOut.SalesChannel)
}

// GetAttributes возвращает Список метаданных доп. полей.
func (invoiceOut InvoiceOut) GetAttributes() Slice[Attribute] {
	return invoiceOut.Attributes
}

// SetVatEnabled устанавливает значение, учитывающее НДС для Заказа Покупателя.
func (invoiceOut *InvoiceOut) SetVatEnabled(vatEnabled bool) *InvoiceOut {
	invoiceOut.VatEnabled = &vatEnabled
	return invoiceOut
}

// SetAgentAccount устанавливает Метаданные счета контрагента.
func (invoiceOut *InvoiceOut) SetAgentAccount(agentAccount *AgentAccount) *InvoiceOut {
	if agentAccount != nil {
		invoiceOut.AgentAccount = agentAccount.Clean()
	}
	return invoiceOut
}

// SetApplicable устанавливает Отметку о проведении.
func (invoiceOut *InvoiceOut) SetApplicable(applicable bool) *InvoiceOut {
	invoiceOut.Applicable = &applicable
	return invoiceOut
}

// SetDemands устанавливает Массив ссылок на связанные отгрузки.
//
// Принимает множество объектов [Demand].
func (invoiceOut *InvoiceOut) SetDemands(demands ...*Demand) *InvoiceOut {
	invoiceOut.Demands.Push(demands...)
	return invoiceOut
}

// SetCode устанавливает Код Счета покупателю.
func (invoiceOut *InvoiceOut) SetCode(code string) *InvoiceOut {
	invoiceOut.Code = &code
	return invoiceOut
}

// SetOrganizationAccount устанавливает Метаданные счета юрлица.
func (invoiceOut *InvoiceOut) SetOrganizationAccount(organizationAccount *AgentAccount) *InvoiceOut {
	if organizationAccount != nil {
		invoiceOut.OrganizationAccount = organizationAccount.Clean()
	}
	return invoiceOut
}

// SetDescription устанавливает Комментарий Счета покупателю.
func (invoiceOut *InvoiceOut) SetDescription(description string) *InvoiceOut {
	invoiceOut.Description = &description
	return invoiceOut
}

// SetExternalCode устанавливает Внешний код Счета покупателю.
func (invoiceOut *InvoiceOut) SetExternalCode(externalCode string) *InvoiceOut {
	invoiceOut.ExternalCode = &externalCode
	return invoiceOut
}

// SetFiles устанавливает Метаданные массива Файлов.
//
// Принимает множество объектов [File].
func (invoiceOut *InvoiceOut) SetFiles(files ...*File) *InvoiceOut {
	invoiceOut.Files = NewMetaArrayFrom(files)
	return invoiceOut
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (invoiceOut *InvoiceOut) SetGroup(group *Group) *InvoiceOut {
	if group != nil {
		invoiceOut.Group = group.Clean()
	}
	return invoiceOut
}

// SetMeta устанавливает Метаданные Счета покупателю.
func (invoiceOut *InvoiceOut) SetMeta(meta *Meta) *InvoiceOut {
	invoiceOut.Meta = meta
	return invoiceOut
}

// SetMoment устанавливает Дату документа.
func (invoiceOut *InvoiceOut) SetMoment(moment time.Time) *InvoiceOut {
	invoiceOut.Moment = NewTimestamp(moment)
	return invoiceOut
}

// SetName устанавливает Наименование Счета покупателю.
func (invoiceOut *InvoiceOut) SetName(name string) *InvoiceOut {
	invoiceOut.Name = &name
	return invoiceOut
}

// SetContract устанавливает Метаданные договора.
//
// Передача nil передаёт сброс значения (null).
func (invoiceOut *InvoiceOut) SetContract(contract *Contract) *InvoiceOut {
	invoiceOut.Contract = NewNullValue(contract)
	return invoiceOut
}

// SetAgent устанавливает Метаданные контрагента [Counterparty] или организации [Organization].
//
// Принимает [Counterparty] или [Organization].
func (invoiceOut *InvoiceOut) SetAgent(agent AgentOrganizationConverter) *InvoiceOut {
	if agent != nil {
		invoiceOut.Agent = agent.AsOrganizationAgent()
	}
	return invoiceOut
}

// SetOrganization устанавливает Метаданные юрлица.
func (invoiceOut *InvoiceOut) SetOrganization(organization *Organization) *InvoiceOut {
	if organization != nil {
		invoiceOut.Organization = organization.Clean()
	}
	return invoiceOut
}

// SetPaymentPlannedMoment устанавливает Планируемую дату оплаты.
func (invoiceOut *InvoiceOut) SetPaymentPlannedMoment(paymentPlannedMoment time.Time) *InvoiceOut {
	invoiceOut.PaymentPlannedMoment = NewTimestamp(paymentPlannedMoment)
	return invoiceOut
}

// SetPositions устанавливает Метаданные позиций Счета покупателю.
//
// Принимает множество объектов [InvoiceOutPosition].
func (invoiceOut *InvoiceOut) SetPositions(positions ...*InvoiceOutPosition) *InvoiceOut {
	invoiceOut.Positions = NewMetaArrayFrom(positions)
	return invoiceOut
}

// SetProject устанавливает Метаданные проекта.
//
// Передача nil передаёт сброс значения (null).
func (invoiceOut *InvoiceOut) SetProject(project *Project) *InvoiceOut {
	invoiceOut.Project = NewNullValue(project)
	return invoiceOut
}

// SetRate устанавливает Валюту.
//
// Передача nil передаёт сброс значения (null).
func (invoiceOut *InvoiceOut) SetRate(rate *Rate) *InvoiceOut {
	invoiceOut.Rate = NewNullValue(rate)
	return invoiceOut
}

// SetShared устанавливает флаг общего доступа.
func (invoiceOut *InvoiceOut) SetShared(shared bool) *InvoiceOut {
	invoiceOut.Shared = &shared
	return invoiceOut
}

// SetState устанавливает Метаданные статуса Счета покупателю.
//
// Передача nil передаёт сброс значения (null).
func (invoiceOut *InvoiceOut) SetState(state *State) *InvoiceOut {
	invoiceOut.State = NewNullValue(state)
	return invoiceOut
}

// SetStore устанавливает Метаданные склада.
//
// Передача nil передаёт сброс значения (null).
func (invoiceOut *InvoiceOut) SetStore(store *Store) *InvoiceOut {
	if store != nil {
		invoiceOut.Store = NewNullValue(store)
	}
	return invoiceOut
}

// SetSyncID устанавливает ID синхронизации.
func (invoiceOut *InvoiceOut) SetSyncID(syncID uuid.UUID) *InvoiceOut {
	invoiceOut.SyncID = &syncID
	return invoiceOut
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (invoiceOut *InvoiceOut) SetOwner(owner *Employee) *InvoiceOut {
	if owner != nil {
		invoiceOut.Owner = owner.Clean()
	}
	return invoiceOut
}

// SetVatIncluded устанавливает флаг включения НДС в цену.
func (invoiceOut *InvoiceOut) SetVatIncluded(vatIncluded bool) *InvoiceOut {
	invoiceOut.VatIncluded = &vatIncluded
	return invoiceOut
}

// SetCustomerOrder устанавливает Ссылку на Заказ Покупателя, с которым связан этот Счет покупателю.
func (invoiceOut *InvoiceOut) SetCustomerOrder(customerOrder *CustomerOrder) *InvoiceOut {
	if customerOrder != nil {
		invoiceOut.CustomerOrder = customerOrder.Clean()
	}
	return invoiceOut
}

// SetPayments устанавливает Метаданные ссылок на связанные платежи.
//
// Принимает множество объектов, реализующих интерфейс [PaymentConverter].
func (invoiceOut *InvoiceOut) SetPayments(payments ...PaymentConverter) *InvoiceOut {
	invoiceOut.Payments = NewPaymentsFrom(payments)
	return invoiceOut
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (invoiceOut *InvoiceOut) SetAttributes(attributes ...*Attribute) *InvoiceOut {
	invoiceOut.Attributes.Push(attributes...)
	return invoiceOut
}

// String реализует интерфейс [fmt.Stringer].
func (invoiceOut InvoiceOut) String() string {
	return Stringify(invoiceOut)
}

// MetaType возвращает код сущности.
func (InvoiceOut) MetaType() MetaType {
	return MetaTypeInvoiceOut
}

// Update shortcut
func (invoiceOut InvoiceOut) Update(ctx context.Context, client *Client, params ...*Params) (*InvoiceOut, *resty.Response, error) {
	return NewInvoiceOutService(client).Update(ctx, invoiceOut.GetID(), &invoiceOut, params...)
}

// Create shortcut
func (invoiceOut InvoiceOut) Create(ctx context.Context, client *Client, params ...*Params) (*InvoiceOut, *resty.Response, error) {
	return NewInvoiceOutService(client).Create(ctx, &invoiceOut, params...)
}

// Delete shortcut
func (invoiceOut InvoiceOut) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewInvoiceOutService(client).Delete(ctx, invoiceOut.GetID())
}

// InvoiceOutPosition Позиция Счета покупателю.
//
// Код сущности: invoiceposition
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-schet-pokupatelu-scheta-pokupatelqm-pozicii-scheta-pokupatelu
type InvoiceOutPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учётной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	Discount   *float64            `json:"discount,omitempty"`   // Процент скидки или наценки. Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID сущности
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Price      *float64            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Vat        *int                `json:"vat,omitempty"`        // НДС, которым облагается текущая позиция
	VatEnabled *bool               `json:"vatEnabled,omitempty"` // Включен ли НДС для позиции. С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
	Stock      *Stock              `json:"stock,omitempty"`      // Остатки и себестоимость позиции (указывается при наличии параметра запроса `fields=stock`)
}

// GetAccountID возвращает ID учётной записи.
func (invoiceOutPosition InvoiceOutPosition) GetAccountID() uuid.UUID {
	return Deref(invoiceOutPosition.AccountID)
}

// GetAssortment возвращает Метаданные товара/услуги/серии/модификации, которую представляет собой позиция.
func (invoiceOutPosition InvoiceOutPosition) GetAssortment() AssortmentPosition {
	return Deref(invoiceOutPosition.Assortment)
}

// GetDiscount возвращает Процент скидки или наценки.
//
// Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%.
func (invoiceOutPosition InvoiceOutPosition) GetDiscount() float64 {
	return Deref(invoiceOutPosition.Discount)
}

// GetID возвращает ID позиции.
func (invoiceOutPosition InvoiceOutPosition) GetID() uuid.UUID {
	return Deref(invoiceOutPosition.ID)
}

// GetPack возвращает Упаковку Товара.
func (invoiceOutPosition InvoiceOutPosition) GetPack() Pack {
	return Deref(invoiceOutPosition.Pack)
}

// GetPrice возвращает Цену товара/услуги в копейках.
func (invoiceOutPosition InvoiceOutPosition) GetPrice() float64 {
	return Deref(invoiceOutPosition.Price)
}

// GetQuantity возвращает Количество товаров данного вида в позиции.
func (invoiceOutPosition InvoiceOutPosition) GetQuantity() float64 {
	return Deref(invoiceOutPosition.Quantity)
}

// GetVat возвращает НДС, которым облагается текущая позиция.
func (invoiceOutPosition InvoiceOutPosition) GetVat() int {
	return Deref(invoiceOutPosition.Vat)
}

// GetVatEnabled возвращает true, если учитывается НДС.
func (invoiceOutPosition InvoiceOutPosition) GetVatEnabled() bool {
	return Deref(invoiceOutPosition.VatEnabled)
}

// GetStock возвращает Остатки и себестоимость позиции (указывается при наличии параметра запроса `fields=stock`).
func (invoiceOutPosition InvoiceOutPosition) GetStock() Stock {
	return Deref(invoiceOutPosition.Stock)
}

// SetAssortment устанавливает Метаданные товара/услуги, которую представляет собой компонент.
//
// Принимает объект, реализующий интерфейс [AssortmentConverter].
func (invoiceOutPosition *InvoiceOutPosition) SetAssortment(assortment AssortmentConverter) *InvoiceOutPosition {
	if assortment != nil {
		invoiceOutPosition.Assortment = assortment.AsAssortment()
	}
	return invoiceOutPosition
}

// SetDiscount устанавливает Процент скидки или наценки.
//
// Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%.
func (invoiceOutPosition *InvoiceOutPosition) SetDiscount(discount float64) *InvoiceOutPosition {
	invoiceOutPosition.Discount = &discount
	return invoiceOutPosition
}

// SetPack устанавливает Упаковку Товара.
func (invoiceOutPosition *InvoiceOutPosition) SetPack(pack *Pack) *InvoiceOutPosition {
	if pack != nil {
		invoiceOutPosition.Pack = pack
	}
	return invoiceOutPosition
}

// SetPrice устанавливает Цену товара/услуги в копейках.
func (invoiceOutPosition *InvoiceOutPosition) SetPrice(price float64) *InvoiceOutPosition {
	invoiceOutPosition.Price = &price
	return invoiceOutPosition
}

// SetQuantity устанавливает Количество товаров данного вида в позиции.
func (invoiceOutPosition *InvoiceOutPosition) SetQuantity(quantity float64) *InvoiceOutPosition {
	invoiceOutPosition.Quantity = &quantity
	return invoiceOutPosition
}

// SetVat устанавливает НДС, которым облагается текущая позиция.
func (invoiceOutPosition *InvoiceOutPosition) SetVat(vat int) *InvoiceOutPosition {
	invoiceOutPosition.Vat = &vat
	return invoiceOutPosition
}

// SetVatEnabled устанавливает значение, учитывающее НДС для текущей позиции.
func (invoiceOutPosition *InvoiceOutPosition) SetVatEnabled(vatEnabled bool) *InvoiceOutPosition {
	invoiceOutPosition.VatEnabled = &vatEnabled
	return invoiceOutPosition
}

// String реализует интерфейс [fmt.Stringer].
func (invoiceOutPosition InvoiceOutPosition) String() string {
	return Stringify(invoiceOutPosition)
}

// MetaType возвращает код сущности.
func (InvoiceOutPosition) MetaType() MetaType {
	return MetaTypeInvoicePosition
}

// InvoiceOutService описывает методы сервиса для работы со счетами покупателей.
type InvoiceOutService interface {
	// GetList выполняет запрос на получение списка счетов покупателям.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[InvoiceOut], *resty.Response, error)

	// Create выполняет запрос на создание счета покупателю.
	// Обязательные поля для заполнения:
	//	- name (Номер Счета покупателю)
	//	- organization (Ссылка на ваше юрлицо)
	//	- agent (Ссылка на контрагента (покупателя))
	// Принимает контекст, счет покупателю и опционально объект параметров запроса Params.
	// Возвращает созданный счет покупателю.
	Create(ctx context.Context, invoiceOut *InvoiceOut, params ...*Params) (*InvoiceOut, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение счетов покупателям.
	// Изменяемые счета покупателям должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список счетов покупателям и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых счетов покупателям.
	CreateUpdateMany(ctx context.Context, invoiceOutList Slice[InvoiceOut], params ...*Params) (*Slice[InvoiceOut], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление счетов покупателям.
	// Принимает контекст и множество счетов покупателям.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*InvoiceOut) (*DeleteManyResponse, *resty.Response, error)

	// Delete выполняет запрос на удаление счета покупателю.
	// Принимает контекст и ID счета покупателю.
	// Возвращает «true» в случае успешного удаления счета покупателю.
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельного счета покупателю по ID.
	// Принимает контекст, ID счета покупателю и опционально объект параметров запроса Params.
	// Возвращает найденный счет покупателю.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*InvoiceOut, *resty.Response, error)

	// Update выполняет запрос на изменение счета покупателю.
	// Принимает контекст, счет покупателю и опционально объект параметров запроса Params.
	// Возвращает изменённый счет покупателю.
	Update(ctx context.Context, id uuid.UUID, invoiceOut *InvoiceOut, params ...*Params) (*InvoiceOut, *resty.Response, error)

	// Template выполняет запрос на получение предзаполненного счета покупателю со стандартными полями без связи с какими-либо другими документами.
	// Принимает контекст.
	// Возвращает предзаполненный счет покупателю.
	Template(ctx context.Context) (*InvoiceOut, *resty.Response, error)

	// TemplateBased выполняет запрос на получение шаблона счета покупателю на основе других документов.
	// Основание, на котором может быть создан:
	//	- Заказ покупателя (CustomerOrder)
	//	- Отгрузка (Demand)
	// Принимает контекст и множество документов из списка выше.
	// Возвращает предзаполненный счет покупателю на основании переданных документов.
	TemplateBased(ctx context.Context, basedOn ...MetaOwner) (*InvoiceOut, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных счетов покупателям.
	// Принимает контекст.
	// Возвращает объект метаданных MetaAttributesStatesSharedWrapper.
	GetMetadata(ctx context.Context) (*MetaAttributesStatesSharedWrapper, *resty.Response, error)

	// GetPositionList выполняет запрос на получение списка позиций документа.
	// Принимает контекст, ID документа и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetPositionList(ctx context.Context, id uuid.UUID, params ...*Params) (*List[InvoiceOutPosition], *resty.Response, error)

	// GetPositionByID выполняет запрос на получение отдельной позиции документа по ID.
	// Принимает контекст, ID документа, ID позиции и опционально объект параметров запроса Params.
	// Возвращает найденную позицию.
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params ...*Params) (*InvoiceOutPosition, *resty.Response, error)

	// UpdatePosition выполняет запрос на изменение позиции документа.
	// Принимает контекст, ID документа, ID позиции, позицию документа и опционально объект параметров запроса Params.
	// Возвращает изменённую позицию.
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *InvoiceOutPosition, params ...*Params) (*InvoiceOutPosition, *resty.Response, error)

	// CreatePosition выполняет запрос на добавление позиции документа.
	// Принимает контекст, ID документа, позицию документа и опционально объект параметров запроса Params.
	// Возвращает добавленную позицию.
	CreatePosition(ctx context.Context, id uuid.UUID, position *InvoiceOutPosition, params ...*Params) (*InvoiceOutPosition, *resty.Response, error)

	// CreatePositionMany выполняет запрос на массовое добавление позиций документа.
	// Принимает контекст, ID документа и множество позиций.
	// Возвращает список добавленных позиций.
	CreatePositionMany(ctx context.Context, id uuid.UUID, positions ...*InvoiceOutPosition) (*Slice[InvoiceOutPosition], *resty.Response, error)

	// DeletePosition выполняет запрос на удаление позиции документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает «true» в случае успешного удаления позиции.
	DeletePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (bool, *resty.Response, error)

	// DeletePositionMany выполняет запрос на массовое удаление позиций документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeletePositionMany(ctx context.Context, id uuid.UUID, positions ...*InvoiceOutPosition) (*DeleteManyResponse, *resty.Response, error)

	// GetPositionTrackingCodeList выполняет запрос на получение кодов маркировки позиции документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает объект List.
	GetPositionTrackingCodeList(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (*List[TrackingCode], *resty.Response, error)

	// CreateUpdatePositionTrackingCodeMany выполняет запрос на массовое создание/изменение кодов маркировки позиции документа.
	// Принимает контекст, ID документа, ID позиции и множество кодов маркировки.
	// Возвращает список созданных и/или изменённых кодов маркировки позиции документа.
	CreateUpdatePositionTrackingCodeMany(ctx context.Context, id uuid.UUID, positionID uuid.UUID, trackingCodes ...*TrackingCode) (*Slice[TrackingCode], *resty.Response, error)

	// DeletePositionTrackingCodeMany выполняет запрос на массовое удаление кодов маркировки позиции документа.
	// Принимает контекст, ID документа, ID позиции и множество кодов маркировки.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeletePositionTrackingCodeMany(ctx context.Context, id uuid.UUID, positionID uuid.UUID, trackingCodes ...*TrackingCode) (*DeleteManyResponse, *resty.Response, error)

	// GetAttributeList выполняет запрос на получение списка доп полей.
	// Принимает контекст.
	// Возвращает объект List.
	GetAttributeList(ctx context.Context) (*List[Attribute], *resty.Response, error)

	// GetAttributeByID выполняет запрос на получение отдельного доп поля по ID.
	// Принимает контекст и ID доп поля.
	// Возвращает найденное доп поле.
	GetAttributeByID(ctx context.Context, id uuid.UUID) (*Attribute, *resty.Response, error)

	// CreateAttribute выполняет запрос на создание доп поля.
	// Принимает контекст и доп поле.
	// Возвращает созданное доп поле.
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)

	// CreateUpdateAttributeMany выполняет запрос на массовое создание и/или изменение доп полей.
	// Изменяемые доп поля должны содержать идентификатор в виде метаданных.
	// Принимает контекст и множество доп полей.
	// Возвращает список созданных и/или изменённых доп полей.
	CreateUpdateAttributeMany(ctx context.Context, attributes ...*Attribute) (*Slice[Attribute], *resty.Response, error)

	// UpdateAttribute выполняет запрос на изменения доп поля.
	// Принимает контекст, ID доп поля и доп поле.
	// Возвращает изменённое доп поле.
	UpdateAttribute(ctx context.Context, id uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)

	// DeleteAttribute выполняет запрос на удаление доп поля.
	// Принимает контекст и ID доп поля.
	// Возвращает «true» в случае успешного удаления доп поля.
	DeleteAttribute(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// DeleteAttributeMany выполняет запрос на массовое удаление доп полей.
	// Принимает контекст и множество доп полей.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteAttributeMany(ctx context.Context, attributes ...*Attribute) (*DeleteManyResponse, *resty.Response, error)

	// GetPublicationList выполняет запрос на получение списка публикаций.
	// Принимает контекст и ID документа.
	// Возвращает объект List.
	GetPublicationList(ctx context.Context, id uuid.UUID) (*List[Publication], *resty.Response, error)

	// GetPublicationByID выполняет запрос на получение отдельной публикации по ID.
	// Принимает контекст, ID документа и ID публикации.
	// Возвращает найденную публикацию.
	GetPublicationByID(ctx context.Context, id uuid.UUID, publicationID uuid.UUID) (*Publication, *resty.Response, error)

	// Publish выполняет запрос на создание публикации.
	// Принимает контекст, ID документа и шаблон (CustomTemplate или EmbeddedTemplate)
	// Возвращает созданную публикацию.
	Publish(ctx context.Context, id uuid.UUID, template TemplateConverter) (*Publication, *resty.Response, error)

	// DeletePublication выполняет запрос на удаление публикации.
	// Принимает контекст, ID документа и ID публикации.
	// Возвращает «true» в случае успешного удаления публикации.
	DeletePublication(ctx context.Context, id uuid.UUID, publicationID uuid.UUID) (bool, *resty.Response, error)

	// GetBySyncID выполняет запрос на получение отдельного документа по syncID.
	// Принимает контекст и syncID документа.
	// Возвращает найденный документ.
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*InvoiceOut, *resty.Response, error)

	// DeleteBySyncID выполняет запрос на удаление документа по syncID.
	// Принимает контекст и syncID документа.
	// Возвращает «true» в случае успешного удаления документа.
	DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error)

	// MoveToTrash выполняет запрос на перемещение документа с указанным ID в корзину.
	// Принимает контекст и ID документа.
	// Возвращает «true» в случае успешного перемещения в корзину.
	MoveToTrash(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetStateByID выполняет запрос на получение статуса документа по ID.
	// Принимает контекст и ID статуса.
	// Возвращает найденный статус.
	GetStateByID(ctx context.Context, id uuid.UUID) (*State, *resty.Response, error)

	// CreateState выполняет запрос на создание статуса документа.
	// Принимает контекст и статус.
	// Возвращает созданный статус.
	CreateState(ctx context.Context, state *State) (*State, *resty.Response, error)

	// UpdateState выполняет запрос на изменение статуса документа.
	// Принимает контекст, ID статуса и статус.
	// Возвращает изменённый статус.
	UpdateState(ctx context.Context, id uuid.UUID, state *State) (*State, *resty.Response, error)

	// CreateUpdateStateMany выполняет запрос на массовое создание и/или изменение статусов документа.
	// Принимает контекст и множество статусов.
	// Возвращает список созданных и/или изменённых статусов.
	CreateUpdateStateMany(ctx context.Context, states ...*State) (*Slice[State], *resty.Response, error)

	// DeleteState выполняет запрос на удаление статуса документа.
	// Принимает контекст и ID статуса.
	// Возвращает «true» в случае успешного удаления статуса.
	DeleteState(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetFileList выполняет запрос на получение файлов в виде списка.
	// Принимает контекст и ID сущности/документа.
	// Возвращает объект List.
	GetFileList(ctx context.Context, id uuid.UUID) (*List[File], *resty.Response, error)

	// CreateFile выполняет запрос на добавление файла.
	// Принимает контекст, ID сущности/документа и файл.
	// Возвращает список файлов.
	CreateFile(ctx context.Context, id uuid.UUID, file *File) (*Slice[File], *resty.Response, error)

	// UpdateFileMany выполняет запрос на массовое создание и/или изменение файлов сущности/документа.
	// Принимает контекст, ID сущности/документа и множество файлов.
	// Возвращает созданных и/или изменённых файлов.
	UpdateFileMany(ctx context.Context, id uuid.UUID, files ...*File) (*Slice[File], *resty.Response, error)

	// DeleteFile выполняет запрос на удаление файла сущности/документа.
	// Принимает контекст, ID сущности/документа и ID файла.
	// Возвращает «true» в случае успешного удаления файла.
	DeleteFile(ctx context.Context, id uuid.UUID, fileID uuid.UUID) (bool, *resty.Response, error)

	// DeleteFileMany выполняет запрос на массовое удаление файлов сущности/документа.
	// Принимает контекст, ID сущности/документа и множество файлов.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteFileMany(ctx context.Context, id uuid.UUID, files ...*File) (*DeleteManyResponse, *resty.Response, error)

	// Evaluate выполняет запрос на получение шаблона документа с автозаполнением.
	// Принимает контекст, документ и множество значений Evaluate.
	//
	// Возможные значения типа Evaluate:
	//	- EvaluateDiscount – скидки
	//	- EvaluatePrice    – цены
	//	- EvaluateVat      – ндс
	//	- EvaluateCost     – себестоимость
	// Возвращает шаблон документа с автозаполнением.
	Evaluate(ctx context.Context, entity *InvoiceOut, evaluate ...Evaluate) (*InvoiceOut, *resty.Response, error)
}

const (
	EndpointInvoiceOut = EndpointEntity + string(MetaTypeInvoiceOut)
)

// NewInvoiceOutService принимает [Client] и возвращает сервис для работы со счетами покупателей.
func NewInvoiceOutService(client *Client) InvoiceOutService {
	return newMainService[InvoiceOut, InvoiceOutPosition, MetaAttributesStatesSharedWrapper, any](client, EndpointInvoiceOut)
}
