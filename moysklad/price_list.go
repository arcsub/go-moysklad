package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// PriceList Прайс-лист.
// Ключевое слово: pricelist
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-prajs-list
type PriceList struct {
	Meta         *Meta                         `json:"meta,omitempty"`
	Columns      *PriceListColumns             `json:"columns,omitempty"`
	AccountID    *uuid.UUID                    `json:"accountId,omitempty"`
	Code         *string                       `json:"code,omitempty"`
	Moment       *Timestamp                    `json:"moment,omitempty"`
	Created      *Timestamp                    `json:"created,omitempty"`
	Deleted      *Timestamp                    `json:"deleted,omitempty"`
	Name         *string                       `json:"name,omitempty"`
	ExternalCode *string                       `json:"externalCode,omitempty"`
	Files        *Files                        `json:"files,omitempty"`
	Group        *Group                        `json:"group,omitempty"`
	ID           *uuid.UUID                    `json:"id,omitempty"`
	Updated      *Timestamp                    `json:"updated,omitempty"`
	Applicable   *bool                         `json:"applicable,omitempty"`
	Description  *string                       `json:"description,omitempty"`
	Organization *Organization                 `json:"organization,omitempty"`
	Owner        *Employee                     `json:"owner,omitempty"`
	Positions    *Positions[PriceListPosition] `json:"positions,omitempty"`
	PriceType    *PriceType                    `json:"priceType,omitempty"`
	Printed      *bool                         `json:"printed,omitempty"`
	Published    *bool                         `json:"published,omitempty"`
	Shared       *bool                         `json:"shared,omitempty"`
	State        *State                        `json:"state,omitempty"`
	SyncID       *uuid.UUID                    `json:"syncId,omitempty"`
	Attributes   Attributes                    `json:"attributes,omitempty"`
}

func (p PriceList) String() string {
	return Stringify(p)
}

func (p PriceList) MetaType() MetaType {
	return MetaTypePriceList
}

// PriceListCell Ячейка прайс листа.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-prajs-list-prajs-listy-yachejki
type PriceListCell struct {
	Column *string  `json:"column,omitempty"` // Название столбца, к которому относится данная ячейка
	Sum    *Decimal `json:"sum,omitempty"`    // Числовое значение ячейки
}

func (p PriceListCell) String() string {
	return Stringify(p)
}

type PriceListCells = Slice[PriceListCell]

// PriceListColumn Столбец прайс листа.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-prajs-list-prajs-listy-stolbcy
type PriceListColumn struct {
	Name               *string  `json:"name,omitempty"`               // Название столбца
	PercentageDiscount *float64 `json:"percentageDiscount,omitempty"` // Процентная наценка или скидка по умолчанию для столбца
}

func (p PriceListColumn) String() string {
	return Stringify(p)
}

type PriceListColumns = Slice[PriceListColumn]

// PriceListPosition Позиция прайс листа.
// Ключевое слово: pricelistrow
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-prajs-list-prajs-listy-pozicii-prajs-lista
type PriceListPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Товар/услуга/модификация, которую представляет собой позиция
	Cells      *PriceListCells     `json:"cells,omitempty"`      // Значения столбцов
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка товара
}

func (p PriceListPosition) String() string {
	return Stringify(p)
}

func (p PriceListPosition) MetaType() MetaType {
	return MetaTypePriceListPosition
}

// PriceListService
// Сервис для работы с прайс-листами.
type PriceListService interface {
	GetList(ctx context.Context, params *Params) (*List[PriceList], *resty.Response, error)
	Create(ctx context.Context, priceList *PriceList, params *Params) (*PriceList, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, priceListList []*PriceList, params *Params) (*[]PriceList, *resty.Response, error)
	DeleteMany(ctx context.Context, priceListList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*PriceList, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, priceList *PriceList, params *Params) (*PriceList, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	GetPositions(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[PriceListPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, params *Params) (*PriceListPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, position *PriceListPosition, params *Params) (*PriceListPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id *uuid.UUID, position *PriceListPosition) (*PriceListPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id *uuid.UUID, positions []*PriceListPosition) (*[]PriceListPosition, *resty.Response, error)
	DeletePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID) (bool, *resty.Response, error)
	GetPositionTrackingCodes(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID) (*MetaArray[TrackingCode], *resty.Response, error)
	CreateOrUpdatePositionTrackingCodes(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, trackingCodes TrackingCodes) (*[]TrackingCode, *resty.Response, error)
	DeletePositionTrackingCodes(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, trackingCodes TrackingCodes) (*DeleteManyResponse, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id *uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error)
	UpdateAttribute(ctx context.Context, id *uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*PriceList, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewPriceListService(client *Client) PriceListService {
	e := NewEndpoint(client, "entity/pricelist")
	return newMainService[PriceList, PriceListPosition, MetadataAttributeSharedStates, any](e)
}
