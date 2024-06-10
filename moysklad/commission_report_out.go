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
	Applicable            *bool                                   `json:"applicable,omitempty"`
	OrganizationAccount   *AgentAccount                           `json:"organizationAccount,omitempty"`
	AgentAccount          *AgentAccount                           `json:"agentAccount,omitempty"`
	Organization          *Organization                           `json:"organization,omitempty"`
	VatSum                *float64                                `json:"vatSum,omitempty"`
	Code                  *string                                 `json:"code,omitempty"`
	CommissionPeriodEnd   *Timestamp                              `json:"commissionPeriodEnd,omitempty"`
	Agent                 *Counterparty                           `json:"agent,omitempty"`
	CommitentSum          *float64                                `json:"commitentSum,omitempty"`
	Contract              *Contract                               `json:"contract,omitempty"`
	Created               *Timestamp                              `json:"created,omitempty"`
	Deleted               *Timestamp                              `json:"deleted,omitempty"`
	Description           *string                                 `json:"description,omitempty"`
	ExternalCode          *string                                 `json:"externalCode,omitempty"`
	Files                 *MetaArray[File]                        `json:"files,omitempty"`
	Group                 *Group                                  `json:"group,omitempty"`
	ID                    *uuid.UUID                              `json:"id,omitempty"`
	Meta                  *Meta                                   `json:"meta,omitempty"`
	Moment                *Timestamp                              `json:"moment,omitempty"`
	Name                  *string                                 `json:"name,omitempty"`
	AccountID             *uuid.UUID                              `json:"accountId,omitempty"`
	CommissionPeriodStart *Timestamp                              `json:"commissionPeriodStart,omitempty"`
	Owner                 *Employee                               `json:"owner,omitempty"`
	PayedSum              *float64                                `json:"payedSum,omitempty"`
	Positions             *Positions[CommissionReportOutPosition] `json:"positions,omitempty"`
	Printed               *bool                                   `json:"printed,omitempty"`
	Project               *Project                                `json:"project,omitempty"`
	Published             *bool                                   `json:"published,omitempty"`
	Rate                  *Rate                                   `json:"rate,omitempty"`
	RewardPercent         *float64                                `json:"rewardPercent,omitempty"`
	Payments              Slice[Payment]                          `json:"payments,omitempty"`
	SalesChannel          *SalesChannel                           `json:"salesChannel,omitempty"`
	Shared                *bool                                   `json:"shared,omitempty"`
	State                 *State                                  `json:"state,omitempty"`
	Sum                   *float64                                `json:"sum,omitempty"`
	SyncID                *uuid.UUID                              `json:"syncId,omitempty"`
	Updated               *Timestamp                              `json:"updated,omitempty"`
	VatEnabled            *bool                                   `json:"vatEnabled,omitempty"`
	VatIncluded           *bool                                   `json:"vatIncluded,omitempty"`
	RewardType            RewardType                              `json:"rewardType,omitempty"`
	Attributes            Slice[AttributeValue]                   `json:"attributes,omitempty"`
}

func (commissionReportOut CommissionReportOut) GetApplicable() bool {
	return Deref(commissionReportOut.Applicable)
}

func (commissionReportOut CommissionReportOut) GetOrganizationAccount() AgentAccount {
	return Deref(commissionReportOut.OrganizationAccount)
}

func (commissionReportOut CommissionReportOut) GetAgentAccount() AgentAccount {
	return Deref(commissionReportOut.AgentAccount)
}

func (commissionReportOut CommissionReportOut) GetOrganization() Organization {
	return Deref(commissionReportOut.Organization)
}

func (commissionReportOut CommissionReportOut) GetVatSum() float64 {
	return Deref(commissionReportOut.VatSum)
}

func (commissionReportOut CommissionReportOut) GetCode() string {
	return Deref(commissionReportOut.Code)
}

func (commissionReportOut CommissionReportOut) GetCommissionPeriodEnd() Timestamp {
	return Deref(commissionReportOut.CommissionPeriodEnd)
}

func (commissionReportOut CommissionReportOut) GetAgent() Counterparty {
	return Deref(commissionReportOut.Agent)
}

func (commissionReportOut CommissionReportOut) GetCommitentSum() float64 {
	return Deref(commissionReportOut.CommitentSum)
}

func (commissionReportOut CommissionReportOut) GetContract() Contract {
	return Deref(commissionReportOut.Contract)
}

func (commissionReportOut CommissionReportOut) GetCreated() Timestamp {
	return Deref(commissionReportOut.Created)
}

func (commissionReportOut CommissionReportOut) GetDeleted() Timestamp {
	return Deref(commissionReportOut.Deleted)
}

func (commissionReportOut CommissionReportOut) GetDescription() string {
	return Deref(commissionReportOut.Description)
}

func (commissionReportOut CommissionReportOut) GetExternalCode() string {
	return Deref(commissionReportOut.ExternalCode)
}

func (commissionReportOut CommissionReportOut) GetFiles() MetaArray[File] {
	return Deref(commissionReportOut.Files)
}

func (commissionReportOut CommissionReportOut) GetGroup() Group {
	return Deref(commissionReportOut.Group)
}

func (commissionReportOut CommissionReportOut) GetID() uuid.UUID {
	return Deref(commissionReportOut.ID)
}

func (commissionReportOut CommissionReportOut) GetMeta() Meta {
	return Deref(commissionReportOut.Meta)
}

func (commissionReportOut CommissionReportOut) GetMoment() Timestamp {
	return Deref(commissionReportOut.Moment)
}

func (commissionReportOut CommissionReportOut) GetName() string {
	return Deref(commissionReportOut.Name)
}

func (commissionReportOut CommissionReportOut) GetAccountID() uuid.UUID {
	return Deref(commissionReportOut.AccountID)
}

func (commissionReportOut CommissionReportOut) GetCommissionPeriodStart() Timestamp {
	return Deref(commissionReportOut.CommissionPeriodStart)
}

func (commissionReportOut CommissionReportOut) GetOwner() Employee {
	return Deref(commissionReportOut.Owner)
}

func (commissionReportOut CommissionReportOut) GetPayedSum() float64 {
	return Deref(commissionReportOut.PayedSum)
}

func (commissionReportOut CommissionReportOut) GetPositions() Positions[CommissionReportOutPosition] {
	return Deref(commissionReportOut.Positions)
}

func (commissionReportOut CommissionReportOut) GetPrinted() bool {
	return Deref(commissionReportOut.Printed)
}

func (commissionReportOut CommissionReportOut) GetProject() Project {
	return Deref(commissionReportOut.Project)
}

func (commissionReportOut CommissionReportOut) GetPublished() bool {
	return Deref(commissionReportOut.Published)
}

func (commissionReportOut CommissionReportOut) GetRate() Rate {
	return Deref(commissionReportOut.Rate)
}

func (commissionReportOut CommissionReportOut) GetRewardPercent() float64 {
	return Deref(commissionReportOut.RewardPercent)
}

func (commissionReportOut CommissionReportOut) GetPayments() Slice[Payment] {
	return commissionReportOut.Payments
}

func (commissionReportOut CommissionReportOut) GetSalesChannel() SalesChannel {
	return Deref(commissionReportOut.SalesChannel)
}

func (commissionReportOut CommissionReportOut) GetShared() bool {
	return Deref(commissionReportOut.Shared)
}

func (commissionReportOut CommissionReportOut) GetState() State {
	return Deref(commissionReportOut.State)
}

func (commissionReportOut CommissionReportOut) GetSum() float64 {
	return Deref(commissionReportOut.Sum)
}

func (commissionReportOut CommissionReportOut) GetSyncID() uuid.UUID {
	return Deref(commissionReportOut.SyncID)
}

func (commissionReportOut CommissionReportOut) GetUpdated() Timestamp {
	return Deref(commissionReportOut.Updated)
}

func (commissionReportOut CommissionReportOut) GetVatEnabled() bool {
	return Deref(commissionReportOut.VatEnabled)
}

func (commissionReportOut CommissionReportOut) GetVatIncluded() bool {
	return Deref(commissionReportOut.VatIncluded)
}

func (commissionReportOut CommissionReportOut) GetRewardType() RewardType {
	return commissionReportOut.RewardType
}

func (commissionReportOut CommissionReportOut) GetAttributes() Slice[AttributeValue] {
	return commissionReportOut.Attributes
}

func (commissionReportOut *CommissionReportOut) SetApplicable(applicable bool) *CommissionReportOut {
	commissionReportOut.Applicable = &applicable
	return commissionReportOut
}

func (commissionReportOut *CommissionReportOut) SetOrganizationAccount(organizationAccount *AgentAccount) *CommissionReportOut {
	commissionReportOut.OrganizationAccount = organizationAccount
	return commissionReportOut
}

func (commissionReportOut *CommissionReportOut) SetAgentAccount(agentAccount *AgentAccount) *CommissionReportOut {
	commissionReportOut.AgentAccount = agentAccount
	return commissionReportOut
}

func (commissionReportOut *CommissionReportOut) SetOrganization(organization *Organization) *CommissionReportOut {
	commissionReportOut.Organization = organization
	return commissionReportOut
}

func (commissionReportOut *CommissionReportOut) SetCode(code string) *CommissionReportOut {
	commissionReportOut.Code = &code
	return commissionReportOut
}

func (commissionReportOut *CommissionReportOut) SetCommissionPeriodEnd(commissionPeriodEnd *Timestamp) *CommissionReportOut {
	commissionReportOut.CommissionPeriodEnd = commissionPeriodEnd
	return commissionReportOut
}

func (commissionReportOut *CommissionReportOut) SetAgent(agent *Counterparty) *CommissionReportOut {
	commissionReportOut.Agent = agent
	return commissionReportOut
}

func (commissionReportOut *CommissionReportOut) SetContract(contract *Contract) *CommissionReportOut {
	commissionReportOut.Contract = contract
	return commissionReportOut
}

func (commissionReportOut *CommissionReportOut) SetDescription(description string) *CommissionReportOut {
	commissionReportOut.Description = &description
	return commissionReportOut
}

func (commissionReportOut *CommissionReportOut) SetExternalCode(externalCode string) *CommissionReportOut {
	commissionReportOut.ExternalCode = &externalCode
	return commissionReportOut
}

func (commissionReportOut *CommissionReportOut) SetFiles(files Slice[File]) *CommissionReportOut {
	commissionReportOut.Files = NewMetaArrayRows(files)
	return commissionReportOut
}

func (commissionReportOut *CommissionReportOut) SetGroup(group *Group) *CommissionReportOut {
	commissionReportOut.Group = group
	return commissionReportOut
}

func (commissionReportOut *CommissionReportOut) SetMeta(meta *Meta) *CommissionReportOut {
	commissionReportOut.Meta = meta
	return commissionReportOut
}

func (commissionReportOut *CommissionReportOut) SetMoment(moment *Timestamp) *CommissionReportOut {
	commissionReportOut.Moment = moment
	return commissionReportOut
}

func (commissionReportOut *CommissionReportOut) SetName(name string) *CommissionReportOut {
	commissionReportOut.Name = &name
	return commissionReportOut
}

func (commissionReportOut *CommissionReportOut) SetCommissionPeriodStart(commissionPeriodStart *Timestamp) *CommissionReportOut {
	commissionReportOut.CommissionPeriodStart = commissionPeriodStart
	return commissionReportOut
}

func (commissionReportOut *CommissionReportOut) SetOwner(owner *Employee) *CommissionReportOut {
	commissionReportOut.Owner = owner
	return commissionReportOut
}

func (commissionReportOut *CommissionReportOut) SetPositions(positions *Positions[CommissionReportOutPosition]) *CommissionReportOut {
	commissionReportOut.Positions = positions
	return commissionReportOut
}

func (commissionReportOut *CommissionReportOut) SetProject(project *Project) *CommissionReportOut {
	commissionReportOut.Project = project
	return commissionReportOut
}

func (commissionReportOut *CommissionReportOut) SetRate(rate *Rate) *CommissionReportOut {
	commissionReportOut.Rate = rate
	return commissionReportOut
}

func (commissionReportOut *CommissionReportOut) SetRewardPercent(rewardPercent float64) *CommissionReportOut {
	commissionReportOut.RewardPercent = &rewardPercent
	return commissionReportOut
}

func (commissionReportOut *CommissionReportOut) SetPayments(payments Slice[Payment]) *CommissionReportOut {
	commissionReportOut.Payments = payments
	return commissionReportOut
}

func (commissionReportOut *CommissionReportOut) SetSalesChannel(salesChannel *SalesChannel) *CommissionReportOut {
	commissionReportOut.SalesChannel = salesChannel
	return commissionReportOut
}

func (commissionReportOut *CommissionReportOut) SetShared(shared bool) *CommissionReportOut {
	commissionReportOut.Shared = &shared
	return commissionReportOut
}

func (commissionReportOut *CommissionReportOut) SetState(state *State) *CommissionReportOut {
	commissionReportOut.State = state
	return commissionReportOut
}

func (commissionReportOut *CommissionReportOut) SetSyncID(syncID uuid.UUID) *CommissionReportOut {
	commissionReportOut.SyncID = &syncID
	return commissionReportOut
}

func (commissionReportOut *CommissionReportOut) SetVatEnabled(vatEnabled bool) *CommissionReportOut {
	commissionReportOut.VatEnabled = &vatEnabled
	return commissionReportOut
}

func (commissionReportOut *CommissionReportOut) SetVatIncluded(vatIncluded bool) *CommissionReportOut {
	commissionReportOut.VatIncluded = &vatIncluded
	return commissionReportOut
}

func (commissionReportOut *CommissionReportOut) SetRewardType(rewardType RewardType) *CommissionReportOut {
	commissionReportOut.RewardType = rewardType
	return commissionReportOut
}

func (commissionReportOut *CommissionReportOut) SetAttributes(attributes Slice[AttributeValue]) *CommissionReportOut {
	commissionReportOut.Attributes = attributes
	return commissionReportOut
}

func (commissionReportOut CommissionReportOut) String() string {
	return Stringify(commissionReportOut)
}

func (commissionReportOut CommissionReportOut) MetaType() MetaType {
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
	Price      *float64            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Reward     *float64            `json:"reward,omitempty"`     // Вознаграждение
	Vat        *int                `json:"vat,omitempty"`        // НДС, которым облагается текущая позиция
	VatEnabled *bool               `json:"vatEnabled,omitempty"` // Включен ли НДС для позиции. С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
}

func (commissionReportOutPosition CommissionReportOutPosition) GetAccountID() uuid.UUID {
	return Deref(commissionReportOutPosition.AccountID)
}

func (commissionReportOutPosition CommissionReportOutPosition) GetAssortment() AssortmentPosition {
	return Deref(commissionReportOutPosition.Assortment)
}

func (commissionReportOutPosition CommissionReportOutPosition) GetID() uuid.UUID {
	return Deref(commissionReportOutPosition.ID)
}

func (commissionReportOutPosition CommissionReportOutPosition) GetPack() Pack {
	return Deref(commissionReportOutPosition.Pack)
}

func (commissionReportOutPosition CommissionReportOutPosition) GetMeta() Meta {
	return Deref(commissionReportOutPosition.Meta)
}

func (commissionReportOutPosition CommissionReportOutPosition) GetPrice() float64 {
	return Deref(commissionReportOutPosition.Price)
}

func (commissionReportOutPosition CommissionReportOutPosition) GetQuantity() float64 {
	return Deref(commissionReportOutPosition.Quantity)
}

func (commissionReportOutPosition CommissionReportOutPosition) GetReward() float64 {
	return Deref(commissionReportOutPosition.Reward)
}

func (commissionReportOutPosition CommissionReportOutPosition) GetVat() int {
	return Deref(commissionReportOutPosition.Vat)
}

func (commissionReportOutPosition CommissionReportOutPosition) GetVatEnabled() bool {
	return Deref(commissionReportOutPosition.VatEnabled)
}

func (commissionReportOutPosition *CommissionReportOutPosition) SetAssortment(assortment *AssortmentPosition) *CommissionReportOutPosition {
	commissionReportOutPosition.Assortment = assortment
	return commissionReportOutPosition
}

func (commissionReportOutPosition *CommissionReportOutPosition) SetPack(pack *Pack) *CommissionReportOutPosition {
	commissionReportOutPosition.Pack = pack
	return commissionReportOutPosition
}

func (commissionReportOutPosition *CommissionReportOutPosition) SetMeta(meta *Meta) *CommissionReportOutPosition {
	commissionReportOutPosition.Meta = meta
	return commissionReportOutPosition
}

func (commissionReportOutPosition *CommissionReportOutPosition) SetPrice(price *float64) *CommissionReportOutPosition {
	commissionReportOutPosition.Price = price
	return commissionReportOutPosition
}

func (commissionReportOutPosition *CommissionReportOutPosition) SetQuantity(quantity float64) *CommissionReportOutPosition {
	commissionReportOutPosition.Quantity = &quantity
	return commissionReportOutPosition
}

func (commissionReportOutPosition *CommissionReportOutPosition) SetReward(reward *float64) *CommissionReportOutPosition {
	commissionReportOutPosition.Reward = reward
	return commissionReportOutPosition
}

func (commissionReportOutPosition *CommissionReportOutPosition) SetVat(vat int) *CommissionReportOutPosition {
	commissionReportOutPosition.Vat = &vat
	return commissionReportOutPosition
}

func (commissionReportOutPosition *CommissionReportOutPosition) SetVatEnabled(vatEnabled bool) *CommissionReportOutPosition {
	commissionReportOutPosition.VatEnabled = &vatEnabled
	return commissionReportOutPosition
}

func (commissionReportOutPosition CommissionReportOutPosition) String() string {
	return Stringify(commissionReportOutPosition)
}

func (commissionReportOutPosition CommissionReportOutPosition) MetaType() MetaType {
	return MetaTypeCommissionReportOutPosition
}

// CommissionReportOutService
// Сервис для работы с выданными отчётами комиссионера.
type CommissionReportOutService interface {
	GetList(ctx context.Context, params *Params) (*List[CommissionReportOut], *resty.Response, error)
	Create(ctx context.Context, commissionReportOut *CommissionReportOut, params *Params) (*CommissionReportOut, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, commissionReportOutList []*CommissionReportOut, params *Params) (*[]CommissionReportOut, *resty.Response, error)
	DeleteMany(ctx context.Context, commissionReportOutList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params *Params) (*CommissionReportOut, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, commissionReportOut *CommissionReportOut, params *Params) (*CommissionReportOut, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetaAttributesSharedStatesWrapper, *resty.Response, error)
	GetPositions(ctx context.Context, id uuid.UUID, params *Params) (*MetaArray[CommissionReportOutPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params *Params) (*CommissionReportOutPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *CommissionReportOutPosition, params *Params) (*CommissionReportOutPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id uuid.UUID, position *CommissionReportOutPosition) (*CommissionReportOutPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id uuid.UUID, positions []*CommissionReportOutPosition) (*[]CommissionReportOutPosition, *resty.Response, error)
	DeletePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (bool, *resty.Response, error)
	GetPositionTrackingCodes(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (*MetaArray[TrackingCode], *resty.Response, error)
	CreateOrUpdatePositionTrackingCodes(ctx context.Context, id uuid.UUID, positionID uuid.UUID, trackingCodes Slice[TrackingCode]) (*[]TrackingCode, *resty.Response, error)
	DeletePositionTrackingCodes(ctx context.Context, id uuid.UUID, positionID uuid.UUID, trackingCodes Slice[TrackingCode]) (*DeleteManyResponse, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error)
	UpdateAttribute(ctx context.Context, id uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*CommissionReportOut, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)
	//endpointTemplate[CommissionReportOut]
	GetPublications(ctx context.Context, id uuid.UUID) (*MetaArray[Publication], *resty.Response, error)
	GetPublicationByID(ctx context.Context, id uuid.UUID, publicationID uuid.UUID) (*Publication, *resty.Response, error)
	Publish(ctx context.Context, id uuid.UUID, template Templater) (*Publication, *resty.Response, error)
	DeletePublication(ctx context.Context, id uuid.UUID, publicationID uuid.UUID) (bool, *resty.Response, error)
	MoveToTrash(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetStateByID(ctx context.Context, id uuid.UUID) (*State, *resty.Response, error)
	CreateState(ctx context.Context, state *State) (*State, *resty.Response, error)
	UpdateState(ctx context.Context, id uuid.UUID, state *State) (*State, *resty.Response, error)
	CreateOrUpdateStates(ctx context.Context, states []*State) (*[]State, *resty.Response, error)
	DeleteState(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
}

func NewCommissionReportOutService(client *Client) CommissionReportOutService {
	e := NewEndpoint(client, "entity/commissionreportout")
	return newMainService[CommissionReportOut, CommissionReportOutPosition, MetaAttributesSharedStatesWrapper, any](e)
}
