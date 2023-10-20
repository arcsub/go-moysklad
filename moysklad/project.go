package moysklad

import (
	"github.com/google/uuid"
)

// Project Проект.
// Ключевое слово: project
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-proekt
type Project struct {
	AccountID    *uuid.UUID  `json:"accountId,omitempty"`    // ID учетной записи
	Archived     *bool       `json:"archived,omitempty"`     // Добавлен ли Проект в архив
	Attributes   *Attributes `json:"attributes,omitempty"`   // Коллекция доп. полей
	Code         *string     `json:"code,omitempty"`         // Код Проекта
	Description  *string     `json:"description,omitempty"`  // Описание Проекта
	ExternalCode *string     `json:"externalCode,omitempty"` // Внешний код Проекта
	Group        *Group      `json:"group,omitempty"`        // Метаданные отдела сотрудника
	ID           *uuid.UUID  `json:"id,omitempty"`           // ID сущности
	Meta         *Meta       `json:"meta,omitempty"`         // Метаданные
	Name         *string     `json:"name,omitempty"`         // Наименование Проекта
	Owner        *Employee   `json:"owner,omitempty"`        // Метаданные владельца (Сотрудника)
	Shared       *bool       `json:"shared,omitempty"`       // Общий доступ
	Updated      *Timestamp  `json:"updated,omitempty"`      // Момент последнего обновления сущности
}

func (p Project) String() string {
	return Stringify(p)
}

func (p Project) MetaType() MetaType {
	return MetaTypeProject
}
