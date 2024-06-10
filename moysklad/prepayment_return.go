package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// PrepaymentReturn Возврат предоплаты.
// Ключевое слово: prepaymentreturn
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vozwrat-predoplaty
type PrepaymentReturn struct {
	Agent        *Counterparty                        `json:"agent,omitempty"`
	Organization *Organization                        `json:"organization,omitempty"`
	Applicable   *bool                                `json:"applicable,omitempty"`
	AccountID    *uuid.UUID                           `json:"accountId,omitempty"`
	CashSum      *float64                             `json:"cashSum,omitempty"`
	Code         *string                              `json:"code,omitempty"`
	Created      *Timestamp                           `json:"created,omitempty"`
	Deleted      *Timestamp                           `json:"deleted,omitempty"`
	Description  *string                              `json:"description,omitempty"`
	ExternalCode *string                              `json:"externalCode,omitempty"`
	Files        *MetaArray[File]                     `json:"files,omitempty"`
	Group        *Group                               `json:"group,omitempty"`
	ID           *uuid.UUID                           `json:"id,omitempty"`
	Meta         *Meta                                `json:"meta,omitempty"`
	Moment       *Timestamp                           `json:"moment,omitempty"`
	Name         *string                              `json:"name,omitempty"`
	NoCashSum    *float64                             `json:"noCashSum,omitempty"`
	Owner        *Employee                            `json:"owner,omitempty"`
	VatIncluded  *bool                                `json:"vatIncluded,omitempty"`
	Positions    *Positions[PrepaymentReturnPosition] `json:"positions,omitempty"`
	Prepayment   *Prepayment                          `json:"prepayment,omitempty"`
	Printed      *bool                                `json:"printed,omitempty"`
	Published    *bool                                `json:"published,omitempty"`
	QRSum        *float64                             `json:"qrSum,omitempty"`
	Rate         *Rate                                `json:"rate,omitempty"`
	RetailShift  *RetailShift                         `json:"retailShift,omitempty"`
	RetailStore  *RetailStore                         `json:"retailStore,omitempty"`
	Shared       *bool                                `json:"shared,omitempty"`
	State        *State                               `json:"state,omitempty"`
	Sum          *float64                             `json:"sum,omitempty"`
	SyncID       *uuid.UUID                           `json:"syncId,omitempty"`
	VatSum       *float64                             `json:"vatSum,omitempty"`
	Updated      *Timestamp                           `json:"updated,omitempty"`
	VatEnabled   *bool                                `json:"vatEnabled,omitempty"`
	TaxSystem    TaxSystem                            `json:"taxSystem,omitempty"`
	Attributes   Slice[AttributeValue]                `json:"attributes,omitempty"`
}

func (prepaymentReturn PrepaymentReturn) Clean() *PrepaymentReturn {
	return &PrepaymentReturn{Meta: prepaymentReturn.Meta}
}

func (prepaymentReturn PrepaymentReturn) GetAgent() Counterparty {
	return Deref(prepaymentReturn.Agent)
}

func (prepaymentReturn PrepaymentReturn) GetOrganization() Organization {
	return Deref(prepaymentReturn.Organization)
}

func (prepaymentReturn PrepaymentReturn) GetApplicable() bool {
	return Deref(prepaymentReturn.Applicable)
}

func (prepaymentReturn PrepaymentReturn) GetAccountID() uuid.UUID {
	return Deref(prepaymentReturn.AccountID)
}

func (prepaymentReturn PrepaymentReturn) GetCashSum() float64 {
	return Deref(prepaymentReturn.CashSum)
}

func (prepaymentReturn PrepaymentReturn) GetCode() string {
	return Deref(prepaymentReturn.Code)
}

func (prepaymentReturn PrepaymentReturn) GetCreated() Timestamp {
	return Deref(prepaymentReturn.Created)
}

func (prepaymentReturn PrepaymentReturn) GetDeleted() Timestamp {
	return Deref(prepaymentReturn.Deleted)
}

func (prepaymentReturn PrepaymentReturn) GetDescription() string {
	return Deref(prepaymentReturn.Description)
}

func (prepaymentReturn PrepaymentReturn) GetExternalCode() string {
	return Deref(prepaymentReturn.ExternalCode)
}

func (prepaymentReturn PrepaymentReturn) GetFiles() MetaArray[File] {
	return Deref(prepaymentReturn.Files)
}

func (prepaymentReturn PrepaymentReturn) GetGroup() Group {
	return Deref(prepaymentReturn.Group)
}

func (prepaymentReturn PrepaymentReturn) GetID() uuid.UUID {
	return Deref(prepaymentReturn.ID)
}

func (prepaymentReturn PrepaymentReturn) GetMeta() Meta {
	return Deref(prepaymentReturn.Meta)
}

func (prepaymentReturn PrepaymentReturn) GetMoment() Timestamp {
	return Deref(prepaymentReturn.Moment)
}

func (prepaymentReturn PrepaymentReturn) GetName() string {
	return Deref(prepaymentReturn.Name)
}

func (prepaymentReturn PrepaymentReturn) GetNoCashSum() float64 {
	return Deref(prepaymentReturn.NoCashSum)
}

func (prepaymentReturn PrepaymentReturn) GetOwner() Employee {
	return Deref(prepaymentReturn.Owner)
}

func (prepaymentReturn PrepaymentReturn) GetVatIncluded() bool {
	return Deref(prepaymentReturn.VatIncluded)
}

func (prepaymentReturn PrepaymentReturn) GetPositions() Positions[PrepaymentReturnPosition] {
	return Deref(prepaymentReturn.Positions)
}

func (prepaymentReturn PrepaymentReturn) GetPrepayment() Prepayment {
	return Deref(prepaymentReturn.Prepayment)
}

func (prepaymentReturn PrepaymentReturn) GetPrinted() bool {
	return Deref(prepaymentReturn.Printed)
}

func (prepaymentReturn PrepaymentReturn) GetPublished() bool {
	return Deref(prepaymentReturn.Published)
}

func (prepaymentReturn PrepaymentReturn) GetQRSum() float64 {
	return Deref(prepaymentReturn.QRSum)
}

func (prepaymentReturn PrepaymentReturn) GetRate() Rate {
	return Deref(prepaymentReturn.Rate)
}

func (prepaymentReturn PrepaymentReturn) GetRetailShift() RetailShift {
	return Deref(prepaymentReturn.RetailShift)
}

func (prepaymentReturn PrepaymentReturn) GetRetailStore() RetailStore {
	return Deref(prepaymentReturn.RetailStore)
}

func (prepaymentReturn PrepaymentReturn) GetShared() bool {
	return Deref(prepaymentReturn.Shared)
}

func (prepaymentReturn PrepaymentReturn) GetState() State {
	return Deref(prepaymentReturn.State)
}

func (prepaymentReturn PrepaymentReturn) GetSum() float64 {
	return Deref(prepaymentReturn.Sum)
}

func (prepaymentReturn PrepaymentReturn) GetSyncID() uuid.UUID {
	return Deref(prepaymentReturn.SyncID)
}

func (prepaymentReturn PrepaymentReturn) GetVatSum() float64 {
	return Deref(prepaymentReturn.VatSum)
}

func (prepaymentReturn PrepaymentReturn) GetUpdated() Timestamp {
	return Deref(prepaymentReturn.Updated)
}

func (prepaymentReturn PrepaymentReturn) GetVatEnabled() bool {
	return Deref(prepaymentReturn.VatEnabled)
}

func (prepaymentReturn PrepaymentReturn) GetTaxSystem() TaxSystem {
	return prepaymentReturn.TaxSystem
}

func (prepaymentReturn PrepaymentReturn) GetAttributes() Slice[AttributeValue] {
	return prepaymentReturn.Attributes
}

func (prepaymentReturn *PrepaymentReturn) SetAgent(agent *Counterparty) *PrepaymentReturn {
	prepaymentReturn.Agent = agent.Clean()
	return prepaymentReturn
}

func (prepaymentReturn *PrepaymentReturn) SetOrganization(organization *Organization) *PrepaymentReturn {
	prepaymentReturn.Organization = organization.Clean()
	return prepaymentReturn
}

func (prepaymentReturn *PrepaymentReturn) SetApplicable(applicable bool) *PrepaymentReturn {
	prepaymentReturn.Applicable = &applicable
	return prepaymentReturn
}

func (prepaymentReturn *PrepaymentReturn) SetCashSum(cashSum float64) *PrepaymentReturn {
	prepaymentReturn.CashSum = &cashSum
	return prepaymentReturn
}

func (prepaymentReturn *PrepaymentReturn) SetCode(code string) *PrepaymentReturn {
	prepaymentReturn.Code = &code
	return prepaymentReturn
}

func (prepaymentReturn *PrepaymentReturn) SetDescription(description string) *PrepaymentReturn {
	prepaymentReturn.Description = &description
	return prepaymentReturn
}

func (prepaymentReturn *PrepaymentReturn) SetExternalCode(externalCode string) *PrepaymentReturn {
	prepaymentReturn.ExternalCode = &externalCode
	return prepaymentReturn
}

func (prepaymentReturn *PrepaymentReturn) SetFiles(files Slice[File]) *PrepaymentReturn {
	prepaymentReturn.Files = NewMetaArrayRows(files)
	return prepaymentReturn
}

func (prepaymentReturn *PrepaymentReturn) SetGroup(group *Group) *PrepaymentReturn {
	prepaymentReturn.Group = group.Clean()
	return prepaymentReturn
}

func (prepaymentReturn *PrepaymentReturn) SetMeta(meta *Meta) *PrepaymentReturn {
	prepaymentReturn.Meta = meta
	return prepaymentReturn
}

func (prepaymentReturn *PrepaymentReturn) SetMoment(moment *Timestamp) *PrepaymentReturn {
	prepaymentReturn.Moment = moment
	return prepaymentReturn
}

func (prepaymentReturn *PrepaymentReturn) SetName(name string) *PrepaymentReturn {
	prepaymentReturn.Name = &name
	return prepaymentReturn
}

func (prepaymentReturn *PrepaymentReturn) SetNoCashSum(noCashSum float64) *PrepaymentReturn {
	prepaymentReturn.NoCashSum = &noCashSum
	return prepaymentReturn
}

func (prepaymentReturn *PrepaymentReturn) SetOwner(owner *Employee) *PrepaymentReturn {
	prepaymentReturn.Owner = owner.Clean()
	return prepaymentReturn
}

func (prepaymentReturn *PrepaymentReturn) SetVatIncluded(vatIncluded bool) *PrepaymentReturn {
	prepaymentReturn.VatIncluded = &vatIncluded
	return prepaymentReturn
}

func (prepaymentReturn *PrepaymentReturn) SetPositions(positions *Positions[PrepaymentReturnPosition]) *PrepaymentReturn {
	prepaymentReturn.Positions = positions
	return prepaymentReturn
}

func (prepaymentReturn *PrepaymentReturn) SetPrepayment(prepayment *Prepayment) *PrepaymentReturn {
	prepaymentReturn.Prepayment = prepayment.Clean()
	return prepaymentReturn
}

func (prepaymentReturn *PrepaymentReturn) SetQRSum(qrSum float64) *PrepaymentReturn {
	prepaymentReturn.QRSum = &qrSum
	return prepaymentReturn
}

func (prepaymentReturn *PrepaymentReturn) SetRate(rate *Rate) *PrepaymentReturn {
	prepaymentReturn.Rate = rate
	return prepaymentReturn
}

func (prepaymentReturn *PrepaymentReturn) SetRetailShift(retailShift *RetailShift) *PrepaymentReturn {
	prepaymentReturn.RetailShift = retailShift.Clean()
	return prepaymentReturn
}

func (prepaymentReturn *PrepaymentReturn) SetRetailStore(retailStore *RetailStore) *PrepaymentReturn {
	prepaymentReturn.RetailStore = retailStore.Clean()
	return prepaymentReturn
}

func (prepaymentReturn *PrepaymentReturn) SetShared(shared bool) *PrepaymentReturn {
	prepaymentReturn.Shared = &shared
	return prepaymentReturn
}

func (prepaymentReturn *PrepaymentReturn) SetState(state *State) *PrepaymentReturn {
	prepaymentReturn.State = state.Clean()
	return prepaymentReturn
}

func (prepaymentReturn *PrepaymentReturn) SetSyncID(syncID uuid.UUID) *PrepaymentReturn {
	prepaymentReturn.SyncID = &syncID
	return prepaymentReturn
}

func (prepaymentReturn *PrepaymentReturn) SetVatSum(vatSum float64) *PrepaymentReturn {
	prepaymentReturn.VatSum = &vatSum
	return prepaymentReturn
}

func (prepaymentReturn *PrepaymentReturn) SetVatEnabled(vatEnabled bool) *PrepaymentReturn {
	prepaymentReturn.VatEnabled = &vatEnabled
	return prepaymentReturn
}

func (prepaymentReturn *PrepaymentReturn) SetTaxSystem(taxSystem TaxSystem) *PrepaymentReturn {
	prepaymentReturn.TaxSystem = taxSystem
	return prepaymentReturn
}

func (prepaymentReturn *PrepaymentReturn) SetAttributes(attributes Slice[AttributeValue]) *PrepaymentReturn {
	prepaymentReturn.Attributes = attributes
	return prepaymentReturn
}

func (prepaymentReturn PrepaymentReturn) String() string {
	return Stringify(prepaymentReturn)
}

func (prepaymentReturn PrepaymentReturn) MetaType() MetaType {
	return MetaTypePrepaymentReturn
}

// PrepaymentReturnPosition Позиция Возврата предоплаты.
// Ключевое слово: prepaymentreturnposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vozwrat-predoplaty-atributy-suschnosti-pozicii-vozwrata-predoplaty
type PrepaymentReturnPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`
	Assortment *AssortmentPosition `json:"assortment,omitempty"`
	Discount   *float64            `json:"discount,omitempty"`
	ID         *uuid.UUID          `json:"id,omitempty"`
	Pack       *Pack               `json:"pack,omitempty"`
	Price      *float64            `json:"price,omitempty"`
	Quantity   *float64            `json:"quantity,omitempty"`
	Vat        *int                `json:"vat,omitempty"`
	VatEnabled *bool               `json:"vatEnabled,omitempty"`
}

func (prepaymentReturnPosition PrepaymentReturnPosition) GetAccountID() uuid.UUID {
	return Deref(prepaymentReturnPosition.AccountID)
}

func (prepaymentReturnPosition PrepaymentReturnPosition) GetAssortment() AssortmentPosition {
	return Deref(prepaymentReturnPosition.Assortment)
}

func (prepaymentReturnPosition PrepaymentReturnPosition) GetDiscount() float64 {
	return Deref(prepaymentReturnPosition.Discount)
}

func (prepaymentReturnPosition PrepaymentReturnPosition) GetID() uuid.UUID {
	return Deref(prepaymentReturnPosition.ID)
}

func (prepaymentReturnPosition PrepaymentReturnPosition) GetPack() Pack {
	return Deref(prepaymentReturnPosition.Pack)
}

func (prepaymentReturnPosition PrepaymentReturnPosition) GetPrice() float64 {
	return Deref(prepaymentReturnPosition.Price)
}

func (prepaymentReturnPosition PrepaymentReturnPosition) GetQuantity() float64 {
	return Deref(prepaymentReturnPosition.Quantity)
}

func (prepaymentReturnPosition PrepaymentReturnPosition) GetVat() int {
	return Deref(prepaymentReturnPosition.Vat)
}

func (prepaymentReturnPosition PrepaymentReturnPosition) GetVatEnabled() bool {
	return Deref(prepaymentReturnPosition.VatEnabled)
}

func (prepaymentReturnPosition *PrepaymentReturnPosition) SetAssortment(assortment AsAssortment) *PrepaymentReturnPosition {
	prepaymentReturnPosition.Assortment = assortment.AsAssortment()
	return prepaymentReturnPosition
}

func (prepaymentReturnPosition *PrepaymentReturnPosition) SetDiscount(discount float64) *PrepaymentReturnPosition {
	prepaymentReturnPosition.Discount = &discount
	return prepaymentReturnPosition
}

func (prepaymentReturnPosition *PrepaymentReturnPosition) SetPack(pack *Pack) *PrepaymentReturnPosition {
	prepaymentReturnPosition.Pack = pack
	return prepaymentReturnPosition
}

func (prepaymentReturnPosition *PrepaymentReturnPosition) SetPrice(price float64) *PrepaymentReturnPosition {
	prepaymentReturnPosition.Price = &price
	return prepaymentReturnPosition
}

func (prepaymentReturnPosition *PrepaymentReturnPosition) SetQuantity(quantity float64) *PrepaymentReturnPosition {
	prepaymentReturnPosition.Quantity = &quantity
	return prepaymentReturnPosition
}

func (prepaymentReturnPosition *PrepaymentReturnPosition) SetVat(vat int) *PrepaymentReturnPosition {
	prepaymentReturnPosition.Vat = &vat
	return prepaymentReturnPosition
}

func (prepaymentReturnPosition *PrepaymentReturnPosition) SetVatEnabled(vatEnabled bool) *PrepaymentReturnPosition {
	prepaymentReturnPosition.VatEnabled = &vatEnabled
	return prepaymentReturnPosition
}

func (prepaymentReturnPosition PrepaymentReturnPosition) String() string {
	return Stringify(prepaymentReturnPosition)
}

func (prepaymentReturnPosition PrepaymentReturnPosition) MetaType() MetaType {
	return MetaTypePrepaymentReturnPosition
}

// PrepaymentReturnService
// Сервис для работы с возвратами предоплат.
type PrepaymentReturnService interface {
	GetList(ctx context.Context, params *Params) (*List[PrepaymentReturn], *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params *Params) (*PrepaymentReturn, *resty.Response, error)
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)
	GetAttributeByID(ctx context.Context, id uuid.UUID) (*Attribute, *resty.Response, error)
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)
	CreateAttributes(ctx context.Context, attributeList []*Attribute) (*[]Attribute, *resty.Response, error)
	UpdateAttribute(ctx context.Context, id uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)
	DeleteAttribute(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	DeleteAttributes(ctx context.Context, attributeList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetaAttributesSharedStatesWrapper, *resty.Response, error)
	GetPositions(ctx context.Context, id uuid.UUID, params *Params) (*MetaArray[PrepaymentReturnPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params *Params) (*PrepaymentReturnPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *PrepaymentReturnPosition, params *Params) (*PrepaymentReturnPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id uuid.UUID, position *PrepaymentReturnPosition) (*PrepaymentReturnPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id uuid.UUID, positions []*PrepaymentReturnPosition) (*[]PrepaymentReturnPosition, *resty.Response, error)
	DeletePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (bool, *resty.Response, error)
	GetPositionTrackingCodes(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (*MetaArray[TrackingCode], *resty.Response, error)
	CreateOrUpdatePositionTrackingCodes(ctx context.Context, id uuid.UUID, positionID uuid.UUID, trackingCodes Slice[TrackingCode]) (*[]TrackingCode, *resty.Response, error)
	DeletePositionTrackingCodes(ctx context.Context, id uuid.UUID, positionID uuid.UUID, trackingCodes Slice[TrackingCode]) (*DeleteManyResponse, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*PrepaymentReturn, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error)
	MoveToTrash(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetStateByID(ctx context.Context, id uuid.UUID) (*State, *resty.Response, error)
	CreateState(ctx context.Context, state *State) (*State, *resty.Response, error)
	UpdateState(ctx context.Context, id uuid.UUID, state *State) (*State, *resty.Response, error)
	CreateOrUpdateStates(ctx context.Context, states []*State) (*[]State, *resty.Response, error)
	DeleteState(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
}

func NewPrepaymentReturnService(client *Client) PrepaymentReturnService {
	e := NewEndpoint(client, "entity/prepaymentreturn")
	return newMainService[PrepaymentReturn, PrepaymentReturnPosition, MetaAttributesSharedStatesWrapper, any](e)
}
