package moysklad

import (
	"github.com/google/uuid"
)

// Cashier Кассир.
// Ключевое слово: cashier
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-kassir
type Cashier struct {
	AccountID   *uuid.UUID   `json:"accountId,omitempty"`   // ID учетной записи
	Employee    *Employee    `json:"employee,omitempty"`    // Метаданные сотрудника, которого представляет собой кассир
	ID          *uuid.UUID   `json:"id,omitempty"`          // ID сущности
	Meta        *Meta        `json:"meta,omitempty"`        // Метаданные
	RetailStore *RetailStore `json:"retailStore,omitempty"` // Метаданные точки продаж, к которой прикреплен кассир
}

func (c Cashier) String() string {
	return Stringify(c)
}

func (c Cashier) MetaType() MetaType {
	return MetaTypeCashier
}

type Cashiers MetaArray[Cashier]
