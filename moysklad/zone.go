package moysklad

import (
	"github.com/google/uuid"
)

// Zone Зона склада.
// Ключевое слово: storezone
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-sklad-zony-sklada
type Zone struct {
	AccountID    *uuid.UUID `json:"accountId,omitempty"`    // ID учетной записи
	ExternalCode *string    `json:"externalCode,omitempty"` // Внешний код Зоны
	ID           *uuid.UUID `json:"id,omitempty"`           // ID Зоны
	Meta         *Meta      `json:"meta,omitempty"`         // Метаданные Зоны
	Name         *string    `json:"name,omitempty"`         // Наименование Зоны
	Updated      *Timestamp `json:"updated,omitempty"`      // Момент последнего обновления Зоны
}

func (z Zone) String() string {
	return Stringify(z)
}

func (z Zone) MetaType() MetaType {
	return MetaTypeStoreZone
}
