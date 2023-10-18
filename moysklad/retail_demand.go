package moysklad

import (
	"github.com/google/uuid"
)

// RetailDemand Розничная продажа.
// Ключевое слово: retaildemand
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-roznichnaq-prodazha
type RetailDemand struct {
	AccountId           *uuid.UUID                 `json:"accountId,omitempty"`           // ID учетной записи
	Agent               *Counterparty              `json:"agent,omitempty"`               // Метаданные контрагента
	AgentAccount        *AgentAccount              `json:"agentAccount,omitempty"`        // Метаданные счета контрагента
	Applicable          *bool                      `json:"applicable,omitempty"`          // Отметка о проведении
	Attributes          *Attributes                `json:"attributes,omitempty"`          // Коллекция метаданных доп. полей. Поля объекта
	CashSum             *float64                   `json:"cashSum,omitempty"`             // Оплачено наличными
	CheckNumber         *string                    `json:"checkNumber,omitempty"`         // Номер чека
	CheckSum            *float64                   `json:"checkSum,omitempty"`            // Сумма Чека
	Code                *string                    `json:"code,omitempty"`                // Код Розничной продажи
	Contract            *Contract                  `json:"contract,omitempty"`            // Метаданные договора
	Created             *Timestamp                 `json:"created,omitempty"`             // Дата создания
	CustomerOrder       *CustomerOrder             `json:"customerOrder,omitempty"`       // Метаданные Заказа Покупателя
	Deleted             *Timestamp                 `json:"deleted,omitempty"`             // Момент последнего удаления Розничной продажи
	Description         *string                    `json:"description,omitempty"`         // Комментарий Розничной продажи
	DocumentNumber      *string                    `json:"documentNumber,omitempty"`      // Номер документа
	ExternalCode        *string                    `json:"externalCode,omitempty"`        // Внешний код Розничной продажи
	Files               *Files                     `json:"files,omitempty"`               // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Fiscal              *bool                      `json:"fiscal,omitempty"`              // Отметка о том, был ли использован ФР
	FiscalPrinterInfo   *string                    `json:"fiscalPrinterInfo,omitempty"`   // Информация о фискальном регистраторе
	Group               *Group                     `json:"group,omitempty"`               // Отдел сотрудника
	Id                  *uuid.UUID                 `json:"id,omitempty"`                  // ID сущности
	Meta                *Meta                      `json:"meta,omitempty"`                // Метаданные
	Moment              *Timestamp                 `json:"moment,omitempty"`              // Дата документа
	Name                *string                    `json:"name,omitempty"`                // Наименование
	NoCashSum           *float64                   `json:"noCashSum,omitempty"`           // Оплачено картой
	OfdCode             *string                    `json:"ofdCode,omitempty"`             // Код оператора фискальных данных
	Organization        *Organization              `json:"organization,omitempty"`        // Метаданные юрлица
	OrganizationAccount *AgentAccount              `json:"organizationAccount,omitempty"` // Метаданные счета юрлица
	Owner               *Employee                  `json:"owner,omitempty"`               // Владелец (Сотрудник)
	PayedSum            *float64                   `json:"payedSum,omitempty"`            // Сумма входящих платежей
	Positions           *Positions[RetailPosition] `json:"positions,omitempty"`           // Метаданные позиций Розничной продажи
	PrepaymentCashSum   *float64                   `json:"prepaymentCashSum,omitempty"`   // Предоплата наличными
	PrepaymentNoCashSum *float64                   `json:"prepaymentNoCashSum,omitempty"` // Предоплата картой
	PrepaymentQRSum     *float64                   `json:"prepaymentQrSum,omitempty"`     // Предоплата по QR-коду
	Printed             *bool                      `json:"printed,omitempty"`             // Напечатан ли документ
	Project             *Project                   `json:"project,omitempty"`             // Метаданные проекта
	Published           *bool                      `json:"published,omitempty"`           // Опубликован ли документ
	QRSum               *float64                   `json:"qrSum,omitempty"`               // Оплачено по QR-коду
	Rate                *Rate                      `json:"rate,omitempty"`                // Валюта
	RetailShift         *RetailShift               `json:"retailShift,omitempty"`         // Метаданные Розничной смены
	RetailStore         *RetailStore               `json:"retailStore,omitempty"`         // Метаданные Точки продаж
	SessionNumber       *string                    `json:"sessionNumber,omitempty"`       // Номер сессии
	Shared              *bool                      `json:"shared,omitempty"`              // Общий доступ
	State               *State                     `json:"state,omitempty"`               // Метаданные статуса Розничной продажи
	Store               *Store                     `json:"store,omitempty"`               // Метаданные склада
	Sum                 *float64                   `json:"sum,omitempty"`                 // Сумма
	SyncId              *uuid.UUID                 `json:"syncId,omitempty"`              // ID синхронизации. После заполнения недоступен для изменения
	TaxSystem           TaxSystem                  `json:"taxSystem,omitempty"`           // Код системы налогообложения
	Updated             *Timestamp                 `json:"updated,omitempty"`             // Момент последнего обновления
	VatEnabled          *bool                      `json:"vatEnabled,omitempty"`          // Учитывается ли НДС
	VatIncluded         *bool                      `json:"vatIncluded,omitempty"`         // Включен ли НДС в цену
	VatSum              *float64                   `json:"vatSum,omitempty"`              // Сумма включая НДС
}

func (r RetailDemand) String() string {
	return Stringify(r)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (r RetailDemand) GetMeta() *Meta {
	return r.Meta
}

func (r RetailDemand) MetaType() MetaType {
	return MetaTypeRetailDemand
}

// RetailDemandPosition позиция розничной продажи.
// Ключевое слово: demandposition
type RetailDemandPosition struct {
	RetailPosition
}

func (r RetailDemandPosition) MetaType() MetaType {
	return MetaTypeRetailDemandPosition
}

// RetailDemandTemplateArg
// Документ: Розничная продажа (retaildemand)
// Основание, на котором он может быть создан:
// - Розничная смена
// - Заказ покупателя
type RetailDemandTemplateArg struct {
	RetailShift   *MetaWrapper `json:"retailShift,omitempty"`
	CustomerOrder *MetaWrapper `json:"customerOrder,omitempty"`
}
