package moysklad

import (
	"github.com/google/uuid"
)

// CustomerOrder Заказ покупателя.
// Ключевое слово: customerorder
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-zakaz-pokupatelq
type CustomerOrder struct {
	AccountId             *uuid.UUID                        `json:"accountId,omitempty"`             // ID учетной записи
	Agent                 *Counterparty                     `json:"agent,omitempty"`                 // Метаданные контрагента
	AgentAccount          *AgentAccount                     `json:"agentAccount,omitempty"`          // Метаданные счета контрагента
	Applicable            *bool                             `json:"applicable,omitempty"`            // Отметка о проведении
	Attributes            *Attributes                       `json:"attributes,omitempty"`            // Коллекция метаданных доп. полей. Поля объекта
	Code                  *string                           `json:"code,omitempty"`                  // Код Заказа покупателя
	Contract              *Contract                         `json:"contract,omitempty"`              // Метаданные договора
	Created               *Timestamp                        `json:"created,omitempty"`               // Дата создания
	Deleted               *Timestamp                        `json:"deleted,omitempty"`               // Момент последнего удаления Заказа покупателя
	DeliveryPlannedMoment *Timestamp                        `json:"deliveryPlannedMoment,omitempty"` // Планируемая дата отгрузки
	Description           *string                           `json:"description,omitempty"`           // Комментарий Заказа покупателя
	ExternalCode          *string                           `json:"externalCode,omitempty"`          // Внешний код Заказа покупателя
	Files                 *Files                            `json:"files,omitempty"`                 // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group                 *Group                            `json:"group,omitempty"`                 // Отдел сотрудника
	Id                    *uuid.UUID                        `json:"id,omitempty"`                    // ID сущности
	InvoicedSum           *float64                          `json:"invoicedSum,omitempty"`           // Сумма счетов покупателю
	Meta                  *Meta                             `json:"meta,omitempty"`                  // Метаданные
	Name                  *string                           `json:"name,omitempty"`                  // Наименование
	Moment                *Timestamp                        `json:"moment,omitempty"`                // Дата документа
	Organization          *Organization                     `json:"organization,omitempty"`          // Метаданные юрлица
	OrganizationAccount   *AgentAccount                     `json:"organizationAccount,omitempty"`   // Метаданные счета юрлица
	Owner                 *Employee                         `json:"owner,omitempty"`                 // Владелец (Сотрудник)
	PayedSum              *float64                          `json:"payedSum,omitempty"`              // Сумма входящих платежей по Заказу
	Positions             *Positions[CustomerOrderPosition] `json:"positions,omitempty"`             // Метаданные позиций Заказа покупателя
	Printed               *bool                             `json:"printed,omitempty"`               // Напечатан ли документ
	Project               *Project                          `json:"project,omitempty"`               // Метаданные проекта
	Published             *bool                             `json:"published,omitempty"`             // Опубликован ли документ
	Rate                  *Rate                             `json:"rate,omitempty"`                  // Валюта
	ReservedSum           *float64                          `json:"reservedSum,omitempty"`           // Сумма товаров в резерве
	SalesChannel          *SalesChannel                     `json:"salesChannel,omitempty"`          // Метаданные канала продаж
	Shared                *bool                             `json:"shared,omitempty"`                // Общий доступ
	ShipmentAddress       *string                           `json:"shipmentAddress,omitempty"`       // Адрес доставки Заказа покупателя
	ShipmentAddressFull   *Address                          `json:"shipmentAddressFull,omitempty"`   // Адрес доставки Заказа покупателя с детализацией по отдельным полям
	ShippedSum            *float64                          `json:"shippedSum,omitempty"`            // Сумма отгруженного
	State                 *State                            `json:"state,omitempty"`                 // Метаданные статуса заказа
	Store                 *Store                            `json:"store,omitempty"`                 // Метаданные склада
	Sum                   *float64                          `json:"sum,omitempty"`                   // Сумма
	SyncId                *uuid.UUID                        `json:"syncId,omitempty"`                // ID синхронизации. После заполнения недоступен для изменения
	TaxSystem             TaxSystem                         `json:"taxSystem,omitempty"`             // Код системы налогообложения
	Updated               *Timestamp                        `json:"updated,omitempty"`               // Момент последнего обновления
	VatEnabled            *bool                             `json:"vatEnabled,omitempty"`            // Учитывается ли НДС
	VatIncluded           *bool                             `json:"vatIncluded,omitempty"`           // Включен ли НДС в цену
	VatSum                *float64                          `json:"vatSum,omitempty"`                // Сумма НДС
	PurchaseOrders        *PurchaseOrders                   `json:"purchaseOrders,omitempty"`        // Массив ссылок на связанные заказы поставщикам в формате Метаданных
	Demands               *Demands                          `json:"demands,omitempty"`               // Массив ссылок на связанные отгрузки в формате Метаданных
	Payments              *Payments                         `json:"payments,omitempty"`              // Массив ссылок на связанные платежи в формате Метаданных
	InvoicesOut           *InvoicesOut                      `json:"invoicesOut,omitempty"`           // Массив ссылок на связанные счета покупателям в формате Метаданных
	Moves                 *Moves                            `json:"moves,omitempty"`                 // Массив ссылок на связанные перемещния в формате Метаданных
	Prepayments           *Prepayments                      `json:"prepayments,omitempty"`           // Массив ссылок на связанные предоплаты в формате Метаданных
}

func (c CustomerOrder) String() string {
	return Stringify(c)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (c CustomerOrder) GetMeta() *Meta {
	return c.Meta
}

func (c CustomerOrder) MetaType() MetaType {
	return MetaTypeCustomerOrder
}

type CustomerOrders = Iterator[CustomerOrder]

// ConvertToOperation удовлетворяет интерфейсу OperationInType
//func (c CustomerOrder) ConvertToOperation(linkedSum *float64) (*OperationIn, error) {
//	return convertToOperation[OperationIn](&c, linkedSum) // &OperationIn{}, nil //OperationFromEntity(c, linkedSum)
//}

// CustomerOrderPosition Позиция Заказа покупателя.
// Ключевое слово: customerorderposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-zakaz-pokupatelq-zakazy-pokupatelej-pozicii-zakaza-pokupatelq
type CustomerOrderPosition struct {
	AccountId  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	Discount   *float64            `json:"discount,omitempty"`   // Процент скидки или наценки. Наценка указывается отрицательным числом, т.е. -10 создаст наценку в 10%
	Id         *uuid.UUID          `json:"id,omitempty"`         // ID сущности
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Price      *float64            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Reserve    *float64            `json:"reserve,omitempty"`    // Резерв данной позиции
	Shipped    *float64            `json:"shipped,omitempty"`    // Доставлено
	TaxSystem  GoodTaxSystem       `json:"taxSystem,omitempty"`  // Код системы налогообложения
	Vat        *int                `json:"vat,omitempty"`        // НДС, которым облагается текущая позиция
	VatEnabled *bool               `json:"vatEnabled,omitempty"` // Включен ли НДС для позиции. С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
	Stock      *Stock              `json:"stock,omitempty"`      // Остатки и себестоимость `?fields=stock&expand=positions`
}

func (c CustomerOrderPosition) String() string {
	return Stringify(c)
}

func (c CustomerOrderPosition) MetaType() MetaType {
	return MetaTypeCustomerOrderPosition
}
