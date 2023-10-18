package moysklad

import "github.com/google/uuid"

// Notification TODO: Общие атрибуты уведомлений.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/notification/#uwedomleniq-lenta-uwedomlenij-obschie-atributy-uwedomlenij
type Notification struct {
	AccountId   uuid.UUID `json:"accountId"`   // ID учетной записи
	Created     Timestamp `json:"created"`     // Дата и время формирования Уведомления
	Description string    `json:"description"` // Описание уведомления
	Id          uuid.UUID `json:"id"`          // ID Уведомления
	Meta        Meta      `json:"meta"`        // Метаданные объекта. Содержит тип конкретного уведомления
	Read        bool      `json:"read"`        // Признак того, было ли Уведомление прочитано
	Title       string    `json:"title"`       // Краткий текст уведомления
}

func (n Notification) String() string {
	return Stringify(n)
}

func (n Notification) MetaType() MetaType {
	return MetaTypeNotification
}
