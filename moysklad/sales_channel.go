package moysklad

import (
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
