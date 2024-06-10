package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Prepayment Предоплата.
// Ключевое слово: prepayment
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-predoplata
type Prepayment struct {
	Returns       Slice[PrepaymentReturn]        `json:"returns,omitempty"`
	Owner         *Employee                      `json:"owner,omitempty"`
	Applicable    *bool                          `json:"applicable,omitempty"`
	Agent         *Counterparty                  `json:"agent,omitempty"`
	CashSum       *float64                       `json:"cashSum,omitempty"`
	Code          *string                        `json:"code,omitempty"`
	Created       *Timestamp                     `json:"created,omitempty"`
	CustomerOrder *CustomerOrder                 `json:"customerOrder,omitempty"`
	Deleted       *Timestamp                     `json:"deleted,omitempty"`
	Description   *string                        `json:"description,omitempty"`
	ExternalCode  *string                        `json:"externalCode,omitempty"`
	Files         *MetaArray[File]               `json:"files,omitempty"`
	Group         *Group                         `json:"group,omitempty"`
	ID            *uuid.UUID                     `json:"id,omitempty"`
	Meta          *Meta                          `json:"meta,omitempty"`
	Moment        *Timestamp                     `json:"moment,omitempty"`
	Name          *string                        `json:"name,omitempty"`
	NoCashSum     *float64                       `json:"noCashSum,omitempty"`
	AccountID     *uuid.UUID                     `json:"accountId,omitempty"`
	VatIncluded   *bool                          `json:"vatIncluded,omitempty"`
	Positions     *Positions[PrepaymentPosition] `json:"positions,omitempty"`
	Printed       *bool                          `json:"printed,omitempty"`
	Published     *bool                          `json:"published,omitempty"`
	QRSum         *float64                       `json:"qrSum,omitempty"`
	Rate          *Rate                          `json:"rate,omitempty"`
	RetailShift   *RetailShift                   `json:"retailShift,omitempty"`
	RetailStore   *RetailStore                   `json:"retailStore,omitempty"`
	Organization  *Organization                  `json:"organization,omitempty"`
	Shared        *bool                          `json:"shared,omitempty"`
	State         *State                         `json:"state,omitempty"`
	Sum           *float64                       `json:"sum,omitempty"`
	SyncID        *uuid.UUID                     `json:"syncId,omitempty"`
	VatSum        *float64                       `json:"vatSum,omitempty"`
	Updated       *Timestamp                     `json:"updated,omitempty"`
	VatEnabled    *bool                          `json:"vatEnabled,omitempty"`
	TaxSystem     TaxSystem                      `json:"taxSystem,omitempty"`
	Attributes    Slice[AttributeValue]          `json:"attributes,omitempty"`
}

func (prepayment Prepayment) GetReturns() Slice[PrepaymentReturn] {
	return prepayment.Returns
}

func (prepayment Prepayment) GetOwner() Employee {
	return Deref(prepayment.Owner)
}

func (prepayment Prepayment) GetApplicable() bool {
	return Deref(prepayment.Applicable)
}

func (prepayment Prepayment) GetAgent() Counterparty {
	return Deref(prepayment.Agent)
}

func (prepayment Prepayment) GetCashSum() float64 {
	return Deref(prepayment.CashSum)
}

func (prepayment Prepayment) GetCode() string {
	return Deref(prepayment.Code)
}

func (prepayment Prepayment) GetCreated() Timestamp {
	return Deref(prepayment.Created)
}

func (prepayment Prepayment) GetCustomerOrder() CustomerOrder {
	return Deref(prepayment.CustomerOrder)
}

func (prepayment Prepayment) GetDeleted() Timestamp {
	return Deref(prepayment.Deleted)
}

func (prepayment Prepayment) GetDescription() string {
	return Deref(prepayment.Description)
}

func (prepayment Prepayment) GetExternalCode() string {
	return Deref(prepayment.ExternalCode)
}

func (prepayment Prepayment) GetFiles() MetaArray[File] {
	return Deref(prepayment.Files)
}

func (prepayment Prepayment) GetGroup() Group {
	return Deref(prepayment.Group)
}

func (prepayment Prepayment) GetID() uuid.UUID {
	return Deref(prepayment.ID)
}

func (prepayment Prepayment) GetMeta() Meta {
	return Deref(prepayment.Meta)
}

func (prepayment Prepayment) GetMoment() Timestamp {
	return Deref(prepayment.Moment)
}

func (prepayment Prepayment) GetName() string {
	return Deref(prepayment.Name)
}

func (prepayment Prepayment) GetNoCashSum() float64 {
	return Deref(prepayment.NoCashSum)
}

func (prepayment Prepayment) GetAccountID() uuid.UUID {
	return Deref(prepayment.AccountID)
}

func (prepayment Prepayment) GetVatIncluded() bool {
	return Deref(prepayment.VatIncluded)
}

func (prepayment Prepayment) GetPositions() Positions[PrepaymentPosition] {
	return Deref(prepayment.Positions)
}

func (prepayment Prepayment) GetPrinted() bool {
	return Deref(prepayment.Printed)
}

func (prepayment Prepayment) GetPublished() bool {
	return Deref(prepayment.Published)
}

func (prepayment Prepayment) GetQRSum() float64 {
	return Deref(prepayment.QRSum)
}

func (prepayment Prepayment) GetRate() Rate {
	return Deref(prepayment.Rate)
}

func (prepayment Prepayment) GetRetailShift() RetailShift {
	return Deref(prepayment.RetailShift)
}

func (prepayment Prepayment) GetRetailStore() RetailStore {
	return Deref(prepayment.RetailStore)
}

func (prepayment Prepayment) GetOrganization() Organization {
	return Deref(prepayment.Organization)
}

func (prepayment Prepayment) GetShared() bool {
	return Deref(prepayment.Shared)
}

func (prepayment Prepayment) GetState() State {
	return Deref(prepayment.State)
}

func (prepayment Prepayment) GetSum() float64 {
	return Deref(prepayment.Sum)
}

func (prepayment Prepayment) GetSyncID() uuid.UUID {
	return Deref(prepayment.SyncID)
}

func (prepayment Prepayment) GetVatSum() float64 {
	return Deref(prepayment.VatSum)
}

func (prepayment Prepayment) GetUpdated() Timestamp {
	return Deref(prepayment.Updated)
}

func (prepayment Prepayment) GetVatEnabled() bool {
	return Deref(prepayment.VatEnabled)
}

func (prepayment Prepayment) GetTaxSystem() TaxSystem {
	return prepayment.TaxSystem
}

func (prepayment Prepayment) GetAttributes() Slice[AttributeValue] {
	return prepayment.Attributes
}

func (prepayment *Prepayment) SetReturns(returns Slice[PrepaymentReturn]) *Prepayment {
	prepayment.Returns = returns
	return prepayment
}

func (prepayment *Prepayment) SetOwner(owner *Employee) *Prepayment {
	prepayment.Owner = owner
	return prepayment
}

func (prepayment *Prepayment) SetApplicable(applicable bool) *Prepayment {
	prepayment.Applicable = &applicable
	return prepayment
}

func (prepayment *Prepayment) SetAgent(agent *Counterparty) *Prepayment {
	prepayment.Agent = agent
	return prepayment
}

func (prepayment *Prepayment) SetCashSum(cashSum float64) *Prepayment {
	prepayment.CashSum = &cashSum
	return prepayment
}

func (prepayment *Prepayment) SetCode(code string) *Prepayment {
	prepayment.Code = &code
	return prepayment
}

func (prepayment *Prepayment) SetCustomerOrder(customerOrder *CustomerOrder) *Prepayment {
	prepayment.CustomerOrder = customerOrder
	return prepayment
}

func (prepayment *Prepayment) SetDescription(description string) *Prepayment {
	prepayment.Description = &description
	return prepayment
}

func (prepayment *Prepayment) SetExternalCode(externalCode string) *Prepayment {
	prepayment.ExternalCode = &externalCode
	return prepayment
}

func (prepayment *Prepayment) SetFiles(files Slice[File]) *Prepayment {
	prepayment.Files = NewMetaArrayRows(files)
	return prepayment
}

func (prepayment *Prepayment) SetGroup(group *Group) *Prepayment {
	prepayment.Group = group
	return prepayment
}

func (prepayment *Prepayment) SetMeta(meta *Meta) *Prepayment {
	prepayment.Meta = meta
	return prepayment
}

func (prepayment *Prepayment) SetMoment(moment *Timestamp) *Prepayment {
	prepayment.Moment = moment
	return prepayment
}

func (prepayment *Prepayment) SetName(name string) *Prepayment {
	prepayment.Name = &name
	return prepayment
}

func (prepayment *Prepayment) SetNoCashSum(noCashSum float64) *Prepayment {
	prepayment.NoCashSum = &noCashSum
	return prepayment
}

func (prepayment *Prepayment) SetVatIncluded(vatIncluded bool) *Prepayment {
	prepayment.VatIncluded = &vatIncluded
	return prepayment
}

func (prepayment *Prepayment) SetPositions(positions *Positions[PrepaymentPosition]) *Prepayment {
	prepayment.Positions = positions
	return prepayment
}

func (prepayment *Prepayment) SetQRSum(qrSum float64) *Prepayment {
	prepayment.QRSum = &qrSum
	return prepayment
}

func (prepayment *Prepayment) SetRate(rate *Rate) *Prepayment {
	prepayment.Rate = rate
	return prepayment
}

func (prepayment *Prepayment) SetRetailShift(retailShift *RetailShift) *Prepayment {
	prepayment.RetailShift = retailShift
	return prepayment
}

func (prepayment *Prepayment) SetRetailStore(retailStore *RetailStore) *Prepayment {
	prepayment.RetailStore = retailStore
	return prepayment
}

func (prepayment *Prepayment) SetOrganization(organization *Organization) *Prepayment {
	prepayment.Organization = organization
	return prepayment
}

func (prepayment *Prepayment) SetShared(shared bool) *Prepayment {
	prepayment.Shared = &shared
	return prepayment
}

func (prepayment *Prepayment) SetState(state *State) *Prepayment {
	prepayment.State = state
	return prepayment
}

func (prepayment *Prepayment) SetSyncID(syncID uuid.UUID) *Prepayment {
	prepayment.SyncID = &syncID
	return prepayment
}

func (prepayment *Prepayment) SetVatSum(vatSum float64) *Prepayment {
	prepayment.VatSum = &vatSum
	return prepayment
}

func (prepayment *Prepayment) SetVatEnabled(vatEnabled bool) *Prepayment {
	prepayment.VatEnabled = &vatEnabled
	return prepayment
}

func (prepayment *Prepayment) SetTaxSystem(taxSystem TaxSystem) *Prepayment {
	prepayment.TaxSystem = taxSystem
	return prepayment
}

func (prepayment *Prepayment) SetAttributes(attributes Slice[AttributeValue]) *Prepayment {
	prepayment.Attributes = attributes
	return prepayment
}

func (prepayment Prepayment) String() string {
	return Stringify(prepayment)
}

func (prepayment Prepayment) MetaType() MetaType {
	return MetaTypePrepayment
}

// PrepaymentPosition Позиция Предоплаты.
// Ключевое слово: prepaymentposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-predoplata-predoplaty-pozicii-predoplaty
type PrepaymentPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	Discount   *float64            `json:"discount,omitempty"`   // Процент скидки или наценки. Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Price      *float64            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Vat        *int                `json:"vat,omitempty"`        // НДС, которым облагается текущая позиция
	VatEnabled *bool               `json:"vatEnabled,omitempty"` // Включен ли НДС для позиции. С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
}

func (prepaymentPosition PrepaymentPosition) GetAccountID() uuid.UUID {
	return Deref(prepaymentPosition.AccountID)
}

func (prepaymentPosition PrepaymentPosition) GetAssortment() AssortmentPosition {
	return Deref(prepaymentPosition.Assortment)
}

func (prepaymentPosition PrepaymentPosition) GetDiscount() float64 {
	return Deref(prepaymentPosition.Discount)
}

func (prepaymentPosition PrepaymentPosition) GetID() uuid.UUID {
	return Deref(prepaymentPosition.ID)
}

func (prepaymentPosition PrepaymentPosition) GetPack() Pack {
	return Deref(prepaymentPosition.Pack)
}

func (prepaymentPosition PrepaymentPosition) GetPrice() float64 {
	return Deref(prepaymentPosition.Price)
}

func (prepaymentPosition PrepaymentPosition) GetQuantity() float64 {
	return Deref(prepaymentPosition.Quantity)
}

func (prepaymentPosition PrepaymentPosition) GetVat() int {
	return Deref(prepaymentPosition.Vat)
}

func (prepaymentPosition PrepaymentPosition) GetVatEnabled() bool {
	return Deref(prepaymentPosition.VatEnabled)
}

func (prepaymentPosition *PrepaymentPosition) SetAssortment(assortment *AssortmentPosition) *PrepaymentPosition {
	prepaymentPosition.Assortment = assortment
	return prepaymentPosition
}

func (prepaymentPosition *PrepaymentPosition) SetDiscount(discount float64) *PrepaymentPosition {
	prepaymentPosition.Discount = &discount
	return prepaymentPosition
}

func (prepaymentPosition *PrepaymentPosition) SetPack(pack *Pack) *PrepaymentPosition {
	prepaymentPosition.Pack = pack
	return prepaymentPosition
}

func (prepaymentPosition *PrepaymentPosition) SetPrice(price float64) *PrepaymentPosition {
	prepaymentPosition.Price = &price
	return prepaymentPosition
}

func (prepaymentPosition *PrepaymentPosition) SetQuantity(quantity float64) *PrepaymentPosition {
	prepaymentPosition.Quantity = &quantity
	return prepaymentPosition
}

func (prepaymentPosition *PrepaymentPosition) SetVat(vat int) *PrepaymentPosition {
	prepaymentPosition.Vat = &vat
	return prepaymentPosition
}

func (prepaymentPosition *PrepaymentPosition) SetVatEnabled(vatEnabled bool) *PrepaymentPosition {
	prepaymentPosition.VatEnabled = &vatEnabled
	return prepaymentPosition
}

func (prepaymentPosition PrepaymentPosition) String() string {
	return Stringify(prepaymentPosition)
}

func (prepaymentPosition PrepaymentPosition) MetaType() MetaType {
	return MetaTypePrepaymentPosition
}

// PrepaymentService
// Сервис для работы с предоплатами.
type PrepaymentService interface {
	GetList(ctx context.Context, params *Params) (*List[Prepayment], *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params *Params) (*Prepayment, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetaAttributesSharedStatesWrapper, *resty.Response, error)
	GetPositions(ctx context.Context, id uuid.UUID, params *Params) (*MetaArray[PrepaymentPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params *Params) (*PrepaymentPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *PrepaymentPosition, params *Params) (*PrepaymentPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id uuid.UUID, position *PrepaymentPosition) (*PrepaymentPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id uuid.UUID, positions []*PrepaymentPosition) (*[]PrepaymentPosition, *resty.Response, error)
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
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*Prepayment, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error)
	MoveToTrash(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetStateByID(ctx context.Context, id uuid.UUID) (*State, *resty.Response, error)
	CreateState(ctx context.Context, state *State) (*State, *resty.Response, error)
	UpdateState(ctx context.Context, id uuid.UUID, state *State) (*State, *resty.Response, error)
	CreateOrUpdateStates(ctx context.Context, states []*State) (*[]State, *resty.Response, error)
	DeleteState(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
}

func NewPrepaymentService(client *Client) PrepaymentService {
	e := NewEndpoint(client, "entity/prepayment")
	return newMainService[Prepayment, PrepaymentPosition, MetaAttributesSharedStatesWrapper, any](e)
}
