package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// InternalOrder Внутренний заказ.
// Ключевое слово: internalorder
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vnutrennij-zakaz
type InternalOrder struct {
	AccountID             *uuid.UUID                        `json:"accountId,omitempty"`             // ID учетной записи
	Applicable            *bool                             `json:"applicable,omitempty"`            // Отметка о проведении
	Attributes            *Attributes                       `json:"attributes,omitempty"`            // Коллекция метаданных доп. полей. Поля объекта
	Code                  *string                           `json:"code,omitempty"`                  // Код
	Created               *Timestamp                        `json:"created,omitempty"`               // Дата создания
	Deleted               *Timestamp                        `json:"deleted,omitempty"`               // Момент последнего удаления
	DeliveryPlannedMoment *Timestamp                        `json:"deliveryPlannedMoment,omitempty"` // Планируемая дата приемки
	Description           *string                           `json:"description,omitempty"`           // Комментарий
	ExternalCode          *string                           `json:"externalCode,omitempty"`          // Внешний код
	Files                 *Files                            `json:"files,omitempty"`                 // Метаданные массива Файлов
	Group                 *Group                            `json:"group,omitempty"`                 // Отдел сотрудника
	ID                    *uuid.UUID                        `json:"id,omitempty"`                    // ID сущности
	Meta                  *Meta                             `json:"meta,omitempty"`                  // Метаданные
	Moment                *Timestamp                        `json:"moment,omitempty"`                // Дата документа
	Moves                 *Moves                            `json:"moves,omitempty"`                 // Коллекция метаданных на связанные заказы перемещения
	Name                  *string                           `json:"name,omitempty"`                  // Наименование
	Organization          *Organization                     `json:"organization,omitempty"`          // Метаданные юрлица
	Owner                 *Employee                         `json:"owner,omitempty"`                 // Владелец (Сотрудник)
	Positions             *Positions[InternalOrderPosition] `json:"positions,omitempty"`             // Метаданные позиций Внутреннего заказа
	Printed               *bool                             `json:"printed,omitempty"`               // Напечатан ли документ
	Project               *Project                          `json:"project,omitempty"`               // Метаданные проекта
	Published             *bool                             `json:"published,omitempty"`             // Опубликован ли документ
	PurchaseOrders        *PurchaseOrders                   `json:"purchaseOrders,omitempty"`        // Коллекция метаданных на связанные заказы поставщику
	Rate                  *Rate                             `json:"rate,omitempty"`                  // Валюта
	Shared                *bool                             `json:"shared,omitempty"`                // Общий доступ
	State                 *State                            `json:"state,omitempty"`                 // Метаданные статуса
	Store                 *Store                            `json:"store,omitempty"`                 // Метаданные склада
	Sum                   *Decimal                          `json:"sum,omitempty"`                   // Сумма
	SyncID                *uuid.UUID                        `json:"syncId,omitempty"`                // ID синхронизации. После заполнения недоступен для изменения
	Updated               *Timestamp                        `json:"updated,omitempty"`               // Момент последнего обновления
	VatEnabled            *bool                             `json:"vatEnabled,omitempty"`            // Учитывается ли НДС
	VatIncluded           *bool                             `json:"vatIncluded,omitempty"`           // Включен ли НДС в цену
	VatSum                *Decimal                          `json:"vatSum,omitempty"`                // Сумма включая НДС
}

func (i InternalOrder) String() string {
	return Stringify(i)
}

func (i InternalOrder) MetaType() MetaType {
	return MetaTypeInternalOrder
}

// InternalOrderPosition Позиция Внутреннего заказа.
// Ключевое слово: internalorderposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vnutrennij-zakaz-vnutrennie-zakazy-pozicii-vnutrennego-zakaza
type InternalOrderPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Price      *Decimal            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Vat        *int                `json:"vat,omitempty"`        // НДС, которым облагается текущая позиция
	VatEnabled *bool               `json:"vatEnabled,omitempty"` // Включен ли НДС для позиции. С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
}

func (i InternalOrderPosition) String() string {
	return Stringify(i)
}

func (i InternalOrderPosition) MetaType() MetaType {
	return MetaTypeInternalOrderPosition
}

// InternalOrderService
// Сервис для работы с внутренними заказами.
type InternalOrderService interface {
	GetList(ctx context.Context, params *Params) (*List[InternalOrder], *resty.Response, error)
	Create(ctx context.Context, entity *InternalOrder, params *Params) (*InternalOrder, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, entities []*InternalOrder, params *Params) (*[]InternalOrder, *resty.Response, error)
	DeleteMany(ctx context.Context, entities []*InternalOrder) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*InternalOrder, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, entity *InternalOrder, params *Params) (*InternalOrder, *resty.Response, error)
	//endpointTemplate[InternalOrder]
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	GetPositions(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[InternalOrderPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, params *Params) (*InternalOrderPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, position *InternalOrderPosition, params *Params) (*InternalOrderPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id *uuid.UUID, position *InternalOrderPosition) (*InternalOrderPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id *uuid.UUID, positions []*InternalOrderPosition) (*[]InternalOrderPosition, *resty.Response, error)
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
	DeleteAttributes(ctx context.Context, attributeList []*Attribute) (*DeleteManyResponse, *resty.Response, error)
	GetPublications(ctx context.Context, id *uuid.UUID) (*MetaArray[Publication], *resty.Response, error)
	GetPublicationByID(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (*Publication, *resty.Response, error)
	Publish(ctx context.Context, id *uuid.UUID, template *Templater) (*Publication, *resty.Response, error)
	DeletePublication(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (bool, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*InternalOrder, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewInternalOrderService(client *Client) InternalOrderService {
	e := NewEndpoint(client, "entity/internalorder")
	return newMainService[InternalOrder, InternalOrderPosition, MetadataAttributeSharedStates, any](e)
}
