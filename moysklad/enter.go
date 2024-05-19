package moysklad

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Enter Оприходование.
// Ключевое слово: enter
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-oprihodowanie
type Enter struct {
	AccountID    *uuid.UUID                `json:"accountId,omitempty"`    // ID учетной записи
	Applicable   *bool                     `json:"applicable,omitempty"`   // Отметка о проведении
	Attributes   *Attributes               `json:"attributes,omitempty"`   // Коллекция метаданных доп. полей. Поля объекта
	Code         *string                   `json:"code,omitempty"`         // Код Оприходования
	Created      *Timestamp                `json:"created,omitempty"`      // Дата создания
	Deleted      *Timestamp                `json:"deleted,omitempty"`      // Момент последнего удаления Оприходования
	Description  *string                   `json:"description,omitempty"`  // Комментарий Оприходования
	ExternalCode *string                   `json:"externalCode,omitempty"` // Внешний код Оприходования
	Files        *Files                    `json:"files,omitempty"`        // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group        *Group                    `json:"group,omitempty"`        // Отдел сотрудника
	ID           *uuid.UUID                `json:"id,omitempty"`           // ID сущности
	Meta         *Meta                     `json:"meta,omitempty"`         // Метаданные
	Moment       *Timestamp                `json:"moment,omitempty"`       // Дата документа
	Name         *string                   `json:"name,omitempty"`         // Наименование
	Organization *Organization             `json:"organization,omitempty"` // Метаданные юрлица
	Overhead     *Overhead                 `json:"overhead,omitempty"`     // Накладные расходы. Если Позиции Оприходования не заданы, то накладные расходы нельзя задать
	Owner        *Employee                 `json:"owner,omitempty"`        // Владелец (Сотрудник)
	Positions    *Positions[EnterPosition] `json:"positions,omitempty"`    // Метаданные позиций Оприходования
	Printed      *bool                     `json:"printed,omitempty"`      // Напечатан ли документ
	Project      *Project                  `json:"project,omitempty"`      // Метаданные проекта
	Published    *bool                     `json:"published,omitempty"`    // Опубликован ли документ
	Rate         *Rate                     `json:"rate,omitempty"`         // Валюта
	Shared       *bool                     `json:"shared,omitempty"`       // Общий доступ
	State        *State                    `json:"state,omitempty"`        // Метаданные статуса оприходования
	Store        *Store                    `json:"store,omitempty"`        // Метаданные склада
	Sum          *decimal.Decimal          `json:"sum,omitempty"`          // Сумма
	SyncID       *uuid.UUID                `json:"syncId,omitempty"`       // ID синхронизации. После заполнения недоступен для изменения
	Updated      *Timestamp                `json:"updated,omitempty"`      // Момент последнего обновления
}

func (e Enter) String() string {
	return Stringify(e)
}

func (e Enter) MetaType() MetaType {
	return MetaTypeEnter
}

// EnterPosition Позиция оприходования
// Ключевое слово: enterposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-oprihodowanie-oprihodowaniq-pozicii-oprihodowaniq
type EnterPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	Country    *Country            `json:"country,omitempty"`    // Метаданные страны
	GTD        *GTD                `json:"gtd,omitempty"`        // ГТД
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Overhead   *decimal.Decimal    `json:"overhead,omitempty"`   // Накладные расходы. Если Позиции Оприходования не заданы, то накладные расходы нельзя задать
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Price      *decimal.Decimal    `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Reason     *string             `json:"reason,omitempty"`     // Причина оприходования данной позиции
	Slot       *Slot               `json:"slot,omitempty"`       // Ячейка на складе
	Things     *Things             `json:"things,omitempty"`     // Серийные номера. Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете. В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута.
}

func (e EnterPosition) String() string {
	return Stringify(e)
}

func (e EnterPosition) MetaType() MetaType {
	return MetaTypeEnterPosition
}

// EnterTemplateArg
// Документ: Оприходование (enter)
// Основание, на котором он может быть создан:
// - Инвентаризация(inventory)
type EnterTemplateArg struct {
	Inventory *MetaWrapper `json:"inventory,omitempty"`
}
