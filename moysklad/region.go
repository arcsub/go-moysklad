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

func (region Region) GetAccountID() uuid.UUID {
	return Deref(region.AccountID)
}

func (region Region) GetCode() string {
	return Deref(region.Code)
}

func (region Region) GetExternalCode() string {
	return Deref(region.ExternalCode)
}

func (region Region) GetID() uuid.UUID {
	return Deref(region.ID)
}

func (region Region) GetMeta() Meta {
	return Deref(region.Meta)
}

func (region Region) GetName() string {
	return Deref(region.Name)
}

func (region Region) GetUpdated() Timestamp {
	return Deref(region.Updated)
}

func (region Region) GetVersion() int {
	return Deref(region.Version)
}

func (region *Region) SetCode(code string) *Region {
	region.Code = &code
	return region
}

func (region *Region) SetExternalCode(externalCode string) *Region {
	region.ExternalCode = &externalCode
	return region
}

func (region *Region) SetMeta(meta *Meta) *Region {
	region.Meta = meta
	return region
}

func (region *Region) SetName(name string) *Region {
	region.Name = &name
	return region
}

func (region Region) String() string {
	return Stringify(region)
}

func (region Region) MetaType() MetaType {
	return MetaTypeRegion
}

// RegionService
// Сервис для работы с регионами.
type RegionService interface {
	GetList(ctx context.Context, params *Params) (*List[Region], *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params *Params) (*Region, *resty.Response, error)
}

func NewRegionService(client *Client) RegionService {
	e := NewEndpoint(client, "entity/region")
	return newMainService[Region, any, any, any](e)
}
