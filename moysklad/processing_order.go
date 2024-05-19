package moysklad

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ProcessingOrder Заказ на производство.
// Ключевое слово: processingorder
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-zakaz-na-proizwodstwo
type ProcessingOrder struct {
	AccountID             *uuid.UUID                          `json:"accountId,omitempty"`             // ID учетной записи
	Applicable            *bool                               `json:"applicable,omitempty"`            // Отметка о проведении
	Attributes            *Attributes                         `json:"attributes,omitempty"`            // Коллекция метаданных доп. полей. Поля объекта
	Code                  *string                             `json:"code,omitempty"`                  // Код
	Created               *Timestamp                          `json:"created,omitempty"`               // Дата создания
	Deleted               *Timestamp                          `json:"deleted,omitempty"`               // Момент последнего удаления
	DeliveryPlannedMoment *Timestamp                          `json:"deliveryPlannedMoment,omitempty"` // Планируемая дата производства
	Description           *string                             `json:"description,omitempty"`           // Комментарий
	ExternalCode          *string                             `json:"externalCode,omitempty"`          // Внешний код
	Files                 *Files                              `json:"files,omitempty"`                 // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group                 *Group                              `json:"group,omitempty"`                 // Отдел сотрудника
	ID                    *uuid.UUID                          `json:"id,omitempty"`                    // ID сущности
	Meta                  *Meta                               `json:"meta,omitempty"`                  // Метаданные
	Moment                *Timestamp                          `json:"moment,omitempty"`                // Дата документа
	Name                  *string                             `json:"name,omitempty"`                  // Наименование
	Organization          *Organization                       `json:"organization,omitempty"`          // Метаданные юрлица
	OrganizationAccount   *AgentAccount                       `json:"organizationAccount,omitempty"`   // Метаданные счета юрлица
	Owner                 *Employee                           `json:"owner,omitempty"`                 // Владелец (Сотрудник)
	Positions             *Positions[ProcessingOrderPosition] `json:"positions,omitempty"`             // Метаданные позиций Заказа на производство
	Printed               *bool                               `json:"printed,omitempty"`               // Напечатан ли документ
	ProcessingPlan        *ProcessingPlan                     `json:"processingPlan,omitempty"`        // Метаданные Тех. плана
	Project               *Project                            `json:"project,omitempty"`               // Метаданные проекта
	Published             *bool                               `json:"published,omitempty"`             // Опубликован ли документ
	Quantity              *float64                            `json:"quantity,omitempty"`              // Объем производства
	Shared                *bool                               `json:"shared,omitempty"`                // Общий доступ
	State                 *State                              `json:"state,omitempty"`                 // Метаданные статуса
	Store                 *Store                              `json:"store,omitempty"`                 // Метаданные склада
	SyncID                *uuid.UUID                          `json:"syncId,omitempty"`                // ID синхронизации. После заполнения недоступен для изменения
	Updated               *Timestamp                          `json:"updated,omitempty"`               // Момент последнего обновления
	Processings           *Processings                        `json:"processings,omitempty"`           // Массив ссылок на связанные тех. операции в формате Метаданных
}

func (p ProcessingOrder) String() string {
	return Stringify(p)
}

func (p ProcessingOrder) MetaType() MetaType {
	return MetaTypeProcessingOrder
}

// ProcessingOrderPosition Позиция Заказа на производство.
// Ключевое слово: processingorderposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-zakaz-na-proizwodstwo-zakazy-na-proizwodstwo-pozicii-zakaza-na-proizwodstwo
type ProcessingOrderPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Reserve    *decimal.Decimal    `json:"reserve,omitempty"`    // Резерв данной позиции
}

func (p ProcessingOrderPosition) String() string {
	return Stringify(p)
}

func (p ProcessingOrderPosition) MetaType() MetaType {
	return MetaTypeProcessingOrderPosition
}

// ProcessingOrderTemplateArg
// Документ: Заказ на производство (processingorder)
// Основание, на котором он может быть создан:
// - Техкарта (processingplan)
type ProcessingOrderTemplateArg struct {
	ProcessingPlan *MetaWrapper `json:"processingPlan,omitempty"`
}
