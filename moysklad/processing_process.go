package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"

	"time"
)

// ProcessingProcess Техпроцесс.
//
// Код сущности: processingprocess
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tehprocess
type ProcessingProcess struct {
	AccountID    *string                               `json:"accountId,omitempty"`    // ID учётной записи
	Archived     *bool                                 `json:"archived,omitempty"`     // Добавлен ли Тех. процесс в архив
	Description  *string                               `json:"description,omitempty"`  // Комментарий Тех. процесса
	ExternalCode *string                               `json:"externalCode,omitempty"` // Внешний код Тех. процесса
	Group        *Group                                `json:"group,omitempty"`        // Отдел сотрудника
	ID           *string                               `json:"id,omitempty"`           // ID Тех. процесса
	Meta         *Meta                                 `json:"meta,omitempty"`         // Метаданные Тех. процесса
	Name         *string                               `json:"name,omitempty"`         // Наименование Тех. процесса
	Owner        *Employee                             `json:"owner,omitempty"`        // Метаданные владельца (Сотрудника)        // Владелец (Сотрудник)
	Positions    *MetaArray[ProcessingProcessPosition] `json:"positions,omitempty"`    // Метаданные позиций Тех. процесса
	Shared       *bool                                 `json:"shared,omitempty"`       // Общий доступ       // Общий доступ
	Updated      *Timestamp                            `json:"updated,omitempty"`      // Момент последнего обновления сущности
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (processingProcess ProcessingProcess) Clean() *ProcessingProcess {
	if processingProcess.Meta == nil {
		return nil
	}
	return &ProcessingProcess{Meta: processingProcess.Meta}
}

func (processingProcess ProcessingProcess) GetAccountID() string {
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

func (processingProcess ProcessingProcess) GetID() string {
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

func (processingProcess ProcessingProcess) GetPositions() MetaArray[ProcessingProcessPosition] {
	return Deref(processingProcess.Positions)
}

func (processingProcess ProcessingProcess) GetShared() bool {
	return Deref(processingProcess.Shared)
}

func (processingProcess ProcessingProcess) GetUpdated() time.Time {
	return Deref(processingProcess.Updated).Time()
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

func (processingProcess *ProcessingProcess) SetPositions(positions ...*ProcessingProcessPosition) *ProcessingProcess {
	processingProcess.Positions = NewMetaArrayFrom(positions)
	return processingProcess
}

func (processingProcess *ProcessingProcess) SetShared(shared bool) *ProcessingProcess {
	processingProcess.Shared = &shared
	return processingProcess
}

func (processingProcess ProcessingProcess) String() string {
	return Stringify(processingProcess)
}

// MetaType возвращает код сущности.
func (ProcessingProcess) MetaType() MetaType {
	return MetaTypeProcessingProcess
}

// Update shortcut
func (processingProcess *ProcessingProcess) Update(ctx context.Context, client *Client, params ...*Params) (*ProcessingProcess, *resty.Response, error) {
	return NewProcessingProcessService(client).Update(ctx, processingProcess.GetID(), processingProcess, params...)
}

// Create shortcut
func (processingProcess *ProcessingProcess) Create(ctx context.Context, client *Client, params ...*Params) (*ProcessingProcess, *resty.Response, error) {
	return NewProcessingProcessService(client).Create(ctx, processingProcess, params...)
}

// Delete shortcut
func (processingProcess *ProcessingProcess) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewProcessingProcessService(client).Delete(ctx, processingProcess)
}

// ProcessingProcessPosition Позиция Тех. процесса.
//
// Код сущности: processingprocessposition
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tehprocess-pozicii-tehprocessa
type ProcessingProcessPosition struct {
	AccountID       *string                               `json:"accountId,omitempty"`       // ID учётной записи
	ID              *string                               `json:"id,omitempty"`              // ID позиции
	Meta            *Meta                                 `json:"meta,omitempty"`            // Метаданные позиции Тех. процесса
	ProcessingStage *ProcessingStage                      `json:"processingstage,omitempty"` // Метаданные этапа, который представляет собой позиция
	NextPositions   *MetaArray[ProcessingProcessPosition] `json:"nextPositions,omitempty"`   // Метаданные следующих позиций позиции Техпроцесса
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (processingProcessPosition ProcessingProcessPosition) Clean() *ProcessingProcessPosition {
	if processingProcessPosition.Meta == nil {
		return nil
	}
	return &ProcessingProcessPosition{Meta: processingProcessPosition.Meta}
}

func (processingProcessPosition ProcessingProcessPosition) GetAccountID() string {
	return Deref(processingProcessPosition.AccountID)
}

func (processingProcessPosition ProcessingProcessPosition) GetID() string {
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

// MetaType возвращает код сущности.
func (ProcessingProcessPosition) MetaType() MetaType {
	return MetaTypeProcessingProcessPosition
}

// ProcessingProcessService
// Сервис для работы с тех процессами.
type ProcessingProcessService interface {
	GetList(ctx context.Context, params ...*Params) (*List[ProcessingProcess], *resty.Response, error)
	Create(ctx context.Context, processingProcess *ProcessingProcess, params ...*Params) (*ProcessingProcess, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, processingProcessList Slice[ProcessingProcess], params ...*Params) (*Slice[ProcessingProcess], *resty.Response, error)
	DeleteMany(ctx context.Context, entities ...*ProcessingProcess) (*DeleteManyResponse, *resty.Response, error)
	DeleteByID(ctx context.Context, id string) (bool, *resty.Response, error)

	// Delete выполняет запрос на удаление тех процесса.
	// Принимает контекст и тех процесс.
	// Возвращает «true» в случае успешного удаления тех процесса.
	Delete(ctx context.Context, entity *ProcessingProcess) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id string, params ...*Params) (*ProcessingProcess, *resty.Response, error)
	Update(ctx context.Context, id string, processingProcess *ProcessingProcess, params ...*Params) (*ProcessingProcess, *resty.Response, error)

	// GetPositionList выполняет запрос на получение списка позиций документа.
	// Принимает контекст, ID документа и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetPositionList(ctx context.Context, id string, params ...*Params) (*List[ProcessingProcessPosition], *resty.Response, error)

	GetPositionListAll(ctx context.Context, id string, params ...*Params) (*Slice[ProcessingProcessPosition], *resty.Response, error)

	// GetPositionByID выполняет запрос на получение отдельной позиции документа по ID.
	// Принимает контекст, ID документа, ID позиции и опционально объект параметров запроса Params.
	// Возвращает найденную позицию.
	GetPositionByID(ctx context.Context, id string, positionID string, params ...*Params) (*ProcessingProcessPosition, *resty.Response, error)

	// UpdatePosition выполняет запрос на изменение позиции документа.
	// Принимает контекст, ID документа, ID позиции, позицию документа и опционально объект параметров запроса Params.
	// Возвращает изменённую позицию.
	UpdatePosition(ctx context.Context, id string, positionID string, position *ProcessingProcessPosition, params ...*Params) (*ProcessingProcessPosition, *resty.Response, error)

	// CreatePosition выполняет запрос на добавление позиции документа.
	// Принимает контекст, ID документа, позицию документа и опционально объект параметров запроса Params.
	// Возвращает добавленную позицию.
	CreatePosition(ctx context.Context, id string, position *ProcessingProcessPosition, params ...*Params) (*ProcessingProcessPosition, *resty.Response, error)

	// CreatePositionMany выполняет запрос на массовое добавление позиций документа.
	// Принимает контекст, ID документа и множество позиций.
	// Возвращает список добавленных позиций.
	CreatePositionMany(ctx context.Context, id string, positions ...*ProcessingProcessPosition) (*Slice[ProcessingProcessPosition], *resty.Response, error)

	// DeletePosition выполняет запрос на удаление позиции документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает «true» в случае успешного удаления позиции.
	DeletePosition(ctx context.Context, id string, positionID string) (bool, *resty.Response, error)

	// DeletePositionMany выполняет запрос на массовое удаление позиций документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeletePositionMany(ctx context.Context, id string, positions ...*ProcessingProcessPosition) (*DeleteManyResponse, *resty.Response, error)

	// GetPositionTrackingCodeList выполняет запрос на получение кодов маркировки позиции документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает объект List.
	GetPositionTrackingCodeList(ctx context.Context, id string, positionID string) (*List[TrackingCode], *resty.Response, error)

	// CreateUpdatePositionTrackingCodeMany выполняет запрос на массовое создание/изменение кодов маркировки позиции документа.
	// Принимает контекст, ID документа, ID позиции и множество кодов маркировки.
	// Возвращает список созданных и/или изменённых кодов маркировки позиции документа.
	CreateUpdatePositionTrackingCodeMany(ctx context.Context, id string, positionID string, trackingCodes ...*TrackingCode) (*Slice[TrackingCode], *resty.Response, error)

	// DeletePositionTrackingCodeMany выполняет запрос на массовое удаление кодов маркировки позиции документа.
	// Принимает контекст, ID документа, ID позиции и множество кодов маркировки.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeletePositionTrackingCodeMany(ctx context.Context, id string, positionID string, trackingCodes ...*TrackingCode) (*DeleteManyResponse, *resty.Response, error)

	// GetNamedFilterList выполняет запрос на получение списка фильтров.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetNamedFilterList(ctx context.Context, params ...*Params) (*List[NamedFilter], *resty.Response, error)

	// GetNamedFilterByID выполняет запрос на получение отдельного фильтра по ID.
	// Принимает контекст и ID фильтра.
	// Возвращает найденный фильтр.
	GetNamedFilterByID(ctx context.Context, id string) (*NamedFilter, *resty.Response, error)

	// MoveToTrash выполняет запрос на перемещение документа с указанным ID в корзину.
	// Принимает контекст и ID документа.
	// Возвращает «true» в случае успешного перемещения в корзину.
	MoveToTrash(ctx context.Context, id string) (bool, *resty.Response, error)
}

const (
	EndpointProcessingProcess = EndpointEntity + string(MetaTypeProcessingProcess)
)

func NewProcessingProcessService(client *Client) ProcessingProcessService {
	return newMainService[ProcessingProcess, ProcessingProcessPosition, any, any](client, EndpointProcessingProcess)
}
