package moysklad

import (
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
