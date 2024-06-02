package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// PrepaymentReturn Возврат предоплаты.
// Ключевое слово: prepaymentreturn
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vozwrat-predoplaty
type PrepaymentReturn struct {
	Agent        *Counterparty                        `json:"agent,omitempty"`
	Organization *Organization                        `json:"organization,omitempty"`
	Applicable   *bool                                `json:"applicable,omitempty"`
	AccountID    *uuid.UUID                           `json:"accountId,omitempty"`
	CashSum      *float64                             `json:"cashSum,omitempty"`
	Code         *string                              `json:"code,omitempty"`
	Created      *Timestamp                           `json:"created,omitempty"`
	Deleted      *Timestamp                           `json:"deleted,omitempty"`
	Description  *string                              `json:"description,omitempty"`
	ExternalCode *string                              `json:"externalCode,omitempty"`
	Files        *Files                               `json:"files,omitempty"`
	Group        *Group                               `json:"group,omitempty"`
	ID           *uuid.UUID                           `json:"id,omitempty"`
	Meta         *Meta                                `json:"meta,omitempty"`
	Moment       *Timestamp                           `json:"moment,omitempty"`
	Name         *string                              `json:"name,omitempty"`
	NoCashSum    *float64                             `json:"noCashSum,omitempty"`
	Owner        *Employee                            `json:"owner,omitempty"`
	VatIncluded  *bool                                `json:"vatIncluded,omitempty"`
	Positions    *Positions[PrepaymentReturnPosition] `json:"positions,omitempty"`
	Prepayment   *Prepayment                          `json:"prepayment,omitempty"`
	Printed      *bool                                `json:"printed,omitempty"`
	Published    *bool                                `json:"published,omitempty"`
	QRSum        *float64                             `json:"qrSum,omitempty"`
	Rate         *Rate                                `json:"rate,omitempty"`
	RetailShift  *RetailShift                         `json:"retailShift,omitempty"`
	RetailStore  *RetailStore                         `json:"retailStore,omitempty"`
	Shared       *bool                                `json:"shared,omitempty"`
	State        *State                               `json:"state,omitempty"`
	Sum          *float64                             `json:"sum,omitempty"`
	SyncID       *uuid.UUID                           `json:"syncId,omitempty"`
	VatSum       *float64                             `json:"vatSum,omitempty"`
	Updated      *Timestamp                           `json:"updated,omitempty"`
	VatEnabled   *bool                                `json:"vatEnabled,omitempty"`
	TaxSystem    TaxSystem                            `json:"taxSystem,omitempty"`
	Attributes   Attributes                           `json:"attributes,omitempty"`
}

func (p PrepaymentReturn) String() string {
	return Stringify(p)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (p PrepaymentReturn) GetMeta() Meta {
	return Deref(p.Meta)
}

func (p PrepaymentReturn) MetaType() MetaType {
	return MetaTypePrepaymentReturn
}

type PrepaymentReturns = Slice[PrepaymentReturn]

// PrepaymentReturnPosition Позиция Возврата предоплаты.
// Ключевое слово: prepaymentreturnposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vozwrat-predoplaty-atributy-suschnosti-pozicii-vozwrata-predoplaty
type PrepaymentReturnPosition struct {
	PrepaymentPosition
}

func (p PrepaymentReturnPosition) String() string {
	return Stringify(p)
}

func (p PrepaymentReturnPosition) MetaType() MetaType {
	return MetaTypePrepaymentReturnPosition
}

// PrepaymentReturnService
// Сервис для работы с возвратами предоплат.
type PrepaymentReturnService interface {
	GetList(ctx context.Context, params *Params) (*List[PrepaymentReturn], *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*PrepaymentReturn, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id *uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error)
	UpdateAttribute(ctx context.Context, id *uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	GetPositions(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[PrepaymentReturnPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, params *Params) (*PrepaymentReturnPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, position *PrepaymentReturnPosition, params *Params) (*PrepaymentReturnPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id *uuid.UUID, position *PrepaymentReturnPosition) (*PrepaymentReturnPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id *uuid.UUID, positions []*PrepaymentReturnPosition) (*[]PrepaymentReturnPosition, *resty.Response, error)
	DeletePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID) (bool, *resty.Response, error)
	GetPositionTrackingCodes(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID) (*MetaArray[TrackingCode], *resty.Response, error)
	CreateOrUpdatePositionTrackingCodes(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, trackingCodes TrackingCodes) (*[]TrackingCode, *resty.Response, error)
	DeletePositionTrackingCodes(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, trackingCodes TrackingCodes) (*DeleteManyResponse, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*PrepaymentReturn, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewPrepaymentReturnService(client *Client) PrepaymentReturnService {
	e := NewEndpoint(client, "entity/prepaymentreturn")
	return newMainService[PrepaymentReturn, PrepaymentReturnPosition, MetadataAttributeSharedStates, any](e)
}
