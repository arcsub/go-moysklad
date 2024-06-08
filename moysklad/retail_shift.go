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
	Organization        *Organization     `json:"organization,omitempty"`
	Shared              *bool             `json:"shared,omitempty"`
	AgentAccount        *AgentAccount     `json:"agentAccount,omitempty"`
	VatIncluded         *bool             `json:"vatIncluded,omitempty"`
	BankCommission      *float64          `json:"bankComission,omitempty"`
	BankPercent         *float64          `json:"bankPercent,omitempty"`
	Name                *string           `json:"name,omitempty"`
	CloseDate           *Timestamp        `json:"closeDate,omitempty"`
	Contract            *Contract         `json:"contract,omitempty"`
	Created             *Timestamp        `json:"created,omitempty"`
	Deleted             *Timestamp        `json:"deleted,omitempty"`
	Description         *string           `json:"description,omitempty"`
	ExternalCode        *string           `json:"externalCode,omitempty"`
	Files               *Files            `json:"files,omitempty"`
	Group               *Group            `json:"group,omitempty"`
	ID                  *uuid.UUID        `json:"id,omitempty"`
	Meta                *Meta             `json:"meta,omitempty"`
	AccountID           *uuid.UUID        `json:"accountId,omitempty"`
	Cheque              *Cheque           `json:"cheque,omitempty"`
	Acquire             *Counterparty     `json:"acquire,omitempty"`
	Moment              *Timestamp        `json:"moment,omitempty"`
	OrganizationAccount *AgentAccount     `json:"organizationAccount,omitempty"`
	Owner               *Employee         `json:"owner,omitempty"`
	Payments            *Payments         `json:"paymentOperations,omitempty"`
	Printed             *bool             `json:"printed,omitempty"`
	ProceedsCash        *float64          `json:"proceedsCash,omitempty"`
	ProceedsNoCash      *float64          `json:"proceedsNoCash,omitempty"`
	Published           *bool             `json:"published,omitempty"`
	QRAcquire           *Counterparty     `json:"qrAcquire,omitempty"`
	QRBankCommission    *float64          `json:"qrBankComission,omitempty"`
	QRBankPercent       *float64          `json:"qrBankPercent,omitempty"`
	ReceivedCash        *float64          `json:"receivedCash,omitempty"`
	ReceivedNoCash      *float64          `json:"receivedNoCash,omitempty"`
	RetailStore         *RetailStore      `json:"retailStore,omitempty"`
	Operations          *RetailOperations `json:"operations,omitempty"`
	Store               *Store            `json:"store,omitempty"`
	SyncID              *uuid.UUID        `json:"syncId,omitempty"`
	Updated             *Timestamp        `json:"updated,omitempty"`
	VatEnabled          *bool             `json:"vatEnabled,omitempty"`
	Attributes          Attributes        `json:"attributes,omitempty"`
}

func (r RetailShift) String() string {
	return Stringify(r)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (r RetailShift) GetMeta() Meta {
	return Deref(r.Meta)
}

func (r RetailShift) MetaType() MetaType {
	return MetaTypeRetailShift
}

// Cheque Информация о смене ККТ
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-roznichnaq-smena-roznichnye-smeny-informaciq-o-smene-kkt
type Cheque struct {
	Start ChequeStart `json:"start,omitempty"` // Информация об открытии смены
	End   ChequeEnd   `json:"end,omitempty"`   // Информация о закрытии смены
}

func (c Cheque) String() string {
	return Stringify(c)
}

// ChequeStart Информация об открытии смены ККТ
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-roznichnaq-smena-roznichnye-smeny-informaciq-ob-otkrytii-smeny-kkt
type ChequeStart struct {
	Time            Timestamp `json:"time,omitempty"`
	FnNumber        string    `json:"fnNumber,omitempty"`
	KktRegNumber    string    `json:"kktRegNumber,omitempty"`
	FiscalDocSign   string    `json:"fiscalDocSign,omitempty"`
	ShiftNumber     string    `json:"shiftNumber,omitempty"`
	FiscalDocNumber string    `json:"fiscalDocNumber,omitempty"`
}

func (c ChequeStart) String() string {
	return Stringify(c)
}

// ChequeEnd Информация о закрытии смены ККТ
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-roznichnaq-smena-roznichnye-smeny-informaciq-o-zakrytii-smeny-kkt
type ChequeEnd struct {
	ChequeStart
	ChequesTotal    float64 `json:"chequesTotal,omitempty"`    // Количество чеков за смену
	FiscalDocsTotal float64 `json:"fiscalDocsTotal,omitempty"` // Количество фискальных документов за смену
}

func (c ChequeEnd) String() string {
	return Stringify(c)
}

// RetailShiftService
// Сервис для работы с розничными сменами.
type RetailShiftService interface {
	GetList(ctx context.Context, params *Params) (*List[RetailShift], *resty.Response, error)
	Create(ctx context.Context, retailShift *RetailShift, params *Params) (*RetailShift, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*RetailShift, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, retailShift *RetailShift, params *Params) (*RetailShift, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetaAttributesSharedStatesWrapper, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id *uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error)
	UpdateAttribute(ctx context.Context, id *uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*RetailShift, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
	MoveToTrash(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewRetailShiftService(client *Client) RetailShiftService {
	e := NewEndpoint(client, "entity/retailshift")
	return newMainService[RetailShift, any, MetaAttributesSharedStatesWrapper, any](e)
}
