package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"time"
)

// Service Услуга.
//
// Код сущности: service
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-usluga
type Service struct {
	VatEnabled          *bool                `json:"vatEnabled,omitempty"`          // Включен ли НДС для услуги. С помощью этого флага для услуги можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
	Group               *Group               `json:"group,omitempty"`               // Отдел сотрудника
	Barcodes            Slice[Barcode]       `json:"barcodes,omitempty"`            // Штрихкоды Услуги
	Description         *string              `json:"description,omitempty"`         // Описание Услуги
	ExternalCode        *string              `json:"externalCode,omitempty"`        // Внешний код Услуги
	ID                  *uuid.UUID           `json:"id,omitempty"`                  // ID Услуги
	Meta                *Meta                `json:"meta,omitempty"`                // Метаданные Услуги
	Name                *string              `json:"name,omitempty"`                // Наименование Услуги
	Archived            *bool                `json:"archived,omitempty"`            // Добавлена ли Услуга в архив
	Files               *MetaArray[File]     `json:"files,omitempty"`               // Метаданные массива Файлов (Максимальное количество файлов - 100)
	BuyPrice            *BuyPrice            `json:"buyPrice,omitempty"`            // Закупочная цена
	DiscountProhibited  *bool                `json:"discountProhibited,omitempty"`  // Признак запрета скидок
	EffectiveVat        *int                 `json:"effectiveVat,omitempty"`        // Реальный НДС %
	EffectiveVatEnabled *bool                `json:"effectiveVatEnabled,omitempty"` // Дополнительный признак для определения разграничения реального НДС = 0 или "без НДС". (effectiveVat = 0, effectiveVatEnabled = false) -> "без НДС", (effectiveVat = 0, effectiveVatEnabled = true) -> 0%.
	UseParentVat        *bool                `json:"useParentVat,omitempty"`        // Используется ли ставка НДС родительской группы. Если true для единицы ассортимента будет применена ставка, установленная для родительской группы.
	Code                *string              `json:"code,omitempty"`                // Код Услуги
	MinPrice            *NullValue[MinPrice] `json:"minPrice,omitempty"`            // Минимальная цена
	Owner               *Employee            `json:"owner,omitempty"`               // Метаданные владельца (Сотрудника)
	PathName            *string              `json:"pathName,omitempty"`            // Наименование группы, в которую входит Услуга
	AccountID           *uuid.UUID           `json:"accountId,omitempty"`           // ID учётной записи
	ProductFolder       *ProductFolder       `json:"productFolder,omitempty"`       // Метаданные группы услуги
	SalePrices          Slice[SalePrice]     `json:"salePrices,omitempty"`          // Цены продажи
	Shared              *bool                `json:"shared,omitempty"`              // Общий доступ
	SyncID              *uuid.UUID           `json:"syncId,omitempty"`              // ID синхронизации
	Vat                 *int                 `json:"vat,omitempty"`                 // НДС %
	Uom                 *NullValue[Uom]      `json:"uom,omitempty"`                 // Единица измерения
	Updated             *Timestamp           `json:"updated,omitempty"`             // Момент последнего обновления Услуги
	PaymentItemType     PaymentItem          `json:"paymentItemType,omitempty"`     // Признак предмета расчета
	TaxSystem           TaxSystem            `json:"taxSystem,omitempty"`           // Код системы налогообложения
	Attributes          Slice[Attribute]     `json:"attributes,omitempty"`          // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (service Service) Clean() *Service {
	if service.Meta == nil {
		return nil
	}
	return &Service{Meta: service.Meta}
}

// NewServiceFromAssortment пытается привести переданный в качестве аргумента [AssortmentPosition] к типу [Service].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [Service] или nil в случае неудачи.
func NewServiceFromAssortment(assortmentPosition *AssortmentPosition) *Service {
	return UnmarshalAsType[Service](assortmentPosition)
}

// FromAssortment пытается привести переданный в качестве аргумента [AssortmentPosition] к типу [Service].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [Service] или nil в случае неудачи.
func (service Service) FromAssortment(assortmentPosition *AssortmentPosition) *Service {
	return UnmarshalAsType[Service](assortmentPosition)
}

// AsAssortment реализует интерфейс [AssortmentConverter].
func (service Service) AsAssortment() *AssortmentPosition {
	return &AssortmentPosition{Meta: service.GetMeta()}
}

// GetVatEnabled возвращает true, если учитывается НДС.
//
// С помощью этого флага для товара можно выставлять НДС = 0 или НДС = "без НДС".
//
// (vat = 0, vatEnabled = false) -> vat = "без НДС",
//
// (vat = 0, vatEnabled = true) -> vat = 0%.
func (service Service) GetVatEnabled() bool {
	return Deref(service.VatEnabled)
}

// GetGroup возвращает Отдел сотрудника.
func (service Service) GetGroup() Group {
	return Deref(service.Group)
}

// GetBarcodes возвращает Штрихкоды Услуги.
func (service Service) GetBarcodes() Slice[Barcode] {
	return service.Barcodes
}

// GetDescription возвращает Описание Услуги.
func (service Service) GetDescription() string {
	return Deref(service.Description)
}

// GetExternalCode возвращает Внешний код Услуги.
func (service Service) GetExternalCode() string {
	return Deref(service.ExternalCode)
}

// GetID возвращает ID Услуги.
func (service Service) GetID() uuid.UUID {
	return Deref(service.ID)
}

// GetMeta возвращает Метаданные Услуги.
func (service Service) GetMeta() Meta {
	return Deref(service.Meta)
}

// GetName возвращает Наименование Услуги.
func (service Service) GetName() string {
	return Deref(service.Name)
}

// GetArchived возвращает флаг нахождения в архиве.
func (service Service) GetArchived() bool {
	return Deref(service.Archived)
}

// GetFiles возвращает Метаданные массива Файлов.
func (service Service) GetFiles() MetaArray[File] {
	return Deref(service.Files)
}

// GetBuyPrice возвращает Закупочную цену.
func (service Service) GetBuyPrice() BuyPrice {
	return Deref(service.BuyPrice)
}

// GetDiscountProhibited возвращает Признак запрета скидок.
func (service Service) GetDiscountProhibited() bool {
	return Deref(service.DiscountProhibited)
}

// GetEffectiveVat возвращает Реальный НДС %.
func (service Service) GetEffectiveVat() int {
	return Deref(service.EffectiveVat)
}

// GetEffectiveVatEnabled возвращает Дополнительный признак для определения разграничения реального НДС = 0 или "без НДС".
//
// (effectiveVat = 0, effectiveVatEnabled = false) -> "без НДС"
//
// (effectiveVat = 0, effectiveVatEnabled = true) -> 0%.
func (service Service) GetEffectiveVatEnabled() bool {
	return Deref(service.EffectiveVatEnabled)
}

// GetUseParentVat возвращает флаг использования ставки НДС родительской группы.
//
// Если «true» для единицы ассортимента будет применена ставка, установленная для родительской группы.
func (service Service) GetUseParentVat() bool {
	return Deref(service.UseParentVat)
}

// GetCode возвращает Код Услуги.
func (service Service) GetCode() string {
	return Deref(service.Code)
}

// GetMinPrice возвращает Минимальную цену.
func (service Service) GetMinPrice() MinPrice {
	return Deref(service.MinPrice).getValue()
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (service Service) GetOwner() Employee {
	return Deref(service.Owner)
}

// GetPathName возвращает Наименование группы, в которую входит Услуга.
func (service Service) GetPathName() string {
	return Deref(service.PathName)
}

// GetAccountID возвращает ID учётной записи.
func (service Service) GetAccountID() uuid.UUID {
	return Deref(service.AccountID)
}

// GetProductFolder возвращает Метаданные группы услуги.
func (service Service) GetProductFolder() ProductFolder {
	return Deref(service.ProductFolder)
}

// GetSalePrices возвращает Цены продажи.
func (service Service) GetSalePrices() Slice[SalePrice] {
	return service.SalePrices
}

// GetShared возвращает флаг Общего доступа.
func (service Service) GetShared() bool {
	return Deref(service.Shared)
}

// GetSyncID возвращает ID синхронизации.
func (service Service) GetSyncID() uuid.UUID {
	return Deref(service.SyncID)
}

// GetVat возвращает НДС %.
func (service Service) GetVat() int {
	return Deref(service.Vat)
}

// GetUom возвращает Единицу измерения.
func (service Service) GetUom() Uom {
	return Deref(service.Uom).getValue()
}

// GetUpdated возвращает Момент последнего обновления Услуги.
func (service Service) GetUpdated() time.Time {
	return Deref(service.Updated).Time()
}

// GetPaymentItemType возвращает Признак предмета расчета.
func (service Service) GetPaymentItemType() PaymentItem {
	return service.PaymentItemType
}

// GetTaxSystem возвращает Код системы налогообложения.
func (service Service) GetTaxSystem() TaxSystem {
	return service.TaxSystem
}

// GetAttributes возвращает Список метаданных доп. полей.
func (service Service) GetAttributes() Slice[Attribute] {
	return service.Attributes
}

// SetVatEnabled устанавливает значение, учитывающее НДС для услуги.
func (service *Service) SetVatEnabled(vatEnabled bool) *Service {
	service.VatEnabled = &vatEnabled
	return service
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (service *Service) SetGroup(group *Group) *Service {
	if group != nil {
		service.Group = group.Clean()
	}
	return service
}

// SetBarcodes устанавливает Штрихкоды.
//
// Для обновления списка штрихкодов необходимо передавать их полный список, включающий как старые,
// так и новые значения. Отсутствующие значения штрихкодов при обновлении будут удалены.
// При обновлении списка штрихкодов валидируются только новые значения.
// Ранее сохраненные штрихкоды не валидируются.
// То есть, если один из старых штрихкодов не соответствует требованиям к валидации,
// то ошибки при обновлении списка не будет. Если на вход передан пустой список штрихкодов
// или список из пустых значений, то ранее созданные штрихкоды будут удалены.
//
// Принимает множество объектов [Barcode].
func (service *Service) SetBarcodes(barcodes ...*Barcode) *Service {
	service.Barcodes.Push(barcodes...)
	return service
}

// SetDescription устанавливает Описание услуги.
func (service *Service) SetDescription(description string) *Service {
	service.Description = &description
	return service
}

// SetExternalCode устанавливает Внешний код услуги.
func (service *Service) SetExternalCode(externalCode string) *Service {
	service.ExternalCode = &externalCode
	return service
}

// SetMeta устанавливает Метаданные услуги.
func (service *Service) SetMeta(meta *Meta) *Service {
	service.Meta = meta
	return service
}

// SetName устанавливает Наименование услуги.
func (service *Service) SetName(name string) *Service {
	service.Name = &name
	return service
}

// SetArchived устанавливает флаг нахождения в архиве.
func (service *Service) SetArchived(archived bool) *Service {
	service.Archived = &archived
	return service
}

// SetFiles устанавливает Метаданные массива Файлов.
//
// Принимает множество объектов [File].
func (service *Service) SetFiles(files ...*File) *Service {
	service.Files = NewMetaArrayFrom(files)
	return service
}

// SetBuyPrice устанавливает Закупочную цену.
func (service *Service) SetBuyPrice(buyPrice *BuyPrice) *Service {
	if buyPrice != nil {
		service.BuyPrice = buyPrice
	}
	return service
}

// SetDiscountProhibited устанавливает Признак запрета скидок.
func (service *Service) SetDiscountProhibited(discountProhibited bool) *Service {
	service.DiscountProhibited = &discountProhibited
	return service
}

// SetUseParentVat устанавливает флаг использования ставки НДС родительской группы.
//
// Если «true» для единицы ассортимента будет применена ставка, установленная для родительской группы.
func (service *Service) SetUseParentVat(useParentVat bool) *Service {
	service.UseParentVat = &useParentVat
	return service
}

// SetCode устанавливает Код услуги.
func (service *Service) SetCode(code string) *Service {
	service.Code = &code
	return service
}

// SetMinPrice устанавливает Минимальную цену.
//
// Передача nil передаёт сброс значения (null).
func (service *Service) SetMinPrice(minPrice *MinPrice) *Service {
	service.MinPrice = NewNullValue(minPrice)
	return service
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (service *Service) SetOwner(owner *Employee) *Service {
	if owner != nil {
		service.Owner = owner.Clean()
	}
	return service
}

// SetProductFolder устанавливает Метаданные группы услуги.
//
// Передача nil передаёт сброс значения (null).
func (service *Service) SetProductFolder(productFolder *ProductFolder) *Service {
	if productFolder != nil {
		service.ProductFolder = productFolder.Clean()
	}
	return service
}

// SetSalePrices устанавливает Цены продажи.
//
// Принимает множество объектов [SalePrice].
func (service *Service) SetSalePrices(salePrices ...*SalePrice) *Service {
	service.SalePrices.Push(salePrices...)
	return service
}

// SetShared устанавливает флаг общего доступа.
func (service *Service) SetShared(shared bool) *Service {
	service.Shared = &shared
	return service
}

// SetSyncID устанавливает ID синхронизации.
func (service *Service) SetSyncID(syncID uuid.UUID) *Service {
	service.SyncID = &syncID
	return service
}

// SetVat устанавливает НДС %.
func (service *Service) SetVat(vat int) *Service {
	service.Vat = &vat
	return service
}

// SetUom устанавливает Единицу измерения.
//
// Передача nil передаёт сброс значения (null).
func (service *Service) SetUom(uom *Uom) *Service {
	service.Uom = NewNullValue(uom)
	return service
}

// SetPaymentItemType устанавливает Признак предмета расчета.
func (service *Service) SetPaymentItemType(paymentItemType PaymentItem) *Service {
	service.PaymentItemType = paymentItemType
	return service
}

// SetTaxSystem устанавливает Код системы налогообложения.
func (service *Service) SetTaxSystem(taxSystem TaxSystem) *Service {
	service.TaxSystem = taxSystem
	return service
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (service *Service) SetAttributes(attributes ...*Attribute) *Service {
	service.Attributes.Push(attributes...)
	return service
}

// String реализует интерфейс [fmt.Stringer].
func (service Service) String() string {
	return Stringify(service)
}

// MetaType возвращает код сущности.
func (Service) MetaType() MetaType {
	return MetaTypeService
}

// Update shortcut
func (service *Service) Update(ctx context.Context, client *Client, params ...*Params) (*Service, *resty.Response, error) {
	return NewServiceService(client).Update(ctx, service.GetID(), service, params...)
}

// Create shortcut
func (service *Service) Create(ctx context.Context, client *Client, params ...*Params) (*Service, *resty.Response, error) {
	return NewServiceService(client).Create(ctx, service, params...)
}

// Delete shortcut
func (service *Service) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewServiceService(client).Delete(ctx, service)
}

// ServiceService описывает методы сервиса для работы с услугами.
type ServiceService interface {
	// GetList выполняет запрос на получение списка услуг.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[Service], *resty.Response, error)

	// GetListAll выполняет запрос на получение всех услуг в виде списка.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает список объектов.
	GetListAll(ctx context.Context, params ...*Params) (*Slice[Service], *resty.Response, error)

	// Create выполняет запрос на создание услуги.
	// Обязательные поля для заполнения:
	//	- name (Наименование услуги)
	// Принимает контекст, услугу и опционально объект параметров запроса Params.
	// Возвращает созданную услугу.
	Create(ctx context.Context, service *Service, params ...*Params) (*Service, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение услуг.
	// Изменяемые услуги должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список услуг и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых услуг.
	CreateUpdateMany(ctx context.Context, serviceList Slice[Service], params ...*Params) (*Slice[Service], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление услуг.
	// Принимает контекст и множество услуг.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*Service) (*DeleteManyResponse, *resty.Response, error)

	// DeleteByID выполняет запрос на удаление услуги по ID.
	// Принимает контекст и ID услуги.
	// Возвращает «true» в случае успешного удаления услуги.
	DeleteByID(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// Delete выполняет запрос на удаление услуги.
	// Принимает контекст и услугу.
	// Возвращает «true» в случае успешного удаления услуги.
	Delete(ctx context.Context, entity *Service) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельной услуги по ID.
	// Принимает контекст, ID услуги и опционально объект параметров запроса Params.
	// Возвращает услугу.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Service, *resty.Response, error)

	// Update выполняет запрос на изменение услуги.
	// Принимает контекст, услугу и опционально объект параметров запроса Params.
	// Возвращает изменённую услугу.
	Update(ctx context.Context, id uuid.UUID, service *Service, params ...*Params) (*Service, *resty.Response, error)

	// GetBySyncID выполняет запрос на получение отдельного документа по syncID.
	// Принимает контекст и syncID документа.
	// Возвращает найденный документ.
	GetBySyncID(ctx context.Context, syncID uuid.UUID) (*Service, *resty.Response, error)

	// DeleteBySyncID выполняет запрос на удаление документа по syncID.
	// Принимает контекст и syncID документа.
	// Возвращает «true» в случае успешного удаления документа.
	DeleteBySyncID(ctx context.Context, syncID uuid.UUID) (bool, *resty.Response, error)

	// GetNamedFilterList выполняет запрос на получение списка фильтров.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetNamedFilterList(ctx context.Context, params ...*Params) (*List[NamedFilter], *resty.Response, error)

	// GetNamedFilterByID выполняет запрос на получение отдельного фильтра по ID.
	// Принимает контекст и ID фильтра.
	// Возвращает найденный фильтр.
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)

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
	EndpointService = EndpointEntity + string(MetaTypeService)
)

// NewServiceService принимает [Client] и возвращает сервис для работы с услугами.
func NewServiceService(client *Client) ServiceService {
	return newMainService[Service, any, any, any](client, EndpointService)
}
