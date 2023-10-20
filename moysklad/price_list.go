package moysklad

import (
	"github.com/google/uuid"
)

// PriceList Прайс-лист.
// Ключевое слово: pricelist
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-prajs-list
type PriceList struct {
	AccountID    *uuid.UUID                    `json:"accountId,omitempty"`    // ID учетной записи
	Applicable   *bool                         `json:"applicable,omitempty"`   // Отметка о проведении
	Attributes   *Attributes                   `json:"attributes,omitempty"`   // Коллекция метаданных доп. полей
	Code         *string                       `json:"code,omitempty"`         // Код
	Columns      *PriceListColumns             `json:"columns,omitempty"`      // Массив столбцов описания таблицы
	Created      *Timestamp                    `json:"created,omitempty"`      // Дата создания
	Deleted      *Timestamp                    `json:"deleted,omitempty"`      // Момент последнего удаления
	Description  *string                       `json:"description,omitempty"`  // Комментарий
	ExternalCode *string                       `json:"externalCode,omitempty"` // Внешний код
	Files        *Files                        `json:"files,omitempty"`        // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group        *Group                        `json:"group,omitempty"`        // Отдел сотрудника
	ID           *uuid.UUID                    `json:"id,omitempty"`           // ID сущности
	Meta         *Meta                         `json:"meta,omitempty"`         // Метаданные
	Moment       *Timestamp                    `json:"moment,omitempty"`       // Дата документа
	Name         *string                       `json:"name,omitempty"`         // Наименование
	Organization *Organization                 `json:"organization,omitempty"` // Метаданные юрлица
	Owner        *Employee                     `json:"owner,omitempty"`        // Владелец (Сотрудник)
	Positions    *Positions[PriceListPosition] `json:"positions,omitempty"`    // Метаданные позиций Прайс-листа
	PriceType    *PriceType                    `json:"priceType,omitempty"`    // Объект типа цены
	Printed      *bool                         `json:"printed,omitempty"`      // Напечатан ли документ
	Published    *bool                         `json:"published,omitempty"`    // Опубликован ли документ
	Shared       *bool                         `json:"shared,omitempty"`       // Общий доступ
	State        *State                        `json:"state,omitempty"`        // Метаданные статуса Прайс-листа
	SyncID       *uuid.UUID                    `json:"syncId,omitempty"`       // ID синхронизации. После заполнения недоступен для изменения
	Updated      *Timestamp                    `json:"updated,omitempty"`      // Момент последнего обновления
}

func (p PriceList) String() string {
	return Stringify(p)
}

func (p PriceList) MetaType() MetaType {
	return MetaTypePriceList
}

// PriceListCell Ячейка прайс листа.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-prajs-list-prajs-listy-yachejki
type PriceListCell struct {
	Column *string  `json:"column,omitempty"` // Название столбца, к которому относится данная ячейка
	Sum    *float64 `json:"sum,omitempty"`    // Числовое значение ячейки
}

func (p PriceListCell) String() string {
	return Stringify(p)
}

type PriceListCells = Iterator[PriceListCell]

// PriceListColumn Столбец прайс листа.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-prajs-list-prajs-listy-stolbcy
type PriceListColumn struct {
	Name               *string `json:"name,omitempty"`               // Название столбца
	PercentageDiscount *int    `json:"percentageDiscount,omitempty"` // Процентная наценка или скидка по умолчанию для столбца
}

func (p PriceListColumn) String() string {
	return Stringify(p)
}

type PriceListColumns = Iterator[PriceListColumn]

// PriceListPosition Позиция прайс листа.
// Ключевое слово: pricelistrow
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-prajs-list-prajs-listy-pozicii-prajs-lista
type PriceListPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Товар/услуга/модификация, которую представляет собой позиция
	Cells      *PriceListCells     `json:"cells,omitempty"`      // Значения столбцов
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка товара
}

func (p PriceListPosition) String() string {
	return Stringify(p)
}

func (p PriceListPosition) MetaType() MetaType {
	return MetaTypePriceListPosition
}
