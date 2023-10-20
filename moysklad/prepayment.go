package moysklad

import (
	"github.com/google/uuid"
)

// Prepayment Предоплата.
// Ключевое слово: prepayment
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-predoplata
type Prepayment struct {
	AccountID     *uuid.UUID                     `json:"accountId,omitempty"`     // ID учетной записи
	Agent         *Counterparty                  `json:"agent,omitempty"`         // Ссылка на контрагента
	Applicable    *bool                          `json:"applicable,omitempty"`    // Отметка о проведении
	Attributes    *Attributes                    `json:"attributes,omitempty"`    // Коллекция метаданных доп. полей. Поля объекта
	CashSum       *float64                       `json:"cashSum,omitempty"`       // Оплачено наличными
	Code          *string                        `json:"code,omitempty"`          // Код
	Created       *Timestamp                     `json:"created,omitempty"`       // Дата создания
	CustomerOrder *CustomerOrder                 `json:"customerOrder,omitempty"` // Метаданные Заказа Покупателя
	Deleted       *Timestamp                     `json:"deleted,omitempty"`       // Момент последнего удаления
	Description   *string                        `json:"description,omitempty"`   // Комментарий
	ExternalCode  *string                        `json:"externalCode,omitempty"`  // Внешний код
	Files         *Files                         `json:"files,omitempty"`         // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group         *Group                         `json:"group,omitempty"`         // Отдел сотрудника
	ID            *uuid.UUID                     `json:"id,omitempty"`            // ID сущности
	Meta          *Meta                          `json:"meta,omitempty"`          // Метаданные
	Moment        *Timestamp                     `json:"moment,omitempty"`        // Дата документа
	Name          *string                        `json:"name,omitempty"`          // Наименование
	NoCashSum     *float64                       `json:"noCashSum,omitempty"`     // Оплачено картой
	Organization  *Organization                  `json:"organization,omitempty"`  // Метаданные юрлица
	Owner         *Employee                      `json:"owner,omitempty"`         // Владелец (Сотрудник)
	Positions     *Positions[PrepaymentPosition] `json:"positions,omitempty"`     // Метаданные позиций Предоплаты
	Printed       *bool                          `json:"printed,omitempty"`       // Напечатан ли документ
	Published     *bool                          `json:"published,omitempty"`     // Опубликован ли документ
	QRSum         *float64                       `json:"qrSum,omitempty"`         // Оплачено по QR-коду
	Rate          *Rate                          `json:"rate,omitempty"`          // Валюта
	RetailShift   *RetailShift                   `json:"retailShift,omitempty"`   // Метаданные Розничной смены
	RetailStore   *RetailStore                   `json:"retailStore,omitempty"`   // Метаданные Точки продаж
	Returns       *PrepaymentReturns             `json:"returns,omitempty"`       // Коллекция метаданных на связанные возвраты
	Shared        *bool                          `json:"shared,omitempty"`        // Общий доступ
	State         *State                         `json:"state,omitempty"`         // Метаданные статуса
	Sum           *float64                       `json:"sum,omitempty"`           // Сумма
	SyncID        *uuid.UUID                     `json:"syncId,omitempty"`        // ID синхронизации. После заполнения недоступен для изменения
	TaxSystem     TaxSystem                      `json:"taxSystem,omitempty"`     // Код системы налогообложения
	Updated       *Timestamp                     `json:"updated,omitempty"`       // Момент последнего обновления
	VatEnabled    *bool                          `json:"vatEnabled,omitempty"`    // Учитывается ли НДС
	VatIncluded   *bool                          `json:"vatIncluded,omitempty"`   // Включен ли НДС в цену
	VatSum        *float64                       `json:"vatSum,omitempty"`        // Сумма включая НДС
}

func (p Prepayment) String() string {
	return Stringify(p)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (p Prepayment) GetMeta() *Meta {
	return p.Meta
}

func (p Prepayment) MetaType() MetaType {
	return MetaTypePrepayment
}

type Prepayments = Slice[Prepayment]

// PrepaymentPosition Позиция Предоплаты.
// Ключевое слово: prepaymentposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-predoplata-predoplaty-pozicii-predoplaty
type PrepaymentPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	Discount   *float64            `json:"discount,omitempty"`   // Процент скидки или наценки. Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Price      *float64            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Vat        *int                `json:"vat,omitempty"`        // НДС, которым облагается текущая позиция
	VatEnabled *bool               `json:"vatEnabled,omitempty"` // Включен ли НДС для позиции. С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
}

func (p PrepaymentPosition) String() string {
	return Stringify(p)
}

func (p PrepaymentPosition) MetaType() MetaType {
	return MetaTypePrepaymentPosition
}
