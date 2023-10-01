package moysklad

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

// TaskService
// Сервис для работы с задачами.
type TaskService struct {
	Endpoint
	endpointGetList[Task]
	endpointCreate[Task]
	endpointCreateUpdateDeleteMany[Task]
	endpointDelete
	endpointGetById[Task]
	endpointUpdate[Task]
	endpointNamedFilter
}

func NewTaskService(client *Client) *TaskService {
	e := NewEndpoint(client, "entity/task")
	return &TaskService{
		Endpoint:                       e,
		endpointGetList:                endpointGetList[Task]{e},
		endpointCreate:                 endpointCreate[Task]{e},
		endpointCreateUpdateDeleteMany: endpointCreateUpdateDeleteMany[Task]{e},
		endpointDelete:                 endpointDelete{e},
		endpointGetById:                endpointGetById[Task]{e},
		endpointUpdate:                 endpointUpdate[Task]{e},
		endpointNamedFilter:            endpointNamedFilter{e},
	}
}

// GetNotes Запрос на получение списка всех комментариев данной Задачи.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-zadacha-poluchit-kommentarii-zadachi
func (s *TaskService) GetNotes(ctx context.Context, taskId uuid.UUID, params *Params) (*List[TaskNote], *Response, error) {
	path := fmt.Sprintf("%s/notes", taskId)
	return NewRequestBuilder[List[TaskNote]](s.Endpoint, ctx).WithPath(path).WithParams(params).Get()
}

// CreateNote Запрос на создание нового комментария к Задаче.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-zadacha-sozdat-kommentarij-zadachi
func (s *TaskService) CreateNote(ctx context.Context, taskId uuid.UUID, taskNote *TaskNote) (*TaskNote, *Response, error) {
	path := fmt.Sprintf("%s/notes", taskId)
	return NewRequestBuilder[TaskNote](s.Endpoint, ctx).WithPath(path).Post()
}

// CreateNotes Запрос на создание нескольких комментариев к Задаче.
func (s *TaskService) CreateNotes(ctx context.Context, taskId uuid.UUID, taskNotes []*TaskNote) (*Slice[TaskNote], *Response, error) {
	path := fmt.Sprintf("%s/notes", taskId)
	return NewRequestBuilder[Slice[TaskNote]](s.Endpoint, ctx).WithPath(path).WithBody(taskNotes).Post()
}

// GetNoteById Отдельный комментарий к Задаче с указанным id комментария.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-zadacha-poluchit-kommentarij-k-zadache
func (s *TaskService) GetNoteById(ctx context.Context, taskId, taskNoteId uuid.UUID) (*TaskNote, *Response, error) {
	path := fmt.Sprintf("%s/notes/%s", taskId, taskNoteId)
	return NewRequestBuilder[TaskNote](s.Endpoint, ctx).WithPath(path).Get()
}

// UpdateNote Запрос на обновление отдельного комментария к Задаче.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-zadacha-izmenit-kommentarij-k-zadache
func (s *TaskService) UpdateNote(ctx context.Context, taskId, taskNoteId uuid.UUID, taskNote *TaskNote) (*TaskNote, *Response, error) {
	path := fmt.Sprintf("%s/notes/%s", taskId, taskNoteId)
	return NewRequestBuilder[TaskNote](s.Endpoint, ctx).WithPath(path).WithBody(taskNote).Put()
}

// DeleteNote Запрос на удаление отдельного комментария к Задаче с указанным id.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-zadacha-udalit-kommentarij
func (s *TaskService) DeleteNote(ctx context.Context, taskId, taskNoteId uuid.UUID) (bool, *Response, error) {
	path := fmt.Sprintf("%s/notes/%s", taskId, taskNoteId)
	return NewRequestBuilder[any](s.Endpoint, ctx).WithPath(path).Delete()
}
