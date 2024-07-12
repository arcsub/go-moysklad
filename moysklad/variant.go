package moysklad

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"time"
)

// Variant Модификация.
//
// Код сущности: variant
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-modifikaciq
type Variant struct {
	Archived           *bool                 `json:"archived,omitempty"`           // Добавлен ли товар в архив
	Updated            *Timestamp            `json:"updated,omitempty"`            // Момент последнего обновления Модификации
	AccountID          *uuid.UUID            `json:"accountId,omitempty"`          // ID учётной записи
	Description        *string               `json:"description,omitempty"`        // Описание Модификации
	ExternalCode       *string               `json:"externalCode,omitempty"`       // Внешний код Модификации
	ID                 *uuid.UUID            `json:"id,omitempty"`                 // ID Модификации
	Meta               *Meta                 `json:"meta,omitempty"`               // Метаданные Модификации
	Name               *string               `json:"name,omitempty"`               // Наименование товара с Модификацией
	Code               *string               `json:"code,omitempty"`               // Код Модификации
	Barcodes           Slice[Barcode]        `json:"barcodes,omitempty"`           // Штрихкоды
	DiscountProhibited *bool                 `json:"discountProhibited,omitempty"` // Признак запрета скидок
	Characteristics    Slice[Characteristic] `json:"characteristics,omitempty"`    // Характеристики Модификации
	Images             *MetaArray[Image]     `json:"images,omitempty"`             // Массив метаданных Изображений (Максимальное количество изображений - 10)
	MinPrice           *NullValue[MinPrice]  `json:"minPrice,omitempty"`           // Минимальная цена
	BuyPrice           *BuyPrice             `json:"buyPrice,omitempty"`           // Закупочная цена
	Product            *Product              `json:"product,omitempty"`            // Метаданные товара, к которому привязана Модификация
	SalePrices         Slice[SalePrice]      `json:"salePrices,omitempty"`         // Цены продажи
	Things             Slice[string]         `json:"things,omitempty"`             // Серийные номера
	Packs              Slice[VariantPack]    `json:"packs,omitempty"`              // Упаковки модификации
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (variant Variant) Clean() *Variant {
	if variant.Meta == nil {
		return nil
	}
	return &Variant{Meta: variant.Meta}
}

// NewVariantFromAssortment пытается привести переданный в качестве аргумента [AssortmentPosition] к типу [Variant].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [Variant] или nil в случае неудачи.
func NewVariantFromAssortment(assortmentPosition *AssortmentPosition) *Variant {
	return UnmarshalAsType[Variant](assortmentPosition)
}

// FromAssortment пытается привести переданный в качестве аргумента [AssortmentPosition] к типу [Variant].
//
// Метод гарантирует преобразование в необходимый тип только при идентичных [MetaType].
//
// Возвращает [Variant] или nil в случае неудачи.
func (variant Variant) FromAssortment(assortmentPosition *AssortmentPosition) *Variant {
	return UnmarshalAsType[Variant](assortmentPosition)
}

// AsAssortment реализует интерфейс [AssortmentConverter].
func (variant Variant) AsAssortment() *AssortmentPosition {
	return &AssortmentPosition{Meta: variant.GetMeta()}
}

// GetArchived возвращает флаг нахождения в архиве.
func (variant Variant) GetArchived() bool {
	return Deref(variant.Archived)
}

// GetUpdated возвращает Момент последнего обновления Модификации.
func (variant Variant) GetUpdated() time.Time {
	return Deref(variant.Updated).Time()
}

// GetAccountID возвращает ID учётной записи.
func (variant Variant) GetAccountID() uuid.UUID {
	return Deref(variant.AccountID)
}

// GetDescription возвращает Описание Модификации.
func (variant Variant) GetDescription() string {
	return Deref(variant.Description)
}

// GetExternalCode возвращает Внешний код Модификации.
func (variant Variant) GetExternalCode() string {
	return Deref(variant.ExternalCode)
}

// GetID возвращает ID Модификации.
func (variant Variant) GetID() uuid.UUID {
	return Deref(variant.ID)
}

// GetMeta возвращает Метаданные Модификации.
func (variant Variant) GetMeta() Meta {
	return Deref(variant.Meta)
}

// GetName возвращает Наименование товара с Модификацией.
func (variant Variant) GetName() string {
	return Deref(variant.Name)
}

// GetCode возвращает Код Модификации.
func (variant Variant) GetCode() string {
	return Deref(variant.Code)
}

// GetBarcodes возвращает Штрихкоды.
func (variant Variant) GetBarcodes() Slice[Barcode] {
	return variant.Barcodes
}

// GetDiscountProhibited возвращает Признак запрета скидок.
func (variant Variant) GetDiscountProhibited() bool {
	return Deref(variant.DiscountProhibited)
}

// GetCharacteristics возвращает Характеристики Модификации.
func (variant Variant) GetCharacteristics() Slice[Characteristic] {
	return variant.Characteristics
}

// GetImages возвращает Массив метаданных Изображений.
func (variant Variant) GetImages() MetaArray[Image] {
	return Deref(variant.Images)
}

// GetMinPrice возвращает Минимальную цену.
func (variant Variant) GetMinPrice() MinPrice {
	return Deref(variant.MinPrice).getValue()
}

// GetBuyPrice возвращает Закупочную цену.
func (variant Variant) GetBuyPrice() BuyPrice {
	return Deref(variant.BuyPrice)
}

// GetProduct возвращает Метаданные товара, к которому привязана Модификация.
func (variant Variant) GetProduct() Product {
	return Deref(variant.Product)
}

// GetSalePrices возвращает Цены продажи.
func (variant Variant) GetSalePrices() Slice[SalePrice] {
	return variant.SalePrices
}

// GetThings возвращает Серийные номера.
func (variant Variant) GetThings() Slice[string] {
	return variant.Things
}

// GetPacks возвращает Упаковки модификации.
func (variant Variant) GetPacks() Slice[VariantPack] {
	return variant.Packs
}

// SetArchived устанавливает флаг нахождения в архиве.
func (variant *Variant) SetArchived(archived bool) *Variant {
	variant.Archived = &archived
	return variant
}

// SetDescription устанавливает Описание Модификации.
func (variant *Variant) SetDescription(description string) *Variant {
	variant.Description = &description
	return variant
}

// SetExternalCode устанавливает Внешний код Модификации.
func (variant *Variant) SetExternalCode(externalCode string) *Variant {
	variant.ExternalCode = &externalCode
	return variant
}

// SetMeta устанавливает Метаданные Модификации.
func (variant *Variant) SetMeta(meta *Meta) *Variant {
	variant.Meta = meta
	return variant
}

// SetCode устанавливает Код Модификации.
func (variant *Variant) SetCode(code string) *Variant {
	variant.Code = &code
	return variant
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
func (variant *Variant) SetBarcodes(barcodes ...*Barcode) *Variant {
	variant.Barcodes = barcodes
	return variant
}

// SetDiscountProhibited устанавливает Признак запрета скидок.
func (variant *Variant) SetDiscountProhibited(discountProhibited bool) *Variant {
	variant.DiscountProhibited = &discountProhibited
	return variant
}

// SetCharacteristics устанавливает Характеристики Модификации.
//
// Принимает множество объектов [Characteristic].
func (variant *Variant) SetCharacteristics(characteristics ...*Characteristic) *Variant {
	variant.Characteristics.Push(characteristics...)
	return variant
}

// SetImages устанавливает Массив метаданных Изображений.
//
// Принимает множество объектов [Image].
func (variant *Variant) SetImages(images ...*Image) *Variant {
	variant.Images = NewMetaArrayFrom(images)
	return variant
}

// SetMinPrice устанавливает Минимальную цену.
//
// Передача nil передаёт сброс значения (null).
func (variant *Variant) SetMinPrice(minPrice *MinPrice) *Variant {
	variant.MinPrice = NewNullValue(minPrice)
	return variant
}

// SetBuyPrice устанавливает Закупочную цену.
func (variant *Variant) SetBuyPrice(buyPrice *BuyPrice) *Variant {
	if buyPrice != nil {
		variant.BuyPrice = buyPrice
	}
	return variant
}

// SetProduct устанавливает Метаданные товара, к которому привязана Модификация.
func (variant *Variant) SetProduct(product *Product) *Variant {
	if product != nil {
		variant.Product = product.Clean()
	}
	return variant
}

// SetSalePrices устанавливает Цены продажи.
//
// Принимает множество объектов [SalePrice].
func (variant *Variant) SetSalePrices(salePrices ...*SalePrice) *Variant {
	variant.SalePrices.Push(salePrices...)
	return variant
}

// SetPacks устанавливает Упаковки Товара.
//
// Принимает множество объектов [VariantPack].
func (variant *Variant) SetPacks(packs ...*VariantPack) *Variant {
	variant.Packs.Push(packs...)
	return variant
}

// String реализует интерфейс [fmt.Stringer].
func (variant Variant) String() string {
	return Stringify(variant)
}

// MetaType возвращает код сущности.
func (Variant) MetaType() MetaType {
	return MetaTypeVariant
}

// Update shortcut
func (variant Variant) Update(ctx context.Context, client *Client, params ...*Params) (*Variant, *resty.Response, error) {
	return NewVariantService(client).Update(ctx, variant.GetID(), &variant, params...)
}

// Create shortcut
func (variant Variant) Create(ctx context.Context, client *Client, params ...*Params) (*Variant, *resty.Response, error) {
	return NewVariantService(client).Create(ctx, &variant, params...)
}

// Delete shortcut
func (variant Variant) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewVariantService(client).Delete(ctx, variant.GetID())
}

// VariantPack Упаковка модификации.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-modifikaciq-modifikacii-atributy-wlozhennyh-suschnostej-upakowki-modifikacii
type VariantPack struct {
	ID         *uuid.UUID     `json:"id,omitempty"`         // ID упаковки модификации
	ParentPack *Pack          `json:"parentpack,omitempty"` // Метаданные родительской упаковки (упаковки товара), для которой переопределяется штрихкод
	Barcodes   Slice[Barcode] `json:"barcodes,omitempty"`   // Массив штрихкодов упаковки модификации. Данный массив может содержать только один штрихкод
}

// GetID возвращает ID упаковки модификации.
func (variantPack VariantPack) GetID() uuid.UUID {
	return Deref(variantPack.ID)
}

// GetParentPack возвращает Метаданные родительской упаковки (упаковки товара), для которой переопределяется штрихкод.
func (variantPack VariantPack) GetParentPack() Pack {
	return Deref(variantPack.ParentPack)
}

// GetBarcodes возвращает Массив штрихкодов упаковки модификации. Данный массив может содержать только один штрихкод.
func (variantPack VariantPack) GetBarcodes() Slice[Barcode] {
	return variantPack.Barcodes
}

// SetParentPack устанавливает Метаданные родительской упаковки (упаковки товара), для которой переопределяется штрихкод.
func (variantPack *VariantPack) SetParentPack(pack *Pack) *VariantPack {
	if pack != nil {
		variantPack.ParentPack = pack
	}
	return variantPack
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
func (variantPack *VariantPack) SetBarcodes(barcodes ...*Barcode) *VariantPack {
	variantPack.Barcodes.Push(barcodes...)
	return variantPack
}

// String реализует интерфейс [fmt.Stringer].
func (variantPack VariantPack) String() string {
	return Stringify(variantPack)
}

// Characteristic Характеристика
//
// Код сущности: attributemetadata
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-modifikaciq-modifikacii-atributy-wlozhennyh-suschnostej-metadannye-modifikacij-harakteristiki-modifikacii
type Characteristic struct {
	ID       *uuid.UUID `json:"id,omitempty"`       // ID соответствующей характеристики
	Meta     *Meta      `json:"meta,omitempty"`     // Метаданные характеристики
	Name     *string    `json:"name,omitempty"`     // Наименование характеристики
	Required *bool      `json:"required,omitempty"` // Флаг о том, является ли характеристика обязательной
	Type     *string    `json:"type,omitempty"`     // Тип значения характеристики (значение всегда "string")
	Value    *string    `json:"value,omitempty"`    // Значение характеристики
}

// GetID возвращает ID соответствующей характеристики.
func (characteristic Characteristic) GetID() uuid.UUID {
	return Deref(characteristic.ID)
}

// GetMeta возвращает Метаданные характеристики.
func (characteristic Characteristic) GetMeta() Meta {
	return Deref(characteristic.Meta)
}

// GetName возвращает Наименование характеристики.
func (characteristic Characteristic) GetName() string {
	return Deref(characteristic.Name)
}

// GetRequired возвращает флаг о том, является ли характеристика обязательной.
func (characteristic Characteristic) GetRequired() bool {
	return Deref(characteristic.Required)
}

// GetType возвращает Тип значения характеристики (значение всегда "string").
func (characteristic Characteristic) GetType() string {
	return Deref(characteristic.Type)
}

// GetValue возвращает Значение характеристики.
func (characteristic Characteristic) GetValue() string {
	return Deref(characteristic.Value)
}

// SetMeta устанавливает Метаданные характеристики.
func (characteristic *Characteristic) SetMeta(meta *Meta) *Characteristic {
	characteristic.Meta = meta
	return characteristic
}

// SetName устанавливает Наименование характеристики.
func (characteristic *Characteristic) SetName(name string) *Characteristic {
	characteristic.Name = &name
	return characteristic
}

// SetRequired устанавливает флаг о том, является ли характеристика обязательной.
func (characteristic *Characteristic) SetRequired(required bool) *Characteristic {
	characteristic.Required = &required
	return characteristic
}

// SetValue устанавливает Значение характеристики.
func (characteristic *Characteristic) SetValue(value string) *Characteristic {
	characteristic.Value = &value
	return characteristic
}

// String реализует интерфейс [fmt.Stringer].
func (characteristic Characteristic) String() string {
	return Stringify(characteristic)
}

// MetaType возвращает код сущности.
func (Characteristic) MetaType() MetaType {
	return MetaTypeCharacteristic
}

// VariantService
// Сервис для работы с модификациями.
type VariantService interface {
	// GetList выполняет запрос на получение списка модификаций.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[Variant], *resty.Response, error)

	// Create выполняет запрос на создание заказа модификации.
	// Обязательные поля для заполнения:
	//	- product (Метаданные товара, к которому привязана Модификация)
	//	- characteristics (Характеристики Модификации)
	// Принимает контекст, модификацию и опционально объект параметров запроса Params.
	// Возвращает созданную модификацию.
	Create(ctx context.Context, variant *Variant, params ...*Params) (*Variant, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение модификаций.
	// Изменяемые модификации должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список модификаций и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых модификаций.
	CreateUpdateMany(ctx context.Context, variantList Slice[Variant], params ...*Params) (*Slice[Variant], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление модификаций.
	// Принимает контекст и множество модификаций.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*Variant) (*DeleteManyResponse, *resty.Response, error)

	// Delete выполняет запрос на удаление модификации.
	// Принимает контекст и ID модификации.
	// Возвращает «true» в случае успешного удаления модификации.
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельной модификации по ID.
	// Принимает контекст, ID модификации и опционально объект параметров запроса Params.
	// Возвращает модификацию.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Variant, *resty.Response, error)

	// Update выполняет запрос на изменение модификации.
	// Принимает контекст, модификацию и опционально объект параметров запроса Params.
	// Возвращает изменённую модификацию.
	Update(ctx context.Context, id uuid.UUID, variant *Variant, params ...*Params) (*Variant, *resty.Response, error)

	// GetMetadata выполняет запрос на получение метаданных модификаций.
	// Принимает контекст.
	// Возвращает объект метаданных MetaCharacteristicsWrapper.
	GetMetadata(ctx context.Context) (*MetaCharacteristicsWrapper, *resty.Response, error)

	// GetImageList выполняет запрос на получение изображений модификации в виде списка.
	// Принимает контекст и ID модификации.
	// Возвращает объект List.
	GetImageList(ctx context.Context, id uuid.UUID) (*List[Image], *resty.Response, error)

	// CreateImage выполняет запрос на добавление изображения.
	// Принимает контекст, ID модификации и изображение.
	// Возвращает список изображений.
	CreateImage(ctx context.Context, id uuid.UUID, image *Image) (*Slice[Image], *resty.Response, error)

	// UpdateImageMany выполняет запрос на обновления изображений.
	// Принимает контекст, ID модификации и изображение.
	// Если необходимо оставить некоторые Изображения, то необходимо передать эти изображения.
	// Возвращает список изображений.
	UpdateImageMany(ctx context.Context, id uuid.UUID, images ...*Image) (*Slice[Image], *resty.Response, error)

	// DeleteImage выполняет запрос на удаление изображения модификации.
	// Принимает контекст, ID модификации и ID изображения.
	// Возвращает «true» в случае успешного удаления изображения модификации.
	DeleteImage(ctx context.Context, id uuid.UUID, imageID uuid.UUID) (bool, *resty.Response, error)

	// DeleteImageMany выполняет запрос на массовое удаление изображений модификации.
	// Принимает контекст, ID модификации и множество файлов.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteImageMany(ctx context.Context, id uuid.UUID, images ...*Image) (*DeleteManyResponse, *resty.Response, error)

	// GetNamedFilterList выполняет запрос на получение списка фильтров.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetNamedFilterList(ctx context.Context, params ...*Params) (*List[NamedFilter], *resty.Response, error)

	// GetNamedFilterByID выполняет запрос на получение отдельного фильтра по ID.
	// Принимает контекст и ID фильтра.
	// Возвращает найденный фильтр.
	GetNamedFilterByID(ctx context.Context, id uuid.UUID) (*NamedFilter, *resty.Response, error)

	// CreateCharacteristic выполняет запрос на создание характеристики.
	// Принимает контекст и характеристику.
	// Возвращает созданную характеристику.
	CreateCharacteristic(ctx context.Context, characteristic *Characteristic) (*Characteristic, *resty.Response, error)

	// CreateCharacteristicMany выполняет запрос на массовое создание и/или изменение характеристик.
	// Изменяемые характеристики должны содержать идентификатор в виде метаданных.
	// Принимает контекст и список характеристик.
	// Возвращает список созданных и/или изменённых характеристик.
	CreateCharacteristicMany(ctx context.Context, characteristics ...*Characteristic) (*Slice[Characteristic], *resty.Response, error)

	// GetCharacteristicByID выполняет запрос на получение отдельной характеристики по ID.
	// Принимает контекст, ID характеристики.
	// Возвращает характеристику.
	GetCharacteristicByID(ctx context.Context, id uuid.UUID) (*Characteristic, *resty.Response, error)

	// UpdateCharacteristic выполняет запрос на изменение характеристики.
	// Принимает контекст и характеристику.
	// Возвращает изменённую характеристику.
	UpdateCharacteristic(ctx context.Context, id uuid.UUID, characteristic *Characteristic) (*Characteristic, *resty.Response, error)

	// DeleteCharacteristic выполняет запрос на удаление характеристики.
	// Принимает контекст и ID характеристики.
	// Возвращает «true» в случае успешного удаления характеристики.
	DeleteCharacteristic(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
}

const (
	EndpointVariant                  = EndpointEntity + string(MetaTypeVariant)
	EndpointVariantCharacteristics   = EndpointVariant + "/metadata/characteristics"
	EndpointVariantCharacteristicsID = EndpointVariantCharacteristics + "/%s"
)

type variantService struct {
	Endpoint
	endpointGetList[Variant]
	endpointCreate[Variant]
	endpointCreateUpdateMany[Variant]
	endpointDeleteMany[Variant]
	endpointDelete
	endpointGetByID[Variant]
	endpointUpdate[Variant]
	endpointMetadata[MetaCharacteristicsWrapper]
	endpointImages
	endpointNamedFilter
}

func (service *variantService) CreateCharacteristic(ctx context.Context, characteristic *Characteristic) (*Characteristic, *resty.Response, error) {
	return NewRequestBuilder[Characteristic](service.client, EndpointVariantCharacteristics).Post(ctx, characteristic)
}

func (service *variantService) CreateCharacteristicMany(ctx context.Context, characteristics ...*Characteristic) (*Slice[Characteristic], *resty.Response, error) {
	return NewRequestBuilder[Slice[Characteristic]](service.client, EndpointVariantCharacteristics).Post(ctx, characteristics)
}

func (service *variantService) GetCharacteristicByID(ctx context.Context, id uuid.UUID) (*Characteristic, *resty.Response, error) {
	path := fmt.Sprintf(EndpointVariantCharacteristicsID, id)
	return NewRequestBuilder[Characteristic](service.client, path).Get(ctx)
}

func (service *variantService) UpdateCharacteristic(ctx context.Context, id uuid.UUID, characteristic *Characteristic) (*Characteristic, *resty.Response, error) {
	path := fmt.Sprintf(EndpointVariantCharacteristicsID, id)
	return NewRequestBuilder[Characteristic](service.client, path).Put(ctx, characteristic)
}

func (service *variantService) DeleteCharacteristic(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error) {
	path := fmt.Sprintf(EndpointVariantCharacteristicsID, id)
	return NewRequestBuilder[any](service.client, path).Delete(ctx)
}

// NewVariantService принимает [Client] и возвращает сервис для работы с модификациями.
func NewVariantService(client *Client) VariantService {
	e := NewEndpoint(client, EndpointVariant)
	return &variantService{
		Endpoint:                 e,
		endpointGetList:          endpointGetList[Variant]{e},
		endpointCreate:           endpointCreate[Variant]{e},
		endpointCreateUpdateMany: endpointCreateUpdateMany[Variant]{e},
		endpointDeleteMany:       endpointDeleteMany[Variant]{e},
		endpointDelete:           endpointDelete{e},
		endpointGetByID:          endpointGetByID[Variant]{e},
		endpointUpdate:           endpointUpdate[Variant]{e},
		endpointMetadata:         endpointMetadata[MetaCharacteristicsWrapper]{e},
		endpointImages:           endpointImages{e},
		endpointNamedFilter:      endpointNamedFilter{e},
	}
}
