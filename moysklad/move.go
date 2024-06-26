package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Move Перемещение.
// Ключевое слово: move
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-peremeschenie
type Move struct {
	Moment        *Timestamp                `json:"moment,omitempty"`
	Updated       *Timestamp                `json:"updated,omitempty"`
	AccountID     *uuid.UUID                `json:"accountId,omitempty"`
	Code          *string                   `json:"code,omitempty"`
	Created       *Timestamp                `json:"created,omitempty"`
	Deleted       *Timestamp                `json:"deleted,omitempty"`
	Demand        *Demand                   `json:"demand,omitempty"`
	Description   *string                   `json:"description,omitempty"`
	ExternalCode  *string                   `json:"externalCode,omitempty"`
	Files         *MetaArray[File]          `json:"files,omitempty"`
	Group         *Group                    `json:"group,omitempty"`
	ID            *uuid.UUID                `json:"id,omitempty"`
	InternalOrder *NullValue[InternalOrder] `json:"internalOrder,omitempty"`
	CustomerOrder *NullValue[CustomerOrder] `json:"customerOrder,omitempty"`
	Meta          *Meta                     `json:"meta,omitempty"`
	Name          *string                   `json:"name,omitempty"`
	Organization  *Organization             `json:"organization,omitempty"`
	Applicable    *bool                     `json:"applicable,omitempty"`
	Overhead      *Overhead                 `json:"overhead,omitempty"`
	Owner         *Employee                 `json:"owner,omitempty"`
	Positions     *Positions[MovePosition]  `json:"positions,omitempty"`
	Printed       *bool                     `json:"printed,omitempty"`
	Project       *NullValue[Project]       `json:"project,omitempty"`
	Published     *bool                     `json:"published,omitempty"`
	Rate          *NullValue[Rate]          `json:"rate,omitempty"`
	Shared        *bool                     `json:"shared,omitempty"`
	SourceStore   *Store                    `json:"sourceStore,omitempty"`
	State         *NullValue[State]         `json:"state,omitempty"`
	Sum           *float64                  `json:"sum,omitempty"`
	SyncID        *uuid.UUID                `json:"syncId,omitempty"`
	Supply        *Supply                   `json:"supply,omitempty"`
	TargetStore   *Store                    `json:"targetStore,omitempty"`
	Attributes    Slice[Attribute]          `json:"attributes,omitempty"`
}

// Clean возвращает сущность с единственным заполненным полем Meta
func (move Move) Clean() *Move {
	return &Move{Meta: move.Meta}
}

// AsTaskOperation реализует интерфейс AsTaskOperationInterface
func (move Move) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: move.Meta}
}

func (move Move) GetMoment() Timestamp {
	return Deref(move.Moment)
}

func (move Move) GetUpdated() Timestamp {
	return Deref(move.Updated)
}

func (move Move) GetAccountID() uuid.UUID {
	return Deref(move.AccountID)
}

func (move Move) GetCode() string {
	return Deref(move.Code)
}

func (move Move) GetCreated() Timestamp {
	return Deref(move.Created)
}

func (move Move) GetDeleted() Timestamp {
	return Deref(move.Deleted)
}

func (move Move) GetDescription() string {
	return Deref(move.Description)
}

func (move Move) GetExternalCode() string {
	return Deref(move.ExternalCode)
}

func (move Move) GetFiles() MetaArray[File] {
	return Deref(move.Files)
}

func (move Move) GetGroup() Group {
	return Deref(move.Group)
}

func (move Move) GetID() uuid.UUID {
	return Deref(move.ID)
}

func (move Move) GetInternalOrder() InternalOrder {
	return move.InternalOrder.Get()
}

func (move Move) GetCustomerOrder() CustomerOrder {
	return move.CustomerOrder.Get()
}

func (move Move) GetMeta() Meta {
	return Deref(move.Meta)
}

func (move Move) GetName() string {
	return Deref(move.Name)
}

func (move Move) GetOrganization() Organization {
	return Deref(move.Organization)
}

func (move Move) GetApplicable() bool {
	return Deref(move.Applicable)
}

func (move Move) GetOverhead() Overhead {
	return Deref(move.Overhead)
}

func (move Move) GetOwner() Employee {
	return Deref(move.Owner)
}

func (move Move) GetPositions() Positions[MovePosition] {
	return Deref(move.Positions)
}

func (move Move) GetPrinted() bool {
	return Deref(move.Printed)
}

func (move Move) GetProject() Project {
	return move.Project.Get()
}

func (move Move) GetPublished() bool {
	return Deref(move.Published)
}

func (move Move) GetRate() Rate {
	return move.Rate.Get()
}

func (move Move) GetShared() bool {
	return Deref(move.Shared)
}

func (move Move) GetSourceStore() Store {
	return Deref(move.SourceStore)
}

func (move Move) GetState() State {
	return move.State.Get()
}

func (move Move) GetSum() float64 {
	return Deref(move.Sum)
}

func (move Move) GetSyncID() uuid.UUID {
	return Deref(move.SyncID)
}

func (move Move) GetTargetStore() Store {
	return Deref(move.TargetStore)
}

func (move Move) GetDemand() Demand {
	return Deref(move.Demand)
}

func (move Move) GetSupply() Supply {
	return Deref(move.Supply)
}

func (move Move) GetAttributes() Slice[Attribute] {
	return move.Attributes
}

func (move *Move) SetMoment(moment *Timestamp) *Move {
	move.Moment = moment
	return move
}

func (move *Move) SetCode(code string) *Move {
	move.Code = &code
	return move
}

func (move *Move) SetDescription(description string) *Move {
	move.Description = &description
	return move
}

func (move *Move) SetExternalCode(externalCode string) *Move {
	move.ExternalCode = &externalCode
	return move
}

func (move *Move) SetFiles(files ...*File) *Move {
	move.Files = NewMetaArrayFrom(files)
	return move
}

func (move *Move) SetGroup(group *Group) *Move {
	move.Group = group.Clean()
	return move
}

func (move *Move) SetInternalOrder(internalOrder *InternalOrder) *Move {
	move.InternalOrder = NewNullValueFrom(internalOrder.Clean())
	return move
}

func (move *Move) SetNullInternalOrder() *Move {
	move.InternalOrder = NewNullValue[InternalOrder]()
	return move
}

func (move *Move) SetCustomerOrder(customerOrder *CustomerOrder) *Move {
	move.CustomerOrder = NewNullValueFrom(customerOrder.Clean())
	return move
}

func (move *Move) SetNullCustomerOrder() *Move {
	move.CustomerOrder = NewNullValue[CustomerOrder]()
	return move
}

func (move *Move) SetMeta(meta *Meta) *Move {
	move.Meta = meta
	return move
}

func (move *Move) SetName(name string) *Move {
	move.Name = &name
	return move
}

func (move *Move) SetOrganization(organization *Organization) *Move {
	move.Organization = organization.Clean()
	return move
}

func (move *Move) SetApplicable(applicable bool) *Move {
	move.Applicable = &applicable
	return move
}

func (move *Move) SetOverhead(overhead *Overhead) *Move {
	move.Overhead = overhead
	return move
}

func (move *Move) SetOwner(owner *Employee) *Move {
	move.Owner = owner.Clean()
	return move
}

func (move *Move) SetPositions(positions ...*MovePosition) *Move {
	move.Positions = NewPositionsFrom(positions)
	return move
}

func (move *Move) SetProject(project *Project) *Move {
	move.Project = NewNullValueFrom(project.Clean())
	return move
}

func (move *Move) SetNullProject() *Move {
	move.Project = NewNullValue[Project]()
	return move
}

func (move *Move) SetRate(rate *Rate) *Move {
	move.Rate = NewNullValueFrom(rate)
	return move
}

func (move *Move) SetNullRate() *Move {
	move.Rate = NewNullValue[Rate]()
	return move
}

func (move *Move) SetShared(shared bool) *Move {
	move.Shared = &shared
	return move
}

func (move *Move) SetSourceStore(sourceStore *Store) *Move {
	move.SourceStore = sourceStore
	return move
}

func (move *Move) SetState(state *State) *Move {
	move.State = NewNullValueFrom(state.Clean())
	return move
}

func (move *Move) SetNullState() *Move {
	move.State = NewNullValue[State]()
	return move
}

func (move *Move) SetSyncID(syncID uuid.UUID) *Move {
	move.SyncID = &syncID
	return move
}

func (move *Move) SetTargetStore(targetStore *Store) *Move {
	move.TargetStore = targetStore.Clean()
	return move
}

func (move *Move) SetAttributes(attributes ...*Attribute) *Move {
	move.Attributes = attributes
	return move
}

func (move Move) String() string {
	return Stringify(move)
}

// MetaType возвращает тип сущности.
func (Move) MetaType() MetaType {
	return MetaTypeMove
}

// Update shortcut
func (move Move) Update(ctx context.Context, client *Client, params ...*Params) (*Move, *resty.Response, error) {
	return client.Entity().Move().Update(ctx, move.GetID(), &move, params...)
}

// Create shortcut
func (move Move) Create(ctx context.Context, client *Client, params ...*Params) (*Move, *resty.Response, error) {
	return client.Entity().Move().Create(ctx, &move, params...)
}

// Delete shortcut
func (move Move) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return client.Entity().Move().Delete(ctx, move.GetID())
}

// MovePosition Позиция перемещения.
// Ключевое слово: moveposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-peremeschenie-peremescheniq-pozicii-peremescheniq
type MovePosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Overhead   *float64            `json:"overhead,omitempty"`   // Накладные расходы. Если Позиции Перемещения не заданы, то накладные расходы нельзя задать
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Price      *float64            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе
	SourceSlot *Slot               `json:"sourceSlot,omitempty"` // Ячейка на складе, с которого совершается перемещение
	TargetSlot *Slot               `json:"targetSlot,omitempty"` // Ячейка на складе, на который совершается перемещение
	Things     Slice[string]       `json:"things,omitempty"`     // Серийные номера. Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете. В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута
}

func (movePosition MovePosition) GetAccountID() uuid.UUID {
	return Deref(movePosition.AccountID)
}

func (movePosition MovePosition) GetAssortment() AssortmentPosition {
	return Deref(movePosition.Assortment)
}

func (movePosition MovePosition) GetID() uuid.UUID {
	return Deref(movePosition.ID)
}

func (movePosition MovePosition) GetOverhead() float64 {
	return Deref(movePosition.Overhead)
}

func (movePosition MovePosition) GetPack() Pack {
	return Deref(movePosition.Pack)
}

func (movePosition MovePosition) GetPrice() float64 {
	return Deref(movePosition.Price)
}

func (movePosition MovePosition) GetQuantity() float64 {
	return Deref(movePosition.Quantity)
}

func (movePosition MovePosition) GetSourceSlot() Slot {
	return Deref(movePosition.SourceSlot)
}

func (movePosition MovePosition) GetTargetSlot() Slot {
	return Deref(movePosition.TargetSlot)
}

func (movePosition MovePosition) GetThings() Slice[string] {
	return movePosition.Things
}

func (movePosition *MovePosition) SetAssortment(assortment AsAssortment) *MovePosition {
	movePosition.Assortment = assortment.AsAssortment()
	return movePosition
}

func (movePosition *MovePosition) SetPack(pack *Pack) *MovePosition {
	movePosition.Pack = pack
	return movePosition
}

func (movePosition *MovePosition) SetPrice(price float64) *MovePosition {
	movePosition.Price = &price
	return movePosition
}

func (movePosition *MovePosition) SetQuantity(quantity float64) *MovePosition {
	movePosition.Quantity = &quantity
	return movePosition
}

func (movePosition *MovePosition) SetSourceSlot(sourceSlot *Slot) *MovePosition {
	movePosition.SourceSlot = sourceSlot.Clean()
	return movePosition
}

func (movePosition *MovePosition) SetTargetSlot(targetSlot *Slot) *MovePosition {
	movePosition.TargetSlot = targetSlot.Clean()
	return movePosition
}

func (movePosition *MovePosition) SetThings(things ...string) *MovePosition {
	movePosition.Things = NewSliceFrom(things)
	return movePosition
}

func (movePosition MovePosition) String() string {
	return Stringify(movePosition)
}

// MetaType возвращает тип сущности.
func (MovePosition) MetaType() MetaType {
	return MetaTypeMovePosition
}

// MoveService
// Сервис для работы со перемещениями.
type MoveService interface {
	GetList(ctx context.Context, params ...*Params) (*List[Move], *resty.Response, error)
	Create(ctx context.Context, move *Move, params ...*Params) (*Move, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, moveList Slice[Move], params ...*Params) (*Slice[Move], *resty.Response, error)
	DeleteMany(ctx context.Context, entities ...*Move) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Move, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, move *Move, params ...*Params) (*Move, *resty.Response, error)
	Template(ctx context.Context) (*Move, *resty.Response, error)
	TemplateBased(ctx context.Context, basedOn ...MetaOwner) (*Move, *resty.Response, error)
	GetMetadata(ctx context.Context) (*MetaAttributesSharedStatesWrapper, *resty.Response, error)
	GetPositions(ctx context.Context, id uuid.UUID, params ...*Params) (*MetaArray[MovePosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params ...*Params) (*MovePosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *MovePosition, params ...*Params) (*MovePosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id uuid.UUID, position *MovePosition) (*MovePosition, *resty.Response, error)
	CreatePositionMany(ctx context.Context, id uuid.UUID, positions ...*MovePosition) (*Slice[MovePosition], *resty.Response, error)
	DeletePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (bool, *resty.Response, error)
	DeletePositionMany(ctx context.Context, id uuid.UUID, entities ...*MovePosition) (*DeleteManyResponse, *resty.Response, error)
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
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*Move, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error)
	MoveToTrash(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetStateByID(ctx context.Context, id uuid.UUID) (*State, *resty.Response, error)
	CreateState(ctx context.Context, state *State) (*State, *resty.Response, error)
	UpdateState(ctx context.Context, id uuid.UUID, state *State) (*State, *resty.Response, error)
	CreateUpdateStateMany(ctx context.Context, states ...*State) (*Slice[State], *resty.Response, error)
	DeleteState(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetFiles(ctx context.Context, id uuid.UUID) (*MetaArray[File], *resty.Response, error)
	CreateFile(ctx context.Context, id uuid.UUID, file *File) (*Slice[File], *resty.Response, error)
	UpdateFileMany(ctx context.Context, id uuid.UUID, files ...*File) (*Slice[File], *resty.Response, error)
	DeleteFile(ctx context.Context, id uuid.UUID, fileID uuid.UUID) (bool, *resty.Response, error)
	DeleteFileMany(ctx context.Context, id uuid.UUID, files ...*File) (*DeleteManyResponse, *resty.Response, error)
	Evaluate(ctx context.Context, entity *Move, evaluate ...Evaluate) (*Move, *resty.Response, error)
}

func NewMoveService(client *Client) MoveService {
	e := NewEndpoint(client, "entity/move")
	return newMainService[Move, MovePosition, MetaAttributesSharedStatesWrapper, any](e)
}
