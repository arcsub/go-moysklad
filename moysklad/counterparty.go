package moysklad

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Counterparty Контрагент.
// Ключевое слово: counterparty
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent
type Counterparty struct {
	AccountID          *uuid.UUID                `json:"accountId,omitempty"`          // ID учетной записи
	Accounts           *MetaArray[AgentAccount]  `json:"accounts,omitempty"`           // Массив счетов Контрагентов
	ActualAddress      *string                   `json:"actualAddress,omitempty"`      // Фактический адрес Контрагента
	ActualAddressFull  *Address                  `json:"actualAddressFull,omitempty"`  // Фактический адрес Контрагента с детализацией по отдельным полям
	Archived           *bool                     `json:"archived,omitempty"`           // Добавлен ли Контрагент в архив
	Attributes         *Attributes               `json:"attributes,omitempty"`         // Массив метаданных доп. полей
	BonusPoints        *int                      `json:"bonusPoints,omitempty"`        // Бонусные баллы по активной бонусной программе
	BonusProgram       *BonusProgram             `json:"bonusProgram,omitempty"`       // Метаданные активной Бонусной программы
	Code               *string                   `json:"code,omitempty"`               // Код Контрагента
	CompanyType        CompanyType               `json:"companyType,omitempty"`        // Тип Контрагента. В зависимости от значения данного поля набор выводимых реквизитов контрагента может меняться
	ContactPersons     *MetaArray[ContactPerson] `json:"contactpersons,omitempty"`     // Массив контактных лиц фирмы Контрагента
	Created            *Timestamp                `json:"created,omitempty"`            // Момент создания
	Description        *string                   `json:"description,omitempty"`        // Комментарий к Контрагенту
	DiscountCardNumber *string                   `json:"discountCardNumber,omitempty"` // Номер дисконтной карты Контрагента
	Discounts          *Discounts                `json:"discounts,omitempty"`          // Массив скидок Контрагента. Массив может содержать персональные и накопительные скидки. Персональная скидка выводится, если хотя бы раз изменялся процент скидки для контрагента, значение будет указано в поле personalDiscount
	Email              *string                   `json:"email,omitempty"`              // Адрес электронной почты
	ExternalCode       *string                   `json:"externalCode,omitempty"`       // Внешний код Контрагента
	Fax                *string                   `json:"fax,omitempty"`                // Номер факса
	Files              *Files                    `json:"files,omitempty"`              // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group              *Group                    `json:"group,omitempty"`              // Отдел сотрудника
	ID                 *uuid.UUID                `json:"id,omitempty"`                 // ID сущности
	Meta               *Meta                     `json:"meta,omitempty"`               // Метаданные
	Name               *string                   `json:"name,omitempty"`               // Наименование
	Notes              *MetaArray[Note]          `json:"notes,omitempty"`              // Массив событий Контрагента
	Owner              *Employee                 `json:"owner,omitempty"`              // Владелец (Сотрудник)
	Phone              *string                   `json:"phone,omitempty"`              // Номер городского телефона
	PriceType          *PriceType                `json:"priceType,omitempty"`          // Тип цены Контрагента
	SalesAmount        *decimal.Decimal          `json:"salesAmount,omitempty"`        // Сумма продаж
	Shared             *bool                     `json:"shared,omitempty"`             // Общий доступ
	State              *State                    `json:"state,omitempty"`              // Метаданные Статуса Контрагента
	SyncID             *uuid.UUID                `json:"syncId,omitempty"`             // ID синхронизации
	Tags               *Tags                     `json:"tags,omitempty"`               // Группы контрагента
	Updated            *Timestamp                `json:"updated,omitempty"`            // Момент последнего обновления Контрагента
	BirthDate          *Timestamp                `json:"birthDate,omitempty"`          // Дата рождения [15-08-2023]
	CertificateDate    *Timestamp                `json:"certificateDate,omitempty"`    // Дата свидетельства
	CertificateNumber  *string                   `json:"certificateNumber,omitempty"`  // Номер свидетельства
	INN                *string                   `json:"inn,omitempty"`                // ИНН
	KPP                *string                   `json:"kpp,omitempty"`                // КПП
	LegalAddress       *string                   `json:"legalAddress,omitempty"`       // Юридический адрес
	LegalAddressFull   *Address                  `json:"legalAddressFull,omitempty"`   // Структурированный Юридический адрес
	LegalTitle         *string                   `json:"legalTitle,omitempty"`         // Полное наименование
	OGRN               *string                   `json:"ogrn,omitempty"`               // ОГРН
	OGRNIP             *string                   `json:"ogrnip,omitempty"`             // ОГРНИП
	OKPO               *string                   `json:"okpo,omitempty"`               // ОКПО
	Sex                SexType                   `json:"sex,omitempty"`                // Пол Контрагента [15-08-2023]
}

func (c Counterparty) String() string {
	return Stringify(c)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (c Counterparty) GetMeta() *Meta {
	return c.Meta
}

func (c Counterparty) MetaType() MetaType {
	return MetaTypeCounterparty
}

type DiscountData struct {
	Discount             *Discount `json:"discount,omitempty"`             // Скидка
	PersonalDiscount     *float64  `json:"personalDiscount,omitempty"`     // Значение персональной скидки
	DemandSumCorrection  *float64  `json:"demandSumCorrection,omitempty"`  // Коррекция суммы накоплений по скидке
	AccumulationDiscount *float64  `json:"accumulationDiscount,omitempty"` // Значение накопительной скидки
}

func (d DiscountData) String() string {
	return Stringify(d)
}

// CounterpartySettings Настройки справочника контрагентов
// Ключевое слово: counterpartysettings
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-nastrojki-sprawochnika-kontragentow
type CounterpartySettings struct {
	Meta            *Meta            `json:"meta,omitempty"`            // Метаданные
	CreateShared    *bool            `json:"createShared,omitempty"`    // Создавать новые документы с меткой «Общий»
	UniqueCodeRules *UniqueCodeRules `json:"uniqueCodeRules,omitempty"` // Настройки кодов контрагентов
}

func (c CounterpartySettings) String() string {
	return Stringify(c)
}

func (c CounterpartySettings) MetaType() MetaType {
	return MetaTypeCounterpartySettings
}

// Note Событие Контрагента.
// Ключевое слово: note
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-kontragenty-attributy-suschnosti-adres-sobytiq-kontragenta
type Note struct {
	AccountID         *uuid.UUID    `json:"accountId,omitempty"`         // ID учетной записи
	Agent             *Counterparty `json:"agent,omitempty"`             // Метаданные Контрагента
	Author            *Employee     `json:"author,omitempty"`            // Метаданные Сотрудника - создателя события (администратор аккаунта, если автор - приложение)
	AuthorApplication *Application  `json:"authorApplication,omitempty"` // Метаданные Приложения - создателя события
	Created           *Timestamp    `json:"created,omitempty"`           // Момент создания события Контрагента
	Description       *string       `json:"description,omitempty"`       // Текст события Контрагента
	ID                *uuid.UUID    `json:"id,omitempty"`                // ID сущности
	Meta              *Meta         `json:"meta,omitempty"`              // Метаданные
}

func (n Note) String() string {
	return Stringify(n)
}

func (n Note) MetaType() MetaType {
	return MetaTypeNote
}

// SexType Пол Контрагента
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-kontragenty-pol-kontragenta
type SexType string

const (
	SexMale   SexType = "MALE"   // Мужской
	SexFemale SexType = "FEMALE" // Женский
)

type Tags = []string
