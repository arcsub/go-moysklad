package moysklad

import (
	"github.com/google/uuid"
)

// State Статус документов.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-statusy-dokumentow
type State struct {
	AccountID  *uuid.UUID `json:"accountId,omitempty"`  // ID учетной записи
	Color      *uint64    `json:"color,omitempty"`      // Цвет Статуса
	EntityType MetaType   `json:"entityType,omitempty"` // Тип сущности, к которой относится Статус (ключевое слово в рамках JSON API)
	ID         *uuid.UUID `json:"id,omitempty"`         // ID сущности
	Meta       *Meta      `json:"meta,omitempty"`       // Метаданные
	Name       *string    `json:"name,omitempty"`       // Наименование Статуса
	StateType  StateType  `json:"stateType,omitempty"`  // Тип Статуса
}

func (state State) GetAccountID() uuid.UUID {
	return Deref(state.AccountID)
}

func (state State) GetColor() uint64 {
	return Deref(state.Color)
}

func (state State) GetEntityType() MetaType {
	return state.EntityType
}

func (state State) GetID() uuid.UUID {
	return Deref(state.ID)
}

func (state State) GetMeta() Meta {
	return Deref(state.Meta)
}

func (state State) GetName() string {
	return Deref(state.Name)
}

func (state State) GetStateType() StateType {
	return state.StateType
}

func (state *State) SetColor(color uint64) *State {
	state.Color = &color
	return state
}

func (state *State) SetMeta(meta *Meta) *State {
	state.Meta = meta
	return state
}

func (state *State) SetName(name string) *State {
	state.Name = &name
	return state
}

func (state *State) SetStateType(stateType StateType) *State {
	state.StateType = stateType
	return state
}

func (state State) String() string {
	return Stringify(state)
}

// StateType Тип статуса.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-statusy-dokumentow-statusy-atributy-suschnosti-tip-statusa
type StateType string

const (
	StateTypeRegular      StateType = "Regular"      // Обычный (значение по умолчанию)
	StateTypeSuccessful   StateType = "Successful"   // Финальный положительный
	StateTypeUnsuccessful StateType = "Unsuccessful" // Финальный отрицательный
)
