package moysklad

import (
	"encoding/json"
	"github.com/google/uuid"
)

type Payment struct {
	// Общие поля
	AccountID      *uuid.UUID    `json:"accountId,omitempty"`      // ID учетной записи
	Agent          *Counterparty `json:"agent,omitempty"`          // Метаданные контрагента
	Applicable     *bool         `json:"applicable,omitempty"`     // Отметка о проведении
	Attributes     Attributes    `json:"attributes,omitempty"`     // Коллекция метаданных доп. полей. Поля объекта
	Code           *string       `json:"code,omitempty"`           // Код выданного
	Contract       *Contract     `json:"contract,omitempty"`       // Метаданные договора
	Created        *Timestamp    `json:"created,omitempty"`        // Дата создания
	Deleted        *Timestamp    `json:"deleted,omitempty"`        // Момент последнего удаления
	Description    *string       `json:"description,omitempty"`    // Комментарий
	ExternalCode   *string       `json:"externalCode,omitempty"`   // Внешний код
	Files          *Files        `json:"files,omitempty"`          // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group          *Group        `json:"group,omitempty"`          // Отдел сотрудника
	ID             *uuid.UUID    `json:"id,omitempty"`             // ID сущности
	Meta           Meta          `json:"meta"`                     // Метаданные
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
	Sum            *Decimal      `json:"sum,omitempty"`            // Сумма
	SyncID         *uuid.UUID    `json:"syncId,omitempty"`         // ID синхронизации. После заполнения недоступен для изменения
	Updated        *Timestamp    `json:"updated,omitempty"`        // Момент последнего обновления
	VatSum         *Decimal      `json:"vatSum,omitempty"`         // Сумма включая НДС
	LinkedSum      *Decimal      `json:"linkedSum,omitempty"`      // Сумма, оплаченная по документу из этого платежа
	Operations     *Operations   `json:"operations,omitempty"`     // Массив ссылок на связанные операции в формате Метаданных

	// сырые данные
	data json.RawMessage
}

// MetaType удовлетворяет интерфейсу MetaTyper
func (p Payment) MetaType() MetaType {
	return p.Meta.Type
}

// Data удовлетворяет интерфейсу DataMetaTyper
func (p Payment) Data() json.RawMessage {
	return p.data
}

func (p Payment) String() string {
	return Stringify(p.Meta)
}

// UnmarshalJSON анмаршалит Входящий платеж, Приходный ордер, при expand=payments
func (p *Payment) UnmarshalJSON(data []byte) (err error) {
	type alias Payment
	var t alias

	if err = json.Unmarshal(data, &t); err != nil {
		return err
	}

	t.data = data
	*p = Payment(t)
	return nil
}

// ConvertToCashIn структурирует позицию в *CashIn
func (p *Payment) ConvertToCashIn() (*CashIn, error) {
	return unmarshalTo[CashIn](p)
}

// ConvertToCashOut структурирует позицию в *CashOut
func (p *Payment) ConvertToCashOut() (*CashOut, error) {
	return unmarshalTo[CashOut](p)
}

// ConvertToPaymentIn структурирует позицию в *PaymentIn
func (p *Payment) ConvertToPaymentIn() (*PaymentIn, error) {
	return unmarshalTo[PaymentIn](p)
}

// ConvertToPaymentOut структурирует позицию в *PaymentOut
func (p *Payment) ConvertToPaymentOut() (*PaymentOut, error) {
	return unmarshalTo[PaymentOut](p)
}

// Payments Входящий платеж, Приходный ордер, Исходящий платеж, Расходный ордер
type Payments []Payment
