package moysklad

import (
	"github.com/google/uuid"
	"time"
)

// ContactPerson Контактное лицо.
//
// Код сущности: contactperson
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-kontragenty-attributy-suschnosti-adres-kontaktnye-lica-kontragentow
type ContactPerson struct {
	AccountID    *uuid.UUID    `json:"accountId,omitempty"`    // ID учётной записи
	Agent        *Counterparty `json:"agent,omitempty"`        // Метаданные контрагента
	Description  *string       `json:"description,omitempty"`  // Описание контактного лица
	Email        *string       `json:"email,omitempty"`        // Адрес электронной почты контактного лица
	ExternalCode *string       `json:"externalCode,omitempty"` // Внешний код контактного лица
	ID           *uuid.UUID    `json:"id,omitempty"`           // ID Контактного лица
	Meta         *Meta         `json:"meta,omitempty"`         // Метаданные Контактного лица Контрагента
	Name         *string       `json:"name,omitempty"`         // ФИО контактного лица
	Phone        *string       `json:"phone,omitempty"`        // Номер телефона контактного лица
	Position     *string       `json:"position,omitempty"`     // Должность контактного лица
	Updated      *Timestamp    `json:"updated,omitempty"`      // Момент последнего обновления
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (contactPerson ContactPerson) Clean() *ContactPerson {
	if contactPerson.Meta == nil {
		return nil
	}
	return &ContactPerson{Meta: contactPerson.Meta}
}

// GetAccountID возвращает ID учётной записи.
func (contactPerson ContactPerson) GetAccountID() uuid.UUID {
	return Deref(contactPerson.AccountID)
}

// GetAgent возвращает Метаданные контрагента.
func (contactPerson ContactPerson) GetAgent() Counterparty {
	return Deref(contactPerson.Agent)
}

// GetDescription возвращает Описание контактного лица.
func (contactPerson ContactPerson) GetDescription() string {
	return Deref(contactPerson.Description)
}

// GetEmail возвращает Адрес электронной почты контактного лица.
func (contactPerson ContactPerson) GetEmail() string {
	return Deref(contactPerson.Email)
}

// GetExternalCode возвращает Внешний код контактного лица.
func (contactPerson ContactPerson) GetExternalCode() string {
	return Deref(contactPerson.ExternalCode)
}

// GetID возвращает ID Контактного лица.
func (contactPerson ContactPerson) GetID() uuid.UUID {
	return Deref(contactPerson.ID)
}

// GetMeta возвращает Метаданные Контактного лица Контрагента.
func (contactPerson ContactPerson) GetMeta() Meta {
	return Deref(contactPerson.Meta)
}

// GetName возвращает ФИО контактного лица.
func (contactPerson ContactPerson) GetName() string {
	return Deref(contactPerson.Name)
}

// GetPhone возвращает Номер телефона контактного лица.
func (contactPerson ContactPerson) GetPhone() string {
	return Deref(contactPerson.Phone)
}

// GetPosition возвращает Должность контактного лица.
func (contactPerson ContactPerson) GetPosition() string {
	return Deref(contactPerson.Position)
}

// GetUpdated возвращает Момент последнего обновления.
func (contactPerson ContactPerson) GetUpdated() time.Time {
	return Deref(contactPerson.Updated).Time()
}

// SetAgent устанавливает Метаданные контрагента.
func (contactPerson *ContactPerson) SetAgent(agent *Counterparty) *ContactPerson {
	if agent != nil {
		contactPerson.Agent = agent.Clean()
	}
	return contactPerson
}

// SetDescription устанавливает Описание контактного лица.
func (contactPerson *ContactPerson) SetDescription(description string) *ContactPerson {
	contactPerson.Description = &description
	return contactPerson
}

// SetEmail устанавливает Адрес электронной почты контактного лица.
func (contactPerson *ContactPerson) SetEmail(email string) *ContactPerson {
	contactPerson.Email = &email
	return contactPerson
}

// SetExternalCode устанавливает Внешний код контактного лица.
func (contactPerson *ContactPerson) SetExternalCode(externalCode string) *ContactPerson {
	contactPerson.ExternalCode = &externalCode
	return contactPerson
}

// SetMeta устанавливает Метаданные Контактного лица Контрагента.
func (contactPerson *ContactPerson) SetMeta(meta *Meta) *ContactPerson {
	contactPerson.Meta = meta
	return contactPerson
}

// SetName устанавливает ФИО контактного лица.
func (contactPerson *ContactPerson) SetName(name string) *ContactPerson {
	contactPerson.Name = &name
	return contactPerson
}

// SetPhone устанавливает Номер телефона контактного лица.
func (contactPerson *ContactPerson) SetPhone(phone string) *ContactPerson {
	contactPerson.Phone = &phone
	return contactPerson
}

// SetPosition устанавливает Должность контактного лица.
func (contactPerson *ContactPerson) SetPosition(position string) *ContactPerson {
	contactPerson.Position = &position
	return contactPerson
}

// String реализует интерфейс [fmt.Stringer].
func (contactPerson ContactPerson) String() string {
	return Stringify(contactPerson)
}

// MetaType возвращает код сущности.
func (ContactPerson) MetaType() MetaType {
	return MetaTypeContactPerson
}
