package moysklad

import (
	"github.com/google/uuid"
)

// RetailSalesReturn Розничный возврат.
// Ключевое слово: retailsalesreturn
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-roznichnyj-wozwrat
type RetailSalesReturn struct {
	AccountId           *uuid.UUID                 `json:"accountId,omitempty"`           // ID учетной записи
	Agent               *Counterparty              `json:"agent,omitempty"`               // Метаданные контрагента
	AgentAccount        *AgentAccount              `json:"agentAccount,omitempty"`        // Метаданные счета контрагента
	Applicable          *bool                      `json:"applicable,omitempty"`          // Отметка о проведении
	Attributes          *Attributes                `json:"attributes,omitempty"`          // Коллекция метаданных доп. полей. Поля объекта
	CashSum             *float64                   `json:"cashSum,omitempty"`             // Оплачено наличными
	Code                *string                    `json:"code,omitempty"`                // Код Розничного возврата
	Contract            *Contract                  `json:"contract,omitempty"`            // Метаданные договора
	Created             *Timestamp                 `json:"created,omitempty"`             // Дата создания
	Deleted             *Timestamp                 `json:"deleted,omitempty"`             // Момент последнего удаления Розничного возврата
	Demand              *RetailDemand              `json:"demand,omitempty"`              // Метаданные позиций Розничного возврата, по которой произошел возврат
	Description         *string                    `json:"description,omitempty"`         // Комментарий Розничного возврата
	ExternalCode        *string                    `json:"externalCode,omitempty"`        // Внешний код Розничного возврата
	Group               *Group                     `json:"group,omitempty"`               // Отдел сотрудника
	Id                  *uuid.UUID                 `json:"id,omitempty"`                  // ID сущности
	Meta                *Meta                      `json:"meta,omitempty"`                // Метаданные
	Moment              *Timestamp                 `json:"moment,omitempty"`              // Дата документа
	Name                *string                    `json:"name,omitempty"`                // Наименование
	NoCashSum           *float64                   `json:"noCashSum,omitempty"`           // Оплачено картой
	Organization        *Organization              `json:"organization,omitempty"`        // Метаданные юрлица
	OrganizationAccount *AgentAccount              `json:"organizationAccount,omitempty"` // Метаданные счета юрлица
	Owner               *Employee                  `json:"owner,omitempty"`               // Владелец (Сотрудник)
	Positions           *Positions[RetailPosition] `json:"positions,omitempty"`           // Метаданные позиций Розничного возврата
	Printed             *bool                      `json:"printed,omitempty"`             // Напечатан ли документ
	Project             *Project                   `json:"project,omitempty"`             // Метаданные проекта
	Published           *bool                      `json:"published,omitempty"`           // Опубликован ли документ
	QrSum               *float64                   `json:"qrSum,omitempty"`               // Оплачено по QR-коду
	Rate                *Rate                      `json:"rate,omitempty"`                // Валюта
	RetailShift         *RetailShift               `json:"retailShift,omitempty"`         // Метаданные Розничной смены
	RetailStore         *RetailStore               `json:"retailStore,omitempty"`         // Метаданные Точки продаж
	Shared              *bool                      `json:"shared,omitempty"`              // Общий доступ
	State               *State                     `json:"state,omitempty"`               // Метаданные статуса Розничного возврата
	Store               *Store                     `json:"store,omitempty"`               // Метаданные склада
	Sum                 *float64                   `json:"sum,omitempty"`                 // Сумма
	SyncId              *uuid.UUID                 `json:"syncId,omitempty"`              // ID синхронизации. После заполнения недоступен для изменения
	TaxSystem           TaxSystem                  `json:"taxSystem,omitempty"`           // Код системы налогообложения
	Updated             *Timestamp                 `json:"updated,omitempty"`             // Момент последнего обновления
	VatEnabled          *bool                      `json:"vatEnabled,omitempty"`          // Учитывается ли НДС
	VatIncluded         *bool                      `json:"vatIncluded,omitempty"`         // Включен ли НДС в цену
	VatSum              *float64                   `json:"vatSum,omitempty"`              // Сумма включая НДС
}

func (r RetailSalesReturn) String() string {
	return Stringify(r)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (r RetailSalesReturn) GetMeta() *Meta {
	return r.Meta
}

func (r RetailSalesReturn) MetaType() MetaType {
	return MetaTypeRetailSalesReturn
}

// RetailSalesReturnPosition позиция розничного возврата.
// Ключевое слово: salesreturnposition
type RetailSalesReturnPosition struct {
	RetailPosition
}

func (r RetailSalesReturnPosition) MetaType() MetaType {
	return MetaTypeRetailSalesReturnPosition
}
