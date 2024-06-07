package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// CustomerOrder Заказ покупателя.
// Ключевое слово: customerorder
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-zakaz-pokupatelq
type CustomerOrder struct {
	OrganizationAccount   *AgentAccount                     `json:"organizationAccount,omitempty"`
	Project               *Project                          `json:"project,omitempty"`
	AgentAccount          *AgentAccount                     `json:"agentAccount,omitempty"`
	Applicable            *bool                             `json:"applicable,omitempty"`
	Moves                 Moves                             `json:"moves,omitempty"`
	Code                  *string                           `json:"code,omitempty"`
	Agent                 *Counterparty                     `json:"agent,omitempty"`
	Created               *Timestamp                        `json:"created,omitempty"`
	Deleted               *Timestamp                        `json:"deleted,omitempty"`
	DeliveryPlannedMoment *Timestamp                        `json:"deliveryPlannedMoment,omitempty"`
	Description           *string                           `json:"description,omitempty"`
	ExternalCode          *string                           `json:"externalCode,omitempty"`
	Files                 *Files                            `json:"files,omitempty"`
	Group                 *Group                            `json:"group,omitempty"`
	ID                    *uuid.UUID                        `json:"id,omitempty"`
	InvoicedSum           *float64                          `json:"invoicedSum,omitempty"`
	Meta                  *Meta                             `json:"meta,omitempty"`
	Name                  *string                           `json:"name,omitempty"`
	Moment                *Timestamp                        `json:"moment,omitempty"`
	Organization          *Organization                     `json:"organization,omitempty"`
	Printed               *bool                             `json:"printed,omitempty"`
	Owner                 *Employee                         `json:"owner,omitempty"`
	PayedSum              *float64                          `json:"payedSum,omitempty"`
	Positions             *Positions[CustomerOrderPosition] `json:"positions,omitempty"`
	AccountID             *uuid.UUID                        `json:"accountId,omitempty"`
	Contract              *Contract                         `json:"contract,omitempty"`
	Published             *bool                             `json:"published,omitempty"`
	Rate                  *Rate                             `json:"rate,omitempty"`
	ReservedSum           *float64                          `json:"reservedSum,omitempty"`
	SalesChannel          *SalesChannel                     `json:"salesChannel,omitempty"`
	Shared                *bool                             `json:"shared,omitempty"`
	ShipmentAddress       *string                           `json:"shipmentAddress,omitempty"`
	ShipmentAddressFull   *Address                          `json:"shipmentAddressFull,omitempty"`
	ShippedSum            *float64                          `json:"shippedSum,omitempty"`
	State                 *State                            `json:"state,omitempty"`
	Store                 *Store                            `json:"store,omitempty"`
	Sum                   *float64                          `json:"sum,omitempty"`
	SyncID                *uuid.UUID                        `json:"syncId,omitempty"`
	Updated               *Timestamp                        `json:"updated,omitempty"`
	VatEnabled            *bool                             `json:"vatEnabled,omitempty"`
	VatIncluded           *bool                             `json:"vatIncluded,omitempty"`
	VatSum                *float64                          `json:"vatSum,omitempty"`
	Prepayments           Prepayments                       `json:"prepayments,omitempty"`
	PurchaseOrders        PurchaseOrders                    `json:"purchaseOrders,omitempty"`
	Demands               Demands                           `json:"demands,omitempty"`
	Payments              Payments                          `json:"payments,omitempty"`
	InvoicesOut           InvoicesOut                       `json:"invoicesOut,omitempty"`
	TaxSystem             TaxSystem                         `json:"taxSystem,omitempty"`
	Attributes            Attributes                        `json:"attributes,omitempty"`
}

func (customerOrder CustomerOrder) GetOrganizationAccount() AgentAccount {
	return Deref(customerOrder.OrganizationAccount)
}

func (customerOrder CustomerOrder) GetProject() Project {
	return Deref(customerOrder.Project)
}

func (customerOrder CustomerOrder) GetAgentAccount() AgentAccount {
	return Deref(customerOrder.AgentAccount)
}

func (customerOrder CustomerOrder) GetApplicable() bool {
	return Deref(customerOrder.Applicable)
}

func (customerOrder CustomerOrder) GetMoves() Moves {
	return customerOrder.Moves
}

func (customerOrder CustomerOrder) GetCode() string {
	return Deref(customerOrder.Code)
}

func (customerOrder CustomerOrder) GetAgent() Counterparty {
	return Deref(customerOrder.Agent)
}

func (customerOrder CustomerOrder) GetCreated() Timestamp {
	return Deref(customerOrder.Created)
}

func (customerOrder CustomerOrder) GetDeleted() Timestamp {
	return Deref(customerOrder.Deleted)
}

func (customerOrder CustomerOrder) GetDeliveryPlannedMoment() Timestamp {
	return Deref(customerOrder.DeliveryPlannedMoment)
}

func (customerOrder CustomerOrder) GetDescription() string {
	return Deref(customerOrder.Description)
}

func (customerOrder CustomerOrder) GetExternalCode() string {
	return Deref(customerOrder.ExternalCode)
}

func (customerOrder CustomerOrder) GetFiles() Files {
	return Deref(customerOrder.Files)
}

func (customerOrder CustomerOrder) GetGroup() Group {
	return Deref(customerOrder.Group)
}

func (customerOrder CustomerOrder) GetID() uuid.UUID {
	return Deref(customerOrder.ID)
}

func (customerOrder CustomerOrder) GetInvoicedSum() float64 {
	return Deref(customerOrder.InvoicedSum)
}

func (customerOrder CustomerOrder) GetMeta() Meta {
	return Deref(customerOrder.Meta)
}

func (customerOrder CustomerOrder) GetName() string {
	return Deref(customerOrder.Name)
}

func (customerOrder CustomerOrder) GetMoment() Timestamp {
	return Deref(customerOrder.Moment)
}

func (customerOrder CustomerOrder) GetOrganization() Organization {
	return Deref(customerOrder.Organization)
}

func (customerOrder CustomerOrder) GetPrinted() bool {
	return Deref(customerOrder.Printed)
}

func (customerOrder CustomerOrder) GetOwner() Employee {
	return Deref(customerOrder.Owner)
}

func (customerOrder CustomerOrder) GetPayedSum() float64 {
	return Deref(customerOrder.PayedSum)
}

func (customerOrder CustomerOrder) GetPositions() Positions[CustomerOrderPosition] {
	return Deref(customerOrder.Positions)
}

func (customerOrder CustomerOrder) GetAccountID() uuid.UUID {
	return Deref(customerOrder.AccountID)
}

func (customerOrder CustomerOrder) GetContract() Contract {
	return Deref(customerOrder.Contract)
}

func (customerOrder CustomerOrder) GetPublished() bool {
	return Deref(customerOrder.Published)
}

func (customerOrder CustomerOrder) GetRate() Rate {
	return Deref(customerOrder.Rate)
}

func (customerOrder CustomerOrder) GetReservedSum() float64 {
	return Deref(customerOrder.ReservedSum)
}

func (customerOrder CustomerOrder) GetSalesChannel() SalesChannel {
	return Deref(customerOrder.SalesChannel)
}

func (customerOrder CustomerOrder) GetShared() bool {
	return Deref(customerOrder.Shared)
}

func (customerOrder CustomerOrder) GetShipmentAddress() string {
	return Deref(customerOrder.ShipmentAddress)
}

func (customerOrder CustomerOrder) GetShipmentAddressFull() Address {
	return Deref(customerOrder.ShipmentAddressFull)
}

func (customerOrder CustomerOrder) GetShippedSum() float64 {
	return Deref(customerOrder.ShippedSum)
}

func (customerOrder CustomerOrder) GetState() State {
	return Deref(customerOrder.State)
}

func (customerOrder CustomerOrder) GetStore() Store {
	return Deref(customerOrder.Store)
}

func (customerOrder CustomerOrder) GetSum() float64 {
	return Deref(customerOrder.Sum)
}

func (customerOrder CustomerOrder) GetSyncID() uuid.UUID {
	return Deref(customerOrder.SyncID)
}

func (customerOrder CustomerOrder) GetPrepayments() Prepayments {
	return customerOrder.Prepayments
}

func (customerOrder CustomerOrder) GetUpdated() Timestamp {
	return Deref(customerOrder.Updated)
}

func (customerOrder CustomerOrder) GetVatEnabled() bool {
	return Deref(customerOrder.VatEnabled)
}

func (customerOrder CustomerOrder) GetVatIncluded() bool {
	return Deref(customerOrder.VatIncluded)
}

func (customerOrder CustomerOrder) GetVatSum() float64 {
	return Deref(customerOrder.VatSum)
}

func (customerOrder CustomerOrder) GetPurchaseOrders() PurchaseOrders {
	return customerOrder.PurchaseOrders
}

func (customerOrder CustomerOrder) GetDemands() Demands {
	return customerOrder.Demands
}

func (customerOrder CustomerOrder) GetPayments() Payments {
	return customerOrder.Payments
}

func (customerOrder CustomerOrder) GetInvoicesOut() InvoicesOut {
	return customerOrder.InvoicesOut
}

func (customerOrder CustomerOrder) GetTaxSystem() TaxSystem {
	return customerOrder.TaxSystem
}

func (customerOrder CustomerOrder) GetAttributes() Attributes {
	return customerOrder.Attributes
}

func (customerOrder *CustomerOrder) SetOrganizationAccount(organizationAccount *AgentAccount) *CustomerOrder {
	customerOrder.OrganizationAccount = organizationAccount
	return customerOrder
}

func (customerOrder *CustomerOrder) SetProject(project *Project) *CustomerOrder {
	customerOrder.Project = project
	return customerOrder
}

func (customerOrder *CustomerOrder) SetAgentAccount(agentAccount *AgentAccount) *CustomerOrder {
	customerOrder.AgentAccount = agentAccount
	return customerOrder
}

func (customerOrder *CustomerOrder) SetApplicable(applicable bool) *CustomerOrder {
	customerOrder.Applicable = &applicable
	return customerOrder
}

func (customerOrder *CustomerOrder) SetMoves(moves Moves) *CustomerOrder {
	customerOrder.Moves = moves
	return customerOrder
}

func (customerOrder *CustomerOrder) SetCode(code string) *CustomerOrder {
	customerOrder.Code = &code
	return customerOrder
}

func (customerOrder *CustomerOrder) SetAgent(agent *Counterparty) *CustomerOrder {
	customerOrder.Agent = agent
	return customerOrder
}

func (customerOrder *CustomerOrder) SetDeliveryPlannedMoment(deliveryPlannedMoment *Timestamp) *CustomerOrder {
	customerOrder.DeliveryPlannedMoment = deliveryPlannedMoment
	return customerOrder
}

func (customerOrder *CustomerOrder) SetDescription(description string) *CustomerOrder {
	customerOrder.Description = &description
	return customerOrder
}

func (customerOrder *CustomerOrder) SetExternalCode(externalCode string) *CustomerOrder {
	customerOrder.ExternalCode = &externalCode
	return customerOrder
}

func (customerOrder *CustomerOrder) SetFiles(files *Files) *CustomerOrder {
	customerOrder.Files = files
	return customerOrder
}

func (customerOrder *CustomerOrder) SetGroup(group *Group) *CustomerOrder {
	customerOrder.Group = group
	return customerOrder
}

func (customerOrder *CustomerOrder) SetMeta(meta *Meta) *CustomerOrder {
	customerOrder.Meta = meta
	return customerOrder
}

func (customerOrder *CustomerOrder) SetName(name string) *CustomerOrder {
	customerOrder.Name = &name
	return customerOrder
}

func (customerOrder *CustomerOrder) SetMoment(moment *Timestamp) *CustomerOrder {
	customerOrder.Moment = moment
	return customerOrder
}

func (customerOrder *CustomerOrder) SetOrganization(organization *Organization) *CustomerOrder {
	customerOrder.Organization = organization
	return customerOrder
}

func (customerOrder *CustomerOrder) SetOwner(owner *Employee) *CustomerOrder {
	customerOrder.Owner = owner
	return customerOrder
}

func (customerOrder *CustomerOrder) SetPositions(positions *Positions[CustomerOrderPosition]) *CustomerOrder {
	customerOrder.Positions = positions
	return customerOrder
}

func (customerOrder *CustomerOrder) SetContract(contract *Contract) *CustomerOrder {
	customerOrder.Contract = contract
	return customerOrder
}

func (customerOrder *CustomerOrder) SetRate(rate *Rate) *CustomerOrder {
	customerOrder.Rate = rate
	return customerOrder
}

func (customerOrder *CustomerOrder) SetSalesChannel(salesChannel *SalesChannel) *CustomerOrder {
	customerOrder.SalesChannel = salesChannel
	return customerOrder
}

func (customerOrder *CustomerOrder) SetShared(shared bool) *CustomerOrder {
	customerOrder.Shared = &shared
	return customerOrder
}

func (customerOrder *CustomerOrder) SetShipmentAddress(shipmentAddress string) *CustomerOrder {
	customerOrder.ShipmentAddress = &shipmentAddress
	return customerOrder
}

func (customerOrder *CustomerOrder) SetShipmentAddressFull(shipmentAddressFull Address) *CustomerOrder {
	customerOrder.ShipmentAddressFull = &shipmentAddressFull
	return customerOrder
}

func (customerOrder *CustomerOrder) SetState(state *State) *CustomerOrder {
	customerOrder.State = state
	return customerOrder
}

func (customerOrder *CustomerOrder) SetStore(store *Store) *CustomerOrder {
	customerOrder.Store = store
	return customerOrder
}

func (customerOrder *CustomerOrder) SetSyncID(syncID *uuid.UUID) *CustomerOrder {
	customerOrder.SyncID = syncID
	return customerOrder
}

func (customerOrder *CustomerOrder) SetPrepayments(prepayments Prepayments) *CustomerOrder {
	customerOrder.Prepayments = prepayments
	return customerOrder
}

func (customerOrder *CustomerOrder) SetVatEnabled(vatEnabled bool) *CustomerOrder {
	customerOrder.VatEnabled = &vatEnabled
	return customerOrder
}

func (customerOrder *CustomerOrder) SetVatIncluded(vatIncluded bool) *CustomerOrder {
	customerOrder.VatIncluded = &vatIncluded
	return customerOrder
}

func (customerOrder *CustomerOrder) SetPurchaseOrders(purchaseOrders PurchaseOrders) *CustomerOrder {
	customerOrder.PurchaseOrders = purchaseOrders
	return customerOrder
}

func (customerOrder *CustomerOrder) SetDemands(demands Demands) *CustomerOrder {
	customerOrder.Demands = demands
	return customerOrder
}

func (customerOrder *CustomerOrder) SetPayments(payments Payments) *CustomerOrder {
	customerOrder.Payments = payments
	return customerOrder
}

func (customerOrder *CustomerOrder) SetInvoicesOut(invoicesOut InvoicesOut) *CustomerOrder {
	customerOrder.InvoicesOut = invoicesOut
	return customerOrder
}

func (customerOrder *CustomerOrder) SetTaxSystem(taxSystem TaxSystem) *CustomerOrder {
	customerOrder.TaxSystem = taxSystem
	return customerOrder
}

func (customerOrder *CustomerOrder) SetAttributes(attributes Attributes) *CustomerOrder {
	customerOrder.Attributes = attributes
	return customerOrder
}

func (customerOrder CustomerOrder) String() string {
	return Stringify(customerOrder)
}

func (customerOrder CustomerOrder) MetaType() MetaType {
	return MetaTypeCustomerOrder
}

type CustomerOrders = Slice[CustomerOrder]

// CustomerOrderPosition Позиция Заказа покупателя.
// Ключевое слово: customerorderposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-zakaz-pokupatelq-zakazy-pokupatelej-pozicii-zakaza-pokupatelq
type CustomerOrderPosition struct {
	Quantity   *float64            `json:"quantity,omitempty"`
	Assortment *AssortmentPosition `json:"assortment,omitempty"`
	Discount   *float64            `json:"discount,omitempty"`
	ID         *uuid.UUID          `json:"id,omitempty"`
	Pack       *Pack               `json:"pack,omitempty"`
	Price      *float64            `json:"price,omitempty"`
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`
	Reserve    *float64            `json:"reserve,omitempty"`
	Shipped    *float64            `json:"shipped,omitempty"`
	Vat        *int                `json:"vat,omitempty"`
	VatEnabled *bool               `json:"vatEnabled,omitempty"`
	Stock      *Stock              `json:"stock,omitempty"`
	TaxSystem  GoodTaxSystem       `json:"taxSystem,omitempty"`
}

func (customerOrderPosition CustomerOrderPosition) GetQuantity() float64 {
	return Deref(customerOrderPosition.Quantity)
}

func (customerOrderPosition CustomerOrderPosition) GetAssortment() AssortmentPosition {
	return Deref(customerOrderPosition.Assortment)
}

func (customerOrderPosition CustomerOrderPosition) GetDiscount() float64 {
	return Deref(customerOrderPosition.Discount)
}

func (customerOrderPosition CustomerOrderPosition) GetID() uuid.UUID {
	return Deref(customerOrderPosition.ID)
}

func (customerOrderPosition CustomerOrderPosition) GetPack() Pack {
	return Deref(customerOrderPosition.Pack)
}

func (customerOrderPosition CustomerOrderPosition) GetPrice() float64 {
	return Deref(customerOrderPosition.Price)
}

func (customerOrderPosition CustomerOrderPosition) GetAccountID() uuid.UUID {
	return Deref(customerOrderPosition.AccountID)
}

func (customerOrderPosition CustomerOrderPosition) GetReserve() float64 {
	return Deref(customerOrderPosition.Reserve)
}

func (customerOrderPosition CustomerOrderPosition) GetShipped() float64 {
	return Deref(customerOrderPosition.Shipped)
}

func (customerOrderPosition CustomerOrderPosition) GetVat() int {
	return Deref(customerOrderPosition.Vat)
}

func (customerOrderPosition CustomerOrderPosition) GetVatEnabled() bool {
	return Deref(customerOrderPosition.VatEnabled)
}

func (customerOrderPosition CustomerOrderPosition) GetStock() Stock {
	return Deref(customerOrderPosition.Stock)
}

func (customerOrderPosition CustomerOrderPosition) GetTaxSystem() GoodTaxSystem {
	return customerOrderPosition.TaxSystem
}

func (customerOrderPosition *CustomerOrderPosition) SetQuantity(quantity float64) *CustomerOrderPosition {
	customerOrderPosition.Quantity = &quantity
	return customerOrderPosition
}

func (customerOrderPosition *CustomerOrderPosition) SetAssortment(assortment *AssortmentPosition) *CustomerOrderPosition {
	customerOrderPosition.Assortment = assortment
	return customerOrderPosition
}

func (customerOrderPosition *CustomerOrderPosition) SetDiscount(discount float64) *CustomerOrderPosition {
	customerOrderPosition.Discount = &discount
	return customerOrderPosition
}

func (customerOrderPosition *CustomerOrderPosition) SetPack(pack *Pack) *CustomerOrderPosition {
	customerOrderPosition.Pack = pack
	return customerOrderPosition
}

func (customerOrderPosition *CustomerOrderPosition) SetPrice(price float64) *CustomerOrderPosition {
	customerOrderPosition.Price = &price
	return customerOrderPosition
}

func (customerOrderPosition *CustomerOrderPosition) SetReserve(reserve float64) *CustomerOrderPosition {
	customerOrderPosition.Reserve = &reserve
	return customerOrderPosition
}

func (customerOrderPosition *CustomerOrderPosition) SetVat(vat int) *CustomerOrderPosition {
	customerOrderPosition.Vat = &vat
	return customerOrderPosition
}

func (customerOrderPosition *CustomerOrderPosition) SetVatEnabled(vatEnabled bool) *CustomerOrderPosition {
	customerOrderPosition.VatEnabled = &vatEnabled
	return customerOrderPosition
}

func (customerOrderPosition *CustomerOrderPosition) SetTaxSystem(taxSystem GoodTaxSystem) *CustomerOrderPosition {
	customerOrderPosition.TaxSystem = taxSystem
	return customerOrderPosition
}

func (customerOrderPosition CustomerOrderPosition) String() string {
	return Stringify(customerOrderPosition)
}

func (customerOrderPosition CustomerOrderPosition) MetaType() MetaType {
	return MetaTypeCustomerOrderPosition
}

// CustomerOrderService
// Сервис для работы с заказами покупателя.
type CustomerOrderService interface {
	GetList(ctx context.Context, params *Params) (*List[CustomerOrder], *resty.Response, error)
	Create(ctx context.Context, customerOrder *CustomerOrder, params *Params) (*CustomerOrder, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, customerOrderList []*CustomerOrder, params *Params) (*[]CustomerOrder, *resty.Response, error)
	DeleteMany(ctx context.Context, customerOrderList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*CustomerOrder, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, customerOrder *CustomerOrder, params *Params) (*CustomerOrder, *resty.Response, error)
	//endpointTemplate[CustomerOrder]
	GetMetadata(ctx context.Context) (*MetadataAttributeSharedStates, *resty.Response, error)
	GetPositions(ctx context.Context, id *uuid.UUID, params *Params) (*MetaArray[CustomerOrderPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, params *Params) (*CustomerOrderPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id *uuid.UUID, positionID *uuid.UUID, position *CustomerOrderPosition, params *Params) (*CustomerOrderPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id *uuid.UUID, position *CustomerOrderPosition) (*CustomerOrderPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id *uuid.UUID, positions []*CustomerOrderPosition) (*[]CustomerOrderPosition, *resty.Response, error)
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
	GetPublications(ctx context.Context, id *uuid.UUID) (*MetaArray[Publication], *resty.Response, error)
	GetPublicationByID(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (*Publication, *resty.Response, error)
	Publish(ctx context.Context, id *uuid.UUID, template *Templater) (*Publication, *resty.Response, error)
	DeletePublication(ctx context.Context, id *uuid.UUID, publicationID *uuid.UUID) (bool, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID *uuid.UUID) (*CustomerOrder, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID *uuid.UUID) (bool, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params *Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id *uuid.UUID) (*NamedFilter, *resty.Response, error)
	MoveToTrash(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetEmbeddedTemplates(ctx context.Context) (*List[EmbeddedTemplate], *resty.Response, error)
	GetEmbeddedTemplateByID(ctx context.Context, id *uuid.UUID) (*EmbeddedTemplate, *resty.Response, error)
	GetCustomTemplates(ctx context.Context) (*List[CustomTemplate], *resty.Response, error)
	GetCustomTemplateByID(ctx context.Context, id *uuid.UUID) (*CustomTemplate, *resty.Response, error)
}

func NewCustomerOrderService(client *Client) CustomerOrderService {
	e := NewEndpoint(client, "entity/customerorder")
	return newMainService[CustomerOrder, CustomerOrderPosition, MetadataAttributeSharedStates, any](e)
}
