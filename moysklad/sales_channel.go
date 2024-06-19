package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// SalesChannel Канал продаж.
// Ключевое слово: saleschannel
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kanal-prodazh
type SalesChannel struct {
	ID               *uuid.UUID       `json:"id,omitempty"`
	Archived         *bool            `json:"archived,omitempty"`
	Code             *string          `json:"code,omitempty"`
	Description      *string          `json:"description,omitempty"`
	ExternalCode     *string          `json:"externalCode,omitempty"`
	Group            *Group           `json:"group,omitempty"`
	AccountID        *uuid.UUID       `json:"accountId,omitempty"`
	Meta             *Meta            `json:"meta,omitempty"`
	Name             *string          `json:"name,omitempty"`
	Owner            *Employee        `json:"owner,omitempty"`
	Shared           *bool            `json:"shared,omitempty"`
	Updated          *Timestamp       `json:"update,omitempty"`
	SalesChannelType SalesChannelType `json:"type,omitempty"`
}

// Clean возвращает сущность с единственным заполненным полем Meta
func (salesChannel SalesChannel) Clean() *SalesChannel {
	return &SalesChannel{Meta: salesChannel.Meta}
}

func (salesChannel SalesChannel) GetID() uuid.UUID {
	return Deref(salesChannel.ID)
}

func (salesChannel SalesChannel) GetArchived() bool {
	return Deref(salesChannel.Archived)
}

func (salesChannel SalesChannel) GetCode() string {
	return Deref(salesChannel.Code)
}

func (salesChannel SalesChannel) GetDescription() string {
	return Deref(salesChannel.Description)
}

func (salesChannel SalesChannel) GetExternalCode() string {
	return Deref(salesChannel.ExternalCode)
}

func (salesChannel SalesChannel) GetGroup() Group {
	return Deref(salesChannel.Group)
}

func (salesChannel SalesChannel) GetAccountID() uuid.UUID {
	return Deref(salesChannel.AccountID)
}

func (salesChannel SalesChannel) GetMeta() Meta {
	return Deref(salesChannel.Meta)
}

func (salesChannel SalesChannel) GetName() string {
	return Deref(salesChannel.Name)
}

func (salesChannel SalesChannel) GetOwner() Employee {
	return Deref(salesChannel.Owner)
}

func (salesChannel SalesChannel) GetShared() bool {
	return Deref(salesChannel.Shared)
}

func (salesChannel SalesChannel) GetUpdated() Timestamp {
	return Deref(salesChannel.Updated)
}

func (salesChannel SalesChannel) GetSalesChannelType() SalesChannelType {
	return salesChannel.SalesChannelType
}

func (salesChannel *SalesChannel) SetArchived(archived bool) *SalesChannel {
	salesChannel.Archived = &archived
	return salesChannel
}

func (salesChannel *SalesChannel) SetCode(code string) *SalesChannel {
	salesChannel.Code = &code
	return salesChannel
}

func (salesChannel *SalesChannel) SetDescription(description string) *SalesChannel {
	salesChannel.Description = &description
	return salesChannel
}

func (salesChannel *SalesChannel) SetExternalCode(externalCode string) *SalesChannel {
	salesChannel.ExternalCode = &externalCode
	return salesChannel
}

func (salesChannel *SalesChannel) SetGroup(group *Group) *SalesChannel {
	salesChannel.Group = group.Clean()
	return salesChannel
}

func (salesChannel *SalesChannel) SetMeta(meta *Meta) *SalesChannel {
	salesChannel.Meta = meta
	return salesChannel
}

func (salesChannel *SalesChannel) SetName(name string) *SalesChannel {
	salesChannel.Name = &name
	return salesChannel
}

func (salesChannel *SalesChannel) SetOwner(owner *Employee) *SalesChannel {
	salesChannel.Owner = owner.Clean()
	return salesChannel
}

func (salesChannel *SalesChannel) SetShared(shared bool) *SalesChannel {
	salesChannel.Shared = &shared
	return salesChannel
}

func (salesChannel *SalesChannel) SetSalesChannelType(salesChannelType SalesChannelType) *SalesChannel {
	salesChannel.SalesChannelType = salesChannelType
	return salesChannel
}

func (salesChannel SalesChannel) String() string {
	return Stringify(salesChannel)
}

// MetaType возвращает тип сущности.
func (SalesChannel) MetaType() MetaType {
	return MetaTypeSalesChannel
}

// Update shortcut
func (salesChannel SalesChannel) Update(ctx context.Context, client *Client, params ...*Params) (*SalesChannel, *resty.Response, error) {
	return client.Entity().SalesChannel().Update(ctx, salesChannel.GetID(), &salesChannel, params...)
}

// Create shortcut
func (salesChannel SalesChannel) Create(ctx context.Context, client *Client, params ...*Params) (*SalesChannel, *resty.Response, error) {
	return client.Entity().SalesChannel().Create(ctx, &salesChannel, params...)
}

// Delete shortcut
func (salesChannel SalesChannel) Delete(ctx context.Context, client *Client) (bool, *resty.Response, error) {
	return client.Entity().SalesChannel().Delete(ctx, salesChannel.GetID())
}

// SalesChannelType Тип канала продаж.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kanal-prodazh-kanaly-prodazh-tip-kanala-prodazh
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

// SalesChannelService
// Сервис для работы с каналами продаж.
type SalesChannelService interface {
	GetList(ctx context.Context, params ...*Params) (*List[SalesChannel], *resty.Response, error)
	Create(ctx context.Context, salesChannel *SalesChannel, params ...*Params) (*SalesChannel, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, salesChannelList Slice[SalesChannel], params ...*Params) (*Slice[SalesChannel], *resty.Response, error)
	DeleteMany(ctx context.Context, entities ...*SalesChannel) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*SalesChannel, *resty.Response, error)
	Update(ctx context.Context, id uuid.UUID, salesChannel *SalesChannel, params ...*Params) (*SalesChannel, *resty.Response, error)
}

func NewSalesChannelService(client *Client) SalesChannelService {
	e := NewEndpoint(client, "entity/saleschannel")
	return newMainService[SalesChannel, any, any, any](e)
}
