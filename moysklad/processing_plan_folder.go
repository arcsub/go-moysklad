package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// ProcessingPlanFolder Группа тех. карт.
// Ключевое слово: processingplanfolder
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-gruppa-tehkart-gruppy-tehkart
type ProcessingPlanFolder struct {
	AccountID    *uuid.UUID `json:"accountId,omitempty"`    // ID учетной записи
	Archived     *bool      `json:"archived,omitempty"`     // Добавлена ли Группа тех. карт в архив
	ExternalCode *string    `json:"externalCode,omitempty"` // Внешний код Группы тех. карт
	Code         *string    `json:"code,omitempty"`         // Код Группы тех. карт
	Description  *string    `json:"description,omitempty"`  // Описание Группы тех. карт
	Group        *Group     `json:"group,omitempty"`        // Отдел сотрудника
	ID           *uuid.UUID `json:"id,omitempty"`           // ID Группы тех. карт
	Meta         *Meta      `json:"meta,omitempty"`         // Метаданные
	Name         *string    `json:"name,omitempty"`         // Наименование
	Owner        *Employee  `json:"owner,omitempty"`        // Владелец (Сотрудник)
	PathName     *string    `json:"pathName,omitempty"`     // Наименование Группы тех. карт, в которую входит данная Группа тех. карт
	Shared       *bool      `json:"shared,omitempty"`       // Общий доступ
	Updated      *Timestamp `json:"updated,omitempty"`      // Момент последнего обновления сущности
}

func (processingPlanFolder ProcessingPlanFolder) Clean() *ProcessingPlanFolder {
	return &ProcessingPlanFolder{Meta: processingPlanFolder.Meta}
}

func (processingPlanFolder ProcessingPlanFolder) GetAccountID() uuid.UUID {
	return Deref(processingPlanFolder.AccountID)
}

func (processingPlanFolder ProcessingPlanFolder) GetArchived() bool {
	return Deref(processingPlanFolder.Archived)
}

func (processingPlanFolder ProcessingPlanFolder) GetExternalCode() string {
	return Deref(processingPlanFolder.ExternalCode)
}

func (processingPlanFolder ProcessingPlanFolder) GetCode() string {
	return Deref(processingPlanFolder.Code)
}

func (processingPlanFolder ProcessingPlanFolder) GetDescription() string {
	return Deref(processingPlanFolder.Description)
}

func (processingPlanFolder ProcessingPlanFolder) GetGroup() Group {
	return Deref(processingPlanFolder.Group)
}

func (processingPlanFolder ProcessingPlanFolder) GetID() uuid.UUID {
	return Deref(processingPlanFolder.ID)
}

func (processingPlanFolder ProcessingPlanFolder) GetMeta() Meta {
	return Deref(processingPlanFolder.Meta)
}

func (processingPlanFolder ProcessingPlanFolder) GetName() string {
	return Deref(processingPlanFolder.Name)
}

func (processingPlanFolder ProcessingPlanFolder) GetOwner() Employee {
	return Deref(processingPlanFolder.Owner)
}

func (processingPlanFolder ProcessingPlanFolder) GetPathName() string {
	return Deref(processingPlanFolder.PathName)
}

func (processingPlanFolder ProcessingPlanFolder) GetShared() bool {
	return Deref(processingPlanFolder.Shared)
}

func (processingPlanFolder ProcessingPlanFolder) GetUpdated() Timestamp {
	return Deref(processingPlanFolder.Updated)
}

func (processingPlanFolder *ProcessingPlanFolder) SetArchived(archived bool) *ProcessingPlanFolder {
	processingPlanFolder.Archived = &archived
	return processingPlanFolder
}

func (processingPlanFolder *ProcessingPlanFolder) SetExternalCode(externalCode string) *ProcessingPlanFolder {
	processingPlanFolder.ExternalCode = &externalCode
	return processingPlanFolder
}

func (processingPlanFolder *ProcessingPlanFolder) SetCode(code string) *ProcessingPlanFolder {
	processingPlanFolder.Code = &code
	return processingPlanFolder
}

func (processingPlanFolder *ProcessingPlanFolder) SetDescription(description string) *ProcessingPlanFolder {
	processingPlanFolder.Description = &description
	return processingPlanFolder
}

func (processingPlanFolder *ProcessingPlanFolder) SetGroup(group *Group) *ProcessingPlanFolder {
	processingPlanFolder.Group = group.Clean()
	return processingPlanFolder
}

func (processingPlanFolder *ProcessingPlanFolder) SetMeta(meta *Meta) *ProcessingPlanFolder {
	processingPlanFolder.Meta = meta
	return processingPlanFolder
}

func (processingPlanFolder *ProcessingPlanFolder) SetName(name string) *ProcessingPlanFolder {
	processingPlanFolder.Name = &name
	return processingPlanFolder
}

func (processingPlanFolder *ProcessingPlanFolder) SetOwner(owner *Employee) *ProcessingPlanFolder {
	processingPlanFolder.Owner = owner.Clean()
	return processingPlanFolder
}

func (processingPlanFolder *ProcessingPlanFolder) SetShared(shared bool) *ProcessingPlanFolder {
	processingPlanFolder.Shared = &shared
	return processingPlanFolder
}

func (processingPlanFolder ProcessingPlanFolder) String() string {
	return Stringify(processingPlanFolder)
}

func (processingPlanFolder ProcessingPlanFolder) MetaType() MetaType {
	return MetaTypeProcessingPlanFolder
}

// ProcessingPlanFolderService
// Сервис для работы с группами техкарт.
type ProcessingPlanFolderService interface {
	GetList(ctx context.Context, params *Params) (*List[ProcessingPlanFolder], *resty.Response, error)
	Create(ctx context.Context, processingPlanFolder *ProcessingPlanFolder, params *Params) (*ProcessingPlanFolder, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params *Params) (*ProcessingPlanFolder, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, processingPlanFolder *ProcessingPlanFolder, params *Params) (*ProcessingPlanFolder, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetaAttributesSharedStatesWrapper, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)
	MoveToTrash(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
}

func NewProcessingPlanFolderService(client *Client) ProcessingPlanFolderService {
	e := NewEndpoint(client, "entity/processingplanfolder")
	return newMainService[ProcessingPlanFolder, any, MetaAttributesSharedStatesWrapper, any](e)
}
