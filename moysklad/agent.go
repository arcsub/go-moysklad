package moysklad

import (
	"github.com/goccy/go-json"
	"github.com/google/uuid"
)

// Agent представляет структуру с общими полями типов Counterparty (Контрагент) и Organization (Юрлицо)
type Agent struct {
	KPP               *string                  `json:"kpp,omitempty"`
	Fax               *string                  `json:"fax,omitempty"`
	Updated           *Timestamp               `json:"updated,omitempty"`
	Shared            *bool                    `json:"shared,omitempty"`
	Name              *string                  `json:"name,omitempty"`
	Meta              *Meta                    `json:"meta,omitempty"`
	ActualAddressFull *Address                 `json:"actualAddressFull,omitempty"`
	LegalAddress      *string                  `json:"legalAddress,omitempty"`
	CertificateDate   *Timestamp               `json:"certificateDate,omitempty"`
	ID                *uuid.UUID               `json:"id,omitempty"`
	OGRN              *string                  `json:"ogrn,omitempty"`
	OKPO              *string                  `json:"okpo,omitempty"`
	ExternalCode      *string                  `json:"externalCode,omitempty"`
	LegalTitle        *string                  `json:"legalTitle,omitempty"`
	LegalAddressFull  *Address                 `json:"legalAddressFull,omitempty"`
	INN               *string                  `json:"inn,omitempty"`
	Phone             *string                  `json:"phone,omitempty"`
	SyncID            *uuid.UUID               `json:"syncId,omitempty"`
	CertificateNumber *string                  `json:"certificateNumber,omitempty"`
	Accounts          *MetaArray[AgentAccount] `json:"accounts,omitempty"`
	Created           *Timestamp               `json:"created,omitempty"`
	Archived          *bool                    `json:"archived,omitempty"`
	Description       *string                  `json:"description,omitempty"`
	Group             *Group                   `json:"group,omitempty"`
	OGRNIP            *string                  `json:"ogrnip,omitempty"`
	BonusProgram      *BonusProgram            `json:"bonusProgram,omitempty"`
	Code              *string                  `json:"code,omitempty"`
	AccountID         *uuid.UUID               `json:"accountId,omitempty"`
	ActualAddress     *string                  `json:"actualAddress,omitempty"`
	BonusPoints       *int                     `json:"bonusPoints,omitempty"`
	Owner             *Employee                `json:"owner,omitempty"`
	Email             *string                  `json:"email,omitempty"`
	CompanyType       CompanyType              `json:"companyType,omitempty"`
	Attributes        Slice[AttributeValue]    `json:"attributes,omitempty"`
	raw               []byte
}

// Clean возвращает сущность с единственным заполненным полем Meta
func (agent Agent) Clean() *Agent {
	return &Agent{Meta: agent.Meta}
}

func (agent Agent) GetKPP() string {
	return Deref(agent.KPP)
}

func (agent Agent) GetFax() string {
	return Deref(agent.Fax)
}

func (agent Agent) GetUpdated() Timestamp {
	return Deref(agent.Updated)
}

func (agent Agent) GetShared() bool {
	return Deref(agent.Shared)
}

func (agent Agent) GetName() string {
	return Deref(agent.Name)
}

func (agent Agent) GetMeta() Meta {
	return Deref(agent.Meta)
}

func (agent Agent) GetActualAddressFull() Address {
	return Deref(agent.ActualAddressFull)
}

func (agent Agent) GetLegalAddress() string {
	return Deref(agent.LegalAddress)
}

func (agent Agent) GetCertificateDate() Timestamp {
	return Deref(agent.CertificateDate)
}

func (agent Agent) GetID() uuid.UUID {
	return Deref(agent.ID)
}

func (agent Agent) GetOGRN() string {
	return Deref(agent.OGRN)
}

func (agent Agent) GetOKPO() string {
	return Deref(agent.OKPO)
}

func (agent Agent) GetExternalCode() string {
	return Deref(agent.ExternalCode)
}

func (agent Agent) GetLegalTitle() string {
	return Deref(agent.LegalTitle)
}

func (agent Agent) GetLegalAddressFull() Address {
	return Deref(agent.LegalAddressFull)
}

func (agent Agent) GetINN() string {
	return Deref(agent.INN)
}

func (agent Agent) GetPhone() string {
	return Deref(agent.Phone)
}

func (agent Agent) GetSyncID() uuid.UUID {
	return Deref(agent.SyncID)
}

func (agent Agent) GetCertificateNumber() string {
	return Deref(agent.CertificateNumber)
}

func (agent Agent) GetAccounts() MetaArray[AgentAccount] {
	return Deref(agent.Accounts)
}

func (agent Agent) GetCreated() Timestamp {
	return Deref(agent.Created)
}

func (agent Agent) GetArchived() bool {
	return Deref(agent.Archived)
}

func (agent Agent) GetDescription() string {
	return Deref(agent.Description)
}

func (agent Agent) GetGroup() Group {
	return Deref(agent.Group)
}

func (agent Agent) GetOGRNIP() string {
	return Deref(agent.OGRNIP)
}

func (agent Agent) GetBonusProgram() BonusProgram {
	return Deref(agent.BonusProgram)
}

func (agent Agent) GetCode() string {
	return Deref(agent.Code)
}

func (agent Agent) GetAccountID() uuid.UUID {
	return Deref(agent.AccountID)
}

func (agent Agent) GetActualAddress() string {
	return Deref(agent.ActualAddress)
}

func (agent Agent) GetBonusPoints() int {
	return Deref(agent.BonusPoints)
}

func (agent Agent) GetOwner() Employee {
	return Deref(agent.Owner)
}

func (agent Agent) GetEmail() string {
	return Deref(agent.Email)
}

func (agent Agent) GetCompanyType() CompanyType {
	return agent.CompanyType
}

func (agent Agent) GetAttributes() Slice[AttributeValue] {
	return agent.Attributes
}

func (agent *Agent) String() string {
	return Stringify(agent.Meta)
}

// MetaType реализует интерфейс MetaTyper
func (agent Agent) MetaType() MetaType {
	return agent.Meta.GetType()
}

// Raw реализует интерфейс RawMetaTyper
func (agent Agent) Raw() []byte {
	return agent.raw
}

// UnmarshalJSON реализует интерфейс json.Unmarshaler
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

// AsCounterparty десериализует объект в тип *Counterparty
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (agent *Agent) AsCounterparty() *Counterparty {
	return UnmarshalAsType[Counterparty](agent)
}

// AsOrganization десериализует объект в тип *Organization
// Метод гарантирует преобразование в необходимый тип только при идентичных MetaType.
// Возвращает nil в случае неудачи.
func (agent *Agent) AsOrganization() *Organization {
	return UnmarshalAsType[Organization](agent)
}
