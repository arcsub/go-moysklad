package moysklad

import (
	"github.com/google/uuid"
)

// ProcessingProcess Техпроцесс.
// Ключевое слово: processingprocess
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tehprocess
type ProcessingProcess struct {
	AccountID    *uuid.UUID                            `json:"accountId,omitempty"`    // ID учетной записи
	Archived     *bool                                 `json:"archived,omitempty"`     // Добавлен ли Тех. процесс в архив
	Description  *string                               `json:"description,omitempty"`  // Комментарий Тех. процесса
	ExternalCode *string                               `json:"externalCode,omitempty"` // Внешний код Тех. процесса
	Group        *Group                                `json:"group,omitempty"`        // Отдел сотрудника
	ID           *uuid.UUID                            `json:"id,omitempty"`           // ID Тех. процесса
	Meta         *Meta                                 `json:"meta,omitempty"`         // Метаданные Тех. процесса
	Name         *string                               `json:"name,omitempty"`         // Наименование Тех. процесса
	Owner        *Employee                             `json:"owner,omitempty"`        // Владелец (Сотрудник)
	Positions    *Positions[ProcessingProcessPosition] `json:"positions,omitempty"`    // Метаданные позиций Тех. процесса
	Shared       *bool                                 `json:"shared,omitempty"`       // Общий доступ
	Updated      *Timestamp                            `json:"updated,omitempty"`      // Момент последнего обновления сущности
}

func (p ProcessingProcess) String() string {
	return Stringify(p)
}

func (p ProcessingProcess) MetaType() MetaType {
	return MetaTypeProcessingProcess
}

// ProcessingProcessPosition Позиция Тех. процесса.
// Ключевое слово: processingprocessposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-teh-process-teh-processy-atributy-wlozhennyh-suschnostej-pozicii-teh-processa
type ProcessingProcessPosition struct {
	AccountID       *uuid.UUID       `json:"accountId,omitempty"`       // ID учетной записи
	ID              *uuid.UUID       `json:"id,omitempty"`              // ID позиции
	Meta            *Meta            `json:"meta,omitempty"`            // Метаданные позиции Тех. процесса
	ProcessingStage *ProcessingStage `json:"processingstage,omitempty"` // Метаданные этапа, который представляет собой позиция
}

func (p ProcessingProcessPosition) String() string {
	return Stringify(p)
}

func (p ProcessingProcessPosition) MetaType() MetaType {
	return MetaTypeProcessingProcessPosition
}
