package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// ProcessingOrder Заказ на производство.
// Ключевое слово: processingorder
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-zakaz-na-proizwodstwo
type ProcessingOrder struct {
	AccountID             *uuid.UUID                          `json:"accountId,omitempty"`             // ID учетной записи
	Applicable            *bool                               `json:"applicable,omitempty"`            // Отметка о проведении
	Attributes            *Attributes                         `json:"attributes,omitempty"`            // Коллекция метаданных доп. полей. Поля объекта
	Code                  *string                             `json:"code,omitempty"`                  // Код
	Created               *Timestamp                          `json:"created,omitempty"`               // Дата создания
	Deleted               *Timestamp                          `json:"deleted,omitempty"`               // Момент последнего удаления
	DeliveryPlannedMoment *Timestamp                          `json:"deliveryPlannedMoment,omitempty"` // Планируемая дата производства
	Description           *string                             `json:"description,omitempty"`           // Комментарий
	ExternalCode          *string                             `json:"externalCode,omitempty"`          // Внешний код
	Files                 *Files                              `json:"files,omitempty"`                 // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group                 *Group                              `json:"group,omitempty"`                 // Отдел сотрудника
	ID                    *uuid.UUID                          `json:"id,omitempty"`                    // ID сущности
	Meta                  *Meta                               `json:"meta,omitempty"`                  // Метаданные
	Moment                *Timestamp                          `json:"moment,omitempty"`                // Дата документа
	Name                  *string                             `json:"name,omitempty"`                  // Наименование
	Organization          *Organization                       `json:"organization,omitempty"`          // Метаданные юрлица
	OrganizationAccount   *AgentAccount                       `json:"organizationAccount,omitempty"`   // Метаданные счета юрлица
	Owner                 *Employee                           `json:"owner,omitempty"`                 // Владелец (Сотрудник)
	Positions             *Positions[ProcessingOrderPosition] `json:"positions,omitempty"`             // Метаданные позиций Заказа на производство
	Printed               *bool                               `json:"printed,omitempty"`               // Напечатан ли документ
	ProcessingPlan        *ProcessingPlan                     `json:"processingPlan,omitempty"`        // Метаданные Тех. плана
	Project               *Project                            `json:"project,omitempty"`               // Метаданные проекта
	Published             *bool                               `json:"published,omitempty"`             // Опубликован ли документ
	Quantity              *float64                            `json:"quantity,omitempty"`              // Объем производства
	Shared                *bool                               `json:"shared,omitempty"`                // Общий доступ
	State                 *State                              `json:"state,omitempty"`                 // Метаданные статуса
	Store                 *Store                              `json:"store,omitempty"`                 // Метаданные склада
	SyncID                *uuid.UUID                          `json:"syncId,omitempty"`                // ID синхронизации. После заполнения недоступен для изменения
	Updated               *Timestamp                          `json:"updated,omitempty"`               // Момент последнего обновления
	Processings           *Processings                        `json:"processings,omitempty"`           // Массив ссылок на связанные тех. операции в формате Метаданных
}

func (p ProcessingOrder) String() string {
	return Stringify(p)
}

func (p ProcessingOrder) MetaType() MetaType {
	return MetaTypeProcessingOrder
}

// ProcessingOrderPosition Позиция Заказа на производство.
// Ключевое слово: processingorderposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-zakaz-na-proizwodstwo-zakazy-na-proizwodstwo-pozicii-zakaza-na-proizwodstwo
type ProcessingOrderPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Reserve    *Decimal            `json:"reserve,omitempty"`    // Резерв данной позиции
}

func (p ProcessingOrderPosition) String() string {
	return Stringify(p)
}

func (p ProcessingOrderPosition) MetaType() MetaType {
	return MetaTypeProcessingOrderPosition
}

// ProcessingOrderTemplateArg
// Документ: Заказ на производство (processingorder)
// Основание, на котором он может быть создан:
// - Техкарта (processingplan)
type ProcessingOrderTemplateArg struct {
	ProcessingPlan *MetaWrapper `json:"processingPlan,omitempty"`
}

// ProcessingOrderService
// Сервис для работы с заказами на производство.
type ProcessingOrderService interface {
	GetList(ctx context.Context, params *Params) (*List[ProcessingOrder], *resty.Response, error)
	Create(ctx context.Context, processingOrder *ProcessingOrder, params *Params) (*ProcessingOrder, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, processingOrderList []*ProcessingOrder, params *Params) (*[]ProcessingOrder, *resty.Response, error)
	DeleteMany(ctx context.Context, processingOrderList []*ProcessingOrder) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*ProcessingOrder, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, processingOrder *ProcessingOrder, params *Params) (*ProcessingOrder, *resty.Response, error)
	//endpointTemplate[ProcessingOrder]
	//endpointTemplateBasedOn[ProcessingOrder, ProcessingOrderTemplateArg]
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	GetPositions(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[ProcessingOrderPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, params *Params) (*ProcessingOrderPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, position *ProcessingOrderPosition, params *Params) (*ProcessingOrderPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id *uuid.UUID, position *ProcessingOrderPosition) (*ProcessingOrderPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id *uuid.UUID, positions []*ProcessingOrderPosition) (*[]ProcessingOrderPosition, *resty.Response, error)
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
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*ProcessingOrder, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewProcessingOrderService(client *Client) ProcessingOrderService {
	e := NewEndpoint(client, "entity/processingorder")
	return newMainService[ProcessingOrder, ProcessingOrderPosition, MetadataAttributeSharedStates, any](e)
}
