package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// ProductFolder Группа товаров.
// Ключевое слово: productfolder
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-gruppa-towarow
type ProductFolder struct {
	Name                *string        `json:"name,omitempty"`
	UseParentVat        *bool          `json:"useParentVat,omitempty"`
	Code                *string        `json:"code,omitempty"`
	Description         *string        `json:"description,omitempty"`
	EffectiveVat        *int           `json:"effectiveVat,omitempty"`
	EffectiveVatEnabled *bool          `json:"effectiveVatEnabled,omitempty"`
	ExternalCode        *string        `json:"externalCode,omitempty"`
	AccountID           *uuid.UUID     `json:"accountId,omitempty"`
	VatEnabled          *bool          `json:"vatEnabled,omitempty"`
	Archived            *bool          `json:"archived,omitempty"`
	Group               *Group         `json:"group,omitempty"`
	Owner               *Employee      `json:"owner,omitempty"`
	PathName            *string        `json:"pathName,omitempty"`
	ProductFolder       *ProductFolder `json:"productFolder,omitempty"`
	Shared              *bool          `json:"shared,omitempty"`
	ID                  *uuid.UUID     `json:"id,omitempty"`
	Updated             *Timestamp     `json:"updated,omitempty"`
	Meta                *Meta          `json:"meta,omitempty"`
	Vat                 *int           `json:"vat,omitempty"`
	TaxSystem           GoodTaxSystem  `json:"taxSystem,omitempty"`
}

func (p ProductFolder) String() string {
	return Stringify(p)
}

func (p ProductFolder) MetaType() MetaType {
	return MetaTypeProductFolder
}

type ProductFolders MetaArray[ProductFolder]

// ProductFolderService
// Сервис для работы с группами товаров.
type ProductFolderService interface {
	GetList(ctx context.Context, params *Params) (*List[ProductFolder], *resty.Response, error)
	Create(ctx context.Context, productFolder *ProductFolder, params *Params) (*ProductFolder, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, productFolderList []*ProductFolder, params *Params) (*[]ProductFolder, *resty.Response, error)
	DeleteMany(ctx context.Context, productFolderList []*ProductFolder) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetadataAttribute, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id *uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error)
	UpdateAttribute(ctx context.Context, id *uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList []*Attribute) (*DeleteManyResponse, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*ProductFolder, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, productFolder *ProductFolder, params *Params) (*ProductFolder, *resty.Response, error)
}

func NewProductFolderService(client *Client) ProductFolderService {
	e := NewEndpoint(client, "entity/productfolder")
	return newMainService[ProductFolder, any, MetadataAttribute, any](e)
}
