package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Loss Списание.
// Ключевое слово: loss
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-spisanie
type Loss struct {
	Organization *Organization            `json:"organization,omitempty"`
	SyncID       *uuid.UUID               `json:"syncId,omitempty"`
	Moment       *Timestamp               `json:"moment,omitempty"`
	Code         *string                  `json:"code,omitempty"`
	Created      *Timestamp               `json:"created,omitempty"`
	Deleted      *Timestamp               `json:"deleted,omitempty"`
	Description  *string                  `json:"description,omitempty"`
	ExternalCode *string                  `json:"externalCode,omitempty"`
	Files        *MetaArray[File]         `json:"files,omitempty"`
	Group        *Group                   `json:"group,omitempty"`
	ID           *uuid.UUID               `json:"id,omitempty"`
	Meta         *Meta                    `json:"meta,omitempty"`
	SalesReturn  *SalesReturn             `json:"salesReturn,omitempty"`
	Applicable   *bool                    `json:"applicable,omitempty"`
	Project      *Project                 `json:"project,omitempty"`
	Owner        *Employee                `json:"owner,omitempty"`
	Positions    *Positions[LossPosition] `json:"positions,omitempty"`
	Printed      *bool                    `json:"printed,omitempty"`
	AccountID    *uuid.UUID               `json:"accountId,omitempty"`
	Published    *bool                    `json:"published,omitempty"`
	Rate         *Rate                    `json:"rate,omitempty"`
	Shared       *bool                    `json:"shared,omitempty"`
	State        *State                   `json:"state,omitempty"`
	Store        *Store                   `json:"store,omitempty"`
	Sum          *float64                 `json:"sum,omitempty"`
	Name         *string                  `json:"name,omitempty"`
	Updated      *Timestamp               `json:"updated,omitempty"`
	Attributes   Slice[AttributeValue]    `json:"attributes,omitempty"`
}

func (loss Loss) Clean() *Loss {
	return &Loss{Meta: loss.Meta}
}

func (loss Loss) GetOrganization() Organization {
	return Deref(loss.Organization)
}

func (loss Loss) GetSyncID() uuid.UUID {
	return Deref(loss.SyncID)
}

func (loss Loss) GetMoment() Timestamp {
	return Deref(loss.Moment)
}

func (loss Loss) GetCode() string {
	return Deref(loss.Code)
}

func (loss Loss) GetCreated() Timestamp {
	return Deref(loss.Created)
}

func (loss Loss) GetDeleted() Timestamp {
	return Deref(loss.Deleted)
}

func (loss Loss) GetDescription() string {
	return Deref(loss.Description)
}

func (loss Loss) GetExternalCode() string {
	return Deref(loss.ExternalCode)
}

func (loss Loss) GetFiles() MetaArray[File] {
	return Deref(loss.Files)
}

func (loss Loss) GetGroup() Group {
	return Deref(loss.Group)
}

func (loss Loss) GetID() uuid.UUID {
	return Deref(loss.ID)
}

func (loss Loss) GetMeta() Meta {
	return Deref(loss.Meta)
}

func (loss Loss) GetSalesReturn() SalesReturn {
	return Deref(loss.SalesReturn)
}

func (loss Loss) GetApplicable() bool {
	return Deref(loss.Applicable)
}

func (loss Loss) GetProject() Project {
	return Deref(loss.Project)
}

func (loss Loss) GetOwner() Employee {
	return Deref(loss.Owner)
}

func (loss Loss) GetPositions() Positions[LossPosition] {
	return Deref(loss.Positions)
}

func (loss Loss) GetPrinted() bool {
	return Deref(loss.Printed)
}

func (loss Loss) GetAccountID() uuid.UUID {
	return Deref(loss.AccountID)
}

func (loss Loss) GetPublished() bool {
	return Deref(loss.Published)
}

func (loss Loss) GetRate() Rate {
	return Deref(loss.Rate)
}

func (loss Loss) GetShared() bool {
	return Deref(loss.Shared)
}

func (loss Loss) GetState() State {
	return Deref(loss.State)
}

func (loss Loss) GetStore() Store {
	return Deref(loss.Store)
}

func (loss Loss) GetSum() float64 {
	return Deref(loss.Sum)
}

func (loss Loss) GetName() string {
	return Deref(loss.Name)
}

func (loss Loss) GetUpdated() Timestamp {
	return Deref(loss.Updated)
}

func (loss Loss) GetAttributes() Slice[AttributeValue] {
	return loss.Attributes
}

func (loss *Loss) SetOrganization(organization *Organization) *Loss {
	loss.Organization = organization.Clean()
	return loss
}

func (loss *Loss) SetSyncID(syncID uuid.UUID) *Loss {
	loss.SyncID = &syncID
	return loss
}

func (loss *Loss) SetMoment(moment *Timestamp) *Loss {
	loss.Moment = moment
	return loss
}

func (loss *Loss) SetCode(code string) *Loss {
	loss.Code = &code
	return loss
}

func (loss *Loss) SetDescription(description string) *Loss {
	loss.Description = &description
	return loss
}

func (loss *Loss) SetExternalCode(externalCode string) *Loss {
	loss.ExternalCode = &externalCode
	return loss
}

func (loss *Loss) SetFiles(files Slice[File]) *Loss {
	loss.Files = NewMetaArrayRows(files)
	return loss
}

func (loss *Loss) SetGroup(group *Group) *Loss {
	loss.Group = group.Clean()
	return loss
}

func (loss *Loss) SetMeta(meta *Meta) *Loss {
	loss.Meta = meta
	return loss
}

func (loss *Loss) SetSalesReturn(salesReturn *SalesReturn) *Loss {
	loss.SalesReturn = salesReturn.Clean()
	return loss
}

func (loss *Loss) SetApplicable(applicable bool) *Loss {
	loss.Applicable = &applicable
	return loss
}

func (loss *Loss) SetProject(project *Project) *Loss {
	loss.Project = project.Clean()
	return loss
}

func (loss *Loss) SetOwner(owner *Employee) *Loss {
	loss.Owner = owner.Clean()
	return loss
}

func (loss *Loss) SetPositions(positions *Positions[LossPosition]) *Loss {
	loss.Positions = positions
	return loss
}

func (loss *Loss) SetRate(rate *Rate) *Loss {
	loss.Rate = rate
	return loss
}

func (loss *Loss) SetShared(shared bool) *Loss {
	loss.Shared = &shared
	return loss
}

func (loss *Loss) SetState(state *State) *Loss {
	loss.State = state.Clean()
	return loss
}

func (loss *Loss) SetStore(store *Store) *Loss {
	loss.Store = store.Clean()
	return loss
}

func (loss *Loss) SetName(name string) *Loss {
	loss.Name = &name
	return loss
}

func (loss *Loss) SetAttributes(attributes Slice[AttributeValue]) *Loss {
	loss.Attributes = attributes
	return loss
}

func (loss Loss) String() string {
	return Stringify(loss)
}

func (loss Loss) MetaType() MetaType {
	return MetaTypeLoss
}

// LossPosition Позиция Списания.
// Ключевое слово: lossposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-spisanie-spisaniq-pozicii-spisaniq
type LossPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID сущности
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Price      *float64            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Reason     *string             `json:"reason,omitempty"`     // Причина списания данной позиции
	Slot       *Slot               `json:"slot,omitempty"`       // Ячейка на складе
	Things     Slice[string]       `json:"things,omitempty"`     // Серийные номера. Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете. В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута.
}

func (lossPosition LossPosition) GetAccountID() uuid.UUID {
	return Deref(lossPosition.AccountID)
}

func (lossPosition LossPosition) GetAssortment() AssortmentPosition {
	return Deref(lossPosition.Assortment)
}

func (lossPosition LossPosition) GetID() uuid.UUID {
	return Deref(lossPosition.ID)
}

func (lossPosition LossPosition) GetPack() Pack {
	return Deref(lossPosition.Pack)
}

func (lossPosition LossPosition) GetPrice() float64 {
	return Deref(lossPosition.Price)
}

func (lossPosition LossPosition) GetQuantity() float64 {
	return Deref(lossPosition.Quantity)
}

func (lossPosition LossPosition) GetReason() string {
	return Deref(lossPosition.Reason)
}

func (lossPosition LossPosition) GetSlot() Slot {
	return Deref(lossPosition.Slot)
}

func (lossPosition LossPosition) GetThings() Slice[string] {
	return lossPosition.Things
}

func (lossPosition *LossPosition) SetAssortment(assortment AsAssortment) *LossPosition {
	lossPosition.Assortment = assortment.AsAssortment()
	return lossPosition
}

func (lossPosition *LossPosition) SetPack(pack *Pack) *LossPosition {
	lossPosition.Pack = pack
	return lossPosition
}

func (lossPosition *LossPosition) SetPrice(price float64) *LossPosition {
	lossPosition.Price = &price
	return lossPosition
}

func (lossPosition *LossPosition) SetQuantity(quantity float64) *LossPosition {
	lossPosition.Quantity = &quantity
	return lossPosition
}

func (lossPosition *LossPosition) SetReason(reason string) *LossPosition {
	lossPosition.Reason = &reason
	return lossPosition
}

func (lossPosition *LossPosition) SetSlot(slot *Slot) *LossPosition {
	lossPosition.Slot = slot.Clean()
	return lossPosition
}

func (lossPosition *LossPosition) SetThings(things Slice[string]) *LossPosition {
	lossPosition.Things = things
	return lossPosition
}

func (lossPosition LossPosition) String() string {
	return Stringify(lossPosition)
}

func (lossPosition LossPosition) MetaType() MetaType {
	return MetaTypeLossPosition
}

// LossTemplateArg
// Документ: Списание (loss)
// Основание, на котором он может быть создан:
// - Возврат покупателя (salesreturn)
// - инвентаризация(inventory)
//type LossTemplateArg struct {
//	SalesReturn *MetaWrapper `json:"salesReturn,omitempty"`
//	Inventory   *MetaWrapper `json:"inventory,omitempty"`
//}

// LossService
// Сервис для работы со списаниями.
type LossService interface {
	GetList(ctx context.Context, params *Params) (*List[Loss], *resty.Response, error)
	Create(ctx context.Context, loss *Loss, params *Params) (*Loss, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, lossList Slice[Loss], params *Params) (*Slice[Loss], *resty.Response, error)
	DeleteMany(ctx context.Context, lossList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params *Params) (*Loss, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, loss *Loss, params *Params) (*Loss, *resty.Response, error)
	//endpointTemplate[Loss]
	//endpointTemplateBasedOn[Loss, LossTemplateArg]
	GetMetadata(ctx context.Context) (*MetaAttributesSharedStatesWrapper, *resty.Response, error)
	GetPositions(ctx context.Context, id uuid.UUID, params *Params) (*MetaArray[LossPosition], *resty.Response, error)
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params *Params) (*LossPosition, *resty.Response, error)
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *LossPosition, params *Params) (*LossPosition, *resty.Response, error)
	CreatePosition(ctx context.Context, id uuid.UUID, position *LossPosition) (*LossPosition, *resty.Response, error)
	CreatePositions(ctx context.Context, id uuid.UUID, positions Slice[LossPosition]) (*Slice[LossPosition], *resty.Response, error)
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
	GetPublications(ctx context.Context, id uuid.UUID) (*MetaArray[Publication], *resty.Response, error)
	GetPublicationByID(ctx context.Context, id uuid.UUID, publicationID uuid.UUID) (*Publication, *resty.Response, error)
	Publish(ctx context.Context, id uuid.UUID, template Templater) (*Publication, *resty.Response, error)
	DeletePublication(ctx context.Context, id uuid.UUID, publicationID uuid.UUID) (bool, *resty.Response, error)
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*Loss, *resty.Response, error)
	DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error)
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
	DeleteFiles(ctx context.Context, id uuid.UUID, files *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
}

func NewLossService(client *Client) LossService {
	e := NewEndpoint(client, "entity/loss")
	return newMainService[Loss, LossPosition, MetaAttributesSharedStatesWrapper, any](e)
}
