package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Thing Серийный номер
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-serijnyj-nomer
type Thing struct {
	AccountId   *uuid.UUID `json:"accountId,omitempty"`   // ID учетной записи
	Description *string    `json:"description,omitempty"` // Описание Серийного номера
	ID          *uuid.UUID `json:"id,omitempty"`          // ID Серийного номера
	Meta        *Meta      `json:"meta,omitempty"`        // Метаданные о Серийном номере
	Name        *string    `json:"name,omitempty"`        // Наименование Серийного номера
}

func (t Thing) String() string {
	return Stringify(t)
}

func (t Thing) MetaType() MetaType {
	return MetaTypeThing
}

// Things серийные номера.
type Things = []string

// ThingService
// Сервис для работы с серийными номерами.
type ThingService interface {
	GetList(ctx context.Context, params *Params) (*List[Thing], *resty.Response, error)
	GetByID(ctx context.Context, id *uuid.UUID, params *Params) (*Thing, *resty.Response, error)
}

func NewThingService(client *Client) ThingService {
	e := NewEndpoint(client, "entity/thing")
	return newMainService[Thing, any, any, any](e)
}
