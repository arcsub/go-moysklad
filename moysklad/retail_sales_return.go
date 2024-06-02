package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// RetailSalesReturn Розничный возврат.
// Ключевое слово: retailsalesreturn
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-roznichnyj-wozwrat
type RetailSalesReturn struct {
	Name                *string                    `json:"name,omitempty"`
	Organization        *Organization              `json:"organization,omitempty"`
	AgentAccount        *AgentAccount              `json:"agentAccount,omitempty"`
	Applicable          *bool                      `json:"applicable,omitempty"`
	VatIncluded         *bool                      `json:"vatIncluded,omitempty"`
	CashSum             *Decimal                   `json:"cashSum,omitempty"`
	Code                *string                    `json:"code,omitempty"`
	Contract            *Contract                  `json:"contract,omitempty"`
	Created             *Timestamp                 `json:"created,omitempty"`
	Deleted             *Timestamp                 `json:"deleted,omitempty"`
	Demand              *RetailDemand              `json:"demand,omitempty"`
	Description         *string                    `json:"description,omitempty"`
	ExternalCode        *string                    `json:"externalCode,omitempty"`
	Group               *Group                     `json:"group,omitempty"`
	ID                  *uuid.UUID                 `json:"id,omitempty"`
	Meta                *Meta                      `json:"meta,omitempty"`
	Moment              *Timestamp                 `json:"moment,omitempty"`
	OrganizationAccount *AgentAccount              `json:"organizationAccount,omitempty"`
	NoCashSum           *Decimal                   `json:"noCashSum,omitempty"`
	SyncID              *uuid.UUID                 `json:"syncId,omitempty"`
	AccountID           *uuid.UUID                 `json:"accountId,omitempty"`
	Owner               *Employee                  `json:"owner,omitempty"`
	Positions           *Positions[RetailPosition] `json:"positions,omitempty"`
	Printed             *bool                      `json:"printed,omitempty"`
	Project             *Project                   `json:"project,omitempty"`
	Published           *bool                      `json:"published,omitempty"`
	QrSum               *Decimal                   `json:"qrSum,omitempty"`
	Rate                *Rate                      `json:"rate,omitempty"`
	RetailShift         *RetailShift               `json:"retailShift,omitempty"`
	RetailStore         *RetailStore               `json:"retailStore,omitempty"`
	Shared              *bool                      `json:"shared,omitempty"`
	State               *State                     `json:"state,omitempty"`
	Store               *Store                     `json:"store,omitempty"`
	Sum                 *Decimal                   `json:"sum,omitempty"`
	Agent               *Counterparty              `json:"agent,omitempty"`
	VatSum              *Decimal                   `json:"vatSum,omitempty"`
	Updated             *Timestamp                 `json:"updated,omitempty"`
	VatEnabled          *bool                      `json:"vatEnabled,omitempty"`
	TaxSystem           TaxSystem                  `json:"taxSystem,omitempty"`
	Attributes          Attributes                 `json:"attributes,omitempty"`
}

func (r RetailSalesReturn) String() string {
	return Stringify(r)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (r RetailSalesReturn) GetMeta() Meta {
	return Deref(r.Meta)
}

func (r RetailSalesReturn) MetaType() MetaType {
	return MetaTypeRetailSalesReturn
}

// RetailSalesReturnPosition позиция розничного возврата.
// Ключевое слово: salesreturnposition
type RetailSalesReturnPosition struct {
	RetailPosition
}

func (r RetailSalesReturnPosition) MetaType() MetaType {
	return MetaTypeRetailSalesReturnPosition
}

// RetailSalesReturnService
// Сервис для работы с розничными возвратами.
type RetailSalesReturnService interface {
	GetList(ctx context.Context, params *Params) (*List[RetailSalesReturn], *resty.Response, error)
	Create(ctx context.Context, retailSalesReturn *RetailSalesReturn, params *Params) (*RetailSalesReturn, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, retailSalesReturnList []*RetailSalesReturn, params *Params) (*[]RetailSalesReturn, *resty.Response, error)
	DeleteMany(ctx context.Context, retailSalesReturnList []*RetailSalesReturn) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*RetailSalesReturn, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, retailSalesReturn *RetailSalesReturn, params *Params) (*RetailSalesReturn, *resty.Response, error)
	//endpointTemplate[RetailSalesReturn]
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	GetPositions(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[RetailPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, params *Params) (*RetailPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, position *RetailPosition, params *Params) (*RetailPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id *uuid.UUID, position *RetailPosition) (*RetailPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id *uuid.UUID, positions []*RetailPosition) (*[]RetailPosition, *resty.Response, error)
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
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*RetailSalesReturn, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
	Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewRetailSalesReturnService(client *Client) RetailSalesReturnService {
	e := NewEndpoint(client, "entity/retailsalesreturn")
	return newMainService[RetailSalesReturn, RetailPosition, MetadataAttributeSharedStates, any](e)
}
