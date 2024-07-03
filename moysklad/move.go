package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"time"
)

// Move Перемещение.
//
// Код сущности: move
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-peremeschenie
type Move struct {
	Moment        *Timestamp                `json:"moment,omitempty"`        // Дата документа
	Updated       *Timestamp                `json:"updated,omitempty"`       // Момент последнего обновления Перемещения
	AccountID     *uuid.UUID                `json:"accountId,omitempty"`     // ID учётной записи
	Code          *string                   `json:"code,omitempty"`          // Код Перемещения
	Created       *Timestamp                `json:"created,omitempty"`       // Дата создания
	Deleted       *Timestamp                `json:"deleted,omitempty"`       // Момент последнего удаления Перемещения
	Demand        *Demand                   `json:"demand,omitempty"`        // Метаданные Отгрузки, связанной с Перемещением
	Description   *string                   `json:"description,omitempty"`   // Комментарий Перемещения
	ExternalCode  *string                   `json:"externalCode,omitempty"`  // Внешний код Перемещения
	Files         *MetaArray[File]          `json:"files,omitempty"`         // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group         *Group                    `json:"group,omitempty"`         // Отдел сотрудника
	ID            *uuid.UUID                `json:"id,omitempty"`            // ID Перемещения
	InternalOrder *NullValue[InternalOrder] `json:"internalOrder,omitempty"` // Метаданные Внутреннего заказа, связанного с Перемещением
	CustomerOrder *NullValue[CustomerOrder] `json:"customerOrder,omitempty"` // Метаданные Заказа покупателя, связанного с Перемещением
	Meta          *Meta                     `json:"meta,omitempty"`          // Метаданные Перемещения
	Name          *string                   `json:"name,omitempty"`          // Наименование Перемещения
	Organization  *Organization             `json:"organization,omitempty"`  // Метаданные юрлица
	Applicable    *bool                     `json:"applicable,omitempty"`    // Отметка о проведении
	Overhead      *Overhead                 `json:"overhead,omitempty"`      // Накладные расходы. Если Позиции Перемещения не заданы, то накладные расходы нельзя задать
	Owner         *Employee                 `json:"owner,omitempty"`         // Метаданные владельца (Сотрудника)
	Positions     *MetaArray[MovePosition]  `json:"positions,omitempty"`     // Метаданные позиций Перемещения
	Printed       *bool                     `json:"printed,omitempty"`       // Напечатан ли документ
	Project       *NullValue[Project]       `json:"project,omitempty"`       // Метаданные проекта
	Published     *bool                     `json:"published,omitempty"`     // Опубликован ли документ
	Rate          *NullValue[Rate]          `json:"rate,omitempty"`          // Валюта
	Shared        *bool                     `json:"shared,omitempty"`        // Общий доступ
	SourceStore   *Store                    `json:"sourceStore,omitempty"`   // Метаданные склада, с которого совершается перемещение
	State         *NullValue[State]         `json:"state,omitempty"`         // Метаданные статуса Перемещения
	Sum           *float64                  `json:"sum,omitempty"`           // Сумма Перемещения в копейках
	SyncID        *uuid.UUID                `json:"syncId,omitempty"`        // ID синхронизации
	Supply        *Supply                   `json:"supply,omitempty"`        // Метаданные Приемки, связанной с Перемещением
	TargetStore   *Store                    `json:"targetStore,omitempty"`   // Метаданные склада, на который совершается перемещение
	Attributes    Slice[Attribute]          `json:"attributes,omitempty"`    // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (move Move) Clean() *Move {
	if move.Meta == nil {
		return nil
	}
	return &Move{Meta: move.Meta}
}

// asTaskOperation реализует интерфейс [TaskOperationInterface].
func (move Move) asTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: move.Meta}
}

// GetMoment возвращает Дату документа.
func (move Move) GetMoment() Timestamp {
	return Deref(move.Moment)
}

// GetUpdated возвращает Момент последнего обновления Перемещения.
func (move Move) GetUpdated() Timestamp {
	return Deref(move.Updated)
}

// GetAccountID возвращает ID учётной записи.
func (move Move) GetAccountID() uuid.UUID {
	return Deref(move.AccountID)
}

// GetCode возвращает Код Перемещения.
func (move Move) GetCode() string {
	return Deref(move.Code)
}

// GetCreated возвращает Дату создания.
func (move Move) GetCreated() Timestamp {
	return Deref(move.Created)
}

// GetDeleted возвращает Момент последнего удаления Перемещения.
func (move Move) GetDeleted() Timestamp {
	return Deref(move.Deleted)
}

// GetDescription возвращает Комментарий Перемещения.
func (move Move) GetDescription() string {
	return Deref(move.Description)
}

// GetExternalCode возвращает Внешний код Перемещения.
func (move Move) GetExternalCode() string {
	return Deref(move.ExternalCode)
}

// GetFiles возвращает Метаданные массива Файлов.
func (move Move) GetFiles() MetaArray[File] {
	return Deref(move.Files)
}

// GetGroup возвращает Отдел сотрудника.
func (move Move) GetGroup() Group {
	return Deref(move.Group)
}

// GetID возвращает ID Перемещения.
func (move Move) GetID() uuid.UUID {
	return Deref(move.ID)
}

// GetInternalOrder возвращает Метаданные Внутреннего заказа, связанного с Перемещением.
func (move Move) GetInternalOrder() InternalOrder {
	return move.InternalOrder.GetValue()
}

// GetCustomerOrder возвращает Ссылку на Заказ Покупателя, связанного с Перемещением.
func (move Move) GetCustomerOrder() CustomerOrder {
	return move.CustomerOrder.GetValue()
}

// GetMeta возвращает Метаданные Перемещения.
func (move Move) GetMeta() Meta {
	return Deref(move.Meta)
}

// GetName возвращает Наименование Перемещения.
func (move Move) GetName() string {
	return Deref(move.Name)
}

// GetOrganization возвращает Метаданные юрлица.
func (move Move) GetOrganization() Organization {
	return Deref(move.Organization)
}

// GetApplicable возвращает Отметку о проведении.
func (move Move) GetApplicable() bool {
	return Deref(move.Applicable)
}

// GetOverhead возвращает Накладные расходы.
//
// Если Позиции Перемещения не заданы, то накладные расходы нельзя задать.
func (move Move) GetOverhead() Overhead {
	return Deref(move.Overhead)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (move Move) GetOwner() Employee {
	return Deref(move.Owner)
}

// GetPositions возвращает Метаданные позиций Перемещения.
func (move Move) GetPositions() MetaArray[MovePosition] {
	return Deref(move.Positions)
}

// GetPrinted возвращает true, если документ напечатан.
func (move Move) GetPrinted() bool {
	return Deref(move.Printed)
}

// GetProject возвращает Метаданные проекта.
func (move Move) GetProject() Project {
	return move.Project.GetValue()
}

// GetPublished возвращает true, если документ опубликован.
func (move Move) GetPublished() bool {
	return Deref(move.Published)
}

// GetRate возвращает Валюту.
func (move Move) GetRate() Rate {
	return move.Rate.GetValue()
}

// GetShared возвращает флаг Общего доступа.
func (move Move) GetShared() bool {
	return Deref(move.Shared)
}

// GetSourceStore возвращает Метаданные склада, с которого совершается перемещение.
func (move Move) GetSourceStore() Store {
	return Deref(move.SourceStore)
}

// GetState возвращает Метаданные статуса Перемещения.
func (move Move) GetState() State {
	return move.State.GetValue()
}

// GetSum возвращает Сумму Перемещения в копейках.
func (move Move) GetSum() float64 {
	return Deref(move.Sum)
}

// GetSyncID возвращает ID синхронизации.
func (move Move) GetSyncID() uuid.UUID {
	return Deref(move.SyncID)
}

// GetTargetStore возвращает Метаданные склада, на который совершается перемещение.
func (move Move) GetTargetStore() Store {
	return Deref(move.TargetStore)
}

// GetDemand возвращает Метаданные Отгрузки, связанной с Перемещением.
func (move Move) GetDemand() Demand {
	return Deref(move.Demand)
}

// GetSupply возвращает Метаданные Приемки, связанной с Перемещением.
func (move Move) GetSupply() Supply {
	return Deref(move.Supply)
}

// GetAttributes возвращает Список метаданных доп. полей.
func (move Move) GetAttributes() Slice[Attribute] {
	return move.Attributes
}

// SetMoment устанавливает Дату документа.
func (move *Move) SetMoment(moment time.Time) *Move {
	move.Moment = NewTimestamp(moment)
	return move
}

// SetCode устанавливает Код Перемещения.
func (move *Move) SetCode(code string) *Move {
	move.Code = &code
	return move
}

// SetDescription устанавливает Комментарий Перемещения.
func (move *Move) SetDescription(description string) *Move {
	move.Description = &description
	return move
}

// SetExternalCode устанавливает Внешний код Перемещения.
func (move *Move) SetExternalCode(externalCode string) *Move {
	move.ExternalCode = &externalCode
	return move
}

// SetFiles устанавливает Метаданные массива Файлов.
//
// Принимает множество объектов [File].
func (move *Move) SetFiles(files ...*File) *Move {
	move.Files = NewMetaArrayFrom(files)
	return move
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (move *Move) SetGroup(group *Group) *Move {
	if group != nil {
		move.Group = group.Clean()
	}
	return move
}

// SetInternalOrder устанавливает Метаданные Внутреннего заказа, связанного с Перемещением.
func (move *Move) SetInternalOrder(internalOrder *InternalOrder) *Move {
	move.InternalOrder = NewNullValue(internalOrder)
	return move
}

// SetCustomerOrder устанавливает Ссылку на Заказ Покупателя, связанного с Перемещением.
func (move *Move) SetCustomerOrder(customerOrder *CustomerOrder) *Move {
	move.CustomerOrder = NewNullValue(customerOrder)
	return move
}

// SetMeta устанавливает Метаданные Перемещения.
func (move *Move) SetMeta(meta *Meta) *Move {
	move.Meta = meta
	return move
}

// SetName устанавливает Наименование Перемещения.
func (move *Move) SetName(name string) *Move {
	move.Name = &name
	return move
}

// SetOrganization устанавливает Метаданные юрлица.
func (move *Move) SetOrganization(organization *Organization) *Move {
	if organization != nil {
		move.Organization = organization.Clean()
	}
	return move
}

// SetApplicable устанавливает Отметку о проведении.
func (move *Move) SetApplicable(applicable bool) *Move {
	move.Applicable = &applicable
	return move
}

// SetOverhead устанавливает Накладные расходы.
//
// Если Позиции Оприходования не заданы, то накладные расходы нельзя задать.
func (move *Move) SetOverhead(overhead *Overhead) *Move {
	if overhead != nil {
		move.Overhead = overhead
	}
	return move
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (move *Move) SetOwner(owner *Employee) *Move {
	if owner != nil {
		move.Owner = owner.Clean()
	}
	return move
}

// SetPositions устанавливает Метаданные позиций Перемещения.
//
// Принимает множество объектов [MovePosition].
func (move *Move) SetPositions(positions ...*MovePosition) *Move {
	move.Positions = NewMetaArrayFrom(positions)
	return move
}

// SetProject устанавливает Метаданные проекта.
//
// Передача nil передаёт сброс значения (null).
func (move *Move) SetProject(project *Project) *Move {
	move.Project = NewNullValue(project)
	return move
}

// SetRate устанавливает Валюту.
//
// Передача nil передаёт сброс значения (null).
func (move *Move) SetRate(rate *Rate) *Move {
	move.Rate = NewNullValue(rate)
	return move
}

// SetShared устанавливает флаг общего доступа.
func (move *Move) SetShared(shared bool) *Move {
	move.Shared = &shared
	return move
}

// SetSourceStore устанавливает Метаданные склада, с которого совершается перемещение.
func (move *Move) SetSourceStore(sourceStore *Store) *Move {
	if sourceStore != nil {
		move.SourceStore = sourceStore
	}
	return move
}

// SetState устанавливает Метаданные статуса Перемещения.
//
// Передача nil передаёт сброс значения (null).
func (move *Move) SetState(state *State) *Move {
	move.State = NewNullValue(state)
	return move
}

// SetSyncID устанавливает ID синхронизации.
func (move *Move) SetSyncID(syncID uuid.UUID) *Move {
	move.SyncID = &syncID
	return move
}

// SetTargetStore устанавливает Метаданные склада, на который совершается перемещение.
func (move *Move) SetTargetStore(targetStore *Store) *Move {
	if targetStore != nil {
		move.TargetStore = targetStore.Clean()
	}
	return move
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (move *Move) SetAttributes(attributes ...*Attribute) *Move {
	move.Attributes.Push(attributes...)
	return move
}

// String реализует интерфейс [fmt.Stringer].
func (move Move) String() string {
	return Stringify(move)
}

// MetaType возвращает код сущности.
func (Move) MetaType() MetaType {
	return MetaTypeMove
}

// Update shortcut
func (move Move) Update(ctx context.Context, client *Client, params ...*Params) (*Move, *resty.Response, error) {
	return NewMoveService(client).Update(ctx, move.GetID(), &move, params...)
}

// Create shortcut
func (move Move) Create(ctx context.Context, client *Client, params ...*Params) (*Move, *resty.Response, error) {
	return NewMoveService(client).Create(ctx, &move, params...)
}

// Delete shortcut
func (move Move) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewMoveService(client).Delete(ctx, move.GetID())
}

// MovePosition Позиция перемещения.
//
// Код сущности: moveposition
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-peremeschenie-peremescheniq-pozicii-peremescheniq
type MovePosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учётной записи
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

// GetAccountID возвращает ID учётной записи.
func (movePosition MovePosition) GetAccountID() uuid.UUID {
	return Deref(movePosition.AccountID)
}

// GetAssortment возвращает Метаданные товара/услуги/серии/модификации, которую представляет собой позиция.
func (movePosition MovePosition) GetAssortment() AssortmentPosition {
	return Deref(movePosition.Assortment)
}

// GetID возвращает ID позиции.
func (movePosition MovePosition) GetID() uuid.UUID {
	return Deref(movePosition.ID)
}

// GetOverhead возвращает Накладные расходы.
func (movePosition MovePosition) GetOverhead() float64 {
	return Deref(movePosition.Overhead)
}

// GetPack возвращает Упаковку Товара.
func (movePosition MovePosition) GetPack() Pack {
	return Deref(movePosition.Pack)
}

// GetPrice возвращает Цену товара/услуги в копейках.
func (movePosition MovePosition) GetPrice() float64 {
	return Deref(movePosition.Price)
}

// GetQuantity возвращает Количество товаров данного вида в позиции.
func (movePosition MovePosition) GetQuantity() float64 {
	return Deref(movePosition.Quantity)
}

// GetSourceSlot возвращает Ячейку на складе, с которого совершается перемещение.
func (movePosition MovePosition) GetSourceSlot() Slot {
	return Deref(movePosition.SourceSlot)
}

// GetTargetSlot возвращает Ячейку на складе, на который совершается перемещение.
func (movePosition MovePosition) GetTargetSlot() Slot {
	return Deref(movePosition.TargetSlot)
}

// GetThings возвращает Серийные номера.
//
// Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете.
// В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута.
func (movePosition MovePosition) GetThings() Slice[string] {
	return movePosition.Things
}

// SetAssortment устанавливает Метаданные товара/услуги, которую представляет собой компонент.
//
// Принимает объект, реализующий интерфейс [AssortmentInterface].
func (movePosition *MovePosition) SetAssortment(assortment AssortmentInterface) *MovePosition {
	if assortment != nil {
		movePosition.Assortment = assortment.asAssortment()
	}
	return movePosition
}

// SetPack устанавливает Упаковку Товара.
func (movePosition *MovePosition) SetPack(pack *Pack) *MovePosition {
	if pack != nil {
		movePosition.Pack = pack
	}
	return movePosition
}

// SetPrice устанавливает Цену товара/услуги в копейках.
func (movePosition *MovePosition) SetPrice(price float64) *MovePosition {
	movePosition.Price = &price
	return movePosition
}

// SetQuantity устанавливает Количество товаров/услуг данного вида в компоненте.
func (movePosition *MovePosition) SetQuantity(quantity float64) *MovePosition {
	movePosition.Quantity = &quantity
	return movePosition
}

// SetSourceSlot устанавливает Ячейку на складе, с которого совершается перемещение.
func (movePosition *MovePosition) SetSourceSlot(sourceSlot *Slot) *MovePosition {
	if sourceSlot != nil {
		movePosition.SourceSlot = sourceSlot.Clean()
	}
	return movePosition
}

// SetTargetSlot устанавливает Ячейку на складе, на который совершается перемещение.
func (movePosition *MovePosition) SetTargetSlot(targetSlot *Slot) *MovePosition {
	if targetSlot != nil {
		movePosition.TargetSlot = targetSlot.Clean()
	}
	return movePosition
}

// SetThings устанавливает Серийные номера.
//
// Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете.
// В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута.
//
// Принимает множество string.
func (movePosition *MovePosition) SetThings(things ...string) *MovePosition {
	movePosition.Things = NewSliceFrom(things)
	return movePosition
}

// String реализует интерфейс [fmt.Stringer].
func (movePosition MovePosition) String() string {
	return Stringify(movePosition)
}

// MetaType возвращает код сущности.
func (MovePosition) MetaType() MetaType {
	return MetaTypeMovePosition
}

// MoveService описывает методы сервиса для работы со перемещениями.
type MoveService interface {
	// GetList выполняет запрос на получение списка перемещений.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[Move], *resty.Response, error)

	// Create выполняет запрос на создание перемещения.
	// Обязательные поля для заполнения:
	//	- organization (Ссылка на ваше юрлицо)
	//	- sourceStore (Ссылка на склад, с которого совершается перемещение)
	//	- targetStore (Ссылка на склад, на который совершается перемещение)
	// Принимает контекст, перемещение и опционально объект параметров запроса Params.
	// Возвращает созданное перемещение.
	Create(ctx context.Context, move *Move, params ...*Params) (*Move, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение перемещений.
	// Изменяемые перемещения должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список перемещений и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых перемещений.
	CreateUpdateMany(ctx context.Context, moveList Slice[Move], params ...*Params) (*Slice[Move], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление перемещений.
	// Принимает контекст и множество перемещений.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*Move) (*DeleteManyResponse, *resty.Response, error)

	// Delete выполняет запрос на удаление перемещения.
	// Принимает контекст и ID перемещения.
	// Возвращает true в случае успешного удаления перемещения.
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельного перемещения по ID.
	// Принимает контекст, ID перемещения и опционально объект параметров запроса Params.
	// Возвращает найденное перемещение.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Move, *resty.Response, error)

	// Update выполняет запрос на изменение перемещения.
	// Принимает контекст, перемещение и опционально объект параметров запроса Params.
	// Возвращает изменённое перемещение.
	Update(ctx context.Context, id uuid.UUID, move *Move, params ...*Params) (*Move, *resty.Response, error)

	// Template выполняет запрос на получение предзаполненного перемещения со стандартными полями без связи с какими-либо другими документами.
	// Принимает контекст.
	// Возвращает предзаполненное перемещение.
	Template(ctx context.Context) (*Move, *resty.Response, error)

	// TemplateBased выполняет запрос на получение шаблона перемещения на основе других документов.
	// Основание, на котором может быть создано:
	//	- Внутренний заказ (InternalOrder)
	//	- Заказ покупателя (CustomerOrder)
	// Принимает контекст и множество документов из списка выше.
	// Возвращает предзаполненное перемещение на основании переданных документов.
	TemplateBased(ctx context.Context, basedOn ...MetaOwner) (*Move, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных перемещений.
	// Принимает контекст.
	// Возвращает объект метаданных MetaAttributesStatesSharedWrapper.
	GetMetadata(ctx context.Context) (*MetaAttributesStatesSharedWrapper, *resty.Response, error)

	// GetPositionList выполняет запрос на получение списка позиций документа.
	// Принимает контекст, ID документа и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetPositionList(ctx context.Context, id uuid.UUID, params ...*Params) (*List[MovePosition], *resty.Response, error)

	// GetPositionByID выполняет запрос на получение отдельной позиции документа по ID.
	// Принимает контекст, ID документа, ID позиции и опционально объект параметров запроса Params.
	// Возвращает найденную позицию.
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params ...*Params) (*MovePosition, *resty.Response, error)

	// UpdatePosition выполняет запрос на изменение позиции документа.
	// Принимает контекст, ID документа, ID позиции, позицию документа и опционально объект параметров запроса Params.
	// Возвращает изменённую позицию.
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *MovePosition, params ...*Params) (*MovePosition, *resty.Response, error)

	// CreatePosition выполняет запрос на добавление позиции документа.
	// Принимает контекст, ID документа, позицию документа и опционально объект параметров запроса Params.
	// Возвращает добавленную позицию.
	CreatePosition(ctx context.Context, id uuid.UUID, position *MovePosition, params ...*Params) (*MovePosition, *resty.Response, error)

	// CreatePositionMany выполняет запрос на массовое добавление позиций документа.
	// Принимает контекст, ID документа и множество позиций.
	// Возвращает список добавленных позиций.
	CreatePositionMany(ctx context.Context, id uuid.UUID, positions ...*MovePosition) (*Slice[MovePosition], *resty.Response, error)

	// DeletePosition выполняет запрос на удаление позиции документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает true в случае успешного удаления позиции.
	DeletePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (bool, *resty.Response, error)

	// DeletePositionMany выполняет запрос на массовое удаление позиций документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeletePositionMany(ctx context.Context, id uuid.UUID, positions ...*MovePosition) (*DeleteManyResponse, *resty.Response, error)

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

	// GetPublicationList выполняет запрос на получение списка публикаций.
	// Принимает контекст и ID документа.
	// Возвращает объект List.
	GetPublicationList(ctx context.Context, id uuid.UUID) (*List[Publication], *resty.Response, error)

	// GetPublicationByID выполняет запрос на получение отдельной публикации по ID.
	// Принимает контекст, ID документа и ID публикации.
	// Возвращает найденную публикацию.
	GetPublicationByID(ctx context.Context, id uuid.UUID, publicationID uuid.UUID) (*Publication, *resty.Response, error)

	// Publish выполняет запрос на создание публикации.
	// Принимает контекст, ID документа и шаблон.
	// Возвращает созданную публикацию.
	Publish(ctx context.Context, id uuid.UUID, template TemplateInterface) (*Publication, *resty.Response, error)

	// DeletePublication выполняет запрос на удаление публикации.
	// Принимает контекст, ID документа и ID публикации.
	// Возвращает true в случае успешного удаления публикации.
	DeletePublication(ctx context.Context, id uuid.UUID, publicationID uuid.UUID) (bool, *resty.Response, error)

	// GetBySyncID выполняет запрос на получение отдельного документа по syncID.
	// Принимает контекст и syncID документа.
	// Возвращает найденный документ.
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*Move, *resty.Response, error)

	// DeleteBySyncID выполняет запрос на удаление документа по syncID.
	// Принимает контекст и syncID документа.
	// Возвращает true в случае успешного удаления документа.
	DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error)

	// MoveToTrash выполняет запрос на перемещение документа с указанным ID в корзину.
	// Принимает контекст и ID документа.
	// Возвращает true в случае успешного перемещения в корзину.
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
	// Возвращает true в случае успешного удаления статуса.
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
	Evaluate(ctx context.Context, entity *Move, evaluate ...Evaluate) (*Move, *resty.Response, error)
}

// NewMoveService принимает [Client] и возвращает сервис для работы со перемещениями.
func NewMoveService(client *Client) MoveService {
	return newMainService[Move, MovePosition, MetaAttributesStatesSharedWrapper, any](NewEndpoint(client, "entity/move"))
}
