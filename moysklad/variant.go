package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Variant Модификация.
// Ключевое слово: variant
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-modifikaciq
type Variant struct {
	Archived           *bool           `json:"archived,omitempty"`
	Updated            *Timestamp      `json:"updated,omitempty"`
	AccountID          *uuid.UUID      `json:"accountId,omitempty"`
	Description        *string         `json:"description,omitempty"`
	ExternalCode       *string         `json:"externalCode,omitempty"`
	ID                 *uuid.UUID      `json:"id,omitempty"`
	Meta               *Meta           `json:"meta,omitempty"`
	Name               *string         `json:"name,omitempty"`
	Code               *string         `json:"code,omitempty"`
	Barcodes           Barcodes        `json:"barcodes,omitempty"`
	DiscountProhibited *bool           `json:"discountProhibited,omitempty"`
	Characteristics    Characteristics `json:"characteristics,omitempty"`
	Images             *Images         `json:"images,omitempty"`
	MinPrice           *MinPrice       `json:"minPrice,omitempty"`
	BuyPrice           *BuyPrice       `json:"buyPrice,omitempty"`
	Product            *Product        `json:"product,omitempty"`
	SalePrices         SalePrices      `json:"salePrices,omitempty"`
	Things             Things          `json:"things,omitempty"`
	Packs              VariantPacks    `json:"packs,omitempty"`
}

func NewVariantFromAssortment(assortmentPosition AssortmentPosition) *Variant {
	return unmarshalAsType[Variant](assortmentPosition)
}

func (variant Variant) FromAssortment(assortmentPosition AssortmentPosition) *Variant {
	return unmarshalAsType[Variant](assortmentPosition)
}

func (variant Variant) GetArchived() bool {
	return Deref(variant.Archived)
}

func (variant Variant) GetUpdated() Timestamp {
	return Deref(variant.Updated)
}

func (variant Variant) GetAccountID() uuid.UUID {
	return Deref(variant.AccountID)
}

func (variant Variant) GetDescription() string {
	return Deref(variant.Description)
}

func (variant Variant) GetExternalCode() string {
	return Deref(variant.ExternalCode)
}

func (variant Variant) GetID() uuid.UUID {
	return Deref(variant.ID)
}

func (variant Variant) GetMeta() Meta {
	return Deref(variant.Meta)
}

func (variant Variant) GetName() string {
	return Deref(variant.Name)
}

func (variant Variant) GetCode() string {
	return Deref(variant.Code)
}

func (variant Variant) GetBarcodes() Barcodes {
	return variant.Barcodes
}

func (variant Variant) GetDiscountProhibited() bool {
	return Deref(variant.DiscountProhibited)
}

func (variant Variant) GetCharacteristics() Characteristics {
	return variant.Characteristics
}

func (variant Variant) GetImages() Images {
	return Deref(variant.Images)
}

func (variant Variant) GetMinPrice() MinPrice {
	return Deref(variant.MinPrice)
}

func (variant Variant) GetBuyPrice() BuyPrice {
	return Deref(variant.BuyPrice)
}

func (variant Variant) GetProduct() Product {
	return Deref(variant.Product)
}

func (variant Variant) GetSalePrices() SalePrices {
	return variant.SalePrices
}

func (variant Variant) GetThings() Things {
	return variant.Things
}

func (variant Variant) GetPacks() VariantPacks {
	return variant.Packs
}

func (variant *Variant) SetArchived(archived bool) *Variant {
	variant.Archived = &archived
	return variant
}

func (variant *Variant) SetDescription(description string) *Variant {
	variant.Description = &description
	return variant
}

func (variant *Variant) SetExternalCode(externalCode string) *Variant {
	variant.ExternalCode = &externalCode
	return variant
}

func (variant *Variant) SetMeta(meta *Meta) *Variant {
	variant.Meta = meta
	return variant
}

func (variant *Variant) SetName(name string) *Variant {
	variant.Name = &name
	return variant
}

func (variant *Variant) SetCode(code string) *Variant {
	variant.Code = &code
	return variant
}

func (variant *Variant) SetBarcodes(barcodes Barcodes) *Variant {
	variant.Barcodes = barcodes
	return variant
}

func (variant *Variant) SetDiscountProhibited(discountProhibited bool) *Variant {
	variant.DiscountProhibited = &discountProhibited
	return variant
}

func (variant *Variant) SetCharacteristics(characteristics Characteristics) *Variant {
	variant.Characteristics = characteristics
	return variant
}

func (variant *Variant) SetImages(images *Images) *Variant {
	variant.Images = images
	return variant
}

func (variant *Variant) SetMinPrice(minPrice *MinPrice) *Variant {
	variant.MinPrice = minPrice
	return variant
}

func (variant *Variant) SetBuyPrice(buyPrice *BuyPrice) *Variant {
	variant.BuyPrice = buyPrice
	return variant
}

func (variant *Variant) SetProduct(product *Product) *Variant {
	variant.Product = product
	return variant
}

func (variant *Variant) SetSalePrices(salePrices SalePrices) *Variant {
	variant.SalePrices = salePrices
	return variant
}

func (variant *Variant) SetPacks(packs VariantPacks) *Variant {
	variant.Packs = packs
	return variant
}

func (variant Variant) String() string {
	return Stringify(variant)
}

func (variant Variant) MetaType() MetaType {
	return MetaTypeVariant
}

// VariantPack Упаковка модификации.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-modifikaciq-modifikacii-atributy-wlozhennyh-suschnostej-upakowki-modifikacii
type VariantPack struct {
	ID         *uuid.UUID `json:"id,omitempty"`
	ParentPack *Pack      `json:"parentpack,omitempty"`
	Barcodes   Barcodes   `json:"barcodes,omitempty"`
}

type VariantPacks = Slice[VariantPack]

func (variantPack VariantPack) GetID() uuid.UUID {
	return Deref(variantPack.ID)
}

func (variantPack VariantPack) GetParentPack() Pack {
	return Deref(variantPack.ParentPack)
}

func (variantPack VariantPack) GetBarcodes() Barcodes {
	return variantPack.Barcodes
}

func (variantPack *VariantPack) SetParentPack(pack *Pack) *VariantPack {
	variantPack.ParentPack = pack
	return variantPack
}

func (variantPack *VariantPack) SetBarcodes(barcodes Barcodes) *VariantPack {
	variantPack.Barcodes = barcodes
	return variantPack
}

func (variantPack VariantPack) String() string {
	return Stringify(variantPack)
}

// Characteristic Характеристика
// Ключевое слово: attributemetadata
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-modifikaciq-modifikacii-atributy-wlozhennyh-suschnostej-metadannye-modifikacij-harakteristiki-modifikacii
type Characteristic struct {
	ID       *uuid.UUID `json:"id,omitempty"`       // ID соответствующей характеристики
	Meta     *Meta      `json:"meta,omitempty"`     // Метаданные характеристики
	Name     *string    `json:"name,omitempty"`     // Наименование характеристики
	Required *bool      `json:"required,omitempty"` // Флаг о том, является ли характеристика обязательной
	Type     *string    `json:"type,omitempty"`     // Тип значения характеристики (значение всегда "string")
	Value    *string    `json:"value,omitempty"`    // Значение характеристики
}

type Characteristics = Slice[Characteristic]

func (characteristic Characteristic) GetID() uuid.UUID {
	return Deref(characteristic.ID)
}

func (characteristic Characteristic) GetMeta() Meta {
	return Deref(characteristic.Meta)
}

func (characteristic Characteristic) GetName() string {
	return Deref(characteristic.Name)
}

func (characteristic Characteristic) GetRequired() bool {
	return Deref(characteristic.Required)
}

func (characteristic Characteristic) GetType() string {
	return Deref(characteristic.Type)
}

func (characteristic Characteristic) GetValue() string {
	return Deref(characteristic.Value)
}

func (characteristic *Characteristic) SetMeta(meta *Meta) *Characteristic {
	characteristic.Meta = meta
	return characteristic
}

func (characteristic *Characteristic) SetName(name string) *Characteristic {
	characteristic.Name = &name
	return characteristic
}

func (characteristic *Characteristic) SetRequired(required bool) *Characteristic {
	characteristic.Required = &required
	return characteristic
}

func (characteristic *Characteristic) SetType(value string) *Characteristic {
	characteristic.Type = &value
	return characteristic
}

func (characteristic *Characteristic) SetValue(value string) *Characteristic {
	characteristic.Value = &value
	return characteristic
}

func (characteristic Characteristic) String() string {
	return Stringify(characteristic)
}

func (characteristic Characteristic) MetaType() MetaType {
	return MetaTypeCharacteristic
}

// VariantService
// Сервис для работы с модификациями.
type VariantService interface {
	GetList(ctx context.Context, params *Params) (*List[Variant], *resty.Response, error)
	Create(ctx context.Context, variant *Variant, params *Params) (*Variant, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, variantList []*Variant, params *Params) (*[]Variant, *resty.Response, error)
	DeleteMany(ctx context.Context, variantList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Variant, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, variant *Variant, params *Params) (*Variant, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetadataVariant, *resty.Response, error)
	GetImages(ctx context.Context, id *uuid.UUID) (*MetaArray[Image], *resty.Response, error)
	CreateImage(ctx context.Context, id *uuid.UUID, image *Image) (*[]*Image, *resty.Response, error)
	UpdateImages(ctx context.Context, id *uuid.UUID, images []*Image) (*[]Image, *resty.Response, error)
	DeleteImage(ctx context.Context, id *uuid.UUID, imageId *uuid.UUID) (bool, *resty.Response, error)
	DeleteImages(ctx context.Context, id *uuid.UUID, images *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
	CreateCharacteristic(ctx context.Context, characteristic *Characteristic) (*Characteristic, *resty.Response, error)
	CreateCharacteristics(ctx context.Context, characteristics []*Characteristic) (*[]Characteristic, *resty.Response, error)
	GetCharacteristicById(ctx context.Context, id *uuid.UUID) (*Characteristic, *resty.Response, error)
	UpdateCharacteristic(ctx context.Context, id *uuid.UUID, characteristic *Characteristic) (*Characteristic, *resty.Response, error)
	DeleteCharacteristic(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

type variantService struct {
	Endpoint
	endpointGetList[Variant]
	endpointCreate[Variant]
	endpointCreateUpdateMany[Variant]
	endpointDeleteMany[Variant]
	endpointDelete
	endpointGetById[Variant]
	endpointUpdate[Variant]
	endpointMetadata[MetadataVariant]
	endpointImages
	endpointNamedFilter
}

func NewVariantService(client *Client) VariantService {
	e := NewEndpoint(client, "entity/variant")
	return &variantService{
		Endpoint:                 e,
		endpointGetList:          endpointGetList[Variant]{e},
		endpointCreate:           endpointCreate[Variant]{e},
		endpointCreateUpdateMany: endpointCreateUpdateMany[Variant]{e},
		endpointDeleteMany:       endpointDeleteMany[Variant]{e},
		endpointDelete:           endpointDelete{e},
		endpointGetById:          endpointGetById[Variant]{e},
		endpointUpdate:           endpointUpdate[Variant]{e},
		endpointMetadata:         endpointMetadata[MetadataVariant]{e},
		endpointImages:           endpointImages{e},
		endpointNamedFilter:      endpointNamedFilter{e},
	}
}

// CreateCharacteristic Создать характеристику.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-harakteristiki-modifikacij-sozdat-harakteristiku
func (s *variantService) CreateCharacteristic(ctx context.Context, characteristic *Characteristic) (*Characteristic, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/characteristics", s.uri)
	return NewRequestBuilder[Characteristic](s.client, path).Post(ctx, characteristic)
}

// CreateCharacteristics Массовое создание Характеристик.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-harakteristiki-modifikacij-massowoe-sozdanie-harakteristik
func (s *variantService) CreateCharacteristics(ctx context.Context, characteristics []*Characteristic) (*[]Characteristic, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/characteristics", s.uri)
	return NewRequestBuilder[[]Characteristic](s.client, path).Post(ctx, characteristics)
}

// GetCharacteristicById Получить Характеристику.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-harakteristiki-modifikacij-poluchit-harakteristiku
func (s *variantService) GetCharacteristicById(ctx context.Context, id *uuid.UUID) (*Characteristic, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/characteristics/%s", s.uri, id)
	return NewRequestBuilder[Characteristic](s.client, path).Get(ctx)
}

// UpdateCharacteristic Изменить характеристику.
func (s *variantService) UpdateCharacteristic(ctx context.Context, id *uuid.UUID, characteristic *Characteristic) (*Characteristic, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/characteristics/%s", s.uri, id)
	return NewRequestBuilder[Characteristic](s.client, path).Put(ctx, characteristic)
}

// DeleteCharacteristic Удалить характеристику.
func (s *variantService) DeleteCharacteristic(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("%s/metadata/characteristics/%s", s.uri, id)
	return NewRequestBuilder[any](s.client, path).Delete(ctx)
}
