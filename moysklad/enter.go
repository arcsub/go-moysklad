package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Enter Оприходование.
// Ключевое слово: enter
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-oprihodowanie
type Enter struct {
	Organization *Organization             `json:"organization,omitempty"`
	Sum          *float64                  `json:"sum,omitempty"`
	Moment       *Timestamp                `json:"moment,omitempty"`
	Code         *string                   `json:"code,omitempty"`
	Created      *Timestamp                `json:"created,omitempty"`
	Deleted      *Timestamp                `json:"deleted,omitempty"`
	Description  *string                   `json:"description,omitempty"`
	ExternalCode *string                   `json:"externalCode,omitempty"`
	Files        *Files                    `json:"files,omitempty"`
	Group        *Group                    `json:"group,omitempty"`
	ID           *uuid.UUID                `json:"id,omitempty"`
	Meta         *Meta                     `json:"meta,omitempty"`
	Updated      *Timestamp                `json:"updated,omitempty"`
	Applicable   *bool                     `json:"applicable,omitempty"`
	Printed      *bool                     `json:"printed,omitempty"`
	Overhead     *Overhead                 `json:"overhead,omitempty"`
	Owner        *Employee                 `json:"owner,omitempty"`
	Positions    *Positions[EnterPosition] `json:"positions,omitempty"`
	AccountID    *uuid.UUID                `json:"accountId,omitempty"`
	Project      *Project                  `json:"project,omitempty"`
	Published    *bool                     `json:"published,omitempty"`
	Rate         *Rate                     `json:"rate,omitempty"`
	Shared       *bool                     `json:"shared,omitempty"`
	State        *State                    `json:"state,omitempty"`
	Store        *Store                    `json:"store,omitempty"`
	Name         *string                   `json:"name,omitempty"`
	SyncID       *uuid.UUID                `json:"syncId,omitempty"`
	Attributes   Attributes                `json:"attributes,omitempty"`
}

func (e Enter) String() string {
	return Stringify(e)
}

func (e Enter) MetaType() MetaType {
	return MetaTypeEnter
}

// EnterPosition Позиция оприходования
// Ключевое слово: enterposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-oprihodowanie-oprihodowaniq-pozicii-oprihodowaniq
type EnterPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	Country    *Country            `json:"country,omitempty"`    // Метаданные страны
	GTD        *GTD                `json:"gtd,omitempty"`        // ГТД
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Overhead   *float64            `json:"overhead,omitempty"`   // Накладные расходы. Если Позиции Оприходования не заданы, то накладные расходы нельзя задать
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Price      *float64            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Reason     *string             `json:"reason,omitempty"`     // Причина оприходования данной позиции
	Slot       *Slot               `json:"slot,omitempty"`       // Ячейка на складе
	Things     *Things             `json:"things,omitempty"`     // Серийные номера. Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете. В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута.
}

func (e EnterPosition) String() string {
	return Stringify(e)
}

func (e EnterPosition) MetaType() MetaType {
	return MetaTypeEnterPosition
}

// EnterTemplateArg
// Документ: Оприходование (enter)
// Основание, на котором он может быть создан:
// - Инвентаризация(inventory)
type EnterTemplateArg struct {
	Inventory *MetaWrapper `json:"inventory,omitempty"`
}

// EnterService
// Сервис для работы с оприходованиями.
type EnterService interface {
	GetList(ctx context.Context, params *Params) (*List[Enter], *resty.Response, error)
	Create(ctx context.Context, enter *Enter, params *Params) (*Enter, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, enterList []*Enter, params *Params) (*[]Enter, *resty.Response, error)
	DeleteMany(ctx context.Context, enterList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Enter, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, enter *Enter, params *Params) (*Enter, *resty.Response, error)
	//endpointTemplate[Enter]
	//endpointTemplateBasedOn[Enter, EnterTemplateArg]
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	GetPositions(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[EnterPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, params *Params) (*EnterPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, position *EnterPosition, params *Params) (*EnterPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id *uuid.UUID, position *EnterPosition) (*EnterPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id *uuid.UUID, positions []*EnterPosition) (*[]EnterPosition, *resty.Response, error)
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
	GetPublications(ctx context.Context, id *uuid.UUID) (*MetaArray[Publication], *resty.Response, error)
	GetPublicationByID(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (*Publication, *resty.Response, error)
	Publish(ctx context.Context, id *uuid.UUID, template *Templater) (*Publication, *resty.Response, error)
	DeletePublication(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (bool, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*Enter, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
	MoveToTrash(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewEnterService(client *Client) EnterService {
	e := NewEndpoint(client, "entity/enter")
	return newMainService[Enter, EnterPosition, MetadataAttributeSharedStates, any](e)
}
