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

func (s SalesChannel) String() string {
	return Stringify(s)
}

func (s SalesChannel) MetaType() MetaType {
	return MetaTypeSalesChannel
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
	GetList(ctx context.Context, params *Params) (*List[SalesChannel], *resty.Response, error)
	Create(ctx context.Context, salesChannel *SalesChannel, params *Params) (*SalesChannel, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, salesChannelList []*SalesChannel, params *Params) (*[]SalesChannel, *resty.Response, error)
	DeleteMany(ctx context.Context, salesChannelList []*SalesChannel) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*SalesChannel, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, salesChannel *SalesChannel, params *Params) (*SalesChannel, *resty.Response, error)
}

func NewSalesChannelService(client *Client) SalesChannelService {
	e := NewEndpoint(client, "entity/saleschannel")
	return newMainService[SalesChannel, any, any, any](e)
}
