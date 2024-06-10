package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Project Проект.
// Ключевое слово: project
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-proekt
type Project struct {
	Group        *Group                `json:"group,omitempty"`
	Archived     *bool                 `json:"archived,omitempty"`
	Code         *string               `json:"code,omitempty"`
	Description  *string               `json:"description,omitempty"`
	ExternalCode *string               `json:"externalCode,omitempty"`
	AccountID    *uuid.UUID            `json:"accountId,omitempty"`
	ID           *uuid.UUID            `json:"id,omitempty"`
	Meta         *Meta                 `json:"meta,omitempty"`
	Name         *string               `json:"name,omitempty"`
	Owner        *Employee             `json:"owner,omitempty"`
	Shared       *bool                 `json:"shared,omitempty"`
	Updated      *Timestamp            `json:"updated,omitempty"`
	Attributes   Slice[AttributeValue] `json:"attributes,omitempty"`
}

func (project Project) Clean() *Project {
	return &Project{Meta: project.Meta}
}

func (project Project) GetGroup() Group {
	return Deref(project.Group)
}

func (project Project) GetArchived() bool {
	return Deref(project.Archived)
}

func (project Project) GetCode() string {
	return Deref(project.Code)
}

func (project Project) GetDescription() string {
	return Deref(project.Description)
}

func (project Project) GetExternalCode() string {
	return Deref(project.ExternalCode)
}

func (project Project) GetAccountID() uuid.UUID {
	return Deref(project.AccountID)
}

func (project Project) GetID() uuid.UUID {
	return Deref(project.ID)
}

func (project Project) GetMeta() Meta {
	return Deref(project.Meta)
}

func (project Project) GetName() string {
	return Deref(project.Name)
}

func (project Project) GetOwner() Employee {
	return Deref(project.Owner)
}

func (project Project) GetShared() bool {
	return Deref(project.Shared)
}

func (project Project) GetUpdated() Timestamp {
	return Deref(project.Updated)
}

func (project Project) GetAttributes() Slice[AttributeValue] {
	return project.Attributes
}

func (project *Project) SetGroup(group *Group) *Project {
	project.Group = group.Clean()
	return project
}

func (project *Project) SetArchived(archived bool) *Project {
	project.Archived = &archived
	return project
}

func (project *Project) SetCode(code string) *Project {
	project.Code = &code
	return project
}

func (project *Project) SetDescription(description string) *Project {
	project.Description = &description
	return project
}

func (project *Project) SetExternalCode(externalCode string) *Project {
	project.ExternalCode = &externalCode
	return project
}

func (project *Project) SetMeta(meta *Meta) *Project {
	project.Meta = meta
	return project
}

func (project *Project) SetName(name string) *Project {
	project.Name = &name
	return project
}

func (project *Project) SetOwner(owner *Employee) *Project {
	project.Owner = owner.Clean()
	return project
}

func (project *Project) SetShared(shared bool) *Project {
	project.Shared = &shared
	return project
}

func (project *Project) SetAttributes(attributes Slice[AttributeValue]) *Project {
	project.Attributes = attributes
	return project
}

func (project Project) String() string {
	return Stringify(project)
}

func (project Project) MetaType() MetaType {
	return MetaTypeProject
}

// ProjectService
// Сервис для работы с проектами.
type ProjectService interface {
	GetList(ctx context.Context, params *Params) (*List[Project], *resty.Response, error)
	Create(ctx context.Context, project *Project, params *Params) (*Project, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, projectList []*Project, params *Params) (*[]Project, *resty.Response, error)
	DeleteMany(ctx context.Context, projectList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params *Params) (*Project, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, project *Project, params *Params) (*Project, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetaAttributesSharedWrapper, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error)
	UpdateAttribute(ctx context.Context, id uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)
}

func NewProjectService(client *Client) ProjectService {
	e := NewEndpoint(client, "entity/project")
	return newMainService[Project, any, MetaAttributesSharedWrapper, any](e)
}
