package moysklad

import "github.com/google/uuid"

// ProductionStageCompletion Выполнение этапа производства
// Ключевое слово: productionstagecompletion
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vypolnenie-atapa-proizwodstwa
type ProductionStageCompletion struct {
	AccountId          *uuid.UUID                                    `json:"accountId,omitempty"`          // ID учетной записи
	Created            *Timestamp                                    `json:"created,omitempty"`            // Дата создания
	ExternalCode       *string                                       `json:"externalCode,omitempty"`       // Внешний код Выполнения этапа производства
	Group              *Group                                        `json:"group,omitempty"`              // Отдел сотрудника
	ID                 *uuid.UUID                                    `json:"id,omitempty"`                 // ID Выполнения этапа производства
	LabourUnitCost     *float64                                      `json:"labourUnitCost,omitempty"`     // Оплата труда за единицу объема производства
	Materials          *Positions[ProductionStageCompletionMaterial] `json:"materials,omitempty"`          // Метаданные Материалов выполнения этапа производства
	Meta               *Meta                                         `json:"meta,omitempty"`               // Метаданные Выполнения этапа производства
	Moment             *Timestamp                                    `json:"moment,omitempty"`             // Дата документа
	Name               *string                                       `json:"name,omitempty"`               // Наименование Выполнения этапа производства
	Owner              *Employee                                     `json:"owner,omitempty"`              // Владелец (Сотрудник)
	Performer          *Employee                                     `json:"performer,omitempty"`          // Исполнитель (Сотрудник)
	ProcessingUnitCost *float64                                      `json:"processingUnitCost,omitempty"` // Затраты на единицу объема производства
	ProductionStage    *ProductionStage                              `json:"productionStage,omitempty"`    // Производственный этап
	ProductionVolume   *float64                                      `json:"productionVolume,omitempty"`   // Объем производства
	Products           *Positions[ProductionStageCompletionResult]   `json:"products,omitempty"`           // Метаданные Продуктов выполнения этапа производства. Есть только у последнего этапа
	Shared             *bool                                         `json:"shared,omitempty"`             // Общий доступ
	Updated            *Timestamp                                    `json:"updated,omitempty"`            // Момент последнего обновления Выполнения этапа производства
}

// ProductionStageCompletionMaterial Материалы Выполнения этапа производства
// Ключевое слово: productionstagecompletionmaterial
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vypolnenie-atapa-proizwodstwa-izmenit-vypolnenie-atapa-proizwodstwa-materialy-vypolneniq-atapa-proizwodstwa
type ProductionStageCompletionMaterial struct {
	AccountId        *uuid.UUID          `json:"accountId,omitempty"`        // ID учетной записи
	Assortment       *AssortmentPosition `json:"assortment,omitempty"`       // Метаданные товара/модификации/серии, которую представляет собой позиция
	ConsumedQuantity *float64            `json:"consumedQuantity,omitempty"` // Количество товаров/модификаций данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе
	ID               *uuid.UUID          `json:"id,omitempty"`               // ID позиции
	Things           *Things             `json:"things,omitempty"`           // Серийные номера. Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете. В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута
}

// ProductionStageCompletionResult Продукт Выполнения этапа производства
// Ключевое слово: productionstagecompletionresult
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vypolnenie-atapa-proizwodstwa-izmenit-vypolnenie-atapa-proizwodstwa-produkty-vypolneniq-atapa-proizwodstwa
type ProductionStageCompletionResult struct {
	AccountId        *uuid.UUID          `json:"accountId,omitempty"`        // ID учетной записи
	Assortment       *AssortmentPosition `json:"assortment,omitempty"`       // Метаданные товара/модификации/серии, которую представляет собой позиция
	ID               *uuid.UUID          `json:"id,omitempty"`               // ID позиции
	ProducedQuantity *float64            `json:"producedQuantity,omitempty"` // Количество товаров/модификаций данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе
	Things           *Things             `json:"things,omitempty"`           // Серийные номера. Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете. В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута
}
