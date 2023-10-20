package moysklad

import (
	"github.com/google/uuid"
)

// PrepaymentReturn Возврат предоплаты.
// Ключевое слово: prepaymentreturn
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vozwrat-predoplaty
type PrepaymentReturn struct {
	AccountID    *uuid.UUID                           `json:"accountId,omitempty"`    // ID учетной записи
	Agent        *Counterparty                        `json:"agent,omitempty"`        // Ссылка на контрагента
	Applicable   *bool                                `json:"applicable,omitempty"`   // Отметка о проведении
	Attributes   *Attributes                          `json:"attributes,omitempty"`   // Коллекция метаданных доп. полей. Поля объекта
	CashSum      *float64                             `json:"cashSum,omitempty"`      // Оплачено наличными
	Code         *string                              `json:"code,omitempty"`         // Код
	Created      *Timestamp                           `json:"created,omitempty"`      // Дата создания
	Deleted      *Timestamp                           `json:"deleted,omitempty"`      // Момент последнего удаления
	Description  *string                              `json:"description,omitempty"`  // Комментарий
	ExternalCode *string                              `json:"externalCode,omitempty"` // Внешний код
	Files        *Files                               `json:"files,omitempty"`        // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group        *Group                               `json:"group,omitempty"`        // Отдел сотрудника
	ID           *uuid.UUID                           `json:"id,omitempty"`           // ID сущности
	Meta         *Meta                                `json:"meta,omitempty"`         // Метаданные
	Moment       *Timestamp                           `json:"moment,omitempty"`       // Дата документа
	Name         *string                              `json:"name,omitempty"`         // Наименование
	NoCashSum    *float64                             `json:"noCashSum,omitempty"`    // Оплачено картой
	Organization *Organization                        `json:"organization,omitempty"` // Метаданные юрлица
	Owner        *Employee                            `json:"owner,omitempty"`        // Владелец (Сотрудник)
	Positions    *Positions[PrepaymentReturnPosition] `json:"positions,omitempty"`    // Метаданные позиций
	Prepayment   *Prepayment                          `json:"prepayment,omitempty"`   // Метаданные Предоплаты
	Printed      *bool                                `json:"printed,omitempty"`      // Напечатан ли документ
	Published    *bool                                `json:"published,omitempty"`    // Опубликован ли документ
	QRSum        *float64                             `json:"qrSum,omitempty"`        // Оплачено по QR-коду
	Rate         *Rate                                `json:"rate,omitempty"`         // Валюта
	RetailShift  *RetailShift                         `json:"retailShift,omitempty"`  // Метаданные Розничной смены
	RetailStore  *RetailStore                         `json:"retailStore,omitempty"`  // Метаданные Точки продаж
	Shared       *bool                                `json:"shared,omitempty"`       // Общий доступ
	State        *State                               `json:"state,omitempty"`        // Метаданные статуса
	Sum          *float64                             `json:"sum,omitempty"`          // Сумма
	SyncID       *uuid.UUID                           `json:"syncId,omitempty"`       // ID синхронизации. После заполнения недоступен для изменения
	TaxSystem    TaxSystem                            `json:"taxSystem,omitempty"`    // Код системы налогообложения
	Updated      *Timestamp                           `json:"updated,omitempty"`      // Момент последнего обновления
	VatEnabled   *bool                                `json:"vatEnabled,omitempty"`   // Учитывается ли НДС
	VatIncluded  *bool                                `json:"vatIncluded,omitempty"`  // Включен ли НДС в цену
	VatSum       *float64                             `json:"vatSum,omitempty"`       // Сумма включая НДС
}

func (p PrepaymentReturn) String() string {
	return Stringify(p)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (p PrepaymentReturn) GetMeta() *Meta {
	return p.Meta
}

func (p PrepaymentReturn) MetaType() MetaType {
	return MetaTypePrepaymentReturn
}

type PrepaymentReturns = Iterator[PrepaymentReturn]

// PrepaymentReturnPosition Позиция Возврата предоплаты.
// Ключевое слово: prepaymentreturnposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vozwrat-predoplaty-atributy-suschnosti-pozicii-vozwrata-predoplaty
type PrepaymentReturnPosition struct {
	PrepaymentPosition
}

func (p PrepaymentReturnPosition) String() string {
	return Stringify(p)
}

func (p PrepaymentReturnPosition) MetaType() MetaType {
	return MetaTypePrepaymentReturnPosition
}
