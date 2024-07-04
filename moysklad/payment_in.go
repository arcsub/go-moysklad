package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"time"
)

// PaymentIn Входящий платеж.
//
// Код сущности: paymentin
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vhodqschij-platezh
type PaymentIn struct {
	Meta                *Meta                    `json:"meta,omitempty"`                // Метаданные Входящего платежа
	Applicable          *bool                    `json:"applicable,omitempty"`          // Отметка о проведении
	AgentAccount        *AgentAccount            `json:"agentAccount,omitempty"`        // Метаданные счета контрагента
	Moment              *Timestamp               `json:"moment,omitempty"`              // Дата документа
	Operations          Operations               `json:"operations,omitempty"`          // Массив ссылок на связанные операции
	Name                *string                  `json:"name,omitempty"`                // Наименование Входящего платежа
	Contract            *NullValue[Contract]     `json:"contract,omitempty"`            // Метаданные договора
	Created             *Timestamp               `json:"created,omitempty"`             // Дата создания
	Deleted             *Timestamp               `json:"deleted,omitempty"`             // Момент последнего удаления Входящего платежа
	Description         *string                  `json:"description,omitempty"`         // Комментарий Входящего платежа
	ExternalCode        *string                  `json:"externalCode,omitempty"`        // Внешний код Входящего платежа
	Files               *MetaArray[File]         `json:"files,omitempty"`               // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group               *Group                   `json:"group,omitempty"`               // Отдел сотрудника
	ID                  *uuid.UUID               `json:"id,omitempty"`                  // ID Входящего платежа
	IncomingDate        *Timestamp               `json:"incomingDate,omitempty"`        // Входящая дата
	IncomingNumber      *string                  `json:"incomingNumber,omitempty"`      // Входящий номер
	FactureOut          *FactureOut              `json:"factureOut,omitempty"`          // Ссылка на выданный Счет-фактуру, с которым связан этот платеж
	Agent               *Agent                   `json:"agent,omitempty"`               // Метаданные контрагента
	Code                *string                  `json:"code,omitempty"`                // Код Входящего платежа
	Organization        *Organization            `json:"organization,omitempty"`        // Метаданные юрлица
	OrganizationAccount *AgentAccount            `json:"organizationAccount,omitempty"` // Метаданные счета юрлица
	Owner               *Employee                `json:"owner,omitempty"`               // Метаданные владельца (Сотрудника)
	PaymentPurpose      *string                  `json:"paymentPurpose,omitempty"`      // Назначение платежа
	Printed             *bool                    `json:"printed,omitempty"`             // Напечатан ли документ
	Project             *NullValue[Project]      `json:"project,omitempty"`             // Метаданные проекта
	Published           *bool                    `json:"published,omitempty"`           // Опубликован ли документ
	Rate                *NullValue[Rate]         `json:"rate,omitempty"`                // Валюта
	Shared              *bool                    `json:"shared,omitempty"`              // Общий доступ
	SalesChannel        *NullValue[SalesChannel] `json:"salesChannel,omitempty"`        // Метаданные канала продаж
	State               *NullValue[State]        `json:"state,omitempty"`               // Метаданные статуса Входящего платежа
	Sum                 *float64                 `json:"sum,omitempty"`                 // Сумма Входящего платежа в установленной валюте
	SyncID              *uuid.UUID               `json:"syncId,omitempty"`              // ID синхронизации
	Updated             *Timestamp               `json:"updated,omitempty"`             // Момент последнего обновления Входящего платежа
	AccountID           *uuid.UUID               `json:"accountId,omitempty"`           // ID учётной записи
	Attributes          Slice[Attribute]         `json:"attributes,omitempty"`          // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (paymentIn PaymentIn) Clean() *PaymentIn {
	if paymentIn.Meta == nil {
		return nil
	}
	return &PaymentIn{Meta: paymentIn.Meta}
}

// AsTaskOperation реализует интерфейс [TaskOperationInterface].
func (paymentIn PaymentIn) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: paymentIn.Meta}
}

// asPayment реализует интерфейс [PaymentInterface].
func (paymentIn PaymentIn) asPayment() *Payment {
	return &Payment{Meta: paymentIn.GetMeta()}
}

// GetMeta возвращает Метаданные Входящего платежа.
func (paymentIn PaymentIn) GetMeta() Meta {
	return Deref(paymentIn.Meta)
}

// GetApplicable возвращает Отметку о проведении.
func (paymentIn PaymentIn) GetApplicable() bool {
	return Deref(paymentIn.Applicable)
}

// GetAgentAccount возвращает Метаданные счета контрагента.
func (paymentIn PaymentIn) GetAgentAccount() AgentAccount {
	return Deref(paymentIn.AgentAccount)
}

// GetMoment возвращает Дату документа.
func (paymentIn PaymentIn) GetMoment() Timestamp {
	return Deref(paymentIn.Moment)
}

// GetOperations возвращает Метаданные связанных операций.
//
// Разрешенные типы связанных операций:
//   - CustomerOrder (Заказ покупателя)
//   - PurchaseReturn (Возврат поставщику)
//   - Demand (Отгрузка)
//   - InvoiceOut (Счет покупателю)
//   - CommissionReportIn (Полученный отчет комиссионера)
//   - RetailShift (Смена)
func (paymentIn PaymentIn) GetOperations() Operations {
	return paymentIn.Operations
}

// GetName возвращает Наименование Входящего платежа.
func (paymentIn PaymentIn) GetName() string {
	return Deref(paymentIn.Name)
}

// GetContract возвращает Метаданные договора.
func (paymentIn PaymentIn) GetContract() Contract {
	return paymentIn.Contract.GetValue()
}

// GetCreated возвращает Дату создания.
func (paymentIn PaymentIn) GetCreated() Timestamp {
	return Deref(paymentIn.Created)
}

// GetDeleted возвращает Момент последнего удаления Входящего платежа.
func (paymentIn PaymentIn) GetDeleted() Timestamp {
	return Deref(paymentIn.Deleted)
}

// GetDescription возвращает Комментарий Входящего платежа.
func (paymentIn PaymentIn) GetDescription() string {
	return Deref(paymentIn.Description)
}

// GetExternalCode возвращает Внешний код Входящего платежа.
func (paymentIn PaymentIn) GetExternalCode() string {
	return Deref(paymentIn.ExternalCode)
}

// GetFiles возвращает Метаданные массива Файлов.
func (paymentIn PaymentIn) GetFiles() MetaArray[File] {
	return Deref(paymentIn.Files)
}

// GetGroup возвращает Отдел сотрудника.
func (paymentIn PaymentIn) GetGroup() Group {
	return Deref(paymentIn.Group)
}

// GetID возвращает ID Перемещения.
func (paymentIn PaymentIn) GetID() uuid.UUID {
	return Deref(paymentIn.ID)
}

// GetIncomingDate возвращает Входящую дату.
func (paymentIn PaymentIn) GetIncomingDate() Timestamp {
	return Deref(paymentIn.IncomingDate)
}

// GetIncomingNumber возвращает Входящий номер.
func (paymentIn PaymentIn) GetIncomingNumber() string {
	return Deref(paymentIn.IncomingNumber)
}

// GetFactureOut возвращает Ссылку на полученный счет-фактуру, с которым связан этот платеж.
func (paymentIn PaymentIn) GetFactureOut() FactureOut {
	return Deref(paymentIn.FactureOut)
}

// GetAgent возвращает Метаданные Контрагента.
func (paymentIn PaymentIn) GetAgent() Agent {
	return Deref(paymentIn.Agent)
}

// GetCode возвращает Код Входящего платежа.
func (paymentIn PaymentIn) GetCode() string {
	return Deref(paymentIn.Code)
}

// GetOrganization возвращает Метаданные юрлица.
func (paymentIn PaymentIn) GetOrganization() Organization {
	return Deref(paymentIn.Organization)
}

// GetOrganizationAccount возвращает Метаданные счета юрлица.
func (paymentIn PaymentIn) GetOrganizationAccount() AgentAccount {
	return Deref(paymentIn.OrganizationAccount)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (paymentIn PaymentIn) GetOwner() Employee {
	return Deref(paymentIn.Owner)
}

// GetPaymentPurpose возвращает Назначение платежа.
func (paymentIn PaymentIn) GetPaymentPurpose() string {
	return Deref(paymentIn.PaymentPurpose)
}

// GetPrinted возвращает true, если документ напечатан.
func (paymentIn PaymentIn) GetPrinted() bool {
	return Deref(paymentIn.Printed)
}

// GetProject возвращает Метаданные проекта.
func (paymentIn PaymentIn) GetProject() Project {
	return paymentIn.Project.GetValue()
}

// GetPublished возвращает true, если документ опубликован.
func (paymentIn PaymentIn) GetPublished() bool {
	return Deref(paymentIn.Published)
}

// GetRate возвращает Валюту.
func (paymentIn PaymentIn) GetRate() Rate {
	return paymentIn.Rate.GetValue()
}

// GetShared возвращает флаг Общего доступа.
func (paymentIn PaymentIn) GetShared() bool {
	return Deref(paymentIn.Shared)
}

// GetSalesChannel возвращает Метаданные канала продаж.
func (paymentIn PaymentIn) GetSalesChannel() SalesChannel {
	return paymentIn.SalesChannel.GetValue()
}

// GetState возвращает Метаданные статуса Входящего платежа.
func (paymentIn PaymentIn) GetState() State {
	return paymentIn.State.GetValue()
}

// GetSum возвращает Сумму Входящего платежа в установленной валюте.
func (paymentIn PaymentIn) GetSum() float64 {
	return Deref(paymentIn.Sum)
}

// GetSyncID возвращает ID синхронизации.
func (paymentIn PaymentIn) GetSyncID() uuid.UUID {
	return Deref(paymentIn.SyncID)
}

// GetUpdated возвращает Момент последнего обновления Входящего платежа.
func (paymentIn PaymentIn) GetUpdated() Timestamp {
	return Deref(paymentIn.Updated)
}

// GetAccountID возвращает ID учётной записи.
func (paymentIn PaymentIn) GetAccountID() uuid.UUID {
	return Deref(paymentIn.AccountID)
}

// GetAttributes возвращает Список метаданных доп. полей.
func (paymentIn PaymentIn) GetAttributes() Slice[Attribute] {
	return paymentIn.Attributes
}

// SetMeta устанавливает Метаданные Входящего платежа.
func (paymentIn *PaymentIn) SetMeta(meta *Meta) *PaymentIn {
	paymentIn.Meta = meta
	return paymentIn
}

// SetApplicable устанавливает Отметку о проведении.
func (paymentIn *PaymentIn) SetApplicable(applicable bool) *PaymentIn {
	paymentIn.Applicable = &applicable
	return paymentIn
}

// SetAgentAccount устанавливает Метаданные счета контрагента.
func (paymentIn *PaymentIn) SetAgentAccount(agentAccount *AgentAccount) *PaymentIn {
	if agentAccount != nil {
		paymentIn.AgentAccount = agentAccount.Clean()
	}
	return paymentIn
}

// SetMoment устанавливает Дату документа.
func (paymentIn *PaymentIn) SetMoment(moment time.Time) *PaymentIn {
	paymentIn.Moment = NewTimestamp(moment)
	return paymentIn
}

// SetOperations устанавливает Метаданные связанных операций.
//
// Разрешенные типы связанных операций:
//   - CustomerOrder (Заказ покупателя)
//   - PurchaseReturn (Возврат поставщику)
//   - Demand (Отгрузка)
//   - InvoiceOut (Счет покупателю)
//   - CommissionReportIn (Полученный отчет комиссионера)
//   - RetailShift (Смена)
//
// Принимает множество объектов, реализующих интерфейс [OperationIn].
func (paymentIn *PaymentIn) SetOperations(operations ...OperationIn) *PaymentIn {
	for _, operation := range operations {
		if operation != nil {
			paymentIn.Operations.Push(operation.AsOperationIn())
		}
	}
	return paymentIn
}

// SetName устанавливает Наименование Входящего платежа.
func (paymentIn *PaymentIn) SetName(name string) *PaymentIn {
	paymentIn.Name = &name
	return paymentIn
}

// SetContract устанавливает Метаданные договора.
//
// Передача nil передаёт сброс значения (null).
func (paymentIn *PaymentIn) SetContract(contract *Contract) *PaymentIn {
	paymentIn.Contract = NewNullValue(contract)
	return paymentIn
}

// SetDescription устанавливает Комментарий Входящего платежа.
func (paymentIn *PaymentIn) SetDescription(description string) *PaymentIn {
	paymentIn.Description = &description
	return paymentIn
}

// SetExternalCode устанавливает Внешний код Входящего платежа.
func (paymentIn *PaymentIn) SetExternalCode(externalCode string) *PaymentIn {
	paymentIn.ExternalCode = &externalCode
	return paymentIn
}

// SetFiles устанавливает Метаданные массива Файлов.
//
// Принимает множество объектов [File].
func (paymentIn *PaymentIn) SetFiles(files ...*File) *PaymentIn {
	paymentIn.Files = NewMetaArrayFrom(files)
	return paymentIn
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (paymentIn *PaymentIn) SetGroup(group *Group) *PaymentIn {
	if group != nil {
		paymentIn.Group = group.Clean()
	}
	return paymentIn
}

// SetIncomingDate устанавливает Входящую дату.
func (paymentIn *PaymentIn) SetIncomingDate(incomingDate time.Time) *PaymentIn {
	paymentIn.IncomingDate = NewTimestamp(incomingDate)
	return paymentIn
}

// SetIncomingNumber устанавливает Входящий номер.
func (paymentIn *PaymentIn) SetIncomingNumber(incomingNumber string) *PaymentIn {
	paymentIn.IncomingNumber = &incomingNumber
	return paymentIn
}

// SetFactureOut устанавливает Метаданные Счет-фактуры выданного.
func (paymentIn *PaymentIn) SetFactureOut(factureOut *FactureOut) *PaymentIn {
	if factureOut != nil {
		paymentIn.FactureOut = factureOut.Clean()
	}
	return paymentIn
}

// SetAgent устанавливает Метаданные Контрагента.
//
// Принимает [Counterparty] или [Organization].
func (paymentIn *PaymentIn) SetAgent(agent AgentCounterpartyOrganizationInterface) *PaymentIn {
	if agent != nil {
		paymentIn.Agent = agent.AsCOAgent()
	}
	return paymentIn
}

// SetCode устанавливает Код Входящего платежа.
func (paymentIn *PaymentIn) SetCode(code string) *PaymentIn {
	paymentIn.Code = &code
	return paymentIn
}

// SetOrganization устанавливает Метаданные юрлица.
func (paymentIn *PaymentIn) SetOrganization(organization *Organization) *PaymentIn {
	if organization != nil {
		paymentIn.Organization = organization.Clean()
	}
	return paymentIn
}

// SetOrganizationAccount устанавливает Метаданные счета юрлица.
func (paymentIn *PaymentIn) SetOrganizationAccount(organizationAccount *AgentAccount) *PaymentIn {
	if organizationAccount != nil {
		paymentIn.OrganizationAccount = organizationAccount.Clean()
	}
	return paymentIn
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (paymentIn *PaymentIn) SetOwner(owner *Employee) *PaymentIn {
	if owner != nil {
		paymentIn.Owner = owner.Clean()
	}
	return paymentIn
}

// SetPaymentPurpose устанавливает Назначение платежа.
func (paymentIn *PaymentIn) SetPaymentPurpose(paymentPurpose string) *PaymentIn {
	paymentIn.PaymentPurpose = &paymentPurpose
	return paymentIn
}

// SetProject устанавливает Метаданные проекта.
//
// Передача nil передаёт сброс значения (null).
func (paymentIn *PaymentIn) SetProject(project *Project) *PaymentIn {
	paymentIn.Project = NewNullValue(project)
	return paymentIn
}

// SetRate устанавливает Валюту.
//
// Передача nil передаёт сброс значения (null).
func (paymentIn *PaymentIn) SetRate(rate *Rate) *PaymentIn {
	paymentIn.Rate = NewNullValue(rate)
	return paymentIn
}

// SetShared устанавливает флаг общего доступа.
func (paymentIn *PaymentIn) SetShared(shared bool) *PaymentIn {
	paymentIn.Shared = &shared
	return paymentIn
}

// SetSalesChannel устанавливает Метаданные канала продаж.
//
// Передача nil передаёт сброс значения (null).
func (paymentIn *PaymentIn) SetSalesChannel(salesChannel *SalesChannel) *PaymentIn {
	paymentIn.SalesChannel = NewNullValue(salesChannel)
	return paymentIn
}

// SetState устанавливает Метаданные статуса Входящего платежа.
//
// Передача nil передаёт сброс значения (null).
func (paymentIn *PaymentIn) SetState(state *State) *PaymentIn {
	paymentIn.State = NewNullValue(state)
	return paymentIn
}

// SetSum устанавливает Сумму Входящего платежа в установленной валюте.
func (paymentIn *PaymentIn) SetSum(sum float64) *PaymentIn {
	paymentIn.Sum = &sum
	return paymentIn
}

// SetSyncID устанавливает ID синхронизации.
func (paymentIn *PaymentIn) SetSyncID(syncID uuid.UUID) *PaymentIn {
	paymentIn.SyncID = &syncID
	return paymentIn
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (paymentIn *PaymentIn) SetAttributes(attributes ...*Attribute) *PaymentIn {
	paymentIn.Attributes.Push(attributes...)
	return paymentIn
}

// String реализует интерфейс [fmt.Stringer].
func (paymentIn PaymentIn) String() string {
	return Stringify(paymentIn)
}

// MetaType возвращает код сущности.
func (PaymentIn) MetaType() MetaType {
	return MetaTypePaymentIn
}

// AsOperation возвращает объект [Operation] c полями meta и linkedSum.
//
// Значение поля linkedSum заполняется из поля sum.
func (paymentIn PaymentIn) AsOperation() *Operation {
	return &Operation{Meta: paymentIn.GetMeta(), LinkedSum: paymentIn.GetSum()}
}

// Update shortcut
func (paymentIn PaymentIn) Update(ctx context.Context, client *Client, params ...*Params) (*PaymentIn, *resty.Response, error) {
	return NewPaymentInService(client).Update(ctx, paymentIn.GetID(), &paymentIn, params...)
}

// Create shortcut
func (paymentIn PaymentIn) Create(ctx context.Context, client *Client, params ...*Params) (*PaymentIn, *resty.Response, error) {
	return NewPaymentInService(client).Create(ctx, &paymentIn, params...)
}

// Delete shortcut
func (paymentIn PaymentIn) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewPaymentInService(client).Delete(ctx, paymentIn.GetID())
}

// PaymentInService описывает методы сервиса для работы с входящими платежами.
type PaymentInService interface {
	// GetList выполняет запрос на получение списка входящих платежей.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[PaymentIn], *resty.Response, error)

	// Create выполняет запрос на создание входящего платежа.
	// Обязательные поля для заполнения:
	//	- organization (Ссылка на ваше юрлицо)
	//	- agent (Ссылка на контрагента)
	// Принимает контекст, входящий платеж и опционально объект параметров запроса Params.
	// Возвращает созданный входящий платеж.
	Create(ctx context.Context, paymentIn *PaymentIn, params ...*Params) (*PaymentIn, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение входящих платежей.
	// Изменяемые входящие платежи должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список входящих платежей и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых входящих платежей.
	CreateUpdateMany(ctx context.Context, paymentInList Slice[PaymentIn], params ...*Params) (*Slice[PaymentIn], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление входящих платежей.
	// Принимает контекст и множество входящих платежей.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*PaymentIn) (*DeleteManyResponse, *resty.Response, error)

	// Delete выполняет запрос на удаление входящего платежа.
	// Принимает контекст и ID входящего платежа.
	// Возвращает true в случае успешного удаления входящего платежа.
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельного входящего платежа по ID.
	// Принимает контекст, ID входящего платежа и опционально объект параметров запроса Params.
	// Возвращает найденный входящий платеж.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*PaymentIn, *resty.Response, error)

	// Update выполняет запрос на изменение входящего платежа.
	// Принимает контекст, входящий платеж и опционально объект параметров запроса Params.
	// Возвращает изменённый входящий платеж.
	Update(ctx context.Context, id uuid.UUID, paymentIn *PaymentIn, params ...*Params) (*PaymentIn, *resty.Response, error)

	// Template выполняет запрос на получение предзаполненного входящего платежа со стандартными полями без связи с какими-либо другими документами.
	// Принимает контекст.
	// Возвращает предзаполненный входящий платеж.
	Template(ctx context.Context) (*PaymentIn, *resty.Response, error)

	// TemplateBased выполняет запрос на получение шаблона входящего платежа на основе других документов.
	// Основание, на котором может быть создан:
	//	- Заказ покупателя (CustomerOrder)
	//	- Возврат поставщику (PurchaseReturn)
	//	- Отгрузка (Demand)
	//	- Счет покупателю (InvoiceOut)
	//	- Полученный отчет комиссионера (CommissionReportIn)
	// Принимает контекст и множество документов из списка выше.
	// Возвращает предзаполненный входящий платеж на основании переданных документов.
	TemplateBased(ctx context.Context, basedOn ...MetaOwner) (*PaymentIn, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных входящих платежей.
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
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*PaymentIn, *resty.Response, error)

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
}

// NewPaymentInService принимает [Client] и возвращает сервис для работы с входящими платежами.
func NewPaymentInService(client *Client) PaymentInService {
	return newMainService[PaymentIn, any, MetaAttributesStatesSharedWrapper, any](NewEndpoint(client, "entity/paymentin"))
}
