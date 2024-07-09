package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"time"
)

// Loss Списание.
//
// Код сущности: loss
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-spisanie
type Loss struct {
	Organization *Organization            `json:"organization,omitempty"` // Метаданные юрлица
	SyncID       *uuid.UUID               `json:"syncId,omitempty"`       // ID синхронизации
	Moment       *Timestamp               `json:"moment,omitempty"`       // Дата документа
	Code         *string                  `json:"code,omitempty"`         // Код Списания
	Created      *Timestamp               `json:"created,omitempty"`      // Дата создания
	Deleted      *Timestamp               `json:"deleted,omitempty"`      // Момент последнего удаления Списания
	Description  *string                  `json:"description,omitempty"`  // Комментарий Списания
	ExternalCode *string                  `json:"externalCode,omitempty"` // Внешний код Списания
	Files        *MetaArray[File]         `json:"files,omitempty"`        // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group        *Group                   `json:"group,omitempty"`        // Отдел сотрудника
	ID           *uuid.UUID               `json:"id,omitempty"`           // ID Списания
	Meta         *Meta                    `json:"meta,omitempty"`         // Метаданные Списания
	SalesReturn  *SalesReturn             `json:"salesReturn,omitempty"`  // Ссылка на связанный со списанием возврат покупателя
	Applicable   *bool                    `json:"applicable,omitempty"`   // Отметка о проведении
	Project      *NullValue[Project]      `json:"project,omitempty"`      // Метаданные проекта
	Owner        *Employee                `json:"owner,omitempty"`        // Метаданные владельца (Сотрудника)
	Positions    *MetaArray[LossPosition] `json:"positions,omitempty"`    // Метаданные позиций Списания
	Printed      *bool                    `json:"printed,omitempty"`      // Напечатан ли документ
	AccountID    *uuid.UUID               `json:"accountId,omitempty"`    // ID учётной записи
	Published    *bool                    `json:"published,omitempty"`    // Опубликован ли документ
	Rate         *NullValue[Rate]         `json:"rate,omitempty"`         // Валюта
	Shared       *bool                    `json:"shared,omitempty"`       // Общий доступ
	State        *NullValue[State]        `json:"state,omitempty"`        // Метаданные статуса Списания
	Store        *Store                   `json:"store,omitempty"`        // Метаданные склада
	Sum          *float64                 `json:"sum,omitempty"`          // Сумма Списания в копейках
	Name         *string                  `json:"name,omitempty"`         // Наименование Списания
	Updated      *Timestamp               `json:"updated,omitempty"`      // Момент последнего обновления Списания
	Inventory    *Inventory               `json:"inventory,omitempty"`    // Ссылка на связанную со списанием инвентаризацию
	Attributes   Slice[Attribute]         `json:"attributes,omitempty"`   // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (loss Loss) Clean() *Loss {
	if loss.Meta == nil {
		return nil
	}
	return &Loss{Meta: loss.Meta}
}

// AsTaskOperation реализует интерфейс [TaskOperationConverter].
func (loss Loss) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: loss.Meta}
}

// GetOrganization возвращает Метаданные юрлица.
func (loss Loss) GetOrganization() Organization {
	return Deref(loss.Organization)
}

// GetSyncID возвращает ID синхронизации.
func (loss Loss) GetSyncID() uuid.UUID {
	return Deref(loss.SyncID)
}

// GetMoment возвращает Дату документа.
func (loss Loss) GetMoment() Timestamp {
	return Deref(loss.Moment)
}

// GetCode возвращает Код Списания.
func (loss Loss) GetCode() string {
	return Deref(loss.Code)
}

// GetCreated возвращает Дату создания.
func (loss Loss) GetCreated() Timestamp {
	return Deref(loss.Created)
}

// GetDeleted возвращает Момент последнего удаления Списания.
func (loss Loss) GetDeleted() Timestamp {
	return Deref(loss.Deleted)
}

// GetDescription возвращает Комментарий Списания.
func (loss Loss) GetDescription() string {
	return Deref(loss.Description)
}

// GetExternalCode возвращает Внешний код Списания.
func (loss Loss) GetExternalCode() string {
	return Deref(loss.ExternalCode)
}

// GetFiles возвращает Метаданные массива Файлов.
func (loss Loss) GetFiles() MetaArray[File] {
	return Deref(loss.Files)
}

// GetGroup возвращает Отдел сотрудника.
func (loss Loss) GetGroup() Group {
	return Deref(loss.Group)
}

// GetID возвращает ID Списания.
func (loss Loss) GetID() uuid.UUID {
	return Deref(loss.ID)
}

// GetMeta возвращает Метаданные Списания.
func (loss Loss) GetMeta() Meta {
	return Deref(loss.Meta)
}

// GetSalesReturn возвращает Ссылку на связанный со списанием возврат покупателя.
func (loss Loss) GetSalesReturn() SalesReturn {
	return Deref(loss.SalesReturn)
}

// GetApplicable возвращает Отметку о проведении.
func (loss Loss) GetApplicable() bool {
	return Deref(loss.Applicable)
}

// GetProject возвращает Метаданные проекта.
func (loss Loss) GetProject() Project {
	return Deref(loss.Project).getValue()
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (loss Loss) GetOwner() Employee {
	return Deref(loss.Owner)
}

// GetPositions возвращает Метаданные позиций Списания.
func (loss Loss) GetPositions() MetaArray[LossPosition] {
	return Deref(loss.Positions)
}

// GetPrinted возвращает true, если документ напечатан.
func (loss Loss) GetPrinted() bool {
	return Deref(loss.Printed)
}

// GetAccountID возвращает ID учётной записи.
func (loss Loss) GetAccountID() uuid.UUID {
	return Deref(loss.AccountID)
}

// GetPublished возвращает true, если документ опубликован.
func (loss Loss) GetPublished() bool {
	return Deref(loss.Published)
}

// GetRate возвращает Валюту.
func (loss Loss) GetRate() Rate {
	return Deref(loss.Rate).getValue()
}

// GetShared возвращает флаг Общего доступа.
func (loss Loss) GetShared() bool {
	return Deref(loss.Shared)
}

// GetState возвращает Метаданные статуса Списания.
func (loss Loss) GetState() State {
	return Deref(loss.State).getValue()
}

// GetStore возвращает Метаданные склада.
func (loss Loss) GetStore() Store {
	return Deref(loss.Store)
}

// GetSum возвращает Сумму Списания в копейках.
func (loss Loss) GetSum() float64 {
	return Deref(loss.Sum)
}

// GetName возвращает Наименование Списания.
func (loss Loss) GetName() string {
	return Deref(loss.Name)
}

// GetUpdated возвращает Момент последнего обновления Списания.
func (loss Loss) GetUpdated() Timestamp {
	return Deref(loss.Updated)
}

// GetAttributes возвращает Список метаданных доп. полей.
func (loss Loss) GetAttributes() Slice[Attribute] {
	return loss.Attributes
}

// SetOrganization устанавливает Метаданные юрлица.
func (loss *Loss) SetOrganization(organization *Organization) *Loss {
	if organization != nil {
		loss.Organization = organization.Clean()
	}
	return loss
}

// SetSyncID устанавливает ID синхронизации.
func (loss *Loss) SetSyncID(syncID uuid.UUID) *Loss {
	loss.SyncID = &syncID
	return loss
}

// SetMoment устанавливает Дату документа.
func (loss *Loss) SetMoment(moment time.Time) *Loss {
	loss.Moment = NewTimestamp(moment)
	return loss
}

// SetCode устанавливает Код Списания.
func (loss *Loss) SetCode(code string) *Loss {
	loss.Code = &code
	return loss
}

// SetDescription устанавливает Комментарий Списания.
func (loss *Loss) SetDescription(description string) *Loss {
	loss.Description = &description
	return loss
}

// SetExternalCode устанавливает Внешний код Списания.
func (loss *Loss) SetExternalCode(externalCode string) *Loss {
	loss.ExternalCode = &externalCode
	return loss
}

// SetFiles устанавливает Метаданные массива Файлов.
//
// Принимает множество объектов [File].
func (loss *Loss) SetFiles(files ...*File) *Loss {
	loss.Files = NewMetaArrayFrom(files)
	return loss
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (loss *Loss) SetGroup(group *Group) *Loss {
	if group != nil {
		loss.Group = group.Clean()
	}
	return loss
}

// SetMeta устанавливает Метаданные Списания.
func (loss *Loss) SetMeta(meta *Meta) *Loss {
	loss.Meta = meta
	return loss
}

// SetSalesReturn устанавливает Ссылку на связанный со списанием возврат покупателя.
func (loss *Loss) SetSalesReturn(salesReturn *SalesReturn) *Loss {
	if salesReturn != nil {
		loss.SalesReturn = salesReturn.Clean()
	}
	return loss
}

// SetApplicable устанавливает Отметку о проведении.
func (loss *Loss) SetApplicable(applicable bool) *Loss {
	loss.Applicable = &applicable
	return loss
}

// SetProject устанавливает Метаданные проекта.
//
// Передача nil передаёт сброс значения (null).
func (loss *Loss) SetProject(project *Project) *Loss {
	loss.Project = NewNullValue(project)
	return loss
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (loss *Loss) SetOwner(owner *Employee) *Loss {
	if owner != nil {
		loss.Owner = owner.Clean()
	}
	return loss
}

// SetPositions устанавливает Метаданные позиций Списания.
//
// Принимает множество объектов [LossPosition].
func (loss *Loss) SetPositions(positions ...*LossPosition) *Loss {
	loss.Positions = NewMetaArrayFrom(positions)
	return loss
}

// SetRate устанавливает Валюту.
//
// Передача nil передаёт сброс значения (null).
func (loss *Loss) SetRate(rate *Rate) *Loss {
	loss.Rate = NewNullValue(rate)
	return loss
}

// SetShared устанавливает флаг общего доступа.
func (loss *Loss) SetShared(shared bool) *Loss {
	loss.Shared = &shared
	return loss
}

// SetState устанавливает Метаданные статуса Списания.
//
// Передача nil передаёт сброс значения (null).
func (loss *Loss) SetState(state *State) *Loss {
	loss.State = NewNullValue(state)
	return loss
}

// SetStore устанавливает Метаданные склада.
func (loss *Loss) SetStore(store *Store) *Loss {
	if store != nil {
		loss.Store = store.Clean()
	}
	return loss
}

// SetName устанавливает Наименование Списания.
func (loss *Loss) SetName(name string) *Loss {
	loss.Name = &name
	return loss
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (loss *Loss) SetAttributes(attributes ...*Attribute) *Loss {
	loss.Attributes.Push(attributes...)
	return loss
}

// String реализует интерфейс [fmt.Stringer].
func (loss Loss) String() string {
	return Stringify(loss)
}

// MetaType возвращает код сущности.
func (Loss) MetaType() MetaType {
	return MetaTypeLoss
}

// Update shortcut
func (loss Loss) Update(ctx context.Context, client *Client, params ...*Params) (*Loss, *resty.Response, error) {
	return NewLossService(client).Update(ctx, loss.GetID(), &loss, params...)
}

// Create shortcut
func (loss Loss) Create(ctx context.Context, client *Client, params ...*Params) (*Loss, *resty.Response, error) {
	return NewLossService(client).Create(ctx, &loss, params...)
}

// Delete shortcut
func (loss Loss) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewLossService(client).Delete(ctx, loss.GetID())
}

// LossPosition Позиция Списания.
//
// Код сущности: lossposition
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-spisanie-spisaniq-pozicii-spisaniq
type LossPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учётной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Price      *float64            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Reason     *string             `json:"reason,omitempty"`     // Причина списания данной позиции
	Slot       *Slot               `json:"slot,omitempty"`       // Ячейка на складе
	Things     Slice[string]       `json:"things,omitempty"`     // Серийные номера. Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете. В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута.
}

// GetAccountID возвращает ID учётной записи.
func (lossPosition LossPosition) GetAccountID() uuid.UUID {
	return Deref(lossPosition.AccountID)
}

// GetAssortment возвращает Метаданные товара/услуги/серии/модификации, которую представляет собой позиция.
func (lossPosition LossPosition) GetAssortment() AssortmentPosition {
	return Deref(lossPosition.Assortment)
}

// GetID возвращает ID позиции.
func (lossPosition LossPosition) GetID() uuid.UUID {
	return Deref(lossPosition.ID)
}

// GetPack возвращает Упаковку Товара.
func (lossPosition LossPosition) GetPack() Pack {
	return Deref(lossPosition.Pack)
}

// GetPrice возвращает Цену товара/услуги в копейках.
func (lossPosition LossPosition) GetPrice() float64 {
	return Deref(lossPosition.Price)
}

// GetQuantity возвращает Количество товаров данного вида в позиции.
func (lossPosition LossPosition) GetQuantity() float64 {
	return Deref(lossPosition.Quantity)
}

// GetReason возвращает Причину списания данной позиции.
func (lossPosition LossPosition) GetReason() string {
	return Deref(lossPosition.Reason)
}

// GetSlot возвращает Ячейку на складе.
func (lossPosition LossPosition) GetSlot() Slot {
	return Deref(lossPosition.Slot)
}

// GetThings возвращает Серийные номера.
//
// Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете.
// В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута.
func (lossPosition LossPosition) GetThings() Slice[string] {
	return lossPosition.Things
}

// SetAssortment устанавливает Метаданные товара/услуги, которую представляет собой компонент.
//
// Принимает объект, реализующий интерфейс [AssortmentConverter].
func (lossPosition *LossPosition) SetAssortment(assortment AssortmentConverter) *LossPosition {
	if assortment != nil {
		lossPosition.Assortment = assortment.AsAssortment()
	}
	return lossPosition
}

// SetPack устанавливает Упаковку Товара.
func (lossPosition *LossPosition) SetPack(pack *Pack) *LossPosition {
	if pack != nil {
		lossPosition.Pack = pack
	}
	return lossPosition
}

// SetPrice устанавливает Цену товара/услуги в копейках.
func (lossPosition *LossPosition) SetPrice(price float64) *LossPosition {
	lossPosition.Price = &price
	return lossPosition
}

// SetQuantity устанавливает Количество товаров данного вида в позиции.
//
// Если позиция - товар, у которого включен учет по серийным номерам,
// то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
func (lossPosition *LossPosition) SetQuantity(quantity float64) *LossPosition {
	lossPosition.Quantity = &quantity
	return lossPosition
}

// SetReason устанавливает Причину списания данной позиции.
func (lossPosition *LossPosition) SetReason(reason string) *LossPosition {
	lossPosition.Reason = &reason
	return lossPosition
}

// SetSlot устанавливает Ячейку на складе.
func (lossPosition *LossPosition) SetSlot(slot *Slot) *LossPosition {
	if slot != nil {
		lossPosition.Slot = slot.Clean()
	}
	return lossPosition
}

// SetThings устанавливает Серийные номера.
//
// Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете.
// В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута.
//
// Принимает множество string.
func (lossPosition *LossPosition) SetThings(things ...string) *LossPosition {
	lossPosition.Things = NewSliceFrom(things)
	return lossPosition
}

// String реализует интерфейс [fmt.Stringer].
func (lossPosition LossPosition) String() string {
	return Stringify(lossPosition)
}

// MetaType возвращает код сущности.
func (LossPosition) MetaType() MetaType {
	return MetaTypeLossPosition
}

// LossService описывает методы сервиса для работы со списаниями.
type LossService interface {
	// GetList выполняет запрос на получение списка списаний.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[Loss], *resty.Response, error)

	// Create выполняет запрос на создание списания.
	// Обязательные поля для заполнения:
	//	- organization (Ссылка на ваше юрлицо)
	//	- store (Ссылка на склад)
	// Принимает контекст, списание и опционально объект параметров запроса Params.
	// Возвращает созданное списание.
	Create(ctx context.Context, loss *Loss, params ...*Params) (*Loss, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение списаний.
	// Изменяемые списания должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список списаний и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых списаний.
	CreateUpdateMany(ctx context.Context, lossList Slice[Loss], params ...*Params) (*Slice[Loss], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление списаний.
	// Принимает контекст и множество списаний.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*Loss) (*DeleteManyResponse, *resty.Response, error)

	// Delete выполняет запрос на удаление списания.
	// Принимает контекст и ID списания.
	// Возвращает «true» в случае успешного удаления списания.
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельного списания по ID.
	// Принимает контекст, ID списания и опционально объект параметров запроса Params.
	// Возвращает найденное списание.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Loss, *resty.Response, error)

	// Update выполняет запрос на изменение списания.
	// Принимает контекст, списание и опционально объект параметров запроса Params.
	// Возвращает изменённое списание.
	Update(ctx context.Context, id uuid.UUID, loss *Loss, params ...*Params) (*Loss, *resty.Response, error)

	// Template выполняет запрос на получение предзаполненного списания со стандартными полями без связи с какими-либо другими документами.
	// Принимает контекст.
	// Возвращает предзаполненное списание.
	Template(ctx context.Context) (*Loss, *resty.Response, error)

	// TemplateBased выполняет запрос на получение шаблона списания на основе других документов.
	// Основание, на котором может быть создано:
	//	- Возврат покупателя (SalesReturn)
	//	- Инвентаризация (Inventory)
	// Принимает контекст и множество документов из списка выше.
	// Возвращает предзаполненное списание на основании переданных документов.
	TemplateBased(ctx context.Context, basedOn ...MetaOwner) (*Loss, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных списаний.
	// Принимает контекст.
	// Возвращает объект метаданных MetaAttributesStatesSharedWrapper.
	GetMetadata(ctx context.Context) (*MetaAttributesStatesSharedWrapper, *resty.Response, error)

	// GetPositionList выполняет запрос на получение списка позиций документа.
	// Принимает контекст, ID документа и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetPositionList(ctx context.Context, id uuid.UUID, params ...*Params) (*List[LossPosition], *resty.Response, error)

	// GetPositionByID выполняет запрос на получение отдельной позиции документа по ID.
	// Принимает контекст, ID документа, ID позиции и опционально объект параметров запроса Params.
	// Возвращает найденную позицию.
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params ...*Params) (*LossPosition, *resty.Response, error)

	// UpdatePosition выполняет запрос на изменение позиции документа.
	// Принимает контекст, ID документа, ID позиции, позицию документа и опционально объект параметров запроса Params.
	// Возвращает изменённую позицию.
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *LossPosition, params ...*Params) (*LossPosition, *resty.Response, error)

	// CreatePosition выполняет запрос на добавление позиции документа.
	// Принимает контекст, ID документа, позицию документа и опционально объект параметров запроса Params.
	// Возвращает добавленную позицию.
	CreatePosition(ctx context.Context, id uuid.UUID, position *LossPosition, params ...*Params) (*LossPosition, *resty.Response, error)

	// CreatePositionMany выполняет запрос на массовое добавление позиций документа.
	// Принимает контекст, ID документа и множество позиций.
	// Возвращает список добавленных позиций.
	CreatePositionMany(ctx context.Context, id uuid.UUID, positions ...*LossPosition) (*Slice[LossPosition], *resty.Response, error)

	// DeletePosition выполняет запрос на удаление позиции документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает «true» в случае успешного удаления позиции.
	DeletePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (bool, *resty.Response, error)

	// DeletePositionMany выполняет запрос на массовое удаление позиций документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeletePositionMany(ctx context.Context, id uuid.UUID, positions ...*LossPosition) (*DeleteManyResponse, *resty.Response, error)

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
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*Loss, *resty.Response, error)

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
	Evaluate(ctx context.Context, entity *Loss, evaluate ...Evaluate) (*Loss, *resty.Response, error)
}

// NewLossService принимает [Client] и возвращает сервис для работы со списаниями.
func NewLossService(client *Client) LossService {
	return newMainService[Loss, LossPosition, MetaAttributesStatesSharedWrapper, any](NewEndpoint(client, "entity/loss"))
}
