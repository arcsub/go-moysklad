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
	Name                *string                               `json:"name,omitempty"`
	Organization        *Organization                         `json:"organization,omitempty"`
	AgentAccount        *AgentAccount                         `json:"agentAccount,omitempty"`
	Applicable          *bool                                 `json:"applicable,omitempty"`
	VatIncluded         *bool                                 `json:"vatIncluded,omitempty"`
	CashSum             *float64                              `json:"cashSum,omitempty"`
	Code                *string                               `json:"code,omitempty"`
	Contract            *Contract                             `json:"contract,omitempty"`
	Created             *Timestamp                            `json:"created,omitempty"`
	Deleted             *Timestamp                            `json:"deleted,omitempty"`
	Demand              *RetailDemand                         `json:"demand,omitempty"`
	Description         *string                               `json:"description,omitempty"`
	ExternalCode        *string                               `json:"externalCode,omitempty"`
	Group               *Group                                `json:"group,omitempty"`
	ID                  *uuid.UUID                            `json:"id,omitempty"`
	Meta                *Meta                                 `json:"meta,omitempty"`
	Moment              *Timestamp                            `json:"moment,omitempty"`
	OrganizationAccount *AgentAccount                         `json:"organizationAccount,omitempty"`
	NoCashSum           *float64                              `json:"noCashSum,omitempty"`
	SyncID              *uuid.UUID                            `json:"syncId,omitempty"`
	AccountID           *uuid.UUID                            `json:"accountId,omitempty"`
	Owner               *Employee                             `json:"owner,omitempty"`
	Positions           *Positions[RetailSalesReturnPosition] `json:"positions,omitempty"`
	Printed             *bool                                 `json:"printed,omitempty"`
	Project             *Project                              `json:"project,omitempty"`
	Published           *bool                                 `json:"published,omitempty"`
	QrSum               *float64                              `json:"qrSum,omitempty"`
	Rate                *Rate                                 `json:"rate,omitempty"`
	RetailShift         *RetailShift                          `json:"retailShift,omitempty"`
	RetailStore         *RetailStore                          `json:"retailStore,omitempty"`
	Shared              *bool                                 `json:"shared,omitempty"`
	State               *State                                `json:"state,omitempty"`
	Store               *Store                                `json:"store,omitempty"`
	Sum                 *float64                              `json:"sum,omitempty"`
	Agent               *Counterparty                         `json:"agent,omitempty"`
	VatSum              *float64                              `json:"vatSum,omitempty"`
	Updated             *Timestamp                            `json:"updated,omitempty"`
	VatEnabled          *bool                                 `json:"vatEnabled,omitempty"`
	TaxSystem           TaxSystem                             `json:"taxSystem,omitempty"`
	Attributes          Slice[AttributeValue]                 `json:"attributes,omitempty"`
}

func (retailSalesReturn RetailSalesReturn) Clean() *RetailSalesReturn {
	return &RetailSalesReturn{Meta: retailSalesReturn.Meta}
}

func (retailSalesReturn RetailSalesReturn) GetName() string {
	return Deref(retailSalesReturn.Name)
}

func (retailSalesReturn RetailSalesReturn) GetOrganization() Organization {
	return Deref(retailSalesReturn.Organization)
}

func (retailSalesReturn RetailSalesReturn) GetAgentAccount() AgentAccount {
	return Deref(retailSalesReturn.AgentAccount)
}

func (retailSalesReturn RetailSalesReturn) GetApplicable() bool {
	return Deref(retailSalesReturn.Applicable)
}

func (retailSalesReturn RetailSalesReturn) GetVatIncluded() bool {
	return Deref(retailSalesReturn.VatIncluded)
}

func (retailSalesReturn RetailSalesReturn) GetCashSum() float64 {
	return Deref(retailSalesReturn.CashSum)
}

func (retailSalesReturn RetailSalesReturn) GetCode() string {
	return Deref(retailSalesReturn.Code)
}

func (retailSalesReturn RetailSalesReturn) GetContract() Contract {
	return Deref(retailSalesReturn.Contract)
}

func (retailSalesReturn RetailSalesReturn) GetCreated() Timestamp {
	return Deref(retailSalesReturn.Created)
}

func (retailSalesReturn RetailSalesReturn) GetDeleted() Timestamp {
	return Deref(retailSalesReturn.Deleted)
}

func (retailSalesReturn RetailSalesReturn) GetDemand() RetailDemand {
	return Deref(retailSalesReturn.Demand)
}

func (retailSalesReturn RetailSalesReturn) GetDescription() string {
	return Deref(retailSalesReturn.Description)
}

func (retailSalesReturn RetailSalesReturn) GetExternalCode() string {
	return Deref(retailSalesReturn.ExternalCode)
}

func (retailSalesReturn RetailSalesReturn) GetGroup() Group {
	return Deref(retailSalesReturn.Group)
}

func (retailSalesReturn RetailSalesReturn) GetID() uuid.UUID {
	return Deref(retailSalesReturn.ID)
}

func (retailSalesReturn RetailSalesReturn) GetMeta() Meta {
	return Deref(retailSalesReturn.Meta)
}

func (retailSalesReturn RetailSalesReturn) GetMoment() Timestamp {
	return Deref(retailSalesReturn.Moment)
}

func (retailSalesReturn RetailSalesReturn) GetOrganizationAccount() AgentAccount {
	return Deref(retailSalesReturn.OrganizationAccount)
}

func (retailSalesReturn RetailSalesReturn) GetNoCashSum() float64 {
	return Deref(retailSalesReturn.NoCashSum)
}

func (retailSalesReturn RetailSalesReturn) GetSyncID() uuid.UUID {
	return Deref(retailSalesReturn.SyncID)
}

func (retailSalesReturn RetailSalesReturn) GetAccountID() uuid.UUID {
	return Deref(retailSalesReturn.AccountID)
}

func (retailSalesReturn RetailSalesReturn) GetOwner() Employee {
	return Deref(retailSalesReturn.Owner)
}

func (retailSalesReturn RetailSalesReturn) GetPositions() Positions[RetailSalesReturnPosition] {
	return Deref(retailSalesReturn.Positions)
}

func (retailSalesReturn RetailSalesReturn) GetPrinted() bool {
	return Deref(retailSalesReturn.Printed)
}

func (retailSalesReturn RetailSalesReturn) GetProject() Project {
	return Deref(retailSalesReturn.Project)
}

func (retailSalesReturn RetailSalesReturn) GetPublished() bool {
	return Deref(retailSalesReturn.Published)
}

func (retailSalesReturn RetailSalesReturn) GetQrSum() float64 {
	return Deref(retailSalesReturn.QrSum)
}

func (retailSalesReturn RetailSalesReturn) GetRate() Rate {
	return Deref(retailSalesReturn.Rate)
}

func (retailSalesReturn RetailSalesReturn) GetRetailShift() RetailShift {
	return Deref(retailSalesReturn.RetailShift)
}

func (retailSalesReturn RetailSalesReturn) GetRetailStore() RetailStore {
	return Deref(retailSalesReturn.RetailStore)
}

func (retailSalesReturn RetailSalesReturn) GetShared() bool {
	return Deref(retailSalesReturn.Shared)
}

func (retailSalesReturn RetailSalesReturn) GetState() State {
	return Deref(retailSalesReturn.State)
}

func (retailSalesReturn RetailSalesReturn) GetStore() Store {
	return Deref(retailSalesReturn.Store)
}

func (retailSalesReturn RetailSalesReturn) GetSum() float64 {
	return Deref(retailSalesReturn.Sum)
}

func (retailSalesReturn RetailSalesReturn) GetAgent() Counterparty {
	return Deref(retailSalesReturn.Agent)
}

func (retailSalesReturn RetailSalesReturn) GetVatSum() float64 {
	return Deref(retailSalesReturn.VatSum)
}

func (retailSalesReturn RetailSalesReturn) GetUpdated() Timestamp {
	return Deref(retailSalesReturn.Updated)
}

func (retailSalesReturn RetailSalesReturn) GetVatEnabled() bool {
	return Deref(retailSalesReturn.VatEnabled)
}

func (retailSalesReturn RetailSalesReturn) GetTaxSystem() TaxSystem {
	return retailSalesReturn.TaxSystem
}

func (retailSalesReturn RetailSalesReturn) GetAttributes() Slice[AttributeValue] {
	return retailSalesReturn.Attributes
}

func (retailSalesReturn *RetailSalesReturn) SetName(name string) *RetailSalesReturn {
	retailSalesReturn.Name = &name
	return retailSalesReturn
}

func (retailSalesReturn *RetailSalesReturn) SetOrganization(organization *Organization) *RetailSalesReturn {
	retailSalesReturn.Organization = organization
	return retailSalesReturn
}

func (retailSalesReturn *RetailSalesReturn) SetAgentAccount(agentAccount *AgentAccount) *RetailSalesReturn {
	retailSalesReturn.AgentAccount = agentAccount
	return retailSalesReturn
}

func (retailSalesReturn *RetailSalesReturn) SetApplicable(applicable bool) *RetailSalesReturn {
	retailSalesReturn.Applicable = &applicable
	return retailSalesReturn
}

func (retailSalesReturn *RetailSalesReturn) SetVatIncluded(vatIncluded bool) *RetailSalesReturn {
	retailSalesReturn.VatIncluded = &vatIncluded
	return retailSalesReturn
}

func (retailSalesReturn *RetailSalesReturn) SetCashSum(cashSum float64) *RetailSalesReturn {
	retailSalesReturn.CashSum = &cashSum
	return retailSalesReturn
}

func (retailSalesReturn *RetailSalesReturn) SetCode(code string) *RetailSalesReturn {
	retailSalesReturn.Code = &code
	return retailSalesReturn
}

func (retailSalesReturn *RetailSalesReturn) SetContract(contract *Contract) *RetailSalesReturn {
	retailSalesReturn.Contract = contract
	return retailSalesReturn
}

func (retailSalesReturn *RetailSalesReturn) SetDemand(demand *RetailDemand) *RetailSalesReturn {
	retailSalesReturn.Demand = demand
	return retailSalesReturn
}

func (retailSalesReturn *RetailSalesReturn) SetDescription(description string) *RetailSalesReturn {
	retailSalesReturn.Description = &description
	return retailSalesReturn
}

func (retailSalesReturn *RetailSalesReturn) SetExternalCode(externalCode string) *RetailSalesReturn {
	retailSalesReturn.ExternalCode = &externalCode
	return retailSalesReturn
}

func (retailSalesReturn *RetailSalesReturn) SetGroup(group *Group) *RetailSalesReturn {
	retailSalesReturn.Group = group
	return retailSalesReturn
}

func (retailSalesReturn *RetailSalesReturn) SetMeta(meta *Meta) *RetailSalesReturn {
	retailSalesReturn.Meta = meta
	return retailSalesReturn
}

func (retailSalesReturn *RetailSalesReturn) SetMoment(moment *Timestamp) *RetailSalesReturn {
	retailSalesReturn.Moment = moment
	return retailSalesReturn
}

func (retailSalesReturn *RetailSalesReturn) SetOrganizationAccount(organizationAccount *AgentAccount) *RetailSalesReturn {
	retailSalesReturn.OrganizationAccount = organizationAccount
	return retailSalesReturn
}

func (retailSalesReturn *RetailSalesReturn) SetNoCashSum(noCashSum float64) *RetailSalesReturn {
	retailSalesReturn.NoCashSum = &noCashSum
	return retailSalesReturn
}

func (retailSalesReturn *RetailSalesReturn) SetSyncID(syncID uuid.UUID) *RetailSalesReturn {
	retailSalesReturn.SyncID = &syncID
	return retailSalesReturn
}

func (retailSalesReturn *RetailSalesReturn) SetOwner(owner *Employee) *RetailSalesReturn {
	retailSalesReturn.Owner = owner
	return retailSalesReturn
}

func (retailSalesReturn *RetailSalesReturn) SetPositions(positions *Positions[RetailSalesReturnPosition]) *RetailSalesReturn {
	retailSalesReturn.Positions = positions
	return retailSalesReturn
}

func (retailSalesReturn *RetailSalesReturn) SetProject(project *Project) *RetailSalesReturn {
	retailSalesReturn.Project = project
	return retailSalesReturn
}

func (retailSalesReturn *RetailSalesReturn) SetQrSum(qrSum float64) *RetailSalesReturn {
	retailSalesReturn.QrSum = &qrSum
	return retailSalesReturn
}

func (retailSalesReturn *RetailSalesReturn) SetRate(rate *Rate) *RetailSalesReturn {
	retailSalesReturn.Rate = rate
	return retailSalesReturn
}

func (retailSalesReturn *RetailSalesReturn) SetRetailShift(retailShift *RetailShift) *RetailSalesReturn {
	retailSalesReturn.RetailShift = retailShift
	return retailSalesReturn
}

func (retailSalesReturn *RetailSalesReturn) SetRetailStore(retailStore *RetailStore) *RetailSalesReturn {
	retailSalesReturn.RetailStore = retailStore
	return retailSalesReturn
}

func (retailSalesReturn *RetailSalesReturn) SetShared(shared bool) *RetailSalesReturn {
	retailSalesReturn.Shared = &shared
	return retailSalesReturn
}

func (retailSalesReturn *RetailSalesReturn) SetState(state *State) *RetailSalesReturn {
	retailSalesReturn.State = state
	return retailSalesReturn
}

func (retailSalesReturn *RetailSalesReturn) SetStore(store *Store) *RetailSalesReturn {
	retailSalesReturn.Store = store
	return retailSalesReturn
}

func (retailSalesReturn *RetailSalesReturn) SetAgent(agent *Counterparty) *RetailSalesReturn {
	retailSalesReturn.Agent = agent
	return retailSalesReturn
}

func (retailSalesReturn *RetailSalesReturn) SetVatSum(vatSum float64) *RetailSalesReturn {
	retailSalesReturn.VatSum = &vatSum
	return retailSalesReturn
}

func (retailSalesReturn *RetailSalesReturn) SetVatEnabled(vatEnabled bool) *RetailSalesReturn {
	retailSalesReturn.VatEnabled = &vatEnabled
	return retailSalesReturn
}

func (retailSalesReturn *RetailSalesReturn) SetTaxSystem(taxSystem TaxSystem) *RetailSalesReturn {
	retailSalesReturn.TaxSystem = taxSystem
	return retailSalesReturn
}

func (retailSalesReturn *RetailSalesReturn) SetAttributes(attributes Slice[AttributeValue]) *RetailSalesReturn {
	retailSalesReturn.Attributes = attributes
	return retailSalesReturn
}

func (retailSalesReturn RetailSalesReturn) String() string {
	return Stringify(retailSalesReturn)
}

func (retailSalesReturn RetailSalesReturn) MetaType() MetaType {
	return MetaTypeRetailSalesReturn
}

// RetailSalesReturnPosition позиция розничного возврата.
// Ключевое слово: salesreturnposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-roznichnyj-wozwrat-roznichnye-wozwraty-pozicii-roznichnogo-wozwrata
type RetailSalesReturnPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`
	Assortment *AssortmentPosition `json:"assortment,omitempty"`
	Cost       *float64            `json:"cost,omitempty"`
	Discount   *float64            `json:"discount,omitempty"`
	ID         *uuid.UUID          `json:"id,omitempty"`
	Pack       *Pack               `json:"pack,omitempty"`
	Price      *float64            `json:"price,omitempty"`
	Quantity   *float64            `json:"quantity,omitempty"`
	Vat        *int                `json:"vat,omitempty"`
	VatEnabled *bool               `json:"vatEnabled,omitempty"`
	Stock      *Stock              `json:"stock,omitempty"`
	Things     Slice[string]       `json:"things,omitempty"`
}

func (retailSalesReturnPosition RetailSalesReturnPosition) GetAccountID() uuid.UUID {
	return Deref(retailSalesReturnPosition.AccountID)
}

func (retailSalesReturnPosition RetailSalesReturnPosition) GetAssortment() AssortmentPosition {
	return Deref(retailSalesReturnPosition.Assortment)
}

func (retailSalesReturnPosition RetailSalesReturnPosition) GetCost() float64 {
	return Deref(retailSalesReturnPosition.Cost)
}

func (retailSalesReturnPosition RetailSalesReturnPosition) GetDiscount() float64 {
	return Deref(retailSalesReturnPosition.Discount)
}

func (retailSalesReturnPosition RetailSalesReturnPosition) GetID() uuid.UUID {
	return Deref(retailSalesReturnPosition.ID)
}

func (retailSalesReturnPosition RetailSalesReturnPosition) GetPack() Pack {
	return Deref(retailSalesReturnPosition.Pack)
}

func (retailSalesReturnPosition RetailSalesReturnPosition) GetPrice() float64 {
	return Deref(retailSalesReturnPosition.Price)
}

func (retailSalesReturnPosition RetailSalesReturnPosition) GetQuantity() float64 {
	return Deref(retailSalesReturnPosition.Quantity)
}

func (retailSalesReturnPosition RetailSalesReturnPosition) GetThings() Slice[string] {
	return retailSalesReturnPosition.Things
}

func (retailSalesReturnPosition RetailSalesReturnPosition) GetVat() int {
	return Deref(retailSalesReturnPosition.Vat)
}

func (retailSalesReturnPosition RetailSalesReturnPosition) GetVatEnabled() bool {
	return Deref(retailSalesReturnPosition.VatEnabled)
}

func (retailSalesReturnPosition RetailSalesReturnPosition) GetStock() Stock {
	return Deref(retailSalesReturnPosition.Stock)
}

func (retailSalesReturnPosition *RetailSalesReturnPosition) SetAssortment(assortment AsAssortment) *RetailSalesReturnPosition {
	retailSalesReturnPosition.Assortment = assortment.AsAssortment()
	return retailSalesReturnPosition
}

func (retailSalesReturnPosition *RetailSalesReturnPosition) SetCost(cost float64) *RetailSalesReturnPosition {
	retailSalesReturnPosition.Cost = &cost
	return retailSalesReturnPosition
}

func (retailSalesReturnPosition *RetailSalesReturnPosition) SetDiscount(discount float64) *RetailSalesReturnPosition {
	retailSalesReturnPosition.Discount = &discount
	return retailSalesReturnPosition
}

func (retailSalesReturnPosition *RetailSalesReturnPosition) SetPack(pack *Pack) *RetailSalesReturnPosition {
	retailSalesReturnPosition.Pack = pack
	return retailSalesReturnPosition
}

func (retailSalesReturnPosition *RetailSalesReturnPosition) SetPrice(price float64) *RetailSalesReturnPosition {
	retailSalesReturnPosition.Price = &price
	return retailSalesReturnPosition
}

func (retailSalesReturnPosition *RetailSalesReturnPosition) SetQuantity(quantity float64) *RetailSalesReturnPosition {
	retailSalesReturnPosition.Quantity = &quantity
	return retailSalesReturnPosition
}

func (retailSalesReturnPosition *RetailSalesReturnPosition) SetVat(vat int) *RetailSalesReturnPosition {
	retailSalesReturnPosition.Vat = &vat
	return retailSalesReturnPosition
}

func (retailSalesReturnPosition *RetailSalesReturnPosition) SetVatEnabled(vatEnabled bool) *RetailSalesReturnPosition {
	retailSalesReturnPosition.VatEnabled = &vatEnabled
	return retailSalesReturnPosition
}

func (retailSalesReturnPosition *RetailSalesReturnPosition) SetThings(things Slice[string]) *RetailSalesReturnPosition {
	retailSalesReturnPosition.Things = things
	return retailSalesReturnPosition
}

func (retailSalesReturnPosition RetailSalesReturnPosition) String() string {
	return Stringify(retailSalesReturnPosition)
}

func (retailSalesReturnPosition RetailSalesReturnPosition) MetaType() MetaType {
	return MetaTypeRetailSalesReturnPosition
}

// RetailSalesReturnService
// Сервис для работы с розничными возвратами.
type RetailSalesReturnService interface {
	GetList(ctx context.Context, params *Params) (*List[RetailSalesReturn], *resty.Response, error)
	Create(ctx context.Context, retailSalesReturn *RetailSalesReturn, params *Params) (*RetailSalesReturn, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, retailSalesReturnList []*RetailSalesReturn, params *Params) (*[]RetailSalesReturn, *resty.Response, error)
	DeleteMany(ctx context.Context, retailSalesReturnList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params *Params) (*RetailSalesReturn, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, retailSalesReturn *RetailSalesReturn, params *Params) (*RetailSalesReturn, *resty.Response, error)
	//endpointTemplate[RetailSalesReturn]
	GetMetadata(ctx context.Context) (*MetaAttributesSharedStatesWrapper, *resty.Response, error)
	GetPositions(ctx context.Context, id uuid.UUID, params *Params) (*MetaArray[RetailSalesReturnPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params *Params) (*RetailSalesReturnPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *RetailSalesReturnPosition, params *Params) (*RetailSalesReturnPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id uuid.UUID, position *RetailSalesReturnPosition) (*RetailSalesReturnPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id uuid.UUID, positions []*RetailSalesReturnPosition) (*[]RetailSalesReturnPosition, *resty.Response, error)
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
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*RetailSalesReturn, *resty.Response, error)
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

func NewRetailSalesReturnService(client *Client) RetailSalesReturnService {
	e := NewEndpoint(client, "entity/retailsalesreturn")
	return newMainService[RetailSalesReturn, RetailSalesReturnPosition, MetaAttributesSharedStatesWrapper, any](e)
}
