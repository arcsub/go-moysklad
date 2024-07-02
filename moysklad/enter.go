package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Enter Оприходование.
//
// Код сущности: enter
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-oprihodowanie
type Enter struct {
	Organization *Organization             `json:"organization,omitempty"` // Метаданные юрлица
	Sum          *float64                  `json:"sum,omitempty"`          // Сумма Оприходования в копейках
	Moment       *Timestamp                `json:"moment,omitempty"`       // Дата документа
	Code         *string                   `json:"code,omitempty"`         // Код Оприходования
	Created      *Timestamp                `json:"created,omitempty"`      // Дата создания
	Deleted      *Timestamp                `json:"deleted,omitempty"`      // Момент последнего удаления Оприходования
	Description  *string                   `json:"description,omitempty"`  // Комментарий Оприходования
	ExternalCode *string                   `json:"externalCode,omitempty"` // Внешний код Оприходования
	Files        *MetaArray[File]          `json:"files,omitempty"`        // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group        *Group                    `json:"group,omitempty"`        // Отдел сотрудника
	ID           *uuid.UUID                `json:"id,omitempty"`           // ID Оприходования
	Meta         *Meta                     `json:"meta,omitempty"`         // Метаданные Оприходования
	Updated      *Timestamp                `json:"updated,omitempty"`      // Момент последнего обновления Оприходования
	Applicable   *bool                     `json:"applicable,omitempty"`   // Отметка о проведении
	Printed      *bool                     `json:"printed,omitempty"`      // Напечатан ли документ
	Overhead     *Overhead                 `json:"overhead,omitempty"`     // Накладные расходы. Если Позиции Оприходования не заданы, то накладные расходы нельзя задать
	Owner        *Employee                 `json:"owner,omitempty"`        // Метаданные владельца (Сотрудника)
	Positions    *MetaArray[EnterPosition] `json:"positions,omitempty"`    // Метаданные позиций Оприходования
	AccountID    *uuid.UUID                `json:"accountId,omitempty"`    // ID учётной записи
	Project      *NullValue[Project]       `json:"project,omitempty"`      // Метаданные проекта
	Published    *bool                     `json:"published,omitempty"`    // Опубликован ли документ
	Rate         *NullValue[Rate]          `json:"rate,omitempty"`         // Валюта
	Shared       *bool                     `json:"shared,omitempty"`       // Общий доступ
	State        *NullValue[State]         `json:"state,omitempty"`        // Метаданные статуса оприходования
	Store        *Store                    `json:"store,omitempty"`        // Метаданные склада
	Name         *string                   `json:"name,omitempty"`         // Номер Оприходования
	SyncID       *uuid.UUID                `json:"syncId,omitempty"`       // ID синхронизации
	Attributes   Slice[Attribute]          `json:"attributes,omitempty"`   // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (enter Enter) Clean() *Enter {
	if enter.Meta == nil {
		return nil
	}
	return &Enter{Meta: enter.Meta}
}

// asTaskOperation реализует интерфейс [TaskOperationInterface].
func (enter Enter) asTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: enter.Meta}
}

// GetOrganization возвращает Метаданные юрлица.
func (enter Enter) GetOrganization() Organization {
	return Deref(enter.Organization)
}

// GetSum возвращает Сумму Оприходования в копейках.
func (enter Enter) GetSum() float64 {
	return Deref(enter.Sum)
}

// GetMoment возвращает Дату документа.
func (enter Enter) GetMoment() Timestamp {
	return Deref(enter.Moment)
}

// GetCode возвращает Код Оприходования.
func (enter Enter) GetCode() string {
	return Deref(enter.Code)
}

// GetCreated возвращает Дату создания.
func (enter Enter) GetCreated() Timestamp {
	return Deref(enter.Created)
}

// GetDeleted возвращает Момент последнего удаления Оприходования.
func (enter Enter) GetDeleted() Timestamp {
	return Deref(enter.Deleted)
}

// GetDescription возвращает Комментарий Оприходования.
func (enter Enter) GetDescription() string {
	return Deref(enter.Description)
}

// GetExternalCode возвращает Внешний код Оприходования.
func (enter Enter) GetExternalCode() string {
	return Deref(enter.ExternalCode)
}

// GetFiles возвращает Метаданные массива Файлов.
func (enter Enter) GetFiles() MetaArray[File] {
	return Deref(enter.Files)
}

// GetGroup возвращает Отдел сотрудника.
func (enter Enter) GetGroup() Group {
	return Deref(enter.Group)
}

// GetID возвращает ID Оприходования.
func (enter Enter) GetID() uuid.UUID {
	return Deref(enter.ID)
}

// GetMeta возвращает Метаданные Оприходования.
func (enter Enter) GetMeta() Meta {
	return Deref(enter.Meta)
}

// GetUpdated возвращает Момент последнего обновления Оприходования.
func (enter Enter) GetUpdated() Timestamp {
	return Deref(enter.Updated)
}

// GetApplicable возвращает Отметку о проведении.
func (enter Enter) GetApplicable() bool {
	return Deref(enter.Applicable)
}

// GetPrinted возвращает true, если документ напечатан.
func (enter Enter) GetPrinted() bool {
	return Deref(enter.Printed)
}

// GetOverhead возвращает Накладные расходы.
func (enter Enter) GetOverhead() Overhead {
	return Deref(enter.Overhead)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (enter Enter) GetOwner() Employee {
	return Deref(enter.Owner)
}

// GetPositions возвращает Метаданные позиций Оприходования.
func (enter Enter) GetPositions() MetaArray[EnterPosition] {
	return Deref(enter.Positions)
}

// GetAccountID возвращает ID учётной записи.
func (enter Enter) GetAccountID() uuid.UUID {
	return Deref(enter.AccountID)
}

// GetProject возвращает Метаданные проекта.
func (enter Enter) GetProject() Project {
	return enter.Project.GetValue()
}

// GetPublished возвращает true, если документ опубликован.
func (enter Enter) GetPublished() bool {
	return Deref(enter.Published)
}

// GetRate возвращает Валюту.
func (enter Enter) GetRate() Rate {
	return enter.Rate.GetValue()
}

// GetShared возвращает флаг Общего доступа.
func (enter Enter) GetShared() bool {
	return Deref(enter.Shared)
}

// GetState возвращает Метаданные статуса оприходования.
func (enter Enter) GetState() State {
	return enter.State.GetValue()
}

// GetStore возвращает Метаданные склада.
func (enter Enter) GetStore() Store {
	return Deref(enter.Store)
}

// GetName возвращает Номер Оприходования.
func (enter Enter) GetName() string {
	return Deref(enter.Name)
}

// GetSyncID возвращает ID синхронизации.
func (enter Enter) GetSyncID() uuid.UUID {
	return Deref(enter.SyncID)
}

// GetAttributes возвращает Список метаданных доп. полей.
func (enter Enter) GetAttributes() Slice[Attribute] {
	return enter.Attributes
}

// SetOrganization устанавливает Метаданные юрлица.
func (enter *Enter) SetOrganization(organization *Organization) *Enter {
	if organization != nil {
		enter.Organization = organization.Clean()
	}
	return enter
}

// SetMoment устанавливает Дату документа.
func (enter *Enter) SetMoment(moment *Timestamp) *Enter {
	enter.Moment = moment
	return enter
}

// SetCode устанавливает Код Оприходования.
func (enter *Enter) SetCode(code string) *Enter {
	enter.Code = &code
	return enter
}

// SetDescription устанавливает Комментарий Оприходования.
func (enter *Enter) SetDescription(description string) *Enter {
	enter.Description = &description
	return enter
}

// SetExternalCode устанавливает Внешний код Оприходования.
func (enter *Enter) SetExternalCode(externalCode string) *Enter {
	enter.ExternalCode = &externalCode
	return enter
}

// SetFiles устанавливает Метаданные массива Файлов.
//
// Принимает множество объектов [File].
func (enter *Enter) SetFiles(files ...*File) *Enter {
	enter.Files = NewMetaArrayFrom(files)
	return enter
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (enter *Enter) SetGroup(group *Group) *Enter {
	if group != nil {
		enter.Group = group.Clean()
	}
	return enter
}

// SetMeta устанавливает Метаданные Оприходования.
func (enter *Enter) SetMeta(meta *Meta) *Enter {
	enter.Meta = meta
	return enter
}

// SetApplicable устанавливает Отметку о проведении.
func (enter *Enter) SetApplicable(applicable bool) *Enter {
	enter.Applicable = &applicable
	return enter
}

// SetOverhead устанавливает Накладные расходы.
//
// Если Позиции Оприходования не заданы, то накладные расходы нельзя задать.
func (enter *Enter) SetOverhead(overhead *Overhead) *Enter {
	if overhead != nil {
		enter.Overhead = overhead
	}
	return enter
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (enter *Enter) SetOwner(owner *Employee) *Enter {
	if owner != nil {
		enter.Owner = owner.Clean()
	}
	return enter
}

// SetPositions устанавливает Метаданные позиций Оприходования.
//
// Принимает множество объектов [EnterPosition].
func (enter *Enter) SetPositions(positions ...*EnterPosition) *Enter {
	enter.Positions = NewMetaArrayFrom(positions)
	return enter
}

// SetProject устанавливает Метаданные проекта.
//
// Передача nil передаёт сброс значения (null).
func (enter *Enter) SetProject(project *Project) *Enter {
	enter.Project = NewNullValue(project)
	return enter
}

// SetRate устанавливает Валюту.
//
// Передача nil передаёт сброс значения (null).
func (enter *Enter) SetRate(rate *Rate) *Enter {
	enter.Rate = NewNullValue(rate)
	return enter
}

// SetShared устанавливает флаг общего доступа.
func (enter *Enter) SetShared(shared bool) *Enter {
	enter.Shared = &shared
	return enter
}

// SetState устанавливает Метаданные статуса Оприходования.
//
// Передача nil передаёт сброс значения (null).
func (enter *Enter) SetState(state *State) *Enter {
	enter.State = NewNullValue(state)
	return enter
}

// SetStore устанавливает Метаданные склада.
func (enter *Enter) SetStore(store *Store) *Enter {
	if store != nil {
		enter.Store = store.Clean()
	}
	return enter
}

// SetName устанавливает Номер Оприходования.
func (enter *Enter) SetName(name string) *Enter {
	enter.Name = &name
	return enter
}

// SetSyncID устанавливает ID синхронизации.
func (enter *Enter) SetSyncID(syncID uuid.UUID) *Enter {
	enter.SyncID = &syncID
	return enter
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (enter *Enter) SetAttributes(attributes ...*Attribute) *Enter {
	enter.Attributes.Push(attributes...)
	return enter
}

// String реализует интерфейс [fmt.Stringer].
func (enter Enter) String() string {
	return Stringify(enter)
}

// MetaType возвращает код сущности.
func (Enter) MetaType() MetaType {
	return MetaTypeEnter
}

// Update shortcut
func (enter Enter) Update(ctx context.Context, client *Client, params ...*Params) (*Enter, *resty.Response, error) {
	return NewEnterService(client).Update(ctx, enter.GetID(), &enter, params...)
}

// Create shortcut
func (enter Enter) Create(ctx context.Context, client *Client, params ...*Params) (*Enter, *resty.Response, error) {
	return NewEnterService(client).Create(ctx, &enter, params...)
}

// Delete shortcut
func (enter Enter) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewEnterService(client).Delete(ctx, enter.GetID())
}

// EnterPosition Позиция оприходования.
//
// Код сущности: enterposition
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-oprihodowanie-oprihodowaniq-pozicii-oprihodowaniq
type EnterPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учётной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	Country    *Country            `json:"country,omitempty"`    // Метаданные Страны
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

// GetAccountID возвращает ID учётной записи.
func (enterPosition EnterPosition) GetAccountID() uuid.UUID {
	return Deref(enterPosition.AccountID)
}

// GetAssortment возвращает Метаданные товара/услуги/серии/модификации, которую представляет собой позиция.
func (enterPosition EnterPosition) GetAssortment() AssortmentPosition {
	return Deref(enterPosition.Assortment)
}

// GetCountry возвращает Метаданные Страны.
func (enterPosition EnterPosition) GetCountry() Country {
	return Deref(enterPosition.Country)
}

// GetGTD возвращает ГТД.
func (enterPosition EnterPosition) GetGTD() GTD {
	return Deref(enterPosition.GTD)
}

// GetGTDName возвращает Номер ГТД.
func (enterPosition EnterPosition) GetGTDName() string {
	return Deref(enterPosition.GTD).GetName()
}

// GetID возвращает ID позиции.
func (enterPosition EnterPosition) GetID() uuid.UUID {
	return Deref(enterPosition.ID)
}

// GetOverhead возвращает Накладные расходы.
func (enterPosition EnterPosition) GetOverhead() float64 {
	return Deref(enterPosition.Overhead)
}

// GetPack возвращает Упаковку Товара.
func (enterPosition EnterPosition) GetPack() Pack {
	return Deref(enterPosition.Pack)
}

// GetPrice возвращает Цену товара/услуги в копейках.
func (enterPosition EnterPosition) GetPrice() float64 {
	return Deref(enterPosition.Price)
}

// GetQuantity возвращает Количество товаров/услуг данного вида в позиции.
//
// Если позиция - товар, у которого включен учет по серийным номерам,
// то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
func (enterPosition EnterPosition) GetQuantity() float64 {
	return Deref(enterPosition.Quantity)
}

// GetReason возвращает Причину оприходования данной позиции.
func (enterPosition EnterPosition) GetReason() string {
	return Deref(enterPosition.Reason)
}

// GetSlot возвращает Ячейку на складе.
func (enterPosition EnterPosition) GetSlot() Slot {
	return Deref(enterPosition.Slot)
}

// GetThings возвращает Серийные номера.
//
// Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете.
// В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута.
func (enterPosition EnterPosition) GetThings() Slice[string] {
	return enterPosition.Things
}

// SetAssortment устанавливает Метаданные товара/услуги/серии/модификации, которую представляет собой позиция.
//
// Принимает объект, реализующий интерфейс [AssortmentInterface].
func (enterPosition *EnterPosition) SetAssortment(assortment AssortmentInterface) *EnterPosition {
	if assortment != nil {
		enterPosition.Assortment = assortment.asAssortment()
	}
	return enterPosition
}

// SetCountry устанавливает Метаданные Страны.
func (enterPosition *EnterPosition) SetCountry(country *Country) *EnterPosition {
	if country != nil {
		enterPosition.Country = country.Clean()
	}
	return enterPosition
}

// SetGTD устанавливает ГТД.
func (enterPosition *EnterPosition) SetGTD(gtd *GTD) *EnterPosition {
	if gtd != nil {
		enterPosition.GTD = gtd
	}
	return enterPosition
}

// SetPack устанавливает Упаковку Товара.
func (enterPosition *EnterPosition) SetPack(pack *Pack) *EnterPosition {
	if pack != nil {
		enterPosition.Pack = pack
	}
	return enterPosition
}

// SetPrice устанавливает Цену товара/услуги в копейках.
func (enterPosition *EnterPosition) SetPrice(price float64) *EnterPosition {
	enterPosition.Price = &price
	return enterPosition
}

// SetQuantity устанавливает Количество товаров/услуг данного вида в позиции.
//
// Если позиция - товар, у которого включен учет по серийным номерам,
// то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
func (enterPosition *EnterPosition) SetQuantity(quantity float64) *EnterPosition {
	enterPosition.Quantity = &quantity
	return enterPosition
}

// SetReason устанавливает Причину оприходования данной позиции.
func (enterPosition *EnterPosition) SetReason(reason string) *EnterPosition {
	enterPosition.Reason = &reason
	return enterPosition
}

// SetSlot устанавливает Ячейку на складе.
func (enterPosition *EnterPosition) SetSlot(slot *Slot) *EnterPosition {
	if slot != nil {
		enterPosition.Slot = slot.Clean()
	}
	return enterPosition
}

// SetThings устанавливает Серийные номера.
//
// Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете.
// В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута.
//
// Принимает множество string.
func (enterPosition *EnterPosition) SetThings(things ...string) *EnterPosition {
	enterPosition.Things = NewSliceFrom(things)
	return enterPosition
}

// String реализует интерфейс [fmt.Stringer].
func (enterPosition EnterPosition) String() string {
	return Stringify(enterPosition)
}

// MetaType возвращает код сущности.
func (EnterPosition) MetaType() MetaType {
	return MetaTypeEnterPosition
}

// EnterService описывает методы сервиса для работы с оприходованиями.
type EnterService interface {
	// GetList выполняет запрос на получение списка оприходований.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[Enter], *resty.Response, error)

	// Create выполняет запрос на создание оприходования.
	// Обязательные поля для заполнения:
	//	- organization (Метаданные юрлица)
	//	- store (Метаданные склада)
	// Принимает контекст, оприходование и опционально объект параметров запроса Params.
	// Возвращает созданное оприходование.
	Create(ctx context.Context, enter *Enter, params ...*Params) (*Enter, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение оприходований.
	// Изменяемые оприходования должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список оприходований и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых оприходований.
	CreateUpdateMany(ctx context.Context, enterList Slice[Enter], params ...*Params) (*Slice[Enter], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление оприходований.
	// Принимает контекст и множество оприходований.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*Enter) (*DeleteManyResponse, *resty.Response, error)

	// Delete выполняет запрос на удаление оприходования.
	// Принимает контекст и ID оприходования.
	// Возвращает true в случае успешного удаления оприходования.
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельного оприходования по ID.
	// Принимает контекст, ID оприходования взаиморасчётов и опционально объект параметров запроса Params.
	// Возвращает найденное оприходование.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Enter, *resty.Response, error)

	// Update выполняет запрос на изменение оприходования.
	// Принимает контекст, оприходование и опционально объект параметров запроса Params.
	// Возвращает изменённое оприходование.
	Update(ctx context.Context, id uuid.UUID, enter *Enter, params ...*Params) (*Enter, *resty.Response, error)

	// Template выполняет запрос на получение предзаполненного оприходования со стандартными полями.
	// без связи с какими-либо другими документами.
	// Принимает контекст.
	// Возвращает предзаполненное оприходование.
	Template(ctx context.Context) (*Enter, *resty.Response, error)

	// TemplateBased выполняет запрос на получение шаблона оприходования на основе других документов.
	// Основание, на котором может быть создано:
	//	- Инвентаризация (Inventory)
	// Принимает контекст и множество документов из списка выше.
	// Возвращает предзаполненное оприходование на основании переданных документов.
	TemplateBased(ctx context.Context, basedOn ...MetaOwner) (*Enter, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных оприходований.
	// Принимает контекст.
	// Возвращает объект метаданных MetaAttributesStatesSharedWrapper.
	GetMetadata(ctx context.Context) (*MetaAttributesStatesSharedWrapper, *resty.Response, error)

	// GetPositionList выполняет запрос на получение списка позиций документа.
	// Принимает контекст, ID документа и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetPositionList(ctx context.Context, id uuid.UUID, params ...*Params) (*List[EnterPosition], *resty.Response, error)

	// GetPositionByID выполняет запрос на получение отдельной позиции документа по ID.
	// Принимает контекст, ID документа, ID позиции и опционально объект параметров запроса Params.
	// Возвращает найденную позицию.
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params ...*Params) (*EnterPosition, *resty.Response, error)

	// UpdatePosition выполняет запрос на изменение позиции документа.
	// Принимает контекст, ID документа, ID позиции, позицию документа и опционально объект параметров запроса Params.
	// Возвращает изменённую позицию.
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *EnterPosition, params ...*Params) (*EnterPosition, *resty.Response, error)

	// CreatePosition выполняет запрос на добавление позиции документа.
	// Принимает контекст, ID документа, позицию документа и опционально объект параметров запроса Params.
	// Возвращает добавленную позицию.
	CreatePosition(ctx context.Context, id uuid.UUID, position *EnterPosition, params ...*Params) (*EnterPosition, *resty.Response, error)

	// CreatePositionMany выполняет запрос на массовое добавление позиций документа.
	// Принимает контекст, ID документа и множество позиций.
	// Возвращает список добавленных позиций.
	CreatePositionMany(ctx context.Context, id uuid.UUID, positions ...*EnterPosition) (*Slice[EnterPosition], *resty.Response, error)

	// DeletePosition выполняет запрос на удаление позиции документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает true в случае успешного удаления позиции.
	DeletePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (bool, *resty.Response, error)

	// DeletePositionMany выполняет запрос на массовое удаление позиций документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeletePositionMany(ctx context.Context, id uuid.UUID, positions ...*EnterPosition) (*DeleteManyResponse, *resty.Response, error)

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
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*Enter, *resty.Response, error)

	// DeleteBySyncID выполняет запрос на удаление документа по syncID.
	// Принимает контекст и syncID документа.
	// Возвращает true в случае успешного удаления документа.
	DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error)

	// GetNamedFilterList выполняет запрос на получение списка фильтров.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetNamedFilterList(ctx context.Context, params ...*Params) (*List[NamedFilter], *resty.Response, error)

	// GetNamedFilterByID выполняет запрос на получение отдельного фильтра по ID.
	// Принимает контекст и ID фильтра.
	// Возвращает найденный фильтр.
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)

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
	Evaluate(ctx context.Context, entity *Enter, evaluate ...Evaluate) (*Enter, *resty.Response, error)
}

// NewEnterService принимает [Client] и возвращает сервис для работы с оприходованиями.
func NewEnterService(client *Client) EnterService {
	return newMainService[Enter, EnterPosition, MetaAttributesStatesSharedWrapper, any](NewEndpoint(client, "entity/enter"))
}
