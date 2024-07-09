package moysklad

import (
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"reflect"
)

// Payment общие поля для входящих/исходящих платежей и приходных/расходных ордеров.
type Payment struct {
	AccountID      *uuid.UUID               `json:"accountId,omitempty"`      // ID учётной записи
	Agent          *Counterparty            `json:"agent,omitempty"`          // Метаданные контрагента
	Applicable     *bool                    `json:"applicable,omitempty"`     // Отметка о проведении
	Attributes     Slice[Attribute]         `json:"attributes,omitempty"`     // Коллекция доп. полей
	Code           *string                  `json:"code,omitempty"`           // Код платежа
	Contract       *NullValue[Contract]     `json:"contract,omitempty"`       // Метаданные платежа
	Created        *Timestamp               `json:"created,omitempty"`        // Дата платежа
	Deleted        *Timestamp               `json:"deleted,omitempty"`        // Момент последнего удаления платежа
	Description    *string                  `json:"description,omitempty"`    // Комментарий платежа
	ExternalCode   *string                  `json:"externalCode,omitempty"`   // Внешний код платежа
	Files          *MetaArray[File]         `json:"files,omitempty"`          // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group          *Group                   `json:"group,omitempty"`          // Отдел сотрудника
	ID             *uuid.UUID               `json:"id,omitempty"`             // ID платежа
	Meta           Meta                     `json:"meta"`                     // Метаданные платежа
	Moment         *Timestamp               `json:"moment,omitempty"`         // Дата платежа
	Name           *string                  `json:"name,omitempty"`           // Наименование платежа
	Organization   *Organization            `json:"organization,omitempty"`   // Метаданные юрлица
	Owner          *Employee                `json:"owner,omitempty"`          // Метаданные владельца (Сотрудника)
	PaymentPurpose *string                  `json:"paymentPurpose,omitempty"` // Назначение платежа
	Printed        *bool                    `json:"printed,omitempty"`        // Напечатан ли документ
	Project        *NullValue[Project]      `json:"project,omitempty"`        // Метаданные проекта
	Published      *bool                    `json:"published,omitempty"`      // Опубликован ли документ
	Rate           *NullValue[Rate]         `json:"rate,omitempty"`           // Валюта
	SalesChannel   *NullValue[SalesChannel] `json:"salesChannel,omitempty"`   // Метаданные канала продаж
	Shared         *bool                    `json:"shared,omitempty"`         // Общий доступ
	State          *State                   `json:"state,omitempty"`          // Метаданные статуса платежа
	Sum            *float64                 `json:"sum,omitempty"`            // Сумма платежа в копейках
	SyncID         *uuid.UUID               `json:"syncId,omitempty"`         // ID синхронизации
	Updated        *Timestamp               `json:"updated,omitempty"`        // Момент последнего обновления платежа
	VatSum         *float64                 `json:"vatSum,omitempty"`         // Сумма НДС
	LinkedSum      *float64                 `json:"linkedSum,omitempty"`      // Сумма, оплаченная по документу из этого платежа
	Operations     Operations               `json:"operations,omitempty"`     // Массив ссылок на связанные операции в формате Метаданных
	raw            []byte                   // сырые данные для последующей конвертации в нужный тип
}

// PaymentConverter описывает метод, возвращающий [Payment].
type PaymentConverter interface {
	AsPayment() *Payment
}

// GetAccountID возвращает ID учётной записи.
func (payment Payment) GetAccountID() uuid.UUID {
	return Deref(payment.AccountID)
}

// GetAgent возвращает Метаданные Контрагента.
func (payment Payment) GetAgent() Counterparty {
	return Deref(payment.Agent)
}

// GetApplicable возвращает Отметку о проведении.
func (payment Payment) GetApplicable() bool {
	return Deref(payment.Applicable)
}

// GetAttributes возвращает Список метаданных доп. полей.
func (payment Payment) GetAttributes() Slice[Attribute] {
	return payment.Attributes
}

// GetCode возвращает Код платежа.
func (payment Payment) GetCode() string {
	return Deref(payment.Code)
}

// GetContract возвращает Метаданные договора.
func (payment Payment) GetContract() Contract {
	return Deref(payment.Contract).getValue()
}

// GetCreated возвращает Дату создания.
func (payment Payment) GetCreated() Timestamp {
	return Deref(payment.Created)
}

// GetDeleted возвращает Момент последнего удаления платежа.
func (payment Payment) GetDeleted() Timestamp {
	return Deref(payment.Deleted)
}

// GetDescription возвращает Комментарий платежа.
func (payment Payment) GetDescription() string {
	return Deref(payment.Description)
}

// GetExternalCode возвращает Внешний код платежа.
func (payment Payment) GetExternalCode() string {
	return Deref(payment.ExternalCode)
}

// GetFiles возвращает Метаданные массива Файлов.
func (payment Payment) GetFiles() MetaArray[File] {
	return Deref(payment.Files)
}

// GetGroup возвращает Отдел сотрудника.
func (payment Payment) GetGroup() Group {
	return Deref(payment.Group)
}

// GetID возвращает ID платежа.
func (payment Payment) GetID() uuid.UUID {
	return Deref(payment.ID)
}

// GetMeta возвращает Метаданные платежа.
func (payment Payment) GetMeta() Meta {
	return payment.Meta
}

// GetMoment возвращает Дату платежа.
func (payment Payment) GetMoment() Timestamp {
	return Deref(payment.Moment)
}

// GetName возвращает Наименование платежа.
func (payment Payment) GetName() string {
	return Deref(payment.Name)
}

// GetOrganization возвращает Метаданные юрлица.
func (payment Payment) GetOrganization() Organization {
	return Deref(payment.Organization)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (payment Payment) GetOwner() Employee {
	return Deref(payment.Owner)
}

// GetPaymentPurpose возвращает Назначение платежа.
func (payment Payment) GetPaymentPurpose() string {
	return Deref(payment.PaymentPurpose)
}

// GetPrinted возвращает true, если документ напечатан.
func (payment Payment) GetPrinted() bool {
	return Deref(payment.Printed)
}

// GetProject возвращает Метаданные проекта.
func (payment Payment) GetProject() Project {
	return Deref(payment.Project).getValue()
}

// GetPublished возвращает true, если документ опубликован.
func (payment Payment) GetPublished() bool {
	return Deref(payment.Published)
}

// GetRate возвращает Валюту.
func (payment Payment) GetRate() Rate {
	return Deref(payment.Rate).getValue()
}

// GetSalesChannel возвращает Метаданные канала продаж.
func (payment Payment) GetSalesChannel() SalesChannel {
	return Deref(payment.SalesChannel).getValue()
}

// GetShared возвращает флаг Общего доступа.
func (payment Payment) GetShared() bool {
	return Deref(payment.Shared)
}

// GetState возвращает Метаданные статуса платежа.
func (payment Payment) GetState() State {
	return Deref(payment.State)
}

// GetSum возвращает Сумму платежа в копейках.
func (payment Payment) GetSum() float64 {
	return Deref(payment.Sum)
}

// GetSyncID возвращает ID синхронизации.
func (payment Payment) GetSyncID() uuid.UUID {
	return Deref(payment.SyncID)
}

// GetUpdated возвращает Момент последнего обновления платежа.
func (payment Payment) GetUpdated() Timestamp {
	return Deref(payment.Updated)
}

// GetVatSum возвращает Сумму НДС.
func (payment Payment) GetVatSum() float64 {
	return Deref(payment.VatSum)
}

// GetLinkedSum возвращает Сумму, оплаченную по документу из этого платежа.
func (payment Payment) GetLinkedSum() float64 {
	return Deref(payment.LinkedSum)
}

// GetOperations возвращает Метаданные связанных операций.
func (payment Payment) GetOperations() Operations {
	return payment.Operations
}

// MetaType возвращает код сущности.
func (payment Payment) MetaType() MetaType {
	return payment.Meta.GetType()
}

// Raw реализует интерфейс [RawMetaTyper].
func (payment Payment) Raw() []byte {
	return payment.raw
}

// String реализует интерфейс [fmt.Stringer].
func (payment Payment) String() string {
	return Stringify(payment.Meta)
}

// UnmarshalJSON реализует интерфейс [json.Unmarshaler].
func (payment *Payment) UnmarshalJSON(data []byte) (err error) {
	type alias Payment
	var t alias

	if err = json.Unmarshal(data, &t); err != nil {
		return err
	}

	t.raw = data
	*payment = Payment(t)
	return nil
}

// IsCashIn возвращает true, если объект имеет код сущности [MetaTypeCashIn].
func (payment Payment) IsCashIn() bool {
	return payment.Meta.GetType() == MetaTypeCashIn
}

// IsCashOut возвращает true, если объект имеет код сущности [MetaTypeCashOut].
func (payment Payment) IsCashOut() bool {
	return payment.Meta.GetType() == MetaTypeCashOut
}

// IsPaymentIn возвращает true, если объект имеет код сущности [MetaTypePaymentIn].
func (payment Payment) IsPaymentIn() bool {
	return payment.Meta.GetType() == MetaTypePaymentIn
}

// IsPaymentOut возвращает true, если объект имеет код сущности [MetaTypePaymentOut].
func (payment Payment) IsPaymentOut() bool {
	return payment.Meta.GetType() == MetaTypePaymentOut
}

// AsCashIn пытается привести объект к типу [CashIn].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [CashIn] или nil в случае неудачи.
func (payment *Payment) AsCashIn() *CashIn {
	return UnmarshalAsType[CashIn](payment)
}

// AsCashOut пытается привести объект к типу [CashOut].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [CashOut] или nil в случае неудачи.
func (payment *Payment) AsCashOut() *CashOut {
	return UnmarshalAsType[CashOut](payment)
}

// AsPaymentIn пытается привести объект к типу [PaymentIn].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [PaymentIn] или nil в случае неудачи.
func (payment *Payment) AsPaymentIn() *PaymentIn {
	return UnmarshalAsType[PaymentIn](payment)
}

// AsPaymentOut пытается привести объект к типу [PaymentOut].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [PaymentOut] или nil в случае неудачи.
func (payment *Payment) AsPaymentOut() *PaymentOut {
	return UnmarshalAsType[PaymentOut](payment)
}

// NewPaymentsFrom преобразует список объектов, реализующих интерфейс [PaymentConverter] в список платежей [Payment].
func NewPaymentsFrom[T PaymentConverter](elements []T) Slice[Payment] {
	payments := make(Slice[Payment], 0, len(elements))
	for _, payment := range elements {
		if reflect.ValueOf(payment).Kind() == reflect.Ptr {
			payments.Push(payment.AsPayment())
		}
	}
	return payments
}
