package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// RetailDemand Розничная продажа.
// Ключевое слово: retaildemand
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-roznichnaq-prodazha
type RetailDemand struct {
	AccountID           *uuid.UUID                 `json:"accountId,omitempty"`
	Agent               *Counterparty              `json:"agent,omitempty"`
	AgentAccount        *AgentAccount              `json:"agentAccount,omitempty"`
	Applicable          *bool                      `json:"applicable,omitempty"`
	Attributes          *Attributes                `json:"attributes,omitempty"`
	CashSum             *float64                   `json:"cashSum,omitempty"`
	CheckNumber         *string                    `json:"checkNumber,omitempty"`
	CheckSum            *float64                   `json:"checkSum,omitempty"`
	Code                *string                    `json:"code,omitempty"`
	Contract            *Contract                  `json:"contract,omitempty"`
	Created             *Timestamp                 `json:"created,omitempty"`
	CustomerOrder       *CustomerOrder             `json:"customerOrder,omitempty"`
	Deleted             *Timestamp                 `json:"deleted,omitempty"`
	Description         *string                    `json:"description,omitempty"`
	DocumentNumber      *string                    `json:"documentNumber,omitempty"`
	ExternalCode        *string                    `json:"externalCode,omitempty"`
	Files               *Files                     `json:"files,omitempty"`
	Fiscal              *bool                      `json:"fiscal,omitempty"`
	FiscalPrinterInfo   *string                    `json:"fiscalPrinterInfo,omitempty"`
	Group               *Group                     `json:"group,omitempty"`
	ID                  *uuid.UUID                 `json:"id,omitempty"`
	Meta                *Meta                      `json:"meta,omitempty"`
	Moment              *Timestamp                 `json:"moment,omitempty"`
	Name                *string                    `json:"name,omitempty"`
	NoCashSum           *float64                   `json:"noCashSum,omitempty"`
	OfdCode             *string                    `json:"ofdCode,omitempty"`
	Organization        *Organization              `json:"organization,omitempty"`
	OrganizationAccount *AgentAccount              `json:"organizationAccount,omitempty"`
	Owner               *Employee                  `json:"owner,omitempty"`
	PayedSum            *float64                   `json:"payedSum,omitempty"`
	Positions           *Positions[RetailPosition] `json:"positions,omitempty"`
	PrepaymentCashSum   *float64                   `json:"prepaymentCashSum,omitempty"`
	PrepaymentNoCashSum *float64                   `json:"prepaymentNoCashSum,omitempty"`
	PrepaymentQRSum     *float64                   `json:"prepaymentQrSum,omitempty"`
	Printed             *bool                      `json:"printed,omitempty"`
	Project             *Project                   `json:"project,omitempty"`
	Published           *bool                      `json:"published,omitempty"`
	QRSum               *float64                   `json:"qrSum,omitempty"`
	Rate                *Rate                      `json:"rate,omitempty"`
	RetailShift         *RetailShift               `json:"retailShift,omitempty"`
	RetailStore         *RetailStore               `json:"retailStore,omitempty"`
	SessionNumber       *string                    `json:"sessionNumber,omitempty"`
	Shared              *bool                      `json:"shared,omitempty"`
	State               *State                     `json:"state,omitempty"`
	Store               *Store                     `json:"store,omitempty"`
	Sum                 *float64                   `json:"sum,omitempty"`
	SyncID              *uuid.UUID                 `json:"syncId,omitempty"`
	Updated             *Timestamp                 `json:"updated,omitempty"`
	VatEnabled          *bool                      `json:"vatEnabled,omitempty"`
	VatIncluded         *bool                      `json:"vatIncluded,omitempty"`
	VatSum              *float64                   `json:"vatSum,omitempty"`
	TaxSystem           TaxSystem                  `json:"taxSystem,omitempty"`
}

func (r RetailDemand) String() string {
	return Stringify(r)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (r RetailDemand) GetMeta() *Meta {
	return r.Meta
}

func (r RetailDemand) MetaType() MetaType {
	return MetaTypeRetailDemand
}

// RetailDemandPosition позиция розничной продажи.
// Ключевое слово: demandposition
type RetailDemandPosition struct {
	RetailPosition
}

func (r RetailDemandPosition) MetaType() MetaType {
	return MetaTypeRetailDemandPosition
}

// RetailDemandTemplateArg
// Документ: Розничная продажа (retaildemand)
// Основание, на котором он может быть создан:
// - Розничная смена
// - Заказ покупателя
type RetailDemandTemplateArg struct {
	RetailShift   *MetaWrapper `json:"retailShift,omitempty"`
	CustomerOrder *MetaWrapper `json:"customerOrder,omitempty"`
}

// RetailDemandService
// Сервис для работы с розничными продажами.
type RetailDemandService interface {
	GetList(ctx context.Context, params *Params) (*List[RetailDemand], *resty.Response, error)
	Create(ctx context.Context, retailDemand *RetailDemand, params *Params) (*RetailDemand, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, retailDemandList []*RetailDemand, params *Params) (*[]RetailDemand, *resty.Response, error)
	DeleteMany(ctx context.Context, retailDemandList []*RetailDemand) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*RetailDemand, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, retailDemand *RetailDemand, params *Params) (*RetailDemand, *resty.Response, error)
	//endpointTemplate[RetailDemand]
	//endpointTemplateBasedOn[RetailDemand, RetailDemandTemplateArg]
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
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*RetailDemand, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
	Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewRetailDemandService(client *Client) RetailDemandService {
	e := NewEndpoint(client, "entity/retaildemand")
	return newMainService[RetailDemand, RetailPosition, MetadataAttributeSharedStates, any](e)
}
