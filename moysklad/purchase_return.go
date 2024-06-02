package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// PurchaseReturn Возврат поставщику.
// Ключевое слово: purchasereturn
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vozwrat-postawschiku
type PurchaseReturn struct {
	Printed             *bool                              `json:"printed,omitempty"`
	Supply              *Supply                            `json:"supply,omitempty"`
	AgentAccount        *AgentAccount                      `json:"agentAccount,omitempty"`
	Applicable          *bool                              `json:"applicable,omitempty"`
	Payments            *Payments                          `json:"payments,omitempty"`
	Code                *string                            `json:"code,omitempty"`
	OrganizationAccount *AgentAccount                      `json:"organizationAccount,omitempty"`
	Created             *Timestamp                         `json:"created,omitempty"`
	Deleted             *Timestamp                         `json:"deleted,omitempty"`
	Description         *string                            `json:"description,omitempty"`
	ExternalCode        *string                            `json:"externalCode,omitempty"`
	Files               *Files                             `json:"files,omitempty"`
	Group               *Group                             `json:"group,omitempty"`
	ID                  *uuid.UUID                         `json:"id,omitempty"`
	Meta                *Meta                              `json:"meta,omitempty"`
	Moment              *Timestamp                         `json:"moment,omitempty"`
	Name                *string                            `json:"name,omitempty"`
	AccountID           *uuid.UUID                         `json:"accountId,omitempty"`
	Contract            *Contract                          `json:"contract,omitempty"`
	Agent               *Counterparty                      `json:"agent,omitempty"`
	Organization        *Organization                      `json:"organization,omitempty"`
	Project             *Project                           `json:"project,omitempty"`
	Published           *bool                              `json:"published,omitempty"`
	Rate                *Rate                              `json:"rate,omitempty"`
	Shared              *bool                              `json:"shared,omitempty"`
	State               *State                             `json:"state,omitempty"`
	Store               *Store                             `json:"store,omitempty"`
	Sum                 *Decimal                           `json:"sum,omitempty"`
	SyncID              *uuid.UUID                         `json:"syncId,omitempty"`
	Updated             *Timestamp                         `json:"updated,omitempty"`
	VatEnabled          *bool                              `json:"vatEnabled,omitempty"`
	VatIncluded         *bool                              `json:"vatIncluded,omitempty"`
	VatSum              *Decimal                           `json:"vatSum,omitempty"`
	Positions           *Positions[PurchaseReturnPosition] `json:"positions,omitempty"`
	Owner               *Employee                          `json:"owner,omitempty"`
	FactureIn           *FactureIn                         `json:"factureIn,omitempty"`
	FactureOut          *FactureOut                        `json:"factureOut,omitempty"`
	InvoicedSum         *float64                           `json:"invoicedSum,omitempty"`
	PayedSum            *float64                           `json:"payedSum,omitempty"`
	Attributes          Attributes                         `json:"attributes,omitempty"`
}

func (p PurchaseReturn) String() string {
	return Stringify(p)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (p PurchaseReturn) GetMeta() Meta {
	return Deref(p.Meta)
}

func (p PurchaseReturn) MetaType() MetaType {
	return MetaTypePurchaseReturn
}

type PurchaseReturns []*PurchaseReturn

// PurchaseReturnPosition Позиция Возврата поставщику.
// Ключевое слово: purchasereturnposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vozwrat-postawschiku-vozwraty-postawschikam-pozicii-vozwrata-postawschiku
type PurchaseReturnPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	Discount   *float64            `json:"discount,omitempty"`   // Процент скидки или наценки. Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Price      *float64            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Slot       *Slot               `json:"slot,omitempty"`       // Ячейка на складе
	Things     *Things             `json:"things,omitempty"`     // Серийные номера. Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете. В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута.
	Vat        *int                `json:"vat,omitempty"`        // НДС, которым облагается текущая позиция
	VatEnabled *bool               `json:"vatEnabled,omitempty"` // Включен ли НДС для позиции. С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
	Stock      *Stock              `json:"stock,omitempty"`      // Остатки и себестоимость `?fields=stock&expand=positions`
}

func (p PurchaseReturnPosition) String() string {
	return Stringify(p)
}

func (p PurchaseReturnPosition) MetaType() MetaType {
	return MetaTypePurchaseReturnPosition
}

// PurchaseReturnTemplateArg
// Документ: Возврат поставщику (purchasereturn)
// Основание, на котором он может быть создан:
// - Приемка (supply)
type PurchaseReturnTemplateArg struct {
	Supply *MetaWrapper `json:"supply,omitempty"`
}

// PurchaseReturnService
// Сервис для работы с возвратами поставщикам.
type PurchaseReturnService interface {
	GetList(ctx context.Context, params *Params) (*List[PurchaseReturn], *resty.Response, error)
	Create(ctx context.Context, purchaseReturn *PurchaseReturn, params *Params) (*PurchaseReturn, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, purchaseReturnList []*PurchaseReturn, params *Params) (*[]PurchaseReturn, *resty.Response, error)
	DeleteMany(ctx context.Context, purchaseReturnList []*PurchaseReturn) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*PurchaseReturn, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, purchaseReturn *PurchaseReturn, params *Params) (*PurchaseReturn, *resty.Response, error)
	//endpointTemplate[PurchaseReturn]
	//endpointTemplateBasedOn[PurchaseReturn, PurchaseReturnTemplateArg]
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	GetPositions(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[PurchaseReturnPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, params *Params) (*PurchaseReturnPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, position *PurchaseReturnPosition, params *Params) (*PurchaseReturnPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id *uuid.UUID, position *PurchaseReturnPosition) (*PurchaseReturnPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id *uuid.UUID, positions []*PurchaseReturnPosition) (*[]PurchaseReturnPosition, *resty.Response, error)
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
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*PurchaseReturn, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
	Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewPurchaseReturnService(client *Client) PurchaseReturnService {
	e := NewEndpoint(client, "entity/purchasereturn")
	return newMainService[PurchaseReturn, PurchaseReturnPosition, MetadataAttributeSharedStates, any](e)
}
