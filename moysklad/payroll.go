package moysklad

import (
	"time"
)

// Payroll Начисление зарплаты.
//
// Код сущности: payroll
type Payroll struct {
	Meta         *Meta            `json:"meta,omitempty"`         // Метаданные Начисления зарплаты
	ID           *string          `json:"id,omitempty"`           // ID Начисления зарплаты
	AccountID    *string          `json:"accountId,omitempty"`    // ID учётной записи
	Owner        *Employee        `json:"owner,omitempty"`        // Метаданные владельца (Сотрудника)
	Shared       *bool            `json:"shared,omitempty"`       // Общий доступ
	Group        *Group           `json:"group,omitempty"`        // Отдел сотрудника
	Updated      *Timestamp       `json:"updated,omitempty"`      // Момент последнего обновления Начисления зарплаты
	Name         *string          `json:"name,omitempty"`         // Наименование Начисления зарплаты
	ExternalCode *string          `json:"externalCode,omitempty"` // Внешний код Начисления зарплаты
	Moment       *Timestamp       `json:"moment,omitempty"`       // Дата документа
	Applicable   *bool            `json:"applicable,omitempty"`   // Отметка о проведении
	Sum          *float64         `json:"sum,omitempty"`          // Сумма в копейках
	Organization *Organization    `json:"organization,omitempty"` // Метаданные юрлица
	Created      *Timestamp       `json:"created,omitempty"`      // Момент создания
	Printed      *bool            `json:"printed,omitempty"`      // Напечатан ли документ
	Published    *bool            `json:"published,omitempty"`    // Опубликован ли документ
	Files        *MetaArray[File] `json:"files,omitempty"`        // Метаданные массива Файлов (Максимальное количество файлов - 100)
	State        *State           `json:"state,omitempty"`        // Метаданные статуса Начисления зарплаты
}

// AsTaskOperation реализует интерфейс [TaskOperationConverter].
func (payroll Payroll) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: payroll.Meta}
}

// GetMeta возвращает Метаданные Начисления зарплаты.
func (payroll Payroll) GetMeta() Meta {
	return Deref(payroll.Meta)
}

// GetID возвращает ID Начисления зарплаты.
func (payroll Payroll) GetID() string {
	return Deref(payroll.ID)
}

// GetAccountID возвращает ID учётной записи.
func (payroll Payroll) GetAccountID() string {
	return Deref(payroll.AccountID)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (payroll Payroll) GetOwner() Employee {
	return Deref(payroll.Owner)
}

// GetShared возвращает флаг Общего доступа.
func (payroll Payroll) GetShared() bool {
	return Deref(payroll.Shared)
}

// GetGroup возвращает Отдел сотрудника.
func (payroll Payroll) GetGroup() Group {
	return Deref(payroll.Group)
}

// GetUpdated возвращает Момент последнего обновления Начисления зарплаты.
func (payroll Payroll) GetUpdated() time.Time {
	return Deref(payroll.Updated).Time()
}

// GetName возвращает Наименование Начисления зарплаты.
func (payroll Payroll) GetName() string {
	return Deref(payroll.Name)
}

// GetExternalCode возвращает Внешний код Начисления зарплаты.
func (payroll Payroll) GetExternalCode() string {
	return Deref(payroll.ExternalCode)
}

// GetMoment возвращает Дату документа.
func (payroll Payroll) GetMoment() time.Time {
	return Deref(payroll.Moment).Time()
}

// GetApplicable возвращает Отметку о проведении.
func (payroll Payroll) GetApplicable() bool {
	return Deref(payroll.Applicable)
}

// GetSum возвращает Сумму в копейках.
func (payroll Payroll) GetSum() float64 {
	return Deref(payroll.Sum)
}

// GetOrganization возвращает Метаданные юрлица.
func (payroll Payroll) GetOrganization() Organization {
	return Deref(payroll.Organization)
}

// GetCreated возвращает Дату создания.
func (payroll Payroll) GetCreated() time.Time {
	return Deref(payroll.Created).Time()
}

// GetPrinted возвращает true, если документ напечатан.
func (payroll Payroll) GetPrinted() bool {
	return Deref(payroll.Printed)
}

// GetPublished возвращает true, если документ опубликован.
func (payroll Payroll) GetPublished() bool {
	return Deref(payroll.Published)
}

// GetFiles возвращает Метаданные массива Файлов.
func (payroll Payroll) GetFiles() MetaArray[File] {
	return Deref(payroll.Files)
}

// GetState возвращает Метаданные статуса Начисления зарплаты.
func (payroll Payroll) GetState() State {
	return Deref(payroll.State)
}

// String реализует интерфейс [fmt.Stringer].
func (payroll Payroll) String() string {
	return Stringify(payroll)
}

// MetaType возвращает код сущности.
func (Payroll) MetaType() MetaType {
	return MetaTypePayroll
}
