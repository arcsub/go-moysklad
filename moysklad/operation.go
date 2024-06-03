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

func (o Operation) String() string {
	return Stringify(o.Meta)
}

// MetaType удовлетворяет интерфейсу MetaTyper
func (o Operation) MetaType() MetaType {
	return o.Meta.Type
}

// Raw удовлетворяет интерфейсу RawMetaTyper
func (o Operation) Raw() json.RawMessage {
	return o.data
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (o *Operation) UnmarshalJSON(data []byte) error {
	type alias Operation
	var t alias

	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	t.data = data

	*o = Operation(t)
	return nil
}

func (o Operation) CustomerOrder() *CustomerOrder {
	return unmarshalAsType[CustomerOrder](o)
}

func (o Operation) PurchaseReturn() *PurchaseReturn {
	return unmarshalAsType[PurchaseReturn](o)
}

func (o Operation) Demand() *Demand {
	return unmarshalAsType[Demand](o)
}

func (o Operation) InvoiceOut() *InvoiceOut {
	return unmarshalAsType[InvoiceOut](o)
}

func (o Operation) RetailShift() *RetailShift {
	return unmarshalAsType[RetailShift](o)
}

func (o Operation) CommissionReportIn() *CommissionReportIn {
	return unmarshalAsType[CommissionReportIn](o)
}

func (o Operation) SalesReturn() *SalesReturn {
	return unmarshalAsType[SalesReturn](o)
}

func (o Operation) Supply() *Supply {
	return unmarshalAsType[Supply](o)
}

func (o Operation) InvoiceIn() *InvoiceIn {
	return unmarshalAsType[InvoiceIn](o)
}

func (o Operation) PurchaseOrder() *PurchaseOrder {
	return unmarshalAsType[PurchaseOrder](o)
}

func (o Operation) CommissionReportOut() *CommissionReportOut {
	return unmarshalAsType[CommissionReportOut](o)
}

type Operations []Operation

func (o Operations) FilterCustomerOrder() Slice[CustomerOrder] {
	return filterType[CustomerOrder](o)
}

func (o Operations) FilterPurchaseReturn() Slice[PurchaseReturn] {
	return filterType[PurchaseReturn](o)
}

func (o Operations) FilterDemand() Slice[Demand] {
	return filterType[Demand](o)
}

func (o Operations) FilterInvoiceOut() Slice[InvoiceOut] {
	return filterType[InvoiceOut](o)
}

func (o Operations) FilterCommissionReportIn() Slice[CommissionReportIn] {
	return filterType[CommissionReportIn](o)
}

func (o Operations) FilterSalesReturn() Slice[SalesReturn] {
	return filterType[SalesReturn](o)
}

func (o Operations) FilterSupply() Slice[Supply] {
	return filterType[Supply](o)
}

func (o Operations) FilterInvoiceIn() Slice[InvoiceIn] {
	return filterType[InvoiceIn](o)
}

func (o Operations) FilterPurchaseOrder() Slice[PurchaseOrder] {
	return filterType[PurchaseOrder](o)
}

func (o Operations) FilterCommissionReportOut() Slice[CommissionReportOut] {
	return filterType[CommissionReportOut](o)
}

func (o Operations) FilterRetailShift() Slice[RetailShift] {
	return filterType[RetailShift](o)
}

//
//	func (o Operations) FilterCounterparty() Slice[Counterparty] {
//		return filterType[Counterparty](o)
//	}
//
//	func (o Operations) FilterOrganization() Slice[Organization] {
//		return filterType[Organization](o)
//	}
