package moysklad

import (
	"github.com/google/uuid"
)

// FactureOut Счет-фактура выданный
// Ключевое слово: factureout
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-schet-faktura-wydannyj
type FactureOut struct {
	AccountID       *uuid.UUID       `json:"accountId,omitempty"`       // ID учетной записи
	Agent           *Counterparty    `json:"agent,omitempty"`           // Метаданные контрагента
	Applicable      *bool            `json:"applicable,omitempty"`      // Отметка о проведении
	Attributes      *Attributes      `json:"attributes,omitempty"`      // Коллекция метаданных доп. полей. Поля объекта
	Code            *string          `json:"code,omitempty"`            // Код выданного Счета-фактуры
	Contract        *Contract        `json:"contract,omitempty"`        // Метаданные договора
	Created         *Timestamp       `json:"created,omitempty"`         // Дата создания
	Deleted         *Timestamp       `json:"deleted,omitempty"`         // Момент последнего удаления Счета-фактуры
	Description     *string          `json:"description,omitempty"`     // Комментарий выданного Счета-фактуры
	ExternalCode    *string          `json:"externalCode,omitempty"`    // Внешний код выданного Счета-фактуры
	Files           *Files           `json:"files,omitempty"`           // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group           *Group           `json:"group,omitempty"`           // Отдел сотрудника
	ID              *uuid.UUID       `json:"id,omitempty"`              // ID сущности
	Meta            *Meta            `json:"meta,omitempty"`            // Метаданные
	Moment          *Timestamp       `json:"moment,omitempty"`          // Дата документа
	Name            *string          `json:"name,omitempty"`            // Наименование
	Organization    *Organization    `json:"organization,omitempty"`    // Метаданные юрлица
	Owner           *Employee        `json:"owner,omitempty"`           // Владелец (Сотрудник)
	Printed         *bool            `json:"printed,omitempty"`         // Напечатан ли документ
	Published       *bool            `json:"published,omitempty"`       // Опубликован ли документ
	Rate            *Rate            `json:"rate,omitempty"`            // Валюта
	Shared          *bool            `json:"shared,omitempty"`          // Общий доступ
	State           *State           `json:"state,omitempty"`           // Метаданные статуса Счета-фактуры
	StateContractId *string          `json:"stateContractId,omitempty"` // Идентификатор гос. контракта
	Sum             *float64         `json:"sum,omitempty"`             // Сумма
	SyncID          *uuid.UUID       `json:"syncId,omitempty"`          // ID синхронизации. После заполнения недоступен для изменения
	Updated         *Timestamp       `json:"updated,omitempty"`         // Момент последнего обновления
	Demands         Iterator[Demand] `json:"demands,omitempty"`         // Массив ссылок на связанные отгрузки в формате Метаданных
	Payments        *Payments        `json:"payments,omitempty"`        // Массив ссылок на связанные входящие платежи в формате Метаданных
	Returns         *PurchaseReturns `json:"returns,omitempty"`         // Массив ссылок на связанные возвраты поставщикам в формате Метаданных
	Consignee       *Counterparty    `json:"consignee,omitempty"`       // Грузополучатель
	PaymentNumber   *string          `json:"paymentNumber,omitempty"`   // Название платежного документа
	PaymentDate     *Timestamp       `json:"paymentDate,omitempty"`     // Дата платежного документа
}

func (f FactureOut) String() string {
	return Stringify(f)
}

func (f FactureOut) MetaType() MetaType {
	return MetaTypeFactureOut
}
