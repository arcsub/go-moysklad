package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Consignment Серия.
//
// Код сущности: consignment
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-seriq
type Consignment struct {
	Meta         *Meta               `json:"meta,omitempty"`         // Метаданные Серии
	Barcodes     Slice[Barcode]      `json:"barcodes,omitempty"`     // Штрихкоды
	Code         *string             `json:"code,omitempty"`         // Код Серии
	Description  *string             `json:"description,omitempty"`  // Описание Серии
	ExternalCode *string             `json:"externalCode,omitempty"` // Внешний код Серии
	ID           *uuid.UUID          `json:"id,omitempty"`           // ID Серии
	AccountID    *uuid.UUID          `json:"accountId,omitempty"`    // ID учётной записи
	Name         *string             `json:"name,omitempty"`         // Наименование Серии. "Собирается" и отображается как "Наименование товара / Метка Серии"
	Assortment   *AssortmentPosition `json:"assortment,omitempty"`   // Метаданные товара
	Image        *NullValue[Image]   `json:"image,omitempty"`        // Изображение товара, к которому относится данная серия
	Label        *string             `json:"label,omitempty"`        // Метка Серии
	Updated      *Timestamp          `json:"updated,omitempty"`      // Момент последнего обновления сущности
	Attributes   Slice[Attribute]    `json:"attributes,omitempty"`   // Список метаданных доп. полей
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (consignment Consignment) Clean() *Consignment {
	return &Consignment{Meta: consignment.Meta}
}

// NewConsignmentFromAssortment пытается привести переданный в качестве аргумента [AssortmentPosition] к типу [Consignment].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает:
//   - указатель на [Consignment].
//   - nil в случае неудачи.
func NewConsignmentFromAssortment(assortmentPosition *AssortmentPosition) *Consignment {
	return UnmarshalAsType[Consignment](assortmentPosition)
}

// FromAssortment пытается привести переданный в качестве аргумента [AssortmentPosition] к типу [Consignment].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает:
//   - указатель на [Consignment].
//   - nil в случае неудачи.
func (consignment Consignment) FromAssortment(assortmentPosition *AssortmentPosition) *Consignment {
	return UnmarshalAsType[Consignment](assortmentPosition)
}

// AsAssortment возвращает [AssortmentPosition] с единственным заполненным полем [Meta].
func (consignment Consignment) AsAssortment() *AssortmentPosition {
	return &AssortmentPosition{Meta: consignment.GetMeta()}
}

// GetMeta возвращает Метаданные Серии.
func (consignment Consignment) GetMeta() Meta {
	return Deref(consignment.Meta)
}

// GetBarcodes возвращает Штрихкоды.
func (consignment Consignment) GetBarcodes() Slice[Barcode] {
	return consignment.Barcodes
}

// GetCode возвращает Код серии.
func (consignment Consignment) GetCode() string {
	return Deref(consignment.Code)
}

// GetDescription возвращает Описание Серии.
func (consignment Consignment) GetDescription() string {
	return Deref(consignment.Description)
}

// GetExternalCode возвращает Внешний код Серии.
func (consignment Consignment) GetExternalCode() string {
	return Deref(consignment.ExternalCode)
}

// GetID возвращает ID Серии.
func (consignment Consignment) GetID() uuid.UUID {
	return Deref(consignment.ID)
}

// GetAccountID возвращает ID учётной записи.
func (consignment Consignment) GetAccountID() uuid.UUID {
	return Deref(consignment.AccountID)
}

// GetName возвращает Наименование Серии. "Собирается" и отображается как "Наименование товара / Метка Серии".
func (consignment Consignment) GetName() string {
	return Deref(consignment.Name)
}

// GetAssortment возвращает Метаданные товара.
func (consignment Consignment) GetAssortment() AssortmentPosition {
	return Deref(consignment.Assortment)
}

// GetImage возвращает Изображение товара, к которому относится данная серия.
func (consignment Consignment) GetImage() Image {
	return consignment.Image.Get()
}

// GetLabel возвращает Метку Серии.
func (consignment Consignment) GetLabel() string {
	return Deref(consignment.Label)
}

// GetUpdated возвращает Момент последнего обновления сущности.
func (consignment Consignment) GetUpdated() Timestamp {
	return Deref(consignment.Updated)
}

// GetAttributes возвращает Список метаданных доп. полей.
func (consignment Consignment) GetAttributes() Slice[Attribute] {
	return consignment.Attributes
}

// SetMeta устанавливает Метаданные Серии.
func (consignment *Consignment) SetMeta(meta *Meta) *Consignment {
	consignment.Meta = meta
	return consignment
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
func (consignment *Consignment) SetBarcodes(barcodes ...*Barcode) *Consignment {
	consignment.Barcodes = barcodes
	return consignment
}

// SetCode устанавливает Код Серии.
func (consignment *Consignment) SetCode(code string) *Consignment {
	consignment.Code = &code
	return consignment
}

// SetDescription устанавливает Описание Серии.
func (consignment *Consignment) SetDescription(description string) *Consignment {
	consignment.Description = &description
	return consignment
}

// SetExternalCode устанавливает Внешний код Серии.
func (consignment *Consignment) SetExternalCode(externalCode string) *Consignment {
	consignment.ExternalCode = &externalCode
	return consignment
}

// SetAssortment устанавливает Метаданные товара.
//
// Принимает объект, реализующий интерфейс [AsAssortmentInterface].
func (consignment *Consignment) SetAssortment(assortment AsAssortmentInterface) *Consignment {
	consignment.Assortment = assortment.AsAssortment()
	return consignment
}

// SetLabel устанавливает Метку Серии.
func (consignment *Consignment) SetLabel(label string) *Consignment {
	consignment.Label = &label
	return consignment
}

// SetAttributes устанавливает Список метаданных доп. полей.
//
// Принимает множество объектов [Attribute].
func (consignment *Consignment) SetAttributes(attributes ...*Attribute) *Consignment {
	consignment.Attributes = attributes
	return consignment
}

// String реализует интерфейс [fmt.Stringer].
func (consignment Consignment) String() string {
	return Stringify(consignment)
}

// MetaType возвращает тип сущности.
func (Consignment) MetaType() MetaType {
	return MetaTypeConsignment
}

// Update shortcut
func (consignment Consignment) Update(ctx context.Context, client *Client, params ...*Params) (*Consignment, *resty.Response, error) {
	return NewConsignmentService(client).Update(ctx, consignment.GetID(), &consignment, params...)
}

// Create shortcut
func (consignment Consignment) Create(ctx context.Context, client *Client, params ...*Params) (*Consignment, *resty.Response, error) {
	return NewConsignmentService(client).Create(ctx, &consignment, params...)
}

// Delete shortcut
func (consignment Consignment) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewConsignmentService(client).Delete(ctx, consignment.GetID())
}

// ConsignmentService описывает методы сервиса для работы с сериями.
type ConsignmentService interface {
	// GetList выполняет запрос на получение списка серий.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[Consignment], *resty.Response, error)

	// Create выполняет запрос на создание серии.
	// Обязательные поля для заполнения:
	//	- label (Метка Серии)
	//	- assortment (Метаданные товара)
	// Принимает контекст, серию и опционально объект параметров запроса Params.
	// Возвращает созданную серию.
	Create(ctx context.Context, consignment *Consignment, params ...*Params) (*Consignment, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и обновление серий.
	// Обновляемые серии должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список серий и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или обновлённых серий.
	CreateUpdateMany(ctx context.Context, consignmentList Slice[Consignment], params ...*Params) (*Slice[Consignment], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление серий.
	// Принимает контекст и множество серий.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*Consignment) (*DeleteManyResponse, *resty.Response, error)

	// Delete выполняет запрос на удаление серии.
	// Принимает контекст и ID серии.
	// Возвращает true в случае успешного удаления серии.
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельной серии по ID.
	// Принимает контекст, ID серии и опционально объект параметров запроса Params.
	// Возвращает найденную серию.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Consignment, *resty.Response, error)

	// Update выполняет запрос на обновление серии.
	// Принимает контекст, серию и опционально объект параметров запроса Params.
	// Возвращает обновлённую серию.
	Update(ctx context.Context, id uuid.UUID, consignment *Consignment, params ...*Params) (*Consignment, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных серий.
	// Принимает контекст.
	// Возвращает объект метаданных MetaAttributesWrapper.
	GetMetadata(ctx context.Context) (*MetaAttributesWrapper, *resty.Response, error)

	// GetAttributes выполняет запрос на получение списка доп полей.
	// Принимает контекст.
	// Возвращает объект MetaArray.
	GetAttributes(ctx context.Context) (*MetaArray[Attribute], *resty.Response, error)

	// GetAttributeByID выполняет запрос на получение отдельного доп поля по ID.
	// Принимает контекст и ID доп поля.
	// Возвращает найденное доп поле.
	GetAttributeByID(ctx context.Context, id uuid.UUID) (*Attribute, *resty.Response, error)

	// CreateAttribute выполняет запрос на создание доп поля.
	// Принимает контекст и доп поле.
	// Возвращает созданное доп поле.
	CreateAttribute(ctx context.Context, attribute *Attribute) (*Attribute, *resty.Response, error)

	// CreateUpdateAttributeMany выполняет запрос на массовое создание и обновление доп полей.
	// Обновляемые доп поля должны содержать идентификатор в виде метаданных.
	// Принимает контекст и множество доп полей.
	// Возвращает список созданных и/или обновлённых доп полей.
	CreateUpdateAttributeMany(ctx context.Context, attributes ...*Attribute) (*Slice[Attribute], *resty.Response, error)

	// UpdateAttribute выполняет запрос на изменения доп поля.
	// Принимает контекст, ID доп поля и доп поле.
	// Возвращает изменённое доп поле.
	UpdateAttribute(ctx context.Context, id uuid.UUID, attr *Attribute) (*Attribute, *resty.Response, error)

	// DeleteAttribute выполняет запрос на удаление доп поля.
	// Принимает контекст и ID доп поля.
	// Возвращает true в случае успешного удаления доп поля.
	DeleteAttribute(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// DeleteAttributeMany выполняет запрос на массовое удаление доп полей.
	// Принимает контекст и множество доп полей.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteAttributeMany(ctx context.Context, attributes ...*Attribute) (*DeleteManyResponse, *resty.Response, error)

	// GetNamedFilters выполняет запрос на получение списка фильтров.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetNamedFilters(ctx context.Context, params ...*Params) (*List[NamedFilter], *resty.Response, error)

	// GetNamedFilterByID выполняет запрос на получение отдельного фильтра по ID.
	// Принимает контекст и ID фильтра.
	// Возвращает найденный фильтр.
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)
}

// NewConsignmentService возвращает сервис для работы с сериями.
func NewConsignmentService(client *Client) ConsignmentService {
	return newMainService[Consignment, any, MetaAttributesWrapper, any](NewEndpoint(client, "entity/consignment"))

}
