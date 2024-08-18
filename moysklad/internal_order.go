package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"time"
)

// InternalOrder Внутренний заказ.
//
// Код сущности: internalorder
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vnutrennij-zakaz
type InternalOrder struct {
	Organization          *Organization                     `json:"organization,omitempty"`          // Метаданные юрлица
	Description           *string                           `json:"description,omitempty"`           // Комментарий Внутреннего заказа
	VatSum                *float64                          `json:"vatSum,omitempty"`                // Сумма НДС
	AccountID             *uuid.UUID                        `json:"accountId,omitempty"`             // ID учётной записи
	Created               *Timestamp                        `json:"created,omitempty"`               // Дата создания
	Deleted               *Timestamp                        `json:"deleted,omitempty"`               // Момент последнего удаления Внутреннего заказа
	DeliveryPlannedMoment *Timestamp                        `json:"deliveryPlannedMoment,omitempty"` // Планируемая дата приемки
	Owner                 *Employee                         `json:"owner,omitempty"`                 // Метаданные владельца (Сотрудника)
	ExternalCode          *string                           `json:"externalCode,omitempty"`          // Внешний код Внутреннего заказа
	Files                 *MetaArray[File]                  `json:"files,omitempty"`                 // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group                 *Group                            `json:"group,omitempty"`                 // Отдел сотрудника
	ID                    *uuid.UUID                        `json:"id,omitempty"`                    // ID Внутреннего заказа
	Meta                  *Meta                             `json:"meta,omitempty"`                  // Метаданные Внутреннего заказа
	Positions             *MetaArray[InternalOrderPosition] `json:"positions,omitempty"`             // Метаданные позиций Внутреннего заказа
	Moves                 Slice[Move]                       `json:"moves,omitempty"`                 // Коллекция метаданных на связанные заказы перемещения
	Name                  *string                           `json:"name,omitempty"`                  // Наименование Внутреннего заказа
	Code                  *string                           `json:"code,omitempty"`                  // Код Внутреннего заказа
	Applicable            *bool                             `json:"applicable,omitempty"`            // Отметка о проведении
	Moment                *Timestamp                        `json:"moment,omitempty"`                // Дата документа
	Printed               *bool                             `json:"printed,omitempty"`               // Напечатан ли документ
	Project               *NullValue[Project]               `json:"project,omitempty"`               // Метаданные проекта
	Published             *bool                             `json:"published,omitempty"`             // Опубликован ли документ
	PurchaseOrders        Slice[PurchaseOrder]              `json:"purchaseOrders,omitempty"`        // Коллекция метаданных на связанные заказы поставщикам
	Rate                  *NullValue[Rate]                  `json:"rate,omitempty"`                  // Валюта
	Shared                *bool                             `json:"shared,omitempty"`                // Общий доступ
	State                 *NullValue[State]                 `json:"state,omitempty"`                 // Метаданные статуса Внутреннего заказа
	Store                 *NullValue[Store]                 `json:"store,omitempty"`                 // Метаданные склада
	Sum                   *float64                          `json:"sum,omitempty"`                   // Сумма Внутреннего заказа в копейках
	SyncID                *uuid.UUID                        `json:"syncId,omitempty"`                // ID синхронизации
	Updated               *Timestamp                        `json:"updated,omitempty"`               // Момент последнего обновления Внутреннего заказа
	VatEnabled            *bool                             `json:"vatEnabled,omitempty"`            // Учитывается ли НДС
	VatIncluded           *bool                             `json:"vatIncluded,omitempty"`           // Включен ли НДС в цену
	Attributes            Slice[Attribute]                  `json:"attributes,omitempty"`            // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (internalOrder InternalOrder) Clean() *InternalOrder {
	if internalOrder.Meta == nil {
		return nil
	}
	return &InternalOrder{Meta: internalOrder.Meta}
}

// AsTaskOperation реализует интерфейс [TaskOperationConverter].
func (internalOrder InternalOrder) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: internalOrder.Meta}
}

// GetOrganization возвращает Метаданные юрлица.
func (internalOrder InternalOrder) GetOrganization() Organization {
	return Deref(internalOrder.Organization)
}

// GetDescription возвращает Комментарий Внутреннего заказа.
func (internalOrder InternalOrder) GetDescription() string {
	return Deref(internalOrder.Description)
}

// GetVatSum возвращает Сумму НДС.
func (internalOrder InternalOrder) GetVatSum() float64 {
	return Deref(internalOrder.VatSum)
}

// GetAccountID возвращает ID учётной записи.
func (internalOrder InternalOrder) GetAccountID() uuid.UUID {
	return Deref(internalOrder.AccountID)
}

// GetCreated возвращает Дату создания.
func (internalOrder InternalOrder) GetCreated() time.Time {
	return Deref(internalOrder.Created).Time()
}

// GetDeleted возвращает Момент последнего удаления Внутреннего заказа.
func (internalOrder InternalOrder) GetDeleted() time.Time {
	return Deref(internalOrder.Deleted).Time()
}

// GetDeliveryPlannedMoment возвращает Планируемую дата приемки.
func (internalOrder InternalOrder) GetDeliveryPlannedMoment() time.Time {
	return Deref(internalOrder.DeliveryPlannedMoment).Time()
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (internalOrder InternalOrder) GetOwner() Employee {
	return Deref(internalOrder.Owner)
}

// GetExternalCode возвращает Внешний код Внутреннего заказа.
func (internalOrder InternalOrder) GetExternalCode() string {
	return Deref(internalOrder.ExternalCode)
}

// GetFiles возвращает Метаданные массива Файлов.
func (internalOrder InternalOrder) GetFiles() MetaArray[File] {
	return Deref(internalOrder.Files)
}

// GetGroup возвращает Отдел сотрудника.
func (internalOrder InternalOrder) GetGroup() Group {
	return Deref(internalOrder.Group)
}

// GetID возвращает ID Внутреннего заказа.
func (internalOrder InternalOrder) GetID() uuid.UUID {
	return Deref(internalOrder.ID)
}

// GetMeta возвращает Метаданные Внутреннего заказа.
func (internalOrder InternalOrder) GetMeta() Meta {
	return Deref(internalOrder.Meta)
}

// GetPositions возвращает Метаданные позиций  Внутреннего заказа.
func (internalOrder InternalOrder) GetPositions() MetaArray[InternalOrderPosition] {
	return Deref(internalOrder.Positions)
}

// GetMoves возвращает Массив ссылок на связанные перемещения.
func (internalOrder InternalOrder) GetMoves() Slice[Move] {
	return internalOrder.Moves
}

// GetName возвращает Наименование Внутреннего заказа.
func (internalOrder InternalOrder) GetName() string {
	return Deref(internalOrder.Name)
}

// GetCode возвращает Код Внутреннего заказа.
func (internalOrder InternalOrder) GetCode() string {
	return Deref(internalOrder.Code)
}

// GetApplicable возвращает Отметку о проведении.
func (internalOrder InternalOrder) GetApplicable() bool {
	return Deref(internalOrder.Applicable)
}

// GetMoment возвращает Дату документа.
func (internalOrder InternalOrder) GetMoment() time.Time {
	return Deref(internalOrder.Moment).Time()
}

// GetPrinted возвращает true, если документ напечатан.
func (internalOrder InternalOrder) GetPrinted() bool {
	return Deref(internalOrder.Printed)
}

// GetProject возвращает Метаданные проекта.
func (internalOrder InternalOrder) GetProject() Project {
	return Deref(internalOrder.Project).getValue()
}

// GetPublished возвращает true, если документ опубликован.
func (internalOrder InternalOrder) GetPublished() bool {
	return Deref(internalOrder.Published)
}

// GetPurchaseOrders возвращает Массив ссылок на связанные заказы поставщикам.
func (internalOrder InternalOrder) GetPurchaseOrders() Slice[PurchaseOrder] {
	return internalOrder.PurchaseOrders
}

// GetRate возвращает Валюту.
func (internalOrder InternalOrder) GetRate() Rate {
	return Deref(internalOrder.Rate).getValue()
}

// GetShared возвращает флаг Общего доступа.
func (internalOrder InternalOrder) GetShared() bool {
	return Deref(internalOrder.Shared)
}

// GetState возвращает Метаданные статуса Внутреннего заказа.
func (internalOrder InternalOrder) GetState() State {
	return Deref(internalOrder.State).getValue()
}

// GetStore возвращает Метаданные склада.
func (internalOrder InternalOrder) GetStore() Store {
	return Deref(internalOrder.Store).getValue()
}

// GetSum возвращает Сумму Внутреннего заказа в копейках.
func (internalOrder InternalOrder) GetSum() float64 {
	return Deref(internalOrder.Sum)
}

// GetSyncID возвращает ID синхронизации.
func (internalOrder InternalOrder) GetSyncID() uuid.UUID {
	return Deref(internalOrder.SyncID)
}

// GetUpdated возвращает Момент последнего обновления Внутреннего заказа.
func (internalOrder InternalOrder) GetUpdated() time.Time {
	return Deref(internalOrder.Updated).Time()
}

// GetVatEnabled возвращает true, если учитывается НДС.
func (internalOrder InternalOrder) GetVatEnabled() bool {
	return Deref(internalOrder.VatEnabled)
}

// GetVatIncluded возвращает true, если НДС включен в цену.
func (internalOrder InternalOrder) GetVatIncluded() bool {
	return Deref(internalOrder.VatIncluded)
}

// GetAttributes возвращает Список метаданных доп. полей.
func (internalOrder InternalOrder) GetAttributes() Slice[Attribute] {
	return internalOrder.Attributes
}

// SetOrganization устанавливает Метаданные юрлица.
func (internalOrder *InternalOrder) SetOrganization(organization *Organization) *InternalOrder {
	if organization != nil {
		internalOrder.Organization = organization.Clean()
	}
	return internalOrder
}

// SetDescription устанавливает Комментарий Внутреннего заказа.
func (internalOrder *InternalOrder) SetDescription(description string) *InternalOrder {
	internalOrder.Description = &description
	return internalOrder
}

// SetDeliveryPlannedMoment устанавливает Планируемую дату приемки.
func (internalOrder *InternalOrder) SetDeliveryPlannedMoment(deliveryPlannedMoment time.Time) *InternalOrder {
	internalOrder.DeliveryPlannedMoment = NewTimestamp(deliveryPlannedMoment)
	return internalOrder
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (internalOrder *InternalOrder) SetOwner(owner *Employee) *InternalOrder {
	if owner != nil {
		internalOrder.Owner = owner.Clean()
	}
	return internalOrder
}

// SetExternalCode устанавливает Внешний код Внутреннего заказа.
func (internalOrder *InternalOrder) SetExternalCode(externalCode string) *InternalOrder {
	internalOrder.ExternalCode = &externalCode
	return internalOrder
}

// SetFiles устанавливает Метаданные массива Файлов.
//
// Принимает множество объектов [File].
func (internalOrder *InternalOrder) SetFiles(files ...*File) *InternalOrder {
	internalOrder.Files = NewMetaArrayFrom(files)
	return internalOrder
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (internalOrder *InternalOrder) SetGroup(group *Group) *InternalOrder {
	if group != nil {
		internalOrder.Group = group.Clean()
	}
	return internalOrder
}

// SetMeta устанавливает Метаданные Внутреннего заказа.
func (internalOrder *InternalOrder) SetMeta(meta *Meta) *InternalOrder {
	internalOrder.Meta = meta
	return internalOrder
}

// SetPositions устанавливает Метаданные позиций Внутреннего заказа.
//
// Принимает множество объектов [InternalOrderPosition].
func (internalOrder *InternalOrder) SetPositions(positions ...*InternalOrderPosition) *InternalOrder {
	internalOrder.Positions = NewMetaArrayFrom(positions)
	return internalOrder
}

// SetMoves устанавливает Массив ссылок на связанные перемещения.
//
// Принимает множество объектов [Move].
func (internalOrder *InternalOrder) SetMoves(moves ...*Move) *InternalOrder {
	internalOrder.Moves.Push(moves...)
	return internalOrder
}

// SetName устанавливает Наименование Внутреннего заказа.
func (internalOrder *InternalOrder) SetName(name string) *InternalOrder {
	internalOrder.Name = &name
	return internalOrder
}

// SetCode устанавливает Код Внутреннего заказа.
func (internalOrder *InternalOrder) SetCode(code string) *InternalOrder {
	internalOrder.Code = &code
	return internalOrder
}

// SetApplicable устанавливает Отметку о проведении.
func (internalOrder *InternalOrder) SetApplicable(applicable bool) *InternalOrder {
	internalOrder.Applicable = &applicable
	return internalOrder
}

// SetMoment устанавливает Дату документа.
func (internalOrder *InternalOrder) SetMoment(moment time.Time) *InternalOrder {
	internalOrder.Moment = NewTimestamp(moment)
	return internalOrder
}

// SetProject устанавливает Метаданные проекта.
//
// Передача nil передаёт сброс значения (null).
func (internalOrder *InternalOrder) SetProject(project *Project) *InternalOrder {
	internalOrder.Project = NewNullValue(project)
	return internalOrder
}

// SetPurchaseOrders устанавливает Массив ссылок на связанные заказы поставщикам.
//
// Принимает множество объектов [PurchaseOrder].
func (internalOrder *InternalOrder) SetPurchaseOrders(purchaseOrders ...*PurchaseOrder) *InternalOrder {
	internalOrder.PurchaseOrders.Push(purchaseOrders...)
	return internalOrder
}

// SetRate устанавливает Валюту.
//
// Передача nil передаёт сброс значения (null).
func (internalOrder *InternalOrder) SetRate(rate *Rate) *InternalOrder {
	internalOrder.Rate = NewNullValue(rate)
	return internalOrder
}

// SetShared устанавливает флаг общего доступа.
func (internalOrder *InternalOrder) SetShared(shared bool) *InternalOrder {
	internalOrder.Shared = &shared
	return internalOrder
}

// SetState устанавливает Метаданные статуса Внутреннего заказа.
//
// Передача nil передаёт сброс значения (null).
func (internalOrder *InternalOrder) SetState(state *State) *InternalOrder {
	internalOrder.State = NewNullValue(state)
	return internalOrder
}

// SetStore устанавливает Метаданные склада.
//
// Передача nil передаёт сброс значения (null).
func (internalOrder *InternalOrder) SetStore(store *Store) *InternalOrder {
	internalOrder.Store = NewNullValue(store)
	return internalOrder
}

// SetSyncID устанавливает ID синхронизации.
func (internalOrder *InternalOrder) SetSyncID(syncID uuid.UUID) *InternalOrder {
	internalOrder.SyncID = &syncID
	return internalOrder
}

// SetVatEnabled устанавливает значение, учитывающее НДС.
func (internalOrder *InternalOrder) SetVatEnabled(vatEnabled bool) *InternalOrder {
	internalOrder.VatEnabled = &vatEnabled
	return internalOrder
}

// SetVatIncluded устанавливает флаг включения НДС в цену.
func (internalOrder *InternalOrder) SetVatIncluded(vatIncluded bool) *InternalOrder {
	internalOrder.VatIncluded = &vatIncluded
	return internalOrder
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (internalOrder *InternalOrder) SetAttributes(attributes ...*Attribute) *InternalOrder {
	internalOrder.Attributes.Push(attributes...)
	return internalOrder
}

// String реализует интерфейс [fmt.Stringer].
func (internalOrder InternalOrder) String() string {
	return Stringify(internalOrder)
}

// MetaType возвращает код сущности.
func (InternalOrder) MetaType() MetaType {
	return MetaTypeInternalOrder
}

// Update shortcut
func (internalOrder *InternalOrder) Update(ctx context.Context, client *Client, params ...*Params) (*InternalOrder, *resty.Response, error) {
	return NewInternalOrderService(client).Update(ctx, internalOrder.GetID(), internalOrder, params...)
}

// Create shortcut
func (internalOrder *InternalOrder) Create(ctx context.Context, client *Client, params ...*Params) (*InternalOrder, *resty.Response, error) {
	return NewInternalOrderService(client).Create(ctx, internalOrder, params...)
}

// Delete shortcut
func (internalOrder *InternalOrder) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewInternalOrderService(client).Delete(ctx, internalOrder)
}

// InternalOrderPosition Позиция Внутреннего заказа.
//
// Код сущности: internalorderposition
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vnutrennij-zakaz-vnutrennie-zakazy-pozicii-vnutrennego-zakaza
type InternalOrderPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учётной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Price      *float64            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Vat        *int                `json:"vat,omitempty"`        // НДС, которым облагается текущая позиция
	VatEnabled *bool               `json:"vatEnabled,omitempty"` // Включен ли НДС для позиции. С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
}

// GetAccountID возвращает ID учётной записи.
func (internalOrderPosition InternalOrderPosition) GetAccountID() uuid.UUID {
	return Deref(internalOrderPosition.AccountID)
}

// GetAssortment возвращает Метаданные товара/услуги/серии/модификации, которую представляет собой позиция.
func (internalOrderPosition InternalOrderPosition) GetAssortment() AssortmentPosition {
	return Deref(internalOrderPosition.Assortment)
}

// GetID возвращает ID позиции.
func (internalOrderPosition InternalOrderPosition) GetID() uuid.UUID {
	return Deref(internalOrderPosition.ID)
}

// GetPack возвращает Упаковку Товара.
func (internalOrderPosition InternalOrderPosition) GetPack() Pack {
	return Deref(internalOrderPosition.Pack)
}

// GetPrice возвращает Цену товара/услуги в копейках.
func (internalOrderPosition InternalOrderPosition) GetPrice() float64 {
	return Deref(internalOrderPosition.Price)
}

// GetQuantity возвращает Количество товаров данного вида в позиции.
func (internalOrderPosition InternalOrderPosition) GetQuantity() float64 {
	return Deref(internalOrderPosition.Quantity)
}

// GetVat возвращает НДС, которым облагается текущая позиция.
func (internalOrderPosition InternalOrderPosition) GetVat() int {
	return Deref(internalOrderPosition.Vat)
}

// GetVatEnabled возвращает true, если учитывается НДС.
func (internalOrderPosition InternalOrderPosition) GetVatEnabled() bool {
	return Deref(internalOrderPosition.VatEnabled)
}

// SetAssortment устанавливает Метаданные товара/услуги/серии/модификации, которую представляет собой позиция.
//
// Принимает объект, реализующий интерфейс [AssortmentConverter].
func (internalOrderPosition *InternalOrderPosition) SetAssortment(assortment AssortmentConverter) *InternalOrderPosition {
	if assortment != nil {
		internalOrderPosition.Assortment = assortment.AsAssortment()
	}
	return internalOrderPosition
}

// SetPack устанавливает Упаковку Товара.
func (internalOrderPosition *InternalOrderPosition) SetPack(pack *Pack) *InternalOrderPosition {
	if pack != nil {
		internalOrderPosition.Pack = pack
	}
	return internalOrderPosition
}

// SetPrice устанавливает Цену товара/услуги в копейках.
func (internalOrderPosition *InternalOrderPosition) SetPrice(price float64) *InternalOrderPosition {
	internalOrderPosition.Price = &price
	return internalOrderPosition
}

// SetQuantity устанавливает Количество товаров данного вида в позиции.
func (internalOrderPosition *InternalOrderPosition) SetQuantity(quantity float64) *InternalOrderPosition {
	internalOrderPosition.Quantity = &quantity
	return internalOrderPosition
}

// SetVat устанавливает НДС, которым облагается текущая позиция.
func (internalOrderPosition *InternalOrderPosition) SetVat(vat int) *InternalOrderPosition {
	internalOrderPosition.Vat = &vat
	return internalOrderPosition
}

// SetVatEnabled устанавливает значение, учитывающее НДС для текущей позиции.
func (internalOrderPosition *InternalOrderPosition) SetVatEnabled(vatEnabled bool) *InternalOrderPosition {
	internalOrderPosition.VatEnabled = &vatEnabled
	return internalOrderPosition
}

// String реализует интерфейс [fmt.Stringer].
func (internalOrderPosition InternalOrderPosition) String() string {
	return Stringify(internalOrderPosition)
}

// MetaType возвращает код сущности.
func (InternalOrderPosition) MetaType() MetaType {
	return MetaTypeInternalOrderPosition
}

// InternalOrderService описывает методы сервиса для работы с внутренними заказами.
type InternalOrderService interface {
	// GetList выполняет запрос на получение списка внутренних заказов.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[InternalOrder], *resty.Response, error)

	// GetListAll выполняет запрос на получение всех внутренних заказов в виде списка.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает список объектов.
	GetListAll(ctx context.Context, params ...*Params) (Slice[InternalOrder], *resty.Response, error)

	// Create выполняет запрос на создание внутреннего заказа.
	// Обязательные поля для заполнения:
	//	- organization (Ссылка на ваше юрлицо)
	// Принимает контекст, внутренний заказ и опционально объект параметров запроса Params.
	// Возвращает созданный внутренний заказ.
	Create(ctx context.Context, internalOrder *InternalOrder, params ...*Params) (*InternalOrder, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение внутренних заказов.
	// Изменяемые внутренние заказы должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список внутренних заказов и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых внутренних заказов.
	CreateUpdateMany(ctx context.Context, internalOrderList Slice[InternalOrder], params ...*Params) (*Slice[InternalOrder], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление внутренних заказов.
	// Принимает контекст и множество внутренних заказов.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*InternalOrder) (*DeleteManyResponse, *resty.Response, error)

	// DeleteByID выполняет запрос на удаление внутреннего заказа по ID.
	// Принимает контекст и ID внутреннего заказа.
	// Возвращает «true» в случае успешного удаления внутреннего заказа.
	DeleteByID(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// Delete выполняет запрос на удаление внутреннего заказа.
	// Принимает контекст и внутренний заказ.
	// Возвращает «true» в случае успешного удаления внутреннего заказа.
	Delete(ctx context.Context, entity *InternalOrder) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельного внутреннего заказа по ID.
	// Принимает контекст, ID внутреннего заказа и опционально объект параметров запроса Params.
	// Возвращает найденный внутренний заказ.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*InternalOrder, *resty.Response, error)

	// Update выполняет запрос на изменение внутреннего заказа.
	// Принимает контекст, внутренний заказ и опционально объект параметров запроса Params.
	// Возвращает изменённый внутренний заказ.
	Update(ctx context.Context, id uuid.UUID, internalOrder *InternalOrder, params ...*Params) (*InternalOrder, *resty.Response, error)

	// Template выполняет запрос на получение предзаполненного внутреннего заказа со стандартными полями без связи с какими-либо другими документами.
	// Принимает контекст.
	// Возвращает предзаполненный внутренний заказ.
	Template(ctx context.Context) (*InternalOrder, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных внутренних заказов.
	// Принимает контекст.
	// Возвращает объект метаданных MetaAttributesStatesSharedWrapper.
	GetMetadata(ctx context.Context) (*MetaAttributesStatesSharedWrapper, *resty.Response, error)

	// GetPositionList выполняет запрос на получение списка позиций документа.
	// Принимает контекст, ID документа и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetPositionList(ctx context.Context, id uuid.UUID, params ...*Params) (*List[InternalOrderPosition], *resty.Response, error)

	// GetPositionByID выполняет запрос на получение отдельной позиции документа по ID.
	// Принимает контекст, ID документа, ID позиции и опционально объект параметров запроса Params.
	// Возвращает найденную позицию.
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params ...*Params) (*InternalOrderPosition, *resty.Response, error)

	// UpdatePosition выполняет запрос на изменение позиции документа.
	// Принимает контекст, ID документа, ID позиции, позицию документа и опционально объект параметров запроса Params.
	// Возвращает изменённую позицию.
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *InternalOrderPosition, params ...*Params) (*InternalOrderPosition, *resty.Response, error)

	// CreatePosition выполняет запрос на добавление позиции документа.
	// Принимает контекст, ID документа, позицию документа и опционально объект параметров запроса Params.
	// Возвращает добавленную позицию.
	CreatePosition(ctx context.Context, id uuid.UUID, position *InternalOrderPosition, params ...*Params) (*InternalOrderPosition, *resty.Response, error)

	// CreatePositionMany выполняет запрос на массовое добавление позиций документа.
	// Принимает контекст, ID документа и множество позиций.
	// Возвращает список добавленных позиций.
	CreatePositionMany(ctx context.Context, id uuid.UUID, positions ...*InternalOrderPosition) (*Slice[InternalOrderPosition], *resty.Response, error)

	// DeletePosition выполняет запрос на удаление позиции документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает «true» в случае успешного удаления позиции.
	DeletePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (bool, *resty.Response, error)

	// DeletePositionMany выполняет запрос на массовое удаление позиций документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeletePositionMany(ctx context.Context, id uuid.UUID, positions ...*InternalOrderPosition) (*DeleteManyResponse, *resty.Response, error)

	// GetPositionTrackingCodeList выполняет запрос на получение кодов маркировки позиции документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает объект List.
	GetPositionTrackingCodeList(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (*List[TrackingCode], *resty.Response, error)

	// CreateUpdatePositionTrackingCodeMany выполняет запрос на массовое создание/изменение кодов маркировки позиции документа.
	// Принимает контекст, ID документа, ID позиции и множество кодов маркировки.
	// Возвращает список созданных и/или изменённых кодов маркировки позиции документа.
	CreateUpdatePositionTrackingCodeMany(ctx context.Context, id uuid.UUID, positionID uuid.UUID, trackingCodes ...*TrackingCode) (*Slice[TrackingCode], *resty.Response, error)

	// DeletePositionTrackingCodeMany выполняет запрос на массовое удаление кодов маркировки позиции документа.
	// Принимает контекст, ID документа, ID позиции и множество кодов маркировки.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeletePositionTrackingCodeMany(ctx context.Context, id uuid.UUID, positionID uuid.UUID, trackingCodes ...*TrackingCode) (*DeleteManyResponse, *resty.Response, error)

	// GetAttributeList выполняет запрос на получение списка доп полей.
	// Принимает контекст.
	// Возвращает объект List.
	GetAttributeList(ctx context.Context) (*List[Attribute], *resty.Response, error)

	// GetAttributeByID выполняет запрос на получение отдельного доп поля по ID.
	// Принимает контекст и ID доп поля.
	// Возвращает найденное доп поле.
	GetAttributeByID(ctx context.Context, id uuid.UUID) (*Attribute, *resty.Response, error)

	// CreateAttribute выполняет запрос на создание доп поля.
	// Принимает контекст и доп поле.
	// Возвращает созданное доп поле.
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)

	// CreateUpdateAttributeMany выполняет запрос на массовое создание и/или изменение доп полей.
	// Изменяемые доп поля должны содержать идентификатор в виде метаданных.
	// Принимает контекст и множество доп полей.
	// Возвращает список созданных и/или изменённых доп полей.
	CreateUpdateAttributeMany(ctx context.Context, attributes ...*Attribute) (*Slice[Attribute], *resty.Response, error)

	// UpdateAttribute выполняет запрос на изменения доп поля.
	// Принимает контекст, ID доп поля и доп поле.
	// Возвращает изменённое доп поле.
	UpdateAttribute(ctx context.Context, id uuid.UUID, attribute *Attribute) (*Attribute, *resty.Response, error)

	// DeleteAttribute выполняет запрос на удаление доп поля.
	// Принимает контекст и ID доп поля.
	// Возвращает «true» в случае успешного удаления доп поля.
	DeleteAttribute(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// DeleteAttributeMany выполняет запрос на массовое удаление доп полей.
	// Принимает контекст и множество доп полей.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteAttributeMany(ctx context.Context, attributes ...*Attribute) (*DeleteManyResponse, *resty.Response, error)

	// GetPublicationList выполняет запрос на получение списка публикаций.
	// Принимает контекст и ID документа.
	// Возвращает объект List.
	GetPublicationList(ctx context.Context, id uuid.UUID) (*List[Publication], *resty.Response, error)

	// GetPublicationByID выполняет запрос на получение отдельной публикации по ID.
	// Принимает контекст, ID документа и ID публикации.
	// Возвращает найденную публикацию.
	GetPublicationByID(ctx context.Context, id uuid.UUID, publicationID uuid.UUID) (*Publication, *resty.Response, error)

	// Publish выполняет запрос на создание публикации.
	// Принимает контекст, ID документа и шаблон (CustomTemplate или EmbeddedTemplate)
	// Возвращает созданную публикацию.
	Publish(ctx context.Context, id uuid.UUID, template TemplateConverter) (*Publication, *resty.Response, error)

	// DeletePublication выполняет запрос на удаление публикации.
	// Принимает контекст, ID документа и ID публикации.
	// Возвращает «true» в случае успешного удаления публикации.
	DeletePublication(ctx context.Context, id uuid.UUID, publicationID uuid.UUID) (bool, *resty.Response, error)

	// GetBySyncID выполняет запрос на получение отдельного документа по syncID.
	// Принимает контекст и syncID документа.
	// Возвращает найденный документ.
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*InternalOrder, *resty.Response, error)

	// DeleteBySyncID выполняет запрос на удаление документа по syncID.
	// Принимает контекст и syncID документа.
	// Возвращает «true» в случае успешного удаления документа.
	DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error)

	// MoveToTrash выполняет запрос на перемещение документа с указанным ID в корзину.
	// Принимает контекст и ID документа.
	// Возвращает «true» в случае успешного перемещения в корзину.
	MoveToTrash(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetStateByID выполняет запрос на получение статуса документа по ID.
	// Принимает контекст и ID статуса.
	// Возвращает найденный статус.
	GetStateByID(ctx context.Context, id uuid.UUID) (*State, *resty.Response, error)

	// CreateState выполняет запрос на создание статуса документа.
	// Принимает контекст и статус.
	// Возвращает созданный статус.
	CreateState(ctx context.Context, state *State) (*State, *resty.Response, error)

	// UpdateState выполняет запрос на изменение статуса документа.
	// Принимает контекст, ID статуса и статус.
	// Возвращает изменённый статус.
	UpdateState(ctx context.Context, id uuid.UUID, state *State) (*State, *resty.Response, error)

	// CreateUpdateStateMany выполняет запрос на массовое создание и/или изменение статусов документа.
	// Принимает контекст и множество статусов.
	// Возвращает список созданных и/или изменённых статусов.
	CreateUpdateStateMany(ctx context.Context, states ...*State) (*Slice[State], *resty.Response, error)

	// DeleteState выполняет запрос на удаление статуса документа.
	// Принимает контекст и ID статуса.
	// Возвращает «true» в случае успешного удаления статуса.
	DeleteState(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetFileList выполняет запрос на получение файлов в виде списка.
	// Принимает контекст и ID сущности/документа.
	// Возвращает объект List.
	GetFileList(ctx context.Context, id uuid.UUID) (*List[File], *resty.Response, error)

	// CreateFile выполняет запрос на добавление файла.
	// Принимает контекст, ID сущности/документа и файл.
	// Возвращает список файлов.
	CreateFile(ctx context.Context, id uuid.UUID, file *File) (*Slice[File], *resty.Response, error)

	// UpdateFileMany выполняет запрос на массовое создание и/или изменение файлов сущности/документа.
	// Принимает контекст, ID сущности/документа и множество файлов.
	// Возвращает созданных и/или изменённых файлов.
	UpdateFileMany(ctx context.Context, id uuid.UUID, files ...*File) (*Slice[File], *resty.Response, error)

	// DeleteFile выполняет запрос на удаление файла сущности/документа.
	// Принимает контекст, ID сущности/документа и ID файла.
	// Возвращает «true» в случае успешного удаления файла.
	DeleteFile(ctx context.Context, id uuid.UUID, fileID uuid.UUID) (bool, *resty.Response, error)

	// DeleteFileMany выполняет запрос на массовое удаление файлов сущности/документа.
	// Принимает контекст, ID сущности/документа и множество файлов.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteFileMany(ctx context.Context, id uuid.UUID, files ...*File) (*DeleteManyResponse, *resty.Response, error)

	// Evaluate выполняет запрос на получение шаблона документа с автозаполнением.
	// Принимает контекст, документ и множество значений Evaluate.
	//
	// Возможные значения типа Evaluate:
	//	- EvaluateDiscount – скидки
	//	- EvaluatePrice    – цены
	//	- EvaluateVat      – ндс
	//	- EvaluateCost     – себестоимость
	// Возвращает шаблон документа с автозаполнением.
	Evaluate(ctx context.Context, entity *InternalOrder, evaluate ...Evaluate) (*InternalOrder, *resty.Response, error)
}

const (
	EndpointInternalOrder = EndpointEntity + string(MetaTypeInternalOrder)
)

// NewInternalOrderService принимает [Client] и возвращает сервис для работы с внутренними заказами.
func NewInternalOrderService(client *Client) InternalOrderService {
	return newMainService[InternalOrder, InternalOrderPosition, MetaAttributesStatesSharedWrapper, any](client, EndpointInternalOrder)
}
