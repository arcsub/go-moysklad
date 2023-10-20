package moysklad

import (
	"github.com/google/uuid"
)

// PurchaseReturn Возврат поставщику.
// Ключевое слово: purchasereturn
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vozwrat-postawschiku
type PurchaseReturn struct {
	AccountID           *uuid.UUID                         `json:"accountId,omitempty"`           // ID учетной записи
	Agent               *Counterparty                      `json:"agent,omitempty"`               // Метаданные контрагента
	AgentAccount        *AgentAccount                      `json:"agentAccount,omitempty"`        // Метаданные счета контрагента
	Applicable          *bool                              `json:"applicable,omitempty"`          // Отметка о проведении
	Attributes          *Attributes                        `json:"attributes,omitempty"`          // Коллекция метаданных доп. полей. Поля объекта
	Code                *string                            `json:"code,omitempty"`                // Код
	Contract            *Contract                          `json:"contract,omitempty"`            // Метаданные договора
	Created             *Timestamp                         `json:"created,omitempty"`             // Дата создания
	Deleted             *Timestamp                         `json:"deleted,omitempty"`             // Момент последнего удаления
	Description         *string                            `json:"description,omitempty"`         // Комментарий
	ExternalCode        *string                            `json:"externalCode,omitempty"`        // Внешний код
	Files               *Files                             `json:"files,omitempty"`               // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group               *Group                             `json:"group,omitempty"`               // Отдел сотрудника
	ID                  *uuid.UUID                         `json:"id,omitempty"`                  // ID сущности
	Meta                *Meta                              `json:"meta,omitempty"`                // Метаданные
	Moment              *Timestamp                         `json:"moment,omitempty"`              // Дата документа
	Name                *string                            `json:"name,omitempty"`                // Наименование
	Organization        *Organization                      `json:"organization,omitempty"`        // Метаданные юрлица
	OrganizationAccount *AgentAccount                      `json:"organizationAccount,omitempty"` // Метаданные счета юрлица
	Owner               *Employee                          `json:"owner,omitempty"`               // Владелец (Сотрудник)
	Printed             *bool                              `json:"printed,omitempty"`             // Напечатан ли документ
	Project             *Project                           `json:"project,omitempty"`             // Проект
	Published           *bool                              `json:"published,omitempty"`           // Опубликован ли документ
	Rate                *Rate                              `json:"rate,omitempty"`                // Валюта
	Shared              *bool                              `json:"shared,omitempty"`              // Общий доступ
	State               *State                             `json:"state,omitempty"`               // Метаданные статуса
	Store               *Store                             `json:"store,omitempty"`               // Метаданные склада
	Sum                 *float64                           `json:"sum,omitempty"`                 // Сумма
	SyncID              *uuid.UUID                         `json:"syncId,omitempty"`              // ID синхронизации. После заполнения недоступен для изменения
	Updated             *Timestamp                         `json:"updated,omitempty"`             // Момент последнего обновления
	VatEnabled          *bool                              `json:"vatEnabled,omitempty"`          // Учитывается ли НДС
	VatIncluded         *bool                              `json:"vatIncluded,omitempty"`         // Включен ли НДС в цену
	VatSum              *float64                           `json:"vatSum,omitempty"`              // Сумма включая НДС
	Positions           *Positions[PurchaseReturnPosition] `json:"positions,omitempty"`           // Метаданные позиций Заказа поставщику
	Supply              *Supply                            `json:"supply,omitempty"`              // Ссылка на приемку, по которой произошел возврат в формате Метаданных Поле является необходимым для возврата с основанием.
	FactureIn           *FactureIn                         `json:"factureIn,omitempty"`           // Ссылка на Счет-фактуру полученный в формате Метаданных
	FactureOut          *FactureOut                        `json:"factureOut,omitempty"`          // Ссылка на Счет-фактуру выданный в формате Метаданных
	InvoicedSum         *float64                           `json:"invoicedSum,omitempty"`         // Сумма счетов поставщику
	PayedSum            *float64                           `json:"payedSum,omitempty"`            // Сумма входящих платежей по Счету поставщика
	Payments            *Payments                          `json:"payments,omitempty"`            // Массив ссылок на связанные платежи в формате Метаданных
}

func (p PurchaseReturn) String() string {
	return Stringify(p)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (p PurchaseReturn) GetMeta() *Meta {
	return p.Meta
}

func (p PurchaseReturn) MetaType() MetaType {
	return MetaTypePurchaseReturn
}

type PurchaseReturns = Iterator[PurchaseReturn]

// PurchaseReturnPosition Позиция Возврата поставщику.
// Ключевое слово: purchasereturnposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vozwrat-postawschiku-vozwraty-postawschikam-pozicii-vozwrata-postawschiku
type PurchaseReturnPosition struct {
	AccountID  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	Discount   *float64            `json:"discount,omitempty"`   // Процент скидки или наценки. Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%
	ID         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Price      *float64            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Slot       *Slot               `json:"slot,omitempty"`       // Ячейка на складе
	Things     *Things             `json:"things,omitempty"`     // Серийные номера. Значение данного атрибута игнорируется, если товар позиции не находится на серийном учете. В ином случае количество товаров в позиции будет равно количеству серийных номеров, переданных в значении атрибута.
	Vat        *int                `json:"vat,omitempty"`        // НДС, которым облагается текущая позиция
	VatEnabled *bool               `json:"vatEnabled,omitempty"` // Включен ли НДС для позиции. С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
	Stock      *Stock              `json:"stock,omitempty"`      // Остатки и себестоимость `?fields=stock&expand=positions`
}

func (p PurchaseReturnPosition) String() string {
	return Stringify(p)
}

func (p PurchaseReturnPosition) MetaType() MetaType {
	return MetaTypePurchaseReturnPosition
}

// PurchaseReturnTemplateArg
// Документ: Возврат поставщику (purchasereturn)
// Основание, на котором он может быть создан:
// - Приемка (supply)
type PurchaseReturnTemplateArg struct {
	Supply *MetaWrapper `json:"supply,omitempty"`
}
