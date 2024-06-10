package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// ProcessingStage Этап производства.
// Ключевое слово: processingstage
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-jetap-proizwodstwa
type ProcessingStage struct {
	AccountID    *uuid.UUID `json:"accountId,omitempty"`    // ID учетной записи
	Archived     *bool      `json:"archived,omitempty"`     // Добавлен ли Этап в архив
	Description  *string    `json:"description,omitempty"`  // Комментарий Этапа
	ExternalCode *string    `json:"externalCode,omitempty"` // Внешний код Этапа
	Group        *Group     `json:"group,omitempty"`        // Отдел сотрудника
	ID           *uuid.UUID `json:"id,omitempty"`           // ID Этапа
	Meta         *Meta      `json:"meta,omitempty"`         // Метаданные Этапа
	Name         *string    `json:"name,omitempty"`         // Наименование Этапа
	Owner        *Employee  `json:"owner,omitempty"`        // Владелец (Сотрудник)
	Shared       *bool      `json:"shared,omitempty"`       // Общий доступ
	Updated      *Timestamp `json:"updated,omitempty"`      // Момент последнего обновления сущности
}

func (processingStage ProcessingStage) Clean() *ProcessingStage {
	return &ProcessingStage{Meta: processingStage.Meta}
}

func (processingStage ProcessingStage) GetAccountID() uuid.UUID {
	return Deref(processingStage.AccountID)
}

func (processingStage ProcessingStage) GetArchived() bool {
	return Deref(processingStage.Archived)
}

func (processingStage ProcessingStage) GetDescription() string {
	return Deref(processingStage.Description)
}

func (processingStage ProcessingStage) GetExternalCode() string {
	return Deref(processingStage.ExternalCode)
}

func (processingStage ProcessingStage) GetGroup() Group {
	return Deref(processingStage.Group)
}

func (processingStage ProcessingStage) GetID() uuid.UUID {
	return Deref(processingStage.ID)
}

func (processingStage ProcessingStage) GetMeta() Meta {
	return Deref(processingStage.Meta)
}

func (processingStage ProcessingStage) GetName() string {
	return Deref(processingStage.Name)
}

func (processingStage ProcessingStage) GetOwner() Employee {
	return Deref(processingStage.Owner)
}

func (processingStage ProcessingStage) GetShared() bool {
	return Deref(processingStage.Shared)
}

func (processingStage ProcessingStage) GetUpdated() Timestamp {
	return Deref(processingStage.Updated)
}

func (processingStage *ProcessingStage) SetArchived(archived bool) *ProcessingStage {
	processingStage.Archived = &archived
	return processingStage
}

func (processingStage *ProcessingStage) SetDescription(description string) *ProcessingStage {
	processingStage.Description = &description
	return processingStage
}

func (processingStage *ProcessingStage) SetExternalCode(externalCode string) *ProcessingStage {
	processingStage.ExternalCode = &externalCode
	return processingStage
}

func (processingStage *ProcessingStage) SetGroup(group *Group) *ProcessingStage {
	processingStage.Group = group
	return processingStage
}

func (processingStage *ProcessingStage) SetMeta(maeta *Meta) *ProcessingStage {
	processingStage.Meta = maeta
	return processingStage
}

func (processingStage *ProcessingStage) SetName(name string) *ProcessingStage {
	processingStage.Name = &name
	return processingStage
}

func (processingStage *ProcessingStage) SetOwner(owner *Employee) *ProcessingStage {
	processingStage.Owner = owner
	return processingStage
}

func (processingStage *ProcessingStage) SetShared(shared bool) *ProcessingStage {
	processingStage.Shared = &shared
	return processingStage
}

func (processingStage ProcessingStage) String() string {
	return Stringify(processingStage)
}

func (processingStage ProcessingStage) MetaType() MetaType {
	return MetaTypeProcessingStage
}

// ProcessingStageService
// Сервис для работы с этапами производства.
type ProcessingStageService interface {
	GetList(ctx context.Context, params *Params) (*List[ProcessingStage], *resty.Response, error)
	Create(ctx context.Context, processingStage *ProcessingStage, params *Params) (*ProcessingStage, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, processingStageList []*ProcessingStage, params *Params) (*[]ProcessingStage, *resty.Response, error)
	DeleteMany(ctx context.Context, processingStageList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params *Params) (*ProcessingStage, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, processingStage *ProcessingStage, params *Params) (*ProcessingStage, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)
	MoveToTrash(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
}

func NewProcessingStageService(client *Client) ProcessingStageService {
	e := NewEndpoint(client, "entity/processingstage")
	return newMainService[ProcessingStage, any, any, any](e)
}
