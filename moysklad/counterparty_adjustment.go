package moysklad

import (
	"github.com/google/uuid"
)

// CounterPartyAdjustment Корректировка баланса контрагента.
// Ключевое слово: counterpartyadjustment
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-korrektirowka-balansa-kontragenta
type CounterPartyAdjustment struct {
	AccountId    *uuid.UUID    `json:"accountId,omitempty"`    // ID учетной записи
	Id           *uuid.UUID    `json:"id,omitempty"`           // ID сущности
	Name         *string       `json:"name,omitempty"`         // Наименование
	Meta         *Meta         `json:"meta,omitempty"`         // Метаданные
	Agent        *Counterparty `json:"agent,omitempty"`        // Метаданные контрагента
	Applicable   *bool         `json:"applicable,omitempty"`   // Отметка о проведении
	Attributes   *Attributes   `json:"attributes,omitempty"`   // Коллекция метаданных доп. полей. Поля объекта
	Created      *Timestamp    `json:"created,omitempty"`      // Дата создания
	Deleted      *Timestamp    `json:"deleted,omitempty"`      // Момент последнего удаления Корректировки баланса контрагента
	Description  *string       `json:"description,omitempty"`  // Комментарий Корректировки баланса контрагента
	ExternalCode *string       `json:"externalCode,omitempty"` // Внешний код Корректировки баланса контрагента
	Files        *Files        `json:"files,omitempty"`        // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group        *Group        `json:"group,omitempty"`        // Отдел сотрудника
	Moment       *Timestamp    `json:"moment,omitempty"`       // Дата документа
	Organization *Organization `json:"organization,omitempty"` // Метаданные юрлица
	Owner        *Employee     `json:"owner,omitempty"`        // Владелец (Сотрудник)
	Printed      *bool         `json:"printed,omitempty"`      // Напечатан ли документ
	Published    *bool         `json:"published,omitempty"`    // Опубликован ли документ
	Shared       *bool         `json:"shared,omitempty"`       // Общий доступ
	Sum          *float64      `json:"sum,omitempty"`          // Сумма Корректировки баланса контрагента в копейках
	Updated      *Timestamp    `json:"updated,omitempty"`      // Момент последнего обновления Корректировки баланса контрагента
}

func (c CounterPartyAdjustment) String() string {
	return Stringify(c)
}

func (c CounterPartyAdjustment) MetaType() MetaType {
	return MetaTypeCounterPartyAdjustment
}
