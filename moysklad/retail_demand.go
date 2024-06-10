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
	AccountID           *uuid.UUID                       `json:"accountId,omitempty"`
	Agent               *Counterparty                    `json:"agent,omitempty"`
	AgentAccount        *AgentAccount                    `json:"agentAccount,omitempty"`
	Applicable          *bool                            `json:"applicable,omitempty"`
	CashSum             *float64                         `json:"cashSum,omitempty"`
	CheckNumber         *string                          `json:"checkNumber,omitempty"`
	CheckSum            *float64                         `json:"checkSum,omitempty"`
	Code                *string                          `json:"code,omitempty"`
	Contract            *Contract                        `json:"contract,omitempty"`
	Created             *Timestamp                       `json:"created,omitempty"`
	CustomerOrder       *CustomerOrder                   `json:"customerOrder,omitempty"`
	Deleted             *Timestamp                       `json:"deleted,omitempty"`
	Description         *string                          `json:"description,omitempty"`
	DocumentNumber      *string                          `json:"documentNumber,omitempty"`
	ExternalCode        *string                          `json:"externalCode,omitempty"`
	Files               *MetaArray[File]                 `json:"files,omitempty"`
	Fiscal              *bool                            `json:"fiscal,omitempty"`
	FiscalPrinterInfo   *string                          `json:"fiscalPrinterInfo,omitempty"`
	Group               *Group                           `json:"group,omitempty"`
	ID                  *uuid.UUID                       `json:"id,omitempty"`
	Meta                *Meta                            `json:"meta,omitempty"`
	Moment              *Timestamp                       `json:"moment,omitempty"`
	Name                *string                          `json:"name,omitempty"`
	NoCashSum           *float64                         `json:"noCashSum,omitempty"`
	OfdCode             *string                          `json:"ofdCode,omitempty"`
	Organization        *Organization                    `json:"organization,omitempty"`
	OrganizationAccount *AgentAccount                    `json:"organizationAccount,omitempty"`
	Owner               *Employee                        `json:"owner,omitempty"`
	PayedSum            *float64                         `json:"payedSum,omitempty"`
	Positions           *Positions[RetailDemandPosition] `json:"positions,omitempty"`
	PrepaymentCashSum   *float64                         `json:"prepaymentCashSum,omitempty"`
	PrepaymentNoCashSum *float64                         `json:"prepaymentNoCashSum,omitempty"`
	PrepaymentQRSum     *float64                         `json:"prepaymentQrSum,omitempty"`
	Printed             *bool                            `json:"printed,omitempty"`
	Project             *Project                         `json:"project,omitempty"`
	Published           *bool                            `json:"published,omitempty"`
	QRSum               *float64                         `json:"qrSum,omitempty"`
	Rate                *Rate                            `json:"rate,omitempty"`
	RetailShift         *RetailShift                     `json:"retailShift,omitempty"`
	RetailStore         *RetailStore                     `json:"retailStore,omitempty"`
	SessionNumber       *string                          `json:"sessionNumber,omitempty"`
	Shared              *bool                            `json:"shared,omitempty"`
	State               *State                           `json:"state,omitempty"`
	Store               *Store                           `json:"store,omitempty"`
	Sum                 *float64                         `json:"sum,omitempty"`
	SyncID              *uuid.UUID                       `json:"syncId,omitempty"`
	Updated             *Timestamp                       `json:"updated,omitempty"`
	VatEnabled          *bool                            `json:"vatEnabled,omitempty"`
	VatIncluded         *bool                            `json:"vatIncluded,omitempty"`
	VatSum              *float64                         `json:"vatSum,omitempty"`
	TaxSystem           TaxSystem                        `json:"taxSystem,omitempty"`
	Attributes          Slice[AttributeValue]            `json:"attributes,omitempty"`
}

func (retailDemand RetailDemand) Clean() *RetailDemand {
	return &RetailDemand{Meta: retailDemand.Meta}
}

func (retailDemand RetailDemand) GetAccountID() uuid.UUID {
	return Deref(retailDemand.AccountID)
}

func (retailDemand RetailDemand) GetAgent() Counterparty {
	return Deref(retailDemand.Agent)
}

func (retailDemand RetailDemand) GetAgentAccount() AgentAccount {
	return Deref(retailDemand.AgentAccount)
}

func (retailDemand RetailDemand) GetApplicable() bool {
	return Deref(retailDemand.Applicable)
}

func (retailDemand RetailDemand) GetCashSum() float64 {
	return Deref(retailDemand.CashSum)
}

func (retailDemand RetailDemand) GetCheckNumber() string {
	return Deref(retailDemand.CheckNumber)
}

func (retailDemand RetailDemand) GetCheckSum() float64 {
	return Deref(retailDemand.CheckSum)
}

func (retailDemand RetailDemand) GetCode() string {
	return Deref(retailDemand.Code)
}

func (retailDemand RetailDemand) GetContract() Contract {
	return Deref(retailDemand.Contract)
}

func (retailDemand RetailDemand) GetCreated() Timestamp {
	return Deref(retailDemand.Created)
}

func (retailDemand RetailDemand) GetCustomerOrder() CustomerOrder {
	return Deref(retailDemand.CustomerOrder)
}

func (retailDemand RetailDemand) GetDeleted() Timestamp {
	return Deref(retailDemand.Deleted)
}

func (retailDemand RetailDemand) GetDescription() string {
	return Deref(retailDemand.Description)
}

func (retailDemand RetailDemand) GetDocumentNumber() string {
	return Deref(retailDemand.DocumentNumber)
}

func (retailDemand RetailDemand) GetExternalCode() string {
	return Deref(retailDemand.ExternalCode)
}

func (retailDemand RetailDemand) GetFiles() MetaArray[File] {
	return Deref(retailDemand.Files)
}

func (retailDemand RetailDemand) GetFiscal() bool {
	return Deref(retailDemand.Fiscal)
}

func (retailDemand RetailDemand) GetFiscalPrinterInfo() string {
	return Deref(retailDemand.FiscalPrinterInfo)
}

func (retailDemand RetailDemand) GetGroup() Group {
	return Deref(retailDemand.Group)
}

func (retailDemand RetailDemand) GetID() uuid.UUID {
	return Deref(retailDemand.ID)
}

func (retailDemand RetailDemand) GetMeta() Meta {
	return Deref(retailDemand.Meta)
}

func (retailDemand RetailDemand) GetMoment() Timestamp {
	return Deref(retailDemand.Moment)
}

func (retailDemand RetailDemand) GetName() string {
	return Deref(retailDemand.Name)
}

func (retailDemand RetailDemand) GetNoCashSum() float64 {
	return Deref(retailDemand.NoCashSum)
}

func (retailDemand RetailDemand) GetOfdCode() string {
	return Deref(retailDemand.OfdCode)
}

func (retailDemand RetailDemand) GetOrganization() Organization {
	return Deref(retailDemand.Organization)
}

func (retailDemand RetailDemand) GetOrganizationAccount() AgentAccount {
	return Deref(retailDemand.OrganizationAccount)
}

func (retailDemand RetailDemand) GetOwner() Employee {
	return Deref(retailDemand.Owner)
}

func (retailDemand RetailDemand) GetPayedSum() float64 {
	return Deref(retailDemand.PayedSum)
}

func (retailDemand RetailDemand) GetPositions() Positions[RetailDemandPosition] {
	return Deref(retailDemand.Positions)
}

func (retailDemand RetailDemand) GetPrepaymentCashSum() float64 {
	return Deref(retailDemand.PrepaymentCashSum)
}

func (retailDemand RetailDemand) GetPrepaymentNoCashSum() float64 {
	return Deref(retailDemand.PrepaymentNoCashSum)
}

func (retailDemand RetailDemand) GetPrepaymentQRSum() float64 {
	return Deref(retailDemand.PrepaymentQRSum)
}

func (retailDemand RetailDemand) GetPrinted() bool {
	return Deref(retailDemand.Printed)
}

func (retailDemand RetailDemand) GetProject() Project {
	return Deref(retailDemand.Project)
}

func (retailDemand RetailDemand) GetPublished() bool {
	return Deref(retailDemand.Published)
}

func (retailDemand RetailDemand) GetQRSum() float64 {
	return Deref(retailDemand.QRSum)
}

func (retailDemand RetailDemand) GetRate() Rate {
	return Deref(retailDemand.Rate)
}

func (retailDemand RetailDemand) GetRetailShift() RetailShift {
	return Deref(retailDemand.RetailShift)
}

func (retailDemand RetailDemand) GetRetailStore() RetailStore {
	return Deref(retailDemand.RetailStore)
}

func (retailDemand RetailDemand) GetSessionNumber() string {
	return Deref(retailDemand.SessionNumber)
}

func (retailDemand RetailDemand) GetShared() bool {
	return Deref(retailDemand.Shared)
}

func (retailDemand RetailDemand) GetState() State {
	return Deref(retailDemand.State)
}

func (retailDemand RetailDemand) GetStore() Store {
	return Deref(retailDemand.Store)
}

func (retailDemand RetailDemand) GetSum() float64 {
	return Deref(retailDemand.Sum)
}

func (retailDemand RetailDemand) GetSyncID() uuid.UUID {
	return Deref(retailDemand.SyncID)
}

func (retailDemand RetailDemand) GetUpdated() Timestamp {
	return Deref(retailDemand.Updated)
}

func (retailDemand RetailDemand) GetVatEnabled() bool {
	return Deref(retailDemand.VatEnabled)
}

func (retailDemand RetailDemand) GetVatIncluded() bool {
	return Deref(retailDemand.VatIncluded)
}

func (retailDemand RetailDemand) GetVatSum() float64 {
	return Deref(retailDemand.VatSum)
}

func (retailDemand RetailDemand) GetTaxSystem() TaxSystem {
	return retailDemand.TaxSystem
}

func (retailDemand RetailDemand) GetAttributes() Slice[AttributeValue] {
	return retailDemand.Attributes
}

func (retailDemand *RetailDemand) SetAgent(agent *Counterparty) *RetailDemand {
	retailDemand.Agent = agent
	return retailDemand
}

func (retailDemand *RetailDemand) SetAgentAccount(agentAccount *AgentAccount) *RetailDemand {
	retailDemand.AgentAccount = agentAccount
	return retailDemand
}

func (retailDemand *RetailDemand) SetApplicable(applicable bool) *RetailDemand {
	retailDemand.Applicable = &applicable
	return retailDemand
}

func (retailDemand *RetailDemand) SetCashSum(cashSum float64) *RetailDemand {
	retailDemand.CashSum = &cashSum
	return retailDemand
}

func (retailDemand *RetailDemand) SetCheckNumber(checkNumber string) *RetailDemand {
	retailDemand.CheckNumber = &checkNumber
	return retailDemand
}

func (retailDemand *RetailDemand) SetCheckSum(checkSum float64) *RetailDemand {
	retailDemand.CheckSum = &checkSum
	return retailDemand
}

func (retailDemand *RetailDemand) SetCode(code string) *RetailDemand {
	retailDemand.Code = &code
	return retailDemand
}

func (retailDemand *RetailDemand) SetContract(contract *Contract) *RetailDemand {
	retailDemand.Contract = contract
	return retailDemand
}

func (retailDemand *RetailDemand) SetCustomerOrder(customerOrder *CustomerOrder) *RetailDemand {
	retailDemand.CustomerOrder = customerOrder
	return retailDemand
}

func (retailDemand *RetailDemand) SetDescription(description string) *RetailDemand {
	retailDemand.Description = &description
	return retailDemand
}

func (retailDemand *RetailDemand) SetDocumentNumber(documentNumber string) *RetailDemand {
	retailDemand.DocumentNumber = &documentNumber
	return retailDemand
}

func (retailDemand *RetailDemand) SetExternalCode(externalCode string) *RetailDemand {
	retailDemand.ExternalCode = &externalCode
	return retailDemand
}

func (retailDemand *RetailDemand) SetFiles(files Slice[File]) *RetailDemand {
	retailDemand.Files = NewMetaArrayRows(files)
	return retailDemand
}

func (retailDemand *RetailDemand) SetGroup(group *Group) *RetailDemand {
	retailDemand.Group = group
	return retailDemand
}

func (retailDemand *RetailDemand) SetMeta(meta *Meta) *RetailDemand {
	retailDemand.Meta = meta
	return retailDemand
}

func (retailDemand *RetailDemand) SetMoment(moment *Timestamp) *RetailDemand {
	retailDemand.Moment = moment
	return retailDemand
}

func (retailDemand *RetailDemand) SetName(name string) *RetailDemand {
	retailDemand.Name = &name
	return retailDemand
}

func (retailDemand *RetailDemand) SetNoCashSum(noCashSum float64) *RetailDemand {
	retailDemand.NoCashSum = &noCashSum
	return retailDemand
}

func (retailDemand *RetailDemand) SetOrganization(organization *Organization) *RetailDemand {
	retailDemand.Organization = organization
	return retailDemand
}

func (retailDemand *RetailDemand) SetOrganizationAccount(organizationAccount *AgentAccount) *RetailDemand {
	retailDemand.OrganizationAccount = organizationAccount
	return retailDemand
}

func (retailDemand *RetailDemand) SetOwner(owner *Employee) *RetailDemand {
	retailDemand.Owner = owner
	return retailDemand
}

func (retailDemand *RetailDemand) SetPositions(positions *Positions[RetailDemandPosition]) *RetailDemand {
	retailDemand.Positions = positions
	return retailDemand
}

func (retailDemand *RetailDemand) SetPrepaymentCashSum(prepaymentCashSum float64) *RetailDemand {
	retailDemand.PrepaymentCashSum = &prepaymentCashSum
	return retailDemand
}

func (retailDemand *RetailDemand) SetPrepaymentNoCashSum(prepaymentNoCashSum float64) *RetailDemand {
	retailDemand.PrepaymentNoCashSum = &prepaymentNoCashSum
	return retailDemand
}

func (retailDemand *RetailDemand) SetPrepaymentQRSum(prepaymentQRSum float64) *RetailDemand {
	retailDemand.PrepaymentQRSum = &prepaymentQRSum
	return retailDemand
}

func (retailDemand *RetailDemand) SetProject(project *Project) *RetailDemand {
	retailDemand.Project = project
	return retailDemand
}

func (retailDemand *RetailDemand) SetQRSum(qrSum float64) *RetailDemand {
	retailDemand.QRSum = &qrSum
	return retailDemand
}

func (retailDemand *RetailDemand) SetRate(rate *Rate) *RetailDemand {
	retailDemand.Rate = rate
	return retailDemand
}

func (retailDemand *RetailDemand) SetRetailShift(retailShift *RetailShift) *RetailDemand {
	retailDemand.RetailShift = retailShift
	return retailDemand
}

func (retailDemand *RetailDemand) SetRetailStore(retailStore *RetailStore) *RetailDemand {
	retailDemand.RetailStore = retailStore
	return retailDemand
}

func (retailDemand *RetailDemand) SetSessionNumber(sessionNumber string) *RetailDemand {
	retailDemand.SessionNumber = &sessionNumber
	return retailDemand
}

func (retailDemand *RetailDemand) SetShared(shared bool) *RetailDemand {
	retailDemand.Shared = &shared
	return retailDemand
}

func (retailDemand *RetailDemand) SetState(state *State) *RetailDemand {
	retailDemand.State = state
	return retailDemand
}

func (retailDemand *RetailDemand) SetStore(store *Store) *RetailDemand {
	retailDemand.Store = store
	return retailDemand
}

func (retailDemand *RetailDemand) SetSyncID(syncID uuid.UUID) *RetailDemand {
	retailDemand.SyncID = &syncID
	return retailDemand
}

func (retailDemand *RetailDemand) SetVatEnabled(vatEnabled bool) *RetailDemand {
	retailDemand.VatEnabled = &vatEnabled
	return retailDemand
}

func (retailDemand *RetailDemand) SetVatIncluded(vatIncluded bool) *RetailDemand {
	retailDemand.VatIncluded = &vatIncluded
	return retailDemand
}

func (retailDemand *RetailDemand) SetVatSum(vatSum float64) *RetailDemand {
	retailDemand.VatSum = &vatSum
	return retailDemand
}

func (retailDemand *RetailDemand) SetTaxSystem(taxSystem TaxSystem) *RetailDemand {
	retailDemand.TaxSystem = taxSystem
	return retailDemand
}

func (retailDemand *RetailDemand) SetAttributes(attributes Slice[AttributeValue]) *RetailDemand {
	retailDemand.Attributes = attributes
	return retailDemand
}

func (retailDemand RetailDemand) String() string {
	return Stringify(retailDemand)
}

func (retailDemand RetailDemand) MetaType() MetaType {
	return MetaTypeRetailDemand
}

// RetailDemandPosition позиция розничной продажи.
// Ключевое слово: demandposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-roznichnaq-prodazha-roznichnye-prodazhi-pozicii-roznichnoj-prodazhi
type RetailDemandPosition struct {
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

func (retailDemandPosition RetailDemandPosition) GetAccountID() uuid.UUID {
	return Deref(retailDemandPosition.AccountID)
}

func (retailDemandPosition RetailDemandPosition) GetAssortment() AssortmentPosition {
	return Deref(retailDemandPosition.Assortment)
}

func (retailDemandPosition RetailDemandPosition) GetCost() float64 {
	return Deref(retailDemandPosition.Cost)
}

func (retailDemandPosition RetailDemandPosition) GetDiscount() float64 {
	return Deref(retailDemandPosition.Discount)
}

func (retailDemandPosition RetailDemandPosition) GetID() uuid.UUID {
	return Deref(retailDemandPosition.ID)
}

func (retailDemandPosition RetailDemandPosition) GetPack() Pack {
	return Deref(retailDemandPosition.Pack)
}

func (retailDemandPosition RetailDemandPosition) GetPrice() float64 {
	return Deref(retailDemandPosition.Price)
}

func (retailDemandPosition RetailDemandPosition) GetQuantity() float64 {
	return Deref(retailDemandPosition.Quantity)
}

func (retailDemandPosition RetailDemandPosition) GetThings() Slice[string] {
	return retailDemandPosition.Things
}

func (retailDemandPosition RetailDemandPosition) GetVat() int {
	return Deref(retailDemandPosition.Vat)
}

func (retailDemandPosition RetailDemandPosition) GetVatEnabled() bool {
	return Deref(retailDemandPosition.VatEnabled)
}

func (retailDemandPosition RetailDemandPosition) GetStock() Stock {
	return Deref(retailDemandPosition.Stock)
}

func (retailDemandPosition *RetailDemandPosition) SetAssortment(assortment AsAssortment) *RetailDemandPosition {
	retailDemandPosition.Assortment = assortment.AsAssortment()
	return retailDemandPosition
}

func (retailDemandPosition *RetailDemandPosition) SetCost(cost float64) *RetailDemandPosition {
	retailDemandPosition.Cost = &cost
	return retailDemandPosition
}

func (retailDemandPosition *RetailDemandPosition) SetDiscount(discount float64) *RetailDemandPosition {
	retailDemandPosition.Discount = &discount
	return retailDemandPosition
}

func (retailDemandPosition *RetailDemandPosition) SetPack(pack *Pack) *RetailDemandPosition {
	retailDemandPosition.Pack = pack
	return retailDemandPosition
}

func (retailDemandPosition *RetailDemandPosition) SetPrice(price float64) *RetailDemandPosition {
	retailDemandPosition.Price = &price
	return retailDemandPosition
}

func (retailDemandPosition *RetailDemandPosition) SetQuantity(quantity float64) *RetailDemandPosition {
	retailDemandPosition.Quantity = &quantity
	return retailDemandPosition
}

func (retailDemandPosition *RetailDemandPosition) SetVat(vat int) *RetailDemandPosition {
	retailDemandPosition.Vat = &vat
	return retailDemandPosition
}

func (retailDemandPosition *RetailDemandPosition) SetVatEnabled(vatEnabled bool) *RetailDemandPosition {
	retailDemandPosition.VatEnabled = &vatEnabled
	return retailDemandPosition
}

func (retailDemandPosition *RetailDemandPosition) SetThings(things Slice[string]) *RetailDemandPosition {
	retailDemandPosition.Things = things
	return retailDemandPosition
}

func (retailDemandPosition RetailDemandPosition) String() string {
	return Stringify(retailDemandPosition)
}

func (retailDemandPosition RetailDemandPosition) MetaType() MetaType {
	return MetaTypeRetailDemandPosition
}

// RetailDemandTemplateArg
// Документ: Розничная продажа (retaildemand)
// Основание, на котором он может быть создан:
// - Розничная смена
// - Заказ покупателя
//type RetailDemandTemplateArg struct {
//	RetailShift   *MetaWrapper `json:"retailShift,omitempty"`
//	CustomerOrder *MetaWrapper `json:"customerOrder,omitempty"`
//}

// RetailDemandService
// Сервис для работы с розничными продажами.
type RetailDemandService interface {
	GetList(ctx context.Context, params *Params) (*List[RetailDemand], *resty.Response, error)
	Create(ctx context.Context, retailDemand *RetailDemand, params *Params) (*RetailDemand, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, retailDemandList []*RetailDemand, params *Params) (*[]RetailDemand, *resty.Response, error)
	DeleteMany(ctx context.Context, retailDemandList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params *Params) (*RetailDemand, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, retailDemand *RetailDemand, params *Params) (*RetailDemand, *resty.Response, error)
	//endpointTemplate[RetailDemand]
	//endpointTemplateBasedOn[RetailDemand, RetailDemandTemplateArg]
	GetMetadata(ctx context.Context) (*MetaAttributesSharedStatesWrapper, *resty.Response, error)
	GetPositions(ctx context.Context, id uuid.UUID, params *Params) (*MetaArray[RetailDemandPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params *Params) (*RetailDemandPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *RetailDemandPosition, params *Params) (*RetailDemandPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id uuid.UUID, position *RetailDemandPosition) (*RetailDemandPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id uuid.UUID, positions []*RetailDemandPosition) (*[]RetailDemandPosition, *resty.Response, error)
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
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*RetailDemand, *resty.Response, error)
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

func NewRetailDemandService(client *Client) RetailDemandService {
	e := NewEndpoint(client, "entity/retaildemand")
	return newMainService[RetailDemand, RetailDemandPosition, MetaAttributesSharedStatesWrapper, any](e)
}
