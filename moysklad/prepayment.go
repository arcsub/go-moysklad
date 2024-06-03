package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Prepayment Предоплата.
// Ключевое слово: prepayment
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-predoplata
type Prepayment struct {
	Returns       *PrepaymentReturns             `json:"returns,omitempty"`
	Owner         *Employee                      `json:"owner,omitempty"`
	Applicable    *bool                          `json:"applicable,omitempty"`
	Agent         *Counterparty                  `json:"agent,omitempty"`
	CashSum       *Decimal                       `json:"cashSum,omitempty"`
	Code          *string                        `json:"code,omitempty"`
	Created       *Timestamp                     `json:"created,omitempty"`
	CustomerOrder *CustomerOrder                 `json:"customerOrder,omitempty"`
	Deleted       *Timestamp                     `json:"deleted,omitempty"`
	Description   *string                        `json:"description,omitempty"`
	ExternalCode  *string                        `json:"externalCode,omitempty"`
	Files         *Files                         `json:"files,omitempty"`
	Group         *Group                         `json:"group,omitempty"`
	ID            *uuid.UUID                     `json:"id,omitempty"`
	Meta          *Meta                          `json:"meta,omitempty"`
	Moment        *Timestamp                     `json:"moment,omitempty"`
	Name          *string                        `json:"name,omitempty"`
	NoCashSum     *Decimal                       `json:"noCashSum,omitempty"`
	AccountID     *uuid.UUID                     `json:"accountId,omitempty"`
	VatIncluded   *bool                          `json:"vatIncluded,omitempty"`
	Positions     *Positions[PrepaymentPosition] `json:"positions,omitempty"`
	Printed       *bool                          `json:"printed,omitempty"`
	Published     *bool                          `json:"published,omitempty"`
	QRSum         *Decimal                       `json:"qrSum,omitempty"`
	Rate          *Rate                          `json:"rate,omitempty"`
	RetailShift   *RetailShift                   `json:"retailShift,omitempty"`
	RetailStore   *RetailStore                   `json:"retailStore,omitempty"`
	Organization  *Organization                  `json:"organization,omitempty"`
	Shared        *bool                          `json:"shared,omitempty"`
	State         *State                         `json:"state,omitempty"`
	Sum           *float64                       `json:"sum,omitempty"`
	SyncID        *uuid.UUID                     `json:"syncId,omitempty"`
	VatSum        *float64                       `json:"vatSum,omitempty"`
	Updated       *Timestamp                     `json:"updated,omitempty"`
	VatEnabled    *bool                          `json:"vatEnabled,omitempty"`
	TaxSystem     TaxSystem                      `json:"taxSystem,omitempty"`
	Attributes    Attributes                     `json:"attributes,omitempty"`
}

func (p Prepayment) String() string {
	return Stringify(p)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (p Prepayment) GetMeta() Meta {
	return Deref(p.Meta)
}

func (p Prepayment) MetaType() MetaType {
	return MetaTypePrepayment
}

type Prepayments = Slice[Prepayment]

// PrepaymentPosition Позиция Предоплаты.
// Ключевое слово: prepaymentposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-predoplata-predoplaty-pozicii-predoplaty
type PrepaymentPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	Discount   *float64            `json:"discount,omitempty"`   // Процент скидки или наценки. Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Price      *float64            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Vat        *int                `json:"vat,omitempty"`        // НДС, которым облагается текущая позиция
	VatEnabled *bool               `json:"vatEnabled,omitempty"` // Включен ли НДС для позиции. С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
}

func (p PrepaymentPosition) String() string {
	return Stringify(p)
}

func (p PrepaymentPosition) MetaType() MetaType {
	return MetaTypePrepaymentPosition
}

// PrepaymentService
// Сервис для работы с предоплатами.
type PrepaymentService interface {
	GetList(ctx context.Context, params *Params) (*List[Prepayment], *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Prepayment, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	GetPositions(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[PrepaymentPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, params *Params) (*PrepaymentPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, position *PrepaymentPosition, params *Params) (*PrepaymentPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id *uuid.UUID, position *PrepaymentPosition) (*PrepaymentPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id *uuid.UUID, positions []*PrepaymentPosition) (*[]PrepaymentPosition, *resty.Response, error)
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
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*Prepayment, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	MoveToTrash(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewPrepaymentService(client *Client) PrepaymentService {
	e := NewEndpoint(client, "entity/prepayment")
	return newMainService[Prepayment, PrepaymentPosition, MetadataAttributeSharedStates, any](e)
}
