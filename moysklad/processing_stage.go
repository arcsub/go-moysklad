package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"time"
)

// ProcessingStage Этап производства.
//
// Код сущности: processingstage
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-jetap-proizwodstwa
type ProcessingStage struct {
	AccountID     *uuid.UUID           `json:"accountId,omitempty"`     // ID учётной записи     // ID учётной записи
	AllPerformers *bool                `json:"allPerformers,omitempty"` // Признак доступности назначения на этап любого сотрудника
	Archived      *bool                `json:"archived,omitempty"`      // Добавлен ли Этап в архив
	Description   *string              `json:"description,omitempty"`   // Комментарий Этапа
	ExternalCode  *string              `json:"externalCode,omitempty"`  // Внешний код Этапа
	Group         *Group               `json:"group,omitempty"`         // Отдел сотрудника         // Отдел сотрудника
	ID            *uuid.UUID           `json:"id,omitempty"`            // ID Этапа
	Meta          *Meta                `json:"meta,omitempty"`          // Метаданные Этапа
	Name          *string              `json:"name,omitempty"`          // Наименование Этапа
	Owner         *Employee            `json:"owner,omitempty"`         // Метаданные владельца (Сотрудника)         // Владелец (Сотрудник)
	Performers    *MetaArray[Employee] `json:"performers,omitempty"`    // Метаданные возможных исполнителей
	Shared        *bool                `json:"shared,omitempty"`        // Общий доступ        // Общий доступ
	Updated       *Timestamp           `json:"updated,omitempty"`       // Момент последнего обновления сущности
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (processingStage ProcessingStage) Clean() *ProcessingStage {
	if processingStage.Meta == nil {
		return nil
	}
	return &ProcessingStage{Meta: processingStage.Meta}
}

func (processingStage ProcessingStage) GetAccountID() uuid.UUID {
	return Deref(processingStage.AccountID)
}

func (processingStage ProcessingStage) GetAllPerformers() bool {
	return Deref(processingStage.AllPerformers)
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

func (processingStage ProcessingStage) GetPerformers() MetaArray[Employee] {
	return Deref(processingStage.Performers)
}

func (processingStage ProcessingStage) GetShared() bool {
	return Deref(processingStage.Shared)
}

func (processingStage ProcessingStage) GetUpdated() time.Time {
	return Deref(processingStage.Updated).Time()
}

func (processingStage *ProcessingStage) SetAllPerformers(allPerformers bool) *ProcessingStage {
	processingStage.AllPerformers = &allPerformers
	return processingStage
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
	processingStage.Group = group.Clean()
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
	processingStage.Owner = owner.Clean()
	return processingStage
}

func (processingStage *ProcessingStage) SetPerformers(performers ...*Employee) *ProcessingStage {
	processingStage.Performers = NewMetaArrayFrom(performers)
	return processingStage
}

func (processingStage *ProcessingStage) SetShared(shared bool) *ProcessingStage {
	processingStage.Shared = &shared
	return processingStage
}

func (processingStage ProcessingStage) String() string {
	return Stringify(processingStage)
}

// MetaType возвращает код сущности.
func (ProcessingStage) MetaType() MetaType {
	return MetaTypeProcessingStage
}

// Update shortcut
func (processingStage *ProcessingStage) Update(ctx context.Context, client *Client, params ...*Params) (*ProcessingStage, *resty.Response, error) {
	return NewProcessingStageService(client).Update(ctx, processingStage.GetID(), processingStage, params...)
}

// Create shortcut
func (processingStage *ProcessingStage) Create(ctx context.Context, client *Client, params ...*Params) (*ProcessingStage, *resty.Response, error) {
	return NewProcessingStageService(client).Create(ctx, processingStage, params...)
}

// Delete shortcut
func (processingStage *ProcessingStage) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewProcessingStageService(client).Delete(ctx, processingStage)
}

// ProcessingStageService
// Сервис для работы с этапами производства.
type ProcessingStageService interface {
	GetList(ctx context.Context, params ...*Params) (*List[ProcessingStage], *resty.Response, error)
	Create(ctx context.Context, processingStage *ProcessingStage, params ...*Params) (*ProcessingStage, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, processingStageList Slice[ProcessingStage], params ...*Params) (*Slice[ProcessingStage], *resty.Response, error)
	DeleteMany(ctx context.Context, entities ...*ProcessingStage) (*DeleteManyResponse, *resty.Response, error)
	DeleteByID(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// Delete выполняет запрос на удаление этапа производства.
	// Принимает контекст и этап производства.
	// Возвращает «true» в случае успешного удаления этапа производства.
	Delete(ctx context.Context, entity *ProcessingStage) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*ProcessingStage, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, processingStage *ProcessingStage, params ...*Params) (*ProcessingStage, *resty.Response, error)

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
}

const (
	EndpointProcessingStage = EndpointEntity + string(MetaTypeProcessingStage)
)

func NewProcessingStageService(client *Client) ProcessingStageService {
	return newMainService[ProcessingStage, any, any, any](client, EndpointProcessingStage)
}
