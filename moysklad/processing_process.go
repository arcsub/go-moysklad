package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// ProcessingProcess Техпроцесс.
// Ключевое слово: processingprocess
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tehprocess
type ProcessingProcess struct {
	AccountID    *uuid.UUID                            `json:"accountId,omitempty"`    // ID учетной записи
	Archived     *bool                                 `json:"archived,omitempty"`     // Добавлен ли Тех. процесс в архив
	Description  *string                               `json:"description,omitempty"`  // Комментарий Тех. процесса
	ExternalCode *string                               `json:"externalCode,omitempty"` // Внешний код Тех. процесса
	Group        *Group                                `json:"group,omitempty"`        // Отдел сотрудника
	ID           *uuid.UUID                            `json:"id,omitempty"`           // ID Тех. процесса
	Meta         *Meta                                 `json:"meta,omitempty"`         // Метаданные Тех. процесса
	Name         *string                               `json:"name,omitempty"`         // Наименование Тех. процесса
	Owner        *Employee                             `json:"owner,omitempty"`        // Владелец (Сотрудник)
	Positions    *Positions[ProcessingProcessPosition] `json:"positions,omitempty"`    // Метаданные позиций Тех. процесса
	Shared       *bool                                 `json:"shared,omitempty"`       // Общий доступ
	Updated      *Timestamp                            `json:"updated,omitempty"`      // Момент последнего обновления сущности
}

// Clean возвращает сущность с единственным заполненным полем Meta
func (processingProcess ProcessingProcess) Clean() *ProcessingProcess {
	return &ProcessingProcess{Meta: processingProcess.Meta}
}

func (processingProcess ProcessingProcess) GetAccountID() uuid.UUID {
	return Deref(processingProcess.AccountID)
}

func (processingProcess ProcessingProcess) GetArchived() bool {
	return Deref(processingProcess.Archived)
}

func (processingProcess ProcessingProcess) GetDescription() string {
	return Deref(processingProcess.Description)
}

func (processingProcess ProcessingProcess) GetExternalCode() string {
	return Deref(processingProcess.ExternalCode)
}

func (processingProcess ProcessingProcess) GetGroup() Group {
	return Deref(processingProcess.Group)
}

func (processingProcess ProcessingProcess) GetID() uuid.UUID {
	return Deref(processingProcess.ID)
}

func (processingProcess ProcessingProcess) GetMeta() Meta {
	return Deref(processingProcess.Meta)
}

func (processingProcess ProcessingProcess) GetName() string {
	return Deref(processingProcess.Name)
}

func (processingProcess ProcessingProcess) GetOwner() Employee {
	return Deref(processingProcess.Owner)
}

func (processingProcess ProcessingProcess) GetPositions() Positions[ProcessingProcessPosition] {
	return Deref(processingProcess.Positions)
}

func (processingProcess ProcessingProcess) GetShared() bool {
	return Deref(processingProcess.Shared)
}

func (processingProcess ProcessingProcess) GetUpdated() Timestamp {
	return Deref(processingProcess.Updated)
}

func (processingProcess *ProcessingProcess) SetArchived(archived bool) *ProcessingProcess {
	processingProcess.Archived = &archived
	return processingProcess
}

func (processingProcess *ProcessingProcess) SetDescription(description string) *ProcessingProcess {
	processingProcess.Description = &description
	return processingProcess
}

func (processingProcess *ProcessingProcess) SetExternalCode(externalCode string) *ProcessingProcess {
	processingProcess.ExternalCode = &externalCode
	return processingProcess
}

func (processingProcess *ProcessingProcess) SetGroup(group *Group) *ProcessingProcess {
	processingProcess.Group = group.Clean()
	return processingProcess
}

func (processingProcess *ProcessingProcess) SetMeta(meta *Meta) *ProcessingProcess {
	processingProcess.Meta = meta
	return processingProcess
}

func (processingProcess *ProcessingProcess) SetName(name string) *ProcessingProcess {
	processingProcess.Name = &name
	return processingProcess
}

func (processingProcess *ProcessingProcess) SetOwner(owner *Employee) *ProcessingProcess {
	processingProcess.Owner = owner.Clean()
	return processingProcess
}

func (processingProcess *ProcessingProcess) SetPositions(positions *Positions[ProcessingProcessPosition]) *ProcessingProcess {
	processingProcess.Positions = positions
	return processingProcess
}

func (processingProcess *ProcessingProcess) SetShared(shared bool) *ProcessingProcess {
	processingProcess.Shared = &shared
	return processingProcess
}

func (processingProcess ProcessingProcess) String() string {
	return Stringify(processingProcess)
}

func (processingProcess ProcessingProcess) MetaType() MetaType {
	return MetaTypeProcessingProcess
}

// Update shortcut
func (processingProcess ProcessingProcess) Update(ctx context.Context, client *Client, params ...*Params) (*ProcessingProcess, *resty.Response, error) {
	return client.Entity().ProcessingProcess().Update(ctx, processingProcess.GetID(), &processingProcess, params...)
}

// Create shortcut
func (processingProcess ProcessingProcess) Create(ctx context.Context, client *Client, params ...*Params) (*ProcessingProcess, *resty.Response, error) {
	return client.Entity().ProcessingProcess().Create(ctx, &processingProcess, params...)
}

// Delete shortcut
func (processingProcess ProcessingProcess) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return client.Entity().ProcessingProcess().Delete(ctx, processingProcess.GetID())
}

// ProcessingProcessPosition Позиция Тех. процесса.
// Ключевое слово: processingprocessposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tehprocess-pozicii-tehprocessa
type ProcessingProcessPosition struct {
	AccountID       *uuid.UUID                            `json:"accountId,omitempty"`       // ID учетной записи
	ID              *uuid.UUID                            `json:"id,omitempty"`              // ID позиции
	Meta            *Meta                                 `json:"meta,omitempty"`            // Метаданные позиции Тех. процесса
	ProcessingStage *ProcessingStage                      `json:"processingstage,omitempty"` // Метаданные этапа, который представляет собой позиция
	NextPositions   *MetaArray[ProcessingProcessPosition] `json:"nextPositions,omitempty"`   // Метаданные следующих позиций позиции Техпроцесса
}

// Clean возвращает сущность с единственным заполненным полем Meta
func (processingProcessPosition ProcessingProcessPosition) Clean() *ProcessingProcessPosition {
	return &ProcessingProcessPosition{Meta: processingProcessPosition.Meta}
}

func (processingProcessPosition ProcessingProcessPosition) GetAccountID() uuid.UUID {
	return Deref(processingProcessPosition.AccountID)
}

func (processingProcessPosition ProcessingProcessPosition) GetID() uuid.UUID {
	return Deref(processingProcessPosition.ID)
}

func (processingProcessPosition ProcessingProcessPosition) GetMeta() Meta {
	return Deref(processingProcessPosition.Meta)
}

func (processingProcessPosition ProcessingProcessPosition) GetProcessingStage() ProcessingStage {
	return Deref(processingProcessPosition.ProcessingStage)
}

func (processingProcessPosition ProcessingProcessPosition) GetNextPositions() MetaArray[ProcessingProcessPosition] {
	return Deref(processingProcessPosition.NextPositions)
}

func (processingProcessPosition *ProcessingProcessPosition) SetMeta(meta *Meta) *ProcessingProcessPosition {
	processingProcessPosition.Meta = meta
	return processingProcessPosition
}

func (processingProcessPosition *ProcessingProcessPosition) SetProcessingStage(processingStage *ProcessingStage) *ProcessingProcessPosition {
	processingProcessPosition.ProcessingStage = processingStage.Clean()
	return processingProcessPosition
}

func (processingProcessPosition ProcessingProcessPosition) String() string {
	return Stringify(processingProcessPosition)
}

func (processingProcessPosition ProcessingProcessPosition) MetaType() MetaType {
	return MetaTypeProcessingProcessPosition
}

// ProcessingProcessService
// Сервис для работы с тех процессами.
type ProcessingProcessService interface {
	GetList(ctx context.Context, params ...*Params) (*List[ProcessingProcess], *resty.Response, error)
	Create(ctx context.Context, processingProcess *ProcessingProcess, params ...*Params) (*ProcessingProcess, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, processingProcessList Slice[ProcessingProcess], params ...*Params) (*Slice[ProcessingProcess], *resty.Response, error)
	DeleteMany(ctx context.Context, entities ...*ProcessingProcess) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*ProcessingProcess, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, processingProcess *ProcessingProcess, params ...*Params) (*ProcessingProcess, *resty.Response, error)
	GetPositions(ctx context.Context, id uuid.UUID, params ...*Params) (*MetaArray[ProcessingProcessPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params ...*Params) (*ProcessingProcessPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *ProcessingProcessPosition, params ...*Params) (*ProcessingProcessPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id uuid.UUID, position *ProcessingProcessPosition) (*ProcessingProcessPosition, *resty.Response, error)
	CreatePositionMany(ctx context.Context, id uuid.UUID, positions ...*ProcessingProcessPosition) (*Slice[ProcessingProcessPosition], *resty.Response, error)
	DeletePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (bool, *resty.Response, error)
	DeletePositionMany(ctx context.Context, id uuid.UUID, entities ...*ProcessingProcessPosition) (*DeleteManyResponse, *resty.Response, error)
	GetPositionTrackingCodes(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (*MetaArray[TrackingCode], *resty.Response, error)
	CreateUpdatePositionTrackingCodeMany(ctx context.Context, id uuid.UUID, positionID uuid.UUID, trackingCodes ...*TrackingCode) (*Slice[TrackingCode], *resty.Response, error)
	DeletePositionTrackingCodeMany(ctx context.Context, id uuid.UUID, positionID uuid.UUID, trackingCodes ...*TrackingCode) (*DeleteManyResponse, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params ...*Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)
	MoveToTrash(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
}

func NewProcessingProcessService(client *Client) ProcessingProcessService {
	e := NewEndpoint(client, "entity/processingprocess")
	return newMainService[ProcessingProcess, ProcessingProcessPosition, any, any](e)
}
