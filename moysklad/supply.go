package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Supply Приёмка.
// Ключевое слово: supply
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-priemka
type Supply struct {
	OrganizationAccount *AgentAccount              `json:"organizationAccount,omitempty"`
	Applicable          *bool                      `json:"applicable,omitempty"`
	AgentAccount        *AgentAccount              `json:"agentAccount,omitempty"`
	Overhead            *Overhead                  `json:"overhead,omitempty"`
	Returns             Slice[PurchaseReturn]      `json:"returns,omitempty"`
	Code                *string                    `json:"code,omitempty"`
	Contract            *NullValue[Contract]       `json:"contract,omitempty"`
	Created             *Timestamp                 `json:"created,omitempty"`
	Deleted             *Timestamp                 `json:"deleted,omitempty"`
	Description         *string                    `json:"description,omitempty"`
	ExternalCode        *string                    `json:"externalCode,omitempty"`
	Files               *MetaArray[File]           `json:"files,omitempty"`
	Group               *Group                     `json:"group,omitempty"`
	ID                  *uuid.UUID                 `json:"id,omitempty"`
	IncomingDate        *Timestamp                 `json:"incomingDate,omitempty"`
	Owner               *Employee                  `json:"owner,omitempty"`
	Meta                *Meta                      `json:"meta,omitempty"`
	Moment              *Timestamp                 `json:"moment,omitempty"`
	Name                *string                    `json:"name,omitempty"`
	Organization        *Organization              `json:"organization,omitempty"`
	Payments            Slice[Payment]             `json:"payments,omitempty"`
	Agent               *Counterparty              `json:"agent,omitempty"`
	IncomingNumber      *string                    `json:"incomingNumber,omitempty"`
	PayedSum            *float64                   `json:"payedSum,omitempty"`
	Positions           *Positions[SupplyPosition] `json:"positions,omitempty"`
	Printed             *bool                      `json:"printed,omitempty"`
	Project             *NullValue[Project]        `json:"project,omitempty"`
	Published           *bool                      `json:"published,omitempty"`
	Rate                *NullValue[Rate]           `json:"rate,omitempty"`
	Shared              *bool                      `json:"shared,omitempty"`
	State               *State                     `json:"state,omitempty"`
	Store               *Store                     `json:"store,omitempty"`
	Sum                 *float64                   `json:"sum,omitempty"`
	SyncID              *uuid.UUID                 `json:"syncId,omitempty"`
	Updated             *Timestamp                 `json:"updated,omitempty"`
	VatEnabled          *bool                      `json:"vatEnabled,omitempty"`
	VatIncluded         *bool                      `json:"vatIncluded,omitempty"`
	VatSum              *float64                   `json:"vatSum,omitempty"`
	PurchaseOrder       *PurchaseOrder             `json:"purchaseOrder,omitempty"`
	FactureIn           *FactureIn                 `json:"factureIn,omitempty"`
	InvoicesIn          Slice[InvoiceIn]           `json:"invoicesIn,omitempty"`
	AccountID           *uuid.UUID                 `json:"accountId,omitempty"`
	Attributes          Slice[Attribute]           `json:"attributes,omitempty"`
}

// Clean возвращает сущность с единственным заполненным полем Meta
func (supply Supply) Clean() *Supply {
	return &Supply{Meta: supply.Meta}
}

// AsOperation возвращает объект Operation c полем Meta сущности
func (supply Supply) AsOperation() *Operation {
	return &Operation{Meta: supply.GetMeta()}
}

// AsTaskOperation реализует интерфейс AsTaskOperationInterface
func (supply Supply) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: supply.Meta}
}

func (supply Supply) GetOrganizationAccount() AgentAccount {
	return Deref(supply.OrganizationAccount)
}

func (supply Supply) GetApplicable() bool {
	return Deref(supply.Applicable)
}

func (supply Supply) GetAgentAccount() AgentAccount {
	return Deref(supply.AgentAccount)
}

func (supply Supply) GetOverhead() Overhead {
	return Deref(supply.Overhead)
}

func (supply Supply) GetReturns() Slice[PurchaseReturn] {
	return supply.Returns
}

func (supply Supply) GetCode() string {
	return Deref(supply.Code)
}

func (supply Supply) GetContract() Contract {
	return supply.Contract.Get()
}

func (supply Supply) GetCreated() Timestamp {
	return Deref(supply.Created)
}

func (supply Supply) GetDeleted() Timestamp {
	return Deref(supply.Deleted)
}

func (supply Supply) GetDescription() string {
	return Deref(supply.Description)
}

func (supply Supply) GetExternalCode() string {
	return Deref(supply.ExternalCode)
}

func (supply Supply) GetFiles() MetaArray[File] {
	return Deref(supply.Files)
}

func (supply Supply) GetGroup() Group {
	return Deref(supply.Group)
}

func (supply Supply) GetID() uuid.UUID {
	return Deref(supply.ID)
}

func (supply Supply) GetIncomingDate() Timestamp {
	return Deref(supply.IncomingDate)
}

func (supply Supply) GetOwner() Employee {
	return Deref(supply.Owner)
}

func (supply Supply) GetMeta() Meta {
	return Deref(supply.Meta)
}

func (supply Supply) GetMoment() Timestamp {
	return Deref(supply.Moment)
}

func (supply Supply) GetName() string {
	return Deref(supply.Name)
}

func (supply Supply) GetOrganization() Organization {
	return Deref(supply.Organization)
}

func (supply Supply) GetPayments() Slice[Payment] {
	return supply.Payments
}

func (supply Supply) GetAgent() Counterparty {
	return Deref(supply.Agent)
}

func (supply Supply) GetIncomingNumber() string {
	return Deref(supply.IncomingNumber)
}

func (supply Supply) GetPayedSum() float64 {
	return Deref(supply.PayedSum)
}

func (supply Supply) GetPositions() Positions[SupplyPosition] {
	return Deref(supply.Positions)
}

func (supply Supply) GetPrinted() bool {
	return Deref(supply.Printed)
}

func (supply Supply) GetProject() Project {
	return supply.Project.Get()
}

func (supply Supply) GetPublished() bool {
	return Deref(supply.Published)
}

func (supply Supply) GetRate() Rate {
	return supply.Rate.Get()
}

func (supply Supply) GetShared() bool {
	return Deref(supply.Shared)
}

func (supply Supply) GetState() State {
	return Deref(supply.State)
}

func (supply Supply) GetStore() Store {
	return Deref(supply.Store)
}

func (supply Supply) GetSum() float64 {
	return Deref(supply.Sum)
}

func (supply Supply) GetSyncID() uuid.UUID {
	return Deref(supply.SyncID)
}

func (supply Supply) GetUpdated() Timestamp {
	return Deref(supply.Updated)
}

func (supply Supply) GetVatEnabled() bool {
	return Deref(supply.VatEnabled)
}

func (supply Supply) GetVatIncluded() bool {
	return Deref(supply.VatIncluded)
}

func (supply Supply) GetVatSum() float64 {
	return Deref(supply.VatSum)
}

func (supply Supply) GetPurchaseOrder() PurchaseOrder {
	return Deref(supply.PurchaseOrder)
}

func (supply Supply) GetFactureIn() FactureIn {
	return Deref(supply.FactureIn)
}

func (supply Supply) GetInvoicesIn() Slice[InvoiceIn] {
	return supply.InvoicesIn
}

func (supply Supply) GetAccountID() uuid.UUID {
	return Deref(supply.AccountID)
}

func (supply Supply) GetAttributes() Slice[Attribute] {
	return supply.Attributes
}

func (supply *Supply) SetOrganizationAccount(organizationAccount *AgentAccount) *Supply {
	supply.OrganizationAccount = organizationAccount.Clean()
	return supply
}

func (supply *Supply) SetApplicable(applicable bool) *Supply {
	supply.Applicable = &applicable
	return supply
}

func (supply *Supply) SetAgentAccount(agentAccount *AgentAccount) *Supply {
	supply.AgentAccount = agentAccount.Clean()
	return supply
}

func (supply *Supply) SetOverhead(overhead *Overhead) *Supply {
	supply.Overhead = overhead
	return supply
}

func (supply *Supply) SetReturns(returns ...*PurchaseReturn) *Supply {
	supply.Returns = returns
	return supply
}

func (supply *Supply) SetCode(code string) *Supply {
	supply.Code = &code
	return supply
}

func (supply *Supply) SetContract(contract *Contract) *Supply {
	supply.Contract = NewNullValueFrom(contract.Clean())
	return supply
}

func (supply *Supply) SetDescription(description string) *Supply {
	supply.Description = &description
	return supply
}

func (supply *Supply) SetExternalCode(externalCode string) *Supply {
	supply.ExternalCode = &externalCode
	return supply
}

func (supply *Supply) SetFiles(files ...*File) *Supply {
	supply.Files = NewMetaArrayFrom(files)
	return supply
}

func (supply *Supply) SetGroup(group *Group) *Supply {
	supply.Group = group.Clean()
	return supply
}

func (supply *Supply) SetIncomingDate(incomingDate *Timestamp) *Supply {
	supply.IncomingDate = incomingDate
	return supply
}

func (supply *Supply) SetOwner(owner *Employee) *Supply {
	supply.Owner = owner.Clean()
	return supply
}

func (supply *Supply) SetMeta(meta *Meta) *Supply {
	supply.Meta = meta
	return supply
}

func (supply *Supply) SetMoment(moment *Timestamp) *Supply {
	supply.Moment = moment
	return supply
}

func (supply *Supply) SetName(name string) *Supply {
	supply.Name = &name
	return supply
}

func (supply *Supply) SetOrganization(organization *Organization) *Supply {
	supply.Organization = organization.Clean()
	return supply
}

func (supply *Supply) SetPayments(payments ...*Payment) *Supply {
	supply.Payments = payments
	return supply
}

func (supply *Supply) SetAgent(agent *Counterparty) *Supply {
	supply.Agent = agent.Clean()
	return supply
}

func (supply *Supply) SetIncomingNumber(incomingNumber string) *Supply {
	supply.IncomingNumber = &incomingNumber
	return supply
}

func (supply *Supply) SetPositions(positions ...*SupplyPosition) *Supply {
	supply.Positions = NewPositionsFrom(positions)
	return supply
}

func (supply *Supply) SetProject(project *Project) *Supply {
	supply.Project = NewNullValueFrom(project.Clean())
	return supply
}

func (supply *Supply) SetNullProject() *Supply {
	supply.Project = NewNullValue[Project]()
	return supply
}

func (supply *Supply) SetRate(rate *Rate) *Supply {
	supply.Rate = NewNullValueFrom(rate)
	return supply
}

func (supply *Supply) SetNullRate() *Supply {
	supply.Rate = NewNullValue[Rate]()
	return supply
}

func (supply *Supply) SetShared(shared bool) *Supply {
	supply.Shared = &shared
	return supply
}

func (supply *Supply) SetState(state *State) *Supply {
	supply.State = state.Clean()
	return supply
}

func (supply *Supply) SetStore(store *Store) *Supply {
	supply.Store = store.Clean()
	return supply
}

func (supply *Supply) SetSyncID(syncID uuid.UUID) *Supply {
	supply.SyncID = &syncID
	return supply
}

func (supply *Supply) SetVatEnabled(vatEnabled bool) *Supply {
	supply.VatEnabled = &vatEnabled
	return supply
}

func (supply *Supply) SetVatIncluded(vatIncluded bool) *Supply {
	supply.VatIncluded = &vatIncluded
	return supply
}

func (supply *Supply) SetPurchaseOrder(purchaseOrder *PurchaseOrder) *Supply {
	supply.PurchaseOrder = purchaseOrder.Clean()
	return supply
}

func (supply *Supply) SetFactureIn(factureIn *FactureIn) *Supply {
	supply.FactureIn = factureIn.Clean()
	return supply
}

func (supply *Supply) SetInvoicesIn(invoicesIn ...*InvoiceIn) *Supply {
	supply.InvoicesIn = invoicesIn
	return supply
}

func (supply *Supply) SetAttributes(attributes ...*Attribute) *Supply {
	supply.Attributes = attributes
	return supply
}

func (supply Supply) String() string {
	return Stringify(supply)
}

// MetaType возвращает тип сущности.
func (Supply) MetaType() MetaType {
	return MetaTypeSupply
}

// Update shortcut
func (supply Supply) Update(ctx context.Context, client *Client, params ...*Params) (*Supply, *resty.Response, error) {
	return client.Entity().Supply().Update(ctx, supply.GetID(), &supply, params...)
}

// Create shortcut
func (supply Supply) Create(ctx context.Context, client *Client, params ...*Params) (*Supply, *resty.Response, error) {
	return client.Entity().Supply().Create(ctx, &supply, params...)
}

// Delete shortcut
func (supply Supply) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return client.Entity().Supply().Delete(ctx, supply.GetID())
}

// SupplyPosition Позиция Приемки.
// Ключевое слово: supplyposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-priemka-priemki-pozicii-priemki
type SupplyPosition struct {
	Quantity      *float64            `json:"quantity,omitempty"`
	Pack          *Pack               `json:"pack,omitempty"`
	Country       *Country            `json:"country,omitempty"`
	Discount      *float64            `json:"discount,omitempty"`
	AccountID     *uuid.UUID          `json:"accountId,omitempty"`
	ID            *uuid.UUID          `json:"id,omitempty"`
	Assortment    *AssortmentPosition `json:"assortment,omitempty"`
	Price         *float64            `json:"price,omitempty"`
	GTD           *GTD                `json:"gtd,omitempty"`
	Slot          *Slot               `json:"slot,omitempty"`
	Stock         *Stock              `json:"stock,omitempty"`
	VatEnabled    *bool               `json:"vatEnabled,omitempty"`
	Overhead      *float64            `json:"overhead,omitempty"`
	Vat           *int                `json:"vat,omitempty"`
	TrackingCodes Slice[TrackingCode] `json:"trackingCodes,omitempty"`
	Things        Slice[string]       `json:"things,omitempty"`
}

func (supplyPosition SupplyPosition) GetAccountID() uuid.UUID {
	return Deref(supplyPosition.AccountID)
}

func (supplyPosition SupplyPosition) GetAssortment() AssortmentPosition {
	return Deref(supplyPosition.Assortment)
}

func (supplyPosition SupplyPosition) GetCountry() Country {
	return Deref(supplyPosition.Country)
}

func (supplyPosition SupplyPosition) GetDiscount() float64 {
	return Deref(supplyPosition.Discount)
}

func (supplyPosition SupplyPosition) GetGTD() GTD {
	return Deref(supplyPosition.GTD)
}

func (supplyPosition SupplyPosition) GetID() uuid.UUID {
	return Deref(supplyPosition.ID)
}

func (supplyPosition SupplyPosition) GetPack() Pack {
	return Deref(supplyPosition.Pack)
}

func (supplyPosition SupplyPosition) GetPrice() float64 {
	return Deref(supplyPosition.Price)
}

func (supplyPosition SupplyPosition) GetQuantity() float64 {
	return Deref(supplyPosition.Quantity)
}

func (supplyPosition SupplyPosition) GetSlot() Slot {
	return Deref(supplyPosition.Slot)
}

func (supplyPosition SupplyPosition) GetThings() Slice[string] {
	return supplyPosition.Things
}

func (supplyPosition SupplyPosition) GetTrackingCodes() Slice[TrackingCode] {
	return supplyPosition.TrackingCodes
}

func (supplyPosition SupplyPosition) GetOverhead() float64 {
	return Deref(supplyPosition.Overhead)
}

func (supplyPosition SupplyPosition) GetVat() int {
	return Deref(supplyPosition.Vat)
}

func (supplyPosition SupplyPosition) GetVatEnabled() bool {
	return Deref(supplyPosition.VatEnabled)
}

func (supplyPosition SupplyPosition) GetStock() Stock {
	return Deref(supplyPosition.Stock)
}

func (supplyPosition *SupplyPosition) SetAssortment(assortment AsAssortment) *SupplyPosition {
	supplyPosition.Assortment = assortment.AsAssortment()
	return supplyPosition
}

func (supplyPosition *SupplyPosition) SetCountry(country *Country) *SupplyPosition {
	supplyPosition.Country = country.Clean()
	return supplyPosition
}

func (supplyPosition *SupplyPosition) SetDiscount(discount float64) *SupplyPosition {
	supplyPosition.Discount = &discount
	return supplyPosition
}

func (supplyPosition *SupplyPosition) SetGTD(gtd *GTD) *SupplyPosition {
	supplyPosition.GTD = gtd
	return supplyPosition
}

func (supplyPosition *SupplyPosition) SetPack(pack *Pack) *SupplyPosition {
	supplyPosition.Pack = pack
	return supplyPosition
}

func (supplyPosition *SupplyPosition) SetPrice(price float64) *SupplyPosition {
	supplyPosition.Price = &price
	return supplyPosition
}

func (supplyPosition *SupplyPosition) SetQuantity(quantity float64) *SupplyPosition {
	supplyPosition.Quantity = &quantity
	return supplyPosition
}

func (supplyPosition *SupplyPosition) SetSlot(slot *Slot) *SupplyPosition {
	supplyPosition.Slot = slot.Clean()
	return supplyPosition
}

func (supplyPosition *SupplyPosition) SetThings(things ...string) *SupplyPosition {
	supplyPosition.Things = NewSliceFrom(things)
	return supplyPosition
}

func (supplyPosition *SupplyPosition) SetTrackingCodes(trackingCodes ...*TrackingCode) *SupplyPosition {
	supplyPosition.TrackingCodes = trackingCodes
	return supplyPosition
}

func (supplyPosition *SupplyPosition) SetVat(vat int) *SupplyPosition {
	supplyPosition.Vat = &vat
	return supplyPosition
}

func (supplyPosition *SupplyPosition) SetVatEnabled(vatEnabled bool) *SupplyPosition {
	supplyPosition.VatEnabled = &vatEnabled
	return supplyPosition
}

func (supplyPosition SupplyPosition) String() string {
	return Stringify(supplyPosition)
}

// MetaType возвращает тип сущности.
func (SupplyPosition) MetaType() MetaType {
	return MetaTypeSupplyPosition
}

// SupplyService
// Сервис для работы с приёмками.
type SupplyService interface {
	GetList(ctx context.Context, params ...*Params) (*List[Supply], *resty.Response, error)
	Create(ctx context.Context, supply *Supply, params ...*Params) (*Supply, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, supplyList Slice[Supply], params ...*Params) (*Slice[Supply], *resty.Response, error)
	DeleteMany(ctx context.Context, entities ...*Supply) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Supply, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, supply *Supply, params ...*Params) (*Supply, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetaAttributesSharedStatesWrapper, *resty.Response, error)
	Template(ctx context.Context) (*Supply, *resty.Response, error)
	GetPositions(ctx context.Context, id uuid.UUID, params ...*Params) (*MetaArray[SupplyPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params ...*Params) (*SupplyPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *SupplyPosition, params ...*Params) (*SupplyPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id uuid.UUID, position *SupplyPosition) (*SupplyPosition, *resty.Response, error)
	CreatePositionMany(ctx context.Context, id uuid.UUID, positions ...*SupplyPosition) (*Slice[SupplyPosition], *resty.Response, error)
	DeletePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (bool, *resty.Response, error)
	DeletePositionMany(ctx context.Context, id uuid.UUID, entities ...*SupplyPosition) (*DeleteManyResponse, *resty.Response, error)
	GetPositionTrackingCodes(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (*MetaArray[TrackingCode], *resty.Response, error)
	CreateUpdatePositionTrackingCodeMany(ctx context.Context, id uuid.UUID, positionID uuid.UUID, trackingCodes ...*TrackingCode) (*Slice[TrackingCode], *resty.Response, error)
	DeletePositionTrackingCodeMany(ctx context.Context, id uuid.UUID, positionID uuid.UUID, trackingCodes ...*TrackingCode) (*DeleteManyResponse, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributeMany(ctx context.Context, attributes ...*Attribute) (*Slice[Attribute], *resty.Response, error)
	UpdateAttribute(ctx context.Context, id uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributeMany(ctx context.Context, attributes ...*Attribute) (*DeleteManyResponse, *resty.Response, error)
	GetPublications(ctx context.Context, id uuid.UUID) (*MetaArray[Publication], *resty.Response, error)
	GetPublicationByID(ctx context.Context, id uuid.UUID, publicationID uuid.UUID) (*Publication, *resty.Response, error)
	Publish(ctx context.Context, id uuid.UUID, template TemplateInterface) (*Publication, *resty.Response, error)
	DeletePublication(ctx context.Context, id uuid.UUID, publicationID uuid.UUID) (bool, *resty.Response, error)
	PrintDocument(ctx context.Context, id uuid.UUID, PrintDocumentArg *PrintDocumentArg) (*PrintFile, *resty.Response, error)
	GetFiles(ctx context.Context, id uuid.UUID) (*MetaArray[File], *resty.Response, error)
	CreateFile(ctx context.Context, id uuid.UUID, file *File) (*Slice[File], *resty.Response, error)
	UpdateFileMany(ctx context.Context, id uuid.UUID, files ...*File) (*Slice[File], *resty.Response, error)
	DeleteFile(ctx context.Context, id uuid.UUID, fileID uuid.UUID) (bool, *resty.Response, error)
	DeleteFileMany(ctx context.Context, id uuid.UUID, files ...*File) (*DeleteManyResponse, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params ...*Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)
	GetStateByID(ctx context.Context, id uuid.UUID) (*State, *resty.Response, error)
	CreateState(ctx context.Context, state *State) (*State, *resty.Response, error)
	UpdateState(ctx context.Context, id uuid.UUID, state *State) (*State, *resty.Response, error)
	CreateUpdateStateMany(ctx context.Context, states ...*State) (*Slice[State], *resty.Response, error)
	DeleteState(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*Supply, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error)
	MoveToTrash(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	Evaluate(ctx context.Context, entity *Supply, evaluate ...Evaluate) (*Supply, *resty.Response, error)
}

func NewSupplyService(client *Client) SupplyService {
	e := NewEndpoint(client, "entity/supply")
	return newMainService[Supply, SupplyPosition, MetaAttributesSharedStatesWrapper, any](e)
}
