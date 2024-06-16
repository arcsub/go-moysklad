package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// PurchaseReturn Возврат поставщику.
// Ключевое слово: purchasereturn
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vozwrat-postawschiku
type PurchaseReturn struct {
	Printed             *bool                              `json:"printed,omitempty"`
	Supply              *Supply                            `json:"supply,omitempty"`
	AgentAccount        *AgentAccount                      `json:"agentAccount,omitempty"`
	Applicable          *bool                              `json:"applicable,omitempty"`
	Payments            Slice[Payment]                     `json:"payments,omitempty"`
	Code                *string                            `json:"code,omitempty"`
	OrganizationAccount *AgentAccount                      `json:"organizationAccount,omitempty"`
	Created             *Timestamp                         `json:"created,omitempty"`
	Deleted             *Timestamp                         `json:"deleted,omitempty"`
	Description         *string                            `json:"description,omitempty"`
	ExternalCode        *string                            `json:"externalCode,omitempty"`
	Files               *MetaArray[File]                   `json:"files,omitempty"`
	Group               *Group                             `json:"group,omitempty"`
	ID                  *uuid.UUID                         `json:"id,omitempty"`
	Meta                *Meta                              `json:"meta,omitempty"`
	Moment              *Timestamp                         `json:"moment,omitempty"`
	Name                *string                            `json:"name,omitempty"`
	AccountID           *uuid.UUID                         `json:"accountId,omitempty"`
	Contract            *NullValue[Contract]               `json:"contract,omitempty"`
	Agent               *Counterparty                      `json:"agent,omitempty"`
	Organization        *Organization                      `json:"organization,omitempty"`
	Project             *NullValue[Project]                `json:"project,omitempty"`
	Published           *bool                              `json:"published,omitempty"`
	Rate                *Rate                              `json:"rate,omitempty"`
	Shared              *bool                              `json:"shared,omitempty"`
	State               *State                             `json:"state,omitempty"`
	Store               *Store                             `json:"store,omitempty"`
	Sum                 *float64                           `json:"sum,omitempty"`
	SyncID              *uuid.UUID                         `json:"syncId,omitempty"`
	Updated             *Timestamp                         `json:"updated,omitempty"`
	VatEnabled          *bool                              `json:"vatEnabled,omitempty"`
	VatIncluded         *bool                              `json:"vatIncluded,omitempty"`
	VatSum              *float64                           `json:"vatSum,omitempty"`
	Positions           *Positions[PurchaseReturnPosition] `json:"positions,omitempty"`
	Owner               *Employee                          `json:"owner,omitempty"`
	FactureIn           *FactureIn                         `json:"factureIn,omitempty"`
	FactureOut          *FactureOut                        `json:"factureOut,omitempty"`
	InvoicedSum         *float64                           `json:"invoicedSum,omitempty"`
	PayedSum            *float64                           `json:"payedSum,omitempty"`
	Attributes          Slice[AttributeValue]              `json:"attributes,omitempty"`
}

// Clean возвращает сущность с единственным заполненным полем Meta
func (purchaseReturn PurchaseReturn) Clean() *PurchaseReturn {
	return &PurchaseReturn{Meta: purchaseReturn.Meta}
}

// AsOperation возвращает объект Operation c полем Meta сущности
func (purchaseReturn PurchaseReturn) AsOperation() *Operation {
	return &Operation{Meta: purchaseReturn.GetMeta()}
}

func (purchaseReturn PurchaseReturn) GetPrinted() bool {
	return Deref(purchaseReturn.Printed)
}

func (purchaseReturn PurchaseReturn) GetSupply() Supply {
	return Deref(purchaseReturn.Supply)
}

func (purchaseReturn PurchaseReturn) GetAgentAccount() AgentAccount {
	return Deref(purchaseReturn.AgentAccount)
}

func (purchaseReturn PurchaseReturn) GetApplicable() bool {
	return Deref(purchaseReturn.Applicable)
}

func (purchaseReturn PurchaseReturn) GetPayments() Slice[Payment] {
	return purchaseReturn.Payments
}

func (purchaseReturn PurchaseReturn) GetCode() string {
	return Deref(purchaseReturn.Code)
}

func (purchaseReturn PurchaseReturn) GetOrganizationAccount() AgentAccount {
	return Deref(purchaseReturn.OrganizationAccount)
}

func (purchaseReturn PurchaseReturn) GetCreated() Timestamp {
	return Deref(purchaseReturn.Created)
}

func (purchaseReturn PurchaseReturn) GetDeleted() Timestamp {
	return Deref(purchaseReturn.Deleted)
}

func (purchaseReturn PurchaseReturn) GetDescription() string {
	return Deref(purchaseReturn.Description)
}

func (purchaseReturn PurchaseReturn) GetExternalCode() string {
	return Deref(purchaseReturn.ExternalCode)
}

func (purchaseReturn PurchaseReturn) GetFiles() MetaArray[File] {
	return Deref(purchaseReturn.Files)
}

func (purchaseReturn PurchaseReturn) GetGroup() Group {
	return Deref(purchaseReturn.Group)
}

func (purchaseReturn PurchaseReturn) GetID() uuid.UUID {
	return Deref(purchaseReturn.ID)
}

func (purchaseReturn PurchaseReturn) GetMeta() Meta {
	return Deref(purchaseReturn.Meta)
}

func (purchaseReturn PurchaseReturn) GetMoment() Timestamp {
	return Deref(purchaseReturn.Moment)
}

func (purchaseReturn PurchaseReturn) GetName() string {
	return Deref(purchaseReturn.Name)
}

func (purchaseReturn PurchaseReturn) GetAccountID() uuid.UUID {
	return Deref(purchaseReturn.AccountID)
}

func (purchaseReturn PurchaseReturn) GetContract() Contract {
	return purchaseReturn.Contract.Get()
}

func (purchaseReturn PurchaseReturn) GetAgent() Counterparty {
	return Deref(purchaseReturn.Agent)
}

func (purchaseReturn PurchaseReturn) GetOrganization() Organization {
	return Deref(purchaseReturn.Organization)
}

func (purchaseReturn PurchaseReturn) GetProject() Project {
	return purchaseReturn.Project.Get()
}

func (purchaseReturn PurchaseReturn) GetPublished() bool {
	return Deref(purchaseReturn.Published)
}

func (purchaseReturn PurchaseReturn) GetRate() Rate {
	return Deref(purchaseReturn.Rate)
}

func (purchaseReturn PurchaseReturn) GetShared() bool {
	return Deref(purchaseReturn.Shared)
}

func (purchaseReturn PurchaseReturn) GetState() State {
	return Deref(purchaseReturn.State)
}

func (purchaseReturn PurchaseReturn) GetStore() Store {
	return Deref(purchaseReturn.Store)
}

func (purchaseReturn PurchaseReturn) GetSum() float64 {
	return Deref(purchaseReturn.Sum)
}

func (purchaseReturn PurchaseReturn) GetSyncID() uuid.UUID {
	return Deref(purchaseReturn.SyncID)
}

func (purchaseReturn PurchaseReturn) GetUpdated() Timestamp {
	return Deref(purchaseReturn.Updated)
}

func (purchaseReturn PurchaseReturn) GetVatEnabled() bool {
	return Deref(purchaseReturn.VatEnabled)
}

func (purchaseReturn PurchaseReturn) GetVatIncluded() bool {
	return Deref(purchaseReturn.VatIncluded)
}

func (purchaseReturn PurchaseReturn) GetVatSum() float64 {
	return Deref(purchaseReturn.VatSum)
}

func (purchaseReturn PurchaseReturn) GetPositions() Positions[PurchaseReturnPosition] {
	return Deref(purchaseReturn.Positions)
}

func (purchaseReturn PurchaseReturn) GetOwner() Employee {
	return Deref(purchaseReturn.Owner)
}

func (purchaseReturn PurchaseReturn) GetFactureIn() FactureIn {
	return Deref(purchaseReturn.FactureIn)
}

func (purchaseReturn PurchaseReturn) GetFactureOut() FactureOut {
	return Deref(purchaseReturn.FactureOut)
}

func (purchaseReturn PurchaseReturn) GetInvoicedSum() float64 {
	return Deref(purchaseReturn.InvoicedSum)
}

func (purchaseReturn PurchaseReturn) GetPayedSum() float64 {
	return Deref(purchaseReturn.PayedSum)
}

func (purchaseReturn PurchaseReturn) GetAttributes() Slice[AttributeValue] {
	return purchaseReturn.Attributes
}

func (purchaseReturn *PurchaseReturn) SetSupply(supply *Supply) *PurchaseReturn {
	purchaseReturn.Supply = supply.Clean()
	return purchaseReturn
}

func (purchaseReturn *PurchaseReturn) SetAgentAccount(agentAccount *AgentAccount) *PurchaseReturn {
	purchaseReturn.AgentAccount = agentAccount.Clean()
	return purchaseReturn
}

func (purchaseReturn *PurchaseReturn) SetApplicable(applicable bool) *PurchaseReturn {
	purchaseReturn.Applicable = &applicable
	return purchaseReturn
}

func (purchaseReturn *PurchaseReturn) SetPayments(payments Slice[Payment]) *PurchaseReturn {
	purchaseReturn.Payments = payments
	return purchaseReturn
}

func (purchaseReturn *PurchaseReturn) SetCode(code string) *PurchaseReturn {
	purchaseReturn.Code = &code
	return purchaseReturn
}

func (purchaseReturn *PurchaseReturn) SetOrganizationAccount(organizationAccount *AgentAccount) *PurchaseReturn {
	purchaseReturn.OrganizationAccount = organizationAccount.Clean()
	return purchaseReturn
}

func (purchaseReturn *PurchaseReturn) SetDescription(description string) *PurchaseReturn {
	purchaseReturn.Description = &description
	return purchaseReturn
}

func (purchaseReturn *PurchaseReturn) SetExternalCode(externalCode string) *PurchaseReturn {
	purchaseReturn.ExternalCode = &externalCode
	return purchaseReturn
}

func (purchaseReturn *PurchaseReturn) SetFiles(files Slice[File]) *PurchaseReturn {
	purchaseReturn.Files = NewMetaArrayRows(files)
	return purchaseReturn
}

func (purchaseReturn *PurchaseReturn) SetGroup(group *Group) *PurchaseReturn {
	purchaseReturn.Group = group.Clean()
	return purchaseReturn
}

func (purchaseReturn *PurchaseReturn) SetMeta(meta *Meta) *PurchaseReturn {
	purchaseReturn.Meta = meta
	return purchaseReturn
}

func (purchaseReturn *PurchaseReturn) SetMoment(moment *Timestamp) *PurchaseReturn {
	purchaseReturn.Moment = moment
	return purchaseReturn
}

func (purchaseReturn *PurchaseReturn) SetName(name string) *PurchaseReturn {
	purchaseReturn.Name = &name
	return purchaseReturn
}

func (purchaseReturn *PurchaseReturn) SetContract(contract *Contract) *PurchaseReturn {
	purchaseReturn.Contract = NewNullValueWith(contract.Clean())
	return purchaseReturn
}

func (purchaseReturn *PurchaseReturn) SetAgent(agent *Counterparty) *PurchaseReturn {
	purchaseReturn.Agent = agent.Clean()
	return purchaseReturn
}

func (purchaseReturn *PurchaseReturn) SetOrganization(organization *Organization) *PurchaseReturn {
	purchaseReturn.Organization = organization.Clean()
	return purchaseReturn
}

func (purchaseReturn *PurchaseReturn) SetProject(project *Project) *PurchaseReturn {
	purchaseReturn.Project = NewNullValueWith(project.Clean())
	return purchaseReturn
}

func (purchaseReturn *PurchaseReturn) SetNullProject() *PurchaseReturn {
	purchaseReturn.Project = NewNullValue[Project]()
	return purchaseReturn
}

func (purchaseReturn *PurchaseReturn) SetRate(rate *Rate) *PurchaseReturn {
	purchaseReturn.Rate = rate
	return purchaseReturn
}

func (purchaseReturn *PurchaseReturn) SetShared(shared bool) *PurchaseReturn {
	purchaseReturn.Shared = &shared
	return purchaseReturn
}

func (purchaseReturn *PurchaseReturn) SetState(state *State) *PurchaseReturn {
	purchaseReturn.State = state.Clean()
	return purchaseReturn
}

func (purchaseReturn *PurchaseReturn) SetStore(store *Store) *PurchaseReturn {
	purchaseReturn.Store = store.Clean()
	return purchaseReturn
}

func (purchaseReturn *PurchaseReturn) SetSyncID(syncID uuid.UUID) *PurchaseReturn {
	purchaseReturn.SyncID = &syncID
	return purchaseReturn
}

func (purchaseReturn *PurchaseReturn) SetVatEnabled(vatEnabled bool) *PurchaseReturn {
	purchaseReturn.VatEnabled = &vatEnabled
	return purchaseReturn
}

func (purchaseReturn *PurchaseReturn) SetVatIncluded(vatIncluded bool) *PurchaseReturn {
	purchaseReturn.VatIncluded = &vatIncluded
	return purchaseReturn
}

func (purchaseReturn *PurchaseReturn) SetVatSum(vatSum float64) *PurchaseReturn {
	purchaseReturn.VatSum = &vatSum
	return purchaseReturn
}

func (purchaseReturn *PurchaseReturn) SetPositions(positions *Positions[PurchaseReturnPosition]) *PurchaseReturn {
	purchaseReturn.Positions = positions
	return purchaseReturn
}

func (purchaseReturn *PurchaseReturn) SetOwner(owner *Employee) *PurchaseReturn {
	purchaseReturn.Owner = owner.Clean()
	return purchaseReturn
}

func (purchaseReturn *PurchaseReturn) SetFactureIn(factureIn *FactureIn) *PurchaseReturn {
	purchaseReturn.FactureIn = factureIn.Clean()
	return purchaseReturn
}

func (purchaseReturn *PurchaseReturn) SetFactureOut(factureOut *FactureOut) *PurchaseReturn {
	purchaseReturn.FactureOut = factureOut.Clean()
	return purchaseReturn
}

func (purchaseReturn *PurchaseReturn) SetInvoicedSum(invoicedSum float64) *PurchaseReturn {
	purchaseReturn.InvoicedSum = &invoicedSum
	return purchaseReturn
}

func (purchaseReturn *PurchaseReturn) SetPayedSum(payedSum float64) *PurchaseReturn {
	purchaseReturn.PayedSum = &payedSum
	return purchaseReturn
}

func (purchaseReturn *PurchaseReturn) SetAttributes(attributes Slice[AttributeValue]) *PurchaseReturn {
	purchaseReturn.Attributes = attributes
	return purchaseReturn
}

func (purchaseReturn PurchaseReturn) String() string {
	return Stringify(purchaseReturn)
}

func (purchaseReturn PurchaseReturn) MetaType() MetaType {
	return MetaTypePurchaseReturn
}

// PurchaseReturnPosition Позиция Возврата поставщику.
// Ключевое слово: purchasereturnposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vozwrat-postawschiku-vozwraty-postawschikam-pozicii-vozwrata-postawschiku
type PurchaseReturnPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`
	Assortment *AssortmentPosition `json:"assortment,omitempty"`
	Discount   *float64            `json:"discount,omitempty"`
	ID         *uuid.UUID          `json:"id,omitempty"`
	Pack       *Pack               `json:"pack,omitempty"`
	Price      *float64            `json:"price,omitempty"`
	Quantity   *float64            `json:"quantity,omitempty"`
	Slot       *Slot               `json:"slot,omitempty"`
	Vat        *int                `json:"vat,omitempty"`
	VatEnabled *bool               `json:"vatEnabled,omitempty"`
	Stock      *Stock              `json:"stock,omitempty"`
	Things     Slice[string]       `json:"things,omitempty"`
}

func (purchaseReturnPosition PurchaseReturnPosition) GetAccountID() uuid.UUID {
	return Deref(purchaseReturnPosition.AccountID)
}

func (purchaseReturnPosition PurchaseReturnPosition) GetAssortment() AssortmentPosition {
	return Deref(purchaseReturnPosition.Assortment)
}

func (purchaseReturnPosition PurchaseReturnPosition) GetDiscount() float64 {
	return Deref(purchaseReturnPosition.Discount)
}

func (purchaseReturnPosition PurchaseReturnPosition) GetID() uuid.UUID {
	return Deref(purchaseReturnPosition.ID)
}

func (purchaseReturnPosition PurchaseReturnPosition) GetPack() Pack {
	return Deref(purchaseReturnPosition.Pack)
}

func (purchaseReturnPosition PurchaseReturnPosition) GetPrice() float64 {
	return Deref(purchaseReturnPosition.Price)
}

func (purchaseReturnPosition PurchaseReturnPosition) GetQuantity() float64 {
	return Deref(purchaseReturnPosition.Quantity)
}

func (purchaseReturnPosition PurchaseReturnPosition) GetSlot() Slot {
	return Deref(purchaseReturnPosition.Slot)
}

func (purchaseReturnPosition PurchaseReturnPosition) GetVat() int {
	return Deref(purchaseReturnPosition.Vat)
}

func (purchaseReturnPosition PurchaseReturnPosition) GetVatEnabled() bool {
	return Deref(purchaseReturnPosition.VatEnabled)
}

func (purchaseReturnPosition PurchaseReturnPosition) GetStock() Stock {
	return Deref(purchaseReturnPosition.Stock)
}

func (purchaseReturnPosition PurchaseReturnPosition) GetThings() Slice[string] {
	return purchaseReturnPosition.Things
}

func (purchaseReturnPosition *PurchaseReturnPosition) SetAssortment(assortment AsAssortment) *PurchaseReturnPosition {
	purchaseReturnPosition.Assortment = assortment.AsAssortment()
	return purchaseReturnPosition
}

func (purchaseReturnPosition *PurchaseReturnPosition) SetDiscount(discount float64) *PurchaseReturnPosition {
	purchaseReturnPosition.Discount = &discount
	return purchaseReturnPosition
}

func (purchaseReturnPosition *PurchaseReturnPosition) SetPack(pack *Pack) *PurchaseReturnPosition {
	purchaseReturnPosition.Pack = pack
	return purchaseReturnPosition
}

func (purchaseReturnPosition *PurchaseReturnPosition) SetPrice(price float64) *PurchaseReturnPosition {
	purchaseReturnPosition.Price = &price
	return purchaseReturnPosition
}

func (purchaseReturnPosition *PurchaseReturnPosition) SetQuantity(quantity float64) *PurchaseReturnPosition {
	purchaseReturnPosition.Quantity = &quantity
	return purchaseReturnPosition
}

func (purchaseReturnPosition *PurchaseReturnPosition) SetSlot(slot *Slot) *PurchaseReturnPosition {
	purchaseReturnPosition.Slot = slot.Clean()
	return purchaseReturnPosition
}

func (purchaseReturnPosition *PurchaseReturnPosition) SetVat(vat int) *PurchaseReturnPosition {
	purchaseReturnPosition.Vat = &vat
	return purchaseReturnPosition
}

func (purchaseReturnPosition *PurchaseReturnPosition) SetVatEnabled(vatEnabled bool) *PurchaseReturnPosition {
	purchaseReturnPosition.VatEnabled = &vatEnabled
	return purchaseReturnPosition
}

func (purchaseReturnPosition *PurchaseReturnPosition) SetThings(things Slice[string]) *PurchaseReturnPosition {
	purchaseReturnPosition.Things = things
	return purchaseReturnPosition
}

func (purchaseReturnPosition PurchaseReturnPosition) String() string {
	return Stringify(purchaseReturnPosition)
}

func (purchaseReturnPosition PurchaseReturnPosition) MetaType() MetaType {
	return MetaTypePurchaseReturnPosition
}

// PurchaseReturnTemplateArg
// Документ: Возврат поставщику (purchasereturn)
// Основание, на котором он может быть создан:
// - Приемка (supply)
//type PurchaseReturnTemplateArg struct {
//	Supply *MetaWrapper `json:"supply,omitempty"`
//}

// PurchaseReturnService
// Сервис для работы с возвратами поставщикам.
type PurchaseReturnService interface {
	GetList(ctx context.Context, params ...*Params) (*List[PurchaseReturn], *resty.Response, error)
	Create(ctx context.Context, purchaseReturn *PurchaseReturn, params ...*Params) (*PurchaseReturn, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, purchaseReturnList Slice[PurchaseReturn], params ...*Params) (*Slice[PurchaseReturn], *resty.Response, error)
	DeleteMany(ctx context.Context, purchaseReturnList []MetaWrapper) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*PurchaseReturn, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, purchaseReturn *PurchaseReturn, params ...*Params) (*PurchaseReturn, *resty.Response, error)
	Template(ctx context.Context) (*PurchaseReturn, *resty.Response, error)
	TemplateBased(ctx context.Context, basedOn ...MetaOwner) (*PurchaseReturn, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetaAttributesSharedStatesWrapper, *resty.Response, error)
	GetPositions(ctx context.Context, id uuid.UUID, params ...*Params) (*MetaArray[PurchaseReturnPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params ...*Params) (*PurchaseReturnPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *PurchaseReturnPosition, params ...*Params) (*PurchaseReturnPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id uuid.UUID, position *PurchaseReturnPosition) (*PurchaseReturnPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id uuid.UUID, positions Slice[PurchaseReturnPosition]) (*Slice[PurchaseReturnPosition], *resty.Response, error)
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
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*PurchaseReturn, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error)
	GetNamedFilters(ctx context.Context, params ...*Params) (*List[NamedFilter], *resty.Response, error)
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)
	MoveToTrash(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
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

func NewPurchaseReturnService(client *Client) PurchaseReturnService {
	e := NewEndpoint(client, "entity/purchasereturn")
	return newMainService[PurchaseReturn, PurchaseReturnPosition, MetaAttributesSharedStatesWrapper, any](e)
}
