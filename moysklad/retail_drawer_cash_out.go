package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"time"
)

// RetailDrawerCashOut Выплата денег.
//
// Код сущности: retaildrawercashout
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vyplata-deneg
type RetailDrawerCashOut struct {
	Meta         *Meta             `json:"meta,omitempty"`         // Метаданные Выплаты денег
	Applicable   *bool             `json:"applicable,omitempty"`   // Отметка о проведении
	Moment       *Timestamp        `json:"moment,omitempty"`       // Дата документа
	Name         *string           `json:"name,omitempty"`         // Наименование Выплаты денег
	Code         *string           `json:"code,omitempty"`         // Код Выплаты денег
	Created      *Timestamp        `json:"created,omitempty"`      // Дата создания
	Deleted      *Timestamp        `json:"deleted,omitempty"`      // Момент последнего удаления Выплаты денег
	Description  *string           `json:"description,omitempty"`  // Комментарий Выплаты денег
	ExternalCode *string           `json:"externalCode,omitempty"` // Внешний код Выплаты денег
	Files        *MetaArray[File]  `json:"files,omitempty"`        // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group        *Group            `json:"group,omitempty"`        // Отдел сотрудника
	ID           *uuid.UUID        `json:"id,omitempty"`           // ID Выплаты денег
	RetailShift  *RetailShift      `json:"retailShift,omitempty"`  // Ссылка на розничную смену, в рамках которой было выполнено Внесение денег
	Agent        *Employee         `json:"agent,omitempty"`        // Ссылка на сотрудника, которому была совершена Выплата
	AccountID    *uuid.UUID        `json:"accountId,omitempty"`    // ID учётной записи
	Organization *Organization     `json:"organization,omitempty"` // Метаданные юрлица
	Owner        *Employee         `json:"owner,omitempty"`        // Метаданные владельца (Сотрудника)
	Printed      *bool             `json:"printed,omitempty"`      // Напечатан ли документ
	Published    *bool             `json:"published,omitempty"`    // Опубликован ли документ
	Rate         *NullValue[Rate]  `json:"rate,omitempty"`         // Валюта
	Shared       *bool             `json:"shared,omitempty"`       // Общий доступ
	State        *NullValue[State] `json:"state,omitempty"`        // Метаданные статуса Выплаты денег
	Sum          *float64          `json:"sum,omitempty"`          // Сумма Выплаты денег установленной валюте
	SyncID       *uuid.UUID        `json:"syncId,omitempty"`       // ID синхронизации
	Updated      *Timestamp        `json:"updated,omitempty"`      // Момент последнего обновления Выплаты денег
	Attributes   Slice[Attribute]  `json:"attributes,omitempty"`   // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (retailDrawerCashOut RetailDrawerCashOut) Clean() *RetailDrawerCashOut {
	if retailDrawerCashOut.Meta == nil {
		return nil
	}
	return &RetailDrawerCashOut{Meta: retailDrawerCashOut.Meta}
}

// AsTaskOperation реализует интерфейс [TaskOperationConverter].
func (retailDrawerCashOut RetailDrawerCashOut) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: retailDrawerCashOut.Meta}
}

// GetMeta возвращает Метаданные Выплаты денег.
func (retailDrawerCashOut RetailDrawerCashOut) GetMeta() Meta {
	return Deref(retailDrawerCashOut.Meta)
}

// GetApplicable возвращает Отметку о проведении.
func (retailDrawerCashOut RetailDrawerCashOut) GetApplicable() bool {
	return Deref(retailDrawerCashOut.Applicable)
}

// GetMoment возвращает Дату документа.
func (retailDrawerCashOut RetailDrawerCashOut) GetMoment() time.Time {
	return Deref(retailDrawerCashOut.Moment).Time()
}

// GetName возвращает Наименование Выплаты денег.
func (retailDrawerCashOut RetailDrawerCashOut) GetName() string {
	return Deref(retailDrawerCashOut.Name)
}

// GetCode возвращает Код Выплаты денег.
func (retailDrawerCashOut RetailDrawerCashOut) GetCode() string {
	return Deref(retailDrawerCashOut.Code)
}

// GetCreated возвращает Дату создания.
func (retailDrawerCashOut RetailDrawerCashOut) GetCreated() time.Time {
	return Deref(retailDrawerCashOut.Created).Time()
}

// GetDeleted возвращает Момент последнего удаления Выплаты денег.
func (retailDrawerCashOut RetailDrawerCashOut) GetDeleted() time.Time {
	return Deref(retailDrawerCashOut.Deleted).Time()
}

// GetDescription возвращает Комментарий Выплаты денег.
func (retailDrawerCashOut RetailDrawerCashOut) GetDescription() string {
	return Deref(retailDrawerCashOut.Description)
}

// GetExternalCode возвращает Внешний код Выплаты денег.
func (retailDrawerCashOut RetailDrawerCashOut) GetExternalCode() string {
	return Deref(retailDrawerCashOut.ExternalCode)
}

// GetFiles возвращает Метаданные массива Файлов.
func (retailDrawerCashOut RetailDrawerCashOut) GetFiles() MetaArray[File] {
	return Deref(retailDrawerCashOut.Files)
}

// GetGroup возвращает Отдел сотрудника.
func (retailDrawerCashOut RetailDrawerCashOut) GetGroup() Group {
	return Deref(retailDrawerCashOut.Group)
}

// GetID возвращает ID Выплаты денег.
func (retailDrawerCashOut RetailDrawerCashOut) GetID() uuid.UUID {
	return Deref(retailDrawerCashOut.ID)
}

// GetRetailShift возвращает Ссылку на розничную смену, в рамках которой было выполнено Внесение денег.
func (retailDrawerCashOut RetailDrawerCashOut) GetRetailShift() RetailShift {
	return Deref(retailDrawerCashOut.RetailShift)
}

// GetAgent возвращает Ссылку на сотрудника, которому была совершена Выплата.
func (retailDrawerCashOut RetailDrawerCashOut) GetAgent() Employee {
	return Deref(retailDrawerCashOut.Agent)
}

// GetAccountID возвращает ID учётной записи.
func (retailDrawerCashOut RetailDrawerCashOut) GetAccountID() uuid.UUID {
	return Deref(retailDrawerCashOut.AccountID)
}

// GetOrganization возвращает Метаданные юрлица.
func (retailDrawerCashOut RetailDrawerCashOut) GetOrganization() Organization {
	return Deref(retailDrawerCashOut.Organization)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (retailDrawerCashOut RetailDrawerCashOut) GetOwner() Employee {
	return Deref(retailDrawerCashOut.Owner)
}

// GetPrinted возвращает true, если документ напечатан.
func (retailDrawerCashOut RetailDrawerCashOut) GetPrinted() bool {
	return Deref(retailDrawerCashOut.Printed)
}

// GetPublished возвращает true, если документ опубликован.
func (retailDrawerCashOut RetailDrawerCashOut) GetPublished() bool {
	return Deref(retailDrawerCashOut.Published)
}

// GetRate возвращает Валюту.
func (retailDrawerCashOut RetailDrawerCashOut) GetRate() Rate {
	return Deref(retailDrawerCashOut.Rate).getValue()
}

// GetShared возвращает флаг Общего доступа.
func (retailDrawerCashOut RetailDrawerCashOut) GetShared() bool {
	return Deref(retailDrawerCashOut.Shared)
}

// GetState возвращает Метаданные статуса Выплаты денег.
func (retailDrawerCashOut RetailDrawerCashOut) GetState() State {
	return Deref(retailDrawerCashOut.State).getValue()
}

// GetSum возвращает Сумму Выплаты денег в установленной валюте.
func (retailDrawerCashOut RetailDrawerCashOut) GetSum() float64 {
	return Deref(retailDrawerCashOut.Sum)
}

// GetSyncID возвращает ID синхронизации.
func (retailDrawerCashOut RetailDrawerCashOut) GetSyncID() uuid.UUID {
	return Deref(retailDrawerCashOut.SyncID)
}

// GetUpdated возвращает Момент последнего обновления Выплаты денег.
func (retailDrawerCashOut RetailDrawerCashOut) GetUpdated() time.Time {
	return Deref(retailDrawerCashOut.Updated).Time()
}

// GetAttributes возвращает Список метаданных доп. полей.
func (retailDrawerCashOut RetailDrawerCashOut) GetAttributes() Slice[Attribute] {
	return retailDrawerCashOut.Attributes
}

// SetMeta устанавливает Метаданные Выплаты денег.
func (retailDrawerCashOut *RetailDrawerCashOut) SetMeta(meta *Meta) *RetailDrawerCashOut {
	retailDrawerCashOut.Meta = meta
	return retailDrawerCashOut
}

// SetApplicable устанавливает Отметку о проведении.
func (retailDrawerCashOut *RetailDrawerCashOut) SetApplicable(applicable bool) *RetailDrawerCashOut {
	retailDrawerCashOut.Applicable = &applicable
	return retailDrawerCashOut
}

// SetMoment устанавливает Дату документа.
func (retailDrawerCashOut *RetailDrawerCashOut) SetMoment(moment time.Time) *RetailDrawerCashOut {
	retailDrawerCashOut.Moment = NewTimestamp(moment)
	return retailDrawerCashOut
}

// SetName устанавливает Наименование Выплаты денег.
func (retailDrawerCashOut *RetailDrawerCashOut) SetName(name string) *RetailDrawerCashOut {
	retailDrawerCashOut.Name = &name
	return retailDrawerCashOut
}

// SetCode устанавливает Код Выплаты денег.
func (retailDrawerCashOut *RetailDrawerCashOut) SetCode(code string) *RetailDrawerCashOut {
	retailDrawerCashOut.Code = &code
	return retailDrawerCashOut
}

// SetDescription устанавливает Комментарий Выплаты денег.
func (retailDrawerCashOut *RetailDrawerCashOut) SetDescription(description string) *RetailDrawerCashOut {
	retailDrawerCashOut.Description = &description
	return retailDrawerCashOut
}

// SetExternalCode устанавливает Внешний код Выплаты денег.
func (retailDrawerCashOut *RetailDrawerCashOut) SetExternalCode(externalCode string) *RetailDrawerCashOut {
	retailDrawerCashOut.ExternalCode = &externalCode
	return retailDrawerCashOut
}

// SetFiles устанавливает Метаданные массива Файлов.
//
// Принимает множество объектов [File].
func (retailDrawerCashOut *RetailDrawerCashOut) SetFiles(files ...*File) *RetailDrawerCashOut {
	retailDrawerCashOut.Files = NewMetaArrayFrom(files)
	return retailDrawerCashOut
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (retailDrawerCashOut *RetailDrawerCashOut) SetGroup(group *Group) *RetailDrawerCashOut {
	if group != nil {
		retailDrawerCashOut.Group = group.Clean()
	}
	return retailDrawerCashOut
}

// SetRetailShift устанавливает Ссылку на розничную смену, в рамках которой было выполнено Внесение денег.
func (retailDrawerCashOut *RetailDrawerCashOut) SetRetailShift(retailShift *RetailShift) *RetailDrawerCashOut {
	if retailShift != nil {
		retailDrawerCashOut.RetailShift = retailShift.Clean()
	}
	return retailDrawerCashOut
}

// SetAgent устанавливает Ссылку на сотрудника, которому была совершена Выплата.
func (retailDrawerCashOut *RetailDrawerCashOut) SetAgent(agent *Employee) *RetailDrawerCashOut {
	if agent != nil {
		retailDrawerCashOut.Agent = agent.Clean()
	}
	return retailDrawerCashOut
}

// SetOrganization устанавливает Метаданные юрлица.
func (retailDrawerCashOut *RetailDrawerCashOut) SetOrganization(organization *Organization) *RetailDrawerCashOut {
	if organization != nil {
		retailDrawerCashOut.Organization = organization.Clean()
	}
	return retailDrawerCashOut
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (retailDrawerCashOut *RetailDrawerCashOut) SetOwner(owner *Employee) *RetailDrawerCashOut {
	if owner != nil {
		retailDrawerCashOut.Owner = owner.Clean()
	}
	return retailDrawerCashOut
}

// SetRate устанавливает Валюту.
//
// Передача nil передаёт сброс значения (null).
func (retailDrawerCashOut *RetailDrawerCashOut) SetRate(rate *Rate) *RetailDrawerCashOut {
	retailDrawerCashOut.Rate = NewNullValue(rate)
	return retailDrawerCashOut
}

// SetShared устанавливает флаг общего доступа.
func (retailDrawerCashOut *RetailDrawerCashOut) SetShared(shared bool) *RetailDrawerCashOut {
	retailDrawerCashOut.Shared = &shared
	return retailDrawerCashOut
}

// SetState устанавливает Метаданные статуса Выплаты денег.
//
// Передача nil передаёт сброс значения (null).
func (retailDrawerCashOut *RetailDrawerCashOut) SetState(state *State) *RetailDrawerCashOut {
	retailDrawerCashOut.State = NewNullValue(state)
	return retailDrawerCashOut
}

// SetSyncID устанавливает ID синхронизации.
func (retailDrawerCashOut *RetailDrawerCashOut) SetSyncID(syncID uuid.UUID) *RetailDrawerCashOut {
	retailDrawerCashOut.SyncID = &syncID
	return retailDrawerCashOut
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (retailDrawerCashOut *RetailDrawerCashOut) SetAttributes(attributes ...*Attribute) *RetailDrawerCashOut {
	retailDrawerCashOut.Attributes.Push(attributes...)
	return retailDrawerCashOut
}

// String реализует интерфейс [fmt.Stringer].
func (retailDrawerCashOut RetailDrawerCashOut) String() string {
	return Stringify(retailDrawerCashOut)
}

// MetaType возвращает код сущности.
func (RetailDrawerCashOut) MetaType() MetaType {
	return MetaTypeRetailDrawerCashOut
}

// Update shortcut
func (retailDrawerCashOut *RetailDrawerCashOut) Update(ctx context.Context, client *Client, params ...*Params) (*RetailDrawerCashOut, *resty.Response, error) {
	return NewRetailDrawerCashOutService(client).Update(ctx, retailDrawerCashOut.GetID(), retailDrawerCashOut, params...)
}

// Create shortcut
func (retailDrawerCashOut *RetailDrawerCashOut) Create(ctx context.Context, client *Client, params ...*Params) (*RetailDrawerCashOut, *resty.Response, error) {
	return NewRetailDrawerCashOutService(client).Create(ctx, retailDrawerCashOut, params...)
}

// Delete shortcut
func (retailDrawerCashOut *RetailDrawerCashOut) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewRetailDrawerCashOutService(client).Delete(ctx, retailDrawerCashOut)
}

// RetailDrawerCashOutService описывает методы сервиса для работы с выплатами денег.
type RetailDrawerCashOutService interface {
	// GetList выполняет запрос на получение списка выплат денег.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[RetailDrawerCashOut], *resty.Response, error)

	// GetListAll выполняет запрос на получение всех выплат денег в виде списка.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает список объектов.
	GetListAll(ctx context.Context, params ...*Params) (*Slice[RetailDrawerCashOut], *resty.Response, error)

	// Create выполняет запрос на создание выплаты денег.
	// Обязательные поля для заполнения:
	//	- organization (Ссылка на ваше юрлицо)
	//	- agent (Ссылка на сотрудника, которому была совершена Выплата)
	//	- retailShift (Ссылка на розничную смену)
	// Принимает контекст, выплату денег и опционально объект параметров запроса Params.
	// Возвращает созданную выплату денег.
	Create(ctx context.Context, retailDrawerCashOut *RetailDrawerCashOut, params ...*Params) (*RetailDrawerCashOut, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение выплат денег.
	// Изменяемые выплаты денег должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список выплат денег и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых выплат денег.
	CreateUpdateMany(ctx context.Context, retailDrawerCashOutList Slice[RetailDrawerCashOut], params ...*Params) (*Slice[RetailDrawerCashOut], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление выплат денег.
	// Принимает контекст и множество выплат денег.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*RetailDrawerCashOut) (*DeleteManyResponse, *resty.Response, error)

	// DeleteByID выполняет запрос на удаление выплаты денег по ID.
	// Принимает контекст и ID выплаты денег.
	// Возвращает «true» в случае успешного удаления выплаты денег.
	DeleteByID(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// Delete выполняет запрос на удаление выплаты денег.
	// Принимает контекст и выплату денег.
	// Возвращает «true» в случае успешного удаления выплаты денег.
	Delete(ctx context.Context, entity *RetailDrawerCashOut) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение выплаты денег по ID.
	// Принимает контекст, ID выплаты денег и опционально объект параметров запроса Params.
	// Возвращает выплату денег.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*RetailDrawerCashOut, *resty.Response, error)

	// Update выполняет запрос на изменение выплаты денег.
	// Принимает контекст, выплату денег и опционально объект параметров запроса Params.
	// Возвращает изменённую выплату денег.
	Update(ctx context.Context, id uuid.UUID, retailDrawerCashOut *RetailDrawerCashOut, params ...*Params) (*RetailDrawerCashOut, *resty.Response, error)

	// Template выполняет запрос на получение предзаполненной выплаты денег со стандартными полями без связи с какими-либо другими документами.
	// Принимает контекст.
	// Возвращает предзаполненную выплату денег.
	Template(ctx context.Context) (*RetailDrawerCashOut, *resty.Response, error)

	// TemplateBased выполняет запрос на получение шаблона выплаты денег на основе других документов.
	// Основание, на котором может быть создана:
	//	- Розничная смена (RetailShift)
	// Принимает контекст и один документ из списка выше.
	// Возвращает предзаполненную выплату денег на основании переданных документов.
	TemplateBased(ctx context.Context, basedOn ...MetaOwner) (*RetailDrawerCashOut, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных выплат денег.
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
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*RetailDrawerCashOut, *resty.Response, error)

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
}

const (
	EndpointRetailDrawerCashOut = EndpointEntity + string(MetaTypeRetailDrawerCashOut)
)

// NewRetailDrawerCashOutService принимает [Client] и возвращает сервис для работы с выплатами денег.
func NewRetailDrawerCashOutService(client *Client) RetailDrawerCashOutService {
	return newMainService[RetailDrawerCashOut, any, MetaAttributesStatesSharedWrapper, any](client, EndpointRetailDrawerCashOut)
}
