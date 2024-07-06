package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"time"
)

// PurchaseOrder Заказ поставщику.
//
// Код сущности: purchaseorder
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-zakaz-postawschiku
type PurchaseOrder struct {
	PayedSum              *float64                          `json:"payedSum,omitempty"`              // Сумма входящих платежей по Заказу
	Applicable            *bool                             `json:"applicable,omitempty"`            // Отметка о проведении
	AgentAccount          *AgentAccount                     `json:"agentAccount,omitempty"`          // Метаданные счета контрагента
	Owner                 *Employee                         `json:"owner,omitempty"`                 // Метаданные владельца (Сотрудника)
	InternalOrder         *InternalOrder                    `json:"internalOrder,omitempty"`         // Внутренний заказ, связанный с заказом поставщику
	Code                  *string                           `json:"code,omitempty"`                  // Код Заказа поставщику
	Contract              *NullValue[Contract]              `json:"contract,omitempty"`              // Метаданные договора
	Created               *Timestamp                        `json:"created,omitempty"`               // Дата создания
	Deleted               *Timestamp                        `json:"deleted,omitempty"`               // Момент последнего удаления Заказа поставщику
	DeliveryPlannedMoment *Timestamp                        `json:"deliveryPlannedMoment,omitempty"` // Планируемая дата отгрузки
	OrganizationAccount   *AgentAccount                     `json:"organizationAccount,omitempty"`   // Метаданные счета юрлица
	ExternalCode          *string                           `json:"externalCode,omitempty"`          // Внешний код Заказа поставщику
	AccountID             *uuid.UUID                        `json:"accountId,omitempty"`             // ID учётной записи
	Group                 *Group                            `json:"group,omitempty"`                 // Отдел сотрудника
	ID                    *uuid.UUID                        `json:"id,omitempty"`                    // ID Заказа поставщику
	InvoicedSum           *float64                          `json:"invoicedSum,omitempty"`           // Сумма счетов поставщику
	Meta                  *Meta                             `json:"meta,omitempty"`                  // Метаданные Заказа поставщику
	Moment                *Timestamp                        `json:"moment,omitempty"`                // Дата документа
	Name                  *string                           `json:"name,omitempty"`                  // Наименование Заказа поставщику
	Organization          *Organization                     `json:"organization,omitempty"`          // Метаданные юрлица
	Description           *string                           `json:"description,omitempty"`           // Комментарий Заказа поставщику
	Agent                 *Agent                            `json:"agent,omitempty"`                 // Метаданные контрагента
	Files                 *MetaArray[File]                  `json:"files,omitempty"`                 // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Positions             *MetaArray[PurchaseOrderPosition] `json:"positions,omitempty"`             // Метаданные позиций Заказа поставщику
	Printed               *bool                             `json:"printed,omitempty"`               // Напечатан ли документ
	Project               *NullValue[Project]               `json:"project,omitempty"`               // Метаданные проекта
	Published             *bool                             `json:"published,omitempty"`             // Опубликован ли документ
	Rate                  *NullValue[Rate]                  `json:"rate,omitempty"`                  // Валюта
	Shared                *bool                             `json:"shared,omitempty"`                // Общий доступ
	ShippedSum            *float64                          `json:"shippedSum,omitempty"`            // Сумма принятого
	State                 *NullValue[State]                 `json:"state,omitempty"`                 // Метаданные статуса заказа поставщику
	Store                 *NullValue[Store]                 `json:"store,omitempty"`                 // Метаданные склада
	Sum                   *float64                          `json:"sum,omitempty"`                   // Сумма Заказа поставщику в установленной валюте
	SyncID                *uuid.UUID                        `json:"syncId,omitempty"`                // ID синхронизации
	Updated               *Timestamp                        `json:"updated,omitempty"`               // Момент последнего обновления Заказа поставщику
	VatEnabled            *bool                             `json:"vatEnabled,omitempty"`            // Учитывается ли НДС
	VatIncluded           *bool                             `json:"vatIncluded,omitempty"`           // Включен ли НДС в цену
	VatSum                *float64                          `json:"vatSum,omitempty"`                // Сумма НДС
	WaitSum               *float64                          `json:"waitSum,omitempty"`               // Сумма товаров в пути
	CustomerOrders        Slice[CustomerOrder]              `json:"customerOrders,omitempty"`        // Массив ссылок на связанные заказы покупателей
	InvoicesIn            Slice[InvoiceIn]                  `json:"invoicesIn,omitempty"`            // Массив ссылок на связанные счета поставщиков
	Payments              Slice[Payment]                    `json:"payments,omitempty"`              // Массив ссылок на связанные платежи
	Supplies              Slice[Supply]                     `json:"supplies,omitempty"`              // Массив ссылок на связанные приемки
	Attributes            Slice[Attribute]                  `json:"attributes,omitempty"`            // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (purchaseOrder PurchaseOrder) Clean() *PurchaseOrder {
	if purchaseOrder.Meta == nil {
		return nil
	}
	return &PurchaseOrder{Meta: purchaseOrder.Meta}
}

// AsOperation возвращает объект [Operation] c полем [Meta].
func (purchaseOrder PurchaseOrder) AsOperation() *Operation {
	return &Operation{Meta: purchaseOrder.GetMeta(), LinkedSum: purchaseOrder.GetSum()}
}

// AsTaskOperation реализует интерфейс [TaskOperationInterface].
func (purchaseOrder PurchaseOrder) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: purchaseOrder.Meta}
}

// AsOperationOut реализует интерфейс [OperationOut].
func (purchaseOrder PurchaseOrder) AsOperationOut() *Operation {
	return purchaseOrder.AsOperation()
}

// GetPayedSum возвращает Сумму входящих платежей по Заказу.
func (purchaseOrder PurchaseOrder) GetPayedSum() float64 {
	return Deref(purchaseOrder.PayedSum)
}

// GetApplicable возвращает Отметку о проведении.
func (purchaseOrder PurchaseOrder) GetApplicable() bool {
	return Deref(purchaseOrder.Applicable)
}

// GetAgentAccount возвращает Метаданные счета контрагента.
func (purchaseOrder PurchaseOrder) GetAgentAccount() AgentAccount {
	return Deref(purchaseOrder.AgentAccount)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (purchaseOrder PurchaseOrder) GetOwner() Employee {
	return Deref(purchaseOrder.Owner)
}

// GetInternalOrder возвращает Внутренний заказ, связанный с заказом поставщику.
func (purchaseOrder PurchaseOrder) GetInternalOrder() InternalOrder {
	return Deref(purchaseOrder.InternalOrder)
}

// GetCode возвращает Код Заказа поставщику.
func (purchaseOrder PurchaseOrder) GetCode() string {
	return Deref(purchaseOrder.Code)
}

// GetContract возвращает Метаданные договора.
func (purchaseOrder PurchaseOrder) GetContract() Contract {
	return Deref(purchaseOrder.Contract).GetValue()
}

// GetCreated возвращает Дату создания.
func (purchaseOrder PurchaseOrder) GetCreated() Timestamp {
	return Deref(purchaseOrder.Created)
}

// GetDeleted возвращает Момент последнего удаления Заказа поставщику.
func (purchaseOrder PurchaseOrder) GetDeleted() Timestamp {
	return Deref(purchaseOrder.Deleted)
}

// GetDeliveryPlannedMoment возвращает Планируемую дата отгрузки.
func (purchaseOrder PurchaseOrder) GetDeliveryPlannedMoment() Timestamp {
	return Deref(purchaseOrder.DeliveryPlannedMoment)
}

// GetOrganizationAccount возвращает Метаданные счета юрлица.
func (purchaseOrder PurchaseOrder) GetOrganizationAccount() AgentAccount {
	return Deref(purchaseOrder.OrganizationAccount)
}

// GetExternalCode возвращает Внешний код Заказа поставщику.
func (purchaseOrder PurchaseOrder) GetExternalCode() string {
	return Deref(purchaseOrder.ExternalCode)
}

// GetAccountID возвращает ID учётной записи.
func (purchaseOrder PurchaseOrder) GetAccountID() uuid.UUID {
	return Deref(purchaseOrder.AccountID)
}

// GetGroup возвращает Отдел сотрудника.
func (purchaseOrder PurchaseOrder) GetGroup() Group {
	return Deref(purchaseOrder.Group)
}

// GetID возвращает ID Перемещения.
func (purchaseOrder PurchaseOrder) GetID() uuid.UUID {
	return Deref(purchaseOrder.ID)
}

// GetInvoicedSum возвращает Сумму счетов поставщику.
func (purchaseOrder PurchaseOrder) GetInvoicedSum() float64 {
	return Deref(purchaseOrder.InvoicedSum)
}

// GetMeta возвращает Метаданные Заказа поставщику.
func (purchaseOrder PurchaseOrder) GetMeta() Meta {
	return Deref(purchaseOrder.Meta)
}

// GetMoment возвращает Дату документа.
func (purchaseOrder PurchaseOrder) GetMoment() Timestamp {
	return Deref(purchaseOrder.Moment)
}

// GetName возвращает Наименование Заказа поставщику.
func (purchaseOrder PurchaseOrder) GetName() string {
	return Deref(purchaseOrder.Name)
}

// GetOrganization возвращает Метаданные юрлица.
func (purchaseOrder PurchaseOrder) GetOrganization() Organization {
	return Deref(purchaseOrder.Organization)
}

// GetDescription возвращает Комментарий Заказа поставщику.
func (purchaseOrder PurchaseOrder) GetDescription() string {
	return Deref(purchaseOrder.Description)
}

// GetAgent возвращает Метаданные контрагента.
func (purchaseOrder PurchaseOrder) GetAgent() Agent {
	return Deref(purchaseOrder.Agent)
}

// GetFiles возвращает Метаданные массива Файлов.
func (purchaseOrder PurchaseOrder) GetFiles() MetaArray[File] {
	return Deref(purchaseOrder.Files)
}

// GetPositions возвращает Метаданные позиций Заказа поставщику.
func (purchaseOrder PurchaseOrder) GetPositions() MetaArray[PurchaseOrderPosition] {
	return Deref(purchaseOrder.Positions)
}

// GetPrinted возвращает true, если документ напечатан.
func (purchaseOrder PurchaseOrder) GetPrinted() bool {
	return Deref(purchaseOrder.Printed)
}

// GetProject возвращает Метаданные проекта.
func (purchaseOrder PurchaseOrder) GetProject() Project {
	return Deref(purchaseOrder.Project).GetValue()
}

// GetPublished возвращает true, если документ опубликован.
func (purchaseOrder PurchaseOrder) GetPublished() bool {
	return Deref(purchaseOrder.Published)
}

// GetRate возвращает Валюту.
func (purchaseOrder PurchaseOrder) GetRate() Rate {
	return Deref(purchaseOrder.Rate).GetValue()
}

// GetShared возвращает флаг Общего доступа.
func (purchaseOrder PurchaseOrder) GetShared() bool {
	return Deref(purchaseOrder.Shared)
}

// GetShippedSum возвращает Сумму принятого.
func (purchaseOrder PurchaseOrder) GetShippedSum() float64 {
	return Deref(purchaseOrder.ShippedSum)
}

// GetState возвращает Метаданные статуса заказа поставщику.
func (purchaseOrder PurchaseOrder) GetState() State {
	return Deref(purchaseOrder.State).GetValue()
}

// GetStore возвращает Метаданные склада.
func (purchaseOrder PurchaseOrder) GetStore() Store {
	return Deref(purchaseOrder.Store).GetValue()
}

// GetSum возвращает Сумму Заказа поставщику в установленной валюте.
func (purchaseOrder PurchaseOrder) GetSum() float64 {
	return Deref(purchaseOrder.Sum)
}

// GetSyncID возвращает ID синхронизации.
func (purchaseOrder PurchaseOrder) GetSyncID() uuid.UUID {
	return Deref(purchaseOrder.SyncID)
}

// GetUpdated возвращает Момент последнего обновления заказа поставщику.
func (purchaseOrder PurchaseOrder) GetUpdated() Timestamp {
	return Deref(purchaseOrder.Updated)
}

// GetVatEnabled возвращает true, если учитывается НДС.
func (purchaseOrder PurchaseOrder) GetVatEnabled() bool {
	return Deref(purchaseOrder.VatEnabled)
}

// GetVatIncluded возвращает true, если НДС включен в цену.
func (purchaseOrder PurchaseOrder) GetVatIncluded() bool {
	return Deref(purchaseOrder.VatIncluded)
}

// GetVatSum возвращает Сумму НДС.
func (purchaseOrder PurchaseOrder) GetVatSum() float64 {
	return Deref(purchaseOrder.VatSum)
}

// GetWaitSum возвращает Сумму товаров в пути.
func (purchaseOrder PurchaseOrder) GetWaitSum() float64 {
	return Deref(purchaseOrder.WaitSum)
}

// GetCustomerOrders возвращает Массив ссылок на связанные заказы покупателей.
func (purchaseOrder PurchaseOrder) GetCustomerOrders() Slice[CustomerOrder] {
	return purchaseOrder.CustomerOrders
}

// GetInvoicesIn возвращает Массив ссылок на связанные счета поставщиков.
func (purchaseOrder PurchaseOrder) GetInvoicesIn() Slice[InvoiceIn] {
	return purchaseOrder.InvoicesIn
}

// GetPayments возвращает Массив ссылок на связанные платежи.
func (purchaseOrder PurchaseOrder) GetPayments() Slice[Payment] {
	return purchaseOrder.Payments
}

// GetSupplies возвращает Массив ссылок на связанные приемки.
func (purchaseOrder PurchaseOrder) GetSupplies() Slice[Supply] {
	return purchaseOrder.Supplies
}

// GetAttributes возвращает Список метаданных доп. полей.
func (purchaseOrder PurchaseOrder) GetAttributes() Slice[Attribute] {
	return purchaseOrder.Attributes
}

// SetApplicable устанавливает Отметку о проведении.
func (purchaseOrder *PurchaseOrder) SetApplicable(applicable bool) *PurchaseOrder {
	purchaseOrder.Applicable = &applicable
	return purchaseOrder
}

// SetAgentAccount устанавливает Метаданные счета контрагента.
func (purchaseOrder *PurchaseOrder) SetAgentAccount(agentAccount *AgentAccount) *PurchaseOrder {
	if agentAccount != nil {
		purchaseOrder.AgentAccount = agentAccount.Clean()
	}
	return purchaseOrder
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (purchaseOrder *PurchaseOrder) SetOwner(owner *Employee) *PurchaseOrder {
	if owner != nil {
		purchaseOrder.Owner = owner.Clean()
	}
	return purchaseOrder
}

// SetInternalOrder устанавливает Внутренний заказ, связанный с заказом поставщику.
func (purchaseOrder *PurchaseOrder) SetInternalOrder(internalOrder *InternalOrder) *PurchaseOrder {
	if internalOrder != nil {
		purchaseOrder.InternalOrder = internalOrder.Clean()
	}
	return purchaseOrder
}

// SetCode устанавливает Код Заказа поставщику.
func (purchaseOrder *PurchaseOrder) SetCode(code string) *PurchaseOrder {
	purchaseOrder.Code = &code
	return purchaseOrder
}

// SetContract устанавливает Метаданные договора.
//
// Передача nil передаёт сброс значения (null).
func (purchaseOrder *PurchaseOrder) SetContract(contract *Contract) *PurchaseOrder {
	purchaseOrder.Contract = NewNullValue(contract)
	return purchaseOrder
}

// SetDeliveryPlannedMoment устанавливает Планируемую дата отгрузки.
func (purchaseOrder *PurchaseOrder) SetDeliveryPlannedMoment(deliveryPlannedMoment time.Time) *PurchaseOrder {
	purchaseOrder.DeliveryPlannedMoment = NewTimestamp(deliveryPlannedMoment)
	return purchaseOrder
}

// SetOrganizationAccount устанавливает Метаданные счета юрлица.
func (purchaseOrder *PurchaseOrder) SetOrganizationAccount(organizationAccount *AgentAccount) *PurchaseOrder {
	if organizationAccount != nil {
		purchaseOrder.OrganizationAccount = organizationAccount.Clean()
	}
	return purchaseOrder
}

// SetExternalCode устанавливает Внешний код Заказа поставщику.
func (purchaseOrder *PurchaseOrder) SetExternalCode(externalCode string) *PurchaseOrder {
	purchaseOrder.ExternalCode = &externalCode
	return purchaseOrder
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (purchaseOrder *PurchaseOrder) SetGroup(group *Group) *PurchaseOrder {
	if group != nil {
		purchaseOrder.Group = group.Clean()
	}
	return purchaseOrder
}

// SetMeta устанавливает Метаданные Заказа поставщику.
func (purchaseOrder *PurchaseOrder) SetMeta(meta *Meta) *PurchaseOrder {
	purchaseOrder.Meta = meta
	return purchaseOrder
}

// SetMoment устанавливает Дату документа.
func (purchaseOrder *PurchaseOrder) SetMoment(moment time.Time) *PurchaseOrder {
	purchaseOrder.Moment = NewTimestamp(moment)
	return purchaseOrder
}

// SetName устанавливает Наименование Заказа поставщику.
func (purchaseOrder *PurchaseOrder) SetName(name string) *PurchaseOrder {
	purchaseOrder.Name = &name
	return purchaseOrder
}

// SetOrganization устанавливает Метаданные юрлица.
func (purchaseOrder *PurchaseOrder) SetOrganization(organization *Organization) *PurchaseOrder {
	if organization != nil {
		purchaseOrder.Organization = organization.Clean()
	}
	return purchaseOrder
}

// SetDescription устанавливает Комментарий Заказа поставщику.
func (purchaseOrder *PurchaseOrder) SetDescription(description string) *PurchaseOrder {
	purchaseOrder.Description = &description
	return purchaseOrder
}

// SetAgent устанавливает Метаданные Контрагента.
//
// Принимает [Counterparty] или [Organization].
func (purchaseOrder *PurchaseOrder) SetAgent(agent AgentCounterpartyOrganizationInterface) *PurchaseOrder {
	if agent != nil {
		purchaseOrder.Agent = agent.AsCOAgent()
	}
	return purchaseOrder
}

// SetFiles устанавливает Метаданные массива Файлов.
//
// Принимает множество объектов [File].
func (purchaseOrder *PurchaseOrder) SetFiles(files ...*File) *PurchaseOrder {
	purchaseOrder.Files = NewMetaArrayFrom(files)
	return purchaseOrder
}

// SetPositions устанавливает Метаданные позиций Заказа поставщику.
//
// Принимает множество объектов [PurchaseOrderPosition].
func (purchaseOrder *PurchaseOrder) SetPositions(positions ...*PurchaseOrderPosition) *PurchaseOrder {
	purchaseOrder.Positions = NewMetaArrayFrom(positions)
	return purchaseOrder
}

// SetProject устанавливает Метаданные проекта.
//
// Передача nil передаёт сброс значения (null).
func (purchaseOrder *PurchaseOrder) SetProject(project *Project) *PurchaseOrder {
	purchaseOrder.Project = NewNullValue(project)
	return purchaseOrder
}

// SetRate устанавливает Валюту.
//
// Передача nil передаёт сброс значения (null).
func (purchaseOrder *PurchaseOrder) SetRate(rate *Rate) *PurchaseOrder {
	purchaseOrder.Rate = NewNullValue(rate)
	return purchaseOrder
}

// SetShared устанавливает флаг общего доступа.
func (purchaseOrder *PurchaseOrder) SetShared(shared bool) *PurchaseOrder {
	purchaseOrder.Shared = &shared
	return purchaseOrder
}

// SetState устанавливает Метаданные статуса заказа поставщику.
//
// Передача nil передаёт сброс значения (null).
func (purchaseOrder *PurchaseOrder) SetState(state *State) *PurchaseOrder {
	purchaseOrder.State = NewNullValue(state)
	return purchaseOrder
}

// SetStore устанавливает Метаданные склада.
//
// Передача nil передаёт сброс значения (null).
func (purchaseOrder *PurchaseOrder) SetStore(store *Store) *PurchaseOrder {
	purchaseOrder.Store = NewNullValue(store)
	return purchaseOrder
}

// SetSyncID устанавливает ID синхронизации.
func (purchaseOrder *PurchaseOrder) SetSyncID(syncID uuid.UUID) *PurchaseOrder {
	purchaseOrder.SyncID = &syncID
	return purchaseOrder
}

// SetVatEnabled устанавливает значение, учитывающее НДС для заказа поставщику.
func (purchaseOrder *PurchaseOrder) SetVatEnabled(vatEnabled bool) *PurchaseOrder {
	purchaseOrder.VatEnabled = &vatEnabled
	return purchaseOrder
}

// SetVatIncluded устанавливает флаг включения НДС в цену.
func (purchaseOrder *PurchaseOrder) SetVatIncluded(vatIncluded bool) *PurchaseOrder {
	purchaseOrder.VatIncluded = &vatIncluded
	return purchaseOrder
}

// SetCustomerOrders устанавливает Массив ссылок на связанные заказы покупателей.
//
// Принимает множество объектов [CustomerOrder].
func (purchaseOrder *PurchaseOrder) SetCustomerOrders(customerOrders ...*CustomerOrder) *PurchaseOrder {
	purchaseOrder.CustomerOrders.Push(customerOrders...)
	return purchaseOrder
}

// SetInvoicesIn устанавливает Массив ссылок на связанные счета поставщиков.
//
// Принимает множество объектов [InvoiceIn].
func (purchaseOrder *PurchaseOrder) SetInvoicesIn(invoicesIn ...*InvoiceIn) *PurchaseOrder {
	purchaseOrder.InvoicesIn.Push(invoicesIn...)
	return purchaseOrder
}

// SetPayments устанавливает Метаданные ссылок на связанные платежи.
//
// Принимает множество объектов, реализующих интерфейс [PaymentInterface].
func (purchaseOrder *PurchaseOrder) SetPayments(payments ...PaymentInterface) *PurchaseOrder {
	purchaseOrder.Payments = NewPaymentsFrom(payments)
	return purchaseOrder
}

// SetSupplies устанавливает Массив ссылок на связанные приемки.
//
// Принимает множество объектов [Supply].
func (purchaseOrder *PurchaseOrder) SetSupplies(supplies ...*Supply) *PurchaseOrder {
	purchaseOrder.Supplies.Push(supplies...)
	return purchaseOrder
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (purchaseOrder *PurchaseOrder) SetAttributes(attributes ...*Attribute) *PurchaseOrder {
	purchaseOrder.Attributes.Push(attributes...)
	return purchaseOrder
}

// String реализует интерфейс [fmt.Stringer].
func (purchaseOrder PurchaseOrder) String() string {
	return Stringify(purchaseOrder)
}

// MetaType возвращает код сущности.
func (PurchaseOrder) MetaType() MetaType {
	return MetaTypePurchaseOrder
}

// Update shortcut
func (purchaseOrder PurchaseOrder) Update(ctx context.Context, client *Client, params ...*Params) (*PurchaseOrder, *resty.Response, error) {
	return NewPurchaseOrderService(client).Update(ctx, purchaseOrder.GetID(), &purchaseOrder, params...)
}

// Create shortcut
func (purchaseOrder PurchaseOrder) Create(ctx context.Context, client *Client, params ...*Params) (*PurchaseOrder, *resty.Response, error) {
	return NewPurchaseOrderService(client).Create(ctx, &purchaseOrder, params...)
}

// Delete shortcut
func (purchaseOrder PurchaseOrder) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewPurchaseOrderService(client).Delete(ctx, purchaseOrder.GetID())
}

// PurchaseOrderPosition Позиция Заказа поставщику.
//
// Код сущности: purchaseorderposition
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-zakaz-postawschiku-zakazy-postawschikam-pozicii-zakaza-postawschiku
type PurchaseOrderPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учётной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	Discount   *float64            `json:"discount,omitempty"`   // Процент скидки или наценки. Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Price      *float64            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Shipped    *float64            `json:"shipped,omitempty"`    // Принято
	InTransit  *float64            `json:"inTransit,omitempty"`  // Ожидание
	Vat        *float64            `json:"vat,omitempty"`        // НДС, которым облагается текущая позиция
	VatEnabled *bool               `json:"vatEnabled,omitempty"` // Включен ли НДС для позиции. С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
	Wait       *bool               `json:"wait,omitempty"`       // Ожидается данной позиции
	Stock      *Stock              `json:"stock,omitempty"`      // Остатки и себестоимость `?fields=stock&expand=positions`
}

// GetAccountID возвращает ID учётной записи.
func (purchaseOrderPosition PurchaseOrderPosition) GetAccountID() uuid.UUID {
	return Deref(purchaseOrderPosition.AccountID)
}

// GetAssortment возвращает Метаданные товара/услуги/серии/модификации, которую представляет собой позиция.
func (purchaseOrderPosition PurchaseOrderPosition) GetAssortment() AssortmentPosition {
	return Deref(purchaseOrderPosition.Assortment)
}

// GetDiscount возвращает Процент скидки или наценки.
//
// Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%
func (purchaseOrderPosition PurchaseOrderPosition) GetDiscount() float64 {
	return Deref(purchaseOrderPosition.Discount)
}

// GetID возвращает ID позиции.
func (purchaseOrderPosition PurchaseOrderPosition) GetID() uuid.UUID {
	return Deref(purchaseOrderPosition.ID)
}

// GetPack возвращает Упаковку Товара.
func (purchaseOrderPosition PurchaseOrderPosition) GetPack() Pack {
	return Deref(purchaseOrderPosition.Pack)
}

// GetPrice возвращает Цену товара/услуги в копейках.
func (purchaseOrderPosition PurchaseOrderPosition) GetPrice() float64 {
	return Deref(purchaseOrderPosition.Price)
}

// GetQuantity возвращает Количество товаров данного вида в позиции.
func (purchaseOrderPosition PurchaseOrderPosition) GetQuantity() float64 {
	return Deref(purchaseOrderPosition.Quantity)
}

// GetShipped возвращает Принято.
func (purchaseOrderPosition PurchaseOrderPosition) GetShipped() float64 {
	return Deref(purchaseOrderPosition.Shipped)
}

// GetInTransit возвращает Ожидание.
func (purchaseOrderPosition PurchaseOrderPosition) GetInTransit() float64 {
	return Deref(purchaseOrderPosition.InTransit)
}

// GetVat возвращает НДС, которым облагается текущая позиция.
func (purchaseOrderPosition PurchaseOrderPosition) GetVat() float64 {
	return Deref(purchaseOrderPosition.Vat)
}

// GetVatEnabled возвращает true, если учитывается НДС для позиции.
func (purchaseOrderPosition PurchaseOrderPosition) GetVatEnabled() bool {
	return Deref(purchaseOrderPosition.VatEnabled)
}

// GetWait возвращает признак ожидания данной позиции.
func (purchaseOrderPosition PurchaseOrderPosition) GetWait() bool {
	return Deref(purchaseOrderPosition.Wait)
}

// GetStock возвращает Остатки и себестоимость позиции (указывается при наличии параметра запроса `fields=stock`).
func (purchaseOrderPosition PurchaseOrderPosition) GetStock() Stock {
	return Deref(purchaseOrderPosition.Stock)
}

// SetAssortment устанавливает Метаданные товара/услуги/серии/модификации, которую представляет собой позиция.
//
// Принимает объект, реализующий интерфейс [AssortmentInterface].
func (purchaseOrderPosition *PurchaseOrderPosition) SetAssortment(assortment AssortmentInterface) *PurchaseOrderPosition {
	if assortment != nil {
		purchaseOrderPosition.Assortment = assortment.AsAssortment()
	}
	return purchaseOrderPosition
}

// SetDiscount устанавливает Процент скидки или наценки.
//
// Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%.
func (purchaseOrderPosition *PurchaseOrderPosition) SetDiscount(discount float64) *PurchaseOrderPosition {
	purchaseOrderPosition.Discount = &discount
	return purchaseOrderPosition
}

// SetPack устанавливает Упаковку Товара.
func (purchaseOrderPosition *PurchaseOrderPosition) SetPack(pack *Pack) *PurchaseOrderPosition {
	purchaseOrderPosition.Pack = pack
	return purchaseOrderPosition
}

// SetPrice устанавливает Цену товара/услуги в копейках.
func (purchaseOrderPosition *PurchaseOrderPosition) SetPrice(price float64) *PurchaseOrderPosition {
	purchaseOrderPosition.Price = &price
	return purchaseOrderPosition
}

// SetQuantity устанавливает Количество товаров данного вида в позиции.
func (purchaseOrderPosition *PurchaseOrderPosition) SetQuantity(quantity float64) *PurchaseOrderPosition {
	purchaseOrderPosition.Quantity = &quantity
	return purchaseOrderPosition
}

// SetShipped устанавливает Принято.
func (purchaseOrderPosition *PurchaseOrderPosition) SetShipped(shipped float64) *PurchaseOrderPosition {
	purchaseOrderPosition.Shipped = &shipped
	return purchaseOrderPosition
}

// SetInTransit устанавливает Ожидание.
func (purchaseOrderPosition *PurchaseOrderPosition) SetInTransit(inTransit float64) *PurchaseOrderPosition {
	purchaseOrderPosition.InTransit = &inTransit
	return purchaseOrderPosition
}

// SetVat устанавливает НДС, которым облагается текущая позиция.
func (purchaseOrderPosition *PurchaseOrderPosition) SetVat(vat float64) *PurchaseOrderPosition {
	purchaseOrderPosition.Vat = &vat
	return purchaseOrderPosition
}

// SetVatEnabled устанавливает значение, учитывающее НДС для текущей позиции.
func (purchaseOrderPosition *PurchaseOrderPosition) SetVatEnabled(vatEnabled bool) *PurchaseOrderPosition {
	purchaseOrderPosition.VatEnabled = &vatEnabled
	return purchaseOrderPosition
}

// SetWait устанавливает признак ожидания данной позиции.
func (purchaseOrderPosition *PurchaseOrderPosition) SetWait(wait bool) *PurchaseOrderPosition {
	purchaseOrderPosition.Wait = &wait
	return purchaseOrderPosition
}

// String реализует интерфейс [fmt.Stringer].
func (purchaseOrderPosition PurchaseOrderPosition) String() string {
	return Stringify(purchaseOrderPosition)
}

// MetaType возвращает код сущности.
func (PurchaseOrderPosition) MetaType() MetaType {
	return MetaTypePurchaseOrderPosition
}

// PurchaseOrderService описывает методы сервиса для работы с заказами поставщикам.
type PurchaseOrderService interface {
	// GetList выполняет запрос на получение списка заказов поставщику.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[PurchaseOrder], *resty.Response, error)

	// Create выполняет запрос на создание заказа поставщику.
	// Обязательные поля для заполнения:
	//	- organization (Ссылка на ваше юрлицо)
	//	- agent (Ссылка на контрагента (поставщика))
	// Принимает контекст, заказ поставщику и опционально объект параметров запроса Params.
	// Возвращает созданный заказ поставщику.
	Create(ctx context.Context, purchaseOrder *PurchaseOrder, params ...*Params) (*PurchaseOrder, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение заказов поставщику.
	// Изменяемые заказы поставщику должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список заказов поставщику и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых заказов поставщику.
	CreateUpdateMany(ctx context.Context, purchaseOrderList Slice[PurchaseOrder], params ...*Params) (*Slice[PurchaseOrder], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление заказов поставщику.
	// Принимает контекст и множество заказов поставщику.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*PurchaseOrder) (*DeleteManyResponse, *resty.Response, error)

	// Delete выполняет запрос на удаление заказа поставщику.
	// Принимает контекст и ID заказа поставщику.
	// Возвращает «true» в случае успешного удаления заказа поставщику.
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельного заказа поставщику по ID.
	// Принимает контекст, ID заказа поставщику и опционально объект параметров запроса Params.
	// Возвращает найденный заказ поставщику.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*PurchaseOrder, *resty.Response, error)

	// Update выполняет запрос на изменение заказа поставщику.
	// Принимает контекст, заказ поставщику и опционально объект параметров запроса Params.
	// Возвращает изменённый заказ поставщику.
	Update(ctx context.Context, id uuid.UUID, purchaseOrder *PurchaseOrder, params ...*Params) (*PurchaseOrder, *resty.Response, error)

	// Template выполняет запрос на получение предзаполненного заказа поставщику со стандартными полями без связи с какими-либо другими документами.
	// Принимает контекст.
	// Возвращает предзаполненный заказ поставщику.
	Template(ctx context.Context) (*PurchaseOrder, *resty.Response, error)

	// TemplateBased выполняет запрос на получение шаблона заказа поставщику на основе других документов.
	// Основание, на котором может быть создан:
	//	- Внутренний заказ (InternalOrder)
	//	- Заказ покупателя (CustomerOrder)
	// Принимает контекст и множество документов из списка выше.
	// Возвращает предзаполненный заказ поставщику на основании переданных документов.
	TemplateBased(ctx context.Context, basedOn ...MetaOwner) (*PurchaseOrder, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных заказов поставщику.
	// Принимает контекст.
	// Возвращает объект метаданных MetaAttributesStatesSharedWrapper.
	GetMetadata(ctx context.Context) (*MetaAttributesStatesSharedWrapper, *resty.Response, error)

	// GetPositionList выполняет запрос на получение списка позиций документа.
	// Принимает контекст, ID документа и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetPositionList(ctx context.Context, id uuid.UUID, params ...*Params) (*List[PurchaseOrderPosition], *resty.Response, error)

	// GetPositionByID выполняет запрос на получение отдельной позиции документа по ID.
	// Принимает контекст, ID документа, ID позиции и опционально объект параметров запроса Params.
	// Возвращает найденную позицию.
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params ...*Params) (*PurchaseOrderPosition, *resty.Response, error)

	// UpdatePosition выполняет запрос на изменение позиции документа.
	// Принимает контекст, ID документа, ID позиции, позицию документа и опционально объект параметров запроса Params.
	// Возвращает изменённую позицию.
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *PurchaseOrderPosition, params ...*Params) (*PurchaseOrderPosition, *resty.Response, error)

	// CreatePosition выполняет запрос на добавление позиции документа.
	// Принимает контекст, ID документа, позицию документа и опционально объект параметров запроса Params.
	// Возвращает добавленную позицию.
	CreatePosition(ctx context.Context, id uuid.UUID, position *PurchaseOrderPosition, params ...*Params) (*PurchaseOrderPosition, *resty.Response, error)

	// CreatePositionMany выполняет запрос на массовое добавление позиций документа.
	// Принимает контекст, ID документа и множество позиций.
	// Возвращает список добавленных позиций.
	CreatePositionMany(ctx context.Context, id uuid.UUID, positions ...*PurchaseOrderPosition) (*Slice[PurchaseOrderPosition], *resty.Response, error)

	// DeletePosition выполняет запрос на удаление позиции документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает «true» в случае успешного удаления позиции.
	DeletePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (bool, *resty.Response, error)

	// DeletePositionMany выполняет запрос на массовое удаление позиций документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeletePositionMany(ctx context.Context, id uuid.UUID, positions ...*PurchaseOrderPosition) (*DeleteManyResponse, *resty.Response, error)

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
	// Принимает контекст, ID документа и шаблон.
	// Возвращает созданную публикацию.
	Publish(ctx context.Context, id uuid.UUID, template TemplateInterface) (*Publication, *resty.Response, error)

	// DeletePublication выполняет запрос на удаление публикации.
	// Принимает контекст, ID документа и ID публикации.
	// Возвращает «true» в случае успешного удаления публикации.
	DeletePublication(ctx context.Context, id uuid.UUID, publicationID uuid.UUID) (bool, *resty.Response, error)

	// GetBySyncID выполняет запрос на получение отдельного документа по syncID.
	// Принимает контекст и syncID документа.
	// Возвращает найденный документ.
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*PurchaseOrder, *resty.Response, error)

	// DeleteBySyncID выполняет запрос на удаление документа по syncID.
	// Принимает контекст и syncID документа.
	// Возвращает «true» в случае успешного удаления документа.
	DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error)

	// GetNamedFilterList выполняет запрос на получение списка фильтров.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetNamedFilterList(ctx context.Context, params ...*Params) (*List[NamedFilter], *resty.Response, error)

	// GetNamedFilterByID выполняет запрос на получение отдельного фильтра по ID.
	// Принимает контекст и ID фильтра.
	// Возвращает найденный фильтр.
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)

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
	Evaluate(ctx context.Context, entity *PurchaseOrder, evaluate ...Evaluate) (*PurchaseOrder, *resty.Response, error)
}

// NewPurchaseOrderService принимает [Client] и возвращает сервис для работы с заказами поставщикам.
func NewPurchaseOrderService(client *Client) PurchaseOrderService {
	return newMainService[PurchaseOrder, PurchaseOrderPosition, MetaAttributesStatesSharedWrapper, any](NewEndpoint(client, "entity/purchaseorder"))
}
