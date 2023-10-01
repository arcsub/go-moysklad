package moysklad

import (
	"github.com/google/uuid"
)

// Task Задача.
// Ключевое слово: task
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-zadacha
type Task struct {
	AccountId         *uuid.UUID    `json:"accountId,omitempty"`         // ID учетной записи
	Agent             *Counterparty `json:"agent,omitempty"`             // Метаданные Контрагента или юрлица, связанного с задачей. Задача может быть привязана либо к конрагенту, либо к юрлицу, либо к документу
	Assignee          *Employee     `json:"assignee,omitempty"`          // Метаданные ответственного за выполнение задачи
	Author            *Employee     `json:"author,omitempty"`            // Метаданные Сотрудника, создавшего задачу (администратор аккаунта, если автор - Приложение)
	AuthorApplication *Application  `json:"authorApplication,omitempty"` // Метаданные Приложения, создавшего задачу
	Completed         *Timestamp    `json:"completed,omitempty"`         // Время выполнения задачи
	Created           *Timestamp    `json:"created,omitempty"`           // Момент создания
	Description       *string       `json:"description,omitempty"`       // Текст задачи
	Done              *bool         `json:"done,omitempty"`              // Отметка о выполнении задачи
	DueToDate         *Timestamp    `json:"dueToDate,omitempty"`         // Срок задачи
	Files             *Files        `json:"files,omitempty"`             // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Id                *uuid.UUID    `json:"id,omitempty"`                // ID Задачи
	Implementer       *Employee     `json:"implementer,omitempty"`       // Метаданные Сотрудника, выполнившего задачу
	Meta              *Meta         `json:"meta,omitempty"`              // Метаданные
	Notes             *TaskNotes    `json:"notes,omitempty"`             // Метаданные комментария к задаче
	//Operation         *Operations   `json:"operation,omitempty"`         // Метаданные Документа, связанного с задачей. Задача может быть привязана либо к конрагенту, либо к юрлицу, либо к документу
	Updated *Timestamp `json:"updated,omitempty"` // Момент последнего обновления Задачи
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

type TaskNotes = Iterator[TaskNote]
