package moysklad

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

// Thing Серийный номер
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-serijnyj-nomer
type Thing struct {
	AccountID   *uuid.UUID `json:"accountId,omitempty"`   // ID учетной записи
	Description *string    `json:"description,omitempty"` // Описание Серийного номера
	ID          *uuid.UUID `json:"id,omitempty"`          // ID Серийного номера
	Meta        *Meta      `json:"meta,omitempty"`        // Метаданные о Серийном номере
	Name        *string    `json:"name,omitempty"`        // Наименование Серийного номера
}

func (thing Thing) GetAccountID() uuid.UUID {
	return Deref(thing.AccountID)
}

func (thing Thing) GetDescription() string {
	return Deref(thing.Description)
}

func (thing Thing) GetID() uuid.UUID {
	return Deref(thing.ID)
}

func (thing Thing) GetMeta() Meta {
	return Deref(thing.Meta)
}

func (thing Thing) GetName() string {
	return Deref(thing.Name)
}

func (thing *Thing) SetDescription(description string) *Thing {
	thing.Description = &description
	return thing
}

func (thing *Thing) SetMeta(meta *Meta) *Thing {
	thing.Meta = meta
	return thing
}

func (thing *Thing) SetName(name string) *Thing {
	thing.Name = &name
	return thing
}

func (thing Thing) String() string {
	return Stringify(thing)
}

// MetaType возвращает тип сущности.
func (Thing) MetaType() MetaType {
	return MetaTypeThing
}

// ThingService
// Сервис для работы с серийными номерами.
type ThingService interface {
	GetList(ctx context.Context, params ...*Params) (*List[Thing], *resty.Response, error)
	GetByID(ctx context.Context, id uuid.UUID, params ...*Params) (*Thing, *resty.Response, error)
}

func NewThingService(client *Client) ThingService {
	e := NewEndpoint(client, "entity/thing")
	return newMainService[Thing, any, any, any](e)
}
