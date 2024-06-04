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

func (cashier Cashier) GetAccountID() uuid.UUID {
	return Deref(cashier.AccountID)
}

func (cashier Cashier) GetEmployee() Employee {
	return Deref(cashier.Employee)
}

func (cashier Cashier) GetID() uuid.UUID {
	return Deref(cashier.ID)
}

func (cashier Cashier) GetMeta() Meta {
	return Deref(cashier.Meta)
}

func (cashier Cashier) GetRetailStore() RetailStore {
	return Deref(cashier.RetailStore)
}

func (cashier *Cashier) SetEmployee(employee *Employee) *Cashier {
	cashier.Employee = employee
	return cashier
}

func (cashier *Cashier) SetMeta(meta *Meta) *Cashier {
	cashier.Meta = meta
	return cashier
}

func (cashier *Cashier) SetRetailStore(retailStore *RetailStore) *Cashier {
	cashier.RetailStore = retailStore
	return cashier
}

func (cashier Cashier) String() string {
	return Stringify(cashier)
}

func (cashier Cashier) MetaType() MetaType {
	return MetaTypeCashier
}

type Cashiers MetaArray[Cashier]
