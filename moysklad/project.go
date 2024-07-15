package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"time"
)

// Project Проект.
//
// Код сущности: project
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-proekt
type Project struct {
	Group        *Group           `json:"group,omitempty"`        // Отдел сотрудника
	Archived     *bool            `json:"archived,omitempty"`     // Добавлен ли Проект в архив
	Code         *string          `json:"code,omitempty"`         // Код Проекта
	Description  *string          `json:"description,omitempty"`  // Описание Проекта
	ExternalCode *string          `json:"externalCode,omitempty"` // Внешний код Проекта
	AccountID    *uuid.UUID       `json:"accountId,omitempty"`    // ID учётной записи
	ID           *uuid.UUID       `json:"id,omitempty"`           // ID проекта
	Meta         *Meta            `json:"meta,omitempty"`         // Метаданные Проекта
	Name         *string          `json:"name,omitempty"`         // Наименование Проекта
	Owner        *Employee        `json:"owner,omitempty"`        // Метаданные владельца (Сотрудника)
	Shared       *bool            `json:"shared,omitempty"`       // Общий доступ
	Updated      *Timestamp       `json:"updated,omitempty"`      // Момент последнего обновления Проекта
	Attributes   Slice[Attribute] `json:"attributes,omitempty"`   // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (project Project) Clean() *Project {
	if project.Meta == nil {
		return nil
	}
	return &Project{Meta: project.Meta}
}

// GetGroup возвращает Отдел сотрудника.
func (project Project) GetGroup() Group {
	return Deref(project.Group)
}

// GetArchived возвращает true, если проект находится в архиве.
func (project Project) GetArchived() bool {
	return Deref(project.Archived)
}

// GetCode возвращает Код проекта.
func (project Project) GetCode() string {
	return Deref(project.Code)
}

// GetDescription возвращает Описание проекта.
func (project Project) GetDescription() string {
	return Deref(project.Description)
}

// GetExternalCode возвращает Внешний код проекта.
func (project Project) GetExternalCode() string {
	return Deref(project.ExternalCode)
}

// GetAccountID возвращает ID учётной записи.
func (project Project) GetAccountID() uuid.UUID {
	return Deref(project.AccountID)
}

// GetID возвращает ID проекта.
func (project Project) GetID() uuid.UUID {
	return Deref(project.ID)
}

// GetMeta возвращает Метаданные проекта.
func (project Project) GetMeta() Meta {
	return Deref(project.Meta)
}

// GetName возвращает Наименование проекта.
func (project Project) GetName() string {
	return Deref(project.Name)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (project Project) GetOwner() Employee {
	return Deref(project.Owner)
}

// GetShared возвращает флаг Общего доступа.
func (project Project) GetShared() bool {
	return Deref(project.Shared)
}

// GetUpdated возвращает Момент последнего обновления проекта.
func (project Project) GetUpdated() time.Time {
	return Deref(project.Updated).Time()
}

// GetAttributes возвращает Список метаданных доп. полей.
func (project Project) GetAttributes() Slice[Attribute] {
	return project.Attributes
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (project *Project) SetGroup(group *Group) *Project {
	if group != nil {
		project.Group = group.Clean()
	}
	return project
}

// SetArchived устанавливает флаг нахождения проекта в архиве.
func (project *Project) SetArchived(archived bool) *Project {
	project.Archived = &archived
	return project
}

// SetCode устанавливает Код проекта.
func (project *Project) SetCode(code string) *Project {
	project.Code = &code
	return project
}

// SetDescription устанавливает Описание проекта.
func (project *Project) SetDescription(description string) *Project {
	project.Description = &description
	return project
}

// SetExternalCode устанавливает Внешний код проекта.
func (project *Project) SetExternalCode(externalCode string) *Project {
	project.ExternalCode = &externalCode
	return project
}

// SetMeta устанавливает Метаданные проекта.
func (project *Project) SetMeta(meta *Meta) *Project {
	project.Meta = meta
	return project
}

// SetName устанавливает Наименование проекта.
func (project *Project) SetName(name string) *Project {
	project.Name = &name
	return project
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (project *Project) SetOwner(owner *Employee) *Project {
	if owner != nil {
		project.Owner = owner.Clean()
	}
	return project
}

// SetShared устанавливает флаг общего доступа.
func (project *Project) SetShared(shared bool) *Project {
	project.Shared = &shared
	return project
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (project *Project) SetAttributes(attributes ...*Attribute) *Project {
	project.Attributes.Push(attributes...)
	return project
}

// String реализует интерфейс [fmt.Stringer].
func (project Project) String() string {
	return Stringify(project)
}

// MetaType возвращает код сущности.
func (Project) MetaType() MetaType {
	return MetaTypeProject
}

// Update shortcut
func (project *Project) Update(ctx context.Context, client *Client, params ...*Params) (*Project, *resty.Response, error) {
	return NewProjectService(client).Update(ctx, project.GetID(), project, params...)
}

// Create shortcut
func (project *Project) Create(ctx context.Context, client *Client, params ...*Params) (*Project, *resty.Response, error) {
	return NewProjectService(client).Create(ctx, project, params...)
}

// Delete shortcut
func (project *Project) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewProjectService(client).Delete(ctx, project)
}

// ProjectService описывает методы сервиса для работы с проектами.
type ProjectService interface {
	// GetList выполняет запрос на получение списка проектов.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[Project], *resty.Response, error)

	// Create выполняет запрос на создание проекта.
	// Обязательные поля для заполнения:
	//	- name (Наименование проекта)
	// Принимает контекст, проект и опционально объект параметров запроса Params.
	// Возвращает созданный проект.
	Create(ctx context.Context, project *Project, params ...*Params) (*Project, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение проектов.
	// Изменяемые проекты должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список проектов и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых проектов.
	CreateUpdateMany(ctx context.Context, projectList Slice[Project], params ...*Params) (*Slice[Project], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление проектов.
	// Принимает контекст и множество проектов.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*Project) (*DeleteManyResponse, *resty.Response, error)

	// DeleteByID выполняет запрос на удаление проекта по ID.
	// Принимает контекст и ID проекта.
	// Возвращает «true» в случае успешного удаления проекта.
	DeleteByID(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// Delete выполняет запрос на удаление проекта.
	// Принимает контекст и проект.
	// Возвращает «true» в случае успешного удаления проекта.
	Delete(ctx context.Context, entity *Project) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельного проекта по ID.
	// Принимает контекст, ID проекта и опционально объект параметров запроса Params.
	// Возвращает найденный проект.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Project, *resty.Response, error)

	// Update выполняет запрос на изменение проекта.
	// Принимает контекст, проект и опционально объект параметров запроса Params.
	// Возвращает изменённый проект.
	Update(ctx context.Context, id uuid.UUID, project *Project, params ...*Params) (*Project, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных проектов.
	// Принимает контекст.
	// Возвращает объект метаданных MetaAttributesSharedWrapper.
	GetMetadata(ctx context.Context) (*MetaAttributesSharedWrapper, *resty.Response, error)

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

	// GetNamedFilterList выполняет запрос на получение списка фильтров.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetNamedFilterList(ctx context.Context, params ...*Params) (*List[NamedFilter], *resty.Response, error)

	// GetNamedFilterByID выполняет запрос на получение отдельного фильтра по ID.
	// Принимает контекст и ID фильтра.
	// Возвращает найденный фильтр.
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)
}

const (
	EndpointProject = EndpointEntity + string(MetaTypeProject)
)

// NewProjectService принимает [Client] и возвращает сервис для работы с проектами.
func NewProjectService(client *Client) ProjectService {
	return newMainService[Project, any, MetaAttributesSharedWrapper, any](client, EndpointProject)
}
