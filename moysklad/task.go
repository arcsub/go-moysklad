package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Task Задача.
// TODO
// Ключевое слово: task
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-zadacha
type Task struct {
	Done              *bool            `json:"done,omitempty"`
	Created           *Timestamp       `json:"created,omitempty"`
	Assignee          *Employee        `json:"assignee,omitempty"`
	Author            *Employee        `json:"author,omitempty"`
	AccountID         *uuid.UUID       `json:"accountId,omitempty"`
	Completed         *Timestamp       `json:"completed,omitempty"`
	Agent             *Counterparty    `json:"agent,omitempty"`
	Description       *string          `json:"description,omitempty"`
	AuthorApplication *Application     `json:"authorApplication,omitempty"`
	DueToDate         *Timestamp       `json:"dueToDate,omitempty"`
	Files             *MetaArray[File] `json:"files,omitempty"`
	ID                *uuid.UUID       `json:"id,omitempty"`
	Implementer       *Employee        `json:"implementer,omitempty"`
	Meta              *Meta            `json:"meta,omitempty"`
	Updated           *Timestamp       `json:"updated,omitempty"`
	Notes             Slice[TaskNote]  `json:"notes,omitempty"`
}

func (t Task) String() string {
	return Stringify(t)
}

func (t Task) MetaType() MetaType {
	return MetaTypeTask
}

// TaskNote Комментарии задачи.
// Ключевое слово: tasknote
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-zadacha-zadachi-kommentarii-zadachi
type TaskNote struct {
	Author            *Meta      `json:"author,omitempty"`            // Метаданные Сотрудника, создавшего комментарий (администратор аккаунта, если автор - приложение)
	AuthorApplication *Meta      `json:"authorApplication,omitempty"` // Метаданные Приложения, создавшего комментарий
	Moment            *Timestamp `json:"moment,omitempty"`            // Момент создания комментария
	Description       *string    `json:"description,omitempty"`       // Текст комментария
}

func (t TaskNote) String() string {
	return Stringify(t)
}

func (t TaskNote) MetaType() MetaType {
	return MetaTypeTaskNote
}

//// TODO
//
//type OperationTask struct {
//	wrapper[operationTask]
//	OperationIn
//	OperationOut
//	OperationRetail
//}
//
//type operationTask struct {
//	Meta *Meta `json:"meta,omitempty"` // Метаданные
//	raw  json.RawMessage
//}
//
//func (o *OperationTask) UnmarshalJSON(data []byte) (err error) {
//	type alias OperationTask
//	var t alias
//
//	if err = json.Unmarshal(data, &t); err != nil {
//		return err
//	}
//
//	t.w.raw = data
//	*o = OperationTask(t)
//
//	return nil
//}
//
//func (o OperationTask) Data() json.RawMessage {
//	return o.w.raw
//}
//
//func (o OperationTask) Meta() Meta {
//	return utils.Deref[Meta](o.w.Meta)
//}
//
//func (o OperationTask) Counterparty() (Counterparty, bool) {
//	return ElementAsType[Counterparty](o)
//}
//
//func (o OperationTask) Organization() (Organization, bool) {
//	return ElementAsType[Organization](o)
//}

// TaskService
// Сервис для работы с задачами.
type TaskService interface {
	GetList(ctx context.Context, params ...*Params) (*List[Task], *resty.Response, error)
	Create(ctx context.Context, task *Task, params ...*Params) (*Task, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, taskList Slice[Task], params ...*Params) (*Slice[Task], *resty.Response, error)
	DeleteMany(ctx context.Context, taskList []MetaWrapper) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Task, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, task *Task, params ...*Params) (*Task, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params ...*Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)
	GetNotes(ctx context.Context, taskID uuid.UUID, params ...*Params) (*List[TaskNote], *resty.Response, error)
	CreateNote(ctx context.Context, taskID uuid.UUID, taskNote *TaskNote) (*TaskNote, *resty.Response, error)
	CreateNotes(ctx context.Context, taskID uuid.UUID, taskNotes Slice[TaskNote]) (*Slice[TaskNote], *resty.Response, error)
	GetNoteByID(ctx context.Context, taskID, taskNoteID uuid.UUID) (*TaskNote, *resty.Response, error)
	UpdateNote(ctx context.Context, taskID, taskNoteID uuid.UUID, taskNote *TaskNote) (*TaskNote, *resty.Response, error)
	DeleteNote(ctx context.Context, taskID, taskNoteID uuid.UUID) (bool, *resty.Response, error)
	GetFiles(ctx context.Context, id uuid.UUID) (*MetaArray[File], *resty.Response, error)
	CreateFile(ctx context.Context, id uuid.UUID, file *File) (*Slice[File], *resty.Response, error)
	UpdateFiles(ctx context.Context, id uuid.UUID, files Slice[File]) (*Slice[File], *resty.Response, error)
	DeleteFile(ctx context.Context, id uuid.UUID, fileID uuid.UUID) (bool, *resty.Response, error)
	DeleteFiles(ctx context.Context, id uuid.UUID, files []MetaWrapper) (*DeleteManyResponse, *resty.Response, error)
}

type taskService struct {
	Endpoint
	endpointGetList[Task]
	endpointCreate[Task]
	endpointCreateUpdateMany[Task]
	endpointDeleteMany[Task]
	endpointDelete
	endpointGetByID[Task]
	endpointUpdate[Task]
	endpointNamedFilter
	endpointFiles
}

func NewTaskService(client *Client) TaskService {
	e := NewEndpoint(client, "entity/task")
	return &taskService{
		Endpoint:                 e,
		endpointGetList:          endpointGetList[Task]{e},
		endpointCreate:           endpointCreate[Task]{e},
		endpointCreateUpdateMany: endpointCreateUpdateMany[Task]{e},
		endpointDeleteMany:       endpointDeleteMany[Task]{e},
		endpointDelete:           endpointDelete{e},
		endpointGetByID:          endpointGetByID[Task]{e},
		endpointUpdate:           endpointUpdate[Task]{e},
		endpointNamedFilter:      endpointNamedFilter{e},
		endpointFiles:            endpointFiles{e},
	}
}

// GetNotes Запрос на получение списка всех комментариев данной Задачи.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-zadacha-poluchit-kommentarii-zadachi
func (service *taskService) GetNotes(ctx context.Context, taskID uuid.UUID, params ...*Params) (*List[TaskNote], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes", service.uri, taskID)
	return NewRequestBuilder[List[TaskNote]](service.client, path).SetParams(params...).Get(ctx)
}

// CreateNote Запрос на создание нового комментария к Задаче.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-zadacha-sozdat-kommentarij-zadachi
func (service *taskService) CreateNote(ctx context.Context, taskID uuid.UUID, taskNote *TaskNote) (*TaskNote, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes", service.uri, taskID)
	return NewRequestBuilder[TaskNote](service.client, path).Post(ctx, taskNote)
}

// CreateNotes Запрос на создание нескольких комментариев к Задаче.
func (service *taskService) CreateNotes(ctx context.Context, taskID uuid.UUID, taskNotes Slice[TaskNote]) (*Slice[TaskNote], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes", service.uri, taskID)
	return NewRequestBuilder[Slice[TaskNote]](service.client, path).Post(ctx, taskNotes)
}

// GetNoteByID Отдельный комментарий к Задаче с указанным id комментария.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-zadacha-poluchit-kommentarij-k-zadache
func (service *taskService) GetNoteByID(ctx context.Context, taskID, taskNoteID uuid.UUID) (*TaskNote, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes/%s", service.uri, taskID, taskNoteID)
	return NewRequestBuilder[TaskNote](service.client, path).Get(ctx)
}

// UpdateNote Запрос на обновление отдельного комментария к Задаче.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-zadacha-izmenit-kommentarij-k-zadache
func (service *taskService) UpdateNote(ctx context.Context, taskID, taskNoteID uuid.UUID, taskNote *TaskNote) (*TaskNote, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes/%s", service.uri, taskID, taskNoteID)
	return NewRequestBuilder[TaskNote](service.client, path).Put(ctx, taskNote)
}

// DeleteNote Запрос на удаление отдельного комментария к Задаче с указанным id.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-zadacha-udalit-kommentarij
func (service *taskService) DeleteNote(ctx context.Context, taskID, taskNoteID uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/notes/%s", taskID, taskNoteID)
	return NewRequestBuilder[any](service.client, path).Delete(ctx)
}
