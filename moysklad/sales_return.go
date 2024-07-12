package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// SalesReturn Возврат покупателя.
//
// Код сущности: salesreturn
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vozwrat-pokupatelq
type SalesReturn struct {
	Positions           *MetaArray[SalesReturnPosition] `json:"positions,omitempty"`           // Метаданные позиций Возврата Покупателя
	VatSum              *float64                        `json:"vatSum,omitempty"`              // Сумма НДС
	AgentAccount        *AgentAccount                   `json:"agentAccount,omitempty"`        // Метаданные счета контрагента
	Applicable          *bool                           `json:"applicable,omitempty"`          // Отметка о проведении
	FactureOut          *FactureOut                     `json:"factureOut,omitempty"`          // Ссылка на Счет-фактуру выданный, с которым связан этот возврат
	Code                *string                         `json:"code,omitempty"`                // Код Возврата Покупателя
	OrganizationAccount *AgentAccount                   `json:"organizationAccount,omitempty"` // Метаданные счета юрлица
	Created             *Timestamp                      `json:"created,omitempty"`             // Дата создания
	Deleted             *Timestamp                      `json:"deleted,omitempty"`             // Момент последнего удаления Возврата Покупателя
	Description         *string                         `json:"description,omitempty"`         // Комментарий Возврата Покупателя
	ExternalCode        *string                         `json:"externalCode,omitempty"`        // Внешний код Возврата Покупателя
	Files               *MetaArray[File]                `json:"files,omitempty"`               // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group               *Group                          `json:"group,omitempty"`               // Отдел сотрудника
	ID                  *uuid.UUID                      `json:"id,omitempty"`                  // ID Возврата Покупателя
	Meta                *Meta                           `json:"meta,omitempty"`                // Метаданные Возврата Покупателя
	Moment              *Timestamp                      `json:"moment,omitempty"`              // Дата документа
	Name                *string                         `json:"name,omitempty"`                // Наименование Возврата Покупателя
	AccountID           *uuid.UUID                      `json:"accountId,omitempty"`           // ID учетной записи
	Contract            *NullValue[Contract]            `json:"contract,omitempty"`            // Метаданные договора
	Agent               *Agent                          `json:"agent,omitempty"`               // Метаданные контрагента
	Organization        *Organization                   `json:"organization,omitempty"`        // Метаданные юрлица
	Printed             *bool                           `json:"printed,omitempty"`             // Напечатан ли документ
	Project             *NullValue[Project]             `json:"project,omitempty"`             // Метаданные проекта
	Published           *bool                           `json:"published,omitempty"`           // Опубликован ли документ
	Rate                *NullValue[Rate]                `json:"rate,omitempty"`                // Валюта
	SalesChannel        *NullValue[SalesChannel]        `json:"salesChannel,omitempty"`        // Метаданные канала продаж
	Shared              *bool                           `json:"shared,omitempty"`              // Общий доступ
	State               *NullValue[State]               `json:"state,omitempty"`               // Метаданные статуса Возврата Покупателя
	Store               *Store                          `json:"store,omitempty"`               // Метаданные склада
	Sum                 *float64                        `json:"sum,omitempty"`                 // Сумма Возврата Покупателя в копейках
	SyncID              *uuid.UUID                      `json:"syncId,omitempty"`              // ID синхронизации
	Updated             *Timestamp                      `json:"updated,omitempty"`             // Момент последнего обновления Возврата Покупателя
	VatEnabled          *bool                           `json:"vatEnabled,omitempty"`          // Учитывается ли НДС
	VatIncluded         *bool                           `json:"vatIncluded,omitempty"`         // Включен ли НДС в цену
	Owner               *Employee                       `json:"owner,omitempty"`               // Владелец (Сотрудник)
	Demand              *Demand                         `json:"demand,omitempty"`              // Ссылка на отгрузку, по которой произошел возврат
	Losses              Slice[Loss]                     `json:"losses,omitempty"`              // Массив ссылок на связанные списания
	Payments            Slice[Payment]                  `json:"payments,omitempty"`            // Массив ссылок на связанные платежи
	PayedSum            *float64                        `json:"payedSum,omitempty"`            // Сумма исходящих платежей по возврату покупателя
	Attributes          Slice[Attribute]                `json:"attributes,omitempty"`          // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (salesReturn SalesReturn) Clean() *SalesReturn {
	return &SalesReturn{Meta: salesReturn.Meta}
}

// AsOperation реализует интерфейс [OperationConverter].
func (salesReturn SalesReturn) AsOperation() *Operation {
	return newOperation(salesReturn)
}

// AsTaskOperation реализует интерфейс [TaskOperationConverter].
func (salesReturn SalesReturn) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: salesReturn.Meta}
}

// GetPositions возвращает Метаданные позиций Возврата Покупателя.
func (salesReturn SalesReturn) GetPositions() MetaArray[SalesReturnPosition] {
	return Deref(salesReturn.Positions)
}

// GetVatSum возвращает Сумму НДС.
func (salesReturn SalesReturn) GetVatSum() float64 {
	return Deref(salesReturn.VatSum)
}

// GetAgentAccount возвращает Метаданные счета контрагента.
func (salesReturn SalesReturn) GetAgentAccount() AgentAccount {
	return Deref(salesReturn.AgentAccount)
}

// GetApplicable возвращает Отметку о проведении.
func (salesReturn SalesReturn) GetApplicable() bool {
	return Deref(salesReturn.Applicable)
}

// GetFactureOut возвращает Ссылку на выданный счет-фактуру, с которым связан этот возврат.
func (salesReturn SalesReturn) GetFactureOut() FactureOut {
	return Deref(salesReturn.FactureOut)
}

// GetCode возвращает Код Возврата Покупателя.
func (salesReturn SalesReturn) GetCode() string {
	return Deref(salesReturn.Code)
}

// GetOrganizationAccount возвращает Метаданные счета юрлица.
func (salesReturn SalesReturn) GetOrganizationAccount() AgentAccount {
	return Deref(salesReturn.OrganizationAccount)
}

// GetCreated возвращает Дату создания.
func (salesReturn SalesReturn) GetCreated() Timestamp {
	return Deref(salesReturn.Created)
}

// GetDeleted возвращает Момент последнего удаления Возврата Покупателя.
func (salesReturn SalesReturn) GetDeleted() Timestamp {
	return Deref(salesReturn.Deleted)
}

// GetDescription возвращает Комментарий Возврата Покупателя.
func (salesReturn SalesReturn) GetDescription() string {
	return Deref(salesReturn.Description)
}

// GetExternalCode возвращает Внешний код Возврата Покупателя.
func (salesReturn SalesReturn) GetExternalCode() string {
	return Deref(salesReturn.ExternalCode)
}

// GetFiles возвращает Метаданные массива Файлов.
func (salesReturn SalesReturn) GetFiles() MetaArray[File] {
	return Deref(salesReturn.Files)
}

// GetGroup возвращает Отдел сотрудника.
func (salesReturn SalesReturn) GetGroup() Group {
	return Deref(salesReturn.Group)
}

// GetID возвращает ID Возврата Покупателя.
func (salesReturn SalesReturn) GetID() uuid.UUID {
	return Deref(salesReturn.ID)
}

// GetMeta возвращает Метаданные Возврата Покупателя.
func (salesReturn SalesReturn) GetMeta() Meta {
	return Deref(salesReturn.Meta)
}

// GetMoment возвращает Дату документа.
func (salesReturn SalesReturn) GetMoment() Timestamp {
	return Deref(salesReturn.Moment)
}

// GetName возвращает Наименование Возврата Покупателя.
func (salesReturn SalesReturn) GetName() string {
	return Deref(salesReturn.Name)
}

// GetAccountID возвращает ID учётной записи.
func (salesReturn SalesReturn) GetAccountID() uuid.UUID {
	return Deref(salesReturn.AccountID)
}

// GetContract возвращает Метаданные договора.
func (salesReturn SalesReturn) GetContract() Contract {
	return salesReturn.Contract.getValue()
}

// GetAgent возвращает Метаданные контрагента.
func (salesReturn SalesReturn) GetAgent() Agent {
	return Deref(salesReturn.Agent)
}

// GetOrganization возвращает Метаданные юрлица.
func (salesReturn SalesReturn) GetOrganization() Organization {
	return Deref(salesReturn.Organization)
}

// GetPrinted возвращает true, если документ напечатан.
func (salesReturn SalesReturn) GetPrinted() bool {
	return Deref(salesReturn.Printed)
}

// GetProject возвращает Метаданные проекта.
func (salesReturn SalesReturn) GetProject() Project {
	return salesReturn.Project.getValue()
}

// GetPublished возвращает true, если документ опубликован.
func (salesReturn SalesReturn) GetPublished() bool {
	return Deref(salesReturn.Published)
}

// GetRate возвращает Валюту.
func (salesReturn SalesReturn) GetRate() Rate {
	return salesReturn.Rate.getValue()
}

// GetSalesChannel возвращает Метаданные канала продаж.
func (salesReturn SalesReturn) GetSalesChannel() SalesChannel {
	return salesReturn.SalesChannel.getValue()
}

// GetShared возвращает флаг Общего доступа.
func (salesReturn SalesReturn) GetShared() bool {
	return Deref(salesReturn.Shared)
}

// GetState возвращает Метаданные статуса Возврата Покупателя.
func (salesReturn SalesReturn) GetState() State {
	return salesReturn.State.getValue()
}

// GetStore возвращает Метаданные склада.
func (salesReturn SalesReturn) GetStore() Store {
	return Deref(salesReturn.Store)
}

// GetSum возвращает Сумму Возврата Покупателя в копейках.
func (salesReturn SalesReturn) GetSum() float64 {
	return Deref(salesReturn.Sum)
}

// GetSyncID возвращает ID синхронизации.
func (salesReturn SalesReturn) GetSyncID() uuid.UUID {
	return Deref(salesReturn.SyncID)
}

// GetUpdated возвращает Момент последнего обновления Входящего платежа.
func (salesReturn SalesReturn) GetUpdated() Timestamp {
	return Deref(salesReturn.Updated)
}

// GetVatEnabled возвращает true, если учитывается НДС.
func (salesReturn SalesReturn) GetVatEnabled() bool {
	return Deref(salesReturn.VatEnabled)
}

// GetVatIncluded возвращает true, если НДС включен в цену.
func (salesReturn SalesReturn) GetVatIncluded() bool {
	return Deref(salesReturn.VatIncluded)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (salesReturn SalesReturn) GetOwner() Employee {
	return Deref(salesReturn.Owner)
}

// GetDemand возвращает ссылку на связанную отгрузку.
func (salesReturn SalesReturn) GetDemand() Demand {
	return Deref(salesReturn.Demand)
}

// GetLosses возвращает Массив ссылок на связанные списания.
func (salesReturn SalesReturn) GetLosses() Slice[Loss] {
	return salesReturn.Losses
}

// GetPayments возвращает Массив ссылок на связанные платежи.
func (salesReturn SalesReturn) GetPayments() Slice[Payment] {
	return salesReturn.Payments
}

// GetPayedSum возвращает Сумму исходящих платежей по возврату покупателя.
func (salesReturn SalesReturn) GetPayedSum() float64 {
	return Deref(salesReturn.PayedSum)
}

// GetAttributes возвращает Список метаданных доп. полей.
func (salesReturn SalesReturn) GetAttributes() Slice[Attribute] {
	return salesReturn.Attributes
}

// SetPositions устанавливает Метаданные позиций Возврата Покупателя.
//
// Принимает множество объектов [SalesReturnPosition].
func (salesReturn *SalesReturn) SetPositions(positions ...*SalesReturnPosition) *SalesReturn {
	salesReturn.Positions = NewMetaArrayFrom(positions)
	return salesReturn
}

// SetAgentAccount устанавливает Метаданные счета контрагента.
func (salesReturn *SalesReturn) SetAgentAccount(agentAccount *AgentAccount) *SalesReturn {
	if agentAccount != nil {
		salesReturn.AgentAccount = agentAccount.Clean()
	}
	return salesReturn
}

// SetApplicable устанавливает Отметку о проведении.
func (salesReturn *SalesReturn) SetApplicable(applicable bool) *SalesReturn {
	salesReturn.Applicable = &applicable
	return salesReturn
}

// SetFactureOut устанавливает Ссылку на Счет-фактуру выданный, с которым связан этот возврат.
func (salesReturn *SalesReturn) SetFactureOut(factureOut *FactureOut) *SalesReturn {
	if factureOut != nil {
		salesReturn.FactureOut = factureOut.Clean()
	}
	return salesReturn
}

// SetCode устанавливает Код Возврата Покупателя.
func (salesReturn *SalesReturn) SetCode(code string) *SalesReturn {
	salesReturn.Code = &code
	return salesReturn
}

// SetOrganizationAccount устанавливает Метаданные счета юрлица.
func (salesReturn *SalesReturn) SetOrganizationAccount(organizationAccount *AgentAccount) *SalesReturn {
	if organizationAccount != nil {
		salesReturn.OrganizationAccount = organizationAccount.Clean()
	}
	return salesReturn
}

// SetDescription устанавливает Комментарий Возврата Покупателя.
func (salesReturn *SalesReturn) SetDescription(description string) *SalesReturn {
	salesReturn.Description = &description
	return salesReturn
}

// SetExternalCode устанавливает Внешний код Возврата Покупателя.
func (salesReturn *SalesReturn) SetExternalCode(externalCode string) *SalesReturn {
	salesReturn.ExternalCode = &externalCode
	return salesReturn
}

// SetFiles устанавливает Метаданные массива Файлов.
//
// Принимает множество объектов [File].
func (salesReturn *SalesReturn) SetFiles(files ...*File) *SalesReturn {
	salesReturn.Files = NewMetaArrayFrom(files)
	return salesReturn
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (salesReturn *SalesReturn) SetGroup(group *Group) *SalesReturn {
	if group != nil {
		salesReturn.Group = group.Clean()
	}
	return salesReturn
}

// SetMeta устанавливает Метаданные Возврата Покупателя.
func (salesReturn *SalesReturn) SetMeta(meta *Meta) *SalesReturn {
	salesReturn.Meta = meta
	return salesReturn
}

// SetMoment устанавливает Дату документа.
func (salesReturn *SalesReturn) SetMoment(moment *Timestamp) *SalesReturn {
	salesReturn.Moment = moment
	return salesReturn
}

// SetName устанавливает Наименование Возврата Покупателя.
func (salesReturn *SalesReturn) SetName(name string) *SalesReturn {
	salesReturn.Name = &name
	return salesReturn
}

// SetContract устанавливает Метаданные договора.
//
// Передача nil передаёт сброс значения (null).
func (salesReturn *SalesReturn) SetContract(contract *Contract) *SalesReturn {
	salesReturn.Contract = NewNullValue(contract)
	return salesReturn
}

// SetAgent устанавливает Метаданные Контрагента, связанного с бонусной операцией.
//
// Принимает [Counterparty] или [Organization].
func (salesReturn *SalesReturn) SetAgent(agent AgentOrganizationConverter) *SalesReturn {
	if agent != nil {
		salesReturn.Agent = agent.AsOrganizationAgent()
	}
	return salesReturn
}

// SetOrganization устанавливает Метаданные юрлица.
func (salesReturn *SalesReturn) SetOrganization(organization *Organization) *SalesReturn {
	if organization != nil {
		salesReturn.Organization = organization.Clean()
	}
	return salesReturn
}

// SetProject устанавливает Метаданные проекта.
//
// Передача nil передаёт сброс значения (null).
func (salesReturn *SalesReturn) SetProject(project *Project) *SalesReturn {
	salesReturn.Project = NewNullValue(project)
	return salesReturn
}

// SetRate устанавливает Валюту.
//
// Передача nil передаёт сброс значения (null).
func (salesReturn *SalesReturn) SetRate(rate *Rate) *SalesReturn {
	salesReturn.Rate = NewNullValue(rate)
	return salesReturn
}

// SetSalesChannel устанавливает Метаданные канала продаж.
//
// Передача nil передаёт сброс значения (null).
func (salesReturn *SalesReturn) SetSalesChannel(salesChannel *SalesChannel) *SalesReturn {
	salesReturn.SalesChannel = NewNullValue(salesChannel)
	return salesReturn
}

// SetShared устанавливает флаг общего доступа.
func (salesReturn *SalesReturn) SetShared(shared bool) *SalesReturn {
	salesReturn.Shared = &shared
	return salesReturn
}

// SetState устанавливает Метаданные статуса Возврата Покупателя.
//
// Передача nil передаёт сброс значения (null).
func (salesReturn *SalesReturn) SetState(state *State) *SalesReturn {
	salesReturn.State = NewNullValue(state)
	return salesReturn
}

// SetStore устанавливает Метаданные склада.
func (salesReturn *SalesReturn) SetStore(store *Store) *SalesReturn {
	if store != nil {
		salesReturn.Store = store.Clean()
	}
	return salesReturn
}

// SetSyncID устанавливает ID синхронизации.
func (salesReturn *SalesReturn) SetSyncID(syncID uuid.UUID) *SalesReturn {
	salesReturn.SyncID = &syncID
	return salesReturn
}

// SetVatEnabled устанавливает значение, учитывающее НДС возврата покупателя.
func (salesReturn *SalesReturn) SetVatEnabled(vatEnabled bool) *SalesReturn {
	salesReturn.VatEnabled = &vatEnabled
	return salesReturn
}

// SetVatIncluded устанавливает флаг включения НДС в цену.
func (salesReturn *SalesReturn) SetVatIncluded(vatIncluded bool) *SalesReturn {
	salesReturn.VatIncluded = &vatIncluded
	return salesReturn
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (salesReturn *SalesReturn) SetOwner(owner *Employee) *SalesReturn {
	if owner != nil {
		salesReturn.Owner = owner.Clean()
	}
	return salesReturn
}

// SetDemand устанавливает Ссылку на отгрузку, по которой произошел возврат.
func (salesReturn *SalesReturn) SetDemand(demand *Demand) *SalesReturn {
	if demand != nil {
		salesReturn.Demand = demand.Clean()
	}
	return salesReturn
}

// SetLosses устанавливает Массив ссылок на связанные списания.
//
// Принимает множество объектов [Loss].
func (salesReturn *SalesReturn) SetLosses(losses ...*Loss) *SalesReturn {
	salesReturn.Losses = losses
	return salesReturn
}

// SetPayments устанавливает Метаданные ссылок на связанные платежи.
//
// Принимает множество объектов, реализующих интерфейс [PaymentConverter].
func (salesReturn *SalesReturn) SetPayments(payments ...PaymentConverter) *SalesReturn {
	salesReturn.Payments = NewPaymentsFrom(payments)
	return salesReturn
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (salesReturn *SalesReturn) SetAttributes(attributes ...*Attribute) *SalesReturn {
	salesReturn.Attributes = attributes
	return salesReturn
}

// String реализует интерфейс [fmt.Stringer].
func (salesReturn SalesReturn) String() string {
	return Stringify(salesReturn)
}

// MetaType возвращает тип сущности.
func (SalesReturn) MetaType() MetaType {
	return MetaTypeSalesReturn
}

// Update shortcut
func (salesReturn SalesReturn) Update(ctx context.Context, client *Client, params ...*Params) (*SalesReturn, *resty.Response, error) {
	return NewSalesReturnService(client).Update(ctx, salesReturn.GetID(), &salesReturn, params...)
}

// Create shortcut
func (salesReturn SalesReturn) Create(ctx context.Context, client *Client, params ...*Params) (*SalesReturn, *resty.Response, error) {
	return NewSalesReturnService(client).Create(ctx, &salesReturn, params...)
}

// Delete shortcut
func (salesReturn SalesReturn) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewSalesReturnService(client).Delete(ctx, salesReturn.GetID())
}

// SalesReturnPosition Позиция Возврата покупателя.
//
// Код сущности: salesreturnposition
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vozwrat-pokupatelq-vozwraty-pokupatelej-pozicii-vozwrata-pokupatelq
type SalesReturnPosition struct {
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Price      *float64            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Cost       *float64            `json:"cost,omitempty"`       // Себестоимость (выводится, если документ был создан без основания)
	Country    *Country            `json:"country,omitempty"`    // Метаданные Страны
	Discount   *float64            `json:"discount,omitempty"`   // Процент скидки или наценки. Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%
	GTD        *GTD                `json:"gtd,omitempty"`        // ГТД
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Slot       *Slot               `json:"slot,omitempty"`       // Ячейка на складе
	Stock      *Stock              `json:"stock,omitempty"`      // Остатки и себестоимость позиции (указывается при наличии параметра запроса `fields=stock`)
	Vat        *int                `json:"vat,omitempty"`        // НДС, которым облагается текущая позиция
	VatEnabled *bool               `json:"vatEnabled,omitempty"` // Включен ли НДС для позиции. С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
	Things     Slice[string]       `json:"things,omitempty"`     // Серийные номера. Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете. В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута.
}

// GetAccountID возвращает ID учётной записи.
func (salesReturnPosition SalesReturnPosition) GetAccountID() uuid.UUID {
	return Deref(salesReturnPosition.AccountID)
}

// GetAssortment возвращает Метаданные товара/услуги/серии/модификации, которую представляет собой позиция.
func (salesReturnPosition SalesReturnPosition) GetAssortment() AssortmentPosition {
	return Deref(salesReturnPosition.Assortment)
}

// GetCost возвращает Себестоимость (выводится, если документ был создан без основания).
func (salesReturnPosition SalesReturnPosition) GetCost() float64 {
	return Deref(salesReturnPosition.Cost)
}

// GetCountry возвращает Метаданные Страны.
func (salesReturnPosition SalesReturnPosition) GetCountry() Country {
	return Deref(salesReturnPosition.Country)
}

// GetDiscount возвращает Процент скидки или наценки.
//
// Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%.
func (salesReturnPosition SalesReturnPosition) GetDiscount() float64 {
	return Deref(salesReturnPosition.Discount)
}

// GetGTD возвращает ГТД.
func (salesReturnPosition SalesReturnPosition) GetGTD() GTD {
	return Deref(salesReturnPosition.GTD)
}

// GetID возвращает ID позиции.
func (salesReturnPosition SalesReturnPosition) GetID() uuid.UUID {
	return Deref(salesReturnPosition.ID)
}

// GetPack возвращает Упаковку Товара.
func (salesReturnPosition SalesReturnPosition) GetPack() Pack {
	return Deref(salesReturnPosition.Pack)
}

// GetPrice возвращает Цену товара/услуги в копейках.
func (salesReturnPosition SalesReturnPosition) GetPrice() float64 {
	return Deref(salesReturnPosition.Price)
}

// GetQuantity возвращает Количество товаров/услуг данного вида в позиции.
//
// Если позиция - товар, у которого включен учет по серийным номерам,
// то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
func (salesReturnPosition SalesReturnPosition) GetQuantity() float64 {
	return Deref(salesReturnPosition.Quantity)
}

// GetSlot возвращает Ячейку на складе.
func (salesReturnPosition SalesReturnPosition) GetSlot() Slot {
	return Deref(salesReturnPosition.Slot)
}

// GetThings возвращает Серийные номера.
//
// Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете.
// В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута.
func (salesReturnPosition SalesReturnPosition) GetThings() Slice[string] {
	return salesReturnPosition.Things
}

// GetVat возвращает НДС, которым облагается текущая позиция.
func (salesReturnPosition SalesReturnPosition) GetVat() int {
	return Deref(salesReturnPosition.Vat)
}

// GetVatEnabled возвращает true, если учитывается НДС.
//
// С помощью этого флага для товара можно выставлять НДС = 0 или НДС = "без НДС".
//
// (vat = 0, vatEnabled = false) -> vat = "без НДС",
//
// (vat = 0, vatEnabled = true) -> vat = 0%.
func (salesReturnPosition SalesReturnPosition) GetVatEnabled() bool {
	return Deref(salesReturnPosition.VatEnabled)
}

// GetStock возвращает Остатки и себестоимость позиции (указывается при наличии параметра запроса `fields=stock`).
func (salesReturnPosition SalesReturnPosition) GetStock() Stock {
	return Deref(salesReturnPosition.Stock)
}

// SetAssortment устанавливает Метаданные товара/услуги/серии/модификации, которую представляет собой позиция.
//
// Принимает объект, реализующий интерфейс [AssortmentConverter].
func (salesReturnPosition *SalesReturnPosition) SetAssortment(assortment AssortmentConverter) *SalesReturnPosition {
	if assortment != nil {
		salesReturnPosition.Assortment = assortment.AsAssortment()
	}
	return salesReturnPosition
}

// SetCost устанавливает Себестоимость (выводится, если документ был создан без основания).
func (salesReturnPosition *SalesReturnPosition) SetCost(cost float64) *SalesReturnPosition {
	salesReturnPosition.Cost = &cost
	return salesReturnPosition
}

// SetCountry устанавливает Метаданные страны.
func (salesReturnPosition *SalesReturnPosition) SetCountry(country *Country) *SalesReturnPosition {
	if country != nil {
		salesReturnPosition.Country = country.Clean()
	}
	return salesReturnPosition
}

// SetDiscount устанавливает Процент скидки или наценки.
//
// Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%.
func (salesReturnPosition *SalesReturnPosition) SetDiscount(discount float64) *SalesReturnPosition {
	salesReturnPosition.Discount = &discount
	return salesReturnPosition
}

// SetGTD устанавливает ГТД.
func (salesReturnPosition *SalesReturnPosition) SetGTD(gtd *GTD) *SalesReturnPosition {
	salesReturnPosition.GTD = gtd
	return salesReturnPosition
}

// SetPack устанавливает Упаковку Товара.
func (salesReturnPosition *SalesReturnPosition) SetPack(pack *Pack) *SalesReturnPosition {
	salesReturnPosition.Pack = pack
	return salesReturnPosition
}

// SetPrice устанавливает Цену товара/услуги в копейках.
func (salesReturnPosition *SalesReturnPosition) SetPrice(price float64) *SalesReturnPosition {
	salesReturnPosition.Price = &price
	return salesReturnPosition
}

// SetQuantity устанавливает Количество товаров данного вида в позиции.
//
// Если позиция - товар, у которого включен учет по серийным номерам,
// то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
func (salesReturnPosition *SalesReturnPosition) SetQuantity(quantity float64) *SalesReturnPosition {
	salesReturnPosition.Quantity = &quantity
	return salesReturnPosition
}

// SetSlot устанавливает Ячейку на складе.
func (salesReturnPosition *SalesReturnPosition) SetSlot(slot *Slot) *SalesReturnPosition {
	if slot != nil {
		salesReturnPosition.Slot = slot.Clean()
	}
	return salesReturnPosition
}

// SetThings устанавливает Серийные номера.
//
// Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете.
// В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута.
//
// Принимает множество string.
func (salesReturnPosition *SalesReturnPosition) SetThings(things ...string) *SalesReturnPosition {
	salesReturnPosition.Things = NewSliceFrom(things)
	return salesReturnPosition
}

// SetVat устанавливает НДС, которым облагается текущая позиция.
func (salesReturnPosition *SalesReturnPosition) SetVat(vat int) *SalesReturnPosition {
	salesReturnPosition.Vat = &vat
	return salesReturnPosition
}

// SetVatEnabled устанавливает значение, учитывающее НДС для Группы товаров.
//
// С помощью этого флага для группы можно выставлять НДС = 0 или НДС = "без НДС".
//
// (vat = 0, vatEnabled = false) -> vat = "без НДС"
//
// (vat = 0, vatEnabled = true) -> vat = 0%.
func (salesReturnPosition *SalesReturnPosition) SetVatEnabled(vatEnabled bool) *SalesReturnPosition {
	salesReturnPosition.VatEnabled = &vatEnabled
	return salesReturnPosition
}

// String реализует интерфейс [fmt.Stringer].
func (salesReturnPosition SalesReturnPosition) String() string {
	return Stringify(salesReturnPosition)
}

// MetaType возвращает тип сущности.
func (SalesReturnPosition) MetaType() MetaType {
	return MetaTypeSalesReturnPosition
}

// SalesReturnService описывает методы сервиса для работы с возвратами покупателей.
type SalesReturnService interface {
	// GetList выполняет запрос на получение списка возвратов покупателя.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[SalesReturn], *resty.Response, error)

	// Create выполняет запрос на создание возврата покупателя.
	// Обязательные поля для заполнения:
	//	- organization (Ссылка на ваше юрлицо)
	//	- store (Ссылка на склад)
	//	- demand (Ссылка на отгрузку)
	//	- agent (Ссылка на контрагента (поставщика))
	// Принимает контекст, возврат покупателя и опционально объект параметров запроса Params.
	// Возвращает созданный возврат покупателя.
	Create(ctx context.Context, salesReturn *SalesReturn, params ...*Params) (*SalesReturn, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение возвратов покупателя.
	// Изменяемые возвраты покупателя должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список возвратов покупателя и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых возвратов покупателя.
	CreateUpdateMany(ctx context.Context, salesReturnList Slice[SalesReturn], params ...*Params) (*Slice[SalesReturn], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление возвратов покупателя.
	// Принимает контекст и множество возвратов покупателя.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*SalesReturn) (*DeleteManyResponse, *resty.Response, error)

	// Delete выполняет запрос на удаление возврата покупателя.
	// Принимает контекст и ID возврата покупателя.
	// Возвращает «true» в случае успешного удаления возврата покупателя.
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельного возврата покупателя по ID.
	// Принимает контекст, ID возврата покупателя и опционально объект параметров запроса Params.
	// Возвращает возврат покупателя.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*SalesReturn, *resty.Response, error)

	// Update выполняет запрос на изменение возврата покупателя.
	// Принимает контекст, возврат покупателя и опционально объект параметров запроса Params.
	// Возвращает изменённый возврат покупателя.
	Update(ctx context.Context, id uuid.UUID, salesReturn *SalesReturn, params ...*Params) (*SalesReturn, *resty.Response, error)

	// Template выполняет запрос на получение предзаполненного возврата покупателя со стандартными полями без связи с какими-либо другими документами.
	// Принимает контекст.
	// Возвращает предзаполненный возврат покупателя.
	Template(ctx context.Context) (*SalesReturn, *resty.Response, error)

	// TemplateBased выполняет запрос на получение шаблона возврата покупателя на основе других документов.
	// Основание, на котором может быть создан:
	//	- Отгрузка (Demand)
	// Принимает контекст и один документ из списка выше.
	// Возвращает предзаполненный возврат покупателя на основании переданных документов.
	TemplateBased(ctx context.Context, basedOn ...MetaOwner) (*SalesReturn, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных возвратов покупателя.
	// Принимает контекст.
	// Возвращает объект метаданных MetaAttributesStatesSharedWrapper.
	GetMetadata(ctx context.Context) (*MetaAttributesStatesSharedWrapper, *resty.Response, error)

	// GetPositionList выполняет запрос на получение списка позиций документа.
	// Принимает контекст, ID документа и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetPositionList(ctx context.Context, id uuid.UUID, params ...*Params) (*List[SalesReturnPosition], *resty.Response, error)

	// GetPositionByID выполняет запрос на получение отдельной позиции документа по ID.
	// Принимает контекст, ID документа, ID позиции и опционально объект параметров запроса Params.
	// Возвращает найденную позицию.
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params ...*Params) (*SalesReturnPosition, *resty.Response, error)

	// UpdatePosition выполняет запрос на изменение позиции документа.
	// Принимает контекст, ID документа, ID позиции, позицию документа и опционально объект параметров запроса Params.
	// Возвращает изменённую позицию.
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *SalesReturnPosition, params ...*Params) (*SalesReturnPosition, *resty.Response, error)

	// CreatePosition выполняет запрос на добавление позиции документа.
	// Принимает контекст, ID документа, позицию документа и опционально объект параметров запроса Params.
	// Возвращает добавленную позицию.
	CreatePosition(ctx context.Context, id uuid.UUID, position *SalesReturnPosition, params ...*Params) (*SalesReturnPosition, *resty.Response, error)

	// CreatePositionMany выполняет запрос на массовое добавление позиций документа.
	// Принимает контекст, ID документа и множество позиций.
	// Возвращает список добавленных позиций.
	CreatePositionMany(ctx context.Context, id uuid.UUID, positions ...*SalesReturnPosition) (*Slice[SalesReturnPosition], *resty.Response, error)

	// DeletePosition выполняет запрос на удаление позиции документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает «true» в случае успешного удаления позиции.
	DeletePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (bool, *resty.Response, error)

	// DeletePositionMany выполняет запрос на массовое удаление позиций документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeletePositionMany(ctx context.Context, id uuid.UUID, entities ...*SalesReturnPosition) (*DeleteManyResponse, *resty.Response, error)

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
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*SalesReturn, *resty.Response, error)

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
	Evaluate(ctx context.Context, entity *SalesReturn, evaluate ...Evaluate) (*SalesReturn, *resty.Response, error)
}

const (
	EndpointSalesReturn = EndpointEntity + string(MetaTypeSalesReturn)
)

// NewSalesReturnService принимает [Client] и возвращает сервис для работы с возвратами покупателей.
func NewSalesReturnService(client *Client) SalesReturnService {
	return newMainService[SalesReturn, SalesReturnPosition, MetaAttributesStatesSharedWrapper, any](client, EndpointSalesReturn)
}
