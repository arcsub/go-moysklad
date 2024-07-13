package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"time"
)

// Demand Отгрузка.
//
// Код сущности: demand
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-otgruzka
type Demand struct {
	AccountID               *uuid.UUID                 `json:"accountId,omitempty"`               // ID учётной записи
	Agent                   *Agent                     `json:"agent,omitempty"`                   // Метаданные контрагента
	AgentAccount            *AgentAccount              `json:"agentAccount,omitempty"`            // Метаданные счета контрагента
	Applicable              *bool                      `json:"applicable,omitempty"`              // Отметка о проведении
	Code                    *string                    `json:"code,omitempty"`                    // Код Отгрузки
	Contract                *NullValue[Contract]       `json:"contract,omitempty"`                // Метаданные договора
	Created                 *Timestamp                 `json:"created,omitempty"`                 // Дата создания
	Deleted                 *Timestamp                 `json:"deleted,omitempty"`                 // Момент последнего удаления Отгрузки
	Description             *string                    `json:"description,omitempty"`             // Комментарий Отгрузки
	ExternalCode            *string                    `json:"externalCode,omitempty"`            // Внешний код Отгрузки
	Files                   *MetaArray[File]           `json:"files,omitempty"`                   // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group                   *Group                     `json:"group,omitempty"`                   // Отдел сотрудника
	ID                      *uuid.UUID                 `json:"id,omitempty"`                      // ID Отгрузки
	Meta                    *Meta                      `json:"meta,omitempty"`                    // Метаданные Отгрузки
	Moment                  *Timestamp                 `json:"moment,omitempty"`                  // Дата документа
	Name                    *string                    `json:"name,omitempty"`                    // Наименование Отгрузки
	Organization            *Organization              `json:"organization,omitempty"`            // Метаданные юрлица
	OrganizationAccount     *AgentAccount              `json:"organizationAccount,omitempty"`     // Метаданные счета юрлица
	Overhead                *Overhead                  `json:"overhead,omitempty"`                // Накладные расходы. Если Позиции Отгрузки не заданы, то накладные расходы нельзя задать
	Owner                   *Employee                  `json:"owner,omitempty"`                   // Метаданные владельца (Сотрудника)
	PayedSum                *float64                   `json:"payedSum,omitempty"`                // Сумма входящих платежей по Отгрузке
	Positions               *MetaArray[DemandPosition] `json:"positions,omitempty"`               // Метаданные позиций Отгрузки
	Printed                 *bool                      `json:"printed,omitempty"`                 // Напечатан ли документ
	Project                 *NullValue[Project]        `json:"project,omitempty"`                 // Метаданные проекта
	Published               *bool                      `json:"published,omitempty"`               // Опубликован ли документ
	Rate                    *NullValue[Rate]           `json:"rate,omitempty"`                    // Валюта
	SalesChannel            *NullValue[SalesChannel]   `json:"salesChannel,omitempty"`            // Метаданные канала продаж
	Shared                  *bool                      `json:"shared,omitempty"`                  // Общий доступ
	ShipmentAddress         *string                    `json:"shipmentAddress,omitempty"`         // Адрес доставки Отгрузки
	ShipmentAddressFull     *Address                   `json:"shipmentAddressFull,omitempty"`     // Адрес доставки Отгрузки с детализацией по отдельным полям.
	State                   *NullValue[State]          `json:"state,omitempty"`                   // Метаданные статуса Отгрузки
	Store                   *Store                     `json:"store,omitempty"`                   // Метаданные склада
	Sum                     *float64                   `json:"sum,omitempty"`                     // Сумма Отгрузки в копейках
	SyncID                  *uuid.UUID                 `json:"syncId,omitempty"`                  // ID синхронизации
	Updated                 *Timestamp                 `json:"updated,omitempty"`                 // Момент последнего обновления Отгрузки
	VatEnabled              *bool                      `json:"vatEnabled,omitempty"`              // Учитывается ли НДС
	VatIncluded             *bool                      `json:"vatIncluded,omitempty"`             // Включен ли НДС в цену
	VatSum                  *float64                   `json:"vatSum,omitempty"`                  // Сумма НДС
	CustomerOrder           *CustomerOrder             `json:"customerOrder,omitempty"`           // Ссылка на Заказ Покупателя, с которым связана эта Отгрузка
	FactureOut              *FactureOut                `json:"factureOut,omitempty"`              // Ссылка на Счет-фактуру выданный, с которым связана эта Отгрузка
	Returns                 Slice[SalesReturn]         `json:"returns,omitempty"`                 // Массив ссылок на связанные возвраты
	Payments                Slice[Payment]             `json:"payments,omitempty"`                // Массив ссылок на связанные платежи
	InvoicesOut             Slice[InvoiceOut]          `json:"invoicesOut,omitempty"`             // Массив ссылок на связанные счета покупателям
	CargoName               *string                    `json:"cargoName,omitempty"`               // Наименование груза
	Carrier                 *Agent                     `json:"carrier,omitempty"`                 // Метаданные перевозчика (контрагент или юрлицо)
	Consignee               *Agent                     `json:"consignee,omitempty"`               // Метаданные грузополучателя (контрагент или юрлицо)
	GoodPackQuantity        *int                       `json:"goodPackQuantity,omitempty"`        // Всего мест
	ShippingInstructions    *string                    `json:"shippingInstructions,omitempty"`    // Указания грузоотправителя
	StateContractID         *string                    `json:"stateContractId,omitempty"`         // Идентификатор государственного контракта, договора (соглашения)
	TransportFacility       *string                    `json:"transportFacility,omitempty"`       // Транспортное средство
	TransportFacilityNumber *string                    `json:"transportFacilityNumber,omitempty"` // Номер автомобиля
	Attributes              Slice[Attribute]           `json:"attributes,omitempty"`              // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (demand Demand) Clean() *Demand {
	if demand.Meta == nil {
		return nil
	}
	return &Demand{Meta: demand.Meta}
}

// AsOperation реализует интерфейс [OperationConverter].
func (demand Demand) AsOperation() *Operation {
	return newOperation(demand)
}

// AsTaskOperation реализует интерфейс [TaskOperationConverter].
func (demand Demand) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: demand.Meta}
}

// AsOperationIn реализует интерфейс [OperationInConverter].
func (demand Demand) AsOperationIn() *Operation {
	return demand.AsOperation()
}

// GetAccountID возвращает ID учётной записи.
func (demand Demand) GetAccountID() uuid.UUID {
	return Deref(demand.AccountID)
}

// GetAgent возвращает Метаданные контрагента (контрагент или юрлицо).
func (demand Demand) GetAgent() Agent {
	return Deref(demand.Agent)
}

// GetAgentAccount возвращает Метаданные счета контрагента.
func (demand Demand) GetAgentAccount() AgentAccount {
	return Deref(demand.AgentAccount)
}

// GetApplicable возвращает Отметку о проведении.
func (demand Demand) GetApplicable() bool {
	return Deref(demand.Applicable)
}

// GetCode возвращает Код Отгрузки.
func (demand Demand) GetCode() string {
	return Deref(demand.Code)
}

// GetContract возвращает Метаданные договора.
func (demand Demand) GetContract() Contract {
	return demand.Contract.getValue()
}

// GetCreated возвращает Дату создания.
func (demand Demand) GetCreated() time.Time {
	return Deref(demand.Created).Time()
}

// GetDeleted возвращает Момент последнего удаления Отгрузки.
func (demand Demand) GetDeleted() time.Time {
	return Deref(demand.Deleted).Time()
}

// GetDescription возвращает Комментарий Отгрузки.
func (demand Demand) GetDescription() string {
	return Deref(demand.Description)
}

// GetExternalCode возвращает Внешний код Отгрузки.
func (demand Demand) GetExternalCode() string {
	return Deref(demand.ExternalCode)
}

// GetFiles возвращает Метаданные массива Файлов.
func (demand Demand) GetFiles() MetaArray[File] {
	return Deref(demand.Files)
}

// GetGroup возвращает Отдел сотрудника.
func (demand Demand) GetGroup() Group {
	return Deref(demand.Group)
}

// GetID возвращает ID Отгрузки.
func (demand Demand) GetID() uuid.UUID {
	return Deref(demand.ID)
}

// GetMeta возвращает Метаданные Отгрузки.
func (demand Demand) GetMeta() Meta {
	return Deref(demand.Meta)
}

// GetMoment возвращает Дату документа.
func (demand Demand) GetMoment() time.Time {
	return Deref(demand.Moment).Time()
}

// GetName возвращает Наименование Отгрузки.
func (demand Demand) GetName() string {
	return Deref(demand.Name)
}

// GetOrganization возвращает Метаданные юрлица.
func (demand Demand) GetOrganization() Organization {
	return Deref(demand.Organization)
}

// GetOrganizationAccount возвращает Метаданные счета юрлица.
func (demand Demand) GetOrganizationAccount() AgentAccount {
	return Deref(demand.OrganizationAccount)
}

// GetOverhead возвращает Накладные расходы.
func (demand Demand) GetOverhead() Overhead {
	return Deref(demand.Overhead)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (demand Demand) GetOwner() Employee {
	return Deref(demand.Owner)
}

// GetPayedSum возвращает Оплаченную сумму.
func (demand Demand) GetPayedSum() float64 {
	return Deref(demand.PayedSum)
}

// GetPositions возвращает Метаданные позиций Отгрузки.
func (demand Demand) GetPositions() MetaArray[DemandPosition] {
	return Deref(demand.Positions)
}

// GetPrinted возвращает true, если документ напечатан.
func (demand Demand) GetPrinted() bool {
	return Deref(demand.Printed)
}

// GetProject возвращает Метаданные проекта.
func (demand Demand) GetProject() Project {
	return demand.Project.getValue()
}

// GetPublished возвращает true, если документ опубликован.
func (demand Demand) GetPublished() bool {
	return Deref(demand.Published)
}

// GetRate возвращает Валюту.
func (demand Demand) GetRate() Rate {
	return demand.Rate.getValue()
}

// GetSalesChannel возвращает Метаданные канала продаж.
func (demand Demand) GetSalesChannel() SalesChannel {
	return demand.SalesChannel.getValue()
}

// GetShared возвращает флаг Общего доступа.
func (demand Demand) GetShared() bool {
	return Deref(demand.Shared)
}

// GetShipmentAddress возвращает Адрес доставки Отгрузки.
func (demand Demand) GetShipmentAddress() string {
	return Deref(demand.ShipmentAddress)
}

// GetShipmentAddressFull возвращает Адрес доставки Отгрузки с детализацией по отдельным полям.
func (demand Demand) GetShipmentAddressFull() Address {
	return Deref(demand.ShipmentAddressFull)
}

// GetState возвращает Метаданные статуса Отгрузки.
func (demand Demand) GetState() State {
	return demand.State.getValue()
}

// GetStore возвращает Метаданные склада.
func (demand Demand) GetStore() Store {
	return Deref(demand.Store)
}

// GetSum возвращает Сумму Отгрузки в копейках.
func (demand Demand) GetSum() float64 {
	return Deref(demand.Sum)
}

// GetSyncID возвращает ID синхронизации.
func (demand Demand) GetSyncID() uuid.UUID {
	return Deref(demand.SyncID)
}

// GetUpdated возвращает Момент последнего обновления Отгрузки.
func (demand Demand) GetUpdated() time.Time {
	return Deref(demand.Updated).Time()
}

// GetVatEnabled возвращает true, если учитывается НДС.
func (demand Demand) GetVatEnabled() bool {
	return Deref(demand.VatEnabled)
}

// GetVatIncluded возвращает true, если НДС включен в цену.
func (demand Demand) GetVatIncluded() bool {
	return Deref(demand.VatIncluded)
}

// GetVatSum возвращает Сумму НДС.
func (demand Demand) GetVatSum() float64 {
	return Deref(demand.VatSum)
}

// GetCustomerOrder возвращает Ссылку на Заказ Покупателя, с которым связана эта Отгрузка.
func (demand Demand) GetCustomerOrder() CustomerOrder {
	return Deref(demand.CustomerOrder)
}

// GetFactureOut возвращает Ссылку на выданный Счет-фактуру, с которым связана эта Отгрузка.
func (demand Demand) GetFactureOut() FactureOut {
	return Deref(demand.FactureOut)
}

// GetReturns возвращает Массив ссылок на связанные возвраты.
func (demand Demand) GetReturns() Slice[SalesReturn] {
	return demand.Returns
}

// GetPayments возвращает Массив ссылок на связанные платежи.
func (demand Demand) GetPayments() Slice[Payment] {
	return demand.Payments
}

// GetInvoicesOut возвращает Массив ссылок на связанные счета покупателям.
func (demand Demand) GetInvoicesOut() Slice[InvoiceOut] {
	return demand.InvoicesOut
}

// GetCargoName возвращает Наименование груза.
func (demand Demand) GetCargoName() string {
	return Deref(demand.CargoName)
}

// GetCarrier возвращает Метаданные перевозчика (контрагент или юрлицо).
func (demand Demand) GetCarrier() Agent {
	return Deref(demand.Carrier)
}

// GetConsignee возвращает Метаданные грузополучателя (контрагент или юрлицо).
func (demand Demand) GetConsignee() Agent {
	return Deref(demand.Consignee)
}

// GetGoodPackQuantity возвращает количество мест.
func (demand Demand) GetGoodPackQuantity() int {
	return Deref(demand.GoodPackQuantity)
}

// GetShippingInstructions возвращает Указания грузоотправителя.
func (demand Demand) GetShippingInstructions() string {
	return Deref(demand.ShippingInstructions)
}

// GetStateContractID возвращает Идентификатор государственного контракта, договора (соглашения).
func (demand Demand) GetStateContractID() string {
	return Deref(demand.StateContractID)
}

// GetTransportFacility возвращает Транспортное средство.
func (demand Demand) GetTransportFacility() string {
	return Deref(demand.TransportFacility)
}

// GetTransportFacilityNumber возвращает Номер автомобиля.
func (demand Demand) GetTransportFacilityNumber() string {
	return Deref(demand.TransportFacilityNumber)
}

// GetAttributes возвращает Список метаданных доп. полей.
func (demand Demand) GetAttributes() Slice[Attribute] {
	return demand.Attributes
}

// SetAgent устанавливает Метаданные Контрагента.
//
// Принимает [Counterparty] или [Organization].
func (demand *Demand) SetAgent(agent AgentOrganizationConverter) *Demand {
	if agent != nil {
		demand.Agent = agent.AsOrganizationAgent()
	}
	return demand
}

// SetAgentAccount устанавливает Метаданные счета контрагента.
func (demand *Demand) SetAgentAccount(agentAccount *AgentAccount) *Demand {
	if agentAccount != nil {
		demand.AgentAccount = agentAccount.Clean()
	}
	return demand
}

// SetApplicable устанавливает Отметку о проведении.
func (demand *Demand) SetApplicable(applicable bool) *Demand {
	demand.Applicable = &applicable
	return demand
}

// SetCode устанавливает Код Отгрузки.
func (demand *Demand) SetCode(code string) *Demand {
	demand.Code = &code
	return demand
}

// SetContract устанавливает Метаданные договора.
//
// Передача nil передаёт сброс значения (null).
func (demand *Demand) SetContract(contract *Contract) *Demand {
	demand.Contract = NewNullValue(contract)
	return demand
}

// SetDescription устанавливает Комментарий Отгрузки.
func (demand *Demand) SetDescription(description string) *Demand {
	demand.Description = &description
	return demand
}

// SetExternalCode устанавливает Внешний код Отгрузки.
func (demand *Demand) SetExternalCode(externalCode string) *Demand {
	demand.ExternalCode = &externalCode
	return demand
}

// SetFiles устанавливает Метаданные массива Файлов.
//
// Принимает множество объектов [File].
func (demand *Demand) SetFiles(files ...*File) *Demand {
	demand.Files = NewMetaArrayFrom(files)
	return demand
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (demand *Demand) SetGroup(group *Group) *Demand {
	if group != nil {
		demand.Group = group.Clean()
	}
	return demand
}

// SetMeta устанавливает Метаданные Отгрузки.
func (demand *Demand) SetMeta(meta *Meta) *Demand {
	demand.Meta = meta
	return demand
}

// SetMoment устанавливает Дату документа.
func (demand *Demand) SetMoment(moment time.Time) *Demand {
	demand.Moment = NewTimestamp(moment)
	return demand
}

// SetName устанавливает Наименование Отгрузки.
func (demand *Demand) SetName(name string) *Demand {
	demand.Name = &name
	return demand
}

// SetOrganization устанавливает Метаданные юрлица.
func (demand *Demand) SetOrganization(organization *Organization) *Demand {
	if organization != nil {
		demand.Organization = organization
	}
	return demand
}

// SetOrganizationAccount устанавливает Метаданные счета юрлица.
func (demand *Demand) SetOrganizationAccount(organizationAccount *AgentAccount) *Demand {
	if organizationAccount != nil {
		demand.OrganizationAccount = organizationAccount
	}
	return demand
}

// SetOverhead устанавливает Накладные расходы. Если Позиции Отгрузки не заданы, то накладные расходы нельзя задать.
func (demand *Demand) SetOverhead(overhead *Overhead) *Demand {
	if overhead != nil {
		demand.Overhead = overhead
	}
	return demand
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (demand *Demand) SetOwner(owner *Employee) *Demand {
	if owner != nil {
		demand.Owner = owner.Clean()
	}
	return demand
}

// SetPositions устанавливает Метаданные позиций Отгрузки.
//
// Принимает множество объектов [DemandPosition].
func (demand *Demand) SetPositions(positions ...*DemandPosition) *Demand {
	demand.Positions = NewMetaArrayFrom(positions)
	return demand
}

// SetProject устанавливает Метаданные проекта.
//
// Передача nil передаёт сброс значения (null).
func (demand *Demand) SetProject(project *Project) *Demand {
	demand.Project = NewNullValue(project)
	return demand
}

// SetRate устанавливает Валюту.
//
// Передача nil передаёт сброс значения (null).
func (demand *Demand) SetRate(rate *Rate) *Demand {
	demand.Rate = NewNullValue(rate)
	return demand
}

// SetSalesChannel устанавливает Метаданные канала продаж.
//
// Передача nil передаёт сброс значения (null).
func (demand *Demand) SetSalesChannel(salesChannel *SalesChannel) *Demand {
	demand.SalesChannel = NewNullValue(salesChannel)
	return demand
}

// SetShared устанавливает флаг общего доступа.
func (demand *Demand) SetShared(shared bool) *Demand {
	demand.Shared = &shared
	return demand
}

// SetShipmentAddress устанавливает Адрес доставки Отгрузки.
func (demand *Demand) SetShipmentAddress(shipmentAddress string) *Demand {
	demand.ShipmentAddress = &shipmentAddress
	return demand
}

// SetShipmentAddressFull устанавливает Адрес доставки Отгрузки с детализацией по отдельным полям.
//
// Передача nil передаёт сброс значения (null).
func (demand *Demand) SetShipmentAddressFull(shipmentAddressFull *Address) *Demand {
	if shipmentAddressFull == nil {
		demand.SetShipmentAddress("")
	} else {
		demand.ShipmentAddressFull = shipmentAddressFull
	}
	return demand
}

// SetState устанавливает Метаданные статуса Отгрузки.
//
// Передача nil передаёт сброс значения (null).
func (demand *Demand) SetState(state *State) *Demand {
	demand.State = NewNullValue(state)
	return demand
}

// SetStore устанавливает Метаданные склада.
func (demand *Demand) SetStore(store *Store) *Demand {
	if store != nil {
		demand.Store = store.Clean()
	}
	return demand
}

// SetSyncID устанавливает ID синхронизации.
func (demand *Demand) SetSyncID(syncID uuid.UUID) *Demand {
	demand.SyncID = &syncID
	return demand
}

// SetVatEnabled устанавливает значение, учитывающее НДС для Заказа покупателя.
func (demand *Demand) SetVatEnabled(vatEnabled bool) *Demand {
	demand.VatEnabled = &vatEnabled
	return demand
}

// SetVatIncluded устанавливает флаг включения НДС в цену.
func (demand *Demand) SetVatIncluded(vatIncluded bool) *Demand {
	demand.VatIncluded = &vatIncluded
	return demand
}

// SetCustomerOrder устанавливает Ссылку на Заказ Покупателя, с которым связана эта Отгрузка.
func (demand *Demand) SetCustomerOrder(customerOrder *CustomerOrder) *Demand {
	if customerOrder != nil {
		demand.CustomerOrder = customerOrder.Clean()
	}
	return demand
}

// SetFactureOut устанавливает Метаданные Счет-фактуры выданного, с которым связана эта Отгрузка.
func (demand *Demand) SetFactureOut(factureOut *FactureOut) *Demand {
	if factureOut != nil {
		demand.FactureOut = factureOut.Clean()
	}
	return demand
}

// SetReturns устанавливает Массив ссылок на связанные возвраты.
//
// Принимает множество объектов [SalesReturn].
func (demand *Demand) SetReturns(returns ...*SalesReturn) *Demand {
	demand.Returns.Push(returns...)
	return demand
}

// SetPayments устанавливает Метаданные ссылок на связанные платежи.
//
// Принимает множество объектов, реализующих интерфейс [PaymentConverter].
func (demand *Demand) SetPayments(payments ...PaymentConverter) *Demand {
	demand.Payments = NewPaymentsFrom(payments)
	return demand
}

// SetInvoicesOut устанавливает Массив ссылок на связанные счета покупателям.
//
// Принимает множество объектов [InvoiceOut].
func (demand *Demand) SetInvoicesOut(invoicesOut ...*InvoiceOut) *Demand {
	demand.InvoicesOut.Push(invoicesOut...)
	return demand
}

// SetCargoName устанавливает Наименование груза.
func (demand *Demand) SetCargoName(cargoName string) *Demand {
	demand.CargoName = &cargoName
	return demand
}

// SetCarrier устанавливает Метаданные перевозчика (контрагент или юрлицо).
//
// Принимает [Counterparty] или [Organization].
func (demand *Demand) SetCarrier(carrier AgentOrganizationConverter) *Demand {
	demand.Carrier = carrier.AsOrganizationAgent()
	return demand
}

// SetConsignee устанавливает Метаданные грузополучателя (контрагент или юрлицо).
//
// Принимает [Counterparty] или [Organization].
func (demand *Demand) SetConsignee(consignee AgentOrganizationConverter) *Demand {
	demand.Consignee = consignee.AsOrganizationAgent()
	return demand
}

// SetGoodPackQuantity устанавливает количество мест.
func (demand *Demand) SetGoodPackQuantity(goodPackQuantity int) *Demand {
	demand.GoodPackQuantity = &goodPackQuantity
	return demand
}

// SetShippingInstructions устанавливает Указания грузоотправителя.
func (demand *Demand) SetShippingInstructions(shippingInstructions string) *Demand {
	demand.ShippingInstructions = &shippingInstructions
	return demand
}

// SetStateContractID устанавливает Идентификатор государственного контракта, договора (соглашения).
func (demand *Demand) SetStateContractID(stateContractID string) *Demand {
	demand.StateContractID = &stateContractID
	return demand
}

// SetTransportFacility устанавливает Транспортное средство.
func (demand *Demand) SetTransportFacility(transportFacility string) *Demand {
	demand.TransportFacility = &transportFacility
	return demand
}

// SetTransportFacilityNumber устанавливает Номер автомобиля.
func (demand *Demand) SetTransportFacilityNumber(transportFacilityNumber string) *Demand {
	demand.TransportFacilityNumber = &transportFacilityNumber
	return demand
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (demand *Demand) SetAttributes(attributes ...*Attribute) *Demand {
	demand.Attributes.Push(attributes...)
	return demand
}

// String реализует интерфейс [fmt.Stringer].
func (demand Demand) String() string {
	return Stringify(demand)
}

// MetaType возвращает код сущности.
func (Demand) MetaType() MetaType {
	return MetaTypeDemand
}

// Update shortcut
func (demand Demand) Update(ctx context.Context, client *Client, params ...*Params) (*Demand, *resty.Response, error) {
	return NewDemandService(client).Update(ctx, demand.GetID(), &demand, params...)
}

// Create shortcut
func (demand Demand) Create(ctx context.Context, client *Client, params ...*Params) (*Demand, *resty.Response, error) {
	return NewDemandService(client).Create(ctx, &demand, params...)
}

// Delete shortcut
func (demand Demand) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewDemandService(client).Delete(ctx, demand.GetID())
}

// DemandPosition Позиция Отгрузки
//
// Код сущности: demandposition
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-otgruzka-otgruzki-pozicii-otgruzki
type DemandPosition struct {
	Slot              *Slot               `json:"slot,omitempty"`               // Ячейка на складе
	Price             *float64            `json:"price,omitempty"`              // Цена товара/услуги в копейках
	Cost              *float64            `json:"cost,omitempty"`               // Себестоимость (только для услуг)
	Discount          *float64            `json:"discount,omitempty"`           // Процент скидки или наценки. Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%
	AccountID         *uuid.UUID          `json:"accountId,omitempty"`          // ID учётной записи
	Pack              *Pack               `json:"pack,omitempty"`               // Упаковка Товара
	Assortment        *AssortmentPosition `json:"assortment,omitempty"`         // Метаданные товара/услуги/серии/модификации/комплекта, которую представляет собой позиция
	Quantity          *float64            `json:"quantity,omitempty"`           // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	ID                *uuid.UUID          `json:"id,omitempty"`                 // ID позиции
	Stock             *Stock              `json:"stock,omitempty"`              // Остатки и себестоимость позиции (указывается при наличии параметра запроса `fields=stock`)
	VatEnabled        *bool               `json:"vatEnabled,omitempty"`         // Включен ли НДС для позиции. С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
	Vat               *int                `json:"vat,omitempty"`                // НДС, которым облагается текущая позиция
	Overhead          *float64            `json:"overhead,omitempty"`           // Накладные расходы
	TrackingCodes1162 Slice[TrackingCode] `json:"trackingCodes_1162,omitempty"` // Коды маркировки товаров в формате тега 1162
	TrackingCodes     Slice[TrackingCode] `json:"trackingCodes,omitempty"`      // Коды маркировки товаров и транспортных упаковок
	Things            Slice[string]       `json:"things,omitempty"`             // Серийные номера. Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете. В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута
}

// GetAccountID возвращает ID учётной записи.
func (demandPosition DemandPosition) GetAccountID() uuid.UUID {
	return Deref(demandPosition.AccountID)
}

// GetAssortment возвращает Метаданные товара/услуги/серии/модификации/комплекта, которую представляет собой позиция.
func (demandPosition DemandPosition) GetAssortment() AssortmentPosition {
	return Deref(demandPosition.Assortment)
}

// GetCost возвращает Себестоимость (только для услуг).
func (demandPosition DemandPosition) GetCost() float64 {
	return Deref(demandPosition.Cost)
}

// GetDiscount возвращает Процент скидки или наценки.
//
// Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%.
func (demandPosition DemandPosition) GetDiscount() float64 {
	return Deref(demandPosition.Discount)
}

// GetID возвращает ID позиции.
func (demandPosition DemandPosition) GetID() uuid.UUID {
	return Deref(demandPosition.ID)
}

// GetPack возвращает Упаковку Товара.
func (demandPosition DemandPosition) GetPack() Pack {
	return Deref(demandPosition.Pack)
}

// GetPrice возвращает Цену товара/услуги в копейках.
func (demandPosition DemandPosition) GetPrice() float64 {
	return Deref(demandPosition.Price)
}

// GetQuantity возвращает Количество товаров/услуг данного вида в позиции.
//
// Если позиция - товар, у которого включен учет по серийным номерам,
// то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
func (demandPosition DemandPosition) GetQuantity() float64 {
	return Deref(demandPosition.Quantity)
}

// GetSlot возвращает Ячейку на складе.
func (demandPosition DemandPosition) GetSlot() Slot {
	return Deref(demandPosition.Slot)
}

// GetThings возвращает Серийные номера.
//
// Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете.
// В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута.
func (demandPosition DemandPosition) GetThings() Slice[string] {
	return demandPosition.Things
}

// GetTrackingCodes возвращает Коды маркировки товаров и транспортных упаковок.
func (demandPosition DemandPosition) GetTrackingCodes() Slice[TrackingCode] {
	return demandPosition.TrackingCodes
}

// GetTrackingCodes1162 возвращает Коды маркировки товаров в формате тега 1162.
func (demandPosition DemandPosition) GetTrackingCodes1162() Slice[TrackingCode] {
	return demandPosition.TrackingCodes1162
}

// GetOverhead возвращает Накладные расходы.
func (demandPosition DemandPosition) GetOverhead() float64 {
	return Deref(demandPosition.Overhead)
}

// GetVat возвращает НДС, которым облагается текущая позиция.
func (demandPosition DemandPosition) GetVat() int {
	return Deref(demandPosition.Vat)
}

// GetVatEnabled возвращает true, если НДС включен для позиции.
//
// С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС".
// (vat = 0, vatEnabled = false) -> vat = "без НДС",
// (vat = 0, vatEnabled = true) -> vat = 0%.
func (demandPosition DemandPosition) GetVatEnabled() bool {
	return Deref(demandPosition.VatEnabled)
}

// GetStock возвращает Остатки и себестоимость позиции (указывается при наличии параметра запроса `fields=stock`).
func (demandPosition DemandPosition) GetStock() Stock {
	return Deref(demandPosition.Stock)
}

// SetAssortment устанавливает Метаданные товара/услуги, которую представляет собой компонент.
//
// Принимает объект, реализующий интерфейс [AssortmentConverter].
func (demandPosition *DemandPosition) SetAssortment(assortment AssortmentConverter) *DemandPosition {
	if assortment != nil {
		demandPosition.Assortment = assortment.AsAssortment()
	}
	return demandPosition
}

// SetCost устанавливает Себестоимость (только для услуг).
func (demandPosition *DemandPosition) SetCost(cost float64) *DemandPosition {
	demandPosition.Cost = &cost
	return demandPosition
}

// SetDiscount устанавливает Процент скидки или наценки.
//
// Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%.
func (demandPosition *DemandPosition) SetDiscount(discount float64) *DemandPosition {
	demandPosition.Discount = &discount
	return demandPosition
}

// SetPack устанавливает Упаковку Товара.
func (demandPosition *DemandPosition) SetPack(pack *Pack) *DemandPosition {
	if pack != nil {
		demandPosition.Pack = pack
	}
	return demandPosition
}

// SetPrice устанавливает Цену товара/услуги в копейках.
func (demandPosition *DemandPosition) SetPrice(price float64) *DemandPosition {
	demandPosition.Price = &price
	return demandPosition
}

// SetQuantity устанавливает Количество товаров данного вида в позиции.
//
// Если позиция - товар, у которого включен учет по серийным номерам,
// то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
func (demandPosition *DemandPosition) SetQuantity(quantity float64) *DemandPosition {
	demandPosition.Quantity = &quantity
	return demandPosition
}

// SetSlot устанавливает Ячейку на складе.
func (demandPosition *DemandPosition) SetSlot(slot *Slot) *DemandPosition {
	if slot != nil {
		demandPosition.Slot = slot.Clean()
	}
	return demandPosition
}

// SetThings устанавливает Серийные номера.
//
// Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете.
// В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута.
//
// Принимает множество string.
func (demandPosition *DemandPosition) SetThings(things ...string) *DemandPosition {
	demandPosition.Things = NewSliceFrom(things)
	return demandPosition
}

// SetTrackingCodes устанавливает Коды маркировки товаров и транспортных упаковок.
//
// Принимает множество объектов [TrackingCode].
func (demandPosition *DemandPosition) SetTrackingCodes(trackingCodes ...*TrackingCode) *DemandPosition {
	demandPosition.TrackingCodes.Push(trackingCodes...)
	return demandPosition
}

// SetTrackingCodes1162 устанавливает Коды маркировки товаров в формате тега 1162.
//
// Принимает множество объектов [TrackingCode].
func (demandPosition *DemandPosition) SetTrackingCodes1162(trackingCodes1162 ...*TrackingCode) *DemandPosition {
	demandPosition.TrackingCodes1162.Push(trackingCodes1162...)
	return demandPosition
}

// SetVat устанавливает НДС, которым облагается текущая позиция.
func (demandPosition *DemandPosition) SetVat(vat int) *DemandPosition {
	demandPosition.Vat = &vat
	return demandPosition
}

// SetVatEnabled устанавливает значение, учитывающее НДС для текущей позиции.
func (demandPosition *DemandPosition) SetVatEnabled(vatEnabled bool) *DemandPosition {
	demandPosition.VatEnabled = &vatEnabled
	return demandPosition
}

// AsSalesReturnPosition преобразует позицию отгрузки в позицию возврата покупателя.
//
// Копирует все поля позиции, кроме ID и AccountID.
func (demandPosition DemandPosition) AsSalesReturnPosition() *SalesReturnPosition {
	salesReturnPosition := &SalesReturnPosition{
		Assortment: demandPosition.Assortment,
		Cost:       demandPosition.Cost,
		Discount:   demandPosition.Discount,
		Pack:       demandPosition.Pack,
		Price:      demandPosition.Price,
		Quantity:   demandPosition.Quantity,
		Slot:       demandPosition.Slot,
		Vat:        demandPosition.Vat,
		VatEnabled: demandPosition.VatEnabled,
		Things:     demandPosition.Things,
	}

	return salesReturnPosition
}

// String реализует интерфейс [fmt.Stringer].
func (demandPosition DemandPosition) String() string {
	return Stringify(demandPosition)
}

// MetaType возвращает код сущности.
func (DemandPosition) MetaType() MetaType {
	return MetaTypeDemandPosition
}

// DemandService описывает методы сервиса для работы с отгрузками.
type DemandService interface {
	// GetList выполняет запрос на получение списка отгрузок.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[Demand], *resty.Response, error)

	// Create выполняет запрос на создание отгрузки.
	// Обязательные поля для заполнения:
	//	- organization (Ссылка на ваше юрлицо)
	//	- agent (Ссылка на контрагента)
	//	- store (Ссылка на склад)
	// Принимает контекст, отгрузку и опционально объект параметров запроса Params.
	// Возвращает созданную отгрузку.
	Create(ctx context.Context, demand *Demand, params ...*Params) (*Demand, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение отгрузок.
	// Изменяемые отгрузки должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список отгрузок и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых отгрузок.
	CreateUpdateMany(ctx context.Context, demandList Slice[Demand], params ...*Params) (*Slice[Demand], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление отгрузок.
	// Принимает контекст и множество отгрузок.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*Demand) (*DeleteManyResponse, *resty.Response, error)

	// Delete выполняет запрос на удаление отгрузки.
	// Принимает контекст и ID отгрузки.
	// Возвращает «true» в случае успешного удаления отгрузки.
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельной отгрузки по ID.
	// Принимает контекст, ID отгрузки взаиморасчётов и опционально объект параметров запроса Params.
	// Возвращает найденную отгрузку.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Demand, *resty.Response, error)

	// Update выполняет запрос на изменение отгрузки.
	// Принимает контекст, отгрузку и опционально объект параметров запроса Params.
	// Возвращает изменённую отгрузку.
	Update(ctx context.Context, id uuid.UUID, demand *Demand, params ...*Params) (*Demand, *resty.Response, error)

	// Template выполняет запрос на получение предзаполненной отгрузки со стандартными полями.
	// без связи с какими-либо другими документами.
	// Принимает контекст.
	// Возвращает предзаполненную отгрузку.
	Template(ctx context.Context) (*Demand, *resty.Response, error)

	// TemplateBased выполняет запрос на получение шаблона отгрузки на основе других документов.
	// Основание, на котором может быть создана:
	//	- Заказ покупателя (CustomerOrder)
	//	- Счет покупателю (InvoiceOut)
	// Принимает контекст и множество документов из списка выше.
	// Возвращает предзаполненную отгрузки на основании переданных документов.
	TemplateBased(ctx context.Context, basedOn ...MetaOwner) (*Demand, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных отгрузок.
	// Принимает контекст.
	// Возвращает объект метаданных MetaAttributesStatesSharedWrapper.
	GetMetadata(ctx context.Context) (*MetaAttributesStatesSharedWrapper, *resty.Response, error)

	// GetPositionList выполняет запрос на получение списка позиций документа.
	// Принимает контекст, ID документа и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetPositionList(ctx context.Context, id uuid.UUID, params ...*Params) (*List[DemandPosition], *resty.Response, error)

	// GetPositionByID выполняет запрос на получение отдельной позиции документа по ID.
	// Принимает контекст, ID документа, ID позиции и опционально объект параметров запроса Params.
	// Возвращает найденную позицию.
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params ...*Params) (*DemandPosition, *resty.Response, error)

	// UpdatePosition выполняет запрос на изменение позиции документа.
	// Принимает контекст, ID документа, ID позиции, позицию документа и опционально объект параметров запроса Params.
	// Возвращает изменённую позицию.
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *DemandPosition, params ...*Params) (*DemandPosition, *resty.Response, error)

	// CreatePosition выполняет запрос на добавление позиции документа.
	// Принимает контекст, ID документа, позицию документа и опционально объект параметров запроса Params.
	// Возвращает добавленную позицию.
	CreatePosition(ctx context.Context, id uuid.UUID, position *DemandPosition, params ...*Params) (*DemandPosition, *resty.Response, error)

	// CreatePositionMany выполняет запрос на массовое добавление позиций документа.
	// Принимает контекст, ID документа и множество позиций.
	// Возвращает список добавленных позиций.
	CreatePositionMany(ctx context.Context, id uuid.UUID, positions ...*DemandPosition) (*Slice[DemandPosition], *resty.Response, error)

	// DeletePosition выполняет запрос на удаление позиции документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает «true» в случае успешного удаления позиции.
	DeletePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (bool, *resty.Response, error)

	// DeletePositionMany выполняет запрос на массовое удаление позиций документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeletePositionMany(ctx context.Context, id uuid.UUID, positions ...*DemandPosition) (*DeleteManyResponse, *resty.Response, error)

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
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*Demand, *resty.Response, error)

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

	// GetEmbeddedTemplateList выполняет запрос на получение списка встроенных шаблонов отгрузок.
	// Принимает контекст.
	// Возвращает объект List.
	GetEmbeddedTemplateList(ctx context.Context) (*List[EmbeddedTemplate], *resty.Response, error)

	// GetEmbeddedTemplateByID выполняет запрос на получение отдельного встроенного шаблона отгрузки по ID.
	// Принимает контекст и ID встроенного шаблона отгрузки.
	// Возвращает найденный встроенный шаблон отгрузки.
	GetEmbeddedTemplateByID(ctx context.Context, id uuid.UUID) (*EmbeddedTemplate, *resty.Response, error)

	// GetCustomTemplateList выполняет запрос на получение списка пользовательских шаблонов отгрузки.
	// Принимает контекст.
	// Возвращает объект List.
	GetCustomTemplateList(ctx context.Context) (*List[CustomTemplate], *resty.Response, error)

	// GetCustomTemplateByID выполняет запрос на получение отдельного пользовательского шаблона отгрузки по ID.
	// Принимает контекст и ID пользовательского шаблона отгрузки.
	// Возвращает найденный пользовательский шаблон отгрузки.
	GetCustomTemplateByID(ctx context.Context, id uuid.UUID) (*CustomTemplate, *resty.Response, error)

	// PrintDocument TODO: выполняет запрос на печать отдельного документа по шаблону печатной формы.
	// Принимает контекст, и ID документа и объект PrintDocumentArg.
	// Возвращает файл
	PrintDocument(ctx context.Context, id uuid.UUID, PrintDocumentArg *PrintDocumentArg) (*PrintFile, *resty.Response, error)

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
	Evaluate(ctx context.Context, entity *Demand, evaluate ...Evaluate) (*Demand, *resty.Response, error)
}

const (
	EndpointDemand = EndpointEntity + string(MetaTypeDemand)
)

// NewDemandService принимает [Client] и возвращает сервис для работы с отгрузками.
func NewDemandService(client *Client) DemandService {
	return newMainService[Demand, DemandPosition, MetaAttributesStatesSharedWrapper, any](client, EndpointDemand)
}
