package moysklad

import (
	"github.com/google/uuid"
)

// InternalOrder Внутренний заказ.
// Ключевое слово: internalorder
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vnutrennij-zakaz
type InternalOrder struct {
	AccountId             *uuid.UUID                        `json:"accountId,omitempty"`             // ID учетной записи
	Applicable            *bool                             `json:"applicable,omitempty"`            // Отметка о проведении
	Attributes            *Attributes                       `json:"attributes,omitempty"`            // Коллекция метаданных доп. полей. Поля объекта
	Code                  *string                           `json:"code,omitempty"`                  // Код
	Created               *Timestamp                        `json:"created,omitempty"`               // Дата создания
	Deleted               *Timestamp                        `json:"deleted,omitempty"`               // Момент последнего удаления
	DeliveryPlannedMoment *Timestamp                        `json:"deliveryPlannedMoment,omitempty"` // Планируемая дата приемки
	Description           *string                           `json:"description,omitempty"`           // Комментарий
	ExternalCode          *string                           `json:"externalCode,omitempty"`          // Внешний код
	Files                 *Files                            `json:"files,omitempty"`                 // Метаданные массива Файлов
	Group                 *Group                            `json:"group,omitempty"`                 // Отдел сотрудника
	Id                    *uuid.UUID                        `json:"id,omitempty"`                    // ID сущности
	Meta                  *Meta                             `json:"meta,omitempty"`                  // Метаданные
	Moment                *Timestamp                        `json:"moment,omitempty"`                // Дата документа
	Moves                 Iterator[Move]                    `json:"moves,omitempty"`                 // Коллекция метаданных на связанные заказы перемещения
	Name                  *string                           `json:"name,omitempty"`                  // Наименование
	Organization          *Organization                     `json:"organization,omitempty"`          // Метаданные юрлица
	Owner                 *Employee                         `json:"owner,omitempty"`                 // Владелец (Сотрудник)
	Positions             *Positions[InternalOrderPosition] `json:"positions,omitempty"`             // Метаданные позиций Внутреннего заказа
	Printed               *bool                             `json:"printed,omitempty"`               // Напечатан ли документ
	Project               *Project                          `json:"project,omitempty"`               // Метаданные проекта
	Published             *bool                             `json:"published,omitempty"`             // Опубликован ли документ
	PurchaseOrders        *PurchaseOrders                   `json:"purchaseOrders,omitempty"`        // Коллекция метаданных на связанные заказы поставщику
	Rate                  *Rate                             `json:"rate,omitempty"`                  // Валюта
	Shared                *bool                             `json:"shared,omitempty"`                // Общий доступ
	State                 *State                            `json:"state,omitempty"`                 // Метаданные статуса
	Store                 *Store                            `json:"store,omitempty"`                 // Метаданные склада
	Sum                   *float64                          `json:"sum,omitempty"`                   // Сумма
	SyncId                *uuid.UUID                        `json:"syncId,omitempty"`                // ID синхронизации. После заполнения недоступен для изменения
	Updated               *Timestamp                        `json:"updated,omitempty"`               // Момент последнего обновления
	VatEnabled            *bool                             `json:"vatEnabled,omitempty"`            // Учитывается ли НДС
	VatIncluded           *bool                             `json:"vatIncluded,omitempty"`           // Включен ли НДС в цену
	VatSum                *float64                          `json:"vatSum,omitempty"`                // Сумма включая НДС
}

func (i InternalOrder) String() string {
	return Stringify(i)
}

func (i InternalOrder) MetaType() MetaType {
	return MetaTypeInternalOrder
}

// InternalOrderPosition Позиция Внутреннего заказа.
// Ключевое слово: internalorderposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vnutrennij-zakaz-vnutrennie-zakazy-pozicii-vnutrennego-zakaza
type InternalOrderPosition struct {
	AccountId  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	Id         *uuid.UUID          `json:"id,omitempty"`         // ID позиции
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Price      *float64            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Vat        *int                `json:"vat,omitempty"`        // НДС, которым облагается текущая позиция
	VatEnabled *bool               `json:"vatEnabled,omitempty"` // Включен ли НДС для позиции. С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
}

func (i InternalOrderPosition) String() string {
	return Stringify(i)
}

func (i InternalOrderPosition) MetaType() MetaType {
	return MetaTypeInternalOrderPosition
}
