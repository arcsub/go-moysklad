package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Bundle Комплект.
// Ключевое слово: bundle
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-komplekt
type Bundle struct {
	Volume              *float64                    `json:"volume,omitempty"`
	SyncID              *uuid.UUID                  `json:"syncId,omitempty"`
	Code                *string                     `json:"code,omitempty"`
	Description         *string                     `json:"description,omitempty"`
	ExternalCode        *string                     `json:"externalCode,omitempty"`
	ID                  *uuid.UUID                  `json:"id,omitempty"`
	Meta                *Meta                       `json:"meta,omitempty"`
	Name                *string                     `json:"name,omitempty"`
	Archived            *bool                       `json:"archived,omitempty"`
	Article             *string                     `json:"article,omitempty"`
	Images              *MetaArray[Image]           `json:"images,omitempty"`
	Components          *Positions[BundleComponent] `json:"components,omitempty"`
	Country             *Country                    `json:"country,omitempty"`
	DiscountProhibited  *bool                       `json:"discountProhibited"`
	EffectiveVat        *int                        `json:"effectiveVat,omitempty"`
	EffectiveVatEnabled *bool                       `json:"effectiveVatEnabled,omitempty"`
	Files               *MetaArray[File]            `json:"files,omitempty"`
	Group               *Group                      `json:"group,omitempty"`
	Vat                 *int                        `json:"vat,omitempty"`
	MinPrice            *MinPrice                   `json:"minPrice,omitempty"`
	Overhead            *BundleOverhead             `json:"overhead,omitempty"`
	Owner               *Employee                   `json:"owner,omitempty"`
	PartialDisposal     *bool                       `json:"partialDisposal,omitempty"`
	PathName            *string                     `json:"pathName,omitempty"`
	Weight              *float64                    `json:"weight,omitempty"`
	ProductFolder       *ProductFolder              `json:"productFolder,omitempty"`
	Shared              *bool                       `json:"shared,omitempty"`
	Updated             *Timestamp                  `json:"updated,omitempty"`
	AccountID           *uuid.UUID                  `json:"accountId,omitempty"`
	Tnved               *string                     `json:"tnved,omitempty"`
	VatEnabled          *bool                       `json:"vatEnabled,omitempty"`
	Uom                 *Uom                        `json:"uom,omitempty"`
	UseParentVat        *bool                       `json:"useParentVat,omitempty"`
	TaxSystem           GoodTaxSystem               `json:"taxSystem,omitempty"`
	TrackingType        TrackingType                `json:"trackingType,omitempty"`
	PaymentItemType     PaymentItem                 `json:"paymentItemType,omitempty"`
	Barcodes            Slice[Barcode]              `json:"barcodes,omitempty"`
	SalePrices          Slice[SalePrice]            `json:"salePrices,omitempty"`
	Attributes          Slice[AttributeValue]       `json:"attributes,omitempty"`
}

func NewBundleFromAssortment(assortmentPosition AssortmentPosition) *Bundle {
	return unmarshalAsType[Bundle](assortmentPosition)
}

func (bundle Bundle) FromAssortment(assortmentPosition AssortmentPosition) *Bundle {
	return unmarshalAsType[Bundle](assortmentPosition)
}

func (bundle Bundle) AsAssortment() *AssortmentPosition {
	return &AssortmentPosition{Meta: bundle.GetMeta()}
}

func (bundle Bundle) GetVolume() float64 {
	return Deref(bundle.Volume)
}

func (bundle Bundle) GetSyncID() uuid.UUID {
	return Deref(bundle.SyncID)
}

func (bundle Bundle) GetCode() string {
	return Deref(bundle.Code)
}

func (bundle Bundle) GetDescription() string {
	return Deref(bundle.Description)
}

func (bundle Bundle) GetExternalCode() string {
	return Deref(bundle.ExternalCode)
}

func (bundle Bundle) GetID() uuid.UUID {
	return Deref(bundle.ID)
}

func (bundle Bundle) GetMeta() Meta {
	return Deref(bundle.Meta)
}

func (bundle Bundle) GetName() string {
	return Deref(bundle.Name)
}

func (bundle Bundle) GetArchived() bool {
	return Deref(bundle.Archived)
}

func (bundle Bundle) GetArticle() string {
	return Deref(bundle.Article)
}

func (bundle Bundle) GetImages() MetaArray[Image] {
	return Deref(bundle.Images)
}

func (bundle Bundle) GetComponents() Positions[BundleComponent] {
	return Deref(bundle.Components)
}

func (bundle Bundle) GetCountry() Country {
	return Deref(bundle.Country)
}

func (bundle Bundle) GetDiscountProhibited() bool {
	return Deref(bundle.DiscountProhibited)
}

func (bundle Bundle) GetEffectiveVat() int {
	return Deref(bundle.EffectiveVat)
}

func (bundle Bundle) GetEffectiveVatEnabled() bool {
	return Deref(bundle.EffectiveVatEnabled)
}

func (bundle Bundle) GetFiles() MetaArray[File] {
	return Deref(bundle.Files)
}

func (bundle Bundle) GetGroup() Group {
	return Deref(bundle.Group)
}

func (bundle Bundle) GetVat() int {
	return Deref(bundle.Vat)
}

func (bundle Bundle) GetMinPrice() MinPrice {
	return Deref(bundle.MinPrice)
}

func (bundle Bundle) GetOverhead() BundleOverhead {
	return Deref(bundle.Overhead)
}

func (bundle Bundle) GetOwner() Employee {
	return Deref(bundle.Owner)
}

func (bundle Bundle) GetPartialDisposal() bool {
	return Deref(bundle.PartialDisposal)
}

func (bundle Bundle) GetPathName() string {
	return Deref(bundle.PathName)
}

func (bundle Bundle) GetWeight() float64 {
	return Deref(bundle.Weight)
}

func (bundle Bundle) GetSalePrices() Slice[SalePrice] {
	return bundle.SalePrices
}

func (bundle Bundle) GetProductFolder() ProductFolder {
	return Deref(bundle.ProductFolder)
}

func (bundle Bundle) GetShared() bool {
	return Deref(bundle.Shared)
}

func (bundle Bundle) GetUpdated() Timestamp {
	return Deref(bundle.Updated)
}

func (bundle Bundle) GetAccountID() uuid.UUID {
	return Deref(bundle.AccountID)
}

func (bundle Bundle) GetTnved() string {
	return Deref(bundle.Tnved)
}

func (bundle Bundle) GetVatEnabled() bool {
	return Deref(bundle.VatEnabled)
}

func (bundle Bundle) GetUom() Uom {
	return Deref(bundle.Uom)
}

func (bundle Bundle) GetBarcodes() Slice[Barcode] {
	return bundle.Barcodes
}

func (bundle Bundle) GetUseParentVat() bool {
	return Deref(bundle.UseParentVat)
}

func (bundle Bundle) GetTaxSystem() GoodTaxSystem {
	return bundle.TaxSystem
}

func (bundle Bundle) GetTrackingType() TrackingType {
	return bundle.TrackingType
}

func (bundle Bundle) GetPaymentItemType() PaymentItem {
	return bundle.PaymentItemType
}

func (bundle Bundle) GetAttributes() Slice[AttributeValue] {
	return bundle.Attributes
}

func (bundle *Bundle) SetVolume(volume float64) *Bundle {
	bundle.Volume = &volume
	return bundle
}

func (bundle *Bundle) SetSyncID(syncID uuid.UUID) *Bundle {
	bundle.SyncID = &syncID
	return bundle
}

func (bundle *Bundle) SetCode(code string) *Bundle {
	bundle.Code = &code
	return bundle
}

func (bundle *Bundle) SetDescription(description string) *Bundle {
	bundle.Description = &description
	return bundle
}

func (bundle *Bundle) SetExternalCode(externalCode string) *Bundle {
	bundle.ExternalCode = &externalCode
	return bundle
}

func (bundle *Bundle) SetMeta(meta *Meta) *Bundle {
	bundle.Meta = meta
	return bundle
}

func (bundle *Bundle) SetName(name string) *Bundle {
	bundle.Name = &name
	return bundle
}

func (bundle *Bundle) SetArchived(archived bool) *Bundle {
	bundle.Archived = &archived
	return bundle
}

func (bundle *Bundle) SetArticle(article string) *Bundle {
	bundle.Article = &article
	return bundle
}

func (bundle *Bundle) SetImages(images Slice[Image]) *Bundle {
	bundle.Images = NewMetaArrayRows(images)
	return bundle
}

func (bundle *Bundle) SetComponents(components *Positions[BundleComponent]) *Bundle {
	bundle.Components = components
	return bundle
}

func (bundle *Bundle) SetCountry(country *Country) *Bundle {
	bundle.Country = country.Clean()
	return bundle
}

func (bundle *Bundle) SetDiscountProhibited(discountProhibited bool) *Bundle {
	bundle.DiscountProhibited = &discountProhibited
	return bundle
}

func (bundle *Bundle) SetFiles(files Slice[File]) *Bundle {
	bundle.Files = NewMetaArrayRows(files)
	return bundle
}

func (bundle *Bundle) SetGroup(group *Group) *Bundle {
	bundle.Group = group.Clean()
	return bundle
}

func (bundle *Bundle) SetVat(vat int) *Bundle {
	bundle.Vat = &vat
	return bundle
}

func (bundle *Bundle) SetMinPrice(minPrice *MinPrice) *Bundle {
	bundle.MinPrice = minPrice
	return bundle
}

func (bundle *Bundle) SetOverhead(overhead *BundleOverhead) *Bundle {
	bundle.Overhead = overhead
	return bundle
}

func (bundle *Bundle) SetOwner(owner *Employee) *Bundle {
	bundle.Owner = owner.Clean()
	return bundle
}

func (bundle *Bundle) SetPartialDisposal(partialDisposal bool) *Bundle {
	bundle.PartialDisposal = &partialDisposal
	return bundle
}

func (bundle *Bundle) SetWeight(weight float64) *Bundle {
	bundle.Weight = &weight
	return bundle
}

func (bundle *Bundle) SetSalePrices(salePrices Slice[SalePrice]) *Bundle {
	bundle.SalePrices = salePrices
	return bundle
}

func (bundle *Bundle) SetProductFolder(productFolder *ProductFolder) *Bundle {
	bundle.ProductFolder = productFolder.Clean()
	return bundle
}

func (bundle *Bundle) SetShared(shared bool) *Bundle {
	bundle.Shared = &shared
	return bundle
}

func (bundle *Bundle) SetTnved(tnved string) *Bundle {
	bundle.Tnved = &tnved
	return bundle
}

func (bundle *Bundle) SetVatEnabled(vatEnabled bool) *Bundle {
	bundle.VatEnabled = &vatEnabled
	return bundle
}

func (bundle *Bundle) SetUom(uom *Uom) *Bundle {
	bundle.Uom = uom.Clean()
	return bundle
}

func (bundle *Bundle) SetBarcodes(barcodes Slice[Barcode]) *Bundle {
	bundle.Barcodes = barcodes
	return bundle
}

func (bundle *Bundle) SetUseParentVat(useParentVat bool) *Bundle {
	bundle.UseParentVat = &useParentVat
	return bundle
}

func (bundle *Bundle) SetTaxSystem(taxSystem GoodTaxSystem) *Bundle {
	bundle.TaxSystem = taxSystem
	return bundle
}

func (bundle *Bundle) SetTrackingType(trackingType TrackingType) *Bundle {
	bundle.TrackingType = trackingType
	return bundle
}

func (bundle *Bundle) SetPaymentItemType(paymentItemType PaymentItem) *Bundle {
	bundle.PaymentItemType = paymentItemType
	return bundle
}

func (bundle *Bundle) SetAttributes(attributes Slice[AttributeValue]) *Bundle {
	bundle.Attributes = attributes
	return bundle
}

func (bundle Bundle) String() string {
	return Stringify(bundle)
}

func (bundle Bundle) MetaType() MetaType {
	return MetaTypeBundle
}

// BundleOverhead Дополнительные расходы
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-komplekt-komplekty-atributy-wlozhennyh-suschnostej-dopolnitel-nye-rashody
type BundleOverhead struct {
	Value    *float64  `json:"value,omitempty"`    // Значение цены
	Currency *Currency `json:"currency,omitempty"` // Ссылка на валюту в формате Метаданных
}

func (bundleOverhead BundleOverhead) GetValue() float64 {
	return Deref(bundleOverhead.Value)
}

func (bundleOverhead BundleOverhead) GetCurrency() Currency {
	return Deref(bundleOverhead.Currency)
}

func (bundleOverhead *BundleOverhead) SetValue(value *float64) *BundleOverhead {
	bundleOverhead.Value = value
	return bundleOverhead
}

func (bundleOverhead *BundleOverhead) SetCurrency(currency *Currency) *BundleOverhead {
	bundleOverhead.Currency = currency
	return bundleOverhead
}

func (bundleOverhead BundleOverhead) String() string {
	return Stringify(bundleOverhead)
}

// BundleComponent Компонент комплекта.
// Ключевое слово: bundlecomponent
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-komplekt-komplekty-komponenty-komplekta
type BundleComponent struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги, которую представляет собой компонент
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID сущности
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в компоненте
}

func (bundleComponent BundleComponent) GetAccountID() uuid.UUID {
	return Deref(bundleComponent.AccountID)
}

func (bundleComponent BundleComponent) GetAssortment() AssortmentPosition {
	return Deref(bundleComponent.Assortment)
}

func (bundleComponent BundleComponent) GetID() uuid.UUID {
	return Deref(bundleComponent.ID)
}

func (bundleComponent BundleComponent) GetQuantity() float64 {
	return Deref(bundleComponent.Quantity)
}

func (bundleComponent *BundleComponent) SetAssortment(assortment MetaOwner) *BundleComponent {
	bundleComponent.Assortment = &AssortmentPosition{Meta: assortment.GetMeta()}
	return bundleComponent
}

func (bundleComponent *BundleComponent) SetQuantity(quantity float64) *BundleComponent {
	bundleComponent.Quantity = &quantity
	return bundleComponent
}

func (bundleComponent BundleComponent) String() string {
	return Stringify(bundleComponent)
}

func (bundleComponent BundleComponent) MetaType() MetaType {
	return MetaTypeBundleComponent
}

// BundleService
// Сервис для работы с комплектами.
type BundleService interface {
	GetList(ctx context.Context, params *Params) (*List[Bundle], *resty.Response, error)
	Create(ctx context.Context, bundle *Bundle, params *Params) (*Bundle, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, bundleList Slice[Bundle], params *Params) (*Slice[Bundle], *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params *Params) (*Bundle, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, bundle *Bundle, params *Params) (*Bundle, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	DeleteMany(ctx context.Context, bundleList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	GetComponents(ctx context.Context, id uuid.UUID) (*List[BundleComponent], *resty.Response, error)
	CreateComponent(ctx context.Context, id uuid.UUID, bundleComponent *BundleComponent) (*BundleComponent, *resty.Response, error)
	GetComponentByID(ctx context.Context, id, componentID uuid.UUID) (*BundleComponent, *resty.Response, error)
	UpdateComponent(ctx context.Context, id, componentID uuid.UUID, bundleComponent *BundleComponent) (*BundleComponent, *resty.Response, error)
	DeleteComponent(ctx context.Context, id, componentID uuid.UUID) (bool, *resty.Response, error)
	GetFiles(ctx context.Context, id uuid.UUID) (*MetaArray[File], *resty.Response, error)
	CreateFile(ctx context.Context, id uuid.UUID, file *File) (*Slice[File], *resty.Response, error)
	UpdateFiles(ctx context.Context, id uuid.UUID, files Slice[File]) (*Slice[File], *resty.Response, error)
	DeleteFile(ctx context.Context, id uuid.UUID, fileID uuid.UUID) (bool, *resty.Response, error)
	DeleteFiles(ctx context.Context, id uuid.UUID, files *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	GetImages(ctx context.Context, id uuid.UUID) (*MetaArray[Image], *resty.Response, error)
	CreateImage(ctx context.Context, id uuid.UUID, image *Image) (*Slice[Image], *resty.Response, error)
	UpdateImages(ctx context.Context, id uuid.UUID, images Slice[Image]) (*Slice[Image], *resty.Response, error)
	DeleteImage(ctx context.Context, id uuid.UUID, imageID uuid.UUID) (bool, *resty.Response, error)
	DeleteImages(ctx context.Context, id uuid.UUID, images *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
}

type bundleService struct {
	Endpoint
	endpointGetList[Bundle]
	endpointCreate[Bundle]
	endpointCreateUpdateMany[Bundle]
	endpointGetByID[Bundle]
	endpointUpdate[Bundle]
	endpointDelete
	endpointDeleteMany[Bundle]
	endpointFiles
	endpointImages
}

func NewBundleService(client *Client) BundleService {
	e := NewEndpoint(client, "entity/bundle")
	return &bundleService{
		Endpoint:                 e,
		endpointGetList:          endpointGetList[Bundle]{e},
		endpointCreate:           endpointCreate[Bundle]{e},
		endpointCreateUpdateMany: endpointCreateUpdateMany[Bundle]{e},
		endpointGetByID:          endpointGetByID[Bundle]{e},
		endpointUpdate:           endpointUpdate[Bundle]{e},
		endpointDelete:           endpointDelete{e},
		endpointDeleteMany:       endpointDeleteMany[Bundle]{e},
		endpointFiles:            endpointFiles{e},
		endpointImages:           endpointImages{e},
	}
}

// GetComponents Получить компоненты Комплекта.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-komplekt-poluchit-komponenty-komplekta
func (service *bundleService) GetComponents(ctx context.Context, id uuid.UUID) (*List[BundleComponent], *resty.Response, error) {
	path := fmt.Sprintf("entity/bundle/%s/components", id)
	return NewRequestBuilder[List[BundleComponent]](service.client, path).Get(ctx)
}

// CreateComponent Добавить компонент Комплекта.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-komplekt-dobawit-komponent-komplekta
func (service *bundleService) CreateComponent(ctx context.Context, id uuid.UUID, bundleComponent *BundleComponent) (*BundleComponent, *resty.Response, error) {
	path := fmt.Sprintf("entity/bundle/%s/components", id)
	return NewRequestBuilder[BundleComponent](service.client, path).Post(ctx, bundleComponent)
}

// GetComponentByID Получить компонент.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-komplekt-poluchit-komponent
func (service *bundleService) GetComponentByID(ctx context.Context, id, componentID uuid.UUID) (*BundleComponent, *resty.Response, error) {
	path := fmt.Sprintf("entity/bundle/%s/components/%s", id, componentID)
	return NewRequestBuilder[BundleComponent](service.client, path).Get(ctx)
}

// UpdateComponent Изменить компонент.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-komplekt-izmenit-komponent
func (service *bundleService) UpdateComponent(ctx context.Context, id, componentID uuid.UUID, bundleComponent *BundleComponent) (*BundleComponent, *resty.Response, error) {
	path := fmt.Sprintf("entity/bundle/%s/components/%s", id, componentID)
	return NewRequestBuilder[BundleComponent](service.client, path).Put(ctx, bundleComponent)
}

// DeleteComponent Удалить компонент.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-komplekt-udalit-komponent
func (service *bundleService) DeleteComponent(ctx context.Context, id, componentID uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("entity/bundle/%s/components/%s", id, componentID)
	return NewRequestBuilder[any](service.client, path).Delete(ctx)
}
