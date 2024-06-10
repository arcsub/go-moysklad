package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Product Товар.
// Ключевое слово: product
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-towar
type Product struct {
	Supplier            *Counterparty         `json:"supplier,omitempty"`
	OnTap               *bool                 `json:"onTap,omitempty"`
	Code                *string               `json:"code,omitempty"`
	Description         *string               `json:"description,omitempty"`
	ExternalCode        *string               `json:"externalCode,omitempty"`
	ID                  *uuid.UUID            `json:"id,omitempty"`
	Meta                *Meta                 `json:"meta,omitempty"`
	Name                *string               `json:"name,omitempty"`
	Alcoholic           *Alcoholic            `json:"alcoholic,omitempty"`
	Archived            *bool                 `json:"archived,omitempty"`
	Article             *string               `json:"article,omitempty"`
	PaymentItemType     PaymentItem           `json:"paymentItemType,omitempty"`
	BuyPrice            *BuyPrice             `json:"buyPrice,omitempty"`
	Country             *Country              `json:"country,omitempty"`
	DiscountProhibited  *bool                 `json:"discountProhibited,omitempty"`
	EffectiveVat        *int                  `json:"effectiveVat,omitempty"`
	EffectiveVatEnabled *bool                 `json:"effectiveVatEnabled,omitempty"`
	Files               *MetaArray[File]      `json:"files,omitempty"`
	Group               *Group                `json:"group,omitempty"`
	Images              *MetaArray[Image]     `json:"images,omitempty"`
	IsSerialTrackable   *bool                 `json:"isSerialTrackable,omitempty"`
	MinPrice            *MinPrice             `json:"minPrice,omitempty"`
	TaxSystem           GoodTaxSystem         `json:"taxSystem,omitempty"`
	UseParentVat        *bool                 `json:"useParentVat,omitempty"`
	Owner               *Employee             `json:"owner,omitempty"`
	Packs               Slice[Pack]           `json:"packs,omitempty"`
	PartialDisposal     *bool                 `json:"partialDisposal,omitempty"`
	PathName            *string               `json:"pathName,omitempty"`
	Weight              *float64              `json:"weight,omitempty"`
	PpeType             *string               `json:"ppeType,omitempty"`
	ProductFolder       *ProductFolder        `json:"productFolder,omitempty"`
	SalePrices          Slice[SalePrice]      `json:"salePrices,omitempty"`
	Shared              *bool                 `json:"shared,omitempty"`
	MinimumBalance      *float64              `json:"minimumBalance,omitempty"`
	SyncID              *uuid.UUID            `json:"syncId,omitempty"`
	AccountID           *uuid.UUID            `json:"accountId,omitempty"`
	Things              Slice[string]         `json:"things,omitempty"`
	Tnved               *string               `json:"tnved,omitempty"`
	VatEnabled          *bool                 `json:"vatEnabled,omitempty"`
	Uom                 *Uom                  `json:"uom,omitempty"`
	Updated             *Timestamp            `json:"updated,omitempty"`
	Barcodes            Slice[Barcode]        `json:"barcodes,omitempty"`
	VariantsCount       *int                  `json:"variantsCount,omitempty"`
	Vat                 *int                  `json:"vat,omitempty"`
	TrackingType        TrackingType          `json:"trackingType,omitempty"`
	Volume              *float64              `json:"volume,omitempty"`
	Attributes          Slice[AttributeValue] `json:"attributes,omitempty"`
}

func (product Product) Clean() *Product {
	return &Product{Meta: product.Meta}
}

func NewProductFromAssortment(assortmentPosition AssortmentPosition) *Product {
	return unmarshalAsType[Product](assortmentPosition)
}

func (product Product) FromAssortment(assortmentPosition AssortmentPosition) *Product {
	return unmarshalAsType[Product](assortmentPosition)
}

func (product Product) AsAssortment() *AssortmentPosition {
	return &AssortmentPosition{Meta: product.GetMeta()}
}

func (product Product) GetSupplier() Counterparty {
	return Deref(product.Supplier)
}

func (product Product) GetOnTap() bool {
	return Deref(product.OnTap)
}

func (product Product) GetCode() string {
	return Deref(product.Code)
}

func (product Product) GetDescription() string {
	return Deref(product.Description)
}

func (product Product) GetExternalCode() string {
	return Deref(product.ExternalCode)
}

func (product Product) GetID() uuid.UUID {
	return Deref(product.ID)
}

func (product Product) GetMeta() Meta {
	return Deref(product.Meta)
}

func (product Product) GetName() string {
	return Deref(product.Name)
}

func (product Product) GetAlcoholic() Alcoholic {
	return Deref(product.Alcoholic)
}

func (product Product) GetArchived() bool {
	return Deref(product.Archived)
}

func (product Product) GetArticle() string {
	return Deref(product.Article)
}

func (product Product) GetPaymentItemType() PaymentItem {
	return product.PaymentItemType
}

func (product Product) GetBuyPrice() BuyPrice {
	return Deref(product.BuyPrice)
}

func (product Product) GetCountry() Country {
	return Deref(product.Country)
}

func (product Product) GetDiscountProhibited() bool {
	return Deref(product.DiscountProhibited)
}

func (product Product) GetEffectiveVat() int {
	return Deref(product.EffectiveVat)
}

func (product Product) GetEffectiveVatEnabled() bool {
	return Deref(product.EffectiveVatEnabled)
}

func (product Product) GetFiles() MetaArray[File] {
	return Deref(product.Files)
}

func (product Product) GetGroup() Group {
	return Deref(product.Group)
}

func (product Product) GetImages() MetaArray[Image] {
	return Deref(product.Images)
}

func (product Product) GetIsSerialTrackable() bool {
	return Deref(product.IsSerialTrackable)
}

func (product Product) GetMinPrice() MinPrice {
	return Deref(product.MinPrice)
}

func (product Product) GetTaxSystem() GoodTaxSystem {
	return product.TaxSystem
}

func (product Product) GetUseParentVat() bool {
	return Deref(product.UseParentVat)
}

func (product Product) GetOwner() Employee {
	return Deref(product.Owner)
}

func (product Product) GetPacks() Slice[Pack] {
	return product.Packs
}

func (product Product) GetPartialDisposal() bool {
	return Deref(product.PartialDisposal)
}

func (product Product) GetPathName() string {
	return Deref(product.PathName)
}

func (product Product) GetWeight() float64 {
	return Deref(product.Weight)
}

func (product Product) GetPpeType() string {
	return Deref(product.PpeType)
}

func (product Product) GetProductFolder() ProductFolder {
	return Deref(product.ProductFolder)
}

func (product Product) GetSalePrices() Slice[SalePrice] {
	return product.SalePrices
}

func (product Product) GetShared() bool {
	return Deref(product.Shared)
}

func (product Product) GetMinimumBalance() float64 {
	return Deref(product.MinimumBalance)
}

func (product Product) GetSyncID() uuid.UUID {
	return Deref(product.SyncID)
}

func (product Product) GetAccountID() uuid.UUID {
	return Deref(product.AccountID)
}

func (product Product) GetThings() Slice[string] {
	return product.Things
}

func (product Product) GetTnved() string {
	return Deref(product.Tnved)
}

func (product Product) GetVatEnabled() bool {
	return Deref(product.VatEnabled)
}

func (product Product) GetUom() Uom {
	return Deref(product.Uom)
}

func (product Product) GetUpdated() Timestamp {
	return Deref(product.Updated)
}

func (product Product) GetBarcodes() Slice[Barcode] {
	return product.Barcodes
}

func (product Product) GetVariantsCount() int {
	return Deref(product.VariantsCount)
}

func (product Product) GetVat() int {
	return Deref(product.Vat)
}

func (product Product) GetTrackingType() TrackingType {
	return product.TrackingType
}

func (product Product) GetVolume() float64 {
	return Deref(product.Volume)
}

func (product Product) GetAttributes() Slice[AttributeValue] {
	return product.Attributes
}

func (product *Product) SetSupplier(supplier *Counterparty) *Product {
	product.Supplier = supplier
	return product
}

func (product *Product) SetOnTap(onTap bool) *Product {
	product.OnTap = &onTap
	return product
}

func (product *Product) SetCode(code string) *Product {
	product.Code = &code
	return product
}

func (product *Product) SetDescription(description string) *Product {
	product.Description = &description
	return product
}

func (product *Product) SetExternalCode(externalCode string) *Product {
	product.ExternalCode = &externalCode
	return product
}

func (product *Product) SetMeta(meta *Meta) *Product {
	product.Meta = meta
	return product
}

func (product *Product) SetName(name string) *Product {
	product.Name = &name
	return product
}

func (product *Product) SetAlcoholic(alcoholic *Alcoholic) *Product {
	product.Alcoholic = alcoholic
	return product
}

func (product *Product) SetArchived(archived bool) *Product {
	product.Archived = &archived
	return product
}

func (product *Product) SetArticle(article string) *Product {
	product.Article = &article
	return product
}

func (product *Product) SetPaymentItemType(paymentItem PaymentItem) *Product {
	product.PaymentItemType = paymentItem
	return product
}

func (product *Product) SetBuyPrice(buyPrice *BuyPrice) *Product {
	product.BuyPrice = buyPrice
	return product
}

func (product *Product) SetCountry(country *Country) *Product {
	product.Country = country
	return product
}

func (product *Product) SetDiscountProhibited(discountProhibited bool) *Product {
	product.DiscountProhibited = &discountProhibited
	return product
}

func (product *Product) SetFiles(files Slice[File]) *Product {
	product.Files = NewMetaArrayRows(files)
	return product
}

func (product *Product) SetGroup(group *Group) *Product {
	product.Group = group
	return product
}

func (product *Product) SetImages(images Slice[Image]) *Product {
	product.Images = NewMetaArrayRows(images)
	return product
}

func (product *Product) SetIsSerialTrackable(isSerialTrackable bool) *Product {
	product.IsSerialTrackable = &isSerialTrackable
	return product
}

func (product *Product) SetMinPrice(minPrice *MinPrice) *Product {
	product.MinPrice = minPrice
	return product
}

func (product *Product) SetTaxSystem(taxSystem GoodTaxSystem) *Product {
	product.TaxSystem = taxSystem
	return product
}

func (product *Product) SetUseParentVat(useParentVat bool) *Product {
	product.UseParentVat = &useParentVat
	return product
}

func (product *Product) SetOwner(owner *Employee) *Product {
	product.Owner = owner
	return product
}

func (product *Product) SetPacks(packs Slice[Pack]) *Product {
	product.Packs = packs
	return product
}

func (product *Product) SetPartialDisposal(partialDisposal bool) *Product {
	product.PartialDisposal = &partialDisposal
	return product
}

func (product *Product) SetWeight(weight float64) *Product {
	product.Weight = &weight
	return product
}

func (product *Product) SetPpeType(ppeType string) *Product {
	product.PpeType = &ppeType
	return product
}

func (product *Product) SetProductFolder(productFolder *ProductFolder) *Product {
	product.ProductFolder = productFolder
	return product
}

func (product *Product) SetSalePrices(salePrices Slice[SalePrice]) *Product {
	product.SalePrices = salePrices
	return product
}

func (product *Product) SetShared(shared bool) *Product {
	product.Shared = &shared
	return product
}

func (product *Product) SetMinimumBalance(minimumBalance float64) *Product {
	product.MinimumBalance = &minimumBalance
	return product
}

func (product *Product) SetSyncID(syncID uuid.UUID) *Product {
	product.SyncID = &syncID
	return product
}

func (product *Product) SetThings(things Slice[string]) *Product {
	product.Things = things
	return product
}

func (product *Product) SetTnved(tnved string) *Product {
	product.Tnved = &tnved
	return product
}

func (product *Product) SetVatEnabled(vatEnabled bool) *Product {
	product.VatEnabled = &vatEnabled
	return product
}

func (product *Product) SetUom(uom *Uom) *Product {
	product.Uom = uom
	return product
}

func (product *Product) SetBarcodes(barcodes Slice[Barcode]) *Product {
	product.Barcodes = barcodes
	return product
}

func (product *Product) SetVat(vat int) *Product {
	product.Vat = &vat
	return product
}

func (product *Product) SetTrackingType(trackingType TrackingType) *Product {
	product.TrackingType = trackingType
	return product
}

func (product *Product) SetVolume(volume float64) *Product {
	product.Volume = &volume
	return product
}

func (product *Product) SetAttributes(attributes Slice[AttributeValue]) *Product {
	product.Attributes = attributes
	return product
}

func (product Product) String() string {
	return Stringify(product)
}

func (product Product) MetaType() MetaType {
	return MetaTypeProduct
}

// Alcoholic Объект, содержащий поля алкогольной продукции
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-towar-towary-atributy-wlozhennyh-suschnostej-ob-ekt-soderzhaschij-polq-alkogol-noj-produkcii
type Alcoholic struct {
	Excise   *bool    `json:"excise,omitempty"`   // Содержит акцизную марку
	Type     *int     `json:"type,omitempty"`     // Код вида продукции
	Strength *float64 `json:"strength,omitempty"` // Крепость
	Volume   *float64 `json:"volume,omitempty"`   // Объём тары
}

func (alcoholic Alcoholic) GetExcise() bool {
	return Deref(alcoholic.Excise)
}

func (alcoholic Alcoholic) GetType() int {
	return Deref(alcoholic.Type)
}

func (alcoholic Alcoholic) GetStrength() float64 {
	return Deref(alcoholic.Strength)
}

func (alcoholic Alcoholic) GetVolume() float64 {
	return Deref(alcoholic.Volume)
}

func (alcoholic *Alcoholic) SetExcise(excise bool) *Alcoholic {
	alcoholic.Excise = &excise
	return alcoholic
}

func (alcoholic *Alcoholic) SetType(value int) *Alcoholic {
	alcoholic.Type = &value
	return alcoholic
}

func (alcoholic *Alcoholic) SetStrength(strength float64) *Alcoholic {
	alcoholic.Strength = &strength
	return alcoholic
}

func (alcoholic *Alcoholic) SetVolume(volume float64) *Alcoholic {
	alcoholic.Volume = &volume
	return alcoholic
}

func (alcoholic Alcoholic) String() string {
	return Stringify(alcoholic)
}

// ProductService
// Сервис для работы с товарами.
type ProductService interface {
	GetList(ctx context.Context, params *Params) (*List[Product], *resty.Response, error)
	Create(ctx context.Context, product *Product, params *Params) (*Product, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, productList []*Product, params *Params) (*[]Product, *resty.Response, error)
	DeleteMany(ctx context.Context, productList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetaAttributesSharedWrapper, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error)
	UpdateAttribute(ctx context.Context, id uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params *Params) (*Product, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, product *Product, params *Params) (*Product, *resty.Response, error)
	GetImages(ctx context.Context, id uuid.UUID) (*MetaArray[Image], *resty.Response, error)
	CreateImage(ctx context.Context, id uuid.UUID, image *Image) (*[]*Image, *resty.Response, error)
	UpdateImages(ctx context.Context, id uuid.UUID, images []*Image) (*[]Image, *resty.Response, error)
	DeleteImage(ctx context.Context, id uuid.UUID, imageId uuid.UUID) (bool, *resty.Response, error)
	DeleteImages(ctx context.Context, id uuid.UUID, images *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*Product, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)
	GetAudit(ctx context.Context, id uuid.UUID, params *Params) (*List[AuditEvent], *resty.Response, error)
	PrintLabel(ctx context.Context, id uuid.UUID, PrintLabelArg *PrintLabelArg) (*PrintFile, *resty.Response, error)
}

func NewProductService(client *Client) ProductService {
	e := NewEndpoint(client, "entity/product")
	return newMainService[Product, any, MetaAttributesSharedWrapper, any](e)
}
