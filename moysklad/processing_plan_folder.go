package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// ProcessingPlanFolder Группа тех. карт.
// Ключевое слово: processingplanfolder
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-gruppa-teh-kart
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

func (p ProcessingPlanFolder) String() string {
	return Stringify(p)
}

func (p ProcessingPlanFolder) MetaType() MetaType {
	return MetaTypeProcessingPlanFolder
}

// ProcessingPlanFolderService
// Сервис для работы с группами техкарт.
type ProcessingPlanFolderService interface {
	GetList(ctx context.Context, params *Params) (*List[ProcessingPlanFolder], *resty.Response, error)
	Create(ctx context.Context, processingPlanFolder *ProcessingPlanFolder, params *Params) (*ProcessingPlanFolder, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*ProcessingPlanFolder, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, processingPlanFolder *ProcessingPlanFolder, params *Params) (*ProcessingPlanFolder, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetaAttributesSharedStatesWrapper, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
	MoveToTrash(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewProcessingPlanFolderService(client *Client) ProcessingPlanFolderService {
	e := NewEndpoint(client, "entity/processingplanfolder")
	return newMainService[ProcessingPlanFolder, any, MetaAttributesSharedStatesWrapper, any](e)
}
