package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Demand Отгрузка.
// Ключевое слово: demand
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-otgruzka
type Demand struct {
	AccountID               *uuid.UUID                 `json:"accountId,omitempty"`
	Agent                   *Counterparty              `json:"agent,omitempty"`
	AgentAccount            *AgentAccount              `json:"agentAccount,omitempty"`
	Applicable              *bool                      `json:"applicable,omitempty"`
	Code                    *string                    `json:"code,omitempty"`
	Contract                *NullValue[Contract]       `json:"contract,omitempty"`
	Created                 *Timestamp                 `json:"created,omitempty"`
	Deleted                 *Timestamp                 `json:"deleted,omitempty"`
	Description             *string                    `json:"description,omitempty"`
	ExternalCode            *string                    `json:"externalCode,omitempty"`
	Files                   *MetaArray[File]           `json:"files,omitempty"`
	Group                   *Group                     `json:"group,omitempty"`
	ID                      *uuid.UUID                 `json:"id,omitempty"`
	Meta                    *Meta                      `json:"meta,omitempty"`
	Moment                  *Timestamp                 `json:"moment,omitempty"`
	Name                    *string                    `json:"name,omitempty"`
	Organization            *Organization              `json:"organization,omitempty"`
	OrganizationAccount     *AgentAccount              `json:"organizationAccount,omitempty"`
	Overhead                *Overhead                  `json:"overhead,omitempty"`
	Owner                   *Employee                  `json:"owner,omitempty"`
	PayedSum                *float64                   `json:"payedSum,omitempty"`
	Positions               *Positions[DemandPosition] `json:"positions,omitempty"`
	Printed                 *bool                      `json:"printed,omitempty"`
	Project                 *NullValue[Project]        `json:"project,omitempty"`
	Published               *bool                      `json:"published,omitempty"`
	Rate                    *Rate                      `json:"rate,omitempty"`
	SalesChannel            *NullValue[SalesChannel]   `json:"salesChannel,omitempty"`
	Shared                  *bool                      `json:"shared,omitempty"`
	ShipmentAddress         *string                    `json:"shipmentAddress,omitempty"`
	ShipmentAddressFull     *Address                   `json:"shipmentAddressFull,omitempty"`
	State                   *NullValue[State]          `json:"state,omitempty"`
	Store                   *Store                     `json:"store,omitempty"`
	Sum                     *float64                   `json:"sum,omitempty"`
	SyncID                  *uuid.UUID                 `json:"syncId,omitempty"`
	Updated                 *Timestamp                 `json:"updated,omitempty"`
	VatEnabled              *bool                      `json:"vatEnabled,omitempty"`
	VatIncluded             *bool                      `json:"vatIncluded,omitempty"`
	VatSum                  *float64                   `json:"vatSum,omitempty"`
	CustomerOrder           *CustomerOrder             `json:"customerOrder,omitempty"`
	FactureOut              *FactureOut                `json:"factureOut,omitempty"`
	Returns                 Slice[SalesReturn]         `json:"returns,omitempty"`
	Payments                Slice[Payment]             `json:"payments,omitempty"`
	InvoicesOut             Slice[InvoiceOut]          `json:"invoicesOut,omitempty"`
	CargoName               *string                    `json:"cargoName,omitempty"`
	Carrier                 *Counterparty              `json:"carrier,omitempty"`
	Consignee               *Counterparty              `json:"consignee,omitempty"`
	GoodPackQuantity        *int                       `json:"goodPackQuantity,omitempty"`
	ShippingInstructions    *string                    `json:"shippingInstructions,omitempty"`
	StateContractID         *string                    `json:"stateContractId,omitempty"`
	TransportFacility       *string                    `json:"transportFacility,omitempty"`
	TransportFacilityNumber *string                    `json:"transportFacilityNumber,omitempty"`
	Attributes              Slice[AttributeValue]      `json:"attributes,omitempty"`
}

// Clean возвращает сущность с единственным заполненным полем Meta
func (demand Demand) Clean() *Demand {
	return &Demand{Meta: demand.Meta}
}

// AsOperation возвращает объект Operation c полем Meta сущности
func (demand Demand) AsOperation() *Operation {
	return &Operation{Meta: demand.GetMeta()}
}

func (demand Demand) GetAccountID() uuid.UUID {
	return Deref(demand.AccountID)
}

func (demand Demand) GetAgent() Counterparty {
	return Deref(demand.Agent)
}

func (demand Demand) GetAgentAccount() AgentAccount {
	return Deref(demand.AgentAccount)
}

func (demand Demand) GetApplicable() bool {
	return Deref(demand.Applicable)
}

func (demand Demand) GetCode() string {
	return Deref(demand.Code)
}

func (demand Demand) GetContract() Contract {
	return demand.Contract.Get()
}

func (demand Demand) GetCreated() Timestamp {
	return Deref(demand.Created)
}

func (demand Demand) GetDeleted() Timestamp {
	return Deref(demand.Deleted)
}

func (demand Demand) GetDescription() string {
	return Deref(demand.Description)
}

func (demand Demand) GetExternalCode() string {
	return Deref(demand.ExternalCode)
}

func (demand Demand) GetFiles() MetaArray[File] {
	return Deref(demand.Files)
}

func (demand Demand) GetGroup() Group {
	return Deref(demand.Group)
}

func (demand Demand) GetID() uuid.UUID {
	return Deref(demand.ID)
}

func (demand Demand) GetMeta() Meta {
	return Deref(demand.Meta)
}

func (demand Demand) GetMoment() Timestamp {
	return Deref(demand.Moment)
}

func (demand Demand) GetName() string {
	return Deref(demand.Name)
}

func (demand Demand) GetOrganization() Organization {
	return Deref(demand.Organization)
}

func (demand Demand) GetOrganizationAccount() AgentAccount {
	return Deref(demand.OrganizationAccount)
}

func (demand Demand) GetOverhead() Overhead {
	return Deref(demand.Overhead)
}

func (demand Demand) GetOwner() Employee {
	return Deref(demand.Owner)
}

func (demand Demand) GetPayedSum() float64 {
	return Deref(demand.PayedSum)
}

func (demand Demand) GetPositions() Positions[DemandPosition] {
	return Deref(demand.Positions)
}

func (demand Demand) GetPrinted() bool {
	return Deref(demand.Printed)
}

func (demand Demand) GetProject() Project {
	return demand.Project.Get()
}

func (demand Demand) GetPublished() bool {
	return Deref(demand.Published)
}

func (demand Demand) GetRate() Rate {
	return Deref(demand.Rate)
}

func (demand Demand) GetSalesChannel() SalesChannel {
	return demand.SalesChannel.Get()
}

func (demand Demand) GetShared() bool {
	return Deref(demand.Shared)
}

func (demand Demand) GetShipmentAddress() string {
	return Deref(demand.ShipmentAddress)
}

func (demand Demand) GetShipmentAddressFull() Address {
	return Deref(demand.ShipmentAddressFull)
}

func (demand Demand) GetState() State {
	return demand.State.Get()
}

func (demand Demand) GetStore() Store {
	return Deref(demand.Store)
}

func (demand Demand) GetSum() float64 {
	return Deref(demand.Sum)
}

func (demand Demand) GetSyncID() uuid.UUID {
	return Deref(demand.SyncID)
}

func (demand Demand) GetUpdated() Timestamp {
	return Deref(demand.Updated)
}

func (demand Demand) GetVatEnabled() bool {
	return Deref(demand.VatEnabled)
}

func (demand Demand) GetVatIncluded() bool {
	return Deref(demand.VatIncluded)
}

func (demand Demand) GetVatSum() float64 {
	return Deref(demand.VatSum)
}

func (demand Demand) GetCustomerOrder() CustomerOrder {
	return Deref(demand.CustomerOrder)
}

func (demand Demand) GetFactureOut() FactureOut {
	return Deref(demand.FactureOut)
}

func (demand Demand) GetReturns() Slice[SalesReturn] {
	return demand.Returns
}

func (demand Demand) GetPayments() Slice[Payment] {
	return demand.Payments
}

func (demand Demand) GetInvoicesOut() Slice[InvoiceOut] {
	return demand.InvoicesOut
}

func (demand Demand) GetCargoName() string {
	return Deref(demand.CargoName)
}

func (demand Demand) GetCarrier() Counterparty {
	return Deref(demand.Carrier)
}

func (demand Demand) GetConsignee() Counterparty {
	return Deref(demand.Consignee)
}

func (demand Demand) GetGoodPackQuantity() int {
	return Deref(demand.GoodPackQuantity)
}

func (demand Demand) GetShippingInstructions() string {
	return Deref(demand.ShippingInstructions)
}

func (demand Demand) GetStateContractID() string {
	return Deref(demand.StateContractID)
}

func (demand Demand) GetTransportFacility() string {
	return Deref(demand.TransportFacility)
}

func (demand Demand) GetTransportFacilityNumber() string {
	return Deref(demand.TransportFacilityNumber)
}

func (demand Demand) GetAttributes() Slice[AttributeValue] {
	return demand.Attributes
}

func (demand *Demand) SetAgent(agent *Counterparty) *Demand {
	demand.Agent = agent.Clean()
	return demand
}

func (demand *Demand) SetAgentAccount(agentAccount *AgentAccount) *Demand {
	demand.AgentAccount = agentAccount.Clean()
	return demand
}

func (demand *Demand) SetApplicable(applicable bool) *Demand {
	demand.Applicable = &applicable
	return demand
}

func (demand *Demand) SetCode(code string) *Demand {
	demand.Code = &code
	return demand
}

func (demand *Demand) SetContract(contract *Contract) *Demand {
	demand.Contract = NewNullValueWith(contract.Clean())
	return demand
}

func (demand *Demand) SetNullContract() *Demand {
	demand.Contract = NewNullValue[Contract]()
	return demand
}

func (demand *Demand) SetDescription(description string) *Demand {
	demand.Description = &description
	return demand
}

func (demand *Demand) SetExternalCode(externalCode string) *Demand {
	demand.ExternalCode = &externalCode
	return demand
}

func (demand *Demand) SetFiles(files Slice[File]) *Demand {
	demand.Files = NewMetaArrayRows(files)
	return demand
}

func (demand *Demand) SetGroup(group *Group) *Demand {
	demand.Group = group.Clean()
	return demand
}

func (demand *Demand) SetMeta(meta *Meta) *Demand {
	demand.Meta = meta
	return demand
}

func (demand *Demand) SetMoment(moment *Timestamp) *Demand {
	demand.Moment = moment
	return demand
}

func (demand *Demand) SetName(name string) *Demand {
	demand.Name = &name
	return demand
}

func (demand *Demand) SetOrganization(organization *Organization) *Demand {
	demand.Organization = organization
	return demand
}

func (demand *Demand) SetOrganizationAccount(organizationAccount *AgentAccount) *Demand {
	demand.OrganizationAccount = organizationAccount
	return demand
}

func (demand *Demand) SetOverhead(overhead *Overhead) *Demand {
	demand.Overhead = overhead
	return demand
}

func (demand *Demand) SetOwner(owner *Employee) *Demand {
	demand.Owner = owner.Clean()
	return demand
}

func (demand *Demand) SetPositions(positions *Positions[DemandPosition]) *Demand {
	demand.Positions = positions
	return demand
}

func (demand *Demand) SetProject(project *Project) *Demand {
	demand.Project = NewNullValueWith(project.Clean())
	return demand
}

func (demand *Demand) SetNullProject() *Demand {
	demand.Project = NewNullValue[Project]()
	return demand
}

func (demand *Demand) SetRate(rate *Rate) *Demand {
	demand.Rate = rate
	return demand
}

func (demand *Demand) SetSalesChannel(salesChannel *SalesChannel) *Demand {
	demand.SalesChannel = NewNullValueWith(salesChannel.Clean())
	return demand
}

func (demand *Demand) SetNullSalesChannel() *Demand {
	demand.SalesChannel = NewNullValue[SalesChannel]()
	return demand
}

func (demand *Demand) SetShared(shared bool) *Demand {
	demand.Shared = &shared
	return demand
}

func (demand *Demand) SetShipmentAddress(shipmentAddress string) *Demand {
	demand.ShipmentAddress = &shipmentAddress
	return demand
}

func (demand *Demand) SetShipmentAddressFull(shipmentAddressFull *Address) *Demand {
	demand.ShipmentAddressFull = shipmentAddressFull
	return demand
}

func (demand *Demand) SetState(state *State) *Demand {
	demand.State = NewNullValueWith(state.Clean())
	return demand
}

func (demand *Demand) SetNullState() *Demand {
	demand.State = NewNullValue[State]()
	return demand
}

func (demand *Demand) SetStore(store *Store) *Demand {
	demand.Store = store.Clean()
	return demand
}

func (demand *Demand) SetSyncID(syncID uuid.UUID) *Demand {
	demand.SyncID = &syncID
	return demand
}

func (demand *Demand) SetVatEnabled(vatEnabled bool) *Demand {
	demand.VatEnabled = &vatEnabled
	return demand
}

func (demand *Demand) SetVatIncluded(vatIncluded bool) *Demand {
	demand.VatIncluded = &vatIncluded
	return demand
}

func (demand *Demand) SetCustomerOrder(customerOrder *CustomerOrder) *Demand {
	demand.CustomerOrder = customerOrder.Clean()
	return demand
}

func (demand *Demand) SetFactureOut(factureOut *FactureOut) *Demand {
	demand.FactureOut = factureOut.Clean()
	return demand
}

func (demand *Demand) SetReturns(returns Slice[SalesReturn]) *Demand {
	demand.Returns = returns
	return demand
}

func (demand *Demand) SetPayments(payments Slice[Payment]) *Demand {
	demand.Payments = payments
	return demand
}

func (demand *Demand) SetInvoicesOut(invoicesOut Slice[InvoiceOut]) *Demand {
	demand.InvoicesOut = invoicesOut
	return demand
}

func (demand *Demand) SetCargoName(cargoName string) *Demand {
	demand.CargoName = &cargoName
	return demand
}

func (demand *Demand) SetCarrier(carrier *Counterparty) *Demand {
	demand.Carrier = carrier.Clean()
	return demand
}

func (demand *Demand) SetConsignee(consignee *Counterparty) *Demand {
	demand.Consignee = consignee.Clean()
	return demand
}

func (demand *Demand) SetGoodPackQuantity(goodPackQuantity int) *Demand {
	demand.GoodPackQuantity = &goodPackQuantity
	return demand
}

func (demand *Demand) SetShippingInstructions(shippingInstructions string) *Demand {
	demand.ShippingInstructions = &shippingInstructions
	return demand
}

func (demand *Demand) SetStateContractID(stateContractID string) *Demand {
	demand.StateContractID = &stateContractID
	return demand
}

func (demand *Demand) SetTransportFacility(transportFacility string) *Demand {
	demand.TransportFacility = &transportFacility
	return demand
}

func (demand *Demand) SetTransportFacilityNumber(transportFacilityNumber string) *Demand {
	demand.TransportFacilityNumber = &transportFacilityNumber
	return demand
}

func (demand *Demand) SetAttributes(attributes Slice[AttributeValue]) *Demand {
	demand.Attributes = attributes
	return demand
}

func (demand Demand) String() string {
	return Stringify(demand)
}

func (demand Demand) MetaType() MetaType {
	return MetaTypeDemand
}

// DemandPosition Позиция Отгрузки
// Ключевое слово: demandposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-otgruzka-otgruzki-pozicii-otgruzki
type DemandPosition struct {
	Slot              *Slot               `json:"slot,omitempty"`
	Price             *float64            `json:"price,omitempty"`
	Cost              *int                `json:"cost,omitempty"`
	Discount          *float64            `json:"discount,omitempty"`
	AccountID         *uuid.UUID          `json:"accountId,omitempty"`
	Pack              *Pack               `json:"pack,omitempty"`
	Assortment        *AssortmentPosition `json:"assortment,omitempty"`
	Quantity          *float64            `json:"quantity,omitempty"`
	ID                *uuid.UUID          `json:"id,omitempty"`
	Stock             *Stock              `json:"stock,omitempty"`
	VatEnabled        *bool               `json:"vatEnabled,omitempty"`
	Vat               *int                `json:"vat,omitempty"`
	Overhead          *float64            `json:"overhead,omitempty"`
	TrackingCodes1162 Slice[TrackingCode] `json:"trackingCodes_1162,omitempty"`
	TrackingCodes     Slice[TrackingCode] `json:"trackingCodes,omitempty"`
	Things            Slice[string]       `json:"things,omitempty"`
}

func (demandPosition DemandPosition) GetAccountID() uuid.UUID {
	return Deref(demandPosition.AccountID)
}

func (demandPosition DemandPosition) GetAssortment() AssortmentPosition {
	return Deref(demandPosition.Assortment)
}

func (demandPosition DemandPosition) GetCost() int {
	return Deref(demandPosition.Cost)
}

func (demandPosition DemandPosition) GetDiscount() float64 {
	return Deref(demandPosition.Discount)
}

func (demandPosition DemandPosition) GetID() uuid.UUID {
	return Deref(demandPosition.ID)
}

func (demandPosition DemandPosition) GetPack() Pack {
	return Deref(demandPosition.Pack)
}

func (demandPosition DemandPosition) GetPrice() float64 {
	return Deref(demandPosition.Price)
}

func (demandPosition DemandPosition) GetQuantity() float64 {
	return Deref(demandPosition.Quantity)
}

func (demandPosition DemandPosition) GetSlot() Slot {
	return Deref(demandPosition.Slot)
}

func (demandPosition DemandPosition) GetThings() Slice[string] {
	return demandPosition.Things
}

func (demandPosition DemandPosition) GetTrackingCodes() Slice[TrackingCode] {
	return demandPosition.TrackingCodes
}

func (demandPosition DemandPosition) GetTrackingCodes1162() Slice[TrackingCode] {
	return demandPosition.TrackingCodes1162
}

func (demandPosition DemandPosition) GetOverhead() float64 {
	return Deref(demandPosition.Overhead)
}

func (demandPosition DemandPosition) GetVat() int {
	return Deref(demandPosition.Vat)
}

func (demandPosition DemandPosition) GetVatEnabled() bool {
	return Deref(demandPosition.VatEnabled)
}

func (demandPosition DemandPosition) GetStock() Stock {
	return Deref(demandPosition.Stock)
}

func (demandPosition *DemandPosition) SetAssortment(assortment AsAssortment) *DemandPosition {
	demandPosition.Assortment = assortment.AsAssortment()
	return demandPosition
}

func (demandPosition *DemandPosition) SetCost(cost int) *DemandPosition {
	demandPosition.Cost = &cost
	return demandPosition
}

func (demandPosition *DemandPosition) SetDiscount(discount float64) *DemandPosition {
	demandPosition.Discount = &discount
	return demandPosition
}

func (demandPosition *DemandPosition) SetPack(pack *Pack) *DemandPosition {
	demandPosition.Pack = pack
	return demandPosition
}

func (demandPosition *DemandPosition) SetPrice(price float64) *DemandPosition {
	demandPosition.Price = &price
	return demandPosition
}

func (demandPosition *DemandPosition) SetQuantity(quantity float64) *DemandPosition {
	demandPosition.Quantity = &quantity
	return demandPosition
}

func (demandPosition *DemandPosition) SetSlot(slot *Slot) *DemandPosition {
	demandPosition.Slot = slot.Clean()
	return demandPosition
}

func (demandPosition *DemandPosition) SetThings(things Slice[string]) *DemandPosition {
	demandPosition.Things = things
	return demandPosition
}

func (demandPosition *DemandPosition) SetTrackingCodes(trackingCodes Slice[TrackingCode]) *DemandPosition {
	demandPosition.TrackingCodes = trackingCodes
	return demandPosition
}

func (demandPosition *DemandPosition) SetTrackingCodes1162(trackingCodes1162 Slice[TrackingCode]) *DemandPosition {
	demandPosition.TrackingCodes1162 = trackingCodes1162
	return demandPosition
}

func (demandPosition *DemandPosition) SetVat(vat int) *DemandPosition {
	demandPosition.Vat = &vat
	return demandPosition
}

func (demandPosition *DemandPosition) SetVatEnabled(vatEnabled bool) *DemandPosition {
	demandPosition.VatEnabled = &vatEnabled
	return demandPosition
}

func (demandPosition DemandPosition) String() string {
	return Stringify(demandPosition)
}

func (demandPosition DemandPosition) MetaType() MetaType {
	return MetaTypeDemandPosition
}

// DemandService
// Сервис для работы с отгрузками.
type DemandService interface {
	GetList(ctx context.Context, params ...*Params) (*List[Demand], *resty.Response, error)
	Create(ctx context.Context, demand *Demand, params ...*Params) (*Demand, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, demandList Slice[Demand], params ...*Params) (*Slice[Demand], *resty.Response, error)
	DeleteMany(ctx context.Context, demandList []MetaWrapper) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Demand, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, demand *Demand, params ...*Params) (*Demand, *resty.Response, error)
	Template(ctx context.Context) (*Demand, *resty.Response, error)
	TemplateBased(ctx context.Context, basedOn ...MetaOwner) (*Demand, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetaAttributesSharedStatesWrapper, *resty.Response, error)
	GetPositions(ctx context.Context, id uuid.UUID, params ...*Params) (*MetaArray[DemandPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params ...*Params) (*DemandPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *DemandPosition, params ...*Params) (*DemandPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id uuid.UUID, position *DemandPosition) (*DemandPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id uuid.UUID, positions Slice[DemandPosition]) (*Slice[DemandPosition], *resty.Response, error)
	DeletePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (bool, *resty.Response, error)
	GetPositionTrackingCodes(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (*MetaArray[TrackingCode], *resty.Response, error)
	CreateOrUpdatePositionTrackingCodes(ctx context.Context, id uuid.UUID, positionID uuid.UUID, trackingCodes Slice[TrackingCode]) (*Slice[TrackingCode], *resty.Response, error)
	DeletePositionTrackingCodes(ctx context.Context, id uuid.UUID, positionID uuid.UUID, trackingCodes Slice[TrackingCode]) (*DeleteManyResponse, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList Slice[Attribute]) (*Slice[Attribute], *resty.Response, error)
	UpdateAttribute(ctx context.Context, id uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList []MetaWrapper) (*DeleteManyResponse, *resty.Response, error)
	GetPublications(ctx context.Context, id uuid.UUID) (*MetaArray[Publication], *resty.Response, error)
	GetPublicationByID(ctx context.Context, id uuid.UUID, publicationID uuid.UUID) (*Publication, *resty.Response, error)
	Publish(ctx context.Context, id uuid.UUID, template Templater) (*Publication, *resty.Response, error)
	DeletePublication(ctx context.Context, id uuid.UUID, publicationID uuid.UUID) (bool, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*Demand, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params ...*Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)
	MoveToTrash(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetEmbeddedTemplates(ctx context.Context) (*List[EmbeddedTemplate], *resty.Response, error)
	GetEmbeddedTemplateByID(ctx context.Context, id uuid.UUID) (*EmbeddedTemplate, *resty.Response, error)
	GetCustomTemplates(ctx context.Context) (*List[CustomTemplate], *resty.Response, error)
	GetCustomTemplateByID(ctx context.Context, id uuid.UUID) (*CustomTemplate, *resty.Response, error)
	PrintDocument(ctx context.Context, id uuid.UUID, PrintDocumentArg *PrintDocumentArg) (*PrintFile, *resty.Response, error)
	GetStateByID(ctx context.Context, id uuid.UUID) (*State, *resty.Response, error)
	CreateState(ctx context.Context, state *State) (*State, *resty.Response, error)
	UpdateState(ctx context.Context, id uuid.UUID, state *State) (*State, *resty.Response, error)
	CreateOrUpdateStates(ctx context.Context, states Slice[State]) (*Slice[State], *resty.Response, error)
	DeleteState(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetFiles(ctx context.Context, id uuid.UUID) (*MetaArray[File], *resty.Response, error)
	CreateFile(ctx context.Context, id uuid.UUID, file *File) (*Slice[File], *resty.Response, error)
	UpdateFiles(ctx context.Context, id uuid.UUID, files Slice[File]) (*Slice[File], *resty.Response, error)
	DeleteFile(ctx context.Context, id uuid.UUID, fileID uuid.UUID) (bool, *resty.Response, error)
	DeleteFiles(ctx context.Context, id uuid.UUID, files []MetaWrapper) (*DeleteManyResponse, *resty.Response, error)
}

func NewDemandService(client *Client) DemandService {
	e := NewEndpoint(client, "entity/demand")
	return newMainService[Demand, DemandPosition, MetaAttributesSharedStatesWrapper, any](e)
}
