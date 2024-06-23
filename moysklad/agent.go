package moysklad

import (
	"github.com/goccy/go-json"
	"github.com/google/uuid"
)

// Agent представляет объект с общими полями типов [Counterparty] (Контрагент) и [Organization] (Юрлицо).
type Agent struct {
	KPP               *string                  `json:"kpp,omitempty"`               // КПП
	Fax               *string                  `json:"fax,omitempty"`               // Номер факса
	Updated           *Timestamp               `json:"updated,omitempty"`           // Момент последнего обновления Контрагента/Юрлица
	Shared            *bool                    `json:"shared,omitempty"`            // Общий доступ [Обязательное при ответе]
	Name              *string                  `json:"name,omitempty"`              // Наименование Контрагента/Юрлица
	Meta              *Meta                    `json:"meta,omitempty"`              // Метаданные Контрагента/Юрлица
	ActualAddressFull *Address                 `json:"actualAddressFull,omitempty"` // Фактический адрес Контрагента с детализацией по отдельным полям
	LegalAddress      *string                  `json:"legalAddress,omitempty"`      // Юридический адрес Контрагента/Юрлица
	CertificateDate   *Timestamp               `json:"certificateDate,omitempty"`   // Дата свидетельства
	ID                *uuid.UUID               `json:"id,omitempty"`                // ID Контрагента/Юрлица
	OGRN              *string                  `json:"ogrn,omitempty"`              // ОГРН
	OKPO              *string                  `json:"okpo,omitempty"`              // ОКПО
	ExternalCode      *string                  `json:"externalCode,omitempty"`      // Внешний код Контрагента/Юрлица
	LegalTitle        *string                  `json:"legalTitle,omitempty"`        // Полное наименование для Контрагента типа [Юридическое лицо]. Игнорируется для Контрагентов типа [Индивидуальный предприниматель, Физическое лицо], если передано одно из значений для ФИО и формируется автоматически на основе получаемых ФИО Контрагента
	LegalAddressFull  *Address                 `json:"legalAddressFull,omitempty"`  // Юридический адрес Контрагента/Юрлица с детализацией по отдельным полям
	INN               *string                  `json:"inn,omitempty"`               // ИНН
	Phone             *string                  `json:"phone,omitempty"`             // Номер городского телефона
	SyncID            *uuid.UUID               `json:"syncId,omitempty"`            // ID синхронизации
	CertificateNumber *string                  `json:"certificateNumber,omitempty"` // Номер свидетельства
	Accounts          *MetaArray[AgentAccount] `json:"accounts,omitempty"`          // Массив счетов Контрагента/Юрлица
	Created           *Timestamp               `json:"created,omitempty"`           // Момент создания
	Archived          *bool                    `json:"archived,omitempty"`          // Добавлен ли Контрагент/Юрлицо в архив
	Description       *string                  `json:"description,omitempty"`       // Комментарий к Контрагенту/Юрлицу
	Group             *Group                   `json:"group,omitempty"`             // Отдел сотрудника
	OGRNIP            *string                  `json:"ogrnip,omitempty"`            // ОГРНИП
	BonusProgram      *BonusProgram            `json:"bonusProgram,omitempty"`      // Метаданные активной Бонусной программы
	Code              *string                  `json:"code,omitempty"`              // Код Контрагента/Юрлица
	AccountID         *uuid.UUID               `json:"accountId,omitempty"`         // ID учетной записи
	ActualAddress     *string                  `json:"actualAddress,omitempty"`     // Фактический адрес Контрагента/Юрлица
	BonusPoints       *int                     `json:"bonusPoints,omitempty"`       // Бонусные баллы по активной бонусной программе
	Owner             *Employee                `json:"owner,omitempty"`             // Метаданные владельца (Сотрудника)
	Email             *string                  `json:"email,omitempty"`             // Адрес электронной почты
	CompanyType       CompanyType              `json:"companyType,omitempty"`       // Тип Контрагента/Юрлица. В зависимости от значения данного поля набор выводимых реквизитов контрагента может меняться.
	Attributes        Slice[Attribute]         `json:"attributes,omitempty"`        // Список метаданных доп. полей
	raw               []byte
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (agent Agent) Clean() *Agent {
	return &Agent{Meta: agent.Meta}
}

// GetKPP возвращает КПП.
func (agent Agent) GetKPP() string {
	return Deref(agent.KPP)
}

// GetFax возвращает Номер факса.
func (agent Agent) GetFax() string {
	return Deref(agent.Fax)
}

// GetUpdated возвращает Момент последнего обновления Контрагента/Юрлица.
func (agent Agent) GetUpdated() Timestamp {
	return Deref(agent.Updated)
}

// GetShared возвращает Общий доступ.
func (agent Agent) GetShared() bool {
	return Deref(agent.Shared)
}

// GetName возвращает Наименование Контрагента/Юрлица.
func (agent Agent) GetName() string {
	return Deref(agent.Name)
}

// GetMeta возвращает Метаданные Контрагента/Юрлица.
func (agent Agent) GetMeta() Meta {
	return Deref(agent.Meta)
}

// GetActualAddressFull возвращает Фактический адрес Контрагента с детализацией по отдельным полям.
func (agent Agent) GetActualAddressFull() Address {
	return Deref(agent.ActualAddressFull)
}

// GetLegalAddress возвращает Юридический адрес Контрагента/Юрлица.
func (agent Agent) GetLegalAddress() string {
	return Deref(agent.LegalAddress)
}

// GetCertificateDate возвращает Дату свидетельства.
func (agent Agent) GetCertificateDate() Timestamp {
	return Deref(agent.CertificateDate)
}

// GetID возвращает ID Контрагента/Юрлица.
func (agent Agent) GetID() uuid.UUID {
	return Deref(agent.ID)
}

// GetOGRN возвращает ОГРН.
func (agent Agent) GetOGRN() string {
	return Deref(agent.OGRN)
}

// GetOKPO возвращает ОКПО.
func (agent Agent) GetOKPO() string {
	return Deref(agent.OKPO)
}

// GetExternalCode возвращает Внешний код Контрагента/Юрлица.
func (agent Agent) GetExternalCode() string {
	return Deref(agent.ExternalCode)
}

// GetLegalTitle возвращает Полное наименование для Контрагента типа [Юридическое лицо].
func (agent Agent) GetLegalTitle() string {
	return Deref(agent.LegalTitle)
}

// GetLegalAddressFull возвращает Юридический адрес Контрагента/Юрлица с детализацией по отдельным полям.
func (agent Agent) GetLegalAddressFull() Address {
	return Deref(agent.LegalAddressFull)
}

// GetINN возвращает ИНН.
func (agent Agent) GetINN() string {
	return Deref(agent.INN)
}

// GetPhone возвращает Номер городского телефона.
func (agent Agent) GetPhone() string {
	return Deref(agent.Phone)
}

// GetSyncID возвращает ID синхронизации.
func (agent Agent) GetSyncID() uuid.UUID {
	return Deref(agent.SyncID)
}

// GetCertificateNumber возвращает Номер свидетельства.
func (agent Agent) GetCertificateNumber() string {
	return Deref(agent.CertificateNumber)
}

// GetAccounts возвращает Массив счетов Контрагента/Юрлица.
func (agent Agent) GetAccounts() MetaArray[AgentAccount] {
	return Deref(agent.Accounts)
}

// GetCreated возвращает Момент создания.
func (agent Agent) GetCreated() Timestamp {
	return Deref(agent.Created)
}

// GetArchived возвращает true, если Контрагент/Юрлицо добавлен в архив.
func (agent Agent) GetArchived() bool {
	return Deref(agent.Archived)
}

// GetDescription возвращает Комментарий к Контрагенту/Юрлицу.
func (agent Agent) GetDescription() string {
	return Deref(agent.Description)
}

// GetGroup возвращает Отдел сотрудника.
func (agent Agent) GetGroup() Group {
	return Deref(agent.Group)
}

// GetOGRNIP возвращает ОГРНИП.
func (agent Agent) GetOGRNIP() string {
	return Deref(agent.OGRNIP)
}

// GetBonusProgram возвращает Метаданные активной Бонусной программы.
func (agent Agent) GetBonusProgram() BonusProgram {
	return Deref(agent.BonusProgram)
}

// GetCode возвращает Код Контрагента/Юрлица.
func (agent Agent) GetCode() string {
	return Deref(agent.Code)
}

// GetAccountID возвращает ID учетной записи.
func (agent Agent) GetAccountID() uuid.UUID {
	return Deref(agent.AccountID)
}

// GetActualAddress возвращает Фактический адрес Контрагента/Юрлица.
func (agent Agent) GetActualAddress() string {
	return Deref(agent.ActualAddress)
}

// GetBonusPoints возвращает Бонусные баллы по активной бонусной программе.
func (agent Agent) GetBonusPoints() int {
	return Deref(agent.BonusPoints)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (agent Agent) GetOwner() Employee {
	return Deref(agent.Owner)
}

// GetEmail возвращает Адрес электронной почты.
func (agent Agent) GetEmail() string {
	return Deref(agent.Email)
}

// GetCompanyType возвращает Тип Контрагента/Юрлица.
func (agent Agent) GetCompanyType() CompanyType {
	return agent.CompanyType
}

// GetAttributes возвращает Список метаданных доп. полей.
func (agent Agent) GetAttributes() Slice[Attribute] {
	return agent.Attributes
}

// String реализует интерфейс [fmt.Stringer].
func (agent *Agent) String() string {
	return Stringify(agent.Meta)
}

// MetaType возвращает тип сущности.
func (agent Agent) MetaType() MetaType {
	return agent.GetMeta().GetType()
}

// Raw реализует интерфейс RawMetaTyper.
func (agent Agent) Raw() []byte {
	return agent.raw
}

// UnmarshalJSON реализует интерфейс [json.Unmarshaler].
func (agent *Agent) UnmarshalJSON(data []byte) error {
	type alias Agent
	var t alias
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	t.raw = data
	*agent = Agent(t)
	return nil
}

// AsCounterparty пытается привести объект к типу [Counterparty].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает:
//   - указатель на [Counterparty].
//   - nil в случае неудачи.
func (agent *Agent) AsCounterparty() *Counterparty {
	return UnmarshalAsType[Counterparty](agent)
}

// AsOrganization пытается привести объект к типу [Organization].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает:
//   - указатель на [Organization].
//   - nil в случае неудачи.
func (agent *Agent) AsOrganization() *Organization {
	return UnmarshalAsType[Organization](agent)
}
