package moysklad

import (
	"github.com/google/uuid"
)

// ContactPerson Контактное лицо
// Ключевое слово: contactperson
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-kontragenty-attributy-suschnosti-adres-kontaktnye-lica-kontragentow
type ContactPerson struct {
	AccountId    *uuid.UUID    `json:"accountId,omitempty"`    // ID учетной записи
	Agent        *Counterparty `json:"agent,omitempty"`        // Метаданные контрагента
	Description  *string       `json:"description,omitempty"`  // Описание контактного лица
	Email        *string       `json:"email,omitempty"`        // Адрес электронной почты контактного лица
	ExternalCode *string       `json:"externalCode,omitempty"` // Внешний код контактного лица
	Id           *uuid.UUID    `json:"id,omitempty"`           // ID сущности
	Meta         *Meta         `json:"meta,omitempty"`         // Метаданные
	Name         *string       `json:"name,omitempty"`         // Наименование
	Phone        *string       `json:"phone,omitempty"`        // Номер телефона контактного лица
	Position     *string       `json:"position,omitempty"`     // Должность контактного лица
	Updated      *Timestamp    `json:"updated,omitempty"`      // Момент последнего обновления
}

func (c ContactPerson) String() string {
	return Stringify(c)
}

func (c ContactPerson) MetaType() MetaType {
	return MetaTypeContactPerson
}
