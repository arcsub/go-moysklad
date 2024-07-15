package moysklad

import (
	"fmt"
	"github.com/google/uuid"
	"image/color"
	"strconv"
	"strings"
)

// State Статус документов.
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-statusy-dokumentow
type State struct {
	AccountID  *uuid.UUID `json:"accountId,omitempty"`  // ID учётной записи
	Color      *uint64    `json:"color,omitempty"`      // Цвет Статуса
	EntityType MetaType   `json:"entityType,omitempty"` // Тип сущности, к которой относится Статус (Код сущности в рамках JSON API)
	ID         *uuid.UUID `json:"id,omitempty"`         // ID Статуса
	Meta       *Meta      `json:"meta,omitempty"`       // Метаданные Статуса
	Name       *string    `json:"name,omitempty"`       // Наименование Статуса
	StateType  StateType  `json:"stateType,omitempty"`  // Тип Статуса
}

// Clean возвращает указатель на объект с единственным заполненным полем [Meta].
//
// Метод позволяет избавиться от лишних данных при передаче запроса.
func (state State) Clean() *State {
	if state.Meta == nil {
		return nil
	}
	return &State{Meta: state.Meta}
}

// GetAccountID возвращает ID учётной записи.
func (state State) GetAccountID() uuid.UUID {
	return Deref(state.AccountID)
}

// GetColor возвращает Цвет Статуса.
func (state State) GetColor() uint64 {
	return Deref(state.Color)
}

// GetEntityType возвращает Тип сущности, к которой относится Статус (Код сущности в рамках JSON API).
func (state State) GetEntityType() MetaType {
	return state.EntityType
}

// GetID возвращает ID Статуса.
func (state State) GetID() uuid.UUID {
	return Deref(state.ID)
}

// GetMeta возвращает Метаданные Статуса.
func (state State) GetMeta() Meta {
	return Deref(state.Meta)
}

// GetName возвращает Наименование Статуса.
func (state State) GetName() string {
	return Deref(state.Name)
}

// GetStateType возвращает Тип Статуса.
func (state State) GetStateType() StateType {
	return state.StateType
}

// SetColor устанавливает Цвет Статуса.
func (state *State) SetColor(color uint64) *State {
	state.Color = &color
	return state
}

// SetColorRBG устанавливает Цвет Статуса.
//
// Принимает RGB формат цвета типа string.
//
// Примеры аргументов: "#E6E6E6" ,"e3e3e3", "FFF"
func (state *State) SetColorRBG(rgbString string) *State {
	clr := RGBtoUint64(rgbString)
	state.Color = &clr
	return state
}

// SetMeta устанавливает Метаданные Статуса.
func (state *State) SetMeta(meta *Meta) *State {
	state.Meta = meta
	return state
}

// SetName устанавливает Наименование Статуса.
func (state *State) SetName(name string) *State {
	state.Name = &name
	return state
}

// SetStateType устанавливает Тип Статуса.
func (state *State) SetStateType(stateType StateType) *State {
	state.StateType = stateType
	return state
}

// SetStateTypeRegular устанавливает Тип Статуса в значение [StateTypeRegular].
func (state *State) SetStateTypeRegular() *State {
	state.StateType = StateTypeRegular
	return state
}

// SetStateTypeSuccessful устанавливает Тип Статуса в значение [StateTypeSuccessful].
func (state *State) SetStateTypeSuccessful() *State {
	state.StateType = StateTypeSuccessful
	return state
}

// SetStateTypeUnsuccessful устанавливает Тип Статуса в значение [StateTypeUnsuccessful].
func (state *State) SetStateTypeUnsuccessful() *State {
	state.StateType = StateTypeUnsuccessful
	return state
}

// String реализует интерфейс [fmt.Stringer].
func (state State) String() string {
	return Stringify(state)
}

// StateType Тип статуса.
//
// Возможные значения:
//   - StateTypeRegular      – Обычный (значение по умолчанию)
//   - StateTypeSuccessful   – Финальный положительный
//   - StateTypeUnsuccessful – Финальный отрицательный
//
// [Документация МойСклад]
//
// [Документация МойСклад]: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-statusy-dokumentow-statusy-atributy-suschnosti-tip-statusa
type StateType string

const (
	StateTypeRegular      StateType = "Regular"      // Обычный (значение по умолчанию)
	StateTypeSuccessful   StateType = "Successful"   // Финальный положительный
	StateTypeUnsuccessful StateType = "Unsuccessful" // Финальный отрицательный
)

// RGBtoUint64 конвертирует код цвета из формата RRGGBB / RGB в uint64.
//
// # Примеры:
//
// RGBtoUint64("#E6E6E6") → 15132390
//
// RGBtoUint64("e3e3e3") → 14935011
//
// RGBtoUint64("FFF") → 16777215
func RGBtoUint64(rgbString string) uint64 {
	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		return 0
	}

	n := strings.Replace(rgbString, "0x", "", -1)
	n = strings.Replace(n, "0X", "", -1)
	n = strings.Replace(n, "#", "", -1)

	c := color.RGBA{A: 0xff}

	switch len(n) {
	case 6: // RRGGBB
		c.R = hexToByte(n[0])<<4 + hexToByte(n[1])
		c.G = hexToByte(n[2])<<4 + hexToByte(n[3])
		c.B = hexToByte(n[4])<<4 + hexToByte(n[5])
	case 3: // RGB
		c.R = hexToByte(n[0]) * 17
		c.G = hexToByte(n[1]) * 17
		c.B = hexToByte(n[2]) * 17
	default: // invalid format
		return 0
	}

	hex := fmt.Sprintf("%02x%02x%02x", c.R, c.G, c.B)

	result, err := strconv.ParseUint(hex, 16, 64)
	if err != nil {
		return 0
	}

	return result
}
