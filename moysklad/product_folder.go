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

// Clean возвращает сущность с единственным заполненным полем Meta
func (productFolder ProductFolder) Clean() *ProductFolder {
	return &ProductFolder{Meta: productFolder.Meta}
}

func (productFolder ProductFolder) GetName() string {
	return Deref(productFolder.Name)
}

func (productFolder ProductFolder) GetUseParentVat() bool {
	return Deref(productFolder.UseParentVat)
}

func (productFolder ProductFolder) GetCode() string {
	return Deref(productFolder.Code)
}

func (productFolder ProductFolder) GetDescription() string {
	return Deref(productFolder.Description)
}

func (productFolder ProductFolder) GetEffectiveVat() int {
	return Deref(productFolder.EffectiveVat)
}

func (productFolder ProductFolder) GetEffectiveVatEnabled() bool {
	return Deref(productFolder.EffectiveVatEnabled)
}

func (productFolder ProductFolder) GetExternalCode() string {
	return Deref(productFolder.ExternalCode)
}

func (productFolder ProductFolder) GetAccountID() uuid.UUID {
	return Deref(productFolder.AccountID)
}

func (productFolder ProductFolder) GetVatEnabled() bool {
	return Deref(productFolder.VatEnabled)
}

func (productFolder ProductFolder) GetArchived() bool {
	return Deref(productFolder.Archived)
}

func (productFolder ProductFolder) GetGroup() Group {
	return Deref(productFolder.Group)
}

func (productFolder ProductFolder) GetOwner() Employee {
	return Deref(productFolder.Owner)
}

func (productFolder ProductFolder) GetPathName() string {
	return Deref(productFolder.PathName)
}

func (productFolder ProductFolder) GetProductFolder() ProductFolder {
	return Deref(productFolder.ProductFolder)
}

func (productFolder ProductFolder) GetShared() bool {
	return Deref(productFolder.Shared)
}

func (productFolder ProductFolder) GetID() uuid.UUID {
	return Deref(productFolder.ID)
}

func (productFolder ProductFolder) GetUpdated() Timestamp {
	return Deref(productFolder.Updated)
}

func (productFolder ProductFolder) GetMeta() Meta {
	return Deref(productFolder.Meta)
}

func (productFolder ProductFolder) GetVat() int {
	return Deref(productFolder.Vat)
}

func (productFolder ProductFolder) GetTaxSystem() GoodTaxSystem {
	return productFolder.TaxSystem
}

func (productFolder *ProductFolder) SetName(name string) *ProductFolder {
	productFolder.Name = &name
	return productFolder
}

func (productFolder *ProductFolder) SetUseParentVat(useParentVat bool) *ProductFolder {
	productFolder.UseParentVat = &useParentVat
	return productFolder
}

func (productFolder *ProductFolder) SetCode(code string) *ProductFolder {
	productFolder.Code = &code
	return productFolder
}

func (productFolder *ProductFolder) SetDescription(description string) *ProductFolder {
	productFolder.Description = &description
	return productFolder
}

func (productFolder *ProductFolder) SetExternalCode(externalCode string) *ProductFolder {
	productFolder.ExternalCode = &externalCode
	return productFolder
}

func (productFolder *ProductFolder) SetVatEnabled(vatEnabled bool) *ProductFolder {
	productFolder.VatEnabled = &vatEnabled
	return productFolder
}

func (productFolder *ProductFolder) SetArchived(archived bool) *ProductFolder {
	productFolder.Archived = &archived
	return productFolder
}

func (productFolder *ProductFolder) SetGroup(group *Group) *ProductFolder {
	productFolder.Group = group.Clean()
	return productFolder
}

func (productFolder *ProductFolder) SetOwner(owner *Employee) *ProductFolder {
	productFolder.Owner = owner.Clean()
	return productFolder
}

func (productFolder *ProductFolder) SetProductFolder(parent *ProductFolder) *ProductFolder {
	productFolder.ProductFolder = parent.Clean()
	return productFolder
}

func (productFolder *ProductFolder) SetShared(shared bool) *ProductFolder {
	productFolder.Shared = &shared
	return productFolder
}

func (productFolder *ProductFolder) SetMeta(meta *Meta) *ProductFolder {
	productFolder.Meta = meta
	return productFolder
}

func (productFolder *ProductFolder) SetVat(vat int) *ProductFolder {
	productFolder.Vat = &vat
	return productFolder
}

func (productFolder *ProductFolder) SetTaxSystem(taxSystem GoodTaxSystem) *ProductFolder {
	productFolder.TaxSystem = taxSystem
	return productFolder
}

func (productFolder ProductFolder) String() string {
	return Stringify(productFolder)
}

func (productFolder ProductFolder) MetaType() MetaType {
	return MetaTypeProductFolder
}

// Update shortcut
func (productFolder ProductFolder) Update(ctx context.Context, client *Client, params ...*Params) (*ProductFolder, *resty.Response, error) {
	return client.Entity().ProductFolder().Update(ctx, productFolder.GetID(), &productFolder, params...)
}

// Create shortcut
func (productFolder ProductFolder) Create(ctx context.Context, client *Client, params ...*Params) (*ProductFolder, *resty.Response, error) {
	return client.Entity().ProductFolder().Create(ctx, &productFolder, params...)
}

// Delete shortcut
func (productFolder ProductFolder) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return client.Entity().ProductFolder().Delete(ctx, productFolder.GetID())
}

// ProductFolderService
// Сервис для работы с группами товаров.
type ProductFolderService interface {
	GetList(ctx context.Context, params ...*Params) (*List[ProductFolder], *resty.Response, error)
	Create(ctx context.Context, productFolder *ProductFolder, params ...*Params) (*ProductFolder, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, productFolderList Slice[ProductFolder], params ...*Params) (*Slice[ProductFolder], *resty.Response, error)
	DeleteMany(ctx context.Context, entities ...ProductFolder) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetaAttributesWrapper, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList Slice[Attribute]) (*Slice[Attribute], *resty.Response, error)
	UpdateAttribute(ctx context.Context, id uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList []MetaWrapper) (*DeleteManyResponse, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*ProductFolder, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, productFolder *ProductFolder, params ...*Params) (*ProductFolder, *resty.Response, error)
}

func NewProductFolderService(client *Client) ProductFolderService {
	e := NewEndpoint(client, "entity/productfolder")
	return newMainService[ProductFolder, any, MetaAttributesWrapper, any](e)
}
