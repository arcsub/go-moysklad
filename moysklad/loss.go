package moysklad

import (
	"github.com/google/uuid"
)

// Loss Списание.
// Ключевое слово: loss
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-spisanie
type Loss struct {
	AccountID    *uuid.UUID               `json:"accountId,omitempty"`    // ID учетной записи
	Applicable   *bool                    `json:"applicable,omitempty"`   // Отметка о проведении
	Attributes   *Attributes              `json:"attributes,omitempty"`   // Коллекция метаданных доп. полей. Поля объекта
	Code         *string                  `json:"code,omitempty"`         // Код
	Created      *Timestamp               `json:"created,omitempty"`      // Дата создания
	Deleted      *Timestamp               `json:"deleted,omitempty"`      // Момент последнего удаления
	Description  *string                  `json:"description,omitempty"`  // Комментарий
	ExternalCode *string                  `json:"externalCode,omitempty"` // Внешний код
	Files        *Files                   `json:"files,omitempty"`        // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group        *Group                   `json:"group,omitempty"`        // Отдел сотрудника
	ID           *uuid.UUID               `json:"id,omitempty"`           // ID сущности
	Meta         *Meta                    `json:"meta,omitempty"`         // Метаданные
	Moment       *Timestamp               `json:"moment,omitempty"`       // Дата документа
	Name         *string                  `json:"name,omitempty"`         // Наименование
	Organization *Organization            `json:"organization,omitempty"` // Метаданные юрлица
	Owner        *Employee                `json:"owner,omitempty"`        // Владелец (Сотрудник)
	Positions    *Positions[LossPosition] `json:"positions,omitempty"`    // Метаданные позиций Списания
	Printed      *bool                    `json:"printed,omitempty"`      // Напечатан ли документ
	Project      *Project                 `json:"project,omitempty"`      // Проект
	Published    *bool                    `json:"published,omitempty"`    // Опубликован ли документ
	Rate         *Rate                    `json:"rate,omitempty"`         // Валюта
	Shared       *bool                    `json:"shared,omitempty"`       // Общий доступ
	State        *State                   `json:"state,omitempty"`        // Метаданные статуса
	Store        *Store                   `json:"store,omitempty"`        // Метаданные склада
	Sum          *Decimal                 `json:"sum,omitempty"`          // Сумма
	SyncID       *uuid.UUID               `json:"syncId,omitempty"`       // ID синхронизации. После заполнения недоступен для изменения
	Updated      *Timestamp               `json:"updated,omitempty"`      // Момент последнего обновления
	SalesReturn  *SalesReturn             `json:"salesReturn,omitempty"`  // Ссылка на связанный со списанием возврат покупателя в формате Метаданных
}

func (l Loss) String() string {
	return Stringify(l)
}

func (l Loss) MetaType() MetaType {
	return MetaTypeLoss
}

type Losses = Slice[Loss]

// LossPosition Позиция Списания.
// Ключевое слово: lossposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-spisanie-spisaniq-pozicii-spisaniq
type LossPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID сущности
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Price      *Decimal            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Reason     *string             `json:"reason,omitempty"`     // Причина списания данной позиции
	Slot       *Slot               `json:"slot,omitempty"`       // Ячейка на складе
	Things     *Things             `json:"things,omitempty"`     // Серийные номера. Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете. В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута.
}

func (l LossPosition) String() string {
	return Stringify(l)
}

func (l LossPosition) MetaType() MetaType {
	return MetaTypeLossPosition
}

// LossTemplateArg
// Документ: Списание (loss)
// Основание, на котором он может быть создан:
// - Возврат покупателя (salesreturn)
// - инвентаризация(inventory)
type LossTemplateArg struct {
	SalesReturn *MetaWrapper `json:"salesReturn,omitempty"`
	Inventory   *MetaWrapper `json:"inventory,omitempty"`
}
