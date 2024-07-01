package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"net/http"
)

// Inventory Инвентаризация.
//
// Код сущности: inventory
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-inwentarizaciq
type Inventory struct {
	Name         *string                       `json:"name,omitempty"`         // Наименование Инвентаризации
	Sum          *float64                      `json:"sum,omitempty"`          // Сумма Инвентаризации в копейках
	Code         *string                       `json:"code,omitempty"`         // Код Инвентаризации
	Created      *Timestamp                    `json:"created,omitempty"`      // Дата создания
	Deleted      *Timestamp                    `json:"deleted,omitempty"`      // Момент последнего удаления Инвентаризации
	Description  *string                       `json:"description,omitempty"`  // Комментарий Инвентаризации
	ExternalCode *string                       `json:"externalCode,omitempty"` // Внешний код Инвентаризации
	Files        *MetaArray[File]              `json:"files,omitempty"`        // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group        *Group                        `json:"group,omitempty"`        // Отдел сотрудника
	ID           *uuid.UUID                    `json:"id,omitempty"`           // ID Инвентаризации
	Updated      *Timestamp                    `json:"updated,omitempty"`      // Момент последнего обновления Инвентаризации
	Meta         *Meta                         `json:"meta,omitempty"`         // Метаданные Инвентаризации
	Owner        *Employee                     `json:"owner,omitempty"`        // Метаданные владельца (Сотрудника)
	Organization *Organization                 `json:"organization,omitempty"` // Метаданные юрлица
	AccountID    *uuid.UUID                    `json:"accountId,omitempty"`    // ID учётной записи
	Positions    *Positions[InventoryPosition] `json:"positions,omitempty"`    // Метаданные позиций Инвентаризации
	Printed      *bool                         `json:"printed,omitempty"`      // Напечатан ли документ
	Published    *bool                         `json:"published,omitempty"`    // Опубликован ли документ
	Shared       *bool                         `json:"shared,omitempty"`       // Общий доступ
	State        *NullValue[State]             `json:"state,omitempty"`        // Метаданные статуса Инвентаризации
	Store        *Store                        `json:"store,omitempty"`        // Метаданные склада
	Moment       *Timestamp                    `json:"moment,omitempty"`       // Дата документа
	SyncID       *uuid.UUID                    `json:"syncId,omitempty"`       // ID синхронизации
	Attributes   Slice[Attribute]              `json:"attributes,omitempty"`   // Список метаданных доп. полей
	Enters       Slice[Enter]                  `json:"enters,omitempty"`       // Список связанных с инвентаризацией оприходований
	Losses       Slice[Loss]                   `json:"losses,omitempty"`       // Список связанных с инвентаризацией списаний
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (inventory Inventory) Clean() *Inventory {
	if inventory.Meta == nil {
		return nil
	}
	return &Inventory{Meta: inventory.Meta}
}

// asTaskOperation реализует интерфейс [TaskOperationInterface].
func (inventory Inventory) asTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: inventory.Meta}
}

// GetName возвращает Наименование Инвентаризации.
func (inventory Inventory) GetName() string {
	return Deref(inventory.Name)
}

// GetSum возвращает Сумму Инвентаризации в копейках.
func (inventory Inventory) GetSum() float64 {
	return Deref(inventory.Sum)
}

// GetCode возвращает Код Инвентаризации.
func (inventory Inventory) GetCode() string {
	return Deref(inventory.Code)
}

// GetCreated возвращает Дату создания.
func (inventory Inventory) GetCreated() Timestamp {
	return Deref(inventory.Created)
}

// GetDeleted возвращает Момент последнего удаления Инвентаризации.
func (inventory Inventory) GetDeleted() Timestamp {
	return Deref(inventory.Deleted)
}

// GetDescription возвращает Комментарий Инвентаризации.
func (inventory Inventory) GetDescription() string {
	return Deref(inventory.Description)
}

// GetExternalCode возвращает Внешний код Инвентаризации.
func (inventory Inventory) GetExternalCode() string {
	return Deref(inventory.ExternalCode)
}

// GetFiles возвращает Метаданные массива Файлов.
func (inventory Inventory) GetFiles() MetaArray[File] {
	return Deref(inventory.Files)
}

// GetGroup возвращает Отдел сотрудника.
func (inventory Inventory) GetGroup() Group {
	return Deref(inventory.Group)
}

// GetID возвращает ID Инвентаризации.
func (inventory Inventory) GetID() uuid.UUID {
	return Deref(inventory.ID)
}

// GetUpdated возвращает Момент последнего обновления Инвентаризации.
func (inventory Inventory) GetUpdated() Timestamp {
	return Deref(inventory.Updated)
}

// GetMeta возвращает Метаданные Инвентаризации.
func (inventory Inventory) GetMeta() Meta {
	return Deref(inventory.Meta)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (inventory Inventory) GetOwner() Employee {
	return Deref(inventory.Owner)
}

// GetOrganization возвращает Метаданные юрлица.
func (inventory Inventory) GetOrganization() Organization {
	return Deref(inventory.Organization)
}

// GetAccountID возвращает ID учётной записи.
func (inventory Inventory) GetAccountID() uuid.UUID {
	return Deref(inventory.AccountID)
}

// GetPositions возвращает Метаданные позиций Инвентаризации.
func (inventory Inventory) GetPositions() Positions[InventoryPosition] {
	return Deref(inventory.Positions)
}

// GetPrinted возвращает true, если документ напечатан.
func (inventory Inventory) GetPrinted() bool {
	return Deref(inventory.Printed)
}

// GetPublished возвращает true, если документ опубликован.
func (inventory Inventory) GetPublished() bool {
	return Deref(inventory.Published)
}

// GetShared возвращает флаг общего доступа.
func (inventory Inventory) GetShared() bool {
	return Deref(inventory.Shared)
}

// GetState возвращает Метаданные статуса Инвентаризации.
func (inventory Inventory) GetState() State {
	return inventory.State.Get()
}

// GetStore возвращает Метаданные склада.
func (inventory Inventory) GetStore() Store {
	return Deref(inventory.Store)
}

// GetMoment возвращает Дату документа.
func (inventory Inventory) GetMoment() Timestamp {
	return Deref(inventory.Moment)
}

// GetSyncID возвращает ID синхронизации.
func (inventory Inventory) GetSyncID() uuid.UUID {
	return Deref(inventory.SyncID)
}

// GetAttributes возвращает Список метаданных доп. полей.
func (inventory Inventory) GetAttributes() Slice[Attribute] {
	return inventory.Attributes
}

// GetEnters возвращает Список связанных с инвентаризацией оприходований.
func (inventory Inventory) GetEnters() Slice[Enter] {
	return inventory.Enters
}

// GetLosses возвращает Список связанных с инвентаризацией списаний.
func (inventory Inventory) GetLosses() Slice[Loss] {
	return inventory.Losses
}

// SetName устанавливает Наименование Инвентаризации.
func (inventory *Inventory) SetName(name string) *Inventory {
	inventory.Name = &name
	return inventory
}

// SetCode устанавливает Код Инвентаризации.
func (inventory *Inventory) SetCode(code string) *Inventory {
	inventory.Code = &code
	return inventory
}

// SetDescription устанавливает Комментарий Инвентаризации.
func (inventory *Inventory) SetDescription(description string) *Inventory {
	inventory.Description = &description
	return inventory
}

// SetExternalCode устанавливает Внешний код Инвентаризации.
func (inventory *Inventory) SetExternalCode(externalCode string) *Inventory {
	inventory.ExternalCode = &externalCode
	return inventory
}

// SetFiles устанавливает Метаданные массива Файлов.
//
// Принимает множество объектов [File].
func (inventory *Inventory) SetFiles(files ...*File) *Inventory {
	inventory.Files = NewMetaArrayFrom(files)
	return inventory
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (inventory *Inventory) SetGroup(group *Group) *Inventory {
	if group != nil {
		inventory.Group = group.Clean()
	}
	return inventory
}

// SetMeta устанавливает Метаданные Инвентаризации.
func (inventory *Inventory) SetMeta(meta *Meta) *Inventory {
	inventory.Meta = meta
	return inventory
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (inventory *Inventory) SetOwner(owner *Employee) *Inventory {
	if owner != nil {
		inventory.Owner = owner.Clean()
	}
	return inventory
}

// SetOrganization устанавливает Метаданные юрлица.
func (inventory *Inventory) SetOrganization(organization *Organization) *Inventory {
	if organization != nil {
		inventory.Organization = organization.Clean()
	}
	return inventory
}

// SetPositions устанавливает Метаданные позиций Инвентаризации.
//
// Принимает множество объектов [InventoryPosition].
func (inventory *Inventory) SetPositions(positions ...*InventoryPosition) *Inventory {
	inventory.Positions = NewPositionsFrom(positions)
	return inventory
}

// SetShared устанавливает флаг общего доступа.
func (inventory *Inventory) SetShared(shared bool) *Inventory {
	inventory.Shared = &shared
	return inventory
}

// SetState устанавливает Метаданные статуса Инвентаризации.
//
// Передача nil передаёт сброс значения (null).
func (inventory *Inventory) SetState(state *State) *Inventory {
	inventory.State = NewNullValue(state)
	return inventory
}

// SetStore устанавливает Метаданные склада.
func (inventory *Inventory) SetStore(store *Store) *Inventory {
	if store != nil {
		inventory.Store = store.Clean()
	}
	return inventory
}

// SetMoment устанавливает Дату документа.
func (inventory *Inventory) SetMoment(moment *Timestamp) *Inventory {
	inventory.Moment = moment
	return inventory
}

// SetSyncID устанавливает ID синхронизации.
func (inventory *Inventory) SetSyncID(syncID uuid.UUID) *Inventory {
	inventory.SyncID = &syncID
	return inventory
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (inventory *Inventory) SetAttributes(attributes ...*Attribute) *Inventory {
	inventory.Attributes.Push(attributes...)
	return inventory
}

// SetEnters устанавливает Список связанных с инвентаризацией оприходований.
//
// Принимает множество объектов [Enter].
func (inventory *Inventory) SetEnters(enters ...*Enter) *Inventory {
	inventory.Enters.Push(enters...)
	return inventory
}

// SetLosses устанавливает Список связанных с инвентаризацией списаний.
//
// Принимает множество объектов [Loss].
func (inventory *Inventory) SetLosses(losses ...*Loss) *Inventory {
	inventory.Losses.Push(losses...)
	return inventory
}

// String реализует интерфейс [fmt.Stringer].
func (inventory Inventory) String() string {
	return Stringify(inventory)
}

// MetaType возвращает код сущности.
func (Inventory) MetaType() MetaType {
	return MetaTypeInventory
}

// Update shortcut
func (inventory Inventory) Update(ctx context.Context, client *Client, params ...*Params) (*Inventory, *resty.Response, error) {
	return NewInventoryService(client).Update(ctx, inventory.GetID(), &inventory, params...)
}

// Create shortcut
func (inventory Inventory) Create(ctx context.Context, client *Client, params ...*Params) (*Inventory, *resty.Response, error) {
	return NewInventoryService(client).Create(ctx, &inventory, params...)
}

// Delete shortcut
func (inventory Inventory) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewInventoryService(client).Delete(ctx, inventory.GetID())
}

// InventoryPosition Позиция Инвентаризации.
//
// Код сущности: inventoryposition
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-inwentarizaciq-inwentarizaciq-pozicii-inwentarizacii
type InventoryPosition struct {
	AccountID          *uuid.UUID          `json:"accountId,omitempty"`          // ID учётной записи
	Assortment         *AssortmentPosition `json:"assortment,omitempty"`         // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	CalculatedQuantity *float64            `json:"calculatedQuantity,omitempty"` // расчетный остаток
	CorrectionAmount   *float64            `json:"correctionAmount,omitempty"`   // разница между расчетным остатком и фактическим
	CorrectionSum      *float64            `json:"correctionSum,omitempty"`      // избыток/недостача
	ID                 *uuid.UUID          `json:"id,omitempty"`                 // ID сущности
	Pack               *Pack               `json:"pack,omitempty"`               // Упаковка Товара
	Price              *float64            `json:"price,omitempty"`              // Цена товара/услуги в копейках
	Quantity           *float64            `json:"quantity,omitempty"`           // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
}

// GetAccountID возвращает ID учётной записи.
func (inventoryPosition InventoryPosition) GetAccountID() uuid.UUID {
	return Deref(inventoryPosition.AccountID)
}

// GetAssortment возвращает Метаданные товара/услуги, которую представляет собой компонент.
func (inventoryPosition InventoryPosition) GetAssortment() AssortmentPosition {
	return Deref(inventoryPosition.Assortment)
}

// GetCalculatedQuantity возвращает расчетный остаток.
func (inventoryPosition InventoryPosition) GetCalculatedQuantity() float64 {
	return Deref(inventoryPosition.CalculatedQuantity)
}

// GetCorrectionAmount возвращает разницу между расчетным остатком и фактическим.
func (inventoryPosition InventoryPosition) GetCorrectionAmount() float64 {
	return Deref(inventoryPosition.CorrectionAmount)
}

// GetCorrectionSum возвращает избыток/недостачу
func (inventoryPosition InventoryPosition) GetCorrectionSum() float64 {
	return Deref(inventoryPosition.CorrectionSum)
}

// GetID возвращает ID позиции.
func (inventoryPosition InventoryPosition) GetID() uuid.UUID {
	return Deref(inventoryPosition.ID)
}

// GetPack возвращает Упаковку Товара.
func (inventoryPosition InventoryPosition) GetPack() Pack {
	return Deref(inventoryPosition.Pack)
}

// GetPrice возвращает Цену товара/услуги в копейках.
func (inventoryPosition InventoryPosition) GetPrice() float64 {
	return Deref(inventoryPosition.Price)
}

// GetQuantity возвращает Количество товаров данного вида в позиции.
//
// Если позиция - товар, у которого включен учет по серийным номерам,
// то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
func (inventoryPosition InventoryPosition) GetQuantity() float64 {
	return Deref(inventoryPosition.Quantity)
}

// SetAssortment устанавливает Метаданные товара/услуги, которую представляет собой компонент.
//
// Принимает объект, реализующий интерфейс [AssortmentInterface].
func (inventoryPosition *InventoryPosition) SetAssortment(assortment AssortmentInterface) *InventoryPosition {
	if assortment != nil {
		inventoryPosition.Assortment = assortment.asAssortment()
	}
	return inventoryPosition
}

// SetCalculatedQuantity устанавливает расчетный остаток.
func (inventoryPosition *InventoryPosition) SetCalculatedQuantity(calculatedQuantity float64) *InventoryPosition {
	inventoryPosition.CalculatedQuantity = &calculatedQuantity
	return inventoryPosition
}

// SetPack устанавливает Упаковку Товара.
func (inventoryPosition *InventoryPosition) SetPack(pack *Pack) *InventoryPosition {
	if pack != nil {
		inventoryPosition.Pack = pack
	}
	return inventoryPosition
}

// SetPrice устанавливает Цену товара/услуги в копейках.
func (inventoryPosition *InventoryPosition) SetPrice(price float64) *InventoryPosition {
	inventoryPosition.Price = &price
	return inventoryPosition
}

// SetQuantity устанавливает Количество товаров данного вида в позиции.
//
// Если позиция - товар, у которого включен учет по серийным номерам,
// то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
func (inventoryPosition *InventoryPosition) SetQuantity(quantity float64) *InventoryPosition {
	inventoryPosition.Quantity = &quantity
	return inventoryPosition
}

// String реализует интерфейс [fmt.Stringer].
func (inventoryPosition InventoryPosition) String() string {
	return Stringify(inventoryPosition)
}

// MetaType возвращает код сущности.
func (InventoryPosition) MetaType() MetaType {
	return MetaTypeInventoryPosition
}

// InventoryService описывает методы сервиса для работы с инвентаризациями.
type InventoryService interface {
	// GetList выполняет запрос на получение списка инвентаризаций.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[Inventory], *resty.Response, error)

	// Create выполняет запрос на создание инвентаризации.
	// Обязательные поля для заполнения:
	//	- organization (Ссылка на ваше юрлицо)
	//	- store (Ссылка на склад)
	// Принимает контекст, инвентаризацию и опционально объект параметров запроса Params.
	// Возвращает созданную инвентаризацию.
	Create(ctx context.Context, inventory *Inventory, params ...*Params) (*Inventory, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение инвентаризаций.
	// Изменяемые инвентаризации должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список инвентаризации и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых инвентаризаций.
	CreateUpdateMany(ctx context.Context, inventoryList Slice[Inventory], params ...*Params) (*Slice[Inventory], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление инвентаризаций.
	// Принимает контекст и множество инвентаризаций.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*Inventory) (*DeleteManyResponse, *resty.Response, error)

	// Delete выполняет запрос на удаление инвентаризации.
	// Принимает контекст и ID инвентаризации.
	// Возвращает true в случае успешного удаления инвентаризации.
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельного инвентаризации по ID.
	// Принимает контекст, ID инвентаризации и опционально объект параметров запроса Params.
	// Возвращает найденную инвентаризацию.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Inventory, *resty.Response, error)

	// Update выполняет запрос на изменение инвентаризации.
	// Принимает контекст, инвентаризацию и опционально объект параметров запроса Params.
	// Возвращает изменённую инвентаризацию.
	Update(ctx context.Context, id uuid.UUID, inventory *Inventory, params ...*Params) (*Inventory, *resty.Response, error)

	// Template выполняет запрос на получение предзаполненной инвентаризации со стандартными полями без связи с какими-либо другими документами.
	// Принимает контекст.
	// Возвращает предзаполненную инвентаризацию.
	Template(ctx context.Context) (*Inventory, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных инвентаризаций.
	// Принимает контекст.
	// Возвращает объект метаданных MetaAttributesSharedStatesWrapper.
	GetMetadata(ctx context.Context) (*MetaAttributesSharedStatesWrapper, *resty.Response, error)

	// GetPositionList выполняет запрос на получение списка позиций документа.
	// Принимает контекст, ID документа и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetPositionList(ctx context.Context, id uuid.UUID, params ...*Params) (*List[InventoryPosition], *resty.Response, error)

	// GetPositionByID выполняет запрос на получение отдельной позиции документа по ID.
	// Принимает контекст, ID документа, ID позиции и опционально объект параметров запроса Params.
	// Возвращает найденную позицию.
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params ...*Params) (*InventoryPosition, *resty.Response, error)

	// UpdatePosition выполняет запрос на изменение позиции документа.
	// Принимает контекст, ID документа, ID позиции, позицию документа и опционально объект параметров запроса Params.
	// Возвращает изменённую позицию.
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *InventoryPosition, params ...*Params) (*InventoryPosition, *resty.Response, error)

	// CreatePosition выполняет запрос на добавление позиции документа.
	// Принимает контекст, ID документа, позицию документа и опционально объект параметров запроса Params.
	// Возвращает добавленную позицию.
	CreatePosition(ctx context.Context, id uuid.UUID, position *InventoryPosition, params ...*Params) (*InventoryPosition, *resty.Response, error)

	// CreatePositionMany выполняет запрос на массовое добавление позиций документа.
	// Принимает контекст, ID документа и множество позиций.
	// Возвращает список добавленных позиций.
	CreatePositionMany(ctx context.Context, id uuid.UUID, positions ...*InventoryPosition) (*Slice[InventoryPosition], *resty.Response, error)

	// DeletePosition выполняет запрос на удаление позиции документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает true в случае успешного удаления позиции.
	DeletePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (bool, *resty.Response, error)

	// DeletePositionMany выполняет запрос на массовое удаление позиций документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeletePositionMany(ctx context.Context, id uuid.UUID, positions ...*InventoryPosition) (*DeleteManyResponse, *resty.Response, error)

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
	// Возвращает true в случае успешного удаления доп поля.
	DeleteAttribute(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// DeleteAttributeMany выполняет запрос на массовое удаление доп полей.
	// Принимает контекст и множество доп полей.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteAttributeMany(ctx context.Context, attributes ...*Attribute) (*DeleteManyResponse, *resty.Response, error)

	// GetBySyncID выполняет запрос на получение отдельного документа по syncID.
	// Принимает контекст и syncID документа.
	// Возвращает найденный документ.
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*Inventory, *resty.Response, error)

	// DeleteBySyncID выполняет запрос на удаление документа по syncID.
	// Принимает контекст и syncID документа.
	// Возвращает true в случае успешного удаления документа.
	DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error)

	// MoveToTrash выполняет запрос на перемещение документа с указанным ID в корзину.
	// Принимает контекст и ID документа.
	// Возвращает true в случае успешного перемещения в корзину.
	MoveToTrash(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// Recalculate выполняет запрос на пересчёт расчётных остатков у позиций инвентаризации.
	// Принимает контекст и ID инвентаризации.
	// Возвращает true в случае успешного перерасчёта.
	Recalculate(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

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
	// Возвращает true в случае успешного удаления файла.
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
	Evaluate(ctx context.Context, entity *Inventory, evaluate ...Evaluate) (*Inventory, *resty.Response, error)
}

type inventoryService struct {
	Endpoint
	endpointGetList[Inventory]
	endpointCreate[Inventory]
	endpointCreateUpdateMany[Inventory]
	endpointDeleteMany[Inventory]
	endpointDelete
	endpointGetByID[Inventory]
	endpointUpdate[Inventory]
	endpointTemplate[Inventory]
	endpointMetadata[MetaAttributesSharedStatesWrapper]
	endpointPositions[InventoryPosition]
	endpointAttributes
	endpointSyncID[Inventory]
	endpointTrash
	endpointStates
	endpointFiles
	endpointEvaluate[Inventory]
}

// NewInventoryService принимает [Client] и возвращает сервис для работы с инвентаризациями.
func NewInventoryService(client *Client) InventoryService {
	e := NewEndpoint(client, "entity/inventory")
	return &inventoryService{
		Endpoint:                 e,
		endpointGetList:          endpointGetList[Inventory]{e},
		endpointCreate:           endpointCreate[Inventory]{e},
		endpointCreateUpdateMany: endpointCreateUpdateMany[Inventory]{e},
		endpointDeleteMany:       endpointDeleteMany[Inventory]{e},
		endpointDelete:           endpointDelete{e},
		endpointGetByID:          endpointGetByID[Inventory]{e},
		endpointUpdate:           endpointUpdate[Inventory]{e},
		endpointTemplate:         endpointTemplate[Inventory]{e},
		endpointMetadata:         endpointMetadata[MetaAttributesSharedStatesWrapper]{e},
		endpointPositions:        endpointPositions[InventoryPosition]{e},
		endpointAttributes:       endpointAttributes{e},
		endpointSyncID:           endpointSyncID[Inventory]{e},
		endpointTrash:            endpointTrash{e},
		endpointStates:           endpointStates{e},
		endpointFiles:            endpointFiles{e},
		endpointEvaluate:         endpointEvaluate[Inventory]{e},
	}
}

func (service *inventoryService) Recalculate(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf("rpc/inventory/%s/recalcCalculatedQuantity", id)
	_, resp, err := NewRequestBuilder[any](service.client, path).Put(ctx, nil)
	if err != nil {
		return false, resp, err
	}
	return resp.StatusCode() == http.StatusCreated, resp, nil
}
