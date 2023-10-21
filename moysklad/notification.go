package moysklad

import "github.com/google/uuid"

// Notification TODO: Общие атрибуты уведомлений.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-lenta-uwedomlenij-obschie-atributy-uwedomlenij
type Notification struct {
	Meta        Meta      `json:"meta"`
	Created     Timestamp `json:"created"`
	Description string    `json:"description"`
	Title       string    `json:"title"`
	AccountID   uuid.UUID `json:"accountId"`
	ID          uuid.UUID `json:"id"`
	Read        bool      `json:"read"`
}

func (n Notification) String() string {
	return Stringify(n)
}

func (n Notification) MetaType() MetaType {
	return MetaTypeNotification
}
