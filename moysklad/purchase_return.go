package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"time"
)

// PurchaseReturn Возврат поставщику.
//
// Код сущности: purchasereturn
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vozwrat-postawschiku
type PurchaseReturn struct {
	Printed             *bool                              `json:"printed,omitempty"`             // Напечатан ли документ
	Supply              *Supply                            `json:"supply,omitempty"`              // Ссылка на приемку, по которой произошел возврат
	AgentAccount        *AgentAccount                      `json:"agentAccount,omitempty"`        // Метаданные счета контрагента
	Applicable          *bool                              `json:"applicable,omitempty"`          // Отметка о проведении
	Payments            Slice[Payment]                     `json:"payments,omitempty"`            // Массив ссылок на связанные платежи
	Code                *string                            `json:"code,omitempty"`                // Код Возврата поставщику
	OrganizationAccount *AgentAccount                      `json:"organizationAccount,omitempty"` // Метаданные счета юрлица
	Created             *Timestamp                         `json:"created,omitempty"`             // Дата создания
	Deleted             *Timestamp                         `json:"deleted,omitempty"`             // Момент последнего удаления Возврата поставщику
	Description         *string                            `json:"description,omitempty"`         // Комментарий Возврата поставщику
	ExternalCode        *string                            `json:"externalCode,omitempty"`        // Внешний код Возврата поставщику
	Files               *MetaArray[File]                   `json:"files,omitempty"`               // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group               *Group                             `json:"group,omitempty"`               // Отдел сотрудника
	ID                  *uuid.UUID                         `json:"id,omitempty"`                  // ID Возврата поставщику
	Meta                *Meta                              `json:"meta,omitempty"`                // Метаданные Возврата поставщику
	Moment              *Timestamp                         `json:"moment,omitempty"`              // Дата документа
	Name                *string                            `json:"name,omitempty"`                // Наименование Возврата поставщику
	AccountID           *uuid.UUID                         `json:"accountId,omitempty"`           // ID учётной записи
	Contract            *NullValue[Contract]               `json:"contract,omitempty"`            // Метаданные договора
	Agent               *Agent                             `json:"agent,omitempty"`               // Метаданные контрагента
	Organization        *Organization                      `json:"organization,omitempty"`        // Метаданные юрлица
	Project             *NullValue[Project]                `json:"project,omitempty"`             // Метаданные проекта
	Published           *bool                              `json:"published,omitempty"`           // Опубликован ли документ
	Rate                *NullValue[Rate]                   `json:"rate,omitempty"`                // Валюта
	Shared              *bool                              `json:"shared,omitempty"`              // Общий доступ
	State               *NullValue[State]                  `json:"state,omitempty"`               // Метаданные статуса Возврата поставщику
	Store               *Store                             `json:"store,omitempty"`               // Метаданные склада
	Sum                 *float64                           `json:"sum,omitempty"`                 // Сумма Возврата поставщику в копейках
	SyncID              *uuid.UUID                         `json:"syncId,omitempty"`              // ID синхронизации
	Updated             *Timestamp                         `json:"updated,omitempty"`             // Момент последнего обновления Возврата поставщику
	VatEnabled          *bool                              `json:"vatEnabled,omitempty"`          // Учитывается ли НДС
	VatIncluded         *bool                              `json:"vatIncluded,omitempty"`         // Включен ли НДС в цену
	VatSum              *float64                           `json:"vatSum,omitempty"`              // Сумма НДС
	Positions           *MetaArray[PurchaseReturnPosition] `json:"positions,omitempty"`           // Ссылка на позиции Возврата поставщику
	Owner               *Employee                          `json:"owner,omitempty"`               // Метаданные владельца (Сотрудника)
	FactureIn           *FactureIn                         `json:"factureIn,omitempty"`           // Ссылка на Счет-фактуру полученный
	FactureOut          *FactureOut                        `json:"factureOut,omitempty"`          // Ссылка на Счет-фактуру выданный
	PayedSum            *float64                           `json:"payedSum,omitempty"`            // Сумма входящих платежей по возврату поставщику
	Attributes          Slice[Attribute]                   `json:"attributes,omitempty"`          // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (purchaseReturn PurchaseReturn) Clean() *PurchaseReturn {
	if purchaseReturn.Meta == nil {
		return nil
	}
	return &PurchaseReturn{Meta: purchaseReturn.Meta}
}

// AsOperation возвращает объект [Operation] c полем [Meta].
func (purchaseReturn PurchaseReturn) AsOperation() *Operation {
	return &Operation{Meta: purchaseReturn.GetMeta(), LinkedSum: purchaseReturn.GetSum()}
}

// AsTaskOperation реализует интерфейс [TaskOperationConverter].
func (purchaseReturn PurchaseReturn) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: purchaseReturn.Meta}
}

// AsOperationIn реализует интерфейс [OperationInConverter].
func (purchaseReturn PurchaseReturn) AsOperationIn() *Operation {
	return purchaseReturn.AsOperation()
}

// GetPrinted возвращает true, если документ напечатан.
func (purchaseReturn PurchaseReturn) GetPrinted() bool {
	return Deref(purchaseReturn.Printed)
}

// GetSupply возвращает Ссылку на приемку, по которой произошел возврат.
func (purchaseReturn PurchaseReturn) GetSupply() Supply {
	return Deref(purchaseReturn.Supply)
}

// GetAgentAccount возвращает Метаданные счета контрагента.
func (purchaseReturn PurchaseReturn) GetAgentAccount() AgentAccount {
	return Deref(purchaseReturn.AgentAccount)
}

// GetApplicable возвращает Отметку о проведении.
func (purchaseReturn PurchaseReturn) GetApplicable() bool {
	return Deref(purchaseReturn.Applicable)
}

// GetPayments возвращает Массив ссылок на связанные платежи.
func (purchaseReturn PurchaseReturn) GetPayments() Slice[Payment] {
	return purchaseReturn.Payments
}

// GetCode возвращает Код Возврата поставщику.
func (purchaseReturn PurchaseReturn) GetCode() string {
	return Deref(purchaseReturn.Code)
}

// GetOrganizationAccount возвращает Метаданные счета юрлица.
func (purchaseReturn PurchaseReturn) GetOrganizationAccount() AgentAccount {
	return Deref(purchaseReturn.OrganizationAccount)
}

// GetCreated возвращает Дату создания.
func (purchaseReturn PurchaseReturn) GetCreated() Timestamp {
	return Deref(purchaseReturn.Created)
}

// GetDeleted возвращает Момент последнего удаления Возврата поставщику.
func (purchaseReturn PurchaseReturn) GetDeleted() Timestamp {
	return Deref(purchaseReturn.Deleted)
}

// GetDescription возвращает Описание Возврата поставщику.
func (purchaseReturn PurchaseReturn) GetDescription() string {
	return Deref(purchaseReturn.Description)
}

// GetExternalCode возвращает Внешний код Возврата поставщику.
func (purchaseReturn PurchaseReturn) GetExternalCode() string {
	return Deref(purchaseReturn.ExternalCode)
}

// GetFiles возвращает Метаданные массива Файлов.
func (purchaseReturn PurchaseReturn) GetFiles() MetaArray[File] {
	return Deref(purchaseReturn.Files)
}

// GetGroup возвращает Отдел сотрудника.
func (purchaseReturn PurchaseReturn) GetGroup() Group {
	return Deref(purchaseReturn.Group)
}

// GetID возвращает ID Возврата поставщику.
func (purchaseReturn PurchaseReturn) GetID() uuid.UUID {
	return Deref(purchaseReturn.ID)
}

// GetMeta возвращает Метаданные Возврата поставщику.
func (purchaseReturn PurchaseReturn) GetMeta() Meta {
	return Deref(purchaseReturn.Meta)
}

// GetMoment возвращает Дату документа.
func (purchaseReturn PurchaseReturn) GetMoment() Timestamp {
	return Deref(purchaseReturn.Moment)
}

// GetName возвращает Наименование Возврата поставщику.
func (purchaseReturn PurchaseReturn) GetName() string {
	return Deref(purchaseReturn.Name)
}

// GetAccountID возвращает ID учётной записи.
func (purchaseReturn PurchaseReturn) GetAccountID() uuid.UUID {
	return Deref(purchaseReturn.AccountID)
}

// GetContract возвращает Метаданные договора.
func (purchaseReturn PurchaseReturn) GetContract() Contract {
	return Deref(purchaseReturn.Contract).GetValue()
}

// GetAgent возвращает Метаданные Контрагента.
func (purchaseReturn PurchaseReturn) GetAgent() Agent {
	return Deref(purchaseReturn.Agent)
}

// GetOrganization возвращает Метаданные юрлица.
func (purchaseReturn PurchaseReturn) GetOrganization() Organization {
	return Deref(purchaseReturn.Organization)
}

// GetProject возвращает Метаданные проекта.
func (purchaseReturn PurchaseReturn) GetProject() Project {
	return purchaseReturn.Project.GetValue()
}

// GetPublished возвращает true, если документ опубликован.
func (purchaseReturn PurchaseReturn) GetPublished() bool {
	return Deref(purchaseReturn.Published)
}

// GetRate возвращает Валюту.
func (purchaseReturn PurchaseReturn) GetRate() Rate {
	return purchaseReturn.Rate.GetValue()
}

// GetShared возвращает флаг Общего доступа.
func (purchaseReturn PurchaseReturn) GetShared() bool {
	return Deref(purchaseReturn.Shared)
}

// GetState возвращает Метаданные статуса Возврата поставщику.
func (purchaseReturn PurchaseReturn) GetState() State {
	return Deref(purchaseReturn.State).GetValue()
}

// GetStore возвращает Метаданные склада.
func (purchaseReturn PurchaseReturn) GetStore() Store {
	return Deref(purchaseReturn.Store)
}

// GetSum возвращает Сумму Возврата поставщику в копейках.
func (purchaseReturn PurchaseReturn) GetSum() float64 {
	return Deref(purchaseReturn.Sum)
}

// GetSyncID возвращает ID синхронизации.
func (purchaseReturn PurchaseReturn) GetSyncID() uuid.UUID {
	return Deref(purchaseReturn.SyncID)
}

// GetUpdated возвращает Момент последнего обновления Возврата поставщику.
func (purchaseReturn PurchaseReturn) GetUpdated() Timestamp {
	return Deref(purchaseReturn.Updated)
}

// GetVatEnabled возвращает true, если учитывается НДС.
func (purchaseReturn PurchaseReturn) GetVatEnabled() bool {
	return Deref(purchaseReturn.VatEnabled)
}

// GetVatIncluded возвращает true, если НДС включен в цену.
func (purchaseReturn PurchaseReturn) GetVatIncluded() bool {
	return Deref(purchaseReturn.VatIncluded)
}

// GetVatSum возвращает Сумму НДС.
func (purchaseReturn PurchaseReturn) GetVatSum() float64 {
	return Deref(purchaseReturn.VatSum)
}

// GetPositions возвращает Метаданные позиций Возврата поставщику.
func (purchaseReturn PurchaseReturn) GetPositions() MetaArray[PurchaseReturnPosition] {
	return Deref(purchaseReturn.Positions)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (purchaseReturn PurchaseReturn) GetOwner() Employee {
	return Deref(purchaseReturn.Owner)
}

// GetFactureIn возвращает Ссылку на Счет-фактуру полученный.
func (purchaseReturn PurchaseReturn) GetFactureIn() FactureIn {
	return Deref(purchaseReturn.FactureIn)
}

// GetFactureOut возвращает Ссылку на Счет-фактуру выданный.
func (purchaseReturn PurchaseReturn) GetFactureOut() FactureOut {
	return Deref(purchaseReturn.FactureOut)
}

// GetPayedSum возвращает Сумму входящих платежей по возврату поставщику.
func (purchaseReturn PurchaseReturn) GetPayedSum() float64 {
	return Deref(purchaseReturn.PayedSum)
}

// GetAttributes возвращает Список метаданных доп. полей.
func (purchaseReturn PurchaseReturn) GetAttributes() Slice[Attribute] {
	return purchaseReturn.Attributes
}

// SetSupply устанавливает Ссылку на приемку, по которой произошел возврат.
func (purchaseReturn *PurchaseReturn) SetSupply(supply *Supply) *PurchaseReturn {
	if supply != nil {
		purchaseReturn.Supply = supply.Clean()
	}
	return purchaseReturn
}

// SetAgentAccount устанавливает Метаданные счета контрагента.
func (purchaseReturn *PurchaseReturn) SetAgentAccount(agentAccount *AgentAccount) *PurchaseReturn {
	if agentAccount != nil {
		purchaseReturn.AgentAccount = agentAccount.Clean()
	}
	return purchaseReturn
}

// SetApplicable устанавливает Отметку о проведении.
func (purchaseReturn *PurchaseReturn) SetApplicable(applicable bool) *PurchaseReturn {
	purchaseReturn.Applicable = &applicable
	return purchaseReturn
}

// SetPayments устанавливает Метаданные ссылок на связанные платежи.
//
// Принимает множество объектов, реализующих интерфейс [PaymentConverter].
func (purchaseReturn *PurchaseReturn) SetPayments(payments ...PaymentConverter) *PurchaseReturn {
	purchaseReturn.Payments = NewPaymentsFrom(payments)
	return purchaseReturn
}

// SetCode устанавливает Код Возврата поставщику.
func (purchaseReturn *PurchaseReturn) SetCode(code string) *PurchaseReturn {
	purchaseReturn.Code = &code
	return purchaseReturn
}

// SetOrganizationAccount устанавливает Метаданные счета юрлица.
func (purchaseReturn *PurchaseReturn) SetOrganizationAccount(organizationAccount *AgentAccount) *PurchaseReturn {
	if organizationAccount != nil {
		purchaseReturn.OrganizationAccount = organizationAccount.Clean()
	}
	return purchaseReturn
}

// SetDescription устанавливает Комментарий Возврата поставщику.
func (purchaseReturn *PurchaseReturn) SetDescription(description string) *PurchaseReturn {
	purchaseReturn.Description = &description
	return purchaseReturn
}

// SetExternalCode устанавливает Внешний код Возврата поставщику.
func (purchaseReturn *PurchaseReturn) SetExternalCode(externalCode string) *PurchaseReturn {
	purchaseReturn.ExternalCode = &externalCode
	return purchaseReturn
}

// SetFiles устанавливает Метаданные массива Файлов.
//
// Принимает множество объектов [File].
func (purchaseReturn *PurchaseReturn) SetFiles(files ...*File) *PurchaseReturn {
	purchaseReturn.Files = NewMetaArrayFrom(files)
	return purchaseReturn
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (purchaseReturn *PurchaseReturn) SetGroup(group *Group) *PurchaseReturn {
	if group != nil {
		purchaseReturn.Group = group.Clean()
	}
	return purchaseReturn
}

// SetMeta устанавливает Метаданные Возврата поставщику.
func (purchaseReturn *PurchaseReturn) SetMeta(meta *Meta) *PurchaseReturn {
	purchaseReturn.Meta = meta
	return purchaseReturn
}

// SetMoment устанавливает Дату документа.
func (purchaseReturn *PurchaseReturn) SetMoment(moment time.Time) *PurchaseReturn {
	purchaseReturn.Moment = NewTimestamp(moment)
	return purchaseReturn
}

// SetName устанавливает Наименование доп. поля.
func (purchaseReturn *PurchaseReturn) SetName(name string) *PurchaseReturn {
	purchaseReturn.Name = &name
	return purchaseReturn
}

// SetContract устанавливает Метаданные договора.
//
// Передача nil передаёт сброс значения (null).
func (purchaseReturn *PurchaseReturn) SetContract(contract *Contract) *PurchaseReturn {
	purchaseReturn.Contract = NewNullValue(contract)
	return purchaseReturn
}

// SetAgent устанавливает Метаданные контрагента.
//
// Принимает [Counterparty] или [Organization].
func (purchaseReturn *PurchaseReturn) SetAgent(agent AgentOrganizationConverter) *PurchaseReturn {
	if agent != nil {
		purchaseReturn.Agent = agent.AsOrganizationAgent()
	}
	return purchaseReturn
}

// SetOrganization устанавливает Метаданные юрлица.
func (purchaseReturn *PurchaseReturn) SetOrganization(organization *Organization) *PurchaseReturn {
	if organization != nil {
		purchaseReturn.Organization = organization.Clean()
	}
	return purchaseReturn
}

// SetProject устанавливает Метаданные проекта.
//
// Передача nil передаёт сброс значения (null).
func (purchaseReturn *PurchaseReturn) SetProject(project *Project) *PurchaseReturn {
	purchaseReturn.Project = NewNullValue(project)
	return purchaseReturn
}

// SetRate устанавливает Валюту.
//
// Передача nil передаёт сброс значения (null).
func (purchaseReturn *PurchaseReturn) SetRate(rate *Rate) *PurchaseReturn {
	purchaseReturn.Rate = NewNullValue(rate)
	return purchaseReturn
}

// SetShared устанавливает флаг общего доступа.
func (purchaseReturn *PurchaseReturn) SetShared(shared bool) *PurchaseReturn {
	purchaseReturn.Shared = &shared
	return purchaseReturn
}

// SetState устанавливает Метаданные статуса Возврата поставщику.
//
// Передача nil передаёт сброс значения (null).
func (purchaseReturn *PurchaseReturn) SetState(state *State) *PurchaseReturn {
	purchaseReturn.State = NewNullValue(state)
	return purchaseReturn
}

// SetStore устанавливает Метаданные склада.
func (purchaseReturn *PurchaseReturn) SetStore(store *Store) *PurchaseReturn {
	if store != nil {
		purchaseReturn.Store = store.Clean()
	}
	return purchaseReturn
}

// SetSyncID устанавливает ID синхронизации.
func (purchaseReturn *PurchaseReturn) SetSyncID(syncID uuid.UUID) *PurchaseReturn {
	purchaseReturn.SyncID = &syncID
	return purchaseReturn
}

// SetVatEnabled устанавливает значение, учитывающее НДС для Возврата поставщику.
func (purchaseReturn *PurchaseReturn) SetVatEnabled(vatEnabled bool) *PurchaseReturn {
	purchaseReturn.VatEnabled = &vatEnabled
	return purchaseReturn
}

// SetVatIncluded устанавливает флаг включения НДС в цену.
func (purchaseReturn *PurchaseReturn) SetVatIncluded(vatIncluded bool) *PurchaseReturn {
	purchaseReturn.VatIncluded = &vatIncluded
	return purchaseReturn
}

// SetPositions устанавливает Метаданные позиций Возврата поставщику.
//
// Принимает множество объектов [PurchaseReturnPosition].
func (purchaseReturn *PurchaseReturn) SetPositions(positions ...*PurchaseReturnPosition) *PurchaseReturn {
	purchaseReturn.Positions = NewMetaArrayFrom(positions)
	return purchaseReturn
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (purchaseReturn *PurchaseReturn) SetOwner(owner *Employee) *PurchaseReturn {
	if owner != nil {
		purchaseReturn.Owner = owner.Clean()
	}
	return purchaseReturn
}

// SetFactureIn устанавливает Метаданные Счет-фактуры полученного.
func (purchaseReturn *PurchaseReturn) SetFactureIn(factureIn *FactureIn) *PurchaseReturn {
	if factureIn != nil {
		purchaseReturn.FactureIn = factureIn.Clean()
	}
	return purchaseReturn
}

// SetFactureOut устанавливает Метаданные Счет-фактуры выданного.
func (purchaseReturn *PurchaseReturn) SetFactureOut(factureOut *FactureOut) *PurchaseReturn {
	if factureOut != nil {
		purchaseReturn.FactureOut = factureOut.Clean()
	}
	return purchaseReturn
}

// SetPayedSum устанавливает Сумму входящих платежей по возврату поставщику.
func (purchaseReturn *PurchaseReturn) SetPayedSum(payedSum float64) *PurchaseReturn {
	purchaseReturn.PayedSum = &payedSum
	return purchaseReturn
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (purchaseReturn *PurchaseReturn) SetAttributes(attributes ...*Attribute) *PurchaseReturn {
	purchaseReturn.Attributes.Push(attributes...)
	return purchaseReturn
}

// String реализует интерфейс [fmt.Stringer].
func (purchaseReturn PurchaseReturn) String() string {
	return Stringify(purchaseReturn)
}

// MetaType возвращает код сущности.
func (PurchaseReturn) MetaType() MetaType {
	return MetaTypePurchaseReturn
}

// Update shortcut
func (purchaseReturn PurchaseReturn) Update(ctx context.Context, client *Client, params ...*Params) (*PurchaseReturn, *resty.Response, error) {
	return NewPurchaseReturnService(client).Update(ctx, purchaseReturn.GetID(), &purchaseReturn, params...)
}

// Create shortcut
func (purchaseReturn PurchaseReturn) Create(ctx context.Context, client *Client, params ...*Params) (*PurchaseReturn, *resty.Response, error) {
	return NewPurchaseReturnService(client).Create(ctx, &purchaseReturn, params...)
}

// Delete shortcut
func (purchaseReturn PurchaseReturn) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewPurchaseReturnService(client).Delete(ctx, purchaseReturn.GetID())
}

// PurchaseReturnPosition Позиция Возврата поставщику.
//
// Код сущности: purchasereturnposition
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vozwrat-postawschiku-vozwraty-postawschikam-pozicii-vozwrata-postawschiku
type PurchaseReturnPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учётной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	Discount   *float64            `json:"discount,omitempty"`   // Процент скидки или наценки. Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Price      *float64            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Slot       *Slot               `json:"slot,omitempty"`       // Ячейка на складе
	Vat        *int                `json:"vat,omitempty"`        // НДС, которым облагается текущая позиция
	VatEnabled *bool               `json:"vatEnabled,omitempty"` // Включен ли НДС для позиции. С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
	Stock      *Stock              `json:"stock,omitempty"`      // Остатки и себестоимость позиции (указывается при наличии параметра запроса `fields=stock`)
	Things     Slice[string]       `json:"things,omitempty"`     // Серийные номера. Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете. В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута.
}

// GetAccountID возвращает ID учётной записи.
func (purchaseReturnPosition PurchaseReturnPosition) GetAccountID() uuid.UUID {
	return Deref(purchaseReturnPosition.AccountID)
}

// GetAssortment возвращает Метаданные товара/услуги/серии/модификации, которую представляет собой позиция.
func (purchaseReturnPosition PurchaseReturnPosition) GetAssortment() AssortmentPosition {
	return Deref(purchaseReturnPosition.Assortment)
}

// GetDiscount возвращает Процент скидки или наценки.
//
// Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%.
func (purchaseReturnPosition PurchaseReturnPosition) GetDiscount() float64 {
	return Deref(purchaseReturnPosition.Discount)
}

// GetID возвращает ID позиции.
func (purchaseReturnPosition PurchaseReturnPosition) GetID() uuid.UUID {
	return Deref(purchaseReturnPosition.ID)
}

// GetPack возвращает Упаковку Товара.
func (purchaseReturnPosition PurchaseReturnPosition) GetPack() Pack {
	return Deref(purchaseReturnPosition.Pack)
}

// GetPrice возвращает Цену товара/услуги в копейках.
func (purchaseReturnPosition PurchaseReturnPosition) GetPrice() float64 {
	return Deref(purchaseReturnPosition.Price)
}

// GetQuantity возвращает Количество товаров данного вида в позиции.
//
// Если позиция - товар, у которого включен учет по серийным номерам,
// то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
func (purchaseReturnPosition PurchaseReturnPosition) GetQuantity() float64 {
	return Deref(purchaseReturnPosition.Quantity)
}

// GetSlot возвращает Ячейку на складе.
func (purchaseReturnPosition PurchaseReturnPosition) GetSlot() Slot {
	return Deref(purchaseReturnPosition.Slot)
}

// GetVat возвращает НДС %.
func (purchaseReturnPosition PurchaseReturnPosition) GetVat() int {
	return Deref(purchaseReturnPosition.Vat)
}

// GetVatEnabled возвращает true, если НДС включен для позиции.
//
// С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС".
// (vat = 0, vatEnabled = false) -> vat = "без НДС",
// (vat = 0, vatEnabled = true) -> vat = 0%.
func (purchaseReturnPosition PurchaseReturnPosition) GetVatEnabled() bool {
	return Deref(purchaseReturnPosition.VatEnabled)
}

// GetStock возвращает Остатки и себестоимость позиции (указывается при наличии параметра запроса `fields=stock`).
func (purchaseReturnPosition PurchaseReturnPosition) GetStock() Stock {
	return Deref(purchaseReturnPosition.Stock)
}

// GetThings возвращает Серийные номера.
//
// Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете.
// В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута.
func (purchaseReturnPosition PurchaseReturnPosition) GetThings() Slice[string] {
	return purchaseReturnPosition.Things
}

// SetAssortment устанавливает Метаданные товара/услуги/серии/модификации, которую представляет собой позиция.
//
// Принимает объект, реализующий интерфейс [AssortmentConverter].
func (purchaseReturnPosition *PurchaseReturnPosition) SetAssortment(assortment AssortmentConverter) *PurchaseReturnPosition {
	if assortment != nil {
		purchaseReturnPosition.Assortment = assortment.AsAssortment()
	}
	return purchaseReturnPosition
}

// SetDiscount устанавливает Процент скидки или наценки.
//
// Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%.
func (purchaseReturnPosition *PurchaseReturnPosition) SetDiscount(discount float64) *PurchaseReturnPosition {
	purchaseReturnPosition.Discount = &discount
	return purchaseReturnPosition
}

// SetPack устанавливает Упаковку Товара.
func (purchaseReturnPosition *PurchaseReturnPosition) SetPack(pack *Pack) *PurchaseReturnPosition {
	if pack != nil {
		purchaseReturnPosition.Pack = pack
	}
	return purchaseReturnPosition
}

// SetPrice устанавливает Цену товара/услуги в копейках.
func (purchaseReturnPosition *PurchaseReturnPosition) SetPrice(price float64) *PurchaseReturnPosition {
	purchaseReturnPosition.Price = &price
	return purchaseReturnPosition
}

// SetQuantity устанавливает Количество товаров данного вида в позиции.
//
// Если позиция - товар, у которого включен учет по серийным номерам,
// то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
func (purchaseReturnPosition *PurchaseReturnPosition) SetQuantity(quantity float64) *PurchaseReturnPosition {
	purchaseReturnPosition.Quantity = &quantity
	return purchaseReturnPosition
}

// SetSlot устанавливает Ячейку на складе.
func (purchaseReturnPosition *PurchaseReturnPosition) SetSlot(slot *Slot) *PurchaseReturnPosition {
	if slot != nil {
		purchaseReturnPosition.Slot = slot.Clean()
	}
	return purchaseReturnPosition
}

// SetVat устанавливает НДС, которым облагается текущая позиция.
func (purchaseReturnPosition *PurchaseReturnPosition) SetVat(vat int) *PurchaseReturnPosition {
	purchaseReturnPosition.Vat = &vat
	return purchaseReturnPosition
}

// SetVatEnabled устанавливает значение, учитывающее НДС для текущей позиции.
func (purchaseReturnPosition *PurchaseReturnPosition) SetVatEnabled(vatEnabled bool) *PurchaseReturnPosition {
	purchaseReturnPosition.VatEnabled = &vatEnabled
	return purchaseReturnPosition
}

// SetThings устанавливает Серийные номера.
//
// Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете.
// В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута.
//
// Принимает множество string.
func (purchaseReturnPosition *PurchaseReturnPosition) SetThings(things ...string) *PurchaseReturnPosition {
	purchaseReturnPosition.Things = NewSliceFrom(things)
	return purchaseReturnPosition
}

// String реализует интерфейс [fmt.Stringer].
func (purchaseReturnPosition PurchaseReturnPosition) String() string {
	return Stringify(purchaseReturnPosition)
}

// MetaType возвращает код сущности.
func (PurchaseReturnPosition) MetaType() MetaType {
	return MetaTypePurchaseReturnPosition
}

// PurchaseReturnService описывает метод сервиса для работы с возвратами поставщикам.
type PurchaseReturnService interface {
	// GetList выполняет запрос на получение списка возвратов поставщику.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[PurchaseReturn], *resty.Response, error)

	// Create выполняет запрос на создание возврата поставщику.
	// Обязательные поля для заполнения:
	//	- organization (Ссылка на ваше юрлицо)
	//	- store (Ссылка на склад)
	//	- agent (Ссылка на контрагента)
	// Принимает контекст, возврат поставщику и опционально объект параметров запроса Params.
	// Возвращает созданный возврат поставщику.
	Create(ctx context.Context, purchaseReturn *PurchaseReturn, params ...*Params) (*PurchaseReturn, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение возвратов поставщику.
	// Изменяемые возвраты поставщику должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список возвратов поставщику и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых возвратов поставщику.
	CreateUpdateMany(ctx context.Context, purchaseReturnList Slice[PurchaseReturn], params ...*Params) (*Slice[PurchaseReturn], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление возвратов поставщику.
	// Принимает контекст и множество возвратов поставщику.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*PurchaseReturn) (*DeleteManyResponse, *resty.Response, error)

	// Delete выполняет запрос на удаление возврата поставщику.
	// Принимает контекст и ID возврата поставщику.
	// Возвращает «true» в случае успешного удаления возврата поставщику.
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельного возврата поставщику по ID.
	// Принимает контекст, ID возврата поставщику и опционально объект параметров запроса Params.
	// Возвращает найденный возврат поставщику.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*PurchaseReturn, *resty.Response, error)

	// Update выполняет запрос на изменение возврата поставщику.
	// Принимает контекст, возврат поставщику и опционально объект параметров запроса Params.
	// Возвращает изменённый возврат поставщику.
	Update(ctx context.Context, id uuid.UUID, purchaseReturn *PurchaseReturn, params ...*Params) (*PurchaseReturn, *resty.Response, error)

	// Template выполняет запрос на получение предзаполненного возврата поставщику со стандартными полями без связи с какими-либо другими документами.
	// Принимает контекст.
	// Возвращает предзаполненный возврат поставщику.
	Template(ctx context.Context) (*PurchaseReturn, *resty.Response, error)

	// TemplateBased выполняет запрос на получение шаблона возврата поставщику на основе других документов.
	// Основание, на котором может быть создан:
	//	- Приемка (Supply)
	// Принимает контекст и множество документов из списка выше.
	// Возвращает предзаполненный возврат поставщику на основании переданных документов.
	TemplateBased(ctx context.Context, basedOn ...MetaOwner) (*PurchaseReturn, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных возвратов поставщику.
	// Принимает контекст.
	// Возвращает объект метаданных MetaAttributesStatesSharedWrapper.
	GetMetadata(ctx context.Context) (*MetaAttributesStatesSharedWrapper, *resty.Response, error)

	// GetPositionList выполняет запрос на получение списка позиций документа.
	// Принимает контекст, ID документа и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetPositionList(ctx context.Context, id uuid.UUID, params ...*Params) (*List[PurchaseReturnPosition], *resty.Response, error)

	// GetPositionByID выполняет запрос на получение отдельной позиции документа по ID.
	// Принимает контекст, ID документа, ID позиции и опционально объект параметров запроса Params.
	// Возвращает найденную позицию.
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params ...*Params) (*PurchaseReturnPosition, *resty.Response, error)

	// UpdatePosition выполняет запрос на изменение позиции документа.
	// Принимает контекст, ID документа, ID позиции, позицию документа и опционально объект параметров запроса Params.
	// Возвращает изменённую позицию.
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *PurchaseReturnPosition, params ...*Params) (*PurchaseReturnPosition, *resty.Response, error)

	// CreatePosition выполняет запрос на добавление позиции документа.
	// Принимает контекст, ID документа, позицию документа и опционально объект параметров запроса Params.
	// Возвращает добавленную позицию.
	CreatePosition(ctx context.Context, id uuid.UUID, position *PurchaseReturnPosition, params ...*Params) (*PurchaseReturnPosition, *resty.Response, error)

	// CreatePositionMany выполняет запрос на массовое добавление позиций документа.
	// Принимает контекст, ID документа и множество позиций.
	// Возвращает список добавленных позиций.
	CreatePositionMany(ctx context.Context, id uuid.UUID, positions ...*PurchaseReturnPosition) (*Slice[PurchaseReturnPosition], *resty.Response, error)

	// DeletePosition выполняет запрос на удаление позиции документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает «true» в случае успешного удаления позиции.
	DeletePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (bool, *resty.Response, error)

	// DeletePositionMany выполняет запрос на массовое удаление позиций документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeletePositionMany(ctx context.Context, id uuid.UUID, positions ...*PurchaseReturnPosition) (*DeleteManyResponse, *resty.Response, error)

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
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*PurchaseReturn, *resty.Response, error)

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
	Evaluate(ctx context.Context, entity *PurchaseReturn, evaluate ...Evaluate) (*PurchaseReturn, *resty.Response, error)
}

// NewPurchaseReturnService принимает [Client] и возвращает сервис для работы с возвратами поставщикам.
func NewPurchaseReturnService(client *Client) PurchaseReturnService {
	return newMainService[PurchaseReturn, PurchaseReturnPosition, MetaAttributesStatesSharedWrapper, any](NewEndpoint(client, "entity/purchasereturn"))
}
