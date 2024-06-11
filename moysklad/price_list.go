package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// PriceList Прайс-лист.
// Ключевое слово: pricelist
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-prajs-list
type PriceList struct {
	Meta         *Meta                         `json:"meta,omitempty"`
	Columns      Slice[PriceListColumn]        `json:"columns,omitempty"`
	AccountID    *uuid.UUID                    `json:"accountId,omitempty"`
	Code         *string                       `json:"code,omitempty"`
	Moment       *Timestamp                    `json:"moment,omitempty"`
	Created      *Timestamp                    `json:"created,omitempty"`
	Deleted      *Timestamp                    `json:"deleted,omitempty"`
	Name         *string                       `json:"name,omitempty"`
	ExternalCode *string                       `json:"externalCode,omitempty"`
	Files        *MetaArray[File]              `json:"files,omitempty"`
	Group        *Group                        `json:"group,omitempty"`
	ID           *uuid.UUID                    `json:"id,omitempty"`
	Updated      *Timestamp                    `json:"updated,omitempty"`
	Applicable   *bool                         `json:"applicable,omitempty"`
	Description  *string                       `json:"description,omitempty"`
	Organization *Organization                 `json:"organization,omitempty"`
	Owner        *Employee                     `json:"owner,omitempty"`
	Positions    *Positions[PriceListPosition] `json:"positions,omitempty"`
	PriceType    *PriceType                    `json:"priceType,omitempty"`
	Printed      *bool                         `json:"printed,omitempty"`
	Published    *bool                         `json:"published,omitempty"`
	Shared       *bool                         `json:"shared,omitempty"`
	State        *State                        `json:"state,omitempty"`
	SyncID       *uuid.UUID                    `json:"syncId,omitempty"`
	Attributes   Slice[AttributeValue]         `json:"attributes,omitempty"`
}

func (priceList PriceList) Clean() *PriceList {
	return &PriceList{Meta: priceList.Meta}
}

func (priceList PriceList) GetMeta() Meta {
	return Deref(priceList.Meta)
}

func (priceList PriceList) GetColumns() Slice[PriceListColumn] {
	return priceList.Columns
}

func (priceList PriceList) GetAccountID() uuid.UUID {
	return Deref(priceList.AccountID)
}

func (priceList PriceList) GetCode() string {
	return Deref(priceList.Code)
}

func (priceList PriceList) GetMoment() Timestamp {
	return Deref(priceList.Moment)
}

func (priceList PriceList) GetCreated() Timestamp {
	return Deref(priceList.Created)
}

func (priceList PriceList) GetDeleted() Timestamp {
	return Deref(priceList.Deleted)
}

func (priceList PriceList) GetName() string {
	return Deref(priceList.Name)
}

func (priceList PriceList) GetExternalCode() string {
	return Deref(priceList.ExternalCode)
}

func (priceList PriceList) GetFiles() MetaArray[File] {
	return Deref(priceList.Files)
}

func (priceList PriceList) GetGroup() Group {
	return Deref(priceList.Group)
}

func (priceList PriceList) GetID() uuid.UUID {
	return Deref(priceList.ID)
}

func (priceList PriceList) GetUpdated() Timestamp {
	return Deref(priceList.Updated)
}

func (priceList PriceList) GetApplicable() bool {
	return Deref(priceList.Applicable)
}

func (priceList PriceList) GetDescription() string {
	return Deref(priceList.Description)
}

func (priceList PriceList) GetOrganization() Organization {
	return Deref(priceList.Organization)
}

func (priceList PriceList) GetOwner() Employee {
	return Deref(priceList.Owner)
}

func (priceList PriceList) GetPositions() Positions[PriceListPosition] {
	return Deref(priceList.Positions)
}

func (priceList PriceList) GetPriceType() PriceType {
	return Deref(priceList.PriceType)
}

func (priceList PriceList) GetPrinted() bool {
	return Deref(priceList.Printed)
}

func (priceList PriceList) GetPublished() bool {
	return Deref(priceList.Published)
}

func (priceList PriceList) GetShared() bool {
	return Deref(priceList.Shared)
}

func (priceList PriceList) GetState() State {
	return Deref(priceList.State)
}

func (priceList PriceList) GetSyncID() uuid.UUID {
	return Deref(priceList.SyncID)
}

func (priceList PriceList) GetAttributes() Slice[AttributeValue] {
	return priceList.Attributes
}

func (priceList *PriceList) SetMeta(meta *Meta) *PriceList {
	priceList.Meta = meta
	return priceList
}

func (priceList *PriceList) SetColumns(columns Slice[PriceListColumn]) *PriceList {
	priceList.Columns = columns
	return priceList
}

func (priceList *PriceList) SetCode(code string) *PriceList {
	priceList.Code = &code
	return priceList
}

func (priceList *PriceList) SetMoment(moment *Timestamp) *PriceList {
	priceList.Moment = moment
	return priceList
}

func (priceList *PriceList) SetName(name string) *PriceList {
	priceList.Name = &name
	return priceList
}

func (priceList *PriceList) SetExternalCode(externalCode string) *PriceList {
	priceList.ExternalCode = &externalCode
	return priceList
}

func (priceList *PriceList) SetFiles(files Slice[File]) *PriceList {
	priceList.Files = NewMetaArrayRows(files)
	return priceList
}

func (priceList *PriceList) SetGroup(group *Group) *PriceList {
	priceList.Group = group.Clean()
	return priceList
}

func (priceList *PriceList) SetApplicable(applicable bool) *PriceList {
	priceList.Applicable = &applicable
	return priceList
}

func (priceList *PriceList) SetDescription(description string) *PriceList {
	priceList.Description = &description
	return priceList
}

func (priceList *PriceList) SetOrganization(organization *Organization) *PriceList {
	priceList.Organization = organization
	return priceList
}

func (priceList *PriceList) SetOwner(owner *Employee) *PriceList {
	priceList.Owner = owner.Clean()
	return priceList
}

func (priceList *PriceList) SetPositions(positions *Positions[PriceListPosition]) *PriceList {
	priceList.Positions = positions
	return priceList
}

func (priceList *PriceList) SetShared(shared bool) *PriceList {
	priceList.Shared = &shared
	return priceList
}

func (priceList *PriceList) SetState(state *State) *PriceList {
	priceList.State = state.Clean()
	return priceList
}

func (priceList *PriceList) SetSyncID(syncID uuid.UUID) *PriceList {
	priceList.SyncID = &syncID
	return priceList
}

func (priceList *PriceList) SetAttributes(attributes Slice[AttributeValue]) *PriceList {
	priceList.Attributes = attributes
	return priceList
}

func (priceList PriceList) String() string {
	return Stringify(priceList)
}

func (priceList PriceList) MetaType() MetaType {
	return MetaTypePriceList
}

// PriceListCell Ячейка прайс листа.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-prajs-list-prajs-listy-yachejki
type PriceListCell struct {
	Column *string  `json:"column,omitempty"` // Название столбца, к которому относится данная ячейка
	Sum    *float64 `json:"sum,omitempty"`    // Числовое значение ячейки
}

func (priceListCell PriceListCell) GetColumn() string {
	return Deref(priceListCell.Column)
}

func (priceListCell PriceListCell) GetSum() float64 {
	return Deref(priceListCell.Sum)
}

func (priceListCell *PriceListCell) SetColumn(column string) *PriceListCell {
	priceListCell.Column = &column
	return priceListCell
}

func (priceListCell *PriceListCell) SetSum(sum float64) *PriceListCell {
	priceListCell.Sum = &sum
	return priceListCell
}

func (priceListCell PriceListCell) String() string {
	return Stringify(priceListCell)
}

// PriceListColumn Столбец прайс листа.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-prajs-list-prajs-listy-stolbcy
type PriceListColumn struct {
	Name               *string  `json:"name,omitempty"`               // Название столбца
	PercentageDiscount *float64 `json:"percentageDiscount,omitempty"` // Процентная наценка или скидка по умолчанию для столбца
}

func (priceListColumn PriceListColumn) GetName() string {
	return Deref(priceListColumn.Name)
}

func (priceListColumn PriceListColumn) GetPercentageDiscount() float64 {
	return Deref(priceListColumn.PercentageDiscount)
}

func (priceListColumn *PriceListColumn) SetName(name string) *PriceListColumn {
	priceListColumn.Name = &name
	return priceListColumn
}

func (priceListColumn *PriceListColumn) SetPercentageDiscount(percentageDiscount float64) *PriceListColumn {
	priceListColumn.PercentageDiscount = &percentageDiscount
	return priceListColumn
}

func (priceListColumn PriceListColumn) String() string {
	return Stringify(priceListColumn)
}

// PriceListPosition Позиция прайс листа.
// Ключевое слово: pricelistrow
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-prajs-list-prajs-listy-pozicii-prajs-lista
type PriceListPosition struct {
	AccountID  *uuid.UUID           `json:"accountId,omitempty"`
	Assortment *AssortmentPosition  `json:"assortment,omitempty"`
	ID         *uuid.UUID           `json:"id,omitempty"`
	Pack       *Pack                `json:"pack,omitempty"`
	Cells      Slice[PriceListCell] `json:"cells,omitempty"`
}

func (priceListPosition PriceListPosition) GetAccountID() uuid.UUID {
	return Deref(priceListPosition.AccountID)
}

func (priceListPosition PriceListPosition) GetAssortment() AssortmentPosition {
	return Deref(priceListPosition.Assortment)
}

func (priceListPosition PriceListPosition) GetCells() Slice[PriceListCell] {
	return priceListPosition.Cells
}

func (priceListPosition PriceListPosition) GetID() uuid.UUID {
	return Deref(priceListPosition.ID)
}

func (priceListPosition PriceListPosition) GetPack() Pack {
	return Deref(priceListPosition.Pack)
}

func (priceListPosition *PriceListPosition) SetAssortment(assortment AsAssortment) *PriceListPosition {
	priceListPosition.Assortment = assortment.AsAssortment()
	return priceListPosition
}

func (priceListPosition *PriceListPosition) SetCells(cells Slice[PriceListCell]) *PriceListPosition {
	priceListPosition.Cells = cells
	return priceListPosition
}

func (priceListPosition *PriceListPosition) SetPack(pack *Pack) *PriceListPosition {
	priceListPosition.Pack = pack
	return priceListPosition
}

func (priceListPosition PriceListPosition) String() string {
	return Stringify(priceListPosition)
}

func (priceListPosition PriceListPosition) MetaType() MetaType {
	return MetaTypePriceListPosition
}

// PriceListService
// Сервис для работы с прайс-листами.
type PriceListService interface {
	GetList(ctx context.Context, params *Params) (*List[PriceList], *resty.Response, error)
	Create(ctx context.Context, priceList *PriceList, params *Params) (*PriceList, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, priceListList Slice[PriceList], params *Params) (*Slice[PriceList], *resty.Response, error)
	DeleteMany(ctx context.Context, priceListList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params *Params) (*PriceList, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, priceList *PriceList, params *Params) (*PriceList, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetaAttributesSharedStatesWrapper, *resty.Response, error)
	GetPositions(ctx context.Context, id uuid.UUID, params *Params) (*MetaArray[PriceListPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params *Params) (*PriceListPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *PriceListPosition, params *Params) (*PriceListPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id uuid.UUID, position *PriceListPosition) (*PriceListPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id uuid.UUID, positions Slice[PriceListPosition]) (*Slice[PriceListPosition], *resty.Response, error)
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
	DeleteAttributes(ctx context.Context, attributeList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*PriceList, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error)
	MoveToTrash(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetStateByID(ctx context.Context, id uuid.UUID) (*State, *resty.Response, error)
	CreateState(ctx context.Context, state *State) (*State, *resty.Response, error)
	UpdateState(ctx context.Context, id uuid.UUID, state *State) (*State, *resty.Response, error)
	CreateOrUpdateStates(ctx context.Context, states Slice[State]) (*Slice[State], *resty.Response, error)
	DeleteState(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
}

func NewPriceListService(client *Client) PriceListService {
	e := NewEndpoint(client, "entity/pricelist")
	return newMainService[PriceList, PriceListPosition, MetaAttributesSharedStatesWrapper, any](e)
}
