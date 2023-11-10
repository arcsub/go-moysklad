package moysklad

import (
	"github.com/google/uuid"
)

// Employee Сотрудник.
// Ключевое слово: employee
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sotrudnik
type Employee struct {
	AccountID    *uuid.UUID  `json:"accountId,omitempty"`    // ID учетной записи
	Archived     *bool       `json:"archived,omitempty"`     // Добавлен ли Сотрудник в архив
	Attributes   *Attributes `json:"attributes,omitempty"`   // Дополнительные поля Сотрудника
	Cashiers     *Cashiers   `json:"cashiers,omitempty"`     // Массив кассиров
	Code         *string     `json:"code,omitempty"`         // Код Сотрудника
	Created      *Timestamp  `json:"created,omitempty"`      // Момент создания Сотрудника
	Description  *string     `json:"description,omitempty"`  // Комментарий к Сотруднику
	Email        *string     `json:"email,omitempty"`        // Электронная почта сотрудника
	ExternalCode *string     `json:"externalCode,omitempty"` // Внешний код Сотрудника
	FirstName    *string     `json:"firstName,omitempty"`    // Имя
	FullName     *string     `json:"fullName,omitempty"`     // Имя Отчество Фамилия
	Group        *Group      `json:"group,omitempty"`        // Отдел сотрудника
	ID           *uuid.UUID  `json:"id,omitempty"`           // ID сущности
	Image        *Image      `json:"image,omitempty"`        // Фотография сотрудника
	INN          *string     `json:"inn,omitempty"`          // ИНН сотрудника (в формате ИНН физического лица)
	LastName     *string     `json:"lastName,omitempty"`     // Фамилия
	Meta         *Meta       `json:"meta,omitempty"`         // Метаданные
	MiddleName   *string     `json:"middleName,omitempty"`   // Отчество
	Name         *string     `json:"name,omitempty"`         // Наименование
	Owner        *Employee   `json:"owner,omitempty"`        // Владелец (Сотрудник)
	Phone        *string     `json:"phone,omitempty"`        // Телефон сотрудника
	Position     *string     `json:"position,omitempty"`     // Должность сотрудника
	Shared       *bool       `json:"shared,omitempty"`       // Общий доступ
	ShortFio     *string     `json:"shortFio,omitempty"`     // Краткое ФИО
	UID          *string     `json:"uid,omitempty"`          // Логин Сотрудника
	Updated      *Timestamp  `json:"updated,omitempty"`      // Момент последнего обновления Сотрудника
}

func (e Employee) String() string {
	return Stringify(e)
}

func (e Employee) MetaType() MetaType {
	return MetaTypeEmployee
}

// MailActivationRequired структура ответа на запрос активации сотрудника
// Если поле равно true и сотрудник ранее не был активен, это означает, что на указанную у сотрудника почту было выслано письмо со ссылкой на вход для сотрудника.
// Если поле равно false, то можно использовать ранее заданный пароль для данного пользователя
type MailActivationRequired struct {
	MailActivationRequired bool `json:"mailActivationRequired"`
}

// EmployeePermission Права Сотрудника.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sotrudnik-rabota-s-prawami-sotrudnika
type EmployeePermission struct {
	AuthorizedIpNetmask *string  `json:"authorizedIpNetmask,omitempty"`
	AuthorizedIpNetwork *string  `json:"authorizedIpNetwork,omitempty"`
	Email               *string  `json:"email,omitempty"`
	Group               *Group   `json:"group,omitempty"`
	IsActive            *bool    `json:"isActive,omitempty"`
	Login               *string  `json:"login,omitempty"`
	Role                *Role    `json:"role,omitempty"`
	AuthorizedHosts     []string `json:"authorizedHosts,omitempty"`
}

func (e EmployeePermission) String() string {
	return Stringify(e)
}
