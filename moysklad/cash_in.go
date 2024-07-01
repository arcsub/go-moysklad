package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// CashIn Приходный ордер.
//
// Код сущности: cashin
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-prihodnyj-order
type CashIn struct {
	Organization   *Organization            `json:"organization,omitempty"`   // Метаданные юрлица
	VatSum         *float64                 `json:"vatSum,omitempty"`         // Сумма НДС
	Applicable     *bool                    `json:"applicable,omitempty"`     // Отметка о проведении
	Moment         *Timestamp               `json:"moment,omitempty"`         // Дата документа
	Code           *string                  `json:"code,omitempty"`           // Код Приходного ордера
	Contract       *NullValue[Contract]     `json:"contract,omitempty"`       // Метаданные договора
	AccountID      *uuid.UUID               `json:"accountId,omitempty"`      // ID учётной записи      // ID учётной записи
	Deleted        *Timestamp               `json:"deleted,omitempty"`        // Момент последнего удаления Приходного ордера
	Description    *string                  `json:"description,omitempty"`    // Комментарий Приходного ордера
	ExternalCode   *string                  `json:"externalCode,omitempty"`   // Внешний код Приходного ордера
	Files          *MetaArray[File]         `json:"files,omitempty"`          // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group          *Group                   `json:"group,omitempty"`          // Отдел сотрудника
	ID             *uuid.UUID               `json:"id,omitempty"`             // ID Приходного ордера
	Meta           *Meta                    `json:"meta,omitempty"`           // Метаданные Приходного ордера
	Operations     Slice[Operation]         `json:"operations,omitempty"`     // Метаданные связанных операций
	Agent          *Agent                   `json:"agent,omitempty"`          // Метаданные контрагента
	Created        *Timestamp               `json:"created,omitempty"`        // Дата создания
	Owner          *Employee                `json:"owner,omitempty"`          // Метаданные владельца (Сотрудника)
	PaymentPurpose *string                  `json:"paymentPurpose,omitempty"` // Основание
	Printed        *bool                    `json:"printed,omitempty"`        // Напечатан ли документ
	Project        *NullValue[Project]      `json:"project,omitempty"`        // Метаданные проекта
	Published      *bool                    `json:"published,omitempty"`      // Опубликован ли документ
	Rate           *NullValue[Rate]         `json:"rate,omitempty"`           // Валюта
	SalesChannel   *NullValue[SalesChannel] `json:"salesChannel,omitempty"`   // Метаданные канала продаж
	Shared         *bool                    `json:"shared,omitempty"`         // Общий доступ
	State          *NullValue[State]        `json:"state,omitempty"`          // Метаданные статуса Приходного ордера
	Sum            *float64                 `json:"sum,omitempty"`            // Сумма Приходного ордера в установленной валюте
	SyncID         *uuid.UUID               `json:"syncId,omitempty"`         // ID синхронизации
	Updated        *Timestamp               `json:"updated,omitempty"`        // Момент последнего обновления Приходного ордера
	Name           *string                  `json:"name,omitempty"`           // Наименование Приходного ордера
	FactureIn      *FactureIn               `json:"factureIn,omitempty"`      // Метаданные Счет-фактуры полученного
	Attributes     Slice[Attribute]         `json:"attributes,omitempty"`     // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (cashIn CashIn) Clean() *CashIn {
	if cashIn.Meta == nil {
		return nil
	}
	return &CashIn{Meta: cashIn.Meta}
}

// asTaskOperation реализует интерфейс [AsTaskOperationInterface].
func (cashIn CashIn) asTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: cashIn.Meta}
}

// asPayment реализует интерфейс AsPaymentInterface.
func (cashIn CashIn) asPayment() *Payment {
	return &Payment{Meta: cashIn.GetMeta()}
}

// GetOrganization возвращает Метаданные юрлица.
func (cashIn CashIn) GetOrganization() Organization {
	return Deref(cashIn.Organization)
}

// GetVatSum возвращает Сумму НДС.
func (cashIn CashIn) GetVatSum() float64 {
	return Deref(cashIn.VatSum)
}

// GetApplicable возвращает Отметку о проведении.
func (cashIn CashIn) GetApplicable() bool {
	return Deref(cashIn.Applicable)
}

// GetMoment возвращает Дату документа.
func (cashIn CashIn) GetMoment() Timestamp {
	return Deref(cashIn.Moment)
}

// GetCode возвращает Код Приходного ордера.
func (cashIn CashIn) GetCode() string {
	return Deref(cashIn.Code)
}

// GetContract возвращает Метаданные договора.
func (cashIn CashIn) GetContract() Contract {
	return cashIn.Contract.Get()
}

// GetAccountID возвращает ID учётной записи.
func (cashIn CashIn) GetAccountID() uuid.UUID {
	return Deref(cashIn.AccountID)
}

// GetDeleted возвращает Момент последнего удаления Приходного ордера.
func (cashIn CashIn) GetDeleted() Timestamp {
	return Deref(cashIn.Deleted)
}

// GetDescription возвращает Комментарий Приходного ордера.
func (cashIn CashIn) GetDescription() string {
	return Deref(cashIn.Description)
}

// GetExternalCode возвращает Внешний код Приходного ордера.
func (cashIn CashIn) GetExternalCode() string {
	return Deref(cashIn.ExternalCode)
}

// GetFiles возвращает Метаданные массива Файлов.
func (cashIn CashIn) GetFiles() MetaArray[File] {
	return Deref(cashIn.Files)
}

// GetGroup возвращает Отдел сотрудника.
func (cashIn CashIn) GetGroup() Group {
	return Deref(cashIn.Group)
}

// GetID возвращает ID Приходного ордера.
func (cashIn CashIn) GetID() uuid.UUID {
	return Deref(cashIn.ID)
}

// GetMeta возвращает Метаданные Приходного ордера.
func (cashIn CashIn) GetMeta() Meta {
	return Deref(cashIn.Meta)
}

// GetOperations возвращает Метаданные связанных операций.
func (cashIn CashIn) GetOperations() Slice[Operation] {
	return cashIn.Operations
}

// GetAgent возвращает Метаданные контрагента.
func (cashIn CashIn) GetAgent() Agent {
	return Deref(cashIn.Agent)
}

// GetCreated возвращает Дату создания.
func (cashIn CashIn) GetCreated() Timestamp {
	return Deref(cashIn.Created)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (cashIn CashIn) GetOwner() Employee {
	return Deref(cashIn.Owner)
}

// GetPaymentPurpose возвращает Основание.
func (cashIn CashIn) GetPaymentPurpose() string {
	return Deref(cashIn.PaymentPurpose)
}

// GetPrinted возвращает true, если документ напечатан.
func (cashIn CashIn) GetPrinted() bool {
	return Deref(cashIn.Printed)
}

// GetProject возвращает Метаданные проекта.
func (cashIn CashIn) GetProject() Project {
	return cashIn.Project.Get()
}

// GetPublished возвращает true, если документ опубликован.
func (cashIn CashIn) GetPublished() bool {
	return Deref(cashIn.Published)
}

// GetRate возвращает Валюту.
func (cashIn CashIn) GetRate() Rate {
	return cashIn.Rate.Get()
}

// GetSalesChannel возвращает Метаданные канала продаж.
func (cashIn CashIn) GetSalesChannel() SalesChannel {
	return cashIn.SalesChannel.Get()
}

// GetShared возвращает флаг общего доступа.
func (cashIn CashIn) GetShared() bool {
	return Deref(cashIn.Shared)
}

// GetState возвращает Метаданные статуса Приходного ордера.
func (cashIn CashIn) GetState() State {
	return cashIn.State.Get()
}

// GetSum возвращает Сумму Приходного ордера в установленной валюте.
func (cashIn CashIn) GetSum() float64 {
	return Deref(cashIn.Sum)
}

// GetSyncID возвращает ID синхронизации.
func (cashIn CashIn) GetSyncID() uuid.UUID {
	return Deref(cashIn.SyncID)
}

// GetUpdated возвращает Момент последнего обновления Приходного ордера.
func (cashIn CashIn) GetUpdated() Timestamp {
	return Deref(cashIn.Updated)
}

// GetName возвращает Наименование Приходного ордера.
func (cashIn CashIn) GetName() string {
	return Deref(cashIn.Name)
}

// GetFactureIn возвращает Метаданные Счет-фактуры полученного.
func (cashIn CashIn) GetFactureIn() FactureIn {
	return Deref(cashIn.FactureIn)
}

// GetAttributes возвращает Список метаданных доп. полей.
func (cashIn CashIn) GetAttributes() Slice[Attribute] {
	return cashIn.Attributes
}

// SetOrganization устанавливает Метаданные юрлица.
func (cashIn *CashIn) SetOrganization(organization *Organization) *CashIn {
	if organization != nil {
		cashIn.Organization = organization.Clean()
	}
	return cashIn
}

// SetVatSum устанавливает Сумму НДС.
func (cashIn *CashIn) SetVatSum(vatSum *float64) *CashIn {
	cashIn.VatSum = vatSum
	return cashIn
}

// SetApplicable устанавливает Отметку о проведении.
func (cashIn *CashIn) SetApplicable(applicable bool) *CashIn {
	cashIn.Applicable = &applicable
	return cashIn
}

// SetMoment устанавливает Дату документа.
func (cashIn *CashIn) SetMoment(moment *Timestamp) *CashIn {
	cashIn.Moment = moment
	return cashIn
}

// SetCode устанавливает Код Приходного ордера.
func (cashIn *CashIn) SetCode(code string) *CashIn {
	cashIn.Code = &code
	return cashIn
}

// SetContract устанавливает Метаданные договора.
//
// Передача nil передаёт сброс значения (null).
func (cashIn *CashIn) SetContract(contract *Contract) *CashIn {
	cashIn.Contract = NewNullValue(contract)
	return cashIn
}

// SetDescription устанавливает Комментарий Приходного ордера.
func (cashIn *CashIn) SetDescription(description string) *CashIn {
	cashIn.Description = &description
	return cashIn
}

// SetExternalCode устанавливает Внешний код Приходного ордера.
func (cashIn *CashIn) SetExternalCode(externalCode string) *CashIn {
	cashIn.ExternalCode = &externalCode
	return cashIn
}

// SetFiles устанавливает Метаданные массива Файлов.
//
// Принимает множество объектов [File].
func (cashIn *CashIn) SetFiles(files ...*File) *CashIn {
	cashIn.Files = NewMetaArrayFrom(files)
	return cashIn
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (cashIn *CashIn) SetGroup(group *Group) *CashIn {
	if group != nil {
		cashIn.Group = group.Clean()
	}
	return cashIn
}

// SetMeta устанавливает Метаданные Приходного ордера.
func (cashIn *CashIn) SetMeta(meta *Meta) *CashIn {
	cashIn.Meta = meta
	return cashIn
}

// SetOperations устанавливает Метаданные связанных операций.
func (cashIn *CashIn) SetOperations(operations ...AsOperationInterface) *CashIn {
	cashIn.Operations = NewOperationsFrom(operations)
	return cashIn
}

// SetAgent устанавливает Метаданные контрагента.
//
// Принимает [Counterparty], [Organization] или [Employee].
func (cashIn *CashIn) SetAgent(agent AgentInterface) *CashIn {
	if agent != nil {
		cashIn.Agent = agent.asAgent()
	}
	return cashIn
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (cashIn *CashIn) SetOwner(owner *Employee) *CashIn {
	if owner != nil {
		cashIn.Owner = owner.Clean()
	}
	return cashIn
}

// SetPaymentPurpose устанавливает Основание.
func (cashIn *CashIn) SetPaymentPurpose(paymentPurpose string) *CashIn {
	cashIn.PaymentPurpose = &paymentPurpose
	return cashIn
}

// SetProject устанавливает Метаданные проекта.
//
// Передача nil передаёт сброс значения (null).
func (cashIn *CashIn) SetProject(project *Project) *CashIn {
	cashIn.Project = NewNullValue(project)
	return cashIn
}

// SetRate устанавливает Валюту.
//
// Передача nil передаёт сброс значения (null).
func (cashIn *CashIn) SetRate(rate *Rate) *CashIn {
	cashIn.Rate = NewNullValue(rate)
	return cashIn
}

// SetSalesChannel устанавливает Метаданные канала продаж.
//
// Передача nil передаёт сброс значения (null).
func (cashIn *CashIn) SetSalesChannel(salesChannel *SalesChannel) *CashIn {
	cashIn.SalesChannel = NewNullValue(salesChannel)
	return cashIn
}

// SetShared устанавливает флаг общего доступа.
func (cashIn *CashIn) SetShared(shared bool) *CashIn {
	cashIn.Shared = &shared
	return cashIn
}

// SetState устанавливает Метаданные статуса Приходного ордера.
//
// Передача nil передаёт сброс значения (null).
func (cashIn *CashIn) SetState(state *State) *CashIn {
	cashIn.State = NewNullValue(state)
	return cashIn
}

// SetSum устанавливает Сумму Приходного ордера в установленной валюте.
func (cashIn *CashIn) SetSum(sum *float64) *CashIn {
	cashIn.Sum = sum
	return cashIn
}

// SetSyncID устанавливает ID синхронизации.
func (cashIn *CashIn) SetSyncID(syncID uuid.UUID) *CashIn {
	cashIn.SyncID = &syncID
	return cashIn
}

// SetName устанавливает Наименование Приходного ордера.
func (cashIn *CashIn) SetName(name string) *CashIn {
	cashIn.Name = &name
	return cashIn
}

// SetFactureIn устанавливает Метаданные Счет-фактуры полученного.
func (cashIn *CashIn) SetFactureIn(factureIn *FactureIn) *CashIn {
	if factureIn != nil {
		cashIn.FactureIn = factureIn.Clean()
	}
	return cashIn
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (cashIn *CashIn) SetAttributes(attributes ...*Attribute) *CashIn {
	cashIn.Attributes.Push(attributes...)
	return cashIn
}

// String реализует интерфейс [fmt.Stringer].
func (cashIn CashIn) String() string {
	return Stringify(cashIn)
}

// MetaType возвращает код сущности.
func (CashIn) MetaType() MetaType {
	return MetaTypeCashIn
}

// AsOperation возвращает объект [Operation] c полями meta и linkedSum.
// Значение поля linkedSum заполняется из поля sum.
func (cashIn CashIn) AsOperation() *Operation {
	return &Operation{Meta: cashIn.GetMeta(), LinkedSum: cashIn.GetSum()}
}

// Update shortcut
func (cashIn CashIn) Update(ctx context.Context, client *Client, params ...*Params) (*CashIn, *resty.Response, error) {
	return NewCashInService(client).Update(ctx, cashIn.GetID(), &cashIn, params...)
}

// Create shortcut
func (cashIn CashIn) Create(ctx context.Context, client *Client, params ...*Params) (*CashIn, *resty.Response, error) {
	return NewCashInService(client).Create(ctx, &cashIn, params...)
}

// Delete shortcut
func (cashIn CashIn) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewCashInService(client).Delete(ctx, cashIn.GetID())
}

// CashInService методы сервиса для работы с приходными ордерами.
type CashInService interface {
	// GetList выполняет запрос на получение списка приходных ордеров.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[CashIn], *resty.Response, error)

	// Create выполняет запрос на создание приходного ордера.
	// Обязательные поля для заполнения:
	//	- organization (Метаданные юрлица)
	//	- agent (Метаданные контрагента)
	// Принимает контекст, приходный ордер и опционально объект параметров запроса Params.
	// Возвращает созданный приходный ордер.
	Create(ctx context.Context, cashIn *CashIn, params ...*Params) (*CashIn, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение приходных ордеров.
	// Изменяемые приходные ордеры должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список приходных ордеров и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых приходных ордеров.
	CreateUpdateMany(ctx context.Context, cashInList Slice[CashIn], params ...*Params) (*Slice[CashIn], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление приходных ордеров.
	// Принимает контекст и множество приходных ордеров.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*CashIn) (*DeleteManyResponse, *resty.Response, error)

	// Delete выполняет запрос на удаление приходного ордера.
	// Принимает контекст и ID приходного ордера.
	// Возвращает true в случае успешного удаления приходного ордера.
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных приходных ордеров.
	// Принимает контекст.
	// Возвращает объект метаданных MetaAttributesSharedStatesWrapper.
	GetMetadata(ctx context.Context) (*MetaAttributesSharedStatesWrapper, *resty.Response, error)

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

	// Template выполняет запрос на получение предзаполненного приходного ордера со стандартными полями
	// без связи с какими-либо другими документами.
	// Принимает контекст.
	// Возвращает предзаполненный приходный ордер.
	Template(ctx context.Context) (*CashIn, *resty.Response, error)

	// TemplateBased выполняет запрос на получение шаблона приходного ордера на основе других документов.
	// Основание, на котором может быть создан:
	//	- Заказ покупателя (CustomerOrder)
	//	- Возврат поставщику (PurchaseReturn)
	//	- Отгрузка (Demand)
	//	- Счет покупателю (InvoiceOut)
	//	- Полученный отчёт комиссионера (CommissionReportIn)
	// Принимает контекст и множество документов из списка выше.
	// Возвращает предзаполненный приходный ордер на основании переданных документов.
	TemplateBased(ctx context.Context, basedOn ...MetaOwner) (*CashIn, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельного приходного ордера по ID.
	// Принимает контекст, ID приходного ордера и опционально объект параметров запроса Params.
	// Возвращает найденный приходный ордер.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*CashIn, *resty.Response, error)

	// Update выполняет запрос на изменение приходного ордера.
	// Принимает контекст, приходный ордер и опционально объект параметров запроса Params.
	// Возвращает изменённый приходный ордер.
	Update(ctx context.Context, id uuid.UUID, cashIn *CashIn, params ...*Params) (*CashIn, *resty.Response, error)

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
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*CashIn, *resty.Response, error)

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

// NewCashInService принимает [Client] и возвращает сервис для работы с приходными ордерами.
func NewCashInService(client *Client) CashInService {
	return newMainService[CashIn, any, MetaAttributesSharedStatesWrapper, any](NewEndpoint(client, "entity/cashin"))
}
