package moysklad

import (
	"github.com/google/uuid"
)

// RetailDrawerCashIn Внесение денег.
// Ключевое слово: retaildrawercashin
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vnesenie-deneg
type RetailDrawerCashIn struct {
	AccountId    *uuid.UUID    `json:"accountId,omitempty"`    // ID учетной записи
	Agent        *Counterparty `json:"agent,omitempty"`        // Метаданные контрагента
	Applicable   *bool         `json:"applicable,omitempty"`   // Отметка о проведении
	Attributes   *Attributes   `json:"attributes,omitempty"`   // Коллекция метаданных доп. полей
	Created      *Timestamp    `json:"created,omitempty"`      // Дата создания
	Deleted      *Timestamp    `json:"deleted,omitempty"`      // Момент последнего удаления Внесения денег
	Description  *string       `json:"description,omitempty"`  // Комментарий Внесения денег
	ExternalCode *string       `json:"externalCode,omitempty"` // Внешний код Внесения денег
	Files        *Files        `json:"files,omitempty"`        // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group        *Group        `json:"group,omitempty"`        // Отдел сотрудника
	Id           *uuid.UUID    `json:"id,omitempty"`           // ID сущности
	Meta         *Meta         `json:"meta,omitempty"`         // Метаданные
	Moment       *Timestamp    `json:"moment,omitempty"`       // Дата документа
	Name         *string       `json:"name,omitempty"`         // Наименование
	Organization *Organization `json:"organization,omitempty"` // Метаданные юрлица
	Owner        *Employee     `json:"owner,omitempty"`        // Владелец (Сотрудник)
	Printed      *bool         `json:"printed,omitempty"`      // Напечатан ли документ
	Published    *bool         `json:"published,omitempty"`    // Опубликован ли документ
	Rate         *Rate         `json:"rate,omitempty"`         // Валюта
	Shared       *bool         `json:"shared,omitempty"`       // Общий доступ
	State        *State        `json:"state,omitempty"`        // Метаданные статуса Внесения денег
	Sum          *float64      `json:"sum,omitempty"`          // Сумма
	SyncId       *uuid.UUID    `json:"syncId,omitempty"`       // ID синхронизации. После заполнения недоступен для изменения
	Updated      *Timestamp    `json:"updated,omitempty"`      // Момент последнего обновления
	RetailShift  *RetailShift  `json:"retailShift,omitempty"`  // Ссылка на розничную смену
}

func (r RetailDrawerCashIn) String() string {
	return Stringify(r)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (r RetailDrawerCashIn) GetMeta() *Meta {
	return r.Meta
}

func (r RetailDrawerCashIn) MetaType() MetaType {
	return MetaTypeRetailDrawerCashIn
}
