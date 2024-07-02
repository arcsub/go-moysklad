package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Counterparty Контрагент.
//
// Код сущности: counterparty
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent
type Counterparty struct {
	Name               *string                     `json:"name,omitempty"`               // Наименование Контрагента
	OKPO               *string                     `json:"okpo,omitempty"`               // ОКПО
	ActualAddress      *string                     `json:"actualAddress,omitempty"`      // Фактический адрес Контрагента
	AccountID          *uuid.UUID                  `json:"accountId,omitempty"`          // ID учётной записи
	Archived           *bool                       `json:"archived,omitempty"`           // Добавлен ли Контрагент в архив
	Notes              *MetaArray[Note]            `json:"notes,omitempty"`              // Массив событий Контрагента
	BonusPoints        *int                        `json:"bonusPoints,omitempty"`        // Бонусные баллы по активной бонусной программе
	BonusProgram       *NullValue[BonusProgram]    `json:"bonusProgram,omitempty"`       // Метаданные активной Бонусной программы
	Code               *string                     `json:"code,omitempty"`               // Код Контрагента
	OGRNIP             *string                     `json:"ogrnip,omitempty"`             // ОГРНИП
	ContactPersons     *MetaArray[ContactPerson]   `json:"contactpersons,omitempty"`     // Массив контактных лиц фирмы Контрагента
	Created            *Timestamp                  `json:"created,omitempty"`            // Момент создания
	Description        *string                     `json:"description,omitempty"`        // Комментарий к Контрагенту
	DiscountCardNumber *string                     `json:"discountCardNumber,omitempty"` // Номер дисконтной карты Контрагента
	Discounts          Slice[CounterpartyDiscount] `json:"discounts,omitempty"`          // Массив скидок Контрагента. Массив может содержать персональные и накопительные скидки. Персональная скидка выводится, если хотя бы раз изменялся процент скидки для контрагента, значение будет указано в поле personalDiscount
	Email              *string                     `json:"email,omitempty"`              // Адрес электронной почты
	ExternalCode       *string                     `json:"externalCode,omitempty"`       // Внешний код Контрагента
	Owner              *Employee                   `json:"owner,omitempty"`              // Метаданные владельца (Сотрудника)
	Files              *MetaArray[File]            `json:"files,omitempty"`              // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group              *Group                      `json:"group,omitempty"`              // Отдел сотрудника
	ID                 *uuid.UUID                  `json:"id,omitempty"`                 // ID Контрагента
	Meta               *Meta                       `json:"meta,omitempty"`               // Метаданные Контрагента
	ActualAddressFull  *Address                    `json:"actualAddressFull,omitempty"`  // Фактический адрес Контрагента с детализацией по отдельным полям
	Accounts           *MetaArray[AgentAccount]    `json:"accounts,omitempty"`           // Массив счетов Контрагентов
	Fax                *string                     `json:"fax,omitempty"`                // Номер факса
	Phone              *string                     `json:"phone,omitempty"`              // Номер городского телефона
	PriceType          *PriceType                  `json:"priceType,omitempty"`          // Тип цены Контрагента
	SalesAmount        *float64                    `json:"salesAmount,omitempty"`        // Сумма продаж
	Shared             *bool                       `json:"shared,omitempty"`             // Общий доступ
	State              *NullValue[State]           `json:"state,omitempty"`              // Метаданные Статуса Контрагента
	SyncID             *uuid.UUID                  `json:"syncId,omitempty"`             // ID синхронизации
	Tags               Slice[string]               `json:"tags,omitempty"`               // Группы контрагента
	Updated            *Timestamp                  `json:"updated,omitempty"`            // Момент последнего обновления Контрагента
	BirthDate          *Timestamp                  `json:"birthDate,omitempty"`          // Дата рождения Контрагента типа [Физическое лицо]. Игнорируется для Контрагентов типов [Индивидуальный предприниматель, Юридическое лицо]
	CertificateDate    *Timestamp                  `json:"certificateDate,omitempty"`    // Дата свидетельства
	CertificateNumber  *string                     `json:"certificateNumber,omitempty"`  // Номер свидетельства
	INN                *string                     `json:"inn,omitempty"`                // ИНН
	KPP                *string                     `json:"kpp,omitempty"`                // КПП
	LegalAddress       *string                     `json:"legalAddress,omitempty"`       // Юридический адрес Контрагента
	LegalFirstName     *string                     `json:"legalFirstName,omitempty"`     // Имя для Контрагента типа [Индивидуальный предприниматель, Физическое лицо]. Игнорируется для Контрагентов типа [Юридическое лицо]
	LegalLastName      *string                     `json:"legalLastName,omitempty"`      // Фамилия для Контрагента типа [Индивидуальный предприниматель, Физическое лицо]. Игнорируется для Контрагентов типа [Юридическое лицо]
	LegalMiddleName    *string                     `json:"legalMiddleName,omitempty"`    // Отчество для Контрагента типа [Индивидуальный предприниматель, Физическое лицо]. Игнорируется для Контрагентов типа [Юридическое лицо]
	LegalAddressFull   *Address                    `json:"legalAddressFull,omitempty"`   // Юридический адрес Контрагента с детализацией по отдельным полям
	LegalTitle         *string                     `json:"legalTitle,omitempty"`         // Полное наименование для Контрагента типа [Юридическое лицо]. Игнорируется для Контрагентов типа [Индивидуальный предприниматель, Физическое лицо], если передано одно из значений для ФИО и формируется автоматически на основе получаемых ФИО Контрагента
	OGRN               *string                     `json:"ogrn,omitempty"`               // ОГРН
	CompanyType        CompanyType                 `json:"companyType,omitempty"`        // Тип Контрагента. В зависимости от значения данного поля набор выводимых реквизитов контрагента может меняться.
	Sex                Sex                         `json:"sex,omitempty"`                // Пол Контрагента
	Attributes         Slice[Attribute]            `json:"attributes,omitempty"`         // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (counterparty Counterparty) Clean() *Counterparty {
	if counterparty.Meta == nil {
		return nil
	}
	return &Counterparty{Meta: counterparty.Meta}
}

// asCOAgent реализует интерфейс [AgentCounterpartyOrganizationInterface].
func (counterparty Counterparty) asCOAgent() *Agent {
	return counterparty.asAgent()
}

// asAgent реализует интерфейс [AgentInterface].
func (counterparty Counterparty) asAgent() *Agent {
	if counterparty.Meta == nil {
		return nil
	}
	return &Agent{Meta: counterparty.Meta}
}

// asCEAgent реализует интерфейс [AgentCounterpartyEmployeeInterface].
func (counterparty Counterparty) asCEAgent() *Agent {
	return counterparty.asCOAgent()
}

// GetName возвращает Наименование Контрагента.
func (counterparty Counterparty) GetName() string {
	return Deref(counterparty.Name)
}

// GetOKPO возвращает ОКПО.
func (counterparty Counterparty) GetOKPO() string {
	return Deref(counterparty.OKPO)
}

// GetActualAddress возвращает Фактический адрес Контрагента.
func (counterparty Counterparty) GetActualAddress() string {
	return Deref(counterparty.ActualAddress)
}

// GetAccountID возвращает ID учётной записи.
func (counterparty Counterparty) GetAccountID() uuid.UUID {
	return Deref(counterparty.AccountID)
}

// GetArchived возвращает true, если Контрагент добавлен в архив.
func (counterparty Counterparty) GetArchived() bool {
	return Deref(counterparty.Archived)
}

// GetNotes возвращает Массив событий Контрагента.
func (counterparty Counterparty) GetNotes() MetaArray[Note] {
	return Deref(counterparty.Notes)
}

// GetBonusPoints возвращает Бонусные баллы по активной бонусной программе.
func (counterparty Counterparty) GetBonusPoints() int {
	return Deref(counterparty.BonusPoints)
}

// GetBonusProgram возвращает Метаданные активной бонусной программы.
func (counterparty Counterparty) GetBonusProgram() BonusProgram {
	return counterparty.BonusProgram.GetValue()
}

// GetCode возвращает Код Контрагента.
func (counterparty Counterparty) GetCode() string {
	return Deref(counterparty.Code)
}

// GetOGRNIP возвращает ОГРНИП.
func (counterparty Counterparty) GetOGRNIP() string {
	return Deref(counterparty.OGRNIP)
}

// GetContactPersons возвращает Массив контактных лиц фирмы Контрагента.
func (counterparty Counterparty) GetContactPersons() MetaArray[ContactPerson] {
	return Deref(counterparty.ContactPersons)
}

// GetCreated возвращает Момент создания.
func (counterparty Counterparty) GetCreated() Timestamp {
	return Deref(counterparty.Created)
}

// GetDescription возвращает Комментарий к Контрагенту.
func (counterparty Counterparty) GetDescription() string {
	return Deref(counterparty.Description)
}

// GetDiscountCardNumber возвращает Номер дисконтной карты Контрагента.
func (counterparty Counterparty) GetDiscountCardNumber() string {
	return Deref(counterparty.DiscountCardNumber)
}

// GetDiscounts возвращает Массив скидок Контрагента.
//
// Массив может содержать персональные и накопительные скидки.
//
// Персональная скидка выводится, если хотя бы раз изменялся процент скидки для контрагента, значение будет указано в поле personalDiscount.
func (counterparty Counterparty) GetDiscounts() Slice[CounterpartyDiscount] {
	return counterparty.Discounts
}

// GetEmail возвращает Адрес электронной почты.
func (counterparty Counterparty) GetEmail() string {
	return Deref(counterparty.Email)
}

// GetExternalCode возвращает Внешний код Контрагента.
func (counterparty Counterparty) GetExternalCode() string {
	return Deref(counterparty.ExternalCode)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (counterparty Counterparty) GetOwner() Employee {
	return Deref(counterparty.Owner)
}

// GetFiles возвращает Метаданные массива Файлов.
func (counterparty Counterparty) GetFiles() MetaArray[File] {
	return Deref(counterparty.Files)
}

// GetGroup возвращает Отдел сотрудника.
func (counterparty Counterparty) GetGroup() Group {
	return Deref(counterparty.Group)
}

// GetID возвращает ID Контрагента.
func (counterparty Counterparty) GetID() uuid.UUID {
	return Deref(counterparty.ID)
}

// GetMeta возвращает Метаданные Контрагента.
func (counterparty Counterparty) GetMeta() Meta {
	return Deref(counterparty.Meta)
}

// GetActualAddressFull возвращает Фактический адрес Контрагента с детализацией по отдельным полям.
func (counterparty Counterparty) GetActualAddressFull() Address {
	return Deref(counterparty.ActualAddressFull)
}

// GetAccounts возвращает Массив счетов Контрагентов.
func (counterparty Counterparty) GetAccounts() MetaArray[AgentAccount] {
	return Deref(counterparty.Accounts)
}

// GetFax возвращает Номер факса.
func (counterparty Counterparty) GetFax() string {
	return Deref(counterparty.Fax)
}

// GetPhone возвращает Номер городского телефона.
func (counterparty Counterparty) GetPhone() string {
	return Deref(counterparty.Phone)
}

// GetPriceType возвращает Тип цены Контрагента.
func (counterparty Counterparty) GetPriceType() PriceType {
	return Deref(counterparty.PriceType)
}

// GetSalesAmount возвращает Сумму продаж.
func (counterparty Counterparty) GetSalesAmount() float64 {
	return Deref(counterparty.SalesAmount)
}

// GetShared возвращает флаг Общего доступа.
func (counterparty Counterparty) GetShared() bool {
	return Deref(counterparty.Shared)
}

// GetState возвращает Метаданные статуса Контрагента.
func (counterparty Counterparty) GetState() State {
	return counterparty.State.GetValue()
}

// GetSyncID возвращает ID синхронизации.
func (counterparty Counterparty) GetSyncID() uuid.UUID {
	return Deref(counterparty.SyncID)
}

// GetTags возвращает Группы контрагента.
func (counterparty Counterparty) GetTags() Slice[string] {
	return counterparty.Tags
}

// GetUpdated возвращает Момент последнего обновления контрагента.
func (counterparty Counterparty) GetUpdated() Timestamp {
	return Deref(counterparty.Updated)
}

// GetBirthDate возвращает Дату рождения Контрагента типа [Физическое лицо].
//
// Игнорируется для Контрагентов типов [Индивидуальный предприниматель, Юридическое лицо].
func (counterparty Counterparty) GetBirthDate() Timestamp {
	return Deref(counterparty.BirthDate)
}

// GetCertificateDate возвращает Дату свидетельства.
func (counterparty Counterparty) GetCertificateDate() Timestamp {
	return Deref(counterparty.CertificateDate)
}

// GetCertificateNumber возвращает Номер свидетельства.
func (counterparty Counterparty) GetCertificateNumber() string {
	return Deref(counterparty.CertificateNumber)
}

// GetINN возвращает ИНН.
func (counterparty Counterparty) GetINN() string {
	return Deref(counterparty.INN)
}

// GetKPP возвращает КПП.
func (counterparty Counterparty) GetKPP() string {
	return Deref(counterparty.KPP)
}

// GetLegalAddress возвращает Юридический адрес Контрагента.
func (counterparty Counterparty) GetLegalAddress() string {
	return Deref(counterparty.LegalAddress)
}

// GetLegalAddressFull возвращает Юридический адрес Контрагента с детализацией по отдельным полям.
func (counterparty Counterparty) GetLegalAddressFull() Address {
	return Deref(counterparty.LegalAddressFull)
}

// GetLegalFirstName возвращает Имя для Контрагента типа [Индивидуальный предприниматель, Физическое лицо].
//
// Игнорируется для Контрагентов типа [Юридическое лицо].
func (counterparty Counterparty) GetLegalFirstName() string {
	return Deref(counterparty.LegalFirstName)
}

// GetLegalLastName возвращает Фамилию для Контрагента типа [Индивидуальный предприниматель, Физическое лицо].
//
// Игнорируется для Контрагентов типа [Юридическое лицо]
func (counterparty Counterparty) GetLegalLastName() string {
	return Deref(counterparty.LegalLastName)
}

// GetLegalMiddleName возвращает Отчество для Контрагента типа [Индивидуальный предприниматель, Физическое лицо].
//
// Игнорируется для Контрагентов типа [Юридическое лицо]
func (counterparty Counterparty) GetLegalMiddleName() string {
	return Deref(counterparty.LegalMiddleName)
}

// GetLegalTitle возвращает полное наименование для Контрагента типа [Юридическое лицо].
//
// Игнорируется для Контрагентов типа [Индивидуальный предприниматель, Физическое лицо],
// если передано одно из значений для ФИО и формируется автоматически на основе получаемых ФИО Контрагента.
func (counterparty Counterparty) GetLegalTitle() string {
	return Deref(counterparty.LegalTitle)
}

// GetOGRN возвращает ОГРН.
func (counterparty Counterparty) GetOGRN() string {
	return Deref(counterparty.OGRN)
}

// GetCompanyType возвращает Тип Контрагента.
func (counterparty Counterparty) GetCompanyType() CompanyType {
	return counterparty.CompanyType
}

// GetSex возвращает Пол Контрагента.
func (counterparty Counterparty) GetSex() Sex {
	return counterparty.Sex
}

// GetAttributes возвращает Список метаданных доп. полей.
func (counterparty Counterparty) GetAttributes() Slice[Attribute] {
	return counterparty.Attributes
}

// SetName устанавливает Наименование Контрагента.
func (counterparty *Counterparty) SetName(name string) *Counterparty {
	counterparty.Name = &name
	return counterparty
}

// SetOKPO устанавливает ОКПО.
func (counterparty *Counterparty) SetOKPO(okpo string) *Counterparty {
	counterparty.OKPO = &okpo
	return counterparty
}

// SetActualAddress устанавливает Фактический адрес Контрагента.
func (counterparty *Counterparty) SetActualAddress(actualAddress string) *Counterparty {
	counterparty.ActualAddress = &actualAddress
	return counterparty
}

// SetArchived устанавливает флаг нахождения Контрагента в архиве.
func (counterparty *Counterparty) SetArchived(archived bool) *Counterparty {
	counterparty.Archived = &archived
	return counterparty
}

// SetNotes устанавливает Массив событий Контрагента.
//
// Принимает множество объектов [Note].
func (counterparty *Counterparty) SetNotes(notes ...*Note) *Counterparty {
	counterparty.Notes = NewMetaArrayFrom(notes)
	return counterparty
}

// SetBonusProgram устанавливает Метаданные бонусной программы.
//
// Передача nil передаёт сброс значения (null).
func (counterparty *Counterparty) SetBonusProgram(bonusProgram *BonusProgram) *Counterparty {
	counterparty.BonusProgram = NewNullValue(bonusProgram)
	return counterparty
}

// SetCode устанавливает Код Контрагента.
func (counterparty *Counterparty) SetCode(code string) *Counterparty {
	counterparty.Code = &code
	return counterparty
}

// SetOGRNIP устанавливает ОГРНИП.
func (counterparty *Counterparty) SetOGRNIP(ogrnip string) *Counterparty {
	counterparty.OGRNIP = &ogrnip
	return counterparty
}

// SetContactPersons устанавливает Массив контактных лиц фирмы Контрагента.
//
// Принимает множество объектов [ContactPerson].
func (counterparty *Counterparty) SetContactPersons(contactPersons ...*ContactPerson) *Counterparty {
	counterparty.ContactPersons = NewMetaArrayFrom(contactPersons)
	return counterparty
}

// SetDescription устанавливает Комментарий к Контрагенту.
func (counterparty *Counterparty) SetDescription(description string) *Counterparty {
	counterparty.Description = &description
	return counterparty
}

// SetDiscountCardNumber устанавливает Номер дисконтной карты Контрагента.
func (counterparty *Counterparty) SetDiscountCardNumber(discountCardNumber string) *Counterparty {
	counterparty.DiscountCardNumber = &discountCardNumber
	return counterparty
}

// SetEmail устанавливает Адрес электронной почты.
func (counterparty *Counterparty) SetEmail(email string) *Counterparty {
	counterparty.Email = &email
	return counterparty
}

// SetExternalCode устанавливает Внешний код Контрагента.
func (counterparty *Counterparty) SetExternalCode(externalCode string) *Counterparty {
	counterparty.ExternalCode = &externalCode
	return counterparty
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (counterparty *Counterparty) SetOwner(owner *Employee) *Counterparty {
	if owner != nil {
		counterparty.Owner = owner.Clean()
	}
	return counterparty
}

// SetFiles устанавливает Метаданные массива Файлов.
//
// Принимает множество объектов [File].
func (counterparty *Counterparty) SetFiles(files ...*File) *Counterparty {
	counterparty.Files = NewMetaArrayFrom(files)
	return counterparty
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (counterparty *Counterparty) SetGroup(group *Group) *Counterparty {
	if group != nil {
		counterparty.Group = group.Clean()
	}
	return counterparty
}

// SetMeta устанавливает Метаданные Контрагента.
func (counterparty *Counterparty) SetMeta(meta *Meta) *Counterparty {
	counterparty.Meta = meta
	return counterparty
}

// SetActualAddressFull устанавливает Фактический адрес Контрагента с детализацией по отдельным полям.
//
// Передача nil передаёт сброс значения (null).
func (counterparty *Counterparty) SetActualAddressFull(actualAddressFull *Address) *Counterparty {
	if actualAddressFull == nil {
		counterparty.SetActualAddress("")
	} else {
		counterparty.ActualAddressFull = actualAddressFull
	}
	return counterparty
}

// SetAccounts устанавливает Массив счетов Контрагентов.
func (counterparty *Counterparty) SetAccounts(accounts ...*AgentAccount) *Counterparty {
	counterparty.Accounts = NewMetaArrayFrom(accounts)
	return counterparty
}

// SetFax устанавливает Номер факса.
func (counterparty *Counterparty) SetFax(fax string) *Counterparty {
	counterparty.Fax = &fax
	return counterparty
}

// SetPhone устанавливает Номер городского телефона.
func (counterparty *Counterparty) SetPhone(phone string) *Counterparty {
	counterparty.Phone = &phone
	return counterparty
}

// SetPriceType устанавливает Тип цены Контрагента.
func (counterparty *Counterparty) SetPriceType(priceType *PriceType) *Counterparty {
	if priceType != nil {
		counterparty.PriceType = priceType.Clean()
	}
	return counterparty
}

// SetShared устанавливает флаг общего доступа.
func (counterparty *Counterparty) SetShared(shared bool) *Counterparty {
	counterparty.Shared = &shared
	return counterparty
}

// SetState устанавливает Метаданные Статуса Контрагента.
func (counterparty *Counterparty) SetState(state *State) *Counterparty {
	counterparty.State = NewNullValue(state)
	return counterparty
}

// SetSyncID устанавливает ID синхронизации.
func (counterparty *Counterparty) SetSyncID(syncID uuid.UUID) *Counterparty {
	counterparty.SyncID = &syncID
	return counterparty
}

// SetTags устанавливает Группы контрагента.
func (counterparty *Counterparty) SetTags(tags ...string) *Counterparty {
	counterparty.Tags = NewSliceFrom(tags)
	return counterparty
}

// SetBirthDate устанавливает Дату рождения Контрагента типа [Физическое лицо].
//
// Игнорируется для Контрагентов типов [Индивидуальный предприниматель, Юридическое лицо].
func (counterparty *Counterparty) SetBirthDate(birthDate *Timestamp) *Counterparty {
	counterparty.BirthDate = birthDate
	return counterparty
}

// SetCertificateDate устанавливает Дату свидетельства.
func (counterparty *Counterparty) SetCertificateDate(certificateDate *Timestamp) *Counterparty {
	counterparty.CertificateDate = certificateDate
	return counterparty
}

// SetCertificateNumber устанавливает Номер свидетельства.
func (counterparty *Counterparty) SetCertificateNumber(certificateNumber string) *Counterparty {
	counterparty.CertificateNumber = &certificateNumber
	return counterparty
}

// SetINN устанавливает ИНН.
func (counterparty *Counterparty) SetINN(inn string) *Counterparty {
	counterparty.INN = &inn
	return counterparty
}

// SetKPP устанавливает КПП.
func (counterparty *Counterparty) SetKPP(kpp string) *Counterparty {
	counterparty.KPP = &kpp
	return counterparty
}

// SetLegalAddress устанавливает Юридический адрес Контрагента.
func (counterparty *Counterparty) SetLegalAddress(legalAddress string) *Counterparty {
	counterparty.LegalAddress = &legalAddress
	return counterparty
}

// SetLegalAddressFull устанавливает Юридический адрес Контрагента с детализацией по отдельным полям.
//
// Передача nil передаёт сброс значения (null).
func (counterparty *Counterparty) SetLegalAddressFull(legalAddressFull *Address) *Counterparty {
	if legalAddressFull == nil {
		counterparty.SetLegalAddress("")
	} else {
		counterparty.LegalAddressFull = legalAddressFull
	}
	return counterparty
}

// SetLegalFirstName устанавливает Имя для Контрагента типа [Индивидуальный предприниматель, Физическое лицо].
//
// Игнорируется для Контрагентов типа [Юридическое лицо].
func (counterparty *Counterparty) SetLegalFirstName(legalFirstName string) *Counterparty {
	counterparty.LegalFirstName = &legalFirstName
	return counterparty
}

// SetLegalLastName устанавливает Фамилию для Контрагента типа [Индивидуальный предприниматель, Физическое лицо].
//
// Игнорируется для Контрагентов типа [Юридическое лицо].
func (counterparty *Counterparty) SetLegalLastName(legalLastName string) *Counterparty {
	counterparty.LegalLastName = &legalLastName
	return counterparty
}

// SetLegalMiddleName устанавливает Отчество для Контрагента типа [Индивидуальный предприниматель, Физическое лицо].
//
// Игнорируется для Контрагентов типа [Юридическое лицо].
func (counterparty *Counterparty) SetLegalMiddleName(legalMiddleName string) *Counterparty {
	counterparty.LegalMiddleName = &legalMiddleName
	return counterparty
}

// SetOGRN устанавливает ОГРН.
func (counterparty *Counterparty) SetOGRN(ogrn string) *Counterparty {
	counterparty.OGRN = &ogrn
	return counterparty
}

// SetCompanyType устанавливает Тип Контрагента.
func (counterparty *Counterparty) SetCompanyType(companyType CompanyType) *Counterparty {
	counterparty.CompanyType = companyType
	return counterparty
}

// SetSex устанавливает Пол Контрагента.
func (counterparty *Counterparty) SetSex(sex Sex) *Counterparty {
	counterparty.Sex = sex
	return counterparty
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (counterparty *Counterparty) SetAttributes(attributes ...*Attribute) *Counterparty {
	counterparty.Attributes.Push(attributes...)
	return counterparty
}

// String реализует интерфейс [fmt.Stringer].
func (counterparty Counterparty) String() string {
	return Stringify(counterparty)
}

// MetaType возвращает код сущности.
func (Counterparty) MetaType() MetaType {
	return MetaTypeCounterparty
}

// Update shortcut
func (counterparty Counterparty) Update(ctx context.Context, client *Client, params ...*Params) (*Counterparty, *resty.Response, error) {
	return NewCounterpartyService(client).Update(ctx, counterparty.GetID(), &counterparty, params...)
}

// Create shortcut
func (counterparty Counterparty) Create(ctx context.Context, client *Client, params ...*Params) (*Counterparty, *resty.Response, error) {
	return NewCounterpartyService(client).Create(ctx, &counterparty, params...)
}

// Delete shortcut
func (counterparty Counterparty) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewCounterpartyService(client).Delete(ctx, counterparty.GetID())
}

// CounterpartyDiscount скидка Контрагента.
type CounterpartyDiscount struct {
	Discount             *MetaWrapper `json:"discount,omitempty"`             // Метаданные Скидки
	PersonalDiscount     *float64     `json:"personalDiscount,omitempty"`     // Значение персональной скидки
	DemandSumCorrection  *float64     `json:"demandSumCorrection,omitempty"`  // Коррекция суммы накоплений по скидке
	AccumulationDiscount *float64     `json:"accumulationDiscount,omitempty"` // Значение накопительной скидки
}

// String реализует интерфейс [fmt.Stringer].
func (discountsData CounterpartyDiscount) String() string {
	return Stringify(discountsData)
}

// CounterpartySettings Настройки справочника контрагентов
//
// Код сущности: counterpartysettings
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-nastrojki-sprawochnika-kontragentow
type CounterpartySettings struct {
	Meta            *Meta            `json:"meta,omitempty"`            // Метаданные Настроек справочника контрагентов
	CreateShared    *bool            `json:"createShared,omitempty"`    // Создавать новые документы с меткой «Общий»
	UniqueCodeRules *UniqueCodeRules `json:"uniqueCodeRules,omitempty"` // Настройки кодов контрагентов
}

// GetMeta возвращает Метаданные Настроек справочника контрагентов.
func (counterpartySettings CounterpartySettings) GetMeta() Meta {
	return Deref(counterpartySettings.Meta)
}

// GetCreateShared возвращает флаг создания новых контрагентов с меткой «Общий».
func (counterpartySettings CounterpartySettings) GetCreateShared() bool {
	return Deref(counterpartySettings.CreateShared)
}

// GetUniqueCodeRules возвращает Настройки кодов контрагентов.
func (counterpartySettings CounterpartySettings) GetUniqueCodeRules() UniqueCodeRules {
	return Deref(counterpartySettings.UniqueCodeRules)
}

// SetCreateShared устанавливает флаг создания новых контрагентов с меткой «Общий».
func (counterpartySettings *CounterpartySettings) SetCreateShared(createShared bool) *CounterpartySettings {
	counterpartySettings.CreateShared = &createShared
	return counterpartySettings
}

// SetUniqueCodeRules устанавливает Настройки кодов контрагентов.
func (counterpartySettings *CounterpartySettings) SetUniqueCodeRules(uniqueCodeRules *UniqueCodeRules) *CounterpartySettings {
	counterpartySettings.UniqueCodeRules = uniqueCodeRules
	return counterpartySettings
}

// String реализует интерфейс [fmt.Stringer].
func (counterpartySettings CounterpartySettings) String() string {
	return Stringify(counterpartySettings)
}

// MetaType возвращает код сущности.
func (CounterpartySettings) MetaType() MetaType {
	return MetaTypeCounterpartySettings
}

// Note Событие Контрагента.
//
// Код сущности: note
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-kontragenty-attributy-suschnosti-adres-sobytiq-kontragenta
type Note struct {
	AccountID         *uuid.UUID    `json:"accountId,omitempty"`         // ID учётной записи
	Agent             *Counterparty `json:"agent,omitempty"`             // Метаданные Контрагента
	Author            *Employee     `json:"author,omitempty"`            // Метаданные Сотрудника - создателя события (администратор аккаунта, если автор - приложение)
	AuthorApplication *Application  `json:"authorApplication,omitempty"` // Метаданные Приложения - создателя события
	Created           *Timestamp    `json:"created,omitempty"`           // Момент создания события Контрагента
	Description       *string       `json:"description,omitempty"`       // Текст события Контрагента
	ID                *uuid.UUID    `json:"id,omitempty"`                // ID события контрагента
	Meta              *Meta         `json:"meta,omitempty"`              // Метаданные события контрагента
}

// GetAccountID возвращает ID учётной записи.
func (note Note) GetAccountID() uuid.UUID {
	return Deref(note.AccountID)
}

// GetAgent возвращает Метаданные Контрагента.
func (note Note) GetAgent() Counterparty {
	return Deref(note.Agent)
}

// GetAuthor возвращает Метаданные Сотрудника - создателя события (администратор аккаунта, если автор - приложение).
func (note Note) GetAuthor() Employee {
	return Deref(note.Author)
}

// GetAuthorApplication возвращает Метаданные Приложения - создателя события.
func (note Note) GetAuthorApplication() Application {
	return Deref(note.AuthorApplication)
}

// GetCreated возвращает Момент создания события Контрагента.
func (note Note) GetCreated() Timestamp {
	return Deref(note.Created)
}

// GetDescription возвращает Текст события Контрагента.
func (note Note) GetDescription() string {
	return Deref(note.Description)
}

// GetID возвращает ID события контрагента.
func (note Note) GetID() uuid.UUID {
	return Deref(note.ID)
}

// GetMeta возвращает Метаданные события контрагента.
func (note Note) GetMeta() Meta {
	return Deref(note.Meta)
}

// SetDescription устанавливает Текст события Контрагента.
func (note *Note) SetDescription(description string) *Note {
	note.Description = &description
	return note
}

// SetMeta устанавливает Метаданные события контрагента.
func (note *Note) SetMeta(meta *Meta) *Note {
	note.Meta = meta
	return note
}

// String реализует интерфейс [fmt.Stringer].
func (note Note) String() string {
	return Stringify(note)
}

// MetaType возвращает код сущности.
func (Note) MetaType() MetaType {
	return MetaTypeNote
}

// Sex Пол Контрагента.
//
// Возможные значения:
//   - Male   – Мужской
//   - Female – Женский
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kontragent-kontragenty-pol-kontragenta
type Sex string

const (
	Male   Sex = "MALE"   // Мужской
	Female Sex = "FEMALE" // Женский
)

// CounterpartyService описывает методы сервиса для работы с контрагентами.
type CounterpartyService interface {
	// GetList выполняет запрос на получение списка контрагентов.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[Counterparty], *resty.Response, error)

	// Create выполняет запрос на создание контрагента.
	// Обязательные поля для заполнения:
	//	- name (Наименование контрагента)
	// Принимает контекст, контрагент и опционально объект параметров запроса Params.
	// Возвращает созданный контрагент.
	Create(ctx context.Context, counterparty *Counterparty, params ...*Params) (*Counterparty, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или контрагентов.
	// Изменяемые контрагенты должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список контрагентов и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых контрагентов.
	CreateUpdateMany(ctx context.Context, counterpartyList Slice[Counterparty], params ...*Params) (*Slice[Counterparty], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление контрагентов.
	// Принимает контекст и множество контрагентов.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*Counterparty) (*DeleteManyResponse, *resty.Response, error)

	// Delete выполняет запрос на удаление контрагента.
	// Принимает контекст и ID контрагента.
	// Возвращает true в случае успешного удаления контрагента.
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельного контрагента по ID.
	// Принимает контекст, ID контрагента и опционально объект параметров запроса Params.
	// Возвращает найденный контрагент.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Counterparty, *resty.Response, error)

	// Update выполняет запрос на изменение контрагента.
	// Принимает контекст, контрагент и опционально объект параметров запроса Params.
	// Возвращает изменённый контрагент.
	Update(ctx context.Context, id uuid.UUID, counterparty *Counterparty, params ...*Params) (*Counterparty, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных контрагентов.
	// Принимает контекст.
	// Возвращает объект метаданных MetaAttributesStatesSharedTagsWrapper.
	GetMetadata(ctx context.Context) (*MetaAttributesStatesSharedTagsWrapper, *resty.Response, error)

	// GetAttributeList выполняет запрос на получение списка доп полей.
	// Принимает контекст.
	// Возвращает объект List.
	GetAttributeList(ctx context.Context) (*List[Attribute], *resty.Response, error)

	// GetAttributeByID выполняет запрос на получение отдельного доп поля по ID.
	// Принимает контекст и ID доп поля.
	// Возвращает найденное доп поле.
	GetAttributeByID(ctx context.Context, id uuid.UUID) (*Attribute, *resty.Response, error)

	// CreateAttribute выполняет запрос на создание доп поля.
	// Принимает контекст и доп поле.
	// Возвращает созданное доп поле.
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)

	// CreateUpdateAttributeMany выполняет запрос на массовое создание и/или изменение доп полей.
	// Изменяемые доп поля должны содержать идентификатор в виде метаданных.
	// Принимает контекст и множество доп полей.
	// Возвращает список созданных и/или изменённых доп полей.
	CreateUpdateAttributeMany(ctx context.Context, attributes ...*Attribute) (*Slice[Attribute], *resty.Response, error)

	// UpdateAttribute выполняет запрос на изменения доп поля.
	// Принимает контекст, ID доп поля и доп поле.
	// Возвращает изменённое доп поле.
	UpdateAttribute(ctx context.Context, id uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)

	// DeleteAttribute выполняет запрос на удаление доп поля.
	// Принимает контекст и ID доп поля.
	// Возвращает true в случае успешного удаления доп поля.
	DeleteAttribute(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// DeleteAttributeMany выполняет запрос на массовое удаление доп полей.
	// Принимает контекст и множество доп полей.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteAttributeMany(ctx context.Context, attributes ...*Attribute) (*DeleteManyResponse, *resty.Response, error)

	// GetSettings выполняет запрос на получение настроек справочника контрагентов.
	// Принимает контекст.
	// Возвращает настройки справочника контрагентов.
	GetSettings(ctx context.Context) (*CounterpartySettings, *resty.Response, error)

	// UpdateSettings выполняет запрос на изменение настроек справочника контрагентов.
	// Принимает контекст и настройки справочника контрагентов.
	// Возвращает изменённые настройки справочника контрагентов.
	UpdateSettings(ctx context.Context, settings *CounterpartySettings) (*CounterpartySettings, *resty.Response, error)

	// GetAccountList выполняет запрос на получение списка счетов контрагента.
	// Принимает контекст и ID контрагента.
	// Возвращает объект List.
	GetAccountList(ctx context.Context, id uuid.UUID) (*List[AgentAccount], *resty.Response, error)

	// GetAccountByID выполняет запрос на получение отдельного счёта контрагента по ID.
	// Принимает контекст, ID контрагента и ID счёта контрагента.
	// Возвращает найденный счёт контрагента.
	GetAccountByID(ctx context.Context, id uuid.UUID, accountID uuid.UUID) (*AgentAccount, *resty.Response, error)

	// UpdateAccountMany выполняет запрос на массовое изменение счетов контрагента.
	// Принимает контекст, ID контрагента и множество счетов контрагента.
	// Возвращает список изменённых счетов контрагента.
	UpdateAccountMany(ctx context.Context, id uuid.UUID, accounts ...*AgentAccount) (*MetaArray[AgentAccount], *resty.Response, error)

	// GetBySyncID выполняет запрос на получение отдельного документа по syncID.
	// Принимает контекст и syncID документа.
	// Возвращает найденный документ.
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*Counterparty, *resty.Response, error)

	// DeleteBySyncID выполняет запрос на удаление документа по syncID.
	// Принимает контекст и syncID документа.
	// Возвращает true в случае успешного удаления документа.
	DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error)

	// GetNamedFilterList выполняет запрос на получение списка фильтров.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetNamedFilterList(ctx context.Context, params ...*Params) (*List[NamedFilter], *resty.Response, error)

	// GetNamedFilterByID выполняет запрос на получение отдельного фильтра по ID.
	// Принимает контекст и ID фильтра.
	// Возвращает найденный фильтр.
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)

	// GetListAsync выполняет запрос на получение списка контрагентов асинхронно.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает сервис для работы с контекстом асинхронного запроса.
	GetListAsync(ctx context.Context, params ...*Params) (AsyncResultService[List[Counterparty]], *resty.Response, error)

	// GetContactPersonList выполняет запрос на получение списка контактных лиц контрагента.
	// Принимает контекст, ID контрагента и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetContactPersonList(ctx context.Context, id uuid.UUID, params ...*Params) (*List[ContactPerson], *resty.Response, error)

	// GetContactPersonByID выполняет запрос на получение отдельного контактного лица контрагента по ID.
	// Принимает контекст, ID контрагента и ID контактного лица.
	// Возвращает найденное контактное лицо.
	GetContactPersonByID(ctx context.Context, id, contactPersonID uuid.UUID) (*ContactPerson, *resty.Response, error)

	// CreateContactPerson выполняет запрос на создание контактного лица контрагента.
	// Принимает контекст, ID контрагента и контактное лицо контрагента.
	// Возвращает созданное контактное лицо контрагента.
	CreateContactPerson(ctx context.Context, id uuid.UUID, contactPerson *ContactPerson) (*Slice[ContactPerson], *resty.Response, error)

	// UpdateContactPerson выполняет запрос на изменение контактного лица контрагента.
	// Принимает контекст, ID контактного лица контрагента и контактное лицо контрагента.
	// Возвращает изменённое контактное лицо контрагента.
	UpdateContactPerson(ctx context.Context, id, contactPersonID uuid.UUID, contactPerson *ContactPerson) (*ContactPerson, *resty.Response, error)

	// GetNoteList выполняет запрос на получение списка событий контрагента.
	// Принимает контекст, ID контрагента и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetNoteList(ctx context.Context, id uuid.UUID) (*List[Note], *resty.Response, error)

	// GetNoteByID выполняет запрос на получение отдельного события контрагента по ID.
	// Принимает контекст, ID контрагента и ID события.
	// Возвращает найденное событие.
	GetNoteByID(ctx context.Context, id, noteID uuid.UUID) (*Note, *resty.Response, error)

	// CreateNote выполняет запрос на создание события контрагента.
	// Принимает контекст, ID контрагента и событие контрагента.
	// Возвращает созданное событие контрагента.
	CreateNote(ctx context.Context, id uuid.UUID, note *Note) (*MetaArray[Note], *resty.Response, error)

	// UpdateNote выполняет запрос на изменение события контрагента.
	// Принимает контекст, ID события контрагента и событие контрагента.
	// Возвращает изменённое событие контрагента.
	UpdateNote(ctx context.Context, id, noteID uuid.UUID, note *Note) (*Note, *resty.Response, error)

	// DeleteNote выполняет запрос на удаление события контрагента.
	// Принимает контекст и ID события контрагента.
	// Возвращает true в случае успешного удаления события контрагента.
	DeleteNote(ctx context.Context, id, noteID uuid.UUID) (bool, *resty.Response, error)

	// GetFileList выполняет запрос на получение файлов в виде списка.
	// Принимает контекст и ID сущности/документа.
	// Возвращает объект List.
	GetFileList(ctx context.Context, id uuid.UUID) (*List[File], *resty.Response, error)

	// CreateFile выполняет запрос на добавление файла.
	// Принимает контекст, ID сущности/документа и файл.
	// Возвращает список файлов.
	CreateFile(ctx context.Context, id uuid.UUID, file *File) (*Slice[File], *resty.Response, error)

	// UpdateFileMany выполняет запрос на массовое создание и/или изменение файлов сущности/документа.
	// Принимает контекст, ID сущности/документа и множество файлов.
	// Возвращает созданных и/или изменённых файлов.
	UpdateFileMany(ctx context.Context, id uuid.UUID, files ...*File) (*Slice[File], *resty.Response, error)

	// DeleteFile выполняет запрос на удаление файла сущности/документа.
	// Принимает контекст, ID сущности/документа и ID файла.
	// Возвращает true в случае успешного удаления файла.
	DeleteFile(ctx context.Context, id uuid.UUID, fileID uuid.UUID) (bool, *resty.Response, error)

	// DeleteFileMany выполняет запрос на массовое удаление файлов сущности/документа.
	// Принимает контекст, ID сущности/документа и множество файлов.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteFileMany(ctx context.Context, id uuid.UUID, files ...*File) (*DeleteManyResponse, *resty.Response, error)
}

type counterpartyService struct {
	Endpoint
	endpointGetList[Counterparty]
	endpointCreate[Counterparty]
	endpointCreateUpdateMany[Counterparty]
	endpointDeleteMany[Counterparty]
	endpointDelete
	endpointGetByID[Counterparty]
	endpointUpdate[Counterparty]
	endpointMetadata[MetaAttributesStatesSharedTagsWrapper]
	endpointAttributes
	endpointSettings[CounterpartySettings]
	endpointAccounts
	endpointSyncID[Counterparty]
	endpointNamedFilter
	endpointStates
	endpointFiles
}

// NewCounterpartyService принимает [Client] и возвращает сервис для работы с контрагентами.
func NewCounterpartyService(client *Client) CounterpartyService {
	e := NewEndpoint(client, "entity/counterparty")
	return &counterpartyService{
		Endpoint:                 e,
		endpointGetList:          endpointGetList[Counterparty]{e},
		endpointCreate:           endpointCreate[Counterparty]{e},
		endpointCreateUpdateMany: endpointCreateUpdateMany[Counterparty]{e},
		endpointDeleteMany:       endpointDeleteMany[Counterparty]{e},
		endpointDelete:           endpointDelete{e},
		endpointGetByID:          endpointGetByID[Counterparty]{e},
		endpointUpdate:           endpointUpdate[Counterparty]{e},
		endpointMetadata:         endpointMetadata[MetaAttributesStatesSharedTagsWrapper]{e},
		endpointAttributes:       endpointAttributes{e},
		endpointSettings:         endpointSettings[CounterpartySettings]{e},
		endpointAccounts:         endpointAccounts{e},
		endpointSyncID:           endpointSyncID[Counterparty]{e},
		endpointNamedFilter:      endpointNamedFilter{e},
		endpointStates:           endpointStates{e},
		endpointFiles:            endpointFiles{e},
	}
}

func (service *counterpartyService) GetListAsync(ctx context.Context, params ...*Params) (AsyncResultService[List[Counterparty]], *resty.Response, error) {
	return NewRequestBuilder[List[Counterparty]](service.client, service.uri).SetParams(params...).Async(ctx)
}

func (service *counterpartyService) GetContactPersonList(ctx context.Context, id uuid.UUID, params ...*Params) (*List[ContactPerson], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/contactpersons", service.uri, id)
	return NewRequestBuilder[List[ContactPerson]](service.client, path).SetParams(params...).Get(ctx)
}

func (service *counterpartyService) GetContactPersonByID(ctx context.Context, id, contactPersonID uuid.UUID) (*ContactPerson, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/contactpersons/%s", service.uri, id, contactPersonID)
	return NewRequestBuilder[ContactPerson](service.client, path).Get(ctx)
}

func (service *counterpartyService) CreateContactPerson(ctx context.Context, id uuid.UUID, contactPerson *ContactPerson) (*Slice[ContactPerson], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/contactpersons", service.uri, id)
	return NewRequestBuilder[Slice[ContactPerson]](service.client, path).Post(ctx, contactPerson)
}

func (service *counterpartyService) UpdateContactPerson(ctx context.Context, id, contactPersonID uuid.UUID, contactPerson *ContactPerson) (*ContactPerson, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/contactpersons/%s", service.uri, id, contactPersonID)
	return NewRequestBuilder[ContactPerson](service.client, path).Put(ctx, contactPerson)
}

func (service *counterpartyService) GetNoteList(ctx context.Context, id uuid.UUID) (*List[Note], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes", service.uri, id)
	return NewRequestBuilder[List[Note]](service.client, path).Get(ctx)
}

func (service *counterpartyService) GetNoteByID(ctx context.Context, id, noteID uuid.UUID) (*Note, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes/%s", service.uri, id, noteID)
	return NewRequestBuilder[Note](service.client, path).Get(ctx)
}

func (service *counterpartyService) CreateNote(ctx context.Context, id uuid.UUID, note *Note) (*MetaArray[Note], *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes", service.uri, id)
	return NewRequestBuilder[MetaArray[Note]](service.client, path).Post(ctx, note)
}

func (service *counterpartyService) UpdateNote(ctx context.Context, id, noteID uuid.UUID, note *Note) (*Note, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes/%s", service.uri, id, noteID)
	return NewRequestBuilder[Note](service.client, path).Put(ctx, note)
}

func (service *counterpartyService) DeleteNote(ctx context.Context, id, noteID uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/%s/notes/%s", service.uri, id, noteID)
	return NewRequestBuilder[any](service.client, path).Delete(ctx)
}
