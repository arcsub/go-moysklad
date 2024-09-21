package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"time"
)

// PrepaymentReturn Возврат предоплаты.
//
// Код сущности: prepaymentreturn
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vozwrat-predoplaty
type PrepaymentReturn struct {
	Agent        *Agent                               `json:"agent,omitempty"`        // Метаданные контрагента
	Organization *Organization                        `json:"organization,omitempty"` // Метаданные юрлица
	Applicable   *bool                                `json:"applicable,omitempty"`   // Отметка о проведении
	AccountID    *uuid.UUID                           `json:"accountId,omitempty"`    // ID учётной записи
	CashSum      *float64                             `json:"cashSum,omitempty"`      // Оплачено наличными
	Code         *string                              `json:"code,omitempty"`         // Код Возврата предоплаты
	Created      *Timestamp                           `json:"created,omitempty"`      // Дата создания
	Deleted      *Timestamp                           `json:"deleted,omitempty"`      // Момент последнего удаления Возврата предоплаты
	Description  *string                              `json:"description,omitempty"`  // Комментарий Возврата предоплаты
	ExternalCode *string                              `json:"externalCode,omitempty"` // Внешний код Возврата предоплаты
	Files        *MetaArray[File]                     `json:"files,omitempty"`        // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group        *Group                               `json:"group,omitempty"`        // Отдел сотрудника
	ID           *uuid.UUID                           `json:"id,omitempty"`           // ID Возврата предоплаты
	Meta         *Meta                                `json:"meta,omitempty"`         // Метаданные Возврата предоплаты
	Moment       *Timestamp                           `json:"moment,omitempty"`       // Дата документа
	Name         *string                              `json:"name,omitempty"`         // Наименование Возврата предоплаты
	NoCashSum    *float64                             `json:"noCashSum,omitempty"`    // Оплачено картой
	Owner        *Employee                            `json:"owner,omitempty"`        // Метаданные владельца (Сотрудника)
	VatIncluded  *bool                                `json:"vatIncluded,omitempty"`  // Включен ли НДС в цену
	Positions    *MetaArray[PrepaymentReturnPosition] `json:"positions,omitempty"`    // Метаданные позиций Возврата предоплаты
	Prepayment   *Prepayment                          `json:"prepayment,omitempty"`   // Метаданные Предоплаты
	Printed      *bool                                `json:"printed,omitempty"`      // Напечатан ли документ
	Published    *bool                                `json:"published,omitempty"`    // Опубликован ли документ
	QRSum        *float64                             `json:"qrSum,omitempty"`        // Оплачено по QR-коду
	Rate         *NullValue[Rate]                     `json:"rate,omitempty"`         // Валюта
	RetailShift  *RetailShift                         `json:"retailShift,omitempty"`  // Метаданные Розничной смены
	RetailStore  *RetailStore                         `json:"retailStore,omitempty"`  // Метаданные Точки продаж
	Shared       *bool                                `json:"shared,omitempty"`       // Общий доступ
	State        *State                               `json:"state,omitempty"`        // Метаданные статуса Возврата предоплаты
	Sum          *float64                             `json:"sum,omitempty"`          // Сумма Возврата предоплаты в копейках
	SyncID       *uuid.UUID                           `json:"syncId,omitempty"`       // ID синхронизации
	VatSum       *float64                             `json:"vatSum,omitempty"`       // Сумма НДС
	Updated      *Timestamp                           `json:"updated,omitempty"`      // Момент последнего обновления Возврата предоплаты
	VatEnabled   *bool                                `json:"vatEnabled,omitempty"`   // Учитывается ли НДС
	TaxSystem    TaxSystem                            `json:"taxSystem,omitempty"`    // Код системы налогообложения
	Attributes   Slice[Attribute]                     `json:"attributes,omitempty"`   // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (prepaymentReturn PrepaymentReturn) Clean() *PrepaymentReturn {
	if prepaymentReturn.Meta == nil {
		return nil
	}
	return &PrepaymentReturn{Meta: prepaymentReturn.Meta}
}

// AsTaskOperation реализует интерфейс [TaskOperationConverter].
func (prepaymentReturn PrepaymentReturn) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: prepaymentReturn.Meta}
}

// GetAgent возвращает Метаданные Контрагента.
func (prepaymentReturn PrepaymentReturn) GetAgent() Agent {
	return Deref(prepaymentReturn.Agent)
}

// GetOrganization возвращает Метаданные юрлица.
func (prepaymentReturn PrepaymentReturn) GetOrganization() Organization {
	return Deref(prepaymentReturn.Organization)
}

// GetApplicable возвращает Отметку о проведении.
func (prepaymentReturn PrepaymentReturn) GetApplicable() bool {
	return Deref(prepaymentReturn.Applicable)
}

// GetAccountID возвращает ID учётной записи.
func (prepaymentReturn PrepaymentReturn) GetAccountID() uuid.UUID {
	return Deref(prepaymentReturn.AccountID)
}

// GetCashSum возвращает Оплачено наличными.
func (prepaymentReturn PrepaymentReturn) GetCashSum() float64 {
	return Deref(prepaymentReturn.CashSum)
}

// GetCode возвращает Код Возврата предоплаты.
func (prepaymentReturn PrepaymentReturn) GetCode() string {
	return Deref(prepaymentReturn.Code)
}

// GetCreated возвращает Дату создания.
func (prepaymentReturn PrepaymentReturn) GetCreated() time.Time {
	return Deref(prepaymentReturn.Created).Time()
}

// GetDeleted возвращает Момент последнего удаления Возврата предоплаты.
func (prepaymentReturn PrepaymentReturn) GetDeleted() time.Time {
	return Deref(prepaymentReturn.Deleted).Time()
}

// GetDescription возвращает Комментарий Возврата предоплаты.
func (prepaymentReturn PrepaymentReturn) GetDescription() string {
	return Deref(prepaymentReturn.Description)
}

// GetExternalCode возвращает Внешний код Возврата предоплаты.
func (prepaymentReturn PrepaymentReturn) GetExternalCode() string {
	return Deref(prepaymentReturn.ExternalCode)
}

// GetFiles возвращает Метаданные массива Файлов.
func (prepaymentReturn PrepaymentReturn) GetFiles() MetaArray[File] {
	return Deref(prepaymentReturn.Files)
}

// GetGroup возвращает Отдел сотрудника.
func (prepaymentReturn PrepaymentReturn) GetGroup() Group {
	return Deref(prepaymentReturn.Group)
}

// GetID возвращает ID Возврата предоплаты.
func (prepaymentReturn PrepaymentReturn) GetID() uuid.UUID {
	return Deref(prepaymentReturn.ID)
}

// GetMeta возвращает Метаданные Возврата предоплаты.
func (prepaymentReturn PrepaymentReturn) GetMeta() Meta {
	return Deref(prepaymentReturn.Meta)
}

// GetMoment возвращает Дату документа.
func (prepaymentReturn PrepaymentReturn) GetMoment() time.Time {
	return Deref(prepaymentReturn.Moment).Time()
}

// GetName возвращает Наименование Возврата предоплаты.
func (prepaymentReturn PrepaymentReturn) GetName() string {
	return Deref(prepaymentReturn.Name)
}

// GetNoCashSum возвращает Оплачено картой.
func (prepaymentReturn PrepaymentReturn) GetNoCashSum() float64 {
	return Deref(prepaymentReturn.NoCashSum)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (prepaymentReturn PrepaymentReturn) GetOwner() Employee {
	return Deref(prepaymentReturn.Owner)
}

// GetVatIncluded возвращает true, если НДС включен в цену.
func (prepaymentReturn PrepaymentReturn) GetVatIncluded() bool {
	return Deref(prepaymentReturn.VatIncluded)
}

// GetPositions возвращает Метаданные позиций Возврата предоплаты.
func (prepaymentReturn PrepaymentReturn) GetPositions() MetaArray[PrepaymentReturnPosition] {
	return Deref(prepaymentReturn.Positions)
}

// GetPrepayment возвращает Метаданные Предоплаты.
func (prepaymentReturn PrepaymentReturn) GetPrepayment() Prepayment {
	return Deref(prepaymentReturn.Prepayment)
}

// GetPrinted возвращает true, если документ напечатан.
func (prepaymentReturn PrepaymentReturn) GetPrinted() bool {
	return Deref(prepaymentReturn.Printed)
}

// GetPublished возвращает true, если документ опубликован.
func (prepaymentReturn PrepaymentReturn) GetPublished() bool {
	return Deref(prepaymentReturn.Published)
}

// GetQRSum возвращает оплачено по QR-коду.
func (prepaymentReturn PrepaymentReturn) GetQRSum() float64 {
	return Deref(prepaymentReturn.QRSum)
}

// GetRate возвращает Валюту.
func (prepaymentReturn PrepaymentReturn) GetRate() Rate {
	return prepaymentReturn.Rate.getValue()
}

// GetRetailShift возвращает Метаданные Розничной смены.
func (prepaymentReturn PrepaymentReturn) GetRetailShift() RetailShift {
	return Deref(prepaymentReturn.RetailShift)
}

// GetRetailStore возвращает Метаданные Точки продаж.
func (prepaymentReturn PrepaymentReturn) GetRetailStore() RetailStore {
	return Deref(prepaymentReturn.RetailStore)
}

// GetShared возвращает флаг Общего доступа.
func (prepaymentReturn PrepaymentReturn) GetShared() bool {
	return Deref(prepaymentReturn.Shared)
}

// GetState возвращает Метаданные статуса Возврата предоплаты.
func (prepaymentReturn PrepaymentReturn) GetState() State {
	return Deref(prepaymentReturn.State)
}

// GetSum возвращает Сумму Возврата предоплаты в копейках.
func (prepaymentReturn PrepaymentReturn) GetSum() float64 {
	return Deref(prepaymentReturn.Sum)
}

// GetSyncID возвращает ID синхронизации.
func (prepaymentReturn PrepaymentReturn) GetSyncID() uuid.UUID {
	return Deref(prepaymentReturn.SyncID)
}

// GetVatSum возвращает Сумму НДС.
func (prepaymentReturn PrepaymentReturn) GetVatSum() float64 {
	return Deref(prepaymentReturn.VatSum)
}

// GetUpdated возвращает Момент последнего обновления Возврата предоплаты.
func (prepaymentReturn PrepaymentReturn) GetUpdated() time.Time {
	return Deref(prepaymentReturn.Updated).Time()
}

// GetVatEnabled возвращает true, если учитывается НДС.
func (prepaymentReturn PrepaymentReturn) GetVatEnabled() bool {
	return Deref(prepaymentReturn.VatEnabled)
}

// GetTaxSystem возвращает Код системы налогообложения.
func (prepaymentReturn PrepaymentReturn) GetTaxSystem() TaxSystem {
	return prepaymentReturn.TaxSystem
}

// GetAttributes возвращает Список метаданных доп. полей.
func (prepaymentReturn PrepaymentReturn) GetAttributes() Slice[Attribute] {
	return prepaymentReturn.Attributes
}

// String реализует интерфейс [fmt.Stringer].
func (prepaymentReturn PrepaymentReturn) String() string {
	return Stringify(prepaymentReturn)
}

// MetaType возвращает код сущности.
func (PrepaymentReturn) MetaType() MetaType {
	return MetaTypePrepaymentReturn
}

// PrepaymentReturnPosition Позиция Возврата предоплаты.
//
// Код сущности: prepaymentreturnposition
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vozwrat-predoplaty-atributy-suschnosti-pozicii-vozwrata-predoplaty
type PrepaymentReturnPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учётной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	Discount   *float64            `json:"discount,omitempty"`   // Процент скидки или наценки. Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Price      *float64            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Vat        *int                `json:"vat,omitempty"`        // НДС, которым облагается текущая позиция
	VatEnabled *bool               `json:"vatEnabled,omitempty"` // Включен ли НДС для позиции. С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
}

// GetAccountID возвращает ID учётной записи.
func (prepaymentReturnPosition PrepaymentReturnPosition) GetAccountID() uuid.UUID {
	return Deref(prepaymentReturnPosition.AccountID)
}

// GetAssortment возвращает Метаданные товара/услуги/серии/модификации, которую представляет собой позиция.
func (prepaymentReturnPosition PrepaymentReturnPosition) GetAssortment() AssortmentPosition {
	return Deref(prepaymentReturnPosition.Assortment)
}

// GetDiscount возвращает Процент скидки или наценки.
//
// Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%.
func (prepaymentReturnPosition PrepaymentReturnPosition) GetDiscount() float64 {
	return Deref(prepaymentReturnPosition.Discount)
}

// GetID возвращает ID позиции.
func (prepaymentReturnPosition PrepaymentReturnPosition) GetID() uuid.UUID {
	return Deref(prepaymentReturnPosition.ID)
}

// GetPack возвращает Упаковку Товара.
func (prepaymentReturnPosition PrepaymentReturnPosition) GetPack() Pack {
	return Deref(prepaymentReturnPosition.Pack)
}

// GetPrice возвращает Цену товара/услуги в копейках.
func (prepaymentReturnPosition PrepaymentReturnPosition) GetPrice() float64 {
	return Deref(prepaymentReturnPosition.Price)
}

// GetQuantity возвращает Количество товаров данного вида в позиции.
func (prepaymentReturnPosition PrepaymentReturnPosition) GetQuantity() float64 {
	return Deref(prepaymentReturnPosition.Quantity)
}

// GetVat возвращает НДС, которым облагается текущая позиция.
func (prepaymentReturnPosition PrepaymentReturnPosition) GetVat() int {
	return Deref(prepaymentReturnPosition.Vat)
}

// GetVatEnabled возвращает true, если НДС включен для позиции.
//
// С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС".
// (vat = 0, vatEnabled = false) -> vat = "без НДС",
// (vat = 0, vatEnabled = true) -> vat = 0%.
func (prepaymentReturnPosition PrepaymentReturnPosition) GetVatEnabled() bool {
	return Deref(prepaymentReturnPosition.VatEnabled)
}

// String реализует интерфейс [fmt.Stringer].
func (prepaymentReturnPosition PrepaymentReturnPosition) String() string {
	return Stringify(prepaymentReturnPosition)
}

// MetaType возвращает код сущности.
func (PrepaymentReturnPosition) MetaType() MetaType {
	return MetaTypePrepaymentReturnPosition
}

// PrepaymentReturnService описывает методы сервиса для работы с возвратами предоплат.
type PrepaymentReturnService interface {
	// GetList выполняет запрос на получение списка возвратов предоплат.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[PrepaymentReturn], *resty.Response, error)

	// GetListAll выполняет запрос на получение всех возвратов предоплат в виде списка.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает список объектов.
	GetListAll(ctx context.Context, params ...*Params) (*Slice[PrepaymentReturn], *resty.Response, error)

	// GetByID выполняет запрос на получение отдельного возврата предоплаты по ID.
	// Принимает контекст, ID возврата предоплаты и опционально объект параметров запроса Params.
	// Возвращает найденный возврат предоплаты.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*PrepaymentReturn, *resty.Response, error)

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

	// GetMetadata выполняет запрос на получение метаданных возвратов предоплаты.
	// Принимает контекст.
	// Возвращает объект метаданных MetaAttributesStatesSharedWrapper.
	GetMetadata(ctx context.Context) (*MetaAttributesStatesSharedWrapper, *resty.Response, error)

	// GetPositionList выполняет запрос на получение списка позиций документа.
	// Принимает контекст, ID документа и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetPositionList(ctx context.Context, id uuid.UUID, params ...*Params) (*List[PrepaymentReturnPosition], *resty.Response, error)

	GetPositionListAll(ctx context.Context, id uuid.UUID, params ...*Params) (*Slice[PrepaymentReturnPosition], *resty.Response, error)

	// GetPositionByID выполняет запрос на получение отдельной позиции документа по ID.
	// Принимает контекст, ID документа, ID позиции и опционально объект параметров запроса Params.
	// Возвращает найденную позицию.
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params ...*Params) (*PrepaymentReturnPosition, *resty.Response, error)

	// GetBySyncID выполняет запрос на получение отдельного документа по syncID.
	// Принимает контекст и syncID документа.
	// Возвращает найденный документ.
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*PrepaymentReturn, *resty.Response, error)

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
	EndpointPrepaymentReturn = EndpointEntity + string(MetaTypePrepaymentReturn)
)

// NewPrepaymentReturnService принимает [Client] и возвращает сервис для работы с возвратами предоплат.
func NewPrepaymentReturnService(client *Client) PrepaymentReturnService {
	return newMainService[PrepaymentReturn, PrepaymentReturnPosition, MetaAttributesStatesSharedWrapper, any](client, EndpointPrepaymentReturn)
}
