package moysklad

import (
	"github.com/google/uuid"
)

// CommissionReportOut Выданный отчет комиссионера.
// Ключевое слово: commissionreportout
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vydannyj-otchet-komissionera
type CommissionReportOut struct {
	AccountId             *uuid.UUID                              `json:"accountId,omitempty"`             // ID учетной записи
	Agent                 *Counterparty                           `json:"agent,omitempty"`                 // Метаданные контрагента
	AgentAccount          *AgentAccount                           `json:"agentAccount,omitempty"`          // Метаданные счета контрагента
	Applicable            *bool                                   `json:"applicable,omitempty"`            // Отметка о проведении
	Attributes            *Attributes                             `json:"attributes,omitempty"`            // Коллекция метаданных доп. полей. Поля объекта
	Code                  *string                                 `json:"code,omitempty"`                  // Код Полученного отчета комиссионера
	CommissionPeriodEnd   *Timestamp                              `json:"commissionPeriodEnd,omitempty"`   // Конец периода
	CommissionPeriodStart *Timestamp                              `json:"commissionPeriodStart,omitempty"` // Начало периода
	CommitentSum          *float64                                `json:"commitentSum,omitempty"`          // Сумма коммитента в установленной валюте
	Contract              *Contract                               `json:"contract,omitempty"`              // Метаданные договора
	Created               *Timestamp                              `json:"created,omitempty"`               // Дата создания
	Deleted               *Timestamp                              `json:"deleted,omitempty"`               // Момент последнего удаления отчета комиссионера
	Description           *string                                 `json:"description,omitempty"`           // Комментарий отчета комиссионера
	ExternalCode          *string                                 `json:"externalCode,omitempty"`          // Внешний код отчета комиссионера
	Files                 *Files                                  `json:"files,omitempty"`                 // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group                 *Group                                  `json:"group,omitempty"`                 // Отдел сотрудника
	Id                    *uuid.UUID                              `json:"id,omitempty"`                    // ID сущности
	Meta                  *Meta                                   `json:"meta,omitempty"`                  // Метаданные
	Moment                *Timestamp                              `json:"moment,omitempty"`                // Дата документа
	Name                  *string                                 `json:"name,omitempty"`                  // Наименование
	Organization          *Organization                           `json:"organization,omitempty"`          // Метаданные юрлица
	OrganizationAccount   *AgentAccount                           `json:"organizationAccount,omitempty"`   // Метаданные счета юрлица
	Owner                 *Employee                               `json:"owner,omitempty"`                 // Владелец (Сотрудник)
	PayedSum              *float64                                `json:"payedSum,omitempty"`              // Оплаченная сумма
	Positions             *Positions[CommissionReportOutPosition] `json:"positions,omitempty"`             // Метаданные позиций отчета
	Printed               *bool                                   `json:"printed,omitempty"`               // Напечатан ли документ
	Project               *Project                                `json:"project,omitempty"`               // Метаданные проекта
	Published             *bool                                   `json:"published,omitempty"`             // Опубликован ли документ
	Rate                  *Rate                                   `json:"rate,omitempty"`                  // Валюта
	RewardPercent         *int                                    `json:"rewardPercent,omitempty"`         // Процент вознаграждения (всегда 0 если вознаграждение не рассчитывается)
	RewardType            RewardType                              `json:"rewardType,omitempty"`            // Тип вознаграждения
	SalesChannel          *SalesChannel                           `json:"salesChannel,omitempty"`          // Метаданные канала продаж
	Shared                *bool                                   `json:"shared,omitempty"`                // Общий доступ
	State                 *State                                  `json:"state,omitempty"`                 // Метаданные статуса отчета комиссионера
	Sum                   *float64                                `json:"sum,omitempty"`                   // Сумма
	SyncId                *uuid.UUID                              `json:"syncId,omitempty"`                // ID синхронизации. После заполнения недоступен для изменения
	Updated               *Timestamp                              `json:"updated,omitempty"`               // Момент последнего обновления
	VatEnabled            *bool                                   `json:"vatEnabled,omitempty"`            // Учитывается ли НДС
	VatIncluded           *bool                                   `json:"vatIncluded,omitempty"`           // Включен ли НДС в цену
	VatSum                *float64                                `json:"vatSum,omitempty"`                // Сумма включая НДС
	Payments              *Payments                               `json:"payments,omitempty"`              // Массив ссылок на связанные платежи в формате Метаданных
}

func (c CommissionReportOut) String() string {
	return Stringify(c)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (c CommissionReportOut) GetMeta() *Meta {
	return c.Meta
}

func (c CommissionReportOut) MetaType() MetaType {
	return MetaTypeCommissionReportOut
}

// CommissionReportOutPosition Позиция Выданного отчета комиссионера.
// Ключевое слово: commissionreportoutposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-vydannyj-otchet-komissionera-vydannye-otchety-komissionera-pozicii-vydannogo-otcheta-komissionera
type CommissionReportOutPosition struct {
	AccountId  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	Id         *uuid.UUID          `json:"id,omitempty"`         // ID сущности
	Pack       *Pack               `json:"pack,omitempty"`       // Упаковка Товара
	Meta       *Meta               `json:"meta,omitempty"`       // Метаданные
	Price      *float64            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Reward     *float64            `json:"reward,omitempty"`     // Вознаграждение
	Vat        *int                `json:"vat,omitempty"`        // НДС, которым облагается текущая позиция
	VatEnabled *bool               `json:"vatEnabled,omitempty"` // Включен ли НДС для позиции. С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
}

func (c CommissionReportOutPosition) String() string {
	return Stringify(c)
}

func (c CommissionReportOutPosition) MetaType() MetaType {
	return MetaTypeCommissionReportOutPosition
}
