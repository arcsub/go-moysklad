package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// SalesReturn Возврат покупателя.
// Ключевое слово: salesreturn
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vozwrat-pokupatelq
type SalesReturn struct {
	AccountID           *uuid.UUID                      `json:"accountId,omitempty"`           // ID учетной записи
	Agent               *Counterparty                   `json:"agent,omitempty"`               // Метаданные контрагента
	AgentAccount        *AgentAccount                   `json:"agentAccount,omitempty"`        // Метаданные счета контрагента
	Applicable          *bool                           `json:"applicable,omitempty"`          // Отметка о проведении
	Attributes          *Attributes                     `json:"attributes,omitempty"`          // Коллекция метаданных доп. полей. Поля объекта
	Code                *string                         `json:"code,omitempty"`                // Код Возврата Покупателя
	Contract            *Contract                       `json:"contract,omitempty"`            // Метаданные договора
	Created             *Timestamp                      `json:"created,omitempty"`             // Дата создания
	Deleted             *Timestamp                      `json:"deleted,omitempty"`             // Момент последнего удаления Возврата Покупателя
	Description         *string                         `json:"description,omitempty"`         // Комментарий Возврата Покупателя
	ExternalCode        *string                         `json:"externalCode,omitempty"`        // Внешний код Возврата Покупателя
	Files               *Files                          `json:"files,omitempty"`               // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group               *Group                          `json:"group,omitempty"`               // Отдел сотрудника
	ID                  *uuid.UUID                      `json:"id,omitempty"`                  // ID сущности
	Meta                *Meta                           `json:"meta,omitempty"`                // Метаданные
	Moment              *Timestamp                      `json:"moment,omitempty"`              // Дата документа
	Name                *string                         `json:"name,omitempty"`                // Наименование
	Organization        *Organization                   `json:"organization,omitempty"`        // Метаданные юрлица
	OrganizationAccount *AgentAccount                   `json:"organizationAccount,omitempty"` // Метаданные счета юрлица
	Owner               *Employee                       `json:"owner,omitempty"`               // Владелец (Сотрудник)
	Positions           *Positions[SalesReturnPosition] `json:"positions,omitempty"`           // Метаданные позиций Возврата Покупателя
	Printed             *bool                           `json:"printed,omitempty"`             // Напечатан ли документ
	Project             *Project                        `json:"project,omitempty"`             // Метаданные проекта
	Published           *bool                           `json:"published,omitempty"`           // Опубликован ли документ
	Rate                *Rate                           `json:"rate,omitempty"`                // Валюта
	SalesChannel        *SalesChannel                   `json:"salesChannel,omitempty"`        // Метаданные канала продаж
	Shared              *bool                           `json:"shared,omitempty"`              // Общий доступ
	State               *State                          `json:"state,omitempty"`               // Метаданные статуса Возврата Покупателя
	Store               *Store                          `json:"store,omitempty"`               // Метаданные склада
	Sum                 *Decimal                        `json:"sum,omitempty"`                 // Сумма
	SyncID              *uuid.UUID                      `json:"syncId,omitempty"`              // ID синхронизации. После заполнения недоступен для изменения
	Updated             *Timestamp                      `json:"updated,omitempty"`             // Момент последнего обновления
	VatEnabled          *bool                           `json:"vatEnabled,omitempty"`          // Учитывается ли НДС
	VatIncluded         *bool                           `json:"vatIncluded,omitempty"`         // Включен ли НДС в цену
	VatSum              *Decimal                        `json:"vatSum,omitempty"`              // Сумма включая НДС
	Demand              *Demand                         `json:"demand,omitempty"`              // Ссылка на отгрузку, по которой произошел возврат в формате Метаданных
	Losses              *Losses                         `json:"losses,omitempty"`              // Массив ссылок на связанные списания в формате Метаданных
	Payments            *Payments                       `json:"payments,omitempty"`            // Массив ссылок на связанные операции в формате Метаданных
	PayedSum            *Decimal                        `json:"payedSum,omitempty"`            // Сумма исходящих платежей по возврату покупателя
	FactureOut          *FactureOut                     `json:"factureOut,omitempty"`          // Ссылка на Счет-фактуру выданный, с которым связан этот возврат, в формате Метаданных
}

func (s SalesReturn) String() string {
	return Stringify(s)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (s SalesReturn) GetMeta() *Meta {
	return s.Meta
}

func (s SalesReturn) MetaType() MetaType {
	return MetaTypeSalesReturn
}

type SalesReturns = Slice[SalesReturn]

// SalesReturnPosition Позиция Возврата покупателя.
// Ключевое слово: salesreturnposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vozwrat-pokupatelq-vozwraty-pokupatelej-pozicii-vozwrata-pokupatelq
type SalesReturnPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	Cost       *Decimal            `json:"cost,omitempty"`       // Себестоимость (выводится, если документ был создан без основания)
	Country    *Country            `json:"country,omitempty"`    // Метаданные Страны
	Discount   *Decimal            `json:"discount,omitempty"`   // Процент скидки или наценки. Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%
	GTD        *GTD                `json:"gtd,omitempty"`        // ГТД
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Price      *Decimal            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Slot       *Slot               `json:"slot,omitempty"`       // Метаданные
	Things     *Things             `json:"things,omitempty"`     // Серийные номера. Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете. В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута.
	Vat        *int                `json:"vat,omitempty"`        // НДС, которым облагается текущая позиция
	VatEnabled *bool               `json:"vatEnabled,omitempty"` // Включен ли НДС для позиции. С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
	Stock      *Stock              `json:"stock,omitempty"`      // Остатки и себестоимость `?fields=stock&expand=positions`
}

func (s SalesReturnPosition) String() string {
	return Stringify(s)
}

func (s SalesReturnPosition) MetaType() MetaType {
	return MetaTypeSalesReturnPosition
}

// SalesReturnTemplateArg
// Документ: Возврат покупателя (salesreturn)
// Основание, на котором он может быть создан:
// - Отгрузка (demand)
// - Розничная продажа (retaildemand)
type SalesReturnTemplateArg struct {
	Demand       *MetaWrapper `json:"demand,omitempty"`
	RetailDemand *MetaWrapper `json:"retailDemand,omitempty"`
}

// SalesReturnService
// Сервис для работы с возвратами покупателей.
type SalesReturnService interface {
	GetList(ctx context.Context, params *Params) (*List[SalesReturn], *resty.Response, error)
	Create(ctx context.Context, salesReturn *SalesReturn, params *Params) (*SalesReturn, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, salesReturnList []*SalesReturn, params *Params) (*[]SalesReturn, *resty.Response, error)
	DeleteMany(ctx context.Context, salesReturnList []*SalesReturn) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*SalesReturn, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, salesReturn *SalesReturn, params *Params) (*SalesReturn, *resty.Response, error)
	//endpointTemplate[SalesReturn]
	//endpointTemplateBasedOn[SalesReturn, SalesReturnTemplateArg]
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	GetPositions(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[SalesReturnPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, params *Params) (*SalesReturnPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, position *SalesReturnPosition, params *Params) (*SalesReturnPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id *uuid.UUID, position *SalesReturnPosition) (*SalesReturnPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id *uuid.UUID, positions []*SalesReturnPosition) (*[]SalesReturnPosition, *resty.Response, error)
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
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*SalesReturn, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
	Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewSalesReturnService(client *Client) SalesReturnService {
	e := NewEndpoint(client, "entity/salesreturn")
	return newMainService[SalesReturn, SalesReturnPosition, MetadataAttributeSharedStates, any](e)
}
