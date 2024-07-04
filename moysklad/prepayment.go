package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Prepayment Предоплата.
//
// Код сущности: prepayment
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-predoplata
type Prepayment struct {
	Returns       Slice[PrepaymentReturn]        `json:"returns,omitempty"`       // Коллекция метаданных на связанные возвраты
	Owner         *Employee                      `json:"owner,omitempty"`         // Метаданные владельца (Сотрудника)
	Applicable    *bool                          `json:"applicable,omitempty"`    // Отметка о проведении
	Agent         *Agent                         `json:"agent,omitempty"`         // Метаданные контрагента
	CashSum       *float64                       `json:"cashSum,omitempty"`       // Оплачено наличными
	Code          *string                        `json:"code,omitempty"`          // Код Предоплаты
	Created       *Timestamp                     `json:"created,omitempty"`       // Дата создания
	CustomerOrder *CustomerOrder                 `json:"customerOrder,omitempty"` // Метаданные Заказа Покупателя
	Deleted       *Timestamp                     `json:"deleted,omitempty"`       // Момент последнего удаления Предоплаты
	Description   *string                        `json:"description,omitempty"`   // Комментарий Предоплаты
	ExternalCode  *string                        `json:"externalCode,omitempty"`  // Внешний код Предоплаты
	Files         *MetaArray[File]               `json:"files,omitempty"`         // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group         *Group                         `json:"group,omitempty"`         // Отдел сотрудника
	ID            *uuid.UUID                     `json:"id,omitempty"`            // ID Предоплаты
	Meta          *Meta                          `json:"meta,omitempty"`          // Метаданные Предоплаты
	Moment        *Timestamp                     `json:"moment,omitempty"`        // Дата документа
	Name          *string                        `json:"name,omitempty"`          // Наименование Предоплаты
	NoCashSum     *float64                       `json:"noCashSum,omitempty"`     // Оплачено картой
	AccountID     *uuid.UUID                     `json:"accountId,omitempty"`     // ID учётной записи
	VatIncluded   *bool                          `json:"vatIncluded,omitempty"`   // Включен ли НДС в цену
	Positions     *MetaArray[PrepaymentPosition] `json:"positions,omitempty"`     // Метаданные позиций Предоплаты
	Printed       *bool                          `json:"printed,omitempty"`       // Напечатан ли документ
	Published     *bool                          `json:"published,omitempty"`     // Опубликован ли документ
	QRSum         *float64                       `json:"qrSum,omitempty"`         // Оплачено по QR-коду
	Rate          *NullValue[Rate]               `json:"rate,omitempty"`          // Валюта
	RetailShift   *RetailShift                   `json:"retailShift,omitempty"`   // Метаданные Розничной смены
	RetailStore   *RetailStore                   `json:"retailStore,omitempty"`   // Метаданные Точки продаж
	Organization  *Organization                  `json:"organization,omitempty"`  // Метаданные юрлица
	Shared        *bool                          `json:"shared,omitempty"`        // Общий доступ
	State         *State                         `json:"state,omitempty"`         // Метаданные статуса Предоплаты
	Sum           *float64                       `json:"sum,omitempty"`           // Сумма Предоплаты в копейках
	SyncID        *uuid.UUID                     `json:"syncId,omitempty"`        // ID синхронизации
	VatSum        *float64                       `json:"vatSum,omitempty"`        // Сумма НДС
	Updated       *Timestamp                     `json:"updated,omitempty"`       // Момент последнего обновления Предоплаты
	VatEnabled    *bool                          `json:"vatEnabled,omitempty"`    // Учитывается ли НДС
	TaxSystem     TaxSystem                      `json:"taxSystem,omitempty"`     // Код системы налогообложения
	Attributes    Slice[Attribute]               `json:"attributes,omitempty"`    // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (prepayment Prepayment) Clean() *Prepayment {
	if prepayment.Meta == nil {
		return nil
	}
	return &Prepayment{Meta: prepayment.Meta}
}

// AsTaskOperation реализует интерфейс [TaskOperationInterface].
func (prepayment Prepayment) AsTaskOperation() *TaskOperation {
	return &TaskOperation{Meta: prepayment.Meta}
}

// GetReturns возвращает Коллекцию метаданных на связанные возвраты.
func (prepayment Prepayment) GetReturns() Slice[PrepaymentReturn] {
	return prepayment.Returns
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (prepayment Prepayment) GetOwner() Employee {
	return Deref(prepayment.Owner)
}

// GetApplicable возвращает Отметку о проведении.
func (prepayment Prepayment) GetApplicable() bool {
	return Deref(prepayment.Applicable)
}

// GetAgent возвращает Метаданные контрагента.
func (prepayment Prepayment) GetAgent() Agent {
	return Deref(prepayment.Agent)
}

// GetCashSum возвращает Оплачено наличными.
func (prepayment Prepayment) GetCashSum() float64 {
	return Deref(prepayment.CashSum)
}

// GetCode возвращает Код Предоплаты.
func (prepayment Prepayment) GetCode() string {
	return Deref(prepayment.Code)
}

// GetCreated возвращает Дату создания.
func (prepayment Prepayment) GetCreated() Timestamp {
	return Deref(prepayment.Created)
}

// GetCustomerOrder возвращает Метаданные Заказа Покупателя.
func (prepayment Prepayment) GetCustomerOrder() CustomerOrder {
	return Deref(prepayment.CustomerOrder)
}

// GetDeleted возвращает Момент последнего удаления Предоплаты.
func (prepayment Prepayment) GetDeleted() Timestamp {
	return Deref(prepayment.Deleted)
}

// GetDescription возвращает Комментарий Предоплаты.
func (prepayment Prepayment) GetDescription() string {
	return Deref(prepayment.Description)
}

// GetExternalCode возвращает Внешний код Предоплаты.
func (prepayment Prepayment) GetExternalCode() string {
	return Deref(prepayment.ExternalCode)
}

// GetFiles возвращает Метаданные массива Файлов.
func (prepayment Prepayment) GetFiles() MetaArray[File] {
	return Deref(prepayment.Files)
}

// GetGroup возвращает Отдел сотрудника.
func (prepayment Prepayment) GetGroup() Group {
	return Deref(prepayment.Group)
}

// GetID возвращает ID Предоплаты.
func (prepayment Prepayment) GetID() uuid.UUID {
	return Deref(prepayment.ID)
}

// GetMeta возвращает Метаданные Предоплаты.
func (prepayment Prepayment) GetMeta() Meta {
	return Deref(prepayment.Meta)
}

// GetMoment возвращает Дату документа.
func (prepayment Prepayment) GetMoment() Timestamp {
	return Deref(prepayment.Moment)
}

// GetName возвращает Наименование Предоплаты.
func (prepayment Prepayment) GetName() string {
	return Deref(prepayment.Name)
}

// GetNoCashSum возвращает Оплачено картой.
func (prepayment Prepayment) GetNoCashSum() float64 {
	return Deref(prepayment.NoCashSum)
}

// GetAccountID возвращает ID учётной записи.
func (prepayment Prepayment) GetAccountID() uuid.UUID {
	return Deref(prepayment.AccountID)
}

// GetVatIncluded возвращает true, если НДС включен в цену.
func (prepayment Prepayment) GetVatIncluded() bool {
	return Deref(prepayment.VatIncluded)
}

// GetPositions возвращает Метаданные позиций Предоплаты.
func (prepayment Prepayment) GetPositions() MetaArray[PrepaymentPosition] {
	return Deref(prepayment.Positions)
}

// GetPrinted возвращает true, если документ напечатан.
func (prepayment Prepayment) GetPrinted() bool {
	return Deref(prepayment.Printed)
}

// GetPublished возвращает true, если документ опубликован.
func (prepayment Prepayment) GetPublished() bool {
	return Deref(prepayment.Published)
}

// GetQRSum возвращает оплачено по QR-коду.
func (prepayment Prepayment) GetQRSum() float64 {
	return Deref(prepayment.QRSum)
}

// GetRate возвращает Валюту.
func (prepayment Prepayment) GetRate() Rate {
	return prepayment.Rate.GetValue()
}

// GetRetailShift возвращает Метаданные Розничной смены.
func (prepayment Prepayment) GetRetailShift() RetailShift {
	return Deref(prepayment.RetailShift)
}

// GetRetailStore возвращает Метаданные Точки продаж.
func (prepayment Prepayment) GetRetailStore() RetailStore {
	return Deref(prepayment.RetailStore)
}

// GetOrganization возвращает Метаданные юрлица.
func (prepayment Prepayment) GetOrganization() Organization {
	return Deref(prepayment.Organization)
}

// GetShared возвращает флаг Общего доступа.
func (prepayment Prepayment) GetShared() bool {
	return Deref(prepayment.Shared)
}

// GetState возвращает Метаданные статуса Предоплаты.
func (prepayment Prepayment) GetState() State {
	return Deref(prepayment.State)
}

// GetSum возвращает Сумму Перемещения в копейках.
func (prepayment Prepayment) GetSum() float64 {
	return Deref(prepayment.Sum)
}

// GetSyncID возвращает ID синхронизации.
func (prepayment Prepayment) GetSyncID() uuid.UUID {
	return Deref(prepayment.SyncID)
}

// GetVatSum возвращает Сумму НДС.
func (prepayment Prepayment) GetVatSum() float64 {
	return Deref(prepayment.VatSum)
}

// GetUpdated возвращает Момент последнего обновления Предоплаты.
func (prepayment Prepayment) GetUpdated() Timestamp {
	return Deref(prepayment.Updated)
}

// GetVatEnabled возвращает true, если учитывается НДС.
func (prepayment Prepayment) GetVatEnabled() bool {
	return Deref(prepayment.VatEnabled)
}

// GetTaxSystem возвращает Код системы налогообложения.
func (prepayment Prepayment) GetTaxSystem() TaxSystem {
	return prepayment.TaxSystem
}

// GetAttributes возвращает Список метаданных доп. полей.
func (prepayment Prepayment) GetAttributes() Slice[Attribute] {
	return prepayment.Attributes
}

// String реализует интерфейс [fmt.Stringer].
func (prepayment Prepayment) String() string {
	return Stringify(prepayment)
}

// MetaType возвращает код сущности.
func (Prepayment) MetaType() MetaType {
	return MetaTypePrepayment
}

// PrepaymentPosition Позиция Предоплаты.
//
// Код сущности: prepaymentposition
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-predoplata-predoplaty-pozicii-predoplaty
type PrepaymentPosition struct {
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
func (prepaymentPosition PrepaymentPosition) GetAccountID() uuid.UUID {
	return Deref(prepaymentPosition.AccountID)
}

// GetAssortment возвращает Метаданные товара/услуги/серии/модификации, которую представляет собой позиция.
func (prepaymentPosition PrepaymentPosition) GetAssortment() AssortmentPosition {
	return Deref(prepaymentPosition.Assortment)
}

// GetDiscount возвращает Процент скидки или наценки.
//
// Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%.
func (prepaymentPosition PrepaymentPosition) GetDiscount() float64 {
	return Deref(prepaymentPosition.Discount)
}

// GetID возвращает ID позиции.
func (prepaymentPosition PrepaymentPosition) GetID() uuid.UUID {
	return Deref(prepaymentPosition.ID)
}

// GetPack возвращает Упаковку Товара.
func (prepaymentPosition PrepaymentPosition) GetPack() Pack {
	return Deref(prepaymentPosition.Pack)
}

// GetPrice возвращает Цену товара/услуги в копейках.
func (prepaymentPosition PrepaymentPosition) GetPrice() float64 {
	return Deref(prepaymentPosition.Price)
}

// GetQuantity возвращает Количество товаров/услуг данного вида в позиции.
//
// Если позиция - товар, у которого включен учет по серийным номерам,
// то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
func (prepaymentPosition PrepaymentPosition) GetQuantity() float64 {
	return Deref(prepaymentPosition.Quantity)
}

// GetVat возвращает НДС, которым облагается текущая позиция.
func (prepaymentPosition PrepaymentPosition) GetVat() int {
	return Deref(prepaymentPosition.Vat)
}

// GetVatEnabled возвращает true, если НДС включен для позиции.
//
// С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС".
// (vat = 0, vatEnabled = false) -> vat = "без НДС",
// (vat = 0, vatEnabled = true) -> vat = 0%.
func (prepaymentPosition PrepaymentPosition) GetVatEnabled() bool {
	return Deref(prepaymentPosition.VatEnabled)
}

// String реализует интерфейс [fmt.Stringer].
func (prepaymentPosition PrepaymentPosition) String() string {
	return Stringify(prepaymentPosition)
}

// MetaType возвращает код сущности.
func (PrepaymentPosition) MetaType() MetaType {
	return MetaTypePrepaymentPosition
}

// PrepaymentService описывает методы сервиса для работы с предоплатами.
type PrepaymentService interface {
	// GetList выполняет запрос на получение списка предоплат.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[Prepayment], *resty.Response, error)

	// Delete выполняет запрос на удаление предоплаты.
	// Принимает контекст и ID предоплаты.
	// Возвращает true в случае успешного удаления предоплаты.
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельной предоплаты по ID.
	// Принимает контекст, ID предоплаты и опционально объект параметров запроса Params.
	// Возвращает найденную предоплату.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Prepayment, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных предоплат.
	// Принимает контекст.
	// Возвращает объект метаданных MetaAttributesStatesSharedWrapper.
	GetMetadata(ctx context.Context) (*MetaAttributesStatesSharedWrapper, *resty.Response, error)

	// GetPositionList выполняет запрос на получение списка позиций документа.
	// Принимает контекст, ID документа и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetPositionList(ctx context.Context, id uuid.UUID, params ...*Params) (*List[PrepaymentPosition], *resty.Response, error)

	// GetPositionByID выполняет запрос на получение отдельной позиции документа по ID.
	// Принимает контекст, ID документа, ID позиции и опционально объект параметров запроса Params.
	// Возвращает найденную позицию.
	GetPositionByID(ctx context.Context, id uuid.UUID, positionID uuid.UUID, params ...*Params) (*PrepaymentPosition, *resty.Response, error)

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
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*Prepayment, *resty.Response, error)

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
}

// NewPrepaymentService принимает [Client] и возвращает сервис для работы с предоплатами.
func NewPrepaymentService(client *Client) PrepaymentService {
	return newMainService[Prepayment, PrepaymentPosition, MetaAttributesStatesSharedWrapper, any](NewEndpoint(client, "entity/prepayment"))
}
