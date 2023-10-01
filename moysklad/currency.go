package moysklad

import (
	"github.com/google/uuid"
)

// Currency Валюта.
// Ключевое слово: currency
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-valuta
type Currency struct {
	Archived       *bool          `json:"archived,omitempty"`       // Добавлена ли Валюта в архив
	Code           *string        `json:"code,omitempty"`           // Цифровой код Валюты
	Default        *bool          `json:"default,omitempty"`        // Является ли валюта валютой учета
	FullName       *string        `json:"fullName,omitempty"`       // Полное наименование Валюты
	Id             *uuid.UUID     `json:"id,omitempty"`             // ID сущности
	Indirect       *bool          `json:"indirect,omitempty"`       // Признак обратного курса Валюты
	ISOCode        *string        `json:"isoCode,omitempty"`        // Буквенный код Валюты
	MajorUnit      *CurrencyUnit  `json:"majorUnit,omitempty"`      // Формы единиц целой части Валюты
	Margin         *float64       `json:"margin,omitempty"`         // Наценка при автоматическом обновлении курса
	Meta           *Meta          `json:"meta,omitempty"`           // Метаданные
	MinorUnit      *CurrencyUnit  `json:"minorUnit,omitempty"`      // Формы единиц дробной части Валюты
	Multiplicity   *int           `json:"multiplicity,omitempty"`   // Кратность курса Валюты
	Name           *string        `json:"name,omitempty"`           // Наименование
	Rate           *float64       `json:"rate,omitempty"`           // Курс Валюты
	RateUpdateType RateUpdateType `json:"rateUpdateType,omitempty"` // Способ обновления курса Валюты
	System         *bool          `json:"system,omitempty"`         // Основана ли валюта на валюте из системного справочника
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
