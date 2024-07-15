package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"time"
)

// RetailSalesReturn Розничный возврат.
//
// Код сущности: retailsalesreturn
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-roznichnyj-wozwrat
type RetailSalesReturn struct {
	Name                *string                               `json:"name,omitempty"`                // Наименование Розничного возврата
	Organization        *Organization                         `json:"organization,omitempty"`        // Метаданные юрлица
	AgentAccount        *AgentAccount                         `json:"agentAccount,omitempty"`        // Метаданные счета контрагента
	Applicable          *bool                                 `json:"applicable,omitempty"`          // Отметка о проведении
	VatIncluded         *bool                                 `json:"vatIncluded,omitempty"`         // Включен ли НДС в цену
	CashSum             *float64                              `json:"cashSum,omitempty"`             // Оплачено наличными
	Code                *string                               `json:"code,omitempty"`                // Код Розничного возврата
	Contract            *NullValue[Contract]                  `json:"contract,omitempty"`            // Метаданные договора
	Created             *Timestamp                            `json:"created,omitempty"`             // Дата создания
	Deleted             *Timestamp                            `json:"deleted,omitempty"`             // Момент последнего удаления Розничного возврата
	Demand              *RetailDemand                         `json:"demand,omitempty"`              // Метаданные розничной продажи, по которой произошел возврат
	Description         *string                               `json:"description,omitempty"`         // Комментарий Розничного возврата
	ExternalCode        *string                               `json:"externalCode,omitempty"`        // Внешний код Розничного возврата
	Group               *Group                                `json:"group,omitempty"`               // Отдел сотрудника
	ID                  *uuid.UUID                            `json:"id,omitempty"`                  // ID Розничного возврата
	Meta                *Meta                                 `json:"meta,omitempty"`                // Метаданные Розничного возврата
	Moment              *Timestamp                            `json:"moment,omitempty"`              // Дата документа
	OrganizationAccount *AgentAccount                         `json:"organizationAccount,omitempty"` // Метаданные счета юрлица
	NoCashSum           *float64                              `json:"noCashSum,omitempty"`           // Оплачено картой
	SyncID              *uuid.UUID                            `json:"syncId,omitempty"`              // ID синхронизации
	AccountID           *uuid.UUID                            `json:"accountId,omitempty"`           // ID учётной записи
	Owner               *Employee                             `json:"owner,omitempty"`               // Метаданные владельца (Сотрудника)
	Positions           *MetaArray[RetailSalesReturnPosition] `json:"positions,omitempty"`           // Метаданные позиций Розничного возврата
	Printed             *bool                                 `json:"printed,omitempty"`             // Напечатан ли документ
	Project             *NullValue[Project]                   `json:"project,omitempty"`             // Метаданные проекта
	Published           *bool                                 `json:"published,omitempty"`           // Опубликован ли документ
	QRSum               *float64                              `json:"qrSum,omitempty"`               // Оплачено по QR-коду
	Rate                *NullValue[Rate]                      `json:"rate,omitempty"`                // Валюта
	RetailShift         *RetailShift                          `json:"retailShift,omitempty"`         // Метаданные Розничной смены
	RetailStore         *RetailStore                          `json:"retailStore,omitempty"`         // Метаданные Точки продаж
	Shared              *bool                                 `json:"shared,omitempty"`              // Общий доступ
	State               *NullValue[State]                     `json:"state,omitempty"`               // Метаданные статуса Розничного возврата
	Store               *Store                                `json:"store,omitempty"`               // Метаданные склада
	Sum                 *float64                              `json:"sum,omitempty"`                 // Сумма Розничного возврата в копейках
	Agent               *Agent                                `json:"agent,omitempty"`               // Метаданные контрагента
	VatSum              *float64                              `json:"vatSum,omitempty"`              // Сумма НДС
	Updated             *Timestamp                            `json:"updated,omitempty"`             // Момент последнего обновления Розничного возврата
	VatEnabled          *bool                                 `json:"vatEnabled,omitempty"`          // Учитывается ли НДС
	TaxSystem           TaxSystem                             `json:"taxSystem,omitempty"`           // Код системы налогообложения
	Attributes          Slice[Attribute]                      `json:"attributes,omitempty"`          // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (retailSalesReturn RetailSalesReturn) Clean() *RetailSalesReturn {
	if retailSalesReturn.Meta == nil {
		return nil
	}
	return &RetailSalesReturn{Meta: retailSalesReturn.Meta}
}

// AsTaskOperation реализует интерфейс [TaskOperationConverter].
func (retailSalesReturn RetailSalesReturn) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: retailSalesReturn.Meta}
}

// GetName возвращает Наименование Розничного возврата.
func (retailSalesReturn RetailSalesReturn) GetName() string {
	return Deref(retailSalesReturn.Name)
}

// GetOrganization возвращает Метаданные юрлица.
func (retailSalesReturn RetailSalesReturn) GetOrganization() Organization {
	return Deref(retailSalesReturn.Organization)
}

// GetAgentAccount возвращает Метаданные счета контрагента.
func (retailSalesReturn RetailSalesReturn) GetAgentAccount() AgentAccount {
	return Deref(retailSalesReturn.AgentAccount)
}

// GetApplicable возвращает Отметку о проведении.
func (retailSalesReturn RetailSalesReturn) GetApplicable() bool {
	return Deref(retailSalesReturn.Applicable)
}

// GetVatIncluded возвращает true, если НДС включен в цену.
func (retailSalesReturn RetailSalesReturn) GetVatIncluded() bool {
	return Deref(retailSalesReturn.VatIncluded)
}

// GetCashSum возвращает Оплачено наличными.
func (retailSalesReturn RetailSalesReturn) GetCashSum() float64 {
	return Deref(retailSalesReturn.CashSum)
}

// GetCode возвращает Код Розничного возврата.
func (retailSalesReturn RetailSalesReturn) GetCode() string {
	return Deref(retailSalesReturn.Code)
}

// GetContract возвращает Метаданные договора.
func (retailSalesReturn RetailSalesReturn) GetContract() Contract {
	return Deref(retailSalesReturn.Contract).getValue()
}

// GetCreated возвращает Дату создания.
func (retailSalesReturn RetailSalesReturn) GetCreated() time.Time {
	return Deref(retailSalesReturn.Created).Time()
}

// GetDeleted возвращает Момент последнего удаления Розничного возврата.
func (retailSalesReturn RetailSalesReturn) GetDeleted() time.Time {
	return Deref(retailSalesReturn.Deleted).Time()
}

// GetDemand возвращает Метаданные розничной продажи, по которой произошел возврат.
func (retailSalesReturn RetailSalesReturn) GetDemand() RetailDemand {
	return Deref(retailSalesReturn.Demand)
}

// GetDescription возвращает Комментарий Розничного возврата.
func (retailSalesReturn RetailSalesReturn) GetDescription() string {
	return Deref(retailSalesReturn.Description)
}

// GetExternalCode возвращает Внешний код Розничного возврата.
func (retailSalesReturn RetailSalesReturn) GetExternalCode() string {
	return Deref(retailSalesReturn.ExternalCode)
}

// GetGroup возвращает Отдел сотрудника.
func (retailSalesReturn RetailSalesReturn) GetGroup() Group {
	return Deref(retailSalesReturn.Group)
}

// GetID возвращает ID Розничного возврата.
func (retailSalesReturn RetailSalesReturn) GetID() uuid.UUID {
	return Deref(retailSalesReturn.ID)
}

// GetMeta возвращает Метаданные Розничного возврата.
func (retailSalesReturn RetailSalesReturn) GetMeta() Meta {
	return Deref(retailSalesReturn.Meta)
}

// GetMoment возвращает Дату документа.
func (retailSalesReturn RetailSalesReturn) GetMoment() time.Time {
	return Deref(retailSalesReturn.Moment).Time()
}

// GetOrganizationAccount возвращает Метаданные счета юрлица.
func (retailSalesReturn RetailSalesReturn) GetOrganizationAccount() AgentAccount {
	return Deref(retailSalesReturn.OrganizationAccount)
}

// GetNoCashSum возвращает Оплачено картой.
func (retailSalesReturn RetailSalesReturn) GetNoCashSum() float64 {
	return Deref(retailSalesReturn.NoCashSum)
}

// GetSyncID возвращает ID синхронизации.
func (retailSalesReturn RetailSalesReturn) GetSyncID() uuid.UUID {
	return Deref(retailSalesReturn.SyncID)
}

// GetAccountID возвращает ID учётной записи.
func (retailSalesReturn RetailSalesReturn) GetAccountID() uuid.UUID {
	return Deref(retailSalesReturn.AccountID)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (retailSalesReturn RetailSalesReturn) GetOwner() Employee {
	return Deref(retailSalesReturn.Owner)
}

// GetPositions возвращает Метаданные позиций Розничного возврата.
func (retailSalesReturn RetailSalesReturn) GetPositions() MetaArray[RetailSalesReturnPosition] {
	return Deref(retailSalesReturn.Positions)
}

// GetPrinted возвращает true, если документ напечатан.
func (retailSalesReturn RetailSalesReturn) GetPrinted() bool {
	return Deref(retailSalesReturn.Printed)
}

// GetProject возвращает Метаданные проекта.
func (retailSalesReturn RetailSalesReturn) GetProject() Project {
	return Deref(retailSalesReturn.Project).getValue()
}

// GetPublished возвращает true, если документ опубликован.
func (retailSalesReturn RetailSalesReturn) GetPublished() bool {
	return Deref(retailSalesReturn.Published)
}

// GetQRSum возвращает оплачено по QR-коду.
func (retailSalesReturn RetailSalesReturn) GetQRSum() float64 {
	return Deref(retailSalesReturn.QRSum)
}

// GetRate возвращает Валюту.
func (retailSalesReturn RetailSalesReturn) GetRate() Rate {
	return Deref(retailSalesReturn.Rate).getValue()
}

// GetRetailShift возвращает Метаданные Розничной смены.
func (retailSalesReturn RetailSalesReturn) GetRetailShift() RetailShift {
	return Deref(retailSalesReturn.RetailShift)
}

// GetRetailStore возвращает Метаданные Точки продаж.
func (retailSalesReturn RetailSalesReturn) GetRetailStore() RetailStore {
	return Deref(retailSalesReturn.RetailStore)
}

// GetShared возвращает флаг Общего доступа.
func (retailSalesReturn RetailSalesReturn) GetShared() bool {
	return Deref(retailSalesReturn.Shared)
}

// GetState возвращает Метаданные статуса Розничного возврата.
func (retailSalesReturn RetailSalesReturn) GetState() State {
	return Deref(retailSalesReturn.State).getValue()
}

// GetStore возвращает Метаданные склада.
func (retailSalesReturn RetailSalesReturn) GetStore() Store {
	return Deref(retailSalesReturn.Store)
}

// GetSum возвращает Сумму Розничного возврата в копейках.
func (retailSalesReturn RetailSalesReturn) GetSum() float64 {
	return Deref(retailSalesReturn.Sum)
}

// GetAgent возвращает Метаданные Контрагента.
func (retailSalesReturn RetailSalesReturn) GetAgent() Agent {
	return Deref(retailSalesReturn.Agent)
}

// GetVatSum возвращает Сумму НДС.
func (retailSalesReturn RetailSalesReturn) GetVatSum() float64 {
	return Deref(retailSalesReturn.VatSum)
}

// GetUpdated возвращает Момент последнего обновления Розничного возврата.
func (retailSalesReturn RetailSalesReturn) GetUpdated() time.Time {
	return Deref(retailSalesReturn.Updated).Time()
}

// GetVatEnabled возвращает true, если учитывается НДС.
func (retailSalesReturn RetailSalesReturn) GetVatEnabled() bool {
	return Deref(retailSalesReturn.VatEnabled)
}

// GetTaxSystem возвращает Код системы налогообложения.
func (retailSalesReturn RetailSalesReturn) GetTaxSystem() TaxSystem {
	return retailSalesReturn.TaxSystem
}

// GetAttributes возвращает Список метаданных доп. полей.
func (retailSalesReturn RetailSalesReturn) GetAttributes() Slice[Attribute] {
	return retailSalesReturn.Attributes
}

// SetName устанавливает Наименование Розничного возврата.
func (retailSalesReturn *RetailSalesReturn) SetName(name string) *RetailSalesReturn {
	retailSalesReturn.Name = &name
	return retailSalesReturn
}

// SetOrganization устанавливает Метаданные юрлица.
func (retailSalesReturn *RetailSalesReturn) SetOrganization(organization *Organization) *RetailSalesReturn {
	if organization != nil {
		retailSalesReturn.Organization = organization.Clean()
	}
	return retailSalesReturn
}

// SetAgentAccount устанавливает Метаданные счета контрагента.
func (retailSalesReturn *RetailSalesReturn) SetAgentAccount(agentAccount *AgentAccount) *RetailSalesReturn {
	if agentAccount != nil {
		retailSalesReturn.AgentAccount = agentAccount.Clean()
	}
	return retailSalesReturn
}

// SetApplicable устанавливает Отметку о проведении.
func (retailSalesReturn *RetailSalesReturn) SetApplicable(applicable bool) *RetailSalesReturn {
	retailSalesReturn.Applicable = &applicable
	return retailSalesReturn
}

// SetVatIncluded устанавливает флаг включения НДС в цену.
func (retailSalesReturn *RetailSalesReturn) SetVatIncluded(vatIncluded bool) *RetailSalesReturn {
	retailSalesReturn.VatIncluded = &vatIncluded
	return retailSalesReturn
}

// SetCashSum устанавливает Оплачено наличными.
func (retailSalesReturn *RetailSalesReturn) SetCashSum(cashSum float64) *RetailSalesReturn {
	retailSalesReturn.CashSum = &cashSum
	return retailSalesReturn
}

// SetCode устанавливает Код Розничного возврата.
func (retailSalesReturn *RetailSalesReturn) SetCode(code string) *RetailSalesReturn {
	retailSalesReturn.Code = &code
	return retailSalesReturn
}

// SetContract устанавливает Метаданные договора.
//
// Передача nil передаёт сброс значения (null).
func (retailSalesReturn *RetailSalesReturn) SetContract(contract *Contract) *RetailSalesReturn {
	retailSalesReturn.Contract = NewNullValue(contract)
	return retailSalesReturn
}

// SetDemand устанавливает Метаданные розничной продажи, по которой произошел возврат.
func (retailSalesReturn *RetailSalesReturn) SetDemand(demand *RetailDemand) *RetailSalesReturn {
	if demand != nil {
		retailSalesReturn.Demand = demand.Clean()
	}
	return retailSalesReturn
}

// SetDescription устанавливает Комментарий Розничного возврата.
func (retailSalesReturn *RetailSalesReturn) SetDescription(description string) *RetailSalesReturn {
	retailSalesReturn.Description = &description
	return retailSalesReturn
}

// SetExternalCode устанавливает Внешний код Розничного возврата.
func (retailSalesReturn *RetailSalesReturn) SetExternalCode(externalCode string) *RetailSalesReturn {
	retailSalesReturn.ExternalCode = &externalCode
	return retailSalesReturn
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (retailSalesReturn *RetailSalesReturn) SetGroup(group *Group) *RetailSalesReturn {
	if group != nil {
		retailSalesReturn.Group = group.Clean()
	}
	return retailSalesReturn
}

// SetMeta устанавливает Метаданные Розничного возврата.
func (retailSalesReturn *RetailSalesReturn) SetMeta(meta *Meta) *RetailSalesReturn {
	retailSalesReturn.Meta = meta
	return retailSalesReturn
}

// SetMoment устанавливает Дату документа.
func (retailSalesReturn *RetailSalesReturn) SetMoment(moment time.Time) *RetailSalesReturn {
	retailSalesReturn.Moment = NewTimestamp(moment)
	return retailSalesReturn
}

// SetOrganizationAccount устанавливает Метаданные счета юрлица.
func (retailSalesReturn *RetailSalesReturn) SetOrganizationAccount(organizationAccount *AgentAccount) *RetailSalesReturn {
	if organizationAccount != nil {
		retailSalesReturn.OrganizationAccount = organizationAccount.Clean()
	}
	return retailSalesReturn
}

// SetNoCashSum устанавливает Оплачено картой.
func (retailSalesReturn *RetailSalesReturn) SetNoCashSum(noCashSum float64) *RetailSalesReturn {
	retailSalesReturn.NoCashSum = &noCashSum
	return retailSalesReturn
}

// SetSyncID устанавливает ID синхронизации.
func (retailSalesReturn *RetailSalesReturn) SetSyncID(syncID uuid.UUID) *RetailSalesReturn {
	retailSalesReturn.SyncID = &syncID
	return retailSalesReturn
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (retailSalesReturn *RetailSalesReturn) SetOwner(owner *Employee) *RetailSalesReturn {
	if owner != nil {
		retailSalesReturn.Owner = owner.Clean()
	}
	return retailSalesReturn
}

// SetPositions устанавливает Метаданные позиций Розничного возврата.
//
// Принимает множество объектов [RetailSalesReturnPosition].
func (retailSalesReturn *RetailSalesReturn) SetPositions(positions ...*RetailSalesReturnPosition) *RetailSalesReturn {
	retailSalesReturn.Positions = NewMetaArrayFrom(positions)
	return retailSalesReturn
}

// SetProject устанавливает Метаданные проекта.
//
// Передача nil передаёт сброс значения (null).
func (retailSalesReturn *RetailSalesReturn) SetProject(project *Project) *RetailSalesReturn {
	retailSalesReturn.Project = NewNullValue(project)
	return retailSalesReturn
}

// SetQRSum устанавливает Оплачено по QR-коду.
func (retailSalesReturn *RetailSalesReturn) SetQRSum(qrSum float64) *RetailSalesReturn {
	retailSalesReturn.QRSum = &qrSum
	return retailSalesReturn
}

// SetRate устанавливает Валюту.
//
// Передача nil передаёт сброс значения (null).
func (retailSalesReturn *RetailSalesReturn) SetRate(rate *Rate) *RetailSalesReturn {
	retailSalesReturn.Rate = NewNullValue(rate)
	return retailSalesReturn
}

// SetRetailShift устанавливает Метаданные Розничной смены.
func (retailSalesReturn *RetailSalesReturn) SetRetailShift(retailShift *RetailShift) *RetailSalesReturn {
	if retailShift != nil {
		retailSalesReturn.RetailShift = retailShift.Clean()
	}
	return retailSalesReturn
}

// SetRetailStore устанавливает Метаданные Точки продаж.
func (retailSalesReturn *RetailSalesReturn) SetRetailStore(retailStore *RetailStore) *RetailSalesReturn {
	if retailStore != nil {
		retailSalesReturn.RetailStore = retailStore.Clean()
	}
	return retailSalesReturn
}

// SetShared устанавливает флаг общего доступа.
func (retailSalesReturn *RetailSalesReturn) SetShared(shared bool) *RetailSalesReturn {
	retailSalesReturn.Shared = &shared
	return retailSalesReturn
}

// SetState устанавливает Метаданные статуса Списания.
//
// Передача nil передаёт сброс значения (null).
func (retailSalesReturn *RetailSalesReturn) SetState(state *State) *RetailSalesReturn {
	retailSalesReturn.State = NewNullValue(state)
	return retailSalesReturn
}

// SetStore устанавливает Метаданные склада.
func (retailSalesReturn *RetailSalesReturn) SetStore(store *Store) *RetailSalesReturn {
	if store != nil {
		retailSalesReturn.Store = store.Clean()
	}
	return retailSalesReturn
}

// SetAgent устанавливает Метаданные контрагента.
//
// Принимает [Counterparty] или [Organization].
func (retailSalesReturn *RetailSalesReturn) SetAgent(agent AgentOrganizationConverter) *RetailSalesReturn {
	if agent != nil {
		retailSalesReturn.Agent = agent.AsOrganizationAgent()
	}
	return retailSalesReturn
}

// SetVatEnabled устанавливает значение, учитывающее НДС.
func (retailSalesReturn *RetailSalesReturn) SetVatEnabled(vatEnabled bool) *RetailSalesReturn {
	retailSalesReturn.VatEnabled = &vatEnabled
	return retailSalesReturn
}

// SetTaxSystem устанавливает Код системы налогообложения.
func (retailSalesReturn *RetailSalesReturn) SetTaxSystem(taxSystem TaxSystem) *RetailSalesReturn {
	retailSalesReturn.TaxSystem = taxSystem
	return retailSalesReturn
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (retailSalesReturn *RetailSalesReturn) SetAttributes(attributes ...*Attribute) *RetailSalesReturn {
	retailSalesReturn.Attributes.Push(attributes...)
	return retailSalesReturn
}

// String реализует интерфейс [fmt.Stringer].
func (retailSalesReturn RetailSalesReturn) String() string {
	return Stringify(retailSalesReturn)
}

// MetaType возвращает код сущности.
func (retailSalesReturn RetailSalesReturn) MetaType() MetaType {
	return MetaTypeRetailSalesReturn
}

// Update shortcut
func (retailSalesReturn *RetailSalesReturn) Update(ctx context.Context, client *Client, params ...*Params) (*RetailSalesReturn, *resty.Response, error) {
	return NewRetailSalesReturnService(client).Update(ctx, retailSalesReturn.GetID(), retailSalesReturn, params...)
}

// Create shortcut
func (retailSalesReturn *RetailSalesReturn) Create(ctx context.Context, client *Client, params ...*Params) (*RetailSalesReturn, *resty.Response, error) {
	return NewRetailSalesReturnService(client).Create(ctx, retailSalesReturn, params...)
}

// Delete shortcut
func (retailSalesReturn *RetailSalesReturn) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewRetailSalesReturnService(client).Delete(ctx, retailSalesReturn)
}

// RetailSalesReturnPosition позиция розничного возврата.
//
// Код сущности: salesreturnposition
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-roznichnyj-wozwrat-roznichnye-wozwraty-pozicii-roznichnogo-wozwrata
type RetailSalesReturnPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учётной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	Cost       *float64            `json:"cost,omitempty"`       // Себестоимость (выводится, если документ был создан без основания)
	Discount   *float64            `json:"discount,omitempty"`   // Процент скидки или наценки. Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Price      *float64            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Vat        *int                `json:"vat,omitempty"`        // НДС, которым облагается текущая позиция
	VatEnabled *bool               `json:"vatEnabled,omitempty"` // Включен ли НДС для позиции. С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
	Stock      *Stock              `json:"stock,omitempty"`      // Остатки и себестоимость позиции (указывается при наличии параметра запроса `fields=stock`)
	Things     Slice[string]       `json:"things,omitempty"`     // Серийные номера. Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете. В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута.
}

// GetAccountID возвращает ID учётной записи.
func (retailSalesReturnPosition RetailSalesReturnPosition) GetAccountID() uuid.UUID {
	return Deref(retailSalesReturnPosition.AccountID)
}

// GetAssortment возвращает Метаданные товара/услуги/серии/модификации, которую представляет собой позиция.
func (retailSalesReturnPosition RetailSalesReturnPosition) GetAssortment() AssortmentPosition {
	return Deref(retailSalesReturnPosition.Assortment)
}

// GetCost возвращает Себестоимость (выводится, если документ был создан без основания).
func (retailSalesReturnPosition RetailSalesReturnPosition) GetCost() float64 {
	return Deref(retailSalesReturnPosition.Cost)
}

// GetDiscount возвращает Процент скидки или наценки.
//
// Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%.
func (retailSalesReturnPosition RetailSalesReturnPosition) GetDiscount() float64 {
	return Deref(retailSalesReturnPosition.Discount)
}

// GetID возвращает ID позиции.
func (retailSalesReturnPosition RetailSalesReturnPosition) GetID() uuid.UUID {
	return Deref(retailSalesReturnPosition.ID)
}

// GetPack возвращает Упаковку Товара.
func (retailSalesReturnPosition RetailSalesReturnPosition) GetPack() Pack {
	return Deref(retailSalesReturnPosition.Pack)
}

// GetPrice возвращает Цену товара/услуги в копейках.
func (retailSalesReturnPosition RetailSalesReturnPosition) GetPrice() float64 {
	return Deref(retailSalesReturnPosition.Price)
}

// GetQuantity возвращает Количество товаров/услуг данного вида в позиции.
//
// Если позиция - товар, у которого включен учет по серийным номерам,
// то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
func (retailSalesReturnPosition RetailSalesReturnPosition) GetQuantity() float64 {
	return Deref(retailSalesReturnPosition.Quantity)
}

// GetThings возвращает Серийные номера.
//
// Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете.
// В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута.
func (retailSalesReturnPosition RetailSalesReturnPosition) GetThings() Slice[string] {
	return retailSalesReturnPosition.Things
}

// GetVat возвращает НДС, которым облагается текущая позиция.
func (retailSalesReturnPosition RetailSalesReturnPosition) GetVat() int {
	return Deref(retailSalesReturnPosition.Vat)
}

// GetVatEnabled возвращает true, если учитывается НДС.
//
// С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС".
//
// (vat = 0, vatEnabled = false) -> vat = "без НДС".
//
// (vat = 0, vatEnabled = true) -> vat = 0%.
func (retailSalesReturnPosition RetailSalesReturnPosition) GetVatEnabled() bool {
	return Deref(retailSalesReturnPosition.VatEnabled)
}

// GetStock возвращает Остатки и себестоимость позиции (указывается при наличии параметра запроса `fields=stock`).
func (retailSalesReturnPosition RetailSalesReturnPosition) GetStock() Stock {
	return Deref(retailSalesReturnPosition.Stock)
}

// SetAssortment устанавливает Метаданные товара/услуги/серии/модификации, которую представляет собой позиция.
//
// Принимает объект, реализующий интерфейс [AssortmentConverter].
func (retailSalesReturnPosition *RetailSalesReturnPosition) SetAssortment(assortment AssortmentConverter) *RetailSalesReturnPosition {
	if assortment != nil {
		retailSalesReturnPosition.Assortment = assortment.AsAssortment()
	}
	return retailSalesReturnPosition
}

// SetCost устанавливает Себестоимость  (выводится, если документ был создан без основания).
func (retailSalesReturnPosition *RetailSalesReturnPosition) SetCost(cost float64) *RetailSalesReturnPosition {
	retailSalesReturnPosition.Cost = &cost
	return retailSalesReturnPosition
}

// SetDiscount устанавливает Процент скидки или наценки.
//
// Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%.
func (retailSalesReturnPosition *RetailSalesReturnPosition) SetDiscount(discount float64) *RetailSalesReturnPosition {
	retailSalesReturnPosition.Discount = &discount
	return retailSalesReturnPosition
}

// SetPack устанавливает Упаковку Товара.
func (retailSalesReturnPosition *RetailSalesReturnPosition) SetPack(pack *Pack) *RetailSalesReturnPosition {
	if pack != nil {
		retailSalesReturnPosition.Pack = pack
	}
	return retailSalesReturnPosition
}

// SetPrice устанавливает Цену товара/услуги в копейках.
func (retailSalesReturnPosition *RetailSalesReturnPosition) SetPrice(price float64) *RetailSalesReturnPosition {
	retailSalesReturnPosition.Price = &price
	return retailSalesReturnPosition
}

// SetQuantity устанавливает Количество товаров данного вида в позиции.
//
// Если позиция - товар, у которого включен учет по серийным номерам,
// то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
func (retailSalesReturnPosition *RetailSalesReturnPosition) SetQuantity(quantity float64) *RetailSalesReturnPosition {
	retailSalesReturnPosition.Quantity = &quantity
	return retailSalesReturnPosition
}

// SetVat устанавливает НДС, которым облагается текущая позиция.
func (retailSalesReturnPosition *RetailSalesReturnPosition) SetVat(vat int) *RetailSalesReturnPosition {
	retailSalesReturnPosition.Vat = &vat
	return retailSalesReturnPosition
}

// SetVatEnabled устанавливает значение, учитывающее НДС для Группы товаров.
//
// С помощью этого флага для группы можно выставлять НДС = 0 или НДС = "без НДС".
//
// (vat = 0, vatEnabled = false) -> vat = "без НДС"
//
// (vat = 0, vatEnabled = true) -> vat = 0%.
func (retailSalesReturnPosition *RetailSalesReturnPosition) SetVatEnabled(vatEnabled bool) *RetailSalesReturnPosition {
	retailSalesReturnPosition.VatEnabled = &vatEnabled
	return retailSalesReturnPosition
}

// SetThings устанавливает Серийные номера.
//
// Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете.
// В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута.
//
// Принимает множество string.
func (retailSalesReturnPosition *RetailSalesReturnPosition) SetThings(things ...string) *RetailSalesReturnPosition {
	retailSalesReturnPosition.Things = NewSliceFrom(things)
	return retailSalesReturnPosition
}

// String реализует интерфейс [fmt.Stringer].
func (retailSalesReturnPosition RetailSalesReturnPosition) String() string {
	return Stringify(retailSalesReturnPosition)
}

// MetaType возвращает код сущности.
func (RetailSalesReturnPosition) MetaType() MetaType {
	return MetaTypeRetailSalesReturnPosition
}

// RetailSalesReturnService описывает методы сервиса для работы с розничными возвратами.
type RetailSalesReturnService interface {
	// GetList выполняет запрос на получение списка розничных возвратов.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[RetailSalesReturn], *resty.Response, error)

	// Create выполняет запрос на создание внесения денег.
	// Обязательные поля для заполнения:
	//	- name -(омер возврата)
	//	- organization (Ссылка на ваше юрлицо)
	//	- store (Ссылка на склад)
	//	- demand (Ссылка на Розничную продажу, по которой произошел возврат, обязательное поле только для возврата на основании)
	//	- retailStore (Ссылка на точку продаж)
	//	- retailShift (Ссылка на Розничную смену, в рамках которой происходит возврат)
	//	- agent (Ссылка на контрагента. Контрагент, указанный в запросе на создание возврата, должен совпадать с контрагентом, указанном в документе, по которому создается возврат)
	//	- cashSum (Оплачено наличными. Поле является необходимым для возврата без основания)
	//	- noCashSum (Оплачено картой. Поле является необходимым для возврата без основания)
	// Принимает контекст, розничный возврат и опционально объект параметров запроса Params.
	// Возвращает созданный розничный возврат.
	Create(ctx context.Context, retailSalesReturn *RetailSalesReturn, params ...*Params) (*RetailSalesReturn, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение розничных возвратов.
	// Изменяемые розничные возвраты должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список розничных возвратов и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых розничных возвратов.
	CreateUpdateMany(ctx context.Context, retailSalesReturnList Slice[RetailSalesReturn], params ...*Params) (*Slice[RetailSalesReturn], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление розничных возвратов.
	// Принимает контекст и множество розничных возвратов.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*RetailSalesReturn) (*DeleteManyResponse, *resty.Response, error)

	// DeleteByID выполняет запрос на удаление розничного возврата по ID.
	// Принимает контекст и ID розничного возврата.
	// Возвращает «true» в случае успешного удаления розничного возврата.
	DeleteByID(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// Delete выполняет запрос на удаление розничного возврата.
	// Принимает контекст и розничный возврат.
	// Возвращает «true» в случае успешного удаления розничного возврата.
	Delete(ctx context.Context, entity *RetailSalesReturn) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение розничного возврата по ID.
	// Принимает контекст, ID розничного возврата и опционально объект параметров запроса Params.
	// Возвращает розничный возврат.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*RetailSalesReturn, *resty.Response, error)

	// Update выполняет запрос на изменение розничного возврата.
	// Принимает контекст, розничный возврат и опционально объект параметров запроса Params.
	// Возвращает изменённый розничный возврат.
	Update(ctx context.Context, id uuid.UUID, retailSalesReturn *RetailSalesReturn, params ...*Params) (*RetailSalesReturn, *resty.Response, error)

	// TemplateBased выполняет запрос на получение шаблона розничного возврата на основе других документов.
	// Основание, на котором может быть создано:
	//	- Розничная смена (RetailShift)
	//	- Розничная продажа (RetailDemand)
	// Принимает контекст и множество документов из списка выше.
	// Возвращает предзаполненный розничный возврат на основании переданных документов.
	TemplateBased(ctx context.Context, basedOn ...MetaOwner) (*RetailSalesReturn, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных розничных возвратов.
	// Принимает контекст.
	// Возвращает объект метаданных MetaAttributesStatesSharedWrapper.
	GetMetadata(ctx context.Context) (*MetaAttributesStatesSharedWrapper, *resty.Response, error)

	// GetPositionList выполняет запрос на получение списка позиций документа.
	// Принимает контекст, ID документа и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetPositionList(ctx context.Context, id uuid.UUID, params ...*Params) (*List[RetailSalesReturnPosition], *resty.Response, error)

	// GetPositionByID выполняет запрос на получение отдельной позиции документа по ID.
	// Принимает контекст, ID документа, ID позиции и опционально объект параметров запроса Params.
	// Возвращает найденную позицию.
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params ...*Params) (*RetailSalesReturnPosition, *resty.Response, error)

	// UpdatePosition выполняет запрос на изменение позиции документа.
	// Принимает контекст, ID документа, ID позиции, позицию документа и опционально объект параметров запроса Params.
	// Возвращает изменённую позицию.
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *RetailSalesReturnPosition, params ...*Params) (*RetailSalesReturnPosition, *resty.Response, error)

	// CreatePosition выполняет запрос на добавление позиции документа.
	// Принимает контекст, ID документа, позицию документа и опционально объект параметров запроса Params.
	// Возвращает добавленную позицию.
	CreatePosition(ctx context.Context, id uuid.UUID, position *RetailSalesReturnPosition, params ...*Params) (*RetailSalesReturnPosition, *resty.Response, error)

	// CreatePositionMany выполняет запрос на массовое добавление позиций документа.
	// Принимает контекст, ID документа и множество позиций.
	// Возвращает список добавленных позиций.
	CreatePositionMany(ctx context.Context, id uuid.UUID, positions ...*RetailSalesReturnPosition) (*Slice[RetailSalesReturnPosition], *resty.Response, error)

	// DeletePosition выполняет запрос на удаление позиции документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает «true» в случае успешного удаления позиции.
	DeletePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (bool, *resty.Response, error)

	// DeletePositionMany выполняет запрос на массовое удаление позиций документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeletePositionMany(ctx context.Context, id uuid.UUID, positions ...*RetailSalesReturnPosition) (*DeleteManyResponse, *resty.Response, error)

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
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*RetailSalesReturn, *resty.Response, error)

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

	// Evaluate выполняет запрос на получение шаблона документа с автозаполнением.
	// Принимает контекст, документ и множество значений Evaluate.
	//
	// Возможные значения типа Evaluate:
	//	- EvaluateDiscount – скидки
	//	- EvaluatePrice    – цены
	//	- EvaluateVat      – ндс
	//	- EvaluateCost     – себестоимость
	// Возвращает шаблон документа с автозаполнением.
	Evaluate(ctx context.Context, entity *RetailSalesReturn, evaluate ...Evaluate) (*RetailSalesReturn, *resty.Response, error)
}

const (
	EndpointRetailSalesReturn = EndpointEntity + string(MetaTypeRetailSalesReturn)
)

// NewRetailSalesReturnService принимает [Client] и возвращает сервис для работы с розничными возвратами.
func NewRetailSalesReturnService(client *Client) RetailSalesReturnService {
	return newMainService[RetailSalesReturn, RetailSalesReturnPosition, MetaAttributesStatesSharedWrapper, any](client, EndpointRetailSalesReturn)
}
