package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Region Регион.
// Ключевое слово: region
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-region
type Region struct {
	AccountID    *uuid.UUID `json:"accountId,omitempty"`    // ID учетной записи
	Code         *string    `json:"code,omitempty"`         // Код Региона
	ExternalCode *string    `json:"externalCode,omitempty"` // Внешний код Региона
	ID           *uuid.UUID `json:"id,omitempty"`           // ID Региона
	Meta         *Meta      `json:"meta,omitempty"`         // Метаданные
	Name         *string    `json:"name,omitempty"`         // Наименование
	Updated      *Timestamp `json:"updated,omitempty"`      // Момент последнего обновления сущности
	Version      *int       `json:"version,omitempty"`      // Версия сущности
}

func (r Region) String() string {
	return Stringify(r)
}

func (r Region) MetaType() MetaType {
	return MetaTypeRegion
}

// RegionService
// Сервис для работы с регионами.
type RegionService interface {
	GetList(ctx context.Context, params *Params) (*List[Region], *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Region, *resty.Response, error)
}

func NewRegionService(client *Client) RegionService {
	e := NewEndpoint(client, "entity/region")
	return newMainService[Region, any, any, any](e)
}
