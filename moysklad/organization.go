package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Organization Юрлицо.
// Ключевое слово: organization
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-jurlico
type Organization struct {
	Updated                *Timestamp               `json:"updated,omitempty"`
	ChiefAccountSign       *Image                   `json:"chiefAccountSign,omitempty"`
	ActualAddressFull      *Address                 `json:"actualAddressFull,omitempty"`
	Archived               *bool                    `json:"archived,omitempty"`
	BonusPoints            *int                     `json:"bonusPoints,omitempty"`
	BonusProgram           *BonusProgram            `json:"bonusProgram,omitempty"`
	ActualAddress          *string                  `json:"actualAddress,omitempty"`
	UTMUrl                 *string                  `json:"utmUrl,omitempty"`
	Created                *Timestamp               `json:"created,omitempty"`
	Description            *string                  `json:"description,omitempty"`
	ExternalCode           *string                  `json:"externalCode,omitempty"`
	Group                  *Group                   `json:"group,omitempty"`
	ID                     *uuid.UUID               `json:"id,omitempty"`
	Meta                   *Meta                    `json:"meta,omitempty"`
	Name                   *string                  `json:"name,omitempty"`
	Owner                  *Employee                `json:"owner,omitempty"`
	Shared                 *bool                    `json:"shared,omitempty"`
	SyncID                 *uuid.UUID               `json:"syncId,omitempty"`
	TrackingContractDate   *Timestamp               `json:"trackingContractDate,omitempty"`
	TrackingContractNumber *string                  `json:"trackingContractNumber,omitempty"`
	CertificateNumber      *string                  `json:"certificateNumber,omitempty"`
	Accounts               *MetaArray[AgentAccount] `json:"accounts,omitempty"`
	Stamp                  *Image                   `json:"stamp,omitempty"`
	CertificateDate        *Timestamp               `json:"certificateDate,omitempty"`
	AccountID              *uuid.UUID               `json:"accountId,omitempty"`
	Code                   *string                  `json:"code,omitempty"`
	ChiefAccountant        *string                  `json:"chiefAccountant,omitempty"`
	Director               *string                  `json:"director,omitempty"`
	DirectorPosition       *string                  `json:"directorPosition,omitempty"`
	DirectorSign           *Image                   `json:"directorSign,omitempty"`
	Email                  *string                  `json:"email,omitempty"`
	Fax                    *string                  `json:"fax,omitempty"`
	FSRARId                *string                  `json:"fsrarId,omitempty"`
	INN                    *string                  `json:"inn,omitempty"`
	IsEGAISEnable          *bool                    `json:"isEgaisEnable,omitempty"`
	KPP                    *string                  `json:"kpp,omitempty"`
	LegalAddress           *string                  `json:"legalAddress,omitempty"`
	LegalAddressFull       *Address                 `json:"legalAddressFull,omitempty"`
	LegalFirstName         *string                  `json:"legalFirstName,omitempty"`
	LegalLastName          *string                  `json:"legalLastName,omitempty"`
	LegalMiddleName        *string                  `json:"legalMiddleName,omitempty"`
	LegalTitle             *string                  `json:"legalTitle,omitempty"`
	OGRN                   *string                  `json:"ogrn,omitempty"`
	OGRNIP                 *string                  `json:"ogrnip,omitempty"`
	OKPO                   *string                  `json:"okpo,omitempty"`
	PayerVat               *bool                    `json:"payerVat,omitempty"`
	Phone                  *string                  `json:"phone,omitempty"`
	CompanyType            CompanyType              `json:"companyType,omitempty"`
	Attributes             Attributes               `json:"attributes,omitempty"`
}

func (o Organization) String() string {
	return Stringify(o)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (o Organization) GetMeta() Meta {
	return Deref(o.Meta)
}

func (o Organization) MetaType() MetaType {
	return MetaTypeOrganization
}

// OrganizationService
// Сервис для работы с юридическими лицами.
type OrganizationService interface {
	GetList(ctx context.Context, params *Params) (*List[Organization], *resty.Response, error)
	Create(ctx context.Context, organization *Organization, params *Params) (*Organization, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, organizationList []*Organization, params *Params) (*[]Organization, *resty.Response, error)
	DeleteMany(ctx context.Context, organizationList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Organization, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, organization *Organization, params *Params) (*Organization, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetaAttributesSharedWrapper, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id *uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error)
	UpdateAttribute(ctx context.Context, id *uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	GetAccounts(ctx context.Context, id *uuid.UUID) (*List[AgentAccount], *resty.Response, error)
	GetAccountByID(ctx context.Context, id *uuid.UUID, accountId *uuid.UUID) (*AgentAccount, *resty.Response, error)
	UpdateAccounts(ctx context.Context, id *uuid.UUID, accounts []*AgentAccount) (*[]AgentAccount, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*Organization, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
}

func NewOrganizationService(client *Client) OrganizationService {
	e := NewEndpoint(client, "entity/organization")
	return newMainService[Organization, any, MetaAttributesSharedWrapper, any](e)
}
