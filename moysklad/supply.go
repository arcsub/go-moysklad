package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"time"
)

// Supply Приёмка.
//
// Код сущности: supply
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-priemka
type Supply struct {
	OrganizationAccount *AgentAccount              `json:"organizationAccount,omitempty"` // Метаданные счета юрлица
	Applicable          *bool                      `json:"applicable,omitempty"`          // Отметка о проведении
	AgentAccount        *AgentAccount              `json:"agentAccount,omitempty"`        // Метаданные счета контрагента
	Overhead            *Overhead                  `json:"overhead,omitempty"`            // Накладные расходы. Если Позиции Приемки не заданы, то накладные расходы нельзя задать
	Returns             Slice[PurchaseReturn]      `json:"returns,omitempty"`             // Массив ссылок на связанные возвраты
	Code                *string                    `json:"code,omitempty"`                // Код Приемки
	Contract            *NullValue[Contract]       `json:"contract,omitempty"`            // Метаданные договора
	Created             *Timestamp                 `json:"created,omitempty"`             // Дата создания
	Deleted             *Timestamp                 `json:"deleted,omitempty"`             // Момент последнего удаления Приемки
	Description         *string                    `json:"description,omitempty"`         // Комментарий Приемки
	ExternalCode        *string                    `json:"externalCode,omitempty"`        // Внешний код Приемки
	Files               *MetaArray[File]           `json:"files,omitempty"`               // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group               *Group                     `json:"group,omitempty"`               // Отдел сотрудника
	ID                  *uuid.UUID                 `json:"id,omitempty"`                  // ID Приемки
	IncomingDate        *Timestamp                 `json:"incomingDate,omitempty"`        // Входящая дата
	Owner               *Employee                  `json:"owner,omitempty"`               // Метаданные владельца (Сотрудника)
	Meta                *Meta                      `json:"meta,omitempty"`                // Метаданные Приемки
	Moment              *Timestamp                 `json:"moment,omitempty"`              // Дата документа
	Name                *string                    `json:"name,omitempty"`                // Наименование Приемки
	Organization        *Organization              `json:"organization,omitempty"`        // Метаданные юрлица
	Payments            Slice[Payment]             `json:"payments,omitempty"`            // Массив ссылок на связанные платежи
	Agent               *Agent                     `json:"agent,omitempty"`               // Метаданные контрагента
	IncomingNumber      *string                    `json:"incomingNumber,omitempty"`      // Входящий номер
	PayedSum            *float64                   `json:"payedSum,omitempty"`            // Сумма входящих платежей по Приемке
	Positions           *MetaArray[SupplyPosition] `json:"positions,omitempty"`           // Метаданные позиций Приемки
	Printed             *bool                      `json:"printed,omitempty"`             // Напечатан ли документ
	Project             *NullValue[Project]        `json:"project,omitempty"`             // Метаданные проекта
	Published           *bool                      `json:"published,omitempty"`           // Опубликован ли документ
	Rate                *NullValue[Rate]           `json:"rate,omitempty"`                // Валюта
	Shared              *bool                      `json:"shared,omitempty"`              // Общий доступ
	State               *NullValue[State]          `json:"state,omitempty"`               // Метаданные статуса Приемки
	Store               *Store                     `json:"store,omitempty"`               // Метаданные склада
	Sum                 *float64                   `json:"sum,omitempty"`                 // Сумма Приемки в копейках
	SyncID              *uuid.UUID                 `json:"syncId,omitempty"`              // ID синхронизации
	Updated             *Timestamp                 `json:"updated,omitempty"`             // Момент последнего обновления Приемки
	VatEnabled          *bool                      `json:"vatEnabled,omitempty"`          // Учитывается ли НДС
	VatIncluded         *bool                      `json:"vatIncluded,omitempty"`         // Включен ли НДС в цену
	VatSum              *float64                   `json:"vatSum,omitempty"`              // Сумма НДС
	PurchaseOrder       *PurchaseOrder             `json:"purchaseOrder,omitempty"`       // Ссылка на связанный заказ поставщику
	FactureIn           *FactureIn                 `json:"factureIn,omitempty"`           // Ссылка на Счет-фактуру полученный, с которым связана эта Приемка
	InvoicesIn          Slice[InvoiceIn]           `json:"invoicesIn,omitempty"`          // Массив ссылок на связанные счета поставщиков
	AccountID           *uuid.UUID                 `json:"accountId,omitempty"`           // ID учётной записи
	Attributes          Slice[Attribute]           `json:"attributes,omitempty"`          // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (supply Supply) Clean() *Supply {
	if supply.Meta == nil {
		return nil
	}
	return &Supply{Meta: supply.Meta}
}

// AsOperation реализует интерфейс [OperationConverter].
func (supply Supply) AsOperation() *Operation {
	return newOperation(supply)
}

// AsTaskOperation реализует интерфейс [TaskOperationConverter].
func (supply Supply) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: supply.Meta}
}

// AsOperationOut реализует интерфейс [OperationOutConverter].
func (supply Supply) AsOperationOut() *Operation {
	return supply.AsOperation()
}

// GetOrganizationAccount возвращает Метаданные счета юрлица.
func (supply Supply) GetOrganizationAccount() AgentAccount {
	return Deref(supply.OrganizationAccount)
}

// GetApplicable возвращает Отметку о проведении.
func (supply Supply) GetApplicable() bool {
	return Deref(supply.Applicable)
}

// GetAgentAccount возвращает Метаданные счета контрагента.
func (supply Supply) GetAgentAccount() AgentAccount {
	return Deref(supply.AgentAccount)
}

// GetOverhead возвращает Накладные расходы.
func (supply Supply) GetOverhead() Overhead {
	return Deref(supply.Overhead)
}

// GetReturns возвращает Массив ссылок на связанные возвраты.
func (supply Supply) GetReturns() Slice[PurchaseReturn] {
	return supply.Returns
}

// GetCode возвращает Код Приемки.
func (supply Supply) GetCode() string {
	return Deref(supply.Code)
}

// GetContract возвращает Метаданные договора.
func (supply Supply) GetContract() Contract {
	return Deref(supply.Contract).getValue()
}

// GetCreated возвращает Дату создания.
func (supply Supply) GetCreated() time.Time {
	return Deref(supply.Created).Time()
}

// GetDeleted возвращает Момент последнего удаления Приемки.
func (supply Supply) GetDeleted() time.Time {
	return Deref(supply.Deleted).Time()
}

// GetDescription возвращает Комментарий Приемки.
func (supply Supply) GetDescription() string {
	return Deref(supply.Description)
}

// GetExternalCode возвращает Внешний код Приемки.
func (supply Supply) GetExternalCode() string {
	return Deref(supply.ExternalCode)
}

// GetFiles возвращает Метаданные массива Файлов.
func (supply Supply) GetFiles() MetaArray[File] {
	return Deref(supply.Files)
}

// GetGroup возвращает Отдел сотрудника.
func (supply Supply) GetGroup() Group {
	return Deref(supply.Group)
}

// GetID возвращает ID Приемки.
func (supply Supply) GetID() uuid.UUID {
	return Deref(supply.ID)
}

// GetIncomingDate возвращает Входящую дату.
func (supply Supply) GetIncomingDate() time.Time {
	return Deref(supply.IncomingDate).Time()
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (supply Supply) GetOwner() Employee {
	return Deref(supply.Owner)
}

// GetMeta возвращает Метаданные Приемки.
func (supply Supply) GetMeta() Meta {
	return Deref(supply.Meta)
}

// GetMoment возвращает Дату документа.
func (supply Supply) GetMoment() time.Time {
	return Deref(supply.Moment).Time()
}

// GetName возвращает Наименование Приемки.
func (supply Supply) GetName() string {
	return Deref(supply.Name)
}

// GetOrganization возвращает Метаданные юрлица.
func (supply Supply) GetOrganization() Organization {
	return Deref(supply.Organization)
}

// GetPayments возвращает Массив ссылок на связанные платежи.
func (supply Supply) GetPayments() Slice[Payment] {
	return supply.Payments
}

// GetAgent возвращает Метаданные контрагента.
func (supply Supply) GetAgent() Agent {
	return Deref(supply.Agent)
}

// GetIncomingNumber возвращает Входящий номер.
func (supply Supply) GetIncomingNumber() string {
	return Deref(supply.IncomingNumber)
}

// GetPayedSum возвращает Сумму входящих платежей по Приемке.
func (supply Supply) GetPayedSum() float64 {
	return Deref(supply.PayedSum)
}

// GetPositions возвращает Метаданные позиций Приемки.
func (supply Supply) GetPositions() MetaArray[SupplyPosition] {
	return Deref(supply.Positions)
}

// GetPrinted возвращает true, если документ напечатан.
func (supply Supply) GetPrinted() bool {
	return Deref(supply.Printed)
}

// GetProject возвращает Метаданные проекта.
func (supply Supply) GetProject() Project {
	return supply.Project.getValue()
}

// GetPublished возвращает true, если документ опубликован.
func (supply Supply) GetPublished() bool {
	return Deref(supply.Published)
}

// GetRate возвращает Валюту.
func (supply Supply) GetRate() Rate {
	return supply.Rate.getValue()
}

// GetShared возвращает флаг Общего доступа.
func (supply Supply) GetShared() bool {
	return Deref(supply.Shared)
}

// GetState возвращает Метаданные статуса Приемки.
func (supply Supply) GetState() State {
	return Deref(supply.State).getValue()
}

// GetStore возвращает Метаданные склада.
func (supply Supply) GetStore() Store {
	return Deref(supply.Store)
}

// GetSum возвращает Сумму Приемки в копейках.
func (supply Supply) GetSum() float64 {
	return Deref(supply.Sum)
}

// GetSyncID возвращает ID синхронизации.
func (supply Supply) GetSyncID() uuid.UUID {
	return Deref(supply.SyncID)
}

// GetUpdated возвращает Момент последнего обновления Списания.
func (supply Supply) GetUpdated() time.Time {
	return Deref(supply.Updated).Time()
}

// GetVatEnabled возвращает true, если учитывается НДС.
func (supply Supply) GetVatEnabled() bool {
	return Deref(supply.VatEnabled)
}

// GetVatIncluded возвращает true, если НДС включен в цену.
func (supply Supply) GetVatIncluded() bool {
	return Deref(supply.VatIncluded)
}

// GetVatSum возвращает Сумму НДС.
func (supply Supply) GetVatSum() float64 {
	return Deref(supply.VatSum)
}

// GetPurchaseOrder возвращает Ссылку на связанный заказ поставщику.
func (supply Supply) GetPurchaseOrder() PurchaseOrder {
	return Deref(supply.PurchaseOrder)
}

// GetFactureIn возвращает Ссылку на Счет-фактуру полученный, с которым связана эта Приемка.
func (supply Supply) GetFactureIn() FactureIn {
	return Deref(supply.FactureIn)
}

// GetInvoicesIn возвращает Массив ссылок на связанные счета поставщиков.
func (supply Supply) GetInvoicesIn() Slice[InvoiceIn] {
	return supply.InvoicesIn
}

// GetAccountID возвращает ID учётной записи.
func (supply Supply) GetAccountID() uuid.UUID {
	return Deref(supply.AccountID)
}

// GetAttributes возвращает Список метаданных доп. полей.
func (supply Supply) GetAttributes() Slice[Attribute] {
	return supply.Attributes
}

// SetOrganizationAccount устанавливает Метаданные счета юрлица.
func (supply *Supply) SetOrganizationAccount(organizationAccount *AgentAccount) *Supply {
	if organizationAccount != nil {
		supply.OrganizationAccount = organizationAccount.Clean()
	}
	return supply
}

// SetApplicable устанавливает Отметку о проведении.
func (supply *Supply) SetApplicable(applicable bool) *Supply {
	supply.Applicable = &applicable
	return supply
}

// SetAgentAccount устанавливает Метаданные счета контрагента.
func (supply *Supply) SetAgentAccount(agentAccount *AgentAccount) *Supply {
	if agentAccount != nil {
		supply.AgentAccount = agentAccount.Clean()
	}
	return supply
}

// SetOverhead устанавливает Накладные расходы.
//
// Если Позиции Оприходования не заданы, то накладные расходы нельзя задать.
func (supply *Supply) SetOverhead(overhead *Overhead) *Supply {
	supply.Overhead = overhead
	return supply
}

// SetReturns устанавливает Массив ссылок на связанные возвраты.
//
// Принимает множество объектов [PurchaseReturn].
func (supply *Supply) SetReturns(returns ...*PurchaseReturn) *Supply {
	supply.Returns.Push(returns...)
	return supply
}

// SetCode устанавливает Код Приемки.
func (supply *Supply) SetCode(code string) *Supply {
	supply.Code = &code
	return supply
}

// SetContract устанавливает Метаданные договора.
//
// Передача nil передаёт сброс значения (null).
func (supply *Supply) SetContract(contract *Contract) *Supply {
	supply.Contract = NewNullValue(contract)
	return supply
}

// SetDescription устанавливает Комментарий Приемки.
func (supply *Supply) SetDescription(description string) *Supply {
	supply.Description = &description
	return supply
}

// SetExternalCode устанавливает Внешний код Приемки.
func (supply *Supply) SetExternalCode(externalCode string) *Supply {
	supply.ExternalCode = &externalCode
	return supply
}

// SetFiles устанавливает Метаданные массива Файлов.
//
// Принимает множество объектов [File].
func (supply *Supply) SetFiles(files ...*File) *Supply {
	supply.Files = NewMetaArrayFrom(files)
	return supply
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (supply *Supply) SetGroup(group *Group) *Supply {
	if group != nil {
		supply.Group = group.Clean()
	}
	return supply
}

// SetIncomingDate устанавливает Входящую дату.
func (supply *Supply) SetIncomingDate(incomingDate time.Time) *Supply {
	supply.IncomingDate = NewTimestamp(incomingDate)
	return supply
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (supply *Supply) SetOwner(owner *Employee) *Supply {
	if owner != nil {
		supply.Owner = owner.Clean()
	}
	return supply
}

// SetMeta устанавливает Метаданные Приемки.
func (supply *Supply) SetMeta(meta *Meta) *Supply {
	supply.Meta = meta
	return supply
}

// SetMoment устанавливает Дату документа.
func (supply *Supply) SetMoment(moment time.Time) *Supply {
	supply.Moment = NewTimestamp(moment)
	return supply
}

// SetName устанавливает Наименование Приемки.
func (supply *Supply) SetName(name string) *Supply {
	supply.Name = &name
	return supply
}

// SetOrganization устанавливает Метаданные юрлица.
func (supply *Supply) SetOrganization(organization *Organization) *Supply {
	if organization != nil {
		supply.Organization = organization.Clean()
	}
	return supply
}

// SetPayments устанавливает Метаданные ссылок на связанные входящие платежи.
//
// Принимает множество объектов, реализующих интерфейс [PaymentConverter].
func (supply *Supply) SetPayments(payments ...PaymentConverter) *Supply {
	supply.Payments = NewPaymentsFrom(payments)
	return supply
}

// SetAgent устанавливает Метаданные Контрагента.
//
// Принимает [Counterparty] или [Organization].
func (supply *Supply) SetAgent(agent AgentOrganizationConverter) *Supply {
	if agent != nil {
		supply.Agent = agent.AsOrganizationAgent()
	}
	return supply
}

// SetIncomingNumber устанавливает Входящий номер.
func (supply *Supply) SetIncomingNumber(incomingNumber string) *Supply {
	supply.IncomingNumber = &incomingNumber
	return supply
}

// SetPositions устанавливает Метаданные позиций Приемки.
//
// Принимает множество объектов [SupplyPosition].
func (supply *Supply) SetPositions(positions ...*SupplyPosition) *Supply {
	supply.Positions = NewMetaArrayFrom(positions)
	return supply
}

// SetProject устанавливает Метаданные проекта.
//
// Передача nil передаёт сброс значения (null).
func (supply *Supply) SetProject(project *Project) *Supply {
	supply.Project = NewNullValue(project)
	return supply
}

// SetRate устанавливает Валюту.
//
// Передача nil передаёт сброс значения (null).
func (supply *Supply) SetRate(rate *Rate) *Supply {
	supply.Rate = NewNullValue(rate)
	return supply
}

// SetShared устанавливает флаг общего доступа.
func (supply *Supply) SetShared(shared bool) *Supply {
	supply.Shared = &shared
	return supply
}

// SetState устанавливает Метаданные статуса Приемки.
//
// Передача nil передаёт сброс значения (null).
func (supply *Supply) SetState(state *State) *Supply {
	supply.State = NewNullValue(state)
	return supply
}

// SetStore устанавливает Метаданные склада.
func (supply *Supply) SetStore(store *Store) *Supply {
	if store != nil {
		supply.Store = store.Clean()
	}
	return supply
}

// SetSyncID устанавливает ID синхронизации.
func (supply *Supply) SetSyncID(syncID uuid.UUID) *Supply {
	supply.SyncID = &syncID
	return supply
}

// SetVatEnabled устанавливает значение, учитывающее НДС для приемки.
func (supply *Supply) SetVatEnabled(vatEnabled bool) *Supply {
	supply.VatEnabled = &vatEnabled
	return supply
}

// SetVatIncluded устанавливает флаг включения НДС в цену.
func (supply *Supply) SetVatIncluded(vatIncluded bool) *Supply {
	supply.VatIncluded = &vatIncluded
	return supply
}

// SetPurchaseOrder устанавливает Ссылку на связанный заказ поставщику.
func (supply *Supply) SetPurchaseOrder(purchaseOrder *PurchaseOrder) *Supply {
	if purchaseOrder != nil {
		supply.PurchaseOrder = purchaseOrder.Clean()
	}
	return supply
}

// SetFactureIn устанавливает Ссылку на Счет-фактуру полученный, с которым связана эта Приемка.
func (supply *Supply) SetFactureIn(factureIn *FactureIn) *Supply {
	if factureIn != nil {
		supply.FactureIn = factureIn.Clean()
	}
	return supply
}

// SetInvoicesIn устанавливает Массив ссылок на связанные счета поставщиков.
//
// Принимает множество объектов [InvoiceIn].
func (supply *Supply) SetInvoicesIn(invoicesIn ...*InvoiceIn) *Supply {
	supply.InvoicesIn.Push(invoicesIn...)
	return supply
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (supply *Supply) SetAttributes(attributes ...*Attribute) *Supply {
	supply.Attributes.Push(attributes...)
	return supply
}

// String реализует интерфейс [fmt.Stringer].
func (supply Supply) String() string {
	return Stringify(supply)
}

// MetaType возвращает код сущности.
func (Supply) MetaType() MetaType {
	return MetaTypeSupply
}

// Update shortcut
func (supply *Supply) Update(ctx context.Context, client *Client, params ...*Params) (*Supply, *resty.Response, error) {
	return NewSupplyService(client).Update(ctx, supply.GetID(), supply, params...)
}

// Create shortcut
func (supply *Supply) Create(ctx context.Context, client *Client, params ...*Params) (*Supply, *resty.Response, error) {
	return NewSupplyService(client).Create(ctx, supply, params...)
}

// Delete shortcut
func (supply *Supply) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewSupplyService(client).Delete(ctx, supply)
}

// SupplyPosition Позиция Приемки.
//
// Код сущности: supplyposition
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-priemka-priemki-pozicii-priemki
type SupplyPosition struct {
	Quantity      *float64            `json:"quantity,omitempty"`      // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Pack          *Pack               `json:"pack,omitempty"`          // Упаковка Товара
	Country       *Country            `json:"country,omitempty"`       // Метаданные Страны
	Discount      *float64            `json:"discount,omitempty"`      // Процент скидки или наценки. Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%
	AccountID     *uuid.UUID          `json:"accountId,omitempty"`     // ID учётной записи
	ID            *uuid.UUID          `json:"id,omitempty"`            // ID позиции
	Assortment    *AssortmentPosition `json:"assortment,omitempty"`    // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	Price         *float64            `json:"price,omitempty"`         // Цена товара/услуги в копейках
	GTD           *GTD                `json:"gtd,omitempty"`           // ГТД
	Slot          *Slot               `json:"slot,omitempty"`          // Ячейка на складе
	Stock         *Stock              `json:"stock,omitempty"`         // Остатки и себестоимость позиции (указывается при наличии параметра запроса `fields=stock`)
	VatEnabled    *bool               `json:"vatEnabled,omitempty"`    // Включен ли НДС для позиции. С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
	Overhead      *float64            `json:"overhead,omitempty"`      // Накладные расходы. Если Позиции Приемки не заданы, то накладные расходы нельзя задать.
	Vat           *int                `json:"vat,omitempty"`           // НДС, которым облагается текущая позиция
	TrackingCodes Slice[TrackingCode] `json:"trackingCodes,omitempty"` // Коды маркировки товаров и транспортных упаковок
	Things        Slice[string]       `json:"things,omitempty"`        // Серийные номера. Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете. В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута.
}

// GetAccountID возвращает ID учётной записи.
func (supplyPosition SupplyPosition) GetAccountID() uuid.UUID {
	return Deref(supplyPosition.AccountID)
}

// GetAssortment возвращает Метаданные товара/услуги/серии/модификации, которую представляет собой позиция.
func (supplyPosition SupplyPosition) GetAssortment() AssortmentPosition {
	return Deref(supplyPosition.Assortment)
}

// GetCountry возвращает Метаданные Страны.
func (supplyPosition SupplyPosition) GetCountry() Country {
	return Deref(supplyPosition.Country)
}

// GetDiscount возвращает Процент скидки или наценки.
//
// Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%.
func (supplyPosition SupplyPosition) GetDiscount() float64 {
	return Deref(supplyPosition.Discount)
}

// GetGTD возвращает ГТД.
func (supplyPosition SupplyPosition) GetGTD() GTD {
	return Deref(supplyPosition.GTD)
}

// GetGTDName возвращает Номер ГТД.
func (supplyPosition SupplyPosition) GetGTDName() string {
	return Deref(supplyPosition.GTD).GetName()
}

// GetID возвращает ID позиции.
func (supplyPosition SupplyPosition) GetID() uuid.UUID {
	return Deref(supplyPosition.ID)
}

// GetPack возвращает Упаковку Товара.
func (supplyPosition SupplyPosition) GetPack() Pack {
	return Deref(supplyPosition.Pack)
}

// GetPrice возвращает Цену товара/услуги в копейках.
func (supplyPosition SupplyPosition) GetPrice() float64 {
	return Deref(supplyPosition.Price)
}

// GetQuantity возвращает Количество товаров/услуг данного вида в позиции.
//
// Если позиция - товар, у которого включен учет по серийным номерам,
// то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
func (supplyPosition SupplyPosition) GetQuantity() float64 {
	return Deref(supplyPosition.Quantity)
}

// GetSlot возвращает Ячейку на складе.
func (supplyPosition SupplyPosition) GetSlot() Slot {
	return Deref(supplyPosition.Slot)
}

// GetThings возвращает Серийные номера.
//
// Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете.
// В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута.
func (supplyPosition SupplyPosition) GetThings() Slice[string] {
	return supplyPosition.Things
}

// GetTrackingCodes возвращает Коды маркировки товаров и транспортных упаковок.
func (supplyPosition SupplyPosition) GetTrackingCodes() Slice[TrackingCode] {
	return supplyPosition.TrackingCodes
}

// GetOverhead возвращает Накладные расходы.
func (supplyPosition SupplyPosition) GetOverhead() float64 {
	return Deref(supplyPosition.Overhead)
}

// GetVat возвращает НДС, которым облагается текущая позиция.
func (supplyPosition SupplyPosition) GetVat() int {
	return Deref(supplyPosition.Vat)
}

// GetVatEnabled возвращает true, если учитывается НДС.
//
// С помощью этого флага для товара можно выставлять НДС = 0 или НДС = "без НДС".
//
// (vat = 0, vatEnabled = false) -> vat = "без НДС",
//
// (vat = 0, vatEnabled = true) -> vat = 0%.
func (supplyPosition SupplyPosition) GetVatEnabled() bool {
	return Deref(supplyPosition.VatEnabled)
}

// GetStock возвращает Остатки и себестоимость позиции (указывается при наличии параметра запроса `fields=stock`).
func (supplyPosition SupplyPosition) GetStock() Stock {
	return Deref(supplyPosition.Stock)
}

// SetAssortment устанавливает Метаданные товара/услуги/серии/модификации, которую представляет собой позиция.
//
// Принимает объект, реализующий интерфейс [AssortmentConverter].
func (supplyPosition *SupplyPosition) SetAssortment(assortment AssortmentConverter) *SupplyPosition {
	if assortment != nil {
		supplyPosition.Assortment = assortment.AsAssortment()
	}
	return supplyPosition
}

// SetCountry устанавливает Метаданные Страны.
func (supplyPosition *SupplyPosition) SetCountry(country *Country) *SupplyPosition {
	if country != nil {
		supplyPosition.Country = country.Clean()
	}
	return supplyPosition
}

// SetDiscount устанавливает Процент скидки или наценки.
//
// Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%.
func (supplyPosition *SupplyPosition) SetDiscount(discount float64) *SupplyPosition {
	supplyPosition.Discount = &discount
	return supplyPosition
}

// SetGTD устанавливает ГТД.
func (supplyPosition *SupplyPosition) SetGTD(gtd *GTD) *SupplyPosition {
	supplyPosition.GTD = gtd
	return supplyPosition
}

// SetPack устанавливает Упаковку Товара.
func (supplyPosition *SupplyPosition) SetPack(pack *Pack) *SupplyPosition {
	supplyPosition.Pack = pack
	return supplyPosition
}

// SetPrice устанавливает Цену товара/услуги в копейках.
func (supplyPosition *SupplyPosition) SetPrice(price float64) *SupplyPosition {
	supplyPosition.Price = &price
	return supplyPosition
}

// SetQuantity устанавливает Количество товаров данного вида в позиции.
//
// Если позиция - товар, у которого включен учет по серийным номерам,
// то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
func (supplyPosition *SupplyPosition) SetQuantity(quantity float64) *SupplyPosition {
	supplyPosition.Quantity = &quantity
	return supplyPosition
}

// SetSlot устанавливает Ячейку на складе.
func (supplyPosition *SupplyPosition) SetSlot(slot *Slot) *SupplyPosition {
	if slot != nil {
		supplyPosition.Slot = slot.Clean()
	}
	return supplyPosition
}

// SetThings устанавливает Серийные номера.
//
// Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете.
// В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута.
//
// Принимает множество string.
func (supplyPosition *SupplyPosition) SetThings(things ...string) *SupplyPosition {
	supplyPosition.Things = NewSliceFrom(things)
	return supplyPosition
}

// SetTrackingCodes устанавливает Коды маркировки товаров и транспортных упаковок.
//
// Принимает множество объектов [TrackingCode].
func (supplyPosition *SupplyPosition) SetTrackingCodes(trackingCodes ...*TrackingCode) *SupplyPosition {
	supplyPosition.TrackingCodes.Push(trackingCodes...)
	return supplyPosition
}

// SetVat устанавливает НДС, которым облагается текущая позиция.
func (supplyPosition *SupplyPosition) SetVat(vat int) *SupplyPosition {
	supplyPosition.Vat = &vat
	return supplyPosition
}

// SetVatEnabled устанавливает значение, учитывающее НДС для текущей позиции.
func (supplyPosition *SupplyPosition) SetVatEnabled(vatEnabled bool) *SupplyPosition {
	supplyPosition.VatEnabled = &vatEnabled
	return supplyPosition
}

// String реализует интерфейс [fmt.Stringer].
func (supplyPosition SupplyPosition) String() string {
	return Stringify(supplyPosition)
}

// MetaType возвращает код сущности.
func (SupplyPosition) MetaType() MetaType {
	return MetaTypeSupplyPosition
}

// SupplyService описывает методы сервиса для работы с приёмками.
type SupplyService interface {
	// GetList выполняет запрос на получение списка приемок.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[Supply], *resty.Response, error)

	// GetListAll выполняет запрос на получение всех приемок в виде списка.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает список объектов.
	GetListAll(ctx context.Context, params ...*Params) (*Slice[Supply], *resty.Response, error)

	// Create выполняет запрос на создание приемки.
	// Обязательные поля для заполнения:
	//	- organization (Ссылка на ваше юрлицо)
	//	- agent (Ссылка на контрагента (поставщика))
	//	- store (Ссылка на склад)
	// Принимает контекст, приемку и опционально объект параметров запроса Params.
	// Возвращает созданную приемку.
	Create(ctx context.Context, supply *Supply, params ...*Params) (*Supply, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение приемок.
	// Изменяемые приемки должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список приемок и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых приемок.
	CreateUpdateMany(ctx context.Context, supplyList Slice[Supply], params ...*Params) (*Slice[Supply], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление приемок.
	// Принимает контекст и множество приемок.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*Supply) (*DeleteManyResponse, *resty.Response, error)

	// DeleteByID выполняет запрос на удаление приемки по ID.
	// Принимает контекст и ID приемки.
	// Возвращает «true» в случае успешного удаления приемки.
	DeleteByID(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// Delete выполняет запрос на удаление приемки.
	// Принимает контекст и приемку.
	// Возвращает «true» в случае успешного удаления приемки.
	Delete(ctx context.Context, entity *Supply) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельной приемки по ID.
	// Принимает контекст, ID приемки и опционально объект параметров запроса Params.
	// Возвращает приемку.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Supply, *resty.Response, error)

	// Update выполняет запрос на изменение приемки.
	// Принимает контекст, приемку и опционально объект параметров запроса Params.
	// Возвращает изменённую приемку.
	Update(ctx context.Context, id uuid.UUID, supply *Supply, params ...*Params) (*Supply, *resty.Response, error)

	// Template выполняет запрос на получение предзаполненной приемки со стандартными полями без связи с какими-либо другими документами.
	// Принимает контекст.
	// Возвращает предзаполненную приемку.
	Template(ctx context.Context) (*Supply, *resty.Response, error)

	// TemplateBased выполняет запрос на получение шаблона приемки на основе других документов.
	// Основание, на котором может быть создана:
	//	- Счет поставщика (InvoiceIn)
	//	- Заказ поставщику (PurchaseOrder)
	// Принимает контекст и множество документов из списка выше.
	// Возвращает предзаполненную приемку на основании переданных документов.
	TemplateBased(ctx context.Context, basedOn ...MetaOwner) (*Supply, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных приемок.
	// Принимает контекст.
	// Возвращает объект метаданных MetaAttributesStatesSharedWrapper.
	GetMetadata(ctx context.Context) (*MetaAttributesStatesSharedWrapper, *resty.Response, error)

	// GetPositionList выполняет запрос на получение списка позиций документа.
	// Принимает контекст, ID документа и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetPositionList(ctx context.Context, id uuid.UUID, params ...*Params) (*List[SupplyPosition], *resty.Response, error)

	GetPositionListAll(ctx context.Context, id uuid.UUID, params ...*Params) (*Slice[SupplyPosition], *resty.Response, error)

	// GetPositionByID выполняет запрос на получение отдельной позиции документа по ID.
	// Принимает контекст, ID документа, ID позиции и опционально объект параметров запроса Params.
	// Возвращает найденную позицию.
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params ...*Params) (*SupplyPosition, *resty.Response, error)

	// UpdatePosition выполняет запрос на изменение позиции документа.
	// Принимает контекст, ID документа, ID позиции, позицию документа и опционально объект параметров запроса Params.
	// Возвращает изменённую позицию.
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *SupplyPosition, params ...*Params) (*SupplyPosition, *resty.Response, error)

	// CreatePosition выполняет запрос на добавление позиции документа.
	// Принимает контекст, ID документа, позицию документа и опционально объект параметров запроса Params.
	// Возвращает добавленную позицию.
	CreatePosition(ctx context.Context, id uuid.UUID, position *SupplyPosition, params ...*Params) (*SupplyPosition, *resty.Response, error)

	// CreatePositionMany выполняет запрос на массовое добавление позиций документа.
	// Принимает контекст, ID документа и множество позиций.
	// Возвращает список добавленных позиций.
	CreatePositionMany(ctx context.Context, id uuid.UUID, positions ...*SupplyPosition) (*Slice[SupplyPosition], *resty.Response, error)

	// DeletePosition выполняет запрос на удаление позиции документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает «true» в случае успешного удаления позиции.
	DeletePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (bool, *resty.Response, error)

	// DeletePositionMany выполняет запрос на массовое удаление позиций документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeletePositionMany(ctx context.Context, id uuid.UUID, positions ...*SupplyPosition) (*DeleteManyResponse, *resty.Response, error)

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

	// PrintDocument TODO: выполняет запрос на печать отдельного документа по шаблону печатной формы.
	// Принимает контекст, и ID документа и объект PrintDocumentArg.
	// Возвращает файл
	PrintDocument(ctx context.Context, id uuid.UUID, PrintDocumentArg *PrintDocumentArg) (*PrintFile, *resty.Response, error)

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

	// GetNamedFilterList выполняет запрос на получение списка фильтров.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetNamedFilterList(ctx context.Context, params ...*Params) (*List[NamedFilter], *resty.Response, error)

	// GetNamedFilterByID выполняет запрос на получение отдельного фильтра по ID.
	// Принимает контекст и ID фильтра.
	// Возвращает найденный фильтр.
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)

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

	// GetBySyncID выполняет запрос на получение отдельного документа по syncID.
	// Принимает контекст и syncID документа.
	// Возвращает найденный документ.
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*Supply, *resty.Response, error)

	// DeleteBySyncID выполняет запрос на удаление документа по syncID.
	// Принимает контекст и syncID документа.
	// Возвращает «true» в случае успешного удаления документа.
	DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error)

	// MoveToTrash выполняет запрос на перемещение документа с указанным ID в корзину.
	// Принимает контекст и ID документа.
	// Возвращает «true» в случае успешного перемещения в корзину.
	MoveToTrash(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// Evaluate выполняет запрос на получение шаблона документа с автозаполнением.
	// Принимает контекст, документ и множество значений Evaluate.
	//
	// Возможные значения типа Evaluate:
	//	- EvaluateDiscount – скидки
	//	- EvaluatePrice    – цены
	//	- EvaluateVat      – ндс
	//	- EvaluateCost     – себестоимость
	// Возвращает шаблон документа с автозаполнением.
	Evaluate(ctx context.Context, entity *Supply, evaluate ...Evaluate) (*Supply, *resty.Response, error)
}

const (
	EndpointSupply = EndpointEntity + string(MetaTypeSupply)
)

// NewSupplyService принимает [Client] и возвращает сервис для работы с приёмками.
func NewSupplyService(client *Client) SupplyService {
	return newMainService[Supply, SupplyPosition, MetaAttributesStatesSharedWrapper, any](client, EndpointSupply)
}
