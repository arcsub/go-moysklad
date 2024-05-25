package moysklad

import (
	"github.com/google/uuid"
)

// Move Перемещение.
// Ключевое слово: move
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-peremeschenie
type Move struct {
	AccountID     *uuid.UUID               `json:"accountId,omitempty"`     // ID учетной записи
	Applicable    *bool                    `json:"applicable,omitempty"`    // Отметка о проведении
	Attributes    *Attributes              `json:"attributes,omitempty"`    // Коллекция метаданных доп. полей. Поля объекта
	Code          *string                  `json:"code,omitempty"`          // Код
	Created       *Timestamp               `json:"created,omitempty"`       // Дата создания
	Deleted       *Timestamp               `json:"deleted,omitempty"`       // Момент последнего удаления
	Description   *string                  `json:"description,omitempty"`   // Комментарий
	ExternalCode  *string                  `json:"externalCode,omitempty"`  // Внешний код
	Files         *Files                   `json:"files,omitempty"`         // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group         *Group                   `json:"group,omitempty"`         // Отдел сотрудника
	ID            *uuid.UUID               `json:"id,omitempty"`            // ID сущности
	InternalOrder *InternalOrder           `json:"internalOrder,omitempty"` // Метаданные Внутреннего заказа, связанного с Перемещением
	CustomerOrder *CustomerOrder           `json:"customerOrder,omitempty"` // Метаданные Заказа покупателя, связанного с Перемещением
	Meta          *Meta                    `json:"meta,omitempty"`          // Метаданные
	Moment        *Timestamp               `json:"moment,omitempty"`        // Дата документа
	Name          *string                  `json:"name,omitempty"`          // Наименование
	Organization  *Organization            `json:"organization,omitempty"`  // Метаданные юрлица
	Overhead      *Overhead                `json:"overhead,omitempty"`      // Накладные расходы. Если Позиции Перемещения не заданы, то накладные расходы нельзя задать
	Owner         *Employee                `json:"owner,omitempty"`         // Владелец (Сотрудник)
	Positions     *Positions[MovePosition] `json:"positions,omitempty"`     // Метаданные позиций Перемещения
	Printed       *bool                    `json:"printed,omitempty"`       // Напечатан ли документ
	Project       *Project                 `json:"project,omitempty"`       // Проект
	Published     *bool                    `json:"published,omitempty"`     // Опубликован ли документ
	Rate          *Rate                    `json:"rate,omitempty"`          // Валюта
	Shared        *bool                    `json:"shared,omitempty"`        // Общий доступ
	SourceStore   *Store                   `json:"sourceStore,omitempty"`   // Метаданные склада, с которого совершается перемещение
	State         *State                   `json:"state,omitempty"`         // Метаданные статуса
	Sum           *Decimal                 `json:"sum,omitempty"`           // Сумма
	SyncID        *uuid.UUID               `json:"syncId,omitempty"`        // ID синхронизации. После заполнения недоступен для изменения
	TargetStore   *Store                   `json:"targetStore,omitempty"`   // Метаданные склада, на который совершается перемещение
	Updated       *Timestamp               `json:"updated,omitempty"`       // Момент последнего обновления
}

func (m Move) String() string {
	return Stringify(m)
}

func (m Move) MetaType() MetaType {
	return MetaTypeMove
}

type Moves Slice[Move]

// MovePosition Позиция перемещения.
// Ключевое слово: moveposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-peremeschenie-peremescheniq-pozicii-peremescheniq
type MovePosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Overhead   *Decimal            `json:"overhead,omitempty"`   // Накладные расходы. Если Позиции Перемещения не заданы, то накладные расходы нельзя задать
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Price      *Decimal            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе
	SourceSlot *Slot               `json:"sourceSlot,omitempty"` // Ячейка на складе, с которого совершается перемещение
	TargetSlot *Slot               `json:"targetSlot,omitempty"` // Ячейка на складе, на который совершается перемещение
	Things     *Things             `json:"things,omitempty"`     // Серийные номера. Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете. В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута
}

func (m MovePosition) String() string {
	return Stringify(m)
}

func (m MovePosition) MetaType() MetaType {
	return MetaTypeMovePosition
}

// MoveTemplateArg
// Документ: Перемещение (move)
// Основание, на котором он может быть создан:
// - Внутренний заказ (internalorder)
type MoveTemplateArg struct {
	InternalOrder *MetaWrapper `json:"internalOrder,omitempty"`
}
