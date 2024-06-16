package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Service Услуга.
// Ключевое слово: service
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-usluga
type Service struct {
	VatEnabled          *bool                 `json:"vatEnabled,omitempty"`
	Group               *Group                `json:"group,omitempty"`
	Barcodes            Slice[Barcode]        `json:"barcodes,omitempty"`
	Description         *string               `json:"description,omitempty"`
	ExternalCode        *string               `json:"externalCode,omitempty"`
	ID                  *uuid.UUID            `json:"id,omitempty"`
	Meta                *Meta                 `json:"meta,omitempty"`
	Name                *string               `json:"name,omitempty"`
	Archived            *bool                 `json:"archived,omitempty"`
	Files               *MetaArray[File]      `json:"files,omitempty"`
	BuyPrice            *BuyPrice             `json:"buyPrice,omitempty"`
	DiscountProhibited  *bool                 `json:"discountProhibited,omitempty"`
	EffectiveVat        *int                  `json:"effectiveVat,omitempty"`
	EffectiveVatEnabled *bool                 `json:"effectiveVatEnabled,omitempty"`
	UseParentVat        *bool                 `json:"useParentVat,omitempty"`
	Code                *string               `json:"code,omitempty"`
	MinPrice            *MinPrice             `json:"minPrice,omitempty"`
	Owner               *Employee             `json:"owner,omitempty"`
	PathName            *string               `json:"pathName,omitempty"`
	AccountID           *uuid.UUID            `json:"accountId,omitempty"`
	ProductFolder       *ProductFolder        `json:"productFolder,omitempty"`
	SalePrices          Slice[SalePrice]      `json:"salePrices,omitempty"`
	Shared              *bool                 `json:"shared,omitempty"`
	SyncID              *uuid.UUID            `json:"syncId,omitempty"`
	Vat                 *int                  `json:"vat,omitempty"`
	Uom                 *Uom                  `json:"uom,omitempty"`
	Updated             *Timestamp            `json:"updated,omitempty"`
	PaymentItemType     PaymentItem           `json:"paymentItemType,omitempty"`
	TaxSystem           TaxSystem             `json:"taxSystem,omitempty"`
	Attributes          Slice[AttributeValue] `json:"attributes,omitempty"`
}

// Clean возвращает сущность с единственным заполненным полем Meta
func (service Service) Clean() *Service {
	return &Service{Meta: service.Meta}
}

func NewServiceFromAssortment(assortmentPosition AssortmentPosition) *Service {
	return unmarshalAsType[Service](assortmentPosition)
}

func (service Service) FromAssortment(assortmentPosition AssortmentPosition) *Service {
	return unmarshalAsType[Service](assortmentPosition)
}

func (service Service) GetVatEnabled() bool {
	return Deref(service.VatEnabled)
}

func (service Service) GetGroup() Group {
	return Deref(service.Group)
}

func (service Service) GetBarcodes() Slice[Barcode] {
	return service.Barcodes
}

func (service Service) GetDescription() string {
	return Deref(service.Description)
}

func (service Service) GetExternalCode() string {
	return Deref(service.ExternalCode)
}

func (service Service) GetID() uuid.UUID {
	return Deref(service.ID)
}

func (service Service) GetMeta() Meta {
	return Deref(service.Meta)
}

func (service Service) GetName() string {
	return Deref(service.Name)
}

func (service Service) GetArchived() bool {
	return Deref(service.Archived)
}

func (service Service) GetFiles() MetaArray[File] {
	return Deref(service.Files)
}

func (service Service) GetBuyPrice() BuyPrice {
	return Deref(service.BuyPrice)
}

func (service Service) GetDiscountProhibited() bool {
	return Deref(service.DiscountProhibited)
}

func (service Service) GetEffectiveVat() int {
	return Deref(service.EffectiveVat)
}

func (service Service) GetEffectiveVatEnabled() bool {
	return Deref(service.EffectiveVatEnabled)
}

func (service Service) GetUseParentVat() bool {
	return Deref(service.UseParentVat)
}

func (service Service) GetCode() string {
	return Deref(service.Code)
}

func (service Service) GetMinPrice() MinPrice {
	return Deref(service.MinPrice)
}

func (service Service) GetOwner() Employee {
	return Deref(service.Owner)
}

func (service Service) GetPathName() string {
	return Deref(service.PathName)
}

func (service Service) GetAccountID() uuid.UUID {
	return Deref(service.AccountID)
}

func (service Service) GetProductFolder() ProductFolder {
	return Deref(service.ProductFolder)
}

func (service Service) GetSalePrices() Slice[SalePrice] {
	return service.SalePrices
}

func (service Service) GetShared() bool {
	return Deref(service.Shared)
}

func (service Service) GetSyncID() uuid.UUID {
	return Deref(service.SyncID)
}

func (service Service) GetVat() int {
	return Deref(service.Vat)
}

func (service Service) GetUom() Uom {
	return Deref(service.Uom)
}

func (service Service) GetUpdated() Timestamp {
	return Deref(service.Updated)
}

func (service Service) GetPaymentItemType() PaymentItem {
	return service.PaymentItemType
}

func (service Service) GetTaxSystem() TaxSystem {
	return service.TaxSystem
}

func (service Service) GetAttributes() Slice[AttributeValue] {
	return service.Attributes
}

func (service *Service) SetVatEnabled(vatEnabled bool) *Service {
	service.VatEnabled = &vatEnabled
	return service
}

func (service *Service) SetGroup(group *Group) *Service {
	service.Group = group.Clean()
	return service
}

func (service *Service) SetBarcodes(barcodes Slice[Barcode]) *Service {
	service.Barcodes = barcodes
	return service
}

func (service *Service) SetDescription(description string) *Service {
	service.Description = &description
	return service
}

func (service *Service) SetExternalCode(externalCode string) *Service {
	service.ExternalCode = &externalCode
	return service
}

func (service *Service) SetMeta(meta *Meta) *Service {
	service.Meta = meta
	return service
}

func (service *Service) SetName(name string) *Service {
	service.Name = &name
	return service
}

func (service *Service) SetArchived(archived bool) *Service {
	service.Archived = &archived
	return service
}

func (service *Service) SetFiles(files Slice[File]) *Service {
	service.Files = NewMetaArrayRows(files)
	return service
}

func (service *Service) SetBuyPrice(buyPrice *BuyPrice) *Service {
	service.BuyPrice = buyPrice
	return service
}

func (service *Service) SetDiscountProhibited(discountProhibited bool) *Service {
	service.DiscountProhibited = &discountProhibited
	return service
}

func (service *Service) SetUseParentVat(useParentVat bool) *Service {
	service.UseParentVat = &useParentVat
	return service
}

func (service *Service) SetCode(code string) *Service {
	service.Code = &code
	return service
}

func (service *Service) SetMinPrice(minPrice *MinPrice) *Service {
	service.MinPrice = minPrice
	return service
}

func (service *Service) SetOwner(owner *Employee) *Service {
	service.Owner = owner.Clean()
	return service
}

func (service *Service) SetProductFolder(productFolder *ProductFolder) *Service {
	service.ProductFolder = productFolder.Clean()
	return service
}

func (service *Service) SetSalePrices(salePrices Slice[SalePrice]) *Service {
	service.SalePrices = salePrices
	return service
}

func (service *Service) SetShared(shared bool) *Service {
	service.Shared = &shared
	return service
}

func (service *Service) SetSyncID(syncID uuid.UUID) *Service {
	service.SyncID = &syncID
	return service
}

func (service *Service) SetVat(vat int) *Service {
	service.Vat = &vat
	return service
}

func (service *Service) SetUom(uom *Uom) *Service {
	service.Uom = uom.Clean()
	return service
}

func (service *Service) SetPaymentItemType(paymentItemType PaymentItem) *Service {
	service.PaymentItemType = paymentItemType
	return service
}

func (service *Service) SetTaxSystem(taxSystem TaxSystem) *Service {
	service.TaxSystem = taxSystem
	return service
}

func (service *Service) SetAttributes(attributes Slice[AttributeValue]) *Service {
	service.Attributes = attributes
	return service
}

func (service Service) AsAssortment() *AssortmentPosition {
	return &AssortmentPosition{Meta: service.GetMeta()}
}

func (service Service) String() string {
	return Stringify(service)
}

func (service Service) MetaType() MetaType {
	return MetaTypeService
}

// Update shortcut
func (service Service) Update(ctx context.Context, client *Client, params ...*Params) (*Service, *resty.Response, error) {
	return client.Entity().Service().Update(ctx, service.GetID(), &service, params...)
}

// Create shortcut
func (service Service) Create(ctx context.Context, client *Client, params ...*Params) (*Service, *resty.Response, error) {
	return client.Entity().Service().Create(ctx, &service, params...)
}

// Delete shortcut
func (service Service) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return client.Entity().Service().Delete(ctx, service.GetID())
}

// ServiceService
// Сервис для работы с услугами.
type ServiceService interface {
	GetList(ctx context.Context, params ...*Params) (*List[Service], *resty.Response, error)
	Create(ctx context.Context, service *Service, params ...*Params) (*Service, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, serviceList Slice[Service], params ...*Params) (*Slice[Service], *resty.Response, error)
	DeleteMany(ctx context.Context, serviceList []MetaWrapper) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Service, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, service *Service, params ...*Params) (*Service, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*Service, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params ...*Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)
	GetFiles(ctx context.Context, id uuid.UUID) (*MetaArray[File], *resty.Response, error)
	CreateFile(ctx context.Context, id uuid.UUID, file *File) (*Slice[File], *resty.Response, error)
	UpdateFiles(ctx context.Context, id uuid.UUID, files Slice[File]) (*Slice[File], *resty.Response, error)
	DeleteFile(ctx context.Context, id uuid.UUID, fileID uuid.UUID) (bool, *resty.Response, error)
	DeleteFiles(ctx context.Context, id uuid.UUID, files []MetaWrapper) (*DeleteManyResponse, *resty.Response, error)
}

func NewServiceService(client *Client) ServiceService {
	e := NewEndpoint(client, "entity/service")
	return newMainService[Service, any, any, any](e)
}
