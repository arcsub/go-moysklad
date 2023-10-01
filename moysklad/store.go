package moysklad

import (
	"github.com/google/uuid"
)

// Store Склад.
// Ключевое слово: store
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad
type Store struct {
	AccountId    *uuid.UUID       `json:"accountId,omitempty"`    // ID учетной записи
	Address      *string          `json:"address,omitempty"`      // Адрес склада
	AddressFull  *Address         `json:"addressFull,omitempty"`  // Адрес с детализацией по отдельным полям
	Archived     *bool            `json:"archived,omitempty"`     // Добавлен ли Склад в архив
	Attributes   *Attributes      `json:"attributes,omitempty"`   // Массив метаданных дополнительных полей склада
	Code         *string          `json:"code,omitempty"`         // Код Склада
	Description  *string          `json:"description,omitempty"`  // Комментарий к Складу
	ExternalCode *string          `json:"externalCode,omitempty"` // Внешний код Склада
	Group        *Group           `json:"group,omitempty"`        // Отдел сотрудника
	Id           *uuid.UUID       `json:"id,omitempty"`           // ID сущности
	Meta         *Meta            `json:"meta,omitempty"`         // Метаданные
	Name         *string          `json:"name,omitempty"`         // Наименование Склада
	Owner        *Employee        `json:"owner,omitempty"`        // Владелец (Сотрудник)
	Parent       *Store           `json:"parent,omitempty"`       // Метаданные родительского склада (Группы)
	PathName     *string          `json:"pathName,omitempty"`     // Группа Склада
	Shared       *bool            `json:"shared,omitempty"`       // Общий доступ
	Updated      *Timestamp       `json:"updated,omitempty"`      // Момент последнего обновления Склада
	Zones        *MetaArray[Zone] `json:"zones,omitempty"`        // Зоны склада
	Slots        *MetaArray[Slot] `json:"slots,omitempty"`        // Ячейки склада
}

func (s Store) String() string {
	return Stringify(s)
}

func (s Store) MetaType() MetaType {
	return MetaTypeStore
}
