package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"time"
)

// Region Регион.
//
// Код сущности: region
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-region
type Region struct {
	AccountID    *uuid.UUID `json:"accountId,omitempty"`    // ID учётной записи
	Code         *string    `json:"code,omitempty"`         // Код Региона
	ExternalCode *string    `json:"externalCode,omitempty"` // Внешний код Региона
	ID           *uuid.UUID `json:"id,omitempty"`           // ID Региона
	Meta         *Meta      `json:"meta,omitempty"`         // Метаданные Региона
	Name         *string    `json:"name,omitempty"`         // Наименование Региона
	Updated      *Timestamp `json:"updated,omitempty"`      // Момент последнего обновления Региона
	Version      *int       `json:"version,omitempty"`      // Версия сущности
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (region Region) Clean() *Region {
	if region.Meta == nil {
		return nil
	}
	return &Region{Meta: region.Meta}
}

// GetAccountID возвращает ID учётной записи.
func (region Region) GetAccountID() uuid.UUID {
	return Deref(region.AccountID)
}

// GetCode возвращает Код Региона.
func (region Region) GetCode() string {
	return Deref(region.Code)
}

// GetExternalCode возвращает Внешний код Региона.
func (region Region) GetExternalCode() string {
	return Deref(region.ExternalCode)
}

// GetID возвращает ID Региона.
func (region Region) GetID() uuid.UUID {
	return Deref(region.ID)
}

// GetMeta возвращает Метаданные Региона.
func (region Region) GetMeta() Meta {
	return Deref(region.Meta)
}

// GetName возвращает Наименование Региона.
func (region Region) GetName() string {
	return Deref(region.Name)
}

// GetUpdated возвращает Момент последнего обновления Региона.
func (region Region) GetUpdated() time.Time {
	return Deref(region.Updated).Time()
}

// GetVersion возвращает Версию сущности.
func (region Region) GetVersion() int {
	return Deref(region.Version)
}

// SetCode устанавливает Код Региона.
func (region *Region) SetCode(code string) *Region {
	region.Code = &code
	return region
}

// SetExternalCode устанавливает Внешний код Региона.
func (region *Region) SetExternalCode(externalCode string) *Region {
	region.ExternalCode = &externalCode
	return region
}

// SetMeta устанавливает Метаданные Региона.
func (region *Region) SetMeta(meta *Meta) *Region {
	region.Meta = meta
	return region
}

// SetName устанавливает Наименование Региона.
func (region *Region) SetName(name string) *Region {
	region.Name = &name
	return region
}

// String реализует интерфейс [fmt.Stringer].
func (region Region) String() string {
	return Stringify(region)
}

// MetaType возвращает код сущности.
func (Region) MetaType() MetaType {
	return MetaTypeRegion
}

// RegionService описывает методы сервиса для работы с регионами.
type RegionService interface {
	// GetList выполняет запрос на получение списка регионов.
	// Принимает контекст и опционально объект параметров запроса Params.
	// Возвращает объект List.
	GetList(ctx context.Context, params ...*Params) (*List[Region], *resty.Response, error)

	// GetByID выполняет запрос на получение отдельного региона по ID.
	// Принимает контекст, ID региона и опционально объект параметров запроса Params.
	// Возвращает найденный регион.
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Region, *resty.Response, error)
}

const (
	EndpointRegion = EndpointEntity + string(MetaTypeRegion)
)

// NewRegionService принимает [Client] и возвращает сервис для работы с регионами.
func NewRegionService(client *Client) RegionService {
	return newMainService[Region, any, any, any](client, EndpointRegion)
}
