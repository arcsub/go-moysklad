package moysklad

import (
	"github.com/google/uuid"
)

// SalesChannel Канал продаж.
// Ключевое слово: saleschannel
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kanal-prodazh
type SalesChannel struct {
	AccountID        *uuid.UUID       `json:"accountId,omitempty"`    // ID учетной записи
	Archived         *bool            `json:"archived,omitempty"`     // Добавлен ли Канал продаж в архив
	Code             *string          `json:"code,omitempty"`         // Код Канала продаж
	Description      *string          `json:"description,omitempty"`  // Описание Канала продаж
	ExternalCode     *string          `json:"externalCode,omitempty"` // Внешний код Канала продаж
	Group            *Group           `json:"group,omitempty"`        // Метаданные отдела сотрудника
	ID               *uuid.UUID       `json:"id,omitempty"`           // ID Канала продаж
	Meta             *Meta            `json:"meta,omitempty"`         // Метаданные
	Name             *string          `json:"name,omitempty"`         // Наименование Канала продаж
	Owner            *Employee        `json:"owner,omitempty"`        // Метаданные владельца (Сотрудника)
	Shared           *bool            `json:"shared,omitempty"`       // Общий доступ
	SalesChannelType SalesChannelType `json:"type,omitempty"`         // Тип Канала продаж
	Updated          *Timestamp       `json:"update,omitempty"`       // Момент последнего обновления сущности
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
