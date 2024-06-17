package moysklad

import "github.com/google/uuid"

// Payroll Начисление зарплаты. TODO
// Ключевое слово: payroll
type Payroll struct {
	Meta         *Meta            `json:"meta,omitempty"`
	ID           *uuid.UUID       `json:"id,omitempty"`
	AccountID    *uuid.UUID       `json:"accountId,omitempty"`
	Owner        *Employee        `json:"owner,omitempty"`
	Shared       *bool            `json:"shared,omitempty"`
	Group        *Group           `json:"group,omitempty"`
	Updated      *Timestamp       `json:"updated,omitempty"`
	Name         *string          `json:"name,omitempty"`
	ExternalCode *string          `json:"externalCode,omitempty"`
	Moment       *Timestamp       `json:"moment,omitempty"`
	Applicable   *bool            `json:"applicable,omitempty"`
	Sum          *float64         `json:"sum,omitempty"`
	Organization *Organization    `json:"organization,omitempty"`
	Created      *Timestamp       `json:"created,omitempty"`
	Printed      *bool            `json:"printed,omitempty"`
	Published    *bool            `json:"published,omitempty"`
	Files        *MetaArray[File] `json:"files,omitempty"`
	State        *State           `json:"state,omitempty"`
}

// AsTaskOperation реализует интерфейс AsTaskOperationInterface
func (payroll Payroll) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: payroll.Meta}
}

func (payroll Payroll) GetMeta() Meta {
	return Deref(payroll.Meta)
}

func (payroll Payroll) GetID() uuid.UUID {
	return Deref(payroll.ID)
}

func (payroll Payroll) GetAccountID() uuid.UUID {
	return Deref(payroll.AccountID)
}

func (payroll Payroll) GetOwner() Employee {
	return Deref(payroll.Owner)
}

func (payroll Payroll) GetShared() bool {
	return Deref(payroll.Shared)
}

func (payroll Payroll) GetGroup() Group {
	return Deref(payroll.Group)
}

func (payroll Payroll) GetUpdated() Timestamp {
	return Deref(payroll.Updated)
}

func (payroll Payroll) GetName() string {
	return Deref(payroll.Name)
}

func (payroll Payroll) GetExternalCode() string {
	return Deref(payroll.ExternalCode)
}

func (payroll Payroll) GetMoment() Timestamp {
	return Deref(payroll.Moment)
}

func (payroll Payroll) GetApplicable() bool {
	return Deref(payroll.Applicable)
}

func (payroll Payroll) GetSum() float64 {
	return Deref(payroll.Sum)
}

func (payroll Payroll) GetOrganization() Organization {
	return Deref(payroll.Organization)
}

func (payroll Payroll) GetCreated() Timestamp {
	return Deref(payroll.Created)
}

func (payroll Payroll) GetPrinted() bool {
	return Deref(payroll.Printed)
}

func (payroll Payroll) GetPublished() bool {
	return Deref(payroll.Published)
}

func (payroll Payroll) GetFiles() MetaArray[File] {
	return Deref(payroll.Files)
}

func (payroll Payroll) GetState() State {
	return Deref(payroll.State)
}

func (payroll Payroll) String() string {
	return Stringify(payroll)
}

func (payroll Payroll) MetaType() MetaType {
	return MetaTypePayroll
}
