package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Enter Оприходование.
// Ключевое слово: enter
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-oprihodowanie
type Enter struct {
	Organization *Organization             `json:"organization,omitempty"`
	Sum          *float64                  `json:"sum,omitempty"`
	Moment       *Timestamp                `json:"moment,omitempty"`
	Code         *string                   `json:"code,omitempty"`
	Created      *Timestamp                `json:"created,omitempty"`
	Deleted      *Timestamp                `json:"deleted,omitempty"`
	Description  *string                   `json:"description,omitempty"`
	ExternalCode *string                   `json:"externalCode,omitempty"`
	Files        *MetaArray[File]          `json:"files,omitempty"`
	Group        *Group                    `json:"group,omitempty"`
	ID           *uuid.UUID                `json:"id,omitempty"`
	Meta         *Meta                     `json:"meta,omitempty"`
	Updated      *Timestamp                `json:"updated,omitempty"`
	Applicable   *bool                     `json:"applicable,omitempty"`
	Printed      *bool                     `json:"printed,omitempty"`
	Overhead     *Overhead                 `json:"overhead,omitempty"`
	Owner        *Employee                 `json:"owner,omitempty"`
	Positions    *Positions[EnterPosition] `json:"positions,omitempty"`
	AccountID    *uuid.UUID                `json:"accountId,omitempty"`
	Project      *NullValue[Project]       `json:"project,omitempty"`
	Published    *bool                     `json:"published,omitempty"`
	Rate         *Rate                     `json:"rate,omitempty"`
	Shared       *bool                     `json:"shared,omitempty"`
	State        *NullValue[State]         `json:"state,omitempty"`
	Store        *Store                    `json:"store,omitempty"`
	Name         *string                   `json:"name,omitempty"`
	SyncID       *uuid.UUID                `json:"syncId,omitempty"`
	Attributes   Slice[AttributeValue]     `json:"attributes,omitempty"`
}

// Clean возвращает сущность с единственным заполненным полем Meta
func (enter Enter) Clean() *Enter {
	return &Enter{Meta: enter.Meta}
}

func (enter Enter) GetOrganization() Organization {
	return Deref(enter.Organization)
}

func (enter Enter) GetSum() float64 {
	return Deref(enter.Sum)
}

func (enter Enter) GetMoment() Timestamp {
	return Deref(enter.Moment)
}

func (enter Enter) GetCode() string {
	return Deref(enter.Code)
}

func (enter Enter) GetCreated() Timestamp {
	return Deref(enter.Created)
}

func (enter Enter) GetDeleted() Timestamp {
	return Deref(enter.Deleted)
}

func (enter Enter) GetDescription() string {
	return Deref(enter.Description)
}

func (enter Enter) GetExternalCode() string {
	return Deref(enter.ExternalCode)
}

func (enter Enter) GetFiles() MetaArray[File] {
	return Deref(enter.Files)
}

func (enter Enter) GetGroup() Group {
	return Deref(enter.Group)
}

func (enter Enter) GetID() uuid.UUID {
	return Deref(enter.ID)
}

func (enter Enter) GetMeta() Meta {
	return Deref(enter.Meta)
}

func (enter Enter) GetUpdated() Timestamp {
	return Deref(enter.Updated)
}

func (enter Enter) GetApplicable() bool {
	return Deref(enter.Applicable)
}

func (enter Enter) GetPrinted() bool {
	return Deref(enter.Printed)
}

func (enter Enter) GetOverhead() Overhead {
	return Deref(enter.Overhead)
}

func (enter Enter) GetOwner() Employee {
	return Deref(enter.Owner)
}

func (enter Enter) GetPositions() Positions[EnterPosition] {
	return Deref(enter.Positions)
}

func (enter Enter) GetAccountID() uuid.UUID {
	return Deref(enter.AccountID)
}

func (enter Enter) GetProject() Project {
	return enter.Project.Get()
}

func (enter Enter) GetPublished() bool {
	return Deref(enter.Published)
}

func (enter Enter) GetRate() Rate {
	return Deref(enter.Rate)
}

func (enter Enter) GetShared() bool {
	return Deref(enter.Shared)
}

func (enter Enter) GetState() State {
	return enter.State.Get()
}

func (enter Enter) GetStore() Store {
	return Deref(enter.Store)
}

func (enter Enter) GetName() string {
	return Deref(enter.Name)
}

func (enter Enter) GetSyncID() uuid.UUID {
	return Deref(enter.SyncID)
}

func (enter Enter) GetAttributes() Slice[AttributeValue] {
	return enter.Attributes
}

func (enter *Enter) SetOrganization(organization *Organization) *Enter {
	enter.Organization = organization.Clean()
	return enter
}

func (enter *Enter) SetMoment(moment *Timestamp) *Enter {
	enter.Moment = moment
	return enter
}

func (enter *Enter) SetCode(code string) *Enter {
	enter.Code = &code
	return enter
}

func (enter *Enter) SetDescription(description string) *Enter {
	enter.Description = &description
	return enter
}

func (enter *Enter) SetExternalCode(externalCode string) *Enter {
	enter.ExternalCode = &externalCode
	return enter
}

func (enter *Enter) SetFiles(files Slice[File]) *Enter {
	enter.Files = NewMetaArrayRows(files)
	return enter
}

func (enter *Enter) SetGroup(group *Group) *Enter {
	enter.Group = group.Clean()
	return enter
}

func (enter *Enter) SetMeta(meta *Meta) *Enter {
	enter.Meta = meta
	return enter
}

func (enter *Enter) SetApplicable(applicable bool) *Enter {
	enter.Applicable = &applicable
	return enter
}

func (enter *Enter) SetOverhead(overhead *Overhead) *Enter {
	enter.Overhead = overhead
	return enter
}

func (enter *Enter) SetOwner(owner *Employee) *Enter {
	enter.Owner = owner.Clean()
	return enter
}

func (enter *Enter) SetPositions(positions *Positions[EnterPosition]) *Enter {
	enter.Positions = positions
	return enter
}

func (enter *Enter) SetProject(project *Project) *Enter {
	enter.Project = NewNullValueWith(project.Clean())
	return enter
}

func (enter *Enter) SetNullProject() *Enter {
	enter.Project = NewNullValue[Project]()
	return enter
}

func (enter *Enter) SetRate(rate *Rate) *Enter {
	enter.Rate = rate
	return enter
}

func (enter *Enter) SetShared(shared bool) *Enter {
	enter.Shared = &shared
	return enter
}

func (enter *Enter) SetState(state *State) *Enter {
	enter.State = NewNullValueWith(state.Clean())
	return enter
}

func (enter *Enter) SetNullState() *Enter {
	enter.State = NewNullValue[State]()
	return enter
}

func (enter *Enter) SetStore(store *Store) *Enter {
	enter.Store = store.Clean()
	return enter
}

func (enter *Enter) SetName(name string) *Enter {
	enter.Name = &name
	return enter
}

func (enter *Enter) SetSyncID(syncID uuid.UUID) *Enter {
	enter.SyncID = &syncID
	return enter
}

func (enter *Enter) SetAttributes(attributes Slice[AttributeValue]) *Enter {
	enter.Attributes = attributes
	return enter
}

func (enter Enter) String() string {
	return Stringify(enter)
}

func (enter Enter) MetaType() MetaType {
	return MetaTypeEnter
}

// EnterPosition Позиция оприходования
// Ключевое слово: enterposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-oprihodowanie-oprihodowaniq-pozicii-oprihodowaniq
type EnterPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	Country    *Country            `json:"country,omitempty"`    // Метаданные страны
	GTD        *GTD                `json:"gtd,omitempty"`        // ГТД
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Overhead   *float64            `json:"overhead,omitempty"`   // Накладные расходы. Если Позиции Оприходования не заданы, то накладные расходы нельзя задать
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Price      *float64            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Reason     *string             `json:"reason,omitempty"`     // Причина оприходования данной позиции
	Slot       *Slot               `json:"slot,omitempty"`       // Ячейка на складе
	Things     Slice[string]       `json:"things,omitempty"`     // Серийные номера. Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете. В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута.
}

func (enterPosition EnterPosition) GetAccountID() uuid.UUID {
	return Deref(enterPosition.AccountID)
}

func (enterPosition EnterPosition) GetAssortment() AssortmentPosition {
	return Deref(enterPosition.Assortment)
}

func (enterPosition EnterPosition) GetCountry() Country {
	return Deref(enterPosition.Country)
}

func (enterPosition EnterPosition) GetGTD() GTD {
	return Deref(enterPosition.GTD)
}

func (enterPosition EnterPosition) GetID() uuid.UUID {
	return Deref(enterPosition.ID)
}

func (enterPosition EnterPosition) GetOverhead() float64 {
	return Deref(enterPosition.Overhead)
}

func (enterPosition EnterPosition) GetPack() Pack {
	return Deref(enterPosition.Pack)
}

func (enterPosition EnterPosition) GetPrice() float64 {
	return Deref(enterPosition.Price)
}

func (enterPosition EnterPosition) GetQuantity() float64 {
	return Deref(enterPosition.Quantity)
}

func (enterPosition EnterPosition) GetReason() string {
	return Deref(enterPosition.Reason)
}

func (enterPosition EnterPosition) GetSlot() Slot {
	return Deref(enterPosition.Slot)
}

func (enterPosition EnterPosition) GetThings() Slice[string] {
	return enterPosition.Things
}

func (enterPosition *EnterPosition) SetAssortment(assortment AsAssortment) *EnterPosition {
	enterPosition.Assortment = assortment.AsAssortment()
	return enterPosition
}

func (enterPosition *EnterPosition) SetCountry(country *Country) *EnterPosition {
	enterPosition.Country = country.Clean()
	return enterPosition
}

func (enterPosition *EnterPosition) SetGTD(gtd *GTD) *EnterPosition {
	enterPosition.GTD = gtd
	return enterPosition
}

func (enterPosition *EnterPosition) SetPack(pack *Pack) *EnterPosition {
	enterPosition.Pack = pack
	return enterPosition
}

func (enterPosition *EnterPosition) SetPrice(price float64) *EnterPosition {
	enterPosition.Price = &price
	return enterPosition
}

func (enterPosition *EnterPosition) SetQuantity(quantity float64) *EnterPosition {
	enterPosition.Quantity = &quantity
	return enterPosition
}

func (enterPosition *EnterPosition) SetReason(reason string) *EnterPosition {
	enterPosition.Reason = &reason
	return enterPosition
}

func (enterPosition *EnterPosition) SetSlot(slot *Slot) *EnterPosition {
	enterPosition.Slot = slot.Clean()
	return enterPosition
}

func (enterPosition *EnterPosition) SetThings(things Slice[string]) *EnterPosition {
	enterPosition.Things = things
	return enterPosition
}

func (enterPosition EnterPosition) String() string {
	return Stringify(enterPosition)
}

func (enterPosition EnterPosition) MetaType() MetaType {
	return MetaTypeEnterPosition
}

// EnterService
// Сервис для работы с оприходованиями.
type EnterService interface {
	GetList(ctx context.Context, params ...*Params) (*List[Enter], *resty.Response, error)
	Create(ctx context.Context, enter *Enter, params ...*Params) (*Enter, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, enterList Slice[Enter], params ...*Params) (*Slice[Enter], *resty.Response, error)
	DeleteMany(ctx context.Context, enterList []MetaWrapper) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Enter, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, enter *Enter, params ...*Params) (*Enter, *resty.Response, error)
	Template(ctx context.Context) (*Enter, *resty.Response, error)
	TemplateBased(ctx context.Context, basedOn ...MetaOwner) (*Enter, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetaAttributesSharedStatesWrapper, *resty.Response, error)
	GetPositions(ctx context.Context, id uuid.UUID, params ...*Params) (*MetaArray[EnterPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params ...*Params) (*EnterPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *EnterPosition, params ...*Params) (*EnterPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id uuid.UUID, position *EnterPosition) (*EnterPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id uuid.UUID, positions Slice[EnterPosition]) (*Slice[EnterPosition], *resty.Response, error)
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
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*Enter, *resty.Response, error)
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

func NewEnterService(client *Client) EnterService {
	e := NewEndpoint(client, "entity/enter")
	return newMainService[Enter, EnterPosition, MetaAttributesSharedStatesWrapper, any](e)
}
