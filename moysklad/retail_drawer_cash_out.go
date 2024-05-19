package moysklad

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// RetailDrawerCashOut Выплата денег.
// Ключевое слово: retaildrawercashout
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vyplata-deneg
type RetailDrawerCashOut struct {
	AccountID    *uuid.UUID       `json:"accountId,omitempty"`    // ID учетной записи
	Agent        *Counterparty    `json:"agent,omitempty"`        // Метаданные контрагента
	Applicable   *bool            `json:"applicable,omitempty"`   // Отметка о проведении
	Attributes   *Attributes      `json:"attributes,omitempty"`   // Коллекция метаданных доп. полей
	Code         *string          `json:"code,omitempty"`         // Код Выплаты денег
	Created      *Timestamp       `json:"created,omitempty"`      // Дата создания
	Deleted      *Timestamp       `json:"deleted,omitempty"`      // Момент последнего удаления Выплаты денег
	Description  *string          `json:"description,omitempty"`  // Комментарий Выплаты денег
	ExternalCode *string          `json:"externalCode,omitempty"` // Внешний код Выплаты денег
	Files        *Files           `json:"files,omitempty"`        // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group        *Group           `json:"group,omitempty"`        // Отдел сотрудника
	ID           *uuid.UUID       `json:"id,omitempty"`           // ID сущности
	Meta         *Meta            `json:"meta,omitempty"`         // Метаданные
	Moment       *Timestamp       `json:"moment,omitempty"`       // Дата документа
	Name         *string          `json:"name,omitempty"`         // Наименование
	Organization *Organization    `json:"organization,omitempty"` // Метаданные юрлица
	Owner        *Employee        `json:"owner,omitempty"`        // Владелец (Сотрудник)
	Printed      *bool            `json:"printed,omitempty"`      // Напечатан ли документ
	Published    *bool            `json:"published,omitempty"`    // Опубликован ли документ
	Rate         *Rate            `json:"rate,omitempty"`         // Валюта
	Shared       *bool            `json:"shared,omitempty"`       // Общий доступ
	State        *State           `json:"state,omitempty"`        // Метаданные статуса Выплаты денег
	Sum          *decimal.Decimal `json:"sum,omitempty"`          // Сумма Выплаты денег установленной валюте
	SyncID       *uuid.UUID       `json:"syncId,omitempty"`       // ID синхронизации. После заполнения недоступен для изменения
	Updated      *Timestamp       `json:"updated,omitempty"`      // Момент последнего обновления
	RetailShift  *RetailShift     `json:"retailShift,omitempty"`  // Ссылка на розничную смену, в рамках которой было выполнено Внесение денег в формате Метаданных
}

func (r RetailDrawerCashOut) String() string {
	return Stringify(r)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (r RetailDrawerCashOut) GetMeta() *Meta {
	return r.Meta
}

func (r RetailDrawerCashOut) MetaType() MetaType {
	return MetaTypeRetailDrawerCashOut
}
