package moysklad

import (
	"github.com/google/uuid"
)

// ContactPerson Контактное лицо
// Ключевое слово: contactperson
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-kontragenty-attributy-suschnosti-adres-kontaktnye-lica-kontragentow
type ContactPerson struct {
	AccountID    *uuid.UUID    `json:"accountId,omitempty"`    // ID учетной записи
	Agent        *Counterparty `json:"agent,omitempty"`        // Метаданные контрагента
	Description  *string       `json:"description,omitempty"`  // Описание контактного лица
	Email        *string       `json:"email,omitempty"`        // Адрес электронной почты контактного лица
	ExternalCode *string       `json:"externalCode,omitempty"` // Внешний код контактного лица
	ID           *uuid.UUID    `json:"id,omitempty"`           // ID сущности
	Meta         *Meta         `json:"meta,omitempty"`         // Метаданные
	Name         *string       `json:"name,omitempty"`         // Наименование
	Phone        *string       `json:"phone,omitempty"`        // Номер телефона контактного лица
	Position     *string       `json:"position,omitempty"`     // Должность контактного лица
	Updated      *Timestamp    `json:"updated,omitempty"`      // Момент последнего обновления
}

// Clean возвращает сущность с единственным заполненным полем Meta
func (contactPerson ContactPerson) Clean() *ContactPerson {
	return &ContactPerson{Meta: contactPerson.Meta}
}

func (contactPerson ContactPerson) GetAccountID() uuid.UUID {
	return Deref(contactPerson.AccountID)
}

func (contactPerson ContactPerson) GetAgent() Counterparty {
	return Deref(contactPerson.Agent)
}

func (contactPerson ContactPerson) GetDescription() string {
	return Deref(contactPerson.Description)
}

func (contactPerson ContactPerson) GetEmail() string {
	return Deref(contactPerson.Email)
}

func (contactPerson ContactPerson) GetExternalCode() string {
	return Deref(contactPerson.ExternalCode)
}

func (contactPerson ContactPerson) GetID() uuid.UUID {
	return Deref(contactPerson.ID)
}

func (contactPerson ContactPerson) GetMeta() Meta {
	return Deref(contactPerson.Meta)
}

func (contactPerson ContactPerson) GetName() string {
	return Deref(contactPerson.Name)
}

func (contactPerson ContactPerson) GetPhone() string {
	return Deref(contactPerson.Phone)
}

func (contactPerson ContactPerson) GetPosition() string {
	return Deref(contactPerson.Position)
}

func (contactPerson ContactPerson) GetUpdated() Timestamp {
	return Deref(contactPerson.Updated)
}

func (contactPerson *ContactPerson) SetAgent(agent *Counterparty) *ContactPerson {
	contactPerson.Agent = agent.Clean()
	return contactPerson
}

func (contactPerson *ContactPerson) SetDescription(description string) *ContactPerson {
	contactPerson.Description = &description
	return contactPerson
}

func (contactPerson *ContactPerson) SetEmail(email string) *ContactPerson {
	contactPerson.Email = &email
	return contactPerson
}

func (contactPerson *ContactPerson) SetExternalCode(externalCode string) *ContactPerson {
	contactPerson.ExternalCode = &externalCode
	return contactPerson
}

func (contactPerson *ContactPerson) SetMeta(meta *Meta) *ContactPerson {
	contactPerson.Meta = meta
	return contactPerson
}

func (contactPerson *ContactPerson) SetName(name string) *ContactPerson {
	contactPerson.Name = &name
	return contactPerson
}

func (contactPerson *ContactPerson) SetPhone(phone string) *ContactPerson {
	contactPerson.Phone = &phone
	return contactPerson
}

func (contactPerson *ContactPerson) SetPosition(position string) *ContactPerson {
	contactPerson.Position = &position
	return contactPerson
}

func (contactPerson ContactPerson) String() string {
	return Stringify(contactPerson)
}

// MetaType возвращает тип сущности.
func (ContactPerson) MetaType() MetaType {
	return MetaTypeContactPerson
}
