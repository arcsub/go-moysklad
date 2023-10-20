package moysklad

import (
	"github.com/google/uuid"
)

// Processing Техоперация.
// Ключевое слово: processing
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-tehoperaciq
type Processing struct {
	AccountID           *uuid.UUID                        `json:"accountId,omitempty"`           // ID учетной записи
	Applicable          *bool                             `json:"applicable,omitempty"`          // Отметка о проведении
	Attributes          *Attributes                       `json:"attributes,omitempty"`          // Коллекция метаданных доп. полей. Поля объекта
	Code                *string                           `json:"code,omitempty"`                // Код
	Created             *Timestamp                        `json:"created,omitempty"`             // Дата создания
	Deleted             *Timestamp                        `json:"deleted,omitempty"`             // Момент последнего удаления
	Description         *string                           `json:"description,omitempty"`         // Комментарий
	ExternalCode        *string                           `json:"externalCode,omitempty"`        // Внешний код
	Files               *Files                            `json:"files,omitempty"`               // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group               *Group                            `json:"group,omitempty"`               // Отдел сотрудника
	ID                  *uuid.UUID                        `json:"id,omitempty"`                  // ID сущности
	Materials           Slice[ProcessingPositionMaterial] `json:"materials,omitempty"`           // Список Метаданных материалов Тех. операции
	MaterialsStore      *Store                            `json:"materialsStore,omitempty"`      // Метаданные склада для материалов
	Meta                *Meta                             `json:"meta,omitempty"`                // Метаданные
	Moment              *Timestamp                        `json:"moment,omitempty"`              // Дата документа
	Name                *string                           `json:"name,omitempty"`                // Наименование
	Organization        *Organization                     `json:"organization,omitempty"`        // Метаданные юрлица
	OrganizationAccount *AgentAccount                     `json:"organizationAccount,omitempty"` // Метаданные счета юрлица
	Owner               *Employee                         `json:"owner,omitempty"`               // Владелец (Сотрудник)
	Printed             *bool                             `json:"printed,omitempty"`             // Напечатан ли документ
	ProcessingPlan      *ProcessingPlan                   `json:"processingPlan,omitempty"`      // Метаданные Тех. операции
	ProcessingSum       *float64                          `json:"processingSum,omitempty"`       // Затраты на производство
	Products            Slice[ProcessingPositionProduct]  `json:"products,omitempty"`            // Список Метаданных готовых продуктов Тех. операции
	ProductsStore       *Store                            `json:"productsStore,omitempty"`       // Метаданные склада для продукции
	Project             *Project                          `json:"project,omitempty"`             // Метаданные проекта
	Published           *bool                             `json:"published,omitempty"`           // Опубликован ли документ
	Quantity            *float64                          `json:"quantity,omitempty"`            // Объем производства
	Shared              *bool                             `json:"shared,omitempty"`              // Общий доступ
	State               *State                            `json:"state,omitempty"`               // Метаданные статуса
	SyncID              *uuid.UUID                        `json:"syncId,omitempty"`              // ID синхронизации. После заполнения недоступен для изменения
	Updated             *Timestamp                        `json:"updated,omitempty"`             // Момент последнего обновления
	ProcessingOrder     *ProcessingOrder                  `json:"processingOrder,omitempty"`     // Ссылка на заказ на производство в формате Метаданных
}

func (p Processing) String() string {
	return Stringify(p)
}

func (p Processing) MetaType() MetaType {
	return MetaTypeProcessing
}

type Processings = Iterator[Processing]

// ProcessingPositionMaterial Материал Техоперации.
// Ключевое слово: processingpositionmaterial
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-tehoperaciq-tehoperacii-materialy-tehoperacii
type ProcessingPositionMaterial struct {
	ProcessingPosition
}

func (p ProcessingPositionMaterial) String() string {
	return Stringify(p)
}

func (p ProcessingPositionMaterial) MetaType() MetaType {
	return MetaTypeProcessingPositionMaterial
}

// ProcessingPositionProduct Продукт Техоперации.
// Ключевое слово: processingpositionresult
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-tehoperaciq-tehoperacii-produkty-tehoperacii
type ProcessingPositionProduct struct {
	ProcessingPosition
}

func (p ProcessingPositionProduct) String() string {
	return Stringify(p)
}

func (p ProcessingPositionProduct) MetaType() MetaType {
	return MetaTypeProcessingPositionProduct
}

// ProcessingTemplateArg
// Документ: Техоперация (processing)
// Основание, на котором он может быть создан:
// - Заказ на производство (processingorder)
// - Техкарта (processingplan)
type ProcessingTemplateArg struct {
	ProcessingOrder *MetaWrapper `json:"processingOrder,omitempty"`
	ProcessingPlan  *MetaWrapper `json:"processingPlan,omitempty"`
}
