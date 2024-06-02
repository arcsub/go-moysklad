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
	Images              *Images                     `json:"images,omitempty"`
	Components          *Positions[BundleComponent] `json:"components,omitempty"`
	Country             *Country                    `json:"country,omitempty"`
	DiscountProhibited  *bool                       `json:"discountProhibited"`
	EffectiveVat        *int                        `json:"effectiveVat,omitempty"`
	EffectiveVatEnabled *bool                       `json:"effectiveVatEnabled,omitempty"`
	Files               *Files                      `json:"files,omitempty"`
	Group               *Group                      `json:"group,omitempty"`
	Vat                 *int                        `json:"vat,omitempty"`
	MinPrice            *MinPrice                   `json:"minPrice,omitempty"`
	Overhead            *BundleOverhead             `json:"overhead,omitempty"`
	Owner               *Employee                   `json:"owner,omitempty"`
	PartialDisposal     *bool                       `json:"partialDisposal,omitempty"`
	PathName            *string                     `json:"pathName,omitempty"`
	Weight              *float64                    `json:"weight,omitempty"`
	SalePrices          *SalePrices                 `json:"salePrices,omitempty"`
	ProductFolder       *ProductFolder              `json:"productFolder,omitempty"`
	Shared              *bool                       `json:"shared,omitempty"`
	Updated             *Timestamp                  `json:"updated,omitempty"`
	AccountID           *uuid.UUID                  `json:"accountId,omitempty"`
	Tnved               *string                     `json:"tnved,omitempty"`
	VatEnabled          *bool                       `json:"vatEnabled,omitempty"`
	Uom                 *Uom                        `json:"uom,omitempty"`
	Barcodes            *Barcodes                   `json:"barcodes,omitempty"`
	UseParentVat        *bool                       `json:"useParentVat,omitempty"`
	TaxSystem           GoodTaxSystem               `json:"taxSystem,omitempty"`
	TrackingType        TrackingType                `json:"trackingType,omitempty"`
	PaymentItemType     PaymentItem                 `json:"paymentItemType,omitempty"`
	Attributes          Attributes                  `json:"attributes,omitempty"`
}

func (b Bundle) String() string {
	return Stringify(b)
}

func (b Bundle) MetaType() MetaType {
	return MetaTypeBundle
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (b Bundle) GetMeta() Meta {
	return Deref(b.Meta)
}

func (b Bundle) ConvertToAssortmentPosition() (*AssortmentPosition, error) {
	return convertToAssortmentPosition(b)
}

// BundleOverhead Дополнительные расходы
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-komplekt-komplekty-atributy-wlozhennyh-suschnostej-dopolnitel-nye-rashody
type BundleOverhead struct {
	Value    *Decimal  `json:"value,omitempty"`    // Значение цены
	Currency *Currency `json:"currency,omitempty"` // Ссылка на валюту в формате Метаданных
}

func (b BundleOverhead) String() string {
	return Stringify(b)
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

func (b BundleComponent) String() string {
	return Stringify(b)
}

func (b BundleComponent) MetaType() MetaType {
	return MetaTypeBundleComponent
}

// BundleService
// Сервис для работы с комплектами.
type BundleService interface {
	GetList(ctx context.Context, params *Params) (*List[Bundle], *resty.Response, error)
	Create(ctx context.Context, bundle *Bundle, params *Params) (*Bundle, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, bundleList []*Bundle, params *Params) (*[]Bundle, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Bundle, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, bundle *Bundle, params *Params) (*Bundle, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	DeleteMany(ctx context.Context, bundleList []*Bundle) (*DeleteManyResponse, *resty.Response, error)
	GetComponents(ctx context.Context, id *uuid.UUID) (*List[BundleComponent], *resty.Response, error)
	CreateComponent(ctx context.Context, id *uuid.UUID, bundleComponent *BundleComponent) (*BundleComponent, *resty.Response, error)
	GetComponentById(ctx context.Context, id, componentID *uuid.UUID) (*BundleComponent, *resty.Response, error)
	UpdateComponent(ctx context.Context, id, componentID *uuid.UUID, bundleComponent *BundleComponent) (*BundleComponent, *resty.Response, error)
	DeleteComponent(ctx context.Context, id, componentID *uuid.UUID) (bool, *resty.Response, error)
}

type bundleService struct {
	Endpoint
	endpointGetList[Bundle]
	endpointCreate[Bundle]
	endpointCreateUpdateMany[Bundle]
	endpointGetById[Bundle]
	endpointUpdate[Bundle]
	endpointDelete
	endpointDeleteMany[Bundle]
}

func NewBundleService(client *Client) BundleService {
	e := NewEndpoint(client, "entity/bundle")
	return &bundleService{
		Endpoint:                 e,
		endpointGetList:          endpointGetList[Bundle]{e},
		endpointCreate:           endpointCreate[Bundle]{e},
		endpointCreateUpdateMany: endpointCreateUpdateMany[Bundle]{e},
		endpointGetById:          endpointGetById[Bundle]{e},
		endpointUpdate:           endpointUpdate[Bundle]{e},
		endpointDelete:           endpointDelete{e},
		endpointDeleteMany:       endpointDeleteMany[Bundle]{e},
	}
}

// GetComponents Получить компоненты Комплекта.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-komplekt-poluchit-komponenty-komplekta
func (s *bundleService) GetComponents(ctx context.Context, id *uuid.UUID) (*List[BundleComponent], *resty.Response, error) {
	path := fmt.Sprintf("entity/bundle/%s/components", id)
	return NewRequestBuilder[List[BundleComponent]](s.client, path).Get(ctx)
}

// CreateComponent Добавить компонент Комплекта.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-komplekt-dobawit-komponent-komplekta
func (s *bundleService) CreateComponent(ctx context.Context, id *uuid.UUID, bundleComponent *BundleComponent) (*BundleComponent, *resty.Response, error) {
	path := fmt.Sprintf("entity/bundle/%s/components", id)
	return NewRequestBuilder[BundleComponent](s.client, path).Post(ctx, bundleComponent)
}

// GetComponentById Получить компонент.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-komplekt-poluchit-komponent
func (s *bundleService) GetComponentById(ctx context.Context, id, componentId *uuid.UUID) (*BundleComponent, *resty.Response, error) {
	path := fmt.Sprintf("entity/bundle/%s/components/%s", id, componentId)
	return NewRequestBuilder[BundleComponent](s.client, path).Get(ctx)
}

// UpdateComponent Изменить компонент.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-komplekt-izmenit-komponent
func (s *bundleService) UpdateComponent(ctx context.Context, id, componentId *uuid.UUID, bundleComponent *BundleComponent) (*BundleComponent, *resty.Response, error) {
	path := fmt.Sprintf("entity/bundle/%s/components/%s", id, componentId)
	return NewRequestBuilder[BundleComponent](s.client, path).Put(ctx, bundleComponent)
}

// DeleteComponent Удалить компонент.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-komplekt-udalit-komponent
func (s *bundleService) DeleteComponent(ctx context.Context, id, componentId *uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("entity/bundle/%s/components/%s", id, componentId)
	return NewRequestBuilder[any](s.client, path).Delete(ctx)
}
