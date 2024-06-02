package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Uom Единица измерения.
// Ключевое слово: uom
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-edinica-izmereniq
type Uom struct {
	AccountID    *uuid.UUID `json:"accountId,omitempty"`    // ID учетной записи
	Code         *string    `json:"code,omitempty"`         // Код Единицы измерения
	Description  *string    `json:"description,omitempty"`  // Описание Единциы измерения
	ExternalCode *string    `json:"externalCode,omitempty"` // Внешний код Единицы измерения
	Group        *Group     `json:"group,omitempty"`        // Отдел сотрудника
	ID           *uuid.UUID `json:"id,omitempty"`           // ID сущности
	Meta         *Meta      `json:"meta,omitempty"`         // Метаданные
	Name         *string    `json:"name,omitempty"`         // Наименование Единицы измерения
	Owner        *Employee  `json:"owner,omitempty"`        // Владелец (Сотрудник)
	Shared       *bool      `json:"shared,omitempty"`       // Общий доступ
	Updated      *Timestamp `json:"updated,omitempty"`      // Момент последнего обновления Единицы измерения
}

func (u Uom) String() string {
	return Stringify(u)
}

func (u Uom) MetaType() MetaType {
	return MetaTypeUom
}

// UomService
// Сервис для работы с единицами измерения.
type UomService interface {
	GetList(ctx context.Context, params *Params) (*List[Uom], *resty.Response, error)
	Create(ctx context.Context, uom *Uom, params *Params) (*Uom, *resty.Response, error)
	CreateUpdateMany(ctx context.Context, uomList []*Uom, params *Params) (*[]Uom, *resty.Response, error)
	DeleteMany(ctx context.Context, uomList *DeleteManyRequest) (*DeleteManyResponse, *resty.Response, error)
	Delete(ctx context.Context, id *uuid.UUID) (bool, *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Uom, *resty.Response, error)
	Update(ctx context.Context, id *uuid.UUID, uom *Uom, params *Params) (*Uom, *resty.Response, error)
}

func NewUomService(client *Client) UomService {
	e := NewEndpoint(client, "entity/uom")
	return newMainService[Uom, any, any, any](e)
}
