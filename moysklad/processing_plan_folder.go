package moysklad

import (
	"github.com/google/uuid"
)

// ProcessingPlanFolder Группа тех. карт.
// Ключевое слово: processingplanfolder
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-gruppa-teh-kart
type ProcessingPlanFolder struct {
	AccountId    *uuid.UUID `json:"accountId,omitempty"`    // ID учетной записи
	Archived     *bool      `json:"archived,omitempty"`     // Добавлена ли Группа тех. карт в архив
	ExternalCode *string    `json:"externalCode,omitempty"` // Внешний код Группы тех. карт
	Code         *string    `json:"code,omitempty"`         // Код Группы тех. карт
	Description  *string    `json:"description,omitempty"`  // Описание Группы тех. карт
	Group        *Group     `json:"group,omitempty"`        // Отдел сотрудника
	Id           *uuid.UUID `json:"id,omitempty"`           // ID Группы тех. карт
	Meta         *Meta      `json:"meta,omitempty"`         // Метаданные
	Name         *string    `json:"name,omitempty"`         // Наименование
	Owner        *Employee  `json:"owner,omitempty"`        // Владелец (Сотрудник)
	PathName     *string    `json:"pathName,omitempty"`     // Наименование Группы тех. карт, в которую входит данная Группа тех. карт
	Shared       *bool      `json:"shared,omitempty"`       // Общий доступ
	Updated      *Timestamp `json:"updated,omitempty"`      // Момент последнего обновления сущности
}

func (p ProcessingPlanFolder) String() string {
	return Stringify(p)
}

func (p ProcessingPlanFolder) MetaType() MetaType {
	return MetaTypeProcessingPlanFolder
}
