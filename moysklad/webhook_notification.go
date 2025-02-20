package moysklad

// WebhookNotification уведомление вебхука.
type WebhookNotification struct {
	AuditContext AuditContext `json:"auditContext"`
	Events       Slice[Event] `json:"events"`
}

// AuditContext Контекст аудита, соответствующий событию вебхука.
type AuditContext struct {
	Meta Meta `json:"meta"` // Метаданные контекста аудита
	//Moment Timestamp `json:"moment"` // Дата создания
	UID string `json:"uid"` // Логин Сотрудника
}

// Event Данные о событии, вызвавшем срабатывание вебхука.
type Event struct {
	AccountID     string        `json:"accountId"`     // ID учётной записи
	Action        WebhookAction `json:"action"`        // Действие, которое вызвало срабатывание вебхука
	Meta          Meta          `json:"meta"`          // Метаданные измененной сущности
	UpdatedFields Slice[string] `json:"updatedFields"` // Поля сущности, измененные пользователем
}
