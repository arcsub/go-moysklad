package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"time"
)

// ProcessingOrder Заказ на производство.
//
// Код сущности: processingorder
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-zakaz-na-proizwodstwo
type ProcessingOrder struct {
	Name                  *string                             `json:"name,omitempty"`                  // Наименование Заказа на производство
	Published             *bool                               `json:"published,omitempty"`             // Опубликован ли документ
	Organization          *Organization                       `json:"organization,omitempty"`          // Метаданные юрлица
	Code                  *string                             `json:"code,omitempty"`                  // Код Заказа на производство
	Created               *Timestamp                          `json:"created,omitempty"`               // Дата создания
	Deleted               *Timestamp                          `json:"deleted,omitempty"`               // Момент последнего удаления Заказа на производство
	DeliveryPlannedMoment *Timestamp                          `json:"deliveryPlannedMoment,omitempty"` // Планируемая дата производства
	Description           *string                             `json:"description,omitempty"`           // Комментарий Заказа на производство
	ExternalCode          *string                             `json:"externalCode,omitempty"`          // Внешний код Заказа на производство
	Files                 *MetaArray[File]                    `json:"files,omitempty"`                 // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group                 *Group                              `json:"group,omitempty"`                 // Отдел сотрудника
	ID                    *uuid.UUID                          `json:"id,omitempty"`                    // ID Заказа на производство
	Meta                  *Meta                               `json:"meta,omitempty"`                  // Метаданные Заказа на производство
	Moment                *Timestamp                          `json:"moment,omitempty"`                // Дата документа
	AccountID             *uuid.UUID                          `json:"accountId,omitempty"`             // ID учётной записи
	Owner                 *Employee                           `json:"owner,omitempty"`                 // Метаданные владельца (Сотрудника)
	Positions             *MetaArray[ProcessingOrderPosition] `json:"positions,omitempty"`             // Метаданные позиций Заказа на производство
	Printed               *bool                               `json:"printed,omitempty"`               // Напечатан ли документ
	ProcessingPlan        *ProcessingPlan                     `json:"processingPlan,omitempty"`        // Метаданные Техкарты
	Project               *NullValue[Project]                 `json:"project,omitempty"`               // Метаданные проекта
	Applicable            *bool                               `json:"applicable,omitempty"`            // Отметка о проведении
	Quantity              *float64                            `json:"quantity,omitempty"`              // Объем производства
	Shared                *bool                               `json:"shared,omitempty"`                // Общий доступ
	State                 *NullValue[State]                   `json:"state,omitempty"`                 // Метаданные статуса Заказа на производство
	Store                 *Store                              `json:"store,omitempty"`                 // Метаданные склада
	SyncID                *uuid.UUID                          `json:"syncId,omitempty"`                // ID синхронизации
	Updated               *Timestamp                          `json:"updated,omitempty"`               // Момент последнего обновления Заказа на производство
	Processings           Slice[Processing]                   `json:"processings,omitempty"`           // Массив ссылок на связанные техоперации
	Attributes            Slice[Attribute]                    `json:"attributes,omitempty"`            // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (processingOrder ProcessingOrder) Clean() *ProcessingOrder {
	if processingOrder.Meta == nil {
		return nil
	}
	return &ProcessingOrder{Meta: processingOrder.Meta}
}

// AsTaskOperation реализует интерфейс [TaskOperationConverter].
func (processingOrder ProcessingOrder) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: processingOrder.Meta}
}

// GetName возвращает Наименование Заказа на производство.
func (processingOrder ProcessingOrder) GetName() string {
	return Deref(processingOrder.Name)
}

// GetPublished возвращает true, если документ опубликован.
func (processingOrder ProcessingOrder) GetPublished() bool {
	return Deref(processingOrder.Published)
}

// GetOrganization возвращает Метаданные юрлица.
func (processingOrder ProcessingOrder) GetOrganization() Organization {
	return Deref(processingOrder.Organization)
}

// GetCode возвращает Код Заказа на производство.
func (processingOrder ProcessingOrder) GetCode() string {
	return Deref(processingOrder.Code)
}

// GetCreated возвращает Дату создания.
func (processingOrder ProcessingOrder) GetCreated() time.Time {
	return Deref(processingOrder.Created).Time()
}

// GetDeleted возвращает Момент последнего удаления Заказа на производство.
func (processingOrder ProcessingOrder) GetDeleted() time.Time {
	return Deref(processingOrder.Deleted).Time()
}

// GetDeliveryPlannedMoment возвращает Планируемую дата производства.
func (processingOrder ProcessingOrder) GetDeliveryPlannedMoment() time.Time {
	return Deref(processingOrder.DeliveryPlannedMoment).Time()
}

// GetDescription возвращает Комментарий Заказа на производство.
func (processingOrder ProcessingOrder) GetDescription() string {
	return Deref(processingOrder.Description)
}

// GetExternalCode возвращает Внешний код Заказа на производство.
func (processingOrder ProcessingOrder) GetExternalCode() string {
	return Deref(processingOrder.ExternalCode)
}

// GetFiles возвращает Метаданные массива Файлов.
func (processingOrder ProcessingOrder) GetFiles() MetaArray[File] {
	return Deref(processingOrder.Files)
}

// GetGroup возвращает Отдел сотрудника.
func (processingOrder ProcessingOrder) GetGroup() Group {
	return Deref(processingOrder.Group)
}

// GetID возвращает ID Заказа на производство.
func (processingOrder ProcessingOrder) GetID() uuid.UUID {
	return Deref(processingOrder.ID)
}

// GetMeta возвращает Метаданные Заказа на производство.
func (processingOrder ProcessingOrder) GetMeta() Meta {
	return Deref(processingOrder.Meta)
}

// GetMoment возвращает Дату документа.
func (processingOrder ProcessingOrder) GetMoment() time.Time {
	return Deref(processingOrder.Moment).Time()
}

// GetAccountID возвращает ID учётной записи.
func (processingOrder ProcessingOrder) GetAccountID() uuid.UUID {
	return Deref(processingOrder.AccountID)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (processingOrder ProcessingOrder) GetOwner() Employee {
	return Deref(processingOrder.Owner)
}

// GetPositions возвращает Метаданные позиций Заказа на производство.
func (processingOrder ProcessingOrder) GetPositions() MetaArray[ProcessingOrderPosition] {
	return Deref(processingOrder.Positions)
}

// GetPrinted возвращает true, если документ напечатан.
func (processingOrder ProcessingOrder) GetPrinted() bool {
	return Deref(processingOrder.Printed)
}

// GetProcessingPlan возвращает Метаданные Техкарты.
func (processingOrder ProcessingOrder) GetProcessingPlan() ProcessingPlan {
	return Deref(processingOrder.ProcessingPlan)
}

// GetProject возвращает Метаданные проекта.
func (processingOrder ProcessingOrder) GetProject() Project {
	return Deref(processingOrder.Project).getValue()
}

// GetApplicable возвращает Отметку о проведении.
func (processingOrder ProcessingOrder) GetApplicable() bool {
	return Deref(processingOrder.Applicable)
}

// GetQuantity возвращает Объем производства.
func (processingOrder ProcessingOrder) GetQuantity() float64 {
	return Deref(processingOrder.Quantity)
}

// GetShared возвращает флаг Общего доступа.
func (processingOrder ProcessingOrder) GetShared() bool {
	return Deref(processingOrder.Shared)
}

// GetState возвращает Метаданные статуса Заказа на производство.
func (processingOrder ProcessingOrder) GetState() State {
	return Deref(processingOrder.State).getValue()
}

// GetStore возвращает Метаданные склада.
func (processingOrder ProcessingOrder) GetStore() Store {
	return Deref(processingOrder.Store)
}

// GetSyncID возвращает ID синхронизации.
func (processingOrder ProcessingOrder) GetSyncID() uuid.UUID {
	return Deref(processingOrder.SyncID)
}

// GetUpdated возвращает Момент последнего обновления Заказа на производство.
func (processingOrder ProcessingOrder) GetUpdated() time.Time {
	return Deref(processingOrder.Updated).Time()
}

// GetProcessings возвращает Массив ссылок на связанные техоперации.
func (processingOrder ProcessingOrder) GetProcessings() Slice[Processing] {
	return processingOrder.Processings
}

// GetAttributes возвращает Список метаданных доп. полей.
func (processingOrder ProcessingOrder) GetAttributes() Slice[Attribute] {
	return processingOrder.Attributes
}

// SetName устанавливает Наименование Заказа на производство.
func (processingOrder *ProcessingOrder) SetName(name string) *ProcessingOrder {
	processingOrder.Name = &name
	return processingOrder
}

// SetOrganization устанавливает Метаданные юрлица.
func (processingOrder *ProcessingOrder) SetOrganization(organization *Organization) *ProcessingOrder {
	if organization != nil {
		processingOrder.Organization = organization.Clean()
	}
	return processingOrder
}

// SetCode устанавливает Код Заказа на производство.
func (processingOrder *ProcessingOrder) SetCode(code string) *ProcessingOrder {
	processingOrder.Code = &code
	return processingOrder
}

// SetDeliveryPlannedMoment устанавливает Планируемую дату производства.
func (processingOrder *ProcessingOrder) SetDeliveryPlannedMoment(deliveryPlannedMoment time.Time) *ProcessingOrder {
	processingOrder.DeliveryPlannedMoment = NewTimestamp(deliveryPlannedMoment)
	return processingOrder
}

// SetDescription устанавливает Комментарий Заказа на производство.
func (processingOrder *ProcessingOrder) SetDescription(description string) *ProcessingOrder {
	processingOrder.Description = &description
	return processingOrder
}

// SetExternalCode устанавливает Внешний код Заказа на производство.
func (processingOrder *ProcessingOrder) SetExternalCode(externalCode string) *ProcessingOrder {
	processingOrder.ExternalCode = &externalCode
	return processingOrder
}

// SetFiles устанавливает Метаданные массива Файлов.
//
// Принимает множество объектов [File].
func (processingOrder *ProcessingOrder) SetFiles(files ...*File) *ProcessingOrder {
	processingOrder.Files = NewMetaArrayFrom(files)
	return processingOrder
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (processingOrder *ProcessingOrder) SetGroup(group *Group) *ProcessingOrder {
	if group != nil {
		processingOrder.Group = group.Clean()
	}
	return processingOrder
}

// SetMeta устанавливает Метаданные Заказа на производство.
func (processingOrder *ProcessingOrder) SetMeta(meta *Meta) *ProcessingOrder {
	processingOrder.Meta = meta
	return processingOrder
}

// SetMoment устанавливает Дату документа.
func (processingOrder *ProcessingOrder) SetMoment(moment time.Time) *ProcessingOrder {
	processingOrder.Moment = NewTimestamp(moment)
	return processingOrder
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (processingOrder *ProcessingOrder) SetOwner(owner *Employee) *ProcessingOrder {
	if owner != nil {
		processingOrder.Owner = owner.Clean()
	}
	return processingOrder
}

// SetPositions устанавливает Метаданные позиций Заказа на производство.
//
// Принимает множество объектов [ProcessingOrderPosition].
func (processingOrder *ProcessingOrder) SetPositions(positions ...*ProcessingOrderPosition) *ProcessingOrder {
	processingOrder.Positions = NewMetaArrayFrom(positions)
	return processingOrder
}

// SetProcessingPlan устанавливает Метаданные Техкарты.
func (processingOrder *ProcessingOrder) SetProcessingPlan(processingPlan *ProcessingPlan) *ProcessingOrder {
	if processingPlan != nil {
		processingOrder.ProcessingPlan = processingPlan.Clean()
	}
	return processingOrder
}

// SetProject устанавливает Метаданные проекта.
//
// Передача nil передаёт сброс значения (null).
func (processingOrder *ProcessingOrder) SetProject(project *Project) *ProcessingOrder {
	processingOrder.Project = NewNullValue(project)
	return processingOrder
}

// SetApplicable устанавливает Отметку о проведении.
func (processingOrder *ProcessingOrder) SetApplicable(applicable bool) *ProcessingOrder {
	processingOrder.Applicable = &applicable
	return processingOrder
}

// SetQuantity устанавливает Объем производства.
func (processingOrder *ProcessingOrder) SetQuantity(quantity float64) *ProcessingOrder {
	processingOrder.Quantity = &quantity
	return processingOrder
}

// SetShared устанавливает флаг общего доступа.
func (processingOrder *ProcessingOrder) SetShared(shared bool) *ProcessingOrder {
	processingOrder.Shared = &shared
	return processingOrder
}

// SetState устанавливает Метаданные статуса Заказа на производство.
//
// Передача nil передаёт сброс значения (null).
func (processingOrder *ProcessingOrder) SetState(state *State) *ProcessingOrder {
	processingOrder.State = NewNullValue(state)
	return processingOrder
}

// SetStore устанавливает Метаданные склада.
func (processingOrder *ProcessingOrder) SetStore(store *Store) *ProcessingOrder {
	if store != nil {
		processingOrder.Store = store.Clean()
	}
	return processingOrder
}

// SetSyncID устанавливает ID синхронизации.
func (processingOrder *ProcessingOrder) SetSyncID(syncID uuid.UUID) *ProcessingOrder {
	processingOrder.SyncID = &syncID
	return processingOrder
}

// SetProcessings устанавливает Массив ссылок на связанные техоперации.
//
// Принимает множество объектов [Processing].
func (processingOrder *ProcessingOrder) SetProcessings(processings ...*Processing) *ProcessingOrder {
	processingOrder.Processings.Push(processings...)
	return processingOrder
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (processingOrder *ProcessingOrder) SetAttributes(attributes ...*Attribute) *ProcessingOrder {
	processingOrder.Attributes.Push(attributes...)
	return processingOrder
}

// String реализует интерфейс [fmt.Stringer].
func (processingOrder ProcessingOrder) String() string {
	return Stringify(processingOrder)
}

// MetaType возвращает код сущности.
func (ProcessingOrder) MetaType() MetaType {
	return MetaTypeProcessingOrder
}

// Update shortcut
func (processingOrder ProcessingOrder) Update(ctx context.Context, client *Client, params ...*Params) (*ProcessingOrder, *resty.Response, error) {
	return NewProcessingOrderService(client).Update(ctx, processingOrder.GetID(), &processingOrder, params...)
}

// Create shortcut
func (processingOrder ProcessingOrder) Create(ctx context.Context, client *Client, params ...*Params) (*ProcessingOrder, *resty.Response, error) {
	return NewProcessingOrderService(client).Create(ctx, &processingOrder, params...)
}

// Delete shortcut
func (processingOrder ProcessingOrder) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewProcessingOrderService(client).Delete(ctx, processingOrder.GetID())
}

// ProcessingOrderPosition Позиция Заказа на производство.
//
// Код сущности: processingorderposition
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-zakaz-na-proizwodstwo-zakazy-na-proizwodstwo-pozicii-zakaza-na-proizwodstwo
type ProcessingOrderPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учётной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Reserve    *float64            `json:"reserve,omitempty"`    // Резерв данной позиции
}

// GetAccountID возвращает ID учётной записи.
func (processingOrderPosition ProcessingOrderPosition) GetAccountID() uuid.UUID {
	return Deref(processingOrderPosition.AccountID)
}

// GetAssortment возвращает Метаданные товара/услуги/серии/модификации, которую представляет собой позиция.
func (processingOrderPosition ProcessingOrderPosition) GetAssortment() AssortmentPosition {
	return Deref(processingOrderPosition.Assortment)
}

// GetID возвращает ID позиции.
func (processingOrderPosition ProcessingOrderPosition) GetID() uuid.UUID {
	return Deref(processingOrderPosition.ID)
}

// GetPack возвращает Упаковку Товара.
func (processingOrderPosition ProcessingOrderPosition) GetPack() Pack {
	return Deref(processingOrderPosition.Pack)
}

// GetQuantity возвращает Количество товаров/услуг данного вида в позиции.
//
// Если позиция - товар, у которого включен учет по серийным номерам,
// то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
func (processingOrderPosition ProcessingOrderPosition) GetQuantity() float64 {
	return Deref(processingOrderPosition.Quantity)
}

// GetReserve возвращает Резерв данной позиции.
func (processingOrderPosition ProcessingOrderPosition) GetReserve() float64 {
	return Deref(processingOrderPosition.Reserve)
}

// SetAssortment устанавливает Метаданные товара/услуги/серии/модификации, которую представляет собой позиция.
//
// Принимает объект, реализующий интерфейс [AssortmentConverter].
func (processingOrderPosition *ProcessingOrderPosition) SetAssortment(assortment AssortmentConverter) *ProcessingOrderPosition {
	if assortment != nil {
		processingOrderPosition.Assortment = assortment.AsAssortment()
	}
	return processingOrderPosition
}

// SetPack устанавливает Упаковку Товара.
func (processingOrderPosition *ProcessingOrderPosition) SetPack(pack *Pack) *ProcessingOrderPosition {
	if pack != nil {
		processingOrderPosition.Pack = pack
	}
	return processingOrderPosition
}

// SetQuantity устанавливает Количество товаров/услуг данного вида в позиции.
//
// Если позиция - товар, у которого включен учет по серийным номерам,
// то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
func (processingOrderPosition *ProcessingOrderPosition) SetQuantity(quantity float64) *ProcessingOrderPosition {
	processingOrderPosition.Quantity = &quantity
	return processingOrderPosition
}

// SetReserve устанавливает Резерв данной позиции.
func (processingOrderPosition *ProcessingOrderPosition) SetReserve(reserve float64) *ProcessingOrderPosition {
	processingOrderPosition.Reserve = &reserve
	return processingOrderPosition
}

// String реализует интерфейс [fmt.Stringer].
func (processingOrderPosition ProcessingOrderPosition) String() string {
	return Stringify(processingOrderPosition)
}

// MetaType возвращает код сущности.
func (ProcessingOrderPosition) MetaType() MetaType {
	return MetaTypeProcessingOrderPosition
}

// ProcessingOrderService описывает методы сервиса для работы с заказами на производство.
type ProcessingOrderService interface {
	// GetList выполняет запрос на получение списка заказов на производство.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[ProcessingOrder], *resty.Response, error)

	// Create выполняет запрос на создание заказа на производство.
	// Обязательные поля для заполнения:
	//	- organization (Ссылка на ваше юрлицо)
	//	- processingPlan (Ссылка на техкарту)
	//	- positions (Ссылка на позиции в Заказе)
	// Принимает контекст, заказ на производство и опционально объект параметров запроса Params.
	// Возвращает созданный заказ на производство.
	Create(ctx context.Context, processingOrder *ProcessingOrder, params ...*Params) (*ProcessingOrder, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение заказов на производство.
	// Изменяемые заказы на производство должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список заказов на производство и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых заказов на производство.
	CreateUpdateMany(ctx context.Context, processingOrderList Slice[ProcessingOrder], params ...*Params) (*Slice[ProcessingOrder], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление заказов на производство.
	// Принимает контекст и множество заказов на производство.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*ProcessingOrder) (*DeleteManyResponse, *resty.Response, error)

	// Delete выполняет запрос на удаление заказа на производство.
	// Принимает контекст и ID заказа на производство.
	// Возвращает «true» в случае успешного удаления заказа на производство.
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение заказа на производство по ID.
	// Принимает контекст, ID заказа на производство и опционально объект параметров запроса Params.
	// Возвращает заказ на производство.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*ProcessingOrder, *resty.Response, error)

	// Update выполняет запрос на изменение заказа на производство.
	// Принимает контекст, заказ на производство и опционально объект параметров запроса Params.
	// Возвращает изменённый заказа на производство.
	Update(ctx context.Context, id uuid.UUID, processingOrder *ProcessingOrder, params ...*Params) (*ProcessingOrder, *resty.Response, error)

	// Template выполняет запрос на получение предзаполненного заказа на производство со стандартными полями без связи с какими-либо другими документами.
	// Принимает контекст.
	// Возвращает предзаполненный заказ на производство.
	Template(ctx context.Context) (*ProcessingOrder, *resty.Response, error)

	// TemplateBased выполняет запрос на получение шаблона заказа на производство на основе других документов.
	// Основание, на котором может быть создано:
	//	- Техкарта (ProcessingPlan)
	// Принимает контекст и один документ из списка выше.
	// Возвращает предзаполненнsq заказ на производство на основании переданных документов.
	TemplateBased(ctx context.Context, basedOn ...MetaOwner) (*ProcessingOrder, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных заказов на производство.
	// Принимает контекст.
	// Возвращает объект метаданных MetaAttributesStatesSharedWrapper.
	GetMetadata(ctx context.Context) (*MetaAttributesStatesSharedWrapper, *resty.Response, error)

	// GetPositionList выполняет запрос на получение списка позиций документа.
	// Принимает контекст, ID документа и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetPositionList(ctx context.Context, id uuid.UUID, params ...*Params) (*List[ProcessingOrderPosition], *resty.Response, error)

	// GetPositionByID выполняет запрос на получение отдельной позиции документа по ID.
	// Принимает контекст, ID документа, ID позиции и опционально объект параметров запроса Params.
	// Возвращает найденную позицию.
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params ...*Params) (*ProcessingOrderPosition, *resty.Response, error)

	// UpdatePosition выполняет запрос на изменение позиции документа.
	// Принимает контекст, ID документа, ID позиции, позицию документа и опционально объект параметров запроса Params.
	// Возвращает изменённую позицию.
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *ProcessingOrderPosition, params ...*Params) (*ProcessingOrderPosition, *resty.Response, error)

	// CreatePosition выполняет запрос на добавление позиции документа.
	// Принимает контекст, ID документа, позицию документа и опционально объект параметров запроса Params.
	// Возвращает добавленную позицию.
	CreatePosition(ctx context.Context, id uuid.UUID, position *ProcessingOrderPosition, params ...*Params) (*ProcessingOrderPosition, *resty.Response, error)

	// CreatePositionMany выполняет запрос на массовое добавление позиций документа.
	// Принимает контекст, ID документа и множество позиций.
	// Возвращает список добавленных позиций.
	CreatePositionMany(ctx context.Context, id uuid.UUID, positions ...*ProcessingOrderPosition) (*Slice[ProcessingOrderPosition], *resty.Response, error)

	// DeletePosition выполняет запрос на удаление позиции документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает «true» в случае успешного удаления позиции.
	DeletePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (bool, *resty.Response, error)

	// DeletePositionMany выполняет запрос на массовое удаление позиций документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeletePositionMany(ctx context.Context, id uuid.UUID, positions ...*ProcessingOrderPosition) (*DeleteManyResponse, *resty.Response, error)

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

	// GetBySyncID выполняет запрос на получение отдельного документа по syncID.
	// Принимает контекст и syncID документа.
	// Возвращает найденный документ.
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*ProcessingOrder, *resty.Response, error)

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
}

const (
	EndpointProcessingOrder = EndpointEntity + string(MetaTypeProcessingOrder)
)

// NewProcessingOrderService принимает [Client] и возвращает сервис для работы с заказами на производство.
func NewProcessingOrderService(client *Client) ProcessingOrderService {
	return newMainService[ProcessingOrder, ProcessingOrderPosition, MetaAttributesStatesSharedWrapper, any](client, EndpointProcessingOrder)
}
