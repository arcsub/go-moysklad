package moysklad

import (
	"github.com/google/uuid"
)

// CommissionReportIn Полученный отчет комиссионера.
// Ключевое слово: commissionreportin
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-poluchennyj-otchet-komissionera
type CommissionReportIn struct {
	AccountId                     *uuid.UUID                                   `json:"accountId,omitempty"`                     // ID учетной записи
	Agent                         *Counterparty                                `json:"agent,omitempty"`                         // Метаданные контрагента
	AgentAccount                  *AgentAccount                                `json:"agentAccount,omitempty"`                  // Метаданные счета контрагента
	Applicable                    *bool                                        `json:"applicable,omitempty"`                    // Отметка о проведении
	Attributes                    *Attributes                                  `json:"attributes,omitempty"`                    // Коллекция метаданных доп. полей. Поля объекта
	Code                          *string                                      `json:"code,omitempty"`                          // Код Полученного отчета комиссионера
	CommissionOverhead            *CommissionOverhead                          `json:"commissionOverhead,omitempty"`            // Прочие расходы
	CommissionPeriodEnd           *Timestamp                                   `json:"commissionPeriodEnd,omitempty"`           // Конец периода
	CommissionPeriodStart         *Timestamp                                   `json:"commissionPeriodStart,omitempty"`         // Начало периода
	CommitentSum                  *float64                                     `json:"commitentSum,omitempty"`                  // Сумма коммитента в установленной валюте
	Contract                      *Contract                                    `json:"contract,omitempty"`                      // Метаданные договора
	Created                       *Timestamp                                   `json:"created,omitempty"`                       // Дата создания
	Deleted                       *Timestamp                                   `json:"deleted,omitempty"`                       // Момент последнего удаления отчета комиссионера
	Description                   *string                                      `json:"description,omitempty"`                   // Комментарий отчета комиссионера
	ExternalCode                  *string                                      `json:"externalCode,omitempty"`                  // Внешний код отчета комиссионера
	Files                         *Files                                       `json:"files,omitempty"`                         // Метаданные массива Файлов (Максимальное количество файлов - 100)
	Group                         *Group                                       `json:"group,omitempty"`                         // Отдел сотрудника
	Id                            *uuid.UUID                                   `json:"id,omitempty"`                            // ID сущности
	Meta                          *Meta                                        `json:"meta,omitempty"`                          // Метаданные
	Moment                        *Timestamp                                   `json:"moment,omitempty"`                        // Дата документа
	Name                          *string                                      `json:"name,omitempty"`                          // Наименование
	Organization                  *Organization                                `json:"organization,omitempty"`                  // Метаданные юрлица
	OrganizationAccount           *AgentAccount                                `json:"organizationAccount,omitempty"`           // Метаданные счета юрлица
	Owner                         *Employee                                    `json:"owner,omitempty"`                         // Владелец (Сотрудник)
	PayedSum                      *float64                                     `json:"payedSum,omitempty"`                      // Оплаченная сумма
	Positions                     *Positions[CommissionReportInPosition]       `json:"positions,omitempty"`                     // Метаданные позиций отчета
	Printed                       *bool                                        `json:"printed,omitempty"`                       // Напечатан ли документ
	Project                       *Project                                     `json:"project,omitempty"`                       // Метаданные проекта
	Published                     *bool                                        `json:"published,omitempty"`                     // Опубликован ли документ
	Rate                          *Rate                                        `json:"rate,omitempty"`                          // Валюта
	ReturnToCommissionerPositions *Positions[CommissionReportInReturnPosition] `json:"returnToCommissionerPositions,omitempty"` // Метаданные позиций возврата на склад комиссионера Полученного отчета комиссионера
	RewardPercent                 *int                                         `json:"rewardPercent,omitempty"`                 // Процент вознаграждения (всегда 0 если вознаграждение не рассчитывается)
	RewardType                    RewardType                                   `json:"rewardType,omitempty"`                    // Тип вознаграждения
	SalesChannel                  *SalesChannel                                `json:"salesChannel,omitempty"`                  // Метаданные канала продаж
	Shared                        *bool                                        `json:"shared,omitempty"`                        // Общий доступ
	State                         *State                                       `json:"state,omitempty"`                         // Метаданные статуса отчета комиссионера
	Sum                           *float64                                     `json:"sum,omitempty"`                           // Сумма
	SyncId                        *uuid.UUID                                   `json:"syncId,omitempty"`                        // ID синхронизации. После заполнения недоступен для изменения
	Updated                       *Timestamp                                   `json:"updated,omitempty"`                       // Момент последнего обновления
	VatEnabled                    *bool                                        `json:"vatEnabled,omitempty"`                    // Учитывается ли НДС
	VatIncluded                   *bool                                        `json:"vatIncluded,omitempty"`                   // Включен ли НДС в цену
	VatSum                        *float64                                     `json:"vatSum,omitempty"`                        // Сумма включая НДС
	Payments                      *Payments                                    `json:"payments,omitempty"`                      // Массив ссылок на связанные платежи в формате Метаданных
}

func (c CommissionReportIn) String() string {
	return Stringify(c)
}

// GetMeta удовлетворяет интерфейсу HasMeta
func (c CommissionReportIn) GetMeta() *Meta {
	return c.Meta
}

func (c CommissionReportIn) MetaType() MetaType {
	return MetaTypeCommissionReportIn
}

// CommissionOverhead Прочие расходы
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-poluchennyj-otchet-komissionera-poluchennye-otchety-komissionera-prochie-rashody
type CommissionOverhead struct {
	Sum *float64 `json:"sum,omitempty"` // Сумма в копейках
}

func (c CommissionOverhead) String() string {
	return Stringify(c)
}

// CommissionReportInPosition Позиция Полученного отчета комиссионера.
// Ключевое слово: commissionreportinposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-poluchennyj-otchet-komissionera-poluchennye-otchety-komissionera-pozicii-poluchennogo-otcheta-komissionera
type CommissionReportInPosition struct {
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

func (c CommissionReportInPosition) String() string {
	return Stringify(c)
}

func (c CommissionReportInPosition) MetaType() MetaType {
	return MetaTypeCommissionReportInPosition
}

// CommissionReportInReturnPosition Позиция возврата на склад комиссионера.
// Ключевое слово: commissionreportinreturnedposition
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/documents/#dokumenty-poluchennyj-otchet-komissionera-poluchennye-otchety-komissionera-pozicii-poluchennogo-otcheta-komissionera-ob-ekt-pozicii-wozwrata-na-sklad-komissionera-soderzhit-sleduuschie-polq
type CommissionReportInReturnPosition struct {
	AccountId  *uuid.UUID          `json:"accountId,omitempty"`  // ID учетной записи
	Assortment *AssortmentPosition `json:"assortment,omitempty"` // Метаданные товара/услуги/серии/модификации, которую представляет собой позиция
	Id         *uuid.UUID          `json:"id,omitempty"`         // ID сущности
	Meta       *Meta               `json:"meta,omitempty"`       // Метаданные
	Price      *float64            `json:"price,omitempty"`      // Цена товара/услуги в копейках
	Quantity   *float64            `json:"quantity,omitempty"`   // Количество товаров/услуг данного вида в позиции. Если позиция - товар, у которого включен учет по серийным номерам, то значение в этом поле всегда будет равно количеству серийных номеров для данной позиции в документе.
	Reward     *float64            `json:"reward,omitempty"`     // Вознаграждение
	Vat        *int                `json:"vat,omitempty"`        // НДС, которым облагается текущая позиция
	VatEnabled *bool               `json:"vatEnabled,omitempty"` // Включен ли НДС для позиции. С помощью этого флага для позиции можно выставлять НДС = 0 или НДС = "без НДС". (vat = 0, vatEnabled = false) -> vat = "без НДС", (vat = 0, vatEnabled = true) -> vat = 0%.
}

func (c CommissionReportInReturnPosition) String() string {
	return Stringify(c)
}

func (c CommissionReportInReturnPosition) MetaType() MetaType {
	return MetaTypeCommissionReportInReturnPosition
}
