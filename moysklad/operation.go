package moysklad

import (
	"encoding/json"
	"github.com/google/uuid"
)

// Operation операция, к которой привязан платёж
// Представляет из себя структуру из полей:
// `Meta` для определения типа сущности
// `LinkedSum` для хранения суммы по операции
// `data` для хранения сырых данных
type Operation struct {
	// Общие поля
	AccountID    *uuid.UUID    `json:"accountId,omitempty"`    // ID учетной записи
	Attributes   Attributes    `json:"attributes,omitempty"`   // Коллекция метаданных доп. полей. Поля объекта
	Created      *Timestamp    `json:"created,omitempty"`      // Дата создания
	Deleted      *Timestamp    `json:"deleted,omitempty"`      // Момент последнего удаления
	Description  *string       `json:"description,omitempty"`  // Комментарий
	ExternalCode *string       `json:"externalCode,omitempty"` // Внешний код
	Files        *Files        `json:"files,omitempty"`        // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group        *Group        `json:"group,omitempty"`        // Отдел сотрудника
	ID           *uuid.UUID    `json:"id,omitempty"`           // ID сущности
	Meta         Meta          `json:"meta,omitempty"`         // Метаданные
	Name         *string       `json:"name,omitempty"`         // Наименование
	Organization *Organization `json:"organization,omitempty"` // Метаданные юрлица
	Owner        *Employee     `json:"owner,omitempty"`        // Владелец (Сотрудник)
	Printed      *bool         `json:"printed,omitempty"`      // Напечатан ли документ
	Published    *bool         `json:"published,omitempty"`    // Опубликован ли документ
	Shared       *bool         `json:"shared,omitempty"`       // Общий доступ
	SyncID       *uuid.UUID    `json:"syncId,omitempty"`       // ID синхронизации. После заполнения недоступен для изменения
	Updated      *Timestamp    `json:"updated,omitempty"`      // Момент последнего обновления
	VatEnabled   *bool         `json:"vatEnabled,omitempty"`   // Учитывается ли НДС
	VatIncluded  *bool         `json:"vatIncluded,omitempty"`  // Включен ли НДС в цену
	LinkedSum    *Decimal      `json:"linkedSum,omitempty"`    // Сумма, оплаченная по данному документу
	Payments     *Payments     `json:"payments,omitempty"`     // Массив ссылок на связанные платежи в формате Метаданных

	// сырые данные
	data json.RawMessage
}

func NewOperation(metaOwner MetaOwner) *Operation {
	return &Operation{Meta: metaOwner.GetMeta()}
}

func (operation Operation) String() string {
	return Stringify(operation.Meta)
}

// MetaType удовлетворяет интерфейсу MetaTyper
func (operation Operation) MetaType() MetaType {
	return operation.Meta.Type
}

// Raw удовлетворяет интерфейсу RawMetaTyper
func (operation Operation) Raw() json.RawMessage {
	return operation.data
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (operation *Operation) UnmarshalJSON(data []byte) error {
	type alias Operation
	var t alias

	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	t.data = data

	*operation = Operation(t)
	return nil
}

func (operation Operation) AsCustomerOrder() *CustomerOrder {
	return unmarshalAsType[CustomerOrder](operation)
}

func (operation Operation) AsPurchaseReturn() *PurchaseReturn {
	return unmarshalAsType[PurchaseReturn](operation)
}

func (operation Operation) AsDemand() *Demand {
	return unmarshalAsType[Demand](operation)
}

func (operation Operation) AsInvoiceOut() *InvoiceOut {
	return unmarshalAsType[InvoiceOut](operation)
}

func (operation Operation) AsRetailShift() *RetailShift {
	return unmarshalAsType[RetailShift](operation)
}

func (operation Operation) AsCommissionReportIn() *CommissionReportIn {
	return unmarshalAsType[CommissionReportIn](operation)
}

func (operation Operation) AsSalesReturn() *SalesReturn {
	return unmarshalAsType[SalesReturn](operation)
}

func (operation Operation) AsSupply() *Supply {
	return unmarshalAsType[Supply](operation)
}

func (operation Operation) AsInvoiceIn() *InvoiceIn {
	return unmarshalAsType[InvoiceIn](operation)
}

func (operation Operation) AsPurchaseOrder() *PurchaseOrder {
	return unmarshalAsType[PurchaseOrder](operation)
}

func (operation Operation) AsCommissionReportOut() *CommissionReportOut {
	return unmarshalAsType[CommissionReportOut](operation)
}

type Operations Slice[Operation]

// Push Привязка платежей к документам.
// Необходимо передать *Operation, которые были созданы через NewOperation.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-obschie-swedeniq-priwqzka-platezhej-k-dokumentam
func (operations *Operations) Push(elements ...*Operation) *Operations {
	*operations = append(*operations, elements...)
	return operations
}

func (operations Operations) FilterCustomerOrder() Slice[CustomerOrder] {
	return filterType[CustomerOrder](operations)
}

func (operations Operations) FilterPurchaseReturn() Slice[PurchaseReturn] {
	return filterType[PurchaseReturn](operations)
}

func (operations Operations) FilterDemand() Slice[Demand] {
	return filterType[Demand](operations)
}

func (operations Operations) FilterInvoiceOut() Slice[InvoiceOut] {
	return filterType[InvoiceOut](operations)
}

func (operations Operations) FilterCommissionReportIn() Slice[CommissionReportIn] {
	return filterType[CommissionReportIn](operations)
}

func (operations Operations) FilterSalesReturn() Slice[SalesReturn] {
	return filterType[SalesReturn](operations)
}

func (operations Operations) FilterSupply() Slice[Supply] {
	return filterType[Supply](operations)
}

func (operations Operations) FilterInvoiceIn() Slice[InvoiceIn] {
	return filterType[InvoiceIn](operations)
}

func (operations Operations) FilterPurchaseOrder() Slice[PurchaseOrder] {
	return filterType[PurchaseOrder](operations)
}

func (operations Operations) FilterCommissionReportOut() Slice[CommissionReportOut] {
	return filterType[CommissionReportOut](operations)
}

func (operations Operations) FilterRetailShift() Slice[RetailShift] {
	return filterType[RetailShift](operations)
}

//
//	func (o Operations) FilterCounterparty() Slice[Counterparty] {
//		return filterType[Counterparty](o)
//	}
//
//	func (o Operations) FilterOrganization() Slice[Organization] {
//		return filterType[Organization](o)
//	}
