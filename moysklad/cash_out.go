package moysklad

import (
	"github.com/google/uuid"
)

// CashOut Расходный ордер.
// Ключевое слово: cashout
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-rashodnyj-order
type CashOut struct {
	AccountId      *uuid.UUID    `json:"accountId,omitempty"`      // ID учетной записи
	Agent          *Counterparty `json:"agent,omitempty"`          // Метаданные контрагента
	Applicable     *bool         `json:"applicable,omitempty"`     // Отметка о проведении
	Attributes     *Attributes   `json:"attributes,omitempty"`     // Коллекция метаданных доп. полей. Поля объекта
	Code           *string       `json:"code,omitempty"`           // Код
	Contract       *Contract     `json:"contract,omitempty"`       // Метаданные договора
	Created        *Timestamp    `json:"created,omitempty"`        // Дата создания
	Deleted        *Timestamp    `json:"deleted,omitempty"`        // Момент последнего удаления
	Description    *string       `json:"description,omitempty"`    // Комментарий
	ExpenseItem    *ExpenseItem  `json:"expenseItem,omitempty"`    // Метаданные Статьи расходов
	ExternalCode   *string       `json:"externalCode,omitempty"`   // Внешний код
	Files          *Files        `json:"files,omitempty"`          // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group          *Group        `json:"group,omitempty"`          // Отдел сотрудника
	Id             *uuid.UUID    `json:"id,omitempty"`             // ID сущности
	Meta           *Meta         `json:"meta,omitempty"`           // Метаданные
	Moment         *Timestamp    `json:"moment,omitempty"`         // Дата документа
	Name           *string       `json:"name,omitempty"`           // Наименование
	Organization   *Organization `json:"organization,omitempty"`   // Метаданные юрлица
	Owner          *Employee     `json:"owner,omitempty"`          // Владелец (Сотрудник)
	PaymentPurpose *string       `json:"paymentPurpose,omitempty"` // Основание
	Printed        *bool         `json:"printed,omitempty"`        // Напечатан ли документ
	Project        *Project      `json:"project,omitempty"`        // Метаданные проекта
	Published      *bool         `json:"published,omitempty"`      // Опубликован ли документ
	Rate           *Rate         `json:"rate,omitempty"`           // Валюта
	SalesChannel   *SalesChannel `json:"salesChannel,omitempty"`   // Метаданные канала продаж
	Shared         *bool         `json:"shared,omitempty"`         // Общий доступ
	State          *State        `json:"state,omitempty"`          // Метаданные статуса
	Sum            *float64      `json:"sum,omitempty"`            // Сумма
	SyncId         *uuid.UUID    `json:"syncId,omitempty"`         // ID синхронизации. После заполнения недоступен для изменения
	Updated        *Timestamp    `json:"updated,omitempty"`        // Момент последнего обновления
	VatSum         *float64      `json:"vatSum,omitempty"`         // Сумма включая НДС
	FactureOut     *FactureOut   `json:"factureOut,omitempty"`     // Ссылка на Счет-фактуру выданный, с которым связан этот платеж в формате Метаданных
	//Operations     *Operations   `json:"operations,omitempty"`     // Массив ссылок на связанные операции в формате Метаданных
}

func (c CashOut) String() string {
	return Stringify(c)
}

func (c CashOut) MetaType() MetaType {
	return MetaTypeCashOut
}
