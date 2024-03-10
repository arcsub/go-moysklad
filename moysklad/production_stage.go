package moysklad

import "github.com/google/uuid"

// ProductionStage Производственный этап
// Ключевое слово: productionstage
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-proizwodstwennye-atapy
type ProductionStage struct {
	AccountId          *uuid.UUID                         `json:"accountId,omitempty"`          // ID учетной записи
	ID                 *uuid.UUID                         `json:"id,omitempty"`                 // ID Производственного этапа
	Meta               *Meta                              `json:"meta,omitempty"`               // Метаданные Производственного этапа
	LabourUnitCost     *float64                           `json:"labourUnitCost,omitempty"`     // Затраты на оплату труда за единицу объема производства
	Materials          *Positions[ProductionTaskMaterial] `json:"materials,omitempty"`          // Метаданные Материалов производственного этапа
	OrderingPosition   *int                               `json:"orderingPosition,omitempty"`   // Индекс Производственного этапа в Позиции производственного задания
	Stage              *ProductionStage                   `json:"stage,omitempty"`              // Метаданные Этапа производства
	ProductionRow      *ProductionRow                     `json:"productionRow,omitempty"`      // Метаданные Позиции производственного задания
	TotalQuantity      *float64                           `json:"totalQuantity,omitempty"`      // Объем Производственного этапа. Соответствует объему Позиции производственного задания
	CompletedQuantity  *float64                           `json:"completedQuantity,omitempty"`  // Выполненное количество
	AvailableQuantity  *float64                           `json:"availableQuantity,omitempty"`  // Количество, доступное к выполнению
	BlockedQuantity    *float64                           `json:"blockedQuantity,omitempty"`    // Количество, которое на данный момент выполнять нельзя. Например, ещё не выполнен предыдущий этап
	SkippedQuantity    *float64                           `json:"skippedQuantity,omitempty"`    // Количество, которое не будет выполнено. Например, из-за остановки производства
	ProcessingUnitCost *float64                           `json:"processingUnitCost,omitempty"` // Затраты на единицу объема производства
}

// ProductionTaskMaterial Материал Производственного этапа
// Ключевое слово: productiontaskmaterial
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-proizwodstwennye-atapy-materialy-proizwodstwennogo-atapa
type ProductionTaskMaterial struct {
	AccountId    *uuid.UUID          `json:"accountId,omitempty"`    // ID учетной записи
	Assortment   *AssortmentPosition `json:"assortment,omitempty"`   // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	ID           *uuid.UUID          `json:"id,omitempty"`           // ID позиции
	PlanQuantity *float64            `json:"planQuantity,omitempty"` // Количество товаров/модификаций данного вида в позиции
}
