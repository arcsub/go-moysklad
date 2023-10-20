package moysklad

import (
	"github.com/google/uuid"
)

// InvoiceIn Счет поставщика.
// Ключевое слово: invoicein
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-schet-postawschika
type InvoiceIn struct {
	AccountID            *uuid.UUID                  `json:"accountId,omitempty"`            // ID учетной записи
	Agent                *Counterparty               `json:"agent,omitempty"`                // Метаданные контрагента
	AgentAccount         *AgentAccount               `json:"agentAccount,omitempty"`         // Метаданные счета контрагента
	Applicable           *bool                       `json:"applicable,omitempty"`           // Отметка о проведении
	Attributes           *Attributes                 `json:"attributes,omitempty"`           // Коллекция метаданных доп. полей
	Code                 *string                     `json:"code,omitempty"`                 // Код
	Contract             *Contract                   `json:"contract,omitempty"`             // Метаданные договора
	Created              *Timestamp                  `json:"created,omitempty"`              // Дата создания
	Deleted              *Timestamp                  `json:"deleted,omitempty"`              // Момент последнего удаления
	Description          *string                     `json:"description,omitempty"`          // Комментарий
	ExternalCode         *string                     `json:"externalCode,omitempty"`         // Внешний код
	Files                *Files                      `json:"files,omitempty"`                // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group                *Group                      `json:"group,omitempty"`                // Отдел сотрудника
	ID                   *uuid.UUID                  `json:"id,omitempty"`                   // ID сущности
	IncomingDate         *Timestamp                  `json:"incomingDate,omitempty"`         // Входящая дата
	IncomingNumber       *string                     `json:"incomingNumber,omitempty"`       // Входящий номер
	Meta                 *Meta                       `json:"meta,omitempty"`                 // Метаданные
	Moment               *Timestamp                  `json:"moment,omitempty"`               // Дата документа
	Name                 *string                     `json:"name,omitempty"`                 // Наименование
	Organization         *Organization               `json:"organization,omitempty"`         // Метаданные юрлица
	OrganizationAccount  *AgentAccount               `json:"organizationAccount,omitempty"`  // Метаданные счета юрлица
	Owner                *Employee                   `json:"owner,omitempty"`                // Владелец (Сотрудник)
	PayedSum             *float64                    `json:"payedSum,omitempty"`             // Сумма входящих платежей по Счету
	PaymentPlannedMoment *Timestamp                  `json:"paymentPlannedMoment,omitempty"` // Планируемая дата оплаты
	Positions            *Positions[InvoicePosition] `json:"positions,omitempty"`            // Метаданные позиций
	Printed              *bool                       `json:"printed,omitempty"`              // Напечатан ли документ
	Project              *Project                    `json:"project,omitempty"`              // Проект
	Published            *bool                       `json:"published,omitempty"`            // Опубликован ли документ
	Rate                 *Rate                       `json:"rate,omitempty"`                 // Валюта
	Shared               *bool                       `json:"shared,omitempty"`               // Общий доступ
	ShippedSum           *float64                    `json:"shippedSum,omitempty"`           // Сумма отгруженного
	State                *State                      `json:"state,omitempty"`                // Метаданные статуса
	Store                *Store                      `json:"store,omitempty"`                // Метаданные склада
	Sum                  *float64                    `json:"sum,omitempty"`                  // Сумма
	SyncID               *uuid.UUID                  `json:"syncId,omitempty"`               // ID синхронизации. После заполнения недоступен для изменения
	Updated              *Timestamp                  `json:"updated,omitempty"`              // Момент последнего обновления
	VatEnabled           *bool                       `json:"vatEnabled,omitempty"`           // Учитывается ли НДС
	VatIncluded          *bool                       `json:"vatIncluded,omitempty"`          // Включен ли НДС в цену
	VatSum               *float64                    `json:"vatSum,omitempty"`               // Сумма включая НДС
	Payments             *Payments                   `json:"payments,omitempty"`             // Массив ссылок на связанные операции в формате Метаданных
	PurchaseOrder        *PurchaseOrder              `json:"purchaseOrder,omitempty"`        // Ссылка на связанный заказ поставщику в формате Метаданных
	Supplies             *Supplies                   `json:"supplies,omitempty"`             // Ссылки на связанные приемки в формате Метаданных
}

func (i InvoiceIn) String() string {
	return Stringify(i)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (i InvoiceIn) GetMeta() *Meta {
	return i.Meta
}

func (i InvoiceIn) MetaType() MetaType {
	return MetaTypeInvoiceIn
}

type InvoicesIn = Slice[InvoiceIn]

// InvoiceInPosition Позиция Счета поставщика.
// Ключевое слово: invoiceposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-schet-postawschika-scheta-postawschikow-pozicii-scheta-postawschika
type InvoiceInPosition struct {
	InvoicePosition
}

func (i InvoiceInPosition) String() string {
	return Stringify(i)
}

func (i InvoiceInPosition) MetaType() MetaType {
	return MetaTypeInvoicePosition
}

// InvoiceInTemplateArg
// Документ: Счет поставщика (invoicein)
// Основание, на котором он может быть создан:
// - Заказ поставщику (purchaseorder)
type InvoiceInTemplateArg struct {
	PurchaseOrder *MetaWrapper `json:"purchaseOrder,omitempty"`
}
