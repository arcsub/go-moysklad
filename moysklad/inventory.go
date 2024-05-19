package moysklad

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Inventory Инвентаризация.
// Ключевое слово: inventory
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-inwentarizaciq
type Inventory struct {
	AccountID    *uuid.UUID                    `json:"accountId,omitempty"`    // ID учетной записи
	Attributes   *Attributes                   `json:"attributes,omitempty"`   // Коллекция метаданных доп. полей. Поля объекта
	Code         *string                       `json:"code,omitempty"`         // Код выданного
	Created      *Timestamp                    `json:"created,omitempty"`      // Дата создания
	Deleted      *Timestamp                    `json:"deleted,omitempty"`      // Момент последнего удаления Счета-фактуры полученного
	Description  *string                       `json:"description,omitempty"`  // Комментарий выданного Счета-фактуры полученного
	ExternalCode *string                       `json:"externalCode,omitempty"` // Внешний код выданного Счета-фактуры полученного
	Files        *Files                        `json:"files,omitempty"`        // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group        *Group                        `json:"group,omitempty"`        // Отдел сотрудника
	ID           *uuid.UUID                    `json:"id,omitempty"`           // ID сущности
	Meta         *Meta                         `json:"meta,omitempty"`         // Метаданные
	Moment       *Timestamp                    `json:"moment,omitempty"`       // Дата документа
	Name         *string                       `json:"name,omitempty"`         // Наименование
	Organization *Organization                 `json:"organization,omitempty"` // Метаданные юрлица
	Owner        *Employee                     `json:"owner,omitempty"`        // Владелец (Сотрудник)
	Positions    *Positions[InventoryPosition] `json:"positions,omitempty"`    // Метаданные позиций Инвентаризации
	Printed      *bool                         `json:"printed,omitempty"`      // Напечатан ли документ
	Published    *bool                         `json:"published,omitempty"`    // Опубликован ли документ
	Shared       *bool                         `json:"shared,omitempty"`       // Общий доступ
	State        *State                        `json:"state,omitempty"`        // Метаданные статуса
	Store        *Store                        `json:"store,omitempty"`        // Метаданные склада
	Sum          *decimal.Decimal              `json:"sum,omitempty"`          // Сумма
	SyncID       *uuid.UUID                    `json:"syncId,omitempty"`       // ID синхронизации. После заполнения недоступен для изменения
	Updated      *Timestamp                    `json:"updated,omitempty"`      // Момент последнего обновления
}

func (i Inventory) String() string {
	return Stringify(i)
}

func (i Inventory) MetaType() MetaType {
	return MetaTypeInventory
}

// InventoryPosition Позиция Инвентаризации.
// Ключевое слово: inventoryposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-inwentarizaciq-inwentarizaciq-pozicii-inwentarizacii
type InventoryPosition struct {
	AccountID          *uuid.UUID          `json:"accountId,omitempty"`          // ID учетной записи
	Assortment         *AssortmentPosition `json:"assortment,omitempty"`         // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	CalculatedQuantity *float64            `json:"calculatedQuantity,omitempty"` // расчетный остаток
	CorrectionAmount   *float64            `json:"correctionAmount,omitempty"`   // разница между расчетным остатком и фактическимх
	CorrectionSum      *decimal.Decimal    `json:"correctionSum,omitempty"`      // избыток/недостача
	ID                 *uuid.UUID          `json:"id,omitempty"`                 // ID сущности
	Pack               *Pack               `json:"pack,omitempty"`               // Упаковка Товара
	Price              *decimal.Decimal    `json:"price,omitempty"`              // Цена товара/услуги в копейках
	Quantity           *float64            `json:"quantity,omitempty"`           // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
}

func (i InventoryPosition) String() string {
	return Stringify(i)
}

func (i InventoryPosition) MetaType() MetaType {
	return MetaTypeInventoryPosition
}
