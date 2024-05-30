package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Supply Приёмка.
// Ключевое слово: supply
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-priemka
type Supply struct {
	AccountID           *uuid.UUID                 `json:"accountId,omitempty"`           // ID учетной записи
	Agent               *Counterparty              `json:"agent,omitempty"`               // Метаданные контрагента
	AgentAccount        *AgentAccount              `json:"agentAccount,omitempty"`        // Метаданные счета контрагента
	Applicable          *bool                      `json:"applicable,omitempty"`          // Отметка о проведении
	Attributes          *Attributes                `json:"attributes,omitempty"`          // Коллекция метаданных доп. полей. Поля объекта
	Code                *string                    `json:"code,omitempty"`                // Код Приемки
	Contract            *Contract                  `json:"contract,omitempty"`            // Метаданные договора
	Created             *Timestamp                 `json:"created,omitempty"`             // Дата создания
	Deleted             *Timestamp                 `json:"deleted,omitempty"`             // Момент последнего удаления Приемки
	Description         *string                    `json:"description,omitempty"`         // Комментарий Приемки
	ExternalCode        *string                    `json:"externalCode,omitempty"`        // Внешний код Приемки
	Files               *Files                     `json:"files,omitempty"`               // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group               *Group                     `json:"group,omitempty"`               // Отдел сотрудника
	ID                  *uuid.UUID                 `json:"id,omitempty"`                  // ID сущности
	IncomingDate        *Timestamp                 `json:"incomingDate,omitempty"`        // Входящая дата
	IncomingNumber      *string                    `json:"incomingNumber,omitempty"`      // Входящий номер
	Meta                *Meta                      `json:"meta,omitempty"`                // Метаданные
	Moment              *Timestamp                 `json:"moment,omitempty"`              // Дата документа
	Name                *string                    `json:"name,omitempty"`                // Наименование
	Organization        *Organization              `json:"organization,omitempty"`        // Метаданные юрлица
	OrganizationAccount *AgentAccount              `json:"organizationAccount,omitempty"` // Метаданные счета юрлица
	Overhead            *Overhead                  `json:"overhead,omitempty"`            // Накладные расходы. Если Позиции Приемки не заданы, то накладные расходы нельзя задать
	Owner               *Employee                  `json:"owner,omitempty"`               // Владелец (Сотрудник)
	PayedSum            *Decimal                   `json:"payedSum,omitempty"`            // Сумма входящих платежей по Приемке
	Positions           *Positions[SupplyPosition] `json:"positions,omitempty"`           // Метаданные позиций Приемки
	Printed             *bool                      `json:"printed,omitempty"`             // Напечатан ли документ
	Project             *Project                   `json:"project,omitempty"`             // Метаданные проекта
	Published           *bool                      `json:"published,omitempty"`           // Опубликован ли документ
	Rate                *Rate                      `json:"rate,omitempty"`                // Валюта
	Shared              *bool                      `json:"shared,omitempty"`              // Общий доступ
	State               *State                     `json:"state,omitempty"`               // Метаданные статуса Приемки
	Store               *Store                     `json:"store,omitempty"`               // Метаданные склада
	Sum                 *Decimal                   `json:"sum,omitempty"`                 // Сумма
	SyncID              *uuid.UUID                 `json:"syncId,omitempty"`              // ID синхронизации. После заполнения недоступен для изменения
	Updated             *Timestamp                 `json:"updated,omitempty"`             // Момент последнего обновления
	VatEnabled          *bool                      `json:"vatEnabled,omitempty"`          // Учитывается ли НДС
	VatIncluded         *bool                      `json:"vatIncluded,omitempty"`         // Включен ли НДС в цену
	VatSum              *Decimal                   `json:"vatSum,omitempty"`              // Сумма включая НДС
	PurchaseOrder       *PurchaseOrder             `json:"purchaseOrder,omitempty"`       // Ссылка на связанный заказ поставщику в формате Метаданных
	FactureIn           *FactureIn                 `json:"factureIn,omitempty"`           // Ссылка на Счет-фактуру полученный, с которым связана эта Приемка в формате Метаданных
	InvoicesIn          *InvoicesIn                `json:"invoicesIn,omitempty"`          // Массив ссылок на связанные счета поставщиков в формате Метаданных
	Payments            *Payments                  `json:"payments,omitempty"`            // Массив ссылок на связанные платежи в формате Метаданных
	Returns             *PurchaseReturns           `json:"returns,omitempty"`             // Массив ссылок на связанные возвраты в формате Метаданных
}

func (s Supply) String() string {
	return Stringify(s)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (s Supply) GetMeta() *Meta {
	return s.Meta
}

func (s Supply) MetaType() MetaType {
	return MetaTypeSupply
}

type Supplies Slice[Supply]

// SupplyPosition Позиция Приемки.
// Ключевое слово: supplyposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-priemka-priemki-pozicii-priemki
type SupplyPosition struct {
	AccountID     *uuid.UUID          `json:"accountId,omitempty"`     // ID учетной записи
	Assortment    *AssortmentPosition `json:"assortment,omitempty"`    // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	Country       *Country            `json:"country,omitempty"`       // Метаданные страны
	Discount      *Decimal            `json:"discount,omitempty"`      // Процент скидки или наценки. Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%
	GTD           *GTD                `json:"gtd,omitempty"`           // ГТД
	ID            *uuid.UUID          `json:"id,omitempty"`            // ID позиции
	Pack          *Pack               `json:"pack,omitempty"`          // Упаковка Товара
	Price         *Decimal            `json:"price,omitempty"`         // Цена товара/услуги в копейках
	Quantity      *float64            `json:"quantity,omitempty"`      // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Slot          *Slot               `json:"slot,omitempty"`          // Ячейка на складе
	Things        *Things             `json:"things,omitempty"`        // Серийные номера. Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете. В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута.
	TrackingCodes *TrackingCodes      `json:"trackingCodes,omitempty"` // Коды маркировки товаров и транспортных упаковок
	Overhead      *Decimal            `json:"overhead,omitempty"`      // Накладные расходы. Если Позиции Приемки не заданы, то накладные расходы нельзя задать.
	Vat           *int                `json:"vat,omitempty"`           // НДС, которым облагается текущая позиция
	VatEnabled    *bool               `json:"vatEnabled,omitempty"`    // Включен ли НДС для позиции. С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
	Stock         *Stock              `json:"stock,omitempty"`         // Остатки и себестоимость `?fields=stock&expand=positions`
}

func (s SupplyPosition) String() string {
	return Stringify(s)
}

func (s SupplyPosition) MetaType() MetaType {
	return MetaTypeSupplyPosition
}

// SupplyService
// Сервис для работы с приёмками.
type SupplyService interface {
	GetList(ctx context.Context, params *Params) (*List[Supply], *resty.Response, error)
	Create(ctx context.Context, supply *Supply, params *Params) (*Supply, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, supplyList []*Supply, params *Params) (*[]Supply, *resty.Response, error)
	DeleteMany(ctx context.Context, supplyList []*Supply) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Supply, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, supply *Supply, params *Params) (*Supply, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	//endpointTemplate[Supply]
	GetPositions(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[SupplyPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, params *Params) (*SupplyPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, position *SupplyPosition, params *Params) (*SupplyPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id *uuid.UUID, position *SupplyPosition) (*SupplyPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id *uuid.UUID, positions []*SupplyPosition) (*[]SupplyPosition, *resty.Response, error)
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
	PrintDocument(ctx context.Context, id *uuid.UUID, PrintDocumentArg *PrintDocumentArg) (*PrintFile, *resty.Response, error)
	GetFiles(ctx context.Context, id *uuid.UUID) (*MetaArray[File], *resty.Response, error)
	CreateFile(ctx context.Context, id *uuid.UUID, file *File) (*[]File, *resty.Response, error)
	UpdateFiles(ctx context.Context, id *uuid.UUID, files []*File) (*[]File, *resty.Response, error)
	DeleteFile(ctx context.Context, id *uuid.UUID, fileId *uuid.UUID) (bool, *resty.Response, error)
	DeleteFiles(ctx context.Context, id *uuid.UUID, files []*File) (*DeleteManyResponse, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
	GetStateByID(ctx context.Context, id *uuid.UUID) (*State, *resty.Response, error)
	CreateState(ctx context.Context, state *State) (*State, *resty.Response, error)
	UpdateState(ctx context.Context, id *uuid.UUID, state *State) (*State, *resty.Response, error)
	CreateOrUpdateStates(ctx context.Context, id *uuid.UUID, states []*State) (*[]State, *resty.Response, error)
	DeleteState(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*Supply, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewSupplyService(client *Client) SupplyService {
	e := NewEndpoint(client, "entity/supply")
	return newMainService[Supply, SupplyPosition, MetadataAttributeSharedStates, any](e)
}
