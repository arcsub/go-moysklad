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
	Archived           *bool            `json:"archived,omitempty"`
	Updated            *Timestamp       `json:"updated,omitempty"`
	AccountID          *uuid.UUID       `json:"accountId,omitempty"`
	Description        *string          `json:"description,omitempty"`
	ExternalCode       *string          `json:"externalCode,omitempty"`
	ID                 *uuid.UUID       `json:"id,omitempty"`
	Meta               *Meta            `json:"meta,omitempty"`
	Name               *string          `json:"name,omitempty"`
	Code               *string          `json:"code,omitempty"`
	Barcodes           *Barcodes        `json:"barcodes,omitempty"`
	DiscountProhibited *bool            `json:"discountProhibited,omitempty"`
	Characteristics    *Characteristics `json:"characteristics,omitempty"`
	Images             *Images          `json:"images,omitempty"`
	MinPrice           *MinPrice        `json:"minPrice,omitempty"`
	BuyPrice           *BuyPrice        `json:"buyPrice,omitempty"`
	Product            *Product         `json:"product,omitempty"`
	SalePrices         *SalePrices      `json:"salePrices,omitempty"`
	Things             *Things          `json:"things,omitempty"`
	Packs              []VariantPack    `json:"packs,omitempty"`
}

func (v Variant) String() string {
	return Stringify(v)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (v Variant) GetMeta() *Meta {
	return v.Meta
}

func (v Variant) MetaType() MetaType {
	return MetaTypeVariant
}

func (v Variant) ConvertToAssortmentPosition() (*AssortmentPosition, error) {
	return convertToAssortmentPosition(v)
}

// VariantPack Упаковка модификации.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-modifikaciq-modifikacii-atributy-wlozhennyh-suschnostej-upakowki-modifikacii
type VariantPack struct {
	Barcodes   *Barcodes  `json:"barcodes,omitempty"`   // Массив штрихкодов упаковки модификации. Данный массив может содержать только один штрихкод
	ID         *uuid.UUID `json:"id,omitempty"`         // ID упаковки модификации
	ParentPack *Pack      `json:"parentpack,omitempty"` // Метаданные родительской упаковки (упаковки товара), для которой переопределяется штрихкод
}

func (v VariantPack) String() string {
	return Stringify(v)
}

// VariantService
// Сервис для работы с модификациями.
type VariantService interface {
	GetList(ctx context.Context, params *Params) (*List[Variant], *resty.Response, error)
	Create(ctx context.Context, variant *Variant, params *Params) (*Variant, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, variantList []*Variant, params *Params) (*[]Variant, *resty.Response, error)
	DeleteMany(ctx context.Context, variantList []*Variant) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Variant, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, variant *Variant, params *Params) (*Variant, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetadataVariant, *resty.Response, error)
	GetImages(ctx context.Context, id *uuid.UUID) (*MetaArray[Image], *resty.Response, error)
	CreateImage(ctx context.Context, id *uuid.UUID, image *Image) (*[]*Image, *resty.Response, error)
	UpdateImages(ctx context.Context, id *uuid.UUID, images []*Image) (*[]Image, *resty.Response, error)
	DeleteImage(ctx context.Context, id *uuid.UUID, imageId *uuid.UUID) (bool, *resty.Response, error)
	DeleteImages(ctx context.Context, id *uuid.UUID, images []*Image) (*DeleteManyResponse, *resty.Response, error)
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
