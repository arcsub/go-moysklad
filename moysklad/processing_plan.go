package moysklad

import (
	"github.com/google/uuid"
)

// ProcessingPlan Техкарта.
// Ключевое слово: processingplan
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-tehkarta-tehkarty
type ProcessingPlan struct {
	AccountId         *uuid.UUID                         `json:"accountId,omitempty"`         // ID учетной записи
	Archived          *bool                              `json:"archived,omitempty"`          // Добавлена ли Тех. карта в архив
	Code              *string                            `json:"code,omitempty"`              // Код Тех. карты
	Cost              *float64                           `json:"cost,omitempty"`              // Стоимость производства
	ExternalCode      *string                            `json:"externalCode,omitempty"`      // Внешний код
	Group             *Group                             `json:"group,omitempty"`             // Отдел сотрудника
	Id                *uuid.UUID                         `json:"id,omitempty"`                // ID сущности
	Stages            *MetaArray[ProcessingStage]        `json:"stages,omitempty"`            // Коллекция метаданных этапов Тех. карты
	Materials         *Positions[ProcessingPlanMaterial] `json:"materials,omitempty"`         // Список Метаданных материалов Тех. операции
	Meta              *Meta                              `json:"meta,omitempty"`              // Метаданные
	Name              *string                            `json:"name,omitempty"`              // Наименование
	Owner             *Employee                          `json:"owner,omitempty"`             // Владелец (Сотрудник)
	Parent            *Group                             `json:"parent,omitempty"`            // Метаданные группы Тех. карты
	PathName          *string                            `json:"pathName,omitempty"`          // Наименование группы, в которую входит Тех. карта
	ProcessingProcess *ProcessingProcess                 `json:"processingProcess,omitempty"` // Метаданные Тех. процесса
	Products          *Positions[ProcessingPlanProduct]  `json:"products,omitempty"`          // Коллекция метаданных готовых продуктов Тех. карты
	Shared            *bool                              `json:"shared,omitempty"`            // Общий доступ
	Updated           *Timestamp                         `json:"updated,omitempty"`           // Момент последнего обновления
}

func (p ProcessingPlan) String() string {
	return Stringify(p)
}

func (p ProcessingPlan) MetaType() MetaType {
	return MetaTypeProcessingPlan
}
