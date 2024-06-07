package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Loss Списание.
// Ключевое слово: loss
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-spisanie
type Loss struct {
	Organization *Organization            `json:"organization,omitempty"`
	SyncID       *uuid.UUID               `json:"syncId,omitempty"`
	Moment       *Timestamp               `json:"moment,omitempty"`
	Code         *string                  `json:"code,omitempty"`
	Created      *Timestamp               `json:"created,omitempty"`
	Deleted      *Timestamp               `json:"deleted,omitempty"`
	Description  *string                  `json:"description,omitempty"`
	ExternalCode *string                  `json:"externalCode,omitempty"`
	Files        *Files                   `json:"files,omitempty"`
	Group        *Group                   `json:"group,omitempty"`
	ID           *uuid.UUID               `json:"id,omitempty"`
	Meta         *Meta                    `json:"meta,omitempty"`
	SalesReturn  *SalesReturn             `json:"salesReturn,omitempty"`
	Applicable   *bool                    `json:"applicable,omitempty"`
	Project      *Project                 `json:"project,omitempty"`
	Owner        *Employee                `json:"owner,omitempty"`
	Positions    *Positions[LossPosition] `json:"positions,omitempty"`
	Printed      *bool                    `json:"printed,omitempty"`
	AccountID    *uuid.UUID               `json:"accountId,omitempty"`
	Published    *bool                    `json:"published,omitempty"`
	Rate         *Rate                    `json:"rate,omitempty"`
	Shared       *bool                    `json:"shared,omitempty"`
	State        *State                   `json:"state,omitempty"`
	Store        *Store                   `json:"store,omitempty"`
	Sum          *float64                 `json:"sum,omitempty"`
	Name         *string                  `json:"name,omitempty"`
	Updated      *Timestamp               `json:"updated,omitempty"`
	Attributes   Attributes               `json:"attributes,omitempty"`
}

func (l Loss) String() string {
	return Stringify(l)
}

func (l Loss) MetaType() MetaType {
	return MetaTypeLoss
}

type Losses = Slice[Loss]

// LossPosition Позиция Списания.
// Ключевое слово: lossposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-spisanie-spisaniq-pozicii-spisaniq
type LossPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID сущности
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Price      *float64            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Reason     *string             `json:"reason,omitempty"`     // Причина списания данной позиции
	Slot       *Slot               `json:"slot,omitempty"`       // Ячейка на складе
	Things     *Things             `json:"things,omitempty"`     // Серийные номера. Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете. В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута.
}

func (l LossPosition) String() string {
	return Stringify(l)
}

func (l LossPosition) MetaType() MetaType {
	return MetaTypeLossPosition
}

// LossTemplateArg
// Документ: Списание (loss)
// Основание, на котором он может быть создан:
// - Возврат покупателя (salesreturn)
// - инвентаризация(inventory)
type LossTemplateArg struct {
	SalesReturn *MetaWrapper `json:"salesReturn,omitempty"`
	Inventory   *MetaWrapper `json:"inventory,omitempty"`
}

// LossService
// Сервис для работы со списаниями.
type LossService interface {
	GetList(ctx context.Context, params *Params) (*List[Loss], *resty.Response, error)
	Create(ctx context.Context, loss *Loss, params *Params) (*Loss, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, lossList []*Loss, params *Params) (*[]Loss, *resty.Response, error)
	DeleteMany(ctx context.Context, lossList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Loss, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, loss *Loss, params *Params) (*Loss, *resty.Response, error)
	//endpointTemplate[Loss]
	//endpointTemplateBasedOn[Loss, LossTemplateArg]
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	GetPositions(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[LossPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, params *Params) (*LossPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, position *LossPosition, params *Params) (*LossPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id *uuid.UUID, position *LossPosition) (*LossPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id *uuid.UUID, positions []*LossPosition) (*[]LossPosition, *resty.Response, error)
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
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*Loss, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	MoveToTrash(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewLossService(client *Client) LossService {
	e := NewEndpoint(client, "entity/loss")
	return newMainService[Loss, LossPosition, MetadataAttributeSharedStates, any](e)
}
