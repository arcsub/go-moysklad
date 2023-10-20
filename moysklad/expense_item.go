package moysklad

import (
	"github.com/google/uuid"
)

// ExpenseItem Статья расходов.
// Ключевое слово: expenseitem
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/dictionaries/#suschnosti-stat-q-rashodow
type ExpenseItem struct {
	AccountID    *uuid.UUID `json:"accountId,omitempty"`    // ID учетной записи
	Code         *string    `json:"code,omitempty"`         // Код Статьи расходов
	Description  *string    `json:"description,omitempty"`  // Описание Статьи расходов
	ExternalCode *string    `json:"externalCode,omitempty"` // Внешний код Статьи расходов
	ID           *uuid.UUID `json:"id,omitempty"`           // ID Статьи расходов
	Meta         *Meta      `json:"meta,omitempty"`         // Метаданные о Статье расходов
	Name         *string    `json:"name,omitempty"`         // Наименование Статьи расходов
	Updated      *Timestamp `json:"updated,omitempty"`      // Момент последнего обновления сущности
}

func (e ExpenseItem) String() string {
	return Stringify(e)
}

func (e ExpenseItem) MetaType() MetaType {
	return MetaTypeExpenseItem
}
