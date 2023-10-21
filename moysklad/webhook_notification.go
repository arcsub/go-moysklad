package moysklad

// WebhookNotification уведомление вебхука.
type WebhookNotification struct {
	AuditContext AuditContext `json:"auditContext,omitempty"`
	Events       Slice[Event] `json:"events,omitempty"`
}

// AuditContext Контекст аудита, соответствующий событию вебхука
type AuditContext struct {
	Meta   Meta      `json:"meta,omitempty"`   // Метаданные контекста аудита
	Moment Timestamp `json:"moment,omitempty"` // Дата изменения
	Uid    string    `json:"uid,omitempty"`    // Логин Сотрудника
}

// Event Данные о событии, вызвавшем срабатывание вебхука
type Event struct {
	AccountID     string        `json:"accountId,omitempty"`     // ID учетной записи
	Action        WebhookAction `json:"action,omitempty"`        // Действие, которое вызвало срабатывание вебхука.
	Meta          Meta          `json:"meta,omitempty"`          // Метаданные измененной сущности
	UpdatedFields []string      `json:"updatedFields,omitempty"` // Поля сущности, измененные пользователем
}
