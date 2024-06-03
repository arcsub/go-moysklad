package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Move Перемещение.
// Ключевое слово: move
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-peremeschenie
type Move struct {
	Moment        *Timestamp               `json:"moment,omitempty"`
	Updated       *Timestamp               `json:"updated,omitempty"`
	AccountID     *uuid.UUID               `json:"accountId,omitempty"`
	Code          *string                  `json:"code,omitempty"`
	Created       *Timestamp               `json:"created,omitempty"`
	Deleted       *Timestamp               `json:"deleted,omitempty"`
	Description   *string                  `json:"description,omitempty"`
	ExternalCode  *string                  `json:"externalCode,omitempty"`
	Files         *Files                   `json:"files,omitempty"`
	Group         *Group                   `json:"group,omitempty"`
	ID            *uuid.UUID               `json:"id,omitempty"`
	InternalOrder *InternalOrder           `json:"internalOrder,omitempty"`
	CustomerOrder *CustomerOrder           `json:"customerOrder,omitempty"`
	Meta          *Meta                    `json:"meta,omitempty"`
	Name          *string                  `json:"name,omitempty"`
	Organization  *Organization            `json:"organization,omitempty"`
	Applicable    *bool                    `json:"applicable,omitempty"`
	Overhead      *Overhead                `json:"overhead,omitempty"`
	Owner         *Employee                `json:"owner,omitempty"`
	Positions     *Positions[MovePosition] `json:"positions,omitempty"`
	Printed       *bool                    `json:"printed,omitempty"`
	Project       *Project                 `json:"project,omitempty"`
	Published     *bool                    `json:"published,omitempty"`
	Rate          *Rate                    `json:"rate,omitempty"`
	Shared        *bool                    `json:"shared,omitempty"`
	SourceStore   *Store                   `json:"sourceStore,omitempty"`
	State         *State                   `json:"state,omitempty"`
	Sum           *Decimal                 `json:"sum,omitempty"`
	SyncID        *uuid.UUID               `json:"syncId,omitempty"`
	TargetStore   *Store                   `json:"targetStore,omitempty"`
	Attributes    Attributes               `json:"attributes,omitempty"`
}

func (m Move) String() string {
	return Stringify(m)
}

func (m Move) MetaType() MetaType {
	return MetaTypeMove
}

type Moves Slice[Move]

// MovePosition Позиция перемещения.
// Ключевое слово: moveposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-peremeschenie-peremescheniq-pozicii-peremescheniq
type MovePosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Overhead   *Decimal            `json:"overhead,omitempty"`   // Накладные расходы. Если Позиции Перемещения не заданы, то накладные расходы нельзя задать
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Price      *Decimal            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе
	SourceSlot *Slot               `json:"sourceSlot,omitempty"` // Ячейка на складе, с которого совершается перемещение
	TargetSlot *Slot               `json:"targetSlot,omitempty"` // Ячейка на складе, на который совершается перемещение
	Things     *Things             `json:"things,omitempty"`     // Серийные номера. Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете. В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута
}

func (m MovePosition) String() string {
	return Stringify(m)
}

func (m MovePosition) MetaType() MetaType {
	return MetaTypeMovePosition
}

// MoveTemplateArg
// Документ: Перемещение (move)
// Основание, на котором он может быть создан:
// - Внутренний заказ (internalorder)
type MoveTemplateArg struct {
	InternalOrder *MetaWrapper `json:"internalOrder,omitempty"`
}

// MoveService
// Сервис для работы со перемещениями.
type MoveService interface {
	GetList(ctx context.Context, params *Params) (*List[Move], *resty.Response, error)
	Create(ctx context.Context, move *Move, params *Params) (*Move, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, moveList []*Move, params *Params) (*[]Move, *resty.Response, error)
	DeleteMany(ctx context.Context, moveList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Move, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, move *Move, params *Params) (*Move, *resty.Response, error)
	//endpointTemplate[Move]
	//endpointTemplateBasedOn[Move, MoveTemplateArg]
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	GetPositions(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[MovePosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, params *Params) (*MovePosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, position *MovePosition, params *Params) (*MovePosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id *uuid.UUID, position *MovePosition) (*MovePosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id *uuid.UUID, positions []*MovePosition) (*[]MovePosition, *resty.Response, error)
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
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*Move, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	MoveToTrash(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewMoveService(client *Client) MoveService {
	e := NewEndpoint(client, "entity/move")
	return newMainService[Move, MovePosition, MetadataAttributeSharedStates, any](e)
}
