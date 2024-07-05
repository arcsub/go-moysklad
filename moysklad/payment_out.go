package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"time"
)

// PaymentOut Исходящий платеж.
//
// Код сущности: paymentout
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-ishodqschij-platezh
type PaymentOut struct {
	Moment              *Timestamp               `json:"moment,omitempty"`              // Дата документа
	Applicable          *bool                    `json:"applicable,omitempty"`          // Отметка о проведении
	AgentAccount        *AgentAccount            `json:"agentAccount,omitempty"`        // Метаданные счета контрагента или юр.лица
	Name                *string                  `json:"name,omitempty"`                // Наименование Исходящего платежа
	Operations          Operations               `json:"operations,omitempty"`          // Массив ссылок на связанные операции
	Organization        *Organization            `json:"organization,omitempty"`        // Метаданные юрлица
	Contract            *NullValue[Contract]     `json:"contract,omitempty"`            // Метаданные договора
	Created             *Timestamp               `json:"created,omitempty"`             // Дата создания
	Deleted             *Timestamp               `json:"deleted,omitempty"`             // Момент последнего удаления Исходящего платежа
	Description         *string                  `json:"description,omitempty"`         // Комментарий Исходящего платежа
	ExpenseItem         *ExpenseItem             `json:"expenseItem,omitempty"`         // Метаданные Статьи расходов
	ExternalCode        *string                  `json:"externalCode,omitempty"`        // Внешний код Исходящего платежа
	Files               *MetaArray[File]         `json:"files,omitempty"`               // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group               *Group                   `json:"group,omitempty"`               // Отдел сотрудника
	ID                  *uuid.UUID               `json:"id,omitempty"`                  // ID Исходящего платежа
	Meta                *Meta                    `json:"meta,omitempty"`                // Метаданные Исходящего платежа
	FactureIn           *FactureIn               `json:"factureIn,omitempty"`           // Ссылка на Счет-фактуру полученный, с которым связан этот платеж
	Agent               *Agent                   `json:"agent,omitempty"`               // Метаданные контрагента, сотрудника или юр.лица
	Code                *string                  `json:"code,omitempty"`                // Код Исходящего платежа
	OrganizationAccount *AgentAccount            `json:"organizationAccount,omitempty"` // Метаданные счета юрлица
	Owner               *Employee                `json:"owner,omitempty"`               // Метаданные владельца (Сотрудника)
	PaymentPurpose      *string                  `json:"paymentPurpose,omitempty"`      // Назначение платежа
	Printed             *bool                    `json:"printed,omitempty"`             // Напечатан ли документ
	Project             *NullValue[Project]      `json:"project,omitempty"`             // Метаданные проекта
	Published           *bool                    `json:"published,omitempty"`           // Опубликован ли документ
	Rate                *NullValue[Rate]         `json:"rate,omitempty"`                // Валюта
	SalesChannel        *NullValue[SalesChannel] `json:"salesChannel,omitempty"`        // Метаданные канала продаж
	Shared              *bool                    `json:"shared,omitempty"`              // Общий доступ
	State               *NullValue[State]        `json:"state,omitempty"`               // Метаданные статуса Исходящего платежа
	Sum                 *float64                 `json:"sum,omitempty"`                 // Сумма Исходящего платежа в установленной валюте
	SyncID              *uuid.UUID               `json:"syncId,omitempty"`              // ID синхронизации
	Updated             *Timestamp               `json:"updated,omitempty"`             // Момент последнего обновления Исходящего платежа
	VatSum              *float64                 `json:"vatSum,omitempty"`              // Сумма НДС
	AccountID           *uuid.UUID               `json:"accountId,omitempty"`           // ID учётной записи
	Attributes          Slice[Attribute]         `json:"attributes,omitempty"`          // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (paymentOut PaymentOut) Clean() *PaymentOut {
	if paymentOut.Meta == nil {
		return nil
	}
	return &PaymentOut{Meta: paymentOut.Meta}
}

// AsTaskOperation реализует интерфейс [TaskOperationInterface].
func (paymentOut PaymentOut) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: paymentOut.Meta}
}

// AsOperation возвращает объект [Operation] c полями meta и linkedSum.
//
// Значение поля linkedSum заполняется из поля sum.
func (paymentOut PaymentOut) AsOperation() *Operation {
	return &Operation{Meta: paymentOut.GetMeta(), LinkedSum: paymentOut.GetSum()}
}

// asPayment реализует интерфейс [PaymentInterface].
func (paymentOut PaymentOut) asPayment() *Payment {
	return &Payment{Meta: paymentOut.GetMeta()}
}

// GetMoment возвращает Дату документа.
func (paymentOut PaymentOut) GetMoment() Timestamp {
	return Deref(paymentOut.Moment)
}

// GetApplicable возвращает Отметку о проведении.
func (paymentOut PaymentOut) GetApplicable() bool {
	return Deref(paymentOut.Applicable)
}

// GetAgentAccount возвращает Метаданные счета контрагента.
func (paymentOut PaymentOut) GetAgentAccount() AgentAccount {
	return Deref(paymentOut.AgentAccount)
}

// GetName возвращает Наименование Исходящего платежа.
func (paymentOut PaymentOut) GetName() string {
	return Deref(paymentOut.Name)
}

// GetOperations возвращает Массив ссылок на связанные операции.
//
// Разрешенные типы связанных операций:
//   - SalesReturn (Возврат покупателя)
//   - Supply (Приемка)
//   - InvoiceIn (Счет поставщика)
//   - PurchaseOrder (Заказ поставщику)
//   - CommissionReportOut (Выданный отчет комиссионера)
func (paymentOut PaymentOut) GetOperations() Operations {
	return paymentOut.Operations
}

// GetOrganization возвращает Метаданные юрлица.
func (paymentOut PaymentOut) GetOrganization() Organization {
	return Deref(paymentOut.Organization)
}

// GetContract возвращает Метаданные договора.
func (paymentOut PaymentOut) GetContract() Contract {
	return Deref(paymentOut.Contract).GetValue()
}

// GetCreated возвращает Дату создания.
func (paymentOut PaymentOut) GetCreated() Timestamp {
	return Deref(paymentOut.Created)
}

// GetDeleted возвращает Момент последнего удаления Исходящего платежа.
func (paymentOut PaymentOut) GetDeleted() Timestamp {
	return Deref(paymentOut.Deleted)
}

// GetDescription возвращает Комментарий Исходящего платежа.
func (paymentOut PaymentOut) GetDescription() string {
	return Deref(paymentOut.Description)
}

// GetExpenseItem возвращает Метаданные Статьи расходов.
func (paymentOut PaymentOut) GetExpenseItem() ExpenseItem {
	return Deref(paymentOut.ExpenseItem)
}

// GetExternalCode возвращает Внешний код Исходящего платежа.
func (paymentOut PaymentOut) GetExternalCode() string {
	return Deref(paymentOut.ExternalCode)
}

// GetFiles возвращает Метаданные массива Файлов.
func (paymentOut PaymentOut) GetFiles() MetaArray[File] {
	return Deref(paymentOut.Files)
}

// GetGroup возвращает Отдел сотрудника.
func (paymentOut PaymentOut) GetGroup() Group {
	return Deref(paymentOut.Group)
}

// GetID возвращает ID Исходящего платежа.
func (paymentOut PaymentOut) GetID() uuid.UUID {
	return Deref(paymentOut.ID)
}

// GetMeta возвращает Метаданные Исходящего платежа.
func (paymentOut PaymentOut) GetMeta() Meta {
	return Deref(paymentOut.Meta)
}

// GetFactureIn возвращает Метаданные Счет-фактуры полученного.
func (paymentOut PaymentOut) GetFactureIn() FactureIn {
	return Deref(paymentOut.FactureIn)
}

// GetAgent возвращает Метаданные контрагента, сотрудника или юр.лица.
func (paymentOut PaymentOut) GetAgent() Agent {
	return Deref(paymentOut.Agent)
}

// GetCode возвращает Код Исходящего платежа.
func (paymentOut PaymentOut) GetCode() string {
	return Deref(paymentOut.Code)
}

// GetOrganizationAccount возвращает Метаданные счета юрлица.
func (paymentOut PaymentOut) GetOrganizationAccount() AgentAccount {
	return Deref(paymentOut.OrganizationAccount)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (paymentOut PaymentOut) GetOwner() Employee {
	return Deref(paymentOut.Owner)
}

// GetPaymentPurpose возвращает Назначение платежа.
func (paymentOut PaymentOut) GetPaymentPurpose() string {
	return Deref(paymentOut.PaymentPurpose)
}

// GetPrinted возвращает true, если документ напечатан.
func (paymentOut PaymentOut) GetPrinted() bool {
	return Deref(paymentOut.Printed)
}

// GetProject возвращает Метаданные проекта.
func (paymentOut PaymentOut) GetProject() Project {
	return Deref(paymentOut.Project).GetValue()
}

// GetPublished возвращает true, если документ опубликован.
func (paymentOut PaymentOut) GetPublished() bool {
	return Deref(paymentOut.Published)
}

// GetRate возвращает Курс Валюты.
func (paymentOut PaymentOut) GetRate() Rate {
	return Deref(paymentOut.Rate).GetValue()
}

// GetSalesChannel возвращает Метаданные канала продаж.
func (paymentOut PaymentOut) GetSalesChannel() SalesChannel {
	return Deref(paymentOut.SalesChannel).GetValue()
}

// GetShared возвращает флаг Общего доступа.
func (paymentOut PaymentOut) GetShared() bool {
	return Deref(paymentOut.Shared)
}

// GetState возвращает Метаданные статуса Исходящего платежа.
func (paymentOut PaymentOut) GetState() State {
	return Deref(paymentOut.State).GetValue()
}

// GetSum возвращает Сумму Исходящего платежа в копейках.
func (paymentOut PaymentOut) GetSum() float64 {
	return Deref(paymentOut.Sum)
}

// GetSyncID возвращает ID синхронизации.
func (paymentOut PaymentOut) GetSyncID() uuid.UUID {
	return Deref(paymentOut.SyncID)
}

// GetUpdated возвращает Момент последнего обновления Исходящего платежа.
func (paymentOut PaymentOut) GetUpdated() Timestamp {
	return Deref(paymentOut.Updated)
}

// GetVatSum возвращает Сумму НДС.
func (paymentOut PaymentOut) GetVatSum() float64 {
	return Deref(paymentOut.VatSum)
}

// GetAccountID возвращает ID учётной записи.
func (paymentOut PaymentOut) GetAccountID() uuid.UUID {
	return Deref(paymentOut.AccountID)
}

// GetAttributes возвращает Список метаданных доп. полей.
func (paymentOut PaymentOut) GetAttributes() Slice[Attribute] {
	return paymentOut.Attributes
}

// SetMoment устанавливает Дату документа.
func (paymentOut *PaymentOut) SetMoment(moment time.Time) *PaymentOut {
	paymentOut.Moment = NewTimestamp(moment)
	return paymentOut
}

// SetApplicable устанавливает Отметку о проведении.
func (paymentOut *PaymentOut) SetApplicable(applicable bool) *PaymentOut {
	paymentOut.Applicable = &applicable
	return paymentOut
}

// SetAgentAccount устанавливает Метаданные счета контрагента.
func (paymentOut *PaymentOut) SetAgentAccount(agentAccount *AgentAccount) *PaymentOut {
	if agentAccount != nil {
		paymentOut.AgentAccount = agentAccount.Clean()
	}
	return paymentOut
}

// SetName устанавливает Наименование Исходящего платежа.
func (paymentOut *PaymentOut) SetName(name string) *PaymentOut {
	paymentOut.Name = &name
	return paymentOut
}

// SetOperations устанавливает Метаданные связанных операций.
//
// Разрешенные типы связанных операций:
//   - SalesReturn (Возврат покупателя)
//   - Supply (Приемка)
//   - InvoiceIn (Счет поставщика)
//   - PurchaseOrder (Заказ поставщику)
//   - CommissionReportOut (Выданный отчет комиссионера)
//
// Принимает множество объектов, реализующих интерфейс [OperationOut].
func (paymentOut *PaymentOut) SetOperations(operations ...OperationOut) *PaymentOut {
	for _, operation := range operations {
		if operation != nil {
			paymentOut.Operations.Push(operation.AsOperationOut())
		}
	}
	return paymentOut
}

// SetOrganization устанавливает Метаданные юрлица.
func (paymentOut *PaymentOut) SetOrganization(organization *Organization) *PaymentOut {
	if organization != nil {
		paymentOut.Organization = organization.Clean()
	}
	return paymentOut
}

// SetContract устанавливает Метаданные договора.
//
// Передача nil передаёт сброс значения (null).
func (paymentOut *PaymentOut) SetContract(contract *Contract) *PaymentOut {
	paymentOut.Contract = NewNullValue(contract)
	return paymentOut
}

// SetDescription устанавливает Комментарий Исходящего платежа.
func (paymentOut *PaymentOut) SetDescription(description string) *PaymentOut {
	paymentOut.Description = &description
	return paymentOut
}

// SetExpenseItem устанавливает Метаданные Статьи расходов.
func (paymentOut *PaymentOut) SetExpenseItem(expenseItem *ExpenseItem) *PaymentOut {
	if expenseItem != nil {
		paymentOut.ExpenseItem = expenseItem.Clean()
	}
	return paymentOut
}

// SetExternalCode устанавливает Внешний код Исходящего платежа.
func (paymentOut *PaymentOut) SetExternalCode(externalCode string) *PaymentOut {
	paymentOut.ExternalCode = &externalCode
	return paymentOut
}

// SetFiles устанавливает Метаданные массива Файлов.
//
// Принимает множество объектов [File].
func (paymentOut *PaymentOut) SetFiles(files ...*File) *PaymentOut {
	paymentOut.Files = NewMetaArrayFrom(files)
	return paymentOut
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (paymentOut *PaymentOut) SetGroup(group *Group) *PaymentOut {
	if group != nil {
		paymentOut.Group = group.Clean()
	}
	return paymentOut
}

// SetMeta устанавливает Метаданные Исходящего платежа.
func (paymentOut *PaymentOut) SetMeta(meta *Meta) *PaymentOut {
	paymentOut.Meta = meta
	return paymentOut
}

// SetFactureIn устанавливает Метаданные Счет-фактуры полученного.
func (paymentOut *PaymentOut) SetFactureIn(factureIn *FactureIn) *PaymentOut {
	if factureIn != nil {
		paymentOut.FactureIn = factureIn.Clean()
	}
	return paymentOut
}

// SetAgent устанавливает Метаданные контрагента.
//
// Принимает [Counterparty], [Organization] или [Employee].
func (paymentOut *PaymentOut) SetAgent(agent AgentInterface) *PaymentOut {
	if agent != nil {
		paymentOut.Agent = agent.AsAgent()
	}
	return paymentOut
}

// SetCode устанавливает Код Исходящего платежа.
func (paymentOut *PaymentOut) SetCode(code string) *PaymentOut {
	paymentOut.Code = &code
	return paymentOut
}

// SetOrganizationAccount устанавливает Метаданные счета юрлица.
func (paymentOut *PaymentOut) SetOrganizationAccount(organizationAccount *AgentAccount) *PaymentOut {
	if organizationAccount != nil {
		paymentOut.OrganizationAccount = organizationAccount.Clean()
	}
	return paymentOut
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (paymentOut *PaymentOut) SetOwner(owner *Employee) *PaymentOut {
	if owner != nil {
		paymentOut.Owner = owner.Clean()
	}
	return paymentOut
}

// SetPaymentPurpose устанавливает Назначение платежа.
func (paymentOut *PaymentOut) SetPaymentPurpose(paymentPurpose string) *PaymentOut {
	paymentOut.PaymentPurpose = &paymentPurpose
	return paymentOut
}

// SetProject устанавливает Метаданные проекта.
//
// Передача nil передаёт сброс значения (null).
func (paymentOut *PaymentOut) SetProject(project *Project) *PaymentOut {
	paymentOut.Project = NewNullValue(project)
	return paymentOut
}

// SetRate устанавливает Валюту.
//
// Передача nil передаёт сброс значения (null).
func (paymentOut *PaymentOut) SetRate(rate *Rate) *PaymentOut {
	paymentOut.Rate = NewNullValue(rate)
	return paymentOut
}

// SetSalesChannel устанавливает Метаданные канала продаж.
//
// Передача nil передаёт сброс значения (null).
func (paymentOut *PaymentOut) SetSalesChannel(salesChannel *SalesChannel) *PaymentOut {
	paymentOut.SalesChannel = NewNullValue(salesChannel)
	return paymentOut
}

// SetShared устанавливает флаг общего доступа.
func (paymentOut *PaymentOut) SetShared(shared bool) *PaymentOut {
	paymentOut.Shared = &shared
	return paymentOut
}

// SetState устанавливает Метаданные статуса Исходящего платежа.
//
// Передача nil передаёт сброс значения (null).
func (paymentOut *PaymentOut) SetState(state *State) *PaymentOut {
	paymentOut.State = NewNullValue(state)
	return paymentOut
}

// SetSum устанавливает Сумму Исходящего платежа в установленной валюте.
func (paymentOut *PaymentOut) SetSum(sum float64) *PaymentOut {
	paymentOut.Sum = &sum
	return paymentOut
}

// SetSyncID устанавливает ID синхронизации.
func (paymentOut *PaymentOut) SetSyncID(syncID uuid.UUID) *PaymentOut {
	paymentOut.SyncID = &syncID
	return paymentOut
}

// SetVatSum устанавливает Сумму НДС.
func (paymentOut *PaymentOut) SetVatSum(vatSum float64) *PaymentOut {
	paymentOut.VatSum = &vatSum
	return paymentOut
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (paymentOut *PaymentOut) SetAttributes(attributes ...*Attribute) *PaymentOut {
	paymentOut.Attributes.Push(attributes...)
	return paymentOut
}

// String реализует интерфейс [fmt.Stringer].
func (paymentOut PaymentOut) String() string {
	return Stringify(paymentOut)
}

// MetaType возвращает код сущности.
func (PaymentOut) MetaType() MetaType {
	return MetaTypePaymentOut
}

// Update shortcut
func (paymentOut PaymentOut) Update(ctx context.Context, client *Client, params ...*Params) (*PaymentOut, *resty.Response, error) {
	return NewPaymentOutService(client).Update(ctx, paymentOut.GetID(), &paymentOut, params...)
}

// Create shortcut
func (paymentOut PaymentOut) Create(ctx context.Context, client *Client, params ...*Params) (*PaymentOut, *resty.Response, error) {
	return NewPaymentOutService(client).Create(ctx, &paymentOut, params...)
}

// Delete shortcut
func (paymentOut PaymentOut) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewPaymentOutService(client).Delete(ctx, paymentOut.GetID())
}

// PaymentOutService описывает методы сервиса для работы с исходящими платежами.
type PaymentOutService interface {
	// GetList выполняет запрос на получение списка исходящих платежей.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[PaymentOut], *resty.Response, error)

	// Create выполняет запрос на создание исходящего платежа.
	// Обязательные поля для заполнения:
	//	- organization (Ссылка на ваше юрлицо)
	//	- agent (Ссылка на контрагента)
	//	- expenseItem (Статья расходов)
	// Принимает контекст, исходящий платеж и опционально объект параметров запроса Params.
	// Возвращает созданный исходящий платеж.
	Create(ctx context.Context, paymentOut *PaymentOut, params ...*Params) (*PaymentOut, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение исходящих платежей.
	// Изменяемые исходящие платежи должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список исходящих платежей и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых исходящих платежей.
	CreateUpdateMany(ctx context.Context, paymentOutList Slice[PaymentOut], params ...*Params) (*Slice[PaymentOut], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление исходящих платежей.
	// Принимает контекст и множество исходящих платежей.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*PaymentOut) (*DeleteManyResponse, *resty.Response, error)

	// Delete выполняет запрос на удаление исходящего платежа.
	// Принимает контекст и ID исходящего платежа.
	// Возвращает true в случае успешного удаления исходящего платежа.
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельного исходящего платежа по ID.
	// Принимает контекст, ID исходящего платежа и опционально объект параметров запроса Params.
	// Возвращает найденный исходящий платеж.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*PaymentOut, *resty.Response, error)

	// Update выполняет запрос на изменение исходящего платежа.
	// Принимает контекст, исходящий платеж и опционально объект параметров запроса Params.
	// Возвращает изменённый исходящий платеж.
	Update(ctx context.Context, id uuid.UUID, paymentOut *PaymentOut, params ...*Params) (*PaymentOut, *resty.Response, error)

	// Template выполняет запрос на получение предзаполненного исходящего платежа со стандартными полями без связи с какими-либо другими документами.
	// Принимает контекст.
	// Возвращает предзаполненный исходящий платеж.
	Template(ctx context.Context) (*PaymentOut, *resty.Response, error)

	// TemplateBased выполняет запрос на получение шаблона исходящего платежа на основе других документов.
	// Основание, на котором может быть создан:
	//	- Возврат покупателя (SalesReturn)
	//	- Приемка (Supply)
	//	- Счет поставщика (InvoiceIn)
	//	- Заказ поставщику (PurchaseOrder)
	//	- Выданный отчет комиссионера (CommissionReportOut)
	// Принимает контекст и множество документов из списка выше.
	// Возвращает предзаполненный исходящий платеж на основании переданных документов.
	TemplateBased(ctx context.Context, basedOn ...MetaOwner) (*PaymentOut, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных исходящих платежей.
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
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*PaymentOut, *resty.Response, error)

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

// NewPaymentOutService принимает [Client] и возвращает сервис для работы с исходящими платежами.
func NewPaymentOutService(client *Client) PaymentOutService {
	return newMainService[PaymentOut, any, MetaAttributesStatesSharedWrapper, any](NewEndpoint(client, "entity/paymentout"))
}
