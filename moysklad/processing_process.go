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

func (p ProcessingProcess) String() string {
	return Stringify(p)
}

func (p ProcessingProcess) MetaType() MetaType {
	return MetaTypeProcessingProcess
}

// ProcessingProcessPosition Позиция Тех. процесса.
// Ключевое слово: processingprocessposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-process-teh-processy-atributy-wlozhennyh-suschnostej-pozicii-teh-processa
type ProcessingProcessPosition struct {
	AccountID       *uuid.UUID       `json:"accountId,omitempty"`       // ID учетной записи
	ID              *uuid.UUID       `json:"id,omitempty"`              // ID позиции
	Meta            *Meta            `json:"meta,omitempty"`            // Метаданные позиции Тех. процесса
	ProcessingStage *ProcessingStage `json:"processingstage,omitempty"` // Метаданные этапа, который представляет собой позиция
}

func (p ProcessingProcessPosition) String() string {
	return Stringify(p)
}

func (p ProcessingProcessPosition) MetaType() MetaType {
	return MetaTypeProcessingProcessPosition
}

// ProcessingProcessService
// Сервис для работы с тех процессами.
type ProcessingProcessService interface {
	GetList(ctx context.Context, params *Params) (*List[ProcessingProcess], *resty.Response, error)
	Create(ctx context.Context, processingProcess *ProcessingProcess, params *Params) (*ProcessingProcess, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, processingProcessList []*ProcessingProcess, params *Params) (*[]ProcessingProcess, *resty.Response, error)
	DeleteMany(ctx context.Context, processingProcessList []*ProcessingProcess) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*ProcessingProcess, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, processingProcess *ProcessingProcess, params *Params) (*ProcessingProcess, *resty.Response, error)
	GetPositions(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[ProcessingProcessPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, params *Params) (*ProcessingProcessPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, position *ProcessingProcessPosition, params *Params) (*ProcessingProcessPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id *uuid.UUID, position *ProcessingProcessPosition) (*ProcessingProcessPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id *uuid.UUID, positions []*ProcessingProcessPosition) (*[]ProcessingProcessPosition, *resty.Response, error)
	DeletePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID) (bool, *resty.Response, error)
	GetPositionTrackingCodes(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID) (*MetaArray[TrackingCode], *resty.Response, error)
	CreateOrUpdatePositionTrackingCodes(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, trackingCodes TrackingCodes) (*[]TrackingCode, *resty.Response, error)
	DeletePositionTrackingCodes(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, trackingCodes TrackingCodes) (*DeleteManyResponse, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
	Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewProcessingProcessService(client *Client) ProcessingProcessService {
	e := NewEndpoint(client, "entity/processingprocess")
	return newMainService[ProcessingProcess, ProcessingProcessPosition, any, any](e)
}
