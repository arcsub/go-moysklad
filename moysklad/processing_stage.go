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

func (p ProcessingStage) String() string {
	return Stringify(p)
}

func (p ProcessingStage) MetaType() MetaType {
	return MetaTypeProcessingStage
}

// ProcessingStageService
// Сервис для работы с этапами производства.
type ProcessingStageService interface {
	GetList(ctx context.Context, params *Params) (*List[ProcessingStage], *resty.Response, error)
	Create(ctx context.Context, processingStage *ProcessingStage, params *Params) (*ProcessingStage, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, processingStageList []*ProcessingStage, params *Params) (*[]ProcessingStage, *resty.Response, error)
	DeleteMany(ctx context.Context, processingStageList []*ProcessingStage) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*ProcessingStage, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, processingStage *ProcessingStage, params *Params) (*ProcessingStage, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
	Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewProcessingStageService(client *Client) ProcessingStageService {
	e := NewEndpoint(client, "entity/processingstage")
	return newMainService[ProcessingStage, any, any, any](e)
}
