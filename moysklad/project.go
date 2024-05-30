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
	AccountID    *uuid.UUID  `json:"accountId,omitempty"`    // ID учетной записи
	Archived     *bool       `json:"archived,omitempty"`     // Добавлен ли Проект в архив
	Attributes   *Attributes `json:"attributes,omitempty"`   // Коллекция доп. полей
	Code         *string     `json:"code,omitempty"`         // Код Проекта
	Description  *string     `json:"description,omitempty"`  // Описание Проекта
	ExternalCode *string     `json:"externalCode,omitempty"` // Внешний код Проекта
	Group        *Group      `json:"group,omitempty"`        // Метаданные отдела сотрудника
	ID           *uuid.UUID  `json:"id,omitempty"`           // ID сущности
	Meta         *Meta       `json:"meta,omitempty"`         // Метаданные
	Name         *string     `json:"name,omitempty"`         // Наименование Проекта
	Owner        *Employee   `json:"owner,omitempty"`        // Метаданные владельца (Сотрудника)
	Shared       *bool       `json:"shared,omitempty"`       // Общий доступ
	Updated      *Timestamp  `json:"updated,omitempty"`      // Момент последнего обновления сущности
}

func (p Project) String() string {
	return Stringify(p)
}

func (p Project) MetaType() MetaType {
	return MetaTypeProject
}

// ProjectService
// Сервис для работы с проектами.
type ProjectService interface {
	GetList(ctx context.Context, params *Params) (*List[Project], *resty.Response, error)
	Create(ctx context.Context, project *Project, params *Params) (*Project, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, projectList []*Project, params *Params) (*[]Project, *resty.Response, error)
	DeleteMany(ctx context.Context, projectList []*Project) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Project, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, project *Project, params *Params) (*Project, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetadataAttributeShared, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id *uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error)
	UpdateAttribute(ctx context.Context, id *uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList []*Attribute) (*DeleteManyResponse, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
}

func NewProjectService(client *Client) ProjectService {
	e := NewEndpoint(client, "entity/project")
	return newMainService[Project, any, MetadataAttributeShared, any](e)
}
