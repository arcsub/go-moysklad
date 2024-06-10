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
	Positions           *Positions[SalesReturnPosition] `json:"positions,omitempty"`
	VatSum              *float64                        `json:"vatSum,omitempty"`
	AgentAccount        *AgentAccount                   `json:"agentAccount,omitempty"`
	Applicable          *bool                           `json:"applicable,omitempty"`
	FactureOut          *FactureOut                     `json:"factureOut,omitempty"`
	Code                *string                         `json:"code,omitempty"`
	OrganizationAccount *AgentAccount                   `json:"organizationAccount,omitempty"`
	Created             *Timestamp                      `json:"created,omitempty"`
	Deleted             *Timestamp                      `json:"deleted,omitempty"`
	Description         *string                         `json:"description,omitempty"`
	ExternalCode        *string                         `json:"externalCode,omitempty"`
	Files               *MetaArray[File]                `json:"files,omitempty"`
	Group               *Group                          `json:"group,omitempty"`
	ID                  *uuid.UUID                      `json:"id,omitempty"`
	Meta                *Meta                           `json:"meta,omitempty"`
	Moment              *Timestamp                      `json:"moment,omitempty"`
	Name                *string                         `json:"name,omitempty"`
	AccountID           *uuid.UUID                      `json:"accountId,omitempty"`
	Contract            *Contract                       `json:"contract,omitempty"`
	Agent               *Counterparty                   `json:"agent,omitempty"`
	Organization        *Organization                   `json:"organization,omitempty"`
	Printed             *bool                           `json:"printed,omitempty"`
	Project             *Project                        `json:"project,omitempty"`
	Published           *bool                           `json:"published,omitempty"`
	Rate                *Rate                           `json:"rate,omitempty"`
	SalesChannel        *SalesChannel                   `json:"salesChannel,omitempty"`
	Shared              *bool                           `json:"shared,omitempty"`
	State               *State                          `json:"state,omitempty"`
	Store               *Store                          `json:"store,omitempty"`
	Sum                 *float64                        `json:"sum,omitempty"`
	SyncID              *uuid.UUID                      `json:"syncId,omitempty"`
	Updated             *Timestamp                      `json:"updated,omitempty"`
	VatEnabled          *bool                           `json:"vatEnabled,omitempty"`
	VatIncluded         *bool                           `json:"vatIncluded,omitempty"`
	Owner               *Employee                       `json:"owner,omitempty"`
	Demand              *Demand                         `json:"demand,omitempty"`
	Losses              Slice[Loss]                     `json:"losses,omitempty"`
	Payments            Slice[Payment]                  `json:"payments,omitempty"`
	PayedSum            *float64                        `json:"payedSum,omitempty"`
	Attributes          Slice[AttributeValue]           `json:"attributes,omitempty"`
}

func (salesReturn SalesReturn) GetPositions() Positions[SalesReturnPosition] {
	return Deref(salesReturn.Positions)
}

func (salesReturn SalesReturn) GetVatSum() float64 {
	return Deref(salesReturn.VatSum)
}

func (salesReturn SalesReturn) GetAgentAccount() AgentAccount {
	return Deref(salesReturn.AgentAccount)
}

func (salesReturn SalesReturn) GetApplicable() bool {
	return Deref(salesReturn.Applicable)
}

func (salesReturn SalesReturn) GetFactureOut() FactureOut {
	return Deref(salesReturn.FactureOut)
}

func (salesReturn SalesReturn) GetCode() string {
	return Deref(salesReturn.Code)
}

func (salesReturn SalesReturn) GetOrganizationAccount() AgentAccount {
	return Deref(salesReturn.OrganizationAccount)
}

func (salesReturn SalesReturn) GetCreated() Timestamp {
	return Deref(salesReturn.Created)
}

func (salesReturn SalesReturn) GetDeleted() Timestamp {
	return Deref(salesReturn.Deleted)
}

func (salesReturn SalesReturn) GetDescription() string {
	return Deref(salesReturn.Description)
}

func (salesReturn SalesReturn) GetExternalCode() string {
	return Deref(salesReturn.ExternalCode)
}

func (salesReturn SalesReturn) GetFiles() MetaArray[File] {
	return Deref(salesReturn.Files)
}

func (salesReturn SalesReturn) GetGroup() Group {
	return Deref(salesReturn.Group)
}

func (salesReturn SalesReturn) GetID() uuid.UUID {
	return Deref(salesReturn.ID)
}

func (salesReturn SalesReturn) GetMeta() Meta {
	return Deref(salesReturn.Meta)
}

func (salesReturn SalesReturn) GetMoment() Timestamp {
	return Deref(salesReturn.Moment)
}

func (salesReturn SalesReturn) GetName() string {
	return Deref(salesReturn.Name)
}

func (salesReturn SalesReturn) GetAccountID() uuid.UUID {
	return Deref(salesReturn.AccountID)
}

func (salesReturn SalesReturn) GetContract() Contract {
	return Deref(salesReturn.Contract)
}

func (salesReturn SalesReturn) GetAgent() Counterparty {
	return Deref(salesReturn.Agent)
}

func (salesReturn SalesReturn) GetOrganization() Organization {
	return Deref(salesReturn.Organization)
}

func (salesReturn SalesReturn) GetPrinted() bool {
	return Deref(salesReturn.Printed)
}

func (salesReturn SalesReturn) GetProject() Project {
	return Deref(salesReturn.Project)
}

func (salesReturn SalesReturn) GetPublished() bool {
	return Deref(salesReturn.Published)
}

func (salesReturn SalesReturn) GetRate() Rate {
	return Deref(salesReturn.Rate)
}

func (salesReturn SalesReturn) GetSalesChannel() SalesChannel {
	return Deref(salesReturn.SalesChannel)
}

func (salesReturn SalesReturn) GetShared() bool {
	return Deref(salesReturn.Shared)
}

func (salesReturn SalesReturn) GetState() State {
	return Deref(salesReturn.State)
}

func (salesReturn SalesReturn) GetStore() Store {
	return Deref(salesReturn.Store)
}

func (salesReturn SalesReturn) GetSum() float64 {
	return Deref(salesReturn.Sum)
}

func (salesReturn SalesReturn) GetSyncID() uuid.UUID {
	return Deref(salesReturn.SyncID)
}

func (salesReturn SalesReturn) GetUpdated() Timestamp {
	return Deref(salesReturn.Updated)
}

func (salesReturn SalesReturn) GetVatEnabled() bool {
	return Deref(salesReturn.VatEnabled)
}

func (salesReturn SalesReturn) GetVatIncluded() bool {
	return Deref(salesReturn.VatIncluded)
}

func (salesReturn SalesReturn) GetOwner() Employee {
	return Deref(salesReturn.Owner)
}

func (salesReturn SalesReturn) GetDemand() Demand {
	return Deref(salesReturn.Demand)
}

func (salesReturn SalesReturn) GetLosses() Slice[Loss] {
	return salesReturn.Losses
}

func (salesReturn SalesReturn) GetPayments() Slice[Payment] {
	return salesReturn.Payments
}

func (salesReturn SalesReturn) GetPayedSum() float64 {
	return Deref(salesReturn.PayedSum)
}

func (salesReturn SalesReturn) GetAttributes() Slice[AttributeValue] {
	return salesReturn.Attributes
}

func (salesReturn *SalesReturn) SetPositions(positions *Positions[SalesReturnPosition]) *SalesReturn {
	salesReturn.Positions = positions
	return salesReturn
}

func (salesReturn *SalesReturn) SetAgentAccount(agentAccount *AgentAccount) *SalesReturn {
	salesReturn.AgentAccount = agentAccount
	return salesReturn
}

func (salesReturn *SalesReturn) SetApplicable(applicable bool) *SalesReturn {
	salesReturn.Applicable = &applicable
	return salesReturn
}

func (salesReturn *SalesReturn) SetFactureOut(factureOut *FactureOut) *SalesReturn {
	salesReturn.FactureOut = factureOut
	return salesReturn
}

func (salesReturn *SalesReturn) SetCode(code string) *SalesReturn {
	salesReturn.Code = &code
	return salesReturn
}

func (salesReturn *SalesReturn) SetOrganizationAccount(organizationAccount *AgentAccount) *SalesReturn {
	salesReturn.OrganizationAccount = organizationAccount
	return salesReturn
}

func (salesReturn *SalesReturn) SetDescription(description string) *SalesReturn {
	salesReturn.Description = &description
	return salesReturn
}

func (salesReturn *SalesReturn) SetExternalCode(externalCode string) *SalesReturn {
	salesReturn.ExternalCode = &externalCode
	return salesReturn
}

func (salesReturn *SalesReturn) SetFiles(files Slice[File]) *SalesReturn {
	salesReturn.Files = NewMetaArrayRows(files)
	return salesReturn
}

func (salesReturn *SalesReturn) SetGroup(group *Group) *SalesReturn {
	salesReturn.Group = group
	return salesReturn
}

func (salesReturn *SalesReturn) SetMeta(meta *Meta) *SalesReturn {
	salesReturn.Meta = meta
	return salesReturn
}

func (salesReturn *SalesReturn) SetMoment(moment *Timestamp) *SalesReturn {
	salesReturn.Moment = moment
	return salesReturn
}

func (salesReturn *SalesReturn) SetName(name string) *SalesReturn {
	salesReturn.Name = &name
	return salesReturn
}

func (salesReturn *SalesReturn) SetContract(contract *Contract) *SalesReturn {
	salesReturn.Contract = contract
	return salesReturn
}

func (salesReturn *SalesReturn) SetAgent(agent *Counterparty) *SalesReturn {
	salesReturn.Agent = agent
	return salesReturn
}

func (salesReturn *SalesReturn) SetOrganization(organization *Organization) *SalesReturn {
	salesReturn.Organization = organization
	return salesReturn
}

func (salesReturn *SalesReturn) SetProject(project *Project) *SalesReturn {
	salesReturn.Project = project
	return salesReturn
}

func (salesReturn *SalesReturn) SetRate(rate *Rate) *SalesReturn {
	salesReturn.Rate = rate
	return salesReturn
}

func (salesReturn *SalesReturn) SetSalesChannel(salesChannel *SalesChannel) *SalesReturn {
	salesReturn.SalesChannel = salesChannel
	return salesReturn
}

func (salesReturn *SalesReturn) SetShared(shared bool) *SalesReturn {
	salesReturn.Shared = &shared
	return salesReturn
}

func (salesReturn *SalesReturn) SetState(state *State) *SalesReturn {
	salesReturn.State = state
	return salesReturn
}

func (salesReturn *SalesReturn) SetStore(store *Store) *SalesReturn {
	salesReturn.Store = store
	return salesReturn
}

func (salesReturn *SalesReturn) SetSyncID(syncID uuid.UUID) *SalesReturn {
	salesReturn.SyncID = &syncID
	return salesReturn
}

func (salesReturn *SalesReturn) SetVatEnabled(vatEnabled bool) *SalesReturn {
	salesReturn.VatEnabled = &vatEnabled
	return salesReturn
}

func (salesReturn *SalesReturn) SetVatIncluded(vatIncluded bool) *SalesReturn {
	salesReturn.VatIncluded = &vatIncluded
	return salesReturn
}

func (salesReturn *SalesReturn) SetOwner(owner *Employee) *SalesReturn {
	salesReturn.Owner = owner
	return salesReturn
}

func (salesReturn *SalesReturn) SetDemand(demand *Demand) *SalesReturn {
	salesReturn.Demand = demand
	return salesReturn
}

func (salesReturn *SalesReturn) SetLosses(losses Slice[Loss]) *SalesReturn {
	salesReturn.Losses = losses
	return salesReturn
}

func (salesReturn *SalesReturn) SetPayments(payments Slice[Payment]) *SalesReturn {
	salesReturn.Payments = payments
	return salesReturn
}

func (salesReturn *SalesReturn) SetAttributes(attributes Slice[AttributeValue]) *SalesReturn {
	salesReturn.Attributes = attributes
	return salesReturn
}

func (salesReturn SalesReturn) String() string {
	return Stringify(salesReturn)
}

func (salesReturn SalesReturn) MetaType() MetaType {
	return MetaTypeSalesReturn
}

// SalesReturnPosition Позиция Возврата покупателя.
// Ключевое слово: salesreturnposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vozwrat-pokupatelq-vozwraty-pokupatelej-pozicii-vozwrata-pokupatelq
type SalesReturnPosition struct {
	ID         *uuid.UUID          `json:"id,omitempty"`
	Price      *float64            `json:"price,omitempty"`
	Cost       *float64            `json:"cost,omitempty"`
	Country    *Country            `json:"country,omitempty"`
	Discount   *float64            `json:"discount,omitempty"`
	GTD        *GTD                `json:"gtd,omitempty"`
	Assortment *AssortmentPosition `json:"assortment,omitempty"`
	Pack       *Pack               `json:"pack,omitempty"`
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`
	Quantity   *float64            `json:"quantity,omitempty"`
	Slot       *Slot               `json:"slot,omitempty"`
	Stock      *Stock              `json:"stock,omitempty"`
	Vat        *int                `json:"vat,omitempty"`
	VatEnabled *bool               `json:"vatEnabled,omitempty"`
	Things     Slice[string]       `json:"things,omitempty"`
}

func (salesReturnPosition SalesReturnPosition) GetAccountID() uuid.UUID {
	return Deref(salesReturnPosition.AccountID)
}

func (salesReturnPosition SalesReturnPosition) GetAssortment() AssortmentPosition {
	return Deref(salesReturnPosition.Assortment)
}

func (salesReturnPosition SalesReturnPosition) GetCost() float64 {
	return Deref(salesReturnPosition.Cost)
}

func (salesReturnPosition SalesReturnPosition) GetCountry() Country {
	return Deref(salesReturnPosition.Country)
}

func (salesReturnPosition SalesReturnPosition) GetDiscount() float64 {
	return Deref(salesReturnPosition.Discount)
}

func (salesReturnPosition SalesReturnPosition) GetGTD() GTD {
	return Deref(salesReturnPosition.GTD)
}

func (salesReturnPosition SalesReturnPosition) GetID() uuid.UUID {
	return Deref(salesReturnPosition.ID)
}

func (salesReturnPosition SalesReturnPosition) GetPack() Pack {
	return Deref(salesReturnPosition.Pack)
}

func (salesReturnPosition SalesReturnPosition) GetPrice() float64 {
	return Deref(salesReturnPosition.Price)
}

func (salesReturnPosition SalesReturnPosition) GetQuantity() float64 {
	return Deref(salesReturnPosition.Quantity)
}

func (salesReturnPosition SalesReturnPosition) GetSlot() Slot {
	return Deref(salesReturnPosition.Slot)
}

func (salesReturnPosition SalesReturnPosition) GetThings() Slice[string] {
	return salesReturnPosition.Things
}

func (salesReturnPosition SalesReturnPosition) GetVat() int {
	return Deref(salesReturnPosition.Vat)
}

func (salesReturnPosition SalesReturnPosition) GetVatEnabled() bool {
	return Deref(salesReturnPosition.VatEnabled)
}

func (salesReturnPosition SalesReturnPosition) GetStock() Stock {
	return Deref(salesReturnPosition.Stock)
}

func (salesReturnPosition *SalesReturnPosition) SetAssortment(assortment AsAssortment) *SalesReturnPosition {
	salesReturnPosition.Assortment = assortment.AsAssortment()
	return salesReturnPosition
}

func (salesReturnPosition *SalesReturnPosition) SetCost(cost float64) *SalesReturnPosition {
	salesReturnPosition.Cost = &cost
	return salesReturnPosition
}

func (salesReturnPosition *SalesReturnPosition) SetCountry(country *Country) *SalesReturnPosition {
	salesReturnPosition.Country = country
	return salesReturnPosition
}

func (salesReturnPosition *SalesReturnPosition) SetDiscount(discount float64) *SalesReturnPosition {
	salesReturnPosition.Discount = &discount
	return salesReturnPosition
}

func (salesReturnPosition *SalesReturnPosition) SetGTD(gtd *GTD) *SalesReturnPosition {
	salesReturnPosition.GTD = gtd
	return salesReturnPosition
}

func (salesReturnPosition *SalesReturnPosition) SetPack(pack *Pack) *SalesReturnPosition {
	salesReturnPosition.Pack = pack
	return salesReturnPosition
}

func (salesReturnPosition *SalesReturnPosition) SetPrice(price float64) *SalesReturnPosition {
	salesReturnPosition.Price = &price
	return salesReturnPosition
}

func (salesReturnPosition *SalesReturnPosition) SetQuantity(quantity float64) *SalesReturnPosition {
	salesReturnPosition.Quantity = &quantity
	return salesReturnPosition
}

func (salesReturnPosition *SalesReturnPosition) SetSlot(slot *Slot) *SalesReturnPosition {
	salesReturnPosition.Slot = slot
	return salesReturnPosition
}

func (salesReturnPosition *SalesReturnPosition) SetThings(things Slice[string]) *SalesReturnPosition {
	salesReturnPosition.Things = things
	return salesReturnPosition
}

func (salesReturnPosition *SalesReturnPosition) SetVat(vat int) *SalesReturnPosition {
	salesReturnPosition.Vat = &vat
	return salesReturnPosition
}

func (salesReturnPosition *SalesReturnPosition) SetVatEnabled(vatEnabled bool) *SalesReturnPosition {
	salesReturnPosition.VatEnabled = &vatEnabled
	return salesReturnPosition
}

func (salesReturnPosition SalesReturnPosition) String() string {
	return Stringify(salesReturnPosition)
}

func (salesReturnPosition SalesReturnPosition) MetaType() MetaType {
	return MetaTypeSalesReturnPosition
}

// SalesReturnTemplateArg
// Документ: Возврат покупателя (salesreturn)
// Основание, на котором он может быть создан:
// - Отгрузка (demand)
// - Розничная продажа (retaildemand)
//type SalesReturnTemplateArg struct {
//	Demand       *MetaWrapper `json:"demand,omitempty"`
//	RetailDemand *MetaWrapper `json:"retailDemand,omitempty"`
//}

// SalesReturnService
// Сервис для работы с возвратами покупателей.
type SalesReturnService interface {
	GetList(ctx context.Context, params *Params) (*List[SalesReturn], *resty.Response, error)
	Create(ctx context.Context, salesReturn *SalesReturn, params *Params) (*SalesReturn, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, salesReturnList []*SalesReturn, params *Params) (*[]SalesReturn, *resty.Response, error)
	DeleteMany(ctx context.Context, salesReturnList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params *Params) (*SalesReturn, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, salesReturn *SalesReturn, params *Params) (*SalesReturn, *resty.Response, error)
	//endpointTemplate[SalesReturn]
	//endpointTemplateBasedOn[SalesReturn, SalesReturnTemplateArg]
	GetMetadata(ctx context.Context) (*MetaAttributesSharedStatesWrapper, *resty.Response, error)
	GetPositions(ctx context.Context, id uuid.UUID, params *Params) (*MetaArray[SalesReturnPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params *Params) (*SalesReturnPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *SalesReturnPosition, params *Params) (*SalesReturnPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id uuid.UUID, position *SalesReturnPosition) (*SalesReturnPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id uuid.UUID, positions []*SalesReturnPosition) (*[]SalesReturnPosition, *resty.Response, error)
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
	GetPublications(ctx context.Context, id uuid.UUID) (*MetaArray[Publication], *resty.Response, error)
	GetPublicationByID(ctx context.Context, id uuid.UUID, publicationID uuid.UUID) (*Publication, *resty.Response, error)
	Publish(ctx context.Context, id uuid.UUID, template Templater) (*Publication, *resty.Response, error)
	DeletePublication(ctx context.Context, id uuid.UUID, publicationID uuid.UUID) (bool, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*SalesReturn, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)
	MoveToTrash(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetStateByID(ctx context.Context, id uuid.UUID) (*State, *resty.Response, error)
	CreateState(ctx context.Context, state *State) (*State, *resty.Response, error)
	UpdateState(ctx context.Context, id uuid.UUID, state *State) (*State, *resty.Response, error)
	CreateOrUpdateStates(ctx context.Context, states []*State) (*[]State, *resty.Response, error)
	DeleteState(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
}

func NewSalesReturnService(client *Client) SalesReturnService {
	e := NewEndpoint(client, "entity/salesreturn")
	return newMainService[SalesReturn, SalesReturnPosition, MetaAttributesSharedStatesWrapper, any](e)
}
