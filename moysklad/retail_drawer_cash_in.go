package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"time"
)

// RetailDrawerCashIn Внесение денег.
//
// Код сущности: retaildrawercashin
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vnesenie-deneg
type RetailDrawerCashIn struct {
	Moment       *Timestamp        `json:"moment,omitempty"`       // Дата документа
	Created      *Timestamp        `json:"created,omitempty"`      // Дата создания
	AccountID    *uuid.UUID        `json:"accountId,omitempty"`    // ID учётной записи
	RetailShift  *RetailShift      `json:"retailShift,omitempty"`  // Ссылка на розничную смену, в рамках которой было выполнено Внесение денег
	Name         *string           `json:"name,omitempty"`         // Наименование Внесения денег
	Deleted      *Timestamp        `json:"deleted,omitempty"`      // Момент последнего удаления Внесения денег
	Description  *string           `json:"description,omitempty"`  // Комментарий Внесения денег
	Organization *Organization     `json:"organization,omitempty"` // Метаданные юрлица
	Files        *MetaArray[File]  `json:"files,omitempty"`        // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group        *Group            `json:"group,omitempty"`        // Отдел сотрудника
	ID           *uuid.UUID        `json:"id,omitempty"`           // ID Внесения денег
	Meta         *Meta             `json:"meta,omitempty"`         // Метаданные Внесения денег
	Applicable   *bool             `json:"applicable,omitempty"`   // Отметка о проведении
	Agent        *Employee         `json:"agent,omitempty"`        // Метаданные сотрудника, совершившего Внесение
	ExternalCode *string           `json:"externalCode,omitempty"` // Внешний код Внесения денег
	Owner        *Employee         `json:"owner,omitempty"`        // Метаданные владельца (Сотрудника)
	Printed      *bool             `json:"printed,omitempty"`      // Напечатан ли документ
	Published    *bool             `json:"published,omitempty"`    // Опубликован ли документ
	Rate         *NullValue[Rate]  `json:"rate,omitempty"`         // Валюта
	Shared       *bool             `json:"shared,omitempty"`       // Общий доступ
	State        *NullValue[State] `json:"state,omitempty"`        // Метаданные статуса Внесения денег
	Sum          *float64          `json:"sum,omitempty"`          // Сумма Внесения денег в копейках
	SyncID       *uuid.UUID        `json:"syncId,omitempty"`       // ID синхронизации
	Updated      *Timestamp        `json:"updated,omitempty"`      // Момент последнего обновления Внесения денег
	Attributes   Slice[Attribute]  `json:"attributes,omitempty"`   // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (retailDrawerCashIn RetailDrawerCashIn) Clean() *RetailDrawerCashIn {
	if retailDrawerCashIn.Meta == nil {
		return nil
	}
	return &RetailDrawerCashIn{Meta: retailDrawerCashIn.Meta}
}

// AsTaskOperation реализует интерфейс [TaskOperationConverter].
func (retailDrawerCashIn RetailDrawerCashIn) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: retailDrawerCashIn.Meta}
}

// GetMoment возвращает Дату документа.
func (retailDrawerCashIn RetailDrawerCashIn) GetMoment() time.Time {
	return Deref(retailDrawerCashIn.Moment).Time()
}

// GetCreated возвращает Дату создания.
func (retailDrawerCashIn RetailDrawerCashIn) GetCreated() time.Time {
	return Deref(retailDrawerCashIn.Created).Time()
}

// GetAccountID возвращает ID учётной записи.
func (retailDrawerCashIn RetailDrawerCashIn) GetAccountID() uuid.UUID {
	return Deref(retailDrawerCashIn.AccountID)
}

// GetRetailShift возвращает Ссылку на розничную смену, в рамках которой было выполнено Внесение денег.
func (retailDrawerCashIn RetailDrawerCashIn) GetRetailShift() RetailShift {
	return Deref(retailDrawerCashIn.RetailShift)
}

// GetName возвращает Наименование Внесения денег.
func (retailDrawerCashIn RetailDrawerCashIn) GetName() string {
	return Deref(retailDrawerCashIn.Name)
}

// GetDeleted возвращает Момент последнего удаления Внесения денег.
func (retailDrawerCashIn RetailDrawerCashIn) GetDeleted() time.Time {
	return Deref(retailDrawerCashIn.Deleted).Time()
}

// GetDescription возвращает Комментарий Внесения денег.
func (retailDrawerCashIn RetailDrawerCashIn) GetDescription() string {
	return Deref(retailDrawerCashIn.Description)
}

// GetOrganization возвращает Метаданные юрлица.
func (retailDrawerCashIn RetailDrawerCashIn) GetOrganization() Organization {
	return Deref(retailDrawerCashIn.Organization)
}

// GetFiles возвращает Метаданные массива Файлов.
func (retailDrawerCashIn RetailDrawerCashIn) GetFiles() MetaArray[File] {
	return Deref(retailDrawerCashIn.Files)
}

// GetGroup возвращает Отдел сотрудника.
func (retailDrawerCashIn RetailDrawerCashIn) GetGroup() Group {
	return Deref(retailDrawerCashIn.Group)
}

// GetID возвращает ID Перемещения.
func (retailDrawerCashIn RetailDrawerCashIn) GetID() uuid.UUID {
	return Deref(retailDrawerCashIn.ID)
}

// GetMeta возвращает Метаданные Внесения денег.
func (retailDrawerCashIn RetailDrawerCashIn) GetMeta() Meta {
	return Deref(retailDrawerCashIn.Meta)
}

// GetApplicable возвращает Отметку о проведении.
func (retailDrawerCashIn RetailDrawerCashIn) GetApplicable() bool {
	return Deref(retailDrawerCashIn.Applicable)
}

// GetAgent возвращает Метаданные сотрудника, совершившего Внесение.
func (retailDrawerCashIn RetailDrawerCashIn) GetAgent() Employee {
	return Deref(retailDrawerCashIn.Agent)
}

// GetExternalCode возвращает Внешний код Внесения денег.
func (retailDrawerCashIn RetailDrawerCashIn) GetExternalCode() string {
	return Deref(retailDrawerCashIn.ExternalCode)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (retailDrawerCashIn RetailDrawerCashIn) GetOwner() Employee {
	return Deref(retailDrawerCashIn.Owner)
}

// GetPrinted возвращает true, если документ напечатан.
func (retailDrawerCashIn RetailDrawerCashIn) GetPrinted() bool {
	return Deref(retailDrawerCashIn.Printed)
}

// GetPublished возвращает true, если документ опубликован.
func (retailDrawerCashIn RetailDrawerCashIn) GetPublished() bool {
	return Deref(retailDrawerCashIn.Published)
}

// GetRate возвращает Валюту.
func (retailDrawerCashIn RetailDrawerCashIn) GetRate() Rate {
	return Deref(retailDrawerCashIn.Rate).getValue()
}

// GetShared возвращает флаг Общего доступа.
func (retailDrawerCashIn RetailDrawerCashIn) GetShared() bool {
	return Deref(retailDrawerCashIn.Shared)
}

// GetState возвращает Метаданные статуса Перемещения.
func (retailDrawerCashIn RetailDrawerCashIn) GetState() State {
	return Deref(retailDrawerCashIn.State).getValue()
}

// GetSum возвращает Сумму Перемещения в копейках.
func (retailDrawerCashIn RetailDrawerCashIn) GetSum() float64 {
	return Deref(retailDrawerCashIn.Sum)
}

// GetSyncID возвращает ID синхронизации.
func (retailDrawerCashIn RetailDrawerCashIn) GetSyncID() uuid.UUID {
	return Deref(retailDrawerCashIn.SyncID)
}

// GetUpdated возвращает Момент последнего обновления Внесения денег.
func (retailDrawerCashIn RetailDrawerCashIn) GetUpdated() time.Time {
	return Deref(retailDrawerCashIn.Updated).Time()
}

// GetAttributes возвращает Список метаданных доп. полей.
func (retailDrawerCashIn RetailDrawerCashIn) GetAttributes() Slice[Attribute] {
	return retailDrawerCashIn.Attributes
}

// SetMoment устанавливает Дату документа.
func (retailDrawerCashIn *RetailDrawerCashIn) SetMoment(moment time.Time) *RetailDrawerCashIn {
	retailDrawerCashIn.Moment = NewTimestamp(moment)
	return retailDrawerCashIn
}

// SetRetailShift устанавливает Ссылку на розничную смену, в рамках которой было выполнено Внесение денег.
func (retailDrawerCashIn *RetailDrawerCashIn) SetRetailShift(retailShift *RetailShift) *RetailDrawerCashIn {
	if retailShift != nil {
		retailDrawerCashIn.RetailShift = retailShift.Clean()
	}
	return retailDrawerCashIn
}

// SetName устанавливает Наименование Внесения денег.
func (retailDrawerCashIn *RetailDrawerCashIn) SetName(name string) *RetailDrawerCashIn {
	retailDrawerCashIn.Name = &name
	return retailDrawerCashIn
}

// SetDescription устанавливает Комментарий Внесения денег.
func (retailDrawerCashIn *RetailDrawerCashIn) SetDescription(description string) *RetailDrawerCashIn {
	retailDrawerCashIn.Description = &description
	return retailDrawerCashIn
}

// SetOrganization устанавливает Метаданные юрлица.
func (retailDrawerCashIn *RetailDrawerCashIn) SetOrganization(organization *Organization) *RetailDrawerCashIn {
	if organization != nil {
		retailDrawerCashIn.Organization = organization.Clean()
	}
	return retailDrawerCashIn
}

// SetFiles устанавливает Метаданные массива Файлов.
//
// Принимает множество объектов [File].
func (retailDrawerCashIn *RetailDrawerCashIn) SetFiles(files ...*File) *RetailDrawerCashIn {
	retailDrawerCashIn.Files = NewMetaArrayFrom(files)
	return retailDrawerCashIn
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (retailDrawerCashIn *RetailDrawerCashIn) SetGroup(group *Group) *RetailDrawerCashIn {
	if group != nil {
		retailDrawerCashIn.Group = group.Clean()
	}
	return retailDrawerCashIn
}

// SetMeta устанавливает Метаданные Внесения денег.
func (retailDrawerCashIn *RetailDrawerCashIn) SetMeta(meta *Meta) *RetailDrawerCashIn {
	retailDrawerCashIn.Meta = meta
	return retailDrawerCashIn
}

// SetApplicable устанавливает Отметку о проведении.
func (retailDrawerCashIn *RetailDrawerCashIn) SetApplicable(applicable bool) *RetailDrawerCashIn {
	retailDrawerCashIn.Applicable = &applicable
	return retailDrawerCashIn
}

// SetAgent устанавливает Метаданные сотрудника, совершившего Внесение.
func (retailDrawerCashIn *RetailDrawerCashIn) SetAgent(agent *Employee) *RetailDrawerCashIn {
	if agent != nil {
		retailDrawerCashIn.Agent = agent.Clean()
	}
	return retailDrawerCashIn
}

// SetExternalCode устанавливает Внешний код  Внесения денег.
func (retailDrawerCashIn *RetailDrawerCashIn) SetExternalCode(externalCode string) *RetailDrawerCashIn {
	retailDrawerCashIn.ExternalCode = &externalCode
	return retailDrawerCashIn
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (retailDrawerCashIn *RetailDrawerCashIn) SetOwner(owner *Employee) *RetailDrawerCashIn {
	if owner != nil {
		retailDrawerCashIn.Owner = owner.Clean()
	}
	return retailDrawerCashIn
}

// SetRate устанавливает Валюту.
//
// Передача nil передаёт сброс значения (null).
func (retailDrawerCashIn *RetailDrawerCashIn) SetRate(rate *Rate) *RetailDrawerCashIn {
	retailDrawerCashIn.Rate = NewNullValue(rate)
	return retailDrawerCashIn
}

// SetShared устанавливает флаг общего доступа.
func (retailDrawerCashIn *RetailDrawerCashIn) SetShared(shared bool) *RetailDrawerCashIn {
	retailDrawerCashIn.Shared = &shared
	return retailDrawerCashIn
}

// SetState устанавливает Метаданные статуса Внесения денег.
//
// Передача nil передаёт сброс значения (null).
func (retailDrawerCashIn *RetailDrawerCashIn) SetState(state *State) *RetailDrawerCashIn {
	retailDrawerCashIn.State = NewNullValue(state)
	return retailDrawerCashIn
}

// SetSyncID устанавливает ID синхронизации.
func (retailDrawerCashIn *RetailDrawerCashIn) SetSyncID(syncID uuid.UUID) *RetailDrawerCashIn {
	retailDrawerCashIn.SyncID = &syncID
	return retailDrawerCashIn
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (retailDrawerCashIn *RetailDrawerCashIn) SetAttributes(attributes ...*Attribute) *RetailDrawerCashIn {
	retailDrawerCashIn.Attributes.Push(attributes...)
	return retailDrawerCashIn
}

// String реализует интерфейс [fmt.Stringer].
func (retailDrawerCashIn RetailDrawerCashIn) String() string {
	return Stringify(retailDrawerCashIn)
}

// MetaType возвращает код сущности.
func (RetailDrawerCashIn) MetaType() MetaType {
	return MetaTypeRetailDrawerCashIn
}

// Update shortcut
func (retailDrawerCashIn RetailDrawerCashIn) Update(ctx context.Context, client *Client, params ...*Params) (*RetailDrawerCashIn, *resty.Response, error) {
	return NewRetailDrawerCashInService(client).Update(ctx, retailDrawerCashIn.GetID(), &retailDrawerCashIn, params...)
}

// Create shortcut
func (retailDrawerCashIn RetailDrawerCashIn) Create(ctx context.Context, client *Client, params ...*Params) (*RetailDrawerCashIn, *resty.Response, error) {
	return NewRetailDrawerCashInService(client).Create(ctx, &retailDrawerCashIn, params...)
}

// Delete shortcut
func (retailDrawerCashIn RetailDrawerCashIn) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewRetailDrawerCashInService(client).Delete(ctx, retailDrawerCashIn.GetID())
}

// RetailDrawerCashInService описывает методы сервиса для работы с внесениями денег.
type RetailDrawerCashInService interface {
	// GetList выполняет запрос на получение списка внесений денег.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[RetailDrawerCashIn], *resty.Response, error)

	// Create выполняет запрос на создание внесения денег.
	// Обязательные поля для заполнения:
	//	- organization (Ссылка на ваше юрлицо)
	//	- agent (Ссылка на сотрудника, совершившего Внесение)
	//	- retailShift (Ссылка на розничную смену)
	// Принимает контекст, внесение денег и опционально объект параметров запроса Params.
	// Возвращает созданное внесение денег.
	Create(ctx context.Context, retailDrawerCashIn *RetailDrawerCashIn, params ...*Params) (*RetailDrawerCashIn, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение внесений денег.
	// Изменяемые внесения денег должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список внесений денег и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых внесений денег.
	CreateUpdateMany(ctx context.Context, retailDrawerCashInList Slice[RetailDrawerCashIn], params ...*Params) (*Slice[RetailDrawerCashIn], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление внесений денег.
	// Принимает контекст и множество внесений денег.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*RetailDrawerCashIn) (*DeleteManyResponse, *resty.Response, error)

	// Delete выполняет запрос на удаление внесения денег.
	// Принимает контекст и ID внесения денег.
	// Возвращает «true» в случае успешного удаления внесения денег.
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение внесения денег по ID.
	// Принимает контекст, ID внесения денег и опционально объект параметров запроса Params.
	// Возвращает внесение денег.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*RetailDrawerCashIn, *resty.Response, error)

	// Update выполняет запрос на изменение внесения денег.
	// Принимает контекст, внесение денег и опционально объект параметров запроса Params.
	// Возвращает изменённое внесение денег.
	Update(ctx context.Context, id uuid.UUID, retailDrawerCashIn *RetailDrawerCashIn, params ...*Params) (*RetailDrawerCashIn, *resty.Response, error)

	// Template выполняет запрос на получение предзаполненного внесения денег со стандартными полями без связи с какими-либо другими документами.
	// Принимает контекст.
	// Возвращает предзаполненное внесение денег.
	Template(ctx context.Context) (*RetailDrawerCashIn, *resty.Response, error)

	// TemplateBased выполняет запрос на получение шаблона внесения денег на основе других документов.
	// Основание, на котором может быть создано:
	//	- Розничная смена (RetailShift)
	// Принимает контекст и один документ из списка выше.
	// Возвращает предзаполненное внесение денег на основании переданных документов.
	TemplateBased(ctx context.Context, basedOn ...MetaOwner) (*RetailDrawerCashIn, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных внесений денег.
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
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*RetailDrawerCashIn, *resty.Response, error)

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
	EndpointRetailDrawerCashIn = EndpointEntity + string(MetaTypeRetailDrawerCashIn)
)

// NewRetailDrawerCashInService принимает [Client] и возвращает сервис для работы с внесениями денег.
func NewRetailDrawerCashInService(client *Client) RetailDrawerCashInService {
	return newMainService[RetailDrawerCashIn, any, MetaAttributesStatesSharedWrapper, any](client, EndpointRetailDrawerCashIn)
}
