package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// RetailShift Розничная смена.
// Ключевое слово: retailshift
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-roznichnaq-smena
type RetailShift struct {
	Organization        *Organization          `json:"organization,omitempty"`
	Shared              *bool                  `json:"shared,omitempty"`
	AgentAccount        *AgentAccount          `json:"agentAccount,omitempty"`
	VatIncluded         *bool                  `json:"vatIncluded,omitempty"`
	BankCommission      *float64               `json:"bankComission,omitempty"`
	BankPercent         *float64               `json:"bankPercent,omitempty"`
	Name                *string                `json:"name,omitempty"`
	CloseDate           *Timestamp             `json:"closeDate,omitempty"`
	Contract            *Contract              `json:"contract,omitempty"`
	Created             *Timestamp             `json:"created,omitempty"`
	Deleted             *Timestamp             `json:"deleted,omitempty"`
	Description         *string                `json:"description,omitempty"`
	ExternalCode        *string                `json:"externalCode,omitempty"`
	Files               *MetaArray[File]       `json:"files,omitempty"`
	Group               *Group                 `json:"group,omitempty"`
	ID                  *uuid.UUID             `json:"id,omitempty"`
	Meta                *Meta                  `json:"meta,omitempty"`
	AccountID           *uuid.UUID             `json:"accountId,omitempty"`
	Cheque              *Cheque                `json:"cheque,omitempty"`
	Acquire             *Counterparty          `json:"acquire,omitempty"`
	Moment              *Timestamp             `json:"moment,omitempty"`
	OrganizationAccount *AgentAccount          `json:"organizationAccount,omitempty"`
	Owner               *Employee              `json:"owner,omitempty"`
	Printed             *bool                  `json:"printed,omitempty"`
	ProceedsCash        *float64               `json:"proceedsCash,omitempty"`
	ProceedsNoCash      *float64               `json:"proceedsNoCash,omitempty"`
	Published           *bool                  `json:"published,omitempty"`
	QRAcquire           *Counterparty          `json:"qrAcquire,omitempty"`
	QRBankCommission    *float64               `json:"qrBankComission,omitempty"`
	QRBankPercent       *float64               `json:"qrBankPercent,omitempty"`
	ReceivedCash        *float64               `json:"receivedCash,omitempty"`
	ReceivedNoCash      *float64               `json:"receivedNoCash,omitempty"`
	RetailStore         *RetailStore           `json:"retailStore,omitempty"`
	Store               *Store                 `json:"store,omitempty"`
	SyncID              *uuid.UUID             `json:"syncId,omitempty"`
	Updated             *Timestamp             `json:"updated,omitempty"`
	VatEnabled          *bool                  `json:"vatEnabled,omitempty"`
	PaymentOperations   Slice[Payment]         `json:"paymentOperations,omitempty"`
	Operations          Slice[RetailOperation] `json:"operations,omitempty"`
	Attributes          Slice[AttributeValue]  `json:"attributes,omitempty"`
}

func (retailShift RetailShift) Clean() *RetailShift {
	return &RetailShift{Meta: retailShift.Meta}
}

func (retailShift RetailShift) GetOrganization() Organization {
	return Deref(retailShift.Organization)
}

func (retailShift RetailShift) GetShared() bool {
	return Deref(retailShift.Shared)
}

func (retailShift RetailShift) GetAgentAccount() AgentAccount {
	return Deref(retailShift.AgentAccount)
}

func (retailShift RetailShift) GetVatIncluded() bool {
	return Deref(retailShift.VatIncluded)
}

func (retailShift RetailShift) GetBankCommission() float64 {
	return Deref(retailShift.BankCommission)
}

func (retailShift RetailShift) GetBankPercent() float64 {
	return Deref(retailShift.BankPercent)
}

func (retailShift RetailShift) GetName() string {
	return Deref(retailShift.Name)
}

func (retailShift RetailShift) GetCloseDate() Timestamp {
	return Deref(retailShift.CloseDate)
}

func (retailShift RetailShift) GetContract() Contract {
	return Deref(retailShift.Contract)
}

func (retailShift RetailShift) GetCreated() Timestamp {
	return Deref(retailShift.Created)
}

func (retailShift RetailShift) GetDeleted() Timestamp {
	return Deref(retailShift.Deleted)
}

func (retailShift RetailShift) GetDescription() string {
	return Deref(retailShift.Description)
}

func (retailShift RetailShift) GetExternalCode() string {
	return Deref(retailShift.ExternalCode)
}

func (retailShift RetailShift) GetFiles() MetaArray[File] {
	return Deref(retailShift.Files)
}

func (retailShift RetailShift) GetGroup() Group {
	return Deref(retailShift.Group)
}

func (retailShift RetailShift) GetID() uuid.UUID {
	return Deref(retailShift.ID)
}

func (retailShift RetailShift) GetMeta() Meta {
	return Deref(retailShift.Meta)
}

func (retailShift RetailShift) GetAccountID() uuid.UUID {
	return Deref(retailShift.AccountID)
}

func (retailShift RetailShift) GetCheque() Cheque {
	return Deref(retailShift.Cheque)
}

func (retailShift RetailShift) GetAcquire() Counterparty {
	return Deref(retailShift.Acquire)
}

func (retailShift RetailShift) GetMoment() Timestamp {
	return Deref(retailShift.Moment)
}

func (retailShift RetailShift) GetOrganizationAccount() AgentAccount {
	return Deref(retailShift.OrganizationAccount)
}

func (retailShift RetailShift) GetOwner() Employee {
	return Deref(retailShift.Owner)
}

func (retailShift RetailShift) GetPrinted() bool {
	return Deref(retailShift.Printed)
}

func (retailShift RetailShift) GetProceedsCash() float64 {
	return Deref(retailShift.ProceedsCash)
}

func (retailShift RetailShift) GetProceedsNoCash() float64 {
	return Deref(retailShift.ProceedsNoCash)
}

func (retailShift RetailShift) GetPublished() bool {
	return Deref(retailShift.Published)
}

func (retailShift RetailShift) GetQRAcquire() Counterparty {
	return Deref(retailShift.QRAcquire)
}

func (retailShift RetailShift) GetQRBankCommission() float64 {
	return Deref(retailShift.QRBankCommission)
}

func (retailShift RetailShift) GetQRBankPercent() float64 {
	return Deref(retailShift.QRBankPercent)
}

func (retailShift RetailShift) GetReceivedCash() float64 {
	return Deref(retailShift.ReceivedCash)
}

func (retailShift RetailShift) GetReceivedNoCash() float64 {
	return Deref(retailShift.ReceivedNoCash)
}

func (retailShift RetailShift) GetRetailStore() RetailStore {
	return Deref(retailShift.RetailStore)
}

func (retailShift RetailShift) GetStore() Store {
	return Deref(retailShift.Store)
}

func (retailShift RetailShift) GetSyncID() uuid.UUID {
	return Deref(retailShift.SyncID)
}

func (retailShift RetailShift) GetUpdated() Timestamp {
	return Deref(retailShift.Updated)
}

func (retailShift RetailShift) GetVatEnabled() bool {
	return Deref(retailShift.VatEnabled)
}

func (retailShift RetailShift) GetPayments() Slice[Payment] {
	return retailShift.PaymentOperations
}

func (retailShift RetailShift) GetOperations() Slice[RetailOperation] {
	return retailShift.Operations
}

func (retailShift RetailShift) GetAttributes() Slice[AttributeValue] {
	return retailShift.Attributes
}

func (retailShift *RetailShift) SetOrganization(organization *Organization) *RetailShift {
	retailShift.Organization = organization.Clean()
	return retailShift
}

func (retailShift *RetailShift) SetShared(shared bool) *RetailShift {
	retailShift.Shared = &shared
	return retailShift
}

func (retailShift *RetailShift) SetVatIncluded(vatIncluded bool) *RetailShift {
	retailShift.VatIncluded = &vatIncluded
	return retailShift
}

func (retailShift *RetailShift) SetBankCommission(bankCommission float64) *RetailShift {
	retailShift.BankCommission = &bankCommission
	return retailShift
}

func (retailShift *RetailShift) SetBankPercent(bankPercent float64) *RetailShift {
	retailShift.BankPercent = &bankPercent
	return retailShift
}

func (retailShift *RetailShift) SetName(name string) *RetailShift {
	retailShift.Name = &name
	return retailShift
}

func (retailShift *RetailShift) SetCloseDate(closeDate *Timestamp) *RetailShift {
	retailShift.CloseDate = closeDate
	return retailShift
}

func (retailShift *RetailShift) SetDescription(description string) *RetailShift {
	retailShift.Description = &description
	return retailShift
}

func (retailShift *RetailShift) SetExternalCode(externalCode string) *RetailShift {
	retailShift.ExternalCode = &externalCode
	return retailShift
}

func (retailShift *RetailShift) SetFiles(files Slice[File]) *RetailShift {
	retailShift.Files = NewMetaArrayRows(files)
	return retailShift
}

func (retailShift *RetailShift) SetMeta(meta *Meta) *RetailShift {
	retailShift.Meta = meta
	return retailShift
}

func (retailShift *RetailShift) SetAcquire(acquire *Counterparty) *RetailShift {
	retailShift.Acquire = acquire.Clean()
	return retailShift
}

func (retailShift *RetailShift) SetMoment(moment *Timestamp) *RetailShift {
	retailShift.Moment = moment
	return retailShift
}

func (retailShift *RetailShift) SetOrganizationAccount(organizationAccount *AgentAccount) *RetailShift {
	retailShift.OrganizationAccount = organizationAccount.Clean()
	return retailShift
}

func (retailShift *RetailShift) SetOwner(owner *Employee) *RetailShift {
	retailShift.Owner = owner.Clean()
	return retailShift
}

func (retailShift *RetailShift) SetQRAcquire(qrAcquire *Counterparty) *RetailShift {
	retailShift.QRAcquire = qrAcquire.Clean()
	return retailShift
}

func (retailShift *RetailShift) SetQRBankCommission(qrBankCommission float64) *RetailShift {
	retailShift.QRBankCommission = &qrBankCommission
	return retailShift
}

func (retailShift *RetailShift) SetQRBankPercent(qrBankPercent float64) *RetailShift {
	retailShift.QRBankPercent = &qrBankPercent
	return retailShift
}

func (retailShift *RetailShift) SetRetailStore(retailStore *RetailStore) *RetailShift {
	retailShift.RetailStore = retailStore.Clean()
	return retailShift
}

func (retailShift *RetailShift) SetStore(store *Store) *RetailShift {
	retailShift.Store = store.Clean()
	return retailShift
}

func (retailShift *RetailShift) SetSyncID(syncID uuid.UUID) *RetailShift {
	retailShift.SyncID = &syncID
	return retailShift
}

func (retailShift *RetailShift) SetAttributes(attributes Slice[AttributeValue]) *RetailShift {
	retailShift.Attributes = attributes
	return retailShift
}

func (retailShift RetailShift) String() string {
	return Stringify(retailShift)
}

func (retailShift RetailShift) MetaType() MetaType {
	return MetaTypeRetailShift
}

// Cheque Информация о смене ККТ
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-roznichnaq-smena-roznichnye-smeny-informaciq-o-smene-kkt
type Cheque struct {
	Start ChequeStart `json:"start,omitempty"` // Информация об открытии смены
	End   ChequeEnd   `json:"end,omitempty"`   // Информация о закрытии смены
}

func (cheque Cheque) String() string {
	return Stringify(cheque)
}

// ChequeStart Информация об открытии смены ККТ
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-roznichnaq-smena-roznichnye-smeny-informaciq-ob-otkrytii-smeny-kkt
type ChequeStart struct {
	Time            Timestamp `json:"time,omitempty"`
	FnNumber        string    `json:"fnNumber,omitempty"`
	KKTRegNumber    string    `json:"kktRegNumber,omitempty"`
	FiscalDocSign   string    `json:"fiscalDocSign,omitempty"`
	ShiftNumber     string    `json:"shiftNumber,omitempty"`
	FiscalDocNumber string    `json:"fiscalDocNumber,omitempty"`
}

func (ChequeStart ChequeStart) String() string {
	return Stringify(ChequeStart)
}

// ChequeEnd Информация о закрытии смены ККТ
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-roznichnaq-smena-roznichnye-smeny-informaciq-o-zakrytii-smeny-kkt
type ChequeEnd struct {
	ChequeStart
	ChequesTotal    float64 `json:"chequesTotal,omitempty"`    // Количество чеков за смену
	FiscalDocsTotal float64 `json:"fiscalDocsTotal,omitempty"` // Количество фискальных документов за смену
}

func (chequeEnd ChequeEnd) String() string {
	return Stringify(chequeEnd)
}

// RetailShiftService
// Сервис для работы с розничными сменами.
type RetailShiftService interface {
	GetList(ctx context.Context, params *Params) (*List[RetailShift], *resty.Response, error)
	Create(ctx context.Context, retailShift *RetailShift, params *Params) (*RetailShift, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params *Params) (*RetailShift, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, retailShift *RetailShift, params *Params) (*RetailShift, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetaAttributesSharedStatesWrapper, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error)
	UpdateAttribute(ctx context.Context, id uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*RetailShift, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)
	MoveToTrash(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
}

func NewRetailShiftService(client *Client) RetailShiftService {
	e := NewEndpoint(client, "entity/retailshift")
	return newMainService[RetailShift, any, MetaAttributesSharedStatesWrapper, any](e)
}
