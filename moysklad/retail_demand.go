package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"time"
)

// RetailDemand Розничная продажа.
//
// Код сущности: retaildemand
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-roznichnaq-prodazha
type RetailDemand struct {
	AccountID           *uuid.UUID                       `json:"accountId,omitempty"`           // ID учётной записи
	Agent               *Agent                           `json:"agent,omitempty"`               // Метаданные контрагента
	AgentAccount        *AgentAccount                    `json:"agentAccount,omitempty"`        // Метаданные счета контрагента
	Applicable          *bool                            `json:"applicable,omitempty"`          // Отметка о проведении
	CashSum             *float64                         `json:"cashSum,omitempty"`             // Оплачено наличными
	CheckNumber         *string                          `json:"checkNumber,omitempty"`         // Номер чека
	CheckSum            *float64                         `json:"checkSum,omitempty"`            // Сумма Чека
	Code                *string                          `json:"code,omitempty"`                // Код Розничной продажи
	Contract            *NullValue[Contract]             `json:"contract,omitempty"`            // Метаданные договора
	Created             *Timestamp                       `json:"created,omitempty"`             // Дата создания
	CustomerOrder       *NullValue[CustomerOrder]        `json:"customerOrder,omitempty"`       // Метаданные Заказа Покупателя
	Deleted             *Timestamp                       `json:"deleted,omitempty"`             // Момент последнего удаления Розничной продажи
	Description         *string                          `json:"description,omitempty"`         // Комментарий Розничной продажи
	DocumentNumber      *string                          `json:"documentNumber,omitempty"`      // Номер документа
	ExternalCode        *string                          `json:"externalCode,omitempty"`        // Внешний код Розничной продажи
	Files               *MetaArray[File]                 `json:"files,omitempty"`               // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Fiscal              *bool                            `json:"fiscal,omitempty"`              // Отметка о том, был ли использован ФР
	FiscalPrinterInfo   *string                          `json:"fiscalPrinterInfo,omitempty"`   // Информация о фискальном регистраторе
	Group               *Group                           `json:"group,omitempty"`               // Отдел сотрудника
	ID                  *uuid.UUID                       `json:"id,omitempty"`                  // ID Розничной продажи
	Meta                *Meta                            `json:"meta,omitempty"`                // Метаданные Розничной продажи
	Moment              *Timestamp                       `json:"moment,omitempty"`              // Дата документа
	Name                *string                          `json:"name,omitempty"`                // Наименование Розничной продажи
	NoCashSum           *float64                         `json:"noCashSum,omitempty"`           // Оплачено картой
	OFDCode             *string                          `json:"ofdCode,omitempty"`             // Код оператора фискальных данных
	Organization        *Organization                    `json:"organization,omitempty"`        // Метаданные юрлица
	OrganizationAccount *AgentAccount                    `json:"organizationAccount,omitempty"` // Метаданные счета юрлица
	Owner               *Employee                        `json:"owner,omitempty"`               // Метаданные владельца (Сотрудника)
	PayedSum            *float64                         `json:"payedSum,omitempty"`            // Сумма входящих платежей
	Positions           *MetaArray[RetailDemandPosition] `json:"positions,omitempty"`           // Метаданные позиций Розничной продажи
	PrepaymentCashSum   *float64                         `json:"prepaymentCashSum,omitempty"`   // Предоплата наличными
	PrepaymentNoCashSum *float64                         `json:"prepaymentNoCashSum,omitempty"` // Предоплата картой
	PrepaymentQRSum     *float64                         `json:"prepaymentQrSum,omitempty"`     // Предоплата по QR-коду
	Printed             *bool                            `json:"printed,omitempty"`             // Напечатан ли документ
	Project             *NullValue[Project]              `json:"project,omitempty"`             // Метаданные проекта
	Published           *bool                            `json:"published,omitempty"`           // Опубликован ли документ
	QRSum               *float64                         `json:"qrSum,omitempty"`               // Оплачено по QR-коду
	Rate                *NullValue[Rate]                 `json:"rate,omitempty"`                // Валюта
	RetailShift         *RetailShift                     `json:"retailShift,omitempty"`         // Метаданные Розничной смены
	RetailStore         *RetailStore                     `json:"retailStore,omitempty"`         // Метаданные Точки продаж
	SessionNumber       *string                          `json:"sessionNumber,omitempty"`       // Номер сессии
	Shared              *bool                            `json:"shared,omitempty"`              // Общий доступ
	State               *NullValue[State]                `json:"state,omitempty"`               // Метаданные статуса Розничной продажи
	Store               *Store                           `json:"store,omitempty"`               // Метаданные склада
	Sum                 *float64                         `json:"sum,omitempty"`                 // Сумма Розничной продажи в копейках
	SyncID              *uuid.UUID                       `json:"syncId,omitempty"`              // ID синхронизации
	Updated             *Timestamp                       `json:"updated,omitempty"`             // Момент последнего обновления Розничной продажи
	VatEnabled          *bool                            `json:"vatEnabled,omitempty"`          // Учитывается ли НДС
	VatIncluded         *bool                            `json:"vatIncluded,omitempty"`         // Включен ли НДС в цену
	VatSum              *float64                         `json:"vatSum,omitempty"`              // Сумма НДС
	TaxSystem           TaxSystem                        `json:"taxSystem,omitempty"`           // Код системы налогообложения
	Attributes          Slice[Attribute]                 `json:"attributes,omitempty"`          // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (retailDemand RetailDemand) Clean() *RetailDemand {
	if retailDemand.Meta == nil {
		return nil
	}
	return &RetailDemand{Meta: retailDemand.Meta}
}

// AsTaskOperation реализует интерфейс [TaskOperationConverter].
func (retailDemand RetailDemand) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: retailDemand.Meta}
}

// GetAccountID возвращает ID учётной записи.
func (retailDemand RetailDemand) GetAccountID() uuid.UUID {
	return Deref(retailDemand.AccountID)
}

// GetAgent возвращает Метаданные контрагента.
func (retailDemand RetailDemand) GetAgent() Agent {
	return Deref(retailDemand.Agent)
}

// GetAgentAccount возвращает Метаданные счета контрагента.
func (retailDemand RetailDemand) GetAgentAccount() AgentAccount {
	return Deref(retailDemand.AgentAccount)
}

// GetApplicable возвращает Отметку о проведении.
func (retailDemand RetailDemand) GetApplicable() bool {
	return Deref(retailDemand.Applicable)
}

// GetCashSum возвращает Оплачено наличными.
func (retailDemand RetailDemand) GetCashSum() float64 {
	return Deref(retailDemand.CashSum)
}

// GetCheckNumber возвращает Номер чека.
func (retailDemand RetailDemand) GetCheckNumber() string {
	return Deref(retailDemand.CheckNumber)
}

// GetCheckSum возвращает Сумму Чека.
func (retailDemand RetailDemand) GetCheckSum() float64 {
	return Deref(retailDemand.CheckSum)
}

// GetCode возвращает Код Розничной продажи.
func (retailDemand RetailDemand) GetCode() string {
	return Deref(retailDemand.Code)
}

// GetContract возвращает Метаданные договора.
func (retailDemand RetailDemand) GetContract() Contract {
	return retailDemand.Contract.getValue()
}

// GetCreated возвращает Дату создания.
func (retailDemand RetailDemand) GetCreated() time.Time {
	return Deref(retailDemand.Created).Time()
}

// GetCustomerOrder возвращает Метаданные Заказа Покупателя.
func (retailDemand RetailDemand) GetCustomerOrder() CustomerOrder {
	return Deref(retailDemand.CustomerOrder).getValue()
}

// GetDeleted возвращает Момент последнего удаления Розничной продажи.
func (retailDemand RetailDemand) GetDeleted() time.Time {
	return Deref(retailDemand.Deleted).Time()
}

// GetDescription возвращает Комментарий Розничной продажи.
func (retailDemand RetailDemand) GetDescription() string {
	return Deref(retailDemand.Description)
}

// GetDocumentNumber возвращает Номер документа.
func (retailDemand RetailDemand) GetDocumentNumber() string {
	return Deref(retailDemand.DocumentNumber)
}

// GetExternalCode возвращает Внешний код Розничной продажи.
func (retailDemand RetailDemand) GetExternalCode() string {
	return Deref(retailDemand.ExternalCode)
}

// GetFiles возвращает Метаданные массива Файлов.
func (retailDemand RetailDemand) GetFiles() MetaArray[File] {
	return Deref(retailDemand.Files)
}

// GetFiscal возвращает Отметку о том, был ли использован ФР.
func (retailDemand RetailDemand) GetFiscal() bool {
	return Deref(retailDemand.Fiscal)
}

// GetFiscalPrinterInfo возвращает Информацию о фискальном регистраторе.
func (retailDemand RetailDemand) GetFiscalPrinterInfo() string {
	return Deref(retailDemand.FiscalPrinterInfo)
}

// GetGroup возвращает Отдел сотрудника.
func (retailDemand RetailDemand) GetGroup() Group {
	return Deref(retailDemand.Group)
}

// GetID возвращает ID Розничной продажи.
func (retailDemand RetailDemand) GetID() uuid.UUID {
	return Deref(retailDemand.ID)
}

// GetMeta возвращает Метаданные Розничной продажи.
func (retailDemand RetailDemand) GetMeta() Meta {
	return Deref(retailDemand.Meta)
}

// GetMoment возвращает Дату документа.
func (retailDemand RetailDemand) GetMoment() time.Time {
	return Deref(retailDemand.Moment).Time()
}

// GetName возвращает Наименование Розничной продажи.
func (retailDemand RetailDemand) GetName() string {
	return Deref(retailDemand.Name)
}

// GetNoCashSum возвращает Оплачено картой.
func (retailDemand RetailDemand) GetNoCashSum() float64 {
	return Deref(retailDemand.NoCashSum)
}

// GetOFDCode возвращает Код оператора фискальных данных.
func (retailDemand RetailDemand) GetOFDCode() string {
	return Deref(retailDemand.OFDCode)
}

// GetOrganization возвращает Метаданные юрлица.
func (retailDemand RetailDemand) GetOrganization() Organization {
	return Deref(retailDemand.Organization)
}

// GetOrganizationAccount возвращает Метаданные счета юрлица.
func (retailDemand RetailDemand) GetOrganizationAccount() AgentAccount {
	return Deref(retailDemand.OrganizationAccount)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (retailDemand RetailDemand) GetOwner() Employee {
	return Deref(retailDemand.Owner)
}

// GetPayedSum возвращает Сумму входящих платежей.
func (retailDemand RetailDemand) GetPayedSum() float64 {
	return Deref(retailDemand.PayedSum)
}

// GetPositions возвращает Метаданные позиций Розничной продажи.
func (retailDemand RetailDemand) GetPositions() MetaArray[RetailDemandPosition] {
	return Deref(retailDemand.Positions)
}

// GetPrepaymentCashSum возвращает Предоплату наличными.
func (retailDemand RetailDemand) GetPrepaymentCashSum() float64 {
	return Deref(retailDemand.PrepaymentCashSum)
}

// GetPrepaymentNoCashSum возвращает Предоплату картой.
func (retailDemand RetailDemand) GetPrepaymentNoCashSum() float64 {
	return Deref(retailDemand.PrepaymentNoCashSum)
}

// GetPrepaymentQRSum возвращает Предоплату по QR-коду.
func (retailDemand RetailDemand) GetPrepaymentQRSum() float64 {
	return Deref(retailDemand.PrepaymentQRSum)
}

// GetPrinted возвращает true, если документ напечатан.
func (retailDemand RetailDemand) GetPrinted() bool {
	return Deref(retailDemand.Printed)
}

// GetProject возвращает Метаданные проекта.
func (retailDemand RetailDemand) GetProject() Project {
	return retailDemand.Project.getValue()
}

// GetPublished возвращает true, если документ опубликован.
func (retailDemand RetailDemand) GetPublished() bool {
	return Deref(retailDemand.Published)
}

// GetQRSum возвращает оплачено по QR-коду.
func (retailDemand RetailDemand) GetQRSum() float64 {
	return Deref(retailDemand.QRSum)
}

// GetRate возвращает Валюту.
func (retailDemand RetailDemand) GetRate() Rate {
	return retailDemand.Rate.getValue()
}

// GetRetailShift возвращает Метаданные Розничной смены.
func (retailDemand RetailDemand) GetRetailShift() RetailShift {
	return Deref(retailDemand.RetailShift)
}

// GetRetailStore возвращает Метаданные Точки продаж.
func (retailDemand RetailDemand) GetRetailStore() RetailStore {
	return Deref(retailDemand.RetailStore)
}

// GetSessionNumber возвращает Номер сессии.
func (retailDemand RetailDemand) GetSessionNumber() string {
	return Deref(retailDemand.SessionNumber)
}

// GetShared возвращает флаг Общего доступа.
func (retailDemand RetailDemand) GetShared() bool {
	return Deref(retailDemand.Shared)
}

// GetState возвращает Метаданные статуса Розничной продажи.
func (retailDemand RetailDemand) GetState() State {
	return Deref(retailDemand.State).getValue()
}

// GetStore возвращает Метаданные склада.
func (retailDemand RetailDemand) GetStore() Store {
	return Deref(retailDemand.Store)
}

// GetSum возвращает Сумму Розничной продажи в копейках.
func (retailDemand RetailDemand) GetSum() float64 {
	return Deref(retailDemand.Sum)
}

// GetSyncID возвращает ID синхронизации.
func (retailDemand RetailDemand) GetSyncID() uuid.UUID {
	return Deref(retailDemand.SyncID)
}

// GetUpdated возвращает Момент последнего обновления Розничной продажи.
func (retailDemand RetailDemand) GetUpdated() time.Time {
	return Deref(retailDemand.Updated).Time()
}

// GetVatEnabled возвращает true, если учитывается НДС.
func (retailDemand RetailDemand) GetVatEnabled() bool {
	return Deref(retailDemand.VatEnabled)
}

// GetVatIncluded возвращает true, если НДС включен в цену.
func (retailDemand RetailDemand) GetVatIncluded() bool {
	return Deref(retailDemand.VatIncluded)
}

// GetVatSum возвращает Сумму НДС.
func (retailDemand RetailDemand) GetVatSum() float64 {
	return Deref(retailDemand.VatSum)
}

// GetTaxSystem возвращает Код системы налогообложения.
func (retailDemand RetailDemand) GetTaxSystem() TaxSystem {
	return retailDemand.TaxSystem
}

// GetAttributes возвращает Список метаданных доп. полей.
func (retailDemand RetailDemand) GetAttributes() Slice[Attribute] {
	return retailDemand.Attributes
}

// SetAgent устанавливает Метаданные Контрагента.
//
// Принимает [Counterparty] или [Organization].
func (retailDemand *RetailDemand) SetAgent(agent AgentOrganizationConverter) *RetailDemand {
	if agent != nil {
		retailDemand.Agent = agent.AsOrganizationAgent()
	}
	return retailDemand
}

// SetAgentAccount устанавливает Метаданные счета контрагента.
func (retailDemand *RetailDemand) SetAgentAccount(agentAccount *AgentAccount) *RetailDemand {
	if agentAccount != nil {
		retailDemand.AgentAccount = agentAccount.Clean()
	}
	return retailDemand
}

// SetApplicable устанавливает Отметку о проведении.
func (retailDemand *RetailDemand) SetApplicable(applicable bool) *RetailDemand {
	retailDemand.Applicable = &applicable
	return retailDemand
}

// SetCashSum устанавливает Оплачено наличными.
func (retailDemand *RetailDemand) SetCashSum(cashSum float64) *RetailDemand {
	retailDemand.CashSum = &cashSum
	return retailDemand
}

// SetCheckNumber устанавливает Номер чека.
func (retailDemand *RetailDemand) SetCheckNumber(checkNumber string) *RetailDemand {
	retailDemand.CheckNumber = &checkNumber
	return retailDemand
}

// SetCheckSum устанавливает Сумму чека.
func (retailDemand *RetailDemand) SetCheckSum(checkSum float64) *RetailDemand {
	retailDemand.CheckSum = &checkSum
	return retailDemand
}

// SetCode устанавливает Код Розничной продажи.
func (retailDemand *RetailDemand) SetCode(code string) *RetailDemand {
	retailDemand.Code = &code
	return retailDemand
}

// SetContract устанавливает Метаданные договора.
//
// Передача nil передаёт сброс значения (null).
func (retailDemand *RetailDemand) SetContract(contract *Contract) *RetailDemand {
	retailDemand.Contract = NewNullValue(contract)
	return retailDemand
}

// SetCustomerOrder устанавливает Метаданные Заказа Покупателя.
//
// Передача nil передаёт сброс значения (null).
func (retailDemand *RetailDemand) SetCustomerOrder(customerOrder *CustomerOrder) *RetailDemand {
	retailDemand.CustomerOrder = NewNullValue(customerOrder)
	return retailDemand
}

// SetDescription устанавливает Комментарий Розничной продажи.
func (retailDemand *RetailDemand) SetDescription(description string) *RetailDemand {
	retailDemand.Description = &description
	return retailDemand
}

// SetDocumentNumber устанавливает Номер документа.
func (retailDemand *RetailDemand) SetDocumentNumber(documentNumber string) *RetailDemand {
	retailDemand.DocumentNumber = &documentNumber
	return retailDemand
}

// SetExternalCode устанавливает Внешний код Розничной продажи.
func (retailDemand *RetailDemand) SetExternalCode(externalCode string) *RetailDemand {
	retailDemand.ExternalCode = &externalCode
	return retailDemand
}

// SetFiles устанавливает Метаданные массива Файлов.
//
// Принимает множество объектов [File].
func (retailDemand *RetailDemand) SetFiles(files ...*File) *RetailDemand {
	retailDemand.Files = NewMetaArrayFrom(files)
	return retailDemand
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (retailDemand *RetailDemand) SetGroup(group *Group) *RetailDemand {
	if group != nil {
		retailDemand.Group = group.Clean()
	}
	return retailDemand
}

// SetMeta устанавливает Метаданные Розничной продажи.
func (retailDemand *RetailDemand) SetMeta(meta *Meta) *RetailDemand {
	retailDemand.Meta = meta
	return retailDemand
}

// SetMoment устанавливает Дату документа.
func (retailDemand *RetailDemand) SetMoment(moment time.Time) *RetailDemand {
	retailDemand.Moment = NewTimestamp(moment)
	return retailDemand
}

// SetName устанавливает Наименование Розничной продажи.
func (retailDemand *RetailDemand) SetName(name string) *RetailDemand {
	retailDemand.Name = &name
	return retailDemand
}

// SetNoCashSum устанавливает Оплачено картой.
func (retailDemand *RetailDemand) SetNoCashSum(noCashSum float64) *RetailDemand {
	retailDemand.NoCashSum = &noCashSum
	return retailDemand
}

// SetOrganization устанавливает Метаданные юрлица.
func (retailDemand *RetailDemand) SetOrganization(organization *Organization) *RetailDemand {
	if organization != nil {
		retailDemand.Organization = organization.Clean()
	}
	return retailDemand
}

// SetOrganizationAccount устанавливает Метаданные счета юрлица.
func (retailDemand *RetailDemand) SetOrganizationAccount(organizationAccount *AgentAccount) *RetailDemand {
	if organizationAccount != nil {
		retailDemand.OrganizationAccount = organizationAccount.Clean()
	}
	return retailDemand
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (retailDemand *RetailDemand) SetOwner(owner *Employee) *RetailDemand {
	if owner != nil {
		retailDemand.Owner = owner.Clean()
	}
	return retailDemand
}

// SetPositions устанавливает Метаданные позиций Розничной продажи.
//
// Принимает множество объектов [RetailDemandPosition].
func (retailDemand *RetailDemand) SetPositions(positions ...*RetailDemandPosition) *RetailDemand {
	retailDemand.Positions = NewMetaArrayFrom(positions)
	return retailDemand
}

// SetPrepaymentCashSum устанавливает Предоплату наличными.
func (retailDemand *RetailDemand) SetPrepaymentCashSum(prepaymentCashSum float64) *RetailDemand {
	retailDemand.PrepaymentCashSum = &prepaymentCashSum
	return retailDemand
}

// SetPrepaymentNoCashSum устанавливает Предоплату картой.
func (retailDemand *RetailDemand) SetPrepaymentNoCashSum(prepaymentNoCashSum float64) *RetailDemand {
	retailDemand.PrepaymentNoCashSum = &prepaymentNoCashSum
	return retailDemand
}

// SetPrepaymentQRSum устанавливает Предоплату по QR-коду.
func (retailDemand *RetailDemand) SetPrepaymentQRSum(prepaymentQRSum float64) *RetailDemand {
	retailDemand.PrepaymentQRSum = &prepaymentQRSum
	return retailDemand
}

// SetProject устанавливает Метаданные проекта.
//
// Передача nil передаёт сброс значения (null).
func (retailDemand *RetailDemand) SetProject(project *Project) *RetailDemand {
	retailDemand.Project = NewNullValue(project)
	return retailDemand
}

// SetQRSum устанавливает Оплачено по QR-коду.
func (retailDemand *RetailDemand) SetQRSum(qrSum float64) *RetailDemand {
	retailDemand.QRSum = &qrSum
	return retailDemand
}

// SetRate устанавливает Валюту.
//
// Передача nil передаёт сброс значения (null).
func (retailDemand *RetailDemand) SetRate(rate *Rate) *RetailDemand {
	retailDemand.Rate = NewNullValue(rate)
	return retailDemand
}

// SetRetailShift устанавливает Метаданные Розничной смены.
func (retailDemand *RetailDemand) SetRetailShift(retailShift *RetailShift) *RetailDemand {
	if retailShift != nil {
		retailDemand.RetailShift = retailShift.Clean()
	}
	return retailDemand
}

// SetRetailStore устанавливает Метаданные Точки продаж.
func (retailDemand *RetailDemand) SetRetailStore(retailStore *RetailStore) *RetailDemand {
	if retailStore != nil {
		retailDemand.RetailStore = retailStore.Clean()
	}
	return retailDemand
}

// SetSessionNumber устанавливает Номер сессии.
func (retailDemand *RetailDemand) SetSessionNumber(sessionNumber string) *RetailDemand {
	retailDemand.SessionNumber = &sessionNumber
	return retailDemand
}

// SetShared устанавливает флаг общего доступа.
func (retailDemand *RetailDemand) SetShared(shared bool) *RetailDemand {
	retailDemand.Shared = &shared
	return retailDemand
}

// SetState устанавливает Метаданные статуса Розничной продажи.
//
// Передача nil передаёт сброс значения (null).
func (retailDemand *RetailDemand) SetState(state *State) *RetailDemand {
	retailDemand.State = NewNullValue(state)
	return retailDemand
}

// SetStore устанавливает Метаданные склада.
func (retailDemand *RetailDemand) SetStore(store *Store) *RetailDemand {
	if store != nil {
		retailDemand.Store = store.Clean()
	}
	return retailDemand
}

// SetSyncID устанавливает ID синхронизации.
func (retailDemand *RetailDemand) SetSyncID(syncID uuid.UUID) *RetailDemand {
	retailDemand.SyncID = &syncID
	return retailDemand
}

// SetVatEnabled устанавливает значение, учитывающее НДС для Розничной продажи.
func (retailDemand *RetailDemand) SetVatEnabled(vatEnabled bool) *RetailDemand {
	retailDemand.VatEnabled = &vatEnabled
	return retailDemand
}

// SetVatIncluded устанавливает флаг включения НДС в цену.
func (retailDemand *RetailDemand) SetVatIncluded(vatIncluded bool) *RetailDemand {
	retailDemand.VatIncluded = &vatIncluded
	return retailDemand
}

// SetTaxSystem устанавливает Код системы налогообложения.
func (retailDemand *RetailDemand) SetTaxSystem(taxSystem TaxSystem) *RetailDemand {
	retailDemand.TaxSystem = taxSystem
	return retailDemand
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (retailDemand *RetailDemand) SetAttributes(attributes ...*Attribute) *RetailDemand {
	retailDemand.Attributes.Push(attributes...)
	return retailDemand
}

// String реализует интерфейс [fmt.Stringer].
func (retailDemand RetailDemand) String() string {
	return Stringify(retailDemand)
}

// MetaType возвращает код сущности.
func (RetailDemand) MetaType() MetaType {
	return MetaTypeRetailDemand
}

// Update shortcut
func (retailDemand RetailDemand) Update(ctx context.Context, client *Client, params ...*Params) (*RetailDemand, *resty.Response, error) {
	return NewRetailDemandService(client).Update(ctx, retailDemand.GetID(), &retailDemand, params...)
}

// Create shortcut
func (retailDemand RetailDemand) Create(ctx context.Context, client *Client, params ...*Params) (*RetailDemand, *resty.Response, error) {
	return NewRetailDemandService(client).Create(ctx, &retailDemand, params...)
}

// Delete shortcut
func (retailDemand RetailDemand) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewRetailDemandService(client).Delete(ctx, retailDemand.GetID())
}

// RetailDemandPosition позиция розничной продажи.
//
// Код сущности: demandposition
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-roznichnaq-prodazha-roznichnye-prodazhi-pozicii-roznichnoj-prodazhi
type RetailDemandPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учётной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	Cost       *float64            `json:"cost,omitempty"`       // Себестоимость (только для услуг)
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
func (retailDemandPosition RetailDemandPosition) GetAccountID() uuid.UUID {
	return Deref(retailDemandPosition.AccountID)
}

// GetAssortment возвращает Метаданные товара/услуги/серии/модификации, которую представляет собой позиция.
func (retailDemandPosition RetailDemandPosition) GetAssortment() AssortmentPosition {
	return Deref(retailDemandPosition.Assortment)
}

// GetCost возвращает Себестоимость (только для услуг).
func (retailDemandPosition RetailDemandPosition) GetCost() float64 {
	return Deref(retailDemandPosition.Cost)
}

// GetDiscount возвращает Процент скидки или наценки.
//
// Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%.
func (retailDemandPosition RetailDemandPosition) GetDiscount() float64 {
	return Deref(retailDemandPosition.Discount)
}

// GetID возвращает ID позиции.
func (retailDemandPosition RetailDemandPosition) GetID() uuid.UUID {
	return Deref(retailDemandPosition.ID)
}

// GetPack возвращает Упаковку Товара.
func (retailDemandPosition RetailDemandPosition) GetPack() Pack {
	return Deref(retailDemandPosition.Pack)
}

// GetPrice возвращает Цену товара/услуги в копейках.
func (retailDemandPosition RetailDemandPosition) GetPrice() float64 {
	return Deref(retailDemandPosition.Price)
}

// GetQuantity возвращает Количество товаров/услуг данного вида в позиции.
//
// Если позиция - товар, у которого включен учет по серийным номерам,
// то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
func (retailDemandPosition RetailDemandPosition) GetQuantity() float64 {
	return Deref(retailDemandPosition.Quantity)
}

// GetThings возвращает Серийные номера.
//
// Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете.
// В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута.
func (retailDemandPosition RetailDemandPosition) GetThings() Slice[string] {
	return retailDemandPosition.Things
}

// GetVat возвращает НДС, которым облагается текущая позиция.
func (retailDemandPosition RetailDemandPosition) GetVat() int {
	return Deref(retailDemandPosition.Vat)
}

// GetVatEnabled возвращает true, если учитывается НДС.
func (retailDemandPosition RetailDemandPosition) GetVatEnabled() bool {
	return Deref(retailDemandPosition.VatEnabled)
}

// GetStock возвращает Остатки и себестоимость позиции (указывается при наличии параметра запроса `fields=stock`).
func (retailDemandPosition RetailDemandPosition) GetStock() Stock {
	return Deref(retailDemandPosition.Stock)
}

// SetAssortment устанавливает Метаданные товара/услуги, которую представляет собой компонент.
//
// Принимает объект, реализующий интерфейс [AssortmentConverter].
func (retailDemandPosition *RetailDemandPosition) SetAssortment(assortment AssortmentConverter) *RetailDemandPosition {
	if assortment != nil {
		retailDemandPosition.Assortment = assortment.AsAssortment()
	}
	return retailDemandPosition
}

// SetCost устанавливает Себестоимость (только для услуг).
func (retailDemandPosition *RetailDemandPosition) SetCost(cost float64) *RetailDemandPosition {
	retailDemandPosition.Cost = &cost
	return retailDemandPosition
}

// SetDiscount устанавливает Процент скидки или наценки.
//
// Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%.
func (retailDemandPosition *RetailDemandPosition) SetDiscount(discount float64) *RetailDemandPosition {
	retailDemandPosition.Discount = &discount
	return retailDemandPosition
}

// SetPack устанавливает Упаковку Товара.
func (retailDemandPosition *RetailDemandPosition) SetPack(pack *Pack) *RetailDemandPosition {
	retailDemandPosition.Pack = pack
	return retailDemandPosition
}

// SetPrice устанавливает Цену товара/услуги в копейках.
func (retailDemandPosition *RetailDemandPosition) SetPrice(price float64) *RetailDemandPosition {
	retailDemandPosition.Price = &price
	return retailDemandPosition
}

// SetQuantity устанавливает Количество товаров данного вида в позиции.
//
// Если позиция - товар, у которого включен учет по серийным номерам,
// то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
func (retailDemandPosition *RetailDemandPosition) SetQuantity(quantity float64) *RetailDemandPosition {
	retailDemandPosition.Quantity = &quantity
	return retailDemandPosition
}

// SetVat устанавливает НДС, которым облагается текущая позиция.
func (retailDemandPosition *RetailDemandPosition) SetVat(vat int) *RetailDemandPosition {
	retailDemandPosition.Vat = &vat
	return retailDemandPosition
}

// SetVatEnabled устанавливает значение, учитывающее НДС для текущей позиции.
func (retailDemandPosition *RetailDemandPosition) SetVatEnabled(vatEnabled bool) *RetailDemandPosition {
	retailDemandPosition.VatEnabled = &vatEnabled
	return retailDemandPosition
}

// SetThings устанавливает Серийные номера.
//
// Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете.
// В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута.
//
// Принимает множество string.
func (retailDemandPosition *RetailDemandPosition) SetThings(things ...string) *RetailDemandPosition {
	retailDemandPosition.Things = NewSliceFrom(things)
	return retailDemandPosition
}

// String реализует интерфейс [fmt.Stringer].
func (retailDemandPosition RetailDemandPosition) String() string {
	return Stringify(retailDemandPosition)
}

// MetaType возвращает код сущности.
func (RetailDemandPosition) MetaType() MetaType {
	return MetaTypeRetailDemandPosition
}

// RetailDemandService описывает методы сервиса для работы с розничными продажами.
type RetailDemandService interface {
	// GetList выполняет запрос на получение списка розничных продаж.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[RetailDemand], *resty.Response, error)

	// Create выполняет запрос на создание розничной продажи.
	// Обязательные поля для заполнения:
	//	- retailShift (Ссылка на Розничную смену, в рамках которой происходит продажа)
	// Принимает контекст, розничную продажу и опционально объект параметров запроса Params.
	// Возвращает созданный розничную продажу.
	Create(ctx context.Context, retailDemand *RetailDemand, params ...*Params) (*RetailDemand, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение розничных продаж.
	// Изменяемые розничные продажи должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список розничных продаж и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых розничных продаж.
	CreateUpdateMany(ctx context.Context, retailDemandList Slice[RetailDemand], params ...*Params) (*Slice[RetailDemand], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление розничных продаж.
	// Принимает контекст и множество розничных продаж.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*RetailDemand) (*DeleteManyResponse, *resty.Response, error)

	// Delete выполняет запрос на удаление розничной продажи.
	// Принимает контекст и ID розничной продажи.
	// Возвращает «true» в случае успешного удаления розничной продажи.
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельной розничной продажи по ID.
	// Принимает контекст, ID розничной продажи и опционально объект параметров запроса Params.
	// Возвращает найденную розничную продажу.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*RetailDemand, *resty.Response, error)

	// Update выполняет запрос на изменение розничной продажи.
	// Принимает контекст, розничную продажу и опционально объект параметров запроса Params.
	// Возвращает изменённую розничную продажу.
	Update(ctx context.Context, id uuid.UUID, retailDemand *RetailDemand, params ...*Params) (*RetailDemand, *resty.Response, error)

	// Template выполняет запрос на получение предзаполненной розничной продажи со стандартными полями без связи с какими-либо другими документами.
	// Принимает контекст.
	// Возвращает предзаполненную розничную продажу.
	Template(ctx context.Context) (*RetailDemand, *resty.Response, error)

	// TemplateBased выполняет запрос на получение шаблона розничной продажи на основе других документов.
	// Основание, на котором может быть создана:
	//	- Розничная смена (RetailShift)
	//	- Заказ покупателя (CustomerOrder)
	// Принимает контекст и множество документов из списка выше.
	// Возвращает предзаполненную розничную продажу на основании переданных документов.
	TemplateBased(ctx context.Context, basedOn ...MetaOwner) (*RetailDemand, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных розничных продаж.
	// Принимает контекст.
	// Возвращает объект метаданных MetaAttributesStatesSharedWrapper.
	GetMetadata(ctx context.Context) (*MetaAttributesStatesSharedWrapper, *resty.Response, error)

	// GetPositionList выполняет запрос на получение списка позиций документа.
	// Принимает контекст, ID документа и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetPositionList(ctx context.Context, id uuid.UUID, params ...*Params) (*List[RetailDemandPosition], *resty.Response, error)

	// GetPositionByID выполняет запрос на получение отдельной позиции документа по ID.
	// Принимает контекст, ID документа, ID позиции и опционально объект параметров запроса Params.
	// Возвращает найденную позицию.
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params ...*Params) (*RetailDemandPosition, *resty.Response, error)

	// UpdatePosition выполняет запрос на изменение позиции документа.
	// Принимает контекст, ID документа, ID позиции, позицию документа и опционально объект параметров запроса Params.
	// Возвращает изменённую позицию.
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *RetailDemandPosition, params ...*Params) (*RetailDemandPosition, *resty.Response, error)

	// CreatePosition выполняет запрос на добавление позиции документа.
	// Принимает контекст, ID документа, позицию документа и опционально объект параметров запроса Params.
	// Возвращает добавленную позицию.
	CreatePosition(ctx context.Context, id uuid.UUID, position *RetailDemandPosition, params ...*Params) (*RetailDemandPosition, *resty.Response, error)

	// CreatePositionMany выполняет запрос на массовое добавление позиций документа.
	// Принимает контекст, ID документа и множество позиций.
	// Возвращает список добавленных позиций.
	CreatePositionMany(ctx context.Context, id uuid.UUID, positions ...*RetailDemandPosition) (*Slice[RetailDemandPosition], *resty.Response, error)

	// DeletePosition выполняет запрос на удаление позиции документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает «true» в случае успешного удаления позиции.
	DeletePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (bool, *resty.Response, error)

	// DeletePositionMany выполняет запрос на массовое удаление позиций документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeletePositionMany(ctx context.Context, id uuid.UUID, positions ...*RetailDemandPosition) (*DeleteManyResponse, *resty.Response, error)

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
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*RetailDemand, *resty.Response, error)

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
	Evaluate(ctx context.Context, entity *RetailDemand, evaluate ...Evaluate) (*RetailDemand, *resty.Response, error)
}

const (
	EndpointRetailDemand = EndpointEntity + string(MetaTypeRetailDemand)
)

// NewRetailDemandService принимает [Client] и возвращает сервис для работы с розничными продажами.
func NewRetailDemandService(client *Client) RetailDemandService {
	return newMainService[RetailDemand, RetailDemandPosition, MetaAttributesStatesSharedWrapper, any](client, EndpointRetailDemand)
}
