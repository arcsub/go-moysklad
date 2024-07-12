package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// SalesChannel Канал продаж.
//
// Код сущности: saleschannel
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kanal-prodazh
type SalesChannel struct {
	ID           *uuid.UUID       `json:"id,omitempty"`           // ID Канала продаж
	Archived     *bool            `json:"archived,omitempty"`     // Добавлен ли Канал продаж в архив
	Code         *string          `json:"code,omitempty"`         // Код Канала продаж
	Description  *string          `json:"description,omitempty"`  // Описание Канала продаж
	ExternalCode *string          `json:"externalCode,omitempty"` // Внешний код Канала продаж
	Group        *Group           `json:"group,omitempty"`        // Отдел сотрудника
	AccountID    *uuid.UUID       `json:"accountId,omitempty"`    // ID учётной записи
	Meta         *Meta            `json:"meta,omitempty"`         // Метаданные Канала продаж
	Name         *string          `json:"name,omitempty"`         // Наименование Канала продаж
	Owner        *Employee        `json:"owner,omitempty"`        // Метаданные владельца (Сотрудника)
	Shared       *bool            `json:"shared,omitempty"`       // Общий доступ
	Updated      *Timestamp       `json:"update,omitempty"`       // Момент последнего обновления Канала продаж
	Type         SalesChannelType `json:"type,omitempty"`         // Тип Канала продаж
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (salesChannel SalesChannel) Clean() *SalesChannel {
	if salesChannel.Meta == nil {
		return nil
	}
	return &SalesChannel{Meta: salesChannel.Meta}
}

// GetID возвращает ID Канала продаж.
func (salesChannel SalesChannel) GetID() uuid.UUID {
	return Deref(salesChannel.ID)
}

// GetArchived возвращает true, если Канал продаж находится в архиве.
func (salesChannel SalesChannel) GetArchived() bool {
	return Deref(salesChannel.Archived)
}

// GetCode возвращает Код Канала продаж.
func (salesChannel SalesChannel) GetCode() string {
	return Deref(salesChannel.Code)
}

// GetDescription возвращает Описание Канала продаж.
func (salesChannel SalesChannel) GetDescription() string {
	return Deref(salesChannel.Description)
}

// GetExternalCode возвращает Внешний код Канала продаж.
func (salesChannel SalesChannel) GetExternalCode() string {
	return Deref(salesChannel.ExternalCode)
}

// GetGroup возвращает Отдел сотрудника.
func (salesChannel SalesChannel) GetGroup() Group {
	return Deref(salesChannel.Group)
}

// GetAccountID возвращает ID учётной записи.
func (salesChannel SalesChannel) GetAccountID() uuid.UUID {
	return Deref(salesChannel.AccountID)
}

// GetMeta возвращает Метаданные Канала продаж.
func (salesChannel SalesChannel) GetMeta() Meta {
	return Deref(salesChannel.Meta)
}

// GetName возвращает Наименование Канала продаж.
func (salesChannel SalesChannel) GetName() string {
	return Deref(salesChannel.Name)
}

// GetOwner возвращает Метаданные владельца (Сотрудника).
func (salesChannel SalesChannel) GetOwner() Employee {
	return Deref(salesChannel.Owner)
}

// GetShared возвращает флаг Общего доступа.
func (salesChannel SalesChannel) GetShared() bool {
	return Deref(salesChannel.Shared)
}

// GetUpdated возвращает Момент последнего обновления Канала продаж.
func (salesChannel SalesChannel) GetUpdated() Timestamp {
	return Deref(salesChannel.Updated)
}

// GetType возвращает Тип Канала продаж.
func (salesChannel SalesChannel) GetType() SalesChannelType {
	return salesChannel.Type
}

// SetArchived устанавливает флаг нахождения Канала продаж в архиве.
func (salesChannel *SalesChannel) SetArchived(archived bool) *SalesChannel {
	salesChannel.Archived = &archived
	return salesChannel
}

// SetCode устанавливает Код Канала продаж.
func (salesChannel *SalesChannel) SetCode(code string) *SalesChannel {
	salesChannel.Code = &code
	return salesChannel
}

// SetDescription устанавливает Описание Канала продаж.
func (salesChannel *SalesChannel) SetDescription(description string) *SalesChannel {
	salesChannel.Description = &description
	return salesChannel
}

// SetExternalCode устанавливает Внешний код Канала продаж.
func (salesChannel *SalesChannel) SetExternalCode(externalCode string) *SalesChannel {
	salesChannel.ExternalCode = &externalCode
	return salesChannel
}

// SetGroup устанавливает Метаданные отдела сотрудника.
func (salesChannel *SalesChannel) SetGroup(group *Group) *SalesChannel {
	if group != nil {
		salesChannel.Group = group.Clean()
	}
	return salesChannel
}

// SetMeta устанавливает Метаданные Канала продаж.
func (salesChannel *SalesChannel) SetMeta(meta *Meta) *SalesChannel {
	salesChannel.Meta = meta
	return salesChannel
}

// SetName устанавливает Наименование Канала продаж.
func (salesChannel *SalesChannel) SetName(name string) *SalesChannel {
	salesChannel.Name = &name
	return salesChannel
}

// SetOwner устанавливает Метаданные владельца (Сотрудника).
func (salesChannel *SalesChannel) SetOwner(owner *Employee) *SalesChannel {
	if owner != nil {
		salesChannel.Owner = owner.Clean()
	}
	return salesChannel
}

// SetShared устанавливает флаг общего доступа.
func (salesChannel *SalesChannel) SetShared(shared bool) *SalesChannel {
	salesChannel.Shared = &shared
	return salesChannel
}

// SetType устанавливает Тип Канала продаж.
func (salesChannel *SalesChannel) SetType(salesChannelType SalesChannelType) *SalesChannel {
	salesChannel.Type = salesChannelType
	return salesChannel
}

// SetTypeMessenger устанавливает Тип Канала продаж в значение [SalesChannelTypeMessenger].
func (salesChannel *SalesChannel) SetTypeMessenger() *SalesChannel {
	salesChannel.Type = SalesChannelTypeMessenger
	return salesChannel
}

// SetTypeSocialNetwork устанавливает Тип Канала продаж в значение [SalesChannelTypeSocialNetwork].
func (salesChannel *SalesChannel) SetTypeSocialNetwork() *SalesChannel {
	salesChannel.Type = SalesChannelTypeSocialNetwork
	return salesChannel
}

// SetTypeMarketplace устанавливает Тип Канала продаж в значение [SalesChannelTypeMarketplace].
func (salesChannel *SalesChannel) SetTypeMarketplace() *SalesChannel {
	salesChannel.Type = SalesChannelTypeMarketplace
	return salesChannel
}

// SetTypeEcommerce устанавливает Тип Канала продаж в значение [SalesChannelTypeEcommerce].
func (salesChannel *SalesChannel) SetTypeEcommerce() *SalesChannel {
	salesChannel.Type = SalesChannelTypeEcommerce
	return salesChannel
}

// SetTypeClassifiedAds устанавливает Тип Канала продаж в значение [SalesChannelTypeClassifiedAds].
func (salesChannel *SalesChannel) SetTypeClassifiedAds() *SalesChannel {
	salesChannel.Type = SalesChannelTypeClassifiedAds
	return salesChannel
}

// SetTypeDirectSales устанавливает Тип Канала продаж в значение [SalesChannelTypeDirectSales].
func (salesChannel *SalesChannel) SetTypeDirectSales() *SalesChannel {
	salesChannel.Type = SalesChannelTypeDirectSales
	return salesChannel
}

// SetTypeOther устанавливает Тип Канала продаж в значение [SalesChannelTypeOther].
func (salesChannel *SalesChannel) SetTypeOther() *SalesChannel {
	salesChannel.Type = SalesChannelTypeOther
	return salesChannel
}

// String реализует интерфейс [fmt.Stringer].
func (salesChannel SalesChannel) String() string {
	return Stringify(salesChannel)
}

// MetaType возвращает код сущности.
func (SalesChannel) MetaType() MetaType {
	return MetaTypeSalesChannel
}

// Update shortcut
func (salesChannel SalesChannel) Update(ctx context.Context, client *Client, params ...*Params) (*SalesChannel, *resty.Response, error) {
	return NewSalesChannelService(client).Update(ctx, salesChannel.GetID(), &salesChannel, params...)
}

// Create shortcut
func (salesChannel SalesChannel) Create(ctx context.Context, client *Client, params ...*Params) (*SalesChannel, *resty.Response, error) {
	return NewSalesChannelService(client).Create(ctx, &salesChannel, params...)
}

// Delete shortcut
func (salesChannel SalesChannel) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return NewSalesChannelService(client).Delete(ctx, salesChannel.GetID())
}

// SalesChannelType Тип канала продаж.
//
// Возможные значения:
//   - SalesChannelTypeMessenger     – Мессенджер
//   - SalesChannelTypeSocialNetwork – Социальная сеть
//   - SalesChannelTypeMarketplace   – Маркетплейс
//   - SalesChannelTypeEcommerce     – Интернет-магазин
//   - SalesChannelTypeClassifiedAds – Доска объявлений
//   - SalesChannelTypeDirectSales   – Прямые продажи
//   - SalesChannelTypeOther         – Другое
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kanal-prodazh-kanaly-prodazh-tip-kanala-prodazh
type SalesChannelType string

const (
	SalesChannelTypeMessenger     SalesChannelType = "MESSENGER"      // Мессенджер
	SalesChannelTypeSocialNetwork SalesChannelType = "SOCIAL_NETWORK" // Социальная сеть
	SalesChannelTypeMarketplace   SalesChannelType = "MARKETPLACE"    // Маркетплейс
	SalesChannelTypeEcommerce     SalesChannelType = "ECOMMERCE"      // Интернет-магазин
	SalesChannelTypeClassifiedAds SalesChannelType = "CLASSIFIED_ADS" // Доска объявлений
	SalesChannelTypeDirectSales   SalesChannelType = "DIRECT_SALES"   // Прямые продажи
	SalesChannelTypeOther         SalesChannelType = "OTHER"          // Другое
)

// SalesChannelService описывает методы сервиса для работы с каналами продаж.
type SalesChannelService interface {
	// GetList выполняет запрос на получение списка каналов продаж.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[SalesChannel], *resty.Response, error)

	// Create выполняет запрос на создание канала продаж.
	// Обязательные поля для заполнения:
	//	- name (Наименование Канала продаж)
	//	- type (Тип Канала продаж)
	// Принимает контекст, канал продаж и опционально объект параметров запроса Params.
	// Возвращает созданный канал продаж.
	Create(ctx context.Context, salesChannel *SalesChannel, params ...*Params) (*SalesChannel, *resty.Response, error)

	// CreateUpdateMany выполняет запрос на массовое создание и/или изменение каналов продаж.
	// Изменяемые каналы продаж должны содержать идентификатор в виде метаданных.
	// Принимает контекст, список каналов продаж и опционально объект параметров запроса Params.
	// Возвращает список созданных и/или изменённых каналов продаж.
	CreateUpdateMany(ctx context.Context, salesChannelList Slice[SalesChannel], params ...*Params) (*Slice[SalesChannel], *resty.Response, error)

	// DeleteMany выполняет запрос на массовое удаление каналов продаж.
	// Принимает контекст и множество каналов продаж.
	// Возвращает объект DeleteManyResponse, содержащий информацию об успешном удалении или ошибку.
	DeleteMany(ctx context.Context, entities ...*SalesChannel) (*DeleteManyResponse, *resty.Response, error)

	// Delete выполняет запрос на удаление канала продаж.
	// Принимает контекст и ID канала продаж.
	// Возвращает «true» в случае успешного удаления канала продаж.
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)

	// GetByID выполняет запрос на получение отдельного канала продаж по ID.
	// Принимает контекст, ID канала продаж и опционально объект параметров запроса Params.
	// Возвращает канал продаж.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*SalesChannel, *resty.Response, error)

	// Update выполняет запрос на изменение канала продаж.
	// Принимает контекст, канал продаж и опционально объект параметров запроса Params.
	// Возвращает изменённый канал продаж.
	Update(ctx context.Context, id uuid.UUID, salesChannel *SalesChannel, params ...*Params) (*SalesChannel, *resty.Response, error)
}

const (
	EndpointSalesChannel = EndpointEntity + string(MetaTypeSalesChannel)
)

// NewSalesChannelService принимает [Client] и возвращает сервис для работы с каналами продаж.
func NewSalesChannelService(client *Client) SalesChannelService {
	return newMainService[SalesChannel, any, any, any](client, EndpointSalesChannel)
}
