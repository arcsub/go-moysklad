package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"time"
)

// FactureOut Счет-фактура выданный.
//
// Код сущности: factureout
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-schet-faktura-wydannyj
type FactureOut struct {
	Organization    *Organization         `json:"organization,omitempty"`    // Метаданные юрлица
	Deleted         *Timestamp            `json:"deleted,omitempty"`         // Момент последнего удаления выданного Счета-фактуры
	Applicable      *bool                 `json:"applicable,omitempty"`      // Отметка о проведении
	AccountID       *uuid.UUID            `json:"accountId,omitempty"`       // ID учётной записи
	Code            *string               `json:"code,omitempty"`            // Код выданного Счета-фактуры
	Contract        *NullValue[Contract]  `json:"contract,omitempty"`        // Метаданные договора
	Created         *Timestamp            `json:"created,omitempty"`         // Дата создания
	Owner           *Employee             `json:"owner,omitempty"`           // Метаданные владельца (Сотрудника)
	Description     *string               `json:"description,omitempty"`     // Комментарий выданного Счета-фактуры
	ExternalCode    *string               `json:"externalCode,omitempty"`    // Внешний код выданного Счета-фактуры
	Files           *MetaArray[File]      `json:"files,omitempty"`           // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group           *Group                `json:"group,omitempty"`           // Отдел сотрудника
	ID              *uuid.UUID            `json:"id,omitempty"`              // ID выданного Счета-фактуры
	Printed         *bool                 `json:"printed,omitempty"`         // Напечатан ли документ
	Moment          *Timestamp            `json:"moment,omitempty"`          // Дата документа
	Name            *string               `json:"name,omitempty"`            // Наименование выданного Счета-фактуры
	PaymentDate     *Timestamp            `json:"paymentDate,omitempty"`     // Дата платежного документа
	Agent           *Agent                `json:"agent,omitempty"`           // Метаданные контрагента
	Meta            *Meta                 `json:"meta,omitempty"`            // Метаданные выданного Счета-фактуры
	Published       *bool                 `json:"published,omitempty"`       // Опубликован ли документ
	Rate            *NullValue[Rate]      `json:"rate,omitempty"`            // Валюта
	Shared          *bool                 `json:"shared,omitempty"`          // Общий доступ
	State           *NullValue[State]     `json:"state,omitempty"`           // Метаданные статуса выданного Счета-фактуры
	StateContractID *string               `json:"stateContractId,omitempty"` // Идентификатор государственного контракта, договора (соглашения)
	Sum             *float64              `json:"sum,omitempty"`             // Сумма выданного Счета-фактуры в копейках
	SyncID          *uuid.UUID            `json:"syncId,omitempty"`          // ID синхронизации
	Updated         *Timestamp            `json:"updated,omitempty"`         // Момент последнего обновления выданного Счета-фактуры
	Demands         Slice[Demand]         `json:"demands,omitempty"`         // Массив ссылок на связанные отгрузки
	Payments        Slice[Payment]        `json:"payments,omitempty"`        // Массив ссылок на связанные входящие платежи
	Returns         Slice[PurchaseReturn] `json:"returns,omitempty"`         // Массив ссылок на связанные возвраты поставщикам
	Consignee       *Agent                `json:"consignee,omitempty"`       // Метаданные грузополучателя (контрагент или юрлицо)
	PaymentNumber   *string               `json:"paymentNumber,omitempty"`   // Название платежного документа
	Attributes      Slice[Attribute]      `json:"attributes,omitempty"`      // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (factureOut FactureOut) Clean() *FactureOut {
	if factureOut.Meta == nil {
		return nil
	}
	return &FactureOut{Meta: factureOut.Meta}
}

// AsOperation реализует интерфейс [OperationConverter].
func (factureOut FactureOut) AsOperation() *Operation {
	return newOperation(factureOut)
}

// AsTaskOperation реализует интерфейс [TaskOperationConverter].
func (factureOut FactureOut) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: factureOut.Meta}
}

// GetOrganization возвращает Метаданные юрлица.
func (factureOut FactureOut) GetOrganization() Organization {
	return Deref(factureOut.Organization)
}

// GetDeleted возвращает Момент последнего удаления выданного Счета-фактуры.
func (factureOut FactureOut) GetDeleted() Timestamp {
	return Deref(factureOut.Deleted)
}

// GetApplicable возвращает Отметку о проведении.
func (factureOut FactureOut) GetApplicable() bool {
	return Deref(factureOut.Applicable)
}

// GetAccountID возвращает ID учётной записи.
func (factureOut FactureOut) GetAccountID() uuid.UUID {
	return Deref(factureOut.AccountID)
}

// GetCode возвращает Код выданного Счета-фактуры.
func (factureOut FactureOut) GetCode() string {
	return Deref(factureOut.Code)
}

// GetContract возвращает Метаданные договора.
func (factureOut FactureOut) GetContract() Contract {
	return Deref(factureOut.Contract).getValue()
}

// GetCreated возвращает Дату создания.
func (factureOut FactureOut) GetCreated() Timestamp {
	return Deref(factureOut.Created)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (factureOut FactureOut) GetOwner() Employee {
	return Deref(factureOut.Owner)
}

// GetDescription возвращаетКомментарий выданного Счета-фактуры.
func (factureOut FactureOut) GetDescription() string {
	return Deref(factureOut.Description)
}

// GetExternalCode возвращает Внешний код выданного Счета-фактуры.
func (factureOut FactureOut) GetExternalCode() string {
	return Deref(factureOut.ExternalCode)
}

// GetFiles возвращает Метаданные массива Файлов.
func (factureOut FactureOut) GetFiles() MetaArray[File] {
	return Deref(factureOut.Files)
}

// GetGroup возвращает Отдел сотрудника.
func (factureOut FactureOut) GetGroup() Group {
	return Deref(factureOut.Group)
}

// GetID возвращает ID выданного Счета-фактуры.
func (factureOut FactureOut) GetID() uuid.UUID {
	return Deref(factureOut.ID)
}

// GetPrinted возвращает true, если документ напечатан.
func (factureOut FactureOut) GetPrinted() bool {
	return Deref(factureOut.Printed)
}

// GetMoment возвращает Момент документа.
func (factureOut FactureOut) GetMoment() Timestamp {
	return Deref(factureOut.Moment)
}

// GetName возвращает Наименование выданного Счета-фактуры.
func (factureOut FactureOut) GetName() string {
	return Deref(factureOut.Name)
}

// GetPaymentDate возвращает Дату платежного документа.
func (factureOut FactureOut) GetPaymentDate() Timestamp {
	return Deref(factureOut.PaymentDate)
}

// GetAgent возвращает Метаданные Контрагента.
func (factureOut FactureOut) GetAgent() Agent {
	return Deref(factureOut.Agent)
}

// GetMeta возвращает Метаданные выданного Счета-фактуры.
func (factureOut FactureOut) GetMeta() Meta {
	return Deref(factureOut.Meta)
}

// GetPublished возвращает true, если документ опубликован.
func (factureOut FactureOut) GetPublished() bool {
	return Deref(factureOut.Published)
}

// GetRate возвращает Валюту.
func (factureOut FactureOut) GetRate() Rate {
	return Deref(factureOut.Rate).getValue()
}

// GetShared возвращает флаг Общего доступа.
func (factureOut FactureOut) GetShared() bool {
	return Deref(factureOut.Shared)
}

// GetState возвращает Метаданные статуса выданного Счета-фактуры.
func (factureOut FactureOut) GetState() State {
	return Deref(factureOut.State).getValue()
}

// GetStateContractID возвращает Идентификатор государственного контракта, договора (соглашения).
func (factureOut FactureOut) GetStateContractID() string {
	return Deref(factureOut.StateContractID)
}

// GetSum возвращает Сумму выданного Счета-фактуры в копейках.
func (factureOut FactureOut) GetSum() float64 {
	return Deref(factureOut.Sum)
}

// GetSyncID возвращает ID синхронизации.
func (factureOut FactureOut) GetSyncID() uuid.UUID {
	return Deref(factureOut.SyncID)
}

// GetUpdated возвращает Момент последнего обновления выданного Счета-фактуры.
func (factureOut FactureOut) GetUpdated() Timestamp {
	return Deref(factureOut.Updated)
}

// GetDemands возвращает Массив ссылок на связанные отгрузки.
func (factureOut FactureOut) GetDemands() Slice[Demand] {
	return factureOut.Demands
}

// GetPayments возвращает Массив ссылок на связанные входящие платежи.
func (factureOut FactureOut) GetPayments() Slice[Payment] {
	return factureOut.Payments
}

// GetReturns возвращает Массив ссылок на связанные возвраты поставщикам.
func (factureOut FactureOut) GetReturns() Slice[PurchaseReturn] {
	return factureOut.Returns
}

// GetConsignee возвращает Метаданные грузополучателя (контрагент или юрлицо).
func (factureOut FactureOut) GetConsignee() Agent {
	return Deref(factureOut.Consignee)
}

// GetPaymentNumber возвращает Название платежного документа.
func (factureOut FactureOut) GetPaymentNumber() string {
	return Deref(factureOut.PaymentNumber)
}

// GetAttributes возвращает Список метаданных доп. полей.
func (factureOut FactureOut) GetAttributes() Slice[Attribute] {
	return factureOut.Attributes
}

// SetOrganization устанавливает Метаданные юрлица.
func (factureOut *FactureOut) SetOrganization(organization *Organization) *FactureOut {
	if organization != nil {
		factureOut.Organization = organization.Clean()
	}
	return factureOut
}

// SetApplicable устанавливает Отметку о проведении.
func (factureOut *FactureOut) SetApplicable(applicable bool) *FactureOut {
	factureOut.Applicable = &applicable
	return factureOut
}

// SetCode устанавливает выданного Счета-фактуры.
func (factureOut *FactureOut) SetCode(code string) *FactureOut {
	factureOut.Code = &code
	return factureOut
}

// SetContract устанавливает Метаданные договора.
//
// Передача nil передаёт сброс значения (null).
func (factureOut *FactureOut) SetContract(contract *Contract) *FactureOut {
	factureOut.Contract = NewNullValue(contract)
	return factureOut
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (factureOut *FactureOut) SetOwner(owner *Employee) *FactureOut {
	if owner != nil {
		factureOut.Owner = owner.Clean()
	}
	return factureOut
}

// SetDescription устанавливает Комментарий выданного Счета-фактуры.
func (factureOut *FactureOut) SetDescription(description string) *FactureOut {
	factureOut.Description = &description
	return factureOut
}

// SetExternalCode устанавливает Внешний код выданного Счета-фактуры.
func (factureOut *FactureOut) SetExternalCode(externalCode string) *FactureOut {
	factureOut.ExternalCode = &externalCode
	return factureOut
}

// SetFiles устанавливает Метаданные массива Файлов.
//
// Принимает множество объектов [File].
func (factureOut *FactureOut) SetFiles(files ...*File) *FactureOut {
	factureOut.Files = NewMetaArrayFrom(files)
	return factureOut
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (factureOut *FactureOut) SetGroup(group *Group) *FactureOut {
	if group != nil {
		factureOut.Group = group.Clean()
	}
	return factureOut
}

// SetMoment устанавливает Дату документа.
func (factureOut *FactureOut) SetMoment(moment time.Time) *FactureOut {
	factureOut.Moment = NewTimestamp(moment)
	return factureOut
}

// SetName устанавливает Наименование выданного Счета-фактуры.
func (factureOut *FactureOut) SetName(name string) *FactureOut {
	factureOut.Name = &name
	return factureOut
}

// SetPaymentDate устанавливает Дату платежного документа.
func (factureOut *FactureOut) SetPaymentDate(paymentDate time.Time) *FactureOut {
	factureOut.PaymentDate = NewTimestamp(paymentDate)
	return factureOut
}

// SetAgent устанавливает Метаданные Контрагента, связанного с бонусной операцией.
//
// Принимает [Counterparty] или [Organization].
func (factureOut *FactureOut) SetAgent(agent AgentOrganizationConverter) *FactureOut {
	if agent != nil {
		factureOut.Agent = agent.AsOrganizationAgent()
	}
	return factureOut
}

// SetMeta устанавливает Метаданные выданного Счета-фактуры.
func (factureOut *FactureOut) SetMeta(meta *Meta) *FactureOut {
	factureOut.Meta = meta
	return factureOut
}

// SetRate устанавливает Валюту.
//
// Передача nil передаёт сброс значения (null).
func (factureOut *FactureOut) SetRate(rate *Rate) *FactureOut {
	factureOut.Rate = NewNullValue(rate)
	return factureOut
}

// SetShared устанавливает флаг общего доступа.
func (factureOut *FactureOut) SetShared(shared bool) *FactureOut {
	factureOut.Shared = &shared
	return factureOut
}

// SetState устанавливает Метаданные статуса выданного Счета-фактуры.
//
// Передача nil передаёт сброс значения (null).
func (factureOut *FactureOut) SetState(state *State) *FactureOut {
	factureOut.State = NewNullValue(state)
	return factureOut
}

// SetStateContractID устанавливает Идентификатор государственного контракта, договора (соглашения).
func (factureOut *FactureOut) SetStateContractID(stateContractID string) *FactureOut {
	factureOut.StateContractID = &stateContractID
	return factureOut
}

// SetSyncID устанавливает ID синхронизации.
func (factureOut *FactureOut) SetSyncID(syncID uuid.UUID) *FactureOut {
	factureOut.SyncID = &syncID
	return factureOut
}

// SetDemands устанавливает Массив ссылок на связанные отгрузки.
//
// Принимает множество объектов [Demand].
func (factureOut *FactureOut) SetDemands(demands ...*Demand) *FactureOut {
	factureOut.Demands.Push(demands...)
	return factureOut
}

// SetPayments устанавливает Метаданные ссылок на связанные входящие платежи.
//
// Принимает множество объектов, реализующих интерфейс [PaymentConverter].
func (factureOut *FactureOut) SetPayments(payments ...PaymentConverter) *FactureOut {
	factureOut.Payments = NewPaymentsFrom(payments)
	return factureOut
}

// SetReturns устанавливает Массив ссылок на связанные возвраты.
//
// Принимает множество объектов [SalesReturn].
func (factureOut *FactureOut) SetReturns(returns ...*PurchaseReturn) *FactureOut {
	factureOut.Returns.Push(returns...)
	return factureOut
}

// SetConsignee устанавливает Метаданные грузополучателя (контрагент или юрлицо).
//
// Принимает [Counterparty] или [Organization].
func (factureOut *FactureOut) SetConsignee(consignee AgentOrganizationConverter) *FactureOut {
	if consignee != nil {
		factureOut.Consignee = consignee.AsOrganizationAgent()
	}
	return factureOut
}

// SetPaymentNumber устанавливает Название платежного документа.
func (factureOut *FactureOut) SetPaymentNumber(paymentNumber string) *FactureOut {
	factureOut.PaymentNumber = &paymentNumber
	return factureOut
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (factureOut *FactureOut) SetAttributes(attributes ...*Attribute) *FactureOut {
	factureOut.Attributes.Push(attributes...)
	return factureOut
}

// String реализует интерфейс [fmt.Stringer].
func (factureOut FactureOut) String() string {
	return Stringify(factureOut)
}

// MetaType возвращает код сущности.
func (FactureOut) MetaType() MetaType {
	return MetaTypeFactureOut
}

// Update shortcut
func (factureOut FactureOut) Update(ctx context.Context, client *Client, params ...*Params) (*FactureOut, *resty.Response, error) {
	return NewFactureOutService(client).Update(ctx, factureOut.GetID(), &factureOut, params...)
}

// Create shortcut
func (factureOut FactureOut) Create(ctx context.Context, client *Client, params ...*Params) (*FactureOut, *resty.Response, error) {
	return NewFactureOutService(client).Create(ctx, &factureOut, params...)
}

// Delete shortcut
func (factureOut FactureOut) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewFactureOutService(client).Delete(ctx, factureOut.GetID())
}

// FactureOutService описывает методы сервиса для работы со счетами-фактурами выданными.
type FactureOutService interface {
	// GetList выполняет запрос на получение списка выданных счетов-фактур.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[FactureOut], *resty.Response, error)

	// Create выполняет запрос на создание выданного счета-фактуры.
	// Обязательные поля для заполнения:
	//	- paymentNumber (Название платежного документа)
	//	- paymentDate (Дата платежного документа)
	// Принимает контекст, выданный счет-фактуру и опционально объект параметров запроса Params.
	// Возвращает созданную выданный счет-фактуру.
	Create(ctx context.Context, factureOut *FactureOut, params ...*Params) (*FactureOut, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение выданных счетов-фактур.
	// Изменяемые выданные счета-фактуры должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список выданных счетов-фактур и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых выданных счетов-фактур.
	CreateUpdateMany(ctx context.Context, factureOutList Slice[FactureOut], params ...*Params) (*Slice[FactureOut], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление выданных счетов-фактур.
	// Принимает контекст и множество выданных счетов-фактур.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*FactureOut) (*DeleteManyResponse, *resty.Response, error)

	// Delete выполняет запрос на удаление выданного счета-фактуры.
	// Принимает контекст и ID выданного счета-фактуры.
	// Возвращает «true» в случае успешного удаления выданного счета-фактуры.
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельного выданного счета-фактуры по ID.
	// Принимает контекст, ID выданного счета-фактуры и опционально объект параметров запроса Params.
	// Возвращает найденный выданный счет-фактуру.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*FactureOut, *resty.Response, error)

	// Update выполняет запрос на изменение выданного счета-фактуры.
	// Принимает контекст, выданный счет-фактуру и опционально объект параметров запроса Params.
	// Возвращает изменённую выданный счет-фактуру.
	Update(ctx context.Context, id uuid.UUID, factureOut *FactureOut, params ...*Params) (*FactureOut, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных выданных счетов-фактур.
	// Принимает контекст.
	// Возвращает объект метаданных MetaAttributesStatesSharedWrapper.
	GetMetadata(ctx context.Context) (*MetaAttributesStatesSharedWrapper, *resty.Response, error)

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

	// Template выполняет запрос на получение предзаполненного выданного счета-фактуры со стандартными полями без связи с какими-либо другими документами.
	// Принимает контекст.
	// Возвращает предзаполненный выданный счет-фактуру.
	Template(ctx context.Context) (*FactureOut, *resty.Response, error)

	// TemplateBased выполняет запрос на получение шаблона выданного счета-фактуры на основе других документов.
	// Основание, на котором может быть создан:
	//	- Отгрузка (Demand)
	//	- Возврат поставщику (PurchaseReturn)
	//	- Входящий платеж (PaymentIn)
	// Принимает контекст и множество документов из списка выше.
	// Возвращает предзаполненный выданный счет-фактуру на основании переданных документов.
	TemplateBased(ctx context.Context, basedOn ...MetaOwner) (*FactureOut, *resty.Response, error)

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
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*FactureOut, *resty.Response, error)

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
}

const (
	EndpointFactureOut = EndpointEntity + string(MetaTypeFactureOut)
)

// NewFactureOutService принимает [Client] и возвращает сервис для работы со счетами-фактурами выданными.
func NewFactureOutService(client *Client) FactureOutService {
	return newMainService[FactureOut, any, MetaAttributesStatesSharedWrapper, any](client, EndpointFactureOut)
}
