package moysklad

import (
	"github.com/google/uuid"
)

// State Статус документов.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-statusy-dokumentow
type State struct {
	AccountID  *uuid.UUID `json:"accountId,omitempty"`  // ID учетной записи
	Color      *uint64    `json:"color,omitempty"`      // Цвет Статуса
	EntityType *string    `json:"entityType,omitempty"` // Тип сущности, к которой относится Статус (ключевое слово в рамках JSON API)
	ID         *uuid.UUID `json:"id,omitempty"`         // ID сущности
	Meta       *Meta      `json:"meta,omitempty"`       // Метаданные
	Name       *string    `json:"name,omitempty"`       // Наименование Статуса
	StateType  StateType  `json:"stateType,omitempty"`  // Тип Статуса
}

func (s *State) IsEqual(state *State) bool {
	return s.Meta.IsEqual(state.Meta)
}

func (s State) String() string {
	return Stringify(s)
}

type States = Slice[State]

// StateType Тип статуса.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-statusy-dokumentow-statusy-atributy-suschnosti-tip-statusa
type StateType string

const (
	StateTypeRegular      StateType = "Regular"      // Обычный (значение по умолчанию)
	StateTypeSuccessful   StateType = "Successful"   // Финальный положительный
	StateTypeUnsuccessful StateType = "Unsuccessful" // Финальный отрицательный
)
