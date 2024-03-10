package moysklad

import "github.com/google/uuid"

// ProductionTask Производственное задание
// Ключевое слово: productiontask
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie
type ProductionTask struct {
	AccountId             *uuid.UUID                       `json:"accountId,omitempty"`             // ID учетной записи
	Applicable            *bool                            `json:"applicable,omitempty"`            // Отметка о проведении
	Attributes            *Attributes                      `json:"attributes,omitempty"`            // Коллекция метаданных доп. полей Поля объекта
	Code                  *string                          `json:"code,omitempty"`                  // Код Производственного задания
	Created               *Timestamp                       `json:"created,omitempty"`               // Дата создания
	Deleted               *Timestamp                       `json:"deleted,omitempty"`               // Момент последнего удаления Производственного задания
	DeliveryPlannedMoment *Timestamp                       `json:"deliveryPlannedMoment,omitempty"` // Планируемая дата выполнения
	Description           *string                          `json:"description,omitempty"`           // Комментарий Производственного задания
	ExternalCode          *string                          `json:"externalCode,omitempty"`          // Внешний код Производственного задания
	Files                 *Files                           `json:"files,omitempty"`                 // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group                 *Group                           `json:"group,omitempty"`                 // Отдел сотрудника
	ID                    *uuid.UUID                       `json:"id,omitempty"`                    // ID Производственного задания
	MaterialsStore        *Store                           `json:"materialsStore,omitempty"`        // Метаданные склада материалов
	Meta                  *Meta                            `json:"meta,omitempty"`                  // Метаданные Производственного задания
	Moment                *Timestamp                       `json:"moment,omitempty"`                // Дата документа
	Name                  *string                          `json:"name,omitempty"`                  // Наименование Производственного задания
	Organization          *Organization                    `json:"organization,omitempty"`          // Метаданные юрлица
	Owner                 *Employee                        `json:"owner,omitempty"`                 // Владелец (Сотрудник)
	Printed               *bool                            `json:"printed,omitempty"`               // Напечатан ли документ
	ProductionRows        *Positions[ProductionRow]        `json:"productionRows,omitempty"`        // Метаданные Позиций производственного задания
	ProductionEnd         *Timestamp                       `json:"productionEnd,omitempty"`         // Дата окончания производства
	ProductionStart       *Timestamp                       `json:"productionStart,omitempty"`       // Дата начала производства
	Products              *Positions[ProductionTaskResult] `json:"products,omitempty"`              // Метаданные производимой продукции
	ProductsStore         *Store                           `json:"productsStore,omitempty"`         // Метаданные склада продукции
	Published             *bool                            `json:"published,omitempty"`             // Опубликован ли документ
	Reserve               *bool                            `json:"reserve,omitempty"`               // Флаг резервирования материала Производственного задания
	Shared                *bool                            `json:"shared,omitempty"`                // Общий доступ
	State                 *State                           `json:"state,omitempty"`                 // Метаданные статуса Производственного задания
	Updated               *Timestamp                       `json:"updated,omitempty"`               // Момент последнего обновления Производственного задания
}

// ProductionRow Позиция производственного задания
// Ключевое слово: productionrow
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-proizwodstwennye-zadaniq-pozicii-proizwodstwennogo-zadaniq
type ProductionRow struct {
	AccountId        *uuid.UUID      `json:"accountId,omitempty"`        // ID учетной записи
	ExternalCode     *string         `json:"externalCode,omitempty"`     // Внешний код
	ID               *uuid.UUID      `json:"id,omitempty"`               // ID позиции
	Name             *string         `json:"name,omitempty"`             // Наименование
	ProcessingPlan   *ProcessingPlan `json:"processingPlan,omitempty"`   // Метаданные Техкарты
	ProductionVolume *float64        `json:"productionVolume,omitempty"` // Объем производства.
	Updated          *Timestamp      `json:"updated,omitempty"`          // Момент последнего обновления Производственного задания
}

// ProductionTaskResult Продукт производственного задания
// Ключевое слово: productiontaskresult
// https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-proizwodstwennoe-zadanie-proizwodstwennye-zadaniq-produkty-proizwodstwennogo-zadaniq
type ProductionTaskResult struct {
	AccountId     *uuid.UUID          `json:"accountId,omitempty"`     // ID учетной записи
	Assortment    *AssortmentPosition `json:"assortment,omitempty"`    // Ссылка на товар/серию/модификацию, которую представляет собой позиция.
	ID            *uuid.UUID          `json:"id,omitempty"`            // ID позиции
	PlanQuantity  *float64            `json:"planQuantity,omitempty"`  // Запланированное для производства количество продукта
	ProductionRow *ProductionRow      `json:"productionRow,omitempty"` // Метаданные Позиции производственного задания
}
