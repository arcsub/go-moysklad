package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// CashOut Расходный ордер.
//
// Код сущности: cashout
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-rashodnyj-order
type CashOut struct {
	Name           *string                  `json:"name,omitempty"`           // Наименование Расходного ордера
	Deleted        *Timestamp               `json:"deleted,omitempty"`        // Момент последнего удаления Расходного ордера
	Applicable     *bool                    `json:"applicable,omitempty"`     // Отметка о проведении
	AccountID      *uuid.UUID               `json:"accountId,omitempty"`      // ID учётной записи
	Code           *string                  `json:"code,omitempty"`           // Код Расходного ордера
	Contract       *NullValue[Contract]     `json:"contract,omitempty"`       // Метаданные договора
	Created        *Timestamp               `json:"created,omitempty"`        // Дата создания
	Organization   *Organization            `json:"organization,omitempty"`   // Метаданные юрлица
	Description    *string                  `json:"description,omitempty"`    // Комментарий Расходного ордера
	ExpenseItem    *ExpenseItem             `json:"expenseItem,omitempty"`    // Метаданные Статьи расходов
	ExternalCode   *string                  `json:"externalCode,omitempty"`   // Внешний код Расходного ордера
	Files          *MetaArray[File]         `json:"files,omitempty"`          // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group          *Group                   `json:"group,omitempty"`          // Отдел сотрудника
	Owner          *Employee                `json:"owner,omitempty"`          // Метаданные владельца (Сотрудника)
	Meta           *Meta                    `json:"meta,omitempty"`           // Метаданные Расходного ордера
	Moment         *Timestamp               `json:"moment,omitempty"`         // Дата документа
	Operations     Slice[Operation]         `json:"operations,omitempty"`     // Массив ссылок на связанные операции
	Agent          *Counterparty            `json:"agent,omitempty"`          // Метаданные контрагента
	ID             *uuid.UUID               `json:"id,omitempty"`             // ID Расходного ордера
	PaymentPurpose *string                  `json:"paymentPurpose,omitempty"` // Основание
	Printed        *bool                    `json:"printed,omitempty"`        // Напечатан ли документ
	Project        *NullValue[Project]      `json:"project,omitempty"`        // Метаданные проекта
	Published      *bool                    `json:"published,omitempty"`      // Опубликован ли документ
	Rate           *NullValue[Rate]         `json:"rate,omitempty"`           // Валюта
	SalesChannel   *NullValue[SalesChannel] `json:"salesChannel,omitempty"`   // Метаданные канала продаж
	Shared         *bool                    `json:"shared,omitempty"`         // Общий доступ
	State          *NullValue[State]        `json:"state,omitempty"`          // Метаданные статуса Расходного ордера
	Sum            *float64                 `json:"sum,omitempty"`            // Сумма расходного ордера в установленной валюте
	SyncID         *uuid.UUID               `json:"syncId,omitempty"`         // ID синхронизации
	Updated        *Timestamp               `json:"updated,omitempty"`        // Момент последнего обновления Расходного ордера
	VatSum         *float64                 `json:"vatSum,omitempty"`         // Сумма НДС
	FactureOut     *FactureOut              `json:"factureOut,omitempty"`     // Ссылка на Счет-фактуру полученный, с которым связан этот платеж
	Attributes     Slice[Attribute]         `json:"attributes,omitempty"`     // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (cashOut CashOut) Clean() *CashOut {
	if cashOut.Meta == nil {
		return nil
	}
	return &CashOut{Meta: cashOut.Meta}
}

// AsTaskOperation реализует интерфейс AsTaskOperationInterface.
func (cashOut CashOut) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: cashOut.Meta}
}

// AsPayment реализует интерфейс AsPaymentInterface.
func (cashOut CashOut) AsPayment() *Payment {
	return &Payment{Meta: cashOut.GetMeta()}
}

// GetName возвращает Наименование Расходного ордера.
func (cashOut CashOut) GetName() string {
	return Deref(cashOut.Name)
}

// GetDeleted возвращает Момент последнего удаления Расходного ордера.
func (cashOut CashOut) GetDeleted() Timestamp {
	return Deref(cashOut.Deleted)
}

// GetApplicable возвращает Отметку о проведении.
func (cashOut CashOut) GetApplicable() bool {
	return Deref(cashOut.Applicable)
}

// GetAccountID возвращает ID учётной записи.
func (cashOut CashOut) GetAccountID() uuid.UUID {
	return Deref(cashOut.AccountID)
}

// GetCode возвращает Расходного ордера.
func (cashOut CashOut) GetCode() string {
	return Deref(cashOut.Code)
}

// GetContract возвращает Метаданные договора.
func (cashOut CashOut) GetContract() Contract {
	return cashOut.Contract.Get()
}

// GetCreated возвращает Дату создания.
func (cashOut CashOut) GetCreated() Timestamp {
	return Deref(cashOut.Created)
}

// GetOrganization возвращает Метаданные юрлица.
func (cashOut CashOut) GetOrganization() Organization {
	return Deref(cashOut.Organization)
}

// GetDescription возвращает Комментарий Расходного ордера.
func (cashOut CashOut) GetDescription() string {
	return Deref(cashOut.Description)
}

// GetExpenseItem возвращает Метаданные Статьи расходов.
func (cashOut CashOut) GetExpenseItem() ExpenseItem {
	return Deref(cashOut.ExpenseItem)
}

// GetExternalCode возвращает Внешний код Расходного ордера.
func (cashOut CashOut) GetExternalCode() string {
	return Deref(cashOut.ExternalCode)
}

// GetFiles возвращает Метаданные массива Файлов.
func (cashOut CashOut) GetFiles() MetaArray[File] {
	return Deref(cashOut.Files)
}

// GetGroup возвращает Отдел сотрудника.
func (cashOut CashOut) GetGroup() Group {
	return Deref(cashOut.Group)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (cashOut CashOut) GetOwner() Employee {
	return Deref(cashOut.Owner)
}

// GetMeta возвращает Метаданные Расходного ордера.
func (cashOut CashOut) GetMeta() Meta {
	return Deref(cashOut.Meta)
}

// GetMoment возвращает Дату документа.
func (cashOut CashOut) GetMoment() Timestamp {
	return Deref(cashOut.Moment)
}

// GetOperations возвращает Метаданные связанных операций.
func (cashOut CashOut) GetOperations() Slice[Operation] {
	return cashOut.Operations
}

// GetAgent возвращает Метаданные контрагента.
func (cashOut CashOut) GetAgent() Counterparty {
	return Deref(cashOut.Agent)
}

// GetID возвращает ID Расходного ордера.
func (cashOut CashOut) GetID() uuid.UUID {
	return Deref(cashOut.ID)
}

// GetPaymentPurpose возвращает Основание.
func (cashOut CashOut) GetPaymentPurpose() string {
	return Deref(cashOut.PaymentPurpose)
}

// GetPrinted возвращает true, если документ напечатан.
func (cashOut CashOut) GetPrinted() bool {
	return Deref(cashOut.Printed)
}

// GetProject возвращает Метаданные проекта.
func (cashOut CashOut) GetProject() Project {
	return cashOut.Project.Get()
}

// GetPublished возвращает true, если документ опубликован.
func (cashOut CashOut) GetPublished() bool {
	return Deref(cashOut.Published)
}

// GetRate возвращает Валюту.
func (cashOut CashOut) GetRate() Rate {
	return cashOut.Rate.Get()
}

// GetSalesChannel возвращает Метаданные канала продаж.
func (cashOut CashOut) GetSalesChannel() SalesChannel {
	return cashOut.SalesChannel.Get()
}

// GetShared возвращает флаг общего доступа.
func (cashOut CashOut) GetShared() bool {
	return Deref(cashOut.Shared)
}

// GetState возвращает Метаданные статуса Расходного ордера.
func (cashOut CashOut) GetState() State {
	return cashOut.State.Get()
}

// GetSum возвращает Сумму Расходного ордера в установленной валюте.
func (cashOut CashOut) GetSum() float64 {
	return Deref(cashOut.Sum)
}

// GetSyncID возвращает ID синхронизации.
func (cashOut CashOut) GetSyncID() uuid.UUID {
	return Deref(cashOut.SyncID)
}

// GetUpdated возвращает Момент последнего обновления Расходного ордера.
func (cashOut CashOut) GetUpdated() Timestamp {
	return Deref(cashOut.Updated)
}

// GetVatSum возвращает Сумму НДС.
func (cashOut CashOut) GetVatSum() float64 {
	return Deref(cashOut.VatSum)
}

// GetFactureOut возвращает Ссылку на Счет-фактуру полученный, с которым связан этот платеж.
func (cashOut CashOut) GetFactureOut() FactureOut {
	return Deref(cashOut.FactureOut)
}

// GetAttributes возвращает Список метаданных доп. полей.
func (cashOut CashOut) GetAttributes() Slice[Attribute] {
	return cashOut.Attributes
}

// SetName устанавливает Наименование Расходного ордера.
func (cashOut *CashOut) SetName(name string) *CashOut {
	cashOut.Name = &name
	return cashOut
}

// SetApplicable устанавливает Отметку о проведении.
func (cashOut *CashOut) SetApplicable(applicable bool) *CashOut {
	cashOut.Applicable = &applicable
	return cashOut
}

// SetCode устанавливает Код Расходного ордера.
func (cashOut *CashOut) SetCode(code string) *CashOut {
	cashOut.Code = &code
	return cashOut
}

// SetContract устанавливает Метаданные договора.
//
// Передача nil устанавливает необходимость сброса значения (передача null).
func (cashOut *CashOut) SetContract(contract *Contract) *CashOut {
	if contract == nil {
		cashOut.Contract = NewNullValue[Contract]()
	} else {
		cashOut.Contract = NewNullValueFrom(contract.Clean())
	}
	return cashOut
}

// SetOrganization устанавливает Метаданные юрлица.
func (cashOut *CashOut) SetOrganization(organization *Organization) *CashOut {
	if organization != nil {
		cashOut.Organization = organization.Clean()
	}
	return cashOut
}

// SetDescription устанавливает Комментарий Расходного ордера.
func (cashOut *CashOut) SetDescription(description string) *CashOut {
	cashOut.Description = &description
	return cashOut
}

// SetExpenseItem устанавливает Метаданные Статьи расходов.
func (cashOut *CashOut) SetExpenseItem(expenseItem *ExpenseItem) *CashOut {
	if expenseItem != nil {
		cashOut.ExpenseItem = expenseItem.Clean()
	}
	return cashOut
}

// SetExternalCode устанавливает Внешний код Расходного ордера.
func (cashOut *CashOut) SetExternalCode(externalCode string) *CashOut {
	cashOut.ExternalCode = &externalCode
	return cashOut
}

// SetFiles устанавливает Метаданные массива Файлов.
func (cashOut *CashOut) SetFiles(files ...*File) *CashOut {
	cashOut.Files = NewMetaArrayFrom(files)
	return cashOut
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (cashOut *CashOut) SetGroup(group *Group) *CashOut {
	if group != nil {
		cashOut.Group = group.Clean()
	}
	return cashOut
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (cashOut *CashOut) SetOwner(owner *Employee) *CashOut {
	if owner != nil {
		cashOut.Owner = owner.Clean()
	}
	return cashOut
}

// SetMeta устанавливает Метаданные Расходного ордера.
func (cashOut *CashOut) SetMeta(meta *Meta) *CashOut {
	cashOut.Meta = meta
	return cashOut
}

// SetMoment устанавливает Дату документа.
func (cashOut *CashOut) SetMoment(moment *Timestamp) *CashOut {
	cashOut.Moment = moment
	return cashOut
}

// SetOperations устанавливает Метаданные связанных операций.
func (cashOut *CashOut) SetOperations(operations ...AsOperationInterface) *CashOut {
	cashOut.Operations = NewOperationsFrom(operations)
	return cashOut
}

// SetAgent устанавливает Метаданные контрагента.
func (cashOut *CashOut) SetAgent(agent *Counterparty) *CashOut {
	if agent != nil {
		cashOut.Agent = agent.Clean()
	}
	return cashOut
}

// SetPaymentPurpose устанавливает Основание.
func (cashOut *CashOut) SetPaymentPurpose(paymentPurpose string) *CashOut {
	cashOut.PaymentPurpose = &paymentPurpose
	return cashOut
}

// SetProject устанавливает Метаданные проекта.
//
// Передача nil устанавливает необходимость сброса значения (передача null).
func (cashOut *CashOut) SetProject(project *Project) *CashOut {
	if project == nil {
		cashOut.Project = NewNullValue[Project]()
	} else {
		cashOut.Project = NewNullValueFrom(project.Clean())
	}
	return cashOut
}

// SetRate устанавливает Валюту.
//
// Передача nil устанавливает необходимость сброса значения (передача null).
func (cashOut *CashOut) SetRate(rate *Rate) *CashOut {
	if rate == nil {
		NewNullValue[Rate]()
	} else {
		cashOut.Rate = NewNullValueFrom(rate)
	}
	return cashOut
}

// SetSalesChannel устанавливает Метаданные канала продаж.
//
// Передача nil устанавливает необходимость сброса значения (передача null).
func (cashOut *CashOut) SetSalesChannel(salesChannel *SalesChannel) *CashOut {
	if salesChannel == nil {
		cashOut.SalesChannel = NewNullValue[SalesChannel]()
	} else {
		cashOut.SalesChannel = NewNullValueFrom(salesChannel.Clean())
	}
	return cashOut
}

// SetShared устанавливает флаг общего доступа.
func (cashOut *CashOut) SetShared(shared bool) *CashOut {
	cashOut.Shared = &shared
	return cashOut
}

// SetState устанавливает Метаданные статуса Расходного ордера.
//
// Передача nil устанавливает необходимость сброса значения (передача null).
func (cashOut *CashOut) SetState(state *State) *CashOut {
	if state == nil {
		cashOut.State = NewNullValue[State]()
	} else {
		cashOut.State = NewNullValueFrom(state.Clean())
	}
	return cashOut
}

// SetSum устанавливает Сумму Расходного ордера в установленной валюте.
func (cashOut *CashOut) SetSum(sum float64) *CashOut {
	cashOut.Sum = &sum
	return cashOut
}

// SetSyncID устанавливает ID синхронизации.
func (cashOut *CashOut) SetSyncID(syncID uuid.UUID) *CashOut {
	cashOut.SyncID = &syncID
	return cashOut
}

// SetVatSum устанавливает Сумму НДС.
func (cashOut *CashOut) SetVatSum(vatSum float64) *CashOut {
	cashOut.VatSum = &vatSum
	return cashOut
}

// SetFactureOut устанавливает Метаданные Счет-фактуры выданного.
func (cashOut *CashOut) SetFactureOut(factureOut *FactureOut) *CashOut {
	if factureOut != nil {
		cashOut.FactureOut = factureOut.Clean()
	}
	return cashOut
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (cashOut *CashOut) SetAttributes(attributes ...*Attribute) *CashOut {
	cashOut.Attributes = attributes
	return cashOut
}

// String реализует интерфейс [fmt.Stringer].
func (cashOut CashOut) String() string {
	return Stringify(cashOut)
}

// MetaType возвращает код сущности.
func (CashOut) MetaType() MetaType {
	return MetaTypeCashOut
}

// AsOperation возвращает объект [Operation] c полями meta и linkedSum.
// Значение поля linkedSum заполняется из поля sum.
func (cashOut CashOut) AsOperation() *Operation {
	return &Operation{Meta: cashOut.GetMeta(), LinkedSum: cashOut.GetSum()}
}

// Update shortcut
func (cashOut CashOut) Update(ctx context.Context, client *Client, params ...*Params) (*CashOut, *resty.Response, error) {
	return NewCashOutService(client).Update(ctx, cashOut.GetID(), &cashOut, params...)
}

// Create shortcut
func (cashOut CashOut) Create(ctx context.Context, client *Client, params ...*Params) (*CashOut, *resty.Response, error) {
	return NewCashOutService(client).Create(ctx, &cashOut, params...)
}

// Delete shortcut
func (cashOut CashOut) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewCashOutService(client).Delete(ctx, cashOut.GetID())
}

// CashOutService методы сервиса для работы с расходными ордерами.
type CashOutService interface {
	// GetList выполняет запрос на получение списка расходных ордеров.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[CashOut], *resty.Response, error)

	// Create выполняет запрос на создание расходного ордера.
	// Обязательные поля для заполнения:
	//	- organization (Метаданные юрлица)
	//	- agent (Метаданные контрагента)
	//	- expenseItem (Метаданные Статьи расходов)
	// Принимает контекст, расходный ордер и опционально объект параметров запроса Params.
	// Возвращает созданный расходный ордер.
	Create(ctx context.Context, cashOut *CashOut, params ...*Params) (*CashOut, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение расходных ордеров.
	// Изменяемые приходные ордеры должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список приходных ордеров и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или расходных приходных ордеров.
	CreateUpdateMany(ctx context.Context, cashOutList Slice[CashOut], params ...*Params) (*Slice[CashOut], *resty.Response, error)

	// Delete выполняет запрос на удаление расходного ордера.
	// Принимает контекст и ID расходного ордера.
	// Возвращает true в случае успешного удаления расходного ордера.
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление расходных ордеров.
	// Принимает контекст и множество расходных ордеров.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*CashOut) (*DeleteManyResponse, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных расходных ордеров.
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

	// Template выполняет запрос на получение предзаполненного расходного ордера со стандартными полями
	// без связи с какими-либо другими документами.
	// Принимает контекст.
	// Возвращает предзаполненный расходный ордер.
	Template(ctx context.Context) (*CashOut, *resty.Response, error)

	// TemplateBased выполняет запрос на получение шаблона расходного ордера на основе других документов.
	// Основание, на котором может быть создан:
	//	- Возврат покупателя (SalesReturn)
	//	- Приемка (Supply)
	//	- Счет поставщика (InvoiceIn)
	//	- Заказ поставщику (PurchaseOrder)
	//	- Выданный отчёт комиссионера (CommissionReportOut)
	// Принимает контекст и множество документов из списка выше.
	// Возвращает предзаполненный расходный ордер на основании переданных документов.
	TemplateBased(ctx context.Context, basedOn ...MetaOwner) (*CashOut, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельного расходного ордера по ID.
	// Принимает контекст, ID расходный ордера и опционально объект параметров запроса Params.
	// Возвращает найденный расходный ордер.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*CashOut, *resty.Response, error)

	// Update выполняет запрос на изменение расходного ордера.
	// Принимает контекст, расходный ордер и опционально объект параметров запроса Params.
	// Возвращает изменённый расходный ордер.
	Update(ctx context.Context, id uuid.UUID, cashOut *CashOut, params ...*Params) (*CashOut, *resty.Response, error)

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

	// GetBySyncID выполняет запрос на получение документа по syncID.
	// Принимает контекст и syncID документа.
	// Возвращает найденный документ.
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*CashOut, *resty.Response, error)

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
	// Возвращает объект MetaArray.
	GetFileList(ctx context.Context, id uuid.UUID) (*MetaArray[File], *resty.Response, error)

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

// NewCashOutService принимает [Client] и возвращает сервис для работы с расходными ордерами.
func NewCashOutService(client *Client) CashOutService {
	return newMainService[CashOut, any, MetaAttributesSharedStatesWrapper, any](NewEndpoint(client, "entity/cashout"))
}
