package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// InvoiceIn Счет поставщика.
//
// Код сущности: invoicein
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-schet-postawschika
type InvoiceIn struct {
	OrganizationAccount  *AgentAccount                 `json:"organizationAccount,omitempty"`  // Метаданные счета юрлица
	Created              *Timestamp                    `json:"created,omitempty"`              // Дата создания
	PayedSum             *float64                      `json:"payedSum,omitempty"`             // Сумма входящих платежей по Счету поставщика
	Applicable           *bool                         `json:"applicable,omitempty"`           // Отметка о проведении
	Supplies             Slice[Supply]                 `json:"supplies,omitempty"`             // Ссылки на связанные приемки
	Code                 *string                       `json:"code,omitempty"`                 // Код Счета поставщика
	Contract             *NullValue[Contract]          `json:"contract,omitempty"`             // Метаданные договора
	Owner                *Employee                     `json:"owner,omitempty"`                // Метаданные владельца (Сотрудника)
	Deleted              *Timestamp                    `json:"deleted,omitempty"`              // Момент последнего удаления Счета поставщика
	Description          *string                       `json:"description,omitempty"`          // Комментарий Счета поставщика
	ExternalCode         *string                       `json:"externalCode,omitempty"`         // Внешний код Счета поставщика
	Files                *MetaArray[File]              `json:"files,omitempty"`                // Метаданные массива Файлов (Максимальное количество файлов - 100)
	AccountID            *uuid.UUID                    `json:"accountId,omitempty"`            // ID учётной записи
	ID                   *uuid.UUID                    `json:"id,omitempty"`                   // ID Счета поставщика
	IncomingDate         *Timestamp                    `json:"incomingDate,omitempty"`         // Входящая дата
	IncomingNumber       *string                       `json:"incomingNumber,omitempty"`       // Входящий номер
	Meta                 *Meta                         `json:"meta,omitempty"`                 // Метаданные Счета поставщика
	Moment               *Timestamp                    `json:"moment,omitempty"`               // Дата документа
	Name                 *string                       `json:"name,omitempty"`                 // Наименование Счета поставщика
	Organization         *Organization                 `json:"organization,omitempty"`         // Метаданные юрлица
	Group                *Group                        `json:"group,omitempty"`                // Отдел сотрудника
	Agent                *Agent                        `json:"agent,omitempty"`                // Метаданные контрагента
	AgentAccount         *AgentAccount                 `json:"agentAccount,omitempty"`         // Метаданные счета контрагента
	PaymentPlannedMoment *Timestamp                    `json:"paymentPlannedMoment,omitempty"` // Планируемая дата оплаты
	Positions            *MetaArray[InvoiceInPosition] `json:"positions,omitempty"`            // Метаданные позиций Счета поставщика
	Printed              *bool                         `json:"printed,omitempty"`              // Напечатан ли документ
	Project              *NullValue[Project]           `json:"project,omitempty"`              // Метаданные проекта
	Published            *bool                         `json:"published,omitempty"`            // Опубликован ли документ
	Rate                 *NullValue[Rate]              `json:"rate,omitempty"`                 // Валюта
	Shared               *bool                         `json:"shared,omitempty"`               // Общий доступ
	ShippedSum           *float64                      `json:"shippedSum,omitempty"`           // Сумма отгруженного
	State                *NullValue[State]             `json:"state,omitempty"`                // Метаданные статуса счета поставщика
	Store                *NullValue[Store]             `json:"store,omitempty"`                // Метаданные склада
	Sum                  *float64                      `json:"sum,omitempty"`                  // Сумма Счета в установленной валюте
	SyncID               *uuid.UUID                    `json:"syncId,omitempty"`               // ID синхронизации
	Updated              *Timestamp                    `json:"updated,omitempty"`              // Момент последнего обновления Счета поставщика
	VatEnabled           *bool                         `json:"vatEnabled,omitempty"`           // Учитывается ли НДС
	VatIncluded          *bool                         `json:"vatIncluded,omitempty"`          // Включен ли НДС в цену
	VatSum               *float64                      `json:"vatSum,omitempty"`               // Сумма НДС
	Payments             Slice[Payment]                `json:"payments,omitempty"`             // Массив ссылок на связанные операции
	PurchaseOrder        *PurchaseOrder                `json:"purchaseOrder,omitempty"`        // Ссылка на связанный заказ поставщику
	Attributes           Slice[Attribute]              `json:"attributes,omitempty"`           // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (invoiceIn InvoiceIn) Clean() *InvoiceIn {
	if invoiceIn.Meta == nil {
		return nil
	}
	return &InvoiceIn{Meta: invoiceIn.Meta}
}

// AsOperation возвращает объект [Operation] c полями meta и linkedSum.
//
// Значение поля linkedSum заполняется из поля sum.
func (invoiceIn InvoiceIn) AsOperation() *Operation {
	return &Operation{Meta: invoiceIn.GetMeta(), LinkedSum: invoiceIn.GetSum()}
}

// asTaskOperation реализует интерфейс [TaskOperationInterface].
func (invoiceIn InvoiceIn) asTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: invoiceIn.Meta}
}

// GetOrganizationAccount возвращает Метаданные счета юрлица.
func (invoiceIn InvoiceIn) GetOrganizationAccount() AgentAccount {
	return Deref(invoiceIn.OrganizationAccount)
}

// GetCreated возвращает Дату создания.
func (invoiceIn InvoiceIn) GetCreated() Timestamp {
	return Deref(invoiceIn.Created)
}

// GetPayedSum возвращает Сумму входящих платежей по Счету поставщика.
func (invoiceIn InvoiceIn) GetPayedSum() float64 {
	return Deref(invoiceIn.PayedSum)
}

// GetApplicable возвращает Отметку о проведении.
func (invoiceIn InvoiceIn) GetApplicable() bool {
	return Deref(invoiceIn.Applicable)
}

// GetSupplies возвращает Массив ссылок на связанные приемки.
func (invoiceIn InvoiceIn) GetSupplies() Slice[Supply] {
	return invoiceIn.Supplies
}

// GetCode возвращает Код Счета поставщика.
func (invoiceIn InvoiceIn) GetCode() string {
	return Deref(invoiceIn.Code)
}

// GetContract возвращает Метаданные договора.
func (invoiceIn InvoiceIn) GetContract() Contract {
	return invoiceIn.Contract.GetValue()
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (invoiceIn InvoiceIn) GetOwner() Employee {
	return Deref(invoiceIn.Owner)
}

// GetDeleted возвращает Момент последнего удаления Счета поставщика.
func (invoiceIn InvoiceIn) GetDeleted() Timestamp {
	return Deref(invoiceIn.Deleted)
}

// GetDescription возвращает Комментарий Счета поставщика.
func (invoiceIn InvoiceIn) GetDescription() string {
	return Deref(invoiceIn.Description)
}

// GetExternalCode возвращает Внешний код Счета поставщика.
func (invoiceIn InvoiceIn) GetExternalCode() string {
	return Deref(invoiceIn.ExternalCode)
}

// GetFiles возвращает Метаданные массива Файлов.
func (invoiceIn InvoiceIn) GetFiles() MetaArray[File] {
	return Deref(invoiceIn.Files)
}

// GetAccountID возвращает ID учётной записи.
func (invoiceIn InvoiceIn) GetAccountID() uuid.UUID {
	return Deref(invoiceIn.AccountID)
}

// GetID возвращает ID Счета поставщика.
func (invoiceIn InvoiceIn) GetID() uuid.UUID {
	return Deref(invoiceIn.ID)
}

// GetIncomingDate возвращает Входящую дату.
func (invoiceIn InvoiceIn) GetIncomingDate() Timestamp {
	return Deref(invoiceIn.IncomingDate)
}

// GetIncomingNumber возвращает Входящий номер.
func (invoiceIn InvoiceIn) GetIncomingNumber() string {
	return Deref(invoiceIn.IncomingNumber)
}

// GetMeta возвращает Метаданные Счета поставщика.
func (invoiceIn InvoiceIn) GetMeta() Meta {
	return Deref(invoiceIn.Meta)
}

// GetMoment возвращает Дату документа.
func (invoiceIn InvoiceIn) GetMoment() Timestamp {
	return Deref(invoiceIn.Moment)
}

// GetName возвращает Наименование Счета поставщика.
func (invoiceIn InvoiceIn) GetName() string {
	return Deref(invoiceIn.Name)
}

// GetOrganization возвращает Метаданные юрлица.
func (invoiceIn InvoiceIn) GetOrganization() Organization {
	return Deref(invoiceIn.Organization)
}

// GetGroup возвращает Отдел сотрудника.
func (invoiceIn InvoiceIn) GetGroup() Group {
	return Deref(invoiceIn.Group)
}

// GetAgent возвращает Метаданные Контрагента.
func (invoiceIn InvoiceIn) GetAgent() Agent {
	return Deref(invoiceIn.Agent)
}

// GetAgentAccount возвращает Метаданные счета контрагента.
func (invoiceIn InvoiceIn) GetAgentAccount() AgentAccount {
	return Deref(invoiceIn.AgentAccount)
}

// GetPaymentPlannedMoment возвращает Планируемую дату оплаты.
func (invoiceIn InvoiceIn) GetPaymentPlannedMoment() Timestamp {
	return Deref(invoiceIn.PaymentPlannedMoment)
}

// GetPositions возвращает Метаданные позиций Счета поставщика.
func (invoiceIn InvoiceIn) GetPositions() MetaArray[InvoiceInPosition] {
	return Deref(invoiceIn.Positions)
}

// GetPrinted возвращает true, если документ напечатан.
func (invoiceIn InvoiceIn) GetPrinted() bool {
	return Deref(invoiceIn.Printed)
}

// GetProject возвращает Метаданные проекта.
func (invoiceIn InvoiceIn) GetProject() Project {
	return invoiceIn.Project.GetValue()
}

// GetPublished возвращает true, если документ опубликован.
func (invoiceIn InvoiceIn) GetPublished() bool {
	return Deref(invoiceIn.Published)
}

// GetRate возвращает Валюту.
func (invoiceIn InvoiceIn) GetRate() Rate {
	return invoiceIn.Rate.GetValue()
}

// GetShared возвращает флаг Общего доступа.
func (invoiceIn InvoiceIn) GetShared() bool {
	return Deref(invoiceIn.Shared)
}

// GetShippedSum возвращает Сумму отгруженного.
func (invoiceIn InvoiceIn) GetShippedSum() float64 {
	return Deref(invoiceIn.ShippedSum)
}

// GetState возвращает Метаданные статуса счета поставщика.
func (invoiceIn InvoiceIn) GetState() State {
	return invoiceIn.State.GetValue()
}

// GetStore возвращает Метаданные склада.
func (invoiceIn InvoiceIn) GetStore() Store {
	return invoiceIn.Store.GetValue()
}

// GetSum возвращает Сумму Счета поставщика в установленной валюте.
func (invoiceIn InvoiceIn) GetSum() float64 {
	return Deref(invoiceIn.Sum)
}

// GetSyncID возвращает ID синхронизации.
func (invoiceIn InvoiceIn) GetSyncID() uuid.UUID {
	return Deref(invoiceIn.SyncID)
}

// GetUpdated возвращает Момент последнего обновления Счета поставщика.
func (invoiceIn InvoiceIn) GetUpdated() Timestamp {
	return Deref(invoiceIn.Updated)
}

// GetVatEnabled возвращает true, если учитывается НДС.
func (invoiceIn InvoiceIn) GetVatEnabled() bool {
	return Deref(invoiceIn.VatEnabled)
}

// GetVatIncluded возвращает true, если НДС включен в цену.
func (invoiceIn InvoiceIn) GetVatIncluded() bool {
	return Deref(invoiceIn.VatIncluded)
}

// GetVatSum возвращает Сумму НДС.
func (invoiceIn InvoiceIn) GetVatSum() float64 {
	return Deref(invoiceIn.VatSum)
}

// GetPayments возвращает Массив ссылок на связанные платежи.
func (invoiceIn InvoiceIn) GetPayments() Slice[Payment] {
	return invoiceIn.Payments
}

// GetPurchaseOrder возвращает Ссылку на связанный заказ поставщику.
func (invoiceIn InvoiceIn) GetPurchaseOrder() PurchaseOrder {
	return Deref(invoiceIn.PurchaseOrder)
}

// GetAttributes возвращает Список метаданных доп. полей.
func (invoiceIn InvoiceIn) GetAttributes() Slice[Attribute] {
	return invoiceIn.Attributes
}

// SetOrganizationAccount устанавливает Метаданные счета юрлица.
func (invoiceIn *InvoiceIn) SetOrganizationAccount(organizationAccount *AgentAccount) *InvoiceIn {
	if organizationAccount != nil {
		invoiceIn.OrganizationAccount = organizationAccount.Clean()
	}
	return invoiceIn
}

// SetApplicable устанавливает Отметку о проведении.
func (invoiceIn *InvoiceIn) SetApplicable(applicable bool) *InvoiceIn {
	invoiceIn.Applicable = &applicable
	return invoiceIn
}

// SetSupplies устанавливает Массив ссылок на связанные приемки.
//
// Принимает множество объектов [Supply].
func (invoiceIn *InvoiceIn) SetSupplies(supplies ...*Supply) *InvoiceIn {
	invoiceIn.Supplies.Push(supplies...)
	return invoiceIn
}

// SetCode устанавливает Код Счета поставщика.
func (invoiceIn *InvoiceIn) SetCode(code string) *InvoiceIn {
	invoiceIn.Code = &code
	return invoiceIn
}

// SetContract устанавливает Метаданные договора.
//
// Передача nil передаёт сброс значения (null).
func (invoiceIn *InvoiceIn) SetContract(contract *Contract) *InvoiceIn {
	invoiceIn.Contract = NewNullValue(contract)
	return invoiceIn
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (invoiceIn *InvoiceIn) SetOwner(owner *Employee) *InvoiceIn {
	if owner != nil {
		invoiceIn.Owner = owner.Clean()
	}
	return invoiceIn
}

// SetDescription устанавливает Комментарий Счета поставщика.
func (invoiceIn *InvoiceIn) SetDescription(description string) *InvoiceIn {
	invoiceIn.Description = &description
	return invoiceIn
}

// SetExternalCode устанавливает Внешний код Счета поставщика
func (invoiceIn *InvoiceIn) SetExternalCode(externalCode string) *InvoiceIn {
	invoiceIn.ExternalCode = &externalCode
	return invoiceIn
}

// SetFiles устанавливает Метаданные массива Файлов.
//
// Принимает множество объектов [File].
func (invoiceIn *InvoiceIn) SetFiles(files ...*File) *InvoiceIn {
	invoiceIn.Files = NewMetaArrayFrom(files)
	return invoiceIn
}

// SetIncomingDate устанавливает Входящую дату.
func (invoiceIn *InvoiceIn) SetIncomingDate(incomingDate *Timestamp) *InvoiceIn {
	invoiceIn.IncomingDate = incomingDate
	return invoiceIn
}

// SetIncomingNumber устанавливает Входящий номер.
func (invoiceIn *InvoiceIn) SetIncomingNumber(incomingNumber string) *InvoiceIn {
	invoiceIn.IncomingNumber = &incomingNumber
	return invoiceIn
}

// SetMeta устанавливает Метаданные Счета поставщика.
func (invoiceIn *InvoiceIn) SetMeta(meta *Meta) *InvoiceIn {
	invoiceIn.Meta = meta
	return invoiceIn
}

// SetMoment устанавливает Дату документа.
func (invoiceIn *InvoiceIn) SetMoment(moment *Timestamp) *InvoiceIn {
	invoiceIn.Moment = moment
	return invoiceIn
}

// SetName устанавливает Наименование Счета поставщика.
func (invoiceIn *InvoiceIn) SetName(name string) *InvoiceIn {
	invoiceIn.Name = &name
	return invoiceIn
}

// SetOrganization устанавливает Метаданные юрлица.
func (invoiceIn *InvoiceIn) SetOrganization(organization *Organization) *InvoiceIn {
	if organization != nil {
		invoiceIn.Organization = organization.Clean()
	}
	return invoiceIn
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (invoiceIn *InvoiceIn) SetGroup(group *Group) *InvoiceIn {
	if group != nil {
		invoiceIn.Group = group.Clean()
	}
	return invoiceIn
}

// SetAgent устанавливает Метаданные Контрагента.
//
// Принимает [Counterparty] или [Organization].
func (invoiceIn *InvoiceIn) SetAgent(agent AgentCounterpartyOrganizationInterface) *InvoiceIn {
	if agent != nil {
		invoiceIn.Agent = agent.asCOAgent()
	}
	return invoiceIn
}

// SetAgentAccount устанавливает Метаданные счета контрагента.
func (invoiceIn *InvoiceIn) SetAgentAccount(agentAccount *AgentAccount) *InvoiceIn {
	if agentAccount != nil {
		invoiceIn.AgentAccount = agentAccount.Clean()
	}
	return invoiceIn
}

// SetPaymentPlannedMoment устанавливает Планируемую дату оплаты.
func (invoiceIn *InvoiceIn) SetPaymentPlannedMoment(paymentPlannedMoment *Timestamp) *InvoiceIn {
	invoiceIn.PaymentPlannedMoment = paymentPlannedMoment
	return invoiceIn
}

// SetPositions устанавливает Метаданные позиций Счета поставщика.
//
// Принимает множество объектов [InvoiceInPosition].
func (invoiceIn *InvoiceIn) SetPositions(positions ...*InvoiceInPosition) *InvoiceIn {
	invoiceIn.Positions = NewMetaArrayFrom(positions)
	return invoiceIn
}

// SetProject устанавливает Метаданные проекта.
//
// Передача nil передаёт сброс значения (null).
func (invoiceIn *InvoiceIn) SetProject(project *Project) *InvoiceIn {
	invoiceIn.Project = NewNullValue(project)
	return invoiceIn
}

// SetRate устанавливает Валюту.
//
// Передача nil передаёт сброс значения (null).
func (invoiceIn *InvoiceIn) SetRate(rate *Rate) *InvoiceIn {
	invoiceIn.Rate = NewNullValue(rate)
	return invoiceIn
}

// SetShared устанавливает флаг общего доступа.
func (invoiceIn *InvoiceIn) SetShared(shared bool) *InvoiceIn {
	invoiceIn.Shared = &shared
	return invoiceIn
}

// SetState устанавливает Метаданные статуса счета поставщика.
//
// Передача nil передаёт сброс значения (null).
func (invoiceIn *InvoiceIn) SetState(state *State) *InvoiceIn {
	invoiceIn.State = NewNullValue(state)
	return invoiceIn
}

// SetStore устанавливает Метаданные склада.
//
// Передача nil передаёт сброс значения (null).
func (invoiceIn *InvoiceIn) SetStore(store *Store) *InvoiceIn {
	invoiceIn.Store = NewNullValue(store)
	return invoiceIn
}

// SetSyncID устанавливает ID синхронизации.
func (invoiceIn *InvoiceIn) SetSyncID(syncID uuid.UUID) *InvoiceIn {
	invoiceIn.SyncID = &syncID
	return invoiceIn
}

// SetVatEnabled устанавливает значение, учитывающее НДС для счета поставщика.
func (invoiceIn *InvoiceIn) SetVatEnabled(vatEnabled bool) *InvoiceIn {
	invoiceIn.VatEnabled = &vatEnabled
	return invoiceIn
}

// SetVatIncluded устанавливает флаг включения НДС в цену.
func (invoiceIn *InvoiceIn) SetVatIncluded(vatIncluded bool) *InvoiceIn {
	invoiceIn.VatIncluded = &vatIncluded
	return invoiceIn
}

// SetPayments устанавливает Метаданные ссылок на связанные платежи.
//
// Принимает множество объектов, реализующих интерфейс [AsPaymentInterface].
func (invoiceIn *InvoiceIn) SetPayments(payments ...AsPaymentInterface) *InvoiceIn {
	invoiceIn.Payments = NewPaymentsFrom(payments)
	return invoiceIn
}

// SetPurchaseOrder устанавливает Ссылку на связанный заказ поставщику.
func (invoiceIn *InvoiceIn) SetPurchaseOrder(purchaseOrder *PurchaseOrder) *InvoiceIn {
	if purchaseOrder != nil {
		invoiceIn.PurchaseOrder = purchaseOrder.Clean()
	}
	return invoiceIn
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (invoiceIn *InvoiceIn) SetAttributes(attributes ...*Attribute) *InvoiceIn {
	invoiceIn.Attributes.Push(attributes...)
	return invoiceIn
}

// String реализует интерфейс [fmt.Stringer].
func (invoiceIn InvoiceIn) String() string {
	return Stringify(invoiceIn)
}

// MetaType возвращает код сущности.
func (InvoiceIn) MetaType() MetaType {
	return MetaTypeInvoiceIn
}

// Update shortcut
func (invoiceIn InvoiceIn) Update(ctx context.Context, client *Client, params ...*Params) (*InvoiceIn, *resty.Response, error) {
	return NewInvoiceInService(client).Update(ctx, invoiceIn.GetID(), &invoiceIn, params...)
}

// Create shortcut
func (invoiceIn InvoiceIn) Create(ctx context.Context, client *Client, params ...*Params) (*InvoiceIn, *resty.Response, error) {
	return NewInvoiceInService(client).Create(ctx, &invoiceIn, params...)
}

// Delete shortcut
func (invoiceIn InvoiceIn) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewInvoiceInService(client).Delete(ctx, invoiceIn.GetID())
}

// InvoiceInPosition Позиция Счета поставщика.
//
// Код сущности: invoiceposition
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-schet-postawschika-scheta-postawschikow-pozicii-scheta-postawschika
type InvoiceInPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учётной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	Discount   *float64            `json:"discount,omitempty"`   // Процент скидки или наценки. Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Price      *float64            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Vat        *int                `json:"vat,omitempty"`        // НДС, которым облагается текущая позиция
	VatEnabled *bool               `json:"vatEnabled,omitempty"` // Включен ли НДС для позиции. С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
	Stock      *Stock              `json:"stock,omitempty"`      // Остатки и себестоимость позиции (указывается при наличии параметра запроса `fields=stock`)
}

// GetAccountID возвращает ID учётной записи.
func (invoiceInPosition InvoiceInPosition) GetAccountID() uuid.UUID {
	return Deref(invoiceInPosition.AccountID)
}

// GetAssortment возвращает Метаданные товара/услуги/серии/модификации, которую представляет собой позиция.
func (invoiceInPosition InvoiceInPosition) GetAssortment() AssortmentPosition {
	return Deref(invoiceInPosition.Assortment)
}

// GetDiscount возвращает Процент скидки или наценки.
//
// Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%.
func (invoiceInPosition InvoiceInPosition) GetDiscount() float64 {
	return Deref(invoiceInPosition.Discount)
}

// GetID возвращает ID позиции.
func (invoiceInPosition InvoiceInPosition) GetID() uuid.UUID {
	return Deref(invoiceInPosition.ID)
}

// GetPack возвращает Упаковку Товара.
func (invoiceInPosition InvoiceInPosition) GetPack() Pack {
	return Deref(invoiceInPosition.Pack)
}

// GetPrice возвращает Цену товара/услуги в копейках.
func (invoiceInPosition InvoiceInPosition) GetPrice() float64 {
	return Deref(invoiceInPosition.Price)
}

// GetQuantity возвращает Количество товаров/услуг данного вида в компоненте.
func (invoiceInPosition InvoiceInPosition) GetQuantity() float64 {
	return Deref(invoiceInPosition.Quantity)
}

// GetVat возвращает НДС, которым облагается текущая позиция.
func (invoiceInPosition InvoiceInPosition) GetVat() int {
	return Deref(invoiceInPosition.Vat)
}

// GetVatEnabled возвращает true, если учитывается НДС.
func (invoiceInPosition InvoiceInPosition) GetVatEnabled() bool {
	return Deref(invoiceInPosition.VatEnabled)
}

// GetStock возвращает Остатки и себестоимость позиции (указывается при наличии параметра запроса `fields=stock`).
func (invoiceInPosition InvoiceInPosition) GetStock() Stock {
	return Deref(invoiceInPosition.Stock)
}

// SetAssortment устанавливает Метаданные товара/услуги, которую представляет собой компонент.
//
// Принимает объект, реализующий интерфейс [AssortmentInterface].
func (invoiceInPosition *InvoiceInPosition) SetAssortment(assortment AssortmentInterface) *InvoiceInPosition {
	if assortment != nil {
		invoiceInPosition.Assortment = assortment.asAssortment()
	}
	return invoiceInPosition
}

// SetDiscount устанавливает Процент скидки или наценки.
//
// Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%.
func (invoiceInPosition *InvoiceInPosition) SetDiscount(discount float64) *InvoiceInPosition {
	invoiceInPosition.Discount = &discount
	return invoiceInPosition
}

// SetPack устанавливает Упаковку Товара.
func (invoiceInPosition *InvoiceInPosition) SetPack(pack *Pack) *InvoiceInPosition {
	if pack != nil {
		invoiceInPosition.Pack = pack
	}
	return invoiceInPosition
}

// SetPrice устанавливает Цену товара/услуги в копейках.
func (invoiceInPosition *InvoiceInPosition) SetPrice(price float64) *InvoiceInPosition {
	invoiceInPosition.Price = &price
	return invoiceInPosition
}

// SetQuantity устанавливает Количество товаров данного вида в позиции.
//
// Если позиция - товар, у которого включен учет по серийным номерам,
// то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
func (invoiceInPosition *InvoiceInPosition) SetQuantity(quantity float64) *InvoiceInPosition {
	invoiceInPosition.Quantity = &quantity
	return invoiceInPosition
}

// SetVat устанавливает НДС, которым облагается текущая позиция.
func (invoiceInPosition *InvoiceInPosition) SetVat(vat int) *InvoiceInPosition {
	invoiceInPosition.Vat = &vat
	return invoiceInPosition
}

// SetVatEnabled устанавливает значение, учитывающее НДС для текущей позиции.
func (invoiceInPosition *InvoiceInPosition) SetVatEnabled(vatEnabled bool) *InvoiceInPosition {
	invoiceInPosition.VatEnabled = &vatEnabled
	return invoiceInPosition
}

// String реализует интерфейс [fmt.Stringer].
func (invoiceInPosition InvoiceInPosition) String() string {
	return Stringify(invoiceInPosition)
}

// MetaType возвращает код сущности.
func (InvoiceInPosition) MetaType() MetaType {
	return MetaTypeInvoicePosition
}

// InvoiceInService описывает методы сервиса для работы со счетами поставщиков.
type InvoiceInService interface {
	// GetList выполняет запрос на получение списка счетов поставщиков.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[InvoiceIn], *resty.Response, error)

	// Create выполняет запрос на создание счета поставщика.
	// Обязательные поля для заполнения:
	//	- name (Номер Счета поставщика)
	//	- organization (Ссылка на ваше юрлицо)
	//	- agent (Ссылка на контрагента (поставщика))
	// Принимает контекст, счет поставщика и опционально объект параметров запроса Params.
	// Возвращает созданный счет поставщика.
	Create(ctx context.Context, invoiceIn *InvoiceIn, params ...*Params) (*InvoiceIn, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение счетов поставщиков.
	// Изменяемые счета поставщиков должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список счетов поставщиков и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых счетов поставщиков.
	CreateUpdateMany(ctx context.Context, invoiceInList Slice[InvoiceIn], params ...*Params) (*Slice[InvoiceIn], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление счетов поставщиков.
	// Принимает контекст и множество счетов поставщиков.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*InvoiceIn) (*DeleteManyResponse, *resty.Response, error)

	// Delete выполняет запрос на удаление счета поставщика.
	// Принимает контекст и ID счета поставщика.
	// Возвращает true в случае успешного удаления счета поставщика.
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельного счета поставщика по ID.
	// Принимает контекст, ID счета поставщика и опционально объект параметров запроса Params.
	// Возвращает найденный счет поставщика.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*InvoiceIn, *resty.Response, error)

	// Update выполняет запрос на изменение счета поставщика.
	// Принимает контекст, счет поставщика и опционально объект параметров запроса Params.
	// Возвращает изменённый счет поставщика.
	Update(ctx context.Context, id uuid.UUID, invoiceIn *InvoiceIn, params ...*Params) (*InvoiceIn, *resty.Response, error)

	// Template выполняет запрос на получение предзаполненного счета поставщика со стандартными полями без связи с какими-либо другими документами.
	// Принимает контекст.
	// Возвращает предзаполненный счет поставщика.
	Template(ctx context.Context) (*InvoiceIn, *resty.Response, error)

	// TemplateBased выполняет запрос на получение шаблона счета поставщика на основе других документов.
	// Основание, на котором может быть создан:
	//	- Заказ поставщику (PurchaseOrder)
	//	- Приемка (Supply)
	// Принимает контекст и множество документов из списка выше.
	// Возвращает предзаполненный счет поставщика на основании переданных документов.
	TemplateBased(ctx context.Context, basedOn ...MetaOwner) (*InvoiceIn, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных счетов поставщиков.
	// Принимает контекст.
	// Возвращает объект метаданных MetaAttributesStatesSharedWrapper.
	GetMetadata(ctx context.Context) (*MetaAttributesStatesSharedWrapper, *resty.Response, error)

	// GetPositionList выполняет запрос на получение списка позиций документа.
	// Принимает контекст, ID документа и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetPositionList(ctx context.Context, id uuid.UUID, params ...*Params) (*List[InvoiceInPosition], *resty.Response, error)

	// GetPositionByID выполняет запрос на получение отдельной позиции документа по ID.
	// Принимает контекст, ID документа, ID позиции и опционально объект параметров запроса Params.
	// Возвращает найденную позицию.
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params ...*Params) (*InvoiceInPosition, *resty.Response, error)

	// UpdatePosition выполняет запрос на изменение позиции документа.
	// Принимает контекст, ID документа, ID позиции, позицию документа и опционально объект параметров запроса Params.
	// Возвращает изменённую позицию.
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *InvoiceInPosition, params ...*Params) (*InvoiceInPosition, *resty.Response, error)

	// CreatePosition выполняет запрос на добавление позиции документа.
	// Принимает контекст, ID документа, позицию документа и опционально объект параметров запроса Params.
	// Возвращает добавленную позицию.
	CreatePosition(ctx context.Context, id uuid.UUID, position *InvoiceInPosition, params ...*Params) (*InvoiceInPosition, *resty.Response, error)

	// CreatePositionMany выполняет запрос на массовое добавление позиций документа.
	// Принимает контекст, ID документа и множество позиций.
	// Возвращает список добавленных позиций.
	CreatePositionMany(ctx context.Context, id uuid.UUID, positions ...*InvoiceInPosition) (*Slice[InvoiceInPosition], *resty.Response, error)

	// DeletePosition выполняет запрос на удаление позиции документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает true в случае успешного удаления позиции.
	DeletePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (bool, *resty.Response, error)

	// DeletePositionMany выполняет запрос на массовое удаление позиций документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeletePositionMany(ctx context.Context, id uuid.UUID, positions ...*InvoiceInPosition) (*DeleteManyResponse, *resty.Response, error)

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
	// Возвращает true в случае успешного удаления доп поля.
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
	// Принимает контекст, ID документа и шаблон.
	// Возвращает созданную публикацию.
	Publish(ctx context.Context, id uuid.UUID, template TemplateInterface) (*Publication, *resty.Response, error)

	// DeletePublication выполняет запрос на удаление публикации.
	// Принимает контекст, ID документа и ID публикации.
	// Возвращает true в случае успешного удаления публикации.
	DeletePublication(ctx context.Context, id uuid.UUID, publicationID uuid.UUID) (bool, *resty.Response, error)

	// GetBySyncID выполняет запрос на получение отдельного документа по syncID.
	// Принимает контекст и syncID документа.
	// Возвращает найденный документ.
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*InvoiceIn, *resty.Response, error)

	// DeleteBySyncID выполняет запрос на удаление документа по syncID.
	// Принимает контекст и syncID документа.
	// Возвращает true в случае успешного удаления документа.
	DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error)

	// MoveToTrash выполняет запрос на перемещение документа с указанным ID в корзину.
	// Принимает контекст и ID документа.
	// Возвращает true в случае успешного перемещения в корзину.
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
	// Возвращает true в случае успешного удаления статуса.
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
	// Возвращает true в случае успешного удаления файла.
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
	Evaluate(ctx context.Context, entity *InvoiceIn, evaluate ...Evaluate) (*InvoiceIn, *resty.Response, error)
}

// NewInvoiceInService принимает [Client] и возвращает сервис для работы со счетами поставщиков.
func NewInvoiceInService(client *Client) InvoiceInService {
	return newMainService[InvoiceIn, InvoiceInPosition, MetaAttributesStatesSharedWrapper, any](NewEndpoint(client, "entity/invoicein"))
}
