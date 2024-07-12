package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// ProductFolder Группа товаров.
//
// Код сущности: productfolder
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-gruppa-towarow
type ProductFolder struct {
	Name                *string                   `json:"name,omitempty"`                // Наименование Группы товаров
	UseParentVat        *bool                     `json:"useParentVat,omitempty"`        // Используется ли ставка НДС родительской группы. Если «true» для единицы ассортимента будет применена ставка, установленная для родительской группы.
	Code                *string                   `json:"code,omitempty"`                // Код Группы товаров
	Description         *string                   `json:"description,omitempty"`         // Описание Группы товаров
	EffectiveVat        *int                      `json:"effectiveVat,omitempty"`        // Реальный НДС %
	EffectiveVatEnabled *bool                     `json:"effectiveVatEnabled,omitempty"` // Дополнительный признак для определения разграничения реального НДС = 0 или "без НДС". (effectiveVat = 0, effectiveVatEnabled = false) -> "без НДС", (effectiveVat = 0, effectiveVatEnabled = true) -> 0%.
	ExternalCode        *string                   `json:"externalCode,omitempty"`        // Внешний код Группы товаров
	AccountID           *uuid.UUID                `json:"accountId,omitempty"`           // ID учётной записи
	VatEnabled          *bool                     `json:"vatEnabled,omitempty"`          // Включен ли НДС для группы. С помощью этого флага для группы можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
	Archived            *bool                     `json:"archived,omitempty"`            // Добавлена ли Группа товаров в архив
	Group               *Group                    `json:"group,omitempty"`               // Отдел сотрудника
	Owner               *Employee                 `json:"owner,omitempty"`               // Метаданные владельца (Сотрудника)
	PathName            *string                   `json:"pathName,omitempty"`            // Наименование Группы товаров, в которую входит данная Группа товаров
	ProductFolder       *NullValue[ProductFolder] `json:"productFolder,omitempty"`       // Ссылка на Группу товаров, в которую входит данная Группа товаров, в формате Метаданных
	Shared              *bool                     `json:"shared,omitempty"`              // Общий доступ
	ID                  *uuid.UUID                `json:"id,omitempty"`                  // ID Группы товаров
	Updated             *Timestamp                `json:"updated,omitempty"`             // Момент последнего обновления Группы товаров
	Meta                *Meta                     `json:"meta,omitempty"`                // Метаданные Группы товаров
	Vat                 *int                      `json:"vat,omitempty"`                 // НДС %
	TaxSystem           TaxSystem                 `json:"taxSystem,omitempty"`           // Код системы налогообложения
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (productFolder ProductFolder) Clean() *ProductFolder {
	if productFolder.Meta == nil {
		return nil
	}
	return &ProductFolder{Meta: productFolder.Meta}
}

// GetName возвращает Наименование Группы товаров.
func (productFolder ProductFolder) GetName() string {
	return Deref(productFolder.Name)
}

// GetUseParentVat возвращает флаг использования ставки НДС родительской группы.
//
// Если «true» для единицы ассортимента будет применена ставка, установленная для родительской группы.
func (productFolder ProductFolder) GetUseParentVat() bool {
	return Deref(productFolder.UseParentVat)
}

// GetCode возвращает Код Группы товаров.
func (productFolder ProductFolder) GetCode() string {
	return Deref(productFolder.Code)
}

// GetDescription возвращает Описание Группы товаров.
func (productFolder ProductFolder) GetDescription() string {
	return Deref(productFolder.Description)
}

// GetEffectiveVat возвращает Реальный НДС %.
func (productFolder ProductFolder) GetEffectiveVat() int {
	return Deref(productFolder.EffectiveVat)
}

// GetEffectiveVatEnabled возвращает Дополнительный признак для определения разграничения реального НДС = 0 или "без НДС".
//
// (effectiveVat = 0, effectiveVatEnabled = false) -> "без НДС"
//
// (effectiveVat = 0, effectiveVatEnabled = true) -> 0%.
func (productFolder ProductFolder) GetEffectiveVatEnabled() bool {
	return Deref(productFolder.EffectiveVatEnabled)
}

// GetExternalCode возвращает Внешний код Группы товаров.
func (productFolder ProductFolder) GetExternalCode() string {
	return Deref(productFolder.ExternalCode)
}

// GetAccountID возвращает ID учётной записи.
func (productFolder ProductFolder) GetAccountID() uuid.UUID {
	return Deref(productFolder.AccountID)
}

// GetVatEnabled возвращает true, если учитывается НДС.
//
// С помощью этого флага для группы можно выставлять НДС = 0 или НДС = "без НДС".
//
// (vat = 0, vatEnabled = false) -> vat = "без НДС"
//
// (vat = 0, vatEnabled = true) -> vat = 0%.
func (productFolder ProductFolder) GetVatEnabled() bool {
	return Deref(productFolder.VatEnabled)
}

// GetArchived возвращает true, если группа товаров находится в архиве.
func (productFolder ProductFolder) GetArchived() bool {
	return Deref(productFolder.Archived)
}

// GetGroup возвращает Отдел сотрудника.
func (productFolder ProductFolder) GetGroup() Group {
	return Deref(productFolder.Group)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (productFolder ProductFolder) GetOwner() Employee {
	return Deref(productFolder.Owner)
}

// GetPathName возвращает Наименование родительской Группы товаров.
func (productFolder ProductFolder) GetPathName() string {
	return Deref(productFolder.PathName)
}

// GetProductFolder возвращает Ссылку на Группу товаров, в которую входит данная Группа товаров, в формате Метаданных.
func (productFolder ProductFolder) GetProductFolder() ProductFolder {
	return Deref(productFolder.ProductFolder).getValue()
}

// GetShared возвращает флаг Общего доступа.
func (productFolder ProductFolder) GetShared() bool {
	return Deref(productFolder.Shared)
}

// GetID возвращает ID Группы товаров.
func (productFolder ProductFolder) GetID() uuid.UUID {
	return Deref(productFolder.ID)
}

// GetUpdated возвращает Момент последнего обновления Группы товаров.
func (productFolder ProductFolder) GetUpdated() Timestamp {
	return Deref(productFolder.Updated)
}

// GetMeta возвращает Метаданные Группы товаров.
func (productFolder ProductFolder) GetMeta() Meta {
	return Deref(productFolder.Meta)
}

// GetVat возвращает НДС %.
func (productFolder ProductFolder) GetVat() int {
	return Deref(productFolder.Vat)
}

// GetTaxSystem возвращает Код системы налогообложения.
func (productFolder ProductFolder) GetTaxSystem() TaxSystem {
	return productFolder.TaxSystem
}

// SetName устанавливает Наименование Группы товаров.
func (productFolder *ProductFolder) SetName(name string) *ProductFolder {
	productFolder.Name = &name
	return productFolder
}

// SetUseParentVat устанавливает флаг использования ставки НДС родительской группы.
//
// Если «true» для единицы ассортимента будет применена ставка, установленная для родительской группы.
func (productFolder *ProductFolder) SetUseParentVat(useParentVat bool) *ProductFolder {
	productFolder.UseParentVat = &useParentVat
	return productFolder
}

// SetCode устанавливает Код Группы товаров.
func (productFolder *ProductFolder) SetCode(code string) *ProductFolder {
	productFolder.Code = &code
	return productFolder
}

// SetDescription устанавливает Описание Группы товаров.
func (productFolder *ProductFolder) SetDescription(description string) *ProductFolder {
	productFolder.Description = &description
	return productFolder
}

// SetExternalCode устанавливает Внешний код Группы товаров.
func (productFolder *ProductFolder) SetExternalCode(externalCode string) *ProductFolder {
	productFolder.ExternalCode = &externalCode
	return productFolder
}

// SetVatEnabled устанавливает значение, учитывающее НДС для Группы товаров.
//
// С помощью этого флага для группы можно выставлять НДС = 0 или НДС = "без НДС".
//
// (vat = 0, vatEnabled = false) -> vat = "без НДС"
//
// (vat = 0, vatEnabled = true) -> vat = 0%.
func (productFolder *ProductFolder) SetVatEnabled(vatEnabled bool) *ProductFolder {
	productFolder.VatEnabled = &vatEnabled
	return productFolder
}

// SetArchived устанавливает флаг нахождения группы товаров в архиве.
func (productFolder *ProductFolder) SetArchived(archived bool) *ProductFolder {
	productFolder.Archived = &archived
	return productFolder
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (productFolder *ProductFolder) SetGroup(group *Group) *ProductFolder {
	if group != nil {
		productFolder.Group = group.Clean()
	}
	return productFolder
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (productFolder *ProductFolder) SetOwner(owner *Employee) *ProductFolder {
	if owner != nil {
		productFolder.Owner = owner.Clean()
	}
	return productFolder
}

// SetProductFolder устанавливает Ссылку на Группу товаров, в которую входит данная Группа товаров, в формате Метаданных.
//
// Передача nil передаёт сброс значения (null).
func (productFolder *ProductFolder) SetProductFolder(parent *ProductFolder) *ProductFolder {
	productFolder.ProductFolder = NewNullValue(parent)
	return productFolder
}

// SetShared устанавливает флаг общего доступа.
func (productFolder *ProductFolder) SetShared(shared bool) *ProductFolder {
	productFolder.Shared = &shared
	return productFolder
}

// SetMeta устанавливает Метаданные Группы товаров.
func (productFolder *ProductFolder) SetMeta(meta *Meta) *ProductFolder {
	productFolder.Meta = meta
	return productFolder
}

// SetVat устанавливает НДС %.
func (productFolder *ProductFolder) SetVat(vat int) *ProductFolder {
	productFolder.Vat = &vat
	return productFolder
}

// SetTaxSystem устанавливает Код системы налогообложения.
func (productFolder *ProductFolder) SetTaxSystem(taxSystem TaxSystem) *ProductFolder {
	productFolder.TaxSystem = taxSystem
	return productFolder
}

// String реализует интерфейс [fmt.Stringer].
func (productFolder ProductFolder) String() string {
	return Stringify(productFolder)
}

// MetaType возвращает код сущности.
func (ProductFolder) MetaType() MetaType {
	return MetaTypeProductFolder
}

// Update shortcut
func (productFolder ProductFolder) Update(ctx context.Context, client *Client, params ...*Params) (*ProductFolder, *resty.Response, error) {
	return NewProductFolderService(client).Update(ctx, productFolder.GetID(), &productFolder, params...)
}

// Create shortcut
func (productFolder ProductFolder) Create(ctx context.Context, client *Client, params ...*Params) (*ProductFolder, *resty.Response, error) {
	return NewProductFolderService(client).Create(ctx, &productFolder, params...)
}

// Delete shortcut
func (productFolder ProductFolder) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewProductFolderService(client).Delete(ctx, productFolder.GetID())
}

// ProductFolderService описывает методы сервиса для работы с группами товаров.
type ProductFolderService interface {
	// GetList выполняет запрос на получение списка группы товаров.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[ProductFolder], *resty.Response, error)

	// Create выполняет запрос на создание группы товаров.
	// Обязательные поля для заполнения:
	//	- name (Наименование группы товаров)
	// Принимает контекст, группу товаров и опционально объект параметров запроса Params.
	// Возвращает созданную группу товаров.
	Create(ctx context.Context, productFolder *ProductFolder, params ...*Params) (*ProductFolder, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение групп товаров.
	// Изменяемые группы товаров должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список групп товаров и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых групп товаров.
	CreateUpdateMany(ctx context.Context, productFolderList Slice[ProductFolder], params ...*Params) (*Slice[ProductFolder], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление группы товаров.
	// Принимает контекст и множество групп товаров.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*ProductFolder) (*DeleteManyResponse, *resty.Response, error)

	// Delete выполняет запрос на удаление группы товаров.
	// Принимает контекст и ID группы товаров.
	// Возвращает «true» в случае успешного удаления группы товаров.
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельной группы товаров по ID.
	// Принимает контекст, ID группы товаров и опционально объект параметров запроса Params.
	// Возвращает найденную группу товаров.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*ProductFolder, *resty.Response, error)

	// Update выполняет запрос на изменение группы товаров.
	// Принимает контекст, группу товаров и опционально объект параметров запроса Params.
	// Возвращает изменённую группу товаров.
	Update(ctx context.Context, id uuid.UUID, productFolder *ProductFolder, params ...*Params) (*ProductFolder, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных групп товаров.
	// Принимает контекст.
	// Возвращает объект метаданных MetaAttributesWrapper.
	GetMetadata(ctx context.Context) (*MetaAttributesWrapper, *resty.Response, error)

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
}

const (
	EndpointProductFolder = EndpointEntity + string(MetaTypeProductFolder)
)

// NewProductFolderService принимает [Client] и возвращает сервис для работы с группами товаров.
func NewProductFolderService(client *Client) ProductFolderService {
	return newMainService[ProductFolder, any, MetaAttributesWrapper, any](client, EndpointProductFolder)
}
