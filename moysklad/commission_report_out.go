package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// CommissionReportOut Выданный отчет комиссионера.
// Ключевое слово: commissionreportout
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vydannyj-otchet-komissionera
type CommissionReportOut struct {
	Organization          *Organization                           `json:"organization,omitempty"`
	CommissionPeriodStart *Timestamp                              `json:"commissionPeriodStart,omitempty"`
	AgentAccount          *AgentAccount                           `json:"agentAccount,omitempty"`
	AccountID             *uuid.UUID                              `json:"accountId,omitempty"`
	Attributes            *Attributes                             `json:"attributes,omitempty"`
	Owner                 *Employee                               `json:"owner,omitempty"`
	CommissionPeriodEnd   *Timestamp                              `json:"commissionPeriodEnd,omitempty"`
	OrganizationAccount   *AgentAccount                           `json:"organizationAccount,omitempty"`
	CommitentSum          *Decimal                                `json:"commitentSum,omitempty"`
	Contract              *Contract                               `json:"contract,omitempty"`
	Created               *Timestamp                              `json:"created,omitempty"`
	Deleted               *Timestamp                              `json:"deleted,omitempty"`
	Description           *string                                 `json:"description,omitempty"`
	ExternalCode          *string                                 `json:"externalCode,omitempty"`
	Files                 *Files                                  `json:"files,omitempty"`
	Group                 *Group                                  `json:"group,omitempty"`
	ID                    *uuid.UUID                              `json:"id,omitempty"`
	Meta                  *Meta                                   `json:"meta,omitempty"`
	Moment                *Timestamp                              `json:"moment,omitempty"`
	Name                  *string                                 `json:"name,omitempty"`
	Applicable            *bool                                   `json:"applicable,omitempty"`
	Agent                 *Counterparty                           `json:"agent,omitempty"`
	Code                  *string                                 `json:"code,omitempty"`
	PayedSum              *Decimal                                `json:"payedSum,omitempty"`
	Positions             *Positions[CommissionReportOutPosition] `json:"positions,omitempty"`
	Printed               *bool                                   `json:"printed,omitempty"`
	Project               *Project                                `json:"project,omitempty"`
	Published             *bool                                   `json:"published,omitempty"`
	Rate                  *Rate                                   `json:"rate,omitempty"`
	RewardPercent         *float64                                `json:"rewardPercent,omitempty"`
	Payments              *Payments                               `json:"payments,omitempty"`
	SalesChannel          *SalesChannel                           `json:"salesChannel,omitempty"`
	Shared                *bool                                   `json:"shared,omitempty"`
	State                 *State                                  `json:"state,omitempty"`
	Sum                   *Decimal                                `json:"sum,omitempty"`
	SyncID                *uuid.UUID                              `json:"syncId,omitempty"`
	Updated               *Timestamp                              `json:"updated,omitempty"`
	VatEnabled            *bool                                   `json:"vatEnabled,omitempty"`
	VatIncluded           *bool                                   `json:"vatIncluded,omitempty"`
	VatSum                *Decimal                                `json:"vatSum,omitempty"`
	RewardType            RewardType                              `json:"rewardType,omitempty"`
}

func (c CommissionReportOut) String() string {
	return Stringify(c)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (c CommissionReportOut) GetMeta() *Meta {
	return c.Meta
}

func (c CommissionReportOut) MetaType() MetaType {
	return MetaTypeCommissionReportOut
}

// CommissionReportOutPosition Позиция Выданного отчета комиссионера.
// Ключевое слово: commissionreportoutposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vydannyj-otchet-komissionera-vydannye-otchety-komissionera-pozicii-vydannogo-otcheta-komissionera
type CommissionReportOutPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID сущности
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Meta       *Meta               `json:"meta,omitempty"`       // Метаданные
	Price      *Decimal            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Reward     *Decimal            `json:"reward,omitempty"`     // Вознаграждение
	Vat        *int                `json:"vat,omitempty"`        // НДС, которым облагается текущая позиция
	VatEnabled *bool               `json:"vatEnabled,omitempty"` // Включен ли НДС для позиции. С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
}

func (c CommissionReportOutPosition) String() string {
	return Stringify(c)
}

func (c CommissionReportOutPosition) MetaType() MetaType {
	return MetaTypeCommissionReportOutPosition
}

// CommissionReportOutService
// Сервис для работы с выданными отчётами комиссионера.
type CommissionReportOutService interface {
	GetList(ctx context.Context, params *Params) (*List[CommissionReportOut], *resty.Response, error)
	Create(ctx context.Context, commissionReportOut *CommissionReportOut, params *Params) (*CommissionReportOut, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, commissionReportOutList []*CommissionReportOut, params *Params) (*[]CommissionReportOut, *resty.Response, error)
	DeleteMany(ctx context.Context, commissionReportOutList []*CommissionReportOut) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*CommissionReportOut, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, commissionReportOut *CommissionReportOut, params *Params) (*CommissionReportOut, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	GetPositions(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[CommissionReportOutPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, params *Params) (*CommissionReportOutPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, position *CommissionReportOutPosition, params *Params) (*CommissionReportOutPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id *uuid.UUID, position *CommissionReportOutPosition) (*CommissionReportOutPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id *uuid.UUID, positions []*CommissionReportOutPosition) (*[]CommissionReportOutPosition, *resty.Response, error)
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
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*CommissionReportOut, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
	//endpointTemplate[CommissionReportOut]
	GetPublications(ctx context.Context, id *uuid.UUID) (*MetaArray[Publication], *resty.Response, error)
	GetPublicationByID(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (*Publication, *resty.Response, error)
	Publish(ctx context.Context, id *uuid.UUID, template *Templater) (*Publication, *resty.Response, error)
	DeletePublication(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (bool, *resty.Response, error)
	Remove(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
}

func NewCommissionReportOutService(client *Client) CommissionReportOutService {
	e := NewEndpoint(client, "entity/commissionreportout")
	return newMainService[CommissionReportOut, CommissionReportOutPosition, MetadataAttributeSharedStates, any](e)
}
