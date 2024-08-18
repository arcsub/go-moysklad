package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"time"
)

// PriceList Прайс-лист.
//
// Код сущности: pricelist
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-prajs-list
type PriceList struct {
	Meta         *Meta                         `json:"meta,omitempty"`         // Метаданные Прайс-листа
	Columns      Slice[PriceListColumn]        `json:"columns,omitempty"`      // Массив столбцов описания таблицы
	AccountID    *uuid.UUID                    `json:"accountId,omitempty"`    // ID учётной записи
	Code         *string                       `json:"code,omitempty"`         // Код Прайс-листа
	Moment       *Timestamp                    `json:"moment,omitempty"`       // Дата документа
	Created      *Timestamp                    `json:"created,omitempty"`      // Дата создания
	Deleted      *Timestamp                    `json:"deleted,omitempty"`      // Момент последнего удаления Прайс-листа
	Name         *string                       `json:"name,omitempty"`         // Наименование Прайс-листа
	ExternalCode *string                       `json:"externalCode,omitempty"` // Внешний код Прайс-листа
	Files        *MetaArray[File]              `json:"files,omitempty"`        // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group        *Group                        `json:"group,omitempty"`        // Отдел сотрудника
	ID           *uuid.UUID                    `json:"id,omitempty"`           // ID Прайс-листа
	Updated      *Timestamp                    `json:"updated,omitempty"`      // Момент последнего обновления Прайс-листа
	Applicable   *bool                         `json:"applicable,omitempty"`   // Отметка о проведении
	Description  *string                       `json:"description,omitempty"`  // Комментарий Прайс-листа
	Organization *Organization                 `json:"organization,omitempty"` // Метаданные юрлица
	Owner        *Employee                     `json:"owner,omitempty"`        // Метаданные владельца (Сотрудника)
	Positions    *MetaArray[PriceListPosition] `json:"positions,omitempty"`    // Метаданные позиций Прайс-листа
	PriceType    *PriceType                    `json:"priceType,omitempty"`    // Объект типа цены
	Printed      *bool                         `json:"printed,omitempty"`      // Напечатан ли документ
	Published    *bool                         `json:"published,omitempty"`    // Опубликован ли документ
	Shared       *bool                         `json:"shared,omitempty"`       // Общий доступ
	State        *NullValue[State]             `json:"state,omitempty"`        // Метаданные статуса Прайс-листа
	SyncID       *uuid.UUID                    `json:"syncId,omitempty"`       // ID синхронизации
	Attributes   Slice[Attribute]              `json:"attributes,omitempty"`   // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (priceList PriceList) Clean() *PriceList {
	if priceList.Meta == nil {
		return nil
	}
	return &PriceList{Meta: priceList.Meta}
}

// AsTaskOperation реализует интерфейс [TaskOperationConverter].
func (priceList PriceList) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: priceList.Meta}
}

// GetMeta возвращает Метаданные Прайс-листа.
func (priceList PriceList) GetMeta() Meta {
	return Deref(priceList.Meta)
}

// GetColumns возвращает Массив столбцов описания таблицы.
func (priceList PriceList) GetColumns() Slice[PriceListColumn] {
	return priceList.Columns
}

// GetAccountID возвращает ID учётной записи.
func (priceList PriceList) GetAccountID() uuid.UUID {
	return Deref(priceList.AccountID)
}

// GetCode возвращает Код Прайс-листа.
func (priceList PriceList) GetCode() string {
	return Deref(priceList.Code)
}

// GetMoment возвращает Дату документа.
func (priceList PriceList) GetMoment() time.Time {
	return Deref(priceList.Moment).Time()
}

// GetCreated возвращает Дату создания.
func (priceList PriceList) GetCreated() time.Time {
	return Deref(priceList.Created).Time()
}

// GetDeleted возвращает Момент последнего удаления Прайс-листа.
func (priceList PriceList) GetDeleted() time.Time {
	return Deref(priceList.Deleted).Time()
}

// GetName возвращает Наименование Прайс-листа.
func (priceList PriceList) GetName() string {
	return Deref(priceList.Name)
}

// GetExternalCode возвращает Внешний код Прайс-листа.
func (priceList PriceList) GetExternalCode() string {
	return Deref(priceList.ExternalCode)
}

// GetFiles возвращает Метаданные массива Файлов.
func (priceList PriceList) GetFiles() MetaArray[File] {
	return Deref(priceList.Files)
}

// GetGroup возвращает Отдел сотрудника.
func (priceList PriceList) GetGroup() Group {
	return Deref(priceList.Group)
}

// GetID возвращает ID Прайс-листа.
func (priceList PriceList) GetID() uuid.UUID {
	return Deref(priceList.ID)
}

// GetUpdated возвращает Момент последнего обновления Прайс-листа.
func (priceList PriceList) GetUpdated() time.Time {
	return Deref(priceList.Updated).Time()
}

// GetApplicable возвращает Отметку о проведении.
func (priceList PriceList) GetApplicable() bool {
	return Deref(priceList.Applicable)
}

// GetDescription возвращает Комментарий Прайс-листа.
func (priceList PriceList) GetDescription() string {
	return Deref(priceList.Description)
}

// GetOrganization возвращает Метаданные юрлица.
func (priceList PriceList) GetOrganization() Organization {
	return Deref(priceList.Organization)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (priceList PriceList) GetOwner() Employee {
	return Deref(priceList.Owner)
}

// GetPositions возвращает Метаданные позиций Прайс-листа.
func (priceList PriceList) GetPositions() MetaArray[PriceListPosition] {
	return Deref(priceList.Positions)
}

// GetPriceType возвращает Объект типа цены.
func (priceList PriceList) GetPriceType() PriceType {
	return Deref(priceList.PriceType)
}

// GetPrinted возвращает true, если документ напечатан.
func (priceList PriceList) GetPrinted() bool {
	return Deref(priceList.Printed)
}

// GetPublished возвращает true, если документ опубликован.
func (priceList PriceList) GetPublished() bool {
	return Deref(priceList.Published)
}

// GetShared возвращает флаг Общего доступа.
func (priceList PriceList) GetShared() bool {
	return Deref(priceList.Shared)
}

// GetState возвращает Метаданные статуса Прайс-листа.
func (priceList PriceList) GetState() State {
	return Deref(priceList.State).getValue()
}

// GetSyncID возвращает ID синхронизации.
func (priceList PriceList) GetSyncID() uuid.UUID {
	return Deref(priceList.SyncID)
}

// GetAttributes возвращает Список метаданных доп. полей.
func (priceList PriceList) GetAttributes() Slice[Attribute] {
	return priceList.Attributes
}

// SetMeta устанавливает Метаданные Прайс-листа.
func (priceList *PriceList) SetMeta(meta *Meta) *PriceList {
	priceList.Meta = meta
	return priceList
}

// SetColumns устанавливает Массив столбцов описания таблицы.
//
// Принимает множество объектов [PriceListColumn].
func (priceList *PriceList) SetColumns(columns ...*PriceListColumn) *PriceList {
	priceList.Columns.Push(columns...)
	return priceList
}

// SetCode устанавливает Код Прайс-листа.
func (priceList *PriceList) SetCode(code string) *PriceList {
	priceList.Code = &code
	return priceList
}

// SetMoment устанавливает Дату документа.
func (priceList *PriceList) SetMoment(moment time.Time) *PriceList {
	priceList.Moment = NewTimestamp(moment)
	return priceList
}

// SetName устанавливает Наименование Прайс-листа.
func (priceList *PriceList) SetName(name string) *PriceList {
	priceList.Name = &name
	return priceList
}

// SetExternalCode устанавливает Внешний код Прайс-листа.
func (priceList *PriceList) SetExternalCode(externalCode string) *PriceList {
	priceList.ExternalCode = &externalCode
	return priceList
}

// SetFiles устанавливает Метаданные массива Файлов.
//
// Принимает множество объектов [File].
func (priceList *PriceList) SetFiles(files ...*File) *PriceList {
	priceList.Files = NewMetaArrayFrom(files)
	return priceList
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (priceList *PriceList) SetGroup(group *Group) *PriceList {
	if group != nil {
		priceList.Group = group.Clean()
	}
	return priceList
}

// SetApplicable устанавливает Отметку о проведении.
func (priceList *PriceList) SetApplicable(applicable bool) *PriceList {
	priceList.Applicable = &applicable
	return priceList
}

// SetDescription устанавливает Комментарий Прайс-листа.
func (priceList *PriceList) SetDescription(description string) *PriceList {
	priceList.Description = &description
	return priceList
}

// SetOrganization устанавливает Метаданные юрлица.
func (priceList *PriceList) SetOrganization(organization *Organization) *PriceList {
	if organization != nil {
		priceList.Organization = organization.Clean()
	}
	return priceList
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (priceList *PriceList) SetOwner(owner *Employee) *PriceList {
	if owner != nil {
		priceList.Owner = owner.Clean()
	}
	return priceList
}

// SetPositions устанавливает Метаданные позиций Прайс-листа.
//
// Принимает множество объектов [PriceListPosition].
func (priceList *PriceList) SetPositions(positions ...*PriceListPosition) *PriceList {
	priceList.Positions = NewMetaArrayFrom(positions)
	return priceList
}

// SetShared устанавливает флаг общего доступа.
func (priceList *PriceList) SetShared(shared bool) *PriceList {
	priceList.Shared = &shared
	return priceList
}

// SetState устанавливает Метаданные статуса Прайс-листа.
//
// Передача nil передаёт сброс значения (null).
func (priceList *PriceList) SetState(state *State) *PriceList {
	priceList.State = NewNullValue(state)
	return priceList
}

// SetSyncID устанавливает ID синхронизации.
func (priceList *PriceList) SetSyncID(syncID uuid.UUID) *PriceList {
	priceList.SyncID = &syncID
	return priceList
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (priceList *PriceList) SetAttributes(attributes ...*Attribute) *PriceList {
	priceList.Attributes.Push(attributes...)
	return priceList
}

// String реализует интерфейс [fmt.Stringer].
func (priceList PriceList) String() string {
	return Stringify(priceList)
}

// MetaType возвращает код сущности.
func (PriceList) MetaType() MetaType {
	return MetaTypePriceList
}

// Update shortcut
func (priceList *PriceList) Update(ctx context.Context, client *Client, params ...*Params) (*PriceList, *resty.Response, error) {
	return NewPriceListService(client).Update(ctx, priceList.GetID(), priceList, params...)
}

// Create shortcut
func (priceList *PriceList) Create(ctx context.Context, client *Client, params ...*Params) (*PriceList, *resty.Response, error) {
	return NewPriceListService(client).Create(ctx, priceList, params...)
}

// Delete shortcut
func (priceList *PriceList) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewPriceListService(client).Delete(ctx, priceList)
}

// PriceListCell Ячейка прайс листа.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-prajs-list-prajs-listy-yachejki
type PriceListCell struct {
	Column *string  `json:"column,omitempty"` // Название столбца, к которому относится данная ячейка
	Sum    *float64 `json:"sum,omitempty"`    // Числовое значение ячейки
}

// GetColumn возвращает Название столбца, к которому относится данная ячейка.
func (priceListCell PriceListCell) GetColumn() string {
	return Deref(priceListCell.Column)
}

// GetSum возвращает Числовое значение ячейки.
func (priceListCell PriceListCell) GetSum() float64 {
	return Deref(priceListCell.Sum)
}

// SetColumn устанавливает Название столбца, к которому относится данная ячейка.
func (priceListCell *PriceListCell) SetColumn(column string) *PriceListCell {
	priceListCell.Column = &column
	return priceListCell
}

// SetSum устанавливает Числовое значение ячейки.
func (priceListCell *PriceListCell) SetSum(sum float64) *PriceListCell {
	priceListCell.Sum = &sum
	return priceListCell
}

// String реализует интерфейс [fmt.Stringer].
func (priceListCell PriceListCell) String() string {
	return Stringify(priceListCell)
}

// PriceListColumn Столбец прайс листа.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-prajs-list-prajs-listy-stolbcy
type PriceListColumn struct {
	Name               *string  `json:"name,omitempty"`               // Название столбца
	PercentageDiscount *float64 `json:"percentageDiscount,omitempty"` // Процентная наценка или скидка по умолчанию для столбца
}

// GetName возвращает Название столбца.
func (priceListColumn PriceListColumn) GetName() string {
	return Deref(priceListColumn.Name)
}

// GetPercentageDiscount возвращает Процентную наценку или скидку по умолчанию для столбца.
func (priceListColumn PriceListColumn) GetPercentageDiscount() float64 {
	return Deref(priceListColumn.PercentageDiscount)
}

// SetName устанавливает Название столбца.
func (priceListColumn *PriceListColumn) SetName(name string) *PriceListColumn {
	priceListColumn.Name = &name
	return priceListColumn
}

// SetPercentageDiscount устанавливает Процентную наценку или скидку по умолчанию для столбца.
func (priceListColumn *PriceListColumn) SetPercentageDiscount(percentageDiscount float64) *PriceListColumn {
	priceListColumn.PercentageDiscount = &percentageDiscount
	return priceListColumn
}

// String реализует интерфейс [fmt.Stringer].
func (priceListColumn PriceListColumn) String() string {
	return Stringify(priceListColumn)
}

// PriceListPosition Позиция прайс листа.
//
// Код сущности: pricelistrow
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-prajs-list-prajs-listy-pozicii-prajs-lista
type PriceListPosition struct {
	AccountID  *uuid.UUID           `json:"accountId,omitempty"`  // ID учётной записи
	Assortment *AssortmentPosition  `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	ID         *uuid.UUID           `json:"id,omitempty"`         // ID позиции
	Pack       *Pack                `json:"pack,omitempty"`       // Упаковка Товара
	Cells      Slice[PriceListCell] `json:"cells,omitempty"`      // Массив значений столбцов в позиции Прайс-листа
}

// GetAccountID возвращает ID учётной записи.
func (priceListPosition PriceListPosition) GetAccountID() uuid.UUID {
	return Deref(priceListPosition.AccountID)
}

// GetAssortment возвращает Метаданные товара/услуги/серии/модификации, которую представляет собой позиция.
func (priceListPosition PriceListPosition) GetAssortment() AssortmentPosition {
	return Deref(priceListPosition.Assortment)
}

// GetCells возвращает Массив значений столбцов в позиции Прайс-листа.
func (priceListPosition PriceListPosition) GetCells() Slice[PriceListCell] {
	return priceListPosition.Cells
}

// GetID возвращает ID позиции.
func (priceListPosition PriceListPosition) GetID() uuid.UUID {
	return Deref(priceListPosition.ID)
}

// GetPack возвращает Упаковку Товара.
func (priceListPosition PriceListPosition) GetPack() Pack {
	return Deref(priceListPosition.Pack)
}

// SetAssortment устанавливает Метаданные товара/услуги, которую представляет собой компонент.
//
// Принимает объект, реализующий интерфейс [AssortmentConverter].
func (priceListPosition *PriceListPosition) SetAssortment(assortment AssortmentConverter) *PriceListPosition {
	if assortment != nil {
		priceListPosition.Assortment = assortment.AsAssortment()
	}
	return priceListPosition
}

// SetCells устанавливает Массив значений столбцов в позиции Прайс-листа.
//
// Принимает множество объектов [PriceListCell].
func (priceListPosition *PriceListPosition) SetCells(cells ...*PriceListCell) *PriceListPosition {
	priceListPosition.Cells.Push(cells...)
	return priceListPosition
}

// SetPack устанавливает Упаковку Товара.
func (priceListPosition *PriceListPosition) SetPack(pack *Pack) *PriceListPosition {
	if pack != nil {
		priceListPosition.Pack = pack
	}
	return priceListPosition
}

// String реализует интерфейс [fmt.Stringer].
func (priceListPosition PriceListPosition) String() string {
	return Stringify(priceListPosition)
}

// MetaType возвращает код сущности.
func (PriceListPosition) MetaType() MetaType {
	return MetaTypePriceListPosition
}

// PriceListService описывает методы сервиса для работы с прайс-листами.
type PriceListService interface {
	// GetList выполняет запрос на получение списка прайс-листов.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[PriceList], *resty.Response, error)

	// GetListAll выполняет запрос на получение всех прайс-листов в виде списка.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает список объектов.
	GetListAll(ctx context.Context, params ...*Params) (Slice[PriceList], *resty.Response, error)

	// Create выполняет запрос на создание прайс-листа.
	// Обязательные поля для заполнения:
	//	- columns (Массив объектов, описывающих столбцы нового прайс-листа)
	// Принимает контекст, прайс-лист и опционально объект параметров запроса Params.
	// Возвращает созданный прайс-листа.
	Create(ctx context.Context, priceList *PriceList, params ...*Params) (*PriceList, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение прайс-листа.
	// Изменяемые прайс-листы должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список прайс-листов и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых прайс-листов.
	CreateUpdateMany(ctx context.Context, priceListList Slice[PriceList], params ...*Params) (*Slice[PriceList], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление прайс-листов.
	// Принимает контекст и множество прайс-листов.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*PriceList) (*DeleteManyResponse, *resty.Response, error)

	// DeleteByID выполняет запрос на удаление прайс-листа по ID.
	// Принимает контекст и ID прайс-листа.
	// Возвращает «true» в случае успешного удаления прайс-листа.
	DeleteByID(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// Delete выполняет запрос на удаление прайс-листа.
	// Принимает контекст и прайс-лист.
	// Возвращает «true» в случае успешного удаления прайс-листа.
	Delete(ctx context.Context, entity *PriceList) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельного прайс-листа по ID.
	// Принимает контекст, ID прайс-листа и опционально объект параметров запроса Params.
	// Возвращает найденный прайс-лист.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*PriceList, *resty.Response, error)

	// Update выполняет запрос на изменение прайс-листа.
	// Принимает контекст, прайс-лист и опционально объект параметров запроса Params.
	// Возвращает изменённый прайс-лист.
	Update(ctx context.Context, id uuid.UUID, priceList *PriceList, params ...*Params) (*PriceList, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных прайс-листов.
	// Принимает контекст.
	// Возвращает объект метаданных MetaAttributesStatesSharedWrapper.
	GetMetadata(ctx context.Context) (*MetaAttributesStatesSharedWrapper, *resty.Response, error)

	// GetPositionList выполняет запрос на получение списка позиций документа.
	// Принимает контекст, ID документа и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetPositionList(ctx context.Context, id uuid.UUID, params ...*Params) (*List[PriceListPosition], *resty.Response, error)

	// GetPositionByID выполняет запрос на получение отдельной позиции документа по ID.
	// Принимает контекст, ID документа, ID позиции и опционально объект параметров запроса Params.
	// Возвращает найденную позицию.
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params ...*Params) (*PriceListPosition, *resty.Response, error)

	// UpdatePosition выполняет запрос на изменение позиции документа.
	// Принимает контекст, ID документа, ID позиции, позицию документа и опционально объект параметров запроса Params.
	// Возвращает изменённую позицию.
	UpdatePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID, position *PriceListPosition, params ...*Params) (*PriceListPosition, *resty.Response, error)

	// CreatePosition выполняет запрос на добавление позиции документа.
	// Принимает контекст, ID документа, позицию документа и опционально объект параметров запроса Params.
	// Возвращает добавленную позицию.
	CreatePosition(ctx context.Context, id uuid.UUID, position *PriceListPosition, params ...*Params) (*PriceListPosition, *resty.Response, error)

	// CreatePositionMany выполняет запрос на массовое добавление позиций документа.
	// Принимает контекст, ID документа и множество позиций.
	// Возвращает список добавленных позиций.
	CreatePositionMany(ctx context.Context, id uuid.UUID, positions ...*PriceListPosition) (*Slice[PriceListPosition], *resty.Response, error)

	// DeletePosition выполняет запрос на удаление позиции документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает «true» в случае успешного удаления позиции.
	DeletePosition(ctx context.Context, id uuid.UUID, positionID uuid.UUID) (bool, *resty.Response, error)

	// DeletePositionMany выполняет запрос на массовое удаление позиций документа.
	// Принимает контекст, ID документа и ID позиции.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeletePositionMany(ctx context.Context, id uuid.UUID, positions ...*PriceListPosition) (*DeleteManyResponse, *resty.Response, error)

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
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*PriceList, *resty.Response, error)

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
	EndpointPriceList = EndpointEntity + string(MetaTypePriceList)
)

// NewPriceListService принимает [Client] и возвращает сервис для работы с прайс-листами.
func NewPriceListService(client *Client) PriceListService {
	return newMainService[PriceList, PriceListPosition, MetaAttributesStatesSharedWrapper, any](client, EndpointPriceList)
}
