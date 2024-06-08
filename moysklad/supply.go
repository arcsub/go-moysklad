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
	OrganizationAccount *AgentAccount              `json:"organizationAccount,omitempty"`
	Applicable          *bool                      `json:"applicable,omitempty"`
	AgentAccount        *AgentAccount              `json:"agentAccount,omitempty"`
	Overhead            *Overhead                  `json:"overhead,omitempty"`
	Returns             *PurchaseReturns           `json:"returns,omitempty"`
	Code                *string                    `json:"code,omitempty"`
	Contract            *Contract                  `json:"contract,omitempty"`
	Created             *Timestamp                 `json:"created,omitempty"`
	Deleted             *Timestamp                 `json:"deleted,omitempty"`
	Description         *string                    `json:"description,omitempty"`
	ExternalCode        *string                    `json:"externalCode,omitempty"`
	Files               *Files                     `json:"files,omitempty"`
	Group               *Group                     `json:"group,omitempty"`
	ID                  *uuid.UUID                 `json:"id,omitempty"`
	IncomingDate        *Timestamp                 `json:"incomingDate,omitempty"`
	Owner               *Employee                  `json:"owner,omitempty"`
	Meta                *Meta                      `json:"meta,omitempty"`
	Moment              *Timestamp                 `json:"moment,omitempty"`
	Name                *string                    `json:"name,omitempty"`
	Organization        *Organization              `json:"organization,omitempty"`
	Payments            *Payments                  `json:"payments,omitempty"`
	Agent               *Counterparty              `json:"agent,omitempty"`
	IncomingNumber      *string                    `json:"incomingNumber,omitempty"`
	PayedSum            *float64                   `json:"payedSum,omitempty"`
	Positions           *Positions[SupplyPosition] `json:"positions,omitempty"`
	Printed             *bool                      `json:"printed,omitempty"`
	Project             *Project                   `json:"project,omitempty"`
	Published           *bool                      `json:"published,omitempty"`
	Rate                *Rate                      `json:"rate,omitempty"`
	Shared              *bool                      `json:"shared,omitempty"`
	State               *State                     `json:"state,omitempty"`
	Store               *Store                     `json:"store,omitempty"`
	Sum                 *float64                   `json:"sum,omitempty"`
	SyncID              *uuid.UUID                 `json:"syncId,omitempty"`
	Updated             *Timestamp                 `json:"updated,omitempty"`
	VatEnabled          *bool                      `json:"vatEnabled,omitempty"`
	VatIncluded         *bool                      `json:"vatIncluded,omitempty"`
	VatSum              *float64                   `json:"vatSum,omitempty"`
	PurchaseOrder       *PurchaseOrder             `json:"purchaseOrder,omitempty"`
	FactureIn           *FactureIn                 `json:"factureIn,omitempty"`
	InvoicesIn          Slice[InvoiceIn]           `json:"invoicesIn,omitempty"`
	AccountID           *uuid.UUID                 `json:"accountId,omitempty"`
	Attributes          Attributes                 `json:"attributes,omitempty"`
}

func (s Supply) String() string {
	return Stringify(s)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (s Supply) GetMeta() Meta {
	return Deref(s.Meta)
}

func (s Supply) MetaType() MetaType {
	return MetaTypeSupply
}

// SupplyPosition Позиция Приемки.
// Ключевое слово: supplyposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-priemka-priemki-pozicii-priemki
type SupplyPosition struct {
	AccountID     *uuid.UUID          `json:"accountId,omitempty"`     // ID учетной записи
	Assortment    *AssortmentPosition `json:"assortment,omitempty"`    // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	Country       *Country            `json:"country,omitempty"`       // Метаданные страны
	Discount      *float64            `json:"discount,omitempty"`      // Процент скидки или наценки. Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%
	GTD           *GTD                `json:"gtd,omitempty"`           // ГТД
	ID            *uuid.UUID          `json:"id,omitempty"`            // ID позиции
	Pack          *Pack               `json:"pack,omitempty"`          // Упаковка Товара
	Price         *float64            `json:"price,omitempty"`         // Цена товара/услуги в копейках
	Quantity      *float64            `json:"quantity,omitempty"`      // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Slot          *Slot               `json:"slot,omitempty"`          // Ячейка на складе
	Things        *Things             `json:"things,omitempty"`        // Серийные номера. Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете. В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута.
	TrackingCodes *TrackingCodes      `json:"trackingCodes,omitempty"` // Коды маркировки товаров и транспортных упаковок
	Overhead      *float64            `json:"overhead,omitempty"`      // Накладные расходы. Если Позиции Приемки не заданы, то накладные расходы нельзя задать.
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
	DeleteMany(ctx context.Context, supplyList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Supply, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, supply *Supply, params *Params) (*Supply, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetaAttributesSharedStatesWrapper, *resty.Response, error)
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
	DeleteAttributes(ctx context.Context, attributeList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	GetPublications(ctx context.Context, id *uuid.UUID) (*MetaArray[Publication], *resty.Response, error)
	GetPublicationByID(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (*Publication, *resty.Response, error)
	Publish(ctx context.Context, id *uuid.UUID, template *Templater) (*Publication, *resty.Response, error)
	DeletePublication(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (bool, *resty.Response, error)
	PrintDocument(ctx context.Context, id *uuid.UUID, PrintDocumentArg *PrintDocumentArg) (*PrintFile, *resty.Response, error)
	GetFiles(ctx context.Context, id *uuid.UUID) (*MetaArray[File], *resty.Response, error)
	CreateFile(ctx context.Context, id *uuid.UUID, file *File) (*[]File, *resty.Response, error)
	UpdateFiles(ctx context.Context, id *uuid.UUID, files []*File) (*[]File, *resty.Response, error)
	DeleteFile(ctx context.Context, id *uuid.UUID, fileID *uuid.UUID) (bool, *resty.Response, error)
	DeleteFiles(ctx context.Context, id *uuid.UUID, files *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
	GetStateByID(ctx context.Context, id *uuid.UUID) (*State, *resty.Response, error)
	CreateState(ctx context.Context, state *State) (*State, *resty.Response, error)
	UpdateState(ctx context.Context, id *uuid.UUID, state *State) (*State, *resty.Response, error)
	CreateOrUpdateStates(ctx context.Context, id *uuid.UUID, states []*State) (*[]State, *resty.Response, error)
	DeleteState(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*Supply, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	MoveToTrash(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewSupplyService(client *Client) SupplyService {
	e := NewEndpoint(client, "entity/supply")
	return newMainService[Supply, SupplyPosition, MetaAttributesSharedStatesWrapper, any](e)
}
