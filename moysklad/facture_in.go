package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// FactureIn Счет-фактура полученный.
//
// Код сущности: facturein
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-schet-faktura-poluchennyj
type FactureIn struct {
	Moment         *Timestamp           `json:"moment,omitempty"`         // Дата документа
	Applicable     *bool                `json:"applicable,omitempty"`     // Отметка о проведении
	Name           *string              `json:"name,omitempty"`           // Наименование полученного счета-фактуры
	AccountID      *uuid.UUID           `json:"accountId,omitempty"`      // ID учётной записи
	Code           *string              `json:"code,omitempty"`           // Код полученного счета-фактуры
	Contract       *NullValue[Contract] `json:"contract,omitempty"`       // Метаданные договора
	Created        *Timestamp           `json:"created,omitempty"`        // Дата создания
	Deleted        *Timestamp           `json:"deleted,omitempty"`        // Момент последнего удаления полученного счета-фактуры
	Description    *string              `json:"description,omitempty"`    // Комментарий полученного счета-фактуры
	ExternalCode   *string              `json:"externalCode,omitempty"`   // Внешний код полученного счета-фактуры
	Files          *MetaArray[File]     `json:"files,omitempty"`          // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group          *Group               `json:"group,omitempty"`          // Отдел сотрудника
	ID             *uuid.UUID           `json:"id,omitempty"`             // ID полученного счета-фактуры
	Meta           *Meta                `json:"meta,omitempty"`           // Метаданные полученного счета-фактуры
	IncomingDate   *Timestamp           `json:"incomingDate,omitempty"`   // Входящая дата
	Agent          *Agent               `json:"agent,omitempty"`          // Метаданные контрагента
	Organization   *Organization        `json:"organization,omitempty"`   // Метаданные юрлица
	Owner          *Employee            `json:"owner,omitempty"`          // Метаданные владельца (Сотрудника)
	Printed        *bool                `json:"printed,omitempty"`        // Напечатан ли документ
	Published      *bool                `json:"published,omitempty"`      // Опубликован ли документ
	Rate           *NullValue[Rate]     `json:"rate,omitempty"`           // Валюта
	Shared         *bool                `json:"shared,omitempty"`         // Общий доступ
	State          *NullValue[State]    `json:"state,omitempty"`          // Метаданные статуса полученного счета-фактуры
	Sum            *float64             `json:"sum,omitempty"`            // Сумма полученного счета-фактуры в установленной валюте
	SyncID         *uuid.UUID           `json:"syncId,omitempty"`         // ID синхронизации
	Updated        *Timestamp           `json:"updated,omitempty"`        // Момент последнего обновления полученного счета-фактуры
	Supplies       Slice[Supply]        `json:"supplies,omitempty"`       // Массив ссылок на связанные приемки
	Payments       Slice[Payment]       `json:"payments,omitempty"`       // Массив ссылок на связанные исходящие платежи
	IncomingNumber *string              `json:"incomingNumber,omitempty"` // Входящий номер
	Attributes     Slice[Attribute]     `json:"attributes,omitempty"`     // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (factureIn FactureIn) Clean() *FactureIn {
	if factureIn.Meta == nil {
		return nil
	}
	return &FactureIn{Meta: factureIn.Meta}
}

// AsOperation возвращает объект [Operation] c полями meta и linkedSum.
//
// Значение поля linkedSum заполняется из поля sum.
func (factureIn FactureIn) AsOperation() *Operation {
	return &Operation{Meta: factureIn.GetMeta(), LinkedSum: factureIn.GetSum()}
}

// asTaskOperation реализует интерфейс [TaskOperationInterface].
func (factureIn FactureIn) asTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: factureIn.Meta}
}

// GetMoment возвращает Дату документа.
func (factureIn FactureIn) GetMoment() Timestamp {
	return Deref(factureIn.Moment)
}

// GetApplicable возвращает Отметку о проведении.
func (factureIn FactureIn) GetApplicable() bool {
	return Deref(factureIn.Applicable)
}

// GetName возвращает Наименование полученного счета-фактуры.
func (factureIn FactureIn) GetName() string {
	return Deref(factureIn.Name)
}

// GetAccountID возвращает ID учётной записи.
func (factureIn FactureIn) GetAccountID() uuid.UUID {
	return Deref(factureIn.AccountID)
}

// GetCode возвращает Код полученного счета-фактуры.
func (factureIn FactureIn) GetCode() string {
	return Deref(factureIn.Code)
}

// GetContract возвращает Метаданные договора.
func (factureIn FactureIn) GetContract() Contract {
	return factureIn.Contract.GetValue()
}

// GetCreated возвращает Дату создания.
func (factureIn FactureIn) GetCreated() Timestamp {
	return Deref(factureIn.Created)
}

// GetDeleted возвращает Момент последнего удаления полученного счета-фактуры.
func (factureIn FactureIn) GetDeleted() Timestamp {
	return Deref(factureIn.Deleted)
}

// GetDescription возвращает Комментарий полученного счета-фактуры.
func (factureIn FactureIn) GetDescription() string {
	return Deref(factureIn.Description)
}

// GetExternalCode возвращает Внешний код полученного счета-фактуры.
func (factureIn FactureIn) GetExternalCode() string {
	return Deref(factureIn.ExternalCode)
}

// GetFiles возвращает Метаданные массива Файлов.
func (factureIn FactureIn) GetFiles() MetaArray[File] {
	return Deref(factureIn.Files)
}

// GetGroup возвращает Отдел сотрудника.
func (factureIn FactureIn) GetGroup() Group {
	return Deref(factureIn.Group)
}

// GetID возвращает ID полученного счета-фактуры.
func (factureIn FactureIn) GetID() uuid.UUID {
	return Deref(factureIn.ID)
}

// GetMeta возвращает Метаданные полученного счета-фактуры.
func (factureIn FactureIn) GetMeta() Meta {
	return Deref(factureIn.Meta)
}

// GetIncomingDate возвращает Входящую дату.
func (factureIn FactureIn) GetIncomingDate() Timestamp {
	return Deref(factureIn.IncomingDate)
}

// GetAgent возвращает Метаданные контрагента.
func (factureIn FactureIn) GetAgent() Agent {
	return Deref(factureIn.Agent)
}

// GetOrganization возвращает Метаданные юрлица.
func (factureIn FactureIn) GetOrganization() Organization {
	return Deref(factureIn.Organization)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (factureIn FactureIn) GetOwner() Employee {
	return Deref(factureIn.Owner)
}

// GetPrinted возвращает true, если документ напечатан.
func (factureIn FactureIn) GetPrinted() bool {
	return Deref(factureIn.Printed)
}

// GetPublished возвращает true, если документ опубликован.
func (factureIn FactureIn) GetPublished() bool {
	return Deref(factureIn.Published)
}

// GetRate возвращает Валюту.
func (factureIn FactureIn) GetRate() Rate {
	return factureIn.Rate.GetValue()
}

// GetShared возвращает флаг общего доступа.
func (factureIn FactureIn) GetShared() bool {
	return Deref(factureIn.Shared)
}

// GetState возвращает Метаданные статуса полученного счета-фактуры.
func (factureIn FactureIn) GetState() State {
	return factureIn.State.GetValue()
}

// GetSum возвращает Сумму полученного счета-фактуры в установленной валюте.
func (factureIn FactureIn) GetSum() float64 {
	return Deref(factureIn.Sum)
}

// GetSyncID возвращает ID синхронизации.
func (factureIn FactureIn) GetSyncID() uuid.UUID {
	return Deref(factureIn.SyncID)
}

// GetUpdated возвращает Момент последнего обновления.
func (factureIn FactureIn) GetUpdated() Timestamp {
	return Deref(factureIn.Updated)
}

// GetSupplies возвращает Массив ссылок на связанные приемки.
func (factureIn FactureIn) GetSupplies() Slice[Supply] {
	return factureIn.Supplies
}

// GetPayments возвращает Массив ссылок на связанные платежи.
func (factureIn FactureIn) GetPayments() Slice[Payment] {
	return factureIn.Payments
}

// GetIncomingNumber возвращает Входящий номер.
func (factureIn FactureIn) GetIncomingNumber() string {
	return Deref(factureIn.IncomingNumber)
}

// GetAttributes возвращает Список метаданных доп. полей.
func (factureIn FactureIn) GetAttributes() Slice[Attribute] {
	return factureIn.Attributes
}

// SetMoment устанавливает Дату документа.
func (factureIn *FactureIn) SetMoment(moment *Timestamp) *FactureIn {
	factureIn.Moment = moment
	return factureIn
}

// SetApplicable устанавливает Отметку о проведении.
func (factureIn *FactureIn) SetApplicable(applicable bool) *FactureIn {
	factureIn.Applicable = &applicable
	return factureIn
}

// SetName устанавливает Наименование полученного счета-фактуры.
func (factureIn *FactureIn) SetName(name string) *FactureIn {
	factureIn.Name = &name
	return factureIn
}

// SetCode устанавливает Код полученного счета-фактуры.
func (factureIn *FactureIn) SetCode(code string) *FactureIn {
	factureIn.Code = &code
	return factureIn
}

// SetContract устанавливает Метаданные договора.
//
// Передача nil передаёт сброс значения (null).
func (factureIn *FactureIn) SetContract(contract *Contract) *FactureIn {
	factureIn.Contract = NewNullValue(contract)
	return factureIn
}

// SetDescription устанавливает Комментарий полученного счета-фактуры.
func (factureIn *FactureIn) SetDescription(description string) *FactureIn {
	factureIn.Description = &description
	return factureIn
}

// SetExternalCode устанавливает Внешний код полученного счета-фактуры.
func (factureIn *FactureIn) SetExternalCode(externalCode string) *FactureIn {
	factureIn.ExternalCode = &externalCode
	return factureIn
}

// SetFiles устанавливает Метаданные массива Файлов.
//
// Принимает множество объектов [File].
func (factureIn *FactureIn) SetFiles(files ...*File) *FactureIn {
	factureIn.Files = NewMetaArrayFrom(files)
	return factureIn
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (factureIn *FactureIn) SetGroup(group *Group) *FactureIn {
	if group != nil {
		factureIn.Group = group.Clean()
	}
	return factureIn
}

// SetMeta устанавливает Метаданные полученного счета-фактуры.
func (factureIn *FactureIn) SetMeta(meta *Meta) *FactureIn {
	factureIn.Meta = meta
	return factureIn
}

// SetIncomingDate устанавливает Входящую дату.
func (factureIn *FactureIn) SetIncomingDate(incomingDate *Timestamp) *FactureIn {
	factureIn.IncomingDate = incomingDate
	return factureIn
}

// SetAgent устанавливает Метаданные Контрагента.
//
// Принимает [Counterparty] или [Organization].
func (factureIn *FactureIn) SetAgent(agent AgentCounterpartyOrganizationInterface) *FactureIn {
	if agent != nil {
		factureIn.Agent = agent.asCOAgent()
	}
	return factureIn
}

// SetOrganization устанавливает Метаданные юрлица.
func (factureIn *FactureIn) SetOrganization(organization *Organization) *FactureIn {
	if organization != nil {
		factureIn.Organization = organization.Clean()
	}
	return factureIn
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (factureIn *FactureIn) SetOwner(owner *Employee) *FactureIn {
	if owner != nil {
		factureIn.Owner = owner.Clean()
	}
	return factureIn
}

// SetRate устанавливает Валюту.
//
// Передача nil передаёт сброс значения (null).
func (factureIn *FactureIn) SetRate(rate *Rate) *FactureIn {
	factureIn.Rate = NewNullValue(rate)
	return factureIn
}

// SetShared устанавливает флаг общего доступа.
func (factureIn *FactureIn) SetShared(shared bool) *FactureIn {
	factureIn.Shared = &shared
	return factureIn
}

// SetState устанавливает Метаданные статуса Приходного ордера.
//
// Передача nil передаёт сброс значения (null).
func (factureIn *FactureIn) SetState(state *State) *FactureIn {
	factureIn.State = NewNullValue(state)
	return factureIn
}

// SetSyncID устанавливает ID синхронизации.
func (factureIn *FactureIn) SetSyncID(syncID uuid.UUID) *FactureIn {
	factureIn.SyncID = &syncID
	return factureIn
}

// SetSupplies устанавливает Массив ссылок на связанные приемки.
//
// Принимает множество объектов [Supply].
func (factureIn *FactureIn) SetSupplies(supplies ...*Supply) *FactureIn {
	factureIn.Supplies.Push(supplies...)
	return factureIn
}

// SetPayments устанавливает Массив ссылок на связанные платежи.
//
// Принимает множество объектов [Payment].
func (factureIn *FactureIn) SetPayments(payments ...*Payment) *FactureIn {
	factureIn.Payments.Push(payments...)
	return factureIn
}

// SetIncomingNumber устанавливает Входящий номер.
func (factureIn *FactureIn) SetIncomingNumber(incomingNumber string) *FactureIn {
	factureIn.IncomingNumber = &incomingNumber
	return factureIn
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (factureIn *FactureIn) SetAttributes(attributes ...*Attribute) *FactureIn {
	factureIn.Attributes.Push(attributes...)
	return factureIn
}

// String реализует интерфейс [fmt.Stringer].
func (factureIn FactureIn) String() string {
	return Stringify(factureIn)
}

// MetaType возвращает код сущности.
func (FactureIn) MetaType() MetaType {
	return MetaTypeFactureIn
}

// Update shortcut
func (factureIn FactureIn) Update(ctx context.Context, client *Client, params ...*Params) (*FactureIn, *resty.Response, error) {
	return NewFactureInService(client).Update(ctx, factureIn.GetID(), &factureIn, params...)
}

// Create shortcut
func (factureIn FactureIn) Create(ctx context.Context, client *Client, params ...*Params) (*FactureIn, *resty.Response, error) {
	return NewFactureInService(client).Create(ctx, &factureIn, params...)
}

// Delete shortcut
func (factureIn FactureIn) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewFactureInService(client).Delete(ctx, factureIn.GetID())
}

// FactureInService описывает методы сервиса для работы со счетами-фактурами полученными.
type FactureInService interface {
	// GetList выполняет запрос на получение списка полученных счетов-фактур.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[FactureIn], *resty.Response, error)

	// Create выполняет запрос на создание полученного счета-фактуры.
	// Обязательные поля для заполнения:
	//	- incomingNumber (Входящий номер)
	//	- incomingDate (Входящая дата)
	// Принимает контекст, полученный счет-фактуру и опционально объект параметров запроса Params.
	// Возвращает созданную полученный счет-фактуру.
	Create(ctx context.Context, factureIn *FactureIn, params ...*Params) (*FactureIn, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение полученных счетов-фактур.
	// Изменяемые полученные счета-фактуры должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список полученных счетов-фактур и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых полученных счетов-фактур.
	CreateUpdateMany(ctx context.Context, factureInList Slice[FactureIn], params ...*Params) (*Slice[FactureIn], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление полученных счетов-фактур.
	// Принимает контекст и множество полученных счетов-фактур.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*FactureIn) (*DeleteManyResponse, *resty.Response, error)

	// Delete выполняет запрос на удаление полученного счета-фактуры.
	// Принимает контекст и ID полученного счета-фактуры.
	// Возвращает true в случае успешного удаления полученного счета-фактуры.
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельного полученного счета-фактуры по ID.
	// Принимает контекст, ID полученного счета-фактуры и опционально объект параметров запроса Params.
	// Возвращает найденный полученный счет-фактуру.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*FactureIn, *resty.Response, error)

	// Update выполняет запрос на изменение полученного счета-фактуры.
	// Принимает контекст, полученный счет-фактуру и опционально объект параметров запроса Params.
	// Возвращает изменённую полученный счет-фактуру.
	Update(ctx context.Context, id uuid.UUID, factureIn *FactureIn, params ...*Params) (*FactureIn, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных полученных счетов-фактур.
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

	// Template выполняет запрос на получение предзаполненного полученного счета-фактуры со стандартными полями.
	// без связи с какими-либо другими документами.
	// Принимает контекст.
	// Возвращает предзаполненный полученный счет-фактуру.
	Template(ctx context.Context) (*FactureIn, *resty.Response, error)

	// TemplateBased выполняет запрос на получение шаблона полученного счета-фактуры на основе других документов.
	// Основание, на котором может быть создан:
	//	- Приемка (Supply)
	//	- Исходящий платеж (PaymentOut)
	// Принимает контекст и множество документов из списка выше.
	// Возвращает предзаполненный полученный счет-фактуру на основании переданных документов.
	TemplateBased(ctx context.Context, basedOn ...MetaOwner) (*FactureIn, *resty.Response, error)

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
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*FactureIn, *resty.Response, error)

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

// NewFactureInService принимает [Client] и возвращает сервис для работы со счетами-фактурами полученными.
func NewFactureInService(client *Client) FactureInService {
	return newMainService[FactureIn, any, MetaAttributesStatesSharedWrapper, any](NewEndpoint(client, "entity/facturein"))
}
