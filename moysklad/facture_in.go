package moysklad

import (
	"github.com/google/uuid"
)

// FactureIn Счет-фактура полученный
// Ключевое слово: facturein
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-schet-faktura-poluchennyj
type FactureIn struct {
	AccountId      *uuid.UUID    `json:"accountId,omitempty"`      // ID учетной записи
	Agent          *Counterparty `json:"agent,omitempty"`          // Метаданные контрагента
	Applicable     *bool         `json:"applicable,omitempty"`     // Отметка о проведении
	Attributes     *Attributes   `json:"attributes,omitempty"`     // Коллекция метаданных доп. полей объекта
	Code           *string       `json:"code,omitempty"`           // Код выданного Счета-фактуры полученного
	Contract       *Contract     `json:"contract,omitempty"`       // Метаданные договора
	Created        *Timestamp    `json:"created,omitempty"`        // Дата создания
	Deleted        *Timestamp    `json:"deleted,omitempty"`        // Момент последнего удаления Счета-фактуры полученного
	Description    *string       `json:"description,omitempty"`    // Комментарий выданного Счета-фактуры полученного
	ExternalCode   *string       `json:"externalCode,omitempty"`   // Внешний код выданного Счета-фактуры полученного
	Files          *Files        `json:"files,omitempty"`          // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group          *Group        `json:"group,omitempty"`          // Отдел сотрудника
	Id             *uuid.UUID    `json:"id,omitempty"`             // ID сущности
	Meta           *Meta         `json:"meta,omitempty"`           // Метаданные
	Moment         *Timestamp    `json:"moment,omitempty"`         // Дата документа
	Name           *string       `json:"name,omitempty"`           // Наименование
	Organization   *Organization `json:"organization,omitempty"`   // Метаданные юрлица
	Owner          *Employee     `json:"owner,omitempty"`          // Владелец (Сотрудник)
	Printed        *bool         `json:"printed,omitempty"`        // Напечатан ли документ
	Published      *bool         `json:"published,omitempty"`      // Опубликован ли документ
	Rate           *Rate         `json:"rate,omitempty"`           // Валюта
	Shared         *bool         `json:"shared,omitempty"`         // Общий доступ
	State          *State        `json:"state,omitempty"`          // Метаданные статуса Счета-фактуры полученного
	Sum            *float64      `json:"sum,omitempty"`            // Сумма
	SyncId         *uuid.UUID    `json:"syncId,omitempty"`         // ID синхронизации. После заполнения недоступен для изменения
	Updated        *Timestamp    `json:"updated,omitempty"`        // Момент последнего обновления
	Supplies       *Supplies     `json:"supplies,omitempty"`       // Массив ссылок на связанные приемки в формате Метаданных
	Payments       *Payments     `json:"payments,omitempty"`       // Связанные исходящие платежи и расходные ордеры
	IncomingNumber *string       `json:"incomingNumber,omitempty"` // Входящий номер
	IncomingDate   *Timestamp    `json:"incomingDate,omitempty"`   // Входящая дата
}

func (f FactureIn) String() string {
	return Stringify(f)
}

func (f FactureIn) MetaType() MetaType {
	return MetaTypeFactureIn
}
