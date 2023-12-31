package moysklad

import (
	"github.com/google/uuid"
)

// Currency Валюта.
// Ключевое слово: currency
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-valuta
type Currency struct {
	Margin         *float64       `json:"margin,omitempty"`
	ISOCode        *string        `json:"isoCode,omitempty"`
	Default        *bool          `json:"default,omitempty"`
	FullName       *string        `json:"fullName,omitempty"`
	Archived       *bool          `json:"archived,omitempty"`
	Indirect       *bool          `json:"indirect,omitempty"`
	Code           *string        `json:"code,omitempty"`
	MajorUnit      *CurrencyUnit  `json:"majorUnit,omitempty"`
	ID             *uuid.UUID     `json:"id,omitempty"`
	Meta           *Meta          `json:"meta,omitempty"`
	MinorUnit      *CurrencyUnit  `json:"minorUnit,omitempty"`
	Multiplicity   *int           `json:"multiplicity,omitempty"`
	Name           *string        `json:"name,omitempty"`
	Rate           *float64       `json:"rate,omitempty"`
	System         *bool          `json:"system,omitempty"`
	RateUpdateType RateUpdateType `json:"rateUpdateType,omitempty"`
}

func (c Currency) String() string {
	return Stringify(c)
}

func (c Currency) MetaType() MetaType {
	return MetaTypeCurrency
}

// RateUpdateType Способ обновления курса.
type RateUpdateType string

const (
	RateUpdateTypeAuto   RateUpdateType = "auto"   // Автоматический
	RateUpdateTypeManual RateUpdateType = "manual" // Ручной
)

// CurrencyUnit Формы единиц.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-valuta-formy-edinic
type CurrencyUnit struct {
	Gender *UnitGender `json:"gender,omitempty"` // Грамматический род единицы валюты
	S1     *string     `json:"s1,omitempty'"`    // Форма единицы, используемая при числительном 1
	S2     *string     `json:"s2,omitempty"`     // Форма единицы, используемая при числительном 2
	S5     *string     `json:"s5,omitempty"`     // Форма единицы, используемая при числительном 5
}

func (c CurrencyUnit) String() string {
	return Stringify(c)
}

// UnitGender Грамматический род единицы валюты.
type UnitGender string

const (
	UnitGenderMasculine UnitGender = "masculine" // мужской
	UnitGenderFeminine  UnitGender = "feminine"  // женский
)
