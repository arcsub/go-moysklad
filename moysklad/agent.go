package moysklad

import (
	"github.com/goccy/go-json"
	"github.com/google/uuid"
)

// Agent представляет объект с общими полями объектов:
//   - [Counterparty] (Контрагент)
//   - [Organization] (Юрлицо)
//   - [Employee] (Сотрудник)
type Agent struct {
	Meta         *Meta            `json:"meta,omitempty"`         // Метаданные Контрагента/Юрлица/Сотрудника
	AccountID    *uuid.UUID       `json:"accountId,omitempty"`    // ID учётной записи
	Updated      *Timestamp       `json:"updated,omitempty"`      // Момент последнего обновления Контрагента/Юрлица/Сотрудника
	Archived     *bool            `json:"archived,omitempty"`     // Добавлен ли Контрагент/Юрлицо/Сотрудник в архив
	Phone        *string          `json:"phone,omitempty"`        // Номер телефона
	ID           *uuid.UUID       `json:"id,omitempty"`           // ID Контрагента/Юрлица/Сотрудника
	Attributes   Slice[Attribute] `json:"attributes,omitempty"`   // Список метаданных доп. полей
	Shared       *bool            `json:"shared,omitempty"`       // Общий доступ
	Group        *Group           `json:"group,omitempty"`        // Отдел сотрудника
	Name         *string          `json:"name,omitempty"`         // Наименование Контрагента/Юрлица/Сотрудника
	Created      *Timestamp       `json:"created,omitempty"`      // Момент создания
	Code         *string          `json:"code,omitempty"`         // Код Контрагента/Юрлица/Сотрудника
	Owner        *Employee        `json:"owner,omitempty"`        // Метаданные владельца (Сотрудника)
	Description  *string          `json:"description,omitempty"`  // Комментарий к Контрагенту/Юрлицу/Сотруднику
	INN          *string          `json:"inn,omitempty"`          // ИНН
	Email        *string          `json:"email,omitempty"`        // Адрес электронной почты
	ExternalCode *string          `json:"externalCode,omitempty"` // Внешний код Контрагента/Юрлица/Сотрудника
	raw          []byte           // сырые данные для последующей конвертации в нужный тип
}

// AgentCounterpartyOrganizationInterface описывает метод, который возвращает *Agent
//
// Интерфейс должны реализовывать: [Counterparty] и [Organization].
type AgentCounterpartyOrganizationInterface interface {
	asCOAgent() *Agent
}

// AgentCounterpartyEmployeeInterface описывает метод, который возвращает *Agent
//
// Интерфейс должны реализовывать: [Counterparty] и [Employee].
type AgentCounterpartyEmployeeInterface interface {
	asCEAgent() *Agent
}

// AgentInterface описывает метод, который возвращает *Agent
//
// Интерфейс должны реализовывать: [Counterparty], [Organization] и [Employee].
type AgentInterface interface {
	asAgent() *Agent
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (agent Agent) Clean() *Agent {
	if agent.Meta == nil {
		return nil
	}
	return &Agent{Meta: agent.Meta}
}

// IsCounterparty возвращает true, если объект является контрагентом [Counterparty].
func (agent Agent) IsCounterparty() bool {
	return agent.GetMeta().GetType() == MetaTypeCounterparty
}

// IsOrganization возвращает true, если объект является юрлицом [Organization].
func (agent Agent) IsOrganization() bool {
	return agent.GetMeta().GetType() == MetaTypeOrganization
}

// IsEmployee возвращает true, если объект является сотрудником [Employee].
func (agent Agent) IsEmployee() bool {
	return agent.GetMeta().GetType() == MetaTypeEmployee
}

// GetUpdated возвращает Момент последнего обновления Контрагента/Юрлица/Сотрудника.
func (agent Agent) GetUpdated() Timestamp {
	return Deref(agent.Updated)
}

// GetShared возвращает флаг Общего доступа.
func (agent Agent) GetShared() bool {
	return Deref(agent.Shared)
}

// GetName возвращает Наименование Контрагента/Юрлица/Сотрудника.
func (agent Agent) GetName() string {
	return Deref(agent.Name)
}

// GetMeta возвращает Метаданные Контрагента/Юрлица/Сотрудника.
func (agent Agent) GetMeta() Meta {
	return Deref(agent.Meta)
}

// GetID возвращает ID Контрагента/Юрлица/Сотрудника.
func (agent Agent) GetID() uuid.UUID {
	return Deref(agent.ID)
}

// GetExternalCode возвращает Внешний код Контрагента/Юрлица/Сотрудника.
func (agent Agent) GetExternalCode() string {
	return Deref(agent.ExternalCode)
}

// GetINN возвращает ИНН.
func (agent Agent) GetINN() string {
	return Deref(agent.INN)
}

// GetPhone возвращает Номер городского телефона.
func (agent Agent) GetPhone() string {
	return Deref(agent.Phone)
}

// GetCreated возвращает Дату создания.
func (agent Agent) GetCreated() Timestamp {
	return Deref(agent.Created)
}

// GetArchived возвращает true, если Контрагент/Юрлицо/Сотрудник добавлен в архив.
func (agent Agent) GetArchived() bool {
	return Deref(agent.Archived)
}

// GetDescription возвращает Комментарий к Контрагенту/Юрлицу/Сотруднику.
func (agent Agent) GetDescription() string {
	return Deref(agent.Description)
}

// GetGroup возвращает Отдел сотрудника.
func (agent Agent) GetGroup() Group {
	return Deref(agent.Group)
}

// GetCode возвращает Код Контрагента/Юрлица/Сотрудника.
func (agent Agent) GetCode() string {
	return Deref(agent.Code)
}

// GetAccountID возвращает ID учётной записи.
func (agent Agent) GetAccountID() uuid.UUID {
	return Deref(agent.AccountID)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (agent Agent) GetOwner() Employee {
	return Deref(agent.Owner)
}

// GetEmail возвращает Адрес электронной почты.
func (agent Agent) GetEmail() string {
	return Deref(agent.Email)
}

// GetAttributes возвращает Список метаданных доп. полей.
func (agent Agent) GetAttributes() Slice[Attribute] {
	return agent.Attributes
}

// String реализует интерфейс [fmt.Stringer].
func (agent *Agent) String() string {
	return Stringify(agent.Meta)
}

// MetaType возвращает код сущности.
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
// Возвращает [Counterparty] или nil в случае неудачи.
func (agent *Agent) AsCounterparty() *Counterparty {
	return UnmarshalAsType[Counterparty](agent)
}

// AsOrganization пытается привести объект к типу [Organization].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [Organization] или nil в случае неудачи.
func (agent *Agent) AsOrganization() *Organization {
	return UnmarshalAsType[Organization](agent)
}

// AsEmployee пытается привести объект к типу [Employee].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [Employee] или nil в случае неудачи.
func (agent *Agent) AsEmployee() *Employee {
	return UnmarshalAsType[Employee](agent)
}

// CompanyType Тип Контрагента.
//
// Возможные значения:
//   - CompanyLegal        – Юридическое лицо
//   - CompanyEntrepreneur – Индивидуальный предприниматель
//   - CompanyIndividual   – Физическое лицо
type CompanyType string

const (
	CompanyLegal        CompanyType = "legal"        // Юридическое лицо
	CompanyEntrepreneur CompanyType = "entrepreneur" // Индивидуальный предприниматель
	CompanyIndividual   CompanyType = "individual"   // Физическое лицо
)
