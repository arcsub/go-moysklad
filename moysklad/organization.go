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
	FSRARID                *string                  `json:"fsrarId,omitempty"`
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
	Attributes             Slice[AttributeValue]    `json:"attributes,omitempty"`
}

func (organization Organization) Clean() *Organization {
	return &Organization{Meta: organization.Meta}
}

func (organization Organization) GetUpdated() Timestamp {
	return Deref(organization.Updated)
}

func (organization Organization) GetChiefAccountSign() Image {
	return Deref(organization.ChiefAccountSign)
}

func (organization Organization) GetActualAddressFull() Address {
	return Deref(organization.ActualAddressFull)
}

func (organization Organization) GetArchived() bool {
	return Deref(organization.Archived)
}

func (organization Organization) GetBonusPoints() int {
	return Deref(organization.BonusPoints)
}

func (organization Organization) GetBonusProgram() BonusProgram {
	return Deref(organization.BonusProgram)
}

func (organization Organization) GetActualAddress() string {
	return Deref(organization.ActualAddress)
}

func (organization Organization) GetUTMUrl() string {
	return Deref(organization.UTMUrl)
}

func (organization Organization) GetCreated() Timestamp {
	return Deref(organization.Created)
}

func (organization Organization) GetDescription() string {
	return Deref(organization.Description)
}

func (organization Organization) GetExternalCode() string {
	return Deref(organization.ExternalCode)
}

func (organization Organization) GetGroup() Group {
	return Deref(organization.Group)
}

func (organization Organization) GetID() uuid.UUID {
	return Deref(organization.ID)
}

func (organization Organization) GetMeta() Meta {
	return Deref(organization.Meta)
}

func (organization Organization) GetName() string {
	return Deref(organization.Name)
}

func (organization Organization) GetOwner() Employee {
	return Deref(organization.Owner)
}

func (organization Organization) GetShared() bool {
	return Deref(organization.Shared)
}

func (organization Organization) GetSyncID() uuid.UUID {
	return Deref(organization.SyncID)
}

func (organization Organization) GetTrackingContractDate() Timestamp {
	return Deref(organization.TrackingContractDate)
}

func (organization Organization) GetTrackingContractNumber() string {
	return Deref(organization.TrackingContractNumber)
}

func (organization Organization) GetCertificateNumber() string {
	return Deref(organization.CertificateNumber)
}

func (organization Organization) GetAccounts() MetaArray[AgentAccount] {
	return Deref(organization.Accounts)
}

func (organization Organization) GetStamp() Image {
	return Deref(organization.Stamp)
}

func (organization Organization) GetCertificateDate() Timestamp {
	return Deref(organization.CertificateDate)
}

func (organization Organization) GetAccountID() uuid.UUID {
	return Deref(organization.AccountID)
}

func (organization Organization) GetCode() string {
	return Deref(organization.Code)
}

func (organization Organization) GetChiefAccountant() string {
	return Deref(organization.ChiefAccountant)
}

func (organization Organization) GetDirector() string {
	return Deref(organization.Director)
}

func (organization Organization) GetDirectorPosition() string {
	return Deref(organization.DirectorPosition)
}

func (organization Organization) GetDirectorSign() Image {
	return Deref(organization.DirectorSign)
}

func (organization Organization) GetEmail() string {
	return Deref(organization.Email)
}

func (organization Organization) GetFax() string {
	return Deref(organization.Fax)
}

func (organization Organization) GetFSRARID() string {
	return Deref(organization.FSRARID)
}

func (organization Organization) GetINN() string {
	return Deref(organization.INN)
}

func (organization Organization) GetIsEGAISEnable() bool {
	return Deref(organization.IsEGAISEnable)
}

func (organization Organization) GetKPP() string {
	return Deref(organization.KPP)
}

func (organization Organization) GetLegalAddress() string {
	return Deref(organization.LegalAddress)
}

func (organization Organization) GetLegalAddressFull() Address {
	return Deref(organization.LegalAddressFull)
}

func (organization Organization) GetLegalFirstName() string {
	return Deref(organization.LegalFirstName)
}

func (organization Organization) GetLegalLastName() string {
	return Deref(organization.LegalLastName)
}

func (organization Organization) GetLegalMiddleName() string {
	return Deref(organization.LegalMiddleName)
}

func (organization Organization) GetLegalTitle() string {
	return Deref(organization.LegalTitle)
}

func (organization Organization) GetOGRN() string {
	return Deref(organization.OGRN)
}

func (organization Organization) GetOGRNIP() string {
	return Deref(organization.OGRNIP)
}

func (organization Organization) GetOKPO() string {
	return Deref(organization.OKPO)
}

func (organization Organization) GetPayerVat() bool {
	return Deref(organization.PayerVat)
}

func (organization Organization) GetPhone() string {
	return Deref(organization.Phone)
}

func (organization Organization) GetCompanyType() CompanyType {
	return organization.CompanyType
}

func (organization Organization) GetAttributes() Slice[AttributeValue] {
	return organization.Attributes
}

func (organization *Organization) SetChiefAccountSign(chiefAccountSign *Image) *Organization {
	organization.ChiefAccountSign = chiefAccountSign
	return organization
}

func (organization *Organization) SetActualAddressFull(actualAddressFull *Address) *Organization {
	organization.ActualAddressFull = actualAddressFull
	return organization
}

func (organization *Organization) SetArchived(archived bool) *Organization {
	organization.Archived = &archived
	return organization
}

func (organization *Organization) SetBonusProgram(bonusProgram *BonusProgram) *Organization {
	organization.BonusProgram = bonusProgram.Clean()
	return organization
}

func (organization *Organization) SetActualAddress(actualAddress string) *Organization {
	organization.ActualAddress = &actualAddress
	return organization
}

func (organization *Organization) SetUTMUrl(utmUrl string) *Organization {
	organization.UTMUrl = &utmUrl
	return organization
}

func (organization *Organization) SetDescription(description string) *Organization {
	organization.Description = &description
	return organization
}

func (organization *Organization) SetExternalCode(externalCode string) *Organization {
	organization.ExternalCode = &externalCode
	return organization
}

func (organization *Organization) SetGroup(group *Group) *Organization {
	organization.Group = group.Clean()
	return organization
}

func (organization *Organization) SetMeta(meta *Meta) *Organization {
	organization.Meta = meta
	return organization
}

func (organization *Organization) SetName(name string) *Organization {
	organization.Name = &name
	return organization
}

func (organization *Organization) SetOwner(owner *Employee) *Organization {
	organization.Owner = owner.Clean()
	return organization
}

func (organization *Organization) SetShared(shared bool) *Organization {
	organization.Shared = &shared
	return organization
}

func (organization *Organization) SetSyncID(syncID uuid.UUID) *Organization {
	organization.SyncID = &syncID
	return organization
}

func (organization *Organization) SetTrackingContractDate(trackingContractDate *Timestamp) *Organization {
	organization.TrackingContractDate = trackingContractDate
	return organization
}

func (organization *Organization) SetTrackingContractNumber(trackingContractNumber string) *Organization {
	organization.TrackingContractNumber = &trackingContractNumber
	return organization
}

func (organization *Organization) SetCertificateNumber(certificateNumber string) *Organization {
	organization.CertificateNumber = &certificateNumber
	return organization
}

func (organization *Organization) SetAccounts(accounts Slice[AgentAccount]) *Organization {
	organization.Accounts = NewMetaArrayRows(accounts)
	return organization
}

func (organization *Organization) SetStamp(stamp *Image) *Organization {
	organization.Stamp = stamp
	return organization
}

func (organization *Organization) SetCertificateDate(certificateDate *Timestamp) *Organization {
	organization.CertificateDate = certificateDate
	return organization
}

func (organization *Organization) SetCode(code string) *Organization {
	organization.Code = &code
	return organization
}

func (organization *Organization) SetChiefAccountant(chiefAccountant string) *Organization {
	organization.ChiefAccountant = &chiefAccountant
	return organization
}

func (organization *Organization) SetDirector(director string) *Organization {
	organization.Director = &director
	return organization
}

func (organization *Organization) SetDirectorPosition(directorPosition string) *Organization {
	organization.DirectorPosition = &directorPosition
	return organization
}

func (organization *Organization) SetDirectorSign(directorSign *Image) *Organization {
	organization.DirectorSign = directorSign
	return organization
}

func (organization *Organization) SetEmail(email string) *Organization {
	organization.Email = &email
	return organization
}

func (organization *Organization) SetFax(fax string) *Organization {
	organization.Fax = &fax
	return organization
}

func (organization *Organization) SetFSRARID(fsrarID string) *Organization {
	organization.FSRARID = &fsrarID
	return organization
}

func (organization *Organization) SetINN(inn string) *Organization {
	organization.INN = &inn
	return organization
}

func (organization *Organization) SetIsEGAISEnable(isEGAISEnable bool) *Organization {
	organization.IsEGAISEnable = &isEGAISEnable
	return organization
}

func (organization *Organization) SetKPP(kpp string) *Organization {
	organization.KPP = &kpp
	return organization
}

func (organization *Organization) SetLegalAddress(legalAddress string) *Organization {
	organization.LegalAddress = &legalAddress
	return organization
}

func (organization *Organization) SetLegalAddressFull(legalAddressFull *Address) *Organization {
	organization.LegalAddressFull = legalAddressFull
	return organization
}

func (organization *Organization) SetLegalFirstName(legalFirstName string) *Organization {
	organization.LegalFirstName = &legalFirstName
	return organization
}

func (organization *Organization) SetLegalLastName(legalLastName string) *Organization {
	organization.LegalLastName = &legalLastName
	return organization
}

func (organization *Organization) SetLegalMiddleName(legalMiddleName string) *Organization {
	organization.LegalMiddleName = &legalMiddleName
	return organization
}

func (organization *Organization) SetLegalTitle(legalTitle string) *Organization {
	organization.LegalTitle = &legalTitle
	return organization
}

func (organization *Organization) SetOGRN(ogrn string) *Organization {
	organization.OGRN = &ogrn
	return organization
}

func (organization *Organization) SetOGRNIP(ogrnip string) *Organization {
	organization.OGRNIP = &ogrnip
	return organization
}

func (organization *Organization) SetOKPO(okpo string) *Organization {
	organization.OKPO = &okpo
	return organization
}

func (organization *Organization) SetPayerVat(payerVat bool) *Organization {
	organization.PayerVat = &payerVat
	return organization
}

func (organization *Organization) SetPhone(phone string) *Organization {
	organization.Phone = &phone
	return organization
}

func (organization *Organization) SetCompanyType(companyType CompanyType) *Organization {
	organization.CompanyType = companyType
	return organization
}

func (organization *Organization) SetAttributes(attributes Slice[AttributeValue]) *Organization {
	organization.Attributes = attributes
	return organization
}

func (organization Organization) String() string {
	return Stringify(organization)
}

func (organization Organization) MetaType() MetaType {
	return MetaTypeOrganization
}

// OrganizationService
// Сервис для работы с юридическими лицами.
type OrganizationService interface {
	GetList(ctx context.Context, params *Params) (*List[Organization], *resty.Response, error)
	Create(ctx context.Context, organization *Organization, params *Params) (*Organization, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, organizationList []*Organization, params *Params) (*[]Organization, *resty.Response, error)
	DeleteMany(ctx context.Context, organizationList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params *Params) (*Organization, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, organization *Organization, params *Params) (*Organization, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetaAttributesSharedWrapper, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error)
	UpdateAttribute(ctx context.Context, id uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	GetAccounts(ctx context.Context, id uuid.UUID) (*List[AgentAccount], *resty.Response, error)
	GetAccountByID(ctx context.Context, id uuid.UUID, accountID uuid.UUID) (*AgentAccount, *resty.Response, error)
	UpdateAccounts(ctx context.Context, id uuid.UUID, accounts Slice[AgentAccount]) (*MetaArray[AgentAccount], *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*Organization, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error)
}

func NewOrganizationService(client *Client) OrganizationService {
	e := NewEndpoint(client, "entity/organization")
	return newMainService[Organization, any, MetaAttributesSharedWrapper, any](e)
}
