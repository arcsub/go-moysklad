package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"

	"time"
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
	ID           *string             `json:"id,omitempty"`           // ID Серии
	AccountID    *string             `json:"accountId,omitempty"`    // ID учётной записи
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
	if consignment.Meta == nil {
		return nil
	}
	return &Consignment{Meta: consignment.Meta}
}

// NewConsignmentFromAssortment пытается привести переданный в качестве аргумента [AssortmentPosition] к типу [Consignment].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [Consignment] или nil в случае неудачи.
func NewConsignmentFromAssortment(assortmentPosition *AssortmentPosition) *Consignment {
	return UnmarshalAsType[Consignment](assortmentPosition)
}

// FromAssortment пытается привести переданный в качестве аргумента [AssortmentPosition] к типу [Consignment].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [Consignment] или nil в случае неудачи.
func (consignment Consignment) FromAssortment(assortmentPosition *AssortmentPosition) *Consignment {
	return UnmarshalAsType[Consignment](assortmentPosition)
}

// AsAssortment реализует интерфейс [AssortmentConverter].
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
func (consignment Consignment) GetID() string {
	return Deref(consignment.ID)
}

// GetAccountID возвращает ID учётной записи.
func (consignment Consignment) GetAccountID() string {
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
	return Deref(consignment.Image).getValue()
}

// GetLabel возвращает Метку Серии.
func (consignment Consignment) GetLabel() string {
	return Deref(consignment.Label)
}

// GetUpdated возвращает Момент последнего обновления сущности.
func (consignment Consignment) GetUpdated() time.Time {
	return Deref(consignment.Updated).Time()
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
	consignment.Barcodes.Push(barcodes...)
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
// Принимает объект, реализующий интерфейс [AssortmentConverter].
func (consignment *Consignment) SetAssortment(assortment AssortmentConverter) *Consignment {
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
	consignment.Attributes.Push(attributes...)
	return consignment
}

// String реализует интерфейс [fmt.Stringer].
func (consignment Consignment) String() string {
	return Stringify(consignment)
}

// MetaType возвращает код сущности.
func (Consignment) MetaType() MetaType {
	return MetaTypeConsignment
}

// Update shortcut
func (consignment *Consignment) Update(ctx context.Context, client *Client, params ...func(*Params)) (*Consignment, *resty.Response, error) {
	return NewConsignmentService(client).Update(ctx, consignment.GetID(), consignment, params...)
}

// Create shortcut
func (consignment *Consignment) Create(ctx context.Context, client *Client, params ...func(*Params)) (*Consignment, *resty.Response, error) {
	return NewConsignmentService(client).Create(ctx, consignment, params...)
}

// Delete shortcut
func (consignment *Consignment) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewConsignmentService(client).Delete(ctx, consignment)
}

// ConsignmentService описывает методы сервиса для работы с сериями.
type ConsignmentService interface {
	// GetList выполняет запрос на получение списка серий.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...func(*Params)) (*List[Consignment], *resty.Response, error)

	// GetListAll выполняет запрос на получение всех серий в виде списка.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает список объектов.
	GetListAll(ctx context.Context, params ...func(*Params)) (*Slice[Consignment], *resty.Response, error)

	// Create выполняет запрос на создание серии.
	// Обязательные поля для заполнения:
	//	- label (Метка Серии)
	//	- assortment (Метаданные товара)
	// Принимает контекст, серию и опционально объект параметров запроса Params.
	// Возвращает созданную серию.
	Create(ctx context.Context, consignment *Consignment, params ...func(*Params)) (*Consignment, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение серий.
	// Изменяемые серии должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список серий и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых серий.
	CreateUpdateMany(ctx context.Context, consignmentList Slice[Consignment], params ...func(*Params)) (*Slice[Consignment], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление серий.
	// Принимает контекст и множество серий.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*Consignment) (*DeleteManyResponse, *resty.Response, error)

	// DeleteByID выполняет запрос на удаление серии по ID.
	// Принимает контекст и ID серии.
	// Возвращает «true» в случае успешного удаления серии.
	DeleteByID(ctx context.Context, id string) (bool, *resty.Response, error)

	// Delete выполняет запрос на удаление серии.
	// Принимает контекст и серию.
	// Возвращает «true» в случае успешного удаления серии.
	Delete(ctx context.Context, entity *Consignment) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельной серии по ID.
	// Принимает контекст, ID серии и опционально объект параметров запроса Params.
	// Возвращает найденную серию.
	GetByID(ctx context.Context, id string, params ...func(*Params)) (*Consignment, *resty.Response, error)

	// Update выполняет запрос на изменение серии.
	// Принимает контекст, серию и опционально объект параметров запроса Params.
	// Возвращает изменённую серию.
	Update(ctx context.Context, id string, consignment *Consignment, params ...func(*Params)) (*Consignment, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных серий.
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
	GetAttributeByID(ctx context.Context, id string) (*Attribute, *resty.Response, error)

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
	UpdateAttribute(ctx context.Context, id string, attr *Attribute) (*Attribute, *resty.Response, error)

	// DeleteAttribute выполняет запрос на удаление доп поля.
	// Принимает контекст и ID доп поля.
	// Возвращает «true» в случае успешного удаления доп поля.
	DeleteAttribute(ctx context.Context, id string) (bool, *resty.Response, error)

	// DeleteAttributeMany выполняет запрос на массовое удаление доп полей.
	// Принимает контекст и множество доп полей.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteAttributeMany(ctx context.Context, attributes ...*Attribute) (*DeleteManyResponse, *resty.Response, error)

	// GetNamedFilterList выполняет запрос на получение списка фильтров.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetNamedFilterList(ctx context.Context, params ...func(*Params)) (*List[NamedFilter], *resty.Response, error)

	// GetNamedFilterByID выполняет запрос на получение отдельного фильтра по ID.
	// Принимает контекст и ID фильтра.
	// Возвращает найденный фильтр.
	GetNamedFilterByID(ctx context.Context, id string) (*NamedFilter, *resty.Response, error)
}

const (
	EndpointConsignment = EndpointEntity + string(MetaTypeConsignment)
)

// NewConsignmentService принимает [Client] и возвращает сервис для работы с сериями.
func NewConsignmentService(client *Client) ConsignmentService {
	return newMainService[Consignment, any, MetaAttributesWrapper, any](client, EndpointConsignment)

}
