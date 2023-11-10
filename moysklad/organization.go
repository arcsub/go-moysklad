package moysklad

import (
	"github.com/google/uuid"
)

// Organization Юрлицо.
// Ключевое слово: organization
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-jurlico
type Organization struct {
	CertificateNumber      *string       `json:"certificateNumber,omitempty"`
	Code                   *string       `json:"code,omitempty"`
	ActualAddressFull      *Address      `json:"actualAddressFull,omitempty"`
	Archived               *bool         `json:"archived,omitempty"`
	BonusPoints            *int          `json:"bonusPoints,omitempty"`
	BonusProgram           *BonusProgram `json:"bonusProgram,omitempty"`
	ChiefAccountSign       *Image        `json:"chiefAccountSign,omitempty"`
	UTMUrl                 *string       `json:"utmUrl,omitempty"`
	Created                *Timestamp    `json:"created,omitempty"`
	Description            *string       `json:"description,omitempty"`
	ExternalCode           *string       `json:"externalCode,omitempty"`
	ChiefAccountant        *string       `json:"chiefAccountant,omitempty"`
	ID                     *uuid.UUID    `json:"id,omitempty"`
	Meta                   *Meta         `json:"meta,omitempty"`
	Name                   *string       `json:"name,omitempty"`
	Owner                  *Employee     `json:"owner,omitempty"`
	Shared                 *bool         `json:"shared,omitempty"`
	SyncID                 *uuid.UUID    `json:"syncId,omitempty"`
	TrackingContractDate   *Timestamp    `json:"trackingContractDate,omitempty"`
	TrackingContractNumber *string       `json:"trackingContractNumber,omitempty"`
	AccountID              *uuid.UUID    `json:"accountId,omitempty"`
	Accounts               *MetaWrapper  `json:"accounts,omitempty"`
	Attributes             *Attributes   `json:"attributes,omitempty"`
	CertificateDate        *Timestamp    `json:"certificateDate,omitempty"`
	Updated                *Timestamp    `json:"updated,omitempty"`
	ActualAddress          *string       `json:"actualAddress,omitempty"`
	Group                  *Group        `json:"group,omitempty"`
	Director               *string       `json:"director,omitempty"`
	DirectorPosition       *string       `json:"directorPosition,omitempty"`
	DirectorSign           *Image        `json:"directorSign,omitempty"`
	Email                  *string       `json:"email,omitempty"`
	Fax                    *string       `json:"fax,omitempty"`
	FSRARId                *string       `json:"fsrarId,omitempty"`
	INN                    *string       `json:"inn,omitempty"`
	IsEGAISEnable          *bool         `json:"isEgaisEnable,omitempty"`
	KPP                    *string       `json:"kpp,omitempty"`
	LegalAddress           *string       `json:"legalAddress,omitempty"`
	LegalAddressFull       *Address      `json:"legalAddressFull,omitempty"`
	LegalFirstName         *string       `json:"legalFirstName,omitempty"`
	LegalLastName          *string       `json:"legalLastName,omitempty"`
	LegalMiddleName        *string       `json:"legalMiddleName,omitempty"`
	LegalTitle             *string       `json:"legalTitle,omitempty"`
	OGRN                   *string       `json:"ogrn,omitempty"`
	OGRNIP                 *string       `json:"ogrnip,omitempty"`
	OKPO                   *string       `json:"okpo,omitempty"`
	PayerVat               *bool         `json:"payerVat,omitempty"`
	Phone                  *string       `json:"phone,omitempty"`
	Stamp                  *Image        `json:"stamp,omitempty"`
	CompanyType            CompanyType   `json:"companyType,omitempty"`
}

func (o Organization) String() string {
	return Stringify(o)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (o Organization) GetMeta() *Meta {
	return o.Meta
}

func (o Organization) MetaType() MetaType {
	return MetaTypeOrganization
}
